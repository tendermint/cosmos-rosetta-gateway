package types

import (
	"encoding/hex"
	"encoding/json"
	"github.com/coinbase/rosetta-sdk-go/types"
)

var _ = types.OperationStatus(OperationStatus{})

type OperationStatus struct {
	Status     string `json:"status"`
	Successful bool   `json:"successful"`
}

var _ = types.Currency(Currency{})

type Currency struct {
	Symbol   string                 `json:"symbol"`
	Decimals int32                  `json:"decimals"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type Amount struct {
	Value    string                 `json:"value"`
	Currency *Currency              `json:"currency"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

var _ = types.TransactionIdentifier(TransactionIdentifier{})

type TransactionIdentifier struct {
	// Any transactions that are attributable only to a block (ex: a block event) should use the
	// hash of the block as the identifier.
	Hash string `json:"hash"`
}

var _ = types.OperationIdentifier(OperationIdentifier{})

type OperationIdentifier struct {
	Index        int64  `json:"index"`
	NetworkIndex *int64 `json:"network_index,omitempty"`
}

var _ = types.SubAccountIdentifier(SubAccountIdentifier{})

type SubAccountIdentifier struct {
	Address  string                 `json:"address"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type AccountIdentifier struct {
	Address    string                 `json:"address"`
	SubAccount *SubAccountIdentifier  `json:"sub_account,omitempty"`
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
}

type Operation struct {
	OperationIdentifier *OperationIdentifier   `json:"operation_identifier"`
	RelatedOperations   []*OperationIdentifier `json:"related_operations,omitempty"`
	Type                string                 `json:"type"`
	Status              string                 `json:"status"`
	Account             *AccountIdentifier     `json:"account,omitempty"`
	Amount              *Amount                `json:"amount,omitempty"`
	Metadata            map[string]interface{} `json:"metadata,omitempty"`
}

type Transaction struct {
	TransactionIdentifier *TransactionIdentifier `json:"transaction_identifier"`
	Operations            []*Operation           `json:"operations"`
	// Transactions that are related to other transactions (like a cross-shard transaction) should
	// include the tranaction_identifier of these transactions in the metadata.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

var _ = types.Peer(Peer{})

type Peer struct {
	PeerID   string                 `json:"peer_id"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

var _ = types.SyncStatus(SyncStatus{})

type SyncStatus struct {
	CurrentIndex int64   `json:"current_index"`
	TargetIndex  *int64  `json:"target_index,omitempty"`
	Stage        *string `json:"stage,omitempty"`
}

var _ = types.BlockIdentifier(BlockIdentifier{})

type BlockIdentifier struct {
	Index int64  `json:"index"`
	Hash  string `json:"hash"`
}

type SignatureType string

// List of SignatureType
const (
	Ecdsa           SignatureType = "ecdsa"
	EcdsaRecovery   SignatureType = "ecdsa_recovery"
	Ed25519         SignatureType = "ed25519"
	Schnorr1        SignatureType = "schnorr_1"
	SchnorrPoseidon SignatureType = "schnorr_poseidon"
)

type CurveType string

// List of CurveType
const (
	Secp256k1    CurveType = "secp256k1"
	Secp256r1    CurveType = "secp256r1"
	Edwards25519 CurveType = "edwards25519"
	Tweedle      CurveType = "tweedle"
)

// PublicKey PublicKey contains a public key byte array for a particular CurveType encoded in hex.
// Note that there is no PrivateKey struct as this is NEVER the concern of an implementation.
type PublicKey struct {
	Bytes     []byte    `json:"hex_bytes"`
	CurveType CurveType `json:"curve_type"`
}

// MarshalJSON overrides the default JSON marshaler
// and encodes bytes as hex instead of base64.
func (s *PublicKey) MarshalJSON() ([]byte, error) {
	type Alias PublicKey
	j, err := json.Marshal(struct {
		Bytes string `json:"hex_bytes"`
		*Alias
	}{
		Bytes: hex.EncodeToString(s.Bytes),
		Alias: (*Alias)(s),
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

// UnmarshalJSON overrides the default JSON unmarshaler
// and decodes bytes from hex instead of base64.
func (s *PublicKey) UnmarshalJSON(b []byte) error {
	type Alias PublicKey
	r := struct {
		Bytes string `json:"hex_bytes"`
		*Alias
	}{
		Alias: (*Alias)(s),
	}
	err := json.Unmarshal(b, &r)
	if err != nil {
		return err
	}

	bytes, err := hex.DecodeString(r.Bytes)
	if err != nil {
		return err
	}

	s.Bytes = bytes
	return nil
}

type Signature struct {
	SigningPayload *SigningPayload `json:"signing_payload"`
	PublicKey      *PublicKey      `json:"public_key"`
	SignatureType  SignatureType   `json:"signature_type"`
	Bytes          []byte          `json:"hex_bytes"`
}

// MarshalJSON overrides the default JSON marshaler
// and encodes bytes as hex instead of base64.
func (s *Signature) MarshalJSON() ([]byte, error) {
	type Alias Signature
	j, err := json.Marshal(struct {
		Bytes string `json:"hex_bytes"`
		*Alias
	}{
		Bytes: hex.EncodeToString(s.Bytes),
		Alias: (*Alias)(s),
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

// UnmarshalJSON overrides the default JSON unmarshaler
// and decodes bytes from hex instead of base64.
func (s *Signature) UnmarshalJSON(b []byte) error {
	type Alias Signature
	r := struct {
		Bytes string `json:"hex_bytes"`
		*Alias
	}{
		Alias: (*Alias)(s),
	}
	err := json.Unmarshal(b, &r)
	if err != nil {
		return err
	}

	bytes, err := hex.DecodeString(r.Bytes)
	if err != nil {
		return err
	}

	s.Bytes = bytes
	return nil
}

// SigningPayload SigningPayload is signed by the client with the keypair associated with an
// AccountIdentifier using the specified SignatureType. SignatureType can be optionally populated if
// there is a restriction on the signature scheme that can be used to sign the payload.
type SigningPayload struct {
	AccountIdentifier *AccountIdentifier `json:"account_identifier,omitempty"`
	Bytes             []byte             `json:"hex_bytes"`
	SignatureType     SignatureType      `json:"signature_type,omitempty"`
}

// MarshalJSON overrides the default JSON marshaler
// and encodes bytes as hex instead of base64. It also
// writes the deprecated "address" field to the response.
func (s *SigningPayload) MarshalJSON() ([]byte, error) {
	type Alias SigningPayload
	addressString := ""
	if s.AccountIdentifier != nil {
		addressString = s.AccountIdentifier.Address
	}

	j, err := json.Marshal(struct {
		// [DEPRECATED by `account_identifier` in `v1.4.4`] Address is the network-specific
		// address of the account that should sign the payload.
		Address string `json:"address,omitempty"`
		Bytes   string `json:"hex_bytes"`
		*Alias
	}{
		Address: addressString,
		Bytes:   hex.EncodeToString(s.Bytes),
		Alias:   (*Alias)(s),
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

// UnmarshalJSON overrides the default JSON unmarshaler
// and decodes bytes from hex instead of base64. It also
// reads the deprecated "address" field from the response.
func (s *SigningPayload) UnmarshalJSON(b []byte) error {
	type Alias SigningPayload
	r := struct {
		// [DEPRECATED by `account_identifier` in `v1.4.4`] Address is the network-specific
		// address of the account that should sign the payload.
		Address string `json:"address,omitempty"`
		Bytes   string `json:"hex_bytes"`
		*Alias
	}{
		Alias: (*Alias)(s),
	}
	err := json.Unmarshal(b, &r)
	if err != nil {
		return err
	}

	bytes, err := hex.DecodeString(r.Bytes)
	if err != nil {
		return err
	}
	s.Bytes = bytes

	if s.AccountIdentifier == nil && len(r.Address) > 0 {
		s.AccountIdentifier = &AccountIdentifier{
			Address: r.Address,
		}
	}

	return nil
}

// ConstructionPayloadsResponse ConstructionTransactionResponse is returned by
// `/construction/payloads`. It contains an unsigned transaction blob (that is usually needed to
// construct the a network transaction from a collection of signatures) and an array of payloads
// that must be signed by the caller.
type ConstructionPayloadsResponse struct {
	UnsignedTransaction string            `json:"unsigned_transaction"`
	Payloads            []*SigningPayload `json:"payloads"`
}

type ConstructionPayloadsRequest struct {
	NetworkIdentifier *NetworkIdentifier     `json:"network_identifier"`
	Operations        []*Operation           `json:"operations"`
	Metadata          map[string]interface{} `json:"metadata,omitempty"`
	PublicKeys        []*PublicKey           `json:"public_keys,omitempty"`
}

// NetworkIdentifier The network_identifier specifies which network a particular object is
// associated with.
type NetworkIdentifier struct {
	Blockchain string `json:"blockchain"`
	// If a blockchain has a specific chain-id or network identifier, it should go in this field. It
	// is up to the client to determine which network-specific identifier is mainnet or testnet.
	Network              string                `json:"network"`
	SubNetworkIdentifier *SubNetworkIdentifier `json:"sub_network_identifier,omitempty"`
}

// SubNetworkIdentifier In blockchains with sharded state, the SubNetworkIdentifier is required to
// query some object on a specific shard. This identifier is optional for all non-sharded
// blockchains.
type SubNetworkIdentifier struct {
	Network  string                 `json:"network"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type ConstructionPreprocessRequest struct {
	NetworkIdentifier      *NetworkIdentifier     `json:"network_identifier"`
	Operations             []*Operation           `json:"operations"`
	Metadata               map[string]interface{} `json:"metadata,omitempty"`
	MaxFee                 []*Amount              `json:"max_fee,omitempty"`
	SuggestedFeeMultiplier *float64               `json:"suggested_fee_multiplier,omitempty"`
}
