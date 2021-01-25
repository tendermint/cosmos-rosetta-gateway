package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	crgtypes "github.com/tendermint/cosmos-rosetta-gateway/types"
	"strings"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/tendermint/cosmos-rosetta-gateway/errors"
)

func (on OnlineNetwork) ConstructionCombine(ctx context.Context, request *types.ConstructionCombineRequest) (*types.ConstructionCombineResponse, *types.Error) {
	txBytes, err := hex.DecodeString(request.UnsignedTransaction)
	if err != nil {
		return nil, errors.ToRosetta(err)
	}

	signedTx, err := on.client.SignedTx(ctx, txBytes, toSignatures(request.Signatures))
	if err != nil {
		return nil, errors.ToRosetta(err)
	}

	return &types.ConstructionCombineResponse{
		SignedTransaction: hex.EncodeToString(signedTx),
	}, nil
}

func toSignatures(rosSigs []*types.Signature) []*crgtypes.Signature {
	crgSigs := make([]*crgtypes.Signature, len(rosSigs))

	for i, rosSig := range rosSigs {
		crgSigs[i] = &crgtypes.Signature{
			SigningPayload: &crgtypes.SigningPayload{
				AccountIdentifier: &crgtypes.AccountIdentifier{
					Address:    rosSig.SigningPayload.AccountIdentifier.Address,
					SubAccount: (*crgtypes.SubAccountIdentifier)(rosSig.SigningPayload.AccountIdentifier.SubAccount),
					Metadata:   rosSig.SigningPayload.AccountIdentifier.Metadata,
				},
				Bytes:         rosSig.SigningPayload.Bytes,
				SignatureType: (crgtypes.SignatureType)(rosSig.SigningPayload.SignatureType),
			},
			PublicKey:     nil,
			SignatureType: "",
			Bytes:         nil,
		}
	}

	return crgSigs
}

func (on OnlineNetwork) ConstructionDerive(ctx context.Context, request *types.ConstructionDeriveRequest) (*types.ConstructionDeriveResponse, *types.Error) {
	account, err := on.client.AccountIdentifierFromPublicKey(&crgtypes.PublicKey{
		Bytes:     request.PublicKey.Bytes,
		CurveType: (crgtypes.CurveType)(request.PublicKey.CurveType),
	})
	if err != nil {
		return nil, errors.ToRosetta(err)
	}
	return &types.ConstructionDeriveResponse{
		AccountIdentifier: convertAccountIdentifiers(account)[0],
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
	ops, signers, err := on.client.TxOperationsAndSignersAccountIdentifiers(request.Signed, txBytes)
	if err != nil {
		return nil, errors.ToRosetta(err)
	}
	return &types.ConstructionParseResponse{
		Operations:               convertOperations(ops),
		AccountIdentifierSigners: convertAccountIdentifiers(signers...),
		Metadata:                 nil,
	}, nil

}

func (on OnlineNetwork) ConstructionPayloads(ctx context.Context, request *types.ConstructionPayloadsRequest) (*types.ConstructionPayloadsResponse, *types.Error) {
	payload, err := on.client.ConstructionPayload(ctx, &crgtypes.ConstructionPayloadsRequest{
		NetworkIdentifier: &crgtypes.NetworkIdentifier{
			Blockchain:           request.NetworkIdentifier.Blockchain,
			Network:              request.NetworkIdentifier.Network,
			SubNetworkIdentifier: (*crgtypes.SubNetworkIdentifier)(request.NetworkIdentifier.SubNetworkIdentifier),
		},
		Operations: toOperations(request.Operations...),
		Metadata:   request.Metadata,
		PublicKeys: toPublicKeys(request.PublicKeys),
	})
	if err != nil {
		return nil, errors.ToRosetta(err)
	}
	return &types.ConstructionPayloadsResponse{
		UnsignedTransaction: payload.UnsignedTransaction,
		Payloads:            convertPayloads(payload.Payloads),
	}, nil
}

func convertPayloads(crgPayloads []*crgtypes.SigningPayload) []*types.SigningPayload {
	rosPayloads := make([]*types.SigningPayload, len(crgPayloads))

	for i, crgPayload := range crgPayloads {
		rosPayloads[i] = &types.SigningPayload{
			AccountIdentifier: convertAccountIdentifiers(crgPayload.AccountIdentifier)[0],
			Bytes:             crgPayload.Bytes,
			SignatureType:     (types.SignatureType)(crgPayload.SignatureType),
		}
	}

	return rosPayloads
}

func (on OnlineNetwork) ConstructionPreprocess(ctx context.Context, request *types.ConstructionPreprocessRequest) (*types.ConstructionPreprocessResponse, *types.Error) {
	options, err := on.client.PreprocessOperationsToOptions(ctx, &crgtypes.ConstructionPreprocessRequest{
		NetworkIdentifier: &crgtypes.NetworkIdentifier{
			Blockchain:           request.NetworkIdentifier.Blockchain,
			Network:              request.NetworkIdentifier.Network,
			SubNetworkIdentifier: (*crgtypes.SubNetworkIdentifier)(request.NetworkIdentifier.SubNetworkIdentifier),
		},
		Operations:             toOperations(request.Operations...),
		Metadata:               request.Metadata,
		MaxFee:                 toBalances(request.MaxFee)[0],
		SuggestedFeeMultiplier: request.SuggestedFeeMultiplier,
	})
	if err != nil {
		return nil, errors.ToRosetta(err)
	}

	return &types.ConstructionPreprocessResponse{
		Options: options,
	}, nil
}

func toOperations(rosOps ...*types.Operation) []*crgtypes.Operation {
	crgOps := make()
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
		TransactionIdentifier: (*types.TransactionIdentifier)(res),
		Metadata:              meta,
	}, nil
}
