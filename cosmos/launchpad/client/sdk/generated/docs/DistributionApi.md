# \DistributionApi

All URIs are relative to *https://stargate.cosmos.network*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DistributionCommunityPoolGet**](DistributionApi.md#DistributionCommunityPoolGet) | **Get** /distribution/community_pool | Community pool parameters
[**DistributionDelegatorsDelegatorAddrRewardsGet**](DistributionApi.md#DistributionDelegatorsDelegatorAddrRewardsGet) | **Get** /distribution/delegators/{delegatorAddr}/rewards | Get the total rewards balance from all delegations
[**DistributionDelegatorsDelegatorAddrRewardsPost**](DistributionApi.md#DistributionDelegatorsDelegatorAddrRewardsPost) | **Post** /distribution/delegators/{delegatorAddr}/rewards | Withdraw all the delegator&#39;s delegation rewards
[**DistributionDelegatorsDelegatorAddrRewardsValidatorAddrGet**](DistributionApi.md#DistributionDelegatorsDelegatorAddrRewardsValidatorAddrGet) | **Get** /distribution/delegators/{delegatorAddr}/rewards/{validatorAddr} | Query a delegation reward
[**DistributionDelegatorsDelegatorAddrRewardsValidatorAddrPost**](DistributionApi.md#DistributionDelegatorsDelegatorAddrRewardsValidatorAddrPost) | **Post** /distribution/delegators/{delegatorAddr}/rewards/{validatorAddr} | Withdraw a delegation reward
[**DistributionDelegatorsDelegatorAddrWithdrawAddressGet**](DistributionApi.md#DistributionDelegatorsDelegatorAddrWithdrawAddressGet) | **Get** /distribution/delegators/{delegatorAddr}/withdraw_address | Get the rewards withdrawal address
[**DistributionDelegatorsDelegatorAddrWithdrawAddressPost**](DistributionApi.md#DistributionDelegatorsDelegatorAddrWithdrawAddressPost) | **Post** /distribution/delegators/{delegatorAddr}/withdraw_address | Replace the rewards withdrawal address
[**DistributionParametersGet**](DistributionApi.md#DistributionParametersGet) | **Get** /distribution/parameters | Fee distribution parameters
[**DistributionValidatorsValidatorAddrGet**](DistributionApi.md#DistributionValidatorsValidatorAddrGet) | **Get** /distribution/validators/{validatorAddr} | Validator distribution information
[**DistributionValidatorsValidatorAddrOutstandingRewardsGet**](DistributionApi.md#DistributionValidatorsValidatorAddrOutstandingRewardsGet) | **Get** /distribution/validators/{validatorAddr}/outstanding_rewards | Fee distribution outstanding rewards of a single validator
[**DistributionValidatorsValidatorAddrRewardsGet**](DistributionApi.md#DistributionValidatorsValidatorAddrRewardsGet) | **Get** /distribution/validators/{validatorAddr}/rewards | Commission and self-delegation rewards of a single validator
[**DistributionValidatorsValidatorAddrRewardsPost**](DistributionApi.md#DistributionValidatorsValidatorAddrRewardsPost) | **Post** /distribution/validators/{validatorAddr}/rewards | Withdraw the validator&#39;s rewards



## DistributionCommunityPoolGet

> []Coin DistributionCommunityPoolGet(ctx, )

Community pool parameters

### Required Parameters

This endpoint does not need any parameter.

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


## DistributionDelegatorsDelegatorAddrRewardsGet

> DelegatorTotalRewards DistributionDelegatorsDelegatorAddrRewardsGet(ctx, delegatorAddr)

Get the total rewards balance from all delegations

Get the sum of all the rewards earned by delegations by a single delegator

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**delegatorAddr** | **string**| Bech32 AccAddress of Delegator | 

### Return type

[**DelegatorTotalRewards**](DelegatorTotalRewards.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DistributionDelegatorsDelegatorAddrRewardsPost

> BroadcastTxCommitResult DistributionDelegatorsDelegatorAddrRewardsPost(ctx, delegatorAddr, optional)

Withdraw all the delegator's delegation rewards

Withdraw all the delegator's delegation rewards

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**delegatorAddr** | **string**| Bech32 AccAddress of Delegator | 
 **optional** | ***DistributionDelegatorsDelegatorAddrRewardsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DistributionDelegatorsDelegatorAddrRewardsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withdrawRequestBody** | [**optional.Interface of InlineObject12**](InlineObject12.md)|  | 

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


## DistributionDelegatorsDelegatorAddrRewardsValidatorAddrGet

> []Coin DistributionDelegatorsDelegatorAddrRewardsValidatorAddrGet(ctx, delegatorAddr, validatorAddr)

Query a delegation reward

Query a single delegation reward by a delegator

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**delegatorAddr** | **string**| Bech32 AccAddress of Delegator | 
**validatorAddr** | **string**| Bech32 OperatorAddress of validator | 

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


## DistributionDelegatorsDelegatorAddrRewardsValidatorAddrPost

> BroadcastTxCommitResult DistributionDelegatorsDelegatorAddrRewardsValidatorAddrPost(ctx, delegatorAddr, validatorAddr, optional)

Withdraw a delegation reward

Withdraw a delegator's delegation reward from a single validator

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**delegatorAddr** | **string**| Bech32 AccAddress of Delegator | 
**validatorAddr** | **string**| Bech32 OperatorAddress of validator | 
 **optional** | ***DistributionDelegatorsDelegatorAddrRewardsValidatorAddrPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DistributionDelegatorsDelegatorAddrRewardsValidatorAddrPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **withdrawRequestBody** | [**optional.Interface of InlineObject13**](InlineObject13.md)|  | 

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


## DistributionDelegatorsDelegatorAddrWithdrawAddressGet

> string DistributionDelegatorsDelegatorAddrWithdrawAddressGet(ctx, delegatorAddr)

Get the rewards withdrawal address

Get the delegations' rewards withdrawal address. This is the address in which the user will receive the reward funds

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**delegatorAddr** | **string**| Bech32 AccAddress of Delegator | 

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


## DistributionDelegatorsDelegatorAddrWithdrawAddressPost

> BroadcastTxCommitResult DistributionDelegatorsDelegatorAddrWithdrawAddressPost(ctx, delegatorAddr, optional)

Replace the rewards withdrawal address

Replace the delegations' rewards withdrawal address for a new one.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**delegatorAddr** | **string**| Bech32 AccAddress of Delegator | 
 **optional** | ***DistributionDelegatorsDelegatorAddrWithdrawAddressPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DistributionDelegatorsDelegatorAddrWithdrawAddressPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withdrawRequestBody** | [**optional.Interface of InlineObject14**](InlineObject14.md)|  | 

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


## DistributionParametersGet

> map[string]interface{} DistributionParametersGet(ctx, )

Fee distribution parameters

### Required Parameters

This endpoint does not need any parameter.

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DistributionValidatorsValidatorAddrGet

> ValidatorDistInfo DistributionValidatorsValidatorAddrGet(ctx, validatorAddr)

Validator distribution information

Query the distribution information of a single validator

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**validatorAddr** | **string**| Bech32 OperatorAddress of validator | 

### Return type

[**ValidatorDistInfo**](ValidatorDistInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DistributionValidatorsValidatorAddrOutstandingRewardsGet

> []Coin DistributionValidatorsValidatorAddrOutstandingRewardsGet(ctx, validatorAddr)

Fee distribution outstanding rewards of a single validator

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**validatorAddr** | **string**| Bech32 OperatorAddress of validator | 

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


## DistributionValidatorsValidatorAddrRewardsGet

> []Coin DistributionValidatorsValidatorAddrRewardsGet(ctx, validatorAddr)

Commission and self-delegation rewards of a single validator

Query the commission and self-delegation rewards of validator.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**validatorAddr** | **string**| Bech32 OperatorAddress of validator | 

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


## DistributionValidatorsValidatorAddrRewardsPost

> BroadcastTxCommitResult DistributionValidatorsValidatorAddrRewardsPost(ctx, validatorAddr, optional)

Withdraw the validator's rewards

Withdraw the validator's self-delegation and commissions rewards

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**validatorAddr** | **string**| Bech32 OperatorAddress of validator | 
 **optional** | ***DistributionValidatorsValidatorAddrRewardsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DistributionValidatorsValidatorAddrRewardsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withdrawRequestBody** | [**optional.Interface of InlineObject15**](InlineObject15.md)|  | 

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

