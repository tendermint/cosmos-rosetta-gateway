package launchpad

import (
	"context"
	"strconv"
	"time"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/alttendermint"

	"github.com/coinbase/rosetta-sdk-go/types"
	"golang.org/x/sync/errgroup"
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
		latestBlock  alttendermint.BlockResponse
		genesisBlock alttendermint.BlockResponse
		netInfo      alttendermint.NetInfoResponse
	)

	g, _ := errgroup.WithContext(ctx)
	g.Go(func() (err error) {
		latestBlock, err = l.altTendermint.Block(0)
		return
	})
	g.Go(func() (err error) {
		genesisBlock, err = l.altTendermint.Block(1)
		return
	})
	g.Go(func() (err error) {
		netInfo, err = l.altTendermint.NetInfo()
		return
	})
	if err := g.Wait(); err != nil {
		return nil, ErrNodeConnection
	}

	var peers []*types.Peer
	for _, p := range netInfo.Peers {
		peers = append(peers, &types.Peer{
			PeerID: p.NodeInfo.Id,
		})
	}

	height, err := strconv.ParseUint(latestBlock.Block.Header.Height, 10, 64)
	if err != nil {
		return nil, ErrInterpreting
	}

	t, err := time.Parse(time.RFC3339Nano, latestBlock.Block.Header.Time)
	if err != nil {
		return nil, ErrInterpreting
	}

	return &types.NetworkStatusResponse{
		CurrentBlockIdentifier: &types.BlockIdentifier{
			Index: int64(height),
			Hash:  latestBlock.BlockId.Hash,
		},
		CurrentBlockTimestamp: t.UnixNano() / 1000000,
		GenesisBlockIdentifier: &types.BlockIdentifier{
			Index: 1,
			Hash:  genesisBlock.BlockId.Hash,
		},
		Peers: peers,
	}, nil
}
