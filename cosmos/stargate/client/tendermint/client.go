package tendermint

import "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint"

func NewClient(endpoint string) *tendermint.Client {
	return tendermint.NewClient(endpoint)
}
