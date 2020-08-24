package launchpad

import (
	"context"
	"encoding/hex"
	"testing"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

func TestLaunchpad_ConstructionDerive(t *testing.T) {
	data, err := hex.DecodeString("A2FEB642851ACE7464999E56C8DBFD67C0A145E9")
	require.NoError(t, err)

	properties := rosetta.NetworkProperties{
		Blockchain: "TheBlockchain",
		Network:    "TheNetwork",
		SupportedOperations: []string{
			"Transfer",
			"Reward",
		},
	}

	adapter := NewLaunchpad(TendermintAPI{}, CosmosAPI{}, properties)
	deriveResp, deriveErr := adapter.ConstructionDerive(context.Background(), &types.ConstructionDeriveRequest{
		PublicKey: &types.PublicKey{
			Bytes: data,
		},
	})
	require.Nil(t, deriveErr)
	require.NotNil(t, deriveResp)
	require.Equal(t, "cosmos15tltvs59rt88geyenetv3klavlq2z30fe8z6hj", deriveResp.Address)
}
