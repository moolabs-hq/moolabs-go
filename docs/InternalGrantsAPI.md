# \InternalGrantsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateGrants**](InternalGrantsAPI.md#ActivateGrants) | **Post** /v1/internal/grants/activate | Activate Grants
[**PlanChangeGrantsV1**](InternalGrantsAPI.md#PlanChangeGrantsV1) | **Post** /v1/internal/grants/plan-change | Plan Change Grants
[**RenewGrants**](InternalGrantsAPI.md#RenewGrants) | **Post** /v1/internal/grants/renew | Renew Grants



## ActivateGrants

> interface{} ActivateGrants(ctx).GrantActivationRequest(grantActivationRequest).DryRun(dryRun).IdempotencyKey(idempotencyKey).Authorization(authorization).Execute()

Activate Grants

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
	grantActivationRequest := *openapiclient.NewGrantActivationRequest("TenantId_example", "SubscriptionId_example", "CustomerId_example", "PlanId_example", int32(123), time.Now(), *openapiclient.NewCreditConfigPayload(false, false, false, *openapiclient.NewInitialamount(), *openapiclient.NewRecurringamount(), "Cadence_example", int32(123), false), []openapiclient.EntitlementPayload{*openapiclient.NewEntitlementPayload("EntitlementId_example", "FeatureKey_example", "EntitlementType_example")}, []string{"SubjectKeys_example"}) // GrantActivationRequest | 
	dryRun := true // bool |  (optional) (default to false)
	idempotencyKey := "idempotencyKey_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InternalGrantsAPI.ActivateGrants(context.Background()).GrantActivationRequest(grantActivationRequest).DryRun(dryRun).IdempotencyKey(idempotencyKey).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalGrantsAPI.ActivateGrants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActivateGrants`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `InternalGrantsAPI.ActivateGrants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActivateGrantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantActivationRequest** | [**GrantActivationRequest**](GrantActivationRequest.md) |  | 
 **dryRun** | **bool** |  | [default to false]
 **idempotencyKey** | **string** |  | 
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


## PlanChangeGrantsV1

> interface{} PlanChangeGrantsV1(ctx).GrantPlanChangeRequest(grantPlanChangeRequest).DryRun(dryRun).IdempotencyKey(idempotencyKey).Authorization(authorization).Execute()

Plan Change Grants

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
	grantPlanChangeRequest := *openapiclient.NewGrantPlanChangeRequest("TenantId_example", "SubscriptionId_example", "CustomerId_example", "OldPlanId_example", int32(123), "NewPlanId_example", int32(123), time.Now(), *openapiclient.NewCreditConfigPayload(false, false, false, *openapiclient.NewInitialamount(), *openapiclient.NewRecurringamount(), "Cadence_example", int32(123), false), []openapiclient.EntitlementPayload{*openapiclient.NewEntitlementPayload("EntitlementId_example", "FeatureKey_example", "EntitlementType_example")}, []string{"SubjectKeys_example"}) // GrantPlanChangeRequest | 
	dryRun := true // bool |  (optional) (default to false)
	idempotencyKey := "idempotencyKey_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InternalGrantsAPI.PlanChangeGrantsV1(context.Background()).GrantPlanChangeRequest(grantPlanChangeRequest).DryRun(dryRun).IdempotencyKey(idempotencyKey).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalGrantsAPI.PlanChangeGrantsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlanChangeGrantsV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `InternalGrantsAPI.PlanChangeGrantsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlanChangeGrantsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantPlanChangeRequest** | [**GrantPlanChangeRequest**](GrantPlanChangeRequest.md) |  | 
 **dryRun** | **bool** |  | [default to false]
 **idempotencyKey** | **string** |  | 
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


## RenewGrants

> interface{} RenewGrants(ctx).GrantRenewalRequest(grantRenewalRequest).DryRun(dryRun).IdempotencyKey(idempotencyKey).Authorization(authorization).Execute()

Renew Grants

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
	grantRenewalRequest := *openapiclient.NewGrantRenewalRequest("TenantId_example", "SubscriptionId_example", "CustomerId_example", "PlanId_example", int32(123), time.Now(), *openapiclient.NewCreditConfigPayload(false, false, false, *openapiclient.NewInitialamount(), *openapiclient.NewRecurringamount(), "Cadence_example", int32(123), false), []openapiclient.EntitlementPayload{*openapiclient.NewEntitlementPayload("EntitlementId_example", "FeatureKey_example", "EntitlementType_example")}, []string{"SubjectKeys_example"}) // GrantRenewalRequest | 
	dryRun := true // bool |  (optional) (default to false)
	idempotencyKey := "idempotencyKey_example" // string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InternalGrantsAPI.RenewGrants(context.Background()).GrantRenewalRequest(grantRenewalRequest).DryRun(dryRun).IdempotencyKey(idempotencyKey).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalGrantsAPI.RenewGrants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RenewGrants`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `InternalGrantsAPI.RenewGrants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRenewGrantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantRenewalRequest** | [**GrantRenewalRequest**](GrantRenewalRequest.md) |  | 
 **dryRun** | **bool** |  | [default to false]
 **idempotencyKey** | **string** |  | 
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

