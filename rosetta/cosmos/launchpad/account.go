package launchpad

import (
	"context"

	"github.com/coinbase/rosetta-sdk-go/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	cosmoslp "github.com/tendermint/cosmos-rosetta-gateway/generated/cosmos-launchpad"
)

func (l Launchpad) AccountBalance(ctx context.Context, request *types.AccountBalanceRequest) (
	*types.AccountBalanceResponse, *types.Error) {
	resp, _, err := l.api.Bank.BankBalancesAddressGet(ctx, request.AccountIdentifier.Address)
	if err != nil {
		return nil, ErrNodeConnection
	}

	return &types.AccountBalanceResponse{
		Balances: convertCoinsToRosettaBalances(resp.Result),
	}, nil
}

func convertCoinsToRosettaBalances(coins []cosmoslp.Coin) []*types.Amount {
	var amounts []*types.Amount

	for _, coin := range coins {
		amounts = append(amounts, &types.Amount{
			Value: coin.Amount,
			Currency: &types.Currency{
				Symbol: coin.Denom,
			},
		})
	}

	return amounts
}

type balanceResp struct {
	Result sdk.Coins
}
