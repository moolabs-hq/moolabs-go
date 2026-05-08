# \SubscriptionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSubscription**](SubscriptionsAPI.md#GetSubscription) | **Get** /v1/subscriptions/{subscription_id} | Get Subscription
[**HandleLifecycleEvent**](SubscriptionsAPI.md#HandleLifecycleEvent) | **Post** /v1/subscriptions/lifecycle | Handle Lifecycle Event
[**SyncSubscription**](SubscriptionsAPI.md#SyncSubscription) | **Post** /v1/subscriptions/sync | Sync Subscription



## GetSubscription

> interface{} GetSubscription(ctx, subscriptionId).TenantId(tenantId).AsOf(asOf).Execute()

Get Subscription



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
	subscriptionId := "subscriptionId_example" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	asOf := "asOf_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.GetSubscription(context.Background(), subscriptionId).TenantId(tenantId).AsOf(asOf).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.GetSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubscription`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.GetSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **asOf** | **string** |  | 

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


## HandleLifecycleEvent

> interface{} HandleLifecycleEvent(ctx).TenantId(tenantId).LifecycleEventRequest(lifecycleEventRequest).Execute()

Handle Lifecycle Event



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	lifecycleEventRequest := *openapiclient.NewLifecycleEventRequest("SubscriptionId_example", "EventType_example", "EventId_example", "EffectiveAt_example") // LifecycleEventRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.HandleLifecycleEvent(context.Background()).TenantId(tenantId).LifecycleEventRequest(lifecycleEventRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.HandleLifecycleEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleLifecycleEvent`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.HandleLifecycleEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHandleLifecycleEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **lifecycleEventRequest** | [**LifecycleEventRequest**](LifecycleEventRequest.md) |  | 

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


## SyncSubscription

> interface{} SyncSubscription(ctx).TenantId(tenantId).SubscriptionSyncRequest(subscriptionSyncRequest).Execute()

Sync Subscription



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	subscriptionSyncRequest := *openapiclient.NewSubscriptionSyncRequest("SubscriptionId_example", "BillingAnchor_example", "PlanId_example", "ActiveFrom_example") // SubscriptionSyncRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SyncSubscription(context.Background()).TenantId(tenantId).SubscriptionSyncRequest(subscriptionSyncRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SyncSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncSubscription`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SyncSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **subscriptionSyncRequest** | [**SubscriptionSyncRequest**](SubscriptionSyncRequest.md) |  | 

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

