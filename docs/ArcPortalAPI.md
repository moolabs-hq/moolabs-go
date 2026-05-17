# \ArcPortalAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PortalCreateDispute**](ArcPortalAPI.md#PortalCreateDispute) | **Post** /v1/arc/portal/me/disputes | Portal Create Dispute
[**PortalCreatePtpV1Arc**](ArcPortalAPI.md#PortalCreatePtpV1Arc) | **Post** /v1/arc/portal/me/promises-to-pay | Portal Create Ptp
[**PortalMakePayment**](ArcPortalAPI.md#PortalMakePayment) | **Post** /v1/arc/portal/me/payments | Portal Make Payment
[**PortalMyCases**](ArcPortalAPI.md#PortalMyCases) | **Get** /v1/arc/portal/me/cases | Portal My Cases
[**PortalMyInvoices**](ArcPortalAPI.md#PortalMyInvoices) | **Get** /v1/arc/portal/me/invoices | Portal My Invoices
[**PortalUpdatePreferences**](ArcPortalAPI.md#PortalUpdatePreferences) | **Put** /v1/arc/portal/me/preferences | Portal Update Preferences



## PortalCreateDispute

> interface{} PortalCreateDispute(ctx).Authorization(authorization).PortalDisputeRequest(portalDisputeRequest).Execute()

Portal Create Dispute



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
	authorization := "authorization_example" // string | 
	portalDisputeRequest := *openapiclient.NewPortalDisputeRequest("InvoiceId_example", int32(123), "Description_example") // PortalDisputeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArcPortalAPI.PortalCreateDispute(context.Background()).Authorization(authorization).PortalDisputeRequest(portalDisputeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArcPortalAPI.PortalCreateDispute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PortalCreateDispute`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArcPortalAPI.PortalCreateDispute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPortalCreateDisputeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **portalDisputeRequest** | [**PortalDisputeRequest**](PortalDisputeRequest.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PortalCreatePtpV1Arc

> interface{} PortalCreatePtpV1Arc(ctx).Authorization(authorization).PortalPTPRequest(portalPTPRequest).Execute()

Portal Create Ptp



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
	authorization := "authorization_example" // string | 
	portalPTPRequest := *openapiclient.NewPortalPTPRequest("CaseId_example", int32(123), time.Now()) // PortalPTPRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArcPortalAPI.PortalCreatePtpV1Arc(context.Background()).Authorization(authorization).PortalPTPRequest(portalPTPRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArcPortalAPI.PortalCreatePtpV1Arc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PortalCreatePtpV1Arc`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArcPortalAPI.PortalCreatePtpV1Arc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPortalCreatePtpV1ArcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **portalPTPRequest** | [**PortalPTPRequest**](PortalPTPRequest.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PortalMakePayment

> interface{} PortalMakePayment(ctx).Authorization(authorization).PortalPaymentRequest(portalPaymentRequest).Execute()

Portal Make Payment



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
	authorization := "authorization_example" // string | 
	portalPaymentRequest := *openapiclient.NewPortalPaymentRequest("CaseId_example", int32(123)) // PortalPaymentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArcPortalAPI.PortalMakePayment(context.Background()).Authorization(authorization).PortalPaymentRequest(portalPaymentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArcPortalAPI.PortalMakePayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PortalMakePayment`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArcPortalAPI.PortalMakePayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPortalMakePaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **portalPaymentRequest** | [**PortalPaymentRequest**](PortalPaymentRequest.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PortalMyCases

> interface{} PortalMyCases(ctx).Authorization(authorization).Execute()

Portal My Cases



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
	authorization := "authorization_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArcPortalAPI.PortalMyCases(context.Background()).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArcPortalAPI.PortalMyCases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PortalMyCases`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArcPortalAPI.PortalMyCases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPortalMyCasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

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


## PortalMyInvoices

> interface{} PortalMyInvoices(ctx).Authorization(authorization).Execute()

Portal My Invoices



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
	authorization := "authorization_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArcPortalAPI.PortalMyInvoices(context.Background()).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArcPortalAPI.PortalMyInvoices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PortalMyInvoices`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArcPortalAPI.PortalMyInvoices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPortalMyInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

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


## PortalUpdatePreferences

> interface{} PortalUpdatePreferences(ctx).Authorization(authorization).PortalPreferencesRequest(portalPreferencesRequest).Execute()

Portal Update Preferences



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
	authorization := "authorization_example" // string | 
	portalPreferencesRequest := *openapiclient.NewPortalPreferencesRequest(map[string]interface{}{"key": interface{}(123)}) // PortalPreferencesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArcPortalAPI.PortalUpdatePreferences(context.Background()).Authorization(authorization).PortalPreferencesRequest(portalPreferencesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArcPortalAPI.PortalUpdatePreferences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PortalUpdatePreferences`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArcPortalAPI.PortalUpdatePreferences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPortalUpdatePreferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **portalPreferencesRequest** | [**PortalPreferencesRequest**](PortalPreferencesRequest.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

