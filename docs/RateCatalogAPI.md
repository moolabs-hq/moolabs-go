# \RateCatalogAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkImportRates**](RateCatalogAPI.md#BulkImportRates) | **Post** /api/v1/rates/import | Bulk Import Rates
[**CreateRateEntry**](RateCatalogAPI.md#CreateRateEntry) | **Post** /api/v1/rates | Create Rate Entry
[**GetRateEntry**](RateCatalogAPI.md#GetRateEntry) | **Get** /api/v1/rates/{rate_id} | Get Rate Entry
[**ListCurrentRates**](RateCatalogAPI.md#ListCurrentRates) | **Get** /api/v1/rates/current | List Current Rates
[**RateHistory**](RateCatalogAPI.md#RateHistory) | **Get** /api/v1/rates/history | Rate History
[**RatesChangedSince**](RateCatalogAPI.md#RatesChangedSince) | **Get** /api/v1/rates/changes | Rates Changed Since
[**UpdateRateEntry**](RateCatalogAPI.md#UpdateRateEntry) | **Put** /api/v1/rates/{rate_id} | Update Rate Entry



## BulkImportRates

> RateBulkImportResponse BulkImportRates(ctx).RateBulkImportRequest(rateBulkImportRequest).Execute()

Bulk Import Rates



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
	rateBulkImportRequest := *openapiclient.NewRateBulkImportRequest("TenantId_example", []openapiclient.RateCatalogCreate{*openapiclient.NewRateCatalogCreate("TenantId_example", "Provider_example", "Model_example", "MetricType_example", *openapiclient.NewRatePerUnit(), time.Now())}) // RateBulkImportRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RateCatalogAPI.BulkImportRates(context.Background()).RateBulkImportRequest(rateBulkImportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RateCatalogAPI.BulkImportRates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkImportRates`: RateBulkImportResponse
	fmt.Fprintf(os.Stdout, "Response from `RateCatalogAPI.BulkImportRates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkImportRatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rateBulkImportRequest** | [**RateBulkImportRequest**](RateBulkImportRequest.md) |  | 

### Return type

[**RateBulkImportResponse**](RateBulkImportResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRateEntry

> RateCatalogResponse CreateRateEntry(ctx).RateCatalogCreate(rateCatalogCreate).Execute()

Create Rate Entry



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
	rateCatalogCreate := *openapiclient.NewRateCatalogCreate("TenantId_example", "Provider_example", "Model_example", "MetricType_example", *openapiclient.NewRatePerUnit(), time.Now()) // RateCatalogCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RateCatalogAPI.CreateRateEntry(context.Background()).RateCatalogCreate(rateCatalogCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RateCatalogAPI.CreateRateEntry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRateEntry`: RateCatalogResponse
	fmt.Fprintf(os.Stdout, "Response from `RateCatalogAPI.CreateRateEntry`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRateEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rateCatalogCreate** | [**RateCatalogCreate**](RateCatalogCreate.md) |  | 

### Return type

[**RateCatalogResponse**](RateCatalogResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRateEntry

> RateCatalogResponse GetRateEntry(ctx, rateId).Execute()

Get Rate Entry



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
	rateId := "rateId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RateCatalogAPI.GetRateEntry(context.Background(), rateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RateCatalogAPI.GetRateEntry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRateEntry`: RateCatalogResponse
	fmt.Fprintf(os.Stdout, "Response from `RateCatalogAPI.GetRateEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRateEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RateCatalogResponse**](RateCatalogResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCurrentRates

> []RateCatalogResponse ListCurrentRates(ctx).TenantId(tenantId).Provider(provider).Execute()

List Current Rates



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	provider := "provider_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RateCatalogAPI.ListCurrentRates(context.Background()).TenantId(tenantId).Provider(provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RateCatalogAPI.ListCurrentRates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCurrentRates`: []RateCatalogResponse
	fmt.Fprintf(os.Stdout, "Response from `RateCatalogAPI.ListCurrentRates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCurrentRatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **provider** | **string** |  | 

### Return type

[**[]RateCatalogResponse**](RateCatalogResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RateHistory

> []RateCatalogResponse RateHistory(ctx).TenantId(tenantId).Provider(provider).Model(model).MetricType(metricType).Execute()

Rate History



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	provider := "provider_example" // string |  (optional)
	model := "model_example" // string |  (optional)
	metricType := "metricType_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RateCatalogAPI.RateHistory(context.Background()).TenantId(tenantId).Provider(provider).Model(model).MetricType(metricType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RateCatalogAPI.RateHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RateHistory`: []RateCatalogResponse
	fmt.Fprintf(os.Stdout, "Response from `RateCatalogAPI.RateHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRateHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **provider** | **string** |  | 
 **model** | **string** |  | 
 **metricType** | **string** |  | 

### Return type

[**[]RateCatalogResponse**](RateCatalogResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RatesChangedSince

> []RateCatalogResponse RatesChangedSince(ctx).TenantId(tenantId).Since(since).Execute()

Rates Changed Since



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	since := time.Now() // time.Time | ISO 8601 timestamp; default = 30 days ago (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RateCatalogAPI.RatesChangedSince(context.Background()).TenantId(tenantId).Since(since).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RateCatalogAPI.RatesChangedSince``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RatesChangedSince`: []RateCatalogResponse
	fmt.Fprintf(os.Stdout, "Response from `RateCatalogAPI.RatesChangedSince`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRatesChangedSinceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **since** | **time.Time** | ISO 8601 timestamp; default &#x3D; 30 days ago | 

### Return type

[**[]RateCatalogResponse**](RateCatalogResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRateEntry

> RateCatalogResponse UpdateRateEntry(ctx, rateId).RateCatalogUpdate(rateCatalogUpdate).Execute()

Update Rate Entry



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
	rateId := "rateId_example" // string | 
	rateCatalogUpdate := *openapiclient.NewRateCatalogUpdate() // RateCatalogUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RateCatalogAPI.UpdateRateEntry(context.Background(), rateId).RateCatalogUpdate(rateCatalogUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RateCatalogAPI.UpdateRateEntry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRateEntry`: RateCatalogResponse
	fmt.Fprintf(os.Stdout, "Response from `RateCatalogAPI.UpdateRateEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRateEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rateCatalogUpdate** | [**RateCatalogUpdate**](RateCatalogUpdate.md) |  | 

### Return type

[**RateCatalogResponse**](RateCatalogResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

