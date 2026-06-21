# \WordAddinAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WordOpen**](WordAddinAPI.md#WordOpen) | **Get** /v1/quotes/{quote_id}/contract/editor/word/open | Word Open
[**WordSync**](WordAddinAPI.md#WordSync) | **Post** /v1/quotes/{quote_id}/contract/editor/word/sync | Word Sync



## WordOpen

> interface{} WordOpen(ctx, quoteId).Execute()

Word Open



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
	quoteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WordAddinAPI.WordOpen(context.Background(), quoteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WordAddinAPI.WordOpen``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WordOpen`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WordAddinAPI.WordOpen`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWordOpenRequest struct via the builder pattern


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


## WordSync

> interface{} WordSync(ctx, quoteId).File(file).BaseVersion(baseVersion).XBindToken(xBindToken).Execute()

Word Sync



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
	quoteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	file := os.NewFile(1234, "some_file") // *os.File | The current .docx from the add-in
	baseVersion := int32(56) // int32 | The version_no the add-in derived this doc from
	xBindToken := "xBindToken_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WordAddinAPI.WordSync(context.Background(), quoteId).File(file).BaseVersion(baseVersion).XBindToken(xBindToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WordAddinAPI.WordSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WordSync`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WordAddinAPI.WordSync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWordSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** | The current .docx from the add-in | 
 **baseVersion** | **int32** | The version_no the add-in derived this doc from | 
 **xBindToken** | **string** |  | [default to &quot;&quot;]

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

