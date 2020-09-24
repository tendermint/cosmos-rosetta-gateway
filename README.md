[![PkgGoDev](https://pkg.go.dev/badge/github.com/tendermint/cosmos-rosetta-gateway)](https://pkg.go.dev/github.com/tendermint/cosmos-rosetta-gateway)
[![codecov](https://codecov.io/gh/tendermint/cosmos-rosetta-gateway/branch/develop/graph/badge.svg)](https://codecov.io/gh/tendermint/cosmos-rosetta-gateway)
[![Go Report Card](https://goreportcard.com/badge/github.com/tendermint/cosmos-rosetta-gateway)](https://goreportcard.com/report/tendermint/cosmos-rosetta-gateway)
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

This assumes that you're using a "fresh" Ubuntu or Debian virtual machine.  

You'll need [Go](https://golang.org) and nodejs 14.  We recommend that you use [NVM](https://github.com/nvm-sh/nvm) to install nodejs.  

Install NVM and Go
```bash
# NVM
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.35.3/install.sh | bash
export NVM_DIR="$HOME/.nvm"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # This loads nvm
[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"  # This loads nvm bash_completion

# Go
wget https://golang.org/dl/go1.15.2.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.15.2.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

Get a simple starport chain running in another session
```bash
sudo apt install -y screen git build-essential
screen -S starport
npm i -g @tendermint/starport
starport app github.com/yourname/clay
cd clay
starport serve --verbose
```

Hold `ctrl` and press A, then D. You'll leave the screen session and the blockchain `clay` will run in the background.


Install and use `crd`
```bash
screen -S crd
git clone https://github.com/tendermint/cosmos-rosetta-gateway
cd cosmos-rosetta-gateway
make
./crd
```

Hold `ctrl` and press A, then D. You'll leave the screen session and `crd` will run in the background.

Query the blockchain `clay` using the Rosetta API provided by `crd`

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

You're now making Rosetta API queries on a freshly created Launchpad Cosmos-SDK chain generated with [Starport](https://github.com/tendermint/starport)




### Using library
* Library
The crg library runs as a part of your blockchain's daemon and directly provides a Rosetta-compatible API interface.

### Postman Collection

[Postman](https://postman.io) is a popular tool for working with and designing API systems.  We've made a [postman collection](https://www.postman.com/collections/0bb4205306d904245eee)
