package launchpad

import (
	"github.com/irisnet/irishub/app"
	"github.com/irisnet/irishub/types"
)

const (
	StatusReverted = "Reverted"
	StatusSuccess  = "Success"

	OperationTransfer = "Transfer"

	OptionAddress = "address"
	OptionGas     = "gas"
)

// TransferTxData represents a Tx that sends value.
type TransferTxData struct {
	From   types.AccAddress
	To     types.AccAddress
	Amount types.Coin
}

var Codec = app.MakeLatestCodec()
