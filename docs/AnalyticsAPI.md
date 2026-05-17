# \AnalyticsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentActionAuditV1**](AnalyticsAPI.md#AgentActionAuditV1) | **Get** /v1/arc/analytics/agent-actions | Agent Action Audit
[**AgentPerformanceV1**](AnalyticsAPI.md#AgentPerformanceV1) | **Get** /v1/arc/analytics/agent-performance | Agent Performance
[**AgingReport**](AnalyticsAPI.md#AgingReport) | **Get** /v1/arc/analytics/aging | Aging Report
[**AgingTrend**](AnalyticsAPI.md#AgingTrend) | **Get** /v1/arc/analytics/aging/trend | Aging Trend
[**CeiReport**](AnalyticsAPI.md#CeiReport) | **Get** /v1/arc/analytics/cei | Cei Report
[**CreateAgingSnapshot**](AnalyticsAPI.md#CreateAgingSnapshot) | **Post** /v1/arc/analytics/aging/snapshot | Create Aging Snapshot
[**DashboardOverview**](AnalyticsAPI.md#DashboardOverview) | **Get** /v1/arc/analytics/overview | Dashboard Overview
[**DisputeSummaryV1**](AnalyticsAPI.md#DisputeSummaryV1) | **Get** /v1/arc/analytics/dispute-summary | Dispute Summary
[**DsoReport**](AnalyticsAPI.md#DsoReport) | **Get** /v1/arc/analytics/dso | Dso Report
[**PtpSummaryV1**](AnalyticsAPI.md#PtpSummaryV1) | **Get** /v1/arc/analytics/ptp-summary | Ptp Summary



## AgentActionAuditV1

> interface{} AgentActionAuditV1(ctx).XAPIKey(xAPIKey).AgentType(agentType).CaseId(caseId).Limit(limit).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Agent Action Audit



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
	xAPIKey := "xAPIKey_example" // string | 
	agentType := "agentType_example" // string |  (optional)
	caseId := "caseId_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.AgentActionAuditV1(context.Background()).XAPIKey(xAPIKey).AgentType(agentType).CaseId(caseId).Limit(limit).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.AgentActionAuditV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentActionAuditV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.AgentActionAuditV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgentActionAuditV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKey** | **string** |  | 
 **agentType** | **string** |  | 
 **caseId** | **string** |  | 
 **limit** | **int32** |  | [default to 100]
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


## AgentPerformanceV1

> interface{} AgentPerformanceV1(ctx).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Agent Performance



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
	xAPIKey := "xAPIKey_example" // string | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.AgentPerformanceV1(context.Background()).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.AgentPerformanceV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentPerformanceV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.AgentPerformanceV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgentPerformanceV1Request struct via the builder pattern


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


## AgingReport

> interface{} AgingReport(ctx).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Aging Report



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
	xAPIKey := "xAPIKey_example" // string | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.AgingReport(context.Background()).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.AgingReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgingReport`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.AgingReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgingReportRequest struct via the builder pattern


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


## AgingTrend

> interface{} AgingTrend(ctx).XAPIKey(xAPIKey).Days(days).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Aging Trend



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
	xAPIKey := "xAPIKey_example" // string | 
	days := int32(56) // int32 |  (optional) (default to 30)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.AgingTrend(context.Background()).XAPIKey(xAPIKey).Days(days).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.AgingTrend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgingTrend`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.AgingTrend`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgingTrendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKey** | **string** |  | 
 **days** | **int32** |  | [default to 30]
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


## CeiReport

> interface{} CeiReport(ctx).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Cei Report



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
	xAPIKey := "xAPIKey_example" // string | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.CeiReport(context.Background()).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.CeiReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CeiReport`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.CeiReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCeiReportRequest struct via the builder pattern


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


## CreateAgingSnapshot

> interface{} CreateAgingSnapshot(ctx).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Create Aging Snapshot



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
	xAPIKey := "xAPIKey_example" // string | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.CreateAgingSnapshot(context.Background()).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.CreateAgingSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAgingSnapshot`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.CreateAgingSnapshot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAgingSnapshotRequest struct via the builder pattern


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


## DashboardOverview

> DashboardOverviewResponse DashboardOverview(ctx).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Dashboard Overview



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
	xAPIKey := "xAPIKey_example" // string | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.DashboardOverview(context.Background()).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.DashboardOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DashboardOverview`: DashboardOverviewResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.DashboardOverview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDashboardOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**DashboardOverviewResponse**](DashboardOverviewResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisputeSummaryV1

> interface{} DisputeSummaryV1(ctx).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Dispute Summary



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
	xAPIKey := "xAPIKey_example" // string | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.DisputeSummaryV1(context.Background()).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.DisputeSummaryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisputeSummaryV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.DisputeSummaryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDisputeSummaryV1Request struct via the builder pattern


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


## DsoReport

> interface{} DsoReport(ctx).XAPIKey(xAPIKey).PeriodDays(periodDays).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Dso Report



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
	xAPIKey := "xAPIKey_example" // string | 
	periodDays := int32(56) // int32 |  (optional) (default to 90)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.DsoReport(context.Background()).XAPIKey(xAPIKey).PeriodDays(periodDays).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.DsoReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DsoReport`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.DsoReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDsoReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKey** | **string** |  | 
 **periodDays** | **int32** |  | [default to 90]
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


## PtpSummaryV1

> interface{} PtpSummaryV1(ctx).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Ptp Summary



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
	xAPIKey := "xAPIKey_example" // string | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.PtpSummaryV1(context.Background()).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.PtpSummaryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PtpSummaryV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.PtpSummaryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPtpSummaryV1Request struct via the builder pattern


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

