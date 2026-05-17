# \TenantConfigAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTenantConfig**](TenantConfigAPI.md#GetTenantConfig) | **Get** /v1/tenant/config | Get tenant regional config for SDK auto-discovery



## GetTenantConfig

> TenantConfigResponse GetTenantConfig(ctx).Execute()

Get tenant regional config for SDK auto-discovery



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
	resp, r, err := apiClient.TenantConfigAPI.GetTenantConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantConfigAPI.GetTenantConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantConfig`: TenantConfigResponse
	fmt.Fprintf(os.Stdout, "Response from `TenantConfigAPI.GetTenantConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantConfigRequest struct via the builder pattern


### Return type

[**TenantConfigResponse**](TenantConfigResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

