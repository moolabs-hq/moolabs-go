# \QuotesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApproverV1**](QuotesAPI.md#DeleteApproverV1) | **Delete** /v1/quotes/approver-directory/{user_ref} | Delete Approver
[**DeleteTerritory**](QuotesAPI.md#DeleteTerritory) | **Delete** /v1/quotes/territories/{key} | Delete Territory
[**GetApprovalLevelsV1**](QuotesAPI.md#GetApprovalLevelsV1) | **Get** /v1/quotes/approval-levels | Get Approval Levels
[**GetApproverDirectoryV1**](QuotesAPI.md#GetApproverDirectoryV1) | **Get** /v1/quotes/approver-directory | Get Approver Directory
[**GetLatestQuoteContractUploadV1**](QuotesAPI.md#GetLatestQuoteContractUploadV1) | **Get** /v1/quotes/{quote_id}/contract-uploads/latest | Get Latest Quote Contract Upload
[**GetPricingSnapshotV1**](QuotesAPI.md#GetPricingSnapshotV1) | **Get** /v1/quotes/pricing-snapshots/{snapshot_id} | Get Pricing Snapshot
[**GetQuote**](QuotesAPI.md#GetQuote) | **Get** /v1/quotes/{quote_id} | Get Quote
[**GetQuoteContractPdfV1**](QuotesAPI.md#GetQuoteContractPdfV1) | **Get** /v1/quotes/{quote_id}/contract.pdf | Get Quote Contract Pdf
[**GetQuoteSettings**](QuotesAPI.md#GetQuoteSettings) | **Get** /v1/quotes/settings | Get Quote Settings
[**GetQuoteVersion**](QuotesAPI.md#GetQuoteVersion) | **Get** /v1/quotes/{quote_id}/versions/{version} | Get Quote Version
[**GetRateCardsV1**](QuotesAPI.md#GetRateCardsV1) | **Get** /v1/rate-cards | Get Rate Cards
[**GetTerritories**](QuotesAPI.md#GetTerritories) | **Get** /v1/quotes/territories | Get Territories
[**ListQuoteActivationFailuresV1**](QuotesAPI.md#ListQuoteActivationFailuresV1) | **Get** /v1/quotes/activation-failures | List Quote Activation Failures
[**ListQuoteAgentCardsV1**](QuotesAPI.md#ListQuoteAgentCardsV1) | **Get** /v1/quotes/{quote_id}/agent-cards | List Quote Agent Cards
[**ListQuoteAgentRunsV1**](QuotesAPI.md#ListQuoteAgentRunsV1) | **Get** /v1/quotes/{quote_id}/agent-runs | List Quote Agent Runs
[**ListQuoteRedlines**](QuotesAPI.md#ListQuoteRedlines) | **Get** /v1/quotes/{quote_id}/redlines | List Quote Redlines
[**ListQuotes**](QuotesAPI.md#ListQuotes) | **Get** /v1/quotes | List Quotes
[**PatchRateCardCostV1**](QuotesAPI.md#PatchRateCardCostV1) | **Patch** /v1/rate-cards/{rate_card_id}/cost | Patch Rate Card Cost
[**PostQuote**](QuotesAPI.md#PostQuote) | **Post** /v1/quotes | Post Quote
[**PostQuoteAccept**](QuotesAPI.md#PostQuoteAccept) | **Post** /v1/quotes/{quote_id}/accept | Post Quote Accept
[**PostQuoteActivate**](QuotesAPI.md#PostQuoteActivate) | **Post** /v1/quotes/{quote_id}/activate | Post Quote Activate
[**PostQuoteApprove**](QuotesAPI.md#PostQuoteApprove) | **Post** /v1/quotes/{quote_id}/approve | Post Quote Approve
[**PostQuoteContractUploadV1**](QuotesAPI.md#PostQuoteContractUploadV1) | **Post** /v1/quotes/{quote_id}/contract-uploads | Post Quote Contract Upload
[**PostQuoteLock**](QuotesAPI.md#PostQuoteLock) | **Post** /v1/quotes/{quote_id}/lock | Post Quote Lock
[**PostQuoteRedline**](QuotesAPI.md#PostQuoteRedline) | **Post** /v1/quotes/{quote_id}/redlines | Post Quote Redline
[**PostQuoteRedlineAccept**](QuotesAPI.md#PostQuoteRedlineAccept) | **Post** /v1/quotes/{quote_id}/redlines/{redline_id}/accept | Post Quote Redline Accept
[**PostQuoteRedlineReject**](QuotesAPI.md#PostQuoteRedlineReject) | **Post** /v1/quotes/{quote_id}/redlines/{redline_id}/reject | Post Quote Redline Reject
[**PostQuoteRedlinesFirstPassV1**](QuotesAPI.md#PostQuoteRedlinesFirstPassV1) | **Post** /v1/quotes/{quote_id}/redlines/first-pass | Post Quote Redlines First Pass
[**PostQuoteReject**](QuotesAPI.md#PostQuoteReject) | **Post** /v1/quotes/{quote_id}/reject | Post Quote Reject
[**PostQuoteSend**](QuotesAPI.md#PostQuoteSend) | **Post** /v1/quotes/{quote_id}/send | Post Quote Send
[**PostQuoteSession**](QuotesAPI.md#PostQuoteSession) | **Post** /v1/quotes/{quote_id}/sessions | Post Quote Session
[**PostQuoteSigningWebhook**](QuotesAPI.md#PostQuoteSigningWebhook) | **Post** /v1/quotes/signing/webhooks/{provider} | Post Quote Signing Webhook
[**PostQuoteSubmitApprovalV1**](QuotesAPI.md#PostQuoteSubmitApprovalV1) | **Post** /v1/quotes/{quote_id}/submit-approval | Post Quote Submit Approval
[**PostQuoteUnlock**](QuotesAPI.md#PostQuoteUnlock) | **Post** /v1/quotes/{quote_id}/unlock | Post Quote Unlock
[**PutApprovalLevelsV1**](QuotesAPI.md#PutApprovalLevelsV1) | **Put** /v1/quotes/approval-levels | Put Approval Levels
[**PutApproverV1**](QuotesAPI.md#PutApproverV1) | **Put** /v1/quotes/approver-directory/{user_ref} | Put Approver
[**PutQuoteSettings**](QuotesAPI.md#PutQuoteSettings) | **Put** /v1/quotes/settings | Put Quote Settings
[**PutTerritory**](QuotesAPI.md#PutTerritory) | **Put** /v1/quotes/territories/{key} | Put Territory
[**SyncApproverDirectoryV1**](QuotesAPI.md#SyncApproverDirectoryV1) | **Post** /v1/quotes/approver-directory/sync | Sync Approver Directory



## DeleteApproverV1

> interface{} DeleteApproverV1(ctx, userRef).Execute()

Delete Approver

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
	userRef := "userRef_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.DeleteApproverV1(context.Background(), userRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.DeleteApproverV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteApproverV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.DeleteApproverV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userRef** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApproverV1Request struct via the builder pattern


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


## DeleteTerritory

> interface{} DeleteTerritory(ctx, key).Execute()

Delete Territory

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
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.DeleteTerritory(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.DeleteTerritory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTerritory`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.DeleteTerritory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTerritoryRequest struct via the builder pattern


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


## GetApprovalLevelsV1

> interface{} GetApprovalLevelsV1(ctx).Execute()

Get Approval Levels

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
	resp, r, err := apiClient.QuotesAPI.GetApprovalLevelsV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.GetApprovalLevelsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApprovalLevelsV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.GetApprovalLevelsV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApprovalLevelsV1Request struct via the builder pattern


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


## GetApproverDirectoryV1

> interface{} GetApproverDirectoryV1(ctx).LevelKey(levelKey).TerritoryKey(territoryKey).Execute()

Get Approver Directory

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
	levelKey := "levelKey_example" // string |  (optional)
	territoryKey := "territoryKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.GetApproverDirectoryV1(context.Background()).LevelKey(levelKey).TerritoryKey(territoryKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.GetApproverDirectoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApproverDirectoryV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.GetApproverDirectoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApproverDirectoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **levelKey** | **string** |  | 
 **territoryKey** | **string** |  | 

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


## GetLatestQuoteContractUploadV1

> ContractUploadDetailResponse GetLatestQuoteContractUploadV1(ctx, quoteId).Execute()

Get Latest Quote Contract Upload



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
	quoteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.GetLatestQuoteContractUploadV1(context.Background(), quoteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.GetLatestQuoteContractUploadV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLatestQuoteContractUploadV1`: ContractUploadDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.GetLatestQuoteContractUploadV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestQuoteContractUploadV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContractUploadDetailResponse**](ContractUploadDetailResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPricingSnapshotV1

> PricingSnapshotResponse GetPricingSnapshotV1(ctx, snapshotId).Execute()

Get Pricing Snapshot

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
	snapshotId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.GetPricingSnapshotV1(context.Background(), snapshotId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.GetPricingSnapshotV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPricingSnapshotV1`: PricingSnapshotResponse
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.GetPricingSnapshotV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPricingSnapshotV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PricingSnapshotResponse**](PricingSnapshotResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuote

> QuoteResponse GetQuote(ctx, quoteId).Execute()

Get Quote

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
	quoteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.GetQuote(context.Background(), quoteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.GetQuote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuote`: QuoteResponse
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.GetQuote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQuoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QuoteResponse**](QuoteResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuoteContractPdfV1

> interface{} GetQuoteContractPdfV1(ctx, quoteId).Execute()

Get Quote Contract Pdf

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
	quoteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.GetQuoteContractPdfV1(context.Background(), quoteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.GetQuoteContractPdfV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuoteContractPdfV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.GetQuoteContractPdfV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQuoteContractPdfV1Request struct via the builder pattern


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


## GetQuoteSettings

> QuoteSettingsResponse GetQuoteSettings(ctx).Execute()

Get Quote Settings

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
	resp, r, err := apiClient.QuotesAPI.GetQuoteSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.GetQuoteSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuoteSettings`: QuoteSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.GetQuoteSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetQuoteSettingsRequest struct via the builder pattern


### Return type

[**QuoteSettingsResponse**](QuoteSettingsResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuoteVersion

> QuoteVersionResponse GetQuoteVersion(ctx, quoteId, version).Execute()

Get Quote Version

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
	quoteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	version := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.GetQuoteVersion(context.Background(), quoteId, version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.GetQuoteVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuoteVersion`: QuoteVersionResponse
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.GetQuoteVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 
**version** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQuoteVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**QuoteVersionResponse**](QuoteVersionResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRateCardsV1

> RateCardCostListResponse GetRateCardsV1(ctx).Execute()

Get Rate Cards



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
	resp, r, err := apiClient.QuotesAPI.GetRateCardsV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.GetRateCardsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRateCardsV1`: RateCardCostListResponse
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.GetRateCardsV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRateCardsV1Request struct via the builder pattern


### Return type

[**RateCardCostListResponse**](RateCardCostListResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTerritories

> interface{} GetTerritories(ctx).Execute()

Get Territories

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
	resp, r, err := apiClient.QuotesAPI.GetTerritories(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.GetTerritories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTerritories`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.GetTerritories`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTerritoriesRequest struct via the builder pattern


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


## ListQuoteActivationFailuresV1

> interface{} ListQuoteActivationFailuresV1(ctx).Limit(limit).Offset(offset).Execute()

List Quote Activation Failures

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
	limit := int32(56) // int32 |  (optional) (default to 100)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.ListQuoteActivationFailuresV1(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.ListQuoteActivationFailuresV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListQuoteActivationFailuresV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.ListQuoteActivationFailuresV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListQuoteActivationFailuresV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 100]
 **offset** | **int32** |  | [default to 0]

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


## ListQuoteAgentCardsV1

> interface{} ListQuoteAgentCardsV1(ctx, quoteId).QuoteVersion(quoteVersion).Limit(limit).Offset(offset).Execute()

List Quote Agent Cards

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
	quoteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	quoteVersion := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 20)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.ListQuoteAgentCardsV1(context.Background(), quoteId).QuoteVersion(quoteVersion).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.ListQuoteAgentCardsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListQuoteAgentCardsV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.ListQuoteAgentCardsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListQuoteAgentCardsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **quoteVersion** | **int32** |  | 
 **limit** | **int32** |  | [default to 20]
 **offset** | **int32** |  | [default to 0]

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


## ListQuoteAgentRunsV1

> []QuoteAgentProvenanceRunResponse ListQuoteAgentRunsV1(ctx, quoteId).QuoteVersion(quoteVersion).Limit(limit).Offset(offset).Execute()

List Quote Agent Runs

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
	quoteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	quoteVersion := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.ListQuoteAgentRunsV1(context.Background(), quoteId).QuoteVersion(quoteVersion).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.ListQuoteAgentRunsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListQuoteAgentRunsV1`: []QuoteAgentProvenanceRunResponse
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.ListQuoteAgentRunsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListQuoteAgentRunsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **quoteVersion** | **int32** |  | 
 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** |  | [default to 0]

### Return type

[**[]QuoteAgentProvenanceRunResponse**](QuoteAgentProvenanceRunResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListQuoteRedlines

> []QuoteRedlineResponse ListQuoteRedlines(ctx, quoteId).Status(status).Limit(limit).Offset(offset).Execute()

List Quote Redlines

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
	quoteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	status := "status_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.ListQuoteRedlines(context.Background(), quoteId).Status(status).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.ListQuoteRedlines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListQuoteRedlines`: []QuoteRedlineResponse
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.ListQuoteRedlines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListQuoteRedlinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** |  | 
 **limit** | **int32** |  | [default to 100]
 **offset** | **int32** |  | [default to 0]

### Return type

[**[]QuoteRedlineResponse**](QuoteRedlineResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListQuotes

> QuoteListResponse ListQuotes(ctx).Limit(limit).Offset(offset).State(state).Execute()

List Quotes

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
	limit := int32(56) // int32 |  (optional) (default to 100)
	offset := int32(56) // int32 |  (optional) (default to 0)
	state := "state_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.ListQuotes(context.Background()).Limit(limit).Offset(offset).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.ListQuotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListQuotes`: QuoteListResponse
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.ListQuotes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListQuotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 100]
 **offset** | **int32** |  | [default to 0]
 **state** | **string** |  | 

### Return type

[**QuoteListResponse**](QuoteListResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchRateCardCostV1

> RateCardCostRow PatchRateCardCostV1(ctx, rateCardId).SetRateCardCostRequest(setRateCardCostRequest).Execute()

Patch Rate Card Cost



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
	setRateCardCostRequest := *openapiclient.NewSetRateCardCostRequest() // SetRateCardCostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.PatchRateCardCostV1(context.Background(), rateCardId).SetRateCardCostRequest(setRateCardCostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.PatchRateCardCostV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchRateCardCostV1`: RateCardCostRow
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.PatchRateCardCostV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rateCardId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchRateCardCostV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setRateCardCostRequest** | [**SetRateCardCostRequest**](SetRateCardCostRequest.md) |  | 

### Return type

[**RateCardCostRow**](RateCardCostRow.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostQuote

> QuoteResponse PostQuote(ctx).CreateQuoteRequest(createQuoteRequest).IdempotencyKey(idempotencyKey).Execute()

Post Quote

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
	createQuoteRequest := *openapiclient.NewCreateQuoteRequest(map[string]interface{}{"key": interface{}(123)}) // CreateQuoteRequest | 
	idempotencyKey := "idempotencyKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.PostQuote(context.Background()).CreateQuoteRequest(createQuoteRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.PostQuote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostQuote`: QuoteResponse
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.PostQuote`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostQuoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createQuoteRequest** | [**CreateQuoteRequest**](CreateQuoteRequest.md) |  | 
 **idempotencyKey** | **string** |  | 

### Return type

[**QuoteResponse**](QuoteResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostQuoteAccept

> interface{} PostQuoteAccept(ctx, quoteId).IdempotencyKey(idempotencyKey).SignedDocument(signedDocument).BuyerSignerRef(buyerSignerRef).VerifiedByRef(verifiedByRef).SellerSignerRef(sellerSignerRef).Execute()

Post Quote Accept

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
	quoteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	idempotencyKey := "idempotencyKey_example" // string |  (optional)
	signedDocument := os.NewFile(1234, "some_file") // *os.File |  (optional)
	buyerSignerRef := "buyerSignerRef_example" // string |  (optional)
	verifiedByRef := "verifiedByRef_example" // string |  (optional)
	sellerSignerRef := "sellerSignerRef_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.PostQuoteAccept(context.Background(), quoteId).IdempotencyKey(idempotencyKey).SignedDocument(signedDocument).BuyerSignerRef(buyerSignerRef).VerifiedByRef(verifiedByRef).SellerSignerRef(sellerSignerRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.PostQuoteAccept``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostQuoteAccept`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.PostQuoteAccept`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostQuoteAcceptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** |  | 
 **signedDocument** | ***os.File** |  | 
 **buyerSignerRef** | **string** |  | 
 **verifiedByRef** | **string** |  | 
 **sellerSignerRef** | **string** |  | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostQuoteActivate

> interface{} PostQuoteActivate(ctx, quoteId).IdempotencyKey(idempotencyKey).Execute()

Post Quote Activate

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
	quoteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	idempotencyKey := "idempotencyKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.PostQuoteActivate(context.Background(), quoteId).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.PostQuoteActivate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostQuoteActivate`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.PostQuoteActivate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostQuoteActivateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** |  | 

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


## PostQuoteApprove

> QuoteApprovalResponse PostQuoteApprove(ctx, quoteId).IdempotencyKey(idempotencyKey).QuoteApprovalDecisionRequest(quoteApprovalDecisionRequest).Execute()

Post Quote Approve

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
	quoteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	idempotencyKey := "idempotencyKey_example" // string |  (optional)
	quoteApprovalDecisionRequest := *openapiclient.NewQuoteApprovalDecisionRequest() // QuoteApprovalDecisionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.PostQuoteApprove(context.Background(), quoteId).IdempotencyKey(idempotencyKey).QuoteApprovalDecisionRequest(quoteApprovalDecisionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.PostQuoteApprove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostQuoteApprove`: QuoteApprovalResponse
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.PostQuoteApprove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostQuoteApproveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** |  | 
 **quoteApprovalDecisionRequest** | [**QuoteApprovalDecisionRequest**](QuoteApprovalDecisionRequest.md) |  | 

### Return type

[**QuoteApprovalResponse**](QuoteApprovalResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostQuoteContractUploadV1

> interface{} PostQuoteContractUploadV1(ctx, quoteId).File(file).Execute()

Post Quote Contract Upload



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
	quoteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	file := os.NewFile(1234, "some_file") // *os.File | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.PostQuoteContractUploadV1(context.Background(), quoteId).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.PostQuoteContractUploadV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostQuoteContractUploadV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.PostQuoteContractUploadV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostQuoteContractUploadV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** |  | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostQuoteLock

> QuoteVersionResponse PostQuoteLock(ctx, quoteId).IdempotencyKey(idempotencyKey).Execute()

Post Quote Lock

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
	quoteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	idempotencyKey := "idempotencyKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.PostQuoteLock(context.Background(), quoteId).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.PostQuoteLock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostQuoteLock`: QuoteVersionResponse
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.PostQuoteLock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostQuoteLockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** |  | 

### Return type

[**QuoteVersionResponse**](QuoteVersionResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostQuoteRedline

> QuoteRedlineResponse PostQuoteRedline(ctx, quoteId).CreateQuoteRedlineRequest(createQuoteRedlineRequest).IdempotencyKey(idempotencyKey).Execute()

Post Quote Redline

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
	quoteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	createQuoteRedlineRequest := *openapiclient.NewCreateQuoteRedlineRequest(int32(123), "Source_example", "RedlineText_example") // CreateQuoteRedlineRequest | 
	idempotencyKey := "idempotencyKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.PostQuoteRedline(context.Background(), quoteId).CreateQuoteRedlineRequest(createQuoteRedlineRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.PostQuoteRedline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostQuoteRedline`: QuoteRedlineResponse
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.PostQuoteRedline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostQuoteRedlineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createQuoteRedlineRequest** | [**CreateQuoteRedlineRequest**](CreateQuoteRedlineRequest.md) |  | 
 **idempotencyKey** | **string** |  | 

### Return type

[**QuoteRedlineResponse**](QuoteRedlineResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostQuoteRedlineAccept

> QuoteVersionResponse PostQuoteRedlineAccept(ctx, quoteId, redlineId).IdempotencyKey(idempotencyKey).Execute()

Post Quote Redline Accept

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
	quoteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	redlineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	idempotencyKey := "idempotencyKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.PostQuoteRedlineAccept(context.Background(), quoteId, redlineId).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.PostQuoteRedlineAccept``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostQuoteRedlineAccept`: QuoteVersionResponse
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.PostQuoteRedlineAccept`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 
**redlineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostQuoteRedlineAcceptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **idempotencyKey** | **string** |  | 

### Return type

[**QuoteVersionResponse**](QuoteVersionResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostQuoteRedlineReject

> QuoteRedlineResponse PostQuoteRedlineReject(ctx, quoteId, redlineId).IdempotencyKey(idempotencyKey).QuoteRedlineDecisionRequest(quoteRedlineDecisionRequest).Execute()

Post Quote Redline Reject

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
	quoteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	redlineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	idempotencyKey := "idempotencyKey_example" // string |  (optional)
	quoteRedlineDecisionRequest := *openapiclient.NewQuoteRedlineDecisionRequest() // QuoteRedlineDecisionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.PostQuoteRedlineReject(context.Background(), quoteId, redlineId).IdempotencyKey(idempotencyKey).QuoteRedlineDecisionRequest(quoteRedlineDecisionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.PostQuoteRedlineReject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostQuoteRedlineReject`: QuoteRedlineResponse
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.PostQuoteRedlineReject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 
**redlineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostQuoteRedlineRejectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **idempotencyKey** | **string** |  | 
 **quoteRedlineDecisionRequest** | [**QuoteRedlineDecisionRequest**](QuoteRedlineDecisionRequest.md) |  | 

### Return type

[**QuoteRedlineResponse**](QuoteRedlineResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostQuoteRedlinesFirstPassV1

> FirstPassRedlineResponse PostQuoteRedlinesFirstPassV1(ctx, quoteId).FirstPassRedlineRequest(firstPassRedlineRequest).Execute()

Post Quote Redlines First Pass



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
	quoteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	firstPassRedlineRequest := *openapiclient.NewFirstPassRedlineRequest() // FirstPassRedlineRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.PostQuoteRedlinesFirstPassV1(context.Background(), quoteId).FirstPassRedlineRequest(firstPassRedlineRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.PostQuoteRedlinesFirstPassV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostQuoteRedlinesFirstPassV1`: FirstPassRedlineResponse
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.PostQuoteRedlinesFirstPassV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostQuoteRedlinesFirstPassV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **firstPassRedlineRequest** | [**FirstPassRedlineRequest**](FirstPassRedlineRequest.md) |  | 

### Return type

[**FirstPassRedlineResponse**](FirstPassRedlineResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostQuoteReject

> QuoteApprovalResponse PostQuoteReject(ctx, quoteId).IdempotencyKey(idempotencyKey).QuoteApprovalDecisionRequest(quoteApprovalDecisionRequest).Execute()

Post Quote Reject

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
	quoteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	idempotencyKey := "idempotencyKey_example" // string |  (optional)
	quoteApprovalDecisionRequest := *openapiclient.NewQuoteApprovalDecisionRequest() // QuoteApprovalDecisionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.PostQuoteReject(context.Background(), quoteId).IdempotencyKey(idempotencyKey).QuoteApprovalDecisionRequest(quoteApprovalDecisionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.PostQuoteReject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostQuoteReject`: QuoteApprovalResponse
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.PostQuoteReject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostQuoteRejectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** |  | 
 **quoteApprovalDecisionRequest** | [**QuoteApprovalDecisionRequest**](QuoteApprovalDecisionRequest.md) |  | 

### Return type

[**QuoteApprovalResponse**](QuoteApprovalResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostQuoteSend

> QuoteResponse PostQuoteSend(ctx, quoteId).IdempotencyKey(idempotencyKey).Execute()

Post Quote Send

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
	quoteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	idempotencyKey := "idempotencyKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.PostQuoteSend(context.Background(), quoteId).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.PostQuoteSend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostQuoteSend`: QuoteResponse
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.PostQuoteSend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostQuoteSendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** |  | 

### Return type

[**QuoteResponse**](QuoteResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostQuoteSession

> CreateSessionResponse PostQuoteSession(ctx, quoteId).IdempotencyKey(idempotencyKey).Execute()

Post Quote Session

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
	quoteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	idempotencyKey := "idempotencyKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.PostQuoteSession(context.Background(), quoteId).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.PostQuoteSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostQuoteSession`: CreateSessionResponse
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.PostQuoteSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostQuoteSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** |  | 

### Return type

[**CreateSessionResponse**](CreateSessionResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostQuoteSigningWebhook

> interface{} PostQuoteSigningWebhook(ctx, provider).XQuoteSignature(xQuoteSignature).Execute()

Post Quote Signing Webhook

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
	xQuoteSignature := "xQuoteSignature_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.PostQuoteSigningWebhook(context.Background(), provider).XQuoteSignature(xQuoteSignature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.PostQuoteSigningWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostQuoteSigningWebhook`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.PostQuoteSigningWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostQuoteSigningWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xQuoteSignature** | **string** |  | 

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


## PostQuoteSubmitApprovalV1

> QuoteApprovalResponse PostQuoteSubmitApprovalV1(ctx, quoteId).IdempotencyKey(idempotencyKey).Execute()

Post Quote Submit Approval

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
	quoteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	idempotencyKey := "idempotencyKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.PostQuoteSubmitApprovalV1(context.Background(), quoteId).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.PostQuoteSubmitApprovalV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostQuoteSubmitApprovalV1`: QuoteApprovalResponse
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.PostQuoteSubmitApprovalV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostQuoteSubmitApprovalV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** |  | 

### Return type

[**QuoteApprovalResponse**](QuoteApprovalResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostQuoteUnlock

> QuoteVersionResponse PostQuoteUnlock(ctx, quoteId).IdempotencyKey(idempotencyKey).Execute()

Post Quote Unlock

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
	quoteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	idempotencyKey := "idempotencyKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.PostQuoteUnlock(context.Background(), quoteId).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.PostQuoteUnlock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostQuoteUnlock`: QuoteVersionResponse
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.PostQuoteUnlock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostQuoteUnlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** |  | 

### Return type

[**QuoteVersionResponse**](QuoteVersionResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutApprovalLevelsV1

> interface{} PutApprovalLevelsV1(ctx).ReplaceLevelsRequest(replaceLevelsRequest).Execute()

Put Approval Levels

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
	replaceLevelsRequest := *openapiclient.NewReplaceLevelsRequest([]openapiclient.LevelIn{*openapiclient.NewLevelIn("Key_example", "Name_example", int32(123))}) // ReplaceLevelsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.PutApprovalLevelsV1(context.Background()).ReplaceLevelsRequest(replaceLevelsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.PutApprovalLevelsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutApprovalLevelsV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.PutApprovalLevelsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutApprovalLevelsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **replaceLevelsRequest** | [**ReplaceLevelsRequest**](ReplaceLevelsRequest.md) |  | 

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


## PutApproverV1

> interface{} PutApproverV1(ctx, userRef).ApproverIn(approverIn).Execute()

Put Approver

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
	userRef := "userRef_example" // string | 
	approverIn := *openapiclient.NewApproverIn("LevelKey_example") // ApproverIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.PutApproverV1(context.Background(), userRef).ApproverIn(approverIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.PutApproverV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutApproverV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.PutApproverV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userRef** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutApproverV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approverIn** | [**ApproverIn**](ApproverIn.md) |  | 

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


## PutQuoteSettings

> QuoteSettingsResponse PutQuoteSettings(ctx).UpdateQuoteSettingsRequest(updateQuoteSettingsRequest).Execute()

Put Quote Settings

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
	updateQuoteSettingsRequest := *openapiclient.NewUpdateQuoteSettingsRequest("BookingTrigger_example") // UpdateQuoteSettingsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.PutQuoteSettings(context.Background()).UpdateQuoteSettingsRequest(updateQuoteSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.PutQuoteSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutQuoteSettings`: QuoteSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.PutQuoteSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutQuoteSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateQuoteSettingsRequest** | [**UpdateQuoteSettingsRequest**](UpdateQuoteSettingsRequest.md) |  | 

### Return type

[**QuoteSettingsResponse**](QuoteSettingsResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutTerritory

> interface{} PutTerritory(ctx, key).TerritoryIn(territoryIn).Execute()

Put Territory

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
	key := "key_example" // string | 
	territoryIn := *openapiclient.NewTerritoryIn("Name_example") // TerritoryIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.PutTerritory(context.Background(), key).TerritoryIn(territoryIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.PutTerritory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutTerritory`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.PutTerritory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutTerritoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **territoryIn** | [**TerritoryIn**](TerritoryIn.md) |  | 

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


## SyncApproverDirectoryV1

> interface{} SyncApproverDirectoryV1(ctx).Execute()

Sync Approver Directory

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
	resp, r, err := apiClient.QuotesAPI.SyncApproverDirectoryV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.SyncApproverDirectoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncApproverDirectoryV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.SyncApproverDirectoryV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSyncApproverDirectoryV1Request struct via the builder pattern


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

