// Package cmd exposes Rosetta API for Cosmos SDK as a cli command that can be
// used as a (sub)command to start a standalone Gateway.
package cmd

import (
	"net/http"
	"time"

	"github.com/tendermint/cosmos-rosetta-gateway/rosetta/cosmos/launchpad"

	"github.com/spf13/cobra"
	crghttp "github.com/tendermint/cosmos-rosetta-gateway/http"
)

func New() *cobra.Command {
	c := &cobra.Command{
		Use:   "crg",
		Short: "Run Cosmos Rosetta Gateway as a service",
		RunE:  runHandler,
	}
	return c
}

func runHandler(cmd *cobra.Command, args []string) error {
	c := &http.Client{
		Timeout: time.Minute * 3,
	}
	var (
		supportedOperations = []string{"Transfer", "Reward"}
		blockhain           = "Test"
		network             = "Test"
	)
	h, err := crghttp.New(
		crghttp.Network{
			Blockchain: blockhain,
			Network:    network,

			Options: crghttp.Options{
				SupportedOperations: supportedOperations,
			},

			Adapter: launchpad.NewLaunchpad(c, "http://localhost:1317", blockhain, network, supportedOperations),
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
