# \CreditMemosAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyCreditMemoV1**](CreditMemosAPI.md#ApplyCreditMemoV1) | **Post** /v1/arc/credit-memos/{memo_id}/apply | Apply Credit Memo
[**ApproveCreditMemoV1**](CreditMemosAPI.md#ApproveCreditMemoV1) | **Post** /v1/arc/credit-memos/{memo_id}/approve | Approve Credit Memo
[**CreateCreditMemoV1**](CreditMemosAPI.md#CreateCreditMemoV1) | **Post** /v1/arc/credit-memos | Create Credit Memo
[**GetCreditMemoV1**](CreditMemosAPI.md#GetCreditMemoV1) | **Get** /v1/arc/credit-memos/{memo_id} | Get Credit Memo
[**ListCreditMemosV1Get**](CreditMemosAPI.md#ListCreditMemosV1Get) | **Get** /v1/arc/credit-memos | List Credit Memos
[**UpdateCreditMemoV1**](CreditMemosAPI.md#UpdateCreditMemoV1) | **Patch** /v1/arc/credit-memos/{memo_id} | Update Credit Memo
[**VoidCreditMemoV1**](CreditMemosAPI.md#VoidCreditMemoV1) | **Post** /v1/arc/credit-memos/{memo_id}/void | Void Credit Memo



## ApplyCreditMemoV1

> CreditMemoResponse ApplyCreditMemoV1(ctx, memoId).XAPIKey(xAPIKey).CreditMemoApplyRequest(creditMemoApplyRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Apply Credit Memo



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
	memoId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	creditMemoApplyRequest := *openapiclient.NewCreditMemoApplyRequest("InvoiceId_example", int32(123)) // CreditMemoApplyRequest | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreditMemosAPI.ApplyCreditMemoV1(context.Background(), memoId).XAPIKey(xAPIKey).CreditMemoApplyRequest(creditMemoApplyRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreditMemosAPI.ApplyCreditMemoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplyCreditMemoV1`: CreditMemoResponse
	fmt.Fprintf(os.Stdout, "Response from `CreditMemosAPI.ApplyCreditMemoV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memoId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplyCreditMemoV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **creditMemoApplyRequest** | [**CreditMemoApplyRequest**](CreditMemoApplyRequest.md) |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**CreditMemoResponse**](CreditMemoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApproveCreditMemoV1

> CreditMemoResponse ApproveCreditMemoV1(ctx, memoId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Approve Credit Memo



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
	memoId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreditMemosAPI.ApproveCreditMemoV1(context.Background(), memoId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreditMemosAPI.ApproveCreditMemoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApproveCreditMemoV1`: CreditMemoResponse
	fmt.Fprintf(os.Stdout, "Response from `CreditMemosAPI.ApproveCreditMemoV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memoId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApproveCreditMemoV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**CreditMemoResponse**](CreditMemoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCreditMemoV1

> CreditMemoResponse CreateCreditMemoV1(ctx).XAPIKey(xAPIKey).CreditMemoCreate(creditMemoCreate).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Create Credit Memo



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
	xAPIKey := "xAPIKey_example" // string | 
	creditMemoCreate := *openapiclient.NewCreditMemoCreate("AccountId_example", "Reason_example", int32(123), "CurrencyCode_example") // CreditMemoCreate | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreditMemosAPI.CreateCreditMemoV1(context.Background()).XAPIKey(xAPIKey).CreditMemoCreate(creditMemoCreate).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreditMemosAPI.CreateCreditMemoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCreditMemoV1`: CreditMemoResponse
	fmt.Fprintf(os.Stdout, "Response from `CreditMemosAPI.CreateCreditMemoV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCreditMemoV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKey** | **string** |  | 
 **creditMemoCreate** | [**CreditMemoCreate**](CreditMemoCreate.md) |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**CreditMemoResponse**](CreditMemoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCreditMemoV1

> CreditMemoResponse GetCreditMemoV1(ctx, memoId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Get Credit Memo



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
	memoId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreditMemosAPI.GetCreditMemoV1(context.Background(), memoId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreditMemosAPI.GetCreditMemoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCreditMemoV1`: CreditMemoResponse
	fmt.Fprintf(os.Stdout, "Response from `CreditMemosAPI.GetCreditMemoV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memoId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCreditMemoV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**CreditMemoResponse**](CreditMemoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCreditMemosV1Get

> CreditMemoListResponse ListCreditMemosV1Get(ctx).XAPIKey(xAPIKey).AccountId(accountId).Status(status).Page(page).PageSize(pageSize).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

List Credit Memos



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
	xAPIKey := "xAPIKey_example" // string | 
	accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	status := "status_example" // string |  (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 50)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreditMemosAPI.ListCreditMemosV1Get(context.Background()).XAPIKey(xAPIKey).AccountId(accountId).Status(status).Page(page).PageSize(pageSize).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreditMemosAPI.ListCreditMemosV1Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCreditMemosV1Get`: CreditMemoListResponse
	fmt.Fprintf(os.Stdout, "Response from `CreditMemosAPI.ListCreditMemosV1Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCreditMemosV1GetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKey** | **string** |  | 
 **accountId** | **string** |  | 
 **status** | **string** |  | 
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 50]
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**CreditMemoListResponse**](CreditMemoListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCreditMemoV1

> CreditMemoResponse UpdateCreditMemoV1(ctx, memoId).XAPIKey(xAPIKey).CreditMemoUpdate(creditMemoUpdate).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Update Credit Memo



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
	memoId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	creditMemoUpdate := *openapiclient.NewCreditMemoUpdate() // CreditMemoUpdate | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreditMemosAPI.UpdateCreditMemoV1(context.Background(), memoId).XAPIKey(xAPIKey).CreditMemoUpdate(creditMemoUpdate).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreditMemosAPI.UpdateCreditMemoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCreditMemoV1`: CreditMemoResponse
	fmt.Fprintf(os.Stdout, "Response from `CreditMemosAPI.UpdateCreditMemoV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memoId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCreditMemoV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **creditMemoUpdate** | [**CreditMemoUpdate**](CreditMemoUpdate.md) |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**CreditMemoResponse**](CreditMemoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VoidCreditMemoV1

> CreditMemoResponse VoidCreditMemoV1(ctx, memoId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Void Credit Memo



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
	memoId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreditMemosAPI.VoidCreditMemoV1(context.Background(), memoId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreditMemosAPI.VoidCreditMemoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VoidCreditMemoV1`: CreditMemoResponse
	fmt.Fprintf(os.Stdout, "Response from `CreditMemosAPI.VoidCreditMemoV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memoId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVoidCreditMemoV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**CreditMemoResponse**](CreditMemoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

