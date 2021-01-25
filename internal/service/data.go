package service

import (
	"context"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/tendermint/cosmos-rosetta-gateway/errors"
	crgtypes "github.com/tendermint/cosmos-rosetta-gateway/types"
)

// AccountBalance retrieves the account balance of an address
// rosetta requires us to fetch the block information too
func (on OnlineNetwork) AccountBalance(ctx context.Context, request *types.AccountBalanceRequest) (*types.AccountBalanceResponse, *types.Error) {
	var (
		height int64
		block  crgtypes.BlockResponse
		err    error
	)

	switch {
	case request.BlockIdentifier == nil:
		block, err = on.client.BlockByHeight(ctx, nil)
		if err != nil {
			return nil, errors.ToRosetta(err)
		}
	case request.BlockIdentifier.Hash != nil:
		block, err = on.client.BlockByHash(ctx, *request.BlockIdentifier.Hash)
		if err != nil {
			return nil, errors.ToRosetta(err)
		}
		height = block.Block.Index
	case request.BlockIdentifier.Index != nil:
		height = *request.BlockIdentifier.Index
		block, err = on.client.BlockByHeight(ctx, &height)
		if err != nil {
			return nil, errors.ToRosetta(err)
		}
	}

	accountCoins, err := on.client.Balances(ctx, request.AccountIdentifier.Address, &height)
	if err != nil {
		return nil, errors.ToRosetta(err)
	}

	return &types.AccountBalanceResponse{
		BlockIdentifier: (*types.BlockIdentifier)(block.Block),
		Balances:        convertBalances(accountCoins...),
		Coins:           nil,
		Metadata:        nil,
	}, nil
}

func convertBalances(crgAmts ...*crgtypes.Amount) []*types.Amount {
	rosAmts := make([]*types.Amount, len(crgAmts))

	for i, crgAmt := range crgAmts {
		rosAmts[i] = &types.Amount{
			Value:    crgAmt.Value,
			Currency: (*types.Currency)(crgAmt.Currency),
			Metadata: crgAmt.Metadata,
		}
	}

	return rosAmts
}

// Block gets the transactions in the given block
func (on OnlineNetwork) Block(ctx context.Context, request *types.BlockRequest) (*types.BlockResponse, *types.Error) {
	var (
		blockResponse crgtypes.BlockTransactionsResponse
		err           error
	)
	// block identifier is assumed not to be nil as rosetta will do this check for us
	// check if we have to query via hash or block number
	switch {
	case request.BlockIdentifier.Hash != nil:
		blockResponse, err = on.client.BlockTransactionsByHash(ctx, *request.BlockIdentifier.Hash)
		if err != nil {
			return nil, errors.ToRosetta(err)
		}
	case request.BlockIdentifier.Index != nil:
		blockResponse, err = on.client.BlockTransactionsByHeight(ctx, request.BlockIdentifier.Index)
		if err != nil {
			return nil, errors.ToRosetta(err)
		}
	default:
		err := errors.WrapError(errors.ErrBadArgument, "at least one of hash or index needs to be specified")
		return nil, errors.ToRosetta(err)
	}

	return &types.BlockResponse{
		Block: &types.Block{
			BlockIdentifier:       (*types.BlockIdentifier)(blockResponse.Block),
			ParentBlockIdentifier: (*types.BlockIdentifier)(blockResponse.ParentBlock),
			Timestamp:             blockResponse.MillisecondTimestamp,
			Transactions:          convertTransactions(blockResponse.Transactions...),
			Metadata:              nil,
		},
		OtherTransactions: nil,
	}, nil
}

func convertTransactions(crgTxs ...*crgtypes.Transaction) []*types.Transaction {
	rosTxs := make([]*types.Transaction, len(crgTxs))

	for i, crgTx := range crgTxs {
		rosTxs[i] = &types.Transaction{
			TransactionIdentifier: (*types.TransactionIdentifier)(crgTx.TransactionIdentifier),
			Operations:            convertOperations(crgTx.Operations),
			Metadata:              crgTx.Metadata,
		}
	}

	return rosTxs
}

func convertOperations(crgOps []*crgtypes.Operation) []*types.Operation {
	rosOps := make([]*types.Operation, len(crgOps))

	for i, crgOp := range crgOps {
		rosOps[i] = &types.Operation{
			OperationIdentifier: (*types.OperationIdentifier)(crgOp.OperationIdentifier),
			RelatedOperations:   convertOperationIdentifiers(crgOp.RelatedOperations...),
			Type:                crgOp.Type,
			Status:              crgOp.Status,
			Account:             convertAccountIdentifiers(crgOp.Account)[0],
			Amount:              convertBalances(crgOp.Amount)[0],
			Metadata:            crgOp.Metadata,
		}
	}

	return rosOps
}

func convertAccountIdentifiers(crgAccs ...*crgtypes.AccountIdentifier) []*types.AccountIdentifier {
	rosAccs := make([]*types.AccountIdentifier, len(crgAccs))

	for i, crgAcc := range crgAccs {
		rosAccs[i] = &types.AccountIdentifier{
			Address:    crgAcc.Address,
			SubAccount: (*types.SubAccountIdentifier)(crgAcc.SubAccount),
			Metadata:   crgAcc.Metadata,
		}
	}

	return rosAccs
}

func convertOperationIdentifiers(crgOps ...*crgtypes.OperationIdentifier) []*types.OperationIdentifier {
	rosOps := make([]*types.OperationIdentifier, len(crgOps))

	for i, crgOp := range crgOps {
		rosOps[i] = (*types.OperationIdentifier)(crgOp)
	}

	return rosOps
}

// BlockTransaction gets the given transaction in the specified block, we do not need to check the block itself too
// due to the fact that tendermint achieves instant finality
func (on OnlineNetwork) BlockTransaction(ctx context.Context, request *types.BlockTransactionRequest) (*types.BlockTransactionResponse, *types.Error) {
	tx, err := on.client.GetTx(ctx, request.TransactionIdentifier.Hash)
	if err != nil {
		return nil, errors.ToRosetta(err)
	}

	return &types.BlockTransactionResponse{
		Transaction: convertTransactions(tx)[0],
	}, nil
}

// Mempool fetches the transactions contained in the mempool
func (on OnlineNetwork) Mempool(ctx context.Context, _ *types.NetworkRequest) (*types.MempoolResponse, *types.Error) {
	txs, err := on.client.Mempool(ctx)
	if err != nil {
		return nil, errors.ToRosetta(err)
	}

	return &types.MempoolResponse{
		TransactionIdentifiers: convertTransactionIdentifiers(txs),
	}, nil
}

func convertTransactionIdentifiers(crgTxs []*crgtypes.TransactionIdentifier) []*types.TransactionIdentifier {
	rosTxs := make([]*types.TransactionIdentifier, len(crgTxs))

	for i, crgTx := range crgTxs {
		rosTxs[i] = (*types.TransactionIdentifier)(crgTx)
	}

	return rosTxs
}

// MempoolTransaction fetches a single transaction in the mempool
// NOTE: it is not implemented yet
func (on OnlineNetwork) MempoolTransaction(ctx context.Context, request *types.MempoolTransactionRequest) (*types.MempoolTransactionResponse, *types.Error) {
	tx, err := on.client.GetUnconfirmedTx(ctx, request.TransactionIdentifier.Hash)
	if err != nil {
		return nil, errors.ToRosetta(err)
	}

	return &types.MempoolTransactionResponse{
		Transaction: convertTransactions(tx)[0],
	}, nil
}

func (on OnlineNetwork) NetworkList(_ context.Context, _ *types.MetadataRequest) (*types.NetworkListResponse, *types.Error) {
	return &types.NetworkListResponse{NetworkIdentifiers: []*types.NetworkIdentifier{on.network}}, nil
}

func (on OnlineNetwork) NetworkOptions(_ context.Context, _ *types.NetworkRequest) (*types.NetworkOptionsResponse, *types.Error) {
	return on.networkOptions, nil
}

func (on OnlineNetwork) NetworkStatus(ctx context.Context, _ *types.NetworkRequest) (*types.NetworkStatusResponse, *types.Error) {
	block, err := on.client.BlockByHeight(ctx, nil)
	if err != nil {
		return nil, errors.ToRosetta(err)
	}

	peers, err := on.client.Peers(ctx)
	if err != nil {
		return nil, errors.ToRosetta(err)
	}

	syncStatus, err := on.client.Status(ctx)
	if err != nil {
		return nil, errors.ToRosetta(err)
	}

	return &types.NetworkStatusResponse{
		CurrentBlockIdentifier: (*types.BlockIdentifier)(block.Block),
		CurrentBlockTimestamp:  block.MillisecondTimestamp,
		GenesisBlockIdentifier: on.genesisBlockIdentifier,
		OldestBlockIdentifier:  nil,
		SyncStatus:             (*types.SyncStatus)(syncStatus),
		Peers:                  convertPeers(peers),
	}, nil
}

func convertPeers(crg []*crgtypes.Peer) []*types.Peer {
	rosPeers := make([]*types.Peer, len(crg))

	for i, crgPeer := range crg {
		rosPeers[i] = (*types.Peer)(crgPeer)
	}

	return rosPeers
}
