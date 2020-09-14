package launchpad

import (
	"context"
	"net/http"

	cosmosclient "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/generated"
	tendermintclient "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint/generated"
)

type CosmosAPI struct {
	Auth         CosmosAuthAPI
	Bank         CosmosBankAPI
	Tendermint   CosmosTendermintAPI
	Transactions CosmosTransactionsAPI
}

type CosmosTransactionsAPI interface {
	TxsHashGet(ctx context.Context, hash string) (cosmosclient.TxQuery, *http.Response, error)
	TxsPost(ctx context.Context, txBroadcast cosmosclient.InlineObject) (cosmosclient.InlineResponse2003, *http.Response, error)
}

type CosmosBankAPI interface {
	BankBalancesAddressGet(ctx context.Context, address string) (cosmosclient.InlineResponse2005, *http.Response, error)
}

type CosmosAuthAPI interface {
	AuthAccountsAddressGet(ctx context.Context, address string) (cosmosclient.InlineResponse2006, *http.Response, error)
}

type CosmosTendermintAPI interface {
	NodeInfoGet(ctx context.Context) (cosmosclient.InlineResponse200, *http.Response, error)
}

type TendermintAPI struct {
	Info TendermintInfoAPI
	Tx   TendermintTxAPI
}

type TendermintInfoAPI interface {
	NetInfo(ctx context.Context) (tendermintclient.NetInfoResponse, *http.Response, error)
	Block(ctx context.Context, localVarOptionals *tendermintclient.BlockOpts) (tendermintclient.BlockResponse, *http.Response, error)
	UnconfirmedTxs(ctx context.Context, localVarOptionals *tendermintclient.UnconfirmedTxsOpts) (tendermintclient.UnconfirmedTransactionsResponse, *http.Response, error)
	BlockByHash(ctx context.Context, hash string) (tendermintclient.BlockResponse, *http.Response, error)
	TxSearch(ctx context.Context, query string, localVarOptionals *tendermintclient.TxSearchOpts) (tendermintclient.TxSearchResponse, *http.Response, error)
	Tx(ctx context.Context, hash string, localVarOptionals *tendermintclient.TxOpts) (tendermintclient.TxResponse, *http.Response, error)
}

type TendermintTxAPI interface {
	BroadcastTxAsync(ctx context.Context, tx string) (tendermintclient.BroadcastTxResponse, *http.Response, error)
}
