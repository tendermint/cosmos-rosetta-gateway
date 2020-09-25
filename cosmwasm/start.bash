#!/bin/bash

# turn on bash's job control
set -m

export PATH=$PATH:/wasmd/build
export CHAIN_ID="cosmwasm-coral"
export TESTNET_NAME="coralnet"
export WASMD_VERSION="v0.10.0"
export CONFIG_DIR=".corald"
export BINARY="corald"
export CLI_BINARY="coral"

export COSMJS_VERSION="v0.22.1"
export GENESIS_URL="https://raw.githubusercontent.com/CosmWasm/testnets/master/coralnet/config/genesis.json"
export APP_CONFIG_URL="https://raw.githubusercontent.com/CosmWasm/testnets/master/coralnet/config/app.toml"

export RPC="https://rpc.coralnet.cosmwasm.com:443"
export LCD="https://lcd.coralnet.cosmwasm.com"
export FAUCET="https://faucet.coralnet.cosmwasm.com"
export SEED_NODE="ec488c9215e1917e41bce5ef4b53d39ff6805166@195.201.88.9:26656"


coral config chain-id $CHAIN_ID
coral config trust-node true
coral config node $RPC
coral config output json
coral config indent true
# this is important, so the cli returns after the tx is in a block,
# and subsequent queries return the proper results
coral config broadcast-mode block

# check you can connect
coral query supply total
coral query staking validators
coral query wasm list-code

# create wallet
coral keys add mywallet

export MONIKER=new_validator
# initialize corald configuration
corald init $MONIKER

# get the testnets genesis file
curl -sSL $GENESIS_URL > ~/.corald/config/genesis.json

# get app.toml. Minimum gas prices must be 0.025ushell
curl -sSL $APP_CONFIG_URL > ~/.corald/config/app.toml

# You need to configure p2p seeds
# Either you can insert the seed addresses in $HOMEDIR/.corald/config/config.toml to "seeds"
# For simplicity we will pass the seed ID and domain as argument
# You can get the seed it using command:
## Start corald
corald start --p2p.seeds $SEED_NODE &

# Start the helper process
./cosmos-rosetta-gateway/crg -network Coral -blockchain Wasmd

# the my_helper_process might need to know how to wait on the
# primary process to start before it does its work and returns


# now we bring the primary process back into the foreground
# and leave it there
fg %1
