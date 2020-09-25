package launchpad

import (
	"context"

	"github.com/coinbase/rosetta-sdk-go/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	cryptoamino "github.com/tendermint/tendermint/crypto/encoding/amino"
)

func (l launchpad) ConstructionDerive(ctx context.Context, r *types.ConstructionDeriveRequest) (*types.ConstructionDeriveResponse, *types.Error) {
	if r.PublicKey.CurveType != "secp256k1" {
		return nil, ErrUnsupportedCurve
	}

	pubKey, err := cryptoamino.PubKeyFromBytes(r.PublicKey.Bytes)
	if err != nil {
		return nil, ErrInvalidPubkey
	}

	return &types.ConstructionDeriveResponse{
		Address: sdk.AccAddress(pubKey.Address().Bytes()).String(),
	}, nil
}
