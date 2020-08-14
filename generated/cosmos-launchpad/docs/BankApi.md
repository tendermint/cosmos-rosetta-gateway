# \BankApi

All URIs are relative to *https://stargate.cosmos.network*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BankAccountsAddressTransfersPost**](BankApi.md#BankAccountsAddressTransfersPost) | **Post** /bank/accounts/{address}/transfers | Send coins from one account to another
[**BankBalancesAddressGet**](BankApi.md#BankBalancesAddressGet) | **Get** /bank/balances/{address} | Get the account balances



## BankAccountsAddressTransfersPost

> StdTx BankAccountsAddressTransfersPost(ctx, address, account)

Send coins from one account to another

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string**| Account address in bech32 format | 
**account** | [**InlineObject3**](InlineObject3.md)|  | 

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


## BankBalancesAddressGet

> []Coin BankBalancesAddressGet(ctx, address)

Get the account balances

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string**| Account address in bech32 format | 

### Return type

[**[]Coin**](Coin.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

