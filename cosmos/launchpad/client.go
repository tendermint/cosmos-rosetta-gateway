package launchpad

import (
	"context"
	"net/http"

	cosmosclient "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/generated"
	tendermintclient "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint/generated"
)

type CosmosAPI struct {
	Bank       CosmosBankAPI
	Tendermint CosmosTendermintAPI
}

type CosmosBankAPI interface {
	BankBalancesAddressGet(ctx context.Context, address string) (cosmosclient.InlineResponse2004, *http.Response, error)
}

type CosmosTendermintAPI interface {
	NodeInfoGet(ctx context.Context) (cosmosclient.InlineResponse200, *http.Response, error)
}

type TendermintAPI struct {
	Info TendermintInfoAPI
}

type TendermintInfoAPI interface {
	NetInfo(ctx context.Context) (tendermintclient.NetInfoResponse, *http.Response, error)
	Block(ctx context.Context, localVarOptionals *tendermintclient.BlockOpts) (tendermintclient.BlockResponse, *http.Response, error)
}
