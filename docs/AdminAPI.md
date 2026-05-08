# \AdminAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DiagnoseSubscription**](AdminAPI.md#DiagnoseSubscription) | **Get** /v1/admin/diagnostics/subscription/{subscription_id} | Diagnose Subscription
[**GetLedgerAuditGet**](AdminAPI.md#GetLedgerAuditGet) | **Get** /v1/admin/ledger/audit | Get Ledger Audit
[**GetLedgerBalanceGet**](AdminAPI.md#GetLedgerBalanceGet) | **Get** /v1/admin/ledger/balance | Get Ledger Balance
[**ProcessLifecycleEventManuallyV1Admin**](AdminAPI.md#ProcessLifecycleEventManuallyV1Admin) | **Post** /v1/admin/subscriptions/{subscription_id}/process-lifecycle-event | Process Lifecycle Event Manually
[**RetryUnpriced**](AdminAPI.md#RetryUnpriced) | **Post** /v1/admin/usage/unpriced/retry | Retry Unpriced
[**ReviewQuarantined**](AdminAPI.md#ReviewQuarantined) | **Post** /v1/admin/usage/quarantined/review | Review Quarantined
[**TriggerSubscriptionSync**](AdminAPI.md#TriggerSubscriptionSync) | **Post** /v1/admin/subscriptions/{subscription_id}/sync | Trigger Subscription Sync



## DiagnoseSubscription

> interface{} DiagnoseSubscription(ctx, subscriptionId).TenantId(tenantId).Execute()

Diagnose Subscription



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
	subscriptionId := "subscriptionId_example" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.DiagnoseSubscription(context.Background(), subscriptionId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.DiagnoseSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiagnoseSubscription`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.DiagnoseSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDiagnoseSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** | Tenant identifier | 

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


## GetLedgerAuditGet

> interface{} GetLedgerAuditGet(ctx).TenantId(tenantId).PoolId(poolId).WalletId(walletId).JournalEntryId(journalEntryId).UsageEventId(usageEventId).FromEffectiveAt(fromEffectiveAt).ToEffectiveAt(toEffectiveAt).Limit(limit).Execute()

Get Ledger Audit



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional wallet filter (optional)
	journalEntryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional journal entry filter (optional)
	usageEventId := "usageEventId_example" // string | Optional usage event filter (optional)
	fromEffectiveAt := time.Now() // time.Time | Optional start time filter (optional)
	toEffectiveAt := time.Now() // time.Time | Optional end time filter (optional)
	limit := int32(56) // int32 | Maximum number of entries (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.GetLedgerAuditGet(context.Background()).TenantId(tenantId).PoolId(poolId).WalletId(walletId).JournalEntryId(journalEntryId).UsageEventId(usageEventId).FromEffectiveAt(fromEffectiveAt).ToEffectiveAt(toEffectiveAt).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.GetLedgerAuditGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLedgerAuditGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.GetLedgerAuditGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLedgerAuditGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Tenant identifier | 
 **poolId** | **string** | Pool identifier | 
 **walletId** | **string** | Optional wallet filter | 
 **journalEntryId** | **string** | Optional journal entry filter | 
 **usageEventId** | **string** | Optional usage event filter | 
 **fromEffectiveAt** | **time.Time** | Optional start time filter | 
 **toEffectiveAt** | **time.Time** | Optional end time filter | 
 **limit** | **int32** | Maximum number of entries | [default to 100]

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


## GetLedgerBalanceGet

> interface{} GetLedgerBalanceGet(ctx).TenantId(tenantId).PoolId(poolId).WalletId(walletId).AsOfEffectiveAt(asOfEffectiveAt).AsOfRecordedAt(asOfRecordedAt).Execute()

Get Ledger Balance



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Wallet identifier
	asOfEffectiveAt := time.Now() // time.Time | Effective timestamp (for time travel)
	asOfRecordedAt := time.Now() // time.Time | Recorded timestamp (for time travel) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.GetLedgerBalanceGet(context.Background()).TenantId(tenantId).PoolId(poolId).WalletId(walletId).AsOfEffectiveAt(asOfEffectiveAt).AsOfRecordedAt(asOfRecordedAt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.GetLedgerBalanceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLedgerBalanceGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.GetLedgerBalanceGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLedgerBalanceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Tenant identifier | 
 **poolId** | **string** | Pool identifier | 
 **walletId** | **string** | Wallet identifier | 
 **asOfEffectiveAt** | **time.Time** | Effective timestamp (for time travel) | 
 **asOfRecordedAt** | **time.Time** | Recorded timestamp (for time travel) | 

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


## ProcessLifecycleEventManuallyV1Admin

> interface{} ProcessLifecycleEventManuallyV1Admin(ctx, subscriptionId).TenantId(tenantId).Execute()

Process Lifecycle Event Manually



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
	subscriptionId := "subscriptionId_example" // string | Subscription ID
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.ProcessLifecycleEventManuallyV1Admin(context.Background(), subscriptionId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.ProcessLifecycleEventManuallyV1Admin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProcessLifecycleEventManuallyV1Admin`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.ProcessLifecycleEventManuallyV1Admin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProcessLifecycleEventManuallyV1AdminRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** | Tenant identifier | 

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


## RetryUnpriced

> RetryUnpricedResponse RetryUnpriced(ctx).RetryUnpricedRequest(retryUnpricedRequest).Execute()

Retry Unpriced



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
	retryUnpricedRequest := *openapiclient.NewRetryUnpricedRequest("TenantId_example", "PoolId_example", "UsageEventId_example") // RetryUnpricedRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.RetryUnpriced(context.Background()).RetryUnpricedRequest(retryUnpricedRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.RetryUnpriced``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetryUnpriced`: RetryUnpricedResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.RetryUnpriced`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetryUnpricedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **retryUnpricedRequest** | [**RetryUnpricedRequest**](RetryUnpricedRequest.md) |  | 

### Return type

[**RetryUnpricedResponse**](RetryUnpricedResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReviewQuarantined

> ReviewQuarantinedResponse ReviewQuarantined(ctx).ReviewQuarantinedRequest(reviewQuarantinedRequest).Execute()

Review Quarantined



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
	reviewQuarantinedRequest := *openapiclient.NewReviewQuarantinedRequest("TenantId_example", "PoolId_example", "UsageEventId_example", "Action_example") // ReviewQuarantinedRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.ReviewQuarantined(context.Background()).ReviewQuarantinedRequest(reviewQuarantinedRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.ReviewQuarantined``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReviewQuarantined`: ReviewQuarantinedResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.ReviewQuarantined`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReviewQuarantinedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reviewQuarantinedRequest** | [**ReviewQuarantinedRequest**](ReviewQuarantinedRequest.md) |  | 

### Return type

[**ReviewQuarantinedResponse**](ReviewQuarantinedResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TriggerSubscriptionSync

> interface{} TriggerSubscriptionSync(ctx, subscriptionId).Execute()

Trigger Subscription Sync



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
	subscriptionId := "subscriptionId_example" // string | Subscription ID to sync

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.TriggerSubscriptionSync(context.Background(), subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.TriggerSubscriptionSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TriggerSubscriptionSync`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.TriggerSubscriptionSync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription ID to sync | 

### Other Parameters

Other parameters are passed through a pointer to a apiTriggerSubscriptionSyncRequest struct via the builder pattern


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

