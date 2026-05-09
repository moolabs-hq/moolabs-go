# \MeterSubscriptionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelSubscription**](MeterSubscriptionsAPI.md#CancelSubscription) | **Post** /api/v1/subscriptions/{subscriptionId}/cancel | Cancel subscription
[**ChangeSubscription**](MeterSubscriptionsAPI.md#ChangeSubscription) | **Post** /api/v1/subscriptions/{subscriptionId}/change | Change subscription
[**CreateSubscription**](MeterSubscriptionsAPI.md#CreateSubscription) | **Post** /api/v1/subscriptions | Create subscription
[**CreateSubscriptionAddon**](MeterSubscriptionsAPI.md#CreateSubscriptionAddon) | **Post** /api/v1/subscriptions/{subscriptionId}/addons | Create subscription addon
[**DeleteSubscription**](MeterSubscriptionsAPI.md#DeleteSubscription) | **Delete** /api/v1/subscriptions/{subscriptionId} | Delete subscription
[**EditSubscription**](MeterSubscriptionsAPI.md#EditSubscription) | **Patch** /api/v1/subscriptions/{subscriptionId} | Edit subscription
[**GetSubscriptionAddon**](MeterSubscriptionsAPI.md#GetSubscriptionAddon) | **Get** /api/v1/subscriptions/{subscriptionId}/addons/{subscriptionAddonId} | Get subscription addon
[**GetSubscriptionGet**](MeterSubscriptionsAPI.md#GetSubscriptionGet) | **Get** /api/v1/subscriptions/{subscriptionId} | Get subscription
[**ListSubscriptionAddons**](MeterSubscriptionsAPI.md#ListSubscriptionAddons) | **Get** /api/v1/subscriptions/{subscriptionId}/addons | List subscription addons
[**ListSubscriptions**](MeterSubscriptionsAPI.md#ListSubscriptions) | **Get** /api/v1/subscriptions | List subscriptions
[**MigrateSubscription**](MeterSubscriptionsAPI.md#MigrateSubscription) | **Post** /api/v1/subscriptions/{subscriptionId}/migrate | Migrate subscription
[**RestoreSubscription**](MeterSubscriptionsAPI.md#RestoreSubscription) | **Post** /api/v1/subscriptions/{subscriptionId}/restore | Restore subscription
[**UnscheduleCancelation**](MeterSubscriptionsAPI.md#UnscheduleCancelation) | **Post** /api/v1/subscriptions/{subscriptionId}/unschedule-cancelation | Unschedule cancelation
[**UpdateSubscriptionAddon**](MeterSubscriptionsAPI.md#UpdateSubscriptionAddon) | **Patch** /api/v1/subscriptions/{subscriptionId}/addons/{subscriptionAddonId} | Update subscription addon



## CancelSubscription

> Subscription CancelSubscription(ctx, subscriptionId).CancelSubscriptionRequest(cancelSubscriptionRequest).Execute()

Cancel subscription



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
	subscriptionId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 
	cancelSubscriptionRequest := *openapiclient.NewCancelSubscriptionRequest() // CancelSubscriptionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeterSubscriptionsAPI.CancelSubscription(context.Background(), subscriptionId).CancelSubscriptionRequest(cancelSubscriptionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeterSubscriptionsAPI.CancelSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelSubscription`: Subscription
	fmt.Fprintf(os.Stdout, "Response from `MeterSubscriptionsAPI.CancelSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cancelSubscriptionRequest** | [**CancelSubscriptionRequest**](CancelSubscriptionRequest.md) |  | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangeSubscription

> SubscriptionChangeResponseBody ChangeSubscription(ctx, subscriptionId).SubscriptionChange(subscriptionChange).Execute()

Change subscription



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
	subscriptionId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 
	subscriptionChange := openapiclient.SubscriptionChange{CustomSubscriptionChange: openapiclient.NewCustomSubscriptionChange(openapiclient.SubscriptionTiming{SubscriptionTimingEnum: openapiclient.SubscriptionTimingEnum("immediate")}, *openapiclient.NewCustomPlanInput("Name_example", "USD", "P1M", []openapiclient.PlanPhase{*openapiclient.NewPlanPhase("Key_example", "Name_example", "P1Y", []openapiclient.RateCard{openapiclient.RateCard{RateCardFlatFee: openapiclient.NewRateCardFlatFee("Type_example", "Key_example", "Name_example", "BillingCadence_example", *openapiclient.NewFlatPriceWithPaymentTerm("Type_example", "Amount_example"))}})}))} // SubscriptionChange | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeterSubscriptionsAPI.ChangeSubscription(context.Background(), subscriptionId).SubscriptionChange(subscriptionChange).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeterSubscriptionsAPI.ChangeSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeSubscription`: SubscriptionChangeResponseBody
	fmt.Fprintf(os.Stdout, "Response from `MeterSubscriptionsAPI.ChangeSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionChange** | [**SubscriptionChange**](SubscriptionChange.md) |  | 

### Return type

[**SubscriptionChangeResponseBody**](SubscriptionChangeResponseBody.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSubscription

> Subscription CreateSubscription(ctx).SubscriptionCreate(subscriptionCreate).Execute()

Create subscription

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
	subscriptionCreate := openapiclient.SubscriptionCreate{CustomSubscriptionCreate: openapiclient.NewCustomSubscriptionCreate(*openapiclient.NewCustomPlanInput("Name_example", "USD", "P1M", []openapiclient.PlanPhase{*openapiclient.NewPlanPhase("Key_example", "Name_example", "P1Y", []openapiclient.RateCard{openapiclient.RateCard{RateCardFlatFee: openapiclient.NewRateCardFlatFee("Type_example", "Key_example", "Name_example", "BillingCadence_example", *openapiclient.NewFlatPriceWithPaymentTerm("Type_example", "Amount_example"))}})}))} // SubscriptionCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeterSubscriptionsAPI.CreateSubscription(context.Background()).SubscriptionCreate(subscriptionCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeterSubscriptionsAPI.CreateSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSubscription`: Subscription
	fmt.Fprintf(os.Stdout, "Response from `MeterSubscriptionsAPI.CreateSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionCreate** | [**SubscriptionCreate**](SubscriptionCreate.md) |  | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSubscriptionAddon

> SubscriptionAddon CreateSubscriptionAddon(ctx, subscriptionId).SubscriptionAddonCreate(subscriptionAddonCreate).Execute()

Create subscription addon



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
	subscriptionId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 
	subscriptionAddonCreate := *openapiclient.NewSubscriptionAddonCreate("Name_example", int32(1), openapiclient.SubscriptionTiming{SubscriptionTimingEnum: openapiclient.SubscriptionTimingEnum("immediate")}, *openapiclient.NewAddon2("01G65Z755AFWAKHE12NY0CQ9FH")) // SubscriptionAddonCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeterSubscriptionsAPI.CreateSubscriptionAddon(context.Background(), subscriptionId).SubscriptionAddonCreate(subscriptionAddonCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeterSubscriptionsAPI.CreateSubscriptionAddon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSubscriptionAddon`: SubscriptionAddon
	fmt.Fprintf(os.Stdout, "Response from `MeterSubscriptionsAPI.CreateSubscriptionAddon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubscriptionAddonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionAddonCreate** | [**SubscriptionAddonCreate**](SubscriptionAddonCreate.md) |  | 

### Return type

[**SubscriptionAddon**](SubscriptionAddon.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSubscription

> DeleteSubscription(ctx, subscriptionId).Execute()

Delete subscription



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
	subscriptionId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MeterSubscriptionsAPI.DeleteSubscription(context.Background(), subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeterSubscriptionsAPI.DeleteSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditSubscription

> Subscription EditSubscription(ctx, subscriptionId).SubscriptionEdit(subscriptionEdit).Execute()

Edit subscription



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
	subscriptionId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 
	subscriptionEdit := *openapiclient.NewSubscriptionEdit([]openapiclient.SubscriptionEditOperation{openapiclient.SubscriptionEditOperation{EditSubscriptionAddItem: openapiclient.NewEditSubscriptionAddItem("Op_example", "PhaseKey_example", openapiclient.RateCard{RateCardFlatFee: openapiclient.NewRateCardFlatFee("Type_example", "Key_example", "Name_example", "BillingCadence_example", *openapiclient.NewFlatPriceWithPaymentTerm("Type_example", "Amount_example"))})}}) // SubscriptionEdit | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeterSubscriptionsAPI.EditSubscription(context.Background(), subscriptionId).SubscriptionEdit(subscriptionEdit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeterSubscriptionsAPI.EditSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditSubscription`: Subscription
	fmt.Fprintf(os.Stdout, "Response from `MeterSubscriptionsAPI.EditSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionEdit** | [**SubscriptionEdit**](SubscriptionEdit.md) |  | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscriptionAddon

> SubscriptionAddon GetSubscriptionAddon(ctx, subscriptionId, subscriptionAddonId).Execute()

Get subscription addon



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
	subscriptionId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 
	subscriptionAddonId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeterSubscriptionsAPI.GetSubscriptionAddon(context.Background(), subscriptionId, subscriptionAddonId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeterSubscriptionsAPI.GetSubscriptionAddon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubscriptionAddon`: SubscriptionAddon
	fmt.Fprintf(os.Stdout, "Response from `MeterSubscriptionsAPI.GetSubscriptionAddon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** |  | 
**subscriptionAddonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionAddonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SubscriptionAddon**](SubscriptionAddon.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscriptionGet

> SubscriptionExpanded GetSubscriptionGet(ctx, subscriptionId).At(at).Execute()

Get subscription

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
	subscriptionId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 
	at := time.Now() // time.Time | The time at which the subscription should be queried. If not provided the current time is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeterSubscriptionsAPI.GetSubscriptionGet(context.Background(), subscriptionId).At(at).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeterSubscriptionsAPI.GetSubscriptionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubscriptionGet`: SubscriptionExpanded
	fmt.Fprintf(os.Stdout, "Response from `MeterSubscriptionsAPI.GetSubscriptionGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **at** | **time.Time** | The time at which the subscription should be queried. If not provided the current time is used. | 

### Return type

[**SubscriptionExpanded**](SubscriptionExpanded.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubscriptionAddons

> []SubscriptionAddon ListSubscriptionAddons(ctx, subscriptionId).Execute()

List subscription addons



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
	subscriptionId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeterSubscriptionsAPI.ListSubscriptionAddons(context.Background(), subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeterSubscriptionsAPI.ListSubscriptionAddons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSubscriptionAddons`: []SubscriptionAddon
	fmt.Fprintf(os.Stdout, "Response from `MeterSubscriptionsAPI.ListSubscriptionAddons`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSubscriptionAddonsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]SubscriptionAddon**](SubscriptionAddon.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubscriptions

> SubscriptionPaginatedResponse ListSubscriptions(ctx).CustomerId(customerId).Status(status).Order(order).OrderBy(orderBy).Page(page).PageSize(pageSize).Execute()

List subscriptions



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
	customerId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | Filter by customer ID. (optional)
	status := []openapiclient.SubscriptionStatus{openapiclient.SubscriptionStatus("active")} // []SubscriptionStatus |  (optional)
	order := openapiclient.SortOrder("ASC") // SortOrder | The order direction. (optional)
	orderBy := openapiclient.CustomerSubscriptionOrderBy("activeFrom") // CustomerSubscriptionOrderBy | The order by field. (optional)
	page := int32(56) // int32 | Page index.  Default is 1. (optional) (default to 1)
	pageSize := int32(56) // int32 | The maximum number of items per page.  Default is 100. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeterSubscriptionsAPI.ListSubscriptions(context.Background()).CustomerId(customerId).Status(status).Order(order).OrderBy(orderBy).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeterSubscriptionsAPI.ListSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSubscriptions`: SubscriptionPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `MeterSubscriptionsAPI.ListSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerId** | **string** | Filter by customer ID. | 
 **status** | [**[]SubscriptionStatus**](SubscriptionStatus.md) |  | 
 **order** | [**SortOrder**](SortOrder.md) | The order direction. | 
 **orderBy** | [**CustomerSubscriptionOrderBy**](CustomerSubscriptionOrderBy.md) | The order by field. | 
 **page** | **int32** | Page index.  Default is 1. | [default to 1]
 **pageSize** | **int32** | The maximum number of items per page.  Default is 100. | [default to 100]

### Return type

[**SubscriptionPaginatedResponse**](SubscriptionPaginatedResponse.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MigrateSubscription

> SubscriptionChangeResponseBody MigrateSubscription(ctx, subscriptionId).MigrateSubscriptionRequest(migrateSubscriptionRequest).Execute()

Migrate subscription



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
	subscriptionId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 
	migrateSubscriptionRequest := *openapiclient.NewMigrateSubscriptionRequest() // MigrateSubscriptionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeterSubscriptionsAPI.MigrateSubscription(context.Background(), subscriptionId).MigrateSubscriptionRequest(migrateSubscriptionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeterSubscriptionsAPI.MigrateSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MigrateSubscription`: SubscriptionChangeResponseBody
	fmt.Fprintf(os.Stdout, "Response from `MeterSubscriptionsAPI.MigrateSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMigrateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **migrateSubscriptionRequest** | [**MigrateSubscriptionRequest**](MigrateSubscriptionRequest.md) |  | 

### Return type

[**SubscriptionChangeResponseBody**](SubscriptionChangeResponseBody.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreSubscription

> Subscription RestoreSubscription(ctx, subscriptionId).Execute()

Restore subscription



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
	subscriptionId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeterSubscriptionsAPI.RestoreSubscription(context.Background(), subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeterSubscriptionsAPI.RestoreSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreSubscription`: Subscription
	fmt.Fprintf(os.Stdout, "Response from `MeterSubscriptionsAPI.RestoreSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Subscription**](Subscription.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnscheduleCancelation

> Subscription UnscheduleCancelation(ctx, subscriptionId).Execute()

Unschedule cancelation



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
	subscriptionId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeterSubscriptionsAPI.UnscheduleCancelation(context.Background(), subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeterSubscriptionsAPI.UnscheduleCancelation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnscheduleCancelation`: Subscription
	fmt.Fprintf(os.Stdout, "Response from `MeterSubscriptionsAPI.UnscheduleCancelation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnscheduleCancelationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Subscription**](Subscription.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSubscriptionAddon

> SubscriptionAddon UpdateSubscriptionAddon(ctx, subscriptionId, subscriptionAddonId).SubscriptionAddonUpdate(subscriptionAddonUpdate).Execute()

Update subscription addon



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
	subscriptionId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 
	subscriptionAddonId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 
	subscriptionAddonUpdate := *openapiclient.NewSubscriptionAddonUpdate() // SubscriptionAddonUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeterSubscriptionsAPI.UpdateSubscriptionAddon(context.Background(), subscriptionId, subscriptionAddonId).SubscriptionAddonUpdate(subscriptionAddonUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeterSubscriptionsAPI.UpdateSubscriptionAddon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSubscriptionAddon`: SubscriptionAddon
	fmt.Fprintf(os.Stdout, "Response from `MeterSubscriptionsAPI.UpdateSubscriptionAddon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** |  | 
**subscriptionAddonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSubscriptionAddonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **subscriptionAddonUpdate** | [**SubscriptionAddonUpdate**](SubscriptionAddonUpdate.md) |  | 

### Return type

[**SubscriptionAddon**](SubscriptionAddon.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

