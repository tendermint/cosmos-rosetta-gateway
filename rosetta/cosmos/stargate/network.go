package stargate

import (
	"context"

	"github.com/coinbase/rosetta-sdk-go/types"
)

func (s Stargate) NetworkList(ctx context.Context, request *types.MetadataRequest) (*types.NetworkListResponse, *types.Error) {
	panic("implement me")
}

func (s Stargate) NetworkOptions(ctx context.Context, request *types.NetworkRequest) (*types.NetworkOptionsResponse, *types.Error) {
	panic("implement me")
}

func (s Stargate) NetworkStatus(ctx context.Context, request *types.NetworkRequest) (*types.NetworkStatusResponse, *types.Error) {
	panic("implement me")
}
