# \UnpricedAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListUnpricedEvents**](UnpricedAPI.md#ListUnpricedEvents) | **Get** /v1/unpriced/events | List Unpriced Events
[**ProcessEvents**](UnpricedAPI.md#ProcessEvents) | **Post** /v1/unpriced/process | Process Events
[**ResolveEvent**](UnpricedAPI.md#ResolveEvent) | **Post** /v1/unpriced/resolve/{usage_event_id} | Resolve Event



## ListUnpricedEvents

> interface{} ListUnpricedEvents(ctx).TenantId(tenantId).PoolId(poolId).Limit(limit).Execute()

List Unpriced Events



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter by tenant ID (optional)
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter by pool ID (optional)
	limit := int32(56) // int32 | Maximum events to return (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UnpricedAPI.ListUnpricedEvents(context.Background()).TenantId(tenantId).PoolId(poolId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UnpricedAPI.ListUnpricedEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUnpricedEvents`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `UnpricedAPI.ListUnpricedEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUnpricedEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Filter by tenant ID | 
 **poolId** | **string** | Filter by pool ID | 
 **limit** | **int32** | Maximum events to return | [default to 100]

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


## ProcessEvents

> interface{} ProcessEvents(ctx).TenantId(tenantId).PoolId(poolId).Limit(limit).MaxAttempts(maxAttempts).Execute()

Process Events



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter by tenant ID (optional)
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter by pool ID (optional)
	limit := int32(56) // int32 | Maximum events to process (optional) (default to 100)
	maxAttempts := int32(56) // int32 | Maximum retry attempts (optional) (default to 5)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UnpricedAPI.ProcessEvents(context.Background()).TenantId(tenantId).PoolId(poolId).Limit(limit).MaxAttempts(maxAttempts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UnpricedAPI.ProcessEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProcessEvents`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `UnpricedAPI.ProcessEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProcessEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Filter by tenant ID | 
 **poolId** | **string** | Filter by pool ID | 
 **limit** | **int32** | Maximum events to process | [default to 100]
 **maxAttempts** | **int32** | Maximum retry attempts | [default to 5]

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


## ResolveEvent

> interface{} ResolveEvent(ctx, usageEventId).MaxAttempts(maxAttempts).Execute()

Resolve Event



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
	maxAttempts := int32(56) // int32 | Maximum retry attempts (optional) (default to 5)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UnpricedAPI.ResolveEvent(context.Background(), usageEventId).MaxAttempts(maxAttempts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UnpricedAPI.ResolveEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResolveEvent`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `UnpricedAPI.ResolveEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**usageEventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResolveEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxAttempts** | **int32** | Maximum retry attempts | [default to 5]

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

