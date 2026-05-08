# \RolloverAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProcessRollover**](RolloverAPI.md#ProcessRollover) | **Post** /v1/rollover/process | Process Rollover



## ProcessRollover

> RolloverResponse ProcessRollover(ctx).RolloverRequest(rolloverRequest).Execute()

Process Rollover



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
	rolloverRequest := *openapiclient.NewRolloverRequest("TenantId_example", "SubscriptionId_example") // RolloverRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolloverAPI.ProcessRollover(context.Background()).RolloverRequest(rolloverRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolloverAPI.ProcessRollover``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProcessRollover`: RolloverResponse
	fmt.Fprintf(os.Stdout, "Response from `RolloverAPI.ProcessRollover`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProcessRolloverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rolloverRequest** | [**RolloverRequest**](RolloverRequest.md) |  | 

### Return type

[**RolloverResponse**](RolloverResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

