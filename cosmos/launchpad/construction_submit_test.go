package launchpad

import (
	"encoding/hex"
	"io/ioutil"
	"testing"

	"github.com/cosmos/cosmos-sdk/x/auth/client/rest"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/stretchr/testify/require"
)

func TestLaunchpad_ConstructionSubmit(t *testing.T) {
	t.SkipNow() // TODO bring back.
	bz, err := ioutil.ReadFile("./testdata/test-with-signature-delete.json")
	require.NoError(t, err)

	cdc := simapp.MakeCodec()

	var stdTx rest.BroadcastReq
	err = cdc.UnmarshalJSON(bz, &stdTx)
	require.NoError(t, err)

	// re-encode it via the Amino wire protocol
	txBytes, err := cdc.MarshalJSON(stdTx)
	require.NoError(t, err)
	t.Logf("%s\n", txBytes)

	toString := hex.EncodeToString(txBytes)
	t.Logf("%s\n", toString)

	//testTx := cosmosclient.InlineObject{
	//	Tx: cosmosclient.StdTx{
	//		//Value: stdTx,
	//	},
	//	Mode: "async",
	//}
	//
	//m := mocks.CosmosTransactionsAPI{}
	//m.
	//	On("TxsPost", mock.Anything, testTx).
	//	Return(cosmosclient.BroadcastTxCommitResult{
	//		Hash:   expectedHash,
	//		Height: 10,
	//	}, nil, nil).Once()
	//
	//adapter := NewLaunchpad(TendermintAPI{}, CosmosAPI{Transactions: &m}, rosetta.NetworkProperties{})
	//
	//// re-encode it via the Amino wire protocol
	//txBytes, err := cdc.MarshalBinaryLengthPrefixed(stdTx)
	//require.NoError(t, err)
	//
	//toString := hex.EncodeToString(txBytes)
	//
	//resp, err2 := adapter.ConstructionSubmit(context.Background(), &types.ConstructionSubmitRequest{
	//	SignedTransaction: toString,
	//})
	//fmt.Printf("%v\n", resp)
	//
	//require.Nil(t, err2)
	//require.NotNil(t, resp)
	//require.Equal(t, expectedHash, resp.TransactionIdentifier.Hash)
}
