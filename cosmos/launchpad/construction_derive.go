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

	return &types.ConstructionDeriveResponse{
		Address: sdk.AccAddress(r.PublicKey.Bytes).String(),
	}, nil
}
