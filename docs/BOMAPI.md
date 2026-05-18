# \BOMAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApproveBom**](BOMAPI.md#ApproveBom) | **Post** /api/v1/bom/{bom_id}/approve | Approve BOM (pending_approval → active)
[**AutoDerive**](BOMAPI.md#AutoDerive) | **Post** /api/v1/bom/{bom_id}/derive | Auto-derive BOM from observed cost_events
[**CreateBom**](BOMAPI.md#CreateBom) | **Post** /api/v1/bom | Create BOM (status&#x3D;draft)
[**CreateNewVersion**](BOMAPI.md#CreateNewVersion) | **Put** /api/v1/bom/{bom_id} | Create new BOM version (immutable update)
[**GetBom**](BOMAPI.md#GetBom) | **Get** /api/v1/bom/{bom_id} | Get BOM with components
[**GetObservationCompleteness**](BOMAPI.md#GetObservationCompleteness) | **Get** /api/v1/bom/{bom_id}/observation | Observation completeness vs BOM expected_span_count
[**ListBoms**](BOMAPI.md#ListBoms) | **Get** /api/v1/bom | List BOMs for tenant
[**SimulateCost**](BOMAPI.md#SimulateCost) | **Get** /api/v1/bom/{bom_id}/simulate | Simulate BOM cost with current rate_catalog rates
[**SubmitForApproval**](BOMAPI.md#SubmitForApproval) | **Post** /api/v1/bom/{bom_id}/submit | Submit BOM for approval (draft → pending_approval)



## ApproveBom

> BomOut ApproveBom(ctx, bomId).TenantId(tenantId).AcuteApproveRequest(acuteApproveRequest).Execute()

Approve BOM (pending_approval → active)



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
	bomId := "bomId_example" // string | 
	tenantId := "tenantId_example" // string | Tenant UUID
	acuteApproveRequest := *openapiclient.NewAcuteApproveRequest("ApprovedBy_example") // AcuteApproveRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BOMAPI.ApproveBom(context.Background(), bomId).TenantId(tenantId).AcuteApproveRequest(acuteApproveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BOMAPI.ApproveBom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApproveBom`: BomOut
	fmt.Fprintf(os.Stdout, "Response from `BOMAPI.ApproveBom`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bomId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApproveBomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** | Tenant UUID | 
 **acuteApproveRequest** | [**AcuteApproveRequest**](AcuteApproveRequest.md) |  | 

### Return type

[**BomOut**](BomOut.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutoDerive

> BomOut AutoDerive(ctx, bomId).DeriveRequest(deriveRequest).Execute()

Auto-derive BOM from observed cost_events



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
	bomId := "bomId_example" // string | 
	deriveRequest := *openapiclient.NewDeriveRequest("TenantId_example", "FeatureKey_example", time.Now(), time.Now()) // DeriveRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BOMAPI.AutoDerive(context.Background(), bomId).DeriveRequest(deriveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BOMAPI.AutoDerive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoDerive`: BomOut
	fmt.Fprintf(os.Stdout, "Response from `BOMAPI.AutoDerive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bomId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutoDeriveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deriveRequest** | [**DeriveRequest**](DeriveRequest.md) |  | 

### Return type

[**BomOut**](BomOut.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBom

> BomOut CreateBom(ctx).BomCreateRequest(bomCreateRequest).Execute()

Create BOM (status=draft)



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
	bomCreateRequest := *openapiclient.NewBomCreateRequest("TenantId_example", "FeatureKey_example", "Name_example", []openapiclient.BomComponentIn{*openapiclient.NewBomComponentIn("Provider_example", "Model_example", "MetricType_example", float32(123))}) // BomCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BOMAPI.CreateBom(context.Background()).BomCreateRequest(bomCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BOMAPI.CreateBom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBom`: BomOut
	fmt.Fprintf(os.Stdout, "Response from `BOMAPI.CreateBom`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bomCreateRequest** | [**BomCreateRequest**](BomCreateRequest.md) |  | 

### Return type

[**BomOut**](BomOut.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNewVersion

> BomOut CreateNewVersion(ctx, bomId).TenantId(tenantId).BomUpdateRequest(bomUpdateRequest).Execute()

Create new BOM version (immutable update)



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
	bomId := "bomId_example" // string | 
	tenantId := "tenantId_example" // string | Tenant UUID
	bomUpdateRequest := *openapiclient.NewBomUpdateRequest() // BomUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BOMAPI.CreateNewVersion(context.Background(), bomId).TenantId(tenantId).BomUpdateRequest(bomUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BOMAPI.CreateNewVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNewVersion`: BomOut
	fmt.Fprintf(os.Stdout, "Response from `BOMAPI.CreateNewVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bomId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** | Tenant UUID | 
 **bomUpdateRequest** | [**BomUpdateRequest**](BomUpdateRequest.md) |  | 

### Return type

[**BomOut**](BomOut.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBom

> BomOut GetBom(ctx, bomId).TenantId(tenantId).Execute()

Get BOM with components



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
	bomId := "bomId_example" // string | 
	tenantId := "tenantId_example" // string | Tenant UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BOMAPI.GetBom(context.Background(), bomId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BOMAPI.GetBom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBom`: BomOut
	fmt.Fprintf(os.Stdout, "Response from `BOMAPI.GetBom`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bomId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** | Tenant UUID | 

### Return type

[**BomOut**](BomOut.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObservationCompleteness

> ObservationOut GetObservationCompleteness(ctx, bomId).TenantId(tenantId).PeriodStart(periodStart).PeriodEnd(periodEnd).Execute()

Observation completeness vs BOM expected_span_count



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
	bomId := "bomId_example" // string | 
	tenantId := "tenantId_example" // string | Tenant UUID
	periodStart := time.Now() // time.Time | Period start (ISO datetime)
	periodEnd := time.Now() // time.Time | Period end (ISO datetime)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BOMAPI.GetObservationCompleteness(context.Background(), bomId).TenantId(tenantId).PeriodStart(periodStart).PeriodEnd(periodEnd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BOMAPI.GetObservationCompleteness``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObservationCompleteness`: ObservationOut
	fmt.Fprintf(os.Stdout, "Response from `BOMAPI.GetObservationCompleteness`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bomId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObservationCompletenessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** | Tenant UUID | 
 **periodStart** | **time.Time** | Period start (ISO datetime) | 
 **periodEnd** | **time.Time** | Period end (ISO datetime) | 

### Return type

[**ObservationOut**](ObservationOut.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBoms

> []BomOut ListBoms(ctx).TenantId(tenantId).FeatureKey(featureKey).Status(status).Limit(limit).Offset(offset).Execute()

List BOMs for tenant



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
	tenantId := "tenantId_example" // string | Tenant UUID
	featureKey := "featureKey_example" // string | Filter by feature_key (optional)
	status := "status_example" // string | Filter by status (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BOMAPI.ListBoms(context.Background()).TenantId(tenantId).FeatureKey(featureKey).Status(status).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BOMAPI.ListBoms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBoms`: []BomOut
	fmt.Fprintf(os.Stdout, "Response from `BOMAPI.ListBoms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBomsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Tenant UUID | 
 **featureKey** | **string** | Filter by feature_key | 
 **status** | **string** | Filter by status | 
 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** |  | [default to 0]

### Return type

[**[]BomOut**](BomOut.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SimulateCost

> SimulateOut SimulateCost(ctx, bomId).TenantId(tenantId).BillingEventCount(billingEventCount).Execute()

Simulate BOM cost with current rate_catalog rates



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
	bomId := "bomId_example" // string | 
	tenantId := "tenantId_example" // string | Tenant UUID
	billingEventCount := int32(56) // int32 | Number of billing events to project (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BOMAPI.SimulateCost(context.Background(), bomId).TenantId(tenantId).BillingEventCount(billingEventCount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BOMAPI.SimulateCost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SimulateCost`: SimulateOut
	fmt.Fprintf(os.Stdout, "Response from `BOMAPI.SimulateCost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bomId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSimulateCostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** | Tenant UUID | 
 **billingEventCount** | **int32** | Number of billing events to project | [default to 1]

### Return type

[**SimulateOut**](SimulateOut.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitForApproval

> BomOut SubmitForApproval(ctx, bomId).TenantId(tenantId).Execute()

Submit BOM for approval (draft → pending_approval)



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
	bomId := "bomId_example" // string | 
	tenantId := "tenantId_example" // string | Tenant UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BOMAPI.SubmitForApproval(context.Background(), bomId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BOMAPI.SubmitForApproval``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitForApproval`: BomOut
	fmt.Fprintf(os.Stdout, "Response from `BOMAPI.SubmitForApproval`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bomId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubmitForApprovalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** | Tenant UUID | 

### Return type

[**BomOut**](BomOut.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

