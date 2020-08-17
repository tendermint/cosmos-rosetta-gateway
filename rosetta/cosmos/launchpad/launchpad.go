package launchpad

import (
	"context"
	"net/http"

	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
	cosmosclient "github.com/tendermint/cosmos-rosetta-gateway/rosetta/cosmos/launchpad/client/cosmos/generated"
	tendermintclient "github.com/tendermint/cosmos-rosetta-gateway/rosetta/cosmos/launchpad/client/tendermint/generated"
)

type Launchpad struct {
	cosmos     CosmosAPI
	tendermint TendermintAPI

	properties rosetta.NetworkProperties
}

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

func NewLaunchpad(tendermint TendermintAPI, cosmos CosmosAPI, properties rosetta.NetworkProperties) rosetta.Adapter {
	return &Launchpad{
		tendermint: tendermint,
		cosmos:     cosmos,
		properties: properties,
	}
}
