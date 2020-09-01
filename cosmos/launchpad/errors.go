package launchpad

import "github.com/tendermint/cosmos-rosetta-gateway/rosetta"

var (
	ErrInterpreting     = rosetta.NewError(1, "error interpreting data from node")
	ErrNodeConnection   = rosetta.NewError(2, "error getting data from node")
	ErrTxInterpreation  = rosetta.NewError(3, "error reading tx data")
	ErrTxUnmarshal      = rosetta.NewError(4, "error unmarshalling tx data")
	ErrUnsupportedCurve = rosetta.NewError(5, "unsupported curve, expected secp256k1")
)
