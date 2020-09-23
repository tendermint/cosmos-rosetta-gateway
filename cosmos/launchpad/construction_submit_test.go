package launchpad

import (
	"encoding/hex"
	"github.com/coinbase/rosetta-sdk-go/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/mock"
	clientsdk "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk"
	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/mocks"
	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint"
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
	"io/ioutil"
	"testing"

	"github.com/cosmos/cosmos-sdk/x/auth/client/rest"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/stretchr/testify/require"
)

func TestLaunchpad_ConstructionSubmit(t *testing.T) {
	bz, err := ioutil.ReadFile("./signed-tx.json")
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

	testTx := BroadcastReq{
		Tx:   stdTx,
		Mode: "async",
	}

	m := mocks.SdkClient{}
	m.
		On("PostTx", mock.Anything, testTx).
		Return(sdk.TxResponse{
			TxHash: expectedHash,
			Height: 10,
		}, nil, nil).Once()

	adapter := NewLaunchpad(clientsdk.NewClient(""), tendermint.NewClient(""), rosetta.NetworkProperties{})

	// re-encode it via the Amino wire protocol
	txBytes, err := cdc.MarshalBinaryLengthPrefixed(stdTx)
	require.NoError(t, err)

	toString := hex.EncodeToString(txBytes)

	resp, err2 := adapter.ConstructionSubmit(context.Background(), &types.ConstructionSubmitRequest{
		SignedTransaction: toString,
	})
	fmt.Printf("%v\n", resp)

	require.Nil(t, err2)
	require.NotNil(t, resp)
	require.Equal(t, expectedHash, resp.TransactionIdentifier.Hash)
}
