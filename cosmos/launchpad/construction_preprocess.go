package launchpad

import (
	"context"

	"github.com/coinbase/rosetta-sdk-go/types"
)

func (l Launchpad) ConstructionPreprocess(ctx context.Context, r *types.ConstructionPreprocessRequest) (*types.ConstructionPreprocessResponse, *types.Error) {
	operations := r.Operations
	if len(operations) < 1 {
		return nil, ErrInterpreting
	}

	account := operations[0].Account

	for _, operation := range operations {
		if operation.Account != account {
			return nil, ErrInvalidOperations
		}
	}

	var res = &types.ConstructionPreprocessResponse{
		Options: map[string]interface{}{
			OptionAddress: account.Address,
			OptionGas:     r.SuggestedFeeMultiplier,
			// TODO: Add memo to options
		},
	}
	return res, nil
}
