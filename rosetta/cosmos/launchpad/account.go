package launchpad

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/coinbase/rosetta-sdk-go/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

const AccountSdkHandler = "/bank/balances"

func (l Launchpad) AccountBalance(ctx context.Context, request *types.AccountBalanceRequest) (
	*types.AccountBalanceResponse, *types.Error) {

	addr := fmt.Sprintf("%s%s/%s", l.endpoint, AccountSdkHandler, request.AccountIdentifier.Address)
	resp, err := l.c.Get(addr)
	if err != nil {
		return nil, rosetta.NewError(1, "error getting data from node")
	}
	defer resp.Body.Close()

	var res balanceResp

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, rosetta.NewError(1, "error interpreting data from node")
	}

	return &types.AccountBalanceResponse{
		Balances: convertCoinsToRosettaBalances(res.Result),
	}, nil
}

func convertCoinsToRosettaBalances(coins sdk.Coins) []*types.Amount {
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

type balanceResp struct {
	Result sdk.Coins
}
