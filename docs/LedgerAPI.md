# \LedgerAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLedgerAudit**](LedgerAPI.md#GetLedgerAudit) | **Get** /v1/ledger/audit | Get Ledger Audit
[**GetLedgerBalance**](LedgerAPI.md#GetLedgerBalance) | **Get** /v1/ledger/balance | Get Ledger Balance
[**GetWalletState**](LedgerAPI.md#GetWalletState) | **Get** /v1/ledger/state | Get Wallet State
[**GetWalletStatesBatch**](LedgerAPI.md#GetWalletStatesBatch) | **Post** /v1/ledger/state/batch | Get Wallet States Batch



## GetLedgerAudit

> interface{} GetLedgerAudit(ctx).TenantId(tenantId).PoolId(poolId).WalletId(walletId).JournalEntryId(journalEntryId).UsageEventId(usageEventId).FromEffectiveAt(fromEffectiveAt).ToEffectiveAt(toEffectiveAt).Limit(limit).Execute()

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
	resp, r, err := apiClient.LedgerAPI.GetLedgerAudit(context.Background()).TenantId(tenantId).PoolId(poolId).WalletId(walletId).JournalEntryId(journalEntryId).UsageEventId(usageEventId).FromEffectiveAt(fromEffectiveAt).ToEffectiveAt(toEffectiveAt).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LedgerAPI.GetLedgerAudit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLedgerAudit`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `LedgerAPI.GetLedgerAudit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLedgerAuditRequest struct via the builder pattern


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


## GetLedgerBalance

> interface{} GetLedgerBalance(ctx).TenantId(tenantId).PoolId(poolId).WalletId(walletId).AsOfEffectiveAt(asOfEffectiveAt).AsOfRecordedAt(asOfRecordedAt).Execute()

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
	resp, r, err := apiClient.LedgerAPI.GetLedgerBalance(context.Background()).TenantId(tenantId).PoolId(poolId).WalletId(walletId).AsOfEffectiveAt(asOfEffectiveAt).AsOfRecordedAt(asOfRecordedAt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LedgerAPI.GetLedgerBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLedgerBalance`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `LedgerAPI.GetLedgerBalance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLedgerBalanceRequest struct via the builder pattern


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


## GetWalletState

> interface{} GetWalletState(ctx).WalletId(walletId).TenantId(tenantId).PoolId(poolId).AsOf(asOf).Consistency(consistency).EffectiveAsOf(effectiveAsOf).RecordedAsOf(recordedAsOf).ConsistentView(consistentView).Execute()

Get Wallet State



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier (resolved from auth when absent) (optional)
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Credit pool identifier (resolved from auth when absent) (optional)
	asOf := time.Now() // time.Time | As-of timestamp (for time travel, legacy, use effective_as_of instead) (optional)
	consistency := "consistency_example" // string | Consistency level: 'eventual' or 'STRONG' (optional) (default to "eventual")
	effectiveAsOf := time.Now() // time.Time | Effective as-of timestamp (business time) for time travel (optional)
	recordedAsOf := time.Now() // time.Time | Recorded as-of timestamp (system time) for time travel (optional)
	consistentView := true // bool | Use strong consistency (overrides consistency param if True) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LedgerAPI.GetWalletState(context.Background()).WalletId(walletId).TenantId(tenantId).PoolId(poolId).AsOf(asOf).Consistency(consistency).EffectiveAsOf(effectiveAsOf).RecordedAsOf(recordedAsOf).ConsistentView(consistentView).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LedgerAPI.GetWalletState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletState`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `LedgerAPI.GetWalletState`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **walletId** | **string** | Wallet identifier | 
 **tenantId** | **string** | Tenant identifier (resolved from auth when absent) | 
 **poolId** | **string** | Credit pool identifier (resolved from auth when absent) | 
 **asOf** | **time.Time** | As-of timestamp (for time travel, legacy, use effective_as_of instead) | 
 **consistency** | **string** | Consistency level: &#39;eventual&#39; or &#39;STRONG&#39; | [default to &quot;eventual&quot;]
 **effectiveAsOf** | **time.Time** | Effective as-of timestamp (business time) for time travel | 
 **recordedAsOf** | **time.Time** | Recorded as-of timestamp (system time) for time travel | 
 **consistentView** | **bool** | Use strong consistency (overrides consistency param if True) | 

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


## GetWalletStatesBatch

> BatchWalletStateResponse GetWalletStatesBatch(ctx).BatchWalletStateRequest(batchWalletStateRequest).Execute()

Get Wallet States Batch



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
	batchWalletStateRequest := *openapiclient.NewBatchWalletStateRequest([]openapiclient.BatchWalletStateRequestItem{*openapiclient.NewBatchWalletStateRequestItem("TenantId_example", "PoolId_example", "WalletId_example")}) // BatchWalletStateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LedgerAPI.GetWalletStatesBatch(context.Background()).BatchWalletStateRequest(batchWalletStateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LedgerAPI.GetWalletStatesBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletStatesBatch`: BatchWalletStateResponse
	fmt.Fprintf(os.Stdout, "Response from `LedgerAPI.GetWalletStatesBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletStatesBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchWalletStateRequest** | [**BatchWalletStateRequest**](BatchWalletStateRequest.md) |  | 

### Return type

[**BatchWalletStateResponse**](BatchWalletStateResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

