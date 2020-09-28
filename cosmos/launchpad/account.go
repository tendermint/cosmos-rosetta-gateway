package launchpad

import (
	"context"
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/coinbase/rosetta-sdk-go/types"
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

	block, err := l.tendermint.Block(uint64(resp.Height))
	if err != nil {
		return nil, rosetta.WrapError(ErrNodeConnection, err.Error())
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
