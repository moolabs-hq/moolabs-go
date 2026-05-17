# \InternalAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReplayTenantProvisioningV1**](InternalAPI.md#ReplayTenantProvisioningV1) | **Post** /v1/internal/tenant-provision/replay | Replay Tenant Provisioning
[**TenantProvisionV1**](InternalAPI.md#TenantProvisionV1) | **Post** /v1/internal/tenant-provision | Tenant Provision



## ReplayTenantProvisioningV1

> interface{} ReplayTenantProvisioningV1(ctx).Limit(limit).Execute()

Replay Tenant Provisioning

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
	limit := int32(56) // int32 |  (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InternalAPI.ReplayTenantProvisioningV1(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalAPI.ReplayTenantProvisioningV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplayTenantProvisioningV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `InternalAPI.ReplayTenantProvisioningV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReplayTenantProvisioningV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 20]

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


## TenantProvisionV1

> interface{} TenantProvisionV1(ctx).TenantProvisionRequest(tenantProvisionRequest).Execute()

Tenant Provision

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
	tenantProvisionRequest := *openapiclient.NewTenantProvisionRequest("OrgId_example") // TenantProvisionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InternalAPI.TenantProvisionV1(context.Background()).TenantProvisionRequest(tenantProvisionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalAPI.TenantProvisionV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenantProvisionV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `InternalAPI.TenantProvisionV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantProvisionV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantProvisionRequest** | [**TenantProvisionRequest**](TenantProvisionRequest.md) |  | 

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

