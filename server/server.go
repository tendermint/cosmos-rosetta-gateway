package server

import (
	"fmt"
	assert "github.com/coinbase/rosetta-sdk-go/asserter"
	"github.com/coinbase/rosetta-sdk-go/server"
	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/tendermint/cosmos-rosetta-gateway/internal/service"
	crgtypes "github.com/tendermint/cosmos-rosetta-gateway/types"
	"net/http"
)

type Settings struct {
	Network *types.NetworkIdentifier
	Client  crgtypes.NodeClient
	Listen  string
}

type Handler struct {
	h    http.Handler
	addr string
}

func (h Handler) Start() error {
	return http.ListenAndServe(h.addr, h.h)
}

func NewHandler(settings Settings) (Handler, error) {
	asserter, err := assert.NewServer(
		settings.Client.SupportedOperations(),
		true,
		[]*types.NetworkIdentifier{settings.Network},
		nil,
	)
	if err != nil {
		return Handler{}, fmt.Errorf("cannot build asserter: %w", err)
	}
	adapter, err := service.NewOnlineNetwork(settings.Client, settings.Network)
	if err != nil {
		return Handler{}, err
	}
	h := server.NewRouter(
		server.NewAccountAPIController(adapter, asserter),
		server.NewBlockAPIController(adapter, asserter),
		server.NewNetworkAPIController(adapter, asserter),
		server.NewMempoolAPIController(adapter, asserter),
		server.NewConstructionAPIController(adapter, asserter),
	)

	return Handler{
		h:    h,
		addr: settings.Listen,
	}, nil
}
