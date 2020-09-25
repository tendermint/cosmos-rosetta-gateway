// Package main exposes Rosetta API for Cosmos SDK as a standalone Gateway.
package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad"
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
	"github.com/tendermint/cosmos-rosetta-gateway/service"
)

var (
	flagAppRPC        = flag.String("app-rpc", "localhost:1317", "Application's RPC endpoint.")
	flagTendermintRPC = flag.String("tendermint-rpc", "localhost:26657", "Tendermint's RPC endpoint.")
	flagBlockchain    = flag.String("blockchain", "app", "Application's name (e.g. Cosmos Hub)")
	flagNetworkID     = flag.String("network", "network", "Network's identifier (e.g. cosmos-hub-3, testnet-1, etc)")
	flagOfflineMode   = flag.Bool("offline", false, "Flag that forces the rosetta service to run in offline mode, some endpoints won't work.")
	flagAddrPrefix    = flag.String("prefix", "cosmos", "Bech32 prefix of address (e.g. cosmos, iaa, xrn:)")
)

func main() {
	flag.Parse()

	if err := runHandler(); err != nil {
		fmt.Fprintln(flag.CommandLine.Output(), err)
		os.Exit(2)
	}
}

func runHandler() error {
	properties := rosetta.NetworkProperties{
		Blockchain: *flagBlockchain,
		Network:    *flagNetworkID,

		SupportedOperations: []string{launchpad.OperationTransfer},
	}

	h, err := service.New(
		service.Network{
			Properties: properties,
			Adapter: launchpad.NewLaunchpad(
				launchpad.Options{
					CosmosEndpoint:     *flagAppRPC,
					TendermintEndpoint: *flagTendermintRPC,
					Blockchain:         *flagBlockchain,
					Network:            *flagNetworkID,
					AddrPrefix:         *flagAddrPrefix,
					OfflineMode:        *flagOfflineMode,
				},
			),
		},
	)
	if err != nil {
		return err
	}

	server := &http.Server{
		Handler: h,
		Addr:    ":8080",
	}

	return server.ListenAndServe()
}
