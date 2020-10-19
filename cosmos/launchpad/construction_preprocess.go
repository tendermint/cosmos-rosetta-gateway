package launchpad

import (
	"context"

	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"

	"github.com/coinbase/rosetta-sdk-go/types"
)

func (l launchpad) ConstructionPreprocess(ctx context.Context, r *types.ConstructionPreprocessRequest) (*types.ConstructionPreprocessResponse, *types.Error) {
	operations := r.Operations
	if len(operations) < 1 {
		return nil, ErrInterpreting
	}

	msg, err := getMsgDataFromOperations(operations)
	if err != nil {
		return nil, rosetta.WrapError(ErrInvalidOperation, err.Error())
	}

	signers := msg.GetSigners()
	if len(signers) != 1 {
		return nil, rosetta.WrapError(ErrInvalidAddress, "invalid number of signers")
	}

	var res = &types.ConstructionPreprocessResponse{
		Options: map[string]interface{}{
			OptionAddress: signers[0].String(),
			OptionGas:     200000,
		},
	}
	return res, nil
}
