package launchpad

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/coinbase/rosetta-sdk-go/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

const AccountSdkHandler = "bank/balances/"

func (l Launchpad) AccountBalance(ctx context.Context, request *types.AccountBalanceRequest) (
	*types.AccountBalanceResponse, *types.Error) {

	get, err := http.Get(
		fmt.Sprintf("%s%s%s", l.endpoint, AccountSdkHandler,
			request.AccountIdentifier.Address))
	if err != nil {
		return nil, rosetta.NewError(1, "error getting data from node")
	}

	body := get.Body

	all, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, rosetta.NewError(1, "error reading data from node")
	}
	defer body.Close()

	var res balanceResp
	err = json.Unmarshal(all, &res)
	if err != nil {
		return nil, rosetta.NewError(1, "error interpreting data from node")
	}

	return &types.AccountBalanceResponse{
		BlockIdentifier: nil,
		Balances:        convertCoinsToRosettaBalances(res.Result),
	}, nil
}

func convertCoinsToRosettaBalances(coins sdk.Coins) []*types.Amount {
	var amounts []*types.Amount

	for _, coin := range coins {
		amounts = append(amounts, &types.Amount{
			Value: coin.Amount.String(),
			Currency: &types.Currency{
				Symbol:   coin.Denom,
				Decimals: 0,
			},
		})
	}

	return amounts
}

type balanceResp struct {
	Result sdk.Coins
}
