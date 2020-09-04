package launchpad

import (
	"context"
	"testing"

	"github.com/coinbase/rosetta-sdk-go/types"
	types2 "github.com/cosmos/cosmos-sdk/types"
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
			expectedErr: ErrInvalidOperation,
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
			expectedErr: rosetta.WrapError(ErrInvalidOperation, "the operations are not Transfer"),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			adapter := NewLaunchpad(TendermintAPI{}, CosmosAPI{}, rosetta.NetworkProperties{})
			_, err := adapter.ConstructionPayloads(context.Background(), tt.req)
			require.Equal(t, err, tt.expectedErr)
		})
	}
}

func TestGetSenderByOperations(t *testing.T) {
	ops := []*types.Operation{
		{
			Account: &types.AccountIdentifier{
				Address: "cosmos15tltvs59rt88geyenetv3klavlq2z30fe8z6hj",
			},
			Type: OperationTransfer,
			Amount: &types.Amount{
				Value: "12345",
				Currency: &types.Currency{
					Symbol:   "stake",
					Decimals: 0,
				},
				Metadata: nil,
			},
		},
		{
			Account: &types.AccountIdentifier{
				Address: "cosmos16xyempempp92x9hyzz9wrgf94r6j9h5f06pxxv",
			},
			Type: OperationTransfer,
			Amount: &types.Amount{
				Value: "-12345",
				Currency: &types.Currency{
					Symbol:   "stake",
					Decimals: 0,
				},
				Metadata: nil,
			},
		},
	}

	transferData, err := getFromAndToAddressFromOperations(ops)
	require.NoError(t, err)

	expectedFrom, err := types2.AccAddressFromBech32(ops[1].Account.Address)
	expectedTo, err := types2.AccAddressFromBech32(ops[0].Account.Address)
	require.NoError(t, err)

	require.Equal(t, expectedFrom, transferData.From)
	require.Equal(t, expectedTo, transferData.To)
	require.Equal(t, types2.NewCoin("stake", types2.NewInt(12345)), transferData.Amount)
}
