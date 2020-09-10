package launchpad

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

func TestLaunchpad_ConstructionSubmit(t *testing.T) {
	bz, err := ioutil.ReadFile("./testdata/test-with-signature-delete.json")
	require.NoError(t, err)

	var stdTx auth.StdTx
	cdc := simapp.MakeCodec()
	err = cdc.UnmarshalJSON(bz, &stdTx)
	require.NoError(t, err)

	adapter := NewLaunchpad(TendermintAPI{}, CosmosAPI{}, rosetta.NetworkProperties{})

	// re-encode it via the Amino wire protocol
	txBytes, err := cdc.MarshalBinaryLengthPrefixed(stdTx)
	require.NoError(t, err)

	toString := hex.EncodeToString(txBytes)
	fmt.Printf("\n%s\n", toString)

	// base64 encode the encoded tx bytes
	txBytesBase64 := base64.StdEncoding.EncodeToString(txBytes)

	resp, err2 := adapter.ConstructionSubmit(context.Background(), &types.ConstructionSubmitRequest{
		SignedTransaction: txBytesBase64,
	})
	fmt.Printf("%v\n", resp)

	require.Nil(t, err2)
}
