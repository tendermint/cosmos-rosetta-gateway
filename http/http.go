// Package http exposes Rosetta API over HTTP by wrapping functions from the
// crg/services package.
package http

import (
	"net/http"

	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"

	"github.com/coinbase/rosetta-sdk-go/asserter"
	"github.com/coinbase/rosetta-sdk-go/server"
	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/pkg/errors"
)

type Service struct {
	http.Handler
}

type Options struct {
	Blockchain string
	Network    string
}

func New(opts Options, adapter rosetta.Adapter) (*Service, error) {
	asserter, err := asserter.NewServer(
		[]string{"Transfer", "Reward"},
		false,
		[]*types.NetworkIdentifier{
			{
				Blockchain: opts.Blockchain,
				Network:    opts.Network,
			},
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build asserter")
	}

	h := server.NewRouter(
		server.NewAccountAPIController(adapter, asserter),
	)

	s := &Service{
		Handler: h,
	}

	return s, nil
}
