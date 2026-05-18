# \DisputesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddEvidenceEndpoint**](DisputesAPI.md#AddEvidenceEndpoint) | **Post** /v1/arc/disputes/{dispute_id}/evidence | Add Evidence Endpoint
[**AddExternalDocEvidenceEndpoint**](DisputesAPI.md#AddExternalDocEvidenceEndpoint) | **Post** /v1/arc/disputes/{dispute_id}/evidence/upload | Add External Doc Evidence Endpoint
[**ApproveDisputeEndpoint**](DisputesAPI.md#ApproveDisputeEndpoint) | **Post** /v1/arc/disputes/{dispute_id}/approve | Approve Dispute Endpoint
[**GetDisputeEndpoint**](DisputesAPI.md#GetDisputeEndpoint) | **Get** /v1/arc/disputes/{dispute_id} | Get Dispute Endpoint
[**ListDisputesEndpoint**](DisputesAPI.md#ListDisputesEndpoint) | **Get** /v1/arc/disputes | List Disputes Endpoint
[**ListEvidenceEndpoint**](DisputesAPI.md#ListEvidenceEndpoint) | **Get** /v1/arc/disputes/{dispute_id}/evidence | List Evidence Endpoint
[**RejectDisputeEndpoint**](DisputesAPI.md#RejectDisputeEndpoint) | **Post** /v1/arc/disputes/{dispute_id}/reject | Reject Dispute Endpoint
[**ResolveDisputeEndpoint**](DisputesAPI.md#ResolveDisputeEndpoint) | **Post** /v1/arc/disputes/{dispute_id}/resolve | Resolve Dispute Endpoint
[**ReviewEvidenceEndpoint**](DisputesAPI.md#ReviewEvidenceEndpoint) | **Post** /v1/arc/disputes/{dispute_id}/evidence/{evidence_id}/review | Review Evidence Endpoint
[**UpdateDisputeEndpoint**](DisputesAPI.md#UpdateDisputeEndpoint) | **Patch** /v1/arc/disputes/{dispute_id} | Update Dispute Endpoint



## AddEvidenceEndpoint

> EvidenceResponse AddEvidenceEndpoint(ctx, disputeId).EvidenceCreateRequest(evidenceCreateRequest).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()

Add Evidence Endpoint



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
	disputeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	evidenceCreateRequest := *openapiclient.NewEvidenceCreateRequest("EvidenceType_example") // EvidenceCreateRequest | 
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisputesAPI.AddEvidenceEndpoint(context.Background(), disputeId).EvidenceCreateRequest(evidenceCreateRequest).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisputesAPI.AddEvidenceEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddEvidenceEndpoint`: EvidenceResponse
	fmt.Fprintf(os.Stdout, "Response from `DisputesAPI.AddEvidenceEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**disputeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddEvidenceEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **evidenceCreateRequest** | [**EvidenceCreateRequest**](EvidenceCreateRequest.md) |  | 
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **authorization** | **string** |  | 

### Return type

[**EvidenceResponse**](EvidenceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddExternalDocEvidenceEndpoint

> EvidenceResponse AddExternalDocEvidenceEndpoint(ctx, disputeId).EvidenceType(evidenceType).Attachment(attachment).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Summary(summary).AmountMicros(amountMicros).Execute()

Add External Doc Evidence Endpoint



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
	disputeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	evidenceType := "evidenceType_example" // string | 
	attachment := os.NewFile(1234, "some_file") // *os.File | 
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)
	summary := "summary_example" // string |  (optional)
	amountMicros := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisputesAPI.AddExternalDocEvidenceEndpoint(context.Background(), disputeId).EvidenceType(evidenceType).Attachment(attachment).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Summary(summary).AmountMicros(amountMicros).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisputesAPI.AddExternalDocEvidenceEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddExternalDocEvidenceEndpoint`: EvidenceResponse
	fmt.Fprintf(os.Stdout, "Response from `DisputesAPI.AddExternalDocEvidenceEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**disputeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddExternalDocEvidenceEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **evidenceType** | **string** |  | 
 **attachment** | ***os.File** |  | 
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **authorization** | **string** |  | 
 **summary** | **string** |  | 
 **amountMicros** | **int32** |  | 

### Return type

[**EvidenceResponse**](EvidenceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApproveDisputeEndpoint

> DisputeResponse ApproveDisputeEndpoint(ctx, disputeId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).DisputeApprovalRequest(disputeApprovalRequest).Execute()

Approve Dispute Endpoint



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
	disputeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)
	disputeApprovalRequest := *openapiclient.NewDisputeApprovalRequest() // DisputeApprovalRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisputesAPI.ApproveDisputeEndpoint(context.Background(), disputeId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).DisputeApprovalRequest(disputeApprovalRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisputesAPI.ApproveDisputeEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApproveDisputeEndpoint`: DisputeResponse
	fmt.Fprintf(os.Stdout, "Response from `DisputesAPI.ApproveDisputeEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**disputeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApproveDisputeEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **authorization** | **string** |  | 
 **disputeApprovalRequest** | [**DisputeApprovalRequest**](DisputeApprovalRequest.md) |  | 

### Return type

[**DisputeResponse**](DisputeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDisputeEndpoint

> DisputeResponse GetDisputeEndpoint(ctx, disputeId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()

Get Dispute Endpoint



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
	disputeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisputesAPI.GetDisputeEndpoint(context.Background(), disputeId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisputesAPI.GetDisputeEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDisputeEndpoint`: DisputeResponse
	fmt.Fprintf(os.Stdout, "Response from `DisputesAPI.GetDisputeEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**disputeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDisputeEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **authorization** | **string** |  | 

### Return type

[**DisputeResponse**](DisputeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDisputesEndpoint

> DisputeListResponse ListDisputesEndpoint(ctx).Page(page).PageSize(pageSize).CaseId(caseId).Status(status).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()

List Disputes Endpoint



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
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 50)
	caseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	status := "status_example" // string |  (optional)
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisputesAPI.ListDisputesEndpoint(context.Background()).Page(page).PageSize(pageSize).CaseId(caseId).Status(status).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisputesAPI.ListDisputesEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDisputesEndpoint`: DisputeListResponse
	fmt.Fprintf(os.Stdout, "Response from `DisputesAPI.ListDisputesEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDisputesEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 50]
 **caseId** | **string** |  | 
 **status** | **string** |  | 
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **authorization** | **string** |  | 

### Return type

[**DisputeListResponse**](DisputeListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEvidenceEndpoint

> EvidenceListResponse ListEvidenceEndpoint(ctx, disputeId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()

List Evidence Endpoint



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
	disputeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisputesAPI.ListEvidenceEndpoint(context.Background(), disputeId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisputesAPI.ListEvidenceEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEvidenceEndpoint`: EvidenceListResponse
	fmt.Fprintf(os.Stdout, "Response from `DisputesAPI.ListEvidenceEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**disputeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEvidenceEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **authorization** | **string** |  | 

### Return type

[**EvidenceListResponse**](EvidenceListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RejectDisputeEndpoint

> DisputeResponse RejectDisputeEndpoint(ctx, disputeId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).DisputeApprovalRequest(disputeApprovalRequest).Execute()

Reject Dispute Endpoint



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
	disputeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)
	disputeApprovalRequest := *openapiclient.NewDisputeApprovalRequest() // DisputeApprovalRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisputesAPI.RejectDisputeEndpoint(context.Background(), disputeId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).DisputeApprovalRequest(disputeApprovalRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisputesAPI.RejectDisputeEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RejectDisputeEndpoint`: DisputeResponse
	fmt.Fprintf(os.Stdout, "Response from `DisputesAPI.RejectDisputeEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**disputeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRejectDisputeEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **authorization** | **string** |  | 
 **disputeApprovalRequest** | [**DisputeApprovalRequest**](DisputeApprovalRequest.md) |  | 

### Return type

[**DisputeResponse**](DisputeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResolveDisputeEndpoint

> DisputeResponse ResolveDisputeEndpoint(ctx, disputeId).DisputeResolveRequest(disputeResolveRequest).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()

Resolve Dispute Endpoint



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
	disputeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	disputeResolveRequest := *openapiclient.NewDisputeResolveRequest("Resolution_example") // DisputeResolveRequest | 
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisputesAPI.ResolveDisputeEndpoint(context.Background(), disputeId).DisputeResolveRequest(disputeResolveRequest).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisputesAPI.ResolveDisputeEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResolveDisputeEndpoint`: DisputeResponse
	fmt.Fprintf(os.Stdout, "Response from `DisputesAPI.ResolveDisputeEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**disputeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResolveDisputeEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **disputeResolveRequest** | [**DisputeResolveRequest**](DisputeResolveRequest.md) |  | 
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **authorization** | **string** |  | 

### Return type

[**DisputeResponse**](DisputeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReviewEvidenceEndpoint

> EvidenceResponse ReviewEvidenceEndpoint(ctx, disputeId, evidenceId).EvidenceReviewRequest(evidenceReviewRequest).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()

Review Evidence Endpoint



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
	disputeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	evidenceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	evidenceReviewRequest := *openapiclient.NewEvidenceReviewRequest("ResolutionStatus_example", int32(123)) // EvidenceReviewRequest | 
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisputesAPI.ReviewEvidenceEndpoint(context.Background(), disputeId, evidenceId).EvidenceReviewRequest(evidenceReviewRequest).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisputesAPI.ReviewEvidenceEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReviewEvidenceEndpoint`: EvidenceResponse
	fmt.Fprintf(os.Stdout, "Response from `DisputesAPI.ReviewEvidenceEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**disputeId** | **string** |  | 
**evidenceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReviewEvidenceEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **evidenceReviewRequest** | [**EvidenceReviewRequest**](EvidenceReviewRequest.md) |  | 
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **authorization** | **string** |  | 

### Return type

[**EvidenceResponse**](EvidenceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDisputeEndpoint

> DisputeResponse UpdateDisputeEndpoint(ctx, disputeId).DisputeUpdateRequest(disputeUpdateRequest).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()

Update Dispute Endpoint



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
	disputeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	disputeUpdateRequest := *openapiclient.NewDisputeUpdateRequest() // DisputeUpdateRequest | 
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisputesAPI.UpdateDisputeEndpoint(context.Background(), disputeId).DisputeUpdateRequest(disputeUpdateRequest).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisputesAPI.UpdateDisputeEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDisputeEndpoint`: DisputeResponse
	fmt.Fprintf(os.Stdout, "Response from `DisputesAPI.UpdateDisputeEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**disputeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDisputeEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **disputeUpdateRequest** | [**DisputeUpdateRequest**](DisputeUpdateRequest.md) |  | 
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **authorization** | **string** |  | 

### Return type

[**DisputeResponse**](DisputeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

