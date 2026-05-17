# \SteeringAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApproveTask**](SteeringAPI.md#ApproveTask) | **Post** /v1/arc/tasks/{task_id}/approve | Approve Task
[**SteerAskCustomerV1**](SteeringAPI.md#SteerAskCustomerV1) | **Post** /v1/arc/tasks/{task_id}/steer/ask-customer | Steer Ask Customer
[**SteerEscalate**](SteeringAPI.md#SteerEscalate) | **Post** /v1/arc/tasks/{task_id}/steer/escalate | Steer Escalate
[**SteerOverride**](SteeringAPI.md#SteerOverride) | **Post** /v1/arc/tasks/{task_id}/steer/override | Steer Override
[**SteerReattempt**](SteeringAPI.md#SteerReattempt) | **Post** /v1/arc/tasks/{task_id}/steer/reattempt | Steer Reattempt



## ApproveTask

> ApproveResponse ApproveTask(ctx, taskId).XAPIKey(xAPIKey).ApproveRequest(approveRequest).XTenantId(xTenantId).XOrgId(xOrgId).XActingUser(xActingUser).Execute()

Approve Task

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
	taskId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	approveRequest := *openapiclient.NewApproveRequest() // ApproveRequest | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	xActingUser := "xActingUser_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SteeringAPI.ApproveTask(context.Background(), taskId).XAPIKey(xAPIKey).ApproveRequest(approveRequest).XTenantId(xTenantId).XOrgId(xOrgId).XActingUser(xActingUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SteeringAPI.ApproveTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApproveTask`: ApproveResponse
	fmt.Fprintf(os.Stdout, "Response from `SteeringAPI.ApproveTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApproveTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **approveRequest** | [**ApproveRequest**](ApproveRequest.md) |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **xActingUser** | **string** |  | 

### Return type

[**ApproveResponse**](ApproveResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SteerAskCustomerV1

> AskCustomerResponse SteerAskCustomerV1(ctx, taskId).XAPIKey(xAPIKey).AskCustomerRequest(askCustomerRequest).XTenantId(xTenantId).XOrgId(xOrgId).XActingUser(xActingUser).Execute()

Steer Ask Customer

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
	taskId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	askCustomerRequest := *openapiclient.NewAskCustomerRequest("MissingField_example") // AskCustomerRequest | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	xActingUser := "xActingUser_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SteeringAPI.SteerAskCustomerV1(context.Background(), taskId).XAPIKey(xAPIKey).AskCustomerRequest(askCustomerRequest).XTenantId(xTenantId).XOrgId(xOrgId).XActingUser(xActingUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SteeringAPI.SteerAskCustomerV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SteerAskCustomerV1`: AskCustomerResponse
	fmt.Fprintf(os.Stdout, "Response from `SteeringAPI.SteerAskCustomerV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSteerAskCustomerV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **askCustomerRequest** | [**AskCustomerRequest**](AskCustomerRequest.md) |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **xActingUser** | **string** |  | 

### Return type

[**AskCustomerResponse**](AskCustomerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SteerEscalate

> EscalateResponse SteerEscalate(ctx, taskId).XAPIKey(xAPIKey).EscalateRequest(escalateRequest).XTenantId(xTenantId).XOrgId(xOrgId).XActingUser(xActingUser).Execute()

Steer Escalate

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
	taskId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	escalateRequest := *openapiclient.NewEscalateRequest("Reason_example") // EscalateRequest | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	xActingUser := "xActingUser_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SteeringAPI.SteerEscalate(context.Background(), taskId).XAPIKey(xAPIKey).EscalateRequest(escalateRequest).XTenantId(xTenantId).XOrgId(xOrgId).XActingUser(xActingUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SteeringAPI.SteerEscalate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SteerEscalate`: EscalateResponse
	fmt.Fprintf(os.Stdout, "Response from `SteeringAPI.SteerEscalate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSteerEscalateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **escalateRequest** | [**EscalateRequest**](EscalateRequest.md) |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **xActingUser** | **string** |  | 

### Return type

[**EscalateResponse**](EscalateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SteerOverride

> OverrideResponse SteerOverride(ctx, taskId).XAPIKey(xAPIKey).OverrideRequest(overrideRequest).XTenantId(xTenantId).XOrgId(xOrgId).XActingUser(xActingUser).Execute()

Steer Override

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
	taskId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	overrideRequest := *openapiclient.NewOverrideRequest(map[string]interface{}{"key": interface{}(123)}) // OverrideRequest | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	xActingUser := "xActingUser_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SteeringAPI.SteerOverride(context.Background(), taskId).XAPIKey(xAPIKey).OverrideRequest(overrideRequest).XTenantId(xTenantId).XOrgId(xOrgId).XActingUser(xActingUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SteeringAPI.SteerOverride``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SteerOverride`: OverrideResponse
	fmt.Fprintf(os.Stdout, "Response from `SteeringAPI.SteerOverride`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSteerOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **overrideRequest** | [**OverrideRequest**](OverrideRequest.md) |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **xActingUser** | **string** |  | 

### Return type

[**OverrideResponse**](OverrideResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SteerReattempt

> ReattemptResponse SteerReattempt(ctx, taskId).XAPIKey(xAPIKey).ReattemptRequest(reattemptRequest).XTenantId(xTenantId).XOrgId(xOrgId).XActingUser(xActingUser).Execute()

Steer Reattempt

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
	taskId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	reattemptRequest := *openapiclient.NewReattemptRequest("Hint_example") // ReattemptRequest | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	xActingUser := "xActingUser_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SteeringAPI.SteerReattempt(context.Background(), taskId).XAPIKey(xAPIKey).ReattemptRequest(reattemptRequest).XTenantId(xTenantId).XOrgId(xOrgId).XActingUser(xActingUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SteeringAPI.SteerReattempt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SteerReattempt`: ReattemptResponse
	fmt.Fprintf(os.Stdout, "Response from `SteeringAPI.SteerReattempt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSteerReattemptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **reattemptRequest** | [**ReattemptRequest**](ReattemptRequest.md) |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **xActingUser** | **string** |  | 

### Return type

[**ReattemptResponse**](ReattemptResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

