package rosetta

import "github.com/coinbase/rosetta-sdk-go/types"

// GetOperationsByRelation parses a list of operations and returns them in arrays
// of related operations. For example: An Transfer Operation normally is 2 operations together.
func GetOperationsByRelation(ops []*types.Operation) [][]*types.Operation {
	var relatedOps map[string][]*types.Operation

	for _, op := range ops {
		if op.RelatedOperations != nil {
			relatedOps[] =
		}
	}

	return nil
}
