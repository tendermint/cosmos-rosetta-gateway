FROM archlinux

# CORAL TESTNET RUNTIME ENV

ENV CHAIN_ID cosmwasm-coral
ENV TESTNET_NAME coralnet
ENV WASMD_VERSION v0.10.0
ENV CONFIG_DIR .corald
ENV BINARY corald
ENV CLI_BINARY coral
ENV COSMJS_VERSION v0.22.1
ENV GENESIS_URL https://raw.githubusercontent.com/CosmWasm/testnets/master/coralnet/config/genesis.json
ENV APP_CONFIG_URL https://raw.githubusercontent.com/CosmWasm/testnets/master/coralnet/config/app.toml
ENV RPC https://rpc.coralnet.cosmwasm.com:443
ENV LCD https://lcd.coralnet.cosmwasm.com
ENV FAUCET https://faucet.coralnet.cosmwasm.com
ENV SEED_NODE ec488c9215e1917e41bce5ef4b53d39ff6805166@195.201.88.9:26656

# GET BUILD DEPENDENCIES
RUN pacman -Syyu --noconfirm go base-devel rust git curl wget

# DOWNLOAD AND COMPILE WASMD
RUN git clone https://github.com/CosmWasm/wasmd && \
      git checkout v0.10.0 && \
      cd wasmd && \
      make build-coral 
      
# DOWNLOAD AND COMPILE COSMOS ROSETTA GATEWAY
RUN git clone https://github.com/tendermint/cosmos-rosetta-gateway && \
      cd cosmos-rosetta-gateway && \
      make
      

      
