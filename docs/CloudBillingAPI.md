# \CloudBillingAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfigureSyncApi**](CloudBillingAPI.md#ConfigureSyncApi) | **Post** /api/v1/cloud-billing/sync | Configure scheduled sync (stores config, returns 202)
[**GetImportBatchApi**](CloudBillingAPI.md#GetImportBatchApi) | **Get** /api/v1/cloud-billing/imports/{batch_id} | Get batch details with row-level data
[**ImportBillingBatchApi**](CloudBillingAPI.md#ImportBillingBatchApi) | **Post** /api/v1/cloud-billing/import | Import a billing batch
[**ListImportBatchesApi**](CloudBillingAPI.md#ListImportBatchesApi) | **Get** /api/v1/cloud-billing/imports | List recent import batches
[**ListResourceMapApiV1**](CloudBillingAPI.md#ListResourceMapApiV1) | **Get** /api/v1/cloud-billing/resource-map | List resource_service_map entries
[**UpsertResourceMapApiV1**](CloudBillingAPI.md#UpsertResourceMapApiV1) | **Post** /api/v1/cloud-billing/resource-map | Create or update a resource mapping



## ConfigureSyncApi

> SyncConfigResponse ConfigureSyncApi(ctx).SyncConfigRequest(syncConfigRequest).Execute()

Configure scheduled sync (stores config, returns 202)



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
	syncConfigRequest := *openapiclient.NewSyncConfigRequest("CloudProvider_example") // SyncConfigRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudBillingAPI.ConfigureSyncApi(context.Background()).SyncConfigRequest(syncConfigRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudBillingAPI.ConfigureSyncApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfigureSyncApi`: SyncConfigResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudBillingAPI.ConfigureSyncApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConfigureSyncApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **syncConfigRequest** | [**SyncConfigRequest**](SyncConfigRequest.md) |  | 

### Return type

[**SyncConfigResponse**](SyncConfigResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImportBatchApi

> ImportBatchDetailResponse GetImportBatchApi(ctx, batchId).Execute()

Get batch details with row-level data



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
	batchId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudBillingAPI.GetImportBatchApi(context.Background(), batchId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudBillingAPI.GetImportBatchApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImportBatchApi`: ImportBatchDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudBillingAPI.GetImportBatchApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImportBatchApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ImportBatchDetailResponse**](ImportBatchDetailResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportBillingBatchApi

> ImportBatchResponse ImportBillingBatchApi(ctx).ImportBatchRequest(importBatchRequest).Execute()

Import a billing batch



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
	importBatchRequest := *openapiclient.NewImportBatchRequest("CloudProvider_example", time.Now(), time.Now(), []openapiclient.CloudCostRowInput{*openapiclient.NewCloudCostRowInput("ServiceName_example", *openapiclient.NewCost())}) // ImportBatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudBillingAPI.ImportBillingBatchApi(context.Background()).ImportBatchRequest(importBatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudBillingAPI.ImportBillingBatchApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportBillingBatchApi`: ImportBatchResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudBillingAPI.ImportBillingBatchApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportBillingBatchApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importBatchRequest** | [**ImportBatchRequest**](ImportBatchRequest.md) |  | 

### Return type

[**ImportBatchResponse**](ImportBatchResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListImportBatchesApi

> []ImportBatchSummary ListImportBatchesApi(ctx).CloudProvider(cloudProvider).Limit(limit).Execute()

List recent import batches



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
	cloudProvider := "cloudProvider_example" // string | 'aws', 'gcp', or 'azure'
	limit := int32(56) // int32 |  (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudBillingAPI.ListImportBatchesApi(context.Background()).CloudProvider(cloudProvider).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudBillingAPI.ListImportBatchesApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListImportBatchesApi`: []ImportBatchSummary
	fmt.Fprintf(os.Stdout, "Response from `CloudBillingAPI.ListImportBatchesApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListImportBatchesApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudProvider** | **string** | &#39;aws&#39;, &#39;gcp&#39;, or &#39;azure&#39; | 
 **limit** | **int32** |  | [default to 20]

### Return type

[**[]ImportBatchSummary**](ImportBatchSummary.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListResourceMapApiV1

> []ResourceMapResponse ListResourceMapApiV1(ctx).CloudProvider(cloudProvider).IsActive(isActive).Limit(limit).Offset(offset).Execute()

List resource_service_map entries



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
	cloudProvider := "cloudProvider_example" // string |  (optional)
	isActive := true // bool |  (optional) (default to true)
	limit := int32(56) // int32 |  (optional) (default to 100)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudBillingAPI.ListResourceMapApiV1(context.Background()).CloudProvider(cloudProvider).IsActive(isActive).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudBillingAPI.ListResourceMapApiV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListResourceMapApiV1`: []ResourceMapResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudBillingAPI.ListResourceMapApiV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListResourceMapApiV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudProvider** | **string** |  | 
 **isActive** | **bool** |  | [default to true]
 **limit** | **int32** |  | [default to 100]
 **offset** | **int32** |  | [default to 0]

### Return type

[**[]ResourceMapResponse**](ResourceMapResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertResourceMapApiV1

> ResourceMapResponse UpsertResourceMapApiV1(ctx).ResourceMapCreateRequest(resourceMapCreateRequest).Execute()

Create or update a resource mapping



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
	resourceMapCreateRequest := *openapiclient.NewResourceMapCreateRequest("CloudProvider_example", "ResourceId_example", "ServiceName_example") // ResourceMapCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudBillingAPI.UpsertResourceMapApiV1(context.Background()).ResourceMapCreateRequest(resourceMapCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudBillingAPI.UpsertResourceMapApiV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpsertResourceMapApiV1`: ResourceMapResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudBillingAPI.UpsertResourceMapApiV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpsertResourceMapApiV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceMapCreateRequest** | [**ResourceMapCreateRequest**](ResourceMapCreateRequest.md) |  | 

### Return type

[**ResourceMapResponse**](ResourceMapResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

