package launchpad

import (
	"context"
	"strconv"
	"time"

	"github.com/antihax/optional"
	"github.com/coinbase/rosetta-sdk-go/types"
	"golang.org/x/sync/errgroup"

	tendermintclient "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint/generated"
)

func (l Launchpad) NetworkList(context.Context, *types.MetadataRequest) (*types.NetworkListResponse, *types.Error) {
	return &types.NetworkListResponse{
		NetworkIdentifiers: []*types.NetworkIdentifier{
			{
				Blockchain: l.properties.Blockchain,
				Network:    l.properties.Network,
			},
		},
	}, nil
}

func (l Launchpad) NetworkOptions(ctx context.Context, _ *types.NetworkRequest) (*types.NetworkOptionsResponse, *types.Error) {
	resp, _, err := l.cosmos.Tendermint.NodeInfoGet(ctx)
	if err != nil {
		return nil, ErrNodeConnection
	}

	return &types.NetworkOptionsResponse{
		Version: &types.Version{
			RosettaVersion: "1.2.5",
			NodeVersion:    resp.NodeInfo.Version,
		},
		Allow: &types.Allow{
			OperationStatuses: []*types.OperationStatus{
				{
					Status:     StatusSuccess,
					Successful: true,
				},
				{
					Status:     StatusReverted,
					Successful: false,
				},
			},
			OperationTypes: []string{OperationTransfer},
		},
	}, nil
}

func (l Launchpad) NetworkStatus(ctx context.Context, _ *types.NetworkRequest) (*types.NetworkStatusResponse, *types.Error) {
	var (
		latestBlock  tendermintclient.BlockResponse
		genesisBlock tendermintclient.BlockResponse
		netInfo      tendermintclient.NetInfoResponse
	)

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() (err error) {
		latestBlock, _, err = l.tendermint.Info.Block(ctx, nil)
		return
	})
	g.Go(func() (err error) {
		genesisBlock, _, err = l.tendermint.Info.Block(ctx, &tendermintclient.BlockOpts{
			Height: optional.NewFloat32(1),
		})
		return
	})
	g.Go(func() (err error) {
		netInfo, _, err = l.tendermint.Info.NetInfo(ctx)
		return
	})
	if err := g.Wait(); err != nil {
		return nil, ErrNodeConnection
	}

	var peers []*types.Peer
	for _, p := range netInfo.Result.Peers {
		peers = append(peers, &types.Peer{
			PeerID: p.NodeInfo.Id,
		})
	}

	height, err := strconv.ParseUint(latestBlock.Result.Block.Header.Height, 10, 64)
	if err != nil {
		return nil, ErrInterpreting
	}

	t, err := time.Parse(time.RFC3339Nano, latestBlock.Result.Block.Header.Time)
	if err != nil {
		return nil, ErrInterpreting
	}

	return &types.NetworkStatusResponse{
		CurrentBlockIdentifier: &types.BlockIdentifier{
			Index: int64(height),
			Hash:  latestBlock.Result.BlockId.Hash,
		},
		CurrentBlockTimestamp: t.UnixNano() / 1000000,
		GenesisBlockIdentifier: &types.BlockIdentifier{
			Index: 1,
			Hash:  genesisBlock.Result.BlockId.Hash,
		},
		Peers: peers,
	}, nil
}
