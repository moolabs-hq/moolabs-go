# \CostEventsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IngestEvent**](CostEventsAPI.md#IngestEvent) | **Post** /api/v1/cost/ingest | Ingest Cost Event
[**IngestEventsBatch**](CostEventsAPI.md#IngestEventsBatch) | **Post** /api/v1/cost/ingest/batch | Batch Ingest Cost Events
[**ListCostEventSummariesByUsageEventApiV1**](CostEventsAPI.md#ListCostEventSummariesByUsageEventApiV1) | **Post** /api/v1/cost/events/by-usage-event/summary | List Cost Event Summaries By Usage Event
[**ListCostEventsByUsageEventApiV1**](CostEventsAPI.md#ListCostEventsByUsageEventApiV1) | **Get** /api/v1/cost/events/by-usage-event/{usage_event_id} | List Cost Events By Usage Event
[**ListCostEventsByUsageEventQueryApiV1**](CostEventsAPI.md#ListCostEventsByUsageEventQueryApiV1) | **Get** /api/v1/cost/events/by-usage-event | List Cost Events By Usage Event Query
[**SubmitAdjustment**](CostEventsAPI.md#SubmitAdjustment) | **Post** /api/v1/cost/adjustments | Submit Adjustment



## IngestEvent

> CostEventResponse IngestEvent(ctx).CostEventIngest(costEventIngest).Execute()

Ingest Cost Event

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
	costEventIngest := *openapiclient.NewCostEventIngest("TenantId_example", "IdempotencyKey_example", "Provider_example", *openapiclient.NewObservedTotalCost(), "Currency_example", *openapiclient.NewReportingTotalCost(), time.Now()) // CostEventIngest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CostEventsAPI.IngestEvent(context.Background()).CostEventIngest(costEventIngest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CostEventsAPI.IngestEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IngestEvent`: CostEventResponse
	fmt.Fprintf(os.Stdout, "Response from `CostEventsAPI.IngestEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIngestEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **costEventIngest** | [**CostEventIngest**](CostEventIngest.md) |  | 

### Return type

[**CostEventResponse**](CostEventResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IngestEventsBatch

> interface{} IngestEventsBatch(ctx).BatchIngestRequest(batchIngestRequest).Execute()

Batch Ingest Cost Events



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
	batchIngestRequest := *openapiclient.NewBatchIngestRequest([]openapiclient.CostEventIngest{*openapiclient.NewCostEventIngest("TenantId_example", "IdempotencyKey_example", "Provider_example", *openapiclient.NewObservedTotalCost(), "Currency_example", *openapiclient.NewReportingTotalCost(), time.Now())}) // BatchIngestRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CostEventsAPI.IngestEventsBatch(context.Background()).BatchIngestRequest(batchIngestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CostEventsAPI.IngestEventsBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IngestEventsBatch`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CostEventsAPI.IngestEventsBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIngestEventsBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchIngestRequest** | [**BatchIngestRequest**](BatchIngestRequest.md) |  | 

### Return type

**interface{}**

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCostEventSummariesByUsageEventApiV1

> CostEventSummaryResponse ListCostEventSummariesByUsageEventApiV1(ctx).XTenantID(xTenantID).CostEventSummaryRequest(costEventSummaryRequest).Execute()

List Cost Event Summaries By Usage Event



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
	xTenantID := "xTenantID_example" // string | Tenant UUID
	costEventSummaryRequest := *openapiclient.NewCostEventSummaryRequest([]string{"UsageEventIds_example"}) // CostEventSummaryRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CostEventsAPI.ListCostEventSummariesByUsageEventApiV1(context.Background()).XTenantID(xTenantID).CostEventSummaryRequest(costEventSummaryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CostEventsAPI.ListCostEventSummariesByUsageEventApiV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCostEventSummariesByUsageEventApiV1`: CostEventSummaryResponse
	fmt.Fprintf(os.Stdout, "Response from `CostEventsAPI.ListCostEventSummariesByUsageEventApiV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCostEventSummariesByUsageEventApiV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTenantID** | **string** | Tenant UUID | 
 **costEventSummaryRequest** | [**CostEventSummaryRequest**](CostEventSummaryRequest.md) |  | 

### Return type

[**CostEventSummaryResponse**](CostEventSummaryResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCostEventsByUsageEventApiV1

> []CostEventDetailResponse ListCostEventsByUsageEventApiV1(ctx, usageEventId).XTenantID(xTenantID).Execute()

List Cost Events By Usage Event



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
	usageEventId := "usageEventId_example" // string | 
	xTenantID := "xTenantID_example" // string | Tenant UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CostEventsAPI.ListCostEventsByUsageEventApiV1(context.Background(), usageEventId).XTenantID(xTenantID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CostEventsAPI.ListCostEventsByUsageEventApiV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCostEventsByUsageEventApiV1`: []CostEventDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `CostEventsAPI.ListCostEventsByUsageEventApiV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**usageEventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCostEventsByUsageEventApiV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTenantID** | **string** | Tenant UUID | 

### Return type

[**[]CostEventDetailResponse**](CostEventDetailResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCostEventsByUsageEventQueryApiV1

> []CostEventDetailResponse ListCostEventsByUsageEventQueryApiV1(ctx).UsageEventId(usageEventId).XTenantID(xTenantID).Execute()

List Cost Events By Usage Event Query



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
	usageEventId := "usageEventId_example" // string | 
	xTenantID := "xTenantID_example" // string | Tenant UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CostEventsAPI.ListCostEventsByUsageEventQueryApiV1(context.Background()).UsageEventId(usageEventId).XTenantID(xTenantID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CostEventsAPI.ListCostEventsByUsageEventQueryApiV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCostEventsByUsageEventQueryApiV1`: []CostEventDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `CostEventsAPI.ListCostEventsByUsageEventQueryApiV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCostEventsByUsageEventQueryApiV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **usageEventId** | **string** |  | 
 **xTenantID** | **string** | Tenant UUID | 

### Return type

[**[]CostEventDetailResponse**](CostEventDetailResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitAdjustment

> CostAdjustmentResponse SubmitAdjustment(ctx).CostAdjustmentCreate(costAdjustmentCreate).Execute()

Submit Adjustment



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
	costAdjustmentCreate := *openapiclient.NewCostAdjustmentCreate("TenantId_example", "AdjustmentType_example", *openapiclient.NewAdjustmentAmount(), "Reason_example") // CostAdjustmentCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CostEventsAPI.SubmitAdjustment(context.Background()).CostAdjustmentCreate(costAdjustmentCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CostEventsAPI.SubmitAdjustment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitAdjustment`: CostAdjustmentResponse
	fmt.Fprintf(os.Stdout, "Response from `CostEventsAPI.SubmitAdjustment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitAdjustmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **costAdjustmentCreate** | [**CostAdjustmentCreate**](CostAdjustmentCreate.md) |  | 

### Return type

[**CostAdjustmentResponse**](CostAdjustmentResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

