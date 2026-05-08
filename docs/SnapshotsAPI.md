# \SnapshotsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSnapshot**](SnapshotsAPI.md#CreateSnapshot) | **Post** /v1/snapshots | Create Snapshot
[**GetSnapshotAtEndpoint**](SnapshotsAPI.md#GetSnapshotAtEndpoint) | **Get** /v1/snapshots/wallet/{wallet_id}/at | Get Snapshot At Endpoint
[**ListSnapshotsEndpoint**](SnapshotsAPI.md#ListSnapshotsEndpoint) | **Get** /v1/snapshots/wallet/{wallet_id} | List Snapshots Endpoint
[**ValidateSnapshotEndpoint**](SnapshotsAPI.md#ValidateSnapshotEndpoint) | **Post** /v1/snapshots/{snapshot_id}/validate | Validate Snapshot Endpoint



## CreateSnapshot

> SnapshotResponse CreateSnapshot(ctx).CreateSnapshotRequest(createSnapshotRequest).Execute()

Create Snapshot



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
	createSnapshotRequest := *openapiclient.NewCreateSnapshotRequest("TenantId_example", "PoolId_example", "WalletId_example", time.Now()) // CreateSnapshotRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotsAPI.CreateSnapshot(context.Background()).CreateSnapshotRequest(createSnapshotRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.CreateSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSnapshot`: SnapshotResponse
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.CreateSnapshot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSnapshotRequest** | [**CreateSnapshotRequest**](CreateSnapshotRequest.md) |  | 

### Return type

[**SnapshotResponse**](SnapshotResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSnapshotAtEndpoint

> interface{} GetSnapshotAtEndpoint(ctx, walletId).TenantId(tenantId).PoolId(poolId).AsOfRecordedAt(asOfRecordedAt).Execute()

Get Snapshot At Endpoint



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Wallet identifier
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Pool identifier
	asOfRecordedAt := time.Now() // time.Time | Recorded timestamp

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotsAPI.GetSnapshotAtEndpoint(context.Background(), walletId).TenantId(tenantId).PoolId(poolId).AsOfRecordedAt(asOfRecordedAt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.GetSnapshotAtEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSnapshotAtEndpoint`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.GetSnapshotAtEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** | Wallet identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSnapshotAtEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** | Tenant identifier | 
 **poolId** | **string** | Pool identifier | 
 **asOfRecordedAt** | **time.Time** | Recorded timestamp | 

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


## ListSnapshotsEndpoint

> interface{} ListSnapshotsEndpoint(ctx, walletId).TenantId(tenantId).PoolId(poolId).Limit(limit).Execute()

List Snapshots Endpoint



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
	limit := int32(56) // int32 | Maximum number of snapshots (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotsAPI.ListSnapshotsEndpoint(context.Background(), walletId).TenantId(tenantId).PoolId(poolId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.ListSnapshotsEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSnapshotsEndpoint`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.ListSnapshotsEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** | Wallet identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSnapshotsEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** | Tenant identifier | 
 **poolId** | **string** | Pool identifier | 
 **limit** | **int32** | Maximum number of snapshots | [default to 100]

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


## ValidateSnapshotEndpoint

> ValidationResponse ValidateSnapshotEndpoint(ctx, snapshotId).Execute()

Validate Snapshot Endpoint



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
	snapshotId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Snapshot identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotsAPI.ValidateSnapshotEndpoint(context.Background(), snapshotId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.ValidateSnapshotEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateSnapshotEndpoint`: ValidationResponse
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.ValidateSnapshotEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string** | Snapshot identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateSnapshotEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ValidationResponse**](ValidationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

