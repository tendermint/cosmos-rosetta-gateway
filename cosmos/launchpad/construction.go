package launchpad

import (
	"context"

	"github.com/coinbase/rosetta-sdk-go/types"
)

func (l Launchpad) ConstructionCombine(context.Context, *types.ConstructionCombineRequest) (*types.ConstructionCombineResponse, *types.Error) {
	panic("unimplemented")
}

func (l Launchpad) ConstructionHash(context.Context, *types.ConstructionHashRequest) (*types.TransactionIdentifierResponse, *types.Error) {
	panic("unimplemented")
}

func (l Launchpad) ConstructionMetadata(context.Context, *types.ConstructionMetadataRequest) (*types.ConstructionMetadataResponse, *types.Error) {
	panic("unimplemented")
}

func (l Launchpad) ConstructionPayloads(context.Context, *types.ConstructionPayloadsRequest) (*types.ConstructionPayloadsResponse, *types.Error) {
	panic("unimplemented")
}

func (l Launchpad) ConstructionPreprocess(context.Context, *types.ConstructionPreprocessRequest) (*types.ConstructionPreprocessResponse, *types.Error) {
	panic("unimplemented")
}

func (l Launchpad) ConstructionSubmit(context.Context, *types.ConstructionSubmitRequest) (*types.TransactionIdentifierResponse, *types.Error) {
	panic("unimplemented")
}
