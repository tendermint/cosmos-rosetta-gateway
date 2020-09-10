package launchpad

import (
	"context"
	"encoding/base64"
	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
	"io/ioutil"
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
	bz, err := ioutil.ReadFile("./testdata/signed-tx.json")
	require.NoError(t, err)

	var stdTx auth.StdTx
	cdc := simapp.MakeCodec()
	err = cdc.UnmarshalJSON(bz, &stdTx)
	require.NoError(t, err)

	// re-encode it via the Amino wire protocol
	txBytes, err := cdc.MarshalBinaryLengthPrefixed(stdTx)
	require.NoError(t, err)

	// base64 encode the encoded tx bytes
	txBytesBase64 := base64.StdEncoding.EncodeToString(txBytes)

	adapter := NewLaunchpad(TendermintAPI{}, CosmosAPI{}, properties)
	var combineRes, combineErr = adapter.ConstructionCombine(context.Background(), &types.ConstructionCombineRequest{
		UnsignedTransaction: txBytesBase64,
		Signatures: []*types.Signature{{
			SigningPayload: &types.SigningPayload{
				Address: "cosmos1qrv8g4hwt4z6ds8mednhhgx907wug9d6y8n9jy",
				Bytes:   []byte("KlPOI6frSRdjVHgiBIIpHI2PAQQnCCMSTWgonJQbECRATe8Yf7gqRgAVeLcPiVUbp2oSq2P7pp51f0iCtiA47Q==")},
		},
		}})
	require.Nil(t, combineErr)
	require.NotNil(t, combineRes)
}
