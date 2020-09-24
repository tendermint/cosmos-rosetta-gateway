package launchpad

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk"
	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

func TestLaunchpad_ConstructionHash(t *testing.T) {
	t.SkipNow() // TODO bring back.
	expectedHash := "6f22ea7620ebcb5078d244f06e88dd26906ba1685135bfc34f83fefdd653198a"

	bz, err := ioutil.ReadFile("./testdata/test-with-signature-delete.json")
	require.NoError(t, err)

	var stdTx auth.StdTx
	cdc := simapp.MakeCodec()
	err = cdc.UnmarshalJSON(bz, &stdTx)
	require.NoError(t, err)

	adapter := NewLaunchpad(sdk.NewClient(""), tendermint.NewClient(""), rosetta.NetworkProperties{})

	// re-encode it via the Amino wire protocol
	txBytes, err := cdc.MarshalBinaryLengthPrefixed(stdTx)
	require.NoError(t, err)

	// base64 encode the encoded tx bytes
	txBytesBase64 := base64.StdEncoding.EncodeToString(txBytes)
	fmt.Printf("\n%s\n", txBytesBase64)

	resp, err2 := adapter.ConstructionHash(context.Background(), &types.ConstructionHashRequest{
		SignedTransaction: txBytesBase64,
	})

	require.Nil(t, err2)
	require.Equal(t, expectedHash, resp.TransactionIdentifier.Hash)
}
