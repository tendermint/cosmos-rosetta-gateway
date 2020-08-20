package rosetta

import (
	"github.com/coinbase/rosetta-sdk-go/server"
)

type Adapter interface {
	DataAPI
	ConstructionAPI
}

type DataAPI interface {
	server.NetworkAPIServicer
	server.AccountAPIServicer
	server.MempoolAPIServicer
}

type ConstructionAPI interface {
}

type NetworkProperties struct {
	Blockchain          string
	Network             string
	SupportedOperations []string
}
