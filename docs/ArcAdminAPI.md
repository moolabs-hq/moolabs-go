# \ArcAdminAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BackfillMalformedDisputesV1ArcAdmin**](ArcAdminAPI.md#BackfillMalformedDisputesV1ArcAdmin) | **Post** /v1/arc/admin/inbound-routing/backfill-malformed-disputes | Repair AI-classified disputes the pre-T4 router mis-routed
[**BulkReplayEndpointV1Arc**](ArcAdminAPI.md#BulkReplayEndpointV1Arc) | **Post** /v1/arc/admin/dead-letters/bulk-replay | Bulk Replay Endpoint
[**GetAgentRunStatusV1**](ArcAdminAPI.md#GetAgentRunStatusV1) | **Get** /v1/arc/admin/agent-runs/{run_id} | Get agent run status
[**GetDeadLettersV1**](ArcAdminAPI.md#GetDeadLettersV1) | **Get** /v1/arc/admin/dead-letters | Get Dead Letters
[**GetEmailConfigV1**](ArcAdminAPI.md#GetEmailConfigV1) | **Get** /v1/arc/admin/email-config | Get Email Config
[**GetOutboundFailuresV1**](ArcAdminAPI.md#GetOutboundFailuresV1) | **Get** /v1/arc/admin/outbound-failures | Get Outbound Failures
[**GetPotentialDuplicatesV1**](ArcAdminAPI.md#GetPotentialDuplicatesV1) | **Get** /v1/arc/admin/potential-duplicates | Get Potential Duplicates
[**GetTenantEscalationConfigV1Arc**](ArcAdminAPI.md#GetTenantEscalationConfigV1Arc) | **Get** /v1/arc/admin/tenant-escalation-config | Get Tenant Escalation Config
[**ImportDisputesEndpoint**](ArcAdminAPI.md#ImportDisputesEndpoint) | **Post** /v1/arc/admin/import/disputes | Import Disputes Endpoint
[**ImportInvoicesEndpoint**](ArcAdminAPI.md#ImportInvoicesEndpoint) | **Post** /v1/arc/admin/import/invoices | Import Invoices Endpoint
[**ImportPromisesEndpoint**](ArcAdminAPI.md#ImportPromisesEndpoint) | **Post** /v1/arc/admin/import/promises | Import Promises Endpoint
[**MergeRemittancesEndpoint**](ArcAdminAPI.md#MergeRemittancesEndpoint) | **Post** /v1/arc/admin/remittances/merge | Merge Remittances Endpoint
[**PatchTenantEscalationConfigV1Arc**](ArcAdminAPI.md#PatchTenantEscalationConfigV1Arc) | **Patch** /v1/arc/admin/tenant-escalation-config | Patch Tenant Escalation Config
[**ReclassifyCommunication**](ArcAdminAPI.md#ReclassifyCommunication) | **Post** /v1/arc/admin/communications/{communication_id}/reclassify | Re-enqueue inbound classification for an existing communication
[**ReplayDeadLetterEndpointV1**](ArcAdminAPI.md#ReplayDeadLetterEndpointV1) | **Post** /v1/arc/admin/dead-letters/{event_id}/replay | Replay Dead Letter Endpoint
[**RetryOutboundFailureEndpointV1**](ArcAdminAPI.md#RetryOutboundFailureEndpointV1) | **Post** /v1/arc/admin/outbound-failures/{event_id}/retry | Retry Outbound Failure Endpoint
[**RotateInboundSecretV1Arc**](ArcAdminAPI.md#RotateInboundSecretV1Arc) | **Post** /v1/arc/admin/email-config/rotate-secret | Rotate Inbound Secret
[**RunAgentSync**](ArcAdminAPI.md#RunAgentSync) | **Post** /v1/arc/admin/agents/{agent_type}/run | Run Agent Synchronously
[**SeedTenantEndpoint**](ArcAdminAPI.md#SeedTenantEndpoint) | **Post** /v1/arc/admin/seed | Seed Tenant Endpoint
[**UpsertEmailConfigV1**](ArcAdminAPI.md#UpsertEmailConfigV1) | **Put** /v1/arc/admin/email-config | Upsert Email Config
[**VerifyEmailConfigV1**](ArcAdminAPI.md#VerifyEmailConfigV1) | **Post** /v1/arc/admin/email-config/verify | Verify Email Config
[**VoidRemittanceEndpoint**](ArcAdminAPI.md#VoidRemittanceEndpoint) | **Post** /v1/arc/admin/remittances/{remittance_id}/void | Void Remittance Endpoint



## BackfillMalformedDisputesV1ArcAdmin

> BackfillMalformedDisputesResponse BackfillMalformedDisputesV1ArcAdmin(ctx).BackfillMalformedDisputesRequest(backfillMalformedDisputesRequest).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()

Repair AI-classified disputes the pre-T4 router mis-routed



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
	backfillMalformedDisputesRequest := *openapiclient.NewBackfillMalformedDisputesRequest() // BackfillMalformedDisputesRequest | 
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArcAdminAPI.BackfillMalformedDisputesV1ArcAdmin(context.Background()).BackfillMalformedDisputesRequest(backfillMalformedDisputesRequest).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArcAdminAPI.BackfillMalformedDisputesV1ArcAdmin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackfillMalformedDisputesV1ArcAdmin`: BackfillMalformedDisputesResponse
	fmt.Fprintf(os.Stdout, "Response from `ArcAdminAPI.BackfillMalformedDisputesV1ArcAdmin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackfillMalformedDisputesV1ArcAdminRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **backfillMalformedDisputesRequest** | [**BackfillMalformedDisputesRequest**](BackfillMalformedDisputesRequest.md) |  | 
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **authorization** | **string** |  | 

### Return type

[**BackfillMalformedDisputesResponse**](BackfillMalformedDisputesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkReplayEndpointV1Arc

> interface{} BulkReplayEndpointV1Arc(ctx).BulkReplayRequest(bulkReplayRequest).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()

Bulk Replay Endpoint



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
	bulkReplayRequest := *openapiclient.NewBulkReplayRequest() // BulkReplayRequest | 
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArcAdminAPI.BulkReplayEndpointV1Arc(context.Background()).BulkReplayRequest(bulkReplayRequest).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArcAdminAPI.BulkReplayEndpointV1Arc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkReplayEndpointV1Arc`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArcAdminAPI.BulkReplayEndpointV1Arc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkReplayEndpointV1ArcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkReplayRequest** | [**BulkReplayRequest**](BulkReplayRequest.md) |  | 
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **authorization** | **string** |  | 

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


## GetAgentRunStatusV1

> interface{} GetAgentRunStatusV1(ctx, runId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()

Get agent run status

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
	runId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArcAdminAPI.GetAgentRunStatusV1(context.Background(), runId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArcAdminAPI.GetAgentRunStatusV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentRunStatusV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArcAdminAPI.GetAgentRunStatusV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentRunStatusV1Request struct via the builder pattern


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


## GetDeadLettersV1

> interface{} GetDeadLettersV1(ctx).Page(page).PageSize(pageSize).Source(source).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()

Get Dead Letters



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
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 50)
	source := "source_example" // string |  (optional)
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArcAdminAPI.GetDeadLettersV1(context.Background()).Page(page).PageSize(pageSize).Source(source).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArcAdminAPI.GetDeadLettersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeadLettersV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArcAdminAPI.GetDeadLettersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeadLettersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 50]
 **source** | **string** |  | 
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


## GetEmailConfigV1

> EmailConfigOut GetEmailConfigV1(ctx).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()

Get Email Config



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
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArcAdminAPI.GetEmailConfigV1(context.Background()).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArcAdminAPI.GetEmailConfigV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmailConfigV1`: EmailConfigOut
	fmt.Fprintf(os.Stdout, "Response from `ArcAdminAPI.GetEmailConfigV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailConfigV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **authorization** | **string** |  | 

### Return type

[**EmailConfigOut**](EmailConfigOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOutboundFailuresV1

> interface{} GetOutboundFailuresV1(ctx).Page(page).PageSize(pageSize).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()

Get Outbound Failures



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
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 50)
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArcAdminAPI.GetOutboundFailuresV1(context.Background()).Page(page).PageSize(pageSize).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArcAdminAPI.GetOutboundFailuresV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOutboundFailuresV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArcAdminAPI.GetOutboundFailuresV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOutboundFailuresV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 50]
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


## GetPotentialDuplicatesV1

> interface{} GetPotentialDuplicatesV1(ctx).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()

Get Potential Duplicates



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
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArcAdminAPI.GetPotentialDuplicatesV1(context.Background()).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArcAdminAPI.GetPotentialDuplicatesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPotentialDuplicatesV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArcAdminAPI.GetPotentialDuplicatesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPotentialDuplicatesV1Request struct via the builder pattern


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


## GetTenantEscalationConfigV1Arc

> TenantEscalationConfigOut GetTenantEscalationConfigV1Arc(ctx).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()

Get Tenant Escalation Config



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
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArcAdminAPI.GetTenantEscalationConfigV1Arc(context.Background()).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArcAdminAPI.GetTenantEscalationConfigV1Arc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantEscalationConfigV1Arc`: TenantEscalationConfigOut
	fmt.Fprintf(os.Stdout, "Response from `ArcAdminAPI.GetTenantEscalationConfigV1Arc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantEscalationConfigV1ArcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **authorization** | **string** |  | 

### Return type

[**TenantEscalationConfigOut**](TenantEscalationConfigOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportDisputesEndpoint

> interface{} ImportDisputesEndpoint(ctx).ImportDisputesRequest(importDisputesRequest).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()

Import Disputes Endpoint



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
	importDisputesRequest := *openapiclient.NewImportDisputesRequest([]openapiclient.ImportDisputeItem{*openapiclient.NewImportDisputeItem("CaseId_example", "InvoiceId_example", int32(123))}) // ImportDisputesRequest | 
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArcAdminAPI.ImportDisputesEndpoint(context.Background()).ImportDisputesRequest(importDisputesRequest).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArcAdminAPI.ImportDisputesEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportDisputesEndpoint`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArcAdminAPI.ImportDisputesEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportDisputesEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importDisputesRequest** | [**ImportDisputesRequest**](ImportDisputesRequest.md) |  | 
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **authorization** | **string** |  | 

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


## ImportInvoicesEndpoint

> interface{} ImportInvoicesEndpoint(ctx).ImportInvoicesRequest(importInvoicesRequest).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()

Import Invoices Endpoint



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
	importInvoicesRequest := *openapiclient.NewImportInvoicesRequest([]openapiclient.ImportInvoiceItem{*openapiclient.NewImportInvoiceItem("CaseId_example", "InvoiceId_example", int32(123), time.Now())}) // ImportInvoicesRequest | 
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArcAdminAPI.ImportInvoicesEndpoint(context.Background()).ImportInvoicesRequest(importInvoicesRequest).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArcAdminAPI.ImportInvoicesEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportInvoicesEndpoint`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArcAdminAPI.ImportInvoicesEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportInvoicesEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importInvoicesRequest** | [**ImportInvoicesRequest**](ImportInvoicesRequest.md) |  | 
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **authorization** | **string** |  | 

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


## ImportPromisesEndpoint

> interface{} ImportPromisesEndpoint(ctx).ImportPromisesRequest(importPromisesRequest).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()

Import Promises Endpoint



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
	importPromisesRequest := *openapiclient.NewImportPromisesRequest([]openapiclient.ImportPromiseItem{*openapiclient.NewImportPromiseItem("ExternalPtpId_example", "CaseId_example", int32(123), time.Now())}) // ImportPromisesRequest | 
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArcAdminAPI.ImportPromisesEndpoint(context.Background()).ImportPromisesRequest(importPromisesRequest).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArcAdminAPI.ImportPromisesEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportPromisesEndpoint`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArcAdminAPI.ImportPromisesEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportPromisesEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importPromisesRequest** | [**ImportPromisesRequest**](ImportPromisesRequest.md) |  | 
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **authorization** | **string** |  | 

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


## MergeRemittancesEndpoint

> interface{} MergeRemittancesEndpoint(ctx).MergeRemittancesRequest(mergeRemittancesRequest).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()

Merge Remittances Endpoint



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
	mergeRemittancesRequest := *openapiclient.NewMergeRemittancesRequest("SourceId_example", "TargetId_example") // MergeRemittancesRequest | 
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArcAdminAPI.MergeRemittancesEndpoint(context.Background()).MergeRemittancesRequest(mergeRemittancesRequest).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArcAdminAPI.MergeRemittancesEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MergeRemittancesEndpoint`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArcAdminAPI.MergeRemittancesEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMergeRemittancesEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mergeRemittancesRequest** | [**MergeRemittancesRequest**](MergeRemittancesRequest.md) |  | 
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **authorization** | **string** |  | 

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


## PatchTenantEscalationConfigV1Arc

> TenantEscalationConfigOut PatchTenantEscalationConfigV1Arc(ctx).TenantEscalationConfigPatch(tenantEscalationConfigPatch).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).XUserId(xUserId).Execute()

Patch Tenant Escalation Config



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
	tenantEscalationConfigPatch := *openapiclient.NewTenantEscalationConfigPatch(map[string]interface{}{"key": interface{}(123)}, "ChangeReason_example") // TenantEscalationConfigPatch | 
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)
	xUserId := "xUserId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArcAdminAPI.PatchTenantEscalationConfigV1Arc(context.Background()).TenantEscalationConfigPatch(tenantEscalationConfigPatch).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).XUserId(xUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArcAdminAPI.PatchTenantEscalationConfigV1Arc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchTenantEscalationConfigV1Arc`: TenantEscalationConfigOut
	fmt.Fprintf(os.Stdout, "Response from `ArcAdminAPI.PatchTenantEscalationConfigV1Arc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchTenantEscalationConfigV1ArcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantEscalationConfigPatch** | [**TenantEscalationConfigPatch**](TenantEscalationConfigPatch.md) |  | 
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **authorization** | **string** |  | 
 **xUserId** | **string** |  | 

### Return type

[**TenantEscalationConfigOut**](TenantEscalationConfigOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReclassifyCommunication

> interface{} ReclassifyCommunication(ctx, communicationId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()

Re-enqueue inbound classification for an existing communication



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
	communicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArcAdminAPI.ReclassifyCommunication(context.Background(), communicationId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArcAdminAPI.ReclassifyCommunication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReclassifyCommunication`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArcAdminAPI.ReclassifyCommunication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**communicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReclassifyCommunicationRequest struct via the builder pattern


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


## ReplayDeadLetterEndpointV1

> interface{} ReplayDeadLetterEndpointV1(ctx, eventId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()

Replay Dead Letter Endpoint



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
	eventId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArcAdminAPI.ReplayDeadLetterEndpointV1(context.Background(), eventId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArcAdminAPI.ReplayDeadLetterEndpointV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplayDeadLetterEndpointV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArcAdminAPI.ReplayDeadLetterEndpointV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplayDeadLetterEndpointV1Request struct via the builder pattern


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


## RetryOutboundFailureEndpointV1

> interface{} RetryOutboundFailureEndpointV1(ctx, eventId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()

Retry Outbound Failure Endpoint



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
	eventId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArcAdminAPI.RetryOutboundFailureEndpointV1(context.Background(), eventId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArcAdminAPI.RetryOutboundFailureEndpointV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetryOutboundFailureEndpointV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArcAdminAPI.RetryOutboundFailureEndpointV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetryOutboundFailureEndpointV1Request struct via the builder pattern


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


## RotateInboundSecretV1Arc

> EmailConfigOut RotateInboundSecretV1Arc(ctx).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()

Rotate Inbound Secret



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
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArcAdminAPI.RotateInboundSecretV1Arc(context.Background()).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArcAdminAPI.RotateInboundSecretV1Arc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RotateInboundSecretV1Arc`: EmailConfigOut
	fmt.Fprintf(os.Stdout, "Response from `ArcAdminAPI.RotateInboundSecretV1Arc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRotateInboundSecretV1ArcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **authorization** | **string** |  | 

### Return type

[**EmailConfigOut**](EmailConfigOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunAgentSync

> interface{} RunAgentSync(ctx, agentType).CaseId(caseId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()

Run Agent Synchronously



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
	caseId := "caseId_example" // string | Optional: run only for this case (optional)
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArcAdminAPI.RunAgentSync(context.Background(), agentType).CaseId(caseId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArcAdminAPI.RunAgentSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RunAgentSync`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArcAdminAPI.RunAgentSync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRunAgentSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **caseId** | **string** | Optional: run only for this case | 
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


## SeedTenantEndpoint

> interface{} SeedTenantEndpoint(ctx).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()

Seed Tenant Endpoint



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
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArcAdminAPI.SeedTenantEndpoint(context.Background()).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArcAdminAPI.SeedTenantEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SeedTenantEndpoint`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArcAdminAPI.SeedTenantEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSeedTenantEndpointRequest struct via the builder pattern


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


## UpsertEmailConfigV1

> EmailConfigOut UpsertEmailConfigV1(ctx).EmailConfigUpsert(emailConfigUpsert).ForceSecretOverwrite(forceSecretOverwrite).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()

Upsert Email Config



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
	emailConfigUpsert := *openapiclient.NewEmailConfigUpsert("SenderDomain_example") // EmailConfigUpsert | 
	forceSecretOverwrite := true // bool | Break-glass: allow this PUT to overwrite an existing non-empty ``inbound_secret`` with a different value. Without this flag the server returns 409 to prevent the smoke-harness drift that silently breaks inbound webhooks for real tenants. Use POST /email-config/rotate-secret for the production rotation flow. (optional) (default to false)
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArcAdminAPI.UpsertEmailConfigV1(context.Background()).EmailConfigUpsert(emailConfigUpsert).ForceSecretOverwrite(forceSecretOverwrite).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArcAdminAPI.UpsertEmailConfigV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpsertEmailConfigV1`: EmailConfigOut
	fmt.Fprintf(os.Stdout, "Response from `ArcAdminAPI.UpsertEmailConfigV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpsertEmailConfigV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailConfigUpsert** | [**EmailConfigUpsert**](EmailConfigUpsert.md) |  | 
 **forceSecretOverwrite** | **bool** | Break-glass: allow this PUT to overwrite an existing non-empty &#x60;&#x60;inbound_secret&#x60;&#x60; with a different value. Without this flag the server returns 409 to prevent the smoke-harness drift that silently breaks inbound webhooks for real tenants. Use POST /email-config/rotate-secret for the production rotation flow. | [default to false]
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **authorization** | **string** |  | 

### Return type

[**EmailConfigOut**](EmailConfigOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyEmailConfigV1

> EmailConfigOut VerifyEmailConfigV1(ctx).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()

Verify Email Config



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
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArcAdminAPI.VerifyEmailConfigV1(context.Background()).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArcAdminAPI.VerifyEmailConfigV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyEmailConfigV1`: EmailConfigOut
	fmt.Fprintf(os.Stdout, "Response from `ArcAdminAPI.VerifyEmailConfigV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyEmailConfigV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **authorization** | **string** |  | 

### Return type

[**EmailConfigOut**](EmailConfigOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VoidRemittanceEndpoint

> interface{} VoidRemittanceEndpoint(ctx, remittanceId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()

Void Remittance Endpoint



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
	remittanceId := "remittanceId_example" // string | 
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArcAdminAPI.VoidRemittanceEndpoint(context.Background(), remittanceId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArcAdminAPI.VoidRemittanceEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VoidRemittanceEndpoint`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArcAdminAPI.VoidRemittanceEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**remittanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVoidRemittanceEndpointRequest struct via the builder pattern


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

