package launchpad

import (
	"context"
	"github.com/antihax/optional"
	sdk "github.com/cosmos/cosmos-sdk/types"

	openapi "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint/generated"

	"github.com/coinbase/rosetta-sdk-go/types"
)

func (l Launchpad) AccountBalance(ctx context.Context, request *types.AccountBalanceRequest) (
	*types.AccountBalanceResponse, *types.Error) {
	resp, err := l.altCosmos.GetAuthAccount(ctx, request.AccountIdentifier.Address)
	if err != nil {
		return nil, ErrNodeConnection
	}

	block, _, err := l.tendermint.Info.Block(ctx, &openapi.BlockOpts{
		Height: optional.NewFloat32(float32(resp.Height)),
	})
	if err != nil {
		return nil, ErrNodeConnection
	}

	return &types.AccountBalanceResponse{
		BlockIdentifier: &types.BlockIdentifier{
			Index: resp.Height,
			Hash:  block.Result.BlockId.Hash,
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
