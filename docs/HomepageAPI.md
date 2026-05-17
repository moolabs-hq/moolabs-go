# \HomepageAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHomepageInsights**](HomepageAPI.md#GetHomepageInsights) | **Get** /v1/homepage/insights | Get Homepage Insights



## GetHomepageInsights

> HomepageResponse GetHomepageInsights(ctx).Persona(persona).Execute()

Get Homepage Insights



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
	persona := openapiclient.Persona("finance") // Persona | User persona/role (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HomepageAPI.GetHomepageInsights(context.Background()).Persona(persona).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HomepageAPI.GetHomepageInsights``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHomepageInsights`: HomepageResponse
	fmt.Fprintf(os.Stdout, "Response from `HomepageAPI.GetHomepageInsights`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHomepageInsightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **persona** | [**Persona**](Persona.md) | User persona/role | 

### Return type

[**HomepageResponse**](HomepageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

