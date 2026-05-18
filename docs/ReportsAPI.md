# \ReportsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReportEndpoint**](ReportsAPI.md#CreateReportEndpoint) | **Post** /v1/arc/reports | Create Report Endpoint
[**DownloadReportEndpoint**](ReportsAPI.md#DownloadReportEndpoint) | **Get** /v1/arc/reports/{report_id}/download | Download Report Endpoint



## CreateReportEndpoint

> ReportCreateResponse CreateReportEndpoint(ctx).ReportCreateRequest(reportCreateRequest).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).XUserId(xUserId).XArcRoles(xArcRoles).Execute()

Create Report Endpoint

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
	reportCreateRequest := *openapiclient.NewReportCreateRequest("ReportType_example", "Format_example") // ReportCreateRequest | 
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)
	xUserId := "xUserId_example" // string |  (optional)
	xArcRoles := "xArcRoles_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.CreateReportEndpoint(context.Background()).ReportCreateRequest(reportCreateRequest).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).XUserId(xUserId).XArcRoles(xArcRoles).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.CreateReportEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateReportEndpoint`: ReportCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.CreateReportEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReportEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reportCreateRequest** | [**ReportCreateRequest**](ReportCreateRequest.md) |  | 
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **authorization** | **string** |  | 
 **xUserId** | **string** |  | 
 **xArcRoles** | **string** |  | 

### Return type

[**ReportCreateResponse**](ReportCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadReportEndpoint

> interface{} DownloadReportEndpoint(ctx, reportId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()

Download Report Endpoint

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
	reportId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.DownloadReportEndpoint(context.Background(), reportId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.DownloadReportEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadReportEndpoint`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.DownloadReportEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadReportEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **authorization** | **string** |  | 

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

