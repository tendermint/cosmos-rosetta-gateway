// Package http exposes Rosetta API over HTTP by wrapping functions from the
// crg/services package.
package service

import (
	"fmt"
	"net/http"

	assert "github.com/coinbase/rosetta-sdk-go/asserter"
	"github.com/coinbase/rosetta-sdk-go/server"
	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/pkg/errors"

	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

type Service interface {
	Start() error
}

type service struct {
	h       http.Handler
	options Options
}

type Options struct {
	Port uint32
}

type Network struct {
	Properties rosetta.NetworkProperties
	Adapter    rosetta.Adapter
}

func New(options Options, network Network) (Service, error) {
	asserter, err := assert.NewServer(
		network.Properties.SupportedOperations,
		false,
		[]*types.NetworkIdentifier{
			{
				Blockchain: network.Properties.Blockchain,
				Network:    network.Properties.Network,
			},
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build asserter")
	}

	h := server.NewRouter(
		server.NewAccountAPIController(network.Adapter, asserter),
		server.NewBlockAPIController(network.Adapter, asserter),
		server.NewNetworkAPIController(network.Adapter, asserter),
		server.NewMempoolAPIController(network.Adapter, asserter),
		server.NewConstructionAPIController(network.Adapter, asserter),
	)

	s := &service{
		h:       h,
		options: options,
	}

	return s, nil
}

func (s service) Start() error {
	return http.ListenAndServe(fmt.Sprintf(":%d", s.options.Port), s.h)
}
