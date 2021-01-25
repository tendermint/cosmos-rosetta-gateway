package types

import (
	"context"
	"github.com/coinbase/rosetta-sdk-go/server"
	"github.com/coinbase/rosetta-sdk-go/types"
)

// API defines the exposed APIs
// if the service is online
type API interface {
	DataAPI
	ConstructionAPI
}

// DataAPI defines the full data API implementation
type DataAPI interface {
	server.NetworkAPIServicer
	server.AccountAPIServicer
	server.BlockAPIServicer
	server.MempoolAPIServicer
}

var _ server.ConstructionAPIServicer = ConstructionAPI(nil)

// ConstructionAPI defines the full construction API with
// the online and offline endpoints
type ConstructionAPI interface {
	ConstructionOnlineAPI
	ConstructionOfflineAPI
}

// ConstructionOnlineAPI defines the construction methods
// allowed in an online implementation
type ConstructionOnlineAPI interface {
	ConstructionMetadata(
		context.Context,
		*types.ConstructionMetadataRequest,
	) (*types.ConstructionMetadataResponse, *types.Error)
	ConstructionSubmit(
		context.Context,
		*types.ConstructionSubmitRequest,
	) (*types.TransactionIdentifierResponse, *types.Error)
}

// ConstructionOfflineAPI defines the construction methods
// allowed
type ConstructionOfflineAPI interface {
	ConstructionCombine(
		context.Context,
		*types.ConstructionCombineRequest,
	) (*types.ConstructionCombineResponse, *types.Error)
	ConstructionDerive(
		context.Context,
		*types.ConstructionDeriveRequest,
	) (*types.ConstructionDeriveResponse, *types.Error)
	ConstructionHash(
		context.Context,
		*types.ConstructionHashRequest,
	) (*types.TransactionIdentifierResponse, *types.Error)
	ConstructionParse(
		context.Context,
		*types.ConstructionParseRequest,
	) (*types.ConstructionParseResponse, *types.Error)
	ConstructionPayloads(
		context.Context,
		*types.ConstructionPayloadsRequest,
	) (*types.ConstructionPayloadsResponse, *types.Error)
	ConstructionPreprocess(
		context.Context,
		*types.ConstructionPreprocessRequest,
	) (*types.ConstructionPreprocessResponse, *types.Error)
}
