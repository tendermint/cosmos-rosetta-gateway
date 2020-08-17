# \InfoApi

All URIs are relative to *https://rpc.cosmos.network*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Block**](InfoApi.md#Block) | **Get** /block | Get block at a specified height
[**BlockByHash**](InfoApi.md#BlockByHash) | **Get** /block_by_hash | Get block by hash
[**BlockResults**](InfoApi.md#BlockResults) | **Get** /block_results | Get block results at a specified height
[**Blockchain**](InfoApi.md#Blockchain) | **Get** /blockchain | Get block headers for minHeight &lt;&#x3D; height &lt;&#x3D; maxHeight.
[**BroadcastEvidence**](InfoApi.md#BroadcastEvidence) | **Get** /broadcast_evidence | Broadcast evidence of the misbehavior.
[**Commit**](InfoApi.md#Commit) | **Get** /commit | Get commit results at a specified height
[**ConsensusParams**](InfoApi.md#ConsensusParams) | **Get** /consensus_params | Get consensus parameters
[**ConsensusState**](InfoApi.md#ConsensusState) | **Get** /consensus_state | Get consensus state
[**DumpConsensusState**](InfoApi.md#DumpConsensusState) | **Get** /dump_consensus_state | Get consensus state
[**Genesis**](InfoApi.md#Genesis) | **Get** /genesis | Get Genesis
[**Health**](InfoApi.md#Health) | **Get** /health | Node heartbeat
[**NetInfo**](InfoApi.md#NetInfo) | **Get** /net_info | Network informations
[**NumUnconfirmedTxs**](InfoApi.md#NumUnconfirmedTxs) | **Get** /num_unconfirmed_txs | Get data about unconfirmed transactions
[**Status**](InfoApi.md#Status) | **Get** /status | Node Status
[**Tx**](InfoApi.md#Tx) | **Get** /tx | Get transactions by hash
[**TxSearch**](InfoApi.md#TxSearch) | **Get** /tx_search | Search for transactions
[**UnconfirmedTxs**](InfoApi.md#UnconfirmedTxs) | **Get** /unconfirmed_txs | Get the list of unconfirmed transactions
[**Validators**](InfoApi.md#Validators) | **Get** /validators | Get validator set at a specified height



## Block

> BlockResponse Block(ctx, optional)

Get block at a specified height

Get Block. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BlockOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BlockOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **height** | **optional.Float32**| height to return. If no height is provided, it will fetch the latest block. | [default to 0]

### Return type

[**BlockResponse**](BlockResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BlockByHash

> BlockResponse BlockByHash(ctx, hash)

Get block by hash

Get Block By Hash. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string**| block hash | 

### Return type

[**BlockResponse**](BlockResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BlockResults

> BlockResultsResponse BlockResults(ctx, optional)

Get block results at a specified height

Get block_results. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BlockResultsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BlockResultsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **height** | **optional.Float32**| height to return. If no height is provided, it will fetch informations regarding the latest block. | [default to 0]

### Return type

[**BlockResultsResponse**](BlockResultsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Blockchain

> BlockchainResponse Blockchain(ctx, optional)

Get block headers for minHeight <= height <= maxHeight.

Get Blockchain info. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BlockchainOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BlockchainOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **minHeight** | **optional.Float32**| Minimum block height to return | 
 **maxHeight** | **optional.Float32**| Maximum block height to return | 

### Return type

[**BlockchainResponse**](BlockchainResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BroadcastEvidence

> BroadcastEvidenceResponse BroadcastEvidence(ctx, evidence)

Broadcast evidence of the misbehavior.

Broadcast evidence of the misbehavior. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**evidence** | **string**| JSON evidence | 

### Return type

[**BroadcastEvidenceResponse**](BroadcastEvidenceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Commit

> CommitResponse Commit(ctx, optional)

Get commit results at a specified height

Get Commit. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CommitOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CommitOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **height** | **optional.Float32**| height to return. If no height is provided, it will fetch commit informations regarding the latest block. | [default to 0]

### Return type

[**CommitResponse**](CommitResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsensusParams

> ConsensusParamsResponse ConsensusParams(ctx, optional)

Get consensus parameters

Get consensus parameters. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConsensusParamsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ConsensusParamsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **height** | **optional.Float32**| height to return. If no height is provided, it will fetch commit informations regarding the latest block. | [default to 0]

### Return type

[**ConsensusParamsResponse**](ConsensusParamsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsensusState

> ConsensusStateResponse ConsensusState(ctx, )

Get consensus state

Get consensus state.  Not safe to call from inside the ABCI application during a block execution. 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ConsensusStateResponse**](ConsensusStateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DumpConsensusState

> DumpConsensusResponse DumpConsensusState(ctx, )

Get consensus state

Get consensus state.  Not safe to call from inside the ABCI application during a block execution. 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**DumpConsensusResponse**](DumpConsensusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Genesis

> GenesisResponse Genesis(ctx, )

Get Genesis

Get genesis. 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**GenesisResponse**](GenesisResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Health

> EmptyResponse Health(ctx, )

Node heartbeat

Get node health. Returns empty result (200 OK) on success, no response - in case of an error. 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**EmptyResponse**](EmptyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetInfo

> NetInfoResponse NetInfo(ctx, )

Network informations

Get network info. 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**NetInfoResponse**](NetInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NumUnconfirmedTxs

> NumUnconfirmedTransactionsResponse NumUnconfirmedTxs(ctx, )

Get data about unconfirmed transactions

Get data about unconfirmed transactions 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**NumUnconfirmedTransactionsResponse**](NumUnconfirmedTransactionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Status

> StatusResponse Status(ctx, )

Node Status

Get Tendermint status including node info, pubkey, latest block hash, app hash, block height and time. 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**StatusResponse**](StatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Tx

> TxResponse Tx(ctx, hash, optional)

Get transactions by hash

Get a trasasction 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string**| transaction Hash to retrive | 
 **optional** | ***TxOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TxOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **prove** | **optional.Bool**| Include proofs of the transactions inclusion in the block | [default to false]

### Return type

[**TxResponse**](TxResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TxSearch

> TxSearchResponse TxSearch(ctx, query, optional)

Search for transactions

Search for transactions w/ their results.  See /subscribe for the query syntax. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**query** | **string**| Query | 
 **optional** | ***TxSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TxSearchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **prove** | **optional.Bool**| Include proofs of the transactions inclusion in the block | [default to false]
 **page** | **optional.Float32**| Page number (1-based) | [default to 1]
 **perPage** | **optional.Float32**| Number of entries per page (max: 100) | [default to 30]
 **orderBy** | **optional.String**| Order in which transactions are sorted (\&quot;asc\&quot; or \&quot;desc\&quot;), by height &amp; index. If empty, default sorting will be still applied. | [default to asc]

### Return type

[**TxSearchResponse**](TxSearchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnconfirmedTxs

> UnconfirmedTransactionsResponse UnconfirmedTxs(ctx, optional)

Get the list of unconfirmed transactions

Get list of unconfirmed transactions 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UnconfirmedTxsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UnconfirmedTxsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Float32**| Maximum number of unconfirmed transactions to return (max 100) | [default to 30]

### Return type

[**UnconfirmedTransactionsResponse**](UnconfirmedTransactionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Validators

> ValidatorsResponse Validators(ctx, optional)

Get validator set at a specified height

Get Validators. Validators are sorted by voting power. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidatorsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidatorsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **height** | **optional.Float32**| height to return. If no height is provided, it will fetch validator set which corresponds to the latest block. | [default to 0]
 **page** | **optional.Float32**| Page number (1-based) | [default to 1]
 **perPage** | **optional.Float32**| Number of entries per page (max: 100) | [default to 30]

### Return type

[**ValidatorsResponse**](ValidatorsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

