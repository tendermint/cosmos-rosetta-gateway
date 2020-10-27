package sdk

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	grpctypes "github.com/cosmos/cosmos-sdk/types/grpc"
	auth "github.com/cosmos/cosmos-sdk/x/auth/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	sdktypes "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/types"
)

// unwrapGRPCError is used to unwrap gRPC errors to standard rosetta errors
// TODO
func unwrapGRPCError(err error) error {
	return err
}

type Client struct {
	clientCtx  client.Context
	authClient auth.QueryClient
	bankClient bank.QueryClient
}

func NewClient(endpoint string) Client {
	// create encoding
	encodingConfig := simapp.MakeEncodingConfig()
	// init client context
	initClientCtx := client.Context{}.
		WithJSONMarshaler(encodingConfig.Marshaler).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithAccountRetriever(auth.AccountRetriever{}).
		WithNodeURI(endpoint)

	return Client{
		authClient: auth.NewQueryClient(initClientCtx),
		bankClient: bank.NewQueryClient(initClientCtx),
		clientCtx:  initClientCtx,
	}
}

// GetAuthAccount gets the account information in the specified height
// as of Stargate to fulfill the sdktypes.AccountResponse we need
// to do two queries, one directed to the authentication module
// and the other one to the bank module
func (c Client) GetAuthAccount(ctx context.Context, address string, height int64) (resp sdktypes.AccountResponse, err error) {
	// parse addr
	addr, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return sdktypes.AccountResponse{}, err
	}
	// get account information
	account, err := c.clientCtx.AccountRetriever.GetAccount(c.clientCtx.WithHeight(height), addr) // use the specified height
	if err != nil {
		return sdktypes.AccountResponse{}, unwrapGRPCError(err)
	}
	// get balance information
	// update the context metadata to add the height header information
	ctx = context.WithValue(ctx, grpctypes.GRPCBlockHeightHeader, height)
	balances, err := c.bankClient.AllBalances(ctx, &bank.QueryAllBalancesRequest{
		Address:    addr.String(),
		Pagination: nil,
	})
	if err != nil {
		return sdktypes.AccountResponse{}, unwrapGRPCError(err)
	}
	// transform response
	resp = sdktypes.AccountResponse{
		Height: height,
		Result: sdktypes.Response{
			Type: "", // type does not apply here as it technically is multiple types
			Value: sdktypes.BaseAccount{
				Address: addr.String(),
				Coins:   balances.Balances,
				PubKey: sdktypes.PublicKey{
					Type:  account.GetPubKey().Type(),
					Value: fmt.Sprintf("%x", account.GetPubKey().Bytes()), // is this correct?
				},
				AccountNumber: sdk.NewIntFromUint64(account.GetAccountNumber()).String(),
				Sequence:      sdk.NewIntFromUint64(account.GetSequence()).String(),
			},
		},
	}
	// success
	return resp, nil
}

func (c Client) GetTx(ctx context.Context, hash string) (sdk.TxResponse, error) {
	panic("implement me")
}

func (c Client) PostTx(ctx context.Context, bytes []byte) (sdk.TxResponse, error) {
	panic("implement me")
}

func (c Client) GetNodeInfo(_ context.Context) (rpc.NodeInfoResponse, error) {
	panic("implement me")
}
