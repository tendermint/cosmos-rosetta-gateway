package crg

import (
	"fmt"
	assert "github.com/coinbase/rosetta-sdk-go/asserter"
	"github.com/coinbase/rosetta-sdk-go/server"
	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
	"github.com/tendermint/cosmos-rosetta-gateway/service"
	"net/http"
)

type Settings struct {
	Network *types.NetworkIdentifier
	Client  rosetta.NodeClient
	Listen  string
}

func Serve(settings Settings) error {
	asserter, err := assert.NewServer(
		settings.Client.SupportedOperations(),
		true,
		[]*types.NetworkIdentifier{settings.Network},
		nil,
	)
	if err != nil {
		return fmt.Errorf("cannot build asserter: %w", err)
	}
	adapter, err := service.NewOnlineNetwork(settings.Client, settings.Network)
	if err != nil {
		return err
	}
	h := server.NewRouter(
		server.NewAccountAPIController(adapter, asserter),
		server.NewBlockAPIController(adapter, asserter),
		server.NewNetworkAPIController(adapter, asserter),
		server.NewMempoolAPIController(adapter, asserter),
		server.NewConstructionAPIController(adapter, asserter),
	)

	return http.ListenAndServe(settings.Listen, h)
}
