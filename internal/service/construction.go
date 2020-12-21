package service

import (
	"context"
	"encoding/hex"
	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/tendermint/cosmos-rosetta-gateway/errors"
)

func (on OnlineNetwork) ConstructionCombine(ctx context.Context, request *types.ConstructionCombineRequest) (*types.ConstructionCombineResponse, *types.Error) {
	txBytes, err := hex.DecodeString(request.UnsignedTransaction)
	if err != nil {
		return nil, errors.ToRosetta(err)
	}

	signedTx, err := on.offlineServicer.SignedTx(ctx, txBytes, request.Signatures)
	if err != nil {
		return nil, errors.ToRosetta(err)
	}

	return &types.ConstructionCombineResponse{
		SignedTransaction: hex.EncodeToString(signedTx),
	}, nil
}

func (on OnlineNetwork) ConstructionDerive(ctx context.Context, request *types.ConstructionDeriveRequest) (*types.ConstructionDeriveResponse, *types.Error) {
	panic("implemenet")
}

func (on OnlineNetwork) ConstructionHash(ctx context.Context, request *types.ConstructionHashRequest) (*types.TransactionIdentifierResponse, *types.Error) {
	panic("implemenet")
}

func (on OnlineNetwork) ConstructionMetadata(ctx context.Context, request *types.ConstructionMetadataRequest) (*types.ConstructionMetadataResponse, *types.Error) {
	metadata, err := on.onlineServicer.ConstructionMetadataFromOptions(ctx, request.Options)
	if err != nil {
		return nil, errors.ToRosetta(err)
	}

	return &types.ConstructionMetadataResponse{
		Metadata: metadata,
	}, nil
}

func (on OnlineNetwork) ConstructionParse(ctx context.Context, request *types.ConstructionParseRequest) (*types.ConstructionParseResponse, *types.Error) {
	txBytes, err := hex.DecodeString(request.Transaction)
	if err != nil {
		err := errors.WrapError(errors.ErrInvalidTransaction, err.Error())
		return nil, errors.ToRosetta(err)
	}
	ops, signers, err := on.offlineServicer.TxOperationsAndSignersAccountIdentifiers(request.Signed, txBytes)
	if err != nil {
		return nil, errors.ToRosetta(err)
	}
	return &types.ConstructionParseResponse{
		Operations:               ops,
		AccountIdentifierSigners: signers,
		Metadata:                 nil,
	}, nil

}

func (on OnlineNetwork) ConstructionPayloads(ctx context.Context, request *types.ConstructionPayloadsRequest) (*types.ConstructionPayloadsResponse, *types.Error) {
	payload, err := on.offlineServicer.ConstructionPayload(ctx, request)
	if err != nil {
		return nil, errors.ToRosetta(err)
	}
	return payload, nil
}

func (on OnlineNetwork) ConstructionPreprocess(ctx context.Context, request *types.ConstructionPreprocessRequest) (*types.ConstructionPreprocessResponse, *types.Error) {
	options, err := on.offlineServicer.PreprocessOperationsToOptions(ctx, request)
	if err != nil {
		return nil, errors.ToRosetta(err)
	}

	return &types.ConstructionPreprocessResponse{
		Options: options,
	}, nil
}

func (on OnlineNetwork) ConstructionSubmit(ctx context.Context, request *types.ConstructionSubmitRequest) (*types.TransactionIdentifierResponse, *types.Error) {
	txBytes, err := hex.DecodeString(request.SignedTransaction)
	if err != nil {
		return nil, errors.ToRosetta(err)
	}

	res, meta, err := on.onlineServicer.PostTx(txBytes)
	if err != nil {
		return nil, errors.ToRosetta(err)
	}

	return &types.TransactionIdentifierResponse{
		TransactionIdentifier: res,
		Metadata:              meta,
	}, nil
}
