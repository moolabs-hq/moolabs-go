# \SubjectsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSubject**](SubjectsAPI.md#DeleteSubject) | **Delete** /api/v1/subjects/{subjectIdOrKey} | Delete subject
[**GetSubject**](SubjectsAPI.md#GetSubject) | **Get** /api/v1/subjects/{subjectIdOrKey} | Get subject
[**ListSubjects**](SubjectsAPI.md#ListSubjects) | **Get** /api/v1/subjects | List subjects
[**UpsertSubject**](SubjectsAPI.md#UpsertSubject) | **Post** /api/v1/subjects | Upsert subject



## DeleteSubject

> DeleteSubject(ctx, subjectIdOrKey).Execute()

Delete subject



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
	subjectIdOrKey := "subjectIdOrKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SubjectsAPI.DeleteSubject(context.Background(), subjectIdOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubjectsAPI.DeleteSubject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectIdOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubject

> Subject GetSubject(ctx, subjectIdOrKey).Execute()

Get subject



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
	subjectIdOrKey := "subjectIdOrKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubjectsAPI.GetSubject(context.Background(), subjectIdOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubjectsAPI.GetSubject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubject`: Subject
	fmt.Fprintf(os.Stdout, "Response from `SubjectsAPI.GetSubject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectIdOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Subject**](Subject.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubjects

> []Subject ListSubjects(ctx).Execute()

List subjects



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubjectsAPI.ListSubjects(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubjectsAPI.ListSubjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSubjects`: []Subject
	fmt.Fprintf(os.Stdout, "Response from `SubjectsAPI.ListSubjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSubjectsRequest struct via the builder pattern


### Return type

[**[]Subject**](Subject.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertSubject

> []Subject UpsertSubject(ctx).SubjectUpsert(subjectUpsert).Execute()

Upsert subject



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
	subjectUpsert := []openapiclient.SubjectUpsert{*openapiclient.NewSubjectUpsert("customer-db-id-123")} // []SubjectUpsert | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubjectsAPI.UpsertSubject(context.Background()).SubjectUpsert(subjectUpsert).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubjectsAPI.UpsertSubject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpsertSubject`: []Subject
	fmt.Fprintf(os.Stdout, "Response from `SubjectsAPI.UpsertSubject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpsertSubjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subjectUpsert** | [**[]SubjectUpsert**](SubjectUpsert.md) |  | 

### Return type

[**[]Subject**](Subject.md)

### Authorization

[CloudTokenAuth](../README.md#CloudTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

