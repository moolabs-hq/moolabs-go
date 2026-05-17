# \AcuteAnalyticsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CostByCustomerApi**](AcuteAnalyticsAPI.md#CostByCustomerApi) | **Get** /api/v1/cost/analytics/by-customer | Cost By Customer
[**CostByFeatureApi**](AcuteAnalyticsAPI.md#CostByFeatureApi) | **Get** /api/v1/cost/analytics/by-feature | Cost By Feature
[**CostByModelApi**](AcuteAnalyticsAPI.md#CostByModelApi) | **Get** /api/v1/cost/analytics/by-model | Cost By Model
[**CostByPlanApi**](AcuteAnalyticsAPI.md#CostByPlanApi) | **Get** /api/v1/cost/analytics/by-plan | Cost By Plan
[**CostTrends**](AcuteAnalyticsAPI.md#CostTrends) | **Get** /api/v1/cost/analytics/trends | Cost Trends
[**IngestHealth**](AcuteAnalyticsAPI.md#IngestHealth) | **Get** /api/v1/cost/analytics/ingest/health | Ingest Health
[**Overview**](AcuteAnalyticsAPI.md#Overview) | **Get** /api/v1/cost/analytics/overview | Overview
[**TopConsumersApi**](AcuteAnalyticsAPI.md#TopConsumersApi) | **Get** /api/v1/cost/analytics/top-consumers | Top Consumers
[**UnitEconomicsApi**](AcuteAnalyticsAPI.md#UnitEconomicsApi) | **Get** /api/v1/cost/analytics/unit-economics | Unit Economics



## CostByCustomerApi

> []CustomerCostItem CostByCustomerApi(ctx).StartDate(startDate).EndDate(endDate).Limit(limit).XTenantId(xTenantId).Execute()

Cost By Customer



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
	startDate := time.Now() // string |  (optional)
	endDate := time.Now() // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	xTenantId := "xTenantId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AcuteAnalyticsAPI.CostByCustomerApi(context.Background()).StartDate(startDate).EndDate(endDate).Limit(limit).XTenantId(xTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AcuteAnalyticsAPI.CostByCustomerApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CostByCustomerApi`: []CustomerCostItem
	fmt.Fprintf(os.Stdout, "Response from `AcuteAnalyticsAPI.CostByCustomerApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCostByCustomerApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **limit** | **int32** |  | [default to 100]
 **xTenantId** | **string** |  | 

### Return type

[**[]CustomerCostItem**](CustomerCostItem.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CostByFeatureApi

> []FeatureCostItem CostByFeatureApi(ctx).StartDate(startDate).EndDate(endDate).Limit(limit).XTenantId(xTenantId).Execute()

Cost By Feature



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
	startDate := time.Now() // string |  (optional)
	endDate := time.Now() // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	xTenantId := "xTenantId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AcuteAnalyticsAPI.CostByFeatureApi(context.Background()).StartDate(startDate).EndDate(endDate).Limit(limit).XTenantId(xTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AcuteAnalyticsAPI.CostByFeatureApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CostByFeatureApi`: []FeatureCostItem
	fmt.Fprintf(os.Stdout, "Response from `AcuteAnalyticsAPI.CostByFeatureApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCostByFeatureApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **limit** | **int32** |  | [default to 100]
 **xTenantId** | **string** |  | 

### Return type

[**[]FeatureCostItem**](FeatureCostItem.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CostByModelApi

> []ModelCostItem CostByModelApi(ctx).StartDate(startDate).EndDate(endDate).XTenantId(xTenantId).Execute()

Cost By Model



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
	startDate := time.Now() // string |  (optional)
	endDate := time.Now() // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AcuteAnalyticsAPI.CostByModelApi(context.Background()).StartDate(startDate).EndDate(endDate).XTenantId(xTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AcuteAnalyticsAPI.CostByModelApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CostByModelApi`: []ModelCostItem
	fmt.Fprintf(os.Stdout, "Response from `AcuteAnalyticsAPI.CostByModelApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCostByModelApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **xTenantId** | **string** |  | 

### Return type

[**[]ModelCostItem**](ModelCostItem.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CostByPlanApi

> []map[string]interface{} CostByPlanApi(ctx).StartDate(startDate).EndDate(endDate).XTenantId(xTenantId).Execute()

Cost By Plan



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
	startDate := time.Now() // string |  (optional)
	endDate := time.Now() // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AcuteAnalyticsAPI.CostByPlanApi(context.Background()).StartDate(startDate).EndDate(endDate).XTenantId(xTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AcuteAnalyticsAPI.CostByPlanApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CostByPlanApi`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AcuteAnalyticsAPI.CostByPlanApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCostByPlanApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **xTenantId** | **string** |  | 

### Return type

[**[]map[string]interface{}**](map.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CostTrends

> []TrendDataPoint CostTrends(ctx).StartDate(startDate).EndDate(endDate).Period(period).XTenantId(xTenantId).Execute()

Cost Trends



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
	startDate := time.Now() // string |  (optional)
	endDate := time.Now() // string |  (optional)
	period := "period_example" // string |  (optional) (default to "day")
	xTenantId := "xTenantId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AcuteAnalyticsAPI.CostTrends(context.Background()).StartDate(startDate).EndDate(endDate).Period(period).XTenantId(xTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AcuteAnalyticsAPI.CostTrends``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CostTrends`: []TrendDataPoint
	fmt.Fprintf(os.Stdout, "Response from `AcuteAnalyticsAPI.CostTrends`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCostTrendsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **period** | **string** |  | [default to &quot;day&quot;]
 **xTenantId** | **string** |  | 

### Return type

[**[]TrendDataPoint**](TrendDataPoint.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IngestHealth

> IngestHealth IngestHealth(ctx).XTenantId(xTenantId).Execute()

Ingest Health



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
	xTenantId := "xTenantId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AcuteAnalyticsAPI.IngestHealth(context.Background()).XTenantId(xTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AcuteAnalyticsAPI.IngestHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IngestHealth`: IngestHealth
	fmt.Fprintf(os.Stdout, "Response from `AcuteAnalyticsAPI.IngestHealth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIngestHealthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTenantId** | **string** |  | 

### Return type

[**IngestHealth**](IngestHealth.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Overview

> OverviewResponse Overview(ctx).StartDate(startDate).EndDate(endDate).XTenantId(xTenantId).Execute()

Overview



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
	startDate := time.Now() // string |  (optional)
	endDate := time.Now() // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AcuteAnalyticsAPI.Overview(context.Background()).StartDate(startDate).EndDate(endDate).XTenantId(xTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AcuteAnalyticsAPI.Overview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Overview`: OverviewResponse
	fmt.Fprintf(os.Stdout, "Response from `AcuteAnalyticsAPI.Overview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **xTenantId** | **string** |  | 

### Return type

[**OverviewResponse**](OverviewResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TopConsumersApi

> []TopConsumer TopConsumersApi(ctx).StartDate(startDate).EndDate(endDate).TopN(topN).XTenantId(xTenantId).Execute()

Top Consumers



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
	startDate := time.Now() // string |  (optional)
	endDate := time.Now() // string |  (optional)
	topN := int32(56) // int32 |  (optional) (default to 10)
	xTenantId := "xTenantId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AcuteAnalyticsAPI.TopConsumersApi(context.Background()).StartDate(startDate).EndDate(endDate).TopN(topN).XTenantId(xTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AcuteAnalyticsAPI.TopConsumersApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TopConsumersApi`: []TopConsumer
	fmt.Fprintf(os.Stdout, "Response from `AcuteAnalyticsAPI.TopConsumersApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTopConsumersApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **topN** | **int32** |  | [default to 10]
 **xTenantId** | **string** |  | 

### Return type

[**[]TopConsumer**](TopConsumer.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnitEconomicsApi

> UnitEconomics UnitEconomicsApi(ctx).FeatureKey(featureKey).StartDate(startDate).EndDate(endDate).XTenantId(xTenantId).Execute()

Unit Economics



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
	featureKey := "featureKey_example" // string | Feature key to analyse
	startDate := time.Now() // string |  (optional)
	endDate := time.Now() // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AcuteAnalyticsAPI.UnitEconomicsApi(context.Background()).FeatureKey(featureKey).StartDate(startDate).EndDate(endDate).XTenantId(xTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AcuteAnalyticsAPI.UnitEconomicsApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnitEconomicsApi`: UnitEconomics
	fmt.Fprintf(os.Stdout, "Response from `AcuteAnalyticsAPI.UnitEconomicsApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnitEconomicsApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **featureKey** | **string** | Feature key to analyse | 
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **xTenantId** | **string** |  | 

### Return type

[**UnitEconomics**](UnitEconomics.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

