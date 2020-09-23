package launchpad

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint"

	"github.com/coinbase/rosetta-sdk-go/types"
	"golang.org/x/sync/errgroup"

	cosmosclient "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/generated"
)

func (l Launchpad) Block(ctx context.Context, r *types.BlockRequest) (*types.BlockResponse, *types.Error) {
	if l.properties.OfflineMode {
		return nil, ErrEndpointDisabledOfflineMode
	}

	var (
		blockResp tendermint.BlockResponse
		err       error
	)

	// retrieve the block first.
	if r.BlockIdentifier.Index != nil {
		blockResp, err = l.tendermint.Block(uint64(*r.BlockIdentifier.Index))
	} else {
		blockResp, err = l.tendermint.BlockByHash(HexPrefix(*r.BlockIdentifier.Hash))
	}
	if err != nil {
		return nil, ErrNodeConnection
	}

	// get all transactions for the block.
	var (
		txs []cosmosclient.TxQuery
		m   sync.Mutex
	)
	txsquery := fmt.Sprintf(`tx.height=%s`, blockResp.Block.Header.Height)
	txsResp, err := l.tendermint.TxSearch(txsquery)
	if err != nil {
		return nil, ErrNodeConnection
	}
	g, ctx := errgroup.WithContext(ctx)
	for _, txshort := range txsResp.Txs {
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

	block, err := toBlockIdentifier(blockResp)
	if err != nil {
		return nil, ErrInterpreting
	}

	timestamp, err := time.Parse(time.RFC3339Nano, blockResp.Block.Header.Time)
	if err != nil {
		return nil, ErrInterpreting
	}

	parentBlockId := block // If it does not have parent block it is the same as block 1.
	hasParentBlock := blockResp.Block.Header.LastBlockId.Hash != ""
	if hasParentBlock {
		parentBlockId = &types.BlockIdentifier{
			Index: block.Index - 1,
			Hash:  blockResp.Block.Header.LastBlockId.Hash,
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
	if l.properties.OfflineMode {
		return nil, ErrEndpointDisabledOfflineMode
	}

	tx, err := l.getTxByHash(ctx, r.TransactionIdentifier.Hash)
	if err != nil {
		return nil, err
	}
	return &types.BlockTransactionResponse{
		Transaction: tx,
	}, nil
}
