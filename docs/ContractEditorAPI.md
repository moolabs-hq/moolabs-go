# \ContractEditorAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiRedlineContractV1**](ContractEditorAPI.md#AiRedlineContractV1) | **Post** /v1/quotes/{quote_id}/contract/editor/ai-redline | Ai Redline Contract
[**ApplyContractEditsV1**](ContractEditorAPI.md#ApplyContractEditsV1) | **Post** /v1/quotes/{quote_id}/contract/editor/apply-edits | Apply Contract Edits
[**AskMooRedlineV1**](ContractEditorAPI.md#AskMooRedlineV1) | **Post** /v1/quotes/{quote_id}/contract/editor/redline/ask-moo | Ask Moo Redline
[**EditorCallback**](ContractEditorAPI.md#EditorCallback) | **Post** /v1/quotes/{quote_id}/contract/editor/callback | Editor Callback
[**FetchDocument**](ContractEditorAPI.md#FetchDocument) | **Get** /v1/quotes/{quote_id}/contract/editor/document/{doc_key} | Fetch Document
[**GetDocumentHtmlV1**](ContractEditorAPI.md#GetDocumentHtmlV1) | **Get** /v1/quotes/{quote_id}/contract/editor/document-html/{doc_key} | Get Document Html
[**GetEditorConfig**](ContractEditorAPI.md#GetEditorConfig) | **Get** /v1/quotes/{quote_id}/contract/editor/config | Get Editor Config
[**GetLatestRedlineJobV1**](ContractEditorAPI.md#GetLatestRedlineJobV1) | **Get** /v1/quotes/{quote_id}/contract/editor/ai-redline/latest | Get Latest Redline Job
[**GetRedlineEdits**](ContractEditorAPI.md#GetRedlineEdits) | **Get** /v1/quotes/{quote_id}/contract/editor/edits/{doc_key} | Get Redline Edits
[**GetRedlineJobV1**](ContractEditorAPI.md#GetRedlineJobV1) | **Get** /v1/quotes/{quote_id}/contract/editor/ai-redline/jobs/{job_id} | Get Redline Job
[**SaveRedlineEdit**](ContractEditorAPI.md#SaveRedlineEdit) | **Post** /v1/quotes/{quote_id}/contract/editor/edits | Save Redline Edit
[**UploadContractForEditor**](ContractEditorAPI.md#UploadContractForEditor) | **Post** /v1/quotes/{quote_id}/contract/editor/upload | Upload Contract For Editor



## AiRedlineContractV1

> interface{} AiRedlineContractV1(ctx, quoteId).Execute()

Ai Redline Contract



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
	resp, r, err := apiClient.ContractEditorAPI.AiRedlineContractV1(context.Background(), quoteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractEditorAPI.AiRedlineContractV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiRedlineContractV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ContractEditorAPI.AiRedlineContractV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAiRedlineContractV1Request struct via the builder pattern


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


## ApplyContractEditsV1

> interface{} ApplyContractEditsV1(ctx, quoteId).Execute()

Apply Contract Edits



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
	resp, r, err := apiClient.ContractEditorAPI.ApplyContractEditsV1(context.Background(), quoteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractEditorAPI.ApplyContractEditsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplyContractEditsV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ContractEditorAPI.ApplyContractEditsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplyContractEditsV1Request struct via the builder pattern


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


## AskMooRedlineV1

> interface{} AskMooRedlineV1(ctx, quoteId).AskMooRedlineRequest(askMooRedlineRequest).Execute()

Ask Moo Redline



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
	askMooRedlineRequest := *openapiclient.NewAskMooRedlineRequest("FindingId_example", int32(123), int32(123), "Instruction_example") // AskMooRedlineRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractEditorAPI.AskMooRedlineV1(context.Background(), quoteId).AskMooRedlineRequest(askMooRedlineRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractEditorAPI.AskMooRedlineV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AskMooRedlineV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ContractEditorAPI.AskMooRedlineV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAskMooRedlineV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **askMooRedlineRequest** | [**AskMooRedlineRequest**](AskMooRedlineRequest.md) |  | 

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


## EditorCallback

> interface{} EditorCallback(ctx, quoteId).T(t).Execute()

Editor Callback



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
	t := "t_example" // string | BFF URL token

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractEditorAPI.EditorCallback(context.Background(), quoteId).T(t).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractEditorAPI.EditorCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditorCallback`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ContractEditorAPI.EditorCallback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditorCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t** | **string** | BFF URL token | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchDocument

> interface{} FetchDocument(ctx, quoteId, docKey).T(t).Execute()

Fetch Document



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
	docKey := "docKey_example" // string | 
	t := "t_example" // string | BFF URL token

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractEditorAPI.FetchDocument(context.Background(), quoteId, docKey).T(t).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractEditorAPI.FetchDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchDocument`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ContractEditorAPI.FetchDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 
**docKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **t** | **string** | BFF URL token | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocumentHtmlV1

> interface{} GetDocumentHtmlV1(ctx, quoteId, docKey).Execute()

Get Document Html



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
	docKey := "docKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractEditorAPI.GetDocumentHtmlV1(context.Background(), quoteId, docKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractEditorAPI.GetDocumentHtmlV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDocumentHtmlV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ContractEditorAPI.GetDocumentHtmlV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 
**docKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentHtmlV1Request struct via the builder pattern


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


## GetEditorConfig

> interface{} GetEditorConfig(ctx, quoteId).Execute()

Get Editor Config



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
	resp, r, err := apiClient.ContractEditorAPI.GetEditorConfig(context.Background(), quoteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractEditorAPI.GetEditorConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEditorConfig`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ContractEditorAPI.GetEditorConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEditorConfigRequest struct via the builder pattern


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


## GetLatestRedlineJobV1

> interface{} GetLatestRedlineJobV1(ctx, quoteId).Execute()

Get Latest Redline Job



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
	resp, r, err := apiClient.ContractEditorAPI.GetLatestRedlineJobV1(context.Background(), quoteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractEditorAPI.GetLatestRedlineJobV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLatestRedlineJobV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ContractEditorAPI.GetLatestRedlineJobV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestRedlineJobV1Request struct via the builder pattern


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


## GetRedlineEdits

> interface{} GetRedlineEdits(ctx, quoteId, docKey).Execute()

Get Redline Edits



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
	docKey := "docKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractEditorAPI.GetRedlineEdits(context.Background(), quoteId, docKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractEditorAPI.GetRedlineEdits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedlineEdits`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ContractEditorAPI.GetRedlineEdits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 
**docKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedlineEditsRequest struct via the builder pattern


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


## GetRedlineJobV1

> interface{} GetRedlineJobV1(ctx, quoteId, jobId).Execute()

Get Redline Job



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
	jobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractEditorAPI.GetRedlineJobV1(context.Background(), quoteId, jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractEditorAPI.GetRedlineJobV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedlineJobV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ContractEditorAPI.GetRedlineJobV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedlineJobV1Request struct via the builder pattern


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


## SaveRedlineEdit

> interface{} SaveRedlineEdit(ctx, quoteId).SaveRedlineEditRequest(saveRedlineEditRequest).Execute()

Save Redline Edit



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
	saveRedlineEditRequest := *openapiclient.NewSaveRedlineEditRequest("DocKey_example", int32(123), "Op_example", "Before_example", "After_example") // SaveRedlineEditRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractEditorAPI.SaveRedlineEdit(context.Background(), quoteId).SaveRedlineEditRequest(saveRedlineEditRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractEditorAPI.SaveRedlineEdit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SaveRedlineEdit`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ContractEditorAPI.SaveRedlineEdit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSaveRedlineEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **saveRedlineEditRequest** | [**SaveRedlineEditRequest**](SaveRedlineEditRequest.md) |  | 

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


## UploadContractForEditor

> interface{} UploadContractForEditor(ctx, quoteId).File(file).Execute()

Upload Contract For Editor



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
	file := os.NewFile(1234, "some_file") // *os.File | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractEditorAPI.UploadContractForEditor(context.Background(), quoteId).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractEditorAPI.UploadContractForEditor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadContractForEditor`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ContractEditorAPI.UploadContractForEditor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadContractForEditorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** |  | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

