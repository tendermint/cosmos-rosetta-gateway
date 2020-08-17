package launchpad

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/coinbase/rosetta-sdk-go/types"
	"golang.org/x/sync/errgroup"
)

const (
	// tendermint.
	endpointNetInfo = "/net_info"
	endpointBlock   = "/block"
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

type nodeResponse struct {
	NodeInfo nodeInfo `json:"node_info"`
}

func (l Launchpad) NetworkOptions(ctx context.Context, _ *types.NetworkRequest) (*types.NetworkOptionsResponse, *types.Error) {
	resp, _, err := l.api.Tendermint.NodeInfoGet(ctx)
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
					Status:     "SUCCESS",
					Successful: true,
				},
			},
			OperationTypes: l.properties.SupportedOperations,
		},
	}, nil
}

type blockResponse struct {
	Result result `json:"result"`
}

type result struct {
	BlockID blockID `json:"block_id"`
	Block   block   `json:"block"`
}

type netInfoResponse struct {
	Result netInfoResult `json:"result"`
}

type netInfoResult struct {
	Peers []peer `json:"peers"`
}

func (l Launchpad) NetworkStatus(ctx context.Context, _ *types.NetworkRequest) (*types.NetworkStatusResponse, *types.Error) {
	var (
		latestBlockResp  blockResponse
		genesisBlockResp blockResponse
		netInfoResp      netInfoResponse
	)

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, l.tendermint(endpointBlock), nil)
		if err != nil {
			return err
		}

		resp, err := l.c.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		return json.NewDecoder(resp.Body).Decode(&latestBlockResp)
	})
	g.Go(func() error {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, l.tendermint(endpointNetInfo), nil)
		if err != nil {
			return err
		}
		resp, err := l.c.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		return json.NewDecoder(resp.Body).Decode(&netInfoResp)
	})
	g.Go(func() error {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, l.tendermint(endpointBlock), nil)
		if err != nil {
			return err
		}
		q := req.URL.Query()
		q.Add("height", "1")
		req.URL.RawQuery = q.Encode()

		resp, err := l.c.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		return json.NewDecoder(resp.Body).Decode(&genesisBlockResp)
	})
	if err := g.Wait(); err != nil {
		fmt.Println(err)
		return nil, ErrNodeConnection
	}

	var peers []*types.Peer
	for _, p := range netInfoResp.Result.Peers {
		peers = append(peers, &types.Peer{
			PeerID: p.NodeInfo.ID,
		})
	}

	return &types.NetworkStatusResponse{
		CurrentBlockIdentifier: &types.BlockIdentifier{
			Index: latestBlockResp.Result.Block.Header.Height.Int64(),
			Hash:  latestBlockResp.Result.BlockID.Hash,
		},
		CurrentBlockTimestamp: latestBlockResp.Result.Block.Header.Time.UnixNano() / 1000000,
		GenesisBlockIdentifier: &types.BlockIdentifier{
			Index: 1,
			Hash:  genesisBlockResp.Result.BlockID.Hash,
		},
		Peers: peers,
	}, nil
}
