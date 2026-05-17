# \ArcWebhooksAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MoometerWebhook**](ArcWebhooksAPI.md#MoometerWebhook) | **Post** /v1/arc/webhooks/moometer | Moometer Webhook
[**ResendDeliveryWebhook**](ArcWebhooksAPI.md#ResendDeliveryWebhook) | **Post** /v1/arc/webhooks/resend | Resend Delivery Webhook
[**ResendInboundWebhook**](ArcWebhooksAPI.md#ResendInboundWebhook) | **Post** /v1/arc/webhooks/inbound/resend | Resend Inbound Webhook
[**SendgridWebhook**](ArcWebhooksAPI.md#SendgridWebhook) | **Post** /v1/arc/webhooks/sendgrid | Sendgrid Webhook
[**StripeWebhookPost**](ArcWebhooksAPI.md#StripeWebhookPost) | **Post** /v1/arc/webhooks/stripe | Stripe Webhook
[**TwilioWebhook**](ArcWebhooksAPI.md#TwilioWebhook) | **Post** /v1/arc/webhooks/twilio | Twilio Webhook



## MoometerWebhook

> interface{} MoometerWebhook(ctx).XMooMeterSignature(xMooMeterSignature).Execute()

Moometer Webhook



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
	xMooMeterSignature := "xMooMeterSignature_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArcWebhooksAPI.MoometerWebhook(context.Background()).XMooMeterSignature(xMooMeterSignature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArcWebhooksAPI.MoometerWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MoometerWebhook`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArcWebhooksAPI.MoometerWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMoometerWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xMooMeterSignature** | **string** |  | 

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


## ResendDeliveryWebhook

> interface{} ResendDeliveryWebhook(ctx).XResendSignature(xResendSignature).Execute()

Resend Delivery Webhook



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
	xResendSignature := "xResendSignature_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArcWebhooksAPI.ResendDeliveryWebhook(context.Background()).XResendSignature(xResendSignature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArcWebhooksAPI.ResendDeliveryWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResendDeliveryWebhook`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArcWebhooksAPI.ResendDeliveryWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResendDeliveryWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xResendSignature** | **string** |  | 

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


## ResendInboundWebhook

> interface{} ResendInboundWebhook(ctx).Execute()

Resend Inbound Webhook



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArcWebhooksAPI.ResendInboundWebhook(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArcWebhooksAPI.ResendInboundWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResendInboundWebhook`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArcWebhooksAPI.ResendInboundWebhook`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiResendInboundWebhookRequest struct via the builder pattern


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


## SendgridWebhook

> interface{} SendgridWebhook(ctx).XTwilioEmailEventWebhookSignature(xTwilioEmailEventWebhookSignature).Execute()

Sendgrid Webhook



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
	xTwilioEmailEventWebhookSignature := "xTwilioEmailEventWebhookSignature_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArcWebhooksAPI.SendgridWebhook(context.Background()).XTwilioEmailEventWebhookSignature(xTwilioEmailEventWebhookSignature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArcWebhooksAPI.SendgridWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendgridWebhook`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArcWebhooksAPI.SendgridWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendgridWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTwilioEmailEventWebhookSignature** | **string** |  | 

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


## StripeWebhookPost

> interface{} StripeWebhookPost(ctx).StripeSignature(stripeSignature).Execute()

Stripe Webhook



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
	stripeSignature := "stripeSignature_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArcWebhooksAPI.StripeWebhookPost(context.Background()).StripeSignature(stripeSignature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArcWebhooksAPI.StripeWebhookPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StripeWebhookPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArcWebhooksAPI.StripeWebhookPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStripeWebhookPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stripeSignature** | **string** |  | 

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


## TwilioWebhook

> interface{} TwilioWebhook(ctx).XTwilioSignature(xTwilioSignature).Execute()

Twilio Webhook



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
	xTwilioSignature := "xTwilioSignature_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArcWebhooksAPI.TwilioWebhook(context.Background()).XTwilioSignature(xTwilioSignature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArcWebhooksAPI.TwilioWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TwilioWebhook`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArcWebhooksAPI.TwilioWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTwilioWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTwilioSignature** | **string** |  | 

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

