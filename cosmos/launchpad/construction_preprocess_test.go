package launchpad

import (
	"context"
	"testing"

	"github.com/coinbase/rosetta-sdk-go/types"
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

	adapter := NewLaunchpad(TendermintAPI{}, CosmosAPI{}, properties)
	deriveResp, deriveErr := adapter.ConstructionPreprocess(context.Background(), &types.ConstructionPreprocessRequest{})
	require.Nil(t, deriveErr)
	require.NotNil(t, deriveResp)
}
