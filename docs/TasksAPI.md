# \TasksAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompleteTask**](TasksAPI.md#CompleteTask) | **Post** /v1/arc/tasks/{task_id}/complete | Complete Task
[**GetTask**](TasksAPI.md#GetTask) | **Get** /v1/arc/tasks/{task_id} | Get Task
[**ListTasks**](TasksAPI.md#ListTasks) | **Get** /v1/arc/tasks | List Tasks
[**UpdateTask**](TasksAPI.md#UpdateTask) | **Patch** /v1/arc/tasks/{task_id} | Update Task



## CompleteTask

> TaskResponse CompleteTask(ctx, taskId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).TaskCompleteRequest(taskCompleteRequest).Execute()

Complete Task



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
	taskId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	taskCompleteRequest := *openapiclient.NewTaskCompleteRequest() // TaskCompleteRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.CompleteTask(context.Background(), taskId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).TaskCompleteRequest(taskCompleteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.CompleteTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompleteTask`: TaskResponse
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.CompleteTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompleteTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **taskCompleteRequest** | [**TaskCompleteRequest**](TaskCompleteRequest.md) |  | 

### Return type

[**TaskResponse**](TaskResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTask

> TaskResponse GetTask(ctx, taskId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Get Task



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
	taskId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.GetTask(context.Background(), taskId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.GetTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTask`: TaskResponse
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.GetTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**TaskResponse**](TaskResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTasks

> TaskListResponse ListTasks(ctx).XAPIKey(xAPIKey).Page(page).PageSize(pageSize).Status(status).TaskType(taskType).AssignedTo(assignedTo).CaseId(caseId).Search(search).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

List Tasks



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
	status := "status_example" // string |  (optional)
	taskType := "taskType_example" // string |  (optional)
	assignedTo := "assignedTo_example" // string |  (optional)
	caseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	search := "search_example" // string | Substring match on title/description/customer name (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.ListTasks(context.Background()).XAPIKey(xAPIKey).Page(page).PageSize(pageSize).Status(status).TaskType(taskType).AssignedTo(assignedTo).CaseId(caseId).Search(search).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.ListTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTasks`: TaskListResponse
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.ListTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKey** | **string** |  | 
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 50]
 **status** | **string** |  | 
 **taskType** | **string** |  | 
 **assignedTo** | **string** |  | 
 **caseId** | **string** |  | 
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


## UpdateTask

> TaskResponse UpdateTask(ctx, taskId).XAPIKey(xAPIKey).TaskUpdate(taskUpdate).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Update Task



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
	taskId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	taskUpdate := *openapiclient.NewTaskUpdate() // TaskUpdate | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.UpdateTask(context.Background(), taskId).XAPIKey(xAPIKey).TaskUpdate(taskUpdate).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.UpdateTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTask`: TaskResponse
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.UpdateTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **taskUpdate** | [**TaskUpdate**](TaskUpdate.md) |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**TaskResponse**](TaskResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

