# \OpsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeepHealthCheck**](OpsAPI.md#DeepHealthCheck) | **Get** /health/deep | Deep Health Check



## DeepHealthCheck

> map[string]interface{} DeepHealthCheck(ctx).XMetricsToken(xMetricsToken).Execute()

Deep Health Check



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
	xMetricsToken := "xMetricsToken_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpsAPI.DeepHealthCheck(context.Background()).XMetricsToken(xMetricsToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsAPI.DeepHealthCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeepHealthCheck`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OpsAPI.DeepHealthCheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeepHealthCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xMetricsToken** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

