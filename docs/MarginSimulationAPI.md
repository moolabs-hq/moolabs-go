# \MarginSimulationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompareScenarios**](MarginSimulationAPI.md#CompareScenarios) | **Post** /api/v1/cost/margins/simulate/compare | ACE margin simulation — multi-scenario comparison
[**SimulateMargin**](MarginSimulationAPI.md#SimulateMargin) | **Post** /api/v1/cost/margins/simulate | ACE margin simulation — single scenario



## CompareScenarios

> ScenarioCompareResponse CompareScenarios(ctx).ScenarioRequest(scenarioRequest).Execute()

ACE margin simulation — multi-scenario comparison



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
	scenarioRequest := *openapiclient.NewScenarioRequest([]openapiclient.SimulateMarginRequest{*openapiclient.NewSimulateMarginRequest("TenantId_example", "FeatureKey_example", float32(123), int32(123))}) // ScenarioRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginSimulationAPI.CompareScenarios(context.Background()).ScenarioRequest(scenarioRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginSimulationAPI.CompareScenarios``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompareScenarios`: ScenarioCompareResponse
	fmt.Fprintf(os.Stdout, "Response from `MarginSimulationAPI.CompareScenarios`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompareScenariosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scenarioRequest** | [**ScenarioRequest**](ScenarioRequest.md) |  | 

### Return type

[**ScenarioCompareResponse**](ScenarioCompareResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SimulateMargin

> SimulateMarginResponse SimulateMargin(ctx).SimulateMarginRequest(simulateMarginRequest).Execute()

ACE margin simulation — single scenario



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
	simulateMarginRequest := *openapiclient.NewSimulateMarginRequest("TenantId_example", "FeatureKey_example", float32(123), int32(123)) // SimulateMarginRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginSimulationAPI.SimulateMargin(context.Background()).SimulateMarginRequest(simulateMarginRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginSimulationAPI.SimulateMargin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SimulateMargin`: SimulateMarginResponse
	fmt.Fprintf(os.Stdout, "Response from `MarginSimulationAPI.SimulateMargin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSimulateMarginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **simulateMarginRequest** | [**SimulateMarginRequest**](SimulateMarginRequest.md) |  | 

### Return type

[**SimulateMarginResponse**](SimulateMarginResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

