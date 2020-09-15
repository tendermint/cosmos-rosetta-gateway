package launchpad

import (
	"context"
	"encoding/hex"
	"io/ioutil"
	"testing"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/altsdk"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

func TestLaunchpad_ConstructionCombine(t *testing.T) {
	pubKey, err := cosmostypes.GetPubKeyFromBech32("accpub", "cosmospub1addwnpepq2ngu5spnhp4qyt6zzlvdex5zncn5rrqscw6m9c6tn6hc4za4jyf60z3mtr")
	require.NoError(t, err)

	properties := rosetta.NetworkProperties{
		Blockchain: "TheBlockchain",
		Network:    "TheNetwork",
		SupportedOperations: []string{
			"Transfer",
			"Reward",
		},
	}

	bz, err := ioutil.ReadFile("./testdata/unsigned-tx.json")
	require.NoError(t, err)
	var stdTx auth.StdTx
	codec := simapp.MakeCodec()
	err = codec.UnmarshalJSON(bz, &stdTx)
	require.NoError(t, err)
	txBytes, err := codec.MarshalBinaryLengthPrefixed(stdTx)
	require.NoError(t, err)

	txHex := hex.EncodeToString(txBytes)
	adapter := NewLaunchpad(TendermintAPI{}, CosmosAPI{}, altsdk.NewClient(""), properties)
	var combineRes, combineErr = adapter.ConstructionCombine(context.Background(), &types.ConstructionCombineRequest{
		UnsignedTransaction: txHex,
		Signatures: []*types.Signature{{
			SigningPayload: &types.SigningPayload{
				Address: "cosmos1qrv8g4hwt4z6ds8mednhhgx907wug9d6y8n9jy",
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
	err = codec.UnmarshalBinaryLengthPrefixed(bz, &signedStdTx)
	require.NoError(t, err)
	require.Equal(t, stdTx.GetSigners(), signedStdTx.GetSigners())
}
