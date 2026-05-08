# \ReconstructionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReconstructionRun**](ReconstructionAPI.md#CreateReconstructionRun) | **Post** /v1/reconstruction/runs | Create Reconstruction Run
[**GetAffectedEventsEndpointV1**](ReconstructionAPI.md#GetAffectedEventsEndpointV1) | **Get** /v1/reconstruction/affected-events | Get Affected Events Endpoint
[**GetLateEventsV1**](ReconstructionAPI.md#GetLateEventsV1) | **Get** /v1/reconstruction/late-events | Get Late Events
[**GetWalletSubtree**](ReconstructionAPI.md#GetWalletSubtree) | **Get** /v1/reconstruction/wallets/{wallet_id}/subtree | Get Wallet Subtree



## CreateReconstructionRun

> ReconstructionRunResponse CreateReconstructionRun(ctx).CreateReconstructionRunRequest(createReconstructionRunRequest).Execute()

Create Reconstruction Run



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
	createReconstructionRunRequest := *openapiclient.NewCreateReconstructionRunRequest("TenantId_example", "PoolId_example", "RootWalletId_example", time.Now(), time.Now(), "TriggerEventId_example") // CreateReconstructionRunRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReconstructionAPI.CreateReconstructionRun(context.Background()).CreateReconstructionRunRequest(createReconstructionRunRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReconstructionAPI.CreateReconstructionRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateReconstructionRun`: ReconstructionRunResponse
	fmt.Fprintf(os.Stdout, "Response from `ReconstructionAPI.CreateReconstructionRun`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReconstructionRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createReconstructionRunRequest** | [**CreateReconstructionRunRequest**](CreateReconstructionRunRequest.md) |  | 

### Return type

[**ReconstructionRunResponse**](ReconstructionRunResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAffectedEventsEndpointV1

> interface{} GetAffectedEventsEndpointV1(ctx).TenantId(tenantId).PoolId(poolId).WalletIds(walletIds).FromEffectiveAt(fromEffectiveAt).ToEffectiveAt(toEffectiveAt).Execute()

Get Affected Events Endpoint



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Pool identifier
	walletIds := "walletIds_example" // string | Comma-separated wallet IDs
	fromEffectiveAt := time.Now() // time.Time | Start of effective time window
	toEffectiveAt := time.Now() // time.Time | End of effective time window

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReconstructionAPI.GetAffectedEventsEndpointV1(context.Background()).TenantId(tenantId).PoolId(poolId).WalletIds(walletIds).FromEffectiveAt(fromEffectiveAt).ToEffectiveAt(toEffectiveAt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReconstructionAPI.GetAffectedEventsEndpointV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAffectedEventsEndpointV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ReconstructionAPI.GetAffectedEventsEndpointV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAffectedEventsEndpointV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Tenant identifier | 
 **poolId** | **string** | Pool identifier | 
 **walletIds** | **string** | Comma-separated wallet IDs | 
 **fromEffectiveAt** | **time.Time** | Start of effective time window | 
 **toEffectiveAt** | **time.Time** | End of effective time window | 

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


## GetLateEventsV1

> interface{} GetLateEventsV1(ctx).TenantId(tenantId).PoolId(poolId).LateThresholdSeconds(lateThresholdSeconds).Execute()

Get Late Events



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
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Pool identifier
	lateThresholdSeconds := int32(56) // int32 | LATE event threshold (seconds) (optional) (default to 3600)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReconstructionAPI.GetLateEventsV1(context.Background()).TenantId(tenantId).PoolId(poolId).LateThresholdSeconds(lateThresholdSeconds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReconstructionAPI.GetLateEventsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLateEventsV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ReconstructionAPI.GetLateEventsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLateEventsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Tenant identifier | 
 **poolId** | **string** | Pool identifier | 
 **lateThresholdSeconds** | **int32** | LATE event threshold (seconds) | [default to 3600]

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


## GetWalletSubtree

> interface{} GetWalletSubtree(ctx, walletId).TenantId(tenantId).PoolId(poolId).MaxDepth(maxDepth).Execute()

Get Wallet Subtree



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
	maxDepth := int32(56) // int32 | Maximum depth to traverse (optional) (default to 3)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReconstructionAPI.GetWalletSubtree(context.Background(), walletId).TenantId(tenantId).PoolId(poolId).MaxDepth(maxDepth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReconstructionAPI.GetWalletSubtree``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletSubtree`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ReconstructionAPI.GetWalletSubtree`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** | Wallet identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletSubtreeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** | Tenant identifier | 
 **poolId** | **string** | Pool identifier | 
 **maxDepth** | **int32** | Maximum depth to traverse | [default to 3]

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

