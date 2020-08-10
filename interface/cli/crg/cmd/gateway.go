// Package cmd exposes Rosetta API for Cosmos SDK as a cli command that can be
// used as a (sub)command to start a standalone Gateway.
package cmd

import (
	"net/http"

	"github.com/spf13/cobra"
	crghttp "github.com/tendermint/cosmos-rosetta-gateway/interface/http"
	"github.com/tendermint/cosmos-rosetta-gateway/service"
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
	s := service.New()
	h, err := crghttp.New(s)
	if err != nil {
		return err
	}
	hserver := &http.Server{
		Handler: h,
	}
	return hserver.ListenAndServe()
}
