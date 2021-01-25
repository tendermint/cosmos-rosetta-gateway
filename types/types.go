package types

import (
	"context"
)

// SpecVersion defines the specification of rosetta
const SpecVersion = ""

// NetworkInformationProvider defines the interface used to provide information regarding
// the network and the version of the cosmos sdk used
type NetworkInformationProvider interface {
	// SupportedOperations lists the operations supported by the implementation
	SupportedOperations() []string
	// OperationsStatuses returns the list of statuses supported by the implementation
	OperationStatuses() []*OperationStatus
	// Version returns the version of the node
	Version() string
}

// Client defines the API the client implementation should provide.
type Client interface {
	// Needed if the client needs to perform some action before connecting.
	Bootstrap() error
	// Ready checks if the servicer constraints for queries are satisfied
	// for example the node might still not be ready, it's useful in process
	// when the rosetta instance might come up before the node itself
	// the servicer must return nil if the node is ready
	Ready() error

	// Data API

	// Balances fetches the balance of the given address
	// if height is not nil, then the balance will be displayed
	// at the provided height, otherwise last block balance will be returned
	Balances(ctx context.Context, addr string, height *int64) ([]*Amount, error)
	// BlockByHashAlt gets a block and its transaction at the provided height
	BlockByHash(ctx context.Context, hash string) (BlockResponse, error)
	// BlockByHeightAlt gets a block given its height, if height is nil then last block is returned
	BlockByHeight(ctx context.Context, height *int64) (BlockResponse, error)
	// BlockTransactionsByHash gets the block, parent block and transactions
	// given the block hash.
	BlockTransactionsByHash(ctx context.Context, hash string) (BlockTransactionsResponse, error)
	// BlockTransactionsByHash gets the block, parent block and transactions
	// given the block hash.
	BlockTransactionsByHeight(ctx context.Context, height *int64) (BlockTransactionsResponse, error)
	// GetTx gets a transaction given its hash
	GetTx(ctx context.Context, hash string) (*Transaction, error)
	// GetUnconfirmedTx gets an unconfirmed Tx given its hash
	// NOTE(fdymylja): NOT IMPLEMENTED YET!
	GetUnconfirmedTx(ctx context.Context, hash string) (*Transaction, error)
	// Mempool returns the list of the current non confirmed transactions
	Mempool(ctx context.Context) ([]*TransactionIdentifier, error)
	// Peers gets the peers currently connected to the node
	Peers(ctx context.Context) ([]*Peer, error)
	// Status returns the node status, such as sync data, version etc
	Status(ctx context.Context) (*SyncStatus, error)

	// Construction API

	// PostTx posts txBytes to the node and returns the transaction identifier plus metadata related
	// to the transaction itself.
	PostTx(txBytes []byte) (res *TransactionIdentifier, meta map[string]interface{}, err error)
	// ConstructionMetadataFromOptions
	ConstructionMetadataFromOptions(ctx context.Context, options map[string]interface{}) (meta map[string]interface{}, err error)
	OfflineClient
}

// OfflineClient defines the functionalities supported without having access to the node
type OfflineClient interface {
	NetworkInformationProvider
	// SignedTx returns the signed transaction given the tx bytes (msgs) plus the signatures
	SignedTx(ctx context.Context, txBytes []byte, sigs []*Signature) (signedTxBytes []byte, err error)
	// TxOperationsAndSignersAccountIdentifiers returns the operations related to a transaction and the account
	// identifiers if the transaction is signed
	TxOperationsAndSignersAccountIdentifiers(signed bool, hexBytes []byte) (ops []*Operation, signers []*AccountIdentifier, err error)
	// ConstructionPayload returns the construction payload given the request
	ConstructionPayload(ctx context.Context, req *ConstructionPayloadsRequest) (resp *ConstructionPayloadsResponse, err error)
	// PreprocessOperationsToOptions returns the options given the preprocess operations
	PreprocessOperationsToOptions(ctx context.Context, req *ConstructionPreprocessRequest) (options map[string]interface{}, err error)
	// AccountIdentifierFromPublicKey returns the account identifier given the public key
	AccountIdentifierFromPublicKey(pubKey *PublicKey) (*AccountIdentifier, error)
}

type BlockTransactionsResponse struct {
	BlockResponse
	Transactions []*Transaction
}

type BlockResponse struct {
	Block                *BlockIdentifier
	ParentBlock          *BlockIdentifier
	MillisecondTimestamp int64
	TxCount              int64
}
