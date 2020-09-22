package launchpad

import (
	"context"
	"testing"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/alttendermint"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

func TestLaunchpad_ConstructionPreprocess(t *testing.T) {
	properties := rosetta.NetworkProperties{
		Blockchain: "TheBlockchain",
		Network:    "TheNetwork",
		SupportedOperations: []string{
			"Transfer",
			"Reward",
		},
	}

	ops := []*types.Operation{
		{
			OperationIdentifier: &types.OperationIdentifier{},
			Type:                "Transfer",
			Status:              "Success",
			Account: &types.AccountIdentifier{
				Address: "cosmos12qqzw4tqu32anlcx0a3hupvgdhaf4cc8j9wfyd",
			},
			Amount: &types.Amount{
				Value: "-10",
				Currency: &types.Currency{
					Symbol: "token",
				},
			},
		},
		{
			OperationIdentifier: &types.OperationIdentifier{
				Index: 1,
			},
			Type:   "Transfer",
			Status: "Success",
			Account: &types.AccountIdentifier{
				Address: "cosmos10rpmm9ur87le39hehteha37sg5awdsns6huyvy",
			},
			Amount: &types.Amount{
				Value: "10",
				Currency: &types.Currency{
					Symbol: "token",
				},
			},
		},
	}
	feeMultiplier := float64(200000)

	expOptions := map[string]interface{}{
		OptionAddress: "cosmos12qqzw4tqu32anlcx0a3hupvgdhaf4cc8j9wfyd",
		OptionGas:     &feeMultiplier,
	}

	adapter := NewLaunchpad(TendermintAPI{}, sdk.NewClient(""), alttendermint.NewClient(""), properties)
	deriveResp, deriveErr := adapter.ConstructionPreprocess(context.Background(), &types.ConstructionPreprocessRequest{
		Operations:             ops,
		SuggestedFeeMultiplier: &feeMultiplier,
	})

	require.Nil(t, deriveErr)
	require.NotNil(t, deriveResp)
	if diff := cmp.Diff(deriveResp.Options, expOptions); diff != "" {
		t.Errorf("Options mismatch %s", diff)
	}
}
