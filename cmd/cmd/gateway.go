// Package cmd exposes Rosetta API for Cosmos SDK as a cli command that can be
// used as a (sub)command to start a standalone Gateway.
package cmd

import (
	"net/http"
	"time"

	"github.com/spf13/cobra"

	cosmoslp "github.com/tendermint/cosmos-rosetta-gateway/generated/cosmos-launchpad"
	crghttp "github.com/tendermint/cosmos-rosetta-gateway/http"
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta/cosmos/launchpad"
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
	c := &http.Client{
		Timeout: time.Minute * 3,
	}

	cc := cosmoslp.NewAPIClient(&cosmoslp.Configuration{
		Host:   "localhost:1317",
		Scheme: "http",
	})
	api := launchpad.API{
		Bank:       cc.BankApi,
		Tendermint: cc.TendermintRPCApi,
	}

	properties := rosetta.NetworkProperties{
		Blockchain:          "Test",
		Network:             "Test",
		SupportedOperations: []string{"Transfer", "Reward"},
	}

	h, err := crghttp.New(
		crghttp.Network{
			Properties: properties,
			Adapter: launchpad.NewLaunchpad(
				c, api, "http://localhost:26657", properties),
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
