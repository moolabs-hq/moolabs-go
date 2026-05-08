# \IntegrationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OpenmeterWebhook**](IntegrationsAPI.md#OpenmeterWebhook) | **Post** /v1/integrations/openmeter/webhooks/openmeter | Openmeter Webhook
[**OpenmeterWebhookBatch**](IntegrationsAPI.md#OpenmeterWebhookBatch) | **Post** /v1/integrations/openmeter/webhooks/openmeter/batch | Openmeter Webhook Batch



## OpenmeterWebhook

> interface{} OpenmeterWebhook(ctx).TenantId(tenantId).PoolId(poolId).Execute()

Openmeter Webhook



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.OpenmeterWebhook(context.Background()).TenantId(tenantId).PoolId(poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.OpenmeterWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpenmeterWebhook`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.OpenmeterWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpenmeterWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **poolId** | **string** |  | 

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


## OpenmeterWebhookBatch

> interface{} OpenmeterWebhookBatch(ctx).TenantId(tenantId).PoolId(poolId).Execute()

Openmeter Webhook Batch



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.OpenmeterWebhookBatch(context.Background()).TenantId(tenantId).PoolId(poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.OpenmeterWebhookBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpenmeterWebhookBatch`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.OpenmeterWebhookBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpenmeterWebhookBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **poolId** | **string** |  | 

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

