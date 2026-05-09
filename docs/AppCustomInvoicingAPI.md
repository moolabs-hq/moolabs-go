# \AppCustomInvoicingAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppCustomInvoicingDraftSynchronized**](AppCustomInvoicingAPI.md#AppCustomInvoicingDraftSynchronized) | **Post** /api/v1/apps/custom-invoicing/{invoiceId}/draft/synchronized | Submit draft synchronization results
[**AppCustomInvoicingIssuingSynchronized**](AppCustomInvoicingAPI.md#AppCustomInvoicingIssuingSynchronized) | **Post** /api/v1/apps/custom-invoicing/{invoiceId}/issuing/synchronized | Submit issuing synchronization results
[**AppCustomInvoicingUpdatePaymentStatus**](AppCustomInvoicingAPI.md#AppCustomInvoicingUpdatePaymentStatus) | **Post** /api/v1/apps/custom-invoicing/{invoiceId}/payment/status | Update payment status



## AppCustomInvoicingDraftSynchronized

> AppCustomInvoicingDraftSynchronized(ctx, invoiceId).CustomInvoicingDraftSynchronizedRequest(customInvoicingDraftSynchronizedRequest).Execute()

Submit draft synchronization results

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
	customInvoicingDraftSynchronizedRequest := *openapiclient.NewCustomInvoicingDraftSynchronizedRequest() // CustomInvoicingDraftSynchronizedRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppCustomInvoicingAPI.AppCustomInvoicingDraftSynchronized(context.Background(), invoiceId).CustomInvoicingDraftSynchronizedRequest(customInvoicingDraftSynchronizedRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppCustomInvoicingAPI.AppCustomInvoicingDraftSynchronized``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAppCustomInvoicingDraftSynchronizedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customInvoicingDraftSynchronizedRequest** | [**CustomInvoicingDraftSynchronizedRequest**](CustomInvoicingDraftSynchronizedRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppCustomInvoicingIssuingSynchronized

> AppCustomInvoicingIssuingSynchronized(ctx, invoiceId).CustomInvoicingFinalizedRequest(customInvoicingFinalizedRequest).Execute()

Submit issuing synchronization results

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
	customInvoicingFinalizedRequest := *openapiclient.NewCustomInvoicingFinalizedRequest() // CustomInvoicingFinalizedRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppCustomInvoicingAPI.AppCustomInvoicingIssuingSynchronized(context.Background(), invoiceId).CustomInvoicingFinalizedRequest(customInvoicingFinalizedRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppCustomInvoicingAPI.AppCustomInvoicingIssuingSynchronized``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAppCustomInvoicingIssuingSynchronizedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customInvoicingFinalizedRequest** | [**CustomInvoicingFinalizedRequest**](CustomInvoicingFinalizedRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppCustomInvoicingUpdatePaymentStatus

> AppCustomInvoicingUpdatePaymentStatus(ctx, invoiceId).CustomInvoicingUpdatePaymentStatusRequest(customInvoicingUpdatePaymentStatusRequest).Execute()

Update payment status

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
	r, err := apiClient.AppCustomInvoicingAPI.AppCustomInvoicingUpdatePaymentStatus(context.Background(), invoiceId).CustomInvoicingUpdatePaymentStatusRequest(customInvoicingUpdatePaymentStatusRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppCustomInvoicingAPI.AppCustomInvoicingUpdatePaymentStatus``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAppCustomInvoicingUpdatePaymentStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customInvoicingUpdatePaymentStatusRequest** | [**CustomInvoicingUpdatePaymentStatusRequest**](CustomInvoicingUpdatePaymentStatusRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

