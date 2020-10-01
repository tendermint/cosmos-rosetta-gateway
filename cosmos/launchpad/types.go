package launchpad

import (
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/types"
)

const (
	StatusReverted = "Reverted"
	StatusSuccess  = "Success"

	OperationTransfer = "Transfer"
	OperationDelegate = "Delegate"

	OptionAddress = "address"
	OptionGas     = "gas"
)

// TransferTxData represents a Tx that sends value.
type TransferTxData struct {
	From   types.AccAddress
	To     types.AccAddress
	Amount types.Coin
}

var supportedOps = []string{OperationTransfer, OperationDelegate}

var Codec = simapp.MakeCodec()
