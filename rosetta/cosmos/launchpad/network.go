package launchpad

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/coinbase/rosetta-sdk-go/types"
)

const (
	endpointNodeInfo = "/node_info"
)

func (l Launchpad) NetworkList(ctx context.Context, request *types.MetadataRequest) (*types.NetworkListResponse, *types.Error) {
	return &types.NetworkListResponse{
		NetworkIdentifiers: []*types.NetworkIdentifier{
			{
				Blockchain: l.blockchain,
				Network:    l.network,
			},
		},
	}, nil
}

type nodeInfo struct {
	Version string `json:"version"`
}

type nodeResponse struct {
	NodeInfo nodeInfo `json:"node_info"`
}

func (l Launchpad) NetworkOptions(ctx context.Context, request *types.NetworkRequest) (*types.NetworkOptionsResponse, *types.Error) {
	addr := fmt.Sprintf("%s%s", l.endpoint, endpointNodeInfo)
	resp, err := l.c.Get(addr)
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
			OperationTypes: l.operations,
		},
	}, nil
}

func (l Launchpad) NetworkStatus(ctx context.Context, request *types.NetworkRequest) (*types.NetworkStatusResponse, *types.Error) {
	panic("implement me")
}
