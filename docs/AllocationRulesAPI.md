# \AllocationRulesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAllocationRuleApi**](AllocationRulesAPI.md#CreateAllocationRuleApi) | **Post** /api/v1/allocation-rules | Create an allocation rule
[**DeleteAllocationRuleApi**](AllocationRulesAPI.md#DeleteAllocationRuleApi) | **Delete** /api/v1/allocation-rules/{rule_id} | Soft-delete allocation rule (is_active&#x3D;false)
[**GetAllocationRuleApi**](AllocationRulesAPI.md#GetAllocationRuleApi) | **Get** /api/v1/allocation-rules/{rule_id} | Get allocation rule detail
[**ListAllocationRulesApi**](AllocationRulesAPI.md#ListAllocationRulesApi) | **Get** /api/v1/allocation-rules | List allocation rules for tenant
[**TriggerReallocateApi**](AllocationRulesAPI.md#TriggerReallocateApi) | **Post** /api/v1/allocation-rules/{rule_id}/reallocate | Manually trigger reallocation for a period
[**UpdateAllocationRuleApi**](AllocationRulesAPI.md#UpdateAllocationRuleApi) | **Put** /api/v1/allocation-rules/{rule_id} | Update allocation rule (triggers retroactive recompute signal)



## CreateAllocationRuleApi

> AllocationRuleResponse CreateAllocationRuleApi(ctx).AllocationRuleCreate(allocationRuleCreate).Execute()

Create an allocation rule



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
	allocationRuleCreate := *openapiclient.NewAllocationRuleCreate("Name_example", "RuleType_example", "TargetType_example") // AllocationRuleCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AllocationRulesAPI.CreateAllocationRuleApi(context.Background()).AllocationRuleCreate(allocationRuleCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllocationRulesAPI.CreateAllocationRuleApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAllocationRuleApi`: AllocationRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `AllocationRulesAPI.CreateAllocationRuleApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAllocationRuleApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **allocationRuleCreate** | [**AllocationRuleCreate**](AllocationRuleCreate.md) |  | 

### Return type

[**AllocationRuleResponse**](AllocationRuleResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAllocationRuleApi

> DeleteAllocationRuleApi(ctx, ruleId).Execute()

Soft-delete allocation rule (is_active=false)



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
	ruleId := "ruleId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AllocationRulesAPI.DeleteAllocationRuleApi(context.Background(), ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllocationRulesAPI.DeleteAllocationRuleApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllocationRuleApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllocationRuleApi

> AllocationRuleResponse GetAllocationRuleApi(ctx, ruleId).Execute()

Get allocation rule detail



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
	ruleId := "ruleId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AllocationRulesAPI.GetAllocationRuleApi(context.Background(), ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllocationRulesAPI.GetAllocationRuleApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllocationRuleApi`: AllocationRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `AllocationRulesAPI.GetAllocationRuleApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllocationRuleApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AllocationRuleResponse**](AllocationRuleResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllocationRulesApi

> []AllocationRuleResponse ListAllocationRulesApi(ctx).IsActive(isActive).RuleType(ruleType).Execute()

List allocation rules for tenant



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
	isActive := true // bool |  (optional)
	ruleType := "ruleType_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AllocationRulesAPI.ListAllocationRulesApi(context.Background()).IsActive(isActive).RuleType(ruleType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllocationRulesAPI.ListAllocationRulesApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAllocationRulesApi`: []AllocationRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `AllocationRulesAPI.ListAllocationRulesApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAllocationRulesApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isActive** | **bool** |  | 
 **ruleType** | **string** |  | 

### Return type

[**[]AllocationRuleResponse**](AllocationRuleResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TriggerReallocateApi

> ReallocateResponse TriggerReallocateApi(ctx, ruleId).ReallocateRequest(reallocateRequest).Execute()

Manually trigger reallocation for a period



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
	ruleId := "ruleId_example" // string | 
	reallocateRequest := *openapiclient.NewReallocateRequest(time.Now(), time.Now()) // ReallocateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AllocationRulesAPI.TriggerReallocateApi(context.Background(), ruleId).ReallocateRequest(reallocateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllocationRulesAPI.TriggerReallocateApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TriggerReallocateApi`: ReallocateResponse
	fmt.Fprintf(os.Stdout, "Response from `AllocationRulesAPI.TriggerReallocateApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTriggerReallocateApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reallocateRequest** | [**ReallocateRequest**](ReallocateRequest.md) |  | 

### Return type

[**ReallocateResponse**](ReallocateResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAllocationRuleApi

> AllocationRuleResponse UpdateAllocationRuleApi(ctx, ruleId).AllocationRuleUpdate(allocationRuleUpdate).Execute()

Update allocation rule (triggers retroactive recompute signal)



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
	ruleId := "ruleId_example" // string | 
	allocationRuleUpdate := *openapiclient.NewAllocationRuleUpdate() // AllocationRuleUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AllocationRulesAPI.UpdateAllocationRuleApi(context.Background(), ruleId).AllocationRuleUpdate(allocationRuleUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllocationRulesAPI.UpdateAllocationRuleApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAllocationRuleApi`: AllocationRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `AllocationRulesAPI.UpdateAllocationRuleApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAllocationRuleApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **allocationRuleUpdate** | [**AllocationRuleUpdate**](AllocationRuleUpdate.md) |  | 

### Return type

[**AllocationRuleResponse**](AllocationRuleResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

