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
	// cosmos.
	endpointNodeInfo    = "/node_info"
	endpointBlockLatest = "/blocks/latest"

	// tendermint.
	endpointNetInfo = "/net_info"
	endpointBlock   = "/block"
)

func (l Launchpad) NetworkList(ctx context.Context, request *types.MetadataRequest) (*types.NetworkListResponse, *types.Error) {
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

func (l Launchpad) NetworkOptions(ctx context.Context, request *types.NetworkRequest) (*types.NetworkOptionsResponse, *types.Error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, l.cosmos(endpointNodeInfo), nil)
	if err != nil {
		return nil, ErrNodeConnection
	}
	resp, err := l.c.Do(req)
	if err != nil {
		return nil, ErrNodeConnection
	}
	defer resp.Body.Close()

	var nodeResp nodeResponse
	if err := json.NewDecoder(resp.Body).Decode(&nodeResp); err != nil {
		return nil, ErrInterpreting
	}

	return &types.NetworkOptionsResponse{
		Version: &types.Version{
			RosettaVersion: "1.2.5",
			NodeVersion:    nodeResp.NodeInfo.Version,
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

type latestBlockResponse struct {
	Block block `json:"block"`
}

type genesisResponse struct {
	Result genesisResult `json:"result"`
}

type genesisResult struct {
	BlockID blockID `json:"block_id"`
	Block   block   `json:"block"`
}

type netInfoResponse struct {
	Result netInfoResult `json:"result"`
}

type netInfoResult struct {
	Peers []peer `json:"peers"`
}

func (l Launchpad) NetworkStatus(ctx context.Context, request *types.NetworkRequest) (*types.NetworkStatusResponse, *types.Error) {
	var (
		latestBlockResp   latestBlockResponse
		genesistBlockResp genesisResponse
		netInfoResp       netInfoResponse
	)

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, l.cosmos(endpointBlockLatest), nil)
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
		return json.NewDecoder(resp.Body).Decode(&genesistBlockResp)
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
			Index: latestBlockResp.Block.Header.Height.Int64(),
			Hash:  latestBlockResp.Block.Header.LastBlockID.Hash,
		},
		CurrentBlockTimestamp: latestBlockResp.Block.Header.Time.UnixNano() / 1000000,
		GenesisBlockIdentifier: &types.BlockIdentifier{
			Hash: genesistBlockResp.Result.BlockID.Hash,
		},
		Peers: peers,
	}, nil
}
