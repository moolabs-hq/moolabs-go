# \MeterBillingAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdvanceInvoiceAction**](MeterBillingAPI.md#AdvanceInvoiceAction) | **Post** /api/v1/billing/invoices/{invoiceId}/advance | Advance the invoice&#39;s state to the next status
[**ApproveInvoiceAction**](MeterBillingAPI.md#ApproveInvoiceAction) | **Post** /api/v1/billing/invoices/{invoiceId}/approve | Send the invoice to the customer
[**CreateBillingProfile**](MeterBillingAPI.md#CreateBillingProfile) | **Post** /api/v1/billing/profiles | Create a new billing profile
[**CreatePendingInvoiceLine**](MeterBillingAPI.md#CreatePendingInvoiceLine) | **Post** /api/v1/billing/customers/{customerId}/invoices/pending-lines | Create pending line items
[**DeleteBillingProfile**](MeterBillingAPI.md#DeleteBillingProfile) | **Delete** /api/v1/billing/profiles/{id} | Delete a billing profile
[**DeleteBillingProfileCustomerOverride**](MeterBillingAPI.md#DeleteBillingProfileCustomerOverride) | **Delete** /api/v1/billing/customers/{customerId} | Delete a customer override
[**DeleteInvoice**](MeterBillingAPI.md#DeleteInvoice) | **Delete** /api/v1/billing/invoices/{invoiceId} | Delete an invoice
[**GetBillingProfile**](MeterBillingAPI.md#GetBillingProfile) | **Get** /api/v1/billing/profiles/{id} | Get a billing profile
[**GetBillingProfileCustomerOverride**](MeterBillingAPI.md#GetBillingProfileCustomerOverride) | **Get** /api/v1/billing/customers/{customerId} | Get a customer override
[**GetInvoice**](MeterBillingAPI.md#GetInvoice) | **Get** /api/v1/billing/invoices/{invoiceId} | Get an invoice
[**InvoicePendingLinesAction**](MeterBillingAPI.md#InvoicePendingLinesAction) | **Post** /api/v1/billing/invoices/invoice | Invoice a customer based on the pending line items
[**ListBillingProfileCustomerOverrides**](MeterBillingAPI.md#ListBillingProfileCustomerOverrides) | **Get** /api/v1/billing/customers | List customer overrides
[**ListBillingProfiles**](MeterBillingAPI.md#ListBillingProfiles) | **Get** /api/v1/billing/profiles | List billing profiles
[**ListInvoices**](MeterBillingAPI.md#ListInvoices) | **Get** /api/v1/billing/invoices | List invoices
[**RecalculateInvoiceTaxAction**](MeterBillingAPI.md#RecalculateInvoiceTaxAction) | **Post** /api/v1/billing/invoices/{invoiceId}/taxes/recalculate | Recalculate an invoice&#39;s tax amounts
[**RetryInvoiceAction**](MeterBillingAPI.md#RetryInvoiceAction) | **Post** /api/v1/billing/invoices/{invoiceId}/retry | Retry advancing the invoice after a failed attempt.
[**SimulateInvoice**](MeterBillingAPI.md#SimulateInvoice) | **Post** /api/v1/billing/customers/{customerId}/invoices/simulate | Simulate an invoice for a customer
[**SnapshotQuantitiesInvoiceAction**](MeterBillingAPI.md#SnapshotQuantitiesInvoiceAction) | **Post** /api/v1/billing/invoices/{invoiceId}/snapshot-quantities | Snapshot quantities for usage based line items
[**UpdateBillingProfile**](MeterBillingAPI.md#UpdateBillingProfile) | **Put** /api/v1/billing/profiles/{id} | Update a billing profile
[**UpdateInvoice**](MeterBillingAPI.md#UpdateInvoice) | **Put** /api/v1/billing/invoices/{invoiceId} | Update an invoice
[**UpdateInvoicePaymentStatus**](MeterBillingAPI.md#UpdateInvoicePaymentStatus) | **Post** /api/v1/billing/invoices/{invoiceId}/payment-status | Update invoice payment status
[**UpsertBillingProfileCustomerOverride**](MeterBillingAPI.md#UpsertBillingProfileCustomerOverride) | **Put** /api/v1/billing/customers/{customerId} | Create a new or update a customer override
[**VoidInvoiceAction**](MeterBillingAPI.md#VoidInvoiceAction) | **Post** /api/v1/billing/invoices/{invoiceId}/void | Void an invoice



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
	resp, r, err := apiClient.MeterBillingAPI.AdvanceInvoiceAction(context.Background(), invoiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeterBillingAPI.AdvanceInvoiceAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdvanceInvoiceAction`: Invoice
	fmt.Fprintf(os.Stdout, "Response from `MeterBillingAPI.AdvanceInvoiceAction`: %v\n", resp)
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

[CloudTokenAuth](../README.md#CloudTokenAuth)

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
	resp, r, err := apiClient.MeterBillingAPI.ApproveInvoiceAction(context.Background(), invoiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeterBillingAPI.ApproveInvoiceAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApproveInvoiceAction`: Invoice
	fmt.Fprintf(os.Stdout, "Response from `MeterBillingAPI.ApproveInvoiceAction`: %v\n", resp)
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

[CloudTokenAuth](../README.md#CloudTokenAuth)

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
	resp, r, err := apiClient.MeterBillingAPI.CreateBillingProfile(context.Background()).BillingProfileCreate(billingProfileCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeterBillingAPI.CreateBillingProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBillingProfile`: BillingProfile
	fmt.Fprintf(os.Stdout, "Response from `MeterBillingAPI.CreateBillingProfile`: %v\n", resp)
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

[CloudTokenAuth](../README.md#CloudTokenAuth)

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
	resp, r, err := apiClient.MeterBillingAPI.CreatePendingInvoiceLine(context.Background(), customerId).InvoicePendingLineCreateInput(invoicePendingLineCreateInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeterBillingAPI.CreatePendingInvoiceLine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePendingInvoiceLine`: InvoicePendingLineCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `MeterBillingAPI.CreatePendingInvoiceLine`: %v\n", resp)
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

[CloudTokenAuth](../README.md#CloudTokenAuth)

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
	r, err := apiClient.MeterBillingAPI.DeleteBillingProfile(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeterBillingAPI.DeleteBillingProfile``: %v\n", err)
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

[CloudTokenAuth](../README.md#CloudTokenAuth)

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
	r, err := apiClient.MeterBillingAPI.DeleteBillingProfileCustomerOverride(context.Background(), customerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeterBillingAPI.DeleteBillingProfileCustomerOverride``: %v\n", err)
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

[CloudTokenAuth](../README.md#CloudTokenAuth)

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
	r, err := apiClient.MeterBillingAPI.DeleteInvoice(context.Background(), invoiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeterBillingAPI.DeleteInvoice``: %v\n", err)
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

[CloudTokenAuth](../README.md#CloudTokenAuth)

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
	resp, r, err := apiClient.MeterBillingAPI.GetBillingProfile(context.Background(), id).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeterBillingAPI.GetBillingProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillingProfile`: BillingProfile
	fmt.Fprintf(os.Stdout, "Response from `MeterBillingAPI.GetBillingProfile`: %v\n", resp)
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

[CloudTokenAuth](../README.md#CloudTokenAuth)

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
	resp, r, err := apiClient.MeterBillingAPI.GetBillingProfileCustomerOverride(context.Background(), customerId).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeterBillingAPI.GetBillingProfileCustomerOverride``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillingProfileCustomerOverride`: BillingProfileCustomerOverrideWithDetails
	fmt.Fprintf(os.Stdout, "Response from `MeterBillingAPI.GetBillingProfileCustomerOverride`: %v\n", resp)
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

[CloudTokenAuth](../README.md#CloudTokenAuth)

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
	resp, r, err := apiClient.MeterBillingAPI.GetInvoice(context.Background(), invoiceId).Expand(expand).IncludeDeletedLines(includeDeletedLines).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeterBillingAPI.GetInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvoice`: Invoice
	fmt.Fprintf(os.Stdout, "Response from `MeterBillingAPI.GetInvoice`: %v\n", resp)
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

[CloudTokenAuth](../README.md#CloudTokenAuth)

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
	resp, r, err := apiClient.MeterBillingAPI.InvoicePendingLinesAction(context.Background()).InvoicePendingLinesActionInput(invoicePendingLinesActionInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeterBillingAPI.InvoicePendingLinesAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoicePendingLinesAction`: []Invoice
	fmt.Fprintf(os.Stdout, "Response from `MeterBillingAPI.InvoicePendingLinesAction`: %v\n", resp)
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

[CloudTokenAuth](../README.md#CloudTokenAuth)

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
	resp, r, err := apiClient.MeterBillingAPI.ListBillingProfileCustomerOverrides(context.Background()).BillingProfile(billingProfile).CustomersWithoutPinnedProfile(customersWithoutPinnedProfile).IncludeAllCustomers(includeAllCustomers).CustomerId(customerId).CustomerName(customerName).CustomerKey(customerKey).CustomerPrimaryEmail(customerPrimaryEmail).Expand(expand).Order(order).OrderBy(orderBy).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeterBillingAPI.ListBillingProfileCustomerOverrides``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBillingProfileCustomerOverrides`: BillingProfileCustomerOverrideWithDetailsPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `MeterBillingAPI.ListBillingProfileCustomerOverrides`: %v\n", resp)
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

[CloudTokenAuth](../README.md#CloudTokenAuth)

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
	resp, r, err := apiClient.MeterBillingAPI.ListBillingProfiles(context.Background()).IncludeArchived(includeArchived).Expand(expand).Page(page).PageSize(pageSize).Order(order).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeterBillingAPI.ListBillingProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBillingProfiles`: BillingProfilePaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `MeterBillingAPI.ListBillingProfiles`: %v\n", resp)
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

[CloudTokenAuth](../README.md#CloudTokenAuth)

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
	resp, r, err := apiClient.MeterBillingAPI.ListInvoices(context.Background()).Statuses(statuses).ExtendedStatuses(extendedStatuses).IssuedAfter(issuedAfter).IssuedBefore(issuedBefore).PeriodStartAfter(periodStartAfter).PeriodStartBefore(periodStartBefore).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Expand(expand).Customers(customers).IncludeDeleted(includeDeleted).Page(page).PageSize(pageSize).Order(order).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeterBillingAPI.ListInvoices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInvoices`: InvoicePaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `MeterBillingAPI.ListInvoices`: %v\n", resp)
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

[CloudTokenAuth](../README.md#CloudTokenAuth)

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
	resp, r, err := apiClient.MeterBillingAPI.RecalculateInvoiceTaxAction(context.Background(), invoiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeterBillingAPI.RecalculateInvoiceTaxAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecalculateInvoiceTaxAction`: Invoice
	fmt.Fprintf(os.Stdout, "Response from `MeterBillingAPI.RecalculateInvoiceTaxAction`: %v\n", resp)
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

[CloudTokenAuth](../README.md#CloudTokenAuth)

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
	resp, r, err := apiClient.MeterBillingAPI.RetryInvoiceAction(context.Background(), invoiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeterBillingAPI.RetryInvoiceAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetryInvoiceAction`: Invoice
	fmt.Fprintf(os.Stdout, "Response from `MeterBillingAPI.RetryInvoiceAction`: %v\n", resp)
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

[CloudTokenAuth](../README.md#CloudTokenAuth)

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
	resp, r, err := apiClient.MeterBillingAPI.SimulateInvoice(context.Background(), customerId).InvoiceSimulationInput(invoiceSimulationInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeterBillingAPI.SimulateInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SimulateInvoice`: Invoice
	fmt.Fprintf(os.Stdout, "Response from `MeterBillingAPI.SimulateInvoice`: %v\n", resp)
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

[CloudTokenAuth](../README.md#CloudTokenAuth)

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
	resp, r, err := apiClient.MeterBillingAPI.SnapshotQuantitiesInvoiceAction(context.Background(), invoiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeterBillingAPI.SnapshotQuantitiesInvoiceAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SnapshotQuantitiesInvoiceAction`: Invoice
	fmt.Fprintf(os.Stdout, "Response from `MeterBillingAPI.SnapshotQuantitiesInvoiceAction`: %v\n", resp)
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

[CloudTokenAuth](../README.md#CloudTokenAuth)

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
	resp, r, err := apiClient.MeterBillingAPI.UpdateBillingProfile(context.Background(), id).BillingProfileReplaceUpdateWithWorkflow(billingProfileReplaceUpdateWithWorkflow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeterBillingAPI.UpdateBillingProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBillingProfile`: BillingProfile
	fmt.Fprintf(os.Stdout, "Response from `MeterBillingAPI.UpdateBillingProfile`: %v\n", resp)
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

[CloudTokenAuth](../README.md#CloudTokenAuth)

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
	resp, r, err := apiClient.MeterBillingAPI.UpdateInvoice(context.Background(), invoiceId).InvoiceReplaceUpdate(invoiceReplaceUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeterBillingAPI.UpdateInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInvoice`: Invoice
	fmt.Fprintf(os.Stdout, "Response from `MeterBillingAPI.UpdateInvoice`: %v\n", resp)
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

[CloudTokenAuth](../README.md#CloudTokenAuth)

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
	resp, r, err := apiClient.MeterBillingAPI.UpdateInvoicePaymentStatus(context.Background(), invoiceId).CustomInvoicingUpdatePaymentStatusRequest(customInvoicingUpdatePaymentStatusRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeterBillingAPI.UpdateInvoicePaymentStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInvoicePaymentStatus`: Invoice
	fmt.Fprintf(os.Stdout, "Response from `MeterBillingAPI.UpdateInvoicePaymentStatus`: %v\n", resp)
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

[CloudTokenAuth](../README.md#CloudTokenAuth)

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
	resp, r, err := apiClient.MeterBillingAPI.UpsertBillingProfileCustomerOverride(context.Background(), customerId).BillingProfileCustomerOverrideCreate(billingProfileCustomerOverrideCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeterBillingAPI.UpsertBillingProfileCustomerOverride``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpsertBillingProfileCustomerOverride`: BillingProfileCustomerOverrideWithDetails
	fmt.Fprintf(os.Stdout, "Response from `MeterBillingAPI.UpsertBillingProfileCustomerOverride`: %v\n", resp)
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

[CloudTokenAuth](../README.md#CloudTokenAuth)

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
	resp, r, err := apiClient.MeterBillingAPI.VoidInvoiceAction(context.Background(), invoiceId).VoidInvoiceActionInput(voidInvoiceActionInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeterBillingAPI.VoidInvoiceAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VoidInvoiceAction`: Invoice
	fmt.Fprintf(os.Stdout, "Response from `MeterBillingAPI.VoidInvoiceAction`: %v\n", resp)
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

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

