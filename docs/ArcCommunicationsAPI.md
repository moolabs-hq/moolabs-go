# \ArcCommunicationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApproveCommunication**](ArcCommunicationsAPI.md#ApproveCommunication) | **Post** /v1/arc/cases/{case_id}/communications/{comm_id}/approve | Approve Communication
[**CreateCommunication**](ArcCommunicationsAPI.md#CreateCommunication) | **Post** /v1/arc/cases/{case_id}/communications | Create Communication
[**ListCommunications**](ArcCommunicationsAPI.md#ListCommunications) | **Get** /v1/arc/cases/{case_id}/communications | List Communications
[**RejectCommunication**](ArcCommunicationsAPI.md#RejectCommunication) | **Post** /v1/arc/cases/{case_id}/communications/{comm_id}/reject | Reject Communication
[**ResolveProviderConfirmationV1Arc**](ArcCommunicationsAPI.md#ResolveProviderConfirmationV1Arc) | **Post** /v1/arc/cases/{case_id}/communications/{comm_id}/resolve-provider-confirmation | Resolve Provider Confirmation



## ApproveCommunication

> ApprovalActionResponse ApproveCommunication(ctx, caseId, commId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).ApprovalActionRequest(approvalActionRequest).Execute()

Approve Communication



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
	caseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	commId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)
	approvalActionRequest := *openapiclient.NewApprovalActionRequest() // ApprovalActionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArcCommunicationsAPI.ApproveCommunication(context.Background(), caseId, commId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).ApprovalActionRequest(approvalActionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArcCommunicationsAPI.ApproveCommunication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApproveCommunication`: ApprovalActionResponse
	fmt.Fprintf(os.Stdout, "Response from `ArcCommunicationsAPI.ApproveCommunication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  | 
**commId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApproveCommunicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **authorization** | **string** |  | 
 **approvalActionRequest** | [**ApprovalActionRequest**](ApprovalActionRequest.md) |  | 

### Return type

[**ApprovalActionResponse**](ApprovalActionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCommunication

> CommunicationResponse CreateCommunication(ctx, caseId).CommunicationCreate(communicationCreate).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()

Create Communication



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
	caseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	communicationCreate := *openapiclient.NewCommunicationCreate() // CommunicationCreate | 
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArcCommunicationsAPI.CreateCommunication(context.Background(), caseId).CommunicationCreate(communicationCreate).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArcCommunicationsAPI.CreateCommunication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCommunication`: CommunicationResponse
	fmt.Fprintf(os.Stdout, "Response from `ArcCommunicationsAPI.CreateCommunication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCommunicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **communicationCreate** | [**CommunicationCreate**](CommunicationCreate.md) |  | 
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **authorization** | **string** |  | 

### Return type

[**CommunicationResponse**](CommunicationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCommunications

> CommunicationListResponse ListCommunications(ctx, caseId).Page(page).PageSize(pageSize).Direction(direction).Channel(channel).Status(status).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()

List Communications



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
	caseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 50)
	direction := "direction_example" // string |  (optional)
	channel := "channel_example" // string |  (optional)
	status := "status_example" // string |  (optional)
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArcCommunicationsAPI.ListCommunications(context.Background(), caseId).Page(page).PageSize(pageSize).Direction(direction).Channel(channel).Status(status).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArcCommunicationsAPI.ListCommunications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCommunications`: CommunicationListResponse
	fmt.Fprintf(os.Stdout, "Response from `ArcCommunicationsAPI.ListCommunications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCommunicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 50]
 **direction** | **string** |  | 
 **channel** | **string** |  | 
 **status** | **string** |  | 
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **authorization** | **string** |  | 

### Return type

[**CommunicationListResponse**](CommunicationListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RejectCommunication

> ApprovalActionResponse RejectCommunication(ctx, caseId, commId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).ApprovalActionRequest(approvalActionRequest).Execute()

Reject Communication



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
	caseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	commId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)
	approvalActionRequest := *openapiclient.NewApprovalActionRequest() // ApprovalActionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArcCommunicationsAPI.RejectCommunication(context.Background(), caseId, commId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).ApprovalActionRequest(approvalActionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArcCommunicationsAPI.RejectCommunication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RejectCommunication`: ApprovalActionResponse
	fmt.Fprintf(os.Stdout, "Response from `ArcCommunicationsAPI.RejectCommunication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  | 
**commId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRejectCommunicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **authorization** | **string** |  | 
 **approvalActionRequest** | [**ApprovalActionRequest**](ApprovalActionRequest.md) |  | 

### Return type

[**ApprovalActionResponse**](ApprovalActionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResolveProviderConfirmationV1Arc

> ApprovalActionResponse ResolveProviderConfirmationV1Arc(ctx, caseId, commId).ProviderConfirmationRepairRequest(providerConfirmationRepairRequest).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).XUserId(xUserId).XArcRoles(xArcRoles).XArcProxySecret(xArcProxySecret).Execute()

Resolve Provider Confirmation



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
	caseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	commId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	providerConfirmationRepairRequest := *openapiclient.NewProviderConfirmationRepairRequest("Outcome_example") // ProviderConfirmationRepairRequest | 
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)
	xUserId := "xUserId_example" // string |  (optional)
	xArcRoles := "xArcRoles_example" // string |  (optional)
	xArcProxySecret := "xArcProxySecret_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArcCommunicationsAPI.ResolveProviderConfirmationV1Arc(context.Background(), caseId, commId).ProviderConfirmationRepairRequest(providerConfirmationRepairRequest).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).XUserId(xUserId).XArcRoles(xArcRoles).XArcProxySecret(xArcProxySecret).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArcCommunicationsAPI.ResolveProviderConfirmationV1Arc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResolveProviderConfirmationV1Arc`: ApprovalActionResponse
	fmt.Fprintf(os.Stdout, "Response from `ArcCommunicationsAPI.ResolveProviderConfirmationV1Arc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  | 
**commId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResolveProviderConfirmationV1ArcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **providerConfirmationRepairRequest** | [**ProviderConfirmationRepairRequest**](ProviderConfirmationRepairRequest.md) |  | 
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **authorization** | **string** |  | 
 **xUserId** | **string** |  | 
 **xArcRoles** | **string** |  | 
 **xArcProxySecret** | **string** |  | 

### Return type

[**ApprovalActionResponse**](ApprovalActionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

