[![PkgGoDev](https://pkg.go.dev/badge/github.com/tendermint/cosmos-rosetta-gateway)](https://pkg.go.dev/github.com/tendermint/cosmos-rosetta-gateway)
[![codecov](https://codecov.io/gh/tendermint/cosmos-rosetta-gateway/branch/develop/graph/badge.svg)](https://codecov.io/gh/tendermint/cosmos-rosetta-gateway)
[![license](https://img.shields.io/github/license/tendermint/cosmos-rosetta-gateway.svg)](https://github.com/tendermint/cosmos-rosetta-gateway/blob/develop/LICENSE)
[![LoC](https://tokei.rs/b1/github/tendermint/cosmos-rosetta-gateway)](https://github.com/tendermint/cosmos-rosetta-gateway)

# Cosmos Rosetta API Gateway (crg)
Cosmos Rosetta Gateway is a Rosetta API adapter for Cosmos SDK Chains running Cosmos-SDK **0.39 "Launchpad"** releases.

Support for Cosmos-SDK **0.37** and Cosmos-SDK **0.40 "Stargate"** releases is coming soon.

This repository contains both a library and a standalone binary, named **crg**.  Documentation on how to use the library is coming soon.

### Using crg
The standalone executable will talk to your blockchain's API and serve as an API adapter, making it compatible with the Coinbase Rosetta blockchain API interface.

**Configuring crg**
`crg` can be configured to match your chain using flags

| Task  | Flag  | Description |
|---|---|---|
| Help | --help | explain these flags |
| Port | --port 8080  | the port where the service is listening |
| App RPC Endpoint  | --app-rpc localhost:1317  |       |
| Tendermint RPC Endpoint  | --tendermint-rpc localhost:2665  |       |
| Application Name  | --blockchain gaia  |  example: "Bitcoin |
| Network ID | --network mainnet-1 |  In cosmos, usually the chain's version, like gaia hub 3 |
| Offline Mode | --offline true | this allows crg to run even when there's no chain for it to connect to  |
| Prefix | --prefix cosmos | bech32 prefix for addresses on your blockchain    |

**Quick Test**:

This test assumes that you've got a running blockchain that uses Cosmos-SDK 0.39.x, and that you have Go installed on your computer.


Install and use `crg`
```bash
screen -S crg
git clone https://github.com/tendermint/cosmos-rosetta-gateway
cd cosmos-rosetta-gateway
make dev
```

Hold `ctrl` and press A, then D. You'll leave the screen session and `crg` will run in the background.

`crg` runs an http server on port 8080, and here are some query examples using `curl`:


**Check Network Status**
```bash
curl --location --request POST 'http://localhost:8080/network/status' \
--header 'Content-Type: text/plain' \
--data-raw '{
    "network_identifier": {
        "blockchain": "Test",
        "network": "Test"
    }
}'
```

Success looks like:
```json
{
  "current_block_identifier": {
    "index": 46564,
    "hash": "2F4A700C064C0E66792DB80387035401421A985B1E4E10419E85F24E815E9D86"
  },
  "current_block_timestamp": 1601050588601,
  "genesis_block_identifier": {
    "index": 1,
    "hash": "8FC19EA07352344DA72C1CB141C945A8FC6C9349FD5244DB6B9C891C17747E12"
  },
  "peers": null
}
```


**Get Info on a particular block:**
```bash
curl --location --request POST 'http://localhost:8080/block' \
--header 'Content-Type: text/plain' \
--data-raw '{
    "network_identifier": {
        "blockchain": "Test",
        "network": "Test"
    },
    "block_identifier": {
        "index": 17807
    }
}'
```

Success looks like:
```json
{
  "block": {
    "block_identifier": {
      "index": 17807,
      "hash": "8C78CBFA84AFC57E20E379B1135C7EE6A14CE115291C8241750505D4FFDDA261"
    },
    "parent_block_identifier": {
      "index": 17806,
      "hash": "C0D4B2ED7B3DFEC3BF9673E132E6C32AA6AB3E1D566FB09AEA4292DA5FFDC349"
    },
    "timestamp": 1600905197777,
    "transactions": null
  }
}
```

Additional requests are documented in the cosmos-rosetta-gateway Postman Collecton.


### Postman Collection

[Postman](https://postman.io) is a popular tool for working with and designing API systems.  We've made a [postman collection](https://www.postman.com/collections/0bb4205306d904245eee) that contains a number common calls to both the data and construction APIs.
