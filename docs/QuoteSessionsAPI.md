# \QuoteSessionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PatchQuoteSessionV1**](QuoteSessionsAPI.md#PatchQuoteSessionV1) | **Patch** /v1/quote-sessions/{session_id} | Patch Quote Session
[**PostQuoteSessionPriceV1**](QuoteSessionsAPI.md#PostQuoteSessionPriceV1) | **Post** /v1/quote-sessions/{session_id}/price | Post Quote Session Price
[**PostQuoteSessionShadowPriceV1Quote**](QuoteSessionsAPI.md#PostQuoteSessionShadowPriceV1Quote) | **Post** /v1/quote-sessions/{session_id}/shadow-price | Post Quote Session Shadow Price



## PatchQuoteSessionV1

> CreateSessionResponse PatchQuoteSessionV1(ctx, sessionId).PatchSessionRequest(patchSessionRequest).IdempotencyKey(idempotencyKey).Execute()

Patch Quote Session

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
	sessionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	patchSessionRequest := *openapiclient.NewPatchSessionRequest() // PatchSessionRequest | 
	idempotencyKey := "idempotencyKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuoteSessionsAPI.PatchQuoteSessionV1(context.Background(), sessionId).PatchSessionRequest(patchSessionRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuoteSessionsAPI.PatchQuoteSessionV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchQuoteSessionV1`: CreateSessionResponse
	fmt.Fprintf(os.Stdout, "Response from `QuoteSessionsAPI.PatchQuoteSessionV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchQuoteSessionV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchSessionRequest** | [**PatchSessionRequest**](PatchSessionRequest.md) |  | 
 **idempotencyKey** | **string** |  | 

### Return type

[**CreateSessionResponse**](CreateSessionResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostQuoteSessionPriceV1

> PriceSessionResponse PostQuoteSessionPriceV1(ctx, sessionId).IdempotencyKey(idempotencyKey).Execute()

Post Quote Session Price

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
	sessionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	idempotencyKey := "idempotencyKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuoteSessionsAPI.PostQuoteSessionPriceV1(context.Background(), sessionId).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuoteSessionsAPI.PostQuoteSessionPriceV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostQuoteSessionPriceV1`: PriceSessionResponse
	fmt.Fprintf(os.Stdout, "Response from `QuoteSessionsAPI.PostQuoteSessionPriceV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostQuoteSessionPriceV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** |  | 

### Return type

[**PriceSessionResponse**](PriceSessionResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostQuoteSessionShadowPriceV1Quote

> interface{} PostQuoteSessionShadowPriceV1Quote(ctx, sessionId).Execute()

Post Quote Session Shadow Price

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
	sessionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuoteSessionsAPI.PostQuoteSessionShadowPriceV1Quote(context.Background(), sessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuoteSessionsAPI.PostQuoteSessionShadowPriceV1Quote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostQuoteSessionShadowPriceV1Quote`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `QuoteSessionsAPI.PostQuoteSessionShadowPriceV1Quote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostQuoteSessionShadowPriceV1QuoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

