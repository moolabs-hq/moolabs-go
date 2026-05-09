# \AppStripeAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppStripeWebhook**](AppStripeAPI.md#AppStripeWebhook) | **Post** /api/v1/apps/{id}/stripe/webhook | Stripe webhook
[**CreateStripeCheckoutSession**](AppStripeAPI.md#CreateStripeCheckoutSession) | **Post** /api/v1/stripe/checkout/sessions | Create checkout session
[**UpdateStripeAPIKey**](AppStripeAPI.md#UpdateStripeAPIKey) | **Put** /api/v1/apps/{id}/stripe/api-key | Update Stripe API key



## AppStripeWebhook

> StripeWebhookResponse AppStripeWebhook(ctx, id).StripeWebhookEvent(stripeWebhookEvent).Execute()

Stripe webhook



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
	stripeWebhookEvent := *openapiclient.NewStripeWebhookEvent("Id_example", "Type_example", false, int32(123), *openapiclient.NewStripeWebhookEventData(interface{}(123))) // StripeWebhookEvent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppStripeAPI.AppStripeWebhook(context.Background(), id).StripeWebhookEvent(stripeWebhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppStripeAPI.AppStripeWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppStripeWebhook`: StripeWebhookResponse
	fmt.Fprintf(os.Stdout, "Response from `AppStripeAPI.AppStripeWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStripeWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stripeWebhookEvent** | [**StripeWebhookEvent**](StripeWebhookEvent.md) |  | 

### Return type

[**StripeWebhookResponse**](StripeWebhookResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStripeCheckoutSession

> CreateStripeCheckoutSessionResult CreateStripeCheckoutSession(ctx).CreateStripeCheckoutSessionRequest(createStripeCheckoutSessionRequest).Execute()

Create checkout session



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
	createStripeCheckoutSessionRequest := *openapiclient.NewCreateStripeCheckoutSessionRequest(*openapiclient.NewCreateStripeCheckoutSessionRequestCustomer("01G65Z755AFWAKHE12NY0CQ9FH", "Key_example", "Name_example", *openapiclient.NewCustomerUsageAttribution([]string{"SubjectKeys_example"})), *openapiclient.NewCreateStripeCheckoutSessionRequestOptions()) // CreateStripeCheckoutSessionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppStripeAPI.CreateStripeCheckoutSession(context.Background()).CreateStripeCheckoutSessionRequest(createStripeCheckoutSessionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppStripeAPI.CreateStripeCheckoutSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStripeCheckoutSession`: CreateStripeCheckoutSessionResult
	fmt.Fprintf(os.Stdout, "Response from `AppStripeAPI.CreateStripeCheckoutSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStripeCheckoutSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createStripeCheckoutSessionRequest** | [**CreateStripeCheckoutSessionRequest**](CreateStripeCheckoutSessionRequest.md) |  | 

### Return type

[**CreateStripeCheckoutSessionResult**](CreateStripeCheckoutSessionResult.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStripeAPIKey

> UpdateStripeAPIKey(ctx, id).StripeAPIKeyInput(stripeAPIKeyInput).Execute()

Update Stripe API key



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
	stripeAPIKeyInput := *openapiclient.NewStripeAPIKeyInput("SecretAPIKey_example") // StripeAPIKeyInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppStripeAPI.UpdateStripeAPIKey(context.Background(), id).StripeAPIKeyInput(stripeAPIKeyInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppStripeAPI.UpdateStripeAPIKey``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateStripeAPIKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stripeAPIKeyInput** | [**StripeAPIKeyInput**](StripeAPIKeyInput.md) |  | 

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

