package launchpad

import (
	"context"
	"encoding/base64"
	"encoding/hex"

	openapi "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint/generated"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/coinbase/rosetta-sdk-go/types"
)

func (l Launchpad) Mempool(ctx context.Context, request *types.NetworkRequest) (*types.MempoolResponse, *types.Error) {
	txs, _, err := l.tendermint.Info.UnconfirmedTxs(ctx, nil)
	if err != nil {
		return nil, ErrNodeConnection
	}

	var txsResp []*types.TransactionIdentifier
	for _, tx := range txs.Result.Txs {
		decodeString, err := base64.StdEncoding.DecodeString(tx)
		if err != nil {
			return nil, ErrInterpreting
		}

		txId := &types.TransactionIdentifier{Hash: hex.EncodeToString(tmtypes.Tx(decodeString).Hash())}
		txsResp = append(txsResp, txId)
	}

	return &types.MempoolResponse{
		TransactionIdentifiers: txsResp,
	}, nil
}

func (l Launchpad) MempoolTransaction(ctx context.Context, request *types.MempoolTransactionRequest) (*types.MempoolTransactionResponse, *types.Error) {
	res, _, err := l.tendermint.Info.Tx(ctx, request.TransactionIdentifier.Hash, nil)
	if err != nil {
		return nil, ErrNodeConnection
	}

	theTx := tendermintTransactionToRosetta(res.Result)

	return &types.MempoolTransactionResponse{
		Transaction: theTx,
	}, nil
}

func tendermintTransactionToRosetta(res openapi.TxResponseResult) *types.Transaction {
	return &types.Transaction{
		TransactionIdentifier: &types.TransactionIdentifier{
			Hash: res.Hash,
		},
		Operations: nil, // TODO difficult to get the operations from the mempool (maybe not worth it due to block times).
	}
}
