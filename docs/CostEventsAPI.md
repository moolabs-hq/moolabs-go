# \CostEventsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IngestEvent**](CostEventsAPI.md#IngestEvent) | **Post** /api/v1/cost/ingest | Ingest Cost Event
[**IngestEventsBatch**](CostEventsAPI.md#IngestEventsBatch) | **Post** /api/v1/cost/ingest/batch | Batch Ingest Cost Events
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

[APIKeyHeader](../README.md#APIKeyHeader)

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

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
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

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

