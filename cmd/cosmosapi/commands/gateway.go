// Package cmd exposes Rosetta API for Cosmos SDK as a cli command that can be
// used as a (sub)command to start a standalone Gateway.
package commands

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad"
	cosmoslaunchpadclient "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/generated"
	tendermintlaunchpadclient "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint/generated"
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
	"github.com/tendermint/cosmos-rosetta-gateway/service"
)

const (
	FlagCosmosRpc     = "cosmos-rpc"
	FlagTendermintRpc = "tendermint-rpc"
	FlagBlockchain    = "blockchain"
	FlagNetwork       = "network"
)

func New() *cobra.Command {
	c := &cobra.Command{
		Use:   "crg",
		Short: "Run Cosmos Rosetta Gateway as a service",
		RunE:  runHandler,
	}

	c.Flags().String(FlagCosmosRpc, "localhost:1317", "the cosmos rpc")
	c.Flags().String(FlagTendermintRpc, "localhost:26657", "the tendermint rpc")

	c.Flags().String(FlagBlockchain, "blockchain", "the name of the blockchain (e.g. Bitcoin)")
	c.Flags().String(FlagNetwork, "network", "the name of the network (e.g. mainnet)")

	return c
}

func runHandler(cmd *cobra.Command, args []string) error {
	cosmosRpc, err := cmd.Flags().GetString(FlagCosmosRpc)
	if err != nil {
		return fmt.Errorf("error with cosmos-rpc address")
	}

	tendermintRpc, err := cmd.Flags().GetString(FlagTendermintRpc)
	if err != nil {
		return fmt.Errorf("error with tendermint-rpc address")
	}

	blockchain, err := cmd.Flags().GetString(FlagBlockchain)
	if err != nil {
		return fmt.Errorf("error with blockchain name")
	}

	network, err := cmd.Flags().GetString(FlagNetwork)
	if err != nil {
		return fmt.Errorf("error with network name")
	}

	cosmoslpc := cosmoslaunchpadclient.NewAPIClient(&cosmoslaunchpadclient.Configuration{
		Host:   cosmosRpc,
		Scheme: "http",
	})
	tendermintlpc := tendermintlaunchpadclient.NewAPIClient(&tendermintlaunchpadclient.Configuration{
		Host:   tendermintRpc,
		Scheme: "http",
	})

	cosmoslp := launchpad.CosmosAPI{
		Bank:         cosmoslpc.BankApi,
		Tendermint:   cosmoslpc.TendermintRPCApi,
		Transactions: cosmoslpc.TransactionsApi,
	}
	tendermintlp := launchpad.TendermintAPI{
		Info: tendermintlpc.InfoApi,
	}

	properties := rosetta.NetworkProperties{
		Blockchain:          blockchain,
		Network:             network,
		SupportedOperations: []string{launchpad.OperationTransfer},
	}

	h, err := service.New(
		service.Network{
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
