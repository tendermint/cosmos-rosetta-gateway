package stargate

import (
	"context"

	"github.com/coinbase/rosetta-sdk-go/types"
)

func (s Stargate) AccountBalance(ctx context.Context, request *types.AccountBalanceRequest) (*types.AccountBalanceResponse, *types.Error) {
	panic("implement me")
}
