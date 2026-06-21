# \ClausePacksAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DismissRefinementEndpointV1Clause**](ClausePacksAPI.md#DismissRefinementEndpointV1Clause) | **Post** /v1/clause-packs/dismiss-refinement | Dismiss Refinement Endpoint
[**ForkPackEndpointV1**](ClausePacksAPI.md#ForkPackEndpointV1) | **Post** /v1/clause-packs/fork | Fork Pack Endpoint
[**GetAuditV1**](ClausePacksAPI.md#GetAuditV1) | **Get** /v1/clause-packs/audit/{audit_id} | Get Audit
[**GetEffectivePackEndpointV1**](ClausePacksAPI.md#GetEffectivePackEndpointV1) | **Get** /v1/clause-packs/effective | Get Effective Pack Endpoint
[**GetProposedRefinementsV1Clause**](ClausePacksAPI.md#GetProposedRefinementsV1Clause) | **Get** /v1/clause-packs/proposed-refinements | Get Proposed Refinements
[**PostAuditV1**](ClausePacksAPI.md#PostAuditV1) | **Post** /v1/clause-packs/audit | Post Audit
[**PostQuickRedlineV1Clause**](ClausePacksAPI.md#PostQuickRedlineV1Clause) | **Post** /v1/clause-packs/quick-redline | Post Quick Redline
[**PostTrainingUploadV1Clause**](ClausePacksAPI.md#PostTrainingUploadV1Clause) | **Post** /v1/clause-packs/training-uploads | Post Training Upload
[**RatifyFamilyEndpointV1Clause**](ClausePacksAPI.md#RatifyFamilyEndpointV1Clause) | **Post** /v1/clause-packs/ratify-family | Ratify Family Endpoint
[**RatifyRefinementEndpointV1Clause**](ClausePacksAPI.md#RatifyRefinementEndpointV1Clause) | **Post** /v1/clause-packs/ratify-refinement | Ratify Refinement Endpoint



## DismissRefinementEndpointV1Clause

> interface{} DismissRefinementEndpointV1Clause(ctx).DismissRefinementRequest(dismissRefinementRequest).Execute()

Dismiss Refinement Endpoint



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
	dismissRefinementRequest := *openapiclient.NewDismissRefinementRequest("FeedbackId_example") // DismissRefinementRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClausePacksAPI.DismissRefinementEndpointV1Clause(context.Background()).DismissRefinementRequest(dismissRefinementRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClausePacksAPI.DismissRefinementEndpointV1Clause``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DismissRefinementEndpointV1Clause`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ClausePacksAPI.DismissRefinementEndpointV1Clause`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDismissRefinementEndpointV1ClauseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dismissRefinementRequest** | [**DismissRefinementRequest**](DismissRefinementRequest.md) |  | 

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


## ForkPackEndpointV1

> ClausePack ForkPackEndpointV1(ctx).IndustryVertical(industryVertical).Execute()

Fork Pack Endpoint



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
	industryVertical := "industryVertical_example" // string | Industry vertical to fork

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClausePacksAPI.ForkPackEndpointV1(context.Background()).IndustryVertical(industryVertical).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClausePacksAPI.ForkPackEndpointV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ForkPackEndpointV1`: ClausePack
	fmt.Fprintf(os.Stdout, "Response from `ClausePacksAPI.ForkPackEndpointV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiForkPackEndpointV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **industryVertical** | **string** | Industry vertical to fork | 

### Return type

[**ClausePack**](ClausePack.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuditV1

> map[string]interface{} GetAuditV1(ctx, auditId).Execute()

Get Audit



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
	auditId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClausePacksAPI.GetAuditV1(context.Background(), auditId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClausePacksAPI.GetAuditV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuditV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ClausePacksAPI.GetAuditV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEffectivePackEndpointV1

> ClausePack GetEffectivePackEndpointV1(ctx).IndustryVertical(industryVertical).Execute()

Get Effective Pack Endpoint



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
	industryVertical := "industryVertical_example" // string | Industry vertical, e.g. 'ai-saas'

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClausePacksAPI.GetEffectivePackEndpointV1(context.Background()).IndustryVertical(industryVertical).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClausePacksAPI.GetEffectivePackEndpointV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEffectivePackEndpointV1`: ClausePack
	fmt.Fprintf(os.Stdout, "Response from `ClausePacksAPI.GetEffectivePackEndpointV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEffectivePackEndpointV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **industryVertical** | **string** | Industry vertical, e.g. &#39;ai-saas&#39; | 

### Return type

[**ClausePack**](ClausePack.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProposedRefinementsV1Clause

> []map[string]interface{} GetProposedRefinementsV1Clause(ctx).IndustryVertical(industryVertical).Execute()

Get Proposed Refinements



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
	industryVertical := "industryVertical_example" // string | Filter by vertical (optional) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClausePacksAPI.GetProposedRefinementsV1Clause(context.Background()).IndustryVertical(industryVertical).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClausePacksAPI.GetProposedRefinementsV1Clause``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProposedRefinementsV1Clause`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ClausePacksAPI.GetProposedRefinementsV1Clause`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProposedRefinementsV1ClauseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **industryVertical** | **string** | Filter by vertical (optional) | 

### Return type

[**[]map[string]interface{}**](map.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAuditV1

> interface{} PostAuditV1(ctx).AuditRequest(auditRequest).Execute()

Post Audit



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
	auditRequest := *openapiclient.NewAuditRequest("IndustryVertical_example", []string{"ContractUploadIds_example"}) // AuditRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClausePacksAPI.PostAuditV1(context.Background()).AuditRequest(auditRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClausePacksAPI.PostAuditV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAuditV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ClausePacksAPI.PostAuditV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuditV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **auditRequest** | [**AuditRequest**](AuditRequest.md) |  | 

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


## PostQuickRedlineV1Clause

> QuickRedlineResponse PostQuickRedlineV1Clause(ctx).File(file).IndustryVertical(industryVertical).Execute()

Post Quick Redline



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
	file := os.NewFile(1234, "some_file") // *os.File | 
	industryVertical := "industryVertical_example" // string | Industry vertical, e.g. 'ai-saas' (optional) (default to "ai-saas")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClausePacksAPI.PostQuickRedlineV1Clause(context.Background()).File(file).IndustryVertical(industryVertical).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClausePacksAPI.PostQuickRedlineV1Clause``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostQuickRedlineV1Clause`: QuickRedlineResponse
	fmt.Fprintf(os.Stdout, "Response from `ClausePacksAPI.PostQuickRedlineV1Clause`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostQuickRedlineV1ClauseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 
 **industryVertical** | **string** | Industry vertical, e.g. &#39;ai-saas&#39; | [default to &quot;ai-saas&quot;]

### Return type

[**QuickRedlineResponse**](QuickRedlineResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTrainingUploadV1Clause

> interface{} PostTrainingUploadV1Clause(ctx).File(file).Execute()

Post Training Upload



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
	file := os.NewFile(1234, "some_file") // *os.File | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClausePacksAPI.PostTrainingUploadV1Clause(context.Background()).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClausePacksAPI.PostTrainingUploadV1Clause``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTrainingUploadV1Clause`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ClausePacksAPI.PostTrainingUploadV1Clause`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTrainingUploadV1ClauseRequest struct via the builder pattern


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


## RatifyFamilyEndpointV1Clause

> ClausePack RatifyFamilyEndpointV1Clause(ctx).RatifyFamilyRequest(ratifyFamilyRequest).Execute()

Ratify Family Endpoint



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
	ratifyFamilyRequest := *openapiclient.NewRatifyFamilyRequest("IndustryVertical_example", "FamilyId_example", "Owner_example") // RatifyFamilyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClausePacksAPI.RatifyFamilyEndpointV1Clause(context.Background()).RatifyFamilyRequest(ratifyFamilyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClausePacksAPI.RatifyFamilyEndpointV1Clause``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RatifyFamilyEndpointV1Clause`: ClausePack
	fmt.Fprintf(os.Stdout, "Response from `ClausePacksAPI.RatifyFamilyEndpointV1Clause`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRatifyFamilyEndpointV1ClauseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ratifyFamilyRequest** | [**RatifyFamilyRequest**](RatifyFamilyRequest.md) |  | 

### Return type

[**ClausePack**](ClausePack.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RatifyRefinementEndpointV1Clause

> ClausePack RatifyRefinementEndpointV1Clause(ctx).RatifyRefinementRequest(ratifyRefinementRequest).Execute()

Ratify Refinement Endpoint



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
	ratifyRefinementRequest := *openapiclient.NewRatifyRefinementRequest("FeedbackId_example", "IndustryVertical_example", "Owner_example") // RatifyRefinementRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClausePacksAPI.RatifyRefinementEndpointV1Clause(context.Background()).RatifyRefinementRequest(ratifyRefinementRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClausePacksAPI.RatifyRefinementEndpointV1Clause``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RatifyRefinementEndpointV1Clause`: ClausePack
	fmt.Fprintf(os.Stdout, "Response from `ClausePacksAPI.RatifyRefinementEndpointV1Clause`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRatifyRefinementEndpointV1ClauseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ratifyRefinementRequest** | [**RatifyRefinementRequest**](RatifyRefinementRequest.md) |  | 

### Return type

[**ClausePack**](ClausePack.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

