# \SdkIngestAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IngestSdkSpans**](SdkIngestAPI.md#IngestSdkSpans) | **Post** /api/v1/ingest/batch | Ingest Sdk Batch



## IngestSdkSpans

> SDKBatchIngestResponse IngestSdkSpans(ctx).SDKBatchIngestRequest(sDKBatchIngestRequest).Execute()

Ingest Sdk Batch

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
	sDKBatchIngestRequest := *openapiclient.NewSDKBatchIngestRequest([]openapiclient.SDKEvent{*openapiclient.NewSDKEvent("UsageEventId_example", "RequestId_example", []openapiclient.SDKSpan{*openapiclient.NewSDKSpan("SpanId_example", "Provider_example", "Model_example", "OperationType_example")}, time.Now())}) // SDKBatchIngestRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SdkIngestAPI.IngestSdkSpans(context.Background()).SDKBatchIngestRequest(sDKBatchIngestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SdkIngestAPI.IngestSdkSpans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IngestSdkSpans`: SDKBatchIngestResponse
	fmt.Fprintf(os.Stdout, "Response from `SdkIngestAPI.IngestSdkSpans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIngestSdkSpansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sDKBatchIngestRequest** | [**SDKBatchIngestRequest**](SDKBatchIngestRequest.md) |  | 

### Return type

[**SDKBatchIngestResponse**](SDKBatchIngestResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

