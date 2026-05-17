# \TopupAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePaymentLinkV1**](TopupAPI.md#CreatePaymentLinkV1) | **Post** /v1/topup/payment-link | Create Payment Link
[**TopupCheckout**](TopupAPI.md#TopupCheckout) | **Get** /v1/topup/checkout | Topup Checkout



## CreatePaymentLinkV1

> PaymentLinkResponse CreatePaymentLinkV1(ctx).PaymentLinkRequest(paymentLinkRequest).Execute()

Create Payment Link



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
	paymentLinkRequest := *openapiclient.NewPaymentLinkRequest("Flow_example", "TenantId_example", "CaseId_example", "PtpId_example", "InvoiceId_example", float32(123)) // PaymentLinkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TopupAPI.CreatePaymentLinkV1(context.Background()).PaymentLinkRequest(paymentLinkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TopupAPI.CreatePaymentLinkV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePaymentLinkV1`: PaymentLinkResponse
	fmt.Fprintf(os.Stdout, "Response from `TopupAPI.CreatePaymentLinkV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaymentLinkV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentLinkRequest** | [**PaymentLinkRequest**](PaymentLinkRequest.md) |  | 

### Return type

[**PaymentLinkResponse**](PaymentLinkResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TopupCheckout

> interface{} TopupCheckout(ctx).WalletId(walletId).TenantId(tenantId).PoolId(poolId).AmountUsd(amountUsd).Execute()

Topup Checkout



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Wallet to top up
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant ID (optional; looked up from wallet if omitted) (optional)
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Pool ID (optional; looked up from wallet if omitted) (optional)
	amountUsd := float32(8.14) // float32 | Amount in USD (optional) (default to 50.0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TopupAPI.TopupCheckout(context.Background()).WalletId(walletId).TenantId(tenantId).PoolId(poolId).AmountUsd(amountUsd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TopupAPI.TopupCheckout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TopupCheckout`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TopupAPI.TopupCheckout`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTopupCheckoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **walletId** | **string** | Wallet to top up | 
 **tenantId** | **string** | Tenant ID (optional; looked up from wallet if omitted) | 
 **poolId** | **string** | Pool ID (optional; looked up from wallet if omitted) | 
 **amountUsd** | **float32** | Amount in USD | [default to 50.0]

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

