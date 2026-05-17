# \NetsuiteSyncAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRun**](NetsuiteSyncAPI.md#GetRun) | **Get** /v1/integrations/netsuite/sync/runs/{sync_run_id} | Get Run
[**ListCreditMemosV1**](NetsuiteSyncAPI.md#ListCreditMemosV1) | **Get** /v1/integrations/netsuite/sync/credit-memos | List Credit Memos
[**ListCustomers**](NetsuiteSyncAPI.md#ListCustomers) | **Get** /v1/integrations/netsuite/sync/customers | List Customers
[**ListInvoices**](NetsuiteSyncAPI.md#ListInvoices) | **Get** /v1/integrations/netsuite/sync/invoices | List Invoices
[**ListPayments**](NetsuiteSyncAPI.md#ListPayments) | **Get** /v1/integrations/netsuite/sync/payments | List Payments
[**ListRuns**](NetsuiteSyncAPI.md#ListRuns) | **Get** /v1/integrations/netsuite/sync/runs | List Runs
[**TriggerSync**](NetsuiteSyncAPI.md#TriggerSync) | **Post** /v1/integrations/netsuite/sync/trigger | Trigger Sync



## GetRun

> SyncRunItem GetRun(ctx, syncRunId).Execute()

Get Run

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
	syncRunId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetsuiteSyncAPI.GetRun(context.Background(), syncRunId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetsuiteSyncAPI.GetRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRun`: SyncRunItem
	fmt.Fprintf(os.Stdout, "Response from `NetsuiteSyncAPI.GetRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**syncRunId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SyncRunItem**](SyncRunItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCreditMemosV1

> PaginatedCreditMemos ListCreditMemosV1(ctx).CustomerId(customerId).ModifiedSince(modifiedSince).Limit(limit).Cursor(cursor).Execute()

List Credit Memos

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
	customerId := "customerId_example" // string |  (optional)
	modifiedSince := time.Now() // time.Time |  (optional)
	limit := int32(56) // int32 |  (optional)
	cursor := "cursor_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetsuiteSyncAPI.ListCreditMemosV1(context.Background()).CustomerId(customerId).ModifiedSince(modifiedSince).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetsuiteSyncAPI.ListCreditMemosV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCreditMemosV1`: PaginatedCreditMemos
	fmt.Fprintf(os.Stdout, "Response from `NetsuiteSyncAPI.ListCreditMemosV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCreditMemosV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerId** | **string** |  | 
 **modifiedSince** | **time.Time** |  | 
 **limit** | **int32** |  | 
 **cursor** | **string** |  | 

### Return type

[**PaginatedCreditMemos**](PaginatedCreditMemos.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCustomers

> PaginatedCustomers ListCustomers(ctx).ModifiedSince(modifiedSince).Limit(limit).Cursor(cursor).NetsuiteId(netsuiteId).Execute()

List Customers

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
	modifiedSince := time.Now() // time.Time |  (optional)
	limit := int32(56) // int32 |  (optional)
	cursor := "cursor_example" // string |  (optional)
	netsuiteId := "netsuiteId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetsuiteSyncAPI.ListCustomers(context.Background()).ModifiedSince(modifiedSince).Limit(limit).Cursor(cursor).NetsuiteId(netsuiteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetsuiteSyncAPI.ListCustomers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCustomers`: PaginatedCustomers
	fmt.Fprintf(os.Stdout, "Response from `NetsuiteSyncAPI.ListCustomers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCustomersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modifiedSince** | **time.Time** |  | 
 **limit** | **int32** |  | 
 **cursor** | **string** |  | 
 **netsuiteId** | **string** |  | 

### Return type

[**PaginatedCustomers**](PaginatedCustomers.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInvoices

> PaginatedInvoices ListInvoices(ctx).CustomerId(customerId).Status(status).ModifiedSince(modifiedSince).Limit(limit).Cursor(cursor).Execute()

List Invoices

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
	customerId := "customerId_example" // string |  (optional)
	status := "status_example" // string |  (optional)
	modifiedSince := time.Now() // time.Time |  (optional)
	limit := int32(56) // int32 |  (optional)
	cursor := "cursor_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetsuiteSyncAPI.ListInvoices(context.Background()).CustomerId(customerId).Status(status).ModifiedSince(modifiedSince).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetsuiteSyncAPI.ListInvoices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInvoices`: PaginatedInvoices
	fmt.Fprintf(os.Stdout, "Response from `NetsuiteSyncAPI.ListInvoices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerId** | **string** |  | 
 **status** | **string** |  | 
 **modifiedSince** | **time.Time** |  | 
 **limit** | **int32** |  | 
 **cursor** | **string** |  | 

### Return type

[**PaginatedInvoices**](PaginatedInvoices.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPayments

> PaginatedPayments ListPayments(ctx).CustomerId(customerId).ModifiedSince(modifiedSince).Limit(limit).Cursor(cursor).Execute()

List Payments

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
	customerId := "customerId_example" // string |  (optional)
	modifiedSince := time.Now() // time.Time |  (optional)
	limit := int32(56) // int32 |  (optional)
	cursor := "cursor_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetsuiteSyncAPI.ListPayments(context.Background()).CustomerId(customerId).ModifiedSince(modifiedSince).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetsuiteSyncAPI.ListPayments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPayments`: PaginatedPayments
	fmt.Fprintf(os.Stdout, "Response from `NetsuiteSyncAPI.ListPayments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerId** | **string** |  | 
 **modifiedSince** | **time.Time** |  | 
 **limit** | **int32** |  | 
 **cursor** | **string** |  | 

### Return type

[**PaginatedPayments**](PaginatedPayments.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRuns

> PaginatedSyncRuns ListRuns(ctx).Status(status).Limit(limit).Cursor(cursor).Execute()

List Runs

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
	status := "status_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)
	cursor := "cursor_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetsuiteSyncAPI.ListRuns(context.Background()).Status(status).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetsuiteSyncAPI.ListRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRuns`: PaginatedSyncRuns
	fmt.Fprintf(os.Stdout, "Response from `NetsuiteSyncAPI.ListRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** |  | 
 **limit** | **int32** |  | 
 **cursor** | **string** |  | 

### Return type

[**PaginatedSyncRuns**](PaginatedSyncRuns.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TriggerSync

> AppApiV1IntegrationsNetsuiteSyncRouterTriggerResponse TriggerSync(ctx).TriggerRequest(triggerRequest).Execute()

Trigger Sync

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
	triggerRequest := *openapiclient.NewTriggerRequest() // TriggerRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetsuiteSyncAPI.TriggerSync(context.Background()).TriggerRequest(triggerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetsuiteSyncAPI.TriggerSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TriggerSync`: AppApiV1IntegrationsNetsuiteSyncRouterTriggerResponse
	fmt.Fprintf(os.Stdout, "Response from `NetsuiteSyncAPI.TriggerSync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTriggerSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **triggerRequest** | [**TriggerRequest**](TriggerRequest.md) |  | 

### Return type

[**AppApiV1IntegrationsNetsuiteSyncRouterTriggerResponse**](AppApiV1IntegrationsNetsuiteSyncRouterTriggerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

