package sdk

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/types"
)

func (c Client) GetAuthAccount(ctx context.Context, address string, height int64) (types.Response, error) {
	path := fmt.Sprintf("/bank/accounts/%s?height=%d", address, height)
	r, err := http.Get(c.getEndpoint(path))
	if err != nil {
		return types.Response{}, err
	}
	if r == nil {
		return types.Response{}, fmt.Errorf("unable to fetch data from endpoint %s", c.getEndpoint(path))
	}
	btes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return types.Response{}, err
	}

	var res types.Response
	err = json.Unmarshal(btes, &res)
	if err != nil {
		return types.Response{}, err
	}
	defer r.Body.Close()

	return res, nil
}
