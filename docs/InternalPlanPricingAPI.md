# \InternalPlanPricingAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SyncPlanPricingV1**](InternalPlanPricingAPI.md#SyncPlanPricingV1) | **Post** /v1/internal/plan-pricing/sync | Sync Plan Pricing



## SyncPlanPricingV1

> interface{} SyncPlanPricingV1(ctx).PlanPricingSyncRequest(planPricingSyncRequest).Authorization(authorization).Execute()

Sync Plan Pricing

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
	planPricingSyncRequest := *openapiclient.NewPlanPricingSyncRequest("TenantId_example", "Operation_example", "PlanId_example") // PlanPricingSyncRequest | 
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InternalPlanPricingAPI.SyncPlanPricingV1(context.Background()).PlanPricingSyncRequest(planPricingSyncRequest).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalPlanPricingAPI.SyncPlanPricingV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncPlanPricingV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `InternalPlanPricingAPI.SyncPlanPricingV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncPlanPricingV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **planPricingSyncRequest** | [**PlanPricingSyncRequest**](PlanPricingSyncRequest.md) |  | 
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

