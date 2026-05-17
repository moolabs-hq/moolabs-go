# \TenantAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePortalTokenV1**](TenantAPI.md#DeletePortalTokenV1) | **Delete** /v1/tenant/portal-tokens/{token_id} | Delete Portal Token
[**GetApiKeysV1**](TenantAPI.md#GetApiKeysV1) | **Get** /v1/tenant/api-keys | Get Api Keys
[**GetAuditChanges**](TenantAPI.md#GetAuditChanges) | **Get** /v1/tenant/audit/changes | Get Audit Changes
[**GetAuditEvidence**](TenantAPI.md#GetAuditEvidence) | **Get** /v1/tenant/audit/evidence | Get Audit Evidence
[**GetAuditTraces**](TenantAPI.md#GetAuditTraces) | **Get** /v1/tenant/audit/traces | Get Audit Traces
[**GetAuditWarnings**](TenantAPI.md#GetAuditWarnings) | **Get** /v1/tenant/audit/warnings | Get Audit Warnings
[**GetDomain**](TenantAPI.md#GetDomain) | **Get** /v1/tenant/communications/domain | Get Domain
[**GetFeatureFlagsV1**](TenantAPI.md#GetFeatureFlagsV1) | **Get** /v1/tenant/audit/feature-flags | Get Feature Flags
[**GetIntegrations**](TenantAPI.md#GetIntegrations) | **Get** /v1/tenant/integrations | Get Integrations
[**GetIntegrationsHealth**](TenantAPI.md#GetIntegrationsHealth) | **Get** /v1/tenant/integrations/health | Get Integrations Health
[**GetIntegrationsMapping**](TenantAPI.md#GetIntegrationsMapping) | **Get** /v1/tenant/integrations/mapping | Get Integrations Mapping
[**GetMoometerNamespaceIdV1**](TenantAPI.md#GetMoometerNamespaceIdV1) | **Get** /v1/tenant/moometer-namespace | Get Moometer Namespace Id
[**GetPortalTokensV1**](TenantAPI.md#GetPortalTokensV1) | **Get** /v1/tenant/portal-tokens | Get Portal Tokens
[**GetRevenueRecognitionV1**](TenantAPI.md#GetRevenueRecognitionV1) | **Get** /v1/tenant/revenue-recognition | Get Revenue Recognition
[**GetSender**](TenantAPI.md#GetSender) | **Get** /v1/tenant/communications/sender | Get Sender
[**GetTemplate**](TenantAPI.md#GetTemplate) | **Get** /v1/tenant/communications/templates/{template_id} | Get Template
[**GetTemplatesMeta**](TenantAPI.md#GetTemplatesMeta) | **Get** /v1/tenant/communications/templates/meta | Get Templates Meta
[**GetTenantContext**](TenantAPI.md#GetTenantContext) | **Get** /v1/tenant/context | Get Tenant Context
[**GetWebhook**](TenantAPI.md#GetWebhook) | **Get** /v1/tenant/communications/webhook | Get Webhook
[**GetWebhookLogs**](TenantAPI.md#GetWebhookLogs) | **Get** /v1/tenant/webhook/logs | Get Webhook Logs
[**GetWebhookMetrics**](TenantAPI.md#GetWebhookMetrics) | **Get** /v1/tenant/webhook/metrics | Get Webhook Metrics
[**HubspotConnect**](TenantAPI.md#HubspotConnect) | **Post** /v1/tenant/integrations/hubspot/connect | Hubspot Connect
[**HubspotDisconnect**](TenantAPI.md#HubspotDisconnect) | **Post** /v1/tenant/integrations/hubspot/disconnect | Hubspot Disconnect
[**HubspotOauthCallback**](TenantAPI.md#HubspotOauthCallback) | **Get** /v1/tenant/integrations/hubspot/oauth/callback | Hubspot Oauth Callback
[**HubspotTestConnection**](TenantAPI.md#HubspotTestConnection) | **Post** /v1/tenant/integrations/hubspot/test | Hubspot Test Connection
[**IssueIntegrationKey**](TenantAPI.md#IssueIntegrationKey) | **Post** /v1/tenant/integrations/{provider}/keys | Issue Integration Key
[**ListTemplates**](TenantAPI.md#ListTemplates) | **Get** /v1/tenant/communications/templates | List Templates
[**NetsuiteConnect**](TenantAPI.md#NetsuiteConnect) | **Post** /v1/tenant/integrations/netsuite/connect | Netsuite Connect
[**NetsuiteDisconnect**](TenantAPI.md#NetsuiteDisconnect) | **Post** /v1/tenant/integrations/netsuite/disconnect | Netsuite Disconnect
[**NetsuiteOauthCallback**](TenantAPI.md#NetsuiteOauthCallback) | **Get** /v1/tenant/integrations/netsuite/oauth/callback | Netsuite Oauth Callback
[**NetsuiteTestConnection**](TenantAPI.md#NetsuiteTestConnection) | **Post** /v1/tenant/integrations/netsuite/test | Netsuite Test Connection
[**PostApiKeyV1**](TenantAPI.md#PostApiKeyV1) | **Post** /v1/tenant/api-keys | Post Api Key
[**PostDomainVerify**](TenantAPI.md#PostDomainVerify) | **Post** /v1/tenant/communications/domain/verify | Post Domain Verify
[**PostPortalTokenV1**](TenantAPI.md#PostPortalTokenV1) | **Post** /v1/tenant/portal-tokens | Post Portal Token
[**PostRevokeApiKeyV1**](TenantAPI.md#PostRevokeApiKeyV1) | **Post** /v1/tenant/api-keys/{key_id}/revoke | Post Revoke Api Key
[**PreviewTemplate**](TenantAPI.md#PreviewTemplate) | **Post** /v1/tenant/communications/templates/{template_id}/preview | Preview Template
[**PutDomain**](TenantAPI.md#PutDomain) | **Put** /v1/tenant/communications/domain | Put Domain
[**PutFeatureFlagV1**](TenantAPI.md#PutFeatureFlagV1) | **Put** /v1/tenant/audit/feature-flags/{flag_id} | Put Feature Flag
[**PutIntegration**](TenantAPI.md#PutIntegration) | **Put** /v1/tenant/integrations/{provider} | Put Integration
[**PutSender**](TenantAPI.md#PutSender) | **Put** /v1/tenant/communications/sender | Put Sender
[**PutTemplate**](TenantAPI.md#PutTemplate) | **Put** /v1/tenant/communications/templates/{template_id} | Put Template
[**PutWebhook**](TenantAPI.md#PutWebhook) | **Put** /v1/tenant/communications/webhook | Put Webhook
[**QuickbooksConnect**](TenantAPI.md#QuickbooksConnect) | **Post** /v1/tenant/integrations/quickbooks/connect | Quickbooks Connect
[**QuickbooksDisconnect**](TenantAPI.md#QuickbooksDisconnect) | **Post** /v1/tenant/integrations/quickbooks/disconnect | Quickbooks Disconnect
[**QuickbooksOauthCallback**](TenantAPI.md#QuickbooksOauthCallback) | **Get** /v1/tenant/integrations/quickbooks/oauth/callback | Quickbooks Oauth Callback
[**QuickbooksTestConnection**](TenantAPI.md#QuickbooksTestConnection) | **Post** /v1/tenant/integrations/quickbooks/test | Quickbooks Test Connection
[**RetryFailedDeliveries**](TenantAPI.md#RetryFailedDeliveries) | **Post** /v1/tenant/webhook/retry | Retry Failed Deliveries
[**RevokeAllApiKeysV1TenantDanger**](TenantAPI.md#RevokeAllApiKeysV1TenantDanger) | **Post** /v1/tenant/danger/revoke-all-api-keys | Revoke All Api Keys
[**RevokeAllPortalTokensV1TenantDanger**](TenantAPI.md#RevokeAllPortalTokensV1TenantDanger) | **Post** /v1/tenant/danger/revoke-all-portal-tokens | Revoke All Portal Tokens
[**SalesforceConnect**](TenantAPI.md#SalesforceConnect) | **Post** /v1/tenant/integrations/salesforce/connect | Salesforce Connect
[**SalesforceDisconnect**](TenantAPI.md#SalesforceDisconnect) | **Post** /v1/tenant/integrations/salesforce/disconnect | Salesforce Disconnect
[**SalesforceOauthCallback**](TenantAPI.md#SalesforceOauthCallback) | **Get** /v1/tenant/integrations/salesforce/oauth/callback | Salesforce Oauth Callback
[**SalesforceTestConnection**](TenantAPI.md#SalesforceTestConnection) | **Post** /v1/tenant/integrations/salesforce/test | Salesforce Test Connection
[**TestSendTemplateV1**](TenantAPI.md#TestSendTemplateV1) | **Post** /v1/tenant/communications/templates/{template_id}/test-send | Test Send Template
[**TestWebhook**](TenantAPI.md#TestWebhook) | **Post** /v1/tenant/communications/webhook/test | Test Webhook
[**UpdateRevenueRecognitionV1**](TenantAPI.md#UpdateRevenueRecognitionV1) | **Put** /v1/tenant/revenue-recognition | Update Revenue Recognition
[**XeroConnect**](TenantAPI.md#XeroConnect) | **Post** /v1/tenant/integrations/xero/connect | Xero Connect
[**XeroDisconnect**](TenantAPI.md#XeroDisconnect) | **Post** /v1/tenant/integrations/xero/disconnect | Xero Disconnect
[**XeroOauthCallback**](TenantAPI.md#XeroOauthCallback) | **Get** /v1/tenant/integrations/xero/oauth/callback | Xero Oauth Callback
[**XeroTestConnection**](TenantAPI.md#XeroTestConnection) | **Post** /v1/tenant/integrations/xero/test | Xero Test Connection



## DeletePortalTokenV1

> interface{} DeletePortalTokenV1(ctx, tokenId).RevokePortalTokenRequest(revokePortalTokenRequest).Execute()

Delete Portal Token



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
	tokenId := "tokenId_example" // string | 
	revokePortalTokenRequest := *openapiclient.NewRevokePortalTokenRequest() // RevokePortalTokenRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.DeletePortalTokenV1(context.Background(), tokenId).RevokePortalTokenRequest(revokePortalTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.DeletePortalTokenV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePortalTokenV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.DeletePortalTokenV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePortalTokenV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revokePortalTokenRequest** | [**RevokePortalTokenRequest**](RevokePortalTokenRequest.md) |  | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiKeysV1

> []ApiKeyItem GetApiKeysV1(ctx).Execute()

Get Api Keys



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
	resp, r, err := apiClient.TenantAPI.GetApiKeysV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.GetApiKeysV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiKeysV1`: []ApiKeyItem
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.GetApiKeysV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiKeysV1Request struct via the builder pattern


### Return type

[**[]ApiKeyItem**](ApiKeyItem.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuditChanges

> []ChangeItem GetAuditChanges(ctx).Limit(limit).Section(section).Execute()

Get Audit Changes



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
	limit := int32(56) // int32 |  (optional) (default to 50)
	section := "section_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.GetAuditChanges(context.Background()).Limit(limit).Section(section).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.GetAuditChanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuditChanges`: []ChangeItem
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.GetAuditChanges`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 50]
 **section** | **string** |  | 

### Return type

[**[]ChangeItem**](ChangeItem.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuditEvidence

> interface{} GetAuditEvidence(ctx).Limit(limit).Execute()

Get Audit Evidence



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
	limit := int32(56) // int32 |  (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.GetAuditEvidence(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.GetAuditEvidence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuditEvidence`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.GetAuditEvidence`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditEvidenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 500]

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuditTraces

> TracesResponse GetAuditTraces(ctx).Execute()

Get Audit Traces



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
	resp, r, err := apiClient.TenantAPI.GetAuditTraces(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.GetAuditTraces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuditTraces`: TracesResponse
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.GetAuditTraces`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditTracesRequest struct via the builder pattern


### Return type

[**TracesResponse**](TracesResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuditWarnings

> []WarningItem GetAuditWarnings(ctx).Limit(limit).Execute()

Get Audit Warnings



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
	limit := int32(56) // int32 |  (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.GetAuditWarnings(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.GetAuditWarnings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuditWarnings`: []WarningItem
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.GetAuditWarnings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditWarningsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 20]

### Return type

[**[]WarningItem**](WarningItem.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDomain

> DomainResponse GetDomain(ctx).Execute()

Get Domain



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
	resp, r, err := apiClient.TenantAPI.GetDomain(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.GetDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDomain`: DomainResponse
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.GetDomain`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainRequest struct via the builder pattern


### Return type

[**DomainResponse**](DomainResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeatureFlagsV1

> []FeatureFlagItem GetFeatureFlagsV1(ctx).Execute()

Get Feature Flags



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
	resp, r, err := apiClient.TenantAPI.GetFeatureFlagsV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.GetFeatureFlagsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFeatureFlagsV1`: []FeatureFlagItem
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.GetFeatureFlagsV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeatureFlagsV1Request struct via the builder pattern


### Return type

[**[]FeatureFlagItem**](FeatureFlagItem.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrations

> []ConnectorItem GetIntegrations(ctx).Execute()

Get Integrations



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
	resp, r, err := apiClient.TenantAPI.GetIntegrations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.GetIntegrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrations`: []ConnectorItem
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.GetIntegrations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationsRequest struct via the builder pattern


### Return type

[**[]ConnectorItem**](ConnectorItem.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrationsHealth

> interface{} GetIntegrationsHealth(ctx).Execute()

Get Integrations Health



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
	resp, r, err := apiClient.TenantAPI.GetIntegrationsHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.GetIntegrationsHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrationsHealth`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.GetIntegrationsHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationsHealthRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrationsMapping

> []MeterMappingItem GetIntegrationsMapping(ctx).Execute()

Get Integrations Mapping



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
	resp, r, err := apiClient.TenantAPI.GetIntegrationsMapping(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.GetIntegrationsMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrationsMapping`: []MeterMappingItem
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.GetIntegrationsMapping`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationsMappingRequest struct via the builder pattern


### Return type

[**[]MeterMappingItem**](MeterMappingItem.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMoometerNamespaceIdV1

> interface{} GetMoometerNamespaceIdV1(ctx).Execute()

Get Moometer Namespace Id



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
	resp, r, err := apiClient.TenantAPI.GetMoometerNamespaceIdV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.GetMoometerNamespaceIdV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMoometerNamespaceIdV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.GetMoometerNamespaceIdV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMoometerNamespaceIdV1Request struct via the builder pattern


### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortalTokensV1

> []PortalTokenItem GetPortalTokensV1(ctx).Execute()

Get Portal Tokens



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
	resp, r, err := apiClient.TenantAPI.GetPortalTokensV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.GetPortalTokensV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortalTokensV1`: []PortalTokenItem
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.GetPortalTokensV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortalTokensV1Request struct via the builder pattern


### Return type

[**[]PortalTokenItem**](PortalTokenItem.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRevenueRecognitionV1

> interface{} GetRevenueRecognitionV1(ctx).Execute()

Get Revenue Recognition



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
	resp, r, err := apiClient.TenantAPI.GetRevenueRecognitionV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.GetRevenueRecognitionV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRevenueRecognitionV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.GetRevenueRecognitionV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRevenueRecognitionV1Request struct via the builder pattern


### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSender

> SenderResponse GetSender(ctx).Execute()

Get Sender



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
	resp, r, err := apiClient.TenantAPI.GetSender(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.GetSender``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSender`: SenderResponse
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.GetSender`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSenderRequest struct via the builder pattern


### Return type

[**SenderResponse**](SenderResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplate

> TemplateItem GetTemplate(ctx, templateId).Execute()

Get Template



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
	templateId := "templateId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.GetTemplate(context.Background(), templateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.GetTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplate`: TemplateItem
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.GetTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TemplateItem**](TemplateItem.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplatesMeta

> TemplateMetaResponse GetTemplatesMeta(ctx).Execute()

Get Templates Meta



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
	resp, r, err := apiClient.TenantAPI.GetTemplatesMeta(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.GetTemplatesMeta``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplatesMeta`: TemplateMetaResponse
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.GetTemplatesMeta`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplatesMetaRequest struct via the builder pattern


### Return type

[**TemplateMetaResponse**](TemplateMetaResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenantContext

> interface{} GetTenantContext(ctx).Execute()

Get Tenant Context



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
	resp, r, err := apiClient.TenantAPI.GetTenantContext(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.GetTenantContext``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantContext`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.GetTenantContext`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantContextRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhook

> WebhookResponse GetWebhook(ctx).Execute()

Get Webhook



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
	resp, r, err := apiClient.TenantAPI.GetWebhook(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.GetWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhook`: WebhookResponse
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.GetWebhook`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookRequest struct via the builder pattern


### Return type

[**WebhookResponse**](WebhookResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookLogs

> interface{} GetWebhookLogs(ctx).Limit(limit).Execute()

Get Webhook Logs



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
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.GetWebhookLogs(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.GetWebhookLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhookLogs`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.GetWebhookLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 50]

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookMetrics

> interface{} GetWebhookMetrics(ctx).Execute()

Get Webhook Metrics



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
	resp, r, err := apiClient.TenantAPI.GetWebhookMetrics(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.GetWebhookMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhookMetrics`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.GetWebhookMetrics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookMetricsRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HubspotConnect

> interface{} HubspotConnect(ctx).HubSpotConnectRequest(hubSpotConnectRequest).Execute()

Hubspot Connect

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
	hubSpotConnectRequest := *openapiclient.NewHubSpotConnectRequest("SyncDirection_example", "SyncFrequency_example") // HubSpotConnectRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.HubspotConnect(context.Background()).HubSpotConnectRequest(hubSpotConnectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.HubspotConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HubspotConnect`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.HubspotConnect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHubspotConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hubSpotConnectRequest** | [**HubSpotConnectRequest**](HubSpotConnectRequest.md) |  | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HubspotDisconnect

> interface{} HubspotDisconnect(ctx).Execute()

Hubspot Disconnect

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
	resp, r, err := apiClient.TenantAPI.HubspotDisconnect(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.HubspotDisconnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HubspotDisconnect`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.HubspotDisconnect`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHubspotDisconnectRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HubspotOauthCallback

> interface{} HubspotOauthCallback(ctx).Code(code).State(state).Execute()

Hubspot Oauth Callback

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
	code := "code_example" // string | 
	state := "state_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.HubspotOauthCallback(context.Background()).Code(code).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.HubspotOauthCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HubspotOauthCallback`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.HubspotOauthCallback`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHubspotOauthCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **string** |  | 
 **state** | **string** |  | 

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


## HubspotTestConnection

> interface{} HubspotTestConnection(ctx).Execute()

Hubspot Test Connection

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
	resp, r, err := apiClient.TenantAPI.HubspotTestConnection(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.HubspotTestConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HubspotTestConnection`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.HubspotTestConnection`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHubspotTestConnectionRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssueIntegrationKey

> interface{} IssueIntegrationKey(ctx, provider).Execute()

Issue Integration Key



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
	provider := "provider_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.IssueIntegrationKey(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.IssueIntegrationKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssueIntegrationKey`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.IssueIntegrationKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssueIntegrationKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTemplates

> []TemplateItem ListTemplates(ctx).Execute()

List Templates



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
	resp, r, err := apiClient.TenantAPI.ListTemplates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.ListTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTemplates`: []TemplateItem
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.ListTemplates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListTemplatesRequest struct via the builder pattern


### Return type

[**[]TemplateItem**](TemplateItem.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetsuiteConnect

> interface{} NetsuiteConnect(ctx).NetSuiteConnectRequest(netSuiteConnectRequest).Execute()

Netsuite Connect

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
	netSuiteConnectRequest := *openapiclient.NewNetSuiteConnectRequest("AccountId_example", "ClientId_example", "ClientSecret_example", "SyncDirection_example", "SyncFrequency_example") // NetSuiteConnectRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.NetsuiteConnect(context.Background()).NetSuiteConnectRequest(netSuiteConnectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.NetsuiteConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetsuiteConnect`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.NetsuiteConnect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNetsuiteConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **netSuiteConnectRequest** | [**NetSuiteConnectRequest**](NetSuiteConnectRequest.md) |  | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetsuiteDisconnect

> interface{} NetsuiteDisconnect(ctx).Execute()

Netsuite Disconnect

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
	resp, r, err := apiClient.TenantAPI.NetsuiteDisconnect(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.NetsuiteDisconnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetsuiteDisconnect`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.NetsuiteDisconnect`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNetsuiteDisconnectRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetsuiteOauthCallback

> interface{} NetsuiteOauthCallback(ctx).Code(code).State(state).Execute()

Netsuite Oauth Callback

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
	code := "code_example" // string | 
	state := "state_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.NetsuiteOauthCallback(context.Background()).Code(code).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.NetsuiteOauthCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetsuiteOauthCallback`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.NetsuiteOauthCallback`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNetsuiteOauthCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **string** |  | 
 **state** | **string** |  | 

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


## NetsuiteTestConnection

> interface{} NetsuiteTestConnection(ctx).Execute()

Netsuite Test Connection

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
	resp, r, err := apiClient.TenantAPI.NetsuiteTestConnection(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.NetsuiteTestConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetsuiteTestConnection`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.NetsuiteTestConnection`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNetsuiteTestConnectionRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiKeyV1

> CreateApiKeyResponse PostApiKeyV1(ctx).CreateApiKeyRequest(createApiKeyRequest).Execute()

Post Api Key



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
	createApiKeyRequest := *openapiclient.NewCreateApiKeyRequest() // CreateApiKeyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.PostApiKeyV1(context.Background()).CreateApiKeyRequest(createApiKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.PostApiKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApiKeyV1`: CreateApiKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.PostApiKeyV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiKeyV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createApiKeyRequest** | [**CreateApiKeyRequest**](CreateApiKeyRequest.md) |  | 

### Return type

[**CreateApiKeyResponse**](CreateApiKeyResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDomainVerify

> DomainResponse PostDomainVerify(ctx).Execute()

Post Domain Verify



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
	resp, r, err := apiClient.TenantAPI.PostDomainVerify(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.PostDomainVerify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDomainVerify`: DomainResponse
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.PostDomainVerify`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostDomainVerifyRequest struct via the builder pattern


### Return type

[**DomainResponse**](DomainResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPortalTokenV1

> CreatePortalTokenResponse PostPortalTokenV1(ctx).AppApiV1TenantAccessRouterCreatePortalTokenRequest(appApiV1TenantAccessRouterCreatePortalTokenRequest).Execute()

Post Portal Token



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
	appApiV1TenantAccessRouterCreatePortalTokenRequest := *openapiclient.NewAppApiV1TenantAccessRouterCreatePortalTokenRequest("Subject_example") // AppApiV1TenantAccessRouterCreatePortalTokenRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.PostPortalTokenV1(context.Background()).AppApiV1TenantAccessRouterCreatePortalTokenRequest(appApiV1TenantAccessRouterCreatePortalTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.PostPortalTokenV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPortalTokenV1`: CreatePortalTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.PostPortalTokenV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPortalTokenV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appApiV1TenantAccessRouterCreatePortalTokenRequest** | [**AppApiV1TenantAccessRouterCreatePortalTokenRequest**](AppApiV1TenantAccessRouterCreatePortalTokenRequest.md) |  | 

### Return type

[**CreatePortalTokenResponse**](CreatePortalTokenResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRevokeApiKeyV1

> interface{} PostRevokeApiKeyV1(ctx, keyId).RevokeApiKeyRequest(revokeApiKeyRequest).Execute()

Post Revoke Api Key



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
	keyId := "keyId_example" // string | 
	revokeApiKeyRequest := *openapiclient.NewRevokeApiKeyRequest() // RevokeApiKeyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.PostRevokeApiKeyV1(context.Background(), keyId).RevokeApiKeyRequest(revokeApiKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.PostRevokeApiKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRevokeApiKeyV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.PostRevokeApiKeyV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRevokeApiKeyV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revokeApiKeyRequest** | [**RevokeApiKeyRequest**](RevokeApiKeyRequest.md) |  | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreviewTemplate

> TemplatePreviewResponse PreviewTemplate(ctx, templateId).TemplatePreviewRequest(templatePreviewRequest).Execute()

Preview Template



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
	templateId := "templateId_example" // string | 
	templatePreviewRequest := *openapiclient.NewTemplatePreviewRequest() // TemplatePreviewRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.PreviewTemplate(context.Background(), templateId).TemplatePreviewRequest(templatePreviewRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.PreviewTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PreviewTemplate`: TemplatePreviewResponse
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.PreviewTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **templatePreviewRequest** | [**TemplatePreviewRequest**](TemplatePreviewRequest.md) |  | 

### Return type

[**TemplatePreviewResponse**](TemplatePreviewResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutDomain

> DomainResponse PutDomain(ctx).DomainUpdate(domainUpdate).Execute()

Put Domain



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
	domainUpdate := *openapiclient.NewDomainUpdate() // DomainUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.PutDomain(context.Background()).DomainUpdate(domainUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.PutDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutDomain`: DomainResponse
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.PutDomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainUpdate** | [**DomainUpdate**](DomainUpdate.md) |  | 

### Return type

[**DomainResponse**](DomainResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFeatureFlagV1

> FeatureFlagItem PutFeatureFlagV1(ctx, flagId).FeatureFlagUpdate(featureFlagUpdate).Execute()

Put Feature Flag



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
	flagId := "flagId_example" // string | 
	featureFlagUpdate := *openapiclient.NewFeatureFlagUpdate(false) // FeatureFlagUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.PutFeatureFlagV1(context.Background(), flagId).FeatureFlagUpdate(featureFlagUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.PutFeatureFlagV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutFeatureFlagV1`: FeatureFlagItem
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.PutFeatureFlagV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flagId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFeatureFlagV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **featureFlagUpdate** | [**FeatureFlagUpdate**](FeatureFlagUpdate.md) |  | 

### Return type

[**FeatureFlagItem**](FeatureFlagItem.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutIntegration

> ConnectorItem PutIntegration(ctx, provider).IntegrationUpdate(integrationUpdate).Execute()

Put Integration



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
	provider := "provider_example" // string | 
	integrationUpdate := *openapiclient.NewIntegrationUpdate() // IntegrationUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.PutIntegration(context.Background(), provider).IntegrationUpdate(integrationUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.PutIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutIntegration`: ConnectorItem
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.PutIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **integrationUpdate** | [**IntegrationUpdate**](IntegrationUpdate.md) |  | 

### Return type

[**ConnectorItem**](ConnectorItem.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSender

> SenderResponse PutSender(ctx).SenderUpdate(senderUpdate).Execute()

Put Sender



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
	senderUpdate := *openapiclient.NewSenderUpdate() // SenderUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.PutSender(context.Background()).SenderUpdate(senderUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.PutSender``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSender`: SenderResponse
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.PutSender`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutSenderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **senderUpdate** | [**SenderUpdate**](SenderUpdate.md) |  | 

### Return type

[**SenderResponse**](SenderResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutTemplate

> TemplateItem PutTemplate(ctx, templateId).TemplateUpdate(templateUpdate).Execute()

Put Template



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
	templateId := "templateId_example" // string | 
	templateUpdate := *openapiclient.NewTemplateUpdate() // TemplateUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.PutTemplate(context.Background(), templateId).TemplateUpdate(templateUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.PutTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutTemplate`: TemplateItem
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.PutTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **templateUpdate** | [**TemplateUpdate**](TemplateUpdate.md) |  | 

### Return type

[**TemplateItem**](TemplateItem.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutWebhook

> WebhookResponse PutWebhook(ctx).WebhookUpdate(webhookUpdate).Execute()

Put Webhook



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
	webhookUpdate := *openapiclient.NewWebhookUpdate() // WebhookUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.PutWebhook(context.Background()).WebhookUpdate(webhookUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.PutWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutWebhook`: WebhookResponse
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.PutWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookUpdate** | [**WebhookUpdate**](WebhookUpdate.md) |  | 

### Return type

[**WebhookResponse**](WebhookResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuickbooksConnect

> interface{} QuickbooksConnect(ctx).QuickBooksConnectRequest(quickBooksConnectRequest).Execute()

Quickbooks Connect

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
	quickBooksConnectRequest := *openapiclient.NewQuickBooksConnectRequest("Environment_example", "SyncDirection_example", "SyncFrequency_example") // QuickBooksConnectRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.QuickbooksConnect(context.Background()).QuickBooksConnectRequest(quickBooksConnectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.QuickbooksConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuickbooksConnect`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.QuickbooksConnect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQuickbooksConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quickBooksConnectRequest** | [**QuickBooksConnectRequest**](QuickBooksConnectRequest.md) |  | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuickbooksDisconnect

> interface{} QuickbooksDisconnect(ctx).Execute()

Quickbooks Disconnect

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
	resp, r, err := apiClient.TenantAPI.QuickbooksDisconnect(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.QuickbooksDisconnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuickbooksDisconnect`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.QuickbooksDisconnect`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiQuickbooksDisconnectRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuickbooksOauthCallback

> interface{} QuickbooksOauthCallback(ctx).Code(code).State(state).RealmId(realmId).Execute()

Quickbooks Oauth Callback

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
	code := "code_example" // string | 
	state := "state_example" // string | 
	realmId := "realmId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.QuickbooksOauthCallback(context.Background()).Code(code).State(state).RealmId(realmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.QuickbooksOauthCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuickbooksOauthCallback`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.QuickbooksOauthCallback`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQuickbooksOauthCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **string** |  | 
 **state** | **string** |  | 
 **realmId** | **string** |  | 

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


## QuickbooksTestConnection

> interface{} QuickbooksTestConnection(ctx).Execute()

Quickbooks Test Connection

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
	resp, r, err := apiClient.TenantAPI.QuickbooksTestConnection(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.QuickbooksTestConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuickbooksTestConnection`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.QuickbooksTestConnection`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiQuickbooksTestConnectionRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetryFailedDeliveries

> interface{} RetryFailedDeliveries(ctx).Execute()

Retry Failed Deliveries



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
	resp, r, err := apiClient.TenantAPI.RetryFailedDeliveries(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.RetryFailedDeliveries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetryFailedDeliveries`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.RetryFailedDeliveries`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRetryFailedDeliveriesRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeAllApiKeysV1TenantDanger

> interface{} RevokeAllApiKeysV1TenantDanger(ctx).DangerRequest(dangerRequest).Execute()

Revoke All Api Keys



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
	dangerRequest := *openapiclient.NewDangerRequest("AuditNote_example") // DangerRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.RevokeAllApiKeysV1TenantDanger(context.Background()).DangerRequest(dangerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.RevokeAllApiKeysV1TenantDanger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RevokeAllApiKeysV1TenantDanger`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.RevokeAllApiKeysV1TenantDanger`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRevokeAllApiKeysV1TenantDangerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dangerRequest** | [**DangerRequest**](DangerRequest.md) |  | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeAllPortalTokensV1TenantDanger

> interface{} RevokeAllPortalTokensV1TenantDanger(ctx).DangerRequest(dangerRequest).Execute()

Revoke All Portal Tokens



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
	dangerRequest := *openapiclient.NewDangerRequest("AuditNote_example") // DangerRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.RevokeAllPortalTokensV1TenantDanger(context.Background()).DangerRequest(dangerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.RevokeAllPortalTokensV1TenantDanger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RevokeAllPortalTokensV1TenantDanger`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.RevokeAllPortalTokensV1TenantDanger`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRevokeAllPortalTokensV1TenantDangerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dangerRequest** | [**DangerRequest**](DangerRequest.md) |  | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SalesforceConnect

> interface{} SalesforceConnect(ctx).SalesforceConnectRequest(salesforceConnectRequest).Execute()

Salesforce Connect

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
	salesforceConnectRequest := *openapiclient.NewSalesforceConnectRequest("Environment_example", "SyncDirection_example", "SyncFrequency_example") // SalesforceConnectRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.SalesforceConnect(context.Background()).SalesforceConnectRequest(salesforceConnectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.SalesforceConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SalesforceConnect`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.SalesforceConnect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSalesforceConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **salesforceConnectRequest** | [**SalesforceConnectRequest**](SalesforceConnectRequest.md) |  | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SalesforceDisconnect

> interface{} SalesforceDisconnect(ctx).Execute()

Salesforce Disconnect

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
	resp, r, err := apiClient.TenantAPI.SalesforceDisconnect(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.SalesforceDisconnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SalesforceDisconnect`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.SalesforceDisconnect`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSalesforceDisconnectRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SalesforceOauthCallback

> interface{} SalesforceOauthCallback(ctx).Code(code).State(state).Execute()

Salesforce Oauth Callback

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
	code := "code_example" // string | 
	state := "state_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.SalesforceOauthCallback(context.Background()).Code(code).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.SalesforceOauthCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SalesforceOauthCallback`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.SalesforceOauthCallback`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSalesforceOauthCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **string** |  | 
 **state** | **string** |  | 

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


## SalesforceTestConnection

> interface{} SalesforceTestConnection(ctx).Execute()

Salesforce Test Connection

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
	resp, r, err := apiClient.TenantAPI.SalesforceTestConnection(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.SalesforceTestConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SalesforceTestConnection`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.SalesforceTestConnection`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSalesforceTestConnectionRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestSendTemplateV1

> TemplateTestSendResponse TestSendTemplateV1(ctx, templateId).TemplateTestSendRequest(templateTestSendRequest).Execute()

Test Send Template



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
	templateId := "templateId_example" // string | 
	templateTestSendRequest := *openapiclient.NewTemplateTestSendRequest() // TemplateTestSendRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.TestSendTemplateV1(context.Background(), templateId).TemplateTestSendRequest(templateTestSendRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.TestSendTemplateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestSendTemplateV1`: TemplateTestSendResponse
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.TestSendTemplateV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestSendTemplateV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **templateTestSendRequest** | [**TemplateTestSendRequest**](TemplateTestSendRequest.md) |  | 

### Return type

[**TemplateTestSendResponse**](TemplateTestSendResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestWebhook

> interface{} TestWebhook(ctx).Execute()

Test Webhook



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
	resp, r, err := apiClient.TenantAPI.TestWebhook(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.TestWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestWebhook`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.TestWebhook`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTestWebhookRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRevenueRecognitionV1

> interface{} UpdateRevenueRecognitionV1(ctx).RevenueRecognitionRequest(revenueRecognitionRequest).Execute()

Update Revenue Recognition



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
	revenueRecognitionRequest := *openapiclient.NewRevenueRecognitionRequest() // RevenueRecognitionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.UpdateRevenueRecognitionV1(context.Background()).RevenueRecognitionRequest(revenueRecognitionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.UpdateRevenueRecognitionV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRevenueRecognitionV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.UpdateRevenueRecognitionV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRevenueRecognitionV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revenueRecognitionRequest** | [**RevenueRecognitionRequest**](RevenueRecognitionRequest.md) |  | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## XeroConnect

> interface{} XeroConnect(ctx).XeroConnectRequest(xeroConnectRequest).Execute()

Xero Connect

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
	xeroConnectRequest := *openapiclient.NewXeroConnectRequest("SyncDirection_example", "SyncFrequency_example") // XeroConnectRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.XeroConnect(context.Background()).XeroConnectRequest(xeroConnectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.XeroConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `XeroConnect`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.XeroConnect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiXeroConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroConnectRequest** | [**XeroConnectRequest**](XeroConnectRequest.md) |  | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## XeroDisconnect

> interface{} XeroDisconnect(ctx).Execute()

Xero Disconnect

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
	resp, r, err := apiClient.TenantAPI.XeroDisconnect(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.XeroDisconnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `XeroDisconnect`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.XeroDisconnect`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiXeroDisconnectRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## XeroOauthCallback

> interface{} XeroOauthCallback(ctx).Code(code).State(state).Execute()

Xero Oauth Callback

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
	code := "code_example" // string | 
	state := "state_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.XeroOauthCallback(context.Background()).Code(code).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.XeroOauthCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `XeroOauthCallback`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.XeroOauthCallback`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiXeroOauthCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **string** |  | 
 **state** | **string** |  | 

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


## XeroTestConnection

> interface{} XeroTestConnection(ctx).Execute()

Xero Test Connection

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
	resp, r, err := apiClient.TenantAPI.XeroTestConnection(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.XeroTestConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `XeroTestConnection`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.XeroTestConnection`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiXeroTestConnectionRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

