# \CustomersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomer**](CustomersAPI.md#CreateCustomer) | **Post** /api/v1/customers | Create customer
[**CreateCustomerStripePortalSession**](CustomersAPI.md#CreateCustomerStripePortalSession) | **Post** /api/v1/customers/{customerIdOrKey}/stripe/portal | Create Stripe customer portal session
[**DeleteCustomer**](CustomersAPI.md#DeleteCustomer) | **Delete** /api/v1/customers/{customerIdOrKey} | Delete customer
[**DeleteCustomerAppData**](CustomersAPI.md#DeleteCustomerAppData) | **Delete** /api/v1/customers/{customerIdOrKey}/apps/{appId} | Delete customer app data
[**GetCustomer**](CustomersAPI.md#GetCustomer) | **Get** /api/v1/customers/{customerIdOrKey} | Get customer
[**GetCustomerStripeAppData**](CustomersAPI.md#GetCustomerStripeAppData) | **Get** /api/v1/customers/{customerIdOrKey}/stripe | Get customer stripe app data
[**ListCustomerAppData**](CustomersAPI.md#ListCustomerAppData) | **Get** /api/v1/customers/{customerIdOrKey}/apps | List customer app data
[**ListCustomerSubscriptions**](CustomersAPI.md#ListCustomerSubscriptions) | **Get** /api/v1/customers/{customerIdOrKey}/subscriptions | List customer subscriptions
[**ListCustomers**](CustomersAPI.md#ListCustomers) | **Get** /api/v1/customers | List customers
[**UpdateCustomer**](CustomersAPI.md#UpdateCustomer) | **Put** /api/v1/customers/{customerIdOrKey} | Update customer
[**UpsertCustomerAppData**](CustomersAPI.md#UpsertCustomerAppData) | **Put** /api/v1/customers/{customerIdOrKey}/apps | Upsert customer app data
[**UpsertCustomerStripeAppData**](CustomersAPI.md#UpsertCustomerStripeAppData) | **Put** /api/v1/customers/{customerIdOrKey}/stripe | Upsert customer stripe app data



## CreateCustomer

> Customer CreateCustomer(ctx).CustomerCreate(customerCreate).Execute()

Create customer



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
	customerCreate := *openapiclient.NewCustomerCreate("Name_example", *openapiclient.NewCustomerUsageAttribution([]string{"SubjectKeys_example"})) // CustomerCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.CreateCustomer(context.Background()).CustomerCreate(customerCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.CreateCustomer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomer`: Customer
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.CreateCustomer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerCreate** | [**CustomerCreate**](CustomerCreate.md) |  | 

### Return type

[**Customer**](Customer.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCustomerStripePortalSession

> StripeCustomerPortalSession CreateCustomerStripePortalSession(ctx, customerIdOrKey).CreateStripeCustomerPortalSessionParams(createStripeCustomerPortalSessionParams).Execute()

Create Stripe customer portal session



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
	customerIdOrKey := *openapiclient.NewULIDOrExternalKey() // ULIDOrExternalKey | 
	createStripeCustomerPortalSessionParams := *openapiclient.NewCreateStripeCustomerPortalSessionParams() // CreateStripeCustomerPortalSessionParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.CreateCustomerStripePortalSession(context.Background(), customerIdOrKey).CreateStripeCustomerPortalSessionParams(createStripeCustomerPortalSessionParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.CreateCustomerStripePortalSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomerStripePortalSession`: StripeCustomerPortalSession
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.CreateCustomerStripePortalSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerIdOrKey** | [**ULIDOrExternalKey**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerStripePortalSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createStripeCustomerPortalSessionParams** | [**CreateStripeCustomerPortalSessionParams**](CreateStripeCustomerPortalSessionParams.md) |  | 

### Return type

[**StripeCustomerPortalSession**](StripeCustomerPortalSession.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomer

> DeleteCustomer(ctx, customerIdOrKey).Execute()

Delete customer



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
	customerIdOrKey := *openapiclient.NewULIDOrExternalKey() // ULIDOrExternalKey | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomersAPI.DeleteCustomer(context.Background(), customerIdOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.DeleteCustomer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerIdOrKey** | [**ULIDOrExternalKey**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomerRequest struct via the builder pattern


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


## DeleteCustomerAppData

> DeleteCustomerAppData(ctx, customerIdOrKey, appId).Execute()

Delete customer app data



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
	customerIdOrKey := *openapiclient.NewULIDOrExternalKey() // ULIDOrExternalKey | 
	appId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomersAPI.DeleteCustomerAppData(context.Background(), customerIdOrKey, appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.DeleteCustomerAppData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerIdOrKey** | [**ULIDOrExternalKey**](.md) |  | 
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomerAppDataRequest struct via the builder pattern


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


## GetCustomer

> Customer GetCustomer(ctx, customerIdOrKey).Expand(expand).Execute()

Get customer



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
	customerIdOrKey := *openapiclient.NewULIDOrExternalKey() // ULIDOrExternalKey | 
	expand := []openapiclient.CustomerExpand{openapiclient.CustomerExpand("subscriptions")} // []CustomerExpand | What parts of the customer output to expand (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.GetCustomer(context.Background(), customerIdOrKey).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.GetCustomer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomer`: Customer
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.GetCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerIdOrKey** | [**ULIDOrExternalKey**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | [**[]CustomerExpand**](CustomerExpand.md) | What parts of the customer output to expand | 

### Return type

[**Customer**](Customer.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerStripeAppData

> StripeCustomerAppData GetCustomerStripeAppData(ctx, customerIdOrKey).Execute()

Get customer stripe app data



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
	customerIdOrKey := *openapiclient.NewULIDOrExternalKey() // ULIDOrExternalKey | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.GetCustomerStripeAppData(context.Background(), customerIdOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.GetCustomerStripeAppData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomerStripeAppData`: StripeCustomerAppData
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.GetCustomerStripeAppData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerIdOrKey** | [**ULIDOrExternalKey**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerStripeAppDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StripeCustomerAppData**](StripeCustomerAppData.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCustomerAppData

> CustomerAppDataPaginatedResponse ListCustomerAppData(ctx, customerIdOrKey).Page(page).PageSize(pageSize).Type_(type_).Execute()

List customer app data



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
	customerIdOrKey := *openapiclient.NewULIDOrExternalKey() // ULIDOrExternalKey | 
	page := int32(56) // int32 | Page index.  Default is 1. (optional) (default to 1)
	pageSize := int32(56) // int32 | The maximum number of items per page.  Default is 100. (optional) (default to 100)
	type_ := openapiclient.AppType("stripe") // AppType | Filter customer data by app type. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.ListCustomerAppData(context.Background(), customerIdOrKey).Page(page).PageSize(pageSize).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.ListCustomerAppData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCustomerAppData`: CustomerAppDataPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.ListCustomerAppData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerIdOrKey** | [**ULIDOrExternalKey**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCustomerAppDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page index.  Default is 1. | [default to 1]
 **pageSize** | **int32** | The maximum number of items per page.  Default is 100. | [default to 100]
 **type_** | [**AppType**](AppType.md) | Filter customer data by app type. | 

### Return type

[**CustomerAppDataPaginatedResponse**](CustomerAppDataPaginatedResponse.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCustomerSubscriptions

> SubscriptionPaginatedResponse ListCustomerSubscriptions(ctx, customerIdOrKey).Status(status).Order(order).OrderBy(orderBy).Page(page).PageSize(pageSize).Execute()

List customer subscriptions



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
	customerIdOrKey := *openapiclient.NewULIDOrExternalKey() // ULIDOrExternalKey | 
	status := []openapiclient.SubscriptionStatus{openapiclient.SubscriptionStatus("active")} // []SubscriptionStatus |  (optional)
	order := openapiclient.SortOrder("ASC") // SortOrder | The order direction. (optional)
	orderBy := openapiclient.CustomerSubscriptionOrderBy("activeFrom") // CustomerSubscriptionOrderBy | The order by field. (optional)
	page := int32(56) // int32 | Page index.  Default is 1. (optional) (default to 1)
	pageSize := int32(56) // int32 | The maximum number of items per page.  Default is 100. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.ListCustomerSubscriptions(context.Background(), customerIdOrKey).Status(status).Order(order).OrderBy(orderBy).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.ListCustomerSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCustomerSubscriptions`: SubscriptionPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.ListCustomerSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerIdOrKey** | [**ULIDOrExternalKey**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCustomerSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## ListCustomers

> CustomerPaginatedResponse ListCustomers(ctx).Page(page).PageSize(pageSize).Order(order).OrderBy(orderBy).IncludeDeleted(includeDeleted).Key(key).Name(name).PrimaryEmail(primaryEmail).Subject(subject).PlanKey(planKey).Expand(expand).Execute()

List customers



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
	page := int32(56) // int32 | Page index.  Default is 1. (optional) (default to 1)
	pageSize := int32(56) // int32 | The maximum number of items per page.  Default is 100. (optional) (default to 100)
	order := openapiclient.SortOrder("ASC") // SortOrder | The order direction. (optional)
	orderBy := openapiclient.CustomerOrderBy("id") // CustomerOrderBy | The order by field. (optional)
	includeDeleted := true // bool | Include deleted customers. (optional) (default to false)
	key := "key_example" // string | Filter customers by key. Case-insensitive partial match. (optional)
	name := "name_example" // string | Filter customers by name. Case-insensitive partial match. (optional)
	primaryEmail := "primaryEmail_example" // string | Filter customers by primary email. Case-insensitive partial match. (optional)
	subject := "subject_example" // string | Filter customers by usage attribution subject. Case-insensitive partial match. (optional)
	planKey := "planKey_example" // string | Filter customers by the plan key of their susbcription. (optional)
	expand := []openapiclient.CustomerExpand{openapiclient.CustomerExpand("subscriptions")} // []CustomerExpand | What parts of the list output to expand in listings (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.ListCustomers(context.Background()).Page(page).PageSize(pageSize).Order(order).OrderBy(orderBy).IncludeDeleted(includeDeleted).Key(key).Name(name).PrimaryEmail(primaryEmail).Subject(subject).PlanKey(planKey).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.ListCustomers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCustomers`: CustomerPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.ListCustomers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCustomersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page index.  Default is 1. | [default to 1]
 **pageSize** | **int32** | The maximum number of items per page.  Default is 100. | [default to 100]
 **order** | [**SortOrder**](SortOrder.md) | The order direction. | 
 **orderBy** | [**CustomerOrderBy**](CustomerOrderBy.md) | The order by field. | 
 **includeDeleted** | **bool** | Include deleted customers. | [default to false]
 **key** | **string** | Filter customers by key. Case-insensitive partial match. | 
 **name** | **string** | Filter customers by name. Case-insensitive partial match. | 
 **primaryEmail** | **string** | Filter customers by primary email. Case-insensitive partial match. | 
 **subject** | **string** | Filter customers by usage attribution subject. Case-insensitive partial match. | 
 **planKey** | **string** | Filter customers by the plan key of their susbcription. | 
 **expand** | [**[]CustomerExpand**](CustomerExpand.md) | What parts of the list output to expand in listings | 

### Return type

[**CustomerPaginatedResponse**](CustomerPaginatedResponse.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomer

> Customer UpdateCustomer(ctx, customerIdOrKey).CustomerReplaceUpdate(customerReplaceUpdate).Execute()

Update customer



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
	customerIdOrKey := *openapiclient.NewULIDOrExternalKey() // ULIDOrExternalKey | 
	customerReplaceUpdate := *openapiclient.NewCustomerReplaceUpdate("Name_example", *openapiclient.NewCustomerUsageAttribution([]string{"SubjectKeys_example"})) // CustomerReplaceUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.UpdateCustomer(context.Background(), customerIdOrKey).CustomerReplaceUpdate(customerReplaceUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.UpdateCustomer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCustomer`: Customer
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.UpdateCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerIdOrKey** | [**ULIDOrExternalKey**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customerReplaceUpdate** | [**CustomerReplaceUpdate**](CustomerReplaceUpdate.md) |  | 

### Return type

[**Customer**](Customer.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertCustomerAppData

> []CustomerAppData UpsertCustomerAppData(ctx, customerIdOrKey).CustomerAppDataCreateOrUpdateItem(customerAppDataCreateOrUpdateItem).Execute()

Upsert customer app data



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
	customerIdOrKey := *openapiclient.NewULIDOrExternalKey() // ULIDOrExternalKey | 
	customerAppDataCreateOrUpdateItem := []openapiclient.CustomerAppDataCreateOrUpdateItem{openapiclient.CustomerAppDataCreateOrUpdateItem{CustomInvoicingCustomerAppData: openapiclient.NewCustomInvoicingCustomerAppData("Type_example")}} // []CustomerAppDataCreateOrUpdateItem | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.UpsertCustomerAppData(context.Background(), customerIdOrKey).CustomerAppDataCreateOrUpdateItem(customerAppDataCreateOrUpdateItem).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.UpsertCustomerAppData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpsertCustomerAppData`: []CustomerAppData
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.UpsertCustomerAppData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerIdOrKey** | [**ULIDOrExternalKey**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpsertCustomerAppDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customerAppDataCreateOrUpdateItem** | [**[]CustomerAppDataCreateOrUpdateItem**](CustomerAppDataCreateOrUpdateItem.md) |  | 

### Return type

[**[]CustomerAppData**](CustomerAppData.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertCustomerStripeAppData

> StripeCustomerAppData UpsertCustomerStripeAppData(ctx, customerIdOrKey).StripeCustomerAppDataBase(stripeCustomerAppDataBase).Execute()

Upsert customer stripe app data



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
	customerIdOrKey := *openapiclient.NewULIDOrExternalKey() // ULIDOrExternalKey | 
	stripeCustomerAppDataBase := *openapiclient.NewStripeCustomerAppDataBase("StripeCustomerId_example") // StripeCustomerAppDataBase | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.UpsertCustomerStripeAppData(context.Background(), customerIdOrKey).StripeCustomerAppDataBase(stripeCustomerAppDataBase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.UpsertCustomerStripeAppData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpsertCustomerStripeAppData`: StripeCustomerAppData
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.UpsertCustomerStripeAppData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerIdOrKey** | [**ULIDOrExternalKey**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpsertCustomerStripeAppDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stripeCustomerAppDataBase** | [**StripeCustomerAppDataBase**](StripeCustomerAppDataBase.md) |  | 

### Return type

[**StripeCustomerAppData**](StripeCustomerAppData.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

