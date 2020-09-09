package launchpad

import (
	"context"
	"encoding/base64"
	"io/ioutil"
	"testing"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

func TestLaunchpad_ConstructionHash(t *testing.T) {
	expectedHash := "6f22ea7620ebcb5078d244f06e88dd26906ba1685135bfc34f83fefdd653198a"

	bz, err := ioutil.ReadFile("./testdata/signed-tx.json")
	require.NoError(t, err)

	var stdTx auth.StdTx
	cdc := simapp.MakeCodec()
	err = cdc.UnmarshalJSON(bz, &stdTx)
	require.NoError(t, err)

	adapter := NewLaunchpad(TendermintAPI{}, CosmosAPI{}, rosetta.NetworkProperties{})

	// re-encode it via the Amino wire protocol
	txBytes, err := cdc.MarshalBinaryLengthPrefixed(stdTx)
	require.NoError(t, err)

	// base64 encode the encoded tx bytes
	txBytesBase64 := base64.StdEncoding.EncodeToString(txBytes)

	resp, err2 := adapter.ConstructionHash(context.Background(), &types.ConstructionHashRequest{
		SignedTransaction: txBytesBase64,
	})

	require.Nil(t, err2)
	require.Equal(t, expectedHash, resp.TransactionIdentifier.Hash)
}
