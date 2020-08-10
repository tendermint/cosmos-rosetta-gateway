package launchpad

import (
	"context"
	"github.com/coinbase/rosetta-sdk-go/types"
)

func (l Launchpad) AccountBalance(ctx context.Context, request *types.AccountBalanceRequest) (
	*types.AccountBalanceResponse, *types.Error) {
	panic("implement me")
}
