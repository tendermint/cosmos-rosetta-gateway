// Package main exposes Rosetta API for Cosmos SDK as a standalone Gateway.
package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad"
	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/altsdk"
	cosmoslaunchpadclient "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/generated"
	tendermintlaunchpadclient "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint/generated"
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
	"github.com/tendermint/cosmos-rosetta-gateway/service"
)

var (
	flagAppRPC        = flag.String("app-rpc", "localhost:1317", "Application's RPC endpoint.")
	flagTendermintRPC = flag.String("tendermint-rpc", "localhost:26657", "Tendermint's RPC endpoint.")
	flagBlockchain    = flag.String("blockchain", "app", "Application's name (e.g. Cosmos Hub)")
	flagNetworkID     = flag.String("network", "network", "Network's identifier (e.g. cosmos-hub-3, testnet-1, etc)")
)

func main() {
	flag.Parse()

	if err := runHandler(); err != nil {
		fmt.Fprintln(flag.CommandLine.Output(), err)
		os.Exit(2)
	}
}

func runHandler() error {
	cosmoslpc := cosmoslaunchpadclient.NewAPIClient(&cosmoslaunchpadclient.Configuration{
		Host:   *flagAppRPC,
		Scheme: "http",
	})
	tendermintlpc := tendermintlaunchpadclient.NewAPIClient(&tendermintlaunchpadclient.Configuration{
		Host:   *flagTendermintRPC,
		Scheme: "http",
	})
	altClient := altsdk.NewClient(fmt.Sprintf("http://%s", *flagAppRPC))

	cosmoslp := launchpad.CosmosAPI{
		Auth:       cosmoslpc.AuthApi,
		Bank:       cosmoslpc.BankApi,
		Tendermint: cosmoslpc.TendermintRPCApi,
	}
	tendermintlp := launchpad.TendermintAPI{
		Info: tendermintlpc.InfoApi,
		Tx:   tendermintlpc.TxApi,
	}

	properties := rosetta.NetworkProperties{
		Blockchain:          *flagBlockchain,
		Network:             *flagNetworkID,
		SupportedOperations: []string{launchpad.OperationTransfer},
	}

	h, err := service.New(
		service.Network{
			Properties: properties,
			Adapter:    launchpad.NewLaunchpad(tendermintlp, cosmoslp, altClient, properties),
		},
	)
	// TODO: maybe create some constructor for specific adapters or Factory.
	if err != nil {
		return err
	}

	server := &http.Server{
		Handler: h,
		Addr:    ":8080",
	}

	return server.ListenAndServe()
}
