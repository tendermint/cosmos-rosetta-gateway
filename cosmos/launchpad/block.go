package launchpad

import (
	"context"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/antihax/optional"
	"github.com/coinbase/rosetta-sdk-go/types"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	tendermintclient "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint/generated"
	"github.com/tendermint/cosmos-rosetta-gateway/pkg/ensurehex"
)

func (l Launchpad) Block(ctx context.Context, r *types.BlockRequest) (*types.BlockResponse, *types.Error) {
	var (
		blockResp       tendermintclient.BlockResponse
		parentBlockResp tendermintclient.BlockResponse
		err             error
	)
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

	timestamp, err := time.Parse(time.RFC3339Nano, blockResp.Result.Block.Header.Time)
	if err != nil {
		return nil, ErrInterpreting
	}

	txsquery := fmt.Sprintf(`"tx.height=%s"`, blockResp.Result.Block.Header.Height)
	txsResp, _, err := l.tendermint.Info.TxSearch(ctx, txsquery, nil)
	if err != nil {
		return nil, ErrNodeConnection
	}
	transactions, err := toTransactions(txsResp.Result)
	if err != nil {
		return nil, ErrInterpreting
	}

	block, err := toBlockIdentifier(blockResp.Result)
	if err != nil {
		return nil, ErrInterpreting
	}

	resp := &types.BlockResponse{
		Block: &types.Block{
			BlockIdentifier: block,
			Transactions:    transactions,
			Timestamp:       timestamp.UnixNano() / 1000000,
		},
	}
	hasParentBlock := blockResp.Result.Block.Header.LastBlockId.Hash != ""
	if hasParentBlock {
		parentBlockResp, _, err = l.tendermint.Info.BlockByHash(ctx, ensurehex.String(blockResp.Result.Block.Header.LastBlockId.Hash))
		if err != nil {
			return nil, ErrNodeConnection
		}
		parentBlock, err := toBlockIdentifier(parentBlockResp.Result)
		if err != nil {
			return nil, ErrInterpreting
		}
		resp.Block.ParentBlockIdentifier = parentBlock
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

func toTransactions(result tendermintclient.TxSearchResponseResult) (transactions []*types.Transaction, err error) {
	for _, tx := range result.Txs {
		var operations []*types.Operation
		for _, event := range tx.TxResult.Events {
			switch event.Type {
			case "transfer":
				var (
					recipient string
					sender    string
					amount    string
				)
				for _, attr := range event.Attributes {
					key, value, err := decodeAttr(attr)
					if err != nil {
						return nil, err
					}
					switch key {
					case "recipient":
						recipient = value
					case "sender":
						sender = value
					case "amount":
						amount = value
					}
				}
				if amount != "" {
					coin, err := cosmostypes.ParseCoin(amount)
					if err != nil {
						return nil, err
					}
					amt := &types.Amount{
						Value: fmt.Sprintf("%d", coin.Amount.Int64()),
						Currency: &types.Currency{
							Symbol: coin.Denom,
						},
					}
					opType := strings.Title(event.Type)
					operations = append(operations,
						&types.Operation{
							OperationIdentifier: &types.OperationIdentifier{},
							Type:                opType,
							Status:              "Sent",
							Amount:              amt,
							Account: &types.AccountIdentifier{
								Address: sender,
							},
						},
						&types.Operation{
							OperationIdentifier: &types.OperationIdentifier{
								Index: 1,
							},
							Type:   opType,
							Status: "Received",
							Amount: amt,
							Account: &types.AccountIdentifier{
								Address: recipient,
							},
						},
					)
				}
			}
		}

		transactions = append(transactions, &types.Transaction{
			TransactionIdentifier: &types.TransactionIdentifier{
				Hash: tx.Hash,
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
