package launchpad

import (
	"github.com/cosmos/cosmos-sdk/simapp"
)

const (
	StatusReverted = "Reverted"
	StatusSuccess  = "Success"

	OperationTransfer = "Transfer"
	OperationDelegate = "Delegate"

	OptionAddress = "address"
	OptionGas     = "gas"
)

var supportedOps = []string{OperationTransfer, OperationDelegate}

var Codec = simapp.MakeCodec()
