// Package http exposes Rosetta API over HTTP by wrapping functions from the
// crg/services package.
package http

import (
	"net/http"

	"github.com/coinbase/rosetta-sdk-go/asserter"
	"github.com/coinbase/rosetta-sdk-go/server"
	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/pkg/errors"
	"github.com/tendermint/cosmos-rosetta-gateway/crg/service"
)

type HTTPService struct {
	http.Handler
}

func New(service *service.Service) (*HTTPService, error) {
	asserter, err := asserter.NewServer([]string{}, false, []*types.NetworkIdentifier{})
	if err != nil {
		return nil, errors.Wrap(err, "cannot build asserter")
	}
	h := server.NewRouter(
		server.NewAccountAPIController(service, asserter),
	)
	s := &HTTPService{
		Handler: h,
	}
	return s, nil
}
