# \V1API

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllocateCredits**](V1API.md#AllocateCredits) | **Post** /v1/wallets/allocate | Allocate Credits
[**BindSubjectToWallet**](V1API.md#BindSubjectToWallet) | **Post** /v1/wallets/members | Bind Subject To Wallet
[**CheckTriggerV1**](V1API.md#CheckTriggerV1) | **Post** /v1/auto-topup/rules/{rule_id}/check | Check Trigger
[**CreateFxRateEndpointV1**](V1API.md#CreateFxRateEndpointV1) | **Post** /v1/fx-rates/rates | Create Fx Rate Endpoint
[**CreateMapping**](V1API.md#CreateMapping) | **Post** /v1/mappings | Create Mapping
[**CreateRateCard**](V1API.md#CreateRateCard) | **Post** /v1/rate_cards | Create Rate Card
[**CreateReconstructionRun**](V1API.md#CreateReconstructionRun) | **Post** /v1/reconstruction/runs | Create Reconstruction Run
[**CreateRuleV1**](V1API.md#CreateRuleV1) | **Post** /v1/auto-topup/rules | Create Rule
[**CreateSnapshot**](V1API.md#CreateSnapshot) | **Post** /v1/snapshots | Create Snapshot
[**CreateThresholdEndpoint**](V1API.md#CreateThresholdEndpoint) | **Post** /v1/alerts/thresholds | Create Threshold Endpoint
[**CreateTokenEndpoint**](V1API.md#CreateTokenEndpoint) | **Post** /v1/portal/tokens | Create Token Endpoint
[**CreateValueRecognitionEntryEndpointV1Fx**](V1API.md#CreateValueRecognitionEntryEndpointV1Fx) | **Post** /v1/fx-rates/value-recognition/entry | Create Value Recognition Entry Endpoint
[**CreateWallet**](V1API.md#CreateWallet) | **Post** /v1/wallets | Create Wallet
[**DeleteRuleV1**](V1API.md#DeleteRuleV1) | **Delete** /v1/auto-topup/rules/{rule_id} | Delete Rule
[**DeleteThresholdEndpoint**](V1API.md#DeleteThresholdEndpoint) | **Delete** /v1/alerts/thresholds/{threshold_id} | Delete Threshold Endpoint
[**DiagnoseSubscription**](V1API.md#DiagnoseSubscription) | **Get** /v1/admin/diagnostics/subscription/{subscription_id} | Diagnose Subscription
[**ExplainQuery**](V1API.md#ExplainQuery) | **Post** /v1/performance/queries/explain | Explain Query
[**GetActivityV1**](V1API.md#GetActivityV1) | **Get** /v1/auto-topup/activity | Get Activity
[**GetAffectedEventsEndpointV1**](V1API.md#GetAffectedEventsEndpointV1) | **Get** /v1/reconstruction/affected-events | Get Affected Events Endpoint
[**GetAlertsStateEndpoint**](V1API.md#GetAlertsStateEndpoint) | **Get** /v1/alerts/state | Get Alerts State Endpoint
[**GetConnectionPoolInfoV1**](V1API.md#GetConnectionPoolInfoV1) | **Get** /v1/performance/connection-pool | Get Connection Pool Info
[**GetFxRateAtEndpointV1**](V1API.md#GetFxRateAtEndpointV1) | **Post** /v1/fx-rates/rates/at | Get Fx Rate At Endpoint
[**GetGrant**](V1API.md#GetGrant) | **Get** /v1/grants/{grant_id} | Get Grant
[**GetIndexUsage**](V1API.md#GetIndexUsage) | **Get** /v1/performance/indexes/usage | Get Index Usage
[**GetLateEventsV1**](V1API.md#GetLateEventsV1) | **Get** /v1/reconstruction/late-events | Get Late Events
[**GetLedgerAudit**](V1API.md#GetLedgerAudit) | **Get** /v1/ledger/audit | Get Ledger Audit
[**GetLedgerAuditGet**](V1API.md#GetLedgerAuditGet) | **Get** /v1/admin/ledger/audit | Get Ledger Audit
[**GetLedgerBalance**](V1API.md#GetLedgerBalance) | **Get** /v1/ledger/balance | Get Ledger Balance
[**GetLedgerBalanceGet**](V1API.md#GetLedgerBalanceGet) | **Get** /v1/admin/ledger/balance | Get Ledger Balance
[**GetMissingIndexRecommendations**](V1API.md#GetMissingIndexRecommendations) | **Get** /v1/performance/indexes/missing | Get Missing Index Recommendations
[**GetPerformanceReportEndpoint**](V1API.md#GetPerformanceReportEndpoint) | **Get** /v1/performance/report | Get Performance Report Endpoint
[**GetPortalContext**](V1API.md#GetPortalContext) | **Get** /v1/portal/context | Get Portal Context
[**GetRateCard**](V1API.md#GetRateCard) | **Get** /v1/rate_cards/{rate_card_id} | Get Rate Card
[**GetRuleV1**](V1API.md#GetRuleV1) | **Get** /v1/auto-topup/rules/{rule_id} | Get Rule
[**GetSlowQueries**](V1API.md#GetSlowQueries) | **Get** /v1/performance/queries/slow | Get Slow Queries
[**GetSnapshotAtEndpoint**](V1API.md#GetSnapshotAtEndpoint) | **Get** /v1/snapshots/wallet/{wallet_id}/at | Get Snapshot At Endpoint
[**GetSubscription**](V1API.md#GetSubscription) | **Get** /v1/subscriptions/{subscription_id} | Get Subscription
[**GetTableSizeAnalysis**](V1API.md#GetTableSizeAnalysis) | **Get** /v1/performance/tables/sizes | Get Table Size Analysis
[**GetUsageEvent**](V1API.md#GetUsageEvent) | **Get** /v1/usage/{event_id} | Get Usage Event
[**GetWallet**](V1API.md#GetWallet) | **Get** /v1/wallets/{wallet_id} | Get Wallet
[**GetWalletActivity**](V1API.md#GetWalletActivity) | **Get** /v1/wallets/{wallet_id}/activity | Get Wallet Activity
[**GetWalletMembers**](V1API.md#GetWalletMembers) | **Get** /v1/wallets/{wallet_id}/members | Get Wallet Members
[**GetWalletState**](V1API.md#GetWalletState) | **Get** /v1/ledger/state | Get Wallet State
[**GetWalletStateProjectionEndpointV1**](V1API.md#GetWalletStateProjectionEndpointV1) | **Get** /v1/state-projection/wallet/{wallet_id} | Get Wallet State Projection Endpoint
[**GetWalletSubtree**](V1API.md#GetWalletSubtree) | **Get** /v1/reconstruction/wallets/{wallet_id}/subtree | Get Wallet Subtree
[**HandleLifecycleEvent**](V1API.md#HandleLifecycleEvent) | **Post** /v1/subscriptions/lifecycle | Handle Lifecycle Event
[**InvalidateTokensEndpoint**](V1API.md#InvalidateTokensEndpoint) | **Post** /v1/portal/tokens/invalidate | Invalidate Tokens Endpoint
[**ListGrants**](V1API.md#ListGrants) | **Get** /v1/grants | List Grants
[**ListMappings**](V1API.md#ListMappings) | **Get** /v1/mappings | List Mappings
[**ListOutboxEventsEndpoint**](V1API.md#ListOutboxEventsEndpoint) | **Get** /v1/outbox/events | List Outbox Events Endpoint
[**ListRulesV1**](V1API.md#ListRulesV1) | **Get** /v1/auto-topup/rules | List Rules
[**ListSnapshotsEndpoint**](V1API.md#ListSnapshotsEndpoint) | **Get** /v1/snapshots/wallet/{wallet_id} | List Snapshots Endpoint
[**ListThresholdsEndpoint**](V1API.md#ListThresholdsEndpoint) | **Get** /v1/alerts/thresholds | List Thresholds Endpoint
[**ListTokensEndpoint**](V1API.md#ListTokensEndpoint) | **Get** /v1/portal/tokens | List Tokens Endpoint
[**ListUnpricedEvents**](V1API.md#ListUnpricedEvents) | **Get** /v1/unpriced/events | List Unpriced Events
[**ListUsageEvents**](V1API.md#ListUsageEvents) | **Get** /v1/usage/ | List Usage Events
[**ListWalletMembers**](V1API.md#ListWalletMembers) | **Get** /v1/wallet_members | List Wallet Members
[**ListWallets**](V1API.md#ListWallets) | **Get** /v1/wallets | List Wallets
[**MintGrant**](V1API.md#MintGrant) | **Post** /v1/grants/mint | Mint Grant
[**OpenmeterWebhook**](V1API.md#OpenmeterWebhook) | **Post** /v1/integrations/openmeter/webhooks/openmeter | Openmeter Webhook
[**OpenmeterWebhookBatch**](V1API.md#OpenmeterWebhookBatch) | **Post** /v1/integrations/openmeter/webhooks/openmeter/batch | Openmeter Webhook Batch
[**PatchRuleV1**](V1API.md#PatchRuleV1) | **Patch** /v1/auto-topup/rules/{rule_id} | Patch Rule
[**PaymentSucceededV1**](V1API.md#PaymentSucceededV1) | **Post** /v1/auto-topup/payments/succeeded | Payment Succeeded
[**ProcessEvents**](V1API.md#ProcessEvents) | **Post** /v1/unpriced/process | Process Events
[**ProcessLifecycleEventManuallyV1Admin**](V1API.md#ProcessLifecycleEventManuallyV1Admin) | **Post** /v1/admin/subscriptions/{subscription_id}/process-lifecycle-event | Process Lifecycle Event Manually
[**ProcessOutboxEndpoint**](V1API.md#ProcessOutboxEndpoint) | **Post** /v1/outbox/process | Process Outbox Endpoint
[**ProcessPendingEventsV1**](V1API.md#ProcessPendingEventsV1) | **Post** /v1/rating/process-pending | Process Pending Events
[**ProcessPendingProjectionsEndpointV1**](V1API.md#ProcessPendingProjectionsEndpointV1) | **Post** /v1/state-projection/process | Process Pending Projections Endpoint
[**ProcessRollover**](V1API.md#ProcessRollover) | **Post** /v1/rollover/process | Process Rollover
[**ProcessValueRecognitionV1Fx**](V1API.md#ProcessValueRecognitionV1Fx) | **Post** /v1/fx-rates/value-recognition/process | Process Value Recognition
[**ProjectWalletStateEndpointV1**](V1API.md#ProjectWalletStateEndpointV1) | **Post** /v1/state-projection/project/{wallet_id} | Project Wallet State Endpoint
[**RateEvent**](V1API.md#RateEvent) | **Post** /v1/rating/rate/{usage_event_id} | Rate Event
[**RecordUsage**](V1API.md#RecordUsage) | **Post** /v1/usage/record | Record Usage
[**ResolveEvent**](V1API.md#ResolveEvent) | **Post** /v1/unpriced/resolve/{usage_event_id} | Resolve Event
[**ResolveTenantAndPool**](V1API.md#ResolveTenantAndPool) | **Get** /v1/wallets/resolve | Resolve Tenant And Pool
[**RetryUnpriced**](V1API.md#RetryUnpriced) | **Post** /v1/admin/usage/unpriced/retry | Retry Unpriced
[**ReviewQuarantined**](V1API.md#ReviewQuarantined) | **Post** /v1/admin/usage/quarantined/review | Review Quarantined
[**SyncSubscription**](V1API.md#SyncSubscription) | **Post** /v1/subscriptions/sync | Sync Subscription
[**TriggerSubscriptionSync**](V1API.md#TriggerSubscriptionSync) | **Post** /v1/admin/subscriptions/{subscription_id}/sync | Trigger Subscription Sync
[**UpdateRuleV1**](V1API.md#UpdateRuleV1) | **Put** /v1/auto-topup/rules/{rule_id} | Update Rule
[**UpdateThresholdEndpoint**](V1API.md#UpdateThresholdEndpoint) | **Put** /v1/alerts/thresholds/{threshold_id} | Update Threshold Endpoint
[**UpdateWalletSettings**](V1API.md#UpdateWalletSettings) | **Patch** /v1/wallets/{wallet_id}/settings | Update Wallet Settings
[**UpdateWalletThresholds**](V1API.md#UpdateWalletThresholds) | **Patch** /v1/wallets/{wallet_id}/thresholds | Update Wallet Thresholds
[**ValidateSnapshotEndpoint**](V1API.md#ValidateSnapshotEndpoint) | **Post** /v1/snapshots/{snapshot_id}/validate | Validate Snapshot Endpoint
[**VoidGrant**](V1API.md#VoidGrant) | **Post** /v1/grants/{grant_id}/void | Void Grant
[**WarmRateCardCacheV1**](V1API.md#WarmRateCardCacheV1) | **Post** /v1/performance/cache/warm/rate-cards | Warm Rate Card Cache
[**WarmWalletCache**](V1API.md#WarmWalletCache) | **Post** /v1/performance/cache/warm/wallets | Warm Wallet Cache



## AllocateCredits

> interface{} AllocateCredits(ctx).AllocateCreditsRequest(allocateCreditsRequest).Execute()

Allocate Credits



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
	allocateCreditsRequest := *openapiclient.NewAllocateCreditsRequest("SourceWalletId_example", "TargetWalletId_example", int32(123), "SubjectId_example") // AllocateCreditsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.AllocateCredits(context.Background()).AllocateCreditsRequest(allocateCreditsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.AllocateCredits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AllocateCredits`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.AllocateCredits`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAllocateCreditsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **allocateCreditsRequest** | [**AllocateCreditsRequest**](AllocateCreditsRequest.md) |  | 

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


## BindSubjectToWallet

> interface{} BindSubjectToWallet(ctx).TenantId(tenantId).PoolId(poolId).SubjectId(subjectId).WalletId(walletId).Execute()

Bind Subject To Wallet



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	subjectId := "subjectId_example" // string | 
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.BindSubjectToWallet(context.Background()).TenantId(tenantId).PoolId(poolId).SubjectId(subjectId).WalletId(walletId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.BindSubjectToWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BindSubjectToWallet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.BindSubjectToWallet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBindSubjectToWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **poolId** | **string** |  | 
 **subjectId** | **string** |  | 
 **walletId** | **string** |  | 

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


## CheckTriggerV1

> TriggerResponse CheckTriggerV1(ctx, ruleId).CheckTriggerRequest(checkTriggerRequest).Execute()

Check Trigger



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
	ruleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Auto top-up rule ID
	checkTriggerRequest := *openapiclient.NewCheckTriggerRequest() // CheckTriggerRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CheckTriggerV1(context.Background(), ruleId).CheckTriggerRequest(checkTriggerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CheckTriggerV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckTriggerV1`: TriggerResponse
	fmt.Fprintf(os.Stdout, "Response from `V1API.CheckTriggerV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** | Auto top-up rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckTriggerV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **checkTriggerRequest** | [**CheckTriggerRequest**](CheckTriggerRequest.md) |  | 

### Return type

[**TriggerResponse**](TriggerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFxRateEndpointV1

> FxRateResponse CreateFxRateEndpointV1(ctx).CreateFxRateRequest(createFxRateRequest).Execute()

Create Fx Rate Endpoint



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
	createFxRateRequest := *openapiclient.NewCreateFxRateRequest("TenantId_example", "PoolId_example", int32(123), time.Now()) // CreateFxRateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CreateFxRateEndpointV1(context.Background()).CreateFxRateRequest(createFxRateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CreateFxRateEndpointV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFxRateEndpointV1`: FxRateResponse
	fmt.Fprintf(os.Stdout, "Response from `V1API.CreateFxRateEndpointV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFxRateEndpointV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFxRateRequest** | [**CreateFxRateRequest**](CreateFxRateRequest.md) |  | 

### Return type

[**FxRateResponse**](FxRateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMapping

> MappingResponse CreateMapping(ctx).TenantId(tenantId).CreateMappingRequest(createMappingRequest).Execute()

Create Mapping



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier
	createMappingRequest := *openapiclient.NewCreateMappingRequest("MeterSlug_example", "FeatureKey_example") // CreateMappingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CreateMapping(context.Background()).TenantId(tenantId).CreateMappingRequest(createMappingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CreateMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMapping`: MappingResponse
	fmt.Fprintf(os.Stdout, "Response from `V1API.CreateMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Tenant identifier | 
 **createMappingRequest** | [**CreateMappingRequest**](CreateMappingRequest.md) |  | 

### Return type

[**MappingResponse**](MappingResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRateCard

> RateCardResponse CreateRateCard(ctx).CreateRateCardRequest(createRateCardRequest).Execute()

Create Rate Card



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
	createRateCardRequest := *openapiclient.NewCreateRateCardRequest("TenantId_example", "FeatureKey_example", "Version_example", time.Now(), map[string]interface{}{"key": interface{}(123)}) // CreateRateCardRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CreateRateCard(context.Background()).CreateRateCardRequest(createRateCardRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CreateRateCard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRateCard`: RateCardResponse
	fmt.Fprintf(os.Stdout, "Response from `V1API.CreateRateCard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRateCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRateCardRequest** | [**CreateRateCardRequest**](CreateRateCardRequest.md) |  | 

### Return type

[**RateCardResponse**](RateCardResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateReconstructionRun

> ReconstructionRunResponse CreateReconstructionRun(ctx).CreateReconstructionRunRequest(createReconstructionRunRequest).Execute()

Create Reconstruction Run



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
	createReconstructionRunRequest := *openapiclient.NewCreateReconstructionRunRequest("TenantId_example", "PoolId_example", "RootWalletId_example", time.Now(), time.Now(), "TriggerEventId_example") // CreateReconstructionRunRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CreateReconstructionRun(context.Background()).CreateReconstructionRunRequest(createReconstructionRunRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CreateReconstructionRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateReconstructionRun`: ReconstructionRunResponse
	fmt.Fprintf(os.Stdout, "Response from `V1API.CreateReconstructionRun`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReconstructionRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createReconstructionRunRequest** | [**CreateReconstructionRunRequest**](CreateReconstructionRunRequest.md) |  | 

### Return type

[**ReconstructionRunResponse**](ReconstructionRunResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRuleV1

> AutoTopupRuleResponse CreateRuleV1(ctx).CreateAutoTopupRuleRequest(createAutoTopupRuleRequest).Execute()

Create Rule



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
	createAutoTopupRuleRequest := *openapiclient.NewCreateAutoTopupRuleRequest("TenantId_example", "PoolId_example", "WalletId_example", "TriggerType_example", "TriggerValue_example", int32(123)) // CreateAutoTopupRuleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CreateRuleV1(context.Background()).CreateAutoTopupRuleRequest(createAutoTopupRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CreateRuleV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRuleV1`: AutoTopupRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `V1API.CreateRuleV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRuleV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAutoTopupRuleRequest** | [**CreateAutoTopupRuleRequest**](CreateAutoTopupRuleRequest.md) |  | 

### Return type

[**AutoTopupRuleResponse**](AutoTopupRuleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSnapshot

> SnapshotResponse CreateSnapshot(ctx).CreateSnapshotRequest(createSnapshotRequest).Execute()

Create Snapshot



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
	createSnapshotRequest := *openapiclient.NewCreateSnapshotRequest("TenantId_example", "PoolId_example", "WalletId_example", time.Now()) // CreateSnapshotRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CreateSnapshot(context.Background()).CreateSnapshotRequest(createSnapshotRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CreateSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSnapshot`: SnapshotResponse
	fmt.Fprintf(os.Stdout, "Response from `V1API.CreateSnapshot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSnapshotRequest** | [**CreateSnapshotRequest**](CreateSnapshotRequest.md) |  | 

### Return type

[**SnapshotResponse**](SnapshotResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateThresholdEndpoint

> interface{} CreateThresholdEndpoint(ctx).WalletId(walletId).TenantId(tenantId).PoolId(poolId).CreateThresholdRequest(createThresholdRequest).Execute()

Create Threshold Endpoint



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Wallet identifier
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Pool identifier
	createThresholdRequest := *openapiclient.NewCreateThresholdRequest("ThresholdType_example", int32(123)) // CreateThresholdRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CreateThresholdEndpoint(context.Background()).WalletId(walletId).TenantId(tenantId).PoolId(poolId).CreateThresholdRequest(createThresholdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CreateThresholdEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateThresholdEndpoint`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.CreateThresholdEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateThresholdEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **walletId** | **string** | Wallet identifier | 
 **tenantId** | **string** | Tenant identifier | 
 **poolId** | **string** | Pool identifier | 
 **createThresholdRequest** | [**CreateThresholdRequest**](CreateThresholdRequest.md) |  | 

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


## CreateTokenEndpoint

> interface{} CreateTokenEndpoint(ctx).CreatePortalTokenRequest(createPortalTokenRequest).Execute()

Create Token Endpoint



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
	createPortalTokenRequest := *openapiclient.NewCreatePortalTokenRequest("Subject_example") // CreatePortalTokenRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CreateTokenEndpoint(context.Background()).CreatePortalTokenRequest(createPortalTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CreateTokenEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTokenEndpoint`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.CreateTokenEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTokenEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPortalTokenRequest** | [**CreatePortalTokenRequest**](CreatePortalTokenRequest.md) |  | 

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


## CreateValueRecognitionEntryEndpointV1Fx

> interface{} CreateValueRecognitionEntryEndpointV1Fx(ctx).CreateValueRecognitionRequest(createValueRecognitionRequest).Execute()

Create Value Recognition Entry Endpoint



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
	createValueRecognitionRequest := *openapiclient.NewCreateValueRecognitionRequest("TenantId_example", "PoolId_example", "WalletId_example", "SubjectId_example", int32(123), time.Now()) // CreateValueRecognitionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CreateValueRecognitionEntryEndpointV1Fx(context.Background()).CreateValueRecognitionRequest(createValueRecognitionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CreateValueRecognitionEntryEndpointV1Fx``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateValueRecognitionEntryEndpointV1Fx`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.CreateValueRecognitionEntryEndpointV1Fx`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateValueRecognitionEntryEndpointV1FxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createValueRecognitionRequest** | [**CreateValueRecognitionRequest**](CreateValueRecognitionRequest.md) |  | 

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


## CreateWallet

> WalletResponse CreateWallet(ctx).CreateWalletRequest(createWalletRequest).Execute()

Create Wallet



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
	createWalletRequest := *openapiclient.NewCreateWalletRequest("TenantId_example", "PoolId_example", openapiclient.WalletOwnerType("ORG"), "OwnerId_example") // CreateWalletRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CreateWallet(context.Background()).CreateWalletRequest(createWalletRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CreateWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWallet`: WalletResponse
	fmt.Fprintf(os.Stdout, "Response from `V1API.CreateWallet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWalletRequest** | [**CreateWalletRequest**](CreateWalletRequest.md) |  | 

### Return type

[**WalletResponse**](WalletResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRuleV1

> interface{} DeleteRuleV1(ctx, ruleId).TenantId(tenantId).PoolId(poolId).WalletId(walletId).Execute()

Delete Rule



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
	ruleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Auto top-up rule ID
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional tenant identifier for validation (optional)
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional pool identifier for validation (optional)
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional wallet identifier for validation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.DeleteRuleV1(context.Background(), ruleId).TenantId(tenantId).PoolId(poolId).WalletId(walletId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.DeleteRuleV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRuleV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.DeleteRuleV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** | Auto top-up rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRuleV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** | Optional tenant identifier for validation | 
 **poolId** | **string** | Optional pool identifier for validation | 
 **walletId** | **string** | Optional wallet identifier for validation | 

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


## DeleteThresholdEndpoint

> interface{} DeleteThresholdEndpoint(ctx, thresholdId).WalletId(walletId).TenantId(tenantId).PoolId(poolId).Execute()

Delete Threshold Endpoint



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
	thresholdId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Wallet identifier
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Pool identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.DeleteThresholdEndpoint(context.Background(), thresholdId).WalletId(walletId).TenantId(tenantId).PoolId(poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.DeleteThresholdEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteThresholdEndpoint`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.DeleteThresholdEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thresholdId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteThresholdEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **walletId** | **string** | Wallet identifier | 
 **tenantId** | **string** | Tenant identifier | 
 **poolId** | **string** | Pool identifier | 

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


## DiagnoseSubscription

> interface{} DiagnoseSubscription(ctx, subscriptionId).TenantId(tenantId).Execute()

Diagnose Subscription



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
	subscriptionId := "subscriptionId_example" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.DiagnoseSubscription(context.Background(), subscriptionId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.DiagnoseSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiagnoseSubscription`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.DiagnoseSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDiagnoseSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** | Tenant identifier | 

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


## ExplainQuery

> interface{} ExplainQuery(ctx).QuerySql(querySql).Execute()

Explain Query



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
	querySql := "querySql_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.ExplainQuery(context.Background()).QuerySql(querySql).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.ExplainQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExplainQuery`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.ExplainQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExplainQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **querySql** | **string** |  | 

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


## GetActivityV1

> AutoTopupActivityResponse GetActivityV1(ctx).TenantId(tenantId).PoolId(poolId).WalletId(walletId).Limit(limit).Offset(offset).Execute()

Get Activity



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Pool identifier
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Wallet identifier
	limit := int32(56) // int32 | Maximum number of items to return (optional) (default to 50)
	offset := int32(56) // int32 | Offset for pagination (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.GetActivityV1(context.Background()).TenantId(tenantId).PoolId(poolId).WalletId(walletId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.GetActivityV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActivityV1`: AutoTopupActivityResponse
	fmt.Fprintf(os.Stdout, "Response from `V1API.GetActivityV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetActivityV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Tenant identifier | 
 **poolId** | **string** | Pool identifier | 
 **walletId** | **string** | Wallet identifier | 
 **limit** | **int32** | Maximum number of items to return | [default to 50]
 **offset** | **int32** | Offset for pagination | [default to 0]

### Return type

[**AutoTopupActivityResponse**](AutoTopupActivityResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAffectedEventsEndpointV1

> interface{} GetAffectedEventsEndpointV1(ctx).TenantId(tenantId).PoolId(poolId).WalletIds(walletIds).FromEffectiveAt(fromEffectiveAt).ToEffectiveAt(toEffectiveAt).Execute()

Get Affected Events Endpoint



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Pool identifier
	walletIds := "walletIds_example" // string | Comma-separated wallet IDs
	fromEffectiveAt := time.Now() // time.Time | Start of effective time window
	toEffectiveAt := time.Now() // time.Time | End of effective time window

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.GetAffectedEventsEndpointV1(context.Background()).TenantId(tenantId).PoolId(poolId).WalletIds(walletIds).FromEffectiveAt(fromEffectiveAt).ToEffectiveAt(toEffectiveAt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.GetAffectedEventsEndpointV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAffectedEventsEndpointV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.GetAffectedEventsEndpointV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAffectedEventsEndpointV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Tenant identifier | 
 **poolId** | **string** | Pool identifier | 
 **walletIds** | **string** | Comma-separated wallet IDs | 
 **fromEffectiveAt** | **time.Time** | Start of effective time window | 
 **toEffectiveAt** | **time.Time** | End of effective time window | 

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


## GetAlertsStateEndpoint

> interface{} GetAlertsStateEndpoint(ctx).WalletId(walletId).TenantId(tenantId).PoolId(poolId).Execute()

Get Alerts State Endpoint



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Wallet identifier
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Pool identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.GetAlertsStateEndpoint(context.Background()).WalletId(walletId).TenantId(tenantId).PoolId(poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.GetAlertsStateEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlertsStateEndpoint`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.GetAlertsStateEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertsStateEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **walletId** | **string** | Wallet identifier | 
 **tenantId** | **string** | Tenant identifier | 
 **poolId** | **string** | Pool identifier | 

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


## GetConnectionPoolInfoV1

> interface{} GetConnectionPoolInfoV1(ctx).Execute()

Get Connection Pool Info



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
	resp, r, err := apiClient.V1API.GetConnectionPoolInfoV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.GetConnectionPoolInfoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectionPoolInfoV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.GetConnectionPoolInfoV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionPoolInfoV1Request struct via the builder pattern


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


## GetFxRateAtEndpointV1

> interface{} GetFxRateAtEndpointV1(ctx).GetFxRateRequest(getFxRateRequest).Execute()

Get Fx Rate At Endpoint



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
	getFxRateRequest := *openapiclient.NewGetFxRateRequest("TenantId_example", "PoolId_example", time.Now()) // GetFxRateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.GetFxRateAtEndpointV1(context.Background()).GetFxRateRequest(getFxRateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.GetFxRateAtEndpointV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFxRateAtEndpointV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.GetFxRateAtEndpointV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFxRateAtEndpointV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getFxRateRequest** | [**GetFxRateRequest**](GetFxRateRequest.md) |  | 

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


## GetGrant

> GrantResponse GetGrant(ctx, grantId).Execute()

Get Grant



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
	grantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.GetGrant(context.Background(), grantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.GetGrant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGrant`: GrantResponse
	fmt.Fprintf(os.Stdout, "Response from `V1API.GetGrant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**grantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GrantResponse**](GrantResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndexUsage

> interface{} GetIndexUsage(ctx).Schema(schema).Execute()

Get Index Usage



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
	schema := "schema_example" // string | Schema name (optional) (default to "billing")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.GetIndexUsage(context.Background()).Schema(schema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.GetIndexUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIndexUsage`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.GetIndexUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIndexUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schema** | **string** | Schema name | [default to &quot;billing&quot;]

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


## GetLateEventsV1

> interface{} GetLateEventsV1(ctx).TenantId(tenantId).PoolId(poolId).LateThresholdSeconds(lateThresholdSeconds).Execute()

Get Late Events



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Pool identifier
	lateThresholdSeconds := int32(56) // int32 | LATE event threshold (seconds) (optional) (default to 3600)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.GetLateEventsV1(context.Background()).TenantId(tenantId).PoolId(poolId).LateThresholdSeconds(lateThresholdSeconds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.GetLateEventsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLateEventsV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.GetLateEventsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLateEventsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Tenant identifier | 
 **poolId** | **string** | Pool identifier | 
 **lateThresholdSeconds** | **int32** | LATE event threshold (seconds) | [default to 3600]

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


## GetLedgerAudit

> interface{} GetLedgerAudit(ctx).TenantId(tenantId).PoolId(poolId).WalletId(walletId).JournalEntryId(journalEntryId).UsageEventId(usageEventId).FromEffectiveAt(fromEffectiveAt).ToEffectiveAt(toEffectiveAt).Limit(limit).Execute()

Get Ledger Audit



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Pool identifier
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional wallet filter (optional)
	journalEntryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional journal entry filter (optional)
	usageEventId := "usageEventId_example" // string | Optional usage event filter (optional)
	fromEffectiveAt := time.Now() // time.Time | Optional start time filter (optional)
	toEffectiveAt := time.Now() // time.Time | Optional end time filter (optional)
	limit := int32(56) // int32 | Maximum number of entries (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.GetLedgerAudit(context.Background()).TenantId(tenantId).PoolId(poolId).WalletId(walletId).JournalEntryId(journalEntryId).UsageEventId(usageEventId).FromEffectiveAt(fromEffectiveAt).ToEffectiveAt(toEffectiveAt).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.GetLedgerAudit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLedgerAudit`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.GetLedgerAudit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLedgerAuditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Tenant identifier | 
 **poolId** | **string** | Pool identifier | 
 **walletId** | **string** | Optional wallet filter | 
 **journalEntryId** | **string** | Optional journal entry filter | 
 **usageEventId** | **string** | Optional usage event filter | 
 **fromEffectiveAt** | **time.Time** | Optional start time filter | 
 **toEffectiveAt** | **time.Time** | Optional end time filter | 
 **limit** | **int32** | Maximum number of entries | [default to 100]

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


## GetLedgerAuditGet

> interface{} GetLedgerAuditGet(ctx).TenantId(tenantId).PoolId(poolId).WalletId(walletId).JournalEntryId(journalEntryId).UsageEventId(usageEventId).FromEffectiveAt(fromEffectiveAt).ToEffectiveAt(toEffectiveAt).Limit(limit).Execute()

Get Ledger Audit



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Pool identifier
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional wallet filter (optional)
	journalEntryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional journal entry filter (optional)
	usageEventId := "usageEventId_example" // string | Optional usage event filter (optional)
	fromEffectiveAt := time.Now() // time.Time | Optional start time filter (optional)
	toEffectiveAt := time.Now() // time.Time | Optional end time filter (optional)
	limit := int32(56) // int32 | Maximum number of entries (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.GetLedgerAuditGet(context.Background()).TenantId(tenantId).PoolId(poolId).WalletId(walletId).JournalEntryId(journalEntryId).UsageEventId(usageEventId).FromEffectiveAt(fromEffectiveAt).ToEffectiveAt(toEffectiveAt).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.GetLedgerAuditGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLedgerAuditGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.GetLedgerAuditGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLedgerAuditGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Tenant identifier | 
 **poolId** | **string** | Pool identifier | 
 **walletId** | **string** | Optional wallet filter | 
 **journalEntryId** | **string** | Optional journal entry filter | 
 **usageEventId** | **string** | Optional usage event filter | 
 **fromEffectiveAt** | **time.Time** | Optional start time filter | 
 **toEffectiveAt** | **time.Time** | Optional end time filter | 
 **limit** | **int32** | Maximum number of entries | [default to 100]

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


## GetLedgerBalance

> interface{} GetLedgerBalance(ctx).TenantId(tenantId).PoolId(poolId).WalletId(walletId).AsOfEffectiveAt(asOfEffectiveAt).AsOfRecordedAt(asOfRecordedAt).Execute()

Get Ledger Balance



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Pool identifier
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Wallet identifier
	asOfEffectiveAt := time.Now() // time.Time | Effective timestamp (for time travel)
	asOfRecordedAt := time.Now() // time.Time | Recorded timestamp (for time travel) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.GetLedgerBalance(context.Background()).TenantId(tenantId).PoolId(poolId).WalletId(walletId).AsOfEffectiveAt(asOfEffectiveAt).AsOfRecordedAt(asOfRecordedAt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.GetLedgerBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLedgerBalance`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.GetLedgerBalance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLedgerBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Tenant identifier | 
 **poolId** | **string** | Pool identifier | 
 **walletId** | **string** | Wallet identifier | 
 **asOfEffectiveAt** | **time.Time** | Effective timestamp (for time travel) | 
 **asOfRecordedAt** | **time.Time** | Recorded timestamp (for time travel) | 

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


## GetLedgerBalanceGet

> interface{} GetLedgerBalanceGet(ctx).TenantId(tenantId).PoolId(poolId).WalletId(walletId).AsOfEffectiveAt(asOfEffectiveAt).AsOfRecordedAt(asOfRecordedAt).Execute()

Get Ledger Balance



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Pool identifier
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Wallet identifier
	asOfEffectiveAt := time.Now() // time.Time | Effective timestamp (for time travel)
	asOfRecordedAt := time.Now() // time.Time | Recorded timestamp (for time travel) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.GetLedgerBalanceGet(context.Background()).TenantId(tenantId).PoolId(poolId).WalletId(walletId).AsOfEffectiveAt(asOfEffectiveAt).AsOfRecordedAt(asOfRecordedAt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.GetLedgerBalanceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLedgerBalanceGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.GetLedgerBalanceGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLedgerBalanceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Tenant identifier | 
 **poolId** | **string** | Pool identifier | 
 **walletId** | **string** | Wallet identifier | 
 **asOfEffectiveAt** | **time.Time** | Effective timestamp (for time travel) | 
 **asOfRecordedAt** | **time.Time** | Recorded timestamp (for time travel) | 

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


## GetMissingIndexRecommendations

> interface{} GetMissingIndexRecommendations(ctx).Schema(schema).Execute()

Get Missing Index Recommendations



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
	schema := "schema_example" // string | Schema name (optional) (default to "billing")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.GetMissingIndexRecommendations(context.Background()).Schema(schema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.GetMissingIndexRecommendations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMissingIndexRecommendations`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.GetMissingIndexRecommendations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMissingIndexRecommendationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schema** | **string** | Schema name | [default to &quot;billing&quot;]

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


## GetPerformanceReportEndpoint

> interface{} GetPerformanceReportEndpoint(ctx).Schema(schema).Execute()

Get Performance Report Endpoint



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
	schema := "schema_example" // string | Schema name (optional) (default to "billing")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.GetPerformanceReportEndpoint(context.Background()).Schema(schema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.GetPerformanceReportEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPerformanceReportEndpoint`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.GetPerformanceReportEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPerformanceReportEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schema** | **string** | Schema name | [default to &quot;billing&quot;]

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


## GetPortalContext

> interface{} GetPortalContext(ctx).Execute()

Get Portal Context



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
	resp, r, err := apiClient.V1API.GetPortalContext(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.GetPortalContext``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortalContext`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.GetPortalContext`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortalContextRequest struct via the builder pattern


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


## GetRateCard

> RateCardResponse GetRateCard(ctx, rateCardId).Execute()

Get Rate Card



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
	rateCardId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.GetRateCard(context.Background(), rateCardId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.GetRateCard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRateCard`: RateCardResponse
	fmt.Fprintf(os.Stdout, "Response from `V1API.GetRateCard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rateCardId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRateCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RateCardResponse**](RateCardResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuleV1

> AutoTopupRuleResponse GetRuleV1(ctx, ruleId).TenantId(tenantId).PoolId(poolId).Execute()

Get Rule



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
	ruleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Auto top-up rule ID
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional tenant identifier for validation (optional)
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional pool identifier for validation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.GetRuleV1(context.Background(), ruleId).TenantId(tenantId).PoolId(poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.GetRuleV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRuleV1`: AutoTopupRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `V1API.GetRuleV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** | Auto top-up rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** | Optional tenant identifier for validation | 
 **poolId** | **string** | Optional pool identifier for validation | 

### Return type

[**AutoTopupRuleResponse**](AutoTopupRuleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSlowQueries

> interface{} GetSlowQueries(ctx).MinDurationMs(minDurationMs).Limit(limit).Execute()

Get Slow Queries



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
	minDurationMs := float32(8.14) // float32 | Minimum query duration in milliseconds (optional) (default to 100.0)
	limit := int32(56) // int32 | Maximum number of results (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.GetSlowQueries(context.Background()).MinDurationMs(minDurationMs).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.GetSlowQueries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSlowQueries`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.GetSlowQueries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSlowQueriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **minDurationMs** | **float32** | Minimum query duration in milliseconds | [default to 100.0]
 **limit** | **int32** | Maximum number of results | [default to 20]

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


## GetSnapshotAtEndpoint

> interface{} GetSnapshotAtEndpoint(ctx, walletId).TenantId(tenantId).PoolId(poolId).AsOfRecordedAt(asOfRecordedAt).Execute()

Get Snapshot At Endpoint



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Wallet identifier
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Pool identifier
	asOfRecordedAt := time.Now() // time.Time | Recorded timestamp

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.GetSnapshotAtEndpoint(context.Background(), walletId).TenantId(tenantId).PoolId(poolId).AsOfRecordedAt(asOfRecordedAt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.GetSnapshotAtEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSnapshotAtEndpoint`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.GetSnapshotAtEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** | Wallet identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSnapshotAtEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** | Tenant identifier | 
 **poolId** | **string** | Pool identifier | 
 **asOfRecordedAt** | **time.Time** | Recorded timestamp | 

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


## GetSubscription

> interface{} GetSubscription(ctx, subscriptionId).TenantId(tenantId).AsOf(asOf).Execute()

Get Subscription



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
	subscriptionId := "subscriptionId_example" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	asOf := "asOf_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.GetSubscription(context.Background(), subscriptionId).TenantId(tenantId).AsOf(asOf).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.GetSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubscription`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.GetSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **asOf** | **string** |  | 

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


## GetTableSizeAnalysis

> interface{} GetTableSizeAnalysis(ctx).Schema(schema).Execute()

Get Table Size Analysis



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
	schema := "schema_example" // string | Schema name (optional) (default to "billing")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.GetTableSizeAnalysis(context.Background()).Schema(schema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.GetTableSizeAnalysis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTableSizeAnalysis`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.GetTableSizeAnalysis`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTableSizeAnalysisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schema** | **string** | Schema name | [default to &quot;billing&quot;]

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


## GetUsageEvent

> interface{} GetUsageEvent(ctx, eventId).Execute()

Get Usage Event



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.GetUsageEvent(context.Background(), eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.GetUsageEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsageEvent`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.GetUsageEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetWallet

> WalletResponse GetWallet(ctx, walletId).Execute()

Get Wallet



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.GetWallet(context.Background(), walletId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.GetWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWallet`: WalletResponse
	fmt.Fprintf(os.Stdout, "Response from `V1API.GetWallet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WalletResponse**](WalletResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletActivity

> interface{} GetWalletActivity(ctx, walletId).FromEffectiveAt(fromEffectiveAt).ToEffectiveAt(toEffectiveAt).EffectiveAsOf(effectiveAsOf).RecordedAsOf(recordedAsOf).ConsistentView(consistentView).Limit(limit).Cursor(cursor).Execute()

Get Wallet Activity



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	fromEffectiveAt := time.Now() // time.Time | Start time for activity range (optional)
	toEffectiveAt := time.Now() // time.Time | End time for activity range (optional)
	effectiveAsOf := time.Now() // time.Time | Effective as-of timestamp (business time) for time travel (optional)
	recordedAsOf := time.Now() // time.Time | Recorded as-of timestamp (system time) for time travel (optional)
	consistentView := true // bool | Use strong consistency for reads (optional)
	limit := int32(56) // int32 | Maximum number of items to return (optional) (default to 50)
	cursor := "cursor_example" // string | Pagination cursor (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.GetWalletActivity(context.Background(), walletId).FromEffectiveAt(fromEffectiveAt).ToEffectiveAt(toEffectiveAt).EffectiveAsOf(effectiveAsOf).RecordedAsOf(recordedAsOf).ConsistentView(consistentView).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.GetWalletActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletActivity`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.GetWalletActivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromEffectiveAt** | **time.Time** | Start time for activity range | 
 **toEffectiveAt** | **time.Time** | End time for activity range | 
 **effectiveAsOf** | **time.Time** | Effective as-of timestamp (business time) for time travel | 
 **recordedAsOf** | **time.Time** | Recorded as-of timestamp (system time) for time travel | 
 **consistentView** | **bool** | Use strong consistency for reads | 
 **limit** | **int32** | Maximum number of items to return | [default to 50]
 **cursor** | **string** | Pagination cursor | 

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


## GetWalletMembers

> interface{} GetWalletMembers(ctx, walletId).Execute()

Get Wallet Members



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.GetWalletMembers(context.Background(), walletId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.GetWalletMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletMembers`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.GetWalletMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetWalletState

> interface{} GetWalletState(ctx).TenantId(tenantId).PoolId(poolId).WalletId(walletId).AsOf(asOf).Consistency(consistency).EffectiveAsOf(effectiveAsOf).RecordedAsOf(recordedAsOf).ConsistentView(consistentView).Execute()

Get Wallet State



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Credit pool identifier
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Wallet identifier
	asOf := time.Now() // time.Time | As-of timestamp (for time travel, legacy, use effective_as_of instead) (optional)
	consistency := "consistency_example" // string | Consistency level: 'eventual' or 'STRONG' (optional) (default to "eventual")
	effectiveAsOf := time.Now() // time.Time | Effective as-of timestamp (business time) for time travel (optional)
	recordedAsOf := time.Now() // time.Time | Recorded as-of timestamp (system time) for time travel (optional)
	consistentView := true // bool | Use strong consistency (overrides consistency param if True) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.GetWalletState(context.Background()).TenantId(tenantId).PoolId(poolId).WalletId(walletId).AsOf(asOf).Consistency(consistency).EffectiveAsOf(effectiveAsOf).RecordedAsOf(recordedAsOf).ConsistentView(consistentView).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.GetWalletState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletState`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.GetWalletState`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Tenant identifier | 
 **poolId** | **string** | Credit pool identifier | 
 **walletId** | **string** | Wallet identifier | 
 **asOf** | **time.Time** | As-of timestamp (for time travel, legacy, use effective_as_of instead) | 
 **consistency** | **string** | Consistency level: &#39;eventual&#39; or &#39;STRONG&#39; | [default to &quot;eventual&quot;]
 **effectiveAsOf** | **time.Time** | Effective as-of timestamp (business time) for time travel | 
 **recordedAsOf** | **time.Time** | Recorded as-of timestamp (system time) for time travel | 
 **consistentView** | **bool** | Use strong consistency (overrides consistency param if True) | 

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


## GetWalletStateProjectionEndpointV1

> interface{} GetWalletStateProjectionEndpointV1(ctx, walletId).TenantId(tenantId).PoolId(poolId).AsOf(asOf).Execute()

Get Wallet State Projection Endpoint



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Pool identifier
	asOf := time.Now() // time.Time | As-of timestamp (for time travel) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.GetWalletStateProjectionEndpointV1(context.Background(), walletId).TenantId(tenantId).PoolId(poolId).AsOf(asOf).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.GetWalletStateProjectionEndpointV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletStateProjectionEndpointV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.GetWalletStateProjectionEndpointV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletStateProjectionEndpointV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** | Tenant identifier | 
 **poolId** | **string** | Pool identifier | 
 **asOf** | **time.Time** | As-of timestamp (for time travel) | 

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


## GetWalletSubtree

> interface{} GetWalletSubtree(ctx, walletId).TenantId(tenantId).PoolId(poolId).MaxDepth(maxDepth).Execute()

Get Wallet Subtree



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Wallet identifier
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Pool identifier
	maxDepth := int32(56) // int32 | Maximum depth to traverse (optional) (default to 3)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.GetWalletSubtree(context.Background(), walletId).TenantId(tenantId).PoolId(poolId).MaxDepth(maxDepth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.GetWalletSubtree``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletSubtree`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.GetWalletSubtree`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** | Wallet identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletSubtreeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** | Tenant identifier | 
 **poolId** | **string** | Pool identifier | 
 **maxDepth** | **int32** | Maximum depth to traverse | [default to 3]

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


## HandleLifecycleEvent

> interface{} HandleLifecycleEvent(ctx).TenantId(tenantId).LifecycleEventRequest(lifecycleEventRequest).Execute()

Handle Lifecycle Event



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	lifecycleEventRequest := *openapiclient.NewLifecycleEventRequest("SubscriptionId_example", "EventType_example", "EventId_example", "EffectiveAt_example") // LifecycleEventRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.HandleLifecycleEvent(context.Background()).TenantId(tenantId).LifecycleEventRequest(lifecycleEventRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.HandleLifecycleEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleLifecycleEvent`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.HandleLifecycleEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHandleLifecycleEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **lifecycleEventRequest** | [**LifecycleEventRequest**](LifecycleEventRequest.md) |  | 

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


## InvalidateTokensEndpoint

> interface{} InvalidateTokensEndpoint(ctx).InvalidatePortalTokenRequest(invalidatePortalTokenRequest).Execute()

Invalidate Tokens Endpoint



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
	invalidatePortalTokenRequest := *openapiclient.NewInvalidatePortalTokenRequest() // InvalidatePortalTokenRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.InvalidateTokensEndpoint(context.Background()).InvalidatePortalTokenRequest(invalidatePortalTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.InvalidateTokensEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvalidateTokensEndpoint`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.InvalidateTokensEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvalidateTokensEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invalidatePortalTokenRequest** | [**InvalidatePortalTokenRequest**](InvalidatePortalTokenRequest.md) |  | 

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


## ListGrants

> GrantsListResponse ListGrants(ctx).TenantId(tenantId).PoolId(poolId).WalletId(walletId).SubjectId(subjectId).SubscriptionId(subscriptionId).FeatureKey(featureKey).MeterSlug(meterSlug).AsOf(asOf).EffectiveAsOf(effectiveAsOf).RecordedAsOf(recordedAsOf).ConsistentView(consistentView).Limit(limit).Offset(offset).Execute()

List Grants



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier (optional if subject_id or subscription_id provided) (optional)
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Pool identifier (optional if subject_id or subscription_id provided) (optional)
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Wallet identifier to filter by (optional)
	subjectId := "subjectId_example" // string | Subject identifier (for finding wallet, auto-resolves tenant_id/pool_id) (optional)
	subscriptionId := "subscriptionId_example" // string | Subscription identifier (filters grants by subscription, auto-resolves tenant_id/pool_id) (optional)
	featureKey := "featureKey_example" // string | Feature key for scope filtering (optional)
	meterSlug := "meterSlug_example" // string | Meter slug for scope filtering (optional)
	asOf := time.Now() // time.Time | As-of timestamp for time travel (legacy, use effective_as_of instead) (optional)
	effectiveAsOf := time.Now() // time.Time | Effective as-of timestamp (business time) for time travel (optional)
	recordedAsOf := time.Now() // time.Time | Recorded as-of timestamp (system time) for time travel (optional)
	consistentView := true // bool | Use strong consistency for reads (optional)
	limit := int32(56) // int32 | Maximum number of grants to return (optional) (default to 100)
	offset := int32(56) // int32 | Offset for pagination (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.ListGrants(context.Background()).TenantId(tenantId).PoolId(poolId).WalletId(walletId).SubjectId(subjectId).SubscriptionId(subscriptionId).FeatureKey(featureKey).MeterSlug(meterSlug).AsOf(asOf).EffectiveAsOf(effectiveAsOf).RecordedAsOf(recordedAsOf).ConsistentView(consistentView).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.ListGrants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGrants`: GrantsListResponse
	fmt.Fprintf(os.Stdout, "Response from `V1API.ListGrants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGrantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Tenant identifier (optional if subject_id or subscription_id provided) | 
 **poolId** | **string** | Pool identifier (optional if subject_id or subscription_id provided) | 
 **walletId** | **string** | Wallet identifier to filter by | 
 **subjectId** | **string** | Subject identifier (for finding wallet, auto-resolves tenant_id/pool_id) | 
 **subscriptionId** | **string** | Subscription identifier (filters grants by subscription, auto-resolves tenant_id/pool_id) | 
 **featureKey** | **string** | Feature key for scope filtering | 
 **meterSlug** | **string** | Meter slug for scope filtering | 
 **asOf** | **time.Time** | As-of timestamp for time travel (legacy, use effective_as_of instead) | 
 **effectiveAsOf** | **time.Time** | Effective as-of timestamp (business time) for time travel | 
 **recordedAsOf** | **time.Time** | Recorded as-of timestamp (system time) for time travel | 
 **consistentView** | **bool** | Use strong consistency for reads | 
 **limit** | **int32** | Maximum number of grants to return | [default to 100]
 **offset** | **int32** | Offset for pagination | [default to 0]

### Return type

[**GrantsListResponse**](GrantsListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMappings

> []MappingResponse ListMappings(ctx).TenantId(tenantId).MeterSlug(meterSlug).FeatureKey(featureKey).Execute()

List Mappings



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier
	meterSlug := "meterSlug_example" // string | Filter by meter slug (optional)
	featureKey := "featureKey_example" // string | Filter by feature key (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.ListMappings(context.Background()).TenantId(tenantId).MeterSlug(meterSlug).FeatureKey(featureKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.ListMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMappings`: []MappingResponse
	fmt.Fprintf(os.Stdout, "Response from `V1API.ListMappings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Tenant identifier | 
 **meterSlug** | **string** | Filter by meter slug | 
 **featureKey** | **string** | Filter by feature key | 

### Return type

[**[]MappingResponse**](MappingResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOutboxEventsEndpoint

> interface{} ListOutboxEventsEndpoint(ctx).TenantId(tenantId).EventType(eventType).Status(status).Limit(limit).Execute()

List Outbox Events Endpoint



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter by tenant ID (optional)
	eventType := "eventType_example" // string | Filter by event type (optional)
	status := "status_example" // string | Filter by status: pending, delivered, failed (optional)
	limit := int32(56) // int32 | Maximum events to return (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.ListOutboxEventsEndpoint(context.Background()).TenantId(tenantId).EventType(eventType).Status(status).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.ListOutboxEventsEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOutboxEventsEndpoint`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.ListOutboxEventsEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOutboxEventsEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Filter by tenant ID | 
 **eventType** | **string** | Filter by event type | 
 **status** | **string** | Filter by status: pending, delivered, failed | 
 **limit** | **int32** | Maximum events to return | [default to 100]

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


## ListRulesV1

> ListAutoTopupRulesResponse ListRulesV1(ctx).TenantId(tenantId).PoolId(poolId).WalletId(walletId).Execute()

List Rules



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Pool identifier
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional wallet identifier to filter by (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.ListRulesV1(context.Background()).TenantId(tenantId).PoolId(poolId).WalletId(walletId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.ListRulesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRulesV1`: ListAutoTopupRulesResponse
	fmt.Fprintf(os.Stdout, "Response from `V1API.ListRulesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRulesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Tenant identifier | 
 **poolId** | **string** | Pool identifier | 
 **walletId** | **string** | Optional wallet identifier to filter by | 

### Return type

[**ListAutoTopupRulesResponse**](ListAutoTopupRulesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSnapshotsEndpoint

> interface{} ListSnapshotsEndpoint(ctx, walletId).TenantId(tenantId).PoolId(poolId).Limit(limit).Execute()

List Snapshots Endpoint



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Wallet identifier
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Pool identifier
	limit := int32(56) // int32 | Maximum number of snapshots (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.ListSnapshotsEndpoint(context.Background(), walletId).TenantId(tenantId).PoolId(poolId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.ListSnapshotsEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSnapshotsEndpoint`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.ListSnapshotsEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** | Wallet identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSnapshotsEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** | Tenant identifier | 
 **poolId** | **string** | Pool identifier | 
 **limit** | **int32** | Maximum number of snapshots | [default to 100]

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


## ListThresholdsEndpoint

> interface{} ListThresholdsEndpoint(ctx).WalletId(walletId).TenantId(tenantId).PoolId(poolId).Execute()

List Thresholds Endpoint



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Wallet identifier
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Pool identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.ListThresholdsEndpoint(context.Background()).WalletId(walletId).TenantId(tenantId).PoolId(poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.ListThresholdsEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListThresholdsEndpoint`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.ListThresholdsEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListThresholdsEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **walletId** | **string** | Wallet identifier | 
 **tenantId** | **string** | Tenant identifier | 
 **poolId** | **string** | Pool identifier | 

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


## ListTokensEndpoint

> interface{} ListTokensEndpoint(ctx).Subject(subject).Execute()

List Tokens Endpoint



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
	subject := "subject_example" // string | Filter by subject (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.ListTokensEndpoint(context.Background()).Subject(subject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.ListTokensEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTokensEndpoint`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.ListTokensEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTokensEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subject** | **string** | Filter by subject | 

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


## ListUnpricedEvents

> interface{} ListUnpricedEvents(ctx).TenantId(tenantId).PoolId(poolId).Limit(limit).Execute()

List Unpriced Events



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter by tenant ID (optional)
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter by pool ID (optional)
	limit := int32(56) // int32 | Maximum events to return (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.ListUnpricedEvents(context.Background()).TenantId(tenantId).PoolId(poolId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.ListUnpricedEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUnpricedEvents`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.ListUnpricedEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUnpricedEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Filter by tenant ID | 
 **poolId** | **string** | Filter by pool ID | 
 **limit** | **int32** | Maximum events to return | [default to 100]

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


## ListUsageEvents

> interface{} ListUsageEvents(ctx).TenantId(tenantId).PoolId(poolId).RatingStatus(ratingStatus).FeatureKey(featureKey).MeterSlug(meterSlug).SubjectId(subjectId).FromTime(fromTime).ToTime(toTime).Limit(limit).Offset(offset).Execute()

List Usage Events



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier (optional if subject_id provided) (optional)
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Pool identifier (optional if subject_id provided) (optional)
	ratingStatus := "ratingStatus_example" // string | Filter by rating status (optional)
	featureKey := "featureKey_example" // string | Filter by feature key (optional)
	meterSlug := "meterSlug_example" // string | Filter by meter slug (optional)
	subjectId := "subjectId_example" // string | Filter by subject ID (auto-resolves tenant_id/pool_id if not provided) (optional)
	fromTime := "fromTime_example" // string | Filter events from this time (ISO 8601, e.g., 2026-02-09T00:48:00Z) (optional)
	toTime := "toTime_example" // string | Filter events to this time (ISO 8601, e.g., 2026-02-09T00:50:00Z) (optional)
	limit := int32(56) // int32 | Maximum number of events to return (optional) (default to 100)
	offset := int32(56) // int32 | Offset for pagination (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.ListUsageEvents(context.Background()).TenantId(tenantId).PoolId(poolId).RatingStatus(ratingStatus).FeatureKey(featureKey).MeterSlug(meterSlug).SubjectId(subjectId).FromTime(fromTime).ToTime(toTime).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.ListUsageEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUsageEvents`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.ListUsageEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsageEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Tenant identifier (optional if subject_id provided) | 
 **poolId** | **string** | Pool identifier (optional if subject_id provided) | 
 **ratingStatus** | **string** | Filter by rating status | 
 **featureKey** | **string** | Filter by feature key | 
 **meterSlug** | **string** | Filter by meter slug | 
 **subjectId** | **string** | Filter by subject ID (auto-resolves tenant_id/pool_id if not provided) | 
 **fromTime** | **string** | Filter events from this time (ISO 8601, e.g., 2026-02-09T00:48:00Z) | 
 **toTime** | **string** | Filter events to this time (ISO 8601, e.g., 2026-02-09T00:50:00Z) | 
 **limit** | **int32** | Maximum number of events to return | [default to 100]
 **offset** | **int32** | Offset for pagination | [default to 0]

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


## ListWalletMembers

> interface{} ListWalletMembers(ctx).TenantId(tenantId).PoolId(poolId).SubjectId(subjectId).WalletId(walletId).Execute()

List Wallet Members



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Credit pool identifier
	subjectId := "subjectId_example" // string | Optional subject_id filter (optional)
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional wallet_id filter (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.ListWalletMembers(context.Background()).TenantId(tenantId).PoolId(poolId).SubjectId(subjectId).WalletId(walletId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.ListWalletMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWalletMembers`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.ListWalletMembers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWalletMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Tenant identifier | 
 **poolId** | **string** | Credit pool identifier | 
 **subjectId** | **string** | Optional subject_id filter | 
 **walletId** | **string** | Optional wallet_id filter | 

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


## ListWallets

> interface{} ListWallets(ctx).TenantId(tenantId).PoolId(poolId).SubjectId(subjectId).SubscriptionId(subscriptionId).OwnerType(ownerType).OwnerId(ownerId).EffectiveAsOf(effectiveAsOf).RecordedAsOf(recordedAsOf).ConsistentView(consistentView).Execute()

List Wallets



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier (optional if subject_id or subscription_id provided) (optional)
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Credit pool identifier (optional if subject_id or subscription_id provided) (optional)
	subjectId := "subjectId_example" // string | Subject identifier (if provided, tenant_id and pool_id are auto-resolved) (optional)
	subscriptionId := "subscriptionId_example" // string | Subscription identifier (finds wallets via grants with this subscription_id) (optional)
	ownerType := "ownerType_example" // string | Optional owner_type filter (optional)
	ownerId := "ownerId_example" // string | Optional owner_id filter (optional)
	effectiveAsOf := time.Now() // time.Time | Effective as-of timestamp (business time) for time travel (optional)
	recordedAsOf := time.Now() // time.Time | Recorded as-of timestamp (system time) for time travel (optional)
	consistentView := true // bool | Use strong consistency for reads (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.ListWallets(context.Background()).TenantId(tenantId).PoolId(poolId).SubjectId(subjectId).SubscriptionId(subscriptionId).OwnerType(ownerType).OwnerId(ownerId).EffectiveAsOf(effectiveAsOf).RecordedAsOf(recordedAsOf).ConsistentView(consistentView).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.ListWallets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWallets`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.ListWallets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWalletsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Tenant identifier (optional if subject_id or subscription_id provided) | 
 **poolId** | **string** | Credit pool identifier (optional if subject_id or subscription_id provided) | 
 **subjectId** | **string** | Subject identifier (if provided, tenant_id and pool_id are auto-resolved) | 
 **subscriptionId** | **string** | Subscription identifier (finds wallets via grants with this subscription_id) | 
 **ownerType** | **string** | Optional owner_type filter | 
 **ownerId** | **string** | Optional owner_id filter | 
 **effectiveAsOf** | **time.Time** | Effective as-of timestamp (business time) for time travel | 
 **recordedAsOf** | **time.Time** | Recorded as-of timestamp (system time) for time travel | 
 **consistentView** | **bool** | Use strong consistency for reads | 

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


## MintGrant

> GrantResponse MintGrant(ctx).CreateGrantRequest(createGrantRequest).Execute()

Mint Grant



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
	createGrantRequest := *openapiclient.NewCreateGrantRequest("TenantId_example", "PoolId_example", "WalletId_example", int32(123), time.Now(), openapiclient.GrantSourceType("SUBSCRIPTION_ALLOWANCE"), "SourceId_example") // CreateGrantRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MintGrant(context.Background()).CreateGrantRequest(createGrantRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MintGrant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MintGrant`: GrantResponse
	fmt.Fprintf(os.Stdout, "Response from `V1API.MintGrant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMintGrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createGrantRequest** | [**CreateGrantRequest**](CreateGrantRequest.md) |  | 

### Return type

[**GrantResponse**](GrantResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpenmeterWebhook

> interface{} OpenmeterWebhook(ctx).TenantId(tenantId).PoolId(poolId).Execute()

Openmeter Webhook



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.OpenmeterWebhook(context.Background()).TenantId(tenantId).PoolId(poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OpenmeterWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpenmeterWebhook`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.OpenmeterWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpenmeterWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **poolId** | **string** |  | 

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


## OpenmeterWebhookBatch

> interface{} OpenmeterWebhookBatch(ctx).TenantId(tenantId).PoolId(poolId).Execute()

Openmeter Webhook Batch



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.OpenmeterWebhookBatch(context.Background()).TenantId(tenantId).PoolId(poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.OpenmeterWebhookBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpenmeterWebhookBatch`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.OpenmeterWebhookBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpenmeterWebhookBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **poolId** | **string** |  | 

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


## PatchRuleV1

> AutoTopupRuleResponse PatchRuleV1(ctx, ruleId).TenantId(tenantId).PoolId(poolId).WalletId(walletId).UpdateAutoTopupRuleRequest(updateAutoTopupRuleRequest).Execute()

Patch Rule



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
	ruleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Auto top-up rule ID
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional tenant identifier for validation (optional)
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional pool identifier for validation (optional)
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional wallet identifier for validation (optional)
	updateAutoTopupRuleRequest := *openapiclient.NewUpdateAutoTopupRuleRequest() // UpdateAutoTopupRuleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PatchRuleV1(context.Background(), ruleId).TenantId(tenantId).PoolId(poolId).WalletId(walletId).UpdateAutoTopupRuleRequest(updateAutoTopupRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PatchRuleV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchRuleV1`: AutoTopupRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `V1API.PatchRuleV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** | Auto top-up rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchRuleV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** | Optional tenant identifier for validation | 
 **poolId** | **string** | Optional pool identifier for validation | 
 **walletId** | **string** | Optional wallet identifier for validation | 
 **updateAutoTopupRuleRequest** | [**UpdateAutoTopupRuleRequest**](UpdateAutoTopupRuleRequest.md) |  | 

### Return type

[**AutoTopupRuleResponse**](AutoTopupRuleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentSucceededV1

> PaymentSucceededResponse PaymentSucceededV1(ctx).PaymentSucceededRequest(paymentSucceededRequest).Execute()

Payment Succeeded



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
	paymentSucceededRequest := *openapiclient.NewPaymentSucceededRequest("TenantId_example", "PoolId_example", "WalletId_example", "PaymentId_example", int32(123)) // PaymentSucceededRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PaymentSucceededV1(context.Background()).PaymentSucceededRequest(paymentSucceededRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PaymentSucceededV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentSucceededV1`: PaymentSucceededResponse
	fmt.Fprintf(os.Stdout, "Response from `V1API.PaymentSucceededV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentSucceededV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentSucceededRequest** | [**PaymentSucceededRequest**](PaymentSucceededRequest.md) |  | 

### Return type

[**PaymentSucceededResponse**](PaymentSucceededResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProcessEvents

> interface{} ProcessEvents(ctx).TenantId(tenantId).PoolId(poolId).Limit(limit).MaxAttempts(maxAttempts).Execute()

Process Events



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter by tenant ID (optional)
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter by pool ID (optional)
	limit := int32(56) // int32 | Maximum events to process (optional) (default to 100)
	maxAttempts := int32(56) // int32 | Maximum retry attempts (optional) (default to 5)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.ProcessEvents(context.Background()).TenantId(tenantId).PoolId(poolId).Limit(limit).MaxAttempts(maxAttempts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.ProcessEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProcessEvents`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.ProcessEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProcessEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Filter by tenant ID | 
 **poolId** | **string** | Filter by pool ID | 
 **limit** | **int32** | Maximum events to process | [default to 100]
 **maxAttempts** | **int32** | Maximum retry attempts | [default to 5]

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


## ProcessLifecycleEventManuallyV1Admin

> interface{} ProcessLifecycleEventManuallyV1Admin(ctx, subscriptionId).TenantId(tenantId).Execute()

Process Lifecycle Event Manually



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
	subscriptionId := "subscriptionId_example" // string | Subscription ID
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.ProcessLifecycleEventManuallyV1Admin(context.Background(), subscriptionId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.ProcessLifecycleEventManuallyV1Admin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProcessLifecycleEventManuallyV1Admin`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.ProcessLifecycleEventManuallyV1Admin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProcessLifecycleEventManuallyV1AdminRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** | Tenant identifier | 

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


## ProcessOutboxEndpoint

> interface{} ProcessOutboxEndpoint(ctx).TenantId(tenantId).BatchSize(batchSize).MaxAttempts(maxAttempts).Execute()

Process Outbox Endpoint



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional tenant filter (optional)
	batchSize := int32(56) // int32 | Maximum events to process (optional) (default to 100)
	maxAttempts := int32(56) // int32 | Maximum retry attempts (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.ProcessOutboxEndpoint(context.Background()).TenantId(tenantId).BatchSize(batchSize).MaxAttempts(maxAttempts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.ProcessOutboxEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProcessOutboxEndpoint`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.ProcessOutboxEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProcessOutboxEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Optional tenant filter | 
 **batchSize** | **int32** | Maximum events to process | [default to 100]
 **maxAttempts** | **int32** | Maximum retry attempts | [default to 10]

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


## ProcessPendingEventsV1

> interface{} ProcessPendingEventsV1(ctx).Limit(limit).ShardAttemptOffset(shardAttemptOffset).TenantId(tenantId).PoolId(poolId).Execute()

Process Pending Events



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
	limit := int32(56) // int32 | Number of events to process (optional) (default to 10)
	shardAttemptOffset := int32(56) // int32 | Shard attempt offset (optional) (default to 0)
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter by tenant_id (for testing) (optional)
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter by pool_id (for testing) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.ProcessPendingEventsV1(context.Background()).Limit(limit).ShardAttemptOffset(shardAttemptOffset).TenantId(tenantId).PoolId(poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.ProcessPendingEventsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProcessPendingEventsV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.ProcessPendingEventsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProcessPendingEventsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of events to process | [default to 10]
 **shardAttemptOffset** | **int32** | Shard attempt offset | [default to 0]
 **tenantId** | **string** | Filter by tenant_id (for testing) | 
 **poolId** | **string** | Filter by pool_id (for testing) | 

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


## ProcessPendingProjectionsEndpointV1

> interface{} ProcessPendingProjectionsEndpointV1(ctx).TenantId(tenantId).PoolId(poolId).BatchSize(batchSize).Execute()

Process Pending Projections Endpoint



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional tenant filter (optional)
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional pool filter (optional)
	batchSize := int32(56) // int32 | Maximum number of wallets to process (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.ProcessPendingProjectionsEndpointV1(context.Background()).TenantId(tenantId).PoolId(poolId).BatchSize(batchSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.ProcessPendingProjectionsEndpointV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProcessPendingProjectionsEndpointV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.ProcessPendingProjectionsEndpointV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProcessPendingProjectionsEndpointV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Optional tenant filter | 
 **poolId** | **string** | Optional pool filter | 
 **batchSize** | **int32** | Maximum number of wallets to process | [default to 100]

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


## ProcessRollover

> RolloverResponse ProcessRollover(ctx).RolloverRequest(rolloverRequest).Execute()

Process Rollover



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
	rolloverRequest := *openapiclient.NewRolloverRequest("TenantId_example", "SubscriptionId_example") // RolloverRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.ProcessRollover(context.Background()).RolloverRequest(rolloverRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.ProcessRollover``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProcessRollover`: RolloverResponse
	fmt.Fprintf(os.Stdout, "Response from `V1API.ProcessRollover`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProcessRolloverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rolloverRequest** | [**RolloverRequest**](RolloverRequest.md) |  | 

### Return type

[**RolloverResponse**](RolloverResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProcessValueRecognitionV1Fx

> ValueRecognitionResponse ProcessValueRecognitionV1Fx(ctx).ProcessValueRecognitionRequest(processValueRecognitionRequest).Execute()

Process Value Recognition



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
	processValueRecognitionRequest := *openapiclient.NewProcessValueRecognitionRequest("TenantId_example", "PoolId_example", "JournalEntryId_example") // ProcessValueRecognitionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.ProcessValueRecognitionV1Fx(context.Background()).ProcessValueRecognitionRequest(processValueRecognitionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.ProcessValueRecognitionV1Fx``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProcessValueRecognitionV1Fx`: ValueRecognitionResponse
	fmt.Fprintf(os.Stdout, "Response from `V1API.ProcessValueRecognitionV1Fx`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProcessValueRecognitionV1FxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processValueRecognitionRequest** | [**ProcessValueRecognitionRequest**](ProcessValueRecognitionRequest.md) |  | 

### Return type

[**ValueRecognitionResponse**](ValueRecognitionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectWalletStateEndpointV1

> interface{} ProjectWalletStateEndpointV1(ctx, walletId).TenantId(tenantId).PoolId(poolId).AsOf(asOf).Execute()

Project Wallet State Endpoint



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Pool identifier
	asOf := time.Now() // time.Time | As-of timestamp (for time travel) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.ProjectWalletStateEndpointV1(context.Background(), walletId).TenantId(tenantId).PoolId(poolId).AsOf(asOf).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.ProjectWalletStateEndpointV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectWalletStateEndpointV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.ProjectWalletStateEndpointV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectWalletStateEndpointV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** | Tenant identifier | 
 **poolId** | **string** | Pool identifier | 
 **asOf** | **time.Time** | As-of timestamp (for time travel) | 

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


## RateEvent

> interface{} RateEvent(ctx, usageEventId).Execute()

Rate Event



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
	usageEventId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.RateEvent(context.Background(), usageEventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.RateEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RateEvent`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.RateEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**usageEventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRateEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## RecordUsage

> interface{} RecordUsage(ctx).RecordUsageRequest(recordUsageRequest).Execute()

Record Usage



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
	recordUsageRequest := *openapiclient.NewRecordUsageRequest("TenantId_example", "SubjectId_example", "PoolId_example", "FeatureKey_example", "UsageEventId_example", time.Now(), map[string]interface{}{"key": interface{}(123)}) // RecordUsageRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.RecordUsage(context.Background()).RecordUsageRequest(recordUsageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.RecordUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordUsage`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.RecordUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRecordUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recordUsageRequest** | [**RecordUsageRequest**](RecordUsageRequest.md) |  | 

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


## ResolveEvent

> interface{} ResolveEvent(ctx, usageEventId).MaxAttempts(maxAttempts).Execute()

Resolve Event



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
	usageEventId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	maxAttempts := int32(56) // int32 | Maximum retry attempts (optional) (default to 5)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.ResolveEvent(context.Background(), usageEventId).MaxAttempts(maxAttempts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.ResolveEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResolveEvent`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.ResolveEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**usageEventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResolveEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxAttempts** | **int32** | Maximum retry attempts | [default to 5]

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


## ResolveTenantAndPool

> interface{} ResolveTenantAndPool(ctx).SubjectId(subjectId).Execute()

Resolve Tenant And Pool



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
	subjectId := "subjectId_example" // string | Subject identifier (customer ID or subject key)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.ResolveTenantAndPool(context.Background()).SubjectId(subjectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.ResolveTenantAndPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResolveTenantAndPool`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.ResolveTenantAndPool`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResolveTenantAndPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subjectId** | **string** | Subject identifier (customer ID or subject key) | 

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


## RetryUnpriced

> RetryUnpricedResponse RetryUnpriced(ctx).RetryUnpricedRequest(retryUnpricedRequest).Execute()

Retry Unpriced



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
	retryUnpricedRequest := *openapiclient.NewRetryUnpricedRequest("TenantId_example", "PoolId_example", "UsageEventId_example") // RetryUnpricedRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.RetryUnpriced(context.Background()).RetryUnpricedRequest(retryUnpricedRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.RetryUnpriced``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetryUnpriced`: RetryUnpricedResponse
	fmt.Fprintf(os.Stdout, "Response from `V1API.RetryUnpriced`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetryUnpricedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **retryUnpricedRequest** | [**RetryUnpricedRequest**](RetryUnpricedRequest.md) |  | 

### Return type

[**RetryUnpricedResponse**](RetryUnpricedResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReviewQuarantined

> ReviewQuarantinedResponse ReviewQuarantined(ctx).ReviewQuarantinedRequest(reviewQuarantinedRequest).Execute()

Review Quarantined



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
	reviewQuarantinedRequest := *openapiclient.NewReviewQuarantinedRequest("TenantId_example", "PoolId_example", "UsageEventId_example", "Action_example") // ReviewQuarantinedRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.ReviewQuarantined(context.Background()).ReviewQuarantinedRequest(reviewQuarantinedRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.ReviewQuarantined``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReviewQuarantined`: ReviewQuarantinedResponse
	fmt.Fprintf(os.Stdout, "Response from `V1API.ReviewQuarantined`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReviewQuarantinedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reviewQuarantinedRequest** | [**ReviewQuarantinedRequest**](ReviewQuarantinedRequest.md) |  | 

### Return type

[**ReviewQuarantinedResponse**](ReviewQuarantinedResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncSubscription

> interface{} SyncSubscription(ctx).TenantId(tenantId).SubscriptionSyncRequest(subscriptionSyncRequest).Execute()

Sync Subscription



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	subscriptionSyncRequest := *openapiclient.NewSubscriptionSyncRequest("SubscriptionId_example", "BillingAnchor_example", "PlanId_example", "ActiveFrom_example") // SubscriptionSyncRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.SyncSubscription(context.Background()).TenantId(tenantId).SubscriptionSyncRequest(subscriptionSyncRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.SyncSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncSubscription`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.SyncSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **subscriptionSyncRequest** | [**SubscriptionSyncRequest**](SubscriptionSyncRequest.md) |  | 

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


## TriggerSubscriptionSync

> interface{} TriggerSubscriptionSync(ctx, subscriptionId).Execute()

Trigger Subscription Sync



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
	subscriptionId := "subscriptionId_example" // string | Subscription ID to sync

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.TriggerSubscriptionSync(context.Background(), subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.TriggerSubscriptionSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TriggerSubscriptionSync`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.TriggerSubscriptionSync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription ID to sync | 

### Other Parameters

Other parameters are passed through a pointer to a apiTriggerSubscriptionSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## UpdateRuleV1

> AutoTopupRuleResponse UpdateRuleV1(ctx, ruleId).TenantId(tenantId).PoolId(poolId).WalletId(walletId).UpdateAutoTopupRuleRequest(updateAutoTopupRuleRequest).Execute()

Update Rule



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
	ruleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Auto top-up rule ID
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional tenant identifier for validation (optional)
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional pool identifier for validation (optional)
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional wallet identifier for validation (optional)
	updateAutoTopupRuleRequest := *openapiclient.NewUpdateAutoTopupRuleRequest() // UpdateAutoTopupRuleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UpdateRuleV1(context.Background(), ruleId).TenantId(tenantId).PoolId(poolId).WalletId(walletId).UpdateAutoTopupRuleRequest(updateAutoTopupRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UpdateRuleV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRuleV1`: AutoTopupRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `V1API.UpdateRuleV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** | Auto top-up rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRuleV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** | Optional tenant identifier for validation | 
 **poolId** | **string** | Optional pool identifier for validation | 
 **walletId** | **string** | Optional wallet identifier for validation | 
 **updateAutoTopupRuleRequest** | [**UpdateAutoTopupRuleRequest**](UpdateAutoTopupRuleRequest.md) |  | 

### Return type

[**AutoTopupRuleResponse**](AutoTopupRuleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateThresholdEndpoint

> interface{} UpdateThresholdEndpoint(ctx, thresholdId).WalletId(walletId).TenantId(tenantId).PoolId(poolId).UpdateThresholdRequest(updateThresholdRequest).Execute()

Update Threshold Endpoint



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
	thresholdId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Wallet identifier
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant identifier
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Pool identifier
	updateThresholdRequest := *openapiclient.NewUpdateThresholdRequest() // UpdateThresholdRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UpdateThresholdEndpoint(context.Background(), thresholdId).WalletId(walletId).TenantId(tenantId).PoolId(poolId).UpdateThresholdRequest(updateThresholdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UpdateThresholdEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateThresholdEndpoint`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.UpdateThresholdEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thresholdId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateThresholdEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **walletId** | **string** | Wallet identifier | 
 **tenantId** | **string** | Tenant identifier | 
 **poolId** | **string** | Pool identifier | 
 **updateThresholdRequest** | [**UpdateThresholdRequest**](UpdateThresholdRequest.md) |  | 

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


## UpdateWalletSettings

> WalletResponse UpdateWalletSettings(ctx, walletId).UpdateWalletSettingsRequest(updateWalletSettingsRequest).TenantId(tenantId).PoolId(poolId).Execute()

Update Wallet Settings



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	updateWalletSettingsRequest := *openapiclient.NewUpdateWalletSettingsRequest() // UpdateWalletSettingsRequest | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional tenant identifier for validation (optional)
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional pool identifier for validation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UpdateWalletSettings(context.Background(), walletId).UpdateWalletSettingsRequest(updateWalletSettingsRequest).TenantId(tenantId).PoolId(poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UpdateWalletSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWalletSettings`: WalletResponse
	fmt.Fprintf(os.Stdout, "Response from `V1API.UpdateWalletSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWalletSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateWalletSettingsRequest** | [**UpdateWalletSettingsRequest**](UpdateWalletSettingsRequest.md) |  | 
 **tenantId** | **string** | Optional tenant identifier for validation | 
 **poolId** | **string** | Optional pool identifier for validation | 

### Return type

[**WalletResponse**](WalletResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWalletThresholds

> WalletResponse UpdateWalletThresholds(ctx, walletId).UpdateWalletThresholdsRequest(updateWalletThresholdsRequest).TenantId(tenantId).PoolId(poolId).Execute()

Update Wallet Thresholds



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	updateWalletThresholdsRequest := *openapiclient.NewUpdateWalletThresholdsRequest() // UpdateWalletThresholdsRequest | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional tenant identifier for validation (optional)
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional pool identifier for validation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.UpdateWalletThresholds(context.Background(), walletId).UpdateWalletThresholdsRequest(updateWalletThresholdsRequest).TenantId(tenantId).PoolId(poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UpdateWalletThresholds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWalletThresholds`: WalletResponse
	fmt.Fprintf(os.Stdout, "Response from `V1API.UpdateWalletThresholds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWalletThresholdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateWalletThresholdsRequest** | [**UpdateWalletThresholdsRequest**](UpdateWalletThresholdsRequest.md) |  | 
 **tenantId** | **string** | Optional tenant identifier for validation | 
 **poolId** | **string** | Optional pool identifier for validation | 

### Return type

[**WalletResponse**](WalletResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateSnapshotEndpoint

> ValidationResponse ValidateSnapshotEndpoint(ctx, snapshotId).Execute()

Validate Snapshot Endpoint



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
	snapshotId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Snapshot identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.ValidateSnapshotEndpoint(context.Background(), snapshotId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.ValidateSnapshotEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateSnapshotEndpoint`: ValidationResponse
	fmt.Fprintf(os.Stdout, "Response from `V1API.ValidateSnapshotEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string** | Snapshot identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateSnapshotEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ValidationResponse**](ValidationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VoidGrant

> GrantResponse VoidGrant(ctx, grantId).VoidGrantRequest(voidGrantRequest).Execute()

Void Grant



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
	grantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	voidGrantRequest := *openapiclient.NewVoidGrantRequest() // VoidGrantRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.VoidGrant(context.Background(), grantId).VoidGrantRequest(voidGrantRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.VoidGrant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VoidGrant`: GrantResponse
	fmt.Fprintf(os.Stdout, "Response from `V1API.VoidGrant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**grantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVoidGrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **voidGrantRequest** | [**VoidGrantRequest**](VoidGrantRequest.md) |  | 

### Return type

[**GrantResponse**](GrantResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WarmRateCardCacheV1

> interface{} WarmRateCardCacheV1(ctx).TenantId(tenantId).Limit(limit).Execute()

Warm Rate Card Cache



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional tenant filter (optional)
	limit := int32(56) // int32 | Maximum number of rate cards to warm (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.WarmRateCardCacheV1(context.Background()).TenantId(tenantId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.WarmRateCardCacheV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WarmRateCardCacheV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.WarmRateCardCacheV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWarmRateCardCacheV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Optional tenant filter | 
 **limit** | **int32** | Maximum number of rate cards to warm | [default to 50]

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


## WarmWalletCache

> interface{} WarmWalletCache(ctx).TenantId(tenantId).PoolId(poolId).Limit(limit).Execute()

Warm Wallet Cache



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional tenant filter (optional)
	poolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional pool filter (optional)
	limit := int32(56) // int32 | Maximum number of wallets to warm (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.WarmWalletCache(context.Background()).TenantId(tenantId).PoolId(poolId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.WarmWalletCache``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WarmWalletCache`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.WarmWalletCache`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWarmWalletCacheRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Optional tenant filter | 
 **poolId** | **string** | Optional pool filter | 
 **limit** | **int32** | Maximum number of wallets to warm | [default to 100]

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

