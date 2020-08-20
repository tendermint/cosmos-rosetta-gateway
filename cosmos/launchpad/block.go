package launchpad

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/antihax/optional"
	"github.com/coinbase/rosetta-sdk-go/types"
	cosmosclient "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/generated"
	tendermintclient "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint/generated"
	"github.com/tendermint/cosmos-rosetta-gateway/pkg/ensurehex"
	"golang.org/x/sync/errgroup"
)

func (l Launchpad) Block(ctx context.Context, r *types.BlockRequest) (*types.BlockResponse, *types.Error) {
	var (
		blockResp tendermintclient.BlockResponse
		err       error
	)

	// retrieve the block first.
	if r.BlockIdentifier.Index != nil {
		blockResp, _, err = l.tendermint.Info.Block(ctx, &tendermintclient.BlockOpts{
			Height: optional.NewFloat32(float32(*r.BlockIdentifier.Index)),
		})
	} else {
		blockResp, _, err = l.tendermint.Info.BlockByHash(ctx, ensurehex.String(*r.BlockIdentifier.Hash))
	}
	if err != nil {
		return nil, ErrNodeConnection
	}

	// get all transactions for the block.
	var (
		txs []cosmosclient.TxQuery
		m   sync.Mutex
	)
	txsquery := fmt.Sprintf(`"tx.height=%s"`, blockResp.Result.Block.Header.Height)
	txsResp, _, err := l.tendermint.Info.TxSearch(ctx, txsquery, nil)
	if err != nil {
		return nil, ErrNodeConnection
	}
	g, ctx := errgroup.WithContext(ctx)
	for _, txshort := range txsResp.Result.Txs {
		hash := txshort.Hash
		g.Go(func() error {
			tx, _, err := l.cosmos.Transactions.TxsHashGet(ctx, hash)
			if err != nil {
				return err
			}
			m.Lock()
			defer m.Unlock()
			txs = append(txs, tx)
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return nil, ErrNodeConnection
	}

	// prepare other data.
	transactions, err := toTransactions(txs)
	if err != nil {
		return nil, ErrInterpreting
	}

	block, err := toBlockIdentifier(blockResp.Result)
	if err != nil {
		return nil, ErrInterpreting
	}

	timestamp, err := time.Parse(time.RFC3339Nano, blockResp.Result.Block.Header.Time)
	if err != nil {
		return nil, ErrInterpreting
	}

	var parentBlockId *types.BlockIdentifier
	hasParentBlock := blockResp.Result.Block.Header.LastBlockId.Hash != ""
	if hasParentBlock {
		parentBlockId = &types.BlockIdentifier{
			Index: block.Index - 1,
			Hash:  blockResp.Result.Block.Header.LastBlockId.Hash,
		}
	}

	return &types.BlockResponse{
		Block: &types.Block{
			BlockIdentifier:       block,
			Transactions:          transactions,
			Timestamp:             timestamp.UnixNano() / 1000000,
			ParentBlockIdentifier: parentBlockId,
		},
	}, nil
}

func (l Launchpad) BlockTransaction(ctx context.Context, r *types.BlockTransactionRequest) (*types.BlockTransactionResponse, *types.Error) {
	hash := r.TransactionIdentifier.Hash
	tx, _, err := l.cosmos.Transactions.TxsHashGet(ctx, hash)
	if err != nil {
		return nil, ErrNodeConnection
	}
	return &types.BlockTransactionResponse{
		Transaction: toTransaction(tx),
	}, nil
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
		transactions = append(transactions, toTransaction(tx))
	}
	return
}

func toTransaction(tx cosmosclient.TxQuery) *types.Transaction {
	return &types.Transaction{
		TransactionIdentifier: &types.TransactionIdentifier{
			Hash: tx.Txhash,
		},
		Operations: toOperations(tx.Tx.Value.Msg),
	}
}

func toOperations(msg []cosmosclient.Msg) (operations []*types.Operation) {
	for i, msg := range msg {
		account := msg.Value.Creator
		if account == "" {
			account = msg.Value.FromAddress
		}
		var amount *types.Amount
		amounts := msg.Value.Amount
		if len(amounts) > 0 {
			am := amounts[0]
			amount = &types.Amount{
				Value: am.Amount,
				Currency: &types.Currency{
					Symbol: am.Denom,
				},
			}
		}
		operations = append(operations, &types.Operation{
			OperationIdentifier: &types.OperationIdentifier{
				Index: int64(i),
			},
			Type:   msg.Type,
			Status: "TODO",
			Account: &types.AccountIdentifier{
				Address: account,
			},
			Amount: amount,
		})
	}
	return
}
