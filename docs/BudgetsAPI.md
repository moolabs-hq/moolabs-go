# \BudgetsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcknowledgeAlert**](BudgetsAPI.md#AcknowledgeAlert) | **Post** /api/v1/budgets/{budget_id}/alerts/{alert_id}/acknowledge | Acknowledge an alert
[**CreateBudget**](BudgetsAPI.md#CreateBudget) | **Post** /api/v1/budgets | Create a budget
[**GetBudget**](BudgetsAPI.md#GetBudget) | **Get** /api/v1/budgets/{budget_id} | Get budget with current utilization
[**GetBudgetAlerts**](BudgetsAPI.md#GetBudgetAlerts) | **Get** /api/v1/budgets/{budget_id}/alerts | Get alerts for this budget
[**GetBudgetStatus**](BudgetsAPI.md#GetBudgetStatus) | **Get** /api/v1/budgets/status | Get all budget utilization summary
[**ListBudgets**](BudgetsAPI.md#ListBudgets) | **Get** /api/v1/budgets | List budgets for tenant
[**UpdateBudget**](BudgetsAPI.md#UpdateBudget) | **Put** /api/v1/budgets/{budget_id} | Update a budget



## AcknowledgeAlert

> BudgetAlertResponse AcknowledgeAlert(ctx, budgetId, alertId).Execute()

Acknowledge an alert

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
	budgetId := "budgetId_example" // string | 
	alertId := "alertId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BudgetsAPI.AcknowledgeAlert(context.Background(), budgetId, alertId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BudgetsAPI.AcknowledgeAlert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AcknowledgeAlert`: BudgetAlertResponse
	fmt.Fprintf(os.Stdout, "Response from `BudgetsAPI.AcknowledgeAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**budgetId** | **string** |  | 
**alertId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcknowledgeAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BudgetAlertResponse**](BudgetAlertResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBudget

> BudgetResponse CreateBudget(ctx).BudgetCreate(budgetCreate).Execute()

Create a budget



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
	budgetCreate := *openapiclient.NewBudgetCreate("BudgetName_example", "ScopeType_example", *openapiclient.NewBudgetAmount()) // BudgetCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BudgetsAPI.CreateBudget(context.Background()).BudgetCreate(budgetCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BudgetsAPI.CreateBudget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBudget`: BudgetResponse
	fmt.Fprintf(os.Stdout, "Response from `BudgetsAPI.CreateBudget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **budgetCreate** | [**BudgetCreate**](BudgetCreate.md) |  | 

### Return type

[**BudgetResponse**](BudgetResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBudget

> BudgetWithUtilization GetBudget(ctx, budgetId).PeriodStart(periodStart).PeriodEnd(periodEnd).Execute()

Get budget with current utilization



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
	budgetId := "budgetId_example" // string | 
	periodStart := time.Now() // time.Time |  (optional)
	periodEnd := time.Now() // time.Time |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BudgetsAPI.GetBudget(context.Background(), budgetId).PeriodStart(periodStart).PeriodEnd(periodEnd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BudgetsAPI.GetBudget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBudget`: BudgetWithUtilization
	fmt.Fprintf(os.Stdout, "Response from `BudgetsAPI.GetBudget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**budgetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **periodStart** | **time.Time** |  | 
 **periodEnd** | **time.Time** |  | 

### Return type

[**BudgetWithUtilization**](BudgetWithUtilization.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBudgetAlerts

> []BudgetAlertResponse GetBudgetAlerts(ctx, budgetId).Acknowledged(acknowledged).Limit(limit).Offset(offset).Execute()

Get alerts for this budget



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
	budgetId := "budgetId_example" // string | 
	acknowledged := true // bool | Filter by acknowledgement state (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BudgetsAPI.GetBudgetAlerts(context.Background(), budgetId).Acknowledged(acknowledged).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BudgetsAPI.GetBudgetAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBudgetAlerts`: []BudgetAlertResponse
	fmt.Fprintf(os.Stdout, "Response from `BudgetsAPI.GetBudgetAlerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**budgetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBudgetAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acknowledged** | **bool** | Filter by acknowledgement state | 
 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** |  | [default to 0]

### Return type

[**[]BudgetAlertResponse**](BudgetAlertResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBudgetStatus

> []BudgetStatusItem GetBudgetStatus(ctx).PeriodStart(periodStart).PeriodEnd(periodEnd).Execute()

Get all budget utilization summary



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
	periodStart := time.Now() // time.Time |  (optional)
	periodEnd := time.Now() // time.Time |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BudgetsAPI.GetBudgetStatus(context.Background()).PeriodStart(periodStart).PeriodEnd(periodEnd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BudgetsAPI.GetBudgetStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBudgetStatus`: []BudgetStatusItem
	fmt.Fprintf(os.Stdout, "Response from `BudgetsAPI.GetBudgetStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBudgetStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **periodStart** | **time.Time** |  | 
 **periodEnd** | **time.Time** |  | 

### Return type

[**[]BudgetStatusItem**](BudgetStatusItem.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBudgets

> []BudgetResponse ListBudgets(ctx).IsActive(isActive).ScopeType(scopeType).Execute()

List budgets for tenant



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
	scopeType := "scopeType_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BudgetsAPI.ListBudgets(context.Background()).IsActive(isActive).ScopeType(scopeType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BudgetsAPI.ListBudgets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBudgets`: []BudgetResponse
	fmt.Fprintf(os.Stdout, "Response from `BudgetsAPI.ListBudgets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBudgetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isActive** | **bool** |  | 
 **scopeType** | **string** |  | 

### Return type

[**[]BudgetResponse**](BudgetResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBudget

> BudgetResponse UpdateBudget(ctx, budgetId).BudgetUpdate(budgetUpdate).Execute()

Update a budget



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
	budgetId := "budgetId_example" // string | 
	budgetUpdate := *openapiclient.NewBudgetUpdate() // BudgetUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BudgetsAPI.UpdateBudget(context.Background(), budgetId).BudgetUpdate(budgetUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BudgetsAPI.UpdateBudget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBudget`: BudgetResponse
	fmt.Fprintf(os.Stdout, "Response from `BudgetsAPI.UpdateBudget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**budgetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **budgetUpdate** | [**BudgetUpdate**](BudgetUpdate.md) |  | 

### Return type

[**BudgetResponse**](BudgetResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader), [HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

