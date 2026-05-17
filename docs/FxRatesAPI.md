# \FxRatesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFxRateEndpointV1**](FxRatesAPI.md#CreateFxRateEndpointV1) | **Post** /v1/fx-rates/rates | Create Fx Rate Endpoint
[**CreateValueRecognitionEntryEndpointV1Fx**](FxRatesAPI.md#CreateValueRecognitionEntryEndpointV1Fx) | **Post** /v1/fx-rates/value-recognition/entry | Create Value Recognition Entry Endpoint
[**GetFxRateAtEndpointV1**](FxRatesAPI.md#GetFxRateAtEndpointV1) | **Post** /v1/fx-rates/rates/at | Get Fx Rate At Endpoint
[**ListFxRatesEndpointV1**](FxRatesAPI.md#ListFxRatesEndpointV1) | **Get** /v1/fx-rates/rates | List Fx Rates Endpoint
[**ProcessValueRecognitionV1Fx**](FxRatesAPI.md#ProcessValueRecognitionV1Fx) | **Post** /v1/fx-rates/value-recognition/process | Process Value Recognition



## CreateFxRateEndpointV1

> FxRateResponse CreateFxRateEndpointV1(ctx).CreateFxRateRequest(createFxRateRequest).Execute()

Create Fx Rate Endpoint



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
	createFxRateRequest := *openapiclient.NewCreateFxRateRequest("TenantId_example", "PoolId_example", int32(123), time.Now()) // CreateFxRateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FxRatesAPI.CreateFxRateEndpointV1(context.Background()).CreateFxRateRequest(createFxRateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FxRatesAPI.CreateFxRateEndpointV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFxRateEndpointV1`: FxRateResponse
	fmt.Fprintf(os.Stdout, "Response from `FxRatesAPI.CreateFxRateEndpointV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFxRateEndpointV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFxRateRequest** | [**CreateFxRateRequest**](CreateFxRateRequest.md) |  | 

### Return type

[**FxRateResponse**](FxRateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateValueRecognitionEntryEndpointV1Fx

> interface{} CreateValueRecognitionEntryEndpointV1Fx(ctx).CreateValueRecognitionRequest(createValueRecognitionRequest).Execute()

Create Value Recognition Entry Endpoint



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
	createValueRecognitionRequest := *openapiclient.NewCreateValueRecognitionRequest("TenantId_example", "PoolId_example", "WalletId_example", "SubjectId_example", int32(123), time.Now()) // CreateValueRecognitionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FxRatesAPI.CreateValueRecognitionEntryEndpointV1Fx(context.Background()).CreateValueRecognitionRequest(createValueRecognitionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FxRatesAPI.CreateValueRecognitionEntryEndpointV1Fx``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateValueRecognitionEntryEndpointV1Fx`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FxRatesAPI.CreateValueRecognitionEntryEndpointV1Fx`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateValueRecognitionEntryEndpointV1FxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createValueRecognitionRequest** | [**CreateValueRecognitionRequest**](CreateValueRecognitionRequest.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFxRateAtEndpointV1

> interface{} GetFxRateAtEndpointV1(ctx).GetFxRateRequest(getFxRateRequest).Execute()

Get Fx Rate At Endpoint



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
	getFxRateRequest := *openapiclient.NewGetFxRateRequest("TenantId_example", "PoolId_example", time.Now()) // GetFxRateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FxRatesAPI.GetFxRateAtEndpointV1(context.Background()).GetFxRateRequest(getFxRateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FxRatesAPI.GetFxRateAtEndpointV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFxRateAtEndpointV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FxRatesAPI.GetFxRateAtEndpointV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFxRateAtEndpointV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getFxRateRequest** | [**GetFxRateRequest**](GetFxRateRequest.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFxRatesEndpointV1

> []FxRateResponse ListFxRatesEndpointV1(ctx).TenantId(tenantId).PoolId(poolId).Limit(limit).Execute()

List Fx Rates Endpoint



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Pool identifier
	limit := int32(56) // int32 | Maximum number of rate versions to return (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FxRatesAPI.ListFxRatesEndpointV1(context.Background()).TenantId(tenantId).PoolId(poolId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FxRatesAPI.ListFxRatesEndpointV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFxRatesEndpointV1`: []FxRateResponse
	fmt.Fprintf(os.Stdout, "Response from `FxRatesAPI.ListFxRatesEndpointV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFxRatesEndpointV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Tenant identifier | 
 **poolId** | **string** | Pool identifier | 
 **limit** | **int32** | Maximum number of rate versions to return | [default to 50]

### Return type

[**[]FxRateResponse**](FxRateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProcessValueRecognitionV1Fx

> ValueRecognitionResponse ProcessValueRecognitionV1Fx(ctx).ProcessValueRecognitionRequest(processValueRecognitionRequest).Execute()

Process Value Recognition



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
	processValueRecognitionRequest := *openapiclient.NewProcessValueRecognitionRequest("TenantId_example", "PoolId_example", "JournalEntryId_example") // ProcessValueRecognitionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FxRatesAPI.ProcessValueRecognitionV1Fx(context.Background()).ProcessValueRecognitionRequest(processValueRecognitionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FxRatesAPI.ProcessValueRecognitionV1Fx``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProcessValueRecognitionV1Fx`: ValueRecognitionResponse
	fmt.Fprintf(os.Stdout, "Response from `FxRatesAPI.ProcessValueRecognitionV1Fx`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProcessValueRecognitionV1FxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processValueRecognitionRequest** | [**ProcessValueRecognitionRequest**](ProcessValueRecognitionRequest.md) |  | 

### Return type

[**ValueRecognitionResponse**](ValueRecognitionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

