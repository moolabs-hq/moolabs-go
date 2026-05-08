# \StateProjectionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWalletStateProjectionEndpointV1**](StateProjectionAPI.md#GetWalletStateProjectionEndpointV1) | **Get** /v1/state-projection/wallet/{wallet_id} | Get Wallet State Projection Endpoint
[**ProcessPendingProjectionsEndpointV1**](StateProjectionAPI.md#ProcessPendingProjectionsEndpointV1) | **Post** /v1/state-projection/process | Process Pending Projections Endpoint
[**ProjectWalletStateEndpointV1**](StateProjectionAPI.md#ProjectWalletStateEndpointV1) | **Post** /v1/state-projection/project/{wallet_id} | Project Wallet State Endpoint



## GetWalletStateProjectionEndpointV1

> interface{} GetWalletStateProjectionEndpointV1(ctx, walletId).TenantId(tenantId).PoolId(poolId).AsOf(asOf).Execute()

Get Wallet State Projection Endpoint



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/moolabs/moolabs-go"
)

func main() {
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Pool identifier
	asOf := time.Now() // time.Time | As-of timestamp (for time travel) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StateProjectionAPI.GetWalletStateProjectionEndpointV1(context.Background(), walletId).TenantId(tenantId).PoolId(poolId).AsOf(asOf).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StateProjectionAPI.GetWalletStateProjectionEndpointV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletStateProjectionEndpointV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `StateProjectionAPI.GetWalletStateProjectionEndpointV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletStateProjectionEndpointV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** | Tenant identifier | 
 **poolId** | **string** | Pool identifier | 
 **asOf** | **time.Time** | As-of timestamp (for time travel) | 

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


## ProcessPendingProjectionsEndpointV1

> interface{} ProcessPendingProjectionsEndpointV1(ctx).TenantId(tenantId).PoolId(poolId).BatchSize(batchSize).Execute()

Process Pending Projections Endpoint



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional tenant filter (optional)
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional pool filter (optional)
	batchSize := int32(56) // int32 | Maximum number of wallets to process (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StateProjectionAPI.ProcessPendingProjectionsEndpointV1(context.Background()).TenantId(tenantId).PoolId(poolId).BatchSize(batchSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StateProjectionAPI.ProcessPendingProjectionsEndpointV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProcessPendingProjectionsEndpointV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `StateProjectionAPI.ProcessPendingProjectionsEndpointV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProcessPendingProjectionsEndpointV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Optional tenant filter | 
 **poolId** | **string** | Optional pool filter | 
 **batchSize** | **int32** | Maximum number of wallets to process | [default to 100]

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


## ProjectWalletStateEndpointV1

> interface{} ProjectWalletStateEndpointV1(ctx, walletId).TenantId(tenantId).PoolId(poolId).AsOf(asOf).Execute()

Project Wallet State Endpoint



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/moolabs/moolabs-go"
)

func main() {
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Pool identifier
	asOf := time.Now() // time.Time | As-of timestamp (for time travel) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StateProjectionAPI.ProjectWalletStateEndpointV1(context.Background(), walletId).TenantId(tenantId).PoolId(poolId).AsOf(asOf).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StateProjectionAPI.ProjectWalletStateEndpointV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectWalletStateEndpointV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `StateProjectionAPI.ProjectWalletStateEndpointV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectWalletStateEndpointV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** | Tenant identifier | 
 **poolId** | **string** | Pool identifier | 
 **asOf** | **time.Time** | As-of timestamp (for time travel) | 

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

