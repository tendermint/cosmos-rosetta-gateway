package launchpad

import (
	"context"
	"testing"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

func TestLaunchpad_ConstructionMetadata(t *testing.T) {
	properties := rosetta.NetworkProperties{
		Blockchain: "TheBlockchain",
		Network:    "TheNetwork",
		SupportedOperations: []string{
			"Transfer",
		},
	}

	networkIdentifier := types.NetworkIdentifier{
		Blockchain: "TheBlockchain",
		Network:    "TheNetwork",
	}
	feeMultiplier := float64(200000)
	options := map[string]interface{}{
		OptionAddress: "cosmos12qqzw4tqu32anlcx0a3hupvgdhaf4cc8j9wfyd",
		OptionGas:     &feeMultiplier,
	}

	expMetadata := map[string]interface{}{
		OptionsAccountNumber: 0,
		OptionsSequence:      0,
		OptionsChainId:       "TheNetwork",
	}
	adapter := NewLaunchpad(TendermintAPI{}, CosmosAPI{}, properties)
	metaResp, err := adapter.ConstructionMetadata(context.Background(), &types.ConstructionMetadataRequest{
		NetworkIdentifier: &networkIdentifier,
		Options:           options,
	})

	require.Nil(t, err)
	require.NotNil(t, metaResp)
	if diff := cmp.Diff(metaResp.Metadata, expMetadata); diff != "" {
		t.Errorf("Metadata mismatch %s", diff)
	}
}
