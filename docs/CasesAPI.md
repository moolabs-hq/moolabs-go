# \CasesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCase**](CasesAPI.md#CreateCase) | **Post** /v1/arc/cases | Create Case
[**DunningHistoryV1**](CasesAPI.md#DunningHistoryV1) | **Get** /v1/arc/cases/{case_id}/dunning-history | Dunning History
[**EscalateCase**](CasesAPI.md#EscalateCase) | **Post** /v1/arc/cases/{case_id}/escalate | Escalate Case
[**FlagDisputedV1**](CasesAPI.md#FlagDisputedV1) | **Post** /v1/arc/cases/{case_id}/invoices/{invoice_id}/flag-disputed | Flag Disputed
[**GetCase**](CasesAPI.md#GetCase) | **Get** /v1/arc/cases/{case_id} | Get Case
[**ListCaseTasks**](CasesAPI.md#ListCaseTasks) | **Get** /v1/arc/cases/{case_id}/tasks | List Case Tasks
[**ListCases**](CasesAPI.md#ListCases) | **Get** /v1/arc/cases | List Cases
[**OpenFromScopeV1Arc**](CasesAPI.md#OpenFromScopeV1Arc) | **Post** /v1/arc/cases/open-from-scope | Open From Scope
[**PauseCase**](CasesAPI.md#PauseCase) | **Post** /v1/arc/cases/{case_id}/pause | Pause Case
[**ReScopeCaseV1**](CasesAPI.md#ReScopeCaseV1) | **Post** /v1/arc/cases/{case_id}/re-scope | Re Scope Case
[**RecomputeStrategyV1**](CasesAPI.md#RecomputeStrategyV1) | **Post** /v1/arc/cases/{case_id}/recompute-strategy | Recompute Strategy
[**ResumeCase**](CasesAPI.md#ResumeCase) | **Post** /v1/arc/cases/{case_id}/resume | Resume Case
[**UpdateCase**](CasesAPI.md#UpdateCase) | **Patch** /v1/arc/cases/{case_id} | Update Case
[**WriteOffCaseV1**](CasesAPI.md#WriteOffCaseV1) | **Post** /v1/arc/cases/{case_id}/write-off | Write Off Case



## CreateCase

> CaseResponse CreateCase(ctx).XAPIKey(xAPIKey).CaseCreate(caseCreate).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Create Case



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
	caseCreate := *openapiclient.NewCaseCreate("AccountId_example", "CustomerId_example") // CaseCreate | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CasesAPI.CreateCase(context.Background()).XAPIKey(xAPIKey).CaseCreate(caseCreate).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CasesAPI.CreateCase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCase`: CaseResponse
	fmt.Fprintf(os.Stdout, "Response from `CasesAPI.CreateCase`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKey** | **string** |  | 
 **caseCreate** | [**CaseCreate**](CaseCreate.md) |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**CaseResponse**](CaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DunningHistoryV1

> interface{} DunningHistoryV1(ctx, caseId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Dunning History



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
	caseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CasesAPI.DunningHistoryV1(context.Background(), caseId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CasesAPI.DunningHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DunningHistoryV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CasesAPI.DunningHistoryV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDunningHistoryV1Request struct via the builder pattern


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


## EscalateCase

> CaseResponse EscalateCase(ctx, caseId).XAPIKey(xAPIKey).CaseEscalateRequest(caseEscalateRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Escalate Case



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
	caseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	caseEscalateRequest := *openapiclient.NewCaseEscalateRequest("Reason_example") // CaseEscalateRequest | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CasesAPI.EscalateCase(context.Background(), caseId).XAPIKey(xAPIKey).CaseEscalateRequest(caseEscalateRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CasesAPI.EscalateCase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EscalateCase`: CaseResponse
	fmt.Fprintf(os.Stdout, "Response from `CasesAPI.EscalateCase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEscalateCaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **caseEscalateRequest** | [**CaseEscalateRequest**](CaseEscalateRequest.md) |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**CaseResponse**](CaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlagDisputedV1

> FlagDisputedResponse FlagDisputedV1(ctx, caseId, invoiceId).XAPIKey(xAPIKey).FlagDisputedRequest(flagDisputedRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Flag Disputed



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
	caseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	invoiceId := "invoiceId_example" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	flagDisputedRequest := *openapiclient.NewFlagDisputedRequest(int32(123), "Description_example") // FlagDisputedRequest | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CasesAPI.FlagDisputedV1(context.Background(), caseId, invoiceId).XAPIKey(xAPIKey).FlagDisputedRequest(flagDisputedRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CasesAPI.FlagDisputedV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlagDisputedV1`: FlagDisputedResponse
	fmt.Fprintf(os.Stdout, "Response from `CasesAPI.FlagDisputedV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  | 
**invoiceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlagDisputedV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAPIKey** | **string** |  | 
 **flagDisputedRequest** | [**FlagDisputedRequest**](FlagDisputedRequest.md) |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**FlagDisputedResponse**](FlagDisputedResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCase

> CaseResponse GetCase(ctx, caseId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Get Case



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
	caseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CasesAPI.GetCase(context.Background(), caseId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CasesAPI.GetCase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCase`: CaseResponse
	fmt.Fprintf(os.Stdout, "Response from `CasesAPI.GetCase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**CaseResponse**](CaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCaseTasks

> TaskListResponse ListCaseTasks(ctx, caseId).XAPIKey(xAPIKey).Page(page).PageSize(pageSize).Status(status).TaskType(taskType).AssignedTo(assignedTo).Search(search).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

List Case Tasks



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
	caseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 50)
	status := "status_example" // string |  (optional)
	taskType := "taskType_example" // string |  (optional)
	assignedTo := "assignedTo_example" // string |  (optional)
	search := "search_example" // string | Substring match on title/description/customer name (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CasesAPI.ListCaseTasks(context.Background(), caseId).XAPIKey(xAPIKey).Page(page).PageSize(pageSize).Status(status).TaskType(taskType).AssignedTo(assignedTo).Search(search).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CasesAPI.ListCaseTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCaseTasks`: TaskListResponse
	fmt.Fprintf(os.Stdout, "Response from `CasesAPI.ListCaseTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCaseTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 50]
 **status** | **string** |  | 
 **taskType** | **string** |  | 
 **assignedTo** | **string** |  | 
 **search** | **string** | Substring match on title/description/customer name | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**TaskListResponse**](TaskListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCases

> CaseListResponse ListCases(ctx).XAPIKey(xAPIKey).Page(page).PageSize(pageSize).Status(status).RiskTier(riskTier).AgingBucket(agingBucket).AccountId(accountId).CustomerId(customerId).Q(q).Search(search).HasPtp(hasPtp).SortBy(sortBy).IncludeInvoices(includeInvoices).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

List Cases



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
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 50)
	status := []string{"Inner_example"} // []string |  (optional)
	riskTier := []string{"Inner_example"} // []string |  (optional)
	agingBucket := []string{"Inner_example"} // []string |  (optional)
	accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	customerId := "customerId_example" // string |  (optional)
	q := "q_example" // string | Backwards-compatible free-text search alias. Newer callers should prefer `search`. (optional)
	search := "search_example" // string | Search by account name, customer ID, case ID, or invoice number (optional)
	hasPtp := true // bool | When true, only cases with at least one open PTP; when false, only cases without. Drives the PTP dashboard tile. (optional)
	sortBy := "sortBy_example" // string | Sort order. Accepted values: created-desc (newest first, default), created-asc (oldest first), amount-desc (highest outstanding first), amount-asc (lowest outstanding first), risk-desc (high risk first). Unknown values fall back to the default active-cases-first ordering. (optional)
	includeInvoices := true // bool | Include embedded invoices in each list item (optional) (default to false)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CasesAPI.ListCases(context.Background()).XAPIKey(xAPIKey).Page(page).PageSize(pageSize).Status(status).RiskTier(riskTier).AgingBucket(agingBucket).AccountId(accountId).CustomerId(customerId).Q(q).Search(search).HasPtp(hasPtp).SortBy(sortBy).IncludeInvoices(includeInvoices).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CasesAPI.ListCases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCases`: CaseListResponse
	fmt.Fprintf(os.Stdout, "Response from `CasesAPI.ListCases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKey** | **string** |  | 
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 50]
 **status** | **[]string** |  | 
 **riskTier** | **[]string** |  | 
 **agingBucket** | **[]string** |  | 
 **accountId** | **string** |  | 
 **customerId** | **string** |  | 
 **q** | **string** | Backwards-compatible free-text search alias. Newer callers should prefer &#x60;search&#x60;. | 
 **search** | **string** | Search by account name, customer ID, case ID, or invoice number | 
 **hasPtp** | **bool** | When true, only cases with at least one open PTP; when false, only cases without. Drives the PTP dashboard tile. | 
 **sortBy** | **string** | Sort order. Accepted values: created-desc (newest first, default), created-asc (oldest first), amount-desc (highest outstanding first), amount-asc (lowest outstanding first), risk-desc (high risk first). Unknown values fall back to the default active-cases-first ordering. | 
 **includeInvoices** | **bool** | Include embedded invoices in each list item | [default to false]
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**CaseListResponse**](CaseListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpenFromScopeV1Arc

> CaseResponse OpenFromScopeV1Arc(ctx).XAPIKey(xAPIKey).OpenFromScopeRequest(openFromScopeRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Open From Scope



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
	openFromScopeRequest := *openapiclient.NewOpenFromScopeRequest("AccountId_example", "CustomerId_example") // OpenFromScopeRequest | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CasesAPI.OpenFromScopeV1Arc(context.Background()).XAPIKey(xAPIKey).OpenFromScopeRequest(openFromScopeRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CasesAPI.OpenFromScopeV1Arc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpenFromScopeV1Arc`: CaseResponse
	fmt.Fprintf(os.Stdout, "Response from `CasesAPI.OpenFromScopeV1Arc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpenFromScopeV1ArcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKey** | **string** |  | 
 **openFromScopeRequest** | [**OpenFromScopeRequest**](OpenFromScopeRequest.md) |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**CaseResponse**](CaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PauseCase

> CaseResponse PauseCase(ctx, caseId).XAPIKey(xAPIKey).CasePauseRequest(casePauseRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Pause Case



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
	caseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	casePauseRequest := *openapiclient.NewCasePauseRequest("Reason_example") // CasePauseRequest | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CasesAPI.PauseCase(context.Background(), caseId).XAPIKey(xAPIKey).CasePauseRequest(casePauseRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CasesAPI.PauseCase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PauseCase`: CaseResponse
	fmt.Fprintf(os.Stdout, "Response from `CasesAPI.PauseCase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPauseCaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **casePauseRequest** | [**CasePauseRequest**](CasePauseRequest.md) |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**CaseResponse**](CaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReScopeCaseV1

> CaseResponse ReScopeCaseV1(ctx, caseId).XAPIKey(xAPIKey).ReScopeRequest(reScopeRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Re Scope Case



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
	caseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	reScopeRequest := *openapiclient.NewReScopeRequest() // ReScopeRequest | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CasesAPI.ReScopeCaseV1(context.Background(), caseId).XAPIKey(xAPIKey).ReScopeRequest(reScopeRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CasesAPI.ReScopeCaseV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReScopeCaseV1`: CaseResponse
	fmt.Fprintf(os.Stdout, "Response from `CasesAPI.ReScopeCaseV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReScopeCaseV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **reScopeRequest** | [**ReScopeRequest**](ReScopeRequest.md) |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**CaseResponse**](CaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecomputeStrategyV1

> CaseResponse RecomputeStrategyV1(ctx, caseId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Recompute Strategy



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
	caseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CasesAPI.RecomputeStrategyV1(context.Background(), caseId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CasesAPI.RecomputeStrategyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecomputeStrategyV1`: CaseResponse
	fmt.Fprintf(os.Stdout, "Response from `CasesAPI.RecomputeStrategyV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecomputeStrategyV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**CaseResponse**](CaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumeCase

> CaseResponse ResumeCase(ctx, caseId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Resume Case



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
	caseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CasesAPI.ResumeCase(context.Background(), caseId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CasesAPI.ResumeCase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResumeCase`: CaseResponse
	fmt.Fprintf(os.Stdout, "Response from `CasesAPI.ResumeCase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeCaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**CaseResponse**](CaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCase

> CaseResponse UpdateCase(ctx, caseId).XAPIKey(xAPIKey).CaseUpdate(caseUpdate).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Update Case



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
	caseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	caseUpdate := *openapiclient.NewCaseUpdate() // CaseUpdate | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CasesAPI.UpdateCase(context.Background(), caseId).XAPIKey(xAPIKey).CaseUpdate(caseUpdate).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CasesAPI.UpdateCase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCase`: CaseResponse
	fmt.Fprintf(os.Stdout, "Response from `CasesAPI.UpdateCase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **caseUpdate** | [**CaseUpdate**](CaseUpdate.md) |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**CaseResponse**](CaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WriteOffCaseV1

> CaseResponse WriteOffCaseV1(ctx, caseId).XAPIKey(xAPIKey).CaseWriteOffRequest(caseWriteOffRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Write Off Case



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
	caseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	caseWriteOffRequest := *openapiclient.NewCaseWriteOffRequest("Reason_example") // CaseWriteOffRequest | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CasesAPI.WriteOffCaseV1(context.Background(), caseId).XAPIKey(xAPIKey).CaseWriteOffRequest(caseWriteOffRequest).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CasesAPI.WriteOffCaseV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WriteOffCaseV1`: CaseResponse
	fmt.Fprintf(os.Stdout, "Response from `CasesAPI.WriteOffCaseV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWriteOffCaseV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **caseWriteOffRequest** | [**CaseWriteOffRequest**](CaseWriteOffRequest.md) |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**CaseResponse**](CaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

