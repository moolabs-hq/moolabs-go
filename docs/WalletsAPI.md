# \WalletsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllocateCredits**](WalletsAPI.md#AllocateCredits) | **Post** /v1/wallets/allocate | Allocate Credits
[**BindSubjectToWallet**](WalletsAPI.md#BindSubjectToWallet) | **Post** /v1/wallets/members | Bind Subject To Wallet
[**CreateWallet**](WalletsAPI.md#CreateWallet) | **Post** /v1/wallets | Create Wallet
[**GetWallet**](WalletsAPI.md#GetWallet) | **Get** /v1/wallets/{wallet_id} | Get Wallet
[**GetWalletActivity**](WalletsAPI.md#GetWalletActivity) | **Get** /v1/wallets/{wallet_id}/activity | Get Wallet Activity
[**GetWalletMembers**](WalletsAPI.md#GetWalletMembers) | **Get** /v1/wallets/{wallet_id}/members | Get Wallet Members
[**ListWalletMembers**](WalletsAPI.md#ListWalletMembers) | **Get** /v1/wallet_members | List Wallet Members
[**ListWallets**](WalletsAPI.md#ListWallets) | **Get** /v1/wallets | List Wallets
[**ListWalletsBySubjectsV1**](WalletsAPI.md#ListWalletsBySubjectsV1) | **Post** /v1/wallets/by-subjects | List Wallets By Subjects
[**ResolveTenantAndPool**](WalletsAPI.md#ResolveTenantAndPool) | **Get** /v1/wallets/resolve | Resolve Tenant And Pool
[**UpdateWalletSettings**](WalletsAPI.md#UpdateWalletSettings) | **Patch** /v1/wallets/{wallet_id}/settings | Update Wallet Settings
[**UpdateWalletThresholds**](WalletsAPI.md#UpdateWalletThresholds) | **Patch** /v1/wallets/{wallet_id}/thresholds | Update Wallet Thresholds



## AllocateCredits

> interface{} AllocateCredits(ctx).AllocateCreditsRequest(allocateCreditsRequest).Execute()

Allocate Credits



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
	allocateCreditsRequest := *openapiclient.NewAllocateCreditsRequest("SourceWalletId_example", "TargetWalletId_example", int32(123), "SubjectId_example") // AllocateCreditsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.AllocateCredits(context.Background()).AllocateCreditsRequest(allocateCreditsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.AllocateCredits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AllocateCredits`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.AllocateCredits`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAllocateCreditsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **allocateCreditsRequest** | [**AllocateCreditsRequest**](AllocateCreditsRequest.md) |  | 

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


## BindSubjectToWallet

> interface{} BindSubjectToWallet(ctx).TenantId(tenantId).PoolId(poolId).SubjectId(subjectId).WalletId(walletId).Execute()

Bind Subject To Wallet



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	subjectId := "subjectId_example" // string | 
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.BindSubjectToWallet(context.Background()).TenantId(tenantId).PoolId(poolId).SubjectId(subjectId).WalletId(walletId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.BindSubjectToWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BindSubjectToWallet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.BindSubjectToWallet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBindSubjectToWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **poolId** | **string** |  | 
 **subjectId** | **string** |  | 
 **walletId** | **string** |  | 

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


## CreateWallet

> WalletResponse CreateWallet(ctx).CreateWalletRequest(createWalletRequest).Execute()

Create Wallet



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
	createWalletRequest := *openapiclient.NewCreateWalletRequest("TenantId_example", "PoolId_example", openapiclient.WalletOwnerType("ORG"), "OwnerId_example") // CreateWalletRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.CreateWallet(context.Background()).CreateWalletRequest(createWalletRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.CreateWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWallet`: WalletResponse
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.CreateWallet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWalletRequest** | [**CreateWalletRequest**](CreateWalletRequest.md) |  | 

### Return type

[**WalletResponse**](WalletResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWallet

> WalletResponse GetWallet(ctx, walletId).Execute()

Get Wallet



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.GetWallet(context.Background(), walletId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWallet`: WalletResponse
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetWallet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WalletResponse**](WalletResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletActivity

> interface{} GetWalletActivity(ctx, walletId).FromEffectiveAt(fromEffectiveAt).ToEffectiveAt(toEffectiveAt).EffectiveAsOf(effectiveAsOf).RecordedAsOf(recordedAsOf).ConsistentView(consistentView).Limit(limit).Cursor(cursor).Execute()

Get Wallet Activity



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
	fromEffectiveAt := time.Now() // time.Time | Start time for activity range (optional)
	toEffectiveAt := time.Now() // time.Time | End time for activity range (optional)
	effectiveAsOf := time.Now() // time.Time | Effective as-of timestamp (business time) for time travel (optional)
	recordedAsOf := time.Now() // time.Time | Recorded as-of timestamp (system time) for time travel (optional)
	consistentView := true // bool | Use strong consistency for reads (optional) (default to false)
	limit := int32(56) // int32 | Maximum number of items to return (optional) (default to 50)
	cursor := "cursor_example" // string | Pagination cursor (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.GetWalletActivity(context.Background(), walletId).FromEffectiveAt(fromEffectiveAt).ToEffectiveAt(toEffectiveAt).EffectiveAsOf(effectiveAsOf).RecordedAsOf(recordedAsOf).ConsistentView(consistentView).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetWalletActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletActivity`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetWalletActivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromEffectiveAt** | **time.Time** | Start time for activity range | 
 **toEffectiveAt** | **time.Time** | End time for activity range | 
 **effectiveAsOf** | **time.Time** | Effective as-of timestamp (business time) for time travel | 
 **recordedAsOf** | **time.Time** | Recorded as-of timestamp (system time) for time travel | 
 **consistentView** | **bool** | Use strong consistency for reads | [default to false]
 **limit** | **int32** | Maximum number of items to return | [default to 50]
 **cursor** | **string** | Pagination cursor | 

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


## GetWalletMembers

> interface{} GetWalletMembers(ctx, walletId).Execute()

Get Wallet Members



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.GetWalletMembers(context.Background(), walletId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetWalletMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletMembers`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetWalletMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ListWalletMembers

> interface{} ListWalletMembers(ctx).TenantId(tenantId).PoolId(poolId).SubjectId(subjectId).WalletId(walletId).Execute()

List Wallet Members



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
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Credit pool identifier
	subjectId := "subjectId_example" // string | Optional subject_id filter (optional)
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional wallet_id filter (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.ListWalletMembers(context.Background()).TenantId(tenantId).PoolId(poolId).SubjectId(subjectId).WalletId(walletId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.ListWalletMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWalletMembers`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.ListWalletMembers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWalletMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Tenant identifier | 
 **poolId** | **string** | Credit pool identifier | 
 **subjectId** | **string** | Optional subject_id filter | 
 **walletId** | **string** | Optional wallet_id filter | 

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


## ListWallets

> interface{} ListWallets(ctx).TenantId(tenantId).PoolId(poolId).SubjectId(subjectId).SubscriptionId(subscriptionId).OwnerType(ownerType).OwnerId(ownerId).EffectiveAsOf(effectiveAsOf).RecordedAsOf(recordedAsOf).ConsistentView(consistentView).Execute()

List Wallets



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
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Credit pool identifier (optional if subject_id or subscription_id provided) (optional)
	subjectId := "subjectId_example" // string | Subject identifier (if provided, tenant_id and pool_id are auto-resolved) (optional)
	subscriptionId := "subscriptionId_example" // string | Subscription identifier (finds wallets via grants with this subscription_id) (optional)
	ownerType := "ownerType_example" // string | Optional owner_type filter (optional)
	ownerId := "ownerId_example" // string | Optional owner_id filter (optional)
	effectiveAsOf := time.Now() // time.Time | Effective as-of timestamp (business time) for time travel (optional)
	recordedAsOf := time.Now() // time.Time | Recorded as-of timestamp (system time) for time travel (optional)
	consistentView := true // bool | Use strong consistency for reads (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.ListWallets(context.Background()).TenantId(tenantId).PoolId(poolId).SubjectId(subjectId).SubscriptionId(subscriptionId).OwnerType(ownerType).OwnerId(ownerId).EffectiveAsOf(effectiveAsOf).RecordedAsOf(recordedAsOf).ConsistentView(consistentView).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.ListWallets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWallets`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.ListWallets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWalletsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Tenant identifier (optional if subject_id or subscription_id provided) | 
 **poolId** | **string** | Credit pool identifier (optional if subject_id or subscription_id provided) | 
 **subjectId** | **string** | Subject identifier (if provided, tenant_id and pool_id are auto-resolved) | 
 **subscriptionId** | **string** | Subscription identifier (finds wallets via grants with this subscription_id) | 
 **ownerType** | **string** | Optional owner_type filter | 
 **ownerId** | **string** | Optional owner_id filter | 
 **effectiveAsOf** | **time.Time** | Effective as-of timestamp (business time) for time travel | 
 **recordedAsOf** | **time.Time** | Recorded as-of timestamp (system time) for time travel | 
 **consistentView** | **bool** | Use strong consistency for reads | [default to false]

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


## ListWalletsBySubjectsV1

> interface{} ListWalletsBySubjectsV1(ctx).BatchWalletsBySubjectsRequest(batchWalletsBySubjectsRequest).Execute()

List Wallets By Subjects



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
	batchWalletsBySubjectsRequest := *openapiclient.NewBatchWalletsBySubjectsRequest([]string{"SubjectIds_example"}) // BatchWalletsBySubjectsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.ListWalletsBySubjectsV1(context.Background()).BatchWalletsBySubjectsRequest(batchWalletsBySubjectsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.ListWalletsBySubjectsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWalletsBySubjectsV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.ListWalletsBySubjectsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWalletsBySubjectsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchWalletsBySubjectsRequest** | [**BatchWalletsBySubjectsRequest**](BatchWalletsBySubjectsRequest.md) |  | 

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


## ResolveTenantAndPool

> interface{} ResolveTenantAndPool(ctx).SubjectId(subjectId).Execute()

Resolve Tenant And Pool



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
	subjectId := "subjectId_example" // string | Subject identifier (customer ID or subject key)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.ResolveTenantAndPool(context.Background()).SubjectId(subjectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.ResolveTenantAndPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResolveTenantAndPool`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.ResolveTenantAndPool`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResolveTenantAndPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subjectId** | **string** | Subject identifier (customer ID or subject key) | 

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


## UpdateWalletSettings

> WalletResponse UpdateWalletSettings(ctx, walletId).UpdateWalletSettingsRequest(updateWalletSettingsRequest).TenantId(tenantId).PoolId(poolId).Execute()

Update Wallet Settings



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	updateWalletSettingsRequest := *openapiclient.NewUpdateWalletSettingsRequest() // UpdateWalletSettingsRequest | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional tenant identifier for validation (optional)
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional pool identifier for validation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.UpdateWalletSettings(context.Background(), walletId).UpdateWalletSettingsRequest(updateWalletSettingsRequest).TenantId(tenantId).PoolId(poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.UpdateWalletSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWalletSettings`: WalletResponse
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.UpdateWalletSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWalletSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateWalletSettingsRequest** | [**UpdateWalletSettingsRequest**](UpdateWalletSettingsRequest.md) |  | 
 **tenantId** | **string** | Optional tenant identifier for validation | 
 **poolId** | **string** | Optional pool identifier for validation | 

### Return type

[**WalletResponse**](WalletResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWalletThresholds

> WalletResponse UpdateWalletThresholds(ctx, walletId).UpdateWalletThresholdsRequest(updateWalletThresholdsRequest).TenantId(tenantId).PoolId(poolId).Execute()

Update Wallet Thresholds



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	updateWalletThresholdsRequest := *openapiclient.NewUpdateWalletThresholdsRequest() // UpdateWalletThresholdsRequest | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional tenant identifier for validation (optional)
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional pool identifier for validation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.UpdateWalletThresholds(context.Background(), walletId).UpdateWalletThresholdsRequest(updateWalletThresholdsRequest).TenantId(tenantId).PoolId(poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.UpdateWalletThresholds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWalletThresholds`: WalletResponse
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.UpdateWalletThresholds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWalletThresholdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateWalletThresholdsRequest** | [**UpdateWalletThresholdsRequest**](UpdateWalletThresholdsRequest.md) |  | 
 **tenantId** | **string** | Optional tenant identifier for validation | 
 **poolId** | **string** | Optional pool identifier for validation | 

### Return type

[**WalletResponse**](WalletResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

