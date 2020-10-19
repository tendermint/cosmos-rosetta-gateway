package launchpad

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/x/bank"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/coinbase/rosetta-sdk-go/types"
)

const (
	zerox = "0x"
)

// HexPrefix ensures that string representation of hex starts with 0x.
func HexPrefix(hex string) string {
	if !strings.HasPrefix(hex, zerox) {
		return zerox + hex
	}
	return hex
}

// getTxByHash calls
func (l launchpad) getTxByHash(ctx context.Context, hash string) (*types.Transaction, *types.Error) {
	txQuery, err := l.cosmos.GetTx(ctx, hash)
	if err != nil {
		return nil, ErrNodeConnection
	}
	if txQuery.Tx == nil {
		return nil, ErrInvalidTxHash
	}
	tx := cosmosTxToRosettaTx(txQuery)

	return tx, nil
}

func toBlockIdentifier(result tendermint.BlockResponse) (*types.BlockIdentifier, error) {
	if result.BlockId.Hash == "" {
		return nil, nil
	}
	height, err := strconv.ParseUint(result.Block.Header.Height, 10, 64)
	if err != nil {
		return nil, err
	}
	return &types.BlockIdentifier{
		Index: int64(height),
		Hash:  result.BlockId.Hash,
	}, nil
}

func toTransactions(txs []sdk.TxResponse) (transactions []*types.Transaction, err error) {
	for _, tx := range txs {
		transactions = append(transactions, cosmosTxToRosettaTx(tx))
	}
	return
}

// tendermintTxToRosettaTx converts a Tendermint api TxResponseResult to a Transaction
// in the type expected by Rosetta.
func tendermintTxToRosettaTx(res tendermint.TxResponse) *types.Transaction {
	return &types.Transaction{
		TransactionIdentifier: &types.TransactionIdentifier{
			Hash: res.Hash,
		},
		Operations: nil, // TODO difficult to get the operations from the mempool (maybe not worth it due to block times).
	}
}

// cosmosTxToRosettaTx converts a Cosmos api TxQuery to a Transaction
// in the type expected by Rosetta.
func cosmosTxToRosettaTx(tx sdk.TxResponse) *types.Transaction {
	hasError := tx.Code > 0
	return &types.Transaction{
		TransactionIdentifier: &types.TransactionIdentifier{
			Hash: tx.TxHash,
		},
		Operations: toOperations(tx.Tx.GetMsgs(), hasError, false),
	}
}

func toOperations(msg []sdk.Msg, hasError bool, withoutStatus bool) (operations []*types.Operation) {
	for i, msg := range msg {
		switch msg.Type() {
		case "MsgSend":
			newMsg := msg.(bank.MsgSend)
			fromAddress := newMsg.FromAddress
			toAddress := newMsg.ToAddress
			amounts := newMsg.Amount
			if len(amounts) == 0 {
				continue
			}
			coin := amounts[0]
			sendOp := func(account, amount string, index int) *types.Operation {
				status := StatusSuccess
				if hasError {
					status = StatusReverted
				}
				if withoutStatus {
					status = ""
				}
				return &types.Operation{
					OperationIdentifier: &types.OperationIdentifier{
						Index: int64(index),
					},
					Type:   OperationTransfer,
					Status: status,
					Account: &types.AccountIdentifier{
						Address: account,
					},
					Amount: &types.Amount{
						Value: amount,
						Currency: &types.Currency{
							Symbol: coin.Denom,
						},
					},
				}
			}
			operations = append(operations,
				sendOp(fromAddress.String(), "-"+coin.Amount.String(), i),
				sendOp(toAddress.String(), coin.Amount.String(), i+1),
			)

		case "MsgDelegate":
			newMsg := msg.(staking.MsgDelegate)
			srcAddr := newMsg.DelegatorAddress
			valAddr := newMsg.ValidatorAddress
			delAmt := newMsg.Amount
			sendOp := func(account, amount string, index int) *types.Operation {
				status := StatusSuccess
				if hasError {
					status = StatusReverted
				}
				return &types.Operation{
					OperationIdentifier: &types.OperationIdentifier{
						Index: int64(index),
					},
					Type:   OperationDelegate,
					Status: status,
					Account: &types.AccountIdentifier{
						Address: account,
					},
					Amount: &types.Amount{
						Value: amount,
						Currency: &types.Currency{
							Symbol: delAmt.Denom,
						},
					},
				}
			}
			operations = append(operations,
				sendOp(srcAddr.String(), "-"+delAmt.Amount.String(), i),
				sendOp(valAddr.String(), delAmt.Amount.String(), i+1),
			)
		}
	}
	return
}

func getMsgDataFromOperations(ops []*types.Operation) (sdk.Msg, error) {
	op := ops[0]
	switch op.Type {
	case OperationTransfer:
		return getTransferTxDataFromOperations(ops)
	case OperationDelegate:
		return getDelegateDataFromOperations(ops)
	}

	return nil, fmt.Errorf("unable to iterate operations")
}

// getTransferTxDataFromOperations extracts the from and to addresses from a list of operations.
// We assume that it comes formated in the correct way. And that the balance of the sender is the same
// as the receiver operations.
func getTransferTxDataFromOperations(ops []*types.Operation) (sdk.Msg, error) {
	var (
		from, to sdk.AccAddress
		sendAmt  sdk.Coin
		err      error
	)

	for _, op := range ops {
		if strings.HasPrefix(op.Amount.Value, "-") {
			from, err = sdk.AccAddressFromBech32(op.Account.Address)
			if err != nil {
				return nil, err
			}
		} else {
			to, err = sdk.AccAddressFromBech32(op.Account.Address)
			if err != nil {
				return nil, err
			}

			amount, err := strconv.ParseInt(op.Amount.Value, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid amount")
			}

			sendAmt = sdk.NewCoin(op.Amount.Currency.Symbol, sdk.NewInt(amount))
		}
	}

	msg := bank.NewMsgSend(from, to, sdk.NewCoins(sendAmt))
	return msg, nil
}

func getDelegateDataFromOperations(ops []*types.Operation) (sdk.Msg, error) {
	var (
		delAddr sdk.AccAddress
		valAddr sdk.ValAddress
		delAmt  sdk.Coin
		err     error
	)

	for _, op := range ops {
		if strings.HasPrefix(op.Amount.Value, "-") {
			delAddr, err = sdk.AccAddressFromBech32(op.Account.Address)
			if err != nil {
				return nil, err
			}
		} else {
			valAddr, err = sdk.ValAddressFromBech32(op.Account.Address)
			if err != nil {
				return nil, err
			}

			amount, err := strconv.ParseInt(op.Amount.Value, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid amount")
			}

			delAmt = sdk.NewCoin(op.Amount.Currency.Symbol, sdk.NewInt(amount))
		}
	}

	msg := staking.NewMsgDelegate(delAddr, valAddr, delAmt)
	return msg, nil
}
