package altsdk

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/altsdk/types"
	"io/ioutil"
	"net/http"
)

func (c Client) GetTxs(ctx context.Context, hash string) (types.TxResponse, error) {
	path := fmt.Sprintf("/txs/%s", hash)

	r, err := http.Get(c.getEndpoint(path))
	btes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return types.TxResponse{}, err
	}

	var txRes types.TxResponse
	codec := simapp.MakeCodec()
	if err = codec.UnmarshalJSON(btes, &txRes); err != nil {
		return types.TxResponse{}, err
	}
	return txRes, nil
}
