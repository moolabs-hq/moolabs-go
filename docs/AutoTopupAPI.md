# \AutoTopupAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckTriggerV1**](AutoTopupAPI.md#CheckTriggerV1) | **Post** /v1/auto-topup/rules/{rule_id}/check | Check Trigger
[**CreateRuleV1**](AutoTopupAPI.md#CreateRuleV1) | **Post** /v1/auto-topup/rules | Create Rule
[**DeleteRuleV1**](AutoTopupAPI.md#DeleteRuleV1) | **Delete** /v1/auto-topup/rules/{rule_id} | Delete Rule
[**GetActivityV1**](AutoTopupAPI.md#GetActivityV1) | **Get** /v1/auto-topup/activity | Get Activity
[**GetRuleV1**](AutoTopupAPI.md#GetRuleV1) | **Get** /v1/auto-topup/rules/{rule_id} | Get Rule
[**ListRulesV1**](AutoTopupAPI.md#ListRulesV1) | **Get** /v1/auto-topup/rules | List Rules
[**PatchRuleV1**](AutoTopupAPI.md#PatchRuleV1) | **Patch** /v1/auto-topup/rules/{rule_id} | Patch Rule
[**PaymentSucceededV1**](AutoTopupAPI.md#PaymentSucceededV1) | **Post** /v1/auto-topup/payments/succeeded | Payment Succeeded
[**UpdateRuleV1**](AutoTopupAPI.md#UpdateRuleV1) | **Put** /v1/auto-topup/rules/{rule_id} | Update Rule



## CheckTriggerV1

> TriggerResponse CheckTriggerV1(ctx, ruleId).CheckTriggerRequest(checkTriggerRequest).Execute()

Check Trigger



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
	ruleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Auto top-up rule ID
	checkTriggerRequest := *openapiclient.NewCheckTriggerRequest() // CheckTriggerRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoTopupAPI.CheckTriggerV1(context.Background(), ruleId).CheckTriggerRequest(checkTriggerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTopupAPI.CheckTriggerV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckTriggerV1`: TriggerResponse
	fmt.Fprintf(os.Stdout, "Response from `AutoTopupAPI.CheckTriggerV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** | Auto top-up rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckTriggerV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **checkTriggerRequest** | [**CheckTriggerRequest**](CheckTriggerRequest.md) |  | 

### Return type

[**TriggerResponse**](TriggerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRuleV1

> AutoTopupRuleResponse CreateRuleV1(ctx).CreateAutoTopupRuleRequest(createAutoTopupRuleRequest).Execute()

Create Rule



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
	createAutoTopupRuleRequest := *openapiclient.NewCreateAutoTopupRuleRequest("TenantId_example", "PoolId_example", "WalletId_example", "TriggerType_example", "TriggerValue_example", int32(123)) // CreateAutoTopupRuleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoTopupAPI.CreateRuleV1(context.Background()).CreateAutoTopupRuleRequest(createAutoTopupRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTopupAPI.CreateRuleV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRuleV1`: AutoTopupRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `AutoTopupAPI.CreateRuleV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRuleV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAutoTopupRuleRequest** | [**CreateAutoTopupRuleRequest**](CreateAutoTopupRuleRequest.md) |  | 

### Return type

[**AutoTopupRuleResponse**](AutoTopupRuleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRuleV1

> interface{} DeleteRuleV1(ctx, ruleId).TenantId(tenantId).PoolId(poolId).WalletId(walletId).Execute()

Delete Rule



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
	ruleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Auto top-up rule ID
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional tenant identifier for validation (optional)
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional pool identifier for validation (optional)
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional wallet identifier for validation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoTopupAPI.DeleteRuleV1(context.Background(), ruleId).TenantId(tenantId).PoolId(poolId).WalletId(walletId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTopupAPI.DeleteRuleV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRuleV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoTopupAPI.DeleteRuleV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** | Auto top-up rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRuleV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** | Optional tenant identifier for validation | 
 **poolId** | **string** | Optional pool identifier for validation | 
 **walletId** | **string** | Optional wallet identifier for validation | 

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


## GetActivityV1

> AutoTopupActivityResponse GetActivityV1(ctx).TenantId(tenantId).PoolId(poolId).WalletId(walletId).Limit(limit).Offset(offset).Execute()

Get Activity



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Wallet identifier
	limit := int32(56) // int32 | Maximum number of items to return (optional) (default to 50)
	offset := int32(56) // int32 | Offset for pagination (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoTopupAPI.GetActivityV1(context.Background()).TenantId(tenantId).PoolId(poolId).WalletId(walletId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTopupAPI.GetActivityV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActivityV1`: AutoTopupActivityResponse
	fmt.Fprintf(os.Stdout, "Response from `AutoTopupAPI.GetActivityV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetActivityV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Tenant identifier | 
 **poolId** | **string** | Pool identifier | 
 **walletId** | **string** | Wallet identifier | 
 **limit** | **int32** | Maximum number of items to return | [default to 50]
 **offset** | **int32** | Offset for pagination | [default to 0]

### Return type

[**AutoTopupActivityResponse**](AutoTopupActivityResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuleV1

> AutoTopupRuleResponse GetRuleV1(ctx, ruleId).TenantId(tenantId).PoolId(poolId).Execute()

Get Rule



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
	ruleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Auto top-up rule ID
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional tenant identifier for validation (optional)
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional pool identifier for validation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoTopupAPI.GetRuleV1(context.Background(), ruleId).TenantId(tenantId).PoolId(poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTopupAPI.GetRuleV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRuleV1`: AutoTopupRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `AutoTopupAPI.GetRuleV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** | Auto top-up rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** | Optional tenant identifier for validation | 
 **poolId** | **string** | Optional pool identifier for validation | 

### Return type

[**AutoTopupRuleResponse**](AutoTopupRuleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRulesV1

> ListAutoTopupRulesResponse ListRulesV1(ctx).TenantId(tenantId).PoolId(poolId).WalletId(walletId).Execute()

List Rules



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional wallet identifier to filter by (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoTopupAPI.ListRulesV1(context.Background()).TenantId(tenantId).PoolId(poolId).WalletId(walletId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTopupAPI.ListRulesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRulesV1`: ListAutoTopupRulesResponse
	fmt.Fprintf(os.Stdout, "Response from `AutoTopupAPI.ListRulesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRulesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Tenant identifier | 
 **poolId** | **string** | Pool identifier | 
 **walletId** | **string** | Optional wallet identifier to filter by | 

### Return type

[**ListAutoTopupRulesResponse**](ListAutoTopupRulesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchRuleV1

> AutoTopupRuleResponse PatchRuleV1(ctx, ruleId).TenantId(tenantId).PoolId(poolId).WalletId(walletId).UpdateAutoTopupRuleRequest(updateAutoTopupRuleRequest).Execute()

Patch Rule



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
	ruleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Auto top-up rule ID
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional tenant identifier for validation (optional)
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional pool identifier for validation (optional)
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional wallet identifier for validation (optional)
	updateAutoTopupRuleRequest := *openapiclient.NewUpdateAutoTopupRuleRequest() // UpdateAutoTopupRuleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoTopupAPI.PatchRuleV1(context.Background(), ruleId).TenantId(tenantId).PoolId(poolId).WalletId(walletId).UpdateAutoTopupRuleRequest(updateAutoTopupRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTopupAPI.PatchRuleV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchRuleV1`: AutoTopupRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `AutoTopupAPI.PatchRuleV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** | Auto top-up rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchRuleV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** | Optional tenant identifier for validation | 
 **poolId** | **string** | Optional pool identifier for validation | 
 **walletId** | **string** | Optional wallet identifier for validation | 
 **updateAutoTopupRuleRequest** | [**UpdateAutoTopupRuleRequest**](UpdateAutoTopupRuleRequest.md) |  | 

### Return type

[**AutoTopupRuleResponse**](AutoTopupRuleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentSucceededV1

> PaymentSucceededResponse PaymentSucceededV1(ctx).PaymentSucceededRequest(paymentSucceededRequest).Execute()

Payment Succeeded



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
	paymentSucceededRequest := *openapiclient.NewPaymentSucceededRequest("TenantId_example", "PoolId_example", "WalletId_example", "PaymentId_example", int32(123)) // PaymentSucceededRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoTopupAPI.PaymentSucceededV1(context.Background()).PaymentSucceededRequest(paymentSucceededRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTopupAPI.PaymentSucceededV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentSucceededV1`: PaymentSucceededResponse
	fmt.Fprintf(os.Stdout, "Response from `AutoTopupAPI.PaymentSucceededV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentSucceededV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentSucceededRequest** | [**PaymentSucceededRequest**](PaymentSucceededRequest.md) |  | 

### Return type

[**PaymentSucceededResponse**](PaymentSucceededResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRuleV1

> AutoTopupRuleResponse UpdateRuleV1(ctx, ruleId).TenantId(tenantId).PoolId(poolId).WalletId(walletId).UpdateAutoTopupRuleRequest(updateAutoTopupRuleRequest).Execute()

Update Rule



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
	ruleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Auto top-up rule ID
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional tenant identifier for validation (optional)
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional pool identifier for validation (optional)
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional wallet identifier for validation (optional)
	updateAutoTopupRuleRequest := *openapiclient.NewUpdateAutoTopupRuleRequest() // UpdateAutoTopupRuleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoTopupAPI.UpdateRuleV1(context.Background(), ruleId).TenantId(tenantId).PoolId(poolId).WalletId(walletId).UpdateAutoTopupRuleRequest(updateAutoTopupRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTopupAPI.UpdateRuleV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRuleV1`: AutoTopupRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `AutoTopupAPI.UpdateRuleV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** | Auto top-up rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRuleV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** | Optional tenant identifier for validation | 
 **poolId** | **string** | Optional pool identifier for validation | 
 **walletId** | **string** | Optional wallet identifier for validation | 
 **updateAutoTopupRuleRequest** | [**UpdateAutoTopupRuleRequest**](UpdateAutoTopupRuleRequest.md) |  | 

### Return type

[**AutoTopupRuleResponse**](AutoTopupRuleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

