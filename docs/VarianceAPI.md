# \VarianceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ComputeVariance**](VarianceAPI.md#ComputeVariance) | **Post** /api/v1/variance/compute | Trigger variance computation (operator)
[**GetBomVariance**](VarianceAPI.md#GetBomVariance) | **Get** /api/v1/variance/bom/{bom_id} | BOM-level variance: standard vs actual
[**GetVarianceDecomposition**](VarianceAPI.md#GetVarianceDecomposition) | **Get** /api/v1/variance/decomposition | Full four-way variance decomposition for feature+period
[**GetVarianceSummary**](VarianceAPI.md#GetVarianceSummary) | **Get** /api/v1/variance | Variance summary for all features in period



## ComputeVariance

> ComputeVarianceResponse ComputeVariance(ctx).ComputeVarianceRequest(computeVarianceRequest).Execute()

Trigger variance computation (operator)



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
	computeVarianceRequest := *openapiclient.NewComputeVarianceRequest("TenantId_example", "FeatureKey_example", time.Now(), time.Now(), "BomTemplateId_example") // ComputeVarianceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VarianceAPI.ComputeVariance(context.Background()).ComputeVarianceRequest(computeVarianceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VarianceAPI.ComputeVariance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ComputeVariance`: ComputeVarianceResponse
	fmt.Fprintf(os.Stdout, "Response from `VarianceAPI.ComputeVariance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiComputeVarianceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **computeVarianceRequest** | [**ComputeVarianceRequest**](ComputeVarianceRequest.md) |  | 

### Return type

[**ComputeVarianceResponse**](ComputeVarianceResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBomVariance

> BomVarianceOut GetBomVariance(ctx, bomId).TenantId(tenantId).PeriodStart(periodStart).PeriodEnd(periodEnd).Execute()

BOM-level variance: standard vs actual



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
	resp, r, err := apiClient.VarianceAPI.GetBomVariance(context.Background(), bomId).TenantId(tenantId).PeriodStart(periodStart).PeriodEnd(periodEnd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VarianceAPI.GetBomVariance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBomVariance`: BomVarianceOut
	fmt.Fprintf(os.Stdout, "Response from `VarianceAPI.GetBomVariance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bomId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBomVarianceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** | Tenant UUID | 
 **periodStart** | **time.Time** | Period start (ISO datetime) | 
 **periodEnd** | **time.Time** | Period end (ISO datetime) | 

### Return type

[**BomVarianceOut**](BomVarianceOut.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVarianceDecomposition

> VarianceDecompositionOut GetVarianceDecomposition(ctx).TenantId(tenantId).FeatureKey(featureKey).PeriodStart(periodStart).PeriodEnd(periodEnd).BomTemplateId(bomTemplateId).Execute()

Full four-way variance decomposition for feature+period



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
	tenantId := "tenantId_example" // string | Tenant UUID
	featureKey := "featureKey_example" // string | Feature key
	periodStart := time.Now() // time.Time | Period start (ISO datetime)
	periodEnd := time.Now() // time.Time | Period end (ISO datetime)
	bomTemplateId := "bomTemplateId_example" // string | BOM template UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VarianceAPI.GetVarianceDecomposition(context.Background()).TenantId(tenantId).FeatureKey(featureKey).PeriodStart(periodStart).PeriodEnd(periodEnd).BomTemplateId(bomTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VarianceAPI.GetVarianceDecomposition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVarianceDecomposition`: VarianceDecompositionOut
	fmt.Fprintf(os.Stdout, "Response from `VarianceAPI.GetVarianceDecomposition`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVarianceDecompositionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Tenant UUID | 
 **featureKey** | **string** | Feature key | 
 **periodStart** | **time.Time** | Period start (ISO datetime) | 
 **periodEnd** | **time.Time** | Period end (ISO datetime) | 
 **bomTemplateId** | **string** | BOM template UUID | 

### Return type

[**VarianceDecompositionOut**](VarianceDecompositionOut.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVarianceSummary

> VarianceSummaryResponse GetVarianceSummary(ctx).TenantId(tenantId).PeriodStart(periodStart).PeriodEnd(periodEnd).Execute()

Variance summary for all features in period



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
	tenantId := "tenantId_example" // string | Tenant UUID
	periodStart := time.Now() // time.Time | Period start (ISO datetime)
	periodEnd := time.Now() // time.Time | Period end (ISO datetime)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VarianceAPI.GetVarianceSummary(context.Background()).TenantId(tenantId).PeriodStart(periodStart).PeriodEnd(periodEnd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VarianceAPI.GetVarianceSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVarianceSummary`: VarianceSummaryResponse
	fmt.Fprintf(os.Stdout, "Response from `VarianceAPI.GetVarianceSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVarianceSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Tenant UUID | 
 **periodStart** | **time.Time** | Period start (ISO datetime) | 
 **periodEnd** | **time.Time** | Period end (ISO datetime) | 

### Return type

[**VarianceSummaryResponse**](VarianceSummaryResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

