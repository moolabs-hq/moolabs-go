# \AccountTeamAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountTeamMember**](AccountTeamAPI.md#CreateAccountTeamMember) | **Post** /v1/arc/accounts/{account_id}/team | Create Account Team Member
[**DeleteAccountTeamMember**](AccountTeamAPI.md#DeleteAccountTeamMember) | **Delete** /v1/arc/accounts/{account_id}/team/{member_id} | Delete Account Team Member
[**ListAccountTeam**](AccountTeamAPI.md#ListAccountTeam) | **Get** /v1/arc/accounts/{account_id}/team | List Account Team
[**UpdateAccountTeamMember**](AccountTeamAPI.md#UpdateAccountTeamMember) | **Patch** /v1/arc/accounts/{account_id}/team/{member_id} | Update Account Team Member



## CreateAccountTeamMember

> AccountTeamMemberOut CreateAccountTeamMember(ctx, accountId).XAPIKey(xAPIKey).AccountTeamMemberIn(accountTeamMemberIn).XTenantId(xTenantId).XOrgId(xOrgId).XUserId(xUserId).Execute()

Create Account Team Member



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
	accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	accountTeamMemberIn := *openapiclient.NewAccountTeamMemberIn(openapiclient.AccountTeamRole("ae"), "Name_example") // AccountTeamMemberIn | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	xUserId := "xUserId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountTeamAPI.CreateAccountTeamMember(context.Background(), accountId).XAPIKey(xAPIKey).AccountTeamMemberIn(accountTeamMemberIn).XTenantId(xTenantId).XOrgId(xOrgId).XUserId(xUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountTeamAPI.CreateAccountTeamMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccountTeamMember`: AccountTeamMemberOut
	fmt.Fprintf(os.Stdout, "Response from `AccountTeamAPI.CreateAccountTeamMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountTeamMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **accountTeamMemberIn** | [**AccountTeamMemberIn**](AccountTeamMemberIn.md) |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **xUserId** | **string** |  | 

### Return type

[**AccountTeamMemberOut**](AccountTeamMemberOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccountTeamMember

> DeleteAccountTeamMember(ctx, accountId, memberId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Delete Account Team Member

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
	accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	memberId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccountTeamAPI.DeleteAccountTeamMember(context.Background(), accountId, memberId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountTeamAPI.DeleteAccountTeamMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 
**memberId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountTeamMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccountTeam

> []AccountTeamMemberOut ListAccountTeam(ctx, accountId).XAPIKey(xAPIKey).NotifyOnly(notifyOnly).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

List Account Team



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
	accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	notifyOnly := true // bool |  (optional) (default to false)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountTeamAPI.ListAccountTeam(context.Background(), accountId).XAPIKey(xAPIKey).NotifyOnly(notifyOnly).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountTeamAPI.ListAccountTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAccountTeam`: []AccountTeamMemberOut
	fmt.Fprintf(os.Stdout, "Response from `AccountTeamAPI.ListAccountTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAccountTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **notifyOnly** | **bool** |  | [default to false]
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**[]AccountTeamMemberOut**](AccountTeamMemberOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountTeamMember

> AccountTeamMemberOut UpdateAccountTeamMember(ctx, accountId, memberId).XAPIKey(xAPIKey).AccountTeamMemberPatch(accountTeamMemberPatch).XTenantId(xTenantId).XOrgId(xOrgId).Execute()

Update Account Team Member

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
	accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	memberId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string | 
	accountTeamMemberPatch := *openapiclient.NewAccountTeamMemberPatch() // AccountTeamMemberPatch | 
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountTeamAPI.UpdateAccountTeamMember(context.Background(), accountId, memberId).XAPIKey(xAPIKey).AccountTeamMemberPatch(accountTeamMemberPatch).XTenantId(xTenantId).XOrgId(xOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountTeamAPI.UpdateAccountTeamMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAccountTeamMember`: AccountTeamMemberOut
	fmt.Fprintf(os.Stdout, "Response from `AccountTeamAPI.UpdateAccountTeamMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 
**memberId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountTeamMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAPIKey** | **string** |  | 
 **accountTeamMemberPatch** | [**AccountTeamMemberPatch**](AccountTeamMemberPatch.md) |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 

### Return type

[**AccountTeamMemberOut**](AccountTeamMemberOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

