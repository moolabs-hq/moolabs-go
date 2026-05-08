# \AppsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApp**](AppsAPI.md#GetApp) | **Get** /api/v1/apps/{id} | Get app
[**GetMarketplaceListing**](AppsAPI.md#GetMarketplaceListing) | **Get** /api/v1/marketplace/listings/{type} | Get app details by type
[**ListApps**](AppsAPI.md#ListApps) | **Get** /api/v1/apps | List apps
[**ListMarketplaceListings**](AppsAPI.md#ListMarketplaceListings) | **Get** /api/v1/marketplace/listings | List available apps
[**MarketplaceAppAPIKeyInstall**](AppsAPI.md#MarketplaceAppAPIKeyInstall) | **Post** /api/v1/marketplace/listings/{type}/install/apikey | Install app via API key
[**MarketplaceAppInstall**](AppsAPI.md#MarketplaceAppInstall) | **Post** /api/v1/marketplace/listings/{type}/install | Install app
[**MarketplaceOAuth2InstallAuthorize**](AppsAPI.md#MarketplaceOAuth2InstallAuthorize) | **Get** /api/v1/marketplace/listings/{type}/install/oauth2/authorize | Install app via OAuth2
[**MarketplaceOAuth2InstallGetURL**](AppsAPI.md#MarketplaceOAuth2InstallGetURL) | **Get** /api/v1/marketplace/listings/{type}/install/oauth2 | Get OAuth2 install URL
[**UninstallApp**](AppsAPI.md#UninstallApp) | **Delete** /api/v1/apps/{id} | Uninstall app
[**UpdateApp**](AppsAPI.md#UpdateApp) | **Put** /api/v1/apps/{id} | Update app



## GetApp

> App GetApp(ctx, id).Execute()

Get app



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.GetApp(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.GetApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApp`: App
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.GetApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**App**](App.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketplaceListing

> MarketplaceListing GetMarketplaceListing(ctx, type_).Execute()

Get app details by type



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
	type_ := openapiclient.AppType("stripe") // AppType | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.GetMarketplaceListing(context.Background(), type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.GetMarketplaceListing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketplaceListing`: MarketplaceListing
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.GetMarketplaceListing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | [**AppType**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketplaceListingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MarketplaceListing**](MarketplaceListing.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApps

> AppPaginatedResponse ListApps(ctx).Page(page).PageSize(pageSize).Execute()

List apps



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.ListApps(context.Background()).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.ListApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApps`: AppPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.ListApps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page index.  Default is 1. | [default to 1]
 **pageSize** | **int32** | The maximum number of items per page.  Default is 100. | [default to 100]

### Return type

[**AppPaginatedResponse**](AppPaginatedResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMarketplaceListings

> MarketplaceListingPaginatedResponse ListMarketplaceListings(ctx).Page(page).PageSize(pageSize).Execute()

List available apps



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.ListMarketplaceListings(context.Background()).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.ListMarketplaceListings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMarketplaceListings`: MarketplaceListingPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.ListMarketplaceListings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMarketplaceListingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page index.  Default is 1. | [default to 1]
 **pageSize** | **int32** | The maximum number of items per page.  Default is 100. | [default to 100]

### Return type

[**MarketplaceListingPaginatedResponse**](MarketplaceListingPaginatedResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketplaceAppAPIKeyInstall

> MarketplaceInstallResponse MarketplaceAppAPIKeyInstall(ctx, type_).MarketplaceAppAPIKeyInstallRequest(marketplaceAppAPIKeyInstallRequest).Execute()

Install app via API key



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
	type_ := openapiclient.AppType("stripe") // AppType | The type of the app to install.
	marketplaceAppAPIKeyInstallRequest := *openapiclient.NewMarketplaceAppAPIKeyInstallRequest("ApiKey_example") // MarketplaceAppAPIKeyInstallRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.MarketplaceAppAPIKeyInstall(context.Background(), type_).MarketplaceAppAPIKeyInstallRequest(marketplaceAppAPIKeyInstallRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.MarketplaceAppAPIKeyInstall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceAppAPIKeyInstall`: MarketplaceInstallResponse
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.MarketplaceAppAPIKeyInstall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | [**AppType**](.md) | The type of the app to install. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceAppAPIKeyInstallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **marketplaceAppAPIKeyInstallRequest** | [**MarketplaceAppAPIKeyInstallRequest**](MarketplaceAppAPIKeyInstallRequest.md) |  | 

### Return type

[**MarketplaceInstallResponse**](MarketplaceInstallResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketplaceAppInstall

> MarketplaceInstallResponse MarketplaceAppInstall(ctx, type_).MarketplaceInstallRequestPayload(marketplaceInstallRequestPayload).Execute()

Install app



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
	type_ := openapiclient.AppType("stripe") // AppType | The type of the app to install.
	marketplaceInstallRequestPayload := *openapiclient.NewMarketplaceInstallRequestPayload() // MarketplaceInstallRequestPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.MarketplaceAppInstall(context.Background(), type_).MarketplaceInstallRequestPayload(marketplaceInstallRequestPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.MarketplaceAppInstall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceAppInstall`: MarketplaceInstallResponse
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.MarketplaceAppInstall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | [**AppType**](.md) | The type of the app to install. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceAppInstallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **marketplaceInstallRequestPayload** | [**MarketplaceInstallRequestPayload**](MarketplaceInstallRequestPayload.md) |  | 

### Return type

[**MarketplaceInstallResponse**](MarketplaceInstallResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketplaceOAuth2InstallAuthorize

> UnexpectedProblemResponse MarketplaceOAuth2InstallAuthorize(ctx, type_).State(state).Code(code).Error_(error_).ErrorDescription(errorDescription).ErrorUri(errorUri).Execute()

Install app via OAuth2



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
	type_ := openapiclient.AppType("stripe") // AppType | The type of the app to install.
	state := "state_example" // string | Required if the \"state\" parameter was present in the client authorization request. The exact value received from the client:  Unique, randomly generated, opaque, and non-guessable string that is sent when starting an authentication request and validated when processing the response. (optional)
	code := "code_example" // string | Authorization code which the client will later exchange for an access token. Required with the success response. (optional)
	error_ := openapiclient.OAuth2AuthorizationCodeGrantErrorType("invalid_request") // OAuth2AuthorizationCodeGrantErrorType | Error code. Required with the error response. (optional)
	errorDescription := "errorDescription_example" // string | Optional human-readable text providing additional information, used to assist the client developer in understanding the error that occurred. (optional)
	errorUri := "errorUri_example" // string | Optional uri identifying a human-readable web page with information about the error, used to provide the client developer with additional information about the error (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.MarketplaceOAuth2InstallAuthorize(context.Background(), type_).State(state).Code(code).Error_(error_).ErrorDescription(errorDescription).ErrorUri(errorUri).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.MarketplaceOAuth2InstallAuthorize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceOAuth2InstallAuthorize`: UnexpectedProblemResponse
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.MarketplaceOAuth2InstallAuthorize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | [**AppType**](.md) | The type of the app to install. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceOAuth2InstallAuthorizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **state** | **string** | Required if the \&quot;state\&quot; parameter was present in the client authorization request. The exact value received from the client:  Unique, randomly generated, opaque, and non-guessable string that is sent when starting an authentication request and validated when processing the response. | 
 **code** | **string** | Authorization code which the client will later exchange for an access token. Required with the success response. | 
 **error_** | [**OAuth2AuthorizationCodeGrantErrorType**](OAuth2AuthorizationCodeGrantErrorType.md) | Error code. Required with the error response. | 
 **errorDescription** | **string** | Optional human-readable text providing additional information, used to assist the client developer in understanding the error that occurred. | 
 **errorUri** | **string** | Optional uri identifying a human-readable web page with information about the error, used to provide the client developer with additional information about the error | 

### Return type

[**UnexpectedProblemResponse**](UnexpectedProblemResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketplaceOAuth2InstallGetURL

> ClientAppStartResponse MarketplaceOAuth2InstallGetURL(ctx, type_).Execute()

Get OAuth2 install URL



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
	type_ := openapiclient.AppType("stripe") // AppType | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.MarketplaceOAuth2InstallGetURL(context.Background(), type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.MarketplaceOAuth2InstallGetURL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceOAuth2InstallGetURL`: ClientAppStartResponse
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.MarketplaceOAuth2InstallGetURL`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | [**AppType**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceOAuth2InstallGetURLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClientAppStartResponse**](ClientAppStartResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UninstallApp

> UninstallApp(ctx, id).Execute()

Uninstall app



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppsAPI.UninstallApp(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.UninstallApp``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUninstallAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApp

> App UpdateApp(ctx, id).AppReplaceUpdate(appReplaceUpdate).Execute()

Update app



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
	appReplaceUpdate := openapiclient.AppReplaceUpdate{CustomInvoicingAppReplaceUpdate: openapiclient.NewCustomInvoicingAppReplaceUpdate("Name_example", "Type_example", false, false)} // AppReplaceUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.UpdateApp(context.Background(), id).AppReplaceUpdate(appReplaceUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.UpdateApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApp`: App
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.UpdateApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appReplaceUpdate** | [**AppReplaceUpdate**](AppReplaceUpdate.md) |  | 

### Return type

[**App**](App.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

