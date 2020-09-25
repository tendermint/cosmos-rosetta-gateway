package launchpad

import (
	"context"
	"encoding/hex"
	"fmt"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/mocks"
	"io/ioutil"
	"testing"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint"
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk"

	"github.com/cosmos/cosmos-sdk/x/auth/client/rest"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/stretchr/testify/require"
)

func TestLaunchpad_ConstructionSubmit(t *testing.T) {
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

	expectedHash := "6f22ea7620ebcb5078d244f06e88dd26906ba1685135bfc34f83fefdd653198a"
	m := &mocks.SdkClient{}
	m.
		On("PostTx", context.Background(), bz).
		Return(cosmostypes.TxResponse{
			TxHash: expectedHash,
			Height: 10,
		}, nil, nil).Once()
	properties := rosetta.NetworkProperties{
		Blockchain: "TheBlockchain",
		Network:    "TheNetwork",
		AddrPrefix: "test",
		SupportedOperations: []string{
			"Transfer",
		},
	}

	adapter := NewLaunchpad(m, tendermint.NewClient(""), properties)
	resp, err2 := adapter.ConstructionSubmit(context.Background(), &types.ConstructionSubmitRequest{
		SignedTransaction: toString,
	})
	fmt.Printf("%v\n", resp)

	require.Nil(t, err2)
	require.NotNil(t, resp)
	require.Equal(t, expectedHash, resp.TransactionIdentifier.Hash)
}

func TestLaunchpad_ConstructionSubmit_FailsOfflineMode(t *testing.T) {
	properties := rosetta.NetworkProperties{
		OfflineMode: true,
	}
	adapter := NewLaunchpad(sdk.NewClient(""), tendermint.NewClient(""), properties)

	_, err2 := adapter.ConstructionSubmit(context.Background(), &types.ConstructionSubmitRequest{
		SignedTransaction: "dkajfkdjkads",
	})

	require.Equal(t, ErrEndpointDisabledOfflineMode, err2)
}
