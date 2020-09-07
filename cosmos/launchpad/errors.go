package launchpad

import "github.com/tendermint/cosmos-rosetta-gateway/rosetta"

var (
	ErrInterpreting     = rosetta.NewError(1, "error interpreting data from node")
	ErrNodeConnection   = rosetta.NewError(2, "error getting data from node")
	ErrUnsupportedCurve = rosetta.NewError(3, "unsupported curve, expected secp256k1")
	ErrInvalidRequest   = rosetta.NewError(4, "invalid request")
	ErrInvalidAddress   = rosetta.NewError(6, "invalid address")
)
