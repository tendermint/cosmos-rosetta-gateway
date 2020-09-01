package launchpad

import (
	"context"

	"github.com/coinbase/rosetta-sdk-go/types"
)

func (l Launchpad) ConstructionPayloads(ctx context.Context, req *types.ConstructionPayloadsRequest) (*types.ConstructionPayloadsResponse, *types.Error) {
	// We only support for now Transfer type of operation.
	if len(req.Operations) != 2 {
		return nil, ErrInvalidOperation
	}

	if req.Operations[0].Type != OperationTransfer || req.Operations[1].Type != OperationTransfer {
		return nil, ErrInvalidOperation
	}

	return nil, nil
}
