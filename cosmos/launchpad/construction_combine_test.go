package launchpad

import (
	"context"
	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
	"testing"
)

func TestLaunchpad_ConstructionCombine(t *testing.T) {
	properties := rosetta.NetworkProperties{
		Blockchain: "TheBlockchain",
		Network:    "TheNetwork",
		SupportedOperations: []string{
			"Transfer",
			"Reward",
		},
	}

	adapter := NewLaunchpad(TendermintAPI{}, CosmosAPI{}, properties)
	var combineRes, err = adapter.ConstructionCombine(context.Background(), &types.ConstructionCombineRequest{
		UnsignedTransaction: "test",
		Signatures: []*types.Signature{{
			SigningPayload: &types.SigningPayload{
				Address: "cosmos1qrv8g4hwt4z6ds8mednhhgx907wug9d6y8n9jy",
				Bytes:   []byte("KlPOI6frSRdjVHgiBIIpHI2PAQQnCCMSTWgonJQbECRATe8Yf7gqRgAVeLcPiVUbp2oSq2P7pp51f0iCtiA47Q==")},
		},
		}})
	require.Nil(t, err)
	require.NotNil(t, combineRes)
}
