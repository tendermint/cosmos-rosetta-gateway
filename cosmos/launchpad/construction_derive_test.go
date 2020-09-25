package launchpad

import (
	"context"
	"encoding/hex"
	"testing"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/stretchr/testify/require"
)

func TestLaunchpad_ConstructionDerive(t *testing.T) {
	data, err := hex.DecodeString("eb5ae98721020d5cdec9d5d170a8d6005e6e25505c423c77dd1431e96f294008e0cb5b8d5945")
	require.NoError(t, err)

	properties := properties{
		Blockchain: "TheBlockchain",
		Network:    "TheNetwork",
		AddrPrefix: "test",
	}

	adapter := newAdapter(sdk.NewClient(""), tendermint.NewClient(""), properties)
	deriveResp, deriveErr := adapter.ConstructionDerive(context.Background(), &types.ConstructionDeriveRequest{
		PublicKey: &types.PublicKey{
			Bytes:     data,
			CurveType: "secp256k1",
		},
	})
	require.Nil(t, deriveErr)
	require.NotNil(t, deriveResp)
	require.Equal(t, "test1rlv9zzy95k6wv2r8g7lqpcsu6hrxcpn6tdap74", deriveResp.Address)

	// TODO: Use table driven tests
	// check unsupported curve returns error
	_, deriveErr = adapter.ConstructionDerive(context.Background(), &types.ConstructionDeriveRequest{
		PublicKey: &types.PublicKey{
			Bytes:     data,
			CurveType: "edwards25519",
		},
	})
	require.Equal(t, ErrUnsupportedCurve, deriveErr)
}
