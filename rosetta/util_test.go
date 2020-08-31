package rosetta_test

import (
	"testing"

	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"

	"github.com/stretchr/testify/require"

	"github.com/coinbase/rosetta-sdk-go/types"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad"
)

func TestGetOperationsByRelation(t *testing.T) {
	tests := []struct {
		name       string
		operations []*types.Operation
	}{
		{
			"only 1 operation with no relations",
			[]*types.Operation{
				{
					OperationIdentifier: &types.OperationIdentifier{
						Index:        0,
						NetworkIndex: nil,
					},
					RelatedOperations: nil,
					Type:              launchpad.OperationTransfer,
					Status:            launchpad.StatusSuccess,
					Account: &types.AccountIdentifier{
						Address: "",
					},
					Amount: &types.Amount{
						Value: "1000",
						Currency: &types.Currency{
							Symbol:   "stake",
							Decimals: 0,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			opRelations := rosetta.GetOperationsByRelation(tt.operations)
			require.Equal(t, opRelations[0], []*types.Operation{
				tt.operations[0],
			})
		})
	}
}
