# \AccountsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccount**](AccountsAPI.md#CreateAccount) | **Post** /v1/arc/accounts | Create Account
[**CreateContact**](AccountsAPI.md#CreateContact) | **Post** /v1/arc/accounts/{account_id}/contacts | Create Contact
[**GetAccount**](AccountsAPI.md#GetAccount) | **Get** /v1/arc/accounts/{account_id} | Get Account
[**GetAccountFilterOptionsV1**](AccountsAPI.md#GetAccountFilterOptionsV1) | **Get** /v1/arc/accounts/filter-options | Get Account Filter Options
[**GetAccountOverview**](AccountsAPI.md#GetAccountOverview) | **Get** /v1/arc/accounts/{account_id}/overview | Get Account Overview
[**GetChannelPreferences**](AccountsAPI.md#GetChannelPreferences) | **Get** /v1/arc/accounts/contacts/{contact_id}/preferences | Get Channel Preferences
[**ListAccountInvoices**](AccountsAPI.md#ListAccountInvoices) | **Get** /v1/arc/accounts/{account_id}/invoices | List Account Invoices
[**ListAccounts**](AccountsAPI.md#ListAccounts) | **Get** /v1/arc/accounts | List Accounts
[**ListContacts**](AccountsAPI.md#ListContacts) | **Get** /v1/arc/accounts/{account_id}/contacts | List Contacts
[**OptOutV1**](AccountsAPI.md#OptOutV1) | **Post** /v1/arc/accounts/contacts/{contact_id}/opt-out | Opt Out
[**ResolveJurisdictionV1**](AccountsAPI.md#ResolveJurisdictionV1) | **Post** /v1/arc/accounts/{account_id}/resolve-jurisdiction | Resolve Jurisdiction
[**UpdateAccount**](AccountsAPI.md#UpdateAccount) | **Patch** /v1/arc/accounts/{account_id} | Update Account
[**UpdateChannelPreferences**](AccountsAPI.md#UpdateChannelPreferences) | **Put** /v1/arc/accounts/contacts/{contact_id}/preferences | Update Channel Preferences
[**UpdateContact**](AccountsAPI.md#UpdateContact) | **Patch** /v1/arc/accounts/contacts/{contact_id} | Update Contact



## CreateAccount

> AccountResponse CreateAccount(ctx).XAPIKey(xAPIKey).AccountCreate(accountCreate).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Create Account



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
	xAPIKey := "xAPIKey_example" // string | 
	accountCreate := *openapiclient.NewAccountCreate("CustomerId_example", "LegalName_example") // AccountCreate | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.CreateAccount(context.Background()).XAPIKey(xAPIKey).AccountCreate(accountCreate).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.CreateAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccount`: AccountResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.CreateAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKey** | **string** |  | 
 **accountCreate** | [**AccountCreate**](AccountCreate.md) |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**AccountResponse**](AccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContact

> ContactResponse CreateContact(ctx, accountId).XAPIKey(xAPIKey).ContactCreate(contactCreate).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Create Contact



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
	accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	contactCreate := *openapiclient.NewContactCreate("FullName_example", "Role_example") // ContactCreate | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.CreateContact(context.Background(), accountId).XAPIKey(xAPIKey).ContactCreate(contactCreate).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.CreateContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContact`: ContactResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.CreateContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **contactCreate** | [**ContactCreate**](ContactCreate.md) |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**ContactResponse**](ContactResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccount

> AccountResponse GetAccount(ctx, accountId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Get Account



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
	accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccount(context.Background(), accountId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccount`: AccountResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**AccountResponse**](AccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountFilterOptionsV1

> AccountFilterOptionsResponse GetAccountFilterOptionsV1(ctx).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Get Account Filter Options



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
	xAPIKey := "xAPIKey_example" // string | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccountFilterOptionsV1(context.Background()).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccountFilterOptionsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountFilterOptionsV1`: AccountFilterOptionsResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccountFilterOptionsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountFilterOptionsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**AccountFilterOptionsResponse**](AccountFilterOptionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountOverview

> AccountOverviewResponse GetAccountOverview(ctx, accountId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Get Account Overview



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
	accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccountOverview(context.Background(), accountId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccountOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountOverview`: AccountOverviewResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccountOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**AccountOverviewResponse**](AccountOverviewResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannelPreferences

> ChannelPreferencesResponse GetChannelPreferences(ctx, contactId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Get Channel Preferences



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
	contactId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetChannelPreferences(context.Background(), contactId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetChannelPreferences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannelPreferences`: ChannelPreferencesResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetChannelPreferences`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelPreferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**ChannelPreferencesResponse**](ChannelPreferencesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccountInvoices

> AccountInvoiceListResponse ListAccountInvoices(ctx, accountId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

List Account Invoices



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
	accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.ListAccountInvoices(context.Background(), accountId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.ListAccountInvoices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAccountInvoices`: AccountInvoiceListResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.ListAccountInvoices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAccountInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**AccountInvoiceListResponse**](AccountInvoiceListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccounts

> AccountListResponse ListAccounts(ctx).XAPIKey(xAPIKey).Page(page).PageSize(pageSize).CustomerId(customerId).Search(search).IsStrategic(isStrategic).SourceSystem(sourceSystem).Subsidiary(subsidiary).HasOpenBalance(hasOpenBalance).SortBy(sortBy).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

List Accounts



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
	xAPIKey := "xAPIKey_example" // string | 
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 50)
	customerId := "customerId_example" // string |  (optional)
	search := "search_example" // string | Search by legal name or customer ID (optional)
	isStrategic := true // bool |  (optional)
	sourceSystem := "sourceSystem_example" // string |  (optional)
	subsidiary := []string{"Inner_example"} // []string | Filter by account subsidiary (optional)
	hasOpenBalance := true // bool | Filter by account open balance (optional)
	sortBy := "sortBy_example" // string | created_desc, created_asc, open_balance_desc, open_balance_asc (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.ListAccounts(context.Background()).XAPIKey(xAPIKey).Page(page).PageSize(pageSize).CustomerId(customerId).Search(search).IsStrategic(isStrategic).SourceSystem(sourceSystem).Subsidiary(subsidiary).HasOpenBalance(hasOpenBalance).SortBy(sortBy).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.ListAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAccounts`: AccountListResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.ListAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKey** | **string** |  | 
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 50]
 **customerId** | **string** |  | 
 **search** | **string** | Search by legal name or customer ID | 
 **isStrategic** | **bool** |  | 
 **sourceSystem** | **string** |  | 
 **subsidiary** | **[]string** | Filter by account subsidiary | 
 **hasOpenBalance** | **bool** | Filter by account open balance | 
 **sortBy** | **string** | created_desc, created_asc, open_balance_desc, open_balance_asc | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**AccountListResponse**](AccountListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContacts

> ContactListResponse ListContacts(ctx, accountId).XAPIKey(xAPIKey).Page(page).PageSize(pageSize).Role(role).IsActive(isActive).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

List Contacts



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
	accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 50)
	role := "role_example" // string | Filter by contact role (optional)
	isActive := true // bool | Filter by active status (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.ListContacts(context.Background(), accountId).XAPIKey(xAPIKey).Page(page).PageSize(pageSize).Role(role).IsActive(isActive).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.ListContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListContacts`: ContactListResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.ListContacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListContactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 50]
 **role** | **string** | Filter by contact role | 
 **isActive** | **bool** | Filter by active status | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**ContactListResponse**](ContactListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptOutV1

> OptOutResponse OptOutV1(ctx, contactId).XAPIKey(xAPIKey).OptOutRequest(optOutRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Opt Out



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
	contactId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	optOutRequest := *openapiclient.NewOptOutRequest("Channel_example") // OptOutRequest | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.OptOutV1(context.Background(), contactId).XAPIKey(xAPIKey).OptOutRequest(optOutRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.OptOutV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptOutV1`: OptOutResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.OptOutV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOptOutV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **optOutRequest** | [**OptOutRequest**](OptOutRequest.md) |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**OptOutResponse**](OptOutResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResolveJurisdictionV1

> JurisdictionProfileResponse ResolveJurisdictionV1(ctx, accountId).XAPIKey(xAPIKey).JurisdictionResolutionRequest(jurisdictionResolutionRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Resolve Jurisdiction



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
	accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	jurisdictionResolutionRequest := *openapiclient.NewJurisdictionResolutionRequest("CountryCode_example") // JurisdictionResolutionRequest | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.ResolveJurisdictionV1(context.Background(), accountId).XAPIKey(xAPIKey).JurisdictionResolutionRequest(jurisdictionResolutionRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.ResolveJurisdictionV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResolveJurisdictionV1`: JurisdictionProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.ResolveJurisdictionV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResolveJurisdictionV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **jurisdictionResolutionRequest** | [**JurisdictionResolutionRequest**](JurisdictionResolutionRequest.md) |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**JurisdictionProfileResponse**](JurisdictionProfileResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccount

> AccountResponse UpdateAccount(ctx, accountId).XAPIKey(xAPIKey).AccountUpdate(accountUpdate).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Update Account



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
	accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	accountUpdate := *openapiclient.NewAccountUpdate() // AccountUpdate | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.UpdateAccount(context.Background(), accountId).XAPIKey(xAPIKey).AccountUpdate(accountUpdate).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.UpdateAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAccount`: AccountResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.UpdateAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **accountUpdate** | [**AccountUpdate**](AccountUpdate.md) |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**AccountResponse**](AccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateChannelPreferences

> ChannelPreferencesResponse UpdateChannelPreferences(ctx, contactId).XAPIKey(xAPIKey).ChannelPreferenceUpdate(channelPreferenceUpdate).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Update Channel Preferences



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
	contactId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	channelPreferenceUpdate := *openapiclient.NewChannelPreferenceUpdate([]openapiclient.ChannelPreferenceCreate{*openapiclient.NewChannelPreferenceCreate("Channel_example")}) // ChannelPreferenceUpdate | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.UpdateChannelPreferences(context.Background(), contactId).XAPIKey(xAPIKey).ChannelPreferenceUpdate(channelPreferenceUpdate).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.UpdateChannelPreferences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateChannelPreferences`: ChannelPreferencesResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.UpdateChannelPreferences`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateChannelPreferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **channelPreferenceUpdate** | [**ChannelPreferenceUpdate**](ChannelPreferenceUpdate.md) |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**ChannelPreferencesResponse**](ChannelPreferencesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContact

> ContactResponse UpdateContact(ctx, contactId).XAPIKey(xAPIKey).ContactUpdate(contactUpdate).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Update Contact



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
	contactId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	contactUpdate := *openapiclient.NewContactUpdate() // ContactUpdate | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.UpdateContact(context.Background(), contactId).XAPIKey(xAPIKey).ContactUpdate(contactUpdate).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.UpdateContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateContact`: ContactResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.UpdateContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **contactUpdate** | [**ContactUpdate**](ContactUpdate.md) |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**ContactResponse**](ContactResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

