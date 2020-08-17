# \TxApi

All URIs are relative to *https://rpc.cosmos.network*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BroadcastTxAsync**](TxApi.md#BroadcastTxAsync) | **Get** /broadcast_tx_async | Returns right away, with no response. Does not wait for CheckTx nor DeliverTx results.
[**BroadcastTxCommit**](TxApi.md#BroadcastTxCommit) | **Get** /broadcast_tx_commit | Returns with the responses from CheckTx and DeliverTx.
[**BroadcastTxSync**](TxApi.md#BroadcastTxSync) | **Get** /broadcast_tx_sync | Returns with the response from CheckTx. Does not wait for DeliverTx result.
[**CheckTx**](TxApi.md#CheckTx) | **Get** /check_tx | Checks the transaction without executing it.



## BroadcastTxAsync

> BroadcastTxResponse BroadcastTxAsync(ctx, tx)

Returns right away, with no response. Does not wait for CheckTx nor DeliverTx results.

If you want to be sure that the transaction is included in a block, you can subscribe for the result using JSONRPC via a websocket. See https://docs.tendermint.com/master/app-dev/subscribing-to-events-via-websocket.html If you haven't received anything after a couple of blocks, resend it. If the same happens again, send it to some other node. A few reasons why it could happen:  1. malicious node can drop or pretend it had committed your tx 2. malicious proposer (not necessary the one you're communicating with) can drop transactions, which might become valid in the future (https://github.com/tendermint/tendermint/issues/3322) 3. node can be offline  Please refer to https://docs.tendermint.com/master/tendermint-core/using-tendermint.html#formatting for formatting/encoding rules. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tx** | **string**| The transaction | 

### Return type

[**BroadcastTxResponse**](BroadcastTxResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BroadcastTxCommit

> BroadcastTxCommitResponse BroadcastTxCommit(ctx, tx)

Returns with the responses from CheckTx and DeliverTx.

IMPORTANT: use only for testing and development. In production, use BroadcastTxSync or BroadcastTxAsync. You can subscribe for the transaction result using JSONRPC via a websocket. See https://docs.tendermint.com/master/app-dev/subscribing-to-events-via-websocket.html  CONTRACT: only returns error if mempool.CheckTx() errs or if we timeout waiting for tx to commit.  If CheckTx or DeliverTx fail, no error will be returned, but the returned result will contain a non-OK ABCI code.  Please refer to https://docs.tendermint.com/master/tendermint-core/using-tendermint.html#formatting for formatting/encoding rules. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tx** | **string**| The transaction | 

### Return type

[**BroadcastTxCommitResponse**](BroadcastTxCommitResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BroadcastTxSync

> BroadcastTxResponse BroadcastTxSync(ctx, tx)

Returns with the response from CheckTx. Does not wait for DeliverTx result.

If you want to be sure that the transaction is included in a block, you can subscribe for the result using JSONRPC via a websocket. See https://docs.tendermint.com/master/app-dev/subscribing-to-events-via-websocket.html If you haven't received anything after a couple of blocks, resend it. If the same happens again, send it to some other node. A few reasons why it could happen:  1. malicious node can drop or pretend it had committed your tx 2. malicious proposer (not necessary the one you're communicating with) can drop transactions, which might become valid in the future (https://github.com/tendermint/tendermint/issues/3322)   Please refer to https://docs.tendermint.com/master/tendermint-core/using-tendermint.html#formatting for formatting/encoding rules. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tx** | **string**| The transaction | 

### Return type

[**BroadcastTxResponse**](BroadcastTxResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckTx

> CheckTxResponse CheckTx(ctx, tx)

Checks the transaction without executing it.

The transaction won't be added to the mempool.  Please refer to https://docs.tendermint.com/master/tendermint-core/using-tendermint.html#formatting for formatting/encoding rules. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tx** | **string**| The transaction | 

### Return type

[**CheckTxResponse**](CheckTxResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

