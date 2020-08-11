package stargate

import "github.com/tendermint/cosmos-rosetta-gateway/rosetta"

type Stargate struct {
}

func NewStargate() rosetta.Adapter {
	return &Stargate{}
}
