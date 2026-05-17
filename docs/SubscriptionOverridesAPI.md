# \SubscriptionOverridesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCommercialOverridesV1**](SubscriptionOverridesAPI.md#CreateCommercialOverridesV1) | **Post** /v1/subscription-overrides | Create Commercial Overrides
[**GetCommercialOverridesV1**](SubscriptionOverridesAPI.md#GetCommercialOverridesV1) | **Get** /v1/subscription-overrides/{subscription_id} | Get Commercial Overrides



## CreateCommercialOverridesV1

> interface{} CreateCommercialOverridesV1(ctx).CommercialOverridesInput(commercialOverridesInput).Execute()

Create Commercial Overrides



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
	commercialOverridesInput := *openapiclient.NewCommercialOverridesInput("SubscriptionId_example", "CustomerId_example", "EffectiveFrom_example") // CommercialOverridesInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionOverridesAPI.CreateCommercialOverridesV1(context.Background()).CommercialOverridesInput(commercialOverridesInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionOverridesAPI.CreateCommercialOverridesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCommercialOverridesV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionOverridesAPI.CreateCommercialOverridesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCommercialOverridesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commercialOverridesInput** | [**CommercialOverridesInput**](CommercialOverridesInput.md) |  | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommercialOverridesV1

> interface{} GetCommercialOverridesV1(ctx, subscriptionId).Execute()

Get Commercial Overrides



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
	subscriptionId := "subscriptionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionOverridesAPI.GetCommercialOverridesV1(context.Background(), subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionOverridesAPI.GetCommercialOverridesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCommercialOverridesV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionOverridesAPI.GetCommercialOverridesV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommercialOverridesV1Request struct via the builder pattern


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

