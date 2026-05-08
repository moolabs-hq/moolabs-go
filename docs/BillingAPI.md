# \BillingAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdvanceInvoiceAction**](BillingAPI.md#AdvanceInvoiceAction) | **Post** /api/v1/billing/invoices/{invoiceId}/advance | Advance the invoice&#39;s state to the next status
[**ApproveInvoiceAction**](BillingAPI.md#ApproveInvoiceAction) | **Post** /api/v1/billing/invoices/{invoiceId}/approve | Send the invoice to the customer
[**CreateBillingProfile**](BillingAPI.md#CreateBillingProfile) | **Post** /api/v1/billing/profiles | Create a new billing profile
[**CreatePendingInvoiceLine**](BillingAPI.md#CreatePendingInvoiceLine) | **Post** /api/v1/billing/customers/{customerId}/invoices/pending-lines | Create pending line items
[**DeleteBillingProfile**](BillingAPI.md#DeleteBillingProfile) | **Delete** /api/v1/billing/profiles/{id} | Delete a billing profile
[**DeleteBillingProfileCustomerOverride**](BillingAPI.md#DeleteBillingProfileCustomerOverride) | **Delete** /api/v1/billing/customers/{customerId} | Delete a customer override
[**DeleteInvoice**](BillingAPI.md#DeleteInvoice) | **Delete** /api/v1/billing/invoices/{invoiceId} | Delete an invoice
[**GetBillingProfile**](BillingAPI.md#GetBillingProfile) | **Get** /api/v1/billing/profiles/{id} | Get a billing profile
[**GetBillingProfileCustomerOverride**](BillingAPI.md#GetBillingProfileCustomerOverride) | **Get** /api/v1/billing/customers/{customerId} | Get a customer override
[**GetInvoice**](BillingAPI.md#GetInvoice) | **Get** /api/v1/billing/invoices/{invoiceId} | Get an invoice
[**InvoicePendingLinesAction**](BillingAPI.md#InvoicePendingLinesAction) | **Post** /api/v1/billing/invoices/invoice | Invoice a customer based on the pending line items
[**ListBillingProfileCustomerOverrides**](BillingAPI.md#ListBillingProfileCustomerOverrides) | **Get** /api/v1/billing/customers | List customer overrides
[**ListBillingProfiles**](BillingAPI.md#ListBillingProfiles) | **Get** /api/v1/billing/profiles | List billing profiles
[**ListInvoices**](BillingAPI.md#ListInvoices) | **Get** /api/v1/billing/invoices | List invoices
[**RecalculateInvoiceTaxAction**](BillingAPI.md#RecalculateInvoiceTaxAction) | **Post** /api/v1/billing/invoices/{invoiceId}/taxes/recalculate | Recalculate an invoice&#39;s tax amounts
[**RetryInvoiceAction**](BillingAPI.md#RetryInvoiceAction) | **Post** /api/v1/billing/invoices/{invoiceId}/retry | Retry advancing the invoice after a failed attempt.
[**SimulateInvoice**](BillingAPI.md#SimulateInvoice) | **Post** /api/v1/billing/customers/{customerId}/invoices/simulate | Simulate an invoice for a customer
[**SnapshotQuantitiesInvoiceAction**](BillingAPI.md#SnapshotQuantitiesInvoiceAction) | **Post** /api/v1/billing/invoices/{invoiceId}/snapshot-quantities | Snapshot quantities for usage based line items
[**UpdateBillingProfile**](BillingAPI.md#UpdateBillingProfile) | **Put** /api/v1/billing/profiles/{id} | Update a billing profile
[**UpdateInvoice**](BillingAPI.md#UpdateInvoice) | **Put** /api/v1/billing/invoices/{invoiceId} | Update an invoice
[**UpdateInvoicePaymentStatus**](BillingAPI.md#UpdateInvoicePaymentStatus) | **Post** /api/v1/billing/invoices/{invoiceId}/payment-status | Update invoice payment status
[**UpsertBillingProfileCustomerOverride**](BillingAPI.md#UpsertBillingProfileCustomerOverride) | **Put** /api/v1/billing/customers/{customerId} | Create a new or update a customer override
[**VoidInvoiceAction**](BillingAPI.md#VoidInvoiceAction) | **Post** /api/v1/billing/invoices/{invoiceId}/void | Void an invoice



## AdvanceInvoiceAction

> Invoice AdvanceInvoiceAction(ctx, invoiceId).Execute()

Advance the invoice's state to the next status



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
	invoiceId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.AdvanceInvoiceAction(context.Background(), invoiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.AdvanceInvoiceAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdvanceInvoiceAction`: Invoice
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.AdvanceInvoiceAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdvanceInvoiceActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Invoice**](Invoice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApproveInvoiceAction

> Invoice ApproveInvoiceAction(ctx, invoiceId).Execute()

Send the invoice to the customer



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
	invoiceId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.ApproveInvoiceAction(context.Background(), invoiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.ApproveInvoiceAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApproveInvoiceAction`: Invoice
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.ApproveInvoiceAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApproveInvoiceActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Invoice**](Invoice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBillingProfile

> BillingProfile CreateBillingProfile(ctx).BillingProfileCreate(billingProfileCreate).Execute()

Create a new billing profile



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
	billingProfileCreate := *openapiclient.NewBillingProfileCreate("Name_example", *openapiclient.NewBillingParty(), false, *openapiclient.NewBillingWorkflowCreate(), *openapiclient.NewBillingProfileAppsCreate("01G65Z755AFWAKHE12NY0CQ9FH", "01G65Z755AFWAKHE12NY0CQ9FH", "01G65Z755AFWAKHE12NY0CQ9FH")) // BillingProfileCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.CreateBillingProfile(context.Background()).BillingProfileCreate(billingProfileCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.CreateBillingProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBillingProfile`: BillingProfile
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.CreateBillingProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBillingProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **billingProfileCreate** | [**BillingProfileCreate**](BillingProfileCreate.md) |  | 

### Return type

[**BillingProfile**](BillingProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePendingInvoiceLine

> InvoicePendingLineCreateResponse CreatePendingInvoiceLine(ctx, customerId).InvoicePendingLineCreateInput(invoicePendingLineCreateInput).Execute()

Create pending line items



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
	customerId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 
	invoicePendingLineCreateInput := *openapiclient.NewInvoicePendingLineCreateInput("USD", []openapiclient.InvoicePendingLineCreate{*openapiclient.NewInvoicePendingLineCreate("Name_example", *openapiclient.NewPeriod(time.Now(), time.Now()), time.Now())}) // InvoicePendingLineCreateInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.CreatePendingInvoiceLine(context.Background(), customerId).InvoicePendingLineCreateInput(invoicePendingLineCreateInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.CreatePendingInvoiceLine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePendingInvoiceLine`: InvoicePendingLineCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.CreatePendingInvoiceLine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePendingInvoiceLineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invoicePendingLineCreateInput** | [**InvoicePendingLineCreateInput**](InvoicePendingLineCreateInput.md) |  | 

### Return type

[**InvoicePendingLineCreateResponse**](InvoicePendingLineCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBillingProfile

> DeleteBillingProfile(ctx, id).Execute()

Delete a billing profile



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
	id := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BillingAPI.DeleteBillingProfile(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.DeleteBillingProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBillingProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBillingProfileCustomerOverride

> DeleteBillingProfileCustomerOverride(ctx, customerId).Execute()

Delete a customer override



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
	customerId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BillingAPI.DeleteBillingProfileCustomerOverride(context.Background(), customerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.DeleteBillingProfileCustomerOverride``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBillingProfileCustomerOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInvoice

> DeleteInvoice(ctx, invoiceId).Execute()

Delete an invoice



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
	invoiceId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BillingAPI.DeleteInvoice(context.Background(), invoiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.DeleteInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillingProfile

> BillingProfile GetBillingProfile(ctx, id).Expand(expand).Execute()

Get a billing profile



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
	id := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 
	expand := []openapiclient.BillingProfileExpand{openapiclient.BillingProfileExpand("apps")} // []BillingProfileExpand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetBillingProfile(context.Background(), id).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetBillingProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillingProfile`: BillingProfile
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetBillingProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | [**[]BillingProfileExpand**](BillingProfileExpand.md) |  | 

### Return type

[**BillingProfile**](BillingProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillingProfileCustomerOverride

> BillingProfileCustomerOverrideWithDetails GetBillingProfileCustomerOverride(ctx, customerId).Expand(expand).Execute()

Get a customer override



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
	customerId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 
	expand := []openapiclient.BillingProfileCustomerOverrideExpand{openapiclient.BillingProfileCustomerOverrideExpand("apps")} // []BillingProfileCustomerOverrideExpand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetBillingProfileCustomerOverride(context.Background(), customerId).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetBillingProfileCustomerOverride``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillingProfileCustomerOverride`: BillingProfileCustomerOverrideWithDetails
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetBillingProfileCustomerOverride`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingProfileCustomerOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | [**[]BillingProfileCustomerOverrideExpand**](BillingProfileCustomerOverrideExpand.md) |  | 

### Return type

[**BillingProfileCustomerOverrideWithDetails**](BillingProfileCustomerOverrideWithDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoice

> Invoice GetInvoice(ctx, invoiceId).Expand(expand).IncludeDeletedLines(includeDeletedLines).Execute()

Get an invoice



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
	invoiceId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 
	expand := []openapiclient.InvoiceExpand{openapiclient.InvoiceExpand("lines")} // []InvoiceExpand |  (optional)
	includeDeletedLines := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GetInvoice(context.Background(), invoiceId).Expand(expand).IncludeDeletedLines(includeDeletedLines).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvoice`: Invoice
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | [**[]InvoiceExpand**](InvoiceExpand.md) |  | 
 **includeDeletedLines** | **bool** |  | [default to false]

### Return type

[**Invoice**](Invoice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoicePendingLinesAction

> []Invoice InvoicePendingLinesAction(ctx).InvoicePendingLinesActionInput(invoicePendingLinesActionInput).Execute()

Invoice a customer based on the pending line items



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
	invoicePendingLinesActionInput := *openapiclient.NewInvoicePendingLinesActionInput("01G65Z755AFWAKHE12NY0CQ9FH") // InvoicePendingLinesActionInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.InvoicePendingLinesAction(context.Background()).InvoicePendingLinesActionInput(invoicePendingLinesActionInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.InvoicePendingLinesAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoicePendingLinesAction`: []Invoice
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.InvoicePendingLinesAction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoicePendingLinesActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invoicePendingLinesActionInput** | [**InvoicePendingLinesActionInput**](InvoicePendingLinesActionInput.md) |  | 

### Return type

[**[]Invoice**](Invoice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBillingProfileCustomerOverrides

> BillingProfileCustomerOverrideWithDetailsPaginatedResponse ListBillingProfileCustomerOverrides(ctx).BillingProfile(billingProfile).CustomersWithoutPinnedProfile(customersWithoutPinnedProfile).IncludeAllCustomers(includeAllCustomers).CustomerId(customerId).CustomerName(customerName).CustomerKey(customerKey).CustomerPrimaryEmail(customerPrimaryEmail).Expand(expand).Order(order).OrderBy(orderBy).Page(page).PageSize(pageSize).Execute()

List customer overrides



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
	billingProfile := []string{"01G65Z755AFWAKHE12NY0CQ9FH"} // []string | Filter by billing profile. (optional)
	customersWithoutPinnedProfile := true // bool | Only return customers without pinned billing profiles. This implicitly sets includeAllCustomers to true. (optional)
	includeAllCustomers := true // bool | Include customers without customer overrides.  If set to false only the customers specifically associated with a billing profile will be returned.  If set to true, in case of the default billing profile, all customers will be returned. (optional) (default to true)
	customerId := []string{"01G65Z755AFWAKHE12NY0CQ9FH"} // []string | Filter by customer id. (optional)
	customerName := "customerName_example" // string | Filter by customer name. (optional)
	customerKey := "customerKey_example" // string | Filter by customer key (optional)
	customerPrimaryEmail := "customerPrimaryEmail_example" // string | Filter by customer primary email (optional)
	expand := []openapiclient.BillingProfileCustomerOverrideExpand{openapiclient.BillingProfileCustomerOverrideExpand("apps")} // []BillingProfileCustomerOverrideExpand | Expand the response with additional details. (optional)
	order := openapiclient.SortOrder("ASC") // SortOrder | The order direction. (optional)
	orderBy := openapiclient.BillingProfileCustomerOverrideOrderBy("customerId") // BillingProfileCustomerOverrideOrderBy | The order by field. (optional)
	page := int32(56) // int32 | Page index.  Default is 1. (optional) (default to 1)
	pageSize := int32(56) // int32 | The maximum number of items per page.  Default is 100. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.ListBillingProfileCustomerOverrides(context.Background()).BillingProfile(billingProfile).CustomersWithoutPinnedProfile(customersWithoutPinnedProfile).IncludeAllCustomers(includeAllCustomers).CustomerId(customerId).CustomerName(customerName).CustomerKey(customerKey).CustomerPrimaryEmail(customerPrimaryEmail).Expand(expand).Order(order).OrderBy(orderBy).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.ListBillingProfileCustomerOverrides``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBillingProfileCustomerOverrides`: BillingProfileCustomerOverrideWithDetailsPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.ListBillingProfileCustomerOverrides`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBillingProfileCustomerOverridesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **billingProfile** | **[]string** | Filter by billing profile. | 
 **customersWithoutPinnedProfile** | **bool** | Only return customers without pinned billing profiles. This implicitly sets includeAllCustomers to true. | 
 **includeAllCustomers** | **bool** | Include customers without customer overrides.  If set to false only the customers specifically associated with a billing profile will be returned.  If set to true, in case of the default billing profile, all customers will be returned. | [default to true]
 **customerId** | **[]string** | Filter by customer id. | 
 **customerName** | **string** | Filter by customer name. | 
 **customerKey** | **string** | Filter by customer key | 
 **customerPrimaryEmail** | **string** | Filter by customer primary email | 
 **expand** | [**[]BillingProfileCustomerOverrideExpand**](BillingProfileCustomerOverrideExpand.md) | Expand the response with additional details. | 
 **order** | [**SortOrder**](SortOrder.md) | The order direction. | 
 **orderBy** | [**BillingProfileCustomerOverrideOrderBy**](BillingProfileCustomerOverrideOrderBy.md) | The order by field. | 
 **page** | **int32** | Page index.  Default is 1. | [default to 1]
 **pageSize** | **int32** | The maximum number of items per page.  Default is 100. | [default to 100]

### Return type

[**BillingProfileCustomerOverrideWithDetailsPaginatedResponse**](BillingProfileCustomerOverrideWithDetailsPaginatedResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBillingProfiles

> BillingProfilePaginatedResponse ListBillingProfiles(ctx).IncludeArchived(includeArchived).Expand(expand).Page(page).PageSize(pageSize).Order(order).OrderBy(orderBy).Execute()

List billing profiles



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
	includeArchived := true // bool |  (optional) (default to false)
	expand := []openapiclient.BillingProfileExpand{openapiclient.BillingProfileExpand("apps")} // []BillingProfileExpand |  (optional)
	page := int32(56) // int32 | Page index.  Default is 1. (optional) (default to 1)
	pageSize := int32(56) // int32 | The maximum number of items per page.  Default is 100. (optional) (default to 100)
	order := openapiclient.SortOrder("ASC") // SortOrder | The order direction. (optional)
	orderBy := openapiclient.BillingProfileOrderBy("createdAt") // BillingProfileOrderBy | The order by field. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.ListBillingProfiles(context.Background()).IncludeArchived(includeArchived).Expand(expand).Page(page).PageSize(pageSize).Order(order).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.ListBillingProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBillingProfiles`: BillingProfilePaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.ListBillingProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBillingProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeArchived** | **bool** |  | [default to false]
 **expand** | [**[]BillingProfileExpand**](BillingProfileExpand.md) |  | 
 **page** | **int32** | Page index.  Default is 1. | [default to 1]
 **pageSize** | **int32** | The maximum number of items per page.  Default is 100. | [default to 100]
 **order** | [**SortOrder**](SortOrder.md) | The order direction. | 
 **orderBy** | [**BillingProfileOrderBy**](BillingProfileOrderBy.md) | The order by field. | 

### Return type

[**BillingProfilePaginatedResponse**](BillingProfilePaginatedResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInvoices

> InvoicePaginatedResponse ListInvoices(ctx).Statuses(statuses).ExtendedStatuses(extendedStatuses).IssuedAfter(issuedAfter).IssuedBefore(issuedBefore).PeriodStartAfter(periodStartAfter).PeriodStartBefore(periodStartBefore).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Expand(expand).Customers(customers).IncludeDeleted(includeDeleted).Page(page).PageSize(pageSize).Order(order).OrderBy(orderBy).Execute()

List invoices



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
	statuses := []openapiclient.InvoiceStatus{openapiclient.InvoiceStatus("gathering")} // []InvoiceStatus | Filter by the invoice status. (optional)
	extendedStatuses := []string{"Inner_example"} // []string | Filter by invoice extended statuses (optional)
	issuedAfter := time.Now() // time.Time | Filter by invoice issued time. Inclusive. (optional)
	issuedBefore := time.Now() // time.Time | Filter by invoice issued time. Inclusive. (optional)
	periodStartAfter := time.Now() // time.Time | Filter by period start time. Inclusive. (optional)
	periodStartBefore := time.Now() // time.Time | Filter by period start time. Inclusive. (optional)
	createdAfter := time.Now() // time.Time | Filter by invoice created time. Inclusive. (optional)
	createdBefore := time.Now() // time.Time | Filter by invoice created time. Inclusive. (optional)
	expand := []openapiclient.InvoiceExpand{openapiclient.InvoiceExpand("lines")} // []InvoiceExpand | What parts of the list output to expand in listings (optional)
	customers := []string{"01G65Z755AFWAKHE12NY0CQ9FH"} // []string | Filter by customer ID (optional)
	includeDeleted := true // bool | Include deleted invoices (optional)
	page := int32(56) // int32 | Page index.  Default is 1. (optional) (default to 1)
	pageSize := int32(56) // int32 | The maximum number of items per page.  Default is 100. (optional) (default to 100)
	order := openapiclient.SortOrder("ASC") // SortOrder | The order direction. (optional)
	orderBy := openapiclient.InvoiceOrderBy("customer.name") // InvoiceOrderBy | The order by field. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.ListInvoices(context.Background()).Statuses(statuses).ExtendedStatuses(extendedStatuses).IssuedAfter(issuedAfter).IssuedBefore(issuedBefore).PeriodStartAfter(periodStartAfter).PeriodStartBefore(periodStartBefore).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Expand(expand).Customers(customers).IncludeDeleted(includeDeleted).Page(page).PageSize(pageSize).Order(order).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.ListInvoices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInvoices`: InvoicePaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.ListInvoices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **statuses** | [**[]InvoiceStatus**](InvoiceStatus.md) | Filter by the invoice status. | 
 **extendedStatuses** | **[]string** | Filter by invoice extended statuses | 
 **issuedAfter** | **time.Time** | Filter by invoice issued time. Inclusive. | 
 **issuedBefore** | **time.Time** | Filter by invoice issued time. Inclusive. | 
 **periodStartAfter** | **time.Time** | Filter by period start time. Inclusive. | 
 **periodStartBefore** | **time.Time** | Filter by period start time. Inclusive. | 
 **createdAfter** | **time.Time** | Filter by invoice created time. Inclusive. | 
 **createdBefore** | **time.Time** | Filter by invoice created time. Inclusive. | 
 **expand** | [**[]InvoiceExpand**](InvoiceExpand.md) | What parts of the list output to expand in listings | 
 **customers** | **[]string** | Filter by customer ID | 
 **includeDeleted** | **bool** | Include deleted invoices | 
 **page** | **int32** | Page index.  Default is 1. | [default to 1]
 **pageSize** | **int32** | The maximum number of items per page.  Default is 100. | [default to 100]
 **order** | [**SortOrder**](SortOrder.md) | The order direction. | 
 **orderBy** | [**InvoiceOrderBy**](InvoiceOrderBy.md) | The order by field. | 

### Return type

[**InvoicePaginatedResponse**](InvoicePaginatedResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecalculateInvoiceTaxAction

> Invoice RecalculateInvoiceTaxAction(ctx, invoiceId).Execute()

Recalculate an invoice's tax amounts



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
	invoiceId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.RecalculateInvoiceTaxAction(context.Background(), invoiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.RecalculateInvoiceTaxAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecalculateInvoiceTaxAction`: Invoice
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.RecalculateInvoiceTaxAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecalculateInvoiceTaxActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Invoice**](Invoice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetryInvoiceAction

> Invoice RetryInvoiceAction(ctx, invoiceId).Execute()

Retry advancing the invoice after a failed attempt.



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
	invoiceId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.RetryInvoiceAction(context.Background(), invoiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.RetryInvoiceAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetryInvoiceAction`: Invoice
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.RetryInvoiceAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetryInvoiceActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Invoice**](Invoice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SimulateInvoice

> Invoice SimulateInvoice(ctx, customerId).InvoiceSimulationInput(invoiceSimulationInput).Execute()

Simulate an invoice for a customer



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
	customerId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 
	invoiceSimulationInput := *openapiclient.NewInvoiceSimulationInput("USD", []openapiclient.InvoiceSimulationLine{*openapiclient.NewInvoiceSimulationLine("Name_example", *openapiclient.NewPeriod(time.Now(), time.Now()), time.Now(), "Quantity_example")}) // InvoiceSimulationInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.SimulateInvoice(context.Background(), customerId).InvoiceSimulationInput(invoiceSimulationInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.SimulateInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SimulateInvoice`: Invoice
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.SimulateInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSimulateInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invoiceSimulationInput** | [**InvoiceSimulationInput**](InvoiceSimulationInput.md) |  | 

### Return type

[**Invoice**](Invoice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnapshotQuantitiesInvoiceAction

> Invoice SnapshotQuantitiesInvoiceAction(ctx, invoiceId).Execute()

Snapshot quantities for usage based line items



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
	invoiceId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.SnapshotQuantitiesInvoiceAction(context.Background(), invoiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.SnapshotQuantitiesInvoiceAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SnapshotQuantitiesInvoiceAction`: Invoice
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.SnapshotQuantitiesInvoiceAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSnapshotQuantitiesInvoiceActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Invoice**](Invoice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBillingProfile

> BillingProfile UpdateBillingProfile(ctx, id).BillingProfileReplaceUpdateWithWorkflow(billingProfileReplaceUpdateWithWorkflow).Execute()

Update a billing profile



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
	id := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 
	billingProfileReplaceUpdateWithWorkflow := *openapiclient.NewBillingProfileReplaceUpdateWithWorkflow("Name_example", *openapiclient.NewBillingParty(), false, *openapiclient.NewBillingWorkflow()) // BillingProfileReplaceUpdateWithWorkflow | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.UpdateBillingProfile(context.Background(), id).BillingProfileReplaceUpdateWithWorkflow(billingProfileReplaceUpdateWithWorkflow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.UpdateBillingProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBillingProfile`: BillingProfile
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.UpdateBillingProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBillingProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **billingProfileReplaceUpdateWithWorkflow** | [**BillingProfileReplaceUpdateWithWorkflow**](BillingProfileReplaceUpdateWithWorkflow.md) |  | 

### Return type

[**BillingProfile**](BillingProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInvoice

> Invoice UpdateInvoice(ctx, invoiceId).InvoiceReplaceUpdate(invoiceReplaceUpdate).Execute()

Update an invoice



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
	invoiceId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 
	invoiceReplaceUpdate := *openapiclient.NewInvoiceReplaceUpdate(*openapiclient.NewBillingPartyReplaceUpdate(), *openapiclient.NewBillingPartyReplaceUpdate(), []openapiclient.InvoiceLineReplaceUpdate{*openapiclient.NewInvoiceLineReplaceUpdate("Name_example", *openapiclient.NewPeriod(time.Now(), time.Now()), time.Now())}, *openapiclient.NewInvoiceWorkflowReplaceUpdate(*openapiclient.NewInvoiceWorkflowSettingsReplaceUpdate(*openapiclient.NewInvoiceWorkflowInvoicingSettingsReplaceUpdate(), *openapiclient.NewBillingWorkflowPaymentSettings()))) // InvoiceReplaceUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.UpdateInvoice(context.Background(), invoiceId).InvoiceReplaceUpdate(invoiceReplaceUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.UpdateInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInvoice`: Invoice
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.UpdateInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invoiceReplaceUpdate** | [**InvoiceReplaceUpdate**](InvoiceReplaceUpdate.md) |  | 

### Return type

[**Invoice**](Invoice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInvoicePaymentStatus

> Invoice UpdateInvoicePaymentStatus(ctx, invoiceId).CustomInvoicingUpdatePaymentStatusRequest(customInvoicingUpdatePaymentStatusRequest).Execute()

Update invoice payment status



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
	invoiceId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 
	customInvoicingUpdatePaymentStatusRequest := *openapiclient.NewCustomInvoicingUpdatePaymentStatusRequest(openapiclient.CustomInvoicingPaymentTrigger("paid")) // CustomInvoicingUpdatePaymentStatusRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.UpdateInvoicePaymentStatus(context.Background(), invoiceId).CustomInvoicingUpdatePaymentStatusRequest(customInvoicingUpdatePaymentStatusRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.UpdateInvoicePaymentStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInvoicePaymentStatus`: Invoice
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.UpdateInvoicePaymentStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInvoicePaymentStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customInvoicingUpdatePaymentStatusRequest** | [**CustomInvoicingUpdatePaymentStatusRequest**](CustomInvoicingUpdatePaymentStatusRequest.md) |  | 

### Return type

[**Invoice**](Invoice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertBillingProfileCustomerOverride

> BillingProfileCustomerOverrideWithDetails UpsertBillingProfileCustomerOverride(ctx, customerId).BillingProfileCustomerOverrideCreate(billingProfileCustomerOverrideCreate).Execute()

Create a new or update a customer override



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
	customerId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 
	billingProfileCustomerOverrideCreate := *openapiclient.NewBillingProfileCustomerOverrideCreate() // BillingProfileCustomerOverrideCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.UpsertBillingProfileCustomerOverride(context.Background(), customerId).BillingProfileCustomerOverrideCreate(billingProfileCustomerOverrideCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.UpsertBillingProfileCustomerOverride``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpsertBillingProfileCustomerOverride`: BillingProfileCustomerOverrideWithDetails
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.UpsertBillingProfileCustomerOverride`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpsertBillingProfileCustomerOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **billingProfileCustomerOverrideCreate** | [**BillingProfileCustomerOverrideCreate**](BillingProfileCustomerOverrideCreate.md) |  | 

### Return type

[**BillingProfileCustomerOverrideWithDetails**](BillingProfileCustomerOverrideWithDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VoidInvoiceAction

> Invoice VoidInvoiceAction(ctx, invoiceId).VoidInvoiceActionInput(voidInvoiceActionInput).Execute()

Void an invoice



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
	invoiceId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 
	voidInvoiceActionInput := *openapiclient.NewVoidInvoiceActionInput(*openapiclient.NewVoidInvoiceActionCreate(float64(50), openapiclient.VoidInvoiceLineActionCreate{VoidInvoiceLineDiscardAction: openapiclient.NewVoidInvoiceLineDiscardAction("Type_example")}), "Reason_example") // VoidInvoiceActionInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.VoidInvoiceAction(context.Background(), invoiceId).VoidInvoiceActionInput(voidInvoiceActionInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.VoidInvoiceAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VoidInvoiceAction`: Invoice
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.VoidInvoiceAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVoidInvoiceActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **voidInvoiceActionInput** | [**VoidInvoiceActionInput**](VoidInvoiceActionInput.md) |  | 

### Return type

[**Invoice**](Invoice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

