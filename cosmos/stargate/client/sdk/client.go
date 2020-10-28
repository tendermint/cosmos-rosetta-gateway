package sdk

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	grpctypes "github.com/cosmos/cosmos-sdk/types/grpc"
	auth "github.com/cosmos/cosmos-sdk/x/auth/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	sdktypes "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/types"
	"google.golang.org/grpc"
)

// unwrapGRPCError is used to unwrap gRPC errors to standard rosetta errors
// TODO
func unwrapGRPCError(err error) error {
	return err
}

type Client struct {
	authClient   auth.QueryClient
	bankClient   bank.QueryClient
	encodeConfig types.InterfaceRegistry
}

func NewClient(endpoint string) (Client, error) {
	// instantiate gRPC connection
	grpcConn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	// create interface registry, and register used modules types
	interfaceRegistry := types.NewInterfaceRegistry()
	auth.RegisterInterfaces(interfaceRegistry)
	bank.RegisterInterfaces(interfaceRegistry)
	cryptocodec.RegisterInterfaces(interfaceRegistry)
	return Client{
		authClient:   auth.NewQueryClient(grpcConn),
		bankClient:   bank.NewQueryClient(grpcConn),
		encodeConfig: interfaceRegistry,
	}, nil
}

// GetAuthAccount gets the account information in the specified height
// as of Stargate to fulfill the sdktypes.AccountResponse we need
// to do two queries, one directed to the authentication module
// and the other one to the bank module
func (c Client) GetAuthAccount(ctx context.Context, address string, height int64) (resp sdktypes.AccountResponse, err error) {
	// update the context metadata to add the height header information
	ctx = context.WithValue(ctx, grpctypes.GRPCBlockHeightHeader, height)
	// get account information
	rawAccount, err := c.authClient.Account(ctx, &auth.QueryAccountRequest{Address: address})
	if err != nil {
		return sdktypes.AccountResponse{}, unwrapGRPCError(err)
	}
	// decode any to raw account
	var account auth.AccountI
	err = c.encodeConfig.UnpackAny(rawAccount.Account, &account)
	if err != nil {
		return sdktypes.AccountResponse{}, err
	}
	// get balance information
	balances, err := c.bankClient.AllBalances(ctx, &bank.QueryAllBalancesRequest{
		Address:    address,
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
				Address: address,
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
