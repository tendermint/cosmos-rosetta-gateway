# \TransactionsApi

All URIs are relative to *https://stargate.cosmos.network*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TxsDecodePost**](TransactionsApi.md#TxsDecodePost) | **Post** /txs/decode | Decode a transaction from the Amino wire format
[**TxsEncodePost**](TransactionsApi.md#TxsEncodePost) | **Post** /txs/encode | Encode a transaction to the Amino wire format
[**TxsGet**](TransactionsApi.md#TxsGet) | **Get** /txs | Search transactions
[**TxsHashGet**](TransactionsApi.md#TxsHashGet) | **Get** /txs/{hash} | Get a Tx by hash
[**TxsPost**](TransactionsApi.md#TxsPost) | **Post** /txs | Broadcast a signed tx



## TxsDecodePost

> StdTx TxsDecodePost(ctx, tx)

Decode a transaction from the Amino wire format

Decode a transaction (signed or not) from base64-encoded Amino serialized bytes to JSON

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tx** | [**InlineObject2**](InlineObject2.md)|  | 

### Return type

[**StdTx**](StdTx.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TxsEncodePost

> InlineResponse2003 TxsEncodePost(ctx, tx)

Encode a transaction to the Amino wire format

Encode a transaction (signed or not) from JSON to base64-encoded Amino serialized bytes

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tx** | [**InlineObject1**](InlineObject1.md)|  | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TxsGet

> PaginatedQueryTxs TxsGet(ctx, optional)

Search transactions

Search transactions by events.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TxsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TxsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **messageAction** | **optional.String**| transaction events such as &#39;message.action&#x3D;send&#39; which results in the following endpoint: &#39;GET /txs?message.action&#x3D;send&#39;. note that each module documents its own events. look for xx_events.md in the corresponding cosmos-sdk/docs/spec directory | 
 **messageSender** | **optional.String**| transaction tags with sender: &#39;GET /txs?message.action&#x3D;send&amp;message.sender&#x3D;cosmos16xyempempp92x9hyzz9wrgf94r6j9h5f06pxxv&#39; | 
 **page** | **optional.Int32**| Page number | 
 **limit** | **optional.Int32**| Maximum number of items per page | 
 **txMinheight** | **optional.Int32**| transactions on blocks with height greater or equal this value | 
 **txMaxheight** | **optional.Int32**| transactions on blocks with height less than or equal this value | 

### Return type

[**PaginatedQueryTxs**](PaginatedQueryTxs.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TxsHashGet

> TxQuery TxsHashGet(ctx, hash)

Get a Tx by hash

Retrieve a transaction using its hash.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string**| Tx hash | 

### Return type

[**TxQuery**](TxQuery.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TxsPost

> BroadcastTxCommitResult TxsPost(ctx, txBroadcast)

Broadcast a signed tx

Broadcast a signed tx to a full node

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**txBroadcast** | [**InlineObject**](InlineObject.md)|  | 

### Return type

[**BroadcastTxCommitResult**](BroadcastTxCommitResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

