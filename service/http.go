package service

import (
	"net/http"

	assert "github.com/coinbase/rosetta-sdk-go/asserter"
	"github.com/coinbase/rosetta-sdk-go/server"
	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/pkg/errors"

	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
)

type Service struct {
	http.Handler
}

type Network struct {
	Properties rosetta.NetworkProperties
	Adapter    rosetta.Adapter
}

func New(network Network) (*Service, error) {
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
		server.NewNetworkAPIController(network.Adapter, asserter),
		server.NewMempoolAPIController(network.Adapter, asserter),
	)

	s := &Service{
		Handler: h,
	}

	return s, nil
}
