package launchpad

import (
	"github.com/cosmos/cosmos-sdk/types"
)

const (
	StatusReverted = "Reverted"
	StatusSuccess  = "Success"

	OperationTransfer = "Transfer"

	OptionAddress = "address"
	OptionGas     = "gas"

	OptionsAccountNumber = "account_number"
	OptionsSequence      = "sequence"
	OptionsChainId       = "chain_id"
)

// TransferTxData represents a Tx that sends value.
type TransferTxData struct {
	From   types.AccAddress
	To     types.AccAddress
	Amount types.Coin
}

