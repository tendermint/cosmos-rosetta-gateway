package launchpad

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	cosmostypes "github.com/cosmos/cosmos-sdk/types"

	cosmosclient "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/generated"
	tendermintclient "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint/generated"

	"github.com/coinbase/rosetta-sdk-go/types"
)

const (
	typeMsgSend = "cosmos-sdk/MsgSend"
	zerox       = "0x"
)

// HexPrefix ensures that string representation of hex starts with 0x.
func HexPrefix(hex string) string {
	if !strings.HasPrefix(hex, zerox) {
		return zerox + hex
	}
	return hex
}

// getTxByHash calls
func (l Launchpad) getTxByHash(ctx context.Context, hash string) (*types.Transaction, *types.Error) {
	txQuery, _, err := l.cosmos.Transactions.TxsHashGet(ctx, hash)
	if err != nil {
		return nil, ErrNodeConnection
	}

	tx := cosmosTxToRosettaTx(txQuery)

	return tx, nil
}

func toBlockIdentifier(result tendermintclient.BlockComplete) (*types.BlockIdentifier, error) {
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

func toTransactions(txs []cosmosclient.TxQuery) (transactions []*types.Transaction, err error) {
	for _, tx := range txs {
		transactions = append(transactions, cosmosTxToRosettaTx(tx))
	}
	return
}

// tendermintTxToRosettaTx converts a Tendermint api TxResponseResult to a Transaction
// in the type expected by Rosetta.
func tendermintTxToRosettaTx(res tendermintclient.TxResponseResult) *types.Transaction {
	return &types.Transaction{
		TransactionIdentifier: &types.TransactionIdentifier{
			Hash: res.Hash,
		},
		Operations: nil, // TODO difficult to get the operations from the mempool (maybe not worth it due to block times).
	}
}

// cosmosTxToRosettaTx converts a Cosmos api TxQuery to a Transaction
// in the type expected by Rosetta.
func cosmosTxToRosettaTx(tx cosmosclient.TxQuery) *types.Transaction {
	hasError := tx.Code > 0
	return &types.Transaction{
		TransactionIdentifier: &types.TransactionIdentifier{
			Hash: tx.Txhash,
		},
		Operations: toOperations(tx.Tx.Value.Msg, hasError),
	}
}

func toOperations(msg []cosmosclient.Msg, hasError bool) (operations []*types.Operation) {
	for i, msg := range msg {
		if msg.Type != typeMsgSend {
			continue
		}
		fromAddress := msg.Value.FromAddress
		toAddress := msg.Value.ToAddress
		amounts := msg.Value.Amount
		if len(amounts) == 0 {
			continue
		}
		coin := amounts[0]
		sendOp := func(account, amount string, index int) *types.Operation {
			status := StatusSuccess
			if hasError {
				status = StatusReverted
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
			sendOp(fromAddress, "-"+coin.Amount, i),
			sendOp(toAddress, coin.Amount, i+1),
		)
	}
	return
}

// getTransferTxDataFromOperations extracts the from and to addresses from a list of operations.
// We assume that it comes formated in the correct way. And that the balance of the sender is the same
// as the receiver operations.
func getTransferTxDataFromOperations(ops []*types.Operation) (*TransferTxData, error) {
	var (
		transferData = &TransferTxData{}
		err          error
	)

	for _, op := range ops {
		if strings.HasPrefix(op.Amount.Value, "-") {
			transferData.From, err = cosmostypes.AccAddressFromBech32(op.Account.Address)
			if err != nil {
				return nil, err
			}
		} else {
			transferData.To, err = cosmostypes.AccAddressFromBech32(op.Account.Address)
			if err != nil {
				return nil, err
			}

			amount, err := strconv.ParseInt(op.Amount.Value, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid amount")
			}

			transferData.Amount = cosmostypes.NewCoin(op.Amount.Currency.Symbol, cosmostypes.NewInt(amount))
		}
	}

	return transferData, nil
}
