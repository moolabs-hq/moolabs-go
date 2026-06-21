# \BuyerQuotesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptBuyerQuote**](BuyerQuotesAPI.md#AcceptBuyerQuote) | **Post** /v1/buyer/quotes/accept | Accept Buyer Quote
[**GetBuyerContractPdfV1**](BuyerQuotesAPI.md#GetBuyerContractPdfV1) | **Get** /v1/buyer/quotes/contract.pdf | Get Buyer Contract Pdf
[**GetBuyerQuote**](BuyerQuotesAPI.md#GetBuyerQuote) | **Get** /v1/buyer/quotes | Get Buyer Quote
[**RejectBuyerQuote**](BuyerQuotesAPI.md#RejectBuyerQuote) | **Post** /v1/buyer/quotes/reject | Reject Buyer Quote
[**RequestBuyerQuoteChangesV1**](BuyerQuotesAPI.md#RequestBuyerQuoteChangesV1) | **Post** /v1/buyer/quotes/request-changes | Request Buyer Quote Changes
[**RequestOtp**](BuyerQuotesAPI.md#RequestOtp) | **Post** /v1/buyer/quotes/otp | Request Otp
[**VerifyOtp**](BuyerQuotesAPI.md#VerifyOtp) | **Post** /v1/buyer/quotes/otp/verify | Verify Otp



## AcceptBuyerQuote

> map[string]interface{} AcceptBuyerQuote(ctx).BuyerAcceptRequest(buyerAcceptRequest).IdempotencyKey(idempotencyKey).XQuoteBuyerToken(xQuoteBuyerToken).Execute()

Accept Buyer Quote

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
	buyerAcceptRequest := *openapiclient.NewBuyerAcceptRequest("Code_example") // BuyerAcceptRequest | 
	idempotencyKey := "idempotencyKey_example" // string |  (optional)
	xQuoteBuyerToken := "xQuoteBuyerToken_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuyerQuotesAPI.AcceptBuyerQuote(context.Background()).BuyerAcceptRequest(buyerAcceptRequest).IdempotencyKey(idempotencyKey).XQuoteBuyerToken(xQuoteBuyerToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuyerQuotesAPI.AcceptBuyerQuote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AcceptBuyerQuote`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `BuyerQuotesAPI.AcceptBuyerQuote`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcceptBuyerQuoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **buyerAcceptRequest** | [**BuyerAcceptRequest**](BuyerAcceptRequest.md) |  | 
 **idempotencyKey** | **string** |  | 
 **xQuoteBuyerToken** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBuyerContractPdfV1

> interface{} GetBuyerContractPdfV1(ctx).XQuoteBuyerToken(xQuoteBuyerToken).Execute()

Get Buyer Contract Pdf

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
	xQuoteBuyerToken := "xQuoteBuyerToken_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuyerQuotesAPI.GetBuyerContractPdfV1(context.Background()).XQuoteBuyerToken(xQuoteBuyerToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuyerQuotesAPI.GetBuyerContractPdfV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBuyerContractPdfV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BuyerQuotesAPI.GetBuyerContractPdfV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerContractPdfV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xQuoteBuyerToken** | **string** |  | 

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


## GetBuyerQuote

> BuyerQuoteProjection GetBuyerQuote(ctx).XQuoteBuyerToken(xQuoteBuyerToken).Execute()

Get Buyer Quote

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
	xQuoteBuyerToken := "xQuoteBuyerToken_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuyerQuotesAPI.GetBuyerQuote(context.Background()).XQuoteBuyerToken(xQuoteBuyerToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuyerQuotesAPI.GetBuyerQuote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBuyerQuote`: BuyerQuoteProjection
	fmt.Fprintf(os.Stdout, "Response from `BuyerQuotesAPI.GetBuyerQuote`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerQuoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xQuoteBuyerToken** | **string** |  | 

### Return type

[**BuyerQuoteProjection**](BuyerQuoteProjection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RejectBuyerQuote

> map[string]interface{} RejectBuyerQuote(ctx).BuyerRejectRequest(buyerRejectRequest).XQuoteBuyerToken(xQuoteBuyerToken).Execute()

Reject Buyer Quote

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
	buyerRejectRequest := *openapiclient.NewBuyerRejectRequest() // BuyerRejectRequest | 
	xQuoteBuyerToken := "xQuoteBuyerToken_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuyerQuotesAPI.RejectBuyerQuote(context.Background()).BuyerRejectRequest(buyerRejectRequest).XQuoteBuyerToken(xQuoteBuyerToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuyerQuotesAPI.RejectBuyerQuote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RejectBuyerQuote`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `BuyerQuotesAPI.RejectBuyerQuote`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRejectBuyerQuoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **buyerRejectRequest** | [**BuyerRejectRequest**](BuyerRejectRequest.md) |  | 
 **xQuoteBuyerToken** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestBuyerQuoteChangesV1

> BuyerRequestChangesResponse RequestBuyerQuoteChangesV1(ctx).BuyerRequestChangesRequest(buyerRequestChangesRequest).IdempotencyKey(idempotencyKey).XQuoteBuyerToken(xQuoteBuyerToken).Execute()

Request Buyer Quote Changes

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
	buyerRequestChangesRequest := *openapiclient.NewBuyerRequestChangesRequest("Note_example") // BuyerRequestChangesRequest | 
	idempotencyKey := "idempotencyKey_example" // string |  (optional)
	xQuoteBuyerToken := "xQuoteBuyerToken_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuyerQuotesAPI.RequestBuyerQuoteChangesV1(context.Background()).BuyerRequestChangesRequest(buyerRequestChangesRequest).IdempotencyKey(idempotencyKey).XQuoteBuyerToken(xQuoteBuyerToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuyerQuotesAPI.RequestBuyerQuoteChangesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestBuyerQuoteChangesV1`: BuyerRequestChangesResponse
	fmt.Fprintf(os.Stdout, "Response from `BuyerQuotesAPI.RequestBuyerQuoteChangesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestBuyerQuoteChangesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **buyerRequestChangesRequest** | [**BuyerRequestChangesRequest**](BuyerRequestChangesRequest.md) |  | 
 **idempotencyKey** | **string** |  | 
 **xQuoteBuyerToken** | **string** |  | 

### Return type

[**BuyerRequestChangesResponse**](BuyerRequestChangesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestOtp

> BuyerOtpRequestResponse RequestOtp(ctx).XQuoteBuyerToken(xQuoteBuyerToken).Execute()

Request Otp

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
	xQuoteBuyerToken := "xQuoteBuyerToken_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuyerQuotesAPI.RequestOtp(context.Background()).XQuoteBuyerToken(xQuoteBuyerToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuyerQuotesAPI.RequestOtp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestOtp`: BuyerOtpRequestResponse
	fmt.Fprintf(os.Stdout, "Response from `BuyerQuotesAPI.RequestOtp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestOtpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xQuoteBuyerToken** | **string** |  | 

### Return type

[**BuyerOtpRequestResponse**](BuyerOtpRequestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyOtp

> BuyerOtpVerifyResponse VerifyOtp(ctx).BuyerOtpVerifyRequest(buyerOtpVerifyRequest).XQuoteBuyerToken(xQuoteBuyerToken).Execute()

Verify Otp

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
	buyerOtpVerifyRequest := *openapiclient.NewBuyerOtpVerifyRequest("Code_example") // BuyerOtpVerifyRequest | 
	xQuoteBuyerToken := "xQuoteBuyerToken_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuyerQuotesAPI.VerifyOtp(context.Background()).BuyerOtpVerifyRequest(buyerOtpVerifyRequest).XQuoteBuyerToken(xQuoteBuyerToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuyerQuotesAPI.VerifyOtp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyOtp`: BuyerOtpVerifyResponse
	fmt.Fprintf(os.Stdout, "Response from `BuyerQuotesAPI.VerifyOtp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyOtpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **buyerOtpVerifyRequest** | [**BuyerOtpVerifyRequest**](BuyerOtpVerifyRequest.md) |  | 
 **xQuoteBuyerToken** | **string** |  | 

### Return type

[**BuyerOtpVerifyResponse**](BuyerOtpVerifyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

