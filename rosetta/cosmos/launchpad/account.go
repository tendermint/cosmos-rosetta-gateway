package launchpad

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/coinbase/rosetta-sdk-go/types"
)

func (l Launchpad) AccountBalance(ctx context.Context, request *types.AccountBalanceRequest) (
	*types.AccountBalanceResponse, *types.Error) {

	get, err := http.Get("http://localhost:1317/bank/balances/cosmos15f92rjkapauptyw6lt94rlwq4dcg99nncwc8na")
	if err != nil {
		return nil, &types.Error{Code: 1}
	}

	body := get.Body
	defer body.Close()

	all, _ := ioutil.ReadAll(body)

	var bal balanceResp
	err = json.Unmarshal(all, &bal)
	if err != nil {
		return nil, &types.Error{Code: 1}
	}

	fmt.Printf("%s\n", all)

	return &types.AccountBalanceResponse{
		BlockIdentifier: nil,
		Balances:        nil,
		Coins:           nil,
		Metadata:        nil,
	}, nil
}

type balanceResp struct {
	result string
}