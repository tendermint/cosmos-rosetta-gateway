# \SupplyApi

All URIs are relative to *https://stargate.cosmos.network*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SupplyTotalDenominationGet**](SupplyApi.md#SupplyTotalDenominationGet) | **Get** /supply/total/{denomination} | Total supply of a single coin denomination
[**SupplyTotalGet**](SupplyApi.md#SupplyTotalGet) | **Get** /supply/total | Total supply of coins in the chain



## SupplyTotalDenominationGet

> string SupplyTotalDenominationGet(ctx, denomination)

Total supply of a single coin denomination

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**denomination** | **string**| Coin denomination | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SupplyTotalGet

> Supply SupplyTotalGet(ctx, )

Total supply of coins in the chain

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**Supply**](Supply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

