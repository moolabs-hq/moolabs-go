# \QuoteContractsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddQuoteContractParticipantV1**](QuoteContractsAPI.md#AddQuoteContractParticipantV1) | **Post** /v1/quotes/{quote_id}/contract-thread/participants | Add Quote Contract Participant
[**CreateQuoteContractThreadV1**](QuoteContractsAPI.md#CreateQuoteContractThreadV1) | **Post** /v1/quotes/{quote_id}/contract-thread | Create Quote Contract Thread
[**GetQuoteContractThreadV1**](QuoteContractsAPI.md#GetQuoteContractThreadV1) | **Get** /v1/quotes/{quote_id}/contract-thread | Get Quote Contract Thread
[**ListQuoteContractQuarantineV1**](QuoteContractsAPI.md#ListQuoteContractQuarantineV1) | **Get** /v1/quotes/{quote_id}/contract-thread/quarantine | List Quote Contract Quarantine
[**ListQuoteContractVersionsV1**](QuoteContractsAPI.md#ListQuoteContractVersionsV1) | **Get** /v1/quotes/{quote_id}/contract-thread/versions | List Quote Contract Versions
[**PromoteQuoteContractQuarantineSenderV1**](QuoteContractsAPI.md#PromoteQuoteContractQuarantineSenderV1) | **Post** /v1/quotes/{quote_id}/contract-thread/quarantine/{message_id}/promote | Promote Quote Contract Quarantine Sender
[**RemoveQuoteContractParticipantV1**](QuoteContractsAPI.md#RemoveQuoteContractParticipantV1) | **Delete** /v1/quotes/{quote_id}/contract-thread/participants/{participant_id} | Remove Quote Contract Participant
[**ReprocessQuoteContractQuarantineMessageV1**](QuoteContractsAPI.md#ReprocessQuoteContractQuarantineMessageV1) | **Post** /v1/quotes/{quote_id}/contract-thread/quarantine/{message_id}/reprocess | Reprocess Quote Contract Quarantine Message
[**UploadQuoteContractVersionV1**](QuoteContractsAPI.md#UploadQuoteContractVersionV1) | **Post** /v1/quotes/{quote_id}/contract-thread/versions | Upload Quote Contract Version



## AddQuoteContractParticipantV1

> ContractParticipantResponse AddQuoteContractParticipantV1(ctx, quoteId).AddContractParticipantRequest(addContractParticipantRequest).Execute()

Add Quote Contract Participant

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
	addContractParticipantRequest := *openapiclient.NewAddContractParticipantRequest("Side_example", "Role_example") // AddContractParticipantRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuoteContractsAPI.AddQuoteContractParticipantV1(context.Background(), quoteId).AddContractParticipantRequest(addContractParticipantRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuoteContractsAPI.AddQuoteContractParticipantV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddQuoteContractParticipantV1`: ContractParticipantResponse
	fmt.Fprintf(os.Stdout, "Response from `QuoteContractsAPI.AddQuoteContractParticipantV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddQuoteContractParticipantV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addContractParticipantRequest** | [**AddContractParticipantRequest**](AddContractParticipantRequest.md) |  | 

### Return type

[**ContractParticipantResponse**](ContractParticipantResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateQuoteContractThreadV1

> ContractThreadSnapshotResponse CreateQuoteContractThreadV1(ctx, quoteId).CreateContractThreadRequest(createContractThreadRequest).Execute()

Create Quote Contract Thread

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
	createContractThreadRequest := *openapiclient.NewCreateContractThreadRequest() // CreateContractThreadRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuoteContractsAPI.CreateQuoteContractThreadV1(context.Background(), quoteId).CreateContractThreadRequest(createContractThreadRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuoteContractsAPI.CreateQuoteContractThreadV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateQuoteContractThreadV1`: ContractThreadSnapshotResponse
	fmt.Fprintf(os.Stdout, "Response from `QuoteContractsAPI.CreateQuoteContractThreadV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateQuoteContractThreadV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createContractThreadRequest** | [**CreateContractThreadRequest**](CreateContractThreadRequest.md) |  | 

### Return type

[**ContractThreadSnapshotResponse**](ContractThreadSnapshotResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuoteContractThreadV1

> ContractThreadSnapshotResponse GetQuoteContractThreadV1(ctx, quoteId).Execute()

Get Quote Contract Thread

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
	resp, r, err := apiClient.QuoteContractsAPI.GetQuoteContractThreadV1(context.Background(), quoteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuoteContractsAPI.GetQuoteContractThreadV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuoteContractThreadV1`: ContractThreadSnapshotResponse
	fmt.Fprintf(os.Stdout, "Response from `QuoteContractsAPI.GetQuoteContractThreadV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQuoteContractThreadV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContractThreadSnapshotResponse**](ContractThreadSnapshotResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListQuoteContractQuarantineV1

> []ContractInboundMessageResponse ListQuoteContractQuarantineV1(ctx, quoteId).Execute()

List Quote Contract Quarantine

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
	resp, r, err := apiClient.QuoteContractsAPI.ListQuoteContractQuarantineV1(context.Background(), quoteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuoteContractsAPI.ListQuoteContractQuarantineV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListQuoteContractQuarantineV1`: []ContractInboundMessageResponse
	fmt.Fprintf(os.Stdout, "Response from `QuoteContractsAPI.ListQuoteContractQuarantineV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListQuoteContractQuarantineV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ContractInboundMessageResponse**](ContractInboundMessageResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListQuoteContractVersionsV1

> []ContractVersionResponse ListQuoteContractVersionsV1(ctx, quoteId).Execute()

List Quote Contract Versions

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
	resp, r, err := apiClient.QuoteContractsAPI.ListQuoteContractVersionsV1(context.Background(), quoteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuoteContractsAPI.ListQuoteContractVersionsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListQuoteContractVersionsV1`: []ContractVersionResponse
	fmt.Fprintf(os.Stdout, "Response from `QuoteContractsAPI.ListQuoteContractVersionsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListQuoteContractVersionsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ContractVersionResponse**](ContractVersionResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PromoteQuoteContractQuarantineSenderV1

> ContractParticipantResponse PromoteQuoteContractQuarantineSenderV1(ctx, quoteId, messageId).AddContractParticipantRequest(addContractParticipantRequest).Execute()

Promote Quote Contract Quarantine Sender

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
	messageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	addContractParticipantRequest := *openapiclient.NewAddContractParticipantRequest("Side_example", "Role_example") // AddContractParticipantRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuoteContractsAPI.PromoteQuoteContractQuarantineSenderV1(context.Background(), quoteId, messageId).AddContractParticipantRequest(addContractParticipantRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuoteContractsAPI.PromoteQuoteContractQuarantineSenderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PromoteQuoteContractQuarantineSenderV1`: ContractParticipantResponse
	fmt.Fprintf(os.Stdout, "Response from `QuoteContractsAPI.PromoteQuoteContractQuarantineSenderV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 
**messageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPromoteQuoteContractQuarantineSenderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addContractParticipantRequest** | [**AddContractParticipantRequest**](AddContractParticipantRequest.md) |  | 

### Return type

[**ContractParticipantResponse**](ContractParticipantResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveQuoteContractParticipantV1

> RemoveContractParticipantResponse RemoveQuoteContractParticipantV1(ctx, quoteId, participantId).Execute()

Remove Quote Contract Participant

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
	participantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuoteContractsAPI.RemoveQuoteContractParticipantV1(context.Background(), quoteId, participantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuoteContractsAPI.RemoveQuoteContractParticipantV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveQuoteContractParticipantV1`: RemoveContractParticipantResponse
	fmt.Fprintf(os.Stdout, "Response from `QuoteContractsAPI.RemoveQuoteContractParticipantV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 
**participantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveQuoteContractParticipantV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RemoveContractParticipantResponse**](RemoveContractParticipantResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReprocessQuoteContractQuarantineMessageV1

> InboundEmailProcessResponse ReprocessQuoteContractQuarantineMessageV1(ctx, quoteId, messageId).Execute()

Reprocess Quote Contract Quarantine Message

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
	messageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuoteContractsAPI.ReprocessQuoteContractQuarantineMessageV1(context.Background(), quoteId, messageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuoteContractsAPI.ReprocessQuoteContractQuarantineMessageV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReprocessQuoteContractQuarantineMessageV1`: InboundEmailProcessResponse
	fmt.Fprintf(os.Stdout, "Response from `QuoteContractsAPI.ReprocessQuoteContractQuarantineMessageV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 
**messageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReprocessQuoteContractQuarantineMessageV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InboundEmailProcessResponse**](InboundEmailProcessResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadQuoteContractVersionV1

> ContractVersionResponse UploadQuoteContractVersionV1(ctx, quoteId).File(file).Execute()

Upload Quote Contract Version

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
	resp, r, err := apiClient.QuoteContractsAPI.UploadQuoteContractVersionV1(context.Background(), quoteId).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuoteContractsAPI.UploadQuoteContractVersionV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadQuoteContractVersionV1`: ContractVersionResponse
	fmt.Fprintf(os.Stdout, "Response from `QuoteContractsAPI.UploadQuoteContractVersionV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadQuoteContractVersionV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** |  | 

### Return type

[**ContractVersionResponse**](ContractVersionResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

