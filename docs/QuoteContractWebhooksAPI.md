# \QuoteContractWebhooksAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostQuoteContractInboundEmailV1**](QuoteContractWebhooksAPI.md#PostQuoteContractInboundEmailV1) | **Post** /v1/webhooks/quote-contracts/email | Post Quote Contract Inbound Email



## PostQuoteContractInboundEmailV1

> InboundEmailProcessResponse PostQuoteContractInboundEmailV1(ctx).InboundEmailPayload(inboundEmailPayload).XMoolabsSignature(xMoolabsSignature).Execute()

Post Quote Contract Inbound Email

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
	inboundEmailPayload := *openapiclient.NewInboundEmailPayload("Recipient_example", "FromEmail_example", "MessageId_example") // InboundEmailPayload | 
	xMoolabsSignature := "xMoolabsSignature_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuoteContractWebhooksAPI.PostQuoteContractInboundEmailV1(context.Background()).InboundEmailPayload(inboundEmailPayload).XMoolabsSignature(xMoolabsSignature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuoteContractWebhooksAPI.PostQuoteContractInboundEmailV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostQuoteContractInboundEmailV1`: InboundEmailProcessResponse
	fmt.Fprintf(os.Stdout, "Response from `QuoteContractWebhooksAPI.PostQuoteContractInboundEmailV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostQuoteContractInboundEmailV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inboundEmailPayload** | [**InboundEmailPayload**](InboundEmailPayload.md) |  | 
 **xMoolabsSignature** | **string** |  | 

### Return type

[**InboundEmailProcessResponse**](InboundEmailProcessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

