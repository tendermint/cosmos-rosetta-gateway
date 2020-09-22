package launchpad

import (
	"context"
	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/altsdk/types"
	"net/http"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/alttendermint"

	cosmosclient "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/generated"
	tendermintclient "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint/generated"
)

type CosmosAPI struct {
	Tendermint   CosmosTendermintAPI
	Transactions CosmosTransactionsAPI
}

type CosmosTransactionsAPI interface {
	TxsHashGet(ctx context.Context, hash string) (cosmosclient.TxQuery, *http.Response, error)
	TxsPost(ctx context.Context, txBroadcast cosmosclient.InlineObject) (cosmosclient.BroadcastTxCommitResult, *http.Response, error)
}

type SdkClient interface {
	GetAuthAccount(ctx context.Context, address string) (types.AccountResponse, error)
}

type CosmosTendermintAPI interface {
	NodeInfoGet(ctx context.Context) (cosmosclient.InlineResponse200, *http.Response, error)
}

type TendermintAPI struct {
	Info TendermintInfoAPI
}

type TendermintInfoAPI interface {
	TxSearch(ctx context.Context, query string, localVarOptionals *tendermintclient.TxSearchOpts) (tendermintclient.TxSearchResponse, *http.Response, error)
	Tx(ctx context.Context, hash string, localVarOptionals *tendermintclient.TxOpts) (tendermintclient.TxResponse, *http.Response, error)
}

// New Interfaces.
type TendermintClient interface {
	NetInfo() (alttendermint.NetInfoResponse, error)
	Block(height uint64) (alttendermint.BlockResponse, error)
	BlockByHash(hash string) (alttendermint.BlockResponse, error)
	UnconfirmedTxs() (alttendermint.UnconfirmedTxsResponse, error)
}
