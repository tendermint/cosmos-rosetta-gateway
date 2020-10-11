package launchpad

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
	"strconv"

	"github.com/coinbase/rosetta-sdk-go/types"
)

func (l launchpad) ConstructionPreprocess(ctx context.Context, r *types.ConstructionPreprocessRequest) (*types.ConstructionPreprocessResponse, *types.Error) {
	operations := r.Operations
	if len(operations) < 1 {
		return nil, ErrInterpreting
	}

	txData, err := getTransferTxDataFromOperations(operations)
	if err != nil {
		return nil, rosetta.WrapError(ErrInvalidAddress, err.Error())
	}
	if txData.From == nil {
		return nil, rosetta.WrapError(ErrInvalidAddress, err.Error())
	}

	if len(r.MaxFee) != 1 {
		return nil, rosetta.WrapError(ErrInvalidFee, "multiple fees not supported")
	}

	fee := r.MaxFee[0]
	amount, err := strconv.ParseInt(fee.Value, 10, 64)
	feeCoins := sdk.NewCoin(fee.Currency.Symbol, sdk.NewInt(amount))
	memo, ok := r.Metadata["memo"]
	if !ok {
		return nil, ErrInvalidFee
	}

	var res = &types.ConstructionPreprocessResponse{
		Options: map[string]interface{}{
			OptionAddress: txData.From.String(),
			OptionGas:     r.SuggestedFeeMultiplier,
			OptionFee:     feeCoins.String(),
			OptionMemo:    memo,
		},
	}
	return res, nil
}
