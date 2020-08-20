// Package main exposes Rosetta API for Cosmos SDK as a standalone Gateway.
package main

import (
	"log"

	"github.com/tendermint/cosmos-rosetta-gateway/cmd/cosmosapi/commands"
)

func main() {
	if err := commands.New().Execute(); err != nil {
		log.Fatal(err)
	}
}
