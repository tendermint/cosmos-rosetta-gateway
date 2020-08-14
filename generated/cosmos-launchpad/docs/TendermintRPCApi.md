# \TendermintRPCApi

All URIs are relative to *https://stargate.cosmos.network*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BlocksHeightGet**](TendermintRPCApi.md#BlocksHeightGet) | **Get** /blocks/{height} | Get a block at a certain height
[**BlocksLatestGet**](TendermintRPCApi.md#BlocksLatestGet) | **Get** /blocks/latest | Get the latest block
[**NodeInfoGet**](TendermintRPCApi.md#NodeInfoGet) | **Get** /node_info | The properties of the connected node
[**SyncingGet**](TendermintRPCApi.md#SyncingGet) | **Get** /syncing | Syncing state of node
[**ValidatorsetsHeightGet**](TendermintRPCApi.md#ValidatorsetsHeightGet) | **Get** /validatorsets/{height} | Get a validator set a certain height
[**ValidatorsetsLatestGet**](TendermintRPCApi.md#ValidatorsetsLatestGet) | **Get** /validatorsets/latest | Get the latest validator set



## BlocksHeightGet

> BlockQuery BlocksHeightGet(ctx, height)

Get a block at a certain height

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**height** | **float32**| Block height | 

### Return type

[**BlockQuery**](BlockQuery.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BlocksLatestGet

> BlockQuery BlocksLatestGet(ctx, )

Get the latest block

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**BlockQuery**](BlockQuery.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NodeInfoGet

> InlineResponse200 NodeInfoGet(ctx, )

The properties of the connected node

Information about the connected node

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncingGet

> InlineResponse2001 SyncingGet(ctx, )

Syncing state of node

Get if the node is currently syning with other nodes

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidatorsetsHeightGet

> InlineResponse2002 ValidatorsetsHeightGet(ctx, height)

Get a validator set a certain height

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**height** | **float32**| Block height | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidatorsetsLatestGet

> InlineResponse2002 ValidatorsetsLatestGet(ctx, )

Get the latest validator set

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

