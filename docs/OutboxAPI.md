# \OutboxAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListOutboxEventsEndpoint**](OutboxAPI.md#ListOutboxEventsEndpoint) | **Get** /v1/outbox/events | List Outbox Events Endpoint
[**ProcessOutboxEndpoint**](OutboxAPI.md#ProcessOutboxEndpoint) | **Post** /v1/outbox/process | Process Outbox Endpoint



## ListOutboxEventsEndpoint

> interface{} ListOutboxEventsEndpoint(ctx).TenantId(tenantId).EventType(eventType).Status(status).Limit(limit).Execute()

List Outbox Events Endpoint



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
	eventType := "eventType_example" // string | Filter by event type (optional)
	status := "status_example" // string | Filter by status: pending, delivered, failed (optional)
	limit := int32(56) // int32 | Maximum events to return (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutboxAPI.ListOutboxEventsEndpoint(context.Background()).TenantId(tenantId).EventType(eventType).Status(status).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutboxAPI.ListOutboxEventsEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOutboxEventsEndpoint`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `OutboxAPI.ListOutboxEventsEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOutboxEventsEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Filter by tenant ID | 
 **eventType** | **string** | Filter by event type | 
 **status** | **string** | Filter by status: pending, delivered, failed | 
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


## ProcessOutboxEndpoint

> interface{} ProcessOutboxEndpoint(ctx).TenantId(tenantId).BatchSize(batchSize).MaxAttempts(maxAttempts).Execute()

Process Outbox Endpoint



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
	batchSize := int32(56) // int32 | Maximum events to process (optional) (default to 100)
	maxAttempts := int32(56) // int32 | Maximum retry attempts (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutboxAPI.ProcessOutboxEndpoint(context.Background()).TenantId(tenantId).BatchSize(batchSize).MaxAttempts(maxAttempts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutboxAPI.ProcessOutboxEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProcessOutboxEndpoint`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `OutboxAPI.ProcessOutboxEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProcessOutboxEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Optional tenant filter | 
 **batchSize** | **int32** | Maximum events to process | [default to 100]
 **maxAttempts** | **int32** | Maximum retry attempts | [default to 10]

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

