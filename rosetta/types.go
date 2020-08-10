package rosetta

import (
	"github.com/coinbase/rosetta-sdk-go/server"
)

type Adapter interface {
	DataAPI
}

type DataAPI interface {
	server.AccountAPIServicer
}
