# \StakingApi

All URIs are relative to *https://stargate.cosmos.network*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StakingDelegatorsDelegatorAddrDelegationsGet**](StakingApi.md#StakingDelegatorsDelegatorAddrDelegationsGet) | **Get** /staking/delegators/{delegatorAddr}/delegations | Get all delegations from a delegator
[**StakingDelegatorsDelegatorAddrDelegationsPost**](StakingApi.md#StakingDelegatorsDelegatorAddrDelegationsPost) | **Post** /staking/delegators/{delegatorAddr}/delegations | Submit delegation
[**StakingDelegatorsDelegatorAddrDelegationsValidatorAddrGet**](StakingApi.md#StakingDelegatorsDelegatorAddrDelegationsValidatorAddrGet) | **Get** /staking/delegators/{delegatorAddr}/delegations/{validatorAddr} | Query the current delegation between a delegator and a validator
[**StakingDelegatorsDelegatorAddrRedelegationsPost**](StakingApi.md#StakingDelegatorsDelegatorAddrRedelegationsPost) | **Post** /staking/delegators/{delegatorAddr}/redelegations | Submit a redelegation
[**StakingDelegatorsDelegatorAddrUnbondingDelegationsGet**](StakingApi.md#StakingDelegatorsDelegatorAddrUnbondingDelegationsGet) | **Get** /staking/delegators/{delegatorAddr}/unbonding_delegations | Get all unbonding delegations from a delegator
[**StakingDelegatorsDelegatorAddrUnbondingDelegationsPost**](StakingApi.md#StakingDelegatorsDelegatorAddrUnbondingDelegationsPost) | **Post** /staking/delegators/{delegatorAddr}/unbonding_delegations | Submit an unbonding delegation
[**StakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrGet**](StakingApi.md#StakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrGet) | **Get** /staking/delegators/{delegatorAddr}/unbonding_delegations/{validatorAddr} | Query all unbonding delegations between a delegator and a validator
[**StakingDelegatorsDelegatorAddrValidatorsGet**](StakingApi.md#StakingDelegatorsDelegatorAddrValidatorsGet) | **Get** /staking/delegators/{delegatorAddr}/validators | Query all validators that a delegator is bonded to
[**StakingDelegatorsDelegatorAddrValidatorsValidatorAddrGet**](StakingApi.md#StakingDelegatorsDelegatorAddrValidatorsValidatorAddrGet) | **Get** /staking/delegators/{delegatorAddr}/validators/{validatorAddr} | Query a validator that a delegator is bonded to
[**StakingParametersGet**](StakingApi.md#StakingParametersGet) | **Get** /staking/parameters | Get the current staking parameter values
[**StakingPoolGet**](StakingApi.md#StakingPoolGet) | **Get** /staking/pool | Get the current state of the staking pool
[**StakingRedelegationsGet**](StakingApi.md#StakingRedelegationsGet) | **Get** /staking/redelegations | Get all redelegations (filter by query params)
[**StakingValidatorsGet**](StakingApi.md#StakingValidatorsGet) | **Get** /staking/validators | Get all validator candidates. By default it returns only the bonded validators.
[**StakingValidatorsValidatorAddrDelegationsGet**](StakingApi.md#StakingValidatorsValidatorAddrDelegationsGet) | **Get** /staking/validators/{validatorAddr}/delegations | Get all delegations from a validator
[**StakingValidatorsValidatorAddrGet**](StakingApi.md#StakingValidatorsValidatorAddrGet) | **Get** /staking/validators/{validatorAddr} | Query the information from a single validator
[**StakingValidatorsValidatorAddrUnbondingDelegationsGet**](StakingApi.md#StakingValidatorsValidatorAddrUnbondingDelegationsGet) | **Get** /staking/validators/{validatorAddr}/unbonding_delegations | Get all unbonding delegations from a validator



## StakingDelegatorsDelegatorAddrDelegationsGet

> []Delegation StakingDelegatorsDelegatorAddrDelegationsGet(ctx, delegatorAddr)

Get all delegations from a delegator

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**delegatorAddr** | **string**| Bech32 AccAddress of Delegator | 

### Return type

[**[]Delegation**](Delegation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StakingDelegatorsDelegatorAddrDelegationsPost

> BroadcastTxCommitResult StakingDelegatorsDelegatorAddrDelegationsPost(ctx, delegatorAddr, optional)

Submit delegation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**delegatorAddr** | **string**| Bech32 AccAddress of Delegator | 
 **optional** | ***StakingDelegatorsDelegatorAddrDelegationsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StakingDelegatorsDelegatorAddrDelegationsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **delegation** | [**optional.Interface of InlineObject4**](InlineObject4.md)|  | 

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


## StakingDelegatorsDelegatorAddrDelegationsValidatorAddrGet

> Delegation StakingDelegatorsDelegatorAddrDelegationsValidatorAddrGet(ctx, delegatorAddr, validatorAddr)

Query the current delegation between a delegator and a validator

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**delegatorAddr** | **string**| Bech32 AccAddress of Delegator | 
**validatorAddr** | **string**| Bech32 OperatorAddress of validator | 

### Return type

[**Delegation**](Delegation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StakingDelegatorsDelegatorAddrRedelegationsPost

> StdTx StakingDelegatorsDelegatorAddrRedelegationsPost(ctx, delegatorAddr, optional)

Submit a redelegation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**delegatorAddr** | **string**| Bech32 AccAddress of Delegator | 
 **optional** | ***StakingDelegatorsDelegatorAddrRedelegationsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StakingDelegatorsDelegatorAddrRedelegationsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **delegation** | [**optional.Interface of InlineObject6**](InlineObject6.md)|  | 

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


## StakingDelegatorsDelegatorAddrUnbondingDelegationsGet

> []UnbondingDelegation StakingDelegatorsDelegatorAddrUnbondingDelegationsGet(ctx, delegatorAddr)

Get all unbonding delegations from a delegator

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**delegatorAddr** | **string**| Bech32 AccAddress of Delegator | 

### Return type

[**[]UnbondingDelegation**](UnbondingDelegation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StakingDelegatorsDelegatorAddrUnbondingDelegationsPost

> BroadcastTxCommitResult StakingDelegatorsDelegatorAddrUnbondingDelegationsPost(ctx, delegatorAddr, optional)

Submit an unbonding delegation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**delegatorAddr** | **string**| Bech32 AccAddress of Delegator | 
 **optional** | ***StakingDelegatorsDelegatorAddrUnbondingDelegationsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StakingDelegatorsDelegatorAddrUnbondingDelegationsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **delegation** | [**optional.Interface of InlineObject5**](InlineObject5.md)|  | 

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


## StakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrGet

> UnbondingDelegationPair StakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrGet(ctx, delegatorAddr, validatorAddr)

Query all unbonding delegations between a delegator and a validator

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**delegatorAddr** | **string**| Bech32 AccAddress of Delegator | 
**validatorAddr** | **string**| Bech32 OperatorAddress of validator | 

### Return type

[**UnbondingDelegationPair**](UnbondingDelegationPair.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StakingDelegatorsDelegatorAddrValidatorsGet

> []Validator StakingDelegatorsDelegatorAddrValidatorsGet(ctx, delegatorAddr)

Query all validators that a delegator is bonded to

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**delegatorAddr** | **string**| Bech32 AccAddress of Delegator | 

### Return type

[**[]Validator**](Validator.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StakingDelegatorsDelegatorAddrValidatorsValidatorAddrGet

> Validator StakingDelegatorsDelegatorAddrValidatorsValidatorAddrGet(ctx, delegatorAddr, validatorAddr)

Query a validator that a delegator is bonded to

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**delegatorAddr** | **string**| Bech32 AccAddress of Delegator | 
**validatorAddr** | **string**| Bech32 ValAddress of Delegator | 

### Return type

[**Validator**](Validator.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StakingParametersGet

> InlineResponse2008 StakingParametersGet(ctx, )

Get the current staking parameter values

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


## StakingPoolGet

> InlineResponse2007 StakingPoolGet(ctx, )

Get the current state of the staking pool

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StakingRedelegationsGet

> []Redelegation StakingRedelegationsGet(ctx, optional)

Get all redelegations (filter by query params)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StakingRedelegationsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StakingRedelegationsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **delegator** | **optional.String**| Bech32 AccAddress of Delegator | 
 **validatorFrom** | **optional.String**| Bech32 ValAddress of SrcValidator | 
 **validatorTo** | **optional.String**| Bech32 ValAddress of DstValidator | 

### Return type

[**[]Redelegation**](Redelegation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StakingValidatorsGet

> []Validator StakingValidatorsGet(ctx, optional)

Get all validator candidates. By default it returns only the bonded validators.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StakingValidatorsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StakingValidatorsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **optional.String**| The validator bond status. Must be either &#39;bonded&#39;, &#39;unbonded&#39;, or &#39;unbonding&#39;. | 
 **page** | **optional.Int32**| The page number. | 
 **limit** | **optional.Int32**| The maximum number of items per page. | 

### Return type

[**[]Validator**](Validator.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StakingValidatorsValidatorAddrDelegationsGet

> []Delegation StakingValidatorsValidatorAddrDelegationsGet(ctx, validatorAddr)

Get all delegations from a validator

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**validatorAddr** | **string**| Bech32 OperatorAddress of validator | 

### Return type

[**[]Delegation**](Delegation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StakingValidatorsValidatorAddrGet

> Validator StakingValidatorsValidatorAddrGet(ctx, validatorAddr)

Query the information from a single validator

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**validatorAddr** | **string**| Bech32 OperatorAddress of validator | 

### Return type

[**Validator**](Validator.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StakingValidatorsValidatorAddrUnbondingDelegationsGet

> []UnbondingDelegation StakingValidatorsValidatorAddrUnbondingDelegationsGet(ctx, validatorAddr)

Get all unbonding delegations from a validator

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**validatorAddr** | **string**| Bech32 OperatorAddress of validator | 

### Return type

[**[]UnbondingDelegation**](UnbondingDelegation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

