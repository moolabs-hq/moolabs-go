# \PortalAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTokenEndpoint**](PortalAPI.md#CreateTokenEndpoint) | **Post** /v1/portal/tokens | Create Token Endpoint
[**GetPortalContext**](PortalAPI.md#GetPortalContext) | **Get** /v1/portal/context | Get Portal Context
[**InvalidateTokensEndpoint**](PortalAPI.md#InvalidateTokensEndpoint) | **Post** /v1/portal/tokens/invalidate | Invalidate Tokens Endpoint
[**ListTokensEndpoint**](PortalAPI.md#ListTokensEndpoint) | **Get** /v1/portal/tokens | List Tokens Endpoint



## CreateTokenEndpoint

> interface{} CreateTokenEndpoint(ctx).CreatePortalTokenRequest(createPortalTokenRequest).Execute()

Create Token Endpoint



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
	createPortalTokenRequest := *openapiclient.NewCreatePortalTokenRequest("Subject_example") // CreatePortalTokenRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalAPI.CreateTokenEndpoint(context.Background()).CreatePortalTokenRequest(createPortalTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalAPI.CreateTokenEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTokenEndpoint`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PortalAPI.CreateTokenEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTokenEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPortalTokenRequest** | [**CreatePortalTokenRequest**](CreatePortalTokenRequest.md) |  | 

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


## GetPortalContext

> interface{} GetPortalContext(ctx).Execute()

Get Portal Context



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalAPI.GetPortalContext(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalAPI.GetPortalContext``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortalContext`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PortalAPI.GetPortalContext`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortalContextRequest struct via the builder pattern


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


## InvalidateTokensEndpoint

> interface{} InvalidateTokensEndpoint(ctx).InvalidatePortalTokenRequest(invalidatePortalTokenRequest).Execute()

Invalidate Tokens Endpoint



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
	invalidatePortalTokenRequest := *openapiclient.NewInvalidatePortalTokenRequest() // InvalidatePortalTokenRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalAPI.InvalidateTokensEndpoint(context.Background()).InvalidatePortalTokenRequest(invalidatePortalTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalAPI.InvalidateTokensEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvalidateTokensEndpoint`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PortalAPI.InvalidateTokensEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvalidateTokensEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invalidatePortalTokenRequest** | [**InvalidatePortalTokenRequest**](InvalidatePortalTokenRequest.md) |  | 

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


## ListTokensEndpoint

> interface{} ListTokensEndpoint(ctx).Subject(subject).Execute()

List Tokens Endpoint



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
	subject := "subject_example" // string | Filter by subject (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalAPI.ListTokensEndpoint(context.Background()).Subject(subject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalAPI.ListTokensEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTokensEndpoint`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PortalAPI.ListTokensEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTokensEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subject** | **string** | Filter by subject | 

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

