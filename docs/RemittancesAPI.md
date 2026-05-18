# \RemittancesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRemittance**](RemittancesAPI.md#CreateRemittance) | **Post** /v1/arc/remittances | Create Remittance
[**GetUnappliedCashV1**](RemittancesAPI.md#GetUnappliedCashV1) | **Get** /v1/arc/remittances/{remittance_id}/unapplied-cash | Get Unapplied Cash
[**ResolveUnappliedEndpointV1**](RemittancesAPI.md#ResolveUnappliedEndpointV1) | **Post** /v1/arc/remittances/{remittance_id}/unapplied-cash/{uc_id}/resolve | Resolve Unapplied Endpoint



## CreateRemittance

> RemittanceResponse CreateRemittance(ctx).RemittanceCreateRequest(remittanceCreateRequest).ForceCreate(forceCreate).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()

Create Remittance



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
	remittanceCreateRequest := *openapiclient.NewRemittanceCreateRequest("Source_example", int32(123)) // RemittanceCreateRequest | 
	forceCreate := true // bool |  (optional) (default to false)
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemittancesAPI.CreateRemittance(context.Background()).RemittanceCreateRequest(remittanceCreateRequest).ForceCreate(forceCreate).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemittancesAPI.CreateRemittance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRemittance`: RemittanceResponse
	fmt.Fprintf(os.Stdout, "Response from `RemittancesAPI.CreateRemittance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRemittanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **remittanceCreateRequest** | [**RemittanceCreateRequest**](RemittanceCreateRequest.md) |  | 
 **forceCreate** | **bool** |  | [default to false]
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **authorization** | **string** |  | 

### Return type

[**RemittanceResponse**](RemittanceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUnappliedCashV1

> []UnappliedCashResponse GetUnappliedCashV1(ctx, remittanceId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()

Get Unapplied Cash



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
	remittanceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemittancesAPI.GetUnappliedCashV1(context.Background(), remittanceId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemittancesAPI.GetUnappliedCashV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUnappliedCashV1`: []UnappliedCashResponse
	fmt.Fprintf(os.Stdout, "Response from `RemittancesAPI.GetUnappliedCashV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**remittanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUnappliedCashV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **authorization** | **string** |  | 

### Return type

[**[]UnappliedCashResponse**](UnappliedCashResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResolveUnappliedEndpointV1

> UnappliedCashResponse ResolveUnappliedEndpointV1(ctx, remittanceId, ucId).ResolveUnappliedRequest(resolveUnappliedRequest).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()

Resolve Unapplied Endpoint



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
	remittanceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	ucId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	resolveUnappliedRequest := *openapiclient.NewResolveUnappliedRequest("Disposition_example") // ResolveUnappliedRequest | 
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemittancesAPI.ResolveUnappliedEndpointV1(context.Background(), remittanceId, ucId).ResolveUnappliedRequest(resolveUnappliedRequest).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemittancesAPI.ResolveUnappliedEndpointV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResolveUnappliedEndpointV1`: UnappliedCashResponse
	fmt.Fprintf(os.Stdout, "Response from `RemittancesAPI.ResolveUnappliedEndpointV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**remittanceId** | **string** |  | 
**ucId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResolveUnappliedEndpointV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resolveUnappliedRequest** | [**ResolveUnappliedRequest**](ResolveUnappliedRequest.md) |  | 
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **authorization** | **string** |  | 

### Return type

[**UnappliedCashResponse**](UnappliedCashResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

