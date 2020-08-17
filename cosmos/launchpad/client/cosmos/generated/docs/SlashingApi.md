# \SlashingApi

All URIs are relative to *https://stargate.cosmos.network*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SlashingParametersGet**](SlashingApi.md#SlashingParametersGet) | **Get** /slashing/parameters | Get the current slashing parameters
[**SlashingSigningInfosGet**](SlashingApi.md#SlashingSigningInfosGet) | **Get** /slashing/signing_infos | Get sign info of given all validators
[**SlashingValidatorsValidatorAddrUnjailPost**](SlashingApi.md#SlashingValidatorsValidatorAddrUnjailPost) | **Post** /slashing/validators/{validatorAddr}/unjail | Unjail a jailed validator



## SlashingParametersGet

> InlineResponse2008 SlashingParametersGet(ctx, )

Get the current slashing parameters

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlashingSigningInfosGet

> []SigningInfo SlashingSigningInfosGet(ctx, page, limit)

Get sign info of given all validators

Get sign info of all validators

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**page** | **int32**| Page number | 
**limit** | **int32**| Maximum number of items per page | 

### Return type

[**[]SigningInfo**](SigningInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlashingValidatorsValidatorAddrUnjailPost

> BroadcastTxCommitResult SlashingValidatorsValidatorAddrUnjailPost(ctx, validatorAddr, unjailBody)

Unjail a jailed validator

Send transaction to unjail a jailed validator

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**validatorAddr** | **string**| Bech32 validator address | 
**unjailBody** | [**InlineObject7**](InlineObject7.md)|  | 

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

