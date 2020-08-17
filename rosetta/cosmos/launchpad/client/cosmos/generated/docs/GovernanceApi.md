# \GovernanceApi

All URIs are relative to *https://stargate.cosmos.network*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GovParametersDepositGet**](GovernanceApi.md#GovParametersDepositGet) | **Get** /gov/parameters/deposit | Query governance deposit parameters
[**GovParametersTallyingGet**](GovernanceApi.md#GovParametersTallyingGet) | **Get** /gov/parameters/tallying | Query governance tally parameters
[**GovParametersVotingGet**](GovernanceApi.md#GovParametersVotingGet) | **Get** /gov/parameters/voting | Query governance voting parameters
[**GovProposalsGet**](GovernanceApi.md#GovProposalsGet) | **Get** /gov/proposals | Query proposals
[**GovProposalsParamChangePost**](GovernanceApi.md#GovProposalsParamChangePost) | **Post** /gov/proposals/param_change | Generate a parameter change proposal transaction
[**GovProposalsPost**](GovernanceApi.md#GovProposalsPost) | **Post** /gov/proposals | Submit a proposal
[**GovProposalsProposalIdDepositsDepositorGet**](GovernanceApi.md#GovProposalsProposalIdDepositsDepositorGet) | **Get** /gov/proposals/{proposalId}/deposits/{depositor} | Query deposit
[**GovProposalsProposalIdDepositsGet**](GovernanceApi.md#GovProposalsProposalIdDepositsGet) | **Get** /gov/proposals/{proposalId}/deposits | Query deposits
[**GovProposalsProposalIdDepositsPost**](GovernanceApi.md#GovProposalsProposalIdDepositsPost) | **Post** /gov/proposals/{proposalId}/deposits | Deposit tokens to a proposal
[**GovProposalsProposalIdGet**](GovernanceApi.md#GovProposalsProposalIdGet) | **Get** /gov/proposals/{proposalId} | Query a proposal
[**GovProposalsProposalIdProposerGet**](GovernanceApi.md#GovProposalsProposalIdProposerGet) | **Get** /gov/proposals/{proposalId}/proposer | Query proposer
[**GovProposalsProposalIdTallyGet**](GovernanceApi.md#GovProposalsProposalIdTallyGet) | **Get** /gov/proposals/{proposalId}/tally | Get a proposal&#39;s tally result at the current time
[**GovProposalsProposalIdVotesGet**](GovernanceApi.md#GovProposalsProposalIdVotesGet) | **Get** /gov/proposals/{proposalId}/votes | Query voters
[**GovProposalsProposalIdVotesPost**](GovernanceApi.md#GovProposalsProposalIdVotesPost) | **Post** /gov/proposals/{proposalId}/votes | Vote a proposal
[**GovProposalsProposalIdVotesVoterGet**](GovernanceApi.md#GovProposalsProposalIdVotesVoterGet) | **Get** /gov/proposals/{proposalId}/votes/{voter} | Query vote



## GovParametersDepositGet

> InlineResponse2009 GovParametersDepositGet(ctx, )

Query governance deposit parameters

Query governance deposit parameters. The max_deposit_period units are in nanoseconds.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GovParametersTallyingGet

> map[string]interface{} GovParametersTallyingGet(ctx, )

Query governance tally parameters

Query governance tally parameters

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


## GovParametersVotingGet

> map[string]interface{} GovParametersVotingGet(ctx, )

Query governance voting parameters

Query governance voting parameters. The voting_period units are in nanoseconds.

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


## GovProposalsGet

> []TextProposal GovProposalsGet(ctx, optional)

Query proposals

Query proposals information with parameters

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GovProposalsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GovProposalsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **voter** | **optional.String**| voter address | 
 **depositor** | **optional.String**| depositor address | 
 **status** | **optional.String**| proposal status, valid values can be &#x60;\&quot;deposit_period\&quot;&#x60;, &#x60;\&quot;voting_period\&quot;&#x60;, &#x60;\&quot;passed\&quot;&#x60;, &#x60;\&quot;rejected\&quot;&#x60; | 

### Return type

[**[]TextProposal**](TextProposal.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GovProposalsParamChangePost

> StdTx GovProposalsParamChangePost(ctx, postProposalBody)

Generate a parameter change proposal transaction

Generate a parameter change proposal transaction

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postProposalBody** | [**InlineObject9**](InlineObject9.md)|  | 

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


## GovProposalsPost

> StdTx GovProposalsPost(ctx, postProposalBody)

Submit a proposal

Send transaction to submit a proposal

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postProposalBody** | [**InlineObject8**](InlineObject8.md)|  | 

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


## GovProposalsProposalIdDepositsDepositorGet

> Deposit GovProposalsProposalIdDepositsDepositorGet(ctx, proposalId, depositor)

Query deposit

Query deposit by proposalId and depositor address

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**proposalId** | **string**| proposal id | 
**depositor** | **string**| Bech32 depositor address | 

### Return type

[**Deposit**](Deposit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GovProposalsProposalIdDepositsGet

> []Deposit GovProposalsProposalIdDepositsGet(ctx, proposalId)

Query deposits

Query deposits by proposalId

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**proposalId** | **string**|  | 

### Return type

[**[]Deposit**](Deposit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GovProposalsProposalIdDepositsPost

> BroadcastTxCommitResult GovProposalsProposalIdDepositsPost(ctx, proposalId, postDepositBody)

Deposit tokens to a proposal

Send transaction to deposit tokens to a proposal

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**proposalId** | **string**| proposal id | 
**postDepositBody** | [**InlineObject10**](InlineObject10.md)|  | 

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


## GovProposalsProposalIdGet

> TextProposal GovProposalsProposalIdGet(ctx, proposalId)

Query a proposal

Query a proposal by id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**proposalId** | **string**|  | 

### Return type

[**TextProposal**](TextProposal.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GovProposalsProposalIdProposerGet

> Proposer GovProposalsProposalIdProposerGet(ctx, proposalId)

Query proposer

Query for the proposer for a proposal

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**proposalId** | **string**|  | 

### Return type

[**Proposer**](Proposer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GovProposalsProposalIdTallyGet

> TallyResult GovProposalsProposalIdTallyGet(ctx, proposalId)

Get a proposal's tally result at the current time

Gets a proposal's tally result at the current time. If the proposal is pending deposits (i.e status 'DepositPeriod') it returns an empty tally result.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**proposalId** | **string**| proposal id | 

### Return type

[**TallyResult**](TallyResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GovProposalsProposalIdVotesGet

> []Vote GovProposalsProposalIdVotesGet(ctx, proposalId)

Query voters

Query voters information by proposalId

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**proposalId** | **string**| proposal id | 

### Return type

[**[]Vote**](Vote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GovProposalsProposalIdVotesPost

> BroadcastTxCommitResult GovProposalsProposalIdVotesPost(ctx, proposalId, postVoteBody)

Vote a proposal

Send transaction to vote a proposal

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**proposalId** | **string**| proposal id | 
**postVoteBody** | [**InlineObject11**](InlineObject11.md)|  | 

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


## GovProposalsProposalIdVotesVoterGet

> Vote GovProposalsProposalIdVotesVoterGet(ctx, proposalId, voter)

Query vote

Query vote information by proposal Id and voter address

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**proposalId** | **string**| proposal id | 
**voter** | **string**| Bech32 voter address | 

### Return type

[**Vote**](Vote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

