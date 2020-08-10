// Package service exposes high level controller functions that comforts
// Rossetta API's definition which can be used to create clients such as a
// Rosetta API over HTTP or any other interfaces.
package service

import (
	"context"

	"github.com/coinbase/rosetta-sdk-go/types"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) AccountBalance(context.Context, *types.AccountBalanceRequest) (
	*types.AccountBalanceResponse, *types.Error) {
	return nil, nil
}
