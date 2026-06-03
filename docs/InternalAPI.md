# \InternalAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GrantArcDunningTemplatePermissionV1InternalArc**](InternalAPI.md#GrantArcDunningTemplatePermissionV1InternalArc) | **Post** /v1/internal/arc-dunning-template-permissions/grants | Grant Arc Dunning Template Permission
[**ReplayTenantProvisioningV1**](InternalAPI.md#ReplayTenantProvisioningV1) | **Post** /v1/internal/tenant-provision/replay | Replay Tenant Provisioning
[**RevokeArcDunningTemplatePermissionV1InternalArc**](InternalAPI.md#RevokeArcDunningTemplatePermissionV1InternalArc) | **Post** /v1/internal/arc-dunning-template-permissions/grants/{grant_id}/revoke | Revoke Arc Dunning Template Permission
[**TenantProvisionV1**](InternalAPI.md#TenantProvisionV1) | **Post** /v1/internal/tenant-provision | Tenant Provision



## GrantArcDunningTemplatePermissionV1InternalArc

> interface{} GrantArcDunningTemplatePermissionV1InternalArc(ctx).ArcDunningPermissionGrantRequest(arcDunningPermissionGrantRequest).Execute()

Grant Arc Dunning Template Permission

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
	arcDunningPermissionGrantRequest := *openapiclient.NewArcDunningPermissionGrantRequest("TenantId_example", "OrgId_example", "UserSubject_example", "GrantReason_example") // ArcDunningPermissionGrantRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InternalAPI.GrantArcDunningTemplatePermissionV1InternalArc(context.Background()).ArcDunningPermissionGrantRequest(arcDunningPermissionGrantRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalAPI.GrantArcDunningTemplatePermissionV1InternalArc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GrantArcDunningTemplatePermissionV1InternalArc`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `InternalAPI.GrantArcDunningTemplatePermissionV1InternalArc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGrantArcDunningTemplatePermissionV1InternalArcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arcDunningPermissionGrantRequest** | [**ArcDunningPermissionGrantRequest**](ArcDunningPermissionGrantRequest.md) |  | 

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


## RevokeArcDunningTemplatePermissionV1InternalArc

> interface{} RevokeArcDunningTemplatePermissionV1InternalArc(ctx, grantId).ArcDunningPermissionRevokeRequest(arcDunningPermissionRevokeRequest).Execute()

Revoke Arc Dunning Template Permission

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
	grantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	arcDunningPermissionRevokeRequest := *openapiclient.NewArcDunningPermissionRevokeRequest("RevokeReason_example") // ArcDunningPermissionRevokeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InternalAPI.RevokeArcDunningTemplatePermissionV1InternalArc(context.Background(), grantId).ArcDunningPermissionRevokeRequest(arcDunningPermissionRevokeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalAPI.RevokeArcDunningTemplatePermissionV1InternalArc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RevokeArcDunningTemplatePermissionV1InternalArc`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `InternalAPI.RevokeArcDunningTemplatePermissionV1InternalArc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**grantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeArcDunningTemplatePermissionV1InternalArcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **arcDunningPermissionRevokeRequest** | [**ArcDunningPermissionRevokeRequest**](ArcDunningPermissionRevokeRequest.md) |  | 

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

