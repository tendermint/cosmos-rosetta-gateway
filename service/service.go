// Package service exposes high level controller functions that comforts
// Rossetta API's definition which can be used to create clients such as a
// Rosetta API over HTTP or any other interfaces.
package service

import (
	"context"
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"

	"github.com/coinbase/rosetta-sdk-go/types"
)

type Service struct {
	adapter rosetta.Adapter
}

func New(adapter rosetta.Adapter) *Service {
	return &Service{
		adapter: adapter,
	}
}

func (s *Service) AccountBalance(ctx context.Context, req *types.AccountBalanceRequest) (
	*types.AccountBalanceResponse, *types.Error) {

	resp := s.adapter.AccountBalance(req)

	return resp, nil
}
