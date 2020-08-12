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
}

type ConstructionAPI interface {
}

type NetworkProperties struct {
	// Mandatory properties
	Blockchain          string
	Network             string
	SupportedOperations []string
}
