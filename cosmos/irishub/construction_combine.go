package irishub

import (
	"context"

	"github.com/coinbase/rosetta-sdk-go/types"
)

// ConstructionCombine implements the /construction/combine endpoint.
func (l launchpad) ConstructionCombine(ctx context.Context, r *types.ConstructionCombineRequest,
) (*types.ConstructionCombineResponse, *types.Error) {
	return nil, nil
}
