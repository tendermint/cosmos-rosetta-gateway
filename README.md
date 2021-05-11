
# Cosmos Rosetta Shared Lib

This project provides the shared lib that implements the Coinbase Rosetta API used in Cosmos-SDK. This includes the libraries that are used by different versions of the SDK, including Launchpad and Stargate.


[![Rosetta SDK version](https://img.shields.io/badge/Rosetta%20SDK-v0.6.10-informational)](https://github.com/coinbase/rosetta-sdk-go/releases/tag/v0.6.10)

---

[![PkgGoDev](https://pkg.go.dev/badge/github.com/tendermint/cosmos-rosetta-gateway)](https://pkg.go.dev/github.com/tendermint/cosmos-rosetta-gateway)
[![codecov](https://codecov.io/gh/tendermint/cosmos-rosetta-gateway/branch/develop/graph/badge.svg)](https://codecov.io/gh/tendermint/cosmos-rosetta-gateway)
[![license](https://img.shields.io/github/license/tendermint/cosmos-rosetta-gateway.svg)](https://github.com/tendermint/cosmos-rosetta-gateway/blob/develop/LICENSE)
[![LoC](https://tokei.rs/b1/github/tendermint/cosmos-rosetta-gateway)](https://github.com/tendermint/cosmos-rosetta-gateway)

## Installation 

```bash 
go get -u github.com/tendermint/cosmos-rosetta-gateway
```

## Testing

Testing of the rosetta gateway can be performed by:

- cloning the [cosmos-sdk](https://github.com/cosmos/cosmos-sdk)
- updating the `cosmos-rosetta-gateway` dependency 
- and running `make test-rosetta`
  
## Resources

- Cosmos SDK: https://github.com/cosmos/cosmos-sdk
- Coinbase Rosetta API: https://www.rosetta-api.org

