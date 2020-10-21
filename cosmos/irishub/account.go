package irishub

import (
	"context"
	"strconv"

	"github.com/coinbase/rosetta-sdk-go/types"
	sdk "github.com/irisnet/irishub/types"
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

func (l launchpad) AccountBalance(ctx context.Context, request *types.AccountBalanceRequest) (
	*types.AccountBalanceResponse, *types.Error) {
	if l.properties.OfflineMode {
		return nil, ErrEndpointDisabledOfflineMode
	}

	var reqHeight int64
	if request.BlockIdentifier != nil {
		reqHeight = *request.BlockIdentifier.Index
	}
	resp, err := l.cosmos.GetAuthAccount(ctx, request.AccountIdentifier.Address, reqHeight)
	if err != nil {
		return nil, rosetta.WrapError(ErrNodeConnection, err.Error())
	}

	block, err := l.tendermint.Block(0)
	if err != nil {
		return nil, rosetta.WrapError(ErrNodeConnection, err.Error())
	}

	height, err := strconv.Atoi(block.Block.Header.Height)
	if err != nil {
		return nil, rosetta.WrapError(ErrInterpreting, err.Error())
	}

	return &types.AccountBalanceResponse{
		BlockIdentifier: &types.BlockIdentifier{
			Index: int64(height),
			Hash:  block.BlockMeta.BlockId.Hash,
		},
		Balances: convertCoinsToRosettaBalances(resp.Value.Coins),
		Coins:    []*types.Coin{},
	}, nil
}

func convertCoinsToRosettaBalances(coins []sdk.Coin) []*types.Amount {
	var amounts []*types.Amount

	for _, coin := range coins {
		amounts = append(amounts, &types.Amount{
			Value: coin.Amount.String(),
			Currency: &types.Currency{
				Symbol: coin.Denom,
			},
		})
	}

	return amounts
}
