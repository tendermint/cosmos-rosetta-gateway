[![PkgGoDev](https://pkg.go.dev/badge/github.com/tendermint/cosmos-rosetta-gateway)](https://pkg.go.dev/github.com/tendermint/cosmos-rosetta-gateway)
[![codecov](https://codecov.io/gh/tendermint/cosmos-rosetta-gateway/branch/develop/graph/badge.svg)](https://codecov.io/gh/tendermint/cosmos-rosetta-gateway)
[![Go Report Card](https://goreportcard.com/badge/github.com/tendermint/cosmos-rosetta-gateway)](https://goreportcard.com/report/tendermint/cosmos-rosetta-gateway)
[![license](https://img.shields.io/github/license/tendermint/cosmos-rosetta-gateway.svg)](https://github.com/tendermint/cosmos-rosetta-gateway/blob/develop/LICENSE)
[![LoC](https://tokei.rs/b1/github/tendermint/cosmos-rosetta-gateway)](https://github.com/tendermint/cosmos-rosetta-gateway)

# Cosmos Rosetta API Gateway (crg)
Cosmos Rosetta Gateway is a Rosetta API adapter for Cosmos SDK Chains running Cosmos-SDK **0.39 "Launchpad"** releases. Support for Cosmos-SDK **0.37** and Cosmos-SDK **0.40 "Stargate"** releases is coming soon.

This repository contains both a library and a standalone binary, named **crg**.

### Using crg

CRGis shipped in two forms for maximum flexiblity:

* Standalone executable `crg`
The standalone executable will talk to your blockchain's API and serve as an API adapter, making it compatible with the Coinbase Rosetta blockchain API interface.


* Library
The crg library runs as a part of your blockchain's daemon and directly provides a Rosetta-compatible API interface.
