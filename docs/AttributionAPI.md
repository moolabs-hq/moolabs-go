# \AttributionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGateStatus**](AttributionAPI.md#GetGateStatus) | **Get** /api/v1/attribution/shadow/{run_id}/gate | Evaluate go/no-go gate for a shadow run
[**GetShadowRunMetrics**](AttributionAPI.md#GetShadowRunMetrics) | **Get** /api/v1/attribution/shadow/{run_id}/metrics | Get WAPE, coverage, and algorithm breakdown for a shadow run
[**ListAlgorithms**](AttributionAPI.md#ListAlgorithms) | **Get** /api/v1/attribution/algorithms | List all 12 attribution algorithms
[**ListShadowRuns**](AttributionAPI.md#ListShadowRuns) | **Get** /api/v1/attribution/shadow/runs | List all shadow attribution runs for tenant
[**PromoteShadowRun**](AttributionAPI.md#PromoteShadowRun) | **Post** /api/v1/attribution/shadow/{run_id}/promote | Promote shadow attribution to production (gate must pass)
[**RunShadowAttribution**](AttributionAPI.md#RunShadowAttribution) | **Post** /api/v1/attribution/shadow/run | Trigger shadow attribution run for a period



## GetGateStatus

> GateStatusResponse GetGateStatus(ctx, runId).TenantId(tenantId).Execute()

Evaluate go/no-go gate for a shadow run



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
	runId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "tenantId_example" // string | Tenant UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttributionAPI.GetGateStatus(context.Background(), runId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttributionAPI.GetGateStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGateStatus`: GateStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `AttributionAPI.GetGateStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGateStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** | Tenant UUID | 

### Return type

[**GateStatusResponse**](GateStatusResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShadowRunMetrics

> ShadowRunMetrics GetShadowRunMetrics(ctx, runId).TenantId(tenantId).Execute()

Get WAPE, coverage, and algorithm breakdown for a shadow run



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
	runId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "tenantId_example" // string | Tenant UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttributionAPI.GetShadowRunMetrics(context.Background(), runId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttributionAPI.GetShadowRunMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetShadowRunMetrics`: ShadowRunMetrics
	fmt.Fprintf(os.Stdout, "Response from `AttributionAPI.GetShadowRunMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShadowRunMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** | Tenant UUID | 

### Return type

[**ShadowRunMetrics**](ShadowRunMetrics.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAlgorithms

> []AlgorithmInfo ListAlgorithms(ctx).Execute()

List all 12 attribution algorithms



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
	resp, r, err := apiClient.AttributionAPI.ListAlgorithms(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttributionAPI.ListAlgorithms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAlgorithms`: []AlgorithmInfo
	fmt.Fprintf(os.Stdout, "Response from `AttributionAPI.ListAlgorithms`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAlgorithmsRequest struct via the builder pattern


### Return type

[**[]AlgorithmInfo**](AlgorithmInfo.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListShadowRuns

> []ShadowRunListItem ListShadowRuns(ctx).TenantId(tenantId).Limit(limit).Offset(offset).Execute()

List all shadow attribution runs for tenant



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
	limit := int32(56) // int32 |  (optional) (default to 50)
	offset := int32(56) // int32 | Pagination offset (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttributionAPI.ListShadowRuns(context.Background()).TenantId(tenantId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttributionAPI.ListShadowRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListShadowRuns`: []ShadowRunListItem
	fmt.Fprintf(os.Stdout, "Response from `AttributionAPI.ListShadowRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListShadowRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Tenant UUID | 
 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** | Pagination offset | [default to 0]

### Return type

[**[]ShadowRunListItem**](ShadowRunListItem.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PromoteShadowRun

> PromoteResponse PromoteShadowRun(ctx, runId).TenantId(tenantId).Execute()

Promote shadow attribution to production (gate must pass)



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
	runId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "tenantId_example" // string | Tenant UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttributionAPI.PromoteShadowRun(context.Background(), runId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttributionAPI.PromoteShadowRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PromoteShadowRun`: PromoteResponse
	fmt.Fprintf(os.Stdout, "Response from `AttributionAPI.PromoteShadowRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPromoteShadowRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** | Tenant UUID | 

### Return type

[**PromoteResponse**](PromoteResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunShadowAttribution

> ShadowRunSummary RunShadowAttribution(ctx).ShadowRunRequest(shadowRunRequest).Execute()

Trigger shadow attribution run for a period



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
	shadowRunRequest := *openapiclient.NewShadowRunRequest("TenantId_example", time.Now(), time.Now()) // ShadowRunRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttributionAPI.RunShadowAttribution(context.Background()).ShadowRunRequest(shadowRunRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttributionAPI.RunShadowAttribution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RunShadowAttribution`: ShadowRunSummary
	fmt.Fprintf(os.Stdout, "Response from `AttributionAPI.RunShadowAttribution`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRunShadowAttributionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shadowRunRequest** | [**ShadowRunRequest**](ShadowRunRequest.md) |  | 

### Return type

[**ShadowRunSummary**](ShadowRunSummary.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

