package sdk

import (
	"context"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/simapp"
	"io/ioutil"
	"net/http"
)

func (c Client) GetNodeInfo(ctx context.Context) (rpc.NodeInfoResponse, error) {
	r, err := http.Get(c.getEndpoint("/node_info"))
	if err != nil {
		return rpc.NodeInfoResponse{}, err
	}
	btes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return rpc.NodeInfoResponse{}, err
	}

	var infoRes rpc.NodeInfoResponse
	codec := simapp.MakeCodec()
	if err = codec.UnmarshalJSON(btes, &infoRes); err != nil {
		return infoRes, err
	}

	return infoRes, nil

}
