# \RatingAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProcessPendingEventsV1**](RatingAPI.md#ProcessPendingEventsV1) | **Post** /v1/rating/process-pending | Process Pending Events
[**RateEvent**](RatingAPI.md#RateEvent) | **Post** /v1/rating/rate/{usage_event_id} | Rate Event



## ProcessPendingEventsV1

> interface{} ProcessPendingEventsV1(ctx).Limit(limit).ShardAttemptOffset(shardAttemptOffset).TenantId(tenantId).PoolId(poolId).Execute()

Process Pending Events



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
	limit := int32(56) // int32 | Number of events to process (optional) (default to 10)
	shardAttemptOffset := int32(56) // int32 | Shard attempt offset (optional) (default to 0)
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter by tenant_id (for testing) (optional)
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter by pool_id (for testing) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RatingAPI.ProcessPendingEventsV1(context.Background()).Limit(limit).ShardAttemptOffset(shardAttemptOffset).TenantId(tenantId).PoolId(poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RatingAPI.ProcessPendingEventsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProcessPendingEventsV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `RatingAPI.ProcessPendingEventsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProcessPendingEventsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of events to process | [default to 10]
 **shardAttemptOffset** | **int32** | Shard attempt offset | [default to 0]
 **tenantId** | **string** | Filter by tenant_id (for testing) | 
 **poolId** | **string** | Filter by pool_id (for testing) | 

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


## RateEvent

> interface{} RateEvent(ctx, usageEventId).Execute()

Rate Event



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
	usageEventId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RatingAPI.RateEvent(context.Background(), usageEventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RatingAPI.RateEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RateEvent`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `RatingAPI.RateEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**usageEventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRateEventRequest struct via the builder pattern


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

