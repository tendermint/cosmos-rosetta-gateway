package launchpad

import (
	"context"
	"encoding/hex"
	"io/ioutil"
	"testing"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

func TestLaunchpad_ConstructionCombine(t *testing.T) {
	properties := rosetta.NetworkProperties{
		Blockchain: "TheBlockchain",
		Network:    "TheNetwork",
		AddrPrefix: "test",
		SupportedOperations: []string{
			"Transfer",
			"Reward",
		},
	}

	adapter := NewLaunchpad(sdk.NewClient(""), tendermint.NewClient(""), properties)
	bz, err := ioutil.ReadFile("./testdata/unsigned-tx.json")
	require.NoError(t, err)

	var stdTx auth.StdTx
	codec := simapp.MakeCodec()
	err = codec.UnmarshalJSON(bz, &stdTx)
	require.NoError(t, err)
	txBytes, err := codec.MarshalJSON(stdTx)
	require.NoError(t, err)
	txHex := hex.EncodeToString(txBytes)

	pubKey, err := cosmostypes.GetPubKeyFromBech32("accpub", "testpub1addwnpepq2ngu5spnhp4qyt6zzlvdex5zncn5rrqscw6m9c6tn6hc4za4jyf6dj36w7")
	require.NoError(t, err)
	var combineRes, combineErr = adapter.ConstructionCombine(context.Background(), &types.ConstructionCombineRequest{
		UnsignedTransaction: txHex,
		Signatures: []*types.Signature{{
			SigningPayload: &types.SigningPayload{
				Address: "test1qrv8g4hwt4z6ds8mednhhgx907wug9d6y8n9jy",
				Bytes:   txBytes,
			},
			PublicKey: &types.PublicKey{
				CurveType: types.Secp256k1,
				Bytes:     pubKey.Bytes(),
			},
			SignatureType: types.Ecdsa,
			// uses random bytes as signing is out of scope for rosetta
			Bytes: txBytes,
		},
		}})
	require.Nil(t, combineErr)
	require.NotNil(t, combineRes)

	bz, err = hex.DecodeString(combineRes.SignedTransaction)
	require.NoError(t, err)
	var signedStdTx auth.StdTx
	err = codec.UnmarshalJSON(bz, &signedStdTx)
	require.NoError(t, err)
	require.Equal(t, stdTx.GetSigners(), signedStdTx.GetSigners())
}
