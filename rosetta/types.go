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
	server.BlockAPIServicer
	server.ConstructionAPIServicer
}

type ConstructionAPI interface {
	server.ConstructionAPIServicer
}

type NetworkProperties struct {
	// Mandatory properties
	Blockchain          string
	Network             string
	SupportedOperations []string
	OfflineMode         bool
}
