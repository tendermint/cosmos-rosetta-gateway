package launchpad

import (
	"context"
	"fmt"

	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/coinbase/rosetta-sdk-go/types"
)

func (l Launchpad) Mempool(ctx context.Context, request *types.NetworkRequest) (*types.MempoolResponse, *types.Error) {
	txs, _, err := l.tendermint.Info.UnconfirmedTxs(ctx, nil)
	if err != nil {
		return nil, ErrNodeConnection
	}

	for _, tx := range txs.Result.Txs {
		t := tmtypes.Tx(tx)
		fmt.Printf("%s\n", t.Hash())
	}

	return &types.MempoolResponse{}, nil
}

func (l Launchpad) MempoolTransaction(ctx context.Context, request *types.MempoolTransactionRequest) (*types.MempoolTransactionResponse, *types.Error) {
	panic("implement me")
}
