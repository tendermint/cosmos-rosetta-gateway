[![PkgGoDev](https://pkg.go.dev/badge/github.com/tendermint/cosmos-rosetta-gateway)](https://pkg.go.dev/github.com/tendermint/cosmos-rosetta-gateway)
[![codecov](https://codecov.io/gh/tendermint/cosmos-rosetta-gateway/branch/develop/graph/badge.svg)](https://codecov.io/gh/tendermint/cosmos-rosetta-gateway)
[![license](https://img.shields.io/github/license/tendermint/cosmos-rosetta-gateway.svg)](https://github.com/tendermint/cosmos-rosetta-gateway/blob/develop/LICENSE)
[![LoC](https://tokei.rs/b1/github/tendermint/cosmos-rosetta-gateway)](https://github.com/tendermint/cosmos-rosetta-gateway)

# Cosmos Rosetta API Gateway (crg)
Cosmos Rosetta Gateway is a Rosetta API adapter for Cosmos SDK Chains running Cosmos-SDK **0.39 "Launchpad"** releases.

Support for Cosmos-SDK **0.37** and Cosmos-SDK **0.40 "Stargate"** releases is coming soon.

This repository contains both a library and a standalone binary, named **crg**.

### Using crg

CRGis shipped in two forms for maximum flexiblity:

* Standalone executable `crg`
The standalone executable will talk to your blockchain's API and serve as an API adapter, making it compatible with the Coinbase Rosetta blockchain API interface.

**Quick Test**:

This test assumes that you've got a running blockhain that uses Cosmos-SDK 0.39.x, and that you have Go installed on your computer.

Install and use `crg`
```bash
screen -S crg
git clone https://github.com/tendermint/cosmos-rosetta-gateway
cd cosmos-rosetta-gateway
make dev
./crd
```

Hold `ctrl` and press A, then D. You'll leave the screen session and `crg` will run in the background.

Query the blockchain `clay` using the Rosetta API provided by `crg`

```
curl --location --request POST 'http://localhost:8080/network/status' \
--header 'Content-Type: text/plain' \
--data-raw '{
    "network_identifier": {
        "blockchain": "Test",
        "network": "Test"
    }
}'
```


### Postman Collection

[Postman](https://postman.io) is a popular tool for working with and designing API systems.  We've made a [postman collection](https://www.postman.com/collections/0bb4205306d904245eee)
