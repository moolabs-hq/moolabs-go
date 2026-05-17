# \AlertsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateThresholdEndpoint**](AlertsAPI.md#CreateThresholdEndpoint) | **Post** /v1/alerts/thresholds | Create Threshold Endpoint
[**DeleteThresholdEndpoint**](AlertsAPI.md#DeleteThresholdEndpoint) | **Delete** /v1/alerts/thresholds/{threshold_id} | Delete Threshold Endpoint
[**GetAlertsStateEndpoint**](AlertsAPI.md#GetAlertsStateEndpoint) | **Get** /v1/alerts/state | Get Alerts State Endpoint
[**ListThresholdsEndpoint**](AlertsAPI.md#ListThresholdsEndpoint) | **Get** /v1/alerts/thresholds | List Thresholds Endpoint
[**UpdateThresholdEndpoint**](AlertsAPI.md#UpdateThresholdEndpoint) | **Put** /v1/alerts/thresholds/{threshold_id} | Update Threshold Endpoint



## CreateThresholdEndpoint

> interface{} CreateThresholdEndpoint(ctx).WalletId(walletId).TenantId(tenantId).PoolId(poolId).CreateThresholdRequest(createThresholdRequest).Execute()

Create Threshold Endpoint



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Wallet identifier
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Pool identifier
	createThresholdRequest := *openapiclient.NewCreateThresholdRequest("ThresholdType_example", int32(123)) // CreateThresholdRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.CreateThresholdEndpoint(context.Background()).WalletId(walletId).TenantId(tenantId).PoolId(poolId).CreateThresholdRequest(createThresholdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.CreateThresholdEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateThresholdEndpoint`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.CreateThresholdEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateThresholdEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **walletId** | **string** | Wallet identifier | 
 **tenantId** | **string** | Tenant identifier | 
 **poolId** | **string** | Pool identifier | 
 **createThresholdRequest** | [**CreateThresholdRequest**](CreateThresholdRequest.md) |  | 

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


## DeleteThresholdEndpoint

> interface{} DeleteThresholdEndpoint(ctx, thresholdId).WalletId(walletId).TenantId(tenantId).PoolId(poolId).Execute()

Delete Threshold Endpoint



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
	thresholdId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Wallet identifier
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Pool identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.DeleteThresholdEndpoint(context.Background(), thresholdId).WalletId(walletId).TenantId(tenantId).PoolId(poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.DeleteThresholdEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteThresholdEndpoint`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.DeleteThresholdEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thresholdId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteThresholdEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **walletId** | **string** | Wallet identifier | 
 **tenantId** | **string** | Tenant identifier | 
 **poolId** | **string** | Pool identifier | 

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


## GetAlertsStateEndpoint

> interface{} GetAlertsStateEndpoint(ctx).WalletId(walletId).TenantId(tenantId).PoolId(poolId).Execute()

Get Alerts State Endpoint



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Wallet identifier
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Pool identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.GetAlertsStateEndpoint(context.Background()).WalletId(walletId).TenantId(tenantId).PoolId(poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.GetAlertsStateEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlertsStateEndpoint`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.GetAlertsStateEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertsStateEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **walletId** | **string** | Wallet identifier | 
 **tenantId** | **string** | Tenant identifier | 
 **poolId** | **string** | Pool identifier | 

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


## ListThresholdsEndpoint

> interface{} ListThresholdsEndpoint(ctx).WalletId(walletId).TenantId(tenantId).PoolId(poolId).Execute()

List Thresholds Endpoint



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Wallet identifier
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Pool identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.ListThresholdsEndpoint(context.Background()).WalletId(walletId).TenantId(tenantId).PoolId(poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.ListThresholdsEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListThresholdsEndpoint`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.ListThresholdsEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListThresholdsEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **walletId** | **string** | Wallet identifier | 
 **tenantId** | **string** | Tenant identifier | 
 **poolId** | **string** | Pool identifier | 

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


## UpdateThresholdEndpoint

> interface{} UpdateThresholdEndpoint(ctx, thresholdId).WalletId(walletId).TenantId(tenantId).PoolId(poolId).UpdateThresholdRequest(updateThresholdRequest).Execute()

Update Threshold Endpoint



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
	thresholdId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Wallet identifier
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Pool identifier
	updateThresholdRequest := *openapiclient.NewUpdateThresholdRequest() // UpdateThresholdRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.UpdateThresholdEndpoint(context.Background(), thresholdId).WalletId(walletId).TenantId(tenantId).PoolId(poolId).UpdateThresholdRequest(updateThresholdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.UpdateThresholdEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateThresholdEndpoint`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.UpdateThresholdEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thresholdId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateThresholdEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **walletId** | **string** | Wallet identifier | 
 **tenantId** | **string** | Tenant identifier | 
 **poolId** | **string** | Pool identifier | 
 **updateThresholdRequest** | [**UpdateThresholdRequest**](UpdateThresholdRequest.md) |  | 

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

