package launchpad

import (
	"context"
	"fmt"
	"strconv"
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
// We assume that it comes formated in the correct way. And that the balance of the sender is the same
// as the receiver operations.
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

			amount, err := strconv.ParseInt(op.Amount.Value, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid amount")
			}

			transferData.Amount = cosmostypes.NewCoin(op.Amount.Currency.Symbol, cosmostypes.NewInt(amount))
		}
	}

	return transferData, nil
}
