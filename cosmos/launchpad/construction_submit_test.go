package launchpad

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/stretchr/testify/mock"
	cosmosclient "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/generated"
	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/mocks"
	"io/ioutil"
	"testing"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

func TestLaunchpad_ConstructionSubmit(t *testing.T) {
	expectedHash := "6f22ea7620ebcb5078d244f06e88dd26906ba1685135bfc34f83fefdd653198a"

	bz, err := ioutil.ReadFile("./testdata/test-with-signature-delete.json")
	require.NoError(t, err)

	var stdTx cosmosclient.StdTxValue
	cdc := simapp.MakeCodec()
	err = cdc.UnmarshalJSON(bz, &stdTx)
	require.NoError(t, err)

	testTx := cosmosclient.InlineObject{
		Tx: cosmosclient.StdTx{
			Value: stdTx,
		},
		Mode: "async",
	}

	m := mocks.CosmosTransactionsAPI{}
	m.
		On("TxsPost", mock.Anything, testTx).
		Return(cosmosclient.BroadcastTxCommitResult{
			Hash:   expectedHash,
			Height: 10,
		}, nil, nil).Once()

	adapter := NewLaunchpad(TendermintAPI{}, CosmosAPI{Transactions: &m}, rosetta.NetworkProperties{})

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
	require.NotNil(t, resp)
	require.Equal(t, expectedHash, resp.TransactionIdentifier.Hash)
}
