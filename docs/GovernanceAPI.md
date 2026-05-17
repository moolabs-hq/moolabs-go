# \GovernanceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStrategyEndpoint**](GovernanceAPI.md#CreateStrategyEndpoint) | **Post** /v1/arc/governance/strategies | Create Strategy Endpoint
[**DeleteStrategyEndpoint**](GovernanceAPI.md#DeleteStrategyEndpoint) | **Delete** /v1/arc/governance/strategies/{strategy_id} | Delete Strategy Endpoint
[**GetCashAppConfigEndpointV1Arc**](GovernanceAPI.md#GetCashAppConfigEndpointV1Arc) | **Get** /v1/arc/governance/cash-app-config | Get Cash App Config Endpoint
[**GetStrategyEndpoint**](GovernanceAPI.md#GetStrategyEndpoint) | **Get** /v1/arc/governance/strategies/{strategy_id} | Get Strategy Endpoint
[**KillSwitchEndpointV1Arc**](GovernanceAPI.md#KillSwitchEndpointV1Arc) | **Post** /v1/arc/governance/agent-policies/{agent_type}/kill-switch | Kill Switch Endpoint
[**ListEvaluationsEndpoint**](GovernanceAPI.md#ListEvaluationsEndpoint) | **Get** /v1/arc/governance/evaluations | List Evaluations Endpoint
[**ListPoliciesEndpointV1**](GovernanceAPI.md#ListPoliciesEndpointV1) | **Get** /v1/arc/governance/agent-policies | List Policies Endpoint
[**ListStrategiesEndpoint**](GovernanceAPI.md#ListStrategiesEndpoint) | **Get** /v1/arc/governance/strategies | List Strategies Endpoint
[**ShadowModeEndpointV1Arc**](GovernanceAPI.md#ShadowModeEndpointV1Arc) | **Post** /v1/arc/governance/agent-policies/{agent_type}/shadow-mode | Shadow Mode Endpoint
[**StrategyPreviewEndpointV1**](GovernanceAPI.md#StrategyPreviewEndpointV1) | **Post** /v1/arc/governance/strategy-preview | Strategy Preview Endpoint
[**UpdateCashAppConfigEndpointV1Arc**](GovernanceAPI.md#UpdateCashAppConfigEndpointV1Arc) | **Put** /v1/arc/governance/cash-app-config | Update Cash App Config Endpoint
[**UpdatePolicyEndpointV1**](GovernanceAPI.md#UpdatePolicyEndpointV1) | **Put** /v1/arc/governance/agent-policies/{policy_id} | Update Policy Endpoint
[**UpdateStrategyEndpoint**](GovernanceAPI.md#UpdateStrategyEndpoint) | **Put** /v1/arc/governance/strategies/{strategy_id} | Update Strategy Endpoint



## CreateStrategyEndpoint

> interface{} CreateStrategyEndpoint(ctx).XAPIKey(xAPIKey).StrategyCreateRequest(strategyCreateRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Create Strategy Endpoint



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
	strategyCreateRequest := *openapiclient.NewStrategyCreateRequest("Name_example", "RiskTier_example", []openapiclient.StageSchema{*openapiclient.NewStageSchema(int32(123), "Name_example", int32(123), "Channel_example", "TemplateKey_example")}) // StrategyCreateRequest | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GovernanceAPI.CreateStrategyEndpoint(context.Background()).XAPIKey(xAPIKey).StrategyCreateRequest(strategyCreateRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GovernanceAPI.CreateStrategyEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStrategyEndpoint`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `GovernanceAPI.CreateStrategyEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStrategyEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKey** | **string** |  | 
 **strategyCreateRequest** | [**StrategyCreateRequest**](StrategyCreateRequest.md) |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

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


## DeleteStrategyEndpoint

> interface{} DeleteStrategyEndpoint(ctx, strategyId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Delete Strategy Endpoint



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
	strategyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GovernanceAPI.DeleteStrategyEndpoint(context.Background(), strategyId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GovernanceAPI.DeleteStrategyEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteStrategyEndpoint`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `GovernanceAPI.DeleteStrategyEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**strategyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStrategyEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

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


## GetCashAppConfigEndpointV1Arc

> interface{} GetCashAppConfigEndpointV1Arc(ctx).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Get Cash App Config Endpoint



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
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GovernanceAPI.GetCashAppConfigEndpointV1Arc(context.Background()).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GovernanceAPI.GetCashAppConfigEndpointV1Arc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCashAppConfigEndpointV1Arc`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `GovernanceAPI.GetCashAppConfigEndpointV1Arc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCashAppConfigEndpointV1ArcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

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


## GetStrategyEndpoint

> interface{} GetStrategyEndpoint(ctx, strategyId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Get Strategy Endpoint



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
	strategyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GovernanceAPI.GetStrategyEndpoint(context.Background(), strategyId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GovernanceAPI.GetStrategyEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStrategyEndpoint`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `GovernanceAPI.GetStrategyEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**strategyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStrategyEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

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


## KillSwitchEndpointV1Arc

> interface{} KillSwitchEndpointV1Arc(ctx, agentType).XAPIKey(xAPIKey).KillSwitchRequest(killSwitchRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Kill Switch Endpoint



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
	agentType := "agentType_example" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	killSwitchRequest := *openapiclient.NewKillSwitchRequest(false) // KillSwitchRequest | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GovernanceAPI.KillSwitchEndpointV1Arc(context.Background(), agentType).XAPIKey(xAPIKey).KillSwitchRequest(killSwitchRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GovernanceAPI.KillSwitchEndpointV1Arc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KillSwitchEndpointV1Arc`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `GovernanceAPI.KillSwitchEndpointV1Arc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKillSwitchEndpointV1ArcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **killSwitchRequest** | [**KillSwitchRequest**](KillSwitchRequest.md) |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

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


## ListEvaluationsEndpoint

> interface{} ListEvaluationsEndpoint(ctx).XAPIKey(xAPIKey).AgentType(agentType).EvaluationType(evaluationType).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

List Evaluations Endpoint



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
	agentType := "agentType_example" // string |  (optional)
	evaluationType := "evaluationType_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GovernanceAPI.ListEvaluationsEndpoint(context.Background()).XAPIKey(xAPIKey).AgentType(agentType).EvaluationType(evaluationType).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GovernanceAPI.ListEvaluationsEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEvaluationsEndpoint`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `GovernanceAPI.ListEvaluationsEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEvaluationsEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKey** | **string** |  | 
 **agentType** | **string** |  | 
 **evaluationType** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

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


## ListPoliciesEndpointV1

> interface{} ListPoliciesEndpointV1(ctx).XAPIKey(xAPIKey).AgentType(agentType).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

List Policies Endpoint



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
	agentType := "agentType_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GovernanceAPI.ListPoliciesEndpointV1(context.Background()).XAPIKey(xAPIKey).AgentType(agentType).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GovernanceAPI.ListPoliciesEndpointV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPoliciesEndpointV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `GovernanceAPI.ListPoliciesEndpointV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPoliciesEndpointV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKey** | **string** |  | 
 **agentType** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

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


## ListStrategiesEndpoint

> interface{} ListStrategiesEndpoint(ctx).XAPIKey(xAPIKey).RiskTier(riskTier).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

List Strategies Endpoint



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
	riskTier := "riskTier_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GovernanceAPI.ListStrategiesEndpoint(context.Background()).XAPIKey(xAPIKey).RiskTier(riskTier).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GovernanceAPI.ListStrategiesEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListStrategiesEndpoint`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `GovernanceAPI.ListStrategiesEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListStrategiesEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKey** | **string** |  | 
 **riskTier** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

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


## ShadowModeEndpointV1Arc

> interface{} ShadowModeEndpointV1Arc(ctx, agentType).XAPIKey(xAPIKey).ShadowModeRequest(shadowModeRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Shadow Mode Endpoint



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
	agentType := "agentType_example" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	shadowModeRequest := *openapiclient.NewShadowModeRequest(false) // ShadowModeRequest | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GovernanceAPI.ShadowModeEndpointV1Arc(context.Background(), agentType).XAPIKey(xAPIKey).ShadowModeRequest(shadowModeRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GovernanceAPI.ShadowModeEndpointV1Arc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShadowModeEndpointV1Arc`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `GovernanceAPI.ShadowModeEndpointV1Arc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiShadowModeEndpointV1ArcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **shadowModeRequest** | [**ShadowModeRequest**](ShadowModeRequest.md) |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

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


## StrategyPreviewEndpointV1

> interface{} StrategyPreviewEndpointV1(ctx).XAPIKey(xAPIKey).StrategyPreviewRequest(strategyPreviewRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Strategy Preview Endpoint



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
	strategyPreviewRequest := *openapiclient.NewStrategyPreviewRequest("CaseId_example") // StrategyPreviewRequest | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GovernanceAPI.StrategyPreviewEndpointV1(context.Background()).XAPIKey(xAPIKey).StrategyPreviewRequest(strategyPreviewRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GovernanceAPI.StrategyPreviewEndpointV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StrategyPreviewEndpointV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `GovernanceAPI.StrategyPreviewEndpointV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStrategyPreviewEndpointV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKey** | **string** |  | 
 **strategyPreviewRequest** | [**StrategyPreviewRequest**](StrategyPreviewRequest.md) |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

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


## UpdateCashAppConfigEndpointV1Arc

> interface{} UpdateCashAppConfigEndpointV1Arc(ctx).XAPIKey(xAPIKey).CashAppConfigUpdateRequest(cashAppConfigUpdateRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Update Cash App Config Endpoint



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
	cashAppConfigUpdateRequest := *openapiclient.NewCashAppConfigUpdateRequest() // CashAppConfigUpdateRequest | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GovernanceAPI.UpdateCashAppConfigEndpointV1Arc(context.Background()).XAPIKey(xAPIKey).CashAppConfigUpdateRequest(cashAppConfigUpdateRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GovernanceAPI.UpdateCashAppConfigEndpointV1Arc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCashAppConfigEndpointV1Arc`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `GovernanceAPI.UpdateCashAppConfigEndpointV1Arc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCashAppConfigEndpointV1ArcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKey** | **string** |  | 
 **cashAppConfigUpdateRequest** | [**CashAppConfigUpdateRequest**](CashAppConfigUpdateRequest.md) |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

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


## UpdatePolicyEndpointV1

> interface{} UpdatePolicyEndpointV1(ctx, policyId).XAPIKey(xAPIKey).PolicyUpdateRequest(policyUpdateRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Update Policy Endpoint



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
	policyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	policyUpdateRequest := *openapiclient.NewPolicyUpdateRequest() // PolicyUpdateRequest | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GovernanceAPI.UpdatePolicyEndpointV1(context.Background(), policyId).XAPIKey(xAPIKey).PolicyUpdateRequest(policyUpdateRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GovernanceAPI.UpdatePolicyEndpointV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePolicyEndpointV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `GovernanceAPI.UpdatePolicyEndpointV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePolicyEndpointV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **policyUpdateRequest** | [**PolicyUpdateRequest**](PolicyUpdateRequest.md) |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

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


## UpdateStrategyEndpoint

> interface{} UpdateStrategyEndpoint(ctx, strategyId).XAPIKey(xAPIKey).StrategyUpdateRequest(strategyUpdateRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Update Strategy Endpoint



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
	strategyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	strategyUpdateRequest := *openapiclient.NewStrategyUpdateRequest() // StrategyUpdateRequest | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GovernanceAPI.UpdateStrategyEndpoint(context.Background(), strategyId).XAPIKey(xAPIKey).StrategyUpdateRequest(strategyUpdateRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GovernanceAPI.UpdateStrategyEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateStrategyEndpoint`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `GovernanceAPI.UpdateStrategyEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**strategyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStrategyEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **strategyUpdateRequest** | [**StrategyUpdateRequest**](StrategyUpdateRequest.md) |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

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

