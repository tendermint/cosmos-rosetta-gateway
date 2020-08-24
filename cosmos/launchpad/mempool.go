package launchpad

import (
	"context"
	"encoding/base64"
	"encoding/hex"

	tmtypes "tendermint-v0.33.7/types"

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

	theTx := tendermintTxToRosettaTx(res.Result)

	return &types.MempoolTransactionResponse{
		Transaction: theTx,
	}, nil
}
