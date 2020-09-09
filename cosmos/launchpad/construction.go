package launchpad

import (
	"context"

	"github.com/coinbase/rosetta-sdk-go/types"
)

func (l Launchpad) ConstructionCombine(context.Context, *types.ConstructionCombineRequest) (*types.ConstructionCombineResponse, *types.Error) {
	panic("unimplemented")
}

func (l Launchpad) ConstructionSubmit(context.Context, *types.ConstructionSubmitRequest) (*types.TransactionIdentifierResponse, *types.Error) {
	panic("unimplemented")
}
