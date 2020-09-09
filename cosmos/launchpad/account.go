package launchpad

import (
	"context"
	"strconv"

	"github.com/antihax/optional"

	openapi "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint/generated"

	"github.com/coinbase/rosetta-sdk-go/types"

	client "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/generated"
)

func (l Launchpad) AccountBalance(ctx context.Context, request *types.AccountBalanceRequest) (
	*types.AccountBalanceResponse, *types.Error) {
	resp, _, err := l.cosmos.Bank.BankBalancesAddressGet(ctx, request.AccountIdentifier.Address)
	if err != nil {
		return nil, ErrNodeConnection
	}

	height, err := strconv.ParseInt(resp.Height, 10, 64)
	if err != nil {
		return nil, ErrInterpreting
	}

	block, _, err := l.tendermint.Info.Block(ctx, &openapi.BlockOpts{
		Height: optional.NewFloat32(float32(height)),
	})
	if err != nil {
		return nil, ErrNodeConnection
	}

	return &types.AccountBalanceResponse{
		BlockIdentifier: &types.BlockIdentifier{
			Index: height,
			Hash:  block.Result.BlockId.Hash,
		},
		Balances: convertCoinsToRosettaBalances(resp.Result),
	}, nil
}

func convertCoinsToRosettaBalances(coins []client.Coin) []*types.Amount {
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
