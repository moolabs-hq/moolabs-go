# \MappingsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMapping**](MappingsAPI.md#CreateMapping) | **Post** /v1/mappings | Create Mapping
[**ListMappings**](MappingsAPI.md#ListMappings) | **Get** /v1/mappings | List Mappings



## CreateMapping

> MappingResponse CreateMapping(ctx).TenantId(tenantId).CreateMappingRequest(createMappingRequest).Execute()

Create Mapping



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
	createMappingRequest := *openapiclient.NewCreateMappingRequest("MeterSlug_example", "FeatureKey_example") // CreateMappingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MappingsAPI.CreateMapping(context.Background()).TenantId(tenantId).CreateMappingRequest(createMappingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MappingsAPI.CreateMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMapping`: MappingResponse
	fmt.Fprintf(os.Stdout, "Response from `MappingsAPI.CreateMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Tenant identifier | 
 **createMappingRequest** | [**CreateMappingRequest**](CreateMappingRequest.md) |  | 

### Return type

[**MappingResponse**](MappingResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMappings

> []MappingResponse ListMappings(ctx).TenantId(tenantId).MeterSlug(meterSlug).FeatureKey(featureKey).Execute()

List Mappings



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
	meterSlug := "meterSlug_example" // string | Filter by meter slug (optional)
	featureKey := "featureKey_example" // string | Filter by feature key (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MappingsAPI.ListMappings(context.Background()).TenantId(tenantId).MeterSlug(meterSlug).FeatureKey(featureKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MappingsAPI.ListMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMappings`: []MappingResponse
	fmt.Fprintf(os.Stdout, "Response from `MappingsAPI.ListMappings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Tenant identifier | 
 **meterSlug** | **string** | Filter by meter slug | 
 **featureKey** | **string** | Filter by feature key | 

### Return type

[**[]MappingResponse**](MappingResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

