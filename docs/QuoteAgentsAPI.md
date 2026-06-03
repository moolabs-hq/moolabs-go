# \QuoteAgentsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAgentEvaluationsV1**](QuoteAgentsAPI.md#ListAgentEvaluationsV1) | **Get** /v1/quote-agents/evaluations | List Agent Evaluations
[**ListAgentRunsV1**](QuoteAgentsAPI.md#ListAgentRunsV1) | **Get** /v1/quote-agents/runs | List Agent Runs
[**ListPoliciesV1**](QuoteAgentsAPI.md#ListPoliciesV1) | **Get** /v1/quote-agents/policies | List Policies
[**PatchPolicyV1**](QuoteAgentsAPI.md#PatchPolicyV1) | **Patch** /v1/quote-agents/policies/{agent_name} | Patch Policy



## ListAgentEvaluationsV1

> []QuoteAgentEvaluationResponse ListAgentEvaluationsV1(ctx).Limit(limit).Offset(offset).Execute()

List Agent Evaluations

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/moolabs/moolabs-go"
)

func main() {
	limit := int32(56) // int32 |  (optional) (default to 100)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuoteAgentsAPI.ListAgentEvaluationsV1(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuoteAgentsAPI.ListAgentEvaluationsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAgentEvaluationsV1`: []QuoteAgentEvaluationResponse
	fmt.Fprintf(os.Stdout, "Response from `QuoteAgentsAPI.ListAgentEvaluationsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAgentEvaluationsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 100]
 **offset** | **int32** |  | [default to 0]

### Return type

[**[]QuoteAgentEvaluationResponse**](QuoteAgentEvaluationResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAgentRunsV1

> []QuoteAgentRunResponse ListAgentRunsV1(ctx).Limit(limit).Offset(offset).Execute()

List Agent Runs

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/moolabs/moolabs-go"
)

func main() {
	limit := int32(56) // int32 |  (optional) (default to 100)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuoteAgentsAPI.ListAgentRunsV1(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuoteAgentsAPI.ListAgentRunsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAgentRunsV1`: []QuoteAgentRunResponse
	fmt.Fprintf(os.Stdout, "Response from `QuoteAgentsAPI.ListAgentRunsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAgentRunsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 100]
 **offset** | **int32** |  | [default to 0]

### Return type

[**[]QuoteAgentRunResponse**](QuoteAgentRunResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPoliciesV1

> []QuoteAgentPolicyResponse ListPoliciesV1(ctx).Execute()

List Policies

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/moolabs/moolabs-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuoteAgentsAPI.ListPoliciesV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuoteAgentsAPI.ListPoliciesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPoliciesV1`: []QuoteAgentPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `QuoteAgentsAPI.ListPoliciesV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPoliciesV1Request struct via the builder pattern


### Return type

[**[]QuoteAgentPolicyResponse**](QuoteAgentPolicyResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchPolicyV1

> QuoteAgentPolicyResponse PatchPolicyV1(ctx, agentName).PatchQuoteAgentPolicyRequest(patchQuoteAgentPolicyRequest).Execute()

Patch Policy

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/moolabs/moolabs-go"
)

func main() {
	agentName := "agentName_example" // string | 
	patchQuoteAgentPolicyRequest := *openapiclient.NewPatchQuoteAgentPolicyRequest() // PatchQuoteAgentPolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuoteAgentsAPI.PatchPolicyV1(context.Background(), agentName).PatchQuoteAgentPolicyRequest(patchQuoteAgentPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuoteAgentsAPI.PatchPolicyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchPolicyV1`: QuoteAgentPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `QuoteAgentsAPI.PatchPolicyV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPolicyV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchQuoteAgentPolicyRequest** | [**PatchQuoteAgentPolicyRequest**](PatchQuoteAgentPolicyRequest.md) |  | 

### Return type

[**QuoteAgentPolicyResponse**](QuoteAgentPolicyResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

