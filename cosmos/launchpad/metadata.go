package launchpad

import (
	"fmt"

	"github.com/coinbase/rosetta-sdk-go/types"
)

const (
	// Metadata Keys
	ChainIdKey    = "chain_id"
	SequenceKey   = "sequence"
	AccountNumber = "account_number"
)

type PayloadReqMetadata struct {
	ChainId       string
	Sequence      uint64
	AccountNumber uint64
}

// GetMetadataFromPayloadReq obtains the metadata from the request to /construction/payloads endpoint.
func GetMetadataFromPayloadReq(req *types.ConstructionPayloadsRequest) (*PayloadReqMetadata, error) {
	chainId, ok := req.Metadata[ChainIdKey].(string)
	if !ok {
		return nil, fmt.Errorf("chain_id metadata was not provided")
	}

	sequence, ok := req.Metadata[SequenceKey]
	if !ok {
		return nil, fmt.Errorf("sequence metadata was not provided")
	}
	seqF64, ok := sequence.(float64)
	if !ok {
		return nil, fmt.Errorf("invalid sequence value")
	}

	accountNum, ok := req.Metadata[AccountNumber]
	if !ok {
		return nil, fmt.Errorf("account_number metadata was not provided")
	}
	accF64, ok := accountNum.(float64)
	if !ok {
		return nil, fmt.Errorf("invalid account_number value")
	}

	return &PayloadReqMetadata{
		ChainId:       chainId,
		Sequence:      uint64(seqF64),
		AccountNumber: uint64(accF64),
	}, nil
}
