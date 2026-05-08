# \GrantsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGrant**](GrantsAPI.md#GetGrant) | **Get** /v1/grants/{grant_id} | Get Grant
[**ListGrants**](GrantsAPI.md#ListGrants) | **Get** /v1/grants | List Grants
[**MintGrant**](GrantsAPI.md#MintGrant) | **Post** /v1/grants/mint | Mint Grant
[**VoidGrant**](GrantsAPI.md#VoidGrant) | **Post** /v1/grants/{grant_id}/void | Void Grant



## GetGrant

> GrantResponse GetGrant(ctx, grantId).Execute()

Get Grant



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GrantsAPI.GetGrant(context.Background(), grantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GrantsAPI.GetGrant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGrant`: GrantResponse
	fmt.Fprintf(os.Stdout, "Response from `GrantsAPI.GetGrant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**grantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GrantResponse**](GrantResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGrants

> GrantsListResponse ListGrants(ctx).TenantId(tenantId).PoolId(poolId).WalletId(walletId).SubjectId(subjectId).SubscriptionId(subscriptionId).FeatureKey(featureKey).MeterSlug(meterSlug).AsOf(asOf).EffectiveAsOf(effectiveAsOf).RecordedAsOf(recordedAsOf).ConsistentView(consistentView).Limit(limit).Offset(offset).Execute()

List Grants



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier (optional if subject_id or subscription_id provided) (optional)
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Pool identifier (optional if subject_id or subscription_id provided) (optional)
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Wallet identifier to filter by (optional)
	subjectId := "subjectId_example" // string | Subject identifier (for finding wallet, auto-resolves tenant_id/pool_id) (optional)
	subscriptionId := "subscriptionId_example" // string | Subscription identifier (filters grants by subscription, auto-resolves tenant_id/pool_id) (optional)
	featureKey := "featureKey_example" // string | Feature key for scope filtering (optional)
	meterSlug := "meterSlug_example" // string | Meter slug for scope filtering (optional)
	asOf := time.Now() // time.Time | As-of timestamp for time travel (legacy, use effective_as_of instead) (optional)
	effectiveAsOf := time.Now() // time.Time | Effective as-of timestamp (business time) for time travel (optional)
	recordedAsOf := time.Now() // time.Time | Recorded as-of timestamp (system time) for time travel (optional)
	consistentView := true // bool | Use strong consistency for reads (optional)
	limit := int32(56) // int32 | Maximum number of grants to return (optional) (default to 100)
	offset := int32(56) // int32 | Offset for pagination (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GrantsAPI.ListGrants(context.Background()).TenantId(tenantId).PoolId(poolId).WalletId(walletId).SubjectId(subjectId).SubscriptionId(subscriptionId).FeatureKey(featureKey).MeterSlug(meterSlug).AsOf(asOf).EffectiveAsOf(effectiveAsOf).RecordedAsOf(recordedAsOf).ConsistentView(consistentView).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GrantsAPI.ListGrants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGrants`: GrantsListResponse
	fmt.Fprintf(os.Stdout, "Response from `GrantsAPI.ListGrants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGrantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Tenant identifier (optional if subject_id or subscription_id provided) | 
 **poolId** | **string** | Pool identifier (optional if subject_id or subscription_id provided) | 
 **walletId** | **string** | Wallet identifier to filter by | 
 **subjectId** | **string** | Subject identifier (for finding wallet, auto-resolves tenant_id/pool_id) | 
 **subscriptionId** | **string** | Subscription identifier (filters grants by subscription, auto-resolves tenant_id/pool_id) | 
 **featureKey** | **string** | Feature key for scope filtering | 
 **meterSlug** | **string** | Meter slug for scope filtering | 
 **asOf** | **time.Time** | As-of timestamp for time travel (legacy, use effective_as_of instead) | 
 **effectiveAsOf** | **time.Time** | Effective as-of timestamp (business time) for time travel | 
 **recordedAsOf** | **time.Time** | Recorded as-of timestamp (system time) for time travel | 
 **consistentView** | **bool** | Use strong consistency for reads | 
 **limit** | **int32** | Maximum number of grants to return | [default to 100]
 **offset** | **int32** | Offset for pagination | [default to 0]

### Return type

[**GrantsListResponse**](GrantsListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MintGrant

> GrantResponse MintGrant(ctx).CreateGrantRequest(createGrantRequest).Execute()

Mint Grant



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
	createGrantRequest := *openapiclient.NewCreateGrantRequest("TenantId_example", "PoolId_example", "WalletId_example", int32(123), time.Now(), openapiclient.GrantSourceType("SUBSCRIPTION_ALLOWANCE"), "SourceId_example") // CreateGrantRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GrantsAPI.MintGrant(context.Background()).CreateGrantRequest(createGrantRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GrantsAPI.MintGrant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MintGrant`: GrantResponse
	fmt.Fprintf(os.Stdout, "Response from `GrantsAPI.MintGrant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMintGrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createGrantRequest** | [**CreateGrantRequest**](CreateGrantRequest.md) |  | 

### Return type

[**GrantResponse**](GrantResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VoidGrant

> GrantResponse VoidGrant(ctx, grantId).VoidGrantRequest(voidGrantRequest).Execute()

Void Grant



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
	voidGrantRequest := *openapiclient.NewVoidGrantRequest() // VoidGrantRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GrantsAPI.VoidGrant(context.Background(), grantId).VoidGrantRequest(voidGrantRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GrantsAPI.VoidGrant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VoidGrant`: GrantResponse
	fmt.Fprintf(os.Stdout, "Response from `GrantsAPI.VoidGrant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**grantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVoidGrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **voidGrantRequest** | [**VoidGrantRequest**](VoidGrantRequest.md) |  | 

### Return type

[**GrantResponse**](GrantResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

