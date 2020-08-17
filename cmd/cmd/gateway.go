// Package cmd exposes Rosetta API for Cosmos SDK as a cli command that can be
// used as a (sub)command to start a standalone Gateway.
package cmd

import (
	"net/http"

	"github.com/spf13/cobra"

	crghttp "github.com/tendermint/cosmos-rosetta-gateway/http"
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta/cosmos/launchpad"
	cosmoslaunchpadclient "github.com/tendermint/cosmos-rosetta-gateway/rosetta/cosmos/launchpad/client/cosmos/generated"
	tendermintlaunchpadclient "github.com/tendermint/cosmos-rosetta-gateway/rosetta/cosmos/launchpad/client/tendermint/generated"
)

func New() *cobra.Command {
	c := &cobra.Command{
		Use:   "crg",
		Short: "Run Cosmos Rosetta Gateway as a service",
		RunE:  runHandler,
	}

	return c
}

func runHandler(*cobra.Command, []string) error {
	cosmoslpc := cosmoslaunchpadclient.NewAPIClient(&cosmoslaunchpadclient.Configuration{
		Host:   "localhost:1317",
		Scheme: "http",
	})
	tendermintlpc := tendermintlaunchpadclient.NewAPIClient(&tendermintlaunchpadclient.Configuration{
		Host:   "localhost:26657",
		Scheme: "http",
	})

	cosmoslp := launchpad.CosmosAPI{
		Bank:       cosmoslpc.BankApi,
		Tendermint: cosmoslpc.TendermintRPCApi,
	}
	tendermintlp := launchpad.TendermintAPI{
		Info: tendermintlpc.InfoApi,
	}

	properties := rosetta.NetworkProperties{
		Blockchain:          "Test",
		Network:             "Test",
		SupportedOperations: []string{"Transfer", "Reward"},
	}

	h, err := crghttp.New(
		crghttp.Network{
			Properties: properties,
			Adapter:    launchpad.NewLaunchpad(tendermintlp, cosmoslp, properties),
		},
	) // TODO: maybe create some constructor for specific adapters or Factory.
	if err != nil {
		return err
	}

	hserver := &http.Server{
		Handler: h,
		Addr:    ":8080",
	}
	return hserver.ListenAndServe()
}
