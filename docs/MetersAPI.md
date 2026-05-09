# \MetersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMeter**](MetersAPI.md#CreateMeter) | **Post** /api/v1/meters | Create meter
[**DeleteMeter**](MetersAPI.md#DeleteMeter) | **Delete** /api/v1/meters/{meterIdOrSlug} | Delete meter
[**GetMeter**](MetersAPI.md#GetMeter) | **Get** /api/v1/meters/{meterIdOrSlug} | Get meter
[**GetUsageSummary**](MetersAPI.md#GetUsageSummary) | **Get** /api/v1/usage/summary | Get multi-meter usage summary
[**ListMeterGroupByValues**](MetersAPI.md#ListMeterGroupByValues) | **Get** /api/v1/meters/{meterIdOrSlug}/group-by/{groupByKey}/values | List meter group by values
[**ListMeterSubjects**](MetersAPI.md#ListMeterSubjects) | **Get** /api/v1/meters/{meterIdOrSlug}/subjects | List meter subjects
[**ListMeters**](MetersAPI.md#ListMeters) | **Get** /api/v1/meters | List meters
[**QueryMeter**](MetersAPI.md#QueryMeter) | **Get** /api/v1/meters/{meterIdOrSlug}/query | Query meter
[**QueryMeterPost**](MetersAPI.md#QueryMeterPost) | **Post** /api/v1/meters/{meterIdOrSlug}/query | Query meter
[**TestMeterEvent**](MetersAPI.md#TestMeterEvent) | **Post** /api/v1/meters/{meterIdOrSlug}/test-event | Validate a sample event payload against the meter
[**UpdateMeter**](MetersAPI.md#UpdateMeter) | **Put** /api/v1/meters/{meterIdOrSlug} | Update meter



## CreateMeter

> Meter CreateMeter(ctx).MeterCreate(meterCreate).Execute()

Create meter



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
	meterCreate := *openapiclient.NewMeterCreate("tokens_total", openapiclient.MeterAggregation("SUM"), "prompt") // MeterCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetersAPI.CreateMeter(context.Background()).MeterCreate(meterCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetersAPI.CreateMeter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMeter`: Meter
	fmt.Fprintf(os.Stdout, "Response from `MetersAPI.CreateMeter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMeterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **meterCreate** | [**MeterCreate**](MeterCreate.md) |  | 

### Return type

[**Meter**](Meter.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMeter

> DeleteMeter(ctx, meterIdOrSlug).Execute()

Delete meter



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
	meterIdOrSlug := "meterIdOrSlug_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MetersAPI.DeleteMeter(context.Background(), meterIdOrSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetersAPI.DeleteMeter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meterIdOrSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMeterRequest struct via the builder pattern


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


## GetMeter

> Meter GetMeter(ctx, meterIdOrSlug).Execute()

Get meter



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
	meterIdOrSlug := "meterIdOrSlug_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetersAPI.GetMeter(context.Background(), meterIdOrSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetersAPI.GetMeter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMeter`: Meter
	fmt.Fprintf(os.Stdout, "Response from `MetersAPI.GetMeter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meterIdOrSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMeterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Meter**](Meter.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageSummary

> MeterUsageSummaryResponse GetUsageSummary(ctx).Subject(subject).From(from).To(to).WindowSize(windowSize).MeterSlugs(meterSlugs).Execute()

Get multi-meter usage summary



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
	subject := "subject_example" // string | The subject to query usage for (required).
	from := time.Now() // time.Time | Start of the time range (RFC 3339, required).
	to := time.Now() // time.Time | End of the time range (RFC 3339, required).
	windowSize := openapiclient.UsageSummaryWindowSize("HOUR") // UsageSummaryWindowSize | Window size for breakdown. Omit for totals only. (optional)
	meterSlugs := []string{"Inner_example"} // []string | Filter to specific meter slugs. Omit to query all meters. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetersAPI.GetUsageSummary(context.Background()).Subject(subject).From(from).To(to).WindowSize(windowSize).MeterSlugs(meterSlugs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetersAPI.GetUsageSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsageSummary`: MeterUsageSummaryResponse
	fmt.Fprintf(os.Stdout, "Response from `MetersAPI.GetUsageSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subject** | **string** | The subject to query usage for (required). | 
 **from** | **time.Time** | Start of the time range (RFC 3339, required). | 
 **to** | **time.Time** | End of the time range (RFC 3339, required). | 
 **windowSize** | [**UsageSummaryWindowSize**](UsageSummaryWindowSize.md) | Window size for breakdown. Omit for totals only. | 
 **meterSlugs** | **[]string** | Filter to specific meter slugs. Omit to query all meters. | 

### Return type

[**MeterUsageSummaryResponse**](MeterUsageSummaryResponse.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMeterGroupByValues

> []string ListMeterGroupByValues(ctx, meterIdOrSlug, groupByKey).From(from).To(to).Execute()

List meter group by values



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
	meterIdOrSlug := "meterIdOrSlug_example" // string | 
	groupByKey := "groupByKey_example" // string | 
	from := time.Now() // time.Time | Start date-time in RFC 3339 format.  Inclusive. Defaults to 24 hours ago.  For example: ?from=2025-01-01T00%3A00%3A00.000Z (optional)
	to := time.Now() // time.Time | End date-time in RFC 3339 format.  Inclusive.  For example: ?to=2025-02-01T00%3A00%3A00.000Z (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetersAPI.ListMeterGroupByValues(context.Background(), meterIdOrSlug, groupByKey).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetersAPI.ListMeterGroupByValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMeterGroupByValues`: []string
	fmt.Fprintf(os.Stdout, "Response from `MetersAPI.ListMeterGroupByValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meterIdOrSlug** | **string** |  | 
**groupByKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMeterGroupByValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **from** | **time.Time** | Start date-time in RFC 3339 format.  Inclusive. Defaults to 24 hours ago.  For example: ?from&#x3D;2025-01-01T00%3A00%3A00.000Z | 
 **to** | **time.Time** | End date-time in RFC 3339 format.  Inclusive.  For example: ?to&#x3D;2025-02-01T00%3A00%3A00.000Z | 

### Return type

**[]string**

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMeterSubjects

> []string ListMeterSubjects(ctx, meterIdOrSlug).From(from).To(to).Execute()

List meter subjects



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
	meterIdOrSlug := "meterIdOrSlug_example" // string | 
	from := time.Now() // time.Time | Start date-time in RFC 3339 format.  Inclusive. Defaults to the beginning of time.  For example: ?from=2025-01-01T00%3A00%3A00.000Z (optional)
	to := time.Now() // time.Time | End date-time in RFC 3339 format.  Inclusive.  For example: ?to=2025-02-01T00%3A00%3A00.000Z (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetersAPI.ListMeterSubjects(context.Background(), meterIdOrSlug).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetersAPI.ListMeterSubjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMeterSubjects`: []string
	fmt.Fprintf(os.Stdout, "Response from `MetersAPI.ListMeterSubjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meterIdOrSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMeterSubjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **time.Time** | Start date-time in RFC 3339 format.  Inclusive. Defaults to the beginning of time.  For example: ?from&#x3D;2025-01-01T00%3A00%3A00.000Z | 
 **to** | **time.Time** | End date-time in RFC 3339 format.  Inclusive.  For example: ?to&#x3D;2025-02-01T00%3A00%3A00.000Z | 

### Return type

**[]string**

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMeters

> []Meter ListMeters(ctx).Page(page).PageSize(pageSize).Order(order).OrderBy(orderBy).IncludeDeleted(includeDeleted).Execute()

List meters



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
	orderBy := openapiclient.MeterOrderBy("key") // MeterOrderBy | The order by field. (optional)
	includeDeleted := true // bool | Include deleted meters. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetersAPI.ListMeters(context.Background()).Page(page).PageSize(pageSize).Order(order).OrderBy(orderBy).IncludeDeleted(includeDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetersAPI.ListMeters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMeters`: []Meter
	fmt.Fprintf(os.Stdout, "Response from `MetersAPI.ListMeters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMetersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page index.  Default is 1. | [default to 1]
 **pageSize** | **int32** | The maximum number of items per page.  Default is 100. | [default to 100]
 **order** | [**SortOrder**](SortOrder.md) | The order direction. | 
 **orderBy** | [**MeterOrderBy**](MeterOrderBy.md) | The order by field. | 
 **includeDeleted** | **bool** | Include deleted meters. | [default to false]

### Return type

[**[]Meter**](Meter.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryMeter

> MeterQueryResult QueryMeter(ctx, meterIdOrSlug).ClientId(clientId).From(from).To(to).WindowSize(windowSize).WindowTimeZone(windowTimeZone).Subject(subject).FilterCustomerId(filterCustomerId).FilterGroupBy(filterGroupBy).AdvancedMeterGroupByFilters(advancedMeterGroupByFilters).GroupBy(groupBy).Execute()

Query meter



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
	meterIdOrSlug := "meterIdOrSlug_example" // string | 
	clientId := "clientId_example" // string | Client ID Useful to track progress of a query. (optional)
	from := time.Now() // time.Time | Start date-time in RFC 3339 format.  Inclusive.  For example: ?from=2025-01-01T00%3A00%3A00.000Z (optional)
	to := time.Now() // time.Time | End date-time in RFC 3339 format.  Inclusive.  For example: ?to=2025-02-01T00%3A00%3A00.000Z (optional)
	windowSize := openapiclient.WindowSize("MINUTE") // WindowSize | If not specified, a single usage aggregate will be returned for the entirety of the specified period for each subject and group.  For example: ?windowSize=DAY (optional)
	windowTimeZone := "windowTimeZone_example" // string | The value is the name of the time zone as defined in the IANA Time Zone Database (http://www.iana.org/time-zones). If not specified, the UTC timezone will be used.  For example: ?windowTimeZone=UTC (optional) (default to "UTC")
	subject := []string{"Inner_example"} // []string | Filtering by multiple subjects.  For example: ?subject=subject-1&subject=subject-2 (optional)
	filterCustomerId := []string{"Inner_example"} // []string | Filtering by multiple customers.  For example: ?filterCustomerId=customer-1&filterCustomerId=customer-2 (optional)
	filterGroupBy := map[string]string{"key": map[string]string{"key": "Inner_example"}} // map[string]string | Simple filter for group bys with exact match.  For example: ?filterGroupBy[vendor]=openai&filterGroupBy[model]=gpt-4-turbo (optional)
	advancedMeterGroupByFilters := map[string]FilterString{"key": *openapiclient.NewFilterString()} // map[string]FilterString | Optional advanced meter group by filters. You can use this to filter for values of the meter groupBy fields. (optional)
	groupBy := []string{"Inner_example"} // []string | If not specified a single aggregate will be returned for each subject and time window. `subject` is a reserved group by value.  For example: ?groupBy=subject&groupBy=model (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetersAPI.QueryMeter(context.Background(), meterIdOrSlug).ClientId(clientId).From(from).To(to).WindowSize(windowSize).WindowTimeZone(windowTimeZone).Subject(subject).FilterCustomerId(filterCustomerId).FilterGroupBy(filterGroupBy).AdvancedMeterGroupByFilters(advancedMeterGroupByFilters).GroupBy(groupBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetersAPI.QueryMeter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryMeter`: MeterQueryResult
	fmt.Fprintf(os.Stdout, "Response from `MetersAPI.QueryMeter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meterIdOrSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryMeterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** | Client ID Useful to track progress of a query. | 
 **from** | **time.Time** | Start date-time in RFC 3339 format.  Inclusive.  For example: ?from&#x3D;2025-01-01T00%3A00%3A00.000Z | 
 **to** | **time.Time** | End date-time in RFC 3339 format.  Inclusive.  For example: ?to&#x3D;2025-02-01T00%3A00%3A00.000Z | 
 **windowSize** | [**WindowSize**](WindowSize.md) | If not specified, a single usage aggregate will be returned for the entirety of the specified period for each subject and group.  For example: ?windowSize&#x3D;DAY | 
 **windowTimeZone** | **string** | The value is the name of the time zone as defined in the IANA Time Zone Database (http://www.iana.org/time-zones). If not specified, the UTC timezone will be used.  For example: ?windowTimeZone&#x3D;UTC | [default to &quot;UTC&quot;]
 **subject** | **[]string** | Filtering by multiple subjects.  For example: ?subject&#x3D;subject-1&amp;subject&#x3D;subject-2 | 
 **filterCustomerId** | **[]string** | Filtering by multiple customers.  For example: ?filterCustomerId&#x3D;customer-1&amp;filterCustomerId&#x3D;customer-2 | 
 **filterGroupBy** | **map[string]map[string]string** | Simple filter for group bys with exact match.  For example: ?filterGroupBy[vendor]&#x3D;openai&amp;filterGroupBy[model]&#x3D;gpt-4-turbo | 
 **advancedMeterGroupByFilters** | [**map[string]FilterString**](FilterString.md) | Optional advanced meter group by filters. You can use this to filter for values of the meter groupBy fields. | 
 **groupBy** | **[]string** | If not specified a single aggregate will be returned for each subject and time window. &#x60;subject&#x60; is a reserved group by value.  For example: ?groupBy&#x3D;subject&amp;groupBy&#x3D;model | 

### Return type

[**MeterQueryResult**](MeterQueryResult.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryMeterPost

> MeterQueryResult QueryMeterPost(ctx, meterIdOrSlug).MeterQueryRequest(meterQueryRequest).Execute()

Query meter

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
	meterIdOrSlug := "meterIdOrSlug_example" // string | 
	meterQueryRequest := *openapiclient.NewMeterQueryRequest() // MeterQueryRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetersAPI.QueryMeterPost(context.Background(), meterIdOrSlug).MeterQueryRequest(meterQueryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetersAPI.QueryMeterPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryMeterPost`: MeterQueryResult
	fmt.Fprintf(os.Stdout, "Response from `MetersAPI.QueryMeterPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meterIdOrSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryMeterPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **meterQueryRequest** | [**MeterQueryRequest**](MeterQueryRequest.md) |  | 

### Return type

[**MeterQueryResult**](MeterQueryResult.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestMeterEvent

> TestMeterEvent200Response TestMeterEvent(ctx, meterIdOrSlug).TestMeterEventRequest(testMeterEventRequest).Execute()

Validate a sample event payload against the meter



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
	meterIdOrSlug := "meterIdOrSlug_example" // string | 
	testMeterEventRequest := *openapiclient.NewTestMeterEventRequest(map[string]interface{}{"key": interface{}(123)}) // TestMeterEventRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetersAPI.TestMeterEvent(context.Background(), meterIdOrSlug).TestMeterEventRequest(testMeterEventRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetersAPI.TestMeterEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestMeterEvent`: TestMeterEvent200Response
	fmt.Fprintf(os.Stdout, "Response from `MetersAPI.TestMeterEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meterIdOrSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestMeterEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **testMeterEventRequest** | [**TestMeterEventRequest**](TestMeterEventRequest.md) |  | 

### Return type

[**TestMeterEvent200Response**](TestMeterEvent200Response.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMeter

> Meter UpdateMeter(ctx, meterIdOrSlug).MeterUpdate(meterUpdate).Execute()

Update meter



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
	meterIdOrSlug := "meterIdOrSlug_example" // string | 
	meterUpdate := *openapiclient.NewMeterUpdate() // MeterUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetersAPI.UpdateMeter(context.Background(), meterIdOrSlug).MeterUpdate(meterUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetersAPI.UpdateMeter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMeter`: Meter
	fmt.Fprintf(os.Stdout, "Response from `MetersAPI.UpdateMeter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meterIdOrSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMeterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **meterUpdate** | [**MeterUpdate**](MeterUpdate.md) |  | 

### Return type

[**Meter**](Meter.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

