# \PlansAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AmendPlanEndpointV1**](PlansAPI.md#AmendPlanEndpointV1) | **Post** /v1/arc/cases/{case_id}/payment-plans/{plan_id}/amend | Amend Plan Endpoint
[**CreatePlanEndpointV1**](PlansAPI.md#CreatePlanEndpointV1) | **Post** /v1/arc/cases/{case_id}/payment-plans | Create Plan Endpoint
[**ListInstallmentsEndpointV1**](PlansAPI.md#ListInstallmentsEndpointV1) | **Get** /v1/arc/cases/{case_id}/payment-plans/{plan_id}/installments | List Installments Endpoint
[**RescheduleInstallmentEndpointV1**](PlansAPI.md#RescheduleInstallmentEndpointV1) | **Post** /v1/arc/cases/{case_id}/payment-plans/{plan_id}/installments/{installment_id}/reschedule | Reschedule Installment Endpoint
[**WaiveInstallmentEndpointV1**](PlansAPI.md#WaiveInstallmentEndpointV1) | **Post** /v1/arc/cases/{case_id}/payment-plans/{plan_id}/installments/{installment_id}/waive | Waive Installment Endpoint



## AmendPlanEndpointV1

> interface{} AmendPlanEndpointV1(ctx, caseId, planId).XAPIKey(xAPIKey).PlanAmendRequest(planAmendRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Amend Plan Endpoint



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/moolabs/moolabs-go"
)

func main() {
	caseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	planId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	planAmendRequest := *openapiclient.NewPlanAmendRequest(int32(123), int32(123), time.Now()) // PlanAmendRequest | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlansAPI.AmendPlanEndpointV1(context.Background(), caseId, planId).XAPIKey(xAPIKey).PlanAmendRequest(planAmendRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.AmendPlanEndpointV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AmendPlanEndpointV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.AmendPlanEndpointV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  | 
**planId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAmendPlanEndpointV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAPIKey** | **string** |  | 
 **planAmendRequest** | [**PlanAmendRequest**](PlanAmendRequest.md) |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePlanEndpointV1

> interface{} CreatePlanEndpointV1(ctx, caseId).XAPIKey(xAPIKey).PlanCreateRequest(planCreateRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Create Plan Endpoint



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/moolabs/moolabs-go"
)

func main() {
	caseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	planCreateRequest := *openapiclient.NewPlanCreateRequest(int32(123), int32(123), time.Now(), "ProposedBy_example") // PlanCreateRequest | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlansAPI.CreatePlanEndpointV1(context.Background(), caseId).XAPIKey(xAPIKey).PlanCreateRequest(planCreateRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.CreatePlanEndpointV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePlanEndpointV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.CreatePlanEndpointV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePlanEndpointV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **planCreateRequest** | [**PlanCreateRequest**](PlanCreateRequest.md) |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInstallmentsEndpointV1

> interface{} ListInstallmentsEndpointV1(ctx, caseId, planId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

List Installments Endpoint



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
	planId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlansAPI.ListInstallmentsEndpointV1(context.Background(), caseId, planId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.ListInstallmentsEndpointV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInstallmentsEndpointV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.ListInstallmentsEndpointV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  | 
**planId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListInstallmentsEndpointV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RescheduleInstallmentEndpointV1

> interface{} RescheduleInstallmentEndpointV1(ctx, caseId, planId, installmentId).XAPIKey(xAPIKey).RescheduleRequest(rescheduleRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Reschedule Installment Endpoint



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/moolabs/moolabs-go"
)

func main() {
	caseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	planId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	installmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	rescheduleRequest := *openapiclient.NewRescheduleRequest(time.Now()) // RescheduleRequest | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlansAPI.RescheduleInstallmentEndpointV1(context.Background(), caseId, planId, installmentId).XAPIKey(xAPIKey).RescheduleRequest(rescheduleRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.RescheduleInstallmentEndpointV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RescheduleInstallmentEndpointV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.RescheduleInstallmentEndpointV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  | 
**planId** | **string** |  | 
**installmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRescheduleInstallmentEndpointV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xAPIKey** | **string** |  | 
 **rescheduleRequest** | [**RescheduleRequest**](RescheduleRequest.md) |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WaiveInstallmentEndpointV1

> interface{} WaiveInstallmentEndpointV1(ctx, caseId, planId, installmentId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Waive Installment Endpoint



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
	planId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	installmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlansAPI.WaiveInstallmentEndpointV1(context.Background(), caseId, planId, installmentId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.WaiveInstallmentEndpointV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WaiveInstallmentEndpointV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.WaiveInstallmentEndpointV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  | 
**planId** | **string** |  | 
**installmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWaiveInstallmentEndpointV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

