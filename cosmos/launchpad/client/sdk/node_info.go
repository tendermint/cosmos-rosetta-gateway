package sdk

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/sdk/types"
)

type NodeInfoResponse struct {
}

func (c Client) GetNodeInfo(ctx context.Context) (NodeInfoResponse, error) {
	r, err := http.Get(c.getEndpoint("/node_info"))
	if err != nil {
		return NodeInfoResponse{}, err
	}
	if r == nil {
		return NodeInfoResponse{}, fmt.Errorf("unable to fetch data from endpoint %s", c.getEndpoint("/node_info"))
	}
	btes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return NodeInfoResponse{}, err
	}

	var infoRes NodeInfoResponse
	if err = types.Codec.UnmarshalJSON(btes, &infoRes); err != nil {
		return infoRes, err
	}

	return infoRes, nil
}
