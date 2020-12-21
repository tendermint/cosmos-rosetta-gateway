package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/tendermint/cosmos-rosetta-gateway/errors"
	"strings"
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
	account, err := on.offlineServicer.AccountIdentifierFromPublicKey(request.PublicKey)
	if err != nil {
		return nil, errors.ToRosetta(err)
	}
	return &types.ConstructionDeriveResponse{
		AccountIdentifier: account,
		Metadata:          nil,
	}, nil
}

func (on OnlineNetwork) ConstructionHash(ctx context.Context, request *types.ConstructionHashRequest) (*types.TransactionIdentifierResponse, *types.Error) {
	bz, err := hex.DecodeString(request.SignedTransaction)
	if err != nil {
		return nil, errors.ToRosetta(errors.WrapError(errors.ErrInvalidTransaction, "error decoding tx"))
	}

	hash := sha256.Sum256(bz)
	bzHash := hash[:]
	hashString := hex.EncodeToString(bzHash)

	return &types.TransactionIdentifierResponse{
		TransactionIdentifier: &types.TransactionIdentifier{
			Hash: strings.ToUpper(hashString),
		},
	}, nil
}

func (on OnlineNetwork) ConstructionMetadata(ctx context.Context, request *types.ConstructionMetadataRequest) (*types.ConstructionMetadataResponse, *types.Error) {
	metadata, err := on.client.ConstructionMetadataFromOptions(ctx, request.Options)
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

	res, meta, err := on.client.PostTx(txBytes)
	if err != nil {
		return nil, errors.ToRosetta(err)
	}

	return &types.TransactionIdentifierResponse{
		TransactionIdentifier: res,
		Metadata:              meta,
	}, nil
}
