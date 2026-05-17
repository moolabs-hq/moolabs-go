# \InternalSubscriptionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SyncSubscriptionPost**](InternalSubscriptionsAPI.md#SyncSubscriptionPost) | **Post** /v1/internal/subscriptions/sync | Sync Subscription



## SyncSubscriptionPost

> interface{} SyncSubscriptionPost(ctx).AppApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest(appApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest).Authorization(authorization).Execute()

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
	appApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest := *openapiclient.NewAppApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest("Operation_example", "TenantId_example", "SubscriptionId_example", "CustomerId_example", *openapiclient.NewPlanRefPayload("Id_example", int32(123)), map[string]interface{}{"key": interface{}(123)}) // AppApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest | 
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InternalSubscriptionsAPI.SyncSubscriptionPost(context.Background()).AppApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest(appApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalSubscriptionsAPI.SyncSubscriptionPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncSubscriptionPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `InternalSubscriptionsAPI.SyncSubscriptionPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncSubscriptionPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest** | [**AppApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest**](AppApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest.md) |  | 
 **authorization** | **string** |  | 

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

