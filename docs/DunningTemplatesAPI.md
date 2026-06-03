# \DunningTemplatesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDisclosurePolicyV1ArcDunning**](DunningTemplatesAPI.md#GetDisclosurePolicyV1ArcDunning) | **Get** /v1/arc/dunning-template-disclosure-policy | Get Disclosure Policy
[**GetDunningTemplateRegistryV1Arc**](DunningTemplatesAPI.md#GetDunningTemplateRegistryV1Arc) | **Get** /v1/arc/dunning-template-registry | Get Dunning Template Registry
[**GetDunningTemplateVersionV1**](DunningTemplatesAPI.md#GetDunningTemplateVersionV1) | **Get** /v1/arc/dunning-templates/{template_key}/versions/{version_id} | Get Dunning Template Version
[**GetPaymentInstructionsV1ArcDunning**](DunningTemplatesAPI.md#GetPaymentInstructionsV1ArcDunning) | **Get** /v1/arc/dunning-template-payment-instructions | Get Payment Instructions
[**GetProviderReadinessV1ArcDunning**](DunningTemplatesAPI.md#GetProviderReadinessV1ArcDunning) | **Get** /v1/arc/dunning-template-provider-readiness | Get Provider Readiness
[**GetTemplateKeysV1Arc**](DunningTemplatesAPI.md#GetTemplateKeysV1Arc) | **Get** /v1/arc/dunning-template-keys | Get Template Keys
[**ListDunningTemplateVersionsV1**](DunningTemplatesAPI.md#ListDunningTemplateVersionsV1) | **Get** /v1/arc/dunning-templates/{template_key}/versions | List Dunning Template Versions
[**LockDunningTemplateDraftV1**](DunningTemplatesAPI.md#LockDunningTemplateDraftV1) | **Post** /v1/arc/dunning-templates/{template_key}/drafts/{version_id}/lock | Lock Dunning Template Draft
[**PostArchiveTemplateKeyV1Arc**](DunningTemplatesAPI.md#PostArchiveTemplateKeyV1Arc) | **Post** /v1/arc/dunning-template-keys/{template_key}/archive | Post Archive Template Key
[**PostTemplateKeyV1Arc**](DunningTemplatesAPI.md#PostTemplateKeyV1Arc) | **Post** /v1/arc/dunning-template-keys | Post Template Key
[**PostTemplateTestSendV1ArcDunning**](DunningTemplatesAPI.md#PostTemplateTestSendV1ArcDunning) | **Post** /v1/arc/dunning-template-test-sends | Post Template Test Send
[**PreviewDunningTemplateV1**](DunningTemplatesAPI.md#PreviewDunningTemplateV1) | **Post** /v1/arc/dunning-templates/{template_key}/preview | Preview Dunning Template
[**PublishDunningTemplateV1**](DunningTemplatesAPI.md#PublishDunningTemplateV1) | **Post** /v1/arc/dunning-templates/{template_key}/publish | Publish Dunning Template
[**PutDisclosurePolicyV1ArcDunning**](DunningTemplatesAPI.md#PutDisclosurePolicyV1ArcDunning) | **Put** /v1/arc/dunning-template-disclosure-policy | Put Disclosure Policy
[**PutPaymentInstructionsV1ArcDunning**](DunningTemplatesAPI.md#PutPaymentInstructionsV1ArcDunning) | **Put** /v1/arc/dunning-template-payment-instructions | Put Payment Instructions
[**SaveDunningTemplateDraftV1**](DunningTemplatesAPI.md#SaveDunningTemplateDraftV1) | **Post** /v1/arc/dunning-templates/{template_key}/drafts | Save Dunning Template Draft
[**UnlockDunningTemplateDraftV1**](DunningTemplatesAPI.md#UnlockDunningTemplateDraftV1) | **Post** /v1/arc/dunning-templates/{template_key}/drafts/{version_id}/unlock | Unlock Dunning Template Draft
[**ValidateDunningTemplateV1**](DunningTemplatesAPI.md#ValidateDunningTemplateV1) | **Post** /v1/arc/dunning-templates/{template_key}/validate | Validate Dunning Template



## GetDisclosurePolicyV1ArcDunning

> DisclosurePolicyResponse GetDisclosurePolicyV1ArcDunning(ctx).XArcProxySecret(xArcProxySecret).XArcFeatureFlags(xArcFeatureFlags).XArcRoles(xArcRoles).XArcActorUserId(xArcActorUserId).XArcActorDisplay(xArcActorDisplay).XArcActorEmail(xArcActorEmail).XArcSource(xArcSource).XOrgId(xOrgId).XChatSessionId(xChatSessionId).XRequestId(xRequestId).XCorrelationID(xCorrelationID).XAPIKey(xAPIKey).XTenantId(xTenantId).Authorization(authorization).Execute()

Get Disclosure Policy

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
	xArcProxySecret := "xArcProxySecret_example" // string |  (optional)
	xArcFeatureFlags := "xArcFeatureFlags_example" // string |  (optional)
	xArcRoles := "xArcRoles_example" // string |  (optional)
	xArcActorUserId := "xArcActorUserId_example" // string |  (optional)
	xArcActorDisplay := "xArcActorDisplay_example" // string |  (optional)
	xArcActorEmail := "xArcActorEmail_example" // string |  (optional)
	xArcSource := "xArcSource_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	xChatSessionId := "xChatSessionId_example" // string |  (optional)
	xRequestId := "xRequestId_example" // string |  (optional)
	xCorrelationID := "xCorrelationID_example" // string |  (optional)
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DunningTemplatesAPI.GetDisclosurePolicyV1ArcDunning(context.Background()).XArcProxySecret(xArcProxySecret).XArcFeatureFlags(xArcFeatureFlags).XArcRoles(xArcRoles).XArcActorUserId(xArcActorUserId).XArcActorDisplay(xArcActorDisplay).XArcActorEmail(xArcActorEmail).XArcSource(xArcSource).XOrgId(xOrgId).XChatSessionId(xChatSessionId).XRequestId(xRequestId).XCorrelationID(xCorrelationID).XAPIKey(xAPIKey).XTenantId(xTenantId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DunningTemplatesAPI.GetDisclosurePolicyV1ArcDunning``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDisclosurePolicyV1ArcDunning`: DisclosurePolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `DunningTemplatesAPI.GetDisclosurePolicyV1ArcDunning`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDisclosurePolicyV1ArcDunningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xArcProxySecret** | **string** |  | 
 **xArcFeatureFlags** | **string** |  | 
 **xArcRoles** | **string** |  | 
 **xArcActorUserId** | **string** |  | 
 **xArcActorDisplay** | **string** |  | 
 **xArcActorEmail** | **string** |  | 
 **xArcSource** | **string** |  | 
 **xOrgId** | **string** |  | 
 **xChatSessionId** | **string** |  | 
 **xRequestId** | **string** |  | 
 **xCorrelationID** | **string** |  | 
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **authorization** | **string** |  | 

### Return type

[**DisclosurePolicyResponse**](DisclosurePolicyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDunningTemplateRegistryV1Arc

> interface{} GetDunningTemplateRegistryV1Arc(ctx).XArcProxySecret(xArcProxySecret).XArcFeatureFlags(xArcFeatureFlags).XArcRoles(xArcRoles).XArcActorUserId(xArcActorUserId).XArcActorDisplay(xArcActorDisplay).XArcActorEmail(xArcActorEmail).XArcSource(xArcSource).XOrgId(xOrgId).XChatSessionId(xChatSessionId).XRequestId(xRequestId).XCorrelationID(xCorrelationID).XAPIKey(xAPIKey).XTenantId(xTenantId).Authorization(authorization).Execute()

Get Dunning Template Registry

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
	xArcProxySecret := "xArcProxySecret_example" // string |  (optional)
	xArcFeatureFlags := "xArcFeatureFlags_example" // string |  (optional)
	xArcRoles := "xArcRoles_example" // string |  (optional)
	xArcActorUserId := "xArcActorUserId_example" // string |  (optional)
	xArcActorDisplay := "xArcActorDisplay_example" // string |  (optional)
	xArcActorEmail := "xArcActorEmail_example" // string |  (optional)
	xArcSource := "xArcSource_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	xChatSessionId := "xChatSessionId_example" // string |  (optional)
	xRequestId := "xRequestId_example" // string |  (optional)
	xCorrelationID := "xCorrelationID_example" // string |  (optional)
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DunningTemplatesAPI.GetDunningTemplateRegistryV1Arc(context.Background()).XArcProxySecret(xArcProxySecret).XArcFeatureFlags(xArcFeatureFlags).XArcRoles(xArcRoles).XArcActorUserId(xArcActorUserId).XArcActorDisplay(xArcActorDisplay).XArcActorEmail(xArcActorEmail).XArcSource(xArcSource).XOrgId(xOrgId).XChatSessionId(xChatSessionId).XRequestId(xRequestId).XCorrelationID(xCorrelationID).XAPIKey(xAPIKey).XTenantId(xTenantId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DunningTemplatesAPI.GetDunningTemplateRegistryV1Arc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDunningTemplateRegistryV1Arc`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DunningTemplatesAPI.GetDunningTemplateRegistryV1Arc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDunningTemplateRegistryV1ArcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xArcProxySecret** | **string** |  | 
 **xArcFeatureFlags** | **string** |  | 
 **xArcRoles** | **string** |  | 
 **xArcActorUserId** | **string** |  | 
 **xArcActorDisplay** | **string** |  | 
 **xArcActorEmail** | **string** |  | 
 **xArcSource** | **string** |  | 
 **xOrgId** | **string** |  | 
 **xChatSessionId** | **string** |  | 
 **xRequestId** | **string** |  | 
 **xCorrelationID** | **string** |  | 
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
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


## GetDunningTemplateVersionV1

> TemplateVersionResponse GetDunningTemplateVersionV1(ctx, templateKey, versionId).XArcProxySecret(xArcProxySecret).XArcFeatureFlags(xArcFeatureFlags).XArcRoles(xArcRoles).XArcActorUserId(xArcActorUserId).XArcActorDisplay(xArcActorDisplay).XArcActorEmail(xArcActorEmail).XArcSource(xArcSource).XOrgId(xOrgId).XChatSessionId(xChatSessionId).XRequestId(xRequestId).XCorrelationID(xCorrelationID).XAPIKey(xAPIKey).XTenantId(xTenantId).Authorization(authorization).Execute()

Get Dunning Template Version

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
	templateKey := "templateKey_example" // string | 
	versionId := "versionId_example" // string | 
	xArcProxySecret := "xArcProxySecret_example" // string |  (optional)
	xArcFeatureFlags := "xArcFeatureFlags_example" // string |  (optional)
	xArcRoles := "xArcRoles_example" // string |  (optional)
	xArcActorUserId := "xArcActorUserId_example" // string |  (optional)
	xArcActorDisplay := "xArcActorDisplay_example" // string |  (optional)
	xArcActorEmail := "xArcActorEmail_example" // string |  (optional)
	xArcSource := "xArcSource_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	xChatSessionId := "xChatSessionId_example" // string |  (optional)
	xRequestId := "xRequestId_example" // string |  (optional)
	xCorrelationID := "xCorrelationID_example" // string |  (optional)
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DunningTemplatesAPI.GetDunningTemplateVersionV1(context.Background(), templateKey, versionId).XArcProxySecret(xArcProxySecret).XArcFeatureFlags(xArcFeatureFlags).XArcRoles(xArcRoles).XArcActorUserId(xArcActorUserId).XArcActorDisplay(xArcActorDisplay).XArcActorEmail(xArcActorEmail).XArcSource(xArcSource).XOrgId(xOrgId).XChatSessionId(xChatSessionId).XRequestId(xRequestId).XCorrelationID(xCorrelationID).XAPIKey(xAPIKey).XTenantId(xTenantId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DunningTemplatesAPI.GetDunningTemplateVersionV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDunningTemplateVersionV1`: TemplateVersionResponse
	fmt.Fprintf(os.Stdout, "Response from `DunningTemplatesAPI.GetDunningTemplateVersionV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateKey** | **string** |  | 
**versionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDunningTemplateVersionV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xArcProxySecret** | **string** |  | 
 **xArcFeatureFlags** | **string** |  | 
 **xArcRoles** | **string** |  | 
 **xArcActorUserId** | **string** |  | 
 **xArcActorDisplay** | **string** |  | 
 **xArcActorEmail** | **string** |  | 
 **xArcSource** | **string** |  | 
 **xOrgId** | **string** |  | 
 **xChatSessionId** | **string** |  | 
 **xRequestId** | **string** |  | 
 **xCorrelationID** | **string** |  | 
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **authorization** | **string** |  | 

### Return type

[**TemplateVersionResponse**](TemplateVersionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentInstructionsV1ArcDunning

> PaymentInstructionsResponse GetPaymentInstructionsV1ArcDunning(ctx).XArcProxySecret(xArcProxySecret).XArcFeatureFlags(xArcFeatureFlags).XArcRoles(xArcRoles).XArcActorUserId(xArcActorUserId).XArcActorDisplay(xArcActorDisplay).XArcActorEmail(xArcActorEmail).XArcSource(xArcSource).XOrgId(xOrgId).XChatSessionId(xChatSessionId).XRequestId(xRequestId).XCorrelationID(xCorrelationID).XAPIKey(xAPIKey).XTenantId(xTenantId).Authorization(authorization).Execute()

Get Payment Instructions

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
	xArcProxySecret := "xArcProxySecret_example" // string |  (optional)
	xArcFeatureFlags := "xArcFeatureFlags_example" // string |  (optional)
	xArcRoles := "xArcRoles_example" // string |  (optional)
	xArcActorUserId := "xArcActorUserId_example" // string |  (optional)
	xArcActorDisplay := "xArcActorDisplay_example" // string |  (optional)
	xArcActorEmail := "xArcActorEmail_example" // string |  (optional)
	xArcSource := "xArcSource_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	xChatSessionId := "xChatSessionId_example" // string |  (optional)
	xRequestId := "xRequestId_example" // string |  (optional)
	xCorrelationID := "xCorrelationID_example" // string |  (optional)
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DunningTemplatesAPI.GetPaymentInstructionsV1ArcDunning(context.Background()).XArcProxySecret(xArcProxySecret).XArcFeatureFlags(xArcFeatureFlags).XArcRoles(xArcRoles).XArcActorUserId(xArcActorUserId).XArcActorDisplay(xArcActorDisplay).XArcActorEmail(xArcActorEmail).XArcSource(xArcSource).XOrgId(xOrgId).XChatSessionId(xChatSessionId).XRequestId(xRequestId).XCorrelationID(xCorrelationID).XAPIKey(xAPIKey).XTenantId(xTenantId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DunningTemplatesAPI.GetPaymentInstructionsV1ArcDunning``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPaymentInstructionsV1ArcDunning`: PaymentInstructionsResponse
	fmt.Fprintf(os.Stdout, "Response from `DunningTemplatesAPI.GetPaymentInstructionsV1ArcDunning`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentInstructionsV1ArcDunningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xArcProxySecret** | **string** |  | 
 **xArcFeatureFlags** | **string** |  | 
 **xArcRoles** | **string** |  | 
 **xArcActorUserId** | **string** |  | 
 **xArcActorDisplay** | **string** |  | 
 **xArcActorEmail** | **string** |  | 
 **xArcSource** | **string** |  | 
 **xOrgId** | **string** |  | 
 **xChatSessionId** | **string** |  | 
 **xRequestId** | **string** |  | 
 **xCorrelationID** | **string** |  | 
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **authorization** | **string** |  | 

### Return type

[**PaymentInstructionsResponse**](PaymentInstructionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProviderReadinessV1ArcDunning

> ProviderReadinessResponse GetProviderReadinessV1ArcDunning(ctx).XArcProxySecret(xArcProxySecret).XArcFeatureFlags(xArcFeatureFlags).XArcRoles(xArcRoles).XArcActorUserId(xArcActorUserId).XArcActorDisplay(xArcActorDisplay).XArcActorEmail(xArcActorEmail).XArcSource(xArcSource).XOrgId(xOrgId).XChatSessionId(xChatSessionId).XRequestId(xRequestId).XCorrelationID(xCorrelationID).XAPIKey(xAPIKey).XTenantId(xTenantId).Authorization(authorization).Execute()

Get Provider Readiness

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
	xArcProxySecret := "xArcProxySecret_example" // string |  (optional)
	xArcFeatureFlags := "xArcFeatureFlags_example" // string |  (optional)
	xArcRoles := "xArcRoles_example" // string |  (optional)
	xArcActorUserId := "xArcActorUserId_example" // string |  (optional)
	xArcActorDisplay := "xArcActorDisplay_example" // string |  (optional)
	xArcActorEmail := "xArcActorEmail_example" // string |  (optional)
	xArcSource := "xArcSource_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	xChatSessionId := "xChatSessionId_example" // string |  (optional)
	xRequestId := "xRequestId_example" // string |  (optional)
	xCorrelationID := "xCorrelationID_example" // string |  (optional)
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DunningTemplatesAPI.GetProviderReadinessV1ArcDunning(context.Background()).XArcProxySecret(xArcProxySecret).XArcFeatureFlags(xArcFeatureFlags).XArcRoles(xArcRoles).XArcActorUserId(xArcActorUserId).XArcActorDisplay(xArcActorDisplay).XArcActorEmail(xArcActorEmail).XArcSource(xArcSource).XOrgId(xOrgId).XChatSessionId(xChatSessionId).XRequestId(xRequestId).XCorrelationID(xCorrelationID).XAPIKey(xAPIKey).XTenantId(xTenantId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DunningTemplatesAPI.GetProviderReadinessV1ArcDunning``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProviderReadinessV1ArcDunning`: ProviderReadinessResponse
	fmt.Fprintf(os.Stdout, "Response from `DunningTemplatesAPI.GetProviderReadinessV1ArcDunning`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderReadinessV1ArcDunningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xArcProxySecret** | **string** |  | 
 **xArcFeatureFlags** | **string** |  | 
 **xArcRoles** | **string** |  | 
 **xArcActorUserId** | **string** |  | 
 **xArcActorDisplay** | **string** |  | 
 **xArcActorEmail** | **string** |  | 
 **xArcSource** | **string** |  | 
 **xOrgId** | **string** |  | 
 **xChatSessionId** | **string** |  | 
 **xRequestId** | **string** |  | 
 **xCorrelationID** | **string** |  | 
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **authorization** | **string** |  | 

### Return type

[**ProviderReadinessResponse**](ProviderReadinessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplateKeysV1Arc

> TemplateKeyListResponse GetTemplateKeysV1Arc(ctx).XArcProxySecret(xArcProxySecret).XArcFeatureFlags(xArcFeatureFlags).XArcRoles(xArcRoles).XArcActorUserId(xArcActorUserId).XArcActorDisplay(xArcActorDisplay).XArcActorEmail(xArcActorEmail).XArcSource(xArcSource).XOrgId(xOrgId).XChatSessionId(xChatSessionId).XRequestId(xRequestId).XCorrelationID(xCorrelationID).XAPIKey(xAPIKey).XTenantId(xTenantId).Authorization(authorization).Execute()

Get Template Keys

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
	xArcProxySecret := "xArcProxySecret_example" // string |  (optional)
	xArcFeatureFlags := "xArcFeatureFlags_example" // string |  (optional)
	xArcRoles := "xArcRoles_example" // string |  (optional)
	xArcActorUserId := "xArcActorUserId_example" // string |  (optional)
	xArcActorDisplay := "xArcActorDisplay_example" // string |  (optional)
	xArcActorEmail := "xArcActorEmail_example" // string |  (optional)
	xArcSource := "xArcSource_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	xChatSessionId := "xChatSessionId_example" // string |  (optional)
	xRequestId := "xRequestId_example" // string |  (optional)
	xCorrelationID := "xCorrelationID_example" // string |  (optional)
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DunningTemplatesAPI.GetTemplateKeysV1Arc(context.Background()).XArcProxySecret(xArcProxySecret).XArcFeatureFlags(xArcFeatureFlags).XArcRoles(xArcRoles).XArcActorUserId(xArcActorUserId).XArcActorDisplay(xArcActorDisplay).XArcActorEmail(xArcActorEmail).XArcSource(xArcSource).XOrgId(xOrgId).XChatSessionId(xChatSessionId).XRequestId(xRequestId).XCorrelationID(xCorrelationID).XAPIKey(xAPIKey).XTenantId(xTenantId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DunningTemplatesAPI.GetTemplateKeysV1Arc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplateKeysV1Arc`: TemplateKeyListResponse
	fmt.Fprintf(os.Stdout, "Response from `DunningTemplatesAPI.GetTemplateKeysV1Arc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateKeysV1ArcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xArcProxySecret** | **string** |  | 
 **xArcFeatureFlags** | **string** |  | 
 **xArcRoles** | **string** |  | 
 **xArcActorUserId** | **string** |  | 
 **xArcActorDisplay** | **string** |  | 
 **xArcActorEmail** | **string** |  | 
 **xArcSource** | **string** |  | 
 **xOrgId** | **string** |  | 
 **xChatSessionId** | **string** |  | 
 **xRequestId** | **string** |  | 
 **xCorrelationID** | **string** |  | 
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **authorization** | **string** |  | 

### Return type

[**TemplateKeyListResponse**](TemplateKeyListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDunningTemplateVersionsV1

> TemplateVersionListResponse ListDunningTemplateVersionsV1(ctx, templateKey).XArcProxySecret(xArcProxySecret).XArcFeatureFlags(xArcFeatureFlags).XArcRoles(xArcRoles).XArcActorUserId(xArcActorUserId).XArcActorDisplay(xArcActorDisplay).XArcActorEmail(xArcActorEmail).XArcSource(xArcSource).XOrgId(xOrgId).XChatSessionId(xChatSessionId).XRequestId(xRequestId).XCorrelationID(xCorrelationID).XAPIKey(xAPIKey).XTenantId(xTenantId).Authorization(authorization).Execute()

List Dunning Template Versions

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
	templateKey := "templateKey_example" // string | 
	xArcProxySecret := "xArcProxySecret_example" // string |  (optional)
	xArcFeatureFlags := "xArcFeatureFlags_example" // string |  (optional)
	xArcRoles := "xArcRoles_example" // string |  (optional)
	xArcActorUserId := "xArcActorUserId_example" // string |  (optional)
	xArcActorDisplay := "xArcActorDisplay_example" // string |  (optional)
	xArcActorEmail := "xArcActorEmail_example" // string |  (optional)
	xArcSource := "xArcSource_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	xChatSessionId := "xChatSessionId_example" // string |  (optional)
	xRequestId := "xRequestId_example" // string |  (optional)
	xCorrelationID := "xCorrelationID_example" // string |  (optional)
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DunningTemplatesAPI.ListDunningTemplateVersionsV1(context.Background(), templateKey).XArcProxySecret(xArcProxySecret).XArcFeatureFlags(xArcFeatureFlags).XArcRoles(xArcRoles).XArcActorUserId(xArcActorUserId).XArcActorDisplay(xArcActorDisplay).XArcActorEmail(xArcActorEmail).XArcSource(xArcSource).XOrgId(xOrgId).XChatSessionId(xChatSessionId).XRequestId(xRequestId).XCorrelationID(xCorrelationID).XAPIKey(xAPIKey).XTenantId(xTenantId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DunningTemplatesAPI.ListDunningTemplateVersionsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDunningTemplateVersionsV1`: TemplateVersionListResponse
	fmt.Fprintf(os.Stdout, "Response from `DunningTemplatesAPI.ListDunningTemplateVersionsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDunningTemplateVersionsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xArcProxySecret** | **string** |  | 
 **xArcFeatureFlags** | **string** |  | 
 **xArcRoles** | **string** |  | 
 **xArcActorUserId** | **string** |  | 
 **xArcActorDisplay** | **string** |  | 
 **xArcActorEmail** | **string** |  | 
 **xArcSource** | **string** |  | 
 **xOrgId** | **string** |  | 
 **xChatSessionId** | **string** |  | 
 **xRequestId** | **string** |  | 
 **xCorrelationID** | **string** |  | 
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **authorization** | **string** |  | 

### Return type

[**TemplateVersionListResponse**](TemplateVersionListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LockDunningTemplateDraftV1

> DraftLockResponse LockDunningTemplateDraftV1(ctx, templateKey, versionId).DraftLockRequest(draftLockRequest).XArcProxySecret(xArcProxySecret).XArcFeatureFlags(xArcFeatureFlags).XArcRoles(xArcRoles).XArcActorUserId(xArcActorUserId).XArcActorDisplay(xArcActorDisplay).XArcActorEmail(xArcActorEmail).XArcSource(xArcSource).XOrgId(xOrgId).XChatSessionId(xChatSessionId).XRequestId(xRequestId).XCorrelationID(xCorrelationID).XAPIKey(xAPIKey).XTenantId(xTenantId).Authorization(authorization).Execute()

Lock Dunning Template Draft

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
	templateKey := "templateKey_example" // string | 
	versionId := "versionId_example" // string | 
	draftLockRequest := *openapiclient.NewDraftLockRequest() // DraftLockRequest | 
	xArcProxySecret := "xArcProxySecret_example" // string |  (optional)
	xArcFeatureFlags := "xArcFeatureFlags_example" // string |  (optional)
	xArcRoles := "xArcRoles_example" // string |  (optional)
	xArcActorUserId := "xArcActorUserId_example" // string |  (optional)
	xArcActorDisplay := "xArcActorDisplay_example" // string |  (optional)
	xArcActorEmail := "xArcActorEmail_example" // string |  (optional)
	xArcSource := "xArcSource_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	xChatSessionId := "xChatSessionId_example" // string |  (optional)
	xRequestId := "xRequestId_example" // string |  (optional)
	xCorrelationID := "xCorrelationID_example" // string |  (optional)
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DunningTemplatesAPI.LockDunningTemplateDraftV1(context.Background(), templateKey, versionId).DraftLockRequest(draftLockRequest).XArcProxySecret(xArcProxySecret).XArcFeatureFlags(xArcFeatureFlags).XArcRoles(xArcRoles).XArcActorUserId(xArcActorUserId).XArcActorDisplay(xArcActorDisplay).XArcActorEmail(xArcActorEmail).XArcSource(xArcSource).XOrgId(xOrgId).XChatSessionId(xChatSessionId).XRequestId(xRequestId).XCorrelationID(xCorrelationID).XAPIKey(xAPIKey).XTenantId(xTenantId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DunningTemplatesAPI.LockDunningTemplateDraftV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LockDunningTemplateDraftV1`: DraftLockResponse
	fmt.Fprintf(os.Stdout, "Response from `DunningTemplatesAPI.LockDunningTemplateDraftV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateKey** | **string** |  | 
**versionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLockDunningTemplateDraftV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **draftLockRequest** | [**DraftLockRequest**](DraftLockRequest.md) |  | 
 **xArcProxySecret** | **string** |  | 
 **xArcFeatureFlags** | **string** |  | 
 **xArcRoles** | **string** |  | 
 **xArcActorUserId** | **string** |  | 
 **xArcActorDisplay** | **string** |  | 
 **xArcActorEmail** | **string** |  | 
 **xArcSource** | **string** |  | 
 **xOrgId** | **string** |  | 
 **xChatSessionId** | **string** |  | 
 **xRequestId** | **string** |  | 
 **xCorrelationID** | **string** |  | 
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **authorization** | **string** |  | 

### Return type

[**DraftLockResponse**](DraftLockResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostArchiveTemplateKeyV1Arc

> TemplateKeyResponse PostArchiveTemplateKeyV1Arc(ctx, templateKey).TemplateArchiveRequest(templateArchiveRequest).XArcProxySecret(xArcProxySecret).XArcFeatureFlags(xArcFeatureFlags).XArcRoles(xArcRoles).XArcActorUserId(xArcActorUserId).XArcActorDisplay(xArcActorDisplay).XArcActorEmail(xArcActorEmail).XArcSource(xArcSource).XOrgId(xOrgId).XChatSessionId(xChatSessionId).XRequestId(xRequestId).XCorrelationID(xCorrelationID).XAPIKey(xAPIKey).XTenantId(xTenantId).Authorization(authorization).Execute()

Post Archive Template Key

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
	templateKey := "templateKey_example" // string | 
	templateArchiveRequest := *openapiclient.NewTemplateArchiveRequest("ChangeReason_example") // TemplateArchiveRequest | 
	xArcProxySecret := "xArcProxySecret_example" // string |  (optional)
	xArcFeatureFlags := "xArcFeatureFlags_example" // string |  (optional)
	xArcRoles := "xArcRoles_example" // string |  (optional)
	xArcActorUserId := "xArcActorUserId_example" // string |  (optional)
	xArcActorDisplay := "xArcActorDisplay_example" // string |  (optional)
	xArcActorEmail := "xArcActorEmail_example" // string |  (optional)
	xArcSource := "xArcSource_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	xChatSessionId := "xChatSessionId_example" // string |  (optional)
	xRequestId := "xRequestId_example" // string |  (optional)
	xCorrelationID := "xCorrelationID_example" // string |  (optional)
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DunningTemplatesAPI.PostArchiveTemplateKeyV1Arc(context.Background(), templateKey).TemplateArchiveRequest(templateArchiveRequest).XArcProxySecret(xArcProxySecret).XArcFeatureFlags(xArcFeatureFlags).XArcRoles(xArcRoles).XArcActorUserId(xArcActorUserId).XArcActorDisplay(xArcActorDisplay).XArcActorEmail(xArcActorEmail).XArcSource(xArcSource).XOrgId(xOrgId).XChatSessionId(xChatSessionId).XRequestId(xRequestId).XCorrelationID(xCorrelationID).XAPIKey(xAPIKey).XTenantId(xTenantId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DunningTemplatesAPI.PostArchiveTemplateKeyV1Arc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostArchiveTemplateKeyV1Arc`: TemplateKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `DunningTemplatesAPI.PostArchiveTemplateKeyV1Arc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostArchiveTemplateKeyV1ArcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **templateArchiveRequest** | [**TemplateArchiveRequest**](TemplateArchiveRequest.md) |  | 
 **xArcProxySecret** | **string** |  | 
 **xArcFeatureFlags** | **string** |  | 
 **xArcRoles** | **string** |  | 
 **xArcActorUserId** | **string** |  | 
 **xArcActorDisplay** | **string** |  | 
 **xArcActorEmail** | **string** |  | 
 **xArcSource** | **string** |  | 
 **xOrgId** | **string** |  | 
 **xChatSessionId** | **string** |  | 
 **xRequestId** | **string** |  | 
 **xCorrelationID** | **string** |  | 
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **authorization** | **string** |  | 

### Return type

[**TemplateKeyResponse**](TemplateKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTemplateKeyV1Arc

> TemplateKeyResponse PostTemplateKeyV1Arc(ctx).TemplateKeyCreateRequest(templateKeyCreateRequest).XArcProxySecret(xArcProxySecret).XArcFeatureFlags(xArcFeatureFlags).XArcRoles(xArcRoles).XArcActorUserId(xArcActorUserId).XArcActorDisplay(xArcActorDisplay).XArcActorEmail(xArcActorEmail).XArcSource(xArcSource).XOrgId(xOrgId).XChatSessionId(xChatSessionId).XRequestId(xRequestId).XCorrelationID(xCorrelationID).XAPIKey(xAPIKey).XTenantId(xTenantId).Authorization(authorization).Execute()

Post Template Key

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
	templateKeyCreateRequest := *openapiclient.NewTemplateKeyCreateRequest("CanonicalTemplateKey_example", "DisplayName_example", "ChangeReason_example") // TemplateKeyCreateRequest | 
	xArcProxySecret := "xArcProxySecret_example" // string |  (optional)
	xArcFeatureFlags := "xArcFeatureFlags_example" // string |  (optional)
	xArcRoles := "xArcRoles_example" // string |  (optional)
	xArcActorUserId := "xArcActorUserId_example" // string |  (optional)
	xArcActorDisplay := "xArcActorDisplay_example" // string |  (optional)
	xArcActorEmail := "xArcActorEmail_example" // string |  (optional)
	xArcSource := "xArcSource_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	xChatSessionId := "xChatSessionId_example" // string |  (optional)
	xRequestId := "xRequestId_example" // string |  (optional)
	xCorrelationID := "xCorrelationID_example" // string |  (optional)
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DunningTemplatesAPI.PostTemplateKeyV1Arc(context.Background()).TemplateKeyCreateRequest(templateKeyCreateRequest).XArcProxySecret(xArcProxySecret).XArcFeatureFlags(xArcFeatureFlags).XArcRoles(xArcRoles).XArcActorUserId(xArcActorUserId).XArcActorDisplay(xArcActorDisplay).XArcActorEmail(xArcActorEmail).XArcSource(xArcSource).XOrgId(xOrgId).XChatSessionId(xChatSessionId).XRequestId(xRequestId).XCorrelationID(xCorrelationID).XAPIKey(xAPIKey).XTenantId(xTenantId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DunningTemplatesAPI.PostTemplateKeyV1Arc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTemplateKeyV1Arc`: TemplateKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `DunningTemplatesAPI.PostTemplateKeyV1Arc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTemplateKeyV1ArcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateKeyCreateRequest** | [**TemplateKeyCreateRequest**](TemplateKeyCreateRequest.md) |  | 
 **xArcProxySecret** | **string** |  | 
 **xArcFeatureFlags** | **string** |  | 
 **xArcRoles** | **string** |  | 
 **xArcActorUserId** | **string** |  | 
 **xArcActorDisplay** | **string** |  | 
 **xArcActorEmail** | **string** |  | 
 **xArcSource** | **string** |  | 
 **xOrgId** | **string** |  | 
 **xChatSessionId** | **string** |  | 
 **xRequestId** | **string** |  | 
 **xCorrelationID** | **string** |  | 
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **authorization** | **string** |  | 

### Return type

[**TemplateKeyResponse**](TemplateKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTemplateTestSendV1ArcDunning

> TestSendResponse PostTemplateTestSendV1ArcDunning(ctx).TestSendRequest(testSendRequest).XArcProxySecret(xArcProxySecret).XArcFeatureFlags(xArcFeatureFlags).XArcRoles(xArcRoles).XArcActorUserId(xArcActorUserId).XArcActorDisplay(xArcActorDisplay).XArcActorEmail(xArcActorEmail).XArcSource(xArcSource).XOrgId(xOrgId).XChatSessionId(xChatSessionId).XRequestId(xRequestId).XCorrelationID(xCorrelationID).XAPIKey(xAPIKey).XTenantId(xTenantId).Authorization(authorization).Execute()

Post Template Test Send

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
	testSendRequest := *openapiclient.NewTestSendRequest("VersionId_example", "RecipientEmail_example") // TestSendRequest | 
	xArcProxySecret := "xArcProxySecret_example" // string |  (optional)
	xArcFeatureFlags := "xArcFeatureFlags_example" // string |  (optional)
	xArcRoles := "xArcRoles_example" // string |  (optional)
	xArcActorUserId := "xArcActorUserId_example" // string |  (optional)
	xArcActorDisplay := "xArcActorDisplay_example" // string |  (optional)
	xArcActorEmail := "xArcActorEmail_example" // string |  (optional)
	xArcSource := "xArcSource_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	xChatSessionId := "xChatSessionId_example" // string |  (optional)
	xRequestId := "xRequestId_example" // string |  (optional)
	xCorrelationID := "xCorrelationID_example" // string |  (optional)
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DunningTemplatesAPI.PostTemplateTestSendV1ArcDunning(context.Background()).TestSendRequest(testSendRequest).XArcProxySecret(xArcProxySecret).XArcFeatureFlags(xArcFeatureFlags).XArcRoles(xArcRoles).XArcActorUserId(xArcActorUserId).XArcActorDisplay(xArcActorDisplay).XArcActorEmail(xArcActorEmail).XArcSource(xArcSource).XOrgId(xOrgId).XChatSessionId(xChatSessionId).XRequestId(xRequestId).XCorrelationID(xCorrelationID).XAPIKey(xAPIKey).XTenantId(xTenantId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DunningTemplatesAPI.PostTemplateTestSendV1ArcDunning``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTemplateTestSendV1ArcDunning`: TestSendResponse
	fmt.Fprintf(os.Stdout, "Response from `DunningTemplatesAPI.PostTemplateTestSendV1ArcDunning`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTemplateTestSendV1ArcDunningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **testSendRequest** | [**TestSendRequest**](TestSendRequest.md) |  | 
 **xArcProxySecret** | **string** |  | 
 **xArcFeatureFlags** | **string** |  | 
 **xArcRoles** | **string** |  | 
 **xArcActorUserId** | **string** |  | 
 **xArcActorDisplay** | **string** |  | 
 **xArcActorEmail** | **string** |  | 
 **xArcSource** | **string** |  | 
 **xOrgId** | **string** |  | 
 **xChatSessionId** | **string** |  | 
 **xRequestId** | **string** |  | 
 **xCorrelationID** | **string** |  | 
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **authorization** | **string** |  | 

### Return type

[**TestSendResponse**](TestSendResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreviewDunningTemplateV1

> PreviewTemplateResponse PreviewDunningTemplateV1(ctx, templateKey).PreviewTemplateRequest(previewTemplateRequest).XArcProxySecret(xArcProxySecret).XArcFeatureFlags(xArcFeatureFlags).XArcRoles(xArcRoles).XArcActorUserId(xArcActorUserId).XArcActorDisplay(xArcActorDisplay).XArcActorEmail(xArcActorEmail).XArcSource(xArcSource).XOrgId(xOrgId).XChatSessionId(xChatSessionId).XRequestId(xRequestId).XCorrelationID(xCorrelationID).XAPIKey(xAPIKey).XTenantId(xTenantId).Authorization(authorization).Execute()

Preview Dunning Template

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
	templateKey := "templateKey_example" // string | 
	previewTemplateRequest := *openapiclient.NewPreviewTemplateRequest("VersionId_example") // PreviewTemplateRequest | 
	xArcProxySecret := "xArcProxySecret_example" // string |  (optional)
	xArcFeatureFlags := "xArcFeatureFlags_example" // string |  (optional)
	xArcRoles := "xArcRoles_example" // string |  (optional)
	xArcActorUserId := "xArcActorUserId_example" // string |  (optional)
	xArcActorDisplay := "xArcActorDisplay_example" // string |  (optional)
	xArcActorEmail := "xArcActorEmail_example" // string |  (optional)
	xArcSource := "xArcSource_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	xChatSessionId := "xChatSessionId_example" // string |  (optional)
	xRequestId := "xRequestId_example" // string |  (optional)
	xCorrelationID := "xCorrelationID_example" // string |  (optional)
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DunningTemplatesAPI.PreviewDunningTemplateV1(context.Background(), templateKey).PreviewTemplateRequest(previewTemplateRequest).XArcProxySecret(xArcProxySecret).XArcFeatureFlags(xArcFeatureFlags).XArcRoles(xArcRoles).XArcActorUserId(xArcActorUserId).XArcActorDisplay(xArcActorDisplay).XArcActorEmail(xArcActorEmail).XArcSource(xArcSource).XOrgId(xOrgId).XChatSessionId(xChatSessionId).XRequestId(xRequestId).XCorrelationID(xCorrelationID).XAPIKey(xAPIKey).XTenantId(xTenantId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DunningTemplatesAPI.PreviewDunningTemplateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PreviewDunningTemplateV1`: PreviewTemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `DunningTemplatesAPI.PreviewDunningTemplateV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewDunningTemplateV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **previewTemplateRequest** | [**PreviewTemplateRequest**](PreviewTemplateRequest.md) |  | 
 **xArcProxySecret** | **string** |  | 
 **xArcFeatureFlags** | **string** |  | 
 **xArcRoles** | **string** |  | 
 **xArcActorUserId** | **string** |  | 
 **xArcActorDisplay** | **string** |  | 
 **xArcActorEmail** | **string** |  | 
 **xArcSource** | **string** |  | 
 **xOrgId** | **string** |  | 
 **xChatSessionId** | **string** |  | 
 **xRequestId** | **string** |  | 
 **xCorrelationID** | **string** |  | 
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **authorization** | **string** |  | 

### Return type

[**PreviewTemplateResponse**](PreviewTemplateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublishDunningTemplateV1

> TemplateVersionResponse PublishDunningTemplateV1(ctx, templateKey).PublishTemplateRequest(publishTemplateRequest).XArcProxySecret(xArcProxySecret).XArcFeatureFlags(xArcFeatureFlags).XArcRoles(xArcRoles).XArcActorUserId(xArcActorUserId).XArcActorDisplay(xArcActorDisplay).XArcActorEmail(xArcActorEmail).XArcSource(xArcSource).XOrgId(xOrgId).XChatSessionId(xChatSessionId).XRequestId(xRequestId).XCorrelationID(xCorrelationID).XAPIKey(xAPIKey).XTenantId(xTenantId).Authorization(authorization).Execute()

Publish Dunning Template

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
	templateKey := "templateKey_example" // string | 
	publishTemplateRequest := *openapiclient.NewPublishTemplateRequest("VersionId_example", false, "ChangeReason_example") // PublishTemplateRequest | 
	xArcProxySecret := "xArcProxySecret_example" // string |  (optional)
	xArcFeatureFlags := "xArcFeatureFlags_example" // string |  (optional)
	xArcRoles := "xArcRoles_example" // string |  (optional)
	xArcActorUserId := "xArcActorUserId_example" // string |  (optional)
	xArcActorDisplay := "xArcActorDisplay_example" // string |  (optional)
	xArcActorEmail := "xArcActorEmail_example" // string |  (optional)
	xArcSource := "xArcSource_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	xChatSessionId := "xChatSessionId_example" // string |  (optional)
	xRequestId := "xRequestId_example" // string |  (optional)
	xCorrelationID := "xCorrelationID_example" // string |  (optional)
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DunningTemplatesAPI.PublishDunningTemplateV1(context.Background(), templateKey).PublishTemplateRequest(publishTemplateRequest).XArcProxySecret(xArcProxySecret).XArcFeatureFlags(xArcFeatureFlags).XArcRoles(xArcRoles).XArcActorUserId(xArcActorUserId).XArcActorDisplay(xArcActorDisplay).XArcActorEmail(xArcActorEmail).XArcSource(xArcSource).XOrgId(xOrgId).XChatSessionId(xChatSessionId).XRequestId(xRequestId).XCorrelationID(xCorrelationID).XAPIKey(xAPIKey).XTenantId(xTenantId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DunningTemplatesAPI.PublishDunningTemplateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublishDunningTemplateV1`: TemplateVersionResponse
	fmt.Fprintf(os.Stdout, "Response from `DunningTemplatesAPI.PublishDunningTemplateV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublishDunningTemplateV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **publishTemplateRequest** | [**PublishTemplateRequest**](PublishTemplateRequest.md) |  | 
 **xArcProxySecret** | **string** |  | 
 **xArcFeatureFlags** | **string** |  | 
 **xArcRoles** | **string** |  | 
 **xArcActorUserId** | **string** |  | 
 **xArcActorDisplay** | **string** |  | 
 **xArcActorEmail** | **string** |  | 
 **xArcSource** | **string** |  | 
 **xOrgId** | **string** |  | 
 **xChatSessionId** | **string** |  | 
 **xRequestId** | **string** |  | 
 **xCorrelationID** | **string** |  | 
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **authorization** | **string** |  | 

### Return type

[**TemplateVersionResponse**](TemplateVersionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutDisclosurePolicyV1ArcDunning

> DisclosurePolicyResponse PutDisclosurePolicyV1ArcDunning(ctx).DisclosurePolicyUpdateRequest(disclosurePolicyUpdateRequest).XArcProxySecret(xArcProxySecret).XArcFeatureFlags(xArcFeatureFlags).XArcRoles(xArcRoles).XArcActorUserId(xArcActorUserId).XArcActorDisplay(xArcActorDisplay).XArcActorEmail(xArcActorEmail).XArcSource(xArcSource).XOrgId(xOrgId).XChatSessionId(xChatSessionId).XRequestId(xRequestId).XCorrelationID(xCorrelationID).XAPIKey(xAPIKey).XTenantId(xTenantId).Authorization(authorization).Execute()

Put Disclosure Policy

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
	disclosurePolicyUpdateRequest := *openapiclient.NewDisclosurePolicyUpdateRequest("ChangeReason_example") // DisclosurePolicyUpdateRequest | 
	xArcProxySecret := "xArcProxySecret_example" // string |  (optional)
	xArcFeatureFlags := "xArcFeatureFlags_example" // string |  (optional)
	xArcRoles := "xArcRoles_example" // string |  (optional)
	xArcActorUserId := "xArcActorUserId_example" // string |  (optional)
	xArcActorDisplay := "xArcActorDisplay_example" // string |  (optional)
	xArcActorEmail := "xArcActorEmail_example" // string |  (optional)
	xArcSource := "xArcSource_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	xChatSessionId := "xChatSessionId_example" // string |  (optional)
	xRequestId := "xRequestId_example" // string |  (optional)
	xCorrelationID := "xCorrelationID_example" // string |  (optional)
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DunningTemplatesAPI.PutDisclosurePolicyV1ArcDunning(context.Background()).DisclosurePolicyUpdateRequest(disclosurePolicyUpdateRequest).XArcProxySecret(xArcProxySecret).XArcFeatureFlags(xArcFeatureFlags).XArcRoles(xArcRoles).XArcActorUserId(xArcActorUserId).XArcActorDisplay(xArcActorDisplay).XArcActorEmail(xArcActorEmail).XArcSource(xArcSource).XOrgId(xOrgId).XChatSessionId(xChatSessionId).XRequestId(xRequestId).XCorrelationID(xCorrelationID).XAPIKey(xAPIKey).XTenantId(xTenantId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DunningTemplatesAPI.PutDisclosurePolicyV1ArcDunning``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutDisclosurePolicyV1ArcDunning`: DisclosurePolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `DunningTemplatesAPI.PutDisclosurePolicyV1ArcDunning`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutDisclosurePolicyV1ArcDunningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **disclosurePolicyUpdateRequest** | [**DisclosurePolicyUpdateRequest**](DisclosurePolicyUpdateRequest.md) |  | 
 **xArcProxySecret** | **string** |  | 
 **xArcFeatureFlags** | **string** |  | 
 **xArcRoles** | **string** |  | 
 **xArcActorUserId** | **string** |  | 
 **xArcActorDisplay** | **string** |  | 
 **xArcActorEmail** | **string** |  | 
 **xArcSource** | **string** |  | 
 **xOrgId** | **string** |  | 
 **xChatSessionId** | **string** |  | 
 **xRequestId** | **string** |  | 
 **xCorrelationID** | **string** |  | 
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **authorization** | **string** |  | 

### Return type

[**DisclosurePolicyResponse**](DisclosurePolicyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPaymentInstructionsV1ArcDunning

> PaymentInstructionsResponse PutPaymentInstructionsV1ArcDunning(ctx).PaymentInstructionsUpdateRequest(paymentInstructionsUpdateRequest).XArcProxySecret(xArcProxySecret).XArcFeatureFlags(xArcFeatureFlags).XArcRoles(xArcRoles).XArcActorUserId(xArcActorUserId).XArcActorDisplay(xArcActorDisplay).XArcActorEmail(xArcActorEmail).XArcSource(xArcSource).XOrgId(xOrgId).XChatSessionId(xChatSessionId).XRequestId(xRequestId).XCorrelationID(xCorrelationID).XAPIKey(xAPIKey).XTenantId(xTenantId).Authorization(authorization).Execute()

Put Payment Instructions

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
	paymentInstructionsUpdateRequest := *openapiclient.NewPaymentInstructionsUpdateRequest("PaymentInstructionsText_example", "ChangeReason_example") // PaymentInstructionsUpdateRequest | 
	xArcProxySecret := "xArcProxySecret_example" // string |  (optional)
	xArcFeatureFlags := "xArcFeatureFlags_example" // string |  (optional)
	xArcRoles := "xArcRoles_example" // string |  (optional)
	xArcActorUserId := "xArcActorUserId_example" // string |  (optional)
	xArcActorDisplay := "xArcActorDisplay_example" // string |  (optional)
	xArcActorEmail := "xArcActorEmail_example" // string |  (optional)
	xArcSource := "xArcSource_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	xChatSessionId := "xChatSessionId_example" // string |  (optional)
	xRequestId := "xRequestId_example" // string |  (optional)
	xCorrelationID := "xCorrelationID_example" // string |  (optional)
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DunningTemplatesAPI.PutPaymentInstructionsV1ArcDunning(context.Background()).PaymentInstructionsUpdateRequest(paymentInstructionsUpdateRequest).XArcProxySecret(xArcProxySecret).XArcFeatureFlags(xArcFeatureFlags).XArcRoles(xArcRoles).XArcActorUserId(xArcActorUserId).XArcActorDisplay(xArcActorDisplay).XArcActorEmail(xArcActorEmail).XArcSource(xArcSource).XOrgId(xOrgId).XChatSessionId(xChatSessionId).XRequestId(xRequestId).XCorrelationID(xCorrelationID).XAPIKey(xAPIKey).XTenantId(xTenantId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DunningTemplatesAPI.PutPaymentInstructionsV1ArcDunning``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutPaymentInstructionsV1ArcDunning`: PaymentInstructionsResponse
	fmt.Fprintf(os.Stdout, "Response from `DunningTemplatesAPI.PutPaymentInstructionsV1ArcDunning`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutPaymentInstructionsV1ArcDunningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentInstructionsUpdateRequest** | [**PaymentInstructionsUpdateRequest**](PaymentInstructionsUpdateRequest.md) |  | 
 **xArcProxySecret** | **string** |  | 
 **xArcFeatureFlags** | **string** |  | 
 **xArcRoles** | **string** |  | 
 **xArcActorUserId** | **string** |  | 
 **xArcActorDisplay** | **string** |  | 
 **xArcActorEmail** | **string** |  | 
 **xArcSource** | **string** |  | 
 **xOrgId** | **string** |  | 
 **xChatSessionId** | **string** |  | 
 **xRequestId** | **string** |  | 
 **xCorrelationID** | **string** |  | 
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **authorization** | **string** |  | 

### Return type

[**PaymentInstructionsResponse**](PaymentInstructionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveDunningTemplateDraftV1

> TemplateVersionResponse SaveDunningTemplateDraftV1(ctx, templateKey).DraftSaveRequest(draftSaveRequest).XArcProxySecret(xArcProxySecret).XArcFeatureFlags(xArcFeatureFlags).XArcRoles(xArcRoles).XArcActorUserId(xArcActorUserId).XArcActorDisplay(xArcActorDisplay).XArcActorEmail(xArcActorEmail).XArcSource(xArcSource).XOrgId(xOrgId).XChatSessionId(xChatSessionId).XRequestId(xRequestId).XCorrelationID(xCorrelationID).XAPIKey(xAPIKey).XTenantId(xTenantId).Authorization(authorization).Execute()

Save Dunning Template Draft

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
	templateKey := "templateKey_example" // string | 
	draftSaveRequest := *openapiclient.NewDraftSaveRequest("SubjectTemplate_example", "BodyTemplate_example", "ChangeReason_example") // DraftSaveRequest | 
	xArcProxySecret := "xArcProxySecret_example" // string |  (optional)
	xArcFeatureFlags := "xArcFeatureFlags_example" // string |  (optional)
	xArcRoles := "xArcRoles_example" // string |  (optional)
	xArcActorUserId := "xArcActorUserId_example" // string |  (optional)
	xArcActorDisplay := "xArcActorDisplay_example" // string |  (optional)
	xArcActorEmail := "xArcActorEmail_example" // string |  (optional)
	xArcSource := "xArcSource_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	xChatSessionId := "xChatSessionId_example" // string |  (optional)
	xRequestId := "xRequestId_example" // string |  (optional)
	xCorrelationID := "xCorrelationID_example" // string |  (optional)
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DunningTemplatesAPI.SaveDunningTemplateDraftV1(context.Background(), templateKey).DraftSaveRequest(draftSaveRequest).XArcProxySecret(xArcProxySecret).XArcFeatureFlags(xArcFeatureFlags).XArcRoles(xArcRoles).XArcActorUserId(xArcActorUserId).XArcActorDisplay(xArcActorDisplay).XArcActorEmail(xArcActorEmail).XArcSource(xArcSource).XOrgId(xOrgId).XChatSessionId(xChatSessionId).XRequestId(xRequestId).XCorrelationID(xCorrelationID).XAPIKey(xAPIKey).XTenantId(xTenantId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DunningTemplatesAPI.SaveDunningTemplateDraftV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SaveDunningTemplateDraftV1`: TemplateVersionResponse
	fmt.Fprintf(os.Stdout, "Response from `DunningTemplatesAPI.SaveDunningTemplateDraftV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSaveDunningTemplateDraftV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **draftSaveRequest** | [**DraftSaveRequest**](DraftSaveRequest.md) |  | 
 **xArcProxySecret** | **string** |  | 
 **xArcFeatureFlags** | **string** |  | 
 **xArcRoles** | **string** |  | 
 **xArcActorUserId** | **string** |  | 
 **xArcActorDisplay** | **string** |  | 
 **xArcActorEmail** | **string** |  | 
 **xArcSource** | **string** |  | 
 **xOrgId** | **string** |  | 
 **xChatSessionId** | **string** |  | 
 **xRequestId** | **string** |  | 
 **xCorrelationID** | **string** |  | 
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **authorization** | **string** |  | 

### Return type

[**TemplateVersionResponse**](TemplateVersionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlockDunningTemplateDraftV1

> TemplateVersionResponse UnlockDunningTemplateDraftV1(ctx, templateKey, versionId).DraftUnlockRequest(draftUnlockRequest).XArcProxySecret(xArcProxySecret).XArcFeatureFlags(xArcFeatureFlags).XArcRoles(xArcRoles).XArcActorUserId(xArcActorUserId).XArcActorDisplay(xArcActorDisplay).XArcActorEmail(xArcActorEmail).XArcSource(xArcSource).XOrgId(xOrgId).XChatSessionId(xChatSessionId).XRequestId(xRequestId).XCorrelationID(xCorrelationID).XAPIKey(xAPIKey).XTenantId(xTenantId).Authorization(authorization).Execute()

Unlock Dunning Template Draft

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
	templateKey := "templateKey_example" // string | 
	versionId := "versionId_example" // string | 
	draftUnlockRequest := *openapiclient.NewDraftUnlockRequest("LockId_example", "LockToken_example") // DraftUnlockRequest | 
	xArcProxySecret := "xArcProxySecret_example" // string |  (optional)
	xArcFeatureFlags := "xArcFeatureFlags_example" // string |  (optional)
	xArcRoles := "xArcRoles_example" // string |  (optional)
	xArcActorUserId := "xArcActorUserId_example" // string |  (optional)
	xArcActorDisplay := "xArcActorDisplay_example" // string |  (optional)
	xArcActorEmail := "xArcActorEmail_example" // string |  (optional)
	xArcSource := "xArcSource_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	xChatSessionId := "xChatSessionId_example" // string |  (optional)
	xRequestId := "xRequestId_example" // string |  (optional)
	xCorrelationID := "xCorrelationID_example" // string |  (optional)
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DunningTemplatesAPI.UnlockDunningTemplateDraftV1(context.Background(), templateKey, versionId).DraftUnlockRequest(draftUnlockRequest).XArcProxySecret(xArcProxySecret).XArcFeatureFlags(xArcFeatureFlags).XArcRoles(xArcRoles).XArcActorUserId(xArcActorUserId).XArcActorDisplay(xArcActorDisplay).XArcActorEmail(xArcActorEmail).XArcSource(xArcSource).XOrgId(xOrgId).XChatSessionId(xChatSessionId).XRequestId(xRequestId).XCorrelationID(xCorrelationID).XAPIKey(xAPIKey).XTenantId(xTenantId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DunningTemplatesAPI.UnlockDunningTemplateDraftV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnlockDunningTemplateDraftV1`: TemplateVersionResponse
	fmt.Fprintf(os.Stdout, "Response from `DunningTemplatesAPI.UnlockDunningTemplateDraftV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateKey** | **string** |  | 
**versionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnlockDunningTemplateDraftV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **draftUnlockRequest** | [**DraftUnlockRequest**](DraftUnlockRequest.md) |  | 
 **xArcProxySecret** | **string** |  | 
 **xArcFeatureFlags** | **string** |  | 
 **xArcRoles** | **string** |  | 
 **xArcActorUserId** | **string** |  | 
 **xArcActorDisplay** | **string** |  | 
 **xArcActorEmail** | **string** |  | 
 **xArcSource** | **string** |  | 
 **xOrgId** | **string** |  | 
 **xChatSessionId** | **string** |  | 
 **xRequestId** | **string** |  | 
 **xCorrelationID** | **string** |  | 
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **authorization** | **string** |  | 

### Return type

[**TemplateVersionResponse**](TemplateVersionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateDunningTemplateV1

> ValidateTemplateResponse ValidateDunningTemplateV1(ctx, templateKey).ValidateTemplateRequest(validateTemplateRequest).XArcProxySecret(xArcProxySecret).XArcFeatureFlags(xArcFeatureFlags).XArcRoles(xArcRoles).XArcActorUserId(xArcActorUserId).XArcActorDisplay(xArcActorDisplay).XArcActorEmail(xArcActorEmail).XArcSource(xArcSource).XOrgId(xOrgId).XChatSessionId(xChatSessionId).XRequestId(xRequestId).XCorrelationID(xCorrelationID).XAPIKey(xAPIKey).XTenantId(xTenantId).Authorization(authorization).Execute()

Validate Dunning Template

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
	templateKey := "templateKey_example" // string | 
	validateTemplateRequest := *openapiclient.NewValidateTemplateRequest("SubjectTemplate_example", "BodyTemplate_example") // ValidateTemplateRequest | 
	xArcProxySecret := "xArcProxySecret_example" // string |  (optional)
	xArcFeatureFlags := "xArcFeatureFlags_example" // string |  (optional)
	xArcRoles := "xArcRoles_example" // string |  (optional)
	xArcActorUserId := "xArcActorUserId_example" // string |  (optional)
	xArcActorDisplay := "xArcActorDisplay_example" // string |  (optional)
	xArcActorEmail := "xArcActorEmail_example" // string |  (optional)
	xArcSource := "xArcSource_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	xChatSessionId := "xChatSessionId_example" // string |  (optional)
	xRequestId := "xRequestId_example" // string |  (optional)
	xCorrelationID := "xCorrelationID_example" // string |  (optional)
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DunningTemplatesAPI.ValidateDunningTemplateV1(context.Background(), templateKey).ValidateTemplateRequest(validateTemplateRequest).XArcProxySecret(xArcProxySecret).XArcFeatureFlags(xArcFeatureFlags).XArcRoles(xArcRoles).XArcActorUserId(xArcActorUserId).XArcActorDisplay(xArcActorDisplay).XArcActorEmail(xArcActorEmail).XArcSource(xArcSource).XOrgId(xOrgId).XChatSessionId(xChatSessionId).XRequestId(xRequestId).XCorrelationID(xCorrelationID).XAPIKey(xAPIKey).XTenantId(xTenantId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DunningTemplatesAPI.ValidateDunningTemplateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateDunningTemplateV1`: ValidateTemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `DunningTemplatesAPI.ValidateDunningTemplateV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateDunningTemplateV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **validateTemplateRequest** | [**ValidateTemplateRequest**](ValidateTemplateRequest.md) |  | 
 **xArcProxySecret** | **string** |  | 
 **xArcFeatureFlags** | **string** |  | 
 **xArcRoles** | **string** |  | 
 **xArcActorUserId** | **string** |  | 
 **xArcActorDisplay** | **string** |  | 
 **xArcActorEmail** | **string** |  | 
 **xArcSource** | **string** |  | 
 **xOrgId** | **string** |  | 
 **xChatSessionId** | **string** |  | 
 **xRequestId** | **string** |  | 
 **xCorrelationID** | **string** |  | 
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **authorization** | **string** |  | 

### Return type

[**ValidateTemplateResponse**](ValidateTemplateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

