# \MeterPortalAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePortalToken**](MeterPortalAPI.md#CreatePortalToken) | **Post** /api/v1/portal/tokens | Create consumer portal token
[**InvalidatePortalTokens**](MeterPortalAPI.md#InvalidatePortalTokens) | **Post** /api/v1/portal/tokens/invalidate | Invalidate portal tokens
[**ListPortalTokens**](MeterPortalAPI.md#ListPortalTokens) | **Get** /api/v1/portal/tokens | List consumer portal tokens
[**QueryPortalMeter**](MeterPortalAPI.md#QueryPortalMeter) | **Get** /api/v1/portal/meters/{meterSlug}/query | Query meter Query meter



## CreatePortalToken

> PortalToken CreatePortalToken(ctx).PortalToken(portalToken).Execute()

Create consumer portal token



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
	portalToken := *openapiclient.NewPortalToken("customer-1") // PortalToken | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeterPortalAPI.CreatePortalToken(context.Background()).PortalToken(portalToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeterPortalAPI.CreatePortalToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePortalToken`: PortalToken
	fmt.Fprintf(os.Stdout, "Response from `MeterPortalAPI.CreatePortalToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePortalTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **portalToken** | [**PortalToken**](PortalToken.md) |  | 

### Return type

[**PortalToken**](PortalToken.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvalidatePortalTokens

> InvalidatePortalTokens(ctx).InvalidatePortalTokensRequest(invalidatePortalTokensRequest).Execute()

Invalidate portal tokens



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
	invalidatePortalTokensRequest := *openapiclient.NewInvalidatePortalTokensRequest() // InvalidatePortalTokensRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MeterPortalAPI.InvalidatePortalTokens(context.Background()).InvalidatePortalTokensRequest(invalidatePortalTokensRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeterPortalAPI.InvalidatePortalTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvalidatePortalTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invalidatePortalTokensRequest** | [**InvalidatePortalTokensRequest**](InvalidatePortalTokensRequest.md) |  | 

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


## ListPortalTokens

> []PortalToken ListPortalTokens(ctx).Limit(limit).Execute()

List consumer portal tokens



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
	limit := int32(56) // int32 |  (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeterPortalAPI.ListPortalTokens(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeterPortalAPI.ListPortalTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPortalTokens`: []PortalToken
	fmt.Fprintf(os.Stdout, "Response from `MeterPortalAPI.ListPortalTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPortalTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 25]

### Return type

[**[]PortalToken**](PortalToken.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryPortalMeter

> MeterQueryResult QueryPortalMeter(ctx, meterSlug).ClientId(clientId).From(from).To(to).WindowSize(windowSize).WindowTimeZone(windowTimeZone).FilterCustomerId(filterCustomerId).FilterGroupBy(filterGroupBy).AdvancedMeterGroupByFilters(advancedMeterGroupByFilters).GroupBy(groupBy).Execute()

Query meter Query meter



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
	meterSlug := "meterSlug_example" // string | 
	clientId := "clientId_example" // string | Client ID Useful to track progress of a query. (optional)
	from := time.Now() // time.Time | Start date-time in RFC 3339 format.  Inclusive.  For example: ?from=2025-01-01T00%3A00%3A00.000Z (optional)
	to := time.Now() // time.Time | End date-time in RFC 3339 format.  Inclusive.  For example: ?to=2025-02-01T00%3A00%3A00.000Z (optional)
	windowSize := openapiclient.WindowSize("MINUTE") // WindowSize | If not specified, a single usage aggregate will be returned for the entirety of the specified period for each subject and group.  For example: ?windowSize=DAY (optional)
	windowTimeZone := "windowTimeZone_example" // string | The value is the name of the time zone as defined in the IANA Time Zone Database (http://www.iana.org/time-zones). If not specified, the UTC timezone will be used.  For example: ?windowTimeZone=UTC (optional) (default to "UTC")
	filterCustomerId := []string{"Inner_example"} // []string | Filtering by multiple customers.  For example: ?filterCustomerId=customer-1&filterCustomerId=customer-2 (optional)
	filterGroupBy := map[string]string{"key": map[string]string{"key": "Inner_example"}} // map[string]string | Simple filter for group bys with exact match.  For example: ?filterGroupBy[vendor]=openai&filterGroupBy[model]=gpt-4-turbo (optional)
	advancedMeterGroupByFilters := map[string]FilterString{"key": *openapiclient.NewFilterString()} // map[string]FilterString | Optional advanced meter group by filters. You can use this to filter for values of the meter groupBy fields. (optional)
	groupBy := []string{"Inner_example"} // []string | If not specified a single aggregate will be returned for each subject and time window. `subject` is a reserved group by value.  For example: ?groupBy=subject&groupBy=model (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeterPortalAPI.QueryPortalMeter(context.Background(), meterSlug).ClientId(clientId).From(from).To(to).WindowSize(windowSize).WindowTimeZone(windowTimeZone).FilterCustomerId(filterCustomerId).FilterGroupBy(filterGroupBy).AdvancedMeterGroupByFilters(advancedMeterGroupByFilters).GroupBy(groupBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeterPortalAPI.QueryPortalMeter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryPortalMeter`: MeterQueryResult
	fmt.Fprintf(os.Stdout, "Response from `MeterPortalAPI.QueryPortalMeter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meterSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryPortalMeterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** | Client ID Useful to track progress of a query. | 
 **from** | **time.Time** | Start date-time in RFC 3339 format.  Inclusive.  For example: ?from&#x3D;2025-01-01T00%3A00%3A00.000Z | 
 **to** | **time.Time** | End date-time in RFC 3339 format.  Inclusive.  For example: ?to&#x3D;2025-02-01T00%3A00%3A00.000Z | 
 **windowSize** | [**WindowSize**](WindowSize.md) | If not specified, a single usage aggregate will be returned for the entirety of the specified period for each subject and group.  For example: ?windowSize&#x3D;DAY | 
 **windowTimeZone** | **string** | The value is the name of the time zone as defined in the IANA Time Zone Database (http://www.iana.org/time-zones). If not specified, the UTC timezone will be used.  For example: ?windowTimeZone&#x3D;UTC | [default to &quot;UTC&quot;]
 **filterCustomerId** | **[]string** | Filtering by multiple customers.  For example: ?filterCustomerId&#x3D;customer-1&amp;filterCustomerId&#x3D;customer-2 | 
 **filterGroupBy** | **map[string]map[string]string** | Simple filter for group bys with exact match.  For example: ?filterGroupBy[vendor]&#x3D;openai&amp;filterGroupBy[model]&#x3D;gpt-4-turbo | 
 **advancedMeterGroupByFilters** | [**map[string]FilterString**](FilterString.md) | Optional advanced meter group by filters. You can use this to filter for values of the meter groupBy fields. | 
 **groupBy** | **[]string** | If not specified a single aggregate will be returned for each subject and time window. &#x60;subject&#x60; is a reserved group by value.  For example: ?groupBy&#x3D;subject&amp;groupBy&#x3D;model | 

### Return type

[**MeterQueryResult**](MeterQueryResult.md)

### Authorization

[PortalTokenAuth](../README.md#PortalTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

