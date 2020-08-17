# \AuthApi

All URIs are relative to *https://stargate.cosmos.network*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthAccountsAddressGet**](AuthApi.md#AuthAccountsAddressGet) | **Get** /auth/accounts/{address} | Get the account information on blockchain



## AuthAccountsAddressGet

> InlineResponse2005 AuthAccountsAddressGet(ctx, address)

Get the account information on blockchain

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string**| Account address | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

