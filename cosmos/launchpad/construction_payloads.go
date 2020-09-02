package launchpad

import (
	"context"
	"strings"

	"github.com/cosmos/cosmos-sdk/x/bank"

	"github.com/coinbase/rosetta-sdk-go/types"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
)

func (l Launchpad) ConstructionPayloads(ctx context.Context, req *types.ConstructionPayloadsRequest) (*types.ConstructionPayloadsResponse, *types.Error) {
	// We only support for now Transfer type of operation.
	if len(req.Operations) != 2 {
		return nil, ErrInvalidOperation
	}

	if req.Operations[0].Type != OperationTransfer || req.Operations[1].Type != OperationTransfer {
		return nil, ErrInvalidOperation
	}

	transferData, err := getFromAndToAddressFromOperations(req.Operations)
	if err != nil {
		return nil, ErrInvalidOperation
	}

	bank.NewMsgSend(transferData.From, transferData.To, nil)

	return nil, nil
}

// getFromAndToAddressFromOperations extracts the from and to addresses from a list of operations.
// last is to.
func getFromAndToAddressFromOperations(ops []*types.Operation) (*TransferTxData, error) {
	var (
		transferData = &TransferTxData{}
		err          error
	)

	for _, op := range ops {
		if strings.HasPrefix(op.Amount.Value, "-") {
			transferData.From, err = cosmostypes.AccAddressFromBech32(op.Account.Address)
			if err != nil {
				return nil, err
			}
		} else {
			transferData.To, err = cosmostypes.AccAddressFromBech32(op.Account.Address)
			if err != nil {
				return nil, err
			}
		}
	}

	return transferData, nil
}
