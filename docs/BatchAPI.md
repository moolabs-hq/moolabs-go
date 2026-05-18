# \BatchAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchBackfillManagedPortalSuppressionAckProtectionV1ArcBatch**](BatchAPI.md#BatchBackfillManagedPortalSuppressionAckProtectionV1ArcBatch) | **Post** /v1/arc/batch/backfill-managed-portal-suppression/protection/ack | Batch Backfill Managed Portal Suppression Ack Protection
[**BatchBackfillManagedPortalSuppressionCancelV1ArcBatch**](BatchAPI.md#BatchBackfillManagedPortalSuppressionCancelV1ArcBatch) | **Post** /v1/arc/batch/backfill-managed-portal-suppression/{job_id}/cancel | Batch Backfill Managed Portal Suppression Cancel
[**BatchBackfillManagedPortalSuppressionRunV1ArcBatch**](BatchAPI.md#BatchBackfillManagedPortalSuppressionRunV1ArcBatch) | **Post** /v1/arc/batch/backfill-managed-portal-suppression/run | Batch Backfill Managed Portal Suppression Run
[**BatchBackfillManagedPortalSuppressionRunV1ArcBatchBackfill**](BatchAPI.md#BatchBackfillManagedPortalSuppressionRunV1ArcBatchBackfill) | **Post** /v1/arc/batch/backfill-managed-portal-suppression/run-pending | Batch Backfill Managed Portal Suppression Run
[**BatchBackfillManagedPortalSuppressionStatusV1ArcBatch**](BatchAPI.md#BatchBackfillManagedPortalSuppressionStatusV1ArcBatch) | **Get** /v1/arc/batch/backfill-managed-portal-suppression/{job_id} | Batch Backfill Managed Portal Suppression Status
[**BatchBackfillManagedPortalSuppressionV1ArcBatch**](BatchAPI.md#BatchBackfillManagedPortalSuppressionV1ArcBatch) | **Post** /v1/arc/batch/backfill-managed-portal-suppression | Batch Backfill Managed Portal Suppression
[**BatchBackfillNetsuiteDeferredInvoicesV1ArcBatch**](BatchAPI.md#BatchBackfillNetsuiteDeferredInvoicesV1ArcBatch) | **Post** /v1/arc/batch/backfill-netsuite-deferred-invoices | Batch Backfill Netsuite Deferred Invoices
[**BatchBackfillNetsuiteIdentitiesV1Arc**](BatchAPI.md#BatchBackfillNetsuiteIdentitiesV1Arc) | **Post** /v1/arc/batch/backfill-netsuite-identities | Batch Backfill Netsuite Identities
[**BatchBootstrapNetsuiteAccountsV1Arc**](BatchAPI.md#BatchBootstrapNetsuiteAccountsV1Arc) | **Post** /v1/arc/batch/bootstrap-netsuite-accounts | Batch Bootstrap Netsuite Accounts
[**BatchBulkActionV1**](BatchAPI.md#BatchBulkActionV1) | **Post** /v1/arc/batch/bulk-action | Batch Bulk Action
[**BatchDeleteNetsuiteManagedPortalMappingV1ArcBatch**](BatchAPI.md#BatchDeleteNetsuiteManagedPortalMappingV1ArcBatch) | **Delete** /v1/arc/batch/netsuite-managed-portal-mapping | Batch Delete Netsuite Managed Portal Mapping
[**BatchDeleteNetsuiteManagedPortalMappingV1ArcBatchPost**](BatchAPI.md#BatchDeleteNetsuiteManagedPortalMappingV1ArcBatchPost) | **Post** /v1/arc/batch/netsuite-managed-portal-mapping/delete | Batch Delete Netsuite Managed Portal Mapping
[**BatchLinkNetsuiteAccountV1Arc**](BatchAPI.md#BatchLinkNetsuiteAccountV1Arc) | **Post** /v1/arc/batch/link-netsuite-account | Batch Link Netsuite Account
[**BatchRecalculateAgingV1**](BatchAPI.md#BatchRecalculateAgingV1) | **Post** /v1/arc/batch/recalculate-aging | Batch Recalculate Aging
[**BatchReconcilePendingProviderConfirmationsV1ArcBatch**](BatchAPI.md#BatchReconcilePendingProviderConfirmationsV1ArcBatch) | **Post** /v1/arc/batch/communications/reconcile-pending-provider-confirmations | Batch Reconcile Pending Provider Confirmations
[**BatchSetNetsuiteCutoverGuardV1Arc**](BatchAPI.md#BatchSetNetsuiteCutoverGuardV1Arc) | **Post** /v1/arc/batch/netsuite-cutover-guard | Batch Set Netsuite Cutover Guard
[**BatchSetNetsuiteManagedPortalMappingV1ArcBatch**](BatchAPI.md#BatchSetNetsuiteManagedPortalMappingV1ArcBatch) | **Post** /v1/arc/batch/netsuite-managed-portal-mapping | Batch Set Netsuite Managed Portal Mapping
[**BatchSyncInvoicesV1**](BatchAPI.md#BatchSyncInvoicesV1) | **Post** /v1/arc/batch/sync-invoices | Batch Sync Invoices
[**BatchTriggerNetsuiteSyncV1Arc**](BatchAPI.md#BatchTriggerNetsuiteSyncV1Arc) | **Post** /v1/arc/batch/trigger-netsuite-sync | Batch Trigger Netsuite Sync



## BatchBackfillManagedPortalSuppressionAckProtectionV1ArcBatch

> interface{} BatchBackfillManagedPortalSuppressionAckProtectionV1ArcBatch(ctx).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).XUserId(xUserId).XArcRoles(xArcRoles).XArcProxySecret(xArcProxySecret).NetSuiteManagedPortalProtectionAckRequest(netSuiteManagedPortalProtectionAckRequest).Execute()

Batch Backfill Managed Portal Suppression Ack Protection



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
	xUserId := "xUserId_example" // string |  (optional)
	xArcRoles := "xArcRoles_example" // string |  (optional)
	xArcProxySecret := "xArcProxySecret_example" // string |  (optional)
	netSuiteManagedPortalProtectionAckRequest := *openapiclient.NewNetSuiteManagedPortalProtectionAckRequest() // NetSuiteManagedPortalProtectionAckRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.BatchBackfillManagedPortalSuppressionAckProtectionV1ArcBatch(context.Background()).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).XUserId(xUserId).XArcRoles(xArcRoles).XArcProxySecret(xArcProxySecret).NetSuiteManagedPortalProtectionAckRequest(netSuiteManagedPortalProtectionAckRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.BatchBackfillManagedPortalSuppressionAckProtectionV1ArcBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchBackfillManagedPortalSuppressionAckProtectionV1ArcBatch`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.BatchBackfillManagedPortalSuppressionAckProtectionV1ArcBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchBackfillManagedPortalSuppressionAckProtectionV1ArcBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **authorization** | **string** |  | 
 **xUserId** | **string** |  | 
 **xArcRoles** | **string** |  | 
 **xArcProxySecret** | **string** |  | 
 **netSuiteManagedPortalProtectionAckRequest** | [**NetSuiteManagedPortalProtectionAckRequest**](NetSuiteManagedPortalProtectionAckRequest.md) |  | 

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


## BatchBackfillManagedPortalSuppressionCancelV1ArcBatch

> interface{} BatchBackfillManagedPortalSuppressionCancelV1ArcBatch(ctx, jobId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).XUserId(xUserId).XArcRoles(xArcRoles).XArcProxySecret(xArcProxySecret).NetSuiteManagedPortalBackfillCancelRequest(netSuiteManagedPortalBackfillCancelRequest).Execute()

Batch Backfill Managed Portal Suppression Cancel



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
	jobId := "jobId_example" // string | 
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)
	xUserId := "xUserId_example" // string |  (optional)
	xArcRoles := "xArcRoles_example" // string |  (optional)
	xArcProxySecret := "xArcProxySecret_example" // string |  (optional)
	netSuiteManagedPortalBackfillCancelRequest := *openapiclient.NewNetSuiteManagedPortalBackfillCancelRequest() // NetSuiteManagedPortalBackfillCancelRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.BatchBackfillManagedPortalSuppressionCancelV1ArcBatch(context.Background(), jobId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).XUserId(xUserId).XArcRoles(xArcRoles).XArcProxySecret(xArcProxySecret).NetSuiteManagedPortalBackfillCancelRequest(netSuiteManagedPortalBackfillCancelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.BatchBackfillManagedPortalSuppressionCancelV1ArcBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchBackfillManagedPortalSuppressionCancelV1ArcBatch`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.BatchBackfillManagedPortalSuppressionCancelV1ArcBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchBackfillManagedPortalSuppressionCancelV1ArcBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **authorization** | **string** |  | 
 **xUserId** | **string** |  | 
 **xArcRoles** | **string** |  | 
 **xArcProxySecret** | **string** |  | 
 **netSuiteManagedPortalBackfillCancelRequest** | [**NetSuiteManagedPortalBackfillCancelRequest**](NetSuiteManagedPortalBackfillCancelRequest.md) |  | 

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


## BatchBackfillManagedPortalSuppressionRunV1ArcBatch

> interface{} BatchBackfillManagedPortalSuppressionRunV1ArcBatch(ctx).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).XUserId(xUserId).XArcRoles(xArcRoles).XArcProxySecret(xArcProxySecret).NetSuiteManagedPortalBackfillRunRequest(netSuiteManagedPortalBackfillRunRequest).Execute()

Batch Backfill Managed Portal Suppression Run



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
	xUserId := "xUserId_example" // string |  (optional)
	xArcRoles := "xArcRoles_example" // string |  (optional)
	xArcProxySecret := "xArcProxySecret_example" // string |  (optional)
	netSuiteManagedPortalBackfillRunRequest := *openapiclient.NewNetSuiteManagedPortalBackfillRunRequest() // NetSuiteManagedPortalBackfillRunRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.BatchBackfillManagedPortalSuppressionRunV1ArcBatch(context.Background()).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).XUserId(xUserId).XArcRoles(xArcRoles).XArcProxySecret(xArcProxySecret).NetSuiteManagedPortalBackfillRunRequest(netSuiteManagedPortalBackfillRunRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.BatchBackfillManagedPortalSuppressionRunV1ArcBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchBackfillManagedPortalSuppressionRunV1ArcBatch`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.BatchBackfillManagedPortalSuppressionRunV1ArcBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchBackfillManagedPortalSuppressionRunV1ArcBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **authorization** | **string** |  | 
 **xUserId** | **string** |  | 
 **xArcRoles** | **string** |  | 
 **xArcProxySecret** | **string** |  | 
 **netSuiteManagedPortalBackfillRunRequest** | [**NetSuiteManagedPortalBackfillRunRequest**](NetSuiteManagedPortalBackfillRunRequest.md) |  | 

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


## BatchBackfillManagedPortalSuppressionRunV1ArcBatchBackfill

> interface{} BatchBackfillManagedPortalSuppressionRunV1ArcBatchBackfill(ctx).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).XUserId(xUserId).XArcRoles(xArcRoles).XArcProxySecret(xArcProxySecret).NetSuiteManagedPortalBackfillRunRequest(netSuiteManagedPortalBackfillRunRequest).Execute()

Batch Backfill Managed Portal Suppression Run



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
	xUserId := "xUserId_example" // string |  (optional)
	xArcRoles := "xArcRoles_example" // string |  (optional)
	xArcProxySecret := "xArcProxySecret_example" // string |  (optional)
	netSuiteManagedPortalBackfillRunRequest := *openapiclient.NewNetSuiteManagedPortalBackfillRunRequest() // NetSuiteManagedPortalBackfillRunRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.BatchBackfillManagedPortalSuppressionRunV1ArcBatchBackfill(context.Background()).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).XUserId(xUserId).XArcRoles(xArcRoles).XArcProxySecret(xArcProxySecret).NetSuiteManagedPortalBackfillRunRequest(netSuiteManagedPortalBackfillRunRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.BatchBackfillManagedPortalSuppressionRunV1ArcBatchBackfill``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchBackfillManagedPortalSuppressionRunV1ArcBatchBackfill`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.BatchBackfillManagedPortalSuppressionRunV1ArcBatchBackfill`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchBackfillManagedPortalSuppressionRunV1ArcBatchBackfillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **authorization** | **string** |  | 
 **xUserId** | **string** |  | 
 **xArcRoles** | **string** |  | 
 **xArcProxySecret** | **string** |  | 
 **netSuiteManagedPortalBackfillRunRequest** | [**NetSuiteManagedPortalBackfillRunRequest**](NetSuiteManagedPortalBackfillRunRequest.md) |  | 

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


## BatchBackfillManagedPortalSuppressionStatusV1ArcBatch

> interface{} BatchBackfillManagedPortalSuppressionStatusV1ArcBatch(ctx, jobId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).XArcRoles(xArcRoles).XArcProxySecret(xArcProxySecret).Execute()

Batch Backfill Managed Portal Suppression Status



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
	jobId := "jobId_example" // string | 
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)
	xArcRoles := "xArcRoles_example" // string |  (optional)
	xArcProxySecret := "xArcProxySecret_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.BatchBackfillManagedPortalSuppressionStatusV1ArcBatch(context.Background(), jobId).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).XArcRoles(xArcRoles).XArcProxySecret(xArcProxySecret).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.BatchBackfillManagedPortalSuppressionStatusV1ArcBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchBackfillManagedPortalSuppressionStatusV1ArcBatch`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.BatchBackfillManagedPortalSuppressionStatusV1ArcBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchBackfillManagedPortalSuppressionStatusV1ArcBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **authorization** | **string** |  | 
 **xArcRoles** | **string** |  | 
 **xArcProxySecret** | **string** |  | 

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


## BatchBackfillManagedPortalSuppressionV1ArcBatch

> interface{} BatchBackfillManagedPortalSuppressionV1ArcBatch(ctx).NetSuiteManagedPortalBackfillRequest(netSuiteManagedPortalBackfillRequest).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).XUserId(xUserId).XArcRoles(xArcRoles).XArcProxySecret(xArcProxySecret).Execute()

Batch Backfill Managed Portal Suppression



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
	netSuiteManagedPortalBackfillRequest := *openapiclient.NewNetSuiteManagedPortalBackfillRequest() // NetSuiteManagedPortalBackfillRequest | 
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)
	xUserId := "xUserId_example" // string |  (optional)
	xArcRoles := "xArcRoles_example" // string |  (optional)
	xArcProxySecret := "xArcProxySecret_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.BatchBackfillManagedPortalSuppressionV1ArcBatch(context.Background()).NetSuiteManagedPortalBackfillRequest(netSuiteManagedPortalBackfillRequest).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).XUserId(xUserId).XArcRoles(xArcRoles).XArcProxySecret(xArcProxySecret).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.BatchBackfillManagedPortalSuppressionV1ArcBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchBackfillManagedPortalSuppressionV1ArcBatch`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.BatchBackfillManagedPortalSuppressionV1ArcBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchBackfillManagedPortalSuppressionV1ArcBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **netSuiteManagedPortalBackfillRequest** | [**NetSuiteManagedPortalBackfillRequest**](NetSuiteManagedPortalBackfillRequest.md) |  | 
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **authorization** | **string** |  | 
 **xUserId** | **string** |  | 
 **xArcRoles** | **string** |  | 
 **xArcProxySecret** | **string** |  | 

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


## BatchBackfillNetsuiteDeferredInvoicesV1ArcBatch

> interface{} BatchBackfillNetsuiteDeferredInvoicesV1ArcBatch(ctx).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()

Batch Backfill Netsuite Deferred Invoices



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
	resp, r, err := apiClient.BatchAPI.BatchBackfillNetsuiteDeferredInvoicesV1ArcBatch(context.Background()).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.BatchBackfillNetsuiteDeferredInvoicesV1ArcBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchBackfillNetsuiteDeferredInvoicesV1ArcBatch`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.BatchBackfillNetsuiteDeferredInvoicesV1ArcBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchBackfillNetsuiteDeferredInvoicesV1ArcBatchRequest struct via the builder pattern


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


## BatchBackfillNetsuiteIdentitiesV1Arc

> interface{} BatchBackfillNetsuiteIdentitiesV1Arc(ctx).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()

Batch Backfill Netsuite Identities



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
	resp, r, err := apiClient.BatchAPI.BatchBackfillNetsuiteIdentitiesV1Arc(context.Background()).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.BatchBackfillNetsuiteIdentitiesV1Arc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchBackfillNetsuiteIdentitiesV1Arc`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.BatchBackfillNetsuiteIdentitiesV1Arc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchBackfillNetsuiteIdentitiesV1ArcRequest struct via the builder pattern


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


## BatchBootstrapNetsuiteAccountsV1Arc

> interface{} BatchBootstrapNetsuiteAccountsV1Arc(ctx).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()

Batch Bootstrap Netsuite Accounts



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
	resp, r, err := apiClient.BatchAPI.BatchBootstrapNetsuiteAccountsV1Arc(context.Background()).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.BatchBootstrapNetsuiteAccountsV1Arc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchBootstrapNetsuiteAccountsV1Arc`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.BatchBootstrapNetsuiteAccountsV1Arc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchBootstrapNetsuiteAccountsV1ArcRequest struct via the builder pattern


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


## BatchBulkActionV1

> interface{} BatchBulkActionV1(ctx).BulkActionRequest(bulkActionRequest).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()

Batch Bulk Action



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
	bulkActionRequest := *openapiclient.NewBulkActionRequest("Action_example", []string{"CaseIds_example"}) // BulkActionRequest | 
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.BatchBulkActionV1(context.Background()).BulkActionRequest(bulkActionRequest).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.BatchBulkActionV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchBulkActionV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.BatchBulkActionV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchBulkActionV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkActionRequest** | [**BulkActionRequest**](BulkActionRequest.md) |  | 
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


## BatchDeleteNetsuiteManagedPortalMappingV1ArcBatch

> interface{} BatchDeleteNetsuiteManagedPortalMappingV1ArcBatch(ctx).NetSuiteManagedPortalMappingDeleteRequest(netSuiteManagedPortalMappingDeleteRequest).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).XUserId(xUserId).XArcRoles(xArcRoles).XArcProxySecret(xArcProxySecret).Execute()

Batch Delete Netsuite Managed Portal Mapping



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
	netSuiteManagedPortalMappingDeleteRequest := *openapiclient.NewNetSuiteManagedPortalMappingDeleteRequest("ChangeReason_example") // NetSuiteManagedPortalMappingDeleteRequest | 
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)
	xUserId := "xUserId_example" // string |  (optional)
	xArcRoles := "xArcRoles_example" // string |  (optional)
	xArcProxySecret := "xArcProxySecret_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.BatchDeleteNetsuiteManagedPortalMappingV1ArcBatch(context.Background()).NetSuiteManagedPortalMappingDeleteRequest(netSuiteManagedPortalMappingDeleteRequest).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).XUserId(xUserId).XArcRoles(xArcRoles).XArcProxySecret(xArcProxySecret).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.BatchDeleteNetsuiteManagedPortalMappingV1ArcBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchDeleteNetsuiteManagedPortalMappingV1ArcBatch`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.BatchDeleteNetsuiteManagedPortalMappingV1ArcBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchDeleteNetsuiteManagedPortalMappingV1ArcBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **netSuiteManagedPortalMappingDeleteRequest** | [**NetSuiteManagedPortalMappingDeleteRequest**](NetSuiteManagedPortalMappingDeleteRequest.md) |  | 
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **authorization** | **string** |  | 
 **xUserId** | **string** |  | 
 **xArcRoles** | **string** |  | 
 **xArcProxySecret** | **string** |  | 

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


## BatchDeleteNetsuiteManagedPortalMappingV1ArcBatchPost

> interface{} BatchDeleteNetsuiteManagedPortalMappingV1ArcBatchPost(ctx).NetSuiteManagedPortalMappingDeleteRequest(netSuiteManagedPortalMappingDeleteRequest).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).XUserId(xUserId).XArcRoles(xArcRoles).XArcProxySecret(xArcProxySecret).Execute()

Batch Delete Netsuite Managed Portal Mapping



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
	netSuiteManagedPortalMappingDeleteRequest := *openapiclient.NewNetSuiteManagedPortalMappingDeleteRequest("ChangeReason_example") // NetSuiteManagedPortalMappingDeleteRequest | 
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)
	xUserId := "xUserId_example" // string |  (optional)
	xArcRoles := "xArcRoles_example" // string |  (optional)
	xArcProxySecret := "xArcProxySecret_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.BatchDeleteNetsuiteManagedPortalMappingV1ArcBatchPost(context.Background()).NetSuiteManagedPortalMappingDeleteRequest(netSuiteManagedPortalMappingDeleteRequest).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).XUserId(xUserId).XArcRoles(xArcRoles).XArcProxySecret(xArcProxySecret).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.BatchDeleteNetsuiteManagedPortalMappingV1ArcBatchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchDeleteNetsuiteManagedPortalMappingV1ArcBatchPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.BatchDeleteNetsuiteManagedPortalMappingV1ArcBatchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchDeleteNetsuiteManagedPortalMappingV1ArcBatchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **netSuiteManagedPortalMappingDeleteRequest** | [**NetSuiteManagedPortalMappingDeleteRequest**](NetSuiteManagedPortalMappingDeleteRequest.md) |  | 
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **authorization** | **string** |  | 
 **xUserId** | **string** |  | 
 **xArcRoles** | **string** |  | 
 **xArcProxySecret** | **string** |  | 

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


## BatchLinkNetsuiteAccountV1Arc

> interface{} BatchLinkNetsuiteAccountV1Arc(ctx).LinkNetSuiteAccountRequest(linkNetSuiteAccountRequest).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()

Batch Link Netsuite Account



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
	linkNetSuiteAccountRequest := *openapiclient.NewLinkNetSuiteAccountRequest("AccountId_example", "NetsuiteCustomerId_example") // LinkNetSuiteAccountRequest | 
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.BatchLinkNetsuiteAccountV1Arc(context.Background()).LinkNetSuiteAccountRequest(linkNetSuiteAccountRequest).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.BatchLinkNetsuiteAccountV1Arc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchLinkNetsuiteAccountV1Arc`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.BatchLinkNetsuiteAccountV1Arc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchLinkNetsuiteAccountV1ArcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkNetSuiteAccountRequest** | [**LinkNetSuiteAccountRequest**](LinkNetSuiteAccountRequest.md) |  | 
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


## BatchRecalculateAgingV1

> interface{} BatchRecalculateAgingV1(ctx).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()

Batch Recalculate Aging



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
	resp, r, err := apiClient.BatchAPI.BatchRecalculateAgingV1(context.Background()).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.BatchRecalculateAgingV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchRecalculateAgingV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.BatchRecalculateAgingV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchRecalculateAgingV1Request struct via the builder pattern


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


## BatchReconcilePendingProviderConfirmationsV1ArcBatch

> interface{} BatchReconcilePendingProviderConfirmationsV1ArcBatch(ctx).PendingProviderConfirmationReconcileRequest(pendingProviderConfirmationReconcileRequest).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).XArcRoles(xArcRoles).XArcProxySecret(xArcProxySecret).Execute()

Batch Reconcile Pending Provider Confirmations



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
	pendingProviderConfirmationReconcileRequest := *openapiclient.NewPendingProviderConfirmationReconcileRequest() // PendingProviderConfirmationReconcileRequest | 
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)
	xArcRoles := "xArcRoles_example" // string |  (optional)
	xArcProxySecret := "xArcProxySecret_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.BatchReconcilePendingProviderConfirmationsV1ArcBatch(context.Background()).PendingProviderConfirmationReconcileRequest(pendingProviderConfirmationReconcileRequest).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).XArcRoles(xArcRoles).XArcProxySecret(xArcProxySecret).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.BatchReconcilePendingProviderConfirmationsV1ArcBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchReconcilePendingProviderConfirmationsV1ArcBatch`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.BatchReconcilePendingProviderConfirmationsV1ArcBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchReconcilePendingProviderConfirmationsV1ArcBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pendingProviderConfirmationReconcileRequest** | [**PendingProviderConfirmationReconcileRequest**](PendingProviderConfirmationReconcileRequest.md) |  | 
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **authorization** | **string** |  | 
 **xArcRoles** | **string** |  | 
 **xArcProxySecret** | **string** |  | 

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


## BatchSetNetsuiteCutoverGuardV1Arc

> interface{} BatchSetNetsuiteCutoverGuardV1Arc(ctx).NetSuiteCutoverGuardRequest(netSuiteCutoverGuardRequest).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()

Batch Set Netsuite Cutover Guard



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
	netSuiteCutoverGuardRequest := *openapiclient.NewNetSuiteCutoverGuardRequest(false) // NetSuiteCutoverGuardRequest | 
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.BatchSetNetsuiteCutoverGuardV1Arc(context.Background()).NetSuiteCutoverGuardRequest(netSuiteCutoverGuardRequest).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.BatchSetNetsuiteCutoverGuardV1Arc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchSetNetsuiteCutoverGuardV1Arc`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.BatchSetNetsuiteCutoverGuardV1Arc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchSetNetsuiteCutoverGuardV1ArcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **netSuiteCutoverGuardRequest** | [**NetSuiteCutoverGuardRequest**](NetSuiteCutoverGuardRequest.md) |  | 
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


## BatchSetNetsuiteManagedPortalMappingV1ArcBatch

> interface{} BatchSetNetsuiteManagedPortalMappingV1ArcBatch(ctx).NetSuiteManagedPortalMappingRequest(netSuiteManagedPortalMappingRequest).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).XUserId(xUserId).XArcRoles(xArcRoles).XArcProxySecret(xArcProxySecret).Execute()

Batch Set Netsuite Managed Portal Mapping



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
	netSuiteManagedPortalMappingRequest := *openapiclient.NewNetSuiteManagedPortalMappingRequest("SourceFieldId_example", "ChangeReason_example") // NetSuiteManagedPortalMappingRequest | 
	xAPIKey := "xAPIKey_example" // string |  (optional)
	xTenantId := "xTenantId_example" // string |  (optional)
	xOrgId := "xOrgId_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)
	xUserId := "xUserId_example" // string |  (optional)
	xArcRoles := "xArcRoles_example" // string |  (optional)
	xArcProxySecret := "xArcProxySecret_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.BatchSetNetsuiteManagedPortalMappingV1ArcBatch(context.Background()).NetSuiteManagedPortalMappingRequest(netSuiteManagedPortalMappingRequest).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).XUserId(xUserId).XArcRoles(xArcRoles).XArcProxySecret(xArcProxySecret).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.BatchSetNetsuiteManagedPortalMappingV1ArcBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchSetNetsuiteManagedPortalMappingV1ArcBatch`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.BatchSetNetsuiteManagedPortalMappingV1ArcBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchSetNetsuiteManagedPortalMappingV1ArcBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **netSuiteManagedPortalMappingRequest** | [**NetSuiteManagedPortalMappingRequest**](NetSuiteManagedPortalMappingRequest.md) |  | 
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **authorization** | **string** |  | 
 **xUserId** | **string** |  | 
 **xArcRoles** | **string** |  | 
 **xArcProxySecret** | **string** |  | 

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


## BatchSyncInvoicesV1

> interface{} BatchSyncInvoicesV1(ctx).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).SyncInvoicesRequest(syncInvoicesRequest).Execute()

Batch Sync Invoices



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
	syncInvoicesRequest := *openapiclient.NewSyncInvoicesRequest() // SyncInvoicesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.BatchSyncInvoicesV1(context.Background()).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).SyncInvoicesRequest(syncInvoicesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.BatchSyncInvoicesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchSyncInvoicesV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.BatchSyncInvoicesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchSyncInvoicesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKey** | **string** |  | 
 **xTenantId** | **string** |  | 
 **xOrgId** | **string** |  | 
 **authorization** | **string** |  | 
 **syncInvoicesRequest** | [**SyncInvoicesRequest**](SyncInvoicesRequest.md) |  | 

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


## BatchTriggerNetsuiteSyncV1Arc

> interface{} BatchTriggerNetsuiteSyncV1Arc(ctx).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()

Batch Trigger Netsuite Sync



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
	resp, r, err := apiClient.BatchAPI.BatchTriggerNetsuiteSyncV1Arc(context.Background()).XAPIKey(xAPIKey).XTenantId(xTenantId).XOrgId(xOrgId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.BatchTriggerNetsuiteSyncV1Arc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchTriggerNetsuiteSyncV1Arc`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.BatchTriggerNetsuiteSyncV1Arc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchTriggerNetsuiteSyncV1ArcRequest struct via the builder pattern


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

