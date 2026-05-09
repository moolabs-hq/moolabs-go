# \EventsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IngestEvents**](EventsAPI.md#IngestEvents) | **Post** /api/v1/events | Ingest events  
[**ListEvents**](EventsAPI.md#ListEvents) | **Get** /api/v1/events | List ingested events
[**ListEventsV2**](EventsAPI.md#ListEventsV2) | **Get** /api/v2/events | List ingested events



## IngestEvents

> IngestEvents(ctx).Event(event).Execute()

Ingest events  



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
	event := *openapiclient.NewEvent("5c10fade-1c9e-4d6c-8275-c52c36731d3c", "service-name", "1.0", "com.example.someevent", "customer-id") // Event | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.IngestEvents(context.Background()).Event(event).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.IngestEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIngestEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **event** | [**Event**](Event.md) |  | 

### Return type

 (empty response body)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: application/cloudevents+json, application/cloudevents-batch+json, application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEvents

> []IngestedEvent ListEvents(ctx).ClientId(clientId).IngestedAtFrom(ingestedAtFrom).IngestedAtTo(ingestedAtTo).Id(id).Subject(subject).CustomerId(customerId).From(from).To(to).Limit(limit).Execute()

List ingested events



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
	clientId := "clientId_example" // string | Client ID Useful to track progress of a query. (optional)
	ingestedAtFrom := time.Now() // time.Time | Start date-time in RFC 3339 format.  Inclusive. (optional)
	ingestedAtTo := time.Now() // time.Time | End date-time in RFC 3339 format.  Inclusive. (optional)
	id := "id_example" // string | The event ID.  Accepts partial ID. (optional)
	subject := "subject_example" // string | The event subject.  Accepts partial subject. (optional)
	customerId := []string{"01G65Z755AFWAKHE12NY0CQ9FH"} // []string | The event customer ID. (optional)
	from := time.Now() // time.Time | Start date-time in RFC 3339 format.  Inclusive. (optional)
	to := time.Now() // time.Time | End date-time in RFC 3339 format.  Inclusive. (optional)
	limit := int32(56) // int32 | Number of events to return. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.ListEvents(context.Background()).ClientId(clientId).IngestedAtFrom(ingestedAtFrom).IngestedAtTo(ingestedAtTo).Id(id).Subject(subject).CustomerId(customerId).From(from).To(to).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.ListEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEvents`: []IngestedEvent
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.ListEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** | Client ID Useful to track progress of a query. | 
 **ingestedAtFrom** | **time.Time** | Start date-time in RFC 3339 format.  Inclusive. | 
 **ingestedAtTo** | **time.Time** | End date-time in RFC 3339 format.  Inclusive. | 
 **id** | **string** | The event ID.  Accepts partial ID. | 
 **subject** | **string** | The event subject.  Accepts partial subject. | 
 **customerId** | **[]string** | The event customer ID. | 
 **from** | **time.Time** | Start date-time in RFC 3339 format.  Inclusive. | 
 **to** | **time.Time** | End date-time in RFC 3339 format.  Inclusive. | 
 **limit** | **int32** | Number of events to return. | [default to 100]

### Return type

[**[]IngestedEvent**](IngestedEvent.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEventsV2

> IngestedEventCursorPaginatedResponse ListEventsV2(ctx).Cursor(cursor).Limit(limit).ClientId(clientId).Filter(filter).Execute()

List ingested events



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
	cursor := "cursor_example" // string | The cursor after which to start the pagination. (optional)
	limit := int32(56) // int32 | The limit of the pagination. (optional) (default to 100)
	clientId := "clientId_example" // string | Client ID Useful to track progress of a query. (optional)
	filter := *openapiclient.NewListEventsV2FilterParam() // ListEventsV2FilterParam | The filter for the events encoded as JSON string. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.ListEventsV2(context.Background()).Cursor(cursor).Limit(limit).ClientId(clientId).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.ListEventsV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEventsV2`: IngestedEventCursorPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.ListEventsV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEventsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | The cursor after which to start the pagination. | 
 **limit** | **int32** | The limit of the pagination. | [default to 100]
 **clientId** | **string** | Client ID Useful to track progress of a query. | 
 **filter** | [**ListEventsV2FilterParam**](ListEventsV2FilterParam.md) | The filter for the events encoded as JSON string. | 

### Return type

[**IngestedEventCursorPaginatedResponse**](IngestedEventCursorPaginatedResponse.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

