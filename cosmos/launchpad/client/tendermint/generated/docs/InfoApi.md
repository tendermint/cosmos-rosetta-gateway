# \InfoApi

All URIs are relative to *https://rpc.cosmos.network*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Block**](InfoApi.md#Block) | **Get** /block | Get block at a specified height
[**BlockByHash**](InfoApi.md#BlockByHash) | **Get** /block_by_hash | Get block by hash
[**NetInfo**](InfoApi.md#NetInfo) | **Get** /net_info | Network informations
[**NumUnconfirmedTxs**](InfoApi.md#NumUnconfirmedTxs) | **Get** /num_unconfirmed_txs | Get data about unconfirmed transactions
[**Status**](InfoApi.md#Status) | **Get** /status | Node Status
[**Tx**](InfoApi.md#Tx) | **Get** /tx | Get transactions by hash
[**TxSearch**](InfoApi.md#TxSearch) | **Get** /tx_search | Search for transactions
[**UnconfirmedTxs**](InfoApi.md#UnconfirmedTxs) | **Get** /unconfirmed_txs | Get the list of unconfirmed transactions



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

