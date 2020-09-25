package launchpad

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"testing"

	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/mocks"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk"
	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint"

	"github.com/stretchr/testify/require"
)

func TestLaunchpad_ConstructionSubmit(t *testing.T) {
	bz, err := ioutil.ReadFile("./testdata/signed-tx.json")
	require.NoError(t, err)
	signedTx := hex.EncodeToString(bz)

	var tx map[string]json.RawMessage
	err = json.Unmarshal(bz, &tx)
	require.NoError(t, err)

	bReq := BroadcastReq{
		Tx:   tx["value"],
		Mode: "block",
	}

	byteRequest, err := json.Marshal(bReq)
	require.NoError(t, err)

	expectedHash := "6f22ea7620ebcb5078d244f06e88dd26906ba1685135bfc34f83fefdd653198a"
	m := &mocks.SdkClient{}
	m.
		On("PostTx", context.Background(), byteRequest).
		Return(cosmostypes.TxResponse{
			TxHash: expectedHash,
			Height: 10,
		}, nil, nil).Once()
	properties := properties{
		Blockchain: "TheBlockchain",
		Network:    "TheNetwork",
		AddrPrefix: "test",
	}

	adapter := newAdapter(m, tendermint.NewClient(""), properties)
	resp, err2 := adapter.ConstructionSubmit(context.Background(), &types.ConstructionSubmitRequest{
		SignedTransaction: signedTx,
	})

	require.Nil(t, err2)
	require.NotNil(t, resp)
	require.Equal(t, expectedHash, resp.TransactionIdentifier.Hash)
}

func TestLaunchpad_ConstructionSubmit_FailsOfflineMode(t *testing.T) {
	properties := properties{
		OfflineMode: true,
	}
	adapter := newAdapter(sdk.NewClient(""), tendermint.NewClient(""), properties)

	_, err2 := adapter.ConstructionSubmit(context.Background(), &types.ConstructionSubmitRequest{
		SignedTransaction: "dkajfkdjkads",
	})

	require.Equal(t, ErrEndpointDisabledOfflineMode, err2)
}
