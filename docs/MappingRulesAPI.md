# \MappingRulesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRuleApi**](MappingRulesAPI.md#CreateRuleApi) | **Post** /api/v1/mapping-rules | Create Rule
[**DisableRuleApi**](MappingRulesAPI.md#DisableRuleApi) | **Delete** /api/v1/mapping-rules/{rule_id} | Disable Rule
[**DryRunApiV1**](MappingRulesAPI.md#DryRunApiV1) | **Post** /api/v1/mapping-rules/dry-run | Dry Run
[**EnableTier3ApiV1**](MappingRulesAPI.md#EnableTier3ApiV1) | **Post** /api/v1/mapping-rules/enable-tier3 | Enable Tier3
[**ListRulesApi**](MappingRulesAPI.md#ListRulesApi) | **Get** /api/v1/mapping-rules | List Rules
[**UpdateRuleApi**](MappingRulesAPI.md#UpdateRuleApi) | **Put** /api/v1/mapping-rules/{rule_id} | Update Rule



## CreateRuleApi

> MappingRuleResponse CreateRuleApi(ctx).XTenantID(xTenantID).MappingRuleCreate(mappingRuleCreate).Execute()

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
	xTenantID := "xTenantID_example" // string | 
	mappingRuleCreate := *openapiclient.NewMappingRuleCreate("RuleName_example", "SourceType_example", "SourceField_example", "TargetField_example") // MappingRuleCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MappingRulesAPI.CreateRuleApi(context.Background()).XTenantID(xTenantID).MappingRuleCreate(mappingRuleCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MappingRulesAPI.CreateRuleApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRuleApi`: MappingRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `MappingRulesAPI.CreateRuleApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRuleApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTenantID** | **string** |  | 
 **mappingRuleCreate** | [**MappingRuleCreate**](MappingRuleCreate.md) |  | 

### Return type

[**MappingRuleResponse**](MappingRuleResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableRuleApi

> MappingRuleResponse DisableRuleApi(ctx, ruleId).XTenantID(xTenantID).Execute()

Disable Rule



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
	xTenantID := "xTenantID_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MappingRulesAPI.DisableRuleApi(context.Background(), ruleId).XTenantID(xTenantID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MappingRulesAPI.DisableRuleApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisableRuleApi`: MappingRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `MappingRulesAPI.DisableRuleApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableRuleApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTenantID** | **string** |  | 

### Return type

[**MappingRuleResponse**](MappingRuleResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DryRunApiV1

> DryRunResponse DryRunApiV1(ctx).XTenantID(xTenantID).DryRunRequest(dryRunRequest).Execute()

Dry Run



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
	xTenantID := "xTenantID_example" // string | 
	dryRunRequest := *openapiclient.NewDryRunRequest(map[string]interface{}{"key": interface{}(123)}) // DryRunRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MappingRulesAPI.DryRunApiV1(context.Background()).XTenantID(xTenantID).DryRunRequest(dryRunRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MappingRulesAPI.DryRunApiV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DryRunApiV1`: DryRunResponse
	fmt.Fprintf(os.Stdout, "Response from `MappingRulesAPI.DryRunApiV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDryRunApiV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTenantID** | **string** |  | 
 **dryRunRequest** | [**DryRunRequest**](DryRunRequest.md) |  | 

### Return type

[**DryRunResponse**](DryRunResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableTier3ApiV1

> EnableTier3Response EnableTier3ApiV1(ctx).XTenantID(xTenantID).Execute()

Enable Tier3



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
	xTenantID := "xTenantID_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MappingRulesAPI.EnableTier3ApiV1(context.Background()).XTenantID(xTenantID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MappingRulesAPI.EnableTier3ApiV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnableTier3ApiV1`: EnableTier3Response
	fmt.Fprintf(os.Stdout, "Response from `MappingRulesAPI.EnableTier3ApiV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnableTier3ApiV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTenantID** | **string** |  | 

### Return type

[**EnableTier3Response**](EnableTier3Response.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRulesApi

> []MappingRuleResponse ListRulesApi(ctx).XTenantID(xTenantID).EnabledOnly(enabledOnly).IncludeAudit(includeAudit).Execute()

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
	xTenantID := "xTenantID_example" // string | 
	enabledOnly := true // bool | Return only enabled rules (optional) (default to true)
	includeAudit := true // bool | Include last 5 audit entries per rule (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MappingRulesAPI.ListRulesApi(context.Background()).XTenantID(xTenantID).EnabledOnly(enabledOnly).IncludeAudit(includeAudit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MappingRulesAPI.ListRulesApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRulesApi`: []MappingRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `MappingRulesAPI.ListRulesApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRulesApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTenantID** | **string** |  | 
 **enabledOnly** | **bool** | Return only enabled rules | [default to true]
 **includeAudit** | **bool** | Include last 5 audit entries per rule | [default to false]

### Return type

[**[]MappingRuleResponse**](MappingRuleResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRuleApi

> MappingRuleResponse UpdateRuleApi(ctx, ruleId).XTenantID(xTenantID).MappingRuleUpdate(mappingRuleUpdate).Execute()

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
	ruleId := "ruleId_example" // string | 
	xTenantID := "xTenantID_example" // string | 
	mappingRuleUpdate := *openapiclient.NewMappingRuleUpdate() // MappingRuleUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MappingRulesAPI.UpdateRuleApi(context.Background(), ruleId).XTenantID(xTenantID).MappingRuleUpdate(mappingRuleUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MappingRulesAPI.UpdateRuleApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRuleApi`: MappingRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `MappingRulesAPI.UpdateRuleApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRuleApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTenantID** | **string** |  | 
 **mappingRuleUpdate** | [**MappingRuleUpdate**](MappingRuleUpdate.md) |  | 

### Return type

[**MappingRuleResponse**](MappingRuleResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

