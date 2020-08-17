# \UnsafeApi

All URIs are relative to *https://rpc.cosmos.network*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DialPeers**](UnsafeApi.md#DialPeers) | **Post** /dial_peers | Add Peers/Persistent Peers (unsafe)
[**DialSeeds**](UnsafeApi.md#DialSeeds) | **Post** /dial_seeds | Dial Seeds (Unsafe)



## DialPeers

> DialResp DialPeers(ctx, dialPeersPost)

Add Peers/Persistent Peers (unsafe)

Set a persistent peer, this route in under unsafe, and has to manually enabled to use 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dialPeersPost** | [**DialPeersPost**](DialPeersPost.md)| string of possible peers, bool argument if they should be added as persistent | 

### Return type

[**DialResp**](dialResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DialSeeds

> DialResp DialSeeds(ctx, dialSeedsPost)

Dial Seeds (Unsafe)

Dial a peer, this route in under unsafe, and has to manually enabled to use 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dialSeedsPost** | [**DialSeedsPost**](DialSeedsPost.md)| string of possible peers | 

### Return type

[**DialResp**](dialResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

