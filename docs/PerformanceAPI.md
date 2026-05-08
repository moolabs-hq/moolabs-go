# \PerformanceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExplainQuery**](PerformanceAPI.md#ExplainQuery) | **Post** /v1/performance/queries/explain | Explain Query
[**GetConnectionPoolInfoV1**](PerformanceAPI.md#GetConnectionPoolInfoV1) | **Get** /v1/performance/connection-pool | Get Connection Pool Info
[**GetIndexUsage**](PerformanceAPI.md#GetIndexUsage) | **Get** /v1/performance/indexes/usage | Get Index Usage
[**GetMissingIndexRecommendations**](PerformanceAPI.md#GetMissingIndexRecommendations) | **Get** /v1/performance/indexes/missing | Get Missing Index Recommendations
[**GetPerformanceReportEndpoint**](PerformanceAPI.md#GetPerformanceReportEndpoint) | **Get** /v1/performance/report | Get Performance Report Endpoint
[**GetSlowQueries**](PerformanceAPI.md#GetSlowQueries) | **Get** /v1/performance/queries/slow | Get Slow Queries
[**GetTableSizeAnalysis**](PerformanceAPI.md#GetTableSizeAnalysis) | **Get** /v1/performance/tables/sizes | Get Table Size Analysis
[**WarmRateCardCacheV1**](PerformanceAPI.md#WarmRateCardCacheV1) | **Post** /v1/performance/cache/warm/rate-cards | Warm Rate Card Cache
[**WarmWalletCache**](PerformanceAPI.md#WarmWalletCache) | **Post** /v1/performance/cache/warm/wallets | Warm Wallet Cache



## ExplainQuery

> interface{} ExplainQuery(ctx).QuerySql(querySql).Execute()

Explain Query



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
	querySql := "querySql_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PerformanceAPI.ExplainQuery(context.Background()).QuerySql(querySql).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PerformanceAPI.ExplainQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExplainQuery`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PerformanceAPI.ExplainQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExplainQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **querySql** | **string** |  | 

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


## GetConnectionPoolInfoV1

> interface{} GetConnectionPoolInfoV1(ctx).Execute()

Get Connection Pool Info



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
	resp, r, err := apiClient.PerformanceAPI.GetConnectionPoolInfoV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PerformanceAPI.GetConnectionPoolInfoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectionPoolInfoV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PerformanceAPI.GetConnectionPoolInfoV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionPoolInfoV1Request struct via the builder pattern


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


## GetIndexUsage

> interface{} GetIndexUsage(ctx).Schema(schema).Execute()

Get Index Usage



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
	schema := "schema_example" // string | Schema name (optional) (default to "billing")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PerformanceAPI.GetIndexUsage(context.Background()).Schema(schema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PerformanceAPI.GetIndexUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIndexUsage`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PerformanceAPI.GetIndexUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIndexUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schema** | **string** | Schema name | [default to &quot;billing&quot;]

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


## GetMissingIndexRecommendations

> interface{} GetMissingIndexRecommendations(ctx).Schema(schema).Execute()

Get Missing Index Recommendations



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
	schema := "schema_example" // string | Schema name (optional) (default to "billing")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PerformanceAPI.GetMissingIndexRecommendations(context.Background()).Schema(schema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PerformanceAPI.GetMissingIndexRecommendations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMissingIndexRecommendations`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PerformanceAPI.GetMissingIndexRecommendations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMissingIndexRecommendationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schema** | **string** | Schema name | [default to &quot;billing&quot;]

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


## GetPerformanceReportEndpoint

> interface{} GetPerformanceReportEndpoint(ctx).Schema(schema).Execute()

Get Performance Report Endpoint



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
	schema := "schema_example" // string | Schema name (optional) (default to "billing")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PerformanceAPI.GetPerformanceReportEndpoint(context.Background()).Schema(schema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PerformanceAPI.GetPerformanceReportEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPerformanceReportEndpoint`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PerformanceAPI.GetPerformanceReportEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPerformanceReportEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schema** | **string** | Schema name | [default to &quot;billing&quot;]

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


## GetSlowQueries

> interface{} GetSlowQueries(ctx).MinDurationMs(minDurationMs).Limit(limit).Execute()

Get Slow Queries



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
	minDurationMs := float32(8.14) // float32 | Minimum query duration in milliseconds (optional) (default to 100.0)
	limit := int32(56) // int32 | Maximum number of results (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PerformanceAPI.GetSlowQueries(context.Background()).MinDurationMs(minDurationMs).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PerformanceAPI.GetSlowQueries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSlowQueries`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PerformanceAPI.GetSlowQueries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSlowQueriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **minDurationMs** | **float32** | Minimum query duration in milliseconds | [default to 100.0]
 **limit** | **int32** | Maximum number of results | [default to 20]

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


## GetTableSizeAnalysis

> interface{} GetTableSizeAnalysis(ctx).Schema(schema).Execute()

Get Table Size Analysis



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
	schema := "schema_example" // string | Schema name (optional) (default to "billing")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PerformanceAPI.GetTableSizeAnalysis(context.Background()).Schema(schema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PerformanceAPI.GetTableSizeAnalysis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTableSizeAnalysis`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PerformanceAPI.GetTableSizeAnalysis`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTableSizeAnalysisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schema** | **string** | Schema name | [default to &quot;billing&quot;]

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


## WarmRateCardCacheV1

> interface{} WarmRateCardCacheV1(ctx).TenantId(tenantId).Limit(limit).Execute()

Warm Rate Card Cache



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional tenant filter (optional)
	limit := int32(56) // int32 | Maximum number of rate cards to warm (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PerformanceAPI.WarmRateCardCacheV1(context.Background()).TenantId(tenantId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PerformanceAPI.WarmRateCardCacheV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WarmRateCardCacheV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PerformanceAPI.WarmRateCardCacheV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWarmRateCardCacheV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Optional tenant filter | 
 **limit** | **int32** | Maximum number of rate cards to warm | [default to 50]

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


## WarmWalletCache

> interface{} WarmWalletCache(ctx).TenantId(tenantId).PoolId(poolId).Limit(limit).Execute()

Warm Wallet Cache



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional tenant filter (optional)
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional pool filter (optional)
	limit := int32(56) // int32 | Maximum number of wallets to warm (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PerformanceAPI.WarmWalletCache(context.Background()).TenantId(tenantId).PoolId(poolId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PerformanceAPI.WarmWalletCache``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WarmWalletCache`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PerformanceAPI.WarmWalletCache`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWarmWalletCacheRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Optional tenant filter | 
 **poolId** | **string** | Optional pool filter | 
 **limit** | **int32** | Maximum number of wallets to warm | [default to 100]

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

