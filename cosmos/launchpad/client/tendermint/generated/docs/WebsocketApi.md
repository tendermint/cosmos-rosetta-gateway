# \WebsocketApi

All URIs are relative to *https://rpc.cosmos.network*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Subscribe**](WebsocketApi.md#Subscribe) | **Get** /subscribe | Subscribe for events via WebSocket.
[**Unsubscribe**](WebsocketApi.md#Unsubscribe) | **Get** /unsubscribe | Unsubscribe from event on Websocket
[**UnsubscribeAll**](WebsocketApi.md#UnsubscribeAll) | **Get** /unsubscribe_all | Unsubscribe from all events via WebSocket



## Subscribe

> EmptyResponse Subscribe(ctx, query)

Subscribe for events via WebSocket.

To tell which events you want, you need to provide a query. query is a string, which has a form: \"condition AND condition ...\" (no OR at the moment). condition has a form: \"key operation operand\". key is a string with a restricted set of possible symbols ( \\t\\n\\r\\\\()\"'=>< are not allowed). operation can be \"=\", \"<\", \"<=\", \">\", \">=\", \"CONTAINS\" AND \"EXISTS\". operand can be a string (escaped with single quotes), number, date or time.  Examples:       tm.event = 'NewBlock'               # new blocks       tm.event = 'CompleteProposal'       # node got a complete proposal       tm.event = 'Tx' AND tx.hash = 'XYZ' # single transaction       tm.event = 'Tx' AND tx.height = 5   # all txs of the fifth block       tx.height = 5                       # all txs of the fifth block  Tendermint provides a few predefined keys: tm.event, tx.hash and tx.height. Note for transactions, you can define additional keys by providing events with DeliverTx response.  import (     abci \"github.com/tendermint/tendermint/abci/types\"     \"github.com/tendermint/tendermint/libs/pubsub/query\" )  abci.ResponseDeliverTx{   Events: []abci.Event{       {           Type: \"rewards.withdraw\",           Attributes: kv.Pairs{               kv.Pair{Key: []byte(\"address\"), Value: []byte(\"AddrA\")},               kv.Pair{Key: []byte(\"source\"), Value: []byte(\"SrcX\")},               kv.Pair{Key: []byte(\"amount\"), Value: []byte(\"...\")},               kv.Pair{Key: []byte(\"balance\"), Value: []byte(\"...\")},           },       },       {           Type: \"rewards.withdraw\",           Attributes: kv.Pairs{               kv.Pair{Key: []byte(\"address\"), Value: []byte(\"AddrB\")},               kv.Pair{Key: []byte(\"source\"), Value: []byte(\"SrcY\")},               kv.Pair{Key: []byte(\"amount\"), Value: []byte(\"...\")},               kv.Pair{Key: []byte(\"balance\"), Value: []byte(\"...\")},           },       },       {           Type: \"transfer\",           Attributes: kv.Pairs{               kv.Pair{Key: []byte(\"sender\"), Value: []byte(\"AddrC\")},               kv.Pair{Key: []byte(\"recipient\"), Value: []byte(\"AddrD\")},               kv.Pair{Key: []byte(\"amount\"), Value: []byte(\"...\")},           },       },   }, }  All events are indexed by a composite key of the form {eventType}.{evenAttrKey}. In the above examples, the following keys would be indexed:    - rewards.withdraw.address    - rewards.withdraw.source    - rewards.withdraw.amount    - rewards.withdraw.balance    - transfer.sender    - transfer.recipient    - transfer.amount  Multiple event types with duplicate keys are allowed and are meant to categorize unique and distinct events. In the above example, all events indexed under the key `rewards.withdraw.address` will have the following values stored and queryable:     - AddrA    - AddrB  To create a query for txs where address AddrA withdrew rewards: query.MustParse(\"tm.event = 'Tx' AND rewards.withdraw.address = 'AddrA'\")  To create a query for txs where address AddrA withdrew rewards from source Y: query.MustParse(\"tm.event = 'Tx' AND rewards.withdraw.address = 'AddrA' AND rewards.withdraw.source = 'Y'\")  To create a query for txs where AddrA transferred funds: query.MustParse(\"tm.event = 'Tx' AND transfer.sender = 'AddrA'\")  The following queries would return no results: query.MustParse(\"tm.event = 'Tx' AND transfer.sender = 'AddrZ'\") query.MustParse(\"tm.event = 'Tx' AND rewards.withdraw.address = 'AddrZ'\") query.MustParse(\"tm.event = 'Tx' AND rewards.withdraw.source = 'W'\")  See list of all possible events here https://godoc.org/github.com/tendermint/tendermint/types#pkg-constants  For complete query syntax, check out https://godoc.org/github.com/tendermint/tendermint/libs/pubsub/query.  ```go import rpchttp \"github.com/tendermint/rpc/client/http\" import \"github.com/tendermint/tendermint/types\"  client := rpchttp.New(\"tcp:0.0.0.0:26657\", \"/websocket\") err := client.Start() if err != nil {   handle error } defer client.Stop() ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second) defer cancel() query := \"tm.event = 'Tx' AND tx.height = 3\" txs, err := client.Subscribe(ctx, \"test-client\", query) if err != nil {   handle error }  go func() {  for e := range txs {    fmt.Println(\"got \", e.Data.(types.EventDataTx))    } }() ```  NOTE: if you're not reading events fast enough, Tendermint might terminate the subscription. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**query** | **string**| query is a string, which has a form: \&quot;condition AND condition ...\&quot; (no OR at the moment). condition has a form: \&quot;key operation operand\&quot;. key is a string with a restricted set of possible symbols ( \\t\\n\\r\\\\()\&quot;&#39;&#x3D;&gt;&lt; are not allowed). operation can be \&quot;&#x3D;\&quot;, \&quot;&lt;\&quot;, \&quot;&lt;&#x3D;\&quot;, \&quot;&gt;\&quot;, \&quot;&gt;&#x3D;\&quot;, \&quot;CONTAINS\&quot;. operand can be a string (escaped with single quotes), number, date or time.  | 

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


## Unsubscribe

> EmptyResponse Unsubscribe(ctx, query)

Unsubscribe from event on Websocket

```go client := rpchttp.New(\"tcp:0.0.0.0:26657\", \"/websocket\") err := client.Start() if err != nil {    handle error } defer client.Stop() query := \"tm.event = 'Tx' AND tx.height = 3\" err = client.Unsubscribe(context.Background(), \"test-client\", query) if err != nil {    handle error } ``` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**query** | **string**| query is a string, which has a form: \&quot;condition AND condition ...\&quot; (no OR at the moment). condition has a form: \&quot;key operation operand\&quot;. key is a string with a restricted set of possible symbols ( \\t\\n\\r\\\\()\&quot;&#39;&#x3D;&gt;&lt; are not allowed). operation can be \&quot;&#x3D;\&quot;, \&quot;&lt;\&quot;, \&quot;&lt;&#x3D;\&quot;, \&quot;&gt;\&quot;, \&quot;&gt;&#x3D;\&quot;, \&quot;CONTAINS\&quot;. operand can be a string (escaped with single quotes), number, date or time.  | 

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


## UnsubscribeAll

> EmptyResponse UnsubscribeAll(ctx, )

Unsubscribe from all events via WebSocket

Unsubscribe from all events via WebSocket 

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

