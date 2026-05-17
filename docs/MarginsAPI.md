# \MarginsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCompleteness**](MarginsAPI.md#GetCompleteness) | **Get** /api/v1/cost/margins/completeness | Financial completeness metric
[**GetCreditCostRatioApiV1**](MarginsAPI.md#GetCreditCostRatioApiV1) | **Get** /api/v1/cost/margins/credit-cost-ratio | Daily valued_burn / total_cost ratio time series
[**GetMarginByCustomerApi**](MarginsAPI.md#GetMarginByCustomerApi) | **Get** /api/v1/cost/margins/by-customer | Margin per customer
[**GetMarginByFeatureApi**](MarginsAPI.md#GetMarginByFeatureApi) | **Get** /api/v1/cost/margins/by-feature | Margin per feature
[**GetMarginSnapshot**](MarginsAPI.md#GetMarginSnapshot) | **Get** /api/v1/cost/margins/snapshot | Period margin snapshot with reconciliation breakdown
[**TriggerReconciliation**](MarginsAPI.md#TriggerReconciliation) | **Post** /api/v1/cost/margins/reconcile | Trigger manual reconciliation (operator)
[**TriggerSnapshotCompute**](MarginsAPI.md#TriggerSnapshotCompute) | **Post** /api/v1/cost/margins/snapshot/compute | Trigger snapshot computation (operator)



## GetCompleteness

> CompletenessResponse GetCompleteness(ctx).MatchGrade(matchGrade).IncludePartial(includePartial).PeriodStart(periodStart).PeriodEnd(periodEnd).Execute()

Financial completeness metric



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
	matchGrade := "matchGrade_example" // string | 'financial_only' or 'include_operational' (optional) (default to "financial_only")
	includePartial := true // bool | Include partially_reconciled billing events (optional) (default to false)
	periodStart := time.Now() // time.Time | ISO datetime start (inclusive) (optional)
	periodEnd := time.Now() // time.Time | ISO datetime end (exclusive) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginsAPI.GetCompleteness(context.Background()).MatchGrade(matchGrade).IncludePartial(includePartial).PeriodStart(periodStart).PeriodEnd(periodEnd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginsAPI.GetCompleteness``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompleteness`: CompletenessResponse
	fmt.Fprintf(os.Stdout, "Response from `MarginsAPI.GetCompleteness`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCompletenessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **matchGrade** | **string** | &#39;financial_only&#39; or &#39;include_operational&#39; | [default to &quot;financial_only&quot;]
 **includePartial** | **bool** | Include partially_reconciled billing events | [default to false]
 **periodStart** | **time.Time** | ISO datetime start (inclusive) | 
 **periodEnd** | **time.Time** | ISO datetime end (exclusive) | 

### Return type

[**CompletenessResponse**](CompletenessResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCreditCostRatioApiV1

> CreditCostRatioResponse GetCreditCostRatioApiV1(ctx).MatchGrade(matchGrade).IncludePartial(includePartial).PeriodStart(periodStart).PeriodEnd(periodEnd).Execute()

Daily valued_burn / total_cost ratio time series



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
	matchGrade := "matchGrade_example" // string | 'financial_only' or 'include_operational' (optional) (default to "financial_only")
	includePartial := true // bool | Include partially_reconciled billing events (optional) (default to false)
	periodStart := time.Now() // time.Time | ISO datetime start (inclusive) (optional)
	periodEnd := time.Now() // time.Time | ISO datetime end (exclusive) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginsAPI.GetCreditCostRatioApiV1(context.Background()).MatchGrade(matchGrade).IncludePartial(includePartial).PeriodStart(periodStart).PeriodEnd(periodEnd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginsAPI.GetCreditCostRatioApiV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCreditCostRatioApiV1`: CreditCostRatioResponse
	fmt.Fprintf(os.Stdout, "Response from `MarginsAPI.GetCreditCostRatioApiV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCreditCostRatioApiV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **matchGrade** | **string** | &#39;financial_only&#39; or &#39;include_operational&#39; | [default to &quot;financial_only&quot;]
 **includePartial** | **bool** | Include partially_reconciled billing events | [default to false]
 **periodStart** | **time.Time** | ISO datetime start (inclusive) | 
 **periodEnd** | **time.Time** | ISO datetime end (exclusive) | 

### Return type

[**CreditCostRatioResponse**](CreditCostRatioResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarginByCustomerApi

> CustomerMarginResponse GetMarginByCustomerApi(ctx).MatchGrade(matchGrade).IncludePartial(includePartial).PeriodStart(periodStart).PeriodEnd(periodEnd).Execute()

Margin per customer



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
	matchGrade := "matchGrade_example" // string | 'financial_only' or 'include_operational' (optional) (default to "financial_only")
	includePartial := true // bool | Include partially_reconciled billing events (optional) (default to false)
	periodStart := time.Now() // time.Time | ISO datetime start (inclusive) (optional)
	periodEnd := time.Now() // time.Time | ISO datetime end (exclusive) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginsAPI.GetMarginByCustomerApi(context.Background()).MatchGrade(matchGrade).IncludePartial(includePartial).PeriodStart(periodStart).PeriodEnd(periodEnd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginsAPI.GetMarginByCustomerApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarginByCustomerApi`: CustomerMarginResponse
	fmt.Fprintf(os.Stdout, "Response from `MarginsAPI.GetMarginByCustomerApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarginByCustomerApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **matchGrade** | **string** | &#39;financial_only&#39; or &#39;include_operational&#39; | [default to &quot;financial_only&quot;]
 **includePartial** | **bool** | Include partially_reconciled billing events | [default to false]
 **periodStart** | **time.Time** | ISO datetime start (inclusive) | 
 **periodEnd** | **time.Time** | ISO datetime end (exclusive) | 

### Return type

[**CustomerMarginResponse**](CustomerMarginResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarginByFeatureApi

> FeatureMarginResponse GetMarginByFeatureApi(ctx).MatchGrade(matchGrade).IncludePartial(includePartial).PeriodStart(periodStart).PeriodEnd(periodEnd).Execute()

Margin per feature



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
	matchGrade := "matchGrade_example" // string | 'financial_only' or 'include_operational' (optional) (default to "financial_only")
	includePartial := true // bool | Include partially_reconciled billing events (optional) (default to false)
	periodStart := time.Now() // time.Time | ISO datetime start (inclusive) (optional)
	periodEnd := time.Now() // time.Time | ISO datetime end (exclusive) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginsAPI.GetMarginByFeatureApi(context.Background()).MatchGrade(matchGrade).IncludePartial(includePartial).PeriodStart(periodStart).PeriodEnd(periodEnd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginsAPI.GetMarginByFeatureApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarginByFeatureApi`: FeatureMarginResponse
	fmt.Fprintf(os.Stdout, "Response from `MarginsAPI.GetMarginByFeatureApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarginByFeatureApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **matchGrade** | **string** | &#39;financial_only&#39; or &#39;include_operational&#39; | [default to &quot;financial_only&quot;]
 **includePartial** | **bool** | Include partially_reconciled billing events | [default to false]
 **periodStart** | **time.Time** | ISO datetime start (inclusive) | 
 **periodEnd** | **time.Time** | ISO datetime end (exclusive) | 

### Return type

[**FeatureMarginResponse**](FeatureMarginResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarginSnapshot

> MarginSnapshotResponse GetMarginSnapshot(ctx).PeriodType(periodType).MatchGrade(matchGrade).IncludePartial(includePartial).PeriodStart(periodStart).PeriodEnd(periodEnd).Execute()

Period margin snapshot with reconciliation breakdown



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
	periodType := "periodType_example" // string | 'daily', 'weekly', or 'monthly' (optional) (default to "daily")
	matchGrade := "matchGrade_example" // string | 'financial_only' or 'include_operational' (optional) (default to "financial_only")
	includePartial := true // bool | Include partially_reconciled billing events (optional) (default to false)
	periodStart := time.Now() // time.Time | ISO datetime start (inclusive) (optional)
	periodEnd := time.Now() // time.Time | ISO datetime end (exclusive) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginsAPI.GetMarginSnapshot(context.Background()).PeriodType(periodType).MatchGrade(matchGrade).IncludePartial(includePartial).PeriodStart(periodStart).PeriodEnd(periodEnd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginsAPI.GetMarginSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarginSnapshot`: MarginSnapshotResponse
	fmt.Fprintf(os.Stdout, "Response from `MarginsAPI.GetMarginSnapshot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarginSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **periodType** | **string** | &#39;daily&#39;, &#39;weekly&#39;, or &#39;monthly&#39; | [default to &quot;daily&quot;]
 **matchGrade** | **string** | &#39;financial_only&#39; or &#39;include_operational&#39; | [default to &quot;financial_only&quot;]
 **includePartial** | **bool** | Include partially_reconciled billing events | [default to false]
 **periodStart** | **time.Time** | ISO datetime start (inclusive) | 
 **periodEnd** | **time.Time** | ISO datetime end (exclusive) | 

### Return type

[**MarginSnapshotResponse**](MarginSnapshotResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TriggerReconciliation

> ReconcileResponse TriggerReconciliation(ctx).ReconcileRequest(reconcileRequest).Execute()

Trigger manual reconciliation (operator)



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
	reconcileRequest := *openapiclient.NewReconcileRequest() // ReconcileRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginsAPI.TriggerReconciliation(context.Background()).ReconcileRequest(reconcileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginsAPI.TriggerReconciliation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TriggerReconciliation`: ReconcileResponse
	fmt.Fprintf(os.Stdout, "Response from `MarginsAPI.TriggerReconciliation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTriggerReconciliationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reconcileRequest** | [**ReconcileRequest**](ReconcileRequest.md) |  | 

### Return type

[**ReconcileResponse**](ReconcileResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TriggerSnapshotCompute

> SnapshotComputeResponse TriggerSnapshotCompute(ctx).SnapshotComputeRequest(snapshotComputeRequest).Execute()

Trigger snapshot computation (operator)



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
	snapshotComputeRequest := *openapiclient.NewSnapshotComputeRequest(time.Now(), time.Now()) // SnapshotComputeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginsAPI.TriggerSnapshotCompute(context.Background()).SnapshotComputeRequest(snapshotComputeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginsAPI.TriggerSnapshotCompute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TriggerSnapshotCompute`: SnapshotComputeResponse
	fmt.Fprintf(os.Stdout, "Response from `MarginsAPI.TriggerSnapshotCompute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTriggerSnapshotComputeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **snapshotComputeRequest** | [**SnapshotComputeRequest**](SnapshotComputeRequest.md) |  | 

### Return type

[**SnapshotComputeResponse**](SnapshotComputeResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

