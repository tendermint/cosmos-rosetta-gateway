package launchpad

import (
	"context"
	"testing"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

func TestPayloadsEndpoint_Errors(t *testing.T) {
	tests := []struct {
		name        string
		req         *types.ConstructionPayloadsRequest
		expectedErr *types.Error
	}{
		{
			name: "Invalid num of operations",
			req: &types.ConstructionPayloadsRequest{
				Operations: []*types.Operation{
					{
						Type: OperationTransfer,
					},
				},
			},
		},
		{
			name: "Two operations not equal to transfer",
			req: &types.ConstructionPayloadsRequest{
				Operations: []*types.Operation{
					{
						Type: OperationTransfer,
					},
					{
						Type: "otherType",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			adapter := NewLaunchpad(TendermintAPI{}, CosmosAPI{}, rosetta.NetworkProperties{})
			_, err := adapter.ConstructionPayloads(context.Background(), tt.req)
			require.Equal(t, err, ErrInvalidOperation)
		})
	}
}
