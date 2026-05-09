# \NotificationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNotificationChannel**](NotificationsAPI.md#CreateNotificationChannel) | **Post** /api/v1/notification/channels | Create a notification channel
[**CreateNotificationRule**](NotificationsAPI.md#CreateNotificationRule) | **Post** /api/v1/notification/rules | Create a notification rule
[**DeleteNotificationChannel**](NotificationsAPI.md#DeleteNotificationChannel) | **Delete** /api/v1/notification/channels/{channelId} | Delete a notification channel
[**DeleteNotificationRule**](NotificationsAPI.md#DeleteNotificationRule) | **Delete** /api/v1/notification/rules/{ruleId} | Delete a notification rule
[**GetNotificationChannel**](NotificationsAPI.md#GetNotificationChannel) | **Get** /api/v1/notification/channels/{channelId} | Get notification channel
[**GetNotificationEvent**](NotificationsAPI.md#GetNotificationEvent) | **Get** /api/v1/notification/events/{eventId} | Get notification event
[**GetNotificationRule**](NotificationsAPI.md#GetNotificationRule) | **Get** /api/v1/notification/rules/{ruleId} | Get notification rule
[**ListNotificationChannels**](NotificationsAPI.md#ListNotificationChannels) | **Get** /api/v1/notification/channels | List notification channels
[**ListNotificationEvents**](NotificationsAPI.md#ListNotificationEvents) | **Get** /api/v1/notification/events | List notification events
[**ListNotificationRules**](NotificationsAPI.md#ListNotificationRules) | **Get** /api/v1/notification/rules | List notification rules
[**ResendNotificationEvent**](NotificationsAPI.md#ResendNotificationEvent) | **Post** /api/v1/notification/events/{eventId}/resend | Re-send notification event
[**TestNotificationRule**](NotificationsAPI.md#TestNotificationRule) | **Post** /api/v1/notification/rules/{ruleId}/test | Test notification rule
[**UpdateNotificationChannel**](NotificationsAPI.md#UpdateNotificationChannel) | **Put** /api/v1/notification/channels/{channelId} | Update a notification channel
[**UpdateNotificationRule**](NotificationsAPI.md#UpdateNotificationRule) | **Put** /api/v1/notification/rules/{ruleId} | Update a notification rule



## CreateNotificationChannel

> NotificationChannelWebhook CreateNotificationChannel(ctx).Body(body).Execute()

Create a notification channel



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
	body := NotificationChannelWebhookCreateRequest(987) // NotificationChannelWebhookCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.CreateNotificationChannel(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.CreateNotificationChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNotificationChannel`: NotificationChannelWebhook
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.CreateNotificationChannel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNotificationChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **NotificationChannelWebhookCreateRequest** |  | 

### Return type

[**NotificationChannelWebhook**](NotificationChannelWebhook.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNotificationRule

> NotificationRule CreateNotificationRule(ctx).NotificationRuleCreateRequest(notificationRuleCreateRequest).Execute()

Create a notification rule



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
	notificationRuleCreateRequest := openapiclient.NotificationRuleCreateRequest{NotificationRuleBalanceThresholdCreateRequest: openapiclient.NewNotificationRuleBalanceThresholdCreateRequest("Type_example", "Balance threshold reached", []openapiclient.NotificationRuleBalanceThresholdValue{*openapiclient.NewNotificationRuleBalanceThresholdValue(float64(100), openapiclient.NotificationRuleBalanceThresholdValueType("PERCENT"))}, []string{"01G65Z755AFWAKHE12NY0CQ9FH"})} // NotificationRuleCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.CreateNotificationRule(context.Background()).NotificationRuleCreateRequest(notificationRuleCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.CreateNotificationRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNotificationRule`: NotificationRule
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.CreateNotificationRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNotificationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notificationRuleCreateRequest** | [**NotificationRuleCreateRequest**](NotificationRuleCreateRequest.md) |  | 

### Return type

[**NotificationRule**](NotificationRule.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNotificationChannel

> DeleteNotificationChannel(ctx, channelId).Execute()

Delete a notification channel



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
	channelId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NotificationsAPI.DeleteNotificationChannel(context.Background(), channelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.DeleteNotificationChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNotificationChannelRequest struct via the builder pattern


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


## DeleteNotificationRule

> DeleteNotificationRule(ctx, ruleId).Execute()

Delete a notification rule



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
	ruleId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NotificationsAPI.DeleteNotificationRule(context.Background(), ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.DeleteNotificationRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNotificationRuleRequest struct via the builder pattern


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


## GetNotificationChannel

> NotificationChannelWebhook GetNotificationChannel(ctx, channelId).Execute()

Get notification channel



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
	channelId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.GetNotificationChannel(context.Background(), channelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetNotificationChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationChannel`: NotificationChannelWebhook
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetNotificationChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationChannelWebhook**](NotificationChannelWebhook.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationEvent

> NotificationEvent GetNotificationEvent(ctx, eventId).Execute()

Get notification event



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
	eventId := "eventId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.GetNotificationEvent(context.Background(), eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetNotificationEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationEvent`: NotificationEvent
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetNotificationEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationEvent**](NotificationEvent.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationRule

> NotificationRule GetNotificationRule(ctx, ruleId).Execute()

Get notification rule



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
	ruleId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.GetNotificationRule(context.Background(), ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetNotificationRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationRule`: NotificationRule
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetNotificationRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationRule**](NotificationRule.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNotificationChannels

> NotificationChannelPaginatedResponse ListNotificationChannels(ctx).IncludeDeleted(includeDeleted).IncludeDisabled(includeDisabled).Page(page).PageSize(pageSize).Order(order).OrderBy(orderBy).Execute()

List notification channels



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
	includeDeleted := true // bool | Include deleted notification channels in response.  Usage: `?includeDeleted=true` (optional) (default to false)
	includeDisabled := true // bool | Include disabled notification channels in response.  Usage: `?includeDisabled=false` (optional) (default to false)
	page := int32(56) // int32 | Page index.  Default is 1. (optional) (default to 1)
	pageSize := int32(56) // int32 | The maximum number of items per page.  Default is 100. (optional) (default to 100)
	order := openapiclient.SortOrder("ASC") // SortOrder | The order direction. (optional)
	orderBy := openapiclient.NotificationChannelOrderBy("id") // NotificationChannelOrderBy | The order by field. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.ListNotificationChannels(context.Background()).IncludeDeleted(includeDeleted).IncludeDisabled(includeDisabled).Page(page).PageSize(pageSize).Order(order).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.ListNotificationChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNotificationChannels`: NotificationChannelPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.ListNotificationChannels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNotificationChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeDeleted** | **bool** | Include deleted notification channels in response.  Usage: &#x60;?includeDeleted&#x3D;true&#x60; | [default to false]
 **includeDisabled** | **bool** | Include disabled notification channels in response.  Usage: &#x60;?includeDisabled&#x3D;false&#x60; | [default to false]
 **page** | **int32** | Page index.  Default is 1. | [default to 1]
 **pageSize** | **int32** | The maximum number of items per page.  Default is 100. | [default to 100]
 **order** | [**SortOrder**](SortOrder.md) | The order direction. | 
 **orderBy** | [**NotificationChannelOrderBy**](NotificationChannelOrderBy.md) | The order by field. | 

### Return type

[**NotificationChannelPaginatedResponse**](NotificationChannelPaginatedResponse.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNotificationEvents

> NotificationEventPaginatedResponse ListNotificationEvents(ctx).From(from).To(to).Feature(feature).Subject(subject).Rule(rule).Channel(channel).Page(page).PageSize(pageSize).Order(order).OrderBy(orderBy).Execute()

List notification events



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
	from := time.Now() // time.Time | Start date-time in RFC 3339 format. Inclusive. (optional)
	to := time.Now() // time.Time | End date-time in RFC 3339 format. Inclusive. (optional)
	feature := []string{"Inner_example"} // []string | Filtering by multiple feature ids or keys.  Usage: `?feature=feature-1&feature=feature-2` (optional)
	subject := []string{"Inner_example"} // []string | Filtering by multiple subject ids or keys.  Usage: `?subject=subject-1&subject=subject-2` (optional)
	rule := []string{"01G65Z755AFWAKHE12NY0CQ9FH"} // []string | Filtering by multiple rule ids.  Usage: `?rule=01J8J2XYZ2N5WBYK09EDZFBSZM&rule=01J8J4R4VZH180KRKQ63NB2VA5` (optional)
	channel := []string{"01G65Z755AFWAKHE12NY0CQ9FH"} // []string | Filtering by multiple channel ids.  Usage: `?channel=01J8J4RXH778XB056JS088PCYT&channel=01J8J4S1R1G9EVN62RG23A9M6J` (optional)
	page := int32(56) // int32 | Page index.  Default is 1. (optional) (default to 1)
	pageSize := int32(56) // int32 | The maximum number of items per page.  Default is 100. (optional) (default to 100)
	order := openapiclient.SortOrder("ASC") // SortOrder | The order direction. (optional)
	orderBy := openapiclient.NotificationEventOrderBy("id") // NotificationEventOrderBy | The order by field. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.ListNotificationEvents(context.Background()).From(from).To(to).Feature(feature).Subject(subject).Rule(rule).Channel(channel).Page(page).PageSize(pageSize).Order(order).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.ListNotificationEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNotificationEvents`: NotificationEventPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.ListNotificationEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNotificationEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **time.Time** | Start date-time in RFC 3339 format. Inclusive. | 
 **to** | **time.Time** | End date-time in RFC 3339 format. Inclusive. | 
 **feature** | **[]string** | Filtering by multiple feature ids or keys.  Usage: &#x60;?feature&#x3D;feature-1&amp;feature&#x3D;feature-2&#x60; | 
 **subject** | **[]string** | Filtering by multiple subject ids or keys.  Usage: &#x60;?subject&#x3D;subject-1&amp;subject&#x3D;subject-2&#x60; | 
 **rule** | **[]string** | Filtering by multiple rule ids.  Usage: &#x60;?rule&#x3D;01J8J2XYZ2N5WBYK09EDZFBSZM&amp;rule&#x3D;01J8J4R4VZH180KRKQ63NB2VA5&#x60; | 
 **channel** | **[]string** | Filtering by multiple channel ids.  Usage: &#x60;?channel&#x3D;01J8J4RXH778XB056JS088PCYT&amp;channel&#x3D;01J8J4S1R1G9EVN62RG23A9M6J&#x60; | 
 **page** | **int32** | Page index.  Default is 1. | [default to 1]
 **pageSize** | **int32** | The maximum number of items per page.  Default is 100. | [default to 100]
 **order** | [**SortOrder**](SortOrder.md) | The order direction. | 
 **orderBy** | [**NotificationEventOrderBy**](NotificationEventOrderBy.md) | The order by field. | 

### Return type

[**NotificationEventPaginatedResponse**](NotificationEventPaginatedResponse.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNotificationRules

> NotificationRulePaginatedResponse ListNotificationRules(ctx).IncludeDeleted(includeDeleted).IncludeDisabled(includeDisabled).Feature(feature).Channel(channel).Page(page).PageSize(pageSize).Order(order).OrderBy(orderBy).Execute()

List notification rules



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
	includeDeleted := true // bool | Include deleted notification rules in response.  Usage: `?includeDeleted=true` (optional) (default to false)
	includeDisabled := true // bool | Include disabled notification rules in response.  Usage: `?includeDisabled=false` (optional) (default to false)
	feature := []string{"Inner_example"} // []string | Filtering by multiple feature ids/keys.  Usage: `?feature=feature-1&feature=feature-2` (optional)
	channel := []string{"Inner_example"} // []string | Filtering by multiple notifiaction channel ids.  Usage: `?channel=01ARZ3NDEKTSV4RRFFQ69G5FAV&channel=01J8J2Y5X4NNGQS32CF81W95E3` (optional)
	page := int32(56) // int32 | Page index.  Default is 1. (optional) (default to 1)
	pageSize := int32(56) // int32 | The maximum number of items per page.  Default is 100. (optional) (default to 100)
	order := openapiclient.SortOrder("ASC") // SortOrder | The order direction. (optional)
	orderBy := openapiclient.NotificationRuleOrderBy("id") // NotificationRuleOrderBy | The order by field. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.ListNotificationRules(context.Background()).IncludeDeleted(includeDeleted).IncludeDisabled(includeDisabled).Feature(feature).Channel(channel).Page(page).PageSize(pageSize).Order(order).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.ListNotificationRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNotificationRules`: NotificationRulePaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.ListNotificationRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNotificationRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeDeleted** | **bool** | Include deleted notification rules in response.  Usage: &#x60;?includeDeleted&#x3D;true&#x60; | [default to false]
 **includeDisabled** | **bool** | Include disabled notification rules in response.  Usage: &#x60;?includeDisabled&#x3D;false&#x60; | [default to false]
 **feature** | **[]string** | Filtering by multiple feature ids/keys.  Usage: &#x60;?feature&#x3D;feature-1&amp;feature&#x3D;feature-2&#x60; | 
 **channel** | **[]string** | Filtering by multiple notifiaction channel ids.  Usage: &#x60;?channel&#x3D;01ARZ3NDEKTSV4RRFFQ69G5FAV&amp;channel&#x3D;01J8J2Y5X4NNGQS32CF81W95E3&#x60; | 
 **page** | **int32** | Page index.  Default is 1. | [default to 1]
 **pageSize** | **int32** | The maximum number of items per page.  Default is 100. | [default to 100]
 **order** | [**SortOrder**](SortOrder.md) | The order direction. | 
 **orderBy** | [**NotificationRuleOrderBy**](NotificationRuleOrderBy.md) | The order by field. | 

### Return type

[**NotificationRulePaginatedResponse**](NotificationRulePaginatedResponse.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResendNotificationEvent

> ResendNotificationEvent(ctx, eventId).NotificationEventResendRequest(notificationEventResendRequest).Execute()

Re-send notification event

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
	eventId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 
	notificationEventResendRequest := *openapiclient.NewNotificationEventResendRequest() // NotificationEventResendRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NotificationsAPI.ResendNotificationEvent(context.Background(), eventId).NotificationEventResendRequest(notificationEventResendRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.ResendNotificationEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResendNotificationEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notificationEventResendRequest** | [**NotificationEventResendRequest**](NotificationEventResendRequest.md) |  | 

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


## TestNotificationRule

> NotificationEvent TestNotificationRule(ctx, ruleId).Execute()

Test notification rule



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
	ruleId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.TestNotificationRule(context.Background(), ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.TestNotificationRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestNotificationRule`: NotificationEvent
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.TestNotificationRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestNotificationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationEvent**](NotificationEvent.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNotificationChannel

> NotificationChannelWebhook UpdateNotificationChannel(ctx, channelId).Body(body).Execute()

Update a notification channel



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
	channelId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 
	body := NotificationChannelWebhookCreateRequest(987) // NotificationChannelWebhookCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.UpdateNotificationChannel(context.Background(), channelId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.UpdateNotificationChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNotificationChannel`: NotificationChannelWebhook
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.UpdateNotificationChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNotificationChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **NotificationChannelWebhookCreateRequest** |  | 

### Return type

[**NotificationChannelWebhook**](NotificationChannelWebhook.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNotificationRule

> NotificationRule UpdateNotificationRule(ctx, ruleId).NotificationRuleCreateRequest(notificationRuleCreateRequest).Execute()

Update a notification rule



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
	ruleId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 
	notificationRuleCreateRequest := openapiclient.NotificationRuleCreateRequest{NotificationRuleBalanceThresholdCreateRequest: openapiclient.NewNotificationRuleBalanceThresholdCreateRequest("Type_example", "Balance threshold reached", []openapiclient.NotificationRuleBalanceThresholdValue{*openapiclient.NewNotificationRuleBalanceThresholdValue(float64(100), openapiclient.NotificationRuleBalanceThresholdValueType("PERCENT"))}, []string{"01G65Z755AFWAKHE12NY0CQ9FH"})} // NotificationRuleCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.UpdateNotificationRule(context.Background(), ruleId).NotificationRuleCreateRequest(notificationRuleCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.UpdateNotificationRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNotificationRule`: NotificationRule
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.UpdateNotificationRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNotificationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notificationRuleCreateRequest** | [**NotificationRuleCreateRequest**](NotificationRuleCreateRequest.md) |  | 

### Return type

[**NotificationRule**](NotificationRule.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

