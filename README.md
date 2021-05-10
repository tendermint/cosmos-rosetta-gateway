
# Cosmos Rosetta Shared Lib

Shared lib used in Cosmos-SDK. This includes the libraries that are used by
different versions of the SDK, including Launchpad and Stargate.

## Badges

[![PkgGoDev](https://pkg.go.dev/badge/github.com/tendermint/cosmos-rosetta-gateway)](https://pkg.go.dev/github.com/tendermint/cosmos-rosetta-gateway)
[![codecov](https://codecov.io/gh/tendermint/cosmos-rosetta-gateway/branch/develop/graph/badge.svg)](https://codecov.io/gh/tendermint/cosmos-rosetta-gateway)
[![license](https://img.shields.io/github/license/tendermint/cosmos-rosetta-gateway.svg)](https://github.com/tendermint/cosmos-rosetta-gateway/blob/develop/LICENSE)
[![LoC](https://tokei.rs/b1/github/tendermint/cosmos-rosetta-gateway)](https://github.com/tendermint/cosmos-rosetta-gateway)


[![Rosetta sdk version](https://img.shields.io/badge/Rosetta%20SDK-v0.6.10-informational)](https://github.com/coinbase/rosetta-sdk-go/releases/tag/v0.6.10)
## Installation 

```bash 
go get github.com/tendermint/cosmos-rosetta-gateway
```

## Testing

Testing of the rosetta gateway can be performed by cloning the [cosmos-sdk](https://github.com/cosmos/cosmos-sdk) and running 

```bash
make test-rosetta
```
  
## Related

- Coinbase Rosetta: https://www.rosetta-api.org
- Cosmos SDK: https://github.com/cosmos/cosmos-sdk
