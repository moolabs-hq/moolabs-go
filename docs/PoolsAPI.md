# \PoolsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPool**](PoolsAPI.md#GetPool) | **Get** /v1/pools/{pool_id} | Get Pool
[**GetPoolImpactSummaryV1**](PoolsAPI.md#GetPoolImpactSummaryV1) | **Get** /v1/pools/{pool_id}/impact-summary | Get Pool Impact Summary
[**GetPoolNotificationConfigV1**](PoolsAPI.md#GetPoolNotificationConfigV1) | **Get** /v1/pools/{pool_id}/notification-config | Get Pool Notification Config
[**GetPoolTopupDefaultsV1**](PoolsAPI.md#GetPoolTopupDefaultsV1) | **Get** /v1/pools/{pool_id}/topup-defaults | Get Pool Topup Defaults
[**ListPools**](PoolsAPI.md#ListPools) | **Get** /v1/pools | List Pools
[**UpdatePool**](PoolsAPI.md#UpdatePool) | **Put** /v1/pools/{pool_id} | Update Pool
[**UpdatePoolNotificationConfigV1**](PoolsAPI.md#UpdatePoolNotificationConfigV1) | **Put** /v1/pools/{pool_id}/notification-config | Update Pool Notification Config
[**UpdatePoolTopupDefaultsV1**](PoolsAPI.md#UpdatePoolTopupDefaultsV1) | **Put** /v1/pools/{pool_id}/topup-defaults | Update Pool Topup Defaults



## GetPool

> interface{} GetPool(ctx, poolId).TenantId(tenantId).Execute()

Get Pool



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
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.GetPool(context.Background(), poolId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.GetPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPool`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.GetPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** | Tenant identifier | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPoolImpactSummaryV1

> interface{} GetPoolImpactSummaryV1(ctx, poolId).TenantId(tenantId).Execute()

Get Pool Impact Summary



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
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.GetPoolImpactSummaryV1(context.Background(), poolId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.GetPoolImpactSummaryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPoolImpactSummaryV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.GetPoolImpactSummaryV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPoolImpactSummaryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** | Tenant identifier | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPoolNotificationConfigV1

> interface{} GetPoolNotificationConfigV1(ctx, poolId).TenantId(tenantId).Execute()

Get Pool Notification Config



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
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.GetPoolNotificationConfigV1(context.Background(), poolId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.GetPoolNotificationConfigV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPoolNotificationConfigV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.GetPoolNotificationConfigV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPoolNotificationConfigV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** | Tenant identifier | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPoolTopupDefaultsV1

> interface{} GetPoolTopupDefaultsV1(ctx, poolId).TenantId(tenantId).Execute()

Get Pool Topup Defaults



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
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.GetPoolTopupDefaultsV1(context.Background(), poolId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.GetPoolTopupDefaultsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPoolTopupDefaultsV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.GetPoolTopupDefaultsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPoolTopupDefaultsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** | Tenant identifier | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPools

> interface{} ListPools(ctx).TenantId(tenantId).Execute()

List Pools



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.ListPools(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.ListPools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPools`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.ListPools`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Tenant identifier | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePool

> interface{} UpdatePool(ctx, poolId).TenantId(tenantId).UpdateCreditPoolRequest(updateCreditPoolRequest).Execute()

Update Pool



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
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier
	updateCreditPoolRequest := *openapiclient.NewUpdateCreditPoolRequest() // UpdateCreditPoolRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.UpdatePool(context.Background(), poolId).TenantId(tenantId).UpdateCreditPoolRequest(updateCreditPoolRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.UpdatePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePool`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.UpdatePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** | Tenant identifier | 
 **updateCreditPoolRequest** | [**UpdateCreditPoolRequest**](UpdateCreditPoolRequest.md) |  | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePoolNotificationConfigV1

> interface{} UpdatePoolNotificationConfigV1(ctx, poolId).TenantId(tenantId).PoolNotificationConfigRequest(poolNotificationConfigRequest).Execute()

Update Pool Notification Config



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
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier
	poolNotificationConfigRequest := *openapiclient.NewPoolNotificationConfigRequest([]string{"NotificationEmails_example"}) // PoolNotificationConfigRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.UpdatePoolNotificationConfigV1(context.Background(), poolId).TenantId(tenantId).PoolNotificationConfigRequest(poolNotificationConfigRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.UpdatePoolNotificationConfigV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePoolNotificationConfigV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.UpdatePoolNotificationConfigV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePoolNotificationConfigV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** | Tenant identifier | 
 **poolNotificationConfigRequest** | [**PoolNotificationConfigRequest**](PoolNotificationConfigRequest.md) |  | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePoolTopupDefaultsV1

> interface{} UpdatePoolTopupDefaultsV1(ctx, poolId).TenantId(tenantId).TopupDefaultsRequest(topupDefaultsRequest).Execute()

Update Pool Topup Defaults



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
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier
	topupDefaultsRequest := *openapiclient.NewTopupDefaultsRequest() // TopupDefaultsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.UpdatePoolTopupDefaultsV1(context.Background(), poolId).TenantId(tenantId).TopupDefaultsRequest(topupDefaultsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.UpdatePoolTopupDefaultsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePoolTopupDefaultsV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.UpdatePoolTopupDefaultsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePoolTopupDefaultsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** | Tenant identifier | 
 **topupDefaultsRequest** | [**TopupDefaultsRequest**](TopupDefaultsRequest.md) |  | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

