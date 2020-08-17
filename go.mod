module github.com/tendermint/cosmos-rosetta-gateway

go 1.14

require (
	github.com/coinbase/rosetta-sdk-go v0.3.4
	github.com/cosmos/cosmos-sdk v0.39.1
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.0.0
	github.com/stretchr/testify v1.6.1
	github.com/tendermint/cosmos-rosetta-gateway/rosetta/cosmos/launchpad/client/cosmos/generated v0.0.0-00010101000000-000000000000
	github.com/vektra/mockery/v2 v2.2.1
	golang.org/x/sync v0.0.0-20200317015054-43a5402ce75a
)

replace github.com/tendermint/cosmos-rosetta-gateway/rosetta/cosmos/launchpad/client/cosmos/generated => ./rosetta/cosmos/launchpad/client/cosmos/generated
