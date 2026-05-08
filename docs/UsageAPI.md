# \UsageAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUsageEvent**](UsageAPI.md#GetUsageEvent) | **Get** /v1/usage/{event_id} | Get Usage Event
[**ListUsageEvents**](UsageAPI.md#ListUsageEvents) | **Get** /v1/usage/ | List Usage Events
[**RecordUsage**](UsageAPI.md#RecordUsage) | **Post** /v1/usage/record | Record Usage



## GetUsageEvent

> interface{} GetUsageEvent(ctx, eventId).Execute()

Get Usage Event



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
	eventId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageAPI.GetUsageEvent(context.Background(), eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.GetUsageEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsageEvent`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.GetUsageEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ListUsageEvents

> interface{} ListUsageEvents(ctx).TenantId(tenantId).PoolId(poolId).RatingStatus(ratingStatus).FeatureKey(featureKey).MeterSlug(meterSlug).SubjectId(subjectId).FromTime(fromTime).ToTime(toTime).Limit(limit).Offset(offset).Execute()

List Usage Events



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier (optional if subject_id provided) (optional)
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Pool identifier (optional if subject_id provided) (optional)
	ratingStatus := "ratingStatus_example" // string | Filter by rating status (optional)
	featureKey := "featureKey_example" // string | Filter by feature key (optional)
	meterSlug := "meterSlug_example" // string | Filter by meter slug (optional)
	subjectId := "subjectId_example" // string | Filter by subject ID (auto-resolves tenant_id/pool_id if not provided) (optional)
	fromTime := "fromTime_example" // string | Filter events from this time (ISO 8601, e.g., 2026-02-09T00:48:00Z) (optional)
	toTime := "toTime_example" // string | Filter events to this time (ISO 8601, e.g., 2026-02-09T00:50:00Z) (optional)
	limit := int32(56) // int32 | Maximum number of events to return (optional) (default to 100)
	offset := int32(56) // int32 | Offset for pagination (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageAPI.ListUsageEvents(context.Background()).TenantId(tenantId).PoolId(poolId).RatingStatus(ratingStatus).FeatureKey(featureKey).MeterSlug(meterSlug).SubjectId(subjectId).FromTime(fromTime).ToTime(toTime).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.ListUsageEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUsageEvents`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.ListUsageEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsageEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Tenant identifier (optional if subject_id provided) | 
 **poolId** | **string** | Pool identifier (optional if subject_id provided) | 
 **ratingStatus** | **string** | Filter by rating status | 
 **featureKey** | **string** | Filter by feature key | 
 **meterSlug** | **string** | Filter by meter slug | 
 **subjectId** | **string** | Filter by subject ID (auto-resolves tenant_id/pool_id if not provided) | 
 **fromTime** | **string** | Filter events from this time (ISO 8601, e.g., 2026-02-09T00:48:00Z) | 
 **toTime** | **string** | Filter events to this time (ISO 8601, e.g., 2026-02-09T00:50:00Z) | 
 **limit** | **int32** | Maximum number of events to return | [default to 100]
 **offset** | **int32** | Offset for pagination | [default to 0]

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


## RecordUsage

> interface{} RecordUsage(ctx).RecordUsageRequest(recordUsageRequest).Execute()

Record Usage



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
	recordUsageRequest := *openapiclient.NewRecordUsageRequest("TenantId_example", "SubjectId_example", "PoolId_example", "FeatureKey_example", "UsageEventId_example", time.Now(), map[string]interface{}{"key": interface{}(123)}) // RecordUsageRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageAPI.RecordUsage(context.Background()).RecordUsageRequest(recordUsageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.RecordUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordUsage`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.RecordUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRecordUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recordUsageRequest** | [**RecordUsageRequest**](RecordUsageRequest.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

