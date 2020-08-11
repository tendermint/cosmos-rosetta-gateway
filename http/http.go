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

type Network struct {
	Options Options
	Adapter rosetta.Adapter
}

type Options struct {
	Blockchain          string
	Network             string
	SupportedOperations []string
}

func New(network Network) (*Service, error) {
	asserter, err := asserter.NewServer(
		network.Options.SupportedOperations,
		false,
		[]*types.NetworkIdentifier{
			{
				Blockchain: network.Options.Blockchain,
				Network:    network.Options.Network,
			},
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build asserter")
	}

	h := server.NewRouter(
		server.NewAccountAPIController(network.Adapter, asserter),
	)

	s := &Service{
		Handler: h,
	}

	return s, nil
}
