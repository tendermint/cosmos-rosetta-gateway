package launchpad

import (
	"context"

	"github.com/coinbase/rosetta-sdk-go/types"
	cosmostypes "sdk-v0.39.1/types"
)

func (l Launchpad) ConstructionDerive(ctx context.Context, r *types.ConstructionDeriveRequest) (*types.ConstructionDeriveResponse, *types.Error) {
	return &types.ConstructionDeriveResponse{
		Address: cosmostypes.AccAddress(r.PublicKey.Bytes).String(),
	}, nil
}
