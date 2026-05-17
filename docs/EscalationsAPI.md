# \EscalationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEscalation**](EscalationsAPI.md#CreateEscalation) | **Post** /v1/arc/cases/{case_id}/escalations | Create Escalation
[**DismissEscalation**](EscalationsAPI.md#DismissEscalation) | **Post** /v1/arc/cases/{case_id}/escalations/{escalation_id}/dismiss | Dismiss Escalation
[**DispatchEscalationAction**](EscalationsAPI.md#DispatchEscalationAction) | **Post** /v1/arc/escalations/{escalation_id}/actions | Dispatch Escalation Action
[**GetEscalation**](EscalationsAPI.md#GetEscalation) | **Get** /v1/arc/cases/{case_id}/escalations/{escalation_id} | Get Escalation
[**ListEscalations**](EscalationsAPI.md#ListEscalations) | **Get** /v1/arc/cases/{case_id}/escalations | List Escalations
[**ResolveEscalation**](EscalationsAPI.md#ResolveEscalation) | **Post** /v1/arc/cases/{case_id}/escalations/{escalation_id}/resolve | Resolve Escalation
[**UpdateEscalation**](EscalationsAPI.md#UpdateEscalation) | **Patch** /v1/arc/cases/{case_id}/escalations/{escalation_id} | Update Escalation



## CreateEscalation

> EscalationResponse CreateEscalation(ctx, caseId).XAPIKey(xAPIKey).EscalationCreate(escalationCreate).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Create Escalation



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
	xAPIKey := "xAPIKey_example" // string | 
	escalationCreate := *openapiclient.NewEscalationCreate("Reason_example") // EscalationCreate | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EscalationsAPI.CreateEscalation(context.Background(), caseId).XAPIKey(xAPIKey).EscalationCreate(escalationCreate).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EscalationsAPI.CreateEscalation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEscalation`: EscalationResponse
	fmt.Fprintf(os.Stdout, "Response from `EscalationsAPI.CreateEscalation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEscalationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **escalationCreate** | [**EscalationCreate**](EscalationCreate.md) |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**EscalationResponse**](EscalationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DismissEscalation

> EscalationResponse DismissEscalation(ctx, caseId, escalationId).XAPIKey(xAPIKey).EscalationDismissRequest(escalationDismissRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Dismiss Escalation



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
	escalationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	escalationDismissRequest := *openapiclient.NewEscalationDismissRequest("Reason_example") // EscalationDismissRequest | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EscalationsAPI.DismissEscalation(context.Background(), caseId, escalationId).XAPIKey(xAPIKey).EscalationDismissRequest(escalationDismissRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EscalationsAPI.DismissEscalation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DismissEscalation`: EscalationResponse
	fmt.Fprintf(os.Stdout, "Response from `EscalationsAPI.DismissEscalation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  | 
**escalationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDismissEscalationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAPIKey** | **string** |  | 
 **escalationDismissRequest** | [**EscalationDismissRequest**](EscalationDismissRequest.md) |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**EscalationResponse**](EscalationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DispatchEscalationAction

> EscalationActionResponse DispatchEscalationAction(ctx, escalationId).XAPIKey(xAPIKey).EscalationActionRequest(escalationActionRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Dispatch Escalation Action



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
	escalationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	escalationActionRequest := *openapiclient.NewEscalationActionRequest("ActionType_example") // EscalationActionRequest | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EscalationsAPI.DispatchEscalationAction(context.Background(), escalationId).XAPIKey(xAPIKey).EscalationActionRequest(escalationActionRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EscalationsAPI.DispatchEscalationAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DispatchEscalationAction`: EscalationActionResponse
	fmt.Fprintf(os.Stdout, "Response from `EscalationsAPI.DispatchEscalationAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**escalationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDispatchEscalationActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **escalationActionRequest** | [**EscalationActionRequest**](EscalationActionRequest.md) |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**EscalationActionResponse**](EscalationActionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEscalation

> EscalationResponse GetEscalation(ctx, caseId, escalationId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Get Escalation



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
	escalationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EscalationsAPI.GetEscalation(context.Background(), caseId, escalationId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EscalationsAPI.GetEscalation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEscalation`: EscalationResponse
	fmt.Fprintf(os.Stdout, "Response from `EscalationsAPI.GetEscalation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  | 
**escalationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEscalationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**EscalationResponse**](EscalationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEscalations

> EscalationListResponse ListEscalations(ctx, caseId).XAPIKey(xAPIKey).Page(page).PageSize(pageSize).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

List Escalations



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
	xAPIKey := "xAPIKey_example" // string | 
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 50)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EscalationsAPI.ListEscalations(context.Background(), caseId).XAPIKey(xAPIKey).Page(page).PageSize(pageSize).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EscalationsAPI.ListEscalations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEscalations`: EscalationListResponse
	fmt.Fprintf(os.Stdout, "Response from `EscalationsAPI.ListEscalations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEscalationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 50]
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**EscalationListResponse**](EscalationListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResolveEscalation

> EscalationResponse ResolveEscalation(ctx, caseId, escalationId).XAPIKey(xAPIKey).EscalationResolveRequest(escalationResolveRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Resolve Escalation



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
	escalationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	escalationResolveRequest := *openapiclient.NewEscalationResolveRequest() // EscalationResolveRequest | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EscalationsAPI.ResolveEscalation(context.Background(), caseId, escalationId).XAPIKey(xAPIKey).EscalationResolveRequest(escalationResolveRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EscalationsAPI.ResolveEscalation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResolveEscalation`: EscalationResponse
	fmt.Fprintf(os.Stdout, "Response from `EscalationsAPI.ResolveEscalation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  | 
**escalationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResolveEscalationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAPIKey** | **string** |  | 
 **escalationResolveRequest** | [**EscalationResolveRequest**](EscalationResolveRequest.md) |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**EscalationResponse**](EscalationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEscalation

> EscalationResponse UpdateEscalation(ctx, caseId, escalationId).XAPIKey(xAPIKey).EscalationUpdate(escalationUpdate).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Update Escalation



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
	escalationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	escalationUpdate := *openapiclient.NewEscalationUpdate() // EscalationUpdate | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EscalationsAPI.UpdateEscalation(context.Background(), caseId, escalationId).XAPIKey(xAPIKey).EscalationUpdate(escalationUpdate).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EscalationsAPI.UpdateEscalation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEscalation`: EscalationResponse
	fmt.Fprintf(os.Stdout, "Response from `EscalationsAPI.UpdateEscalation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  | 
**escalationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEscalationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAPIKey** | **string** |  | 
 **escalationUpdate** | [**EscalationUpdate**](EscalationUpdate.md) |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**EscalationResponse**](EscalationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

