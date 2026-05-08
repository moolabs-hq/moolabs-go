# \RateCardsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRateCard**](RateCardsAPI.md#CreateRateCard) | **Post** /v1/rate_cards | Create Rate Card
[**GetRateCard**](RateCardsAPI.md#GetRateCard) | **Get** /v1/rate_cards/{rate_card_id} | Get Rate Card



## CreateRateCard

> RateCardResponse CreateRateCard(ctx).CreateRateCardRequest(createRateCardRequest).Execute()

Create Rate Card



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
	createRateCardRequest := *openapiclient.NewCreateRateCardRequest("TenantId_example", "FeatureKey_example", "Version_example", time.Now(), map[string]interface{}{"key": interface{}(123)}) // CreateRateCardRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RateCardsAPI.CreateRateCard(context.Background()).CreateRateCardRequest(createRateCardRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RateCardsAPI.CreateRateCard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRateCard`: RateCardResponse
	fmt.Fprintf(os.Stdout, "Response from `RateCardsAPI.CreateRateCard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRateCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRateCardRequest** | [**CreateRateCardRequest**](CreateRateCardRequest.md) |  | 

### Return type

[**RateCardResponse**](RateCardResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRateCard

> RateCardResponse GetRateCard(ctx, rateCardId).Execute()

Get Rate Card



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
	rateCardId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RateCardsAPI.GetRateCard(context.Background(), rateCardId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RateCardsAPI.GetRateCard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRateCard`: RateCardResponse
	fmt.Fprintf(os.Stdout, "Response from `RateCardsAPI.GetRateCard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rateCardId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRateCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RateCardResponse**](RateCardResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

