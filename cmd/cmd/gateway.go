// Package cmd exposes Rosetta API for Cosmos SDK as a cli command that can be
// used as a (sub)command to start a standalone Gateway.
package cmd

import (
	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad"
	"net/http"

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
	h, err := crghttp.New(launchpad.Launchpad{}) // TODO: maybe create some constructor for specific adapters or Factory.
	if err != nil {
		return err
	}
	hserver := &http.Server{
		Handler: h,
	}
	return hserver.ListenAndServe()
}
