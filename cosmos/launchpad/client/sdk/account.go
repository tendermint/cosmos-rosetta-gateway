package sdk

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/types"
)

func (c Client) GetAuthAccount(ctx context.Context, address string, height string) (types.AccountResponse, error) {
	var accRes rest.ResponseWithHeight
	path := fmt.Sprintf("/auth/accounts/%s?height=%s", address, height)
	r, err := http.Get(c.getEndpoint(path))
	if err != nil {
		return types.AccountResponse{}, err
	}
	btes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return types.AccountResponse{}, err
	}

	if err = types.Codec.UnmarshalJSON(btes, &accRes); err != nil {
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
