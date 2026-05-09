# \LookupInformationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProgress**](LookupInformationAPI.md#GetProgress) | **Get** /api/v1/info/progress/{id} | Get progress
[**ListCurrencies**](LookupInformationAPI.md#ListCurrencies) | **Get** /api/v1/info/currencies | List supported currencies



## GetProgress

> Progress GetProgress(ctx, id).Execute()

Get progress



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LookupInformationAPI.GetProgress(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LookupInformationAPI.GetProgress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProgress`: Progress
	fmt.Fprintf(os.Stdout, "Response from `LookupInformationAPI.GetProgress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProgressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Progress**](Progress.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCurrencies

> []Currency ListCurrencies(ctx).Execute()

List supported currencies



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LookupInformationAPI.ListCurrencies(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LookupInformationAPI.ListCurrencies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCurrencies`: []Currency
	fmt.Fprintf(os.Stdout, "Response from `LookupInformationAPI.ListCurrencies`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCurrenciesRequest struct via the builder pattern


### Return type

[**[]Currency**](Currency.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

