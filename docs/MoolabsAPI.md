# \MoolabsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCredit**](MoolabsAPI.md#CreateCredit) | **Post** /v1/moolabs/credits | Create Credit
[**GetCredits**](MoolabsAPI.md#GetCredits) | **Get** /v1/moolabs/credits | Get Credits
[**GetCreditsIssuedAggregateV1**](MoolabsAPI.md#GetCreditsIssuedAggregateV1) | **Get** /v1/moolabs/credits-issued/aggregate | Get Credits Issued Aggregate
[**GetCreditsIssuedV1**](MoolabsAPI.md#GetCreditsIssuedV1) | **Get** /v1/moolabs/credits-issued | Get Credits Issued



## CreateCredit

> interface{} CreateCredit(ctx).CreateCreditRequest(createCreditRequest).Execute()

Create Credit



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
	createCreditRequest := *openapiclient.NewCreateCreditRequest("InvoiceId_example", "CustomerId_example", float32(123), "Reason_example") // CreateCreditRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MoolabsAPI.CreateCredit(context.Background()).CreateCreditRequest(createCreditRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MoolabsAPI.CreateCredit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCredit`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `MoolabsAPI.CreateCredit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCreditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCreditRequest** | [**CreateCreditRequest**](CreateCreditRequest.md) |  | 

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


## GetCredits

> interface{} GetCredits(ctx).CustomerId(customerId).InvoiceId(invoiceId).Execute()

Get Credits



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
	customerId := "customerId_example" // string | Customer ID to filter by (optional)
	invoiceId := "invoiceId_example" // string | Invoice ID to filter by (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MoolabsAPI.GetCredits(context.Background()).CustomerId(customerId).InvoiceId(invoiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MoolabsAPI.GetCredits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCredits`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `MoolabsAPI.GetCredits`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCreditsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerId** | **string** | Customer ID to filter by | 
 **invoiceId** | **string** | Invoice ID to filter by | 

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


## GetCreditsIssuedAggregateV1

> interface{} GetCreditsIssuedAggregateV1(ctx).FromDate(fromDate).ToDate(toDate).Execute()

Get Credits Issued Aggregate



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
	fromDate := "fromDate_example" // string | Start date (ISO format) - defaults to start of current month (optional)
	toDate := "toDate_example" // string | End date (ISO format) - defaults to now (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MoolabsAPI.GetCreditsIssuedAggregateV1(context.Background()).FromDate(fromDate).ToDate(toDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MoolabsAPI.GetCreditsIssuedAggregateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCreditsIssuedAggregateV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `MoolabsAPI.GetCreditsIssuedAggregateV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCreditsIssuedAggregateV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromDate** | **string** | Start date (ISO format) - defaults to start of current month | 
 **toDate** | **string** | End date (ISO format) - defaults to now | 

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


## GetCreditsIssuedV1

> interface{} GetCreditsIssuedV1(ctx).CustomerId(customerId).FromDate(fromDate).ToDate(toDate).Execute()

Get Credits Issued



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
	customerId := "customerId_example" // string | Customer ID (subject_id) to filter by (optional)
	fromDate := "fromDate_example" // string | Start date (ISO format) - defaults to start of current month (optional)
	toDate := "toDate_example" // string | End date (ISO format) - defaults to now (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MoolabsAPI.GetCreditsIssuedV1(context.Background()).CustomerId(customerId).FromDate(fromDate).ToDate(toDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MoolabsAPI.GetCreditsIssuedV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCreditsIssuedV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `MoolabsAPI.GetCreditsIssuedV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCreditsIssuedV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerId** | **string** | Customer ID (subject_id) to filter by | 
 **fromDate** | **string** | Start date (ISO format) - defaults to start of current month | 
 **toDate** | **string** | End date (ISO format) - defaults to now | 

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

