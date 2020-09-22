package launchpad

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"sync"
	"time"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/alttendermint"

	"github.com/coinbase/rosetta-sdk-go/types"
	"golang.org/x/sync/errgroup"
)

func (l Launchpad) Block(ctx context.Context, r *types.BlockRequest) (*types.BlockResponse, *types.Error) {
	var (
		blockResp alttendermint.BlockResponse
		err       error
	)

	// retrieve the block first.
	if r.BlockIdentifier.Index != nil {
		blockResp, err = l.altTendermint.Block(uint64(*r.BlockIdentifier.Index))
	} else {
		blockResp, err = l.altTendermint.BlockByHash(HexPrefix(*r.BlockIdentifier.Hash))
	}
	if err != nil {
		return nil, ErrNodeConnection
	}

	// get all transactions for the block.
	var (
		txs []sdk.TxResponse
		m   sync.Mutex
	)
	txsquery := fmt.Sprintf(`"tx.height=%s"`, blockResp.Block.Header.Height)
	txsResp, _, err := l.tendermint.Info.TxSearch(ctx, txsquery, nil)
	if err != nil {
		return nil, ErrNodeConnection
	}
	g, ctx := errgroup.WithContext(ctx)
	for _, txshort := range txsResp.Result.Txs {
		hash := txshort.Hash
		g.Go(func() error {
			tx, err := l.altCosmos.GetTx(ctx, hash)
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
	tx, err := l.getTxByHash(ctx, r.TransactionIdentifier.Hash)
	if err != nil {
		return nil, err
	}
	return &types.BlockTransactionResponse{
		Transaction: tx,
	}, nil
}
