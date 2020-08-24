module github.com/tendermint/cosmos-rosetta-gateway

go 1.14

require (
	github.com/coinbase/rosetta-sdk-go v0.3.4
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.0.0
	github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad v0.0.0-00010101000000-000000000000
	github.com/vektra/mockery/v2 v2.2.1
)

replace github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad => ./cosmos/launchpad
