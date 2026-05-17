# \MoometerProxyAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProxyOpenmeter**](MoometerProxyAPI.md#ProxyOpenmeter) | **Get** /v1/moometer/{path} | Proxy Openmeter
[**ProxyOpenmeter_0**](MoometerProxyAPI.md#ProxyOpenmeter_0) | **Put** /v1/moometer/{path} | Proxy Openmeter
[**ProxyOpenmeter_1**](MoometerProxyAPI.md#ProxyOpenmeter_1) | **Post** /v1/moometer/{path} | Proxy Openmeter
[**ProxyOpenmeter_2**](MoometerProxyAPI.md#ProxyOpenmeter_2) | **Delete** /v1/moometer/{path} | Proxy Openmeter
[**ProxyOpenmeter_3**](MoometerProxyAPI.md#ProxyOpenmeter_3) | **Patch** /v1/moometer/{path} | Proxy Openmeter



## ProxyOpenmeter

> interface{} ProxyOpenmeter(ctx, path).Execute()

Proxy Openmeter



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
	path := "path_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MoometerProxyAPI.ProxyOpenmeter(context.Background(), path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MoometerProxyAPI.ProxyOpenmeter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProxyOpenmeter`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `MoometerProxyAPI.ProxyOpenmeter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProxyOpenmeterRequest struct via the builder pattern


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


## ProxyOpenmeter_0

> interface{} ProxyOpenmeter_0(ctx, path).Execute()

Proxy Openmeter



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
	path := "path_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MoometerProxyAPI.ProxyOpenmeter_0(context.Background(), path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MoometerProxyAPI.ProxyOpenmeter_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProxyOpenmeter_0`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `MoometerProxyAPI.ProxyOpenmeter_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProxyOpenmeter_1Request struct via the builder pattern


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


## ProxyOpenmeter_1

> interface{} ProxyOpenmeter_1(ctx, path).Execute()

Proxy Openmeter



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
	path := "path_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MoometerProxyAPI.ProxyOpenmeter_1(context.Background(), path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MoometerProxyAPI.ProxyOpenmeter_1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProxyOpenmeter_1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `MoometerProxyAPI.ProxyOpenmeter_1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProxyOpenmeter_2Request struct via the builder pattern


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


## ProxyOpenmeter_2

> interface{} ProxyOpenmeter_2(ctx, path).Execute()

Proxy Openmeter



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
	path := "path_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MoometerProxyAPI.ProxyOpenmeter_2(context.Background(), path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MoometerProxyAPI.ProxyOpenmeter_2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProxyOpenmeter_2`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `MoometerProxyAPI.ProxyOpenmeter_2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProxyOpenmeter_3Request struct via the builder pattern


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


## ProxyOpenmeter_3

> interface{} ProxyOpenmeter_3(ctx, path).Execute()

Proxy Openmeter



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
	path := "path_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MoometerProxyAPI.ProxyOpenmeter_3(context.Background(), path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MoometerProxyAPI.ProxyOpenmeter_3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProxyOpenmeter_3`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `MoometerProxyAPI.ProxyOpenmeter_3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProxyOpenmeter_4Request struct via the builder pattern


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

