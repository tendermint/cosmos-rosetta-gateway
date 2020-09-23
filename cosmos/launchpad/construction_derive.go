package launchpad

import (
	"context"

	"github.com/coinbase/rosetta-sdk-go/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (l Launchpad) ConstructionDerive(ctx context.Context, r *types.ConstructionDeriveRequest) (*types.ConstructionDeriveResponse, *types.Error) {
	if r.PublicKey.CurveType != "secp256k1" {
		return nil, ErrUnsupportedCurve
	}

	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(
		l.properties.AddrPrefix,
		l.properties.AddrPrefix+sdk.PrefixPublic)

	return &types.ConstructionDeriveResponse{
		Address: sdk.AccAddress(r.PublicKey.Bytes).String(),
	}, nil
}
