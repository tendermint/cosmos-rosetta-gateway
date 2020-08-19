package launchpad

import (
	"context"
	"encoding/base64"
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
		fmt.Println(3, err)
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
				fmt.Println(5, err)
				return err
			}
			m.Lock()
			defer m.Unlock()
			txs = append(txs, tx)
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		fmt.Println(1, err)
		return nil, ErrNodeConnection
	}

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
		fmt.Println(2, err)
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

	resp := &types.BlockResponse{
		Block: &types.Block{
			BlockIdentifier:       block,
			Transactions:          transactions,
			Timestamp:             timestamp.UnixNano() / 1000000,
			ParentBlockIdentifier: parentBlockId,
		},
	}

	return resp, nil
}

func (l Launchpad) BlockTransaction(ctx context.Context, r *types.BlockTransactionRequest) (*types.BlockTransactionResponse, *types.Error) {
	panic("unimplemented")
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
		var operations []*types.Operation
		for i, msg := range tx.Tx.Value.Msg {
			account := msg.Value.Creator
			if account == "" {
				account = msg.Value.FromAddress
			}
			operation := &types.Operation{
				OperationIdentifier: &types.OperationIdentifier{
					Index: int64(i),
				},
				Type:   msg.Type,
				Status: "TODO",
				Account: &types.AccountIdentifier{
					Address: account,
				},
			}
			amounts := msg.Value.Amount
			if len(amounts) > 0 {
				am := amounts[0]
				operation.Amount = &types.Amount{
					Value: am.Amount,
					Currency: &types.Currency{
						Symbol: am.Denom,
					},
				}
			}
			operations = append(operations, operation)
		}
		transactions = append(transactions, &types.Transaction{
			TransactionIdentifier: &types.TransactionIdentifier{
				Hash: tx.Txhash,
			},
			Operations: operations,
		})
	}
	return
}

func decodeAttr(attr tendermintclient.TxSearchResponseResultTxResultAttributes) (key, value string, err error) {
	keyb, err := base64.StdEncoding.DecodeString(attr.Key)
	if err != nil {
		return "", "", err
	}
	valueb, err := base64.StdEncoding.DecodeString(attr.Value)
	if err != nil {
		return "", "", err
	}
	key = string(keyb)
	value = string(valueb)
	return
}
