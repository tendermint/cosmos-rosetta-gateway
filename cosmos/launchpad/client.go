package launchpad

import (
	"context"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/types"
	"net/http"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/alttendermint"

	tendermintclient "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint/generated"
)

type SdkClient interface {
	GetAuthAccount(ctx context.Context, address string) (types.AccountResponse, error)
	GetTx(ctx context.Context, hash string) (sdk.TxResponse, error)
	PostTx(ctx context.Context, bytes []byte) (sdk.TxResponse, error)
	GetNodeInfo(ctx context.Context) (rpc.NodeInfoResponse, error)
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
