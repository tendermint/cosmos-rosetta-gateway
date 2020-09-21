package altsdk

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/altsdk/types"
	"io/ioutil"
	"net/http"
)

func (c Client) GetAuthAccount(ctx context.Context, address string) (types.AccountResponse, error) {
	var accRes rest.ResponseWithHeight
	path := fmt.Sprintf("/auth/accounts/%s", address)

	r, err := http.Get(c.getEndpoint(path))
	btes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return types.AccountResponse{}, err
	}

	codec := simapp.MakeCodec()
	if err = codec.UnmarshalJSON(btes, &accRes); err != nil {
		return types.AccountResponse{}, err
	}

	var res types.Response
	err = json.Unmarshal(accRes.Result, &res)
	if err != nil {
		return types.AccountResponse{}, err
	}
	defer r.Body.Close()

	return types.AccountResponse{
		Height: accRes.Height,
		Result: res,
	}, nil
}
