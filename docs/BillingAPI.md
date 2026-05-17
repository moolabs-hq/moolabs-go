# \BillingAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetryRemediationEvent**](BillingAPI.md#RetryRemediationEvent) | **Post** /v1/billing/remediation/events/{usage_event_id}/retry | Retry Remediation Event



## RetryRemediationEvent

> interface{} RetryRemediationEvent(ctx, usageEventId).Execute()

Retry Remediation Event



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
	resp, r, err := apiClient.BillingAPI.RetryRemediationEvent(context.Background(), usageEventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.RetryRemediationEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetryRemediationEvent`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.RetryRemediationEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**usageEventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetryRemediationEventRequest struct via the builder pattern


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

