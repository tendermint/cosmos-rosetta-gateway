package launchpad

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/coinbase/rosetta-sdk-go/types"
)

func (l Launchpad) AccountBalance(ctx context.Context, request *types.AccountBalanceRequest) (
	*types.AccountBalanceResponse, *types.Error) {
	resp, err := l.altCosmos.GetAuthAccount(ctx, request.AccountIdentifier.Address)
	if err != nil {
		return nil, ErrNodeConnection
	}

	block, err := l.altTendermint.Block(uint64(resp.Height))
	if err != nil {
		return nil, ErrNodeConnection
	}
	if err != nil {
		return nil, ErrNodeConnection
	}

	return &types.AccountBalanceResponse{
		BlockIdentifier: &types.BlockIdentifier{
			Index: resp.Height,
			Hash:  block.BlockId.Hash,
		},
		Balances: convertCoinsToRosettaBalances(resp.Result.Value.Coins),
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
