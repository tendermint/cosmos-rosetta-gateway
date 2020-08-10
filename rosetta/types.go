package rosetta

import "github.com/coinbase/rosetta-sdk-go/types"

type Rosetta interface {
	DataAPI
}

type DataAPI interface {
	AccountBalance(request *types.AccountBalanceRequest) *types.AccountBalanceResponse
}
