# \EntitlementsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomerEntitlementGrantV2**](EntitlementsAPI.md#CreateCustomerEntitlementGrantV2) | **Post** /api/v2/customers/{customerIdOrKey}/entitlements/{entitlementIdOrFeatureKey}/grants | Create customer entitlement grant
[**CreateCustomerEntitlementV2**](EntitlementsAPI.md#CreateCustomerEntitlementV2) | **Post** /api/v2/customers/{customerIdOrKey}/entitlements | Create a customer entitlement
[**CreateEntitlement**](EntitlementsAPI.md#CreateEntitlement) | **Post** /api/v1/subjects/{subjectIdOrKey}/entitlements | Create a subject entitlement
[**CreateGrant**](EntitlementsAPI.md#CreateGrant) | **Post** /api/v1/subjects/{subjectIdOrKey}/entitlements/{entitlementIdOrFeatureKey}/grants | Create subject entitlement grant
[**DeleteCustomerEntitlementV2**](EntitlementsAPI.md#DeleteCustomerEntitlementV2) | **Delete** /api/v2/customers/{customerIdOrKey}/entitlements/{entitlementIdOrFeatureKey} | Delete customer entitlement
[**DeleteEntitlement**](EntitlementsAPI.md#DeleteEntitlement) | **Delete** /api/v1/subjects/{subjectIdOrKey}/entitlements/{entitlementId} | Delete subject entitlement
[**GetCustomerAccess**](EntitlementsAPI.md#GetCustomerAccess) | **Get** /api/v1/customers/{customerIdOrKey}/access | Get customer access
[**GetCustomerEntitlementHistoryV2**](EntitlementsAPI.md#GetCustomerEntitlementHistoryV2) | **Get** /api/v2/customers/{customerIdOrKey}/entitlements/{entitlementIdOrFeatureKey}/history | Get customer entitlement history
[**GetCustomerEntitlementV2**](EntitlementsAPI.md#GetCustomerEntitlementV2) | **Get** /api/v2/customers/{customerIdOrKey}/entitlements/{entitlementIdOrFeatureKey} | Get customer entitlement
[**GetCustomerEntitlementValue**](EntitlementsAPI.md#GetCustomerEntitlementValue) | **Get** /api/v1/customers/{customerIdOrKey}/entitlements/{featureKey}/value | Get customer entitlement value
[**GetCustomerEntitlementValueV2**](EntitlementsAPI.md#GetCustomerEntitlementValueV2) | **Get** /api/v2/customers/{customerIdOrKey}/entitlements/{entitlementIdOrFeatureKey}/value | Get customer entitlement value
[**GetEntitlement**](EntitlementsAPI.md#GetEntitlement) | **Get** /api/v1/subjects/{subjectIdOrKey}/entitlements/{entitlementId} | Get subject entitlement
[**GetEntitlementById**](EntitlementsAPI.md#GetEntitlementById) | **Get** /api/v1/entitlements/{entitlementId} | Get entitlement by id
[**GetEntitlementByIdV2**](EntitlementsAPI.md#GetEntitlementByIdV2) | **Get** /api/v2/entitlements/{entitlementId} | Get entitlement by id
[**GetEntitlementHistory**](EntitlementsAPI.md#GetEntitlementHistory) | **Get** /api/v1/subjects/{subjectIdOrKey}/entitlements/{entitlementId}/history | Get subject entitlement history
[**GetEntitlementValue**](EntitlementsAPI.md#GetEntitlementValue) | **Get** /api/v1/subjects/{subjectIdOrKey}/entitlements/{entitlementIdOrFeatureKey}/value | Get subject entitlement value
[**ListCustomerEntitlementGrantsV2**](EntitlementsAPI.md#ListCustomerEntitlementGrantsV2) | **Get** /api/v2/customers/{customerIdOrKey}/entitlements/{entitlementIdOrFeatureKey}/grants | List customer entitlement grants
[**ListCustomerEntitlementsV2**](EntitlementsAPI.md#ListCustomerEntitlementsV2) | **Get** /api/v2/customers/{customerIdOrKey}/entitlements | List customer entitlements
[**ListEntitlementGrants**](EntitlementsAPI.md#ListEntitlementGrants) | **Get** /api/v1/subjects/{subjectIdOrKey}/entitlements/{entitlementIdOrFeatureKey}/grants | List subject entitlement grants
[**ListEntitlements**](EntitlementsAPI.md#ListEntitlements) | **Get** /api/v1/entitlements | List all entitlements
[**ListEntitlementsV2**](EntitlementsAPI.md#ListEntitlementsV2) | **Get** /api/v2/entitlements | List all entitlements
[**ListGrantsGet**](EntitlementsAPI.md#ListGrantsGet) | **Get** /api/v1/grants | List grants
[**ListGrantsV2**](EntitlementsAPI.md#ListGrantsV2) | **Get** /api/v2/grants | List grants
[**ListSubjectEntitlements**](EntitlementsAPI.md#ListSubjectEntitlements) | **Get** /api/v1/subjects/{subjectIdOrKey}/entitlements | List subject entitlements
[**OverrideCustomerEntitlementV2**](EntitlementsAPI.md#OverrideCustomerEntitlementV2) | **Put** /api/v2/customers/{customerIdOrKey}/entitlements/{entitlementIdOrFeatureKey}/override | Override customer entitlement
[**OverrideEntitlement**](EntitlementsAPI.md#OverrideEntitlement) | **Put** /api/v1/subjects/{subjectIdOrKey}/entitlements/{entitlementIdOrFeatureKey}/override | Override subject entitlement
[**ResetCustomerEntitlementUsageV2**](EntitlementsAPI.md#ResetCustomerEntitlementUsageV2) | **Post** /api/v2/customers/{customerIdOrKey}/entitlements/{entitlementIdOrFeatureKey}/reset | Reset customer entitlement
[**ResetEntitlementUsage**](EntitlementsAPI.md#ResetEntitlementUsage) | **Post** /api/v1/subjects/{subjectIdOrKey}/entitlements/{entitlementId}/reset | Reset subject entitlement
[**VoidGrantDelete**](EntitlementsAPI.md#VoidGrantDelete) | **Delete** /api/v1/grants/{grantId} | Void grant



## CreateCustomerEntitlementGrantV2

> EntitlementGrantV2 CreateCustomerEntitlementGrantV2(ctx, customerIdOrKey, entitlementIdOrFeatureKey).EntitlementGrantCreateInputV2(entitlementGrantCreateInputV2).Execute()

Create customer entitlement grant



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
	customerIdOrKey := *openapiclient.NewULIDOrExternalKey() // ULIDOrExternalKey | 
	entitlementIdOrFeatureKey := "entitlementIdOrFeatureKey_example" // string | 
	entitlementGrantCreateInputV2 := *openapiclient.NewEntitlementGrantCreateInputV2(float64(100), time.Now()) // EntitlementGrantCreateInputV2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.CreateCustomerEntitlementGrantV2(context.Background(), customerIdOrKey, entitlementIdOrFeatureKey).EntitlementGrantCreateInputV2(entitlementGrantCreateInputV2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.CreateCustomerEntitlementGrantV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomerEntitlementGrantV2`: EntitlementGrantV2
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.CreateCustomerEntitlementGrantV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerIdOrKey** | [**ULIDOrExternalKey**](.md) |  | 
**entitlementIdOrFeatureKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerEntitlementGrantV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **entitlementGrantCreateInputV2** | [**EntitlementGrantCreateInputV2**](EntitlementGrantCreateInputV2.md) |  | 

### Return type

[**EntitlementGrantV2**](EntitlementGrantV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCustomerEntitlementV2

> EntitlementV2 CreateCustomerEntitlementV2(ctx, customerIdOrKey).EntitlementV2CreateInputs(entitlementV2CreateInputs).Execute()

Create a customer entitlement



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
	customerIdOrKey := *openapiclient.NewULIDOrExternalKey() // ULIDOrExternalKey | 
	entitlementV2CreateInputs := openapiclient.EntitlementV2CreateInputs{EntitlementBooleanCreateInputs: openapiclient.NewEntitlementBooleanCreateInputs("Type_example")} // EntitlementV2CreateInputs | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.CreateCustomerEntitlementV2(context.Background(), customerIdOrKey).EntitlementV2CreateInputs(entitlementV2CreateInputs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.CreateCustomerEntitlementV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomerEntitlementV2`: EntitlementV2
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.CreateCustomerEntitlementV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerIdOrKey** | [**ULIDOrExternalKey**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerEntitlementV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entitlementV2CreateInputs** | [**EntitlementV2CreateInputs**](EntitlementV2CreateInputs.md) |  | 

### Return type

[**EntitlementV2**](EntitlementV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEntitlement

> Entitlement CreateEntitlement(ctx, subjectIdOrKey).EntitlementCreateInputs(entitlementCreateInputs).Execute()

Create a subject entitlement



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
	subjectIdOrKey := "subjectIdOrKey_example" // string | 
	entitlementCreateInputs := openapiclient.EntitlementCreateInputs{EntitlementBooleanCreateInputs: openapiclient.NewEntitlementBooleanCreateInputs("Type_example")} // EntitlementCreateInputs | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.CreateEntitlement(context.Background(), subjectIdOrKey).EntitlementCreateInputs(entitlementCreateInputs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.CreateEntitlement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEntitlement`: Entitlement
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.CreateEntitlement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectIdOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entitlementCreateInputs** | [**EntitlementCreateInputs**](EntitlementCreateInputs.md) |  | 

### Return type

[**Entitlement**](Entitlement.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGrant

> EntitlementGrant CreateGrant(ctx, subjectIdOrKey, entitlementIdOrFeatureKey).EntitlementGrantCreateInput(entitlementGrantCreateInput).Execute()

Create subject entitlement grant



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
	subjectIdOrKey := "subjectIdOrKey_example" // string | 
	entitlementIdOrFeatureKey := "entitlementIdOrFeatureKey_example" // string | 
	entitlementGrantCreateInput := *openapiclient.NewEntitlementGrantCreateInput(float64(100), time.Now(), *openapiclient.NewExpirationPeriod(openapiclient.ExpirationDuration("HOUR"), int32(12))) // EntitlementGrantCreateInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.CreateGrant(context.Background(), subjectIdOrKey, entitlementIdOrFeatureKey).EntitlementGrantCreateInput(entitlementGrantCreateInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.CreateGrant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGrant`: EntitlementGrant
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.CreateGrant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectIdOrKey** | **string** |  | 
**entitlementIdOrFeatureKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **entitlementGrantCreateInput** | [**EntitlementGrantCreateInput**](EntitlementGrantCreateInput.md) |  | 

### Return type

[**EntitlementGrant**](EntitlementGrant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomerEntitlementV2

> DeleteCustomerEntitlementV2(ctx, customerIdOrKey, entitlementIdOrFeatureKey).Execute()

Delete customer entitlement



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
	customerIdOrKey := *openapiclient.NewULIDOrExternalKey() // ULIDOrExternalKey | 
	entitlementIdOrFeatureKey := "entitlementIdOrFeatureKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EntitlementsAPI.DeleteCustomerEntitlementV2(context.Background(), customerIdOrKey, entitlementIdOrFeatureKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.DeleteCustomerEntitlementV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerIdOrKey** | [**ULIDOrExternalKey**](.md) |  | 
**entitlementIdOrFeatureKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomerEntitlementV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEntitlement

> DeleteEntitlement(ctx, subjectIdOrKey, entitlementId).Execute()

Delete subject entitlement



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
	subjectIdOrKey := "subjectIdOrKey_example" // string | 
	entitlementId := "entitlementId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EntitlementsAPI.DeleteEntitlement(context.Background(), subjectIdOrKey, entitlementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.DeleteEntitlement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectIdOrKey** | **string** |  | 
**entitlementId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerAccess

> CustomerAccess GetCustomerAccess(ctx, customerIdOrKey).Execute()

Get customer access



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
	customerIdOrKey := *openapiclient.NewULIDOrExternalKey() // ULIDOrExternalKey | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.GetCustomerAccess(context.Background(), customerIdOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.GetCustomerAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomerAccess`: CustomerAccess
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.GetCustomerAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerIdOrKey** | [**ULIDOrExternalKey**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomerAccess**](CustomerAccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerEntitlementHistoryV2

> WindowedBalanceHistory GetCustomerEntitlementHistoryV2(ctx, customerIdOrKey, entitlementIdOrFeatureKey).WindowSize(windowSize).From(from).To(to).WindowTimeZone(windowTimeZone).Execute()

Get customer entitlement history



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
	customerIdOrKey := *openapiclient.NewULIDOrExternalKey() // ULIDOrExternalKey | 
	entitlementIdOrFeatureKey := "entitlementIdOrFeatureKey_example" // string | 
	windowSize := openapiclient.WindowSize("MINUTE") // WindowSize | Windowsize
	from := time.Now() // time.Time | Start of time range to query entitlement: date-time in RFC 3339 format. Defaults to the last reset. Gets truncated to the granularity of the underlying meter. (optional)
	to := time.Now() // time.Time | End of time range to query entitlement: date-time in RFC 3339 format. Defaults to now. If not now then gets truncated to the granularity of the underlying meter. (optional)
	windowTimeZone := "windowTimeZone_example" // string | The timezone used when calculating the windows. (optional) (default to "UTC")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.GetCustomerEntitlementHistoryV2(context.Background(), customerIdOrKey, entitlementIdOrFeatureKey).WindowSize(windowSize).From(from).To(to).WindowTimeZone(windowTimeZone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.GetCustomerEntitlementHistoryV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomerEntitlementHistoryV2`: WindowedBalanceHistory
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.GetCustomerEntitlementHistoryV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerIdOrKey** | [**ULIDOrExternalKey**](.md) |  | 
**entitlementIdOrFeatureKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerEntitlementHistoryV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **windowSize** | [**WindowSize**](WindowSize.md) | Windowsize | 
 **from** | **time.Time** | Start of time range to query entitlement: date-time in RFC 3339 format. Defaults to the last reset. Gets truncated to the granularity of the underlying meter. | 
 **to** | **time.Time** | End of time range to query entitlement: date-time in RFC 3339 format. Defaults to now. If not now then gets truncated to the granularity of the underlying meter. | 
 **windowTimeZone** | **string** | The timezone used when calculating the windows. | [default to &quot;UTC&quot;]

### Return type

[**WindowedBalanceHistory**](WindowedBalanceHistory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerEntitlementV2

> EntitlementV2 GetCustomerEntitlementV2(ctx, customerIdOrKey, entitlementIdOrFeatureKey).Execute()

Get customer entitlement



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
	customerIdOrKey := *openapiclient.NewULIDOrExternalKey() // ULIDOrExternalKey | 
	entitlementIdOrFeatureKey := "entitlementIdOrFeatureKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.GetCustomerEntitlementV2(context.Background(), customerIdOrKey, entitlementIdOrFeatureKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.GetCustomerEntitlementV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomerEntitlementV2`: EntitlementV2
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.GetCustomerEntitlementV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerIdOrKey** | [**ULIDOrExternalKey**](.md) |  | 
**entitlementIdOrFeatureKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerEntitlementV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EntitlementV2**](EntitlementV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerEntitlementValue

> EntitlementValue GetCustomerEntitlementValue(ctx, customerIdOrKey, featureKey).Time(time).Execute()

Get customer entitlement value



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
	customerIdOrKey := *openapiclient.NewULIDOrExternalKey() // ULIDOrExternalKey | 
	featureKey := "featureKey_example" // string | 
	time := time.Now() // time.Time |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.GetCustomerEntitlementValue(context.Background(), customerIdOrKey, featureKey).Time(time).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.GetCustomerEntitlementValue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomerEntitlementValue`: EntitlementValue
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.GetCustomerEntitlementValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerIdOrKey** | [**ULIDOrExternalKey**](.md) |  | 
**featureKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerEntitlementValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **time** | **time.Time** |  | 

### Return type

[**EntitlementValue**](EntitlementValue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerEntitlementValueV2

> EntitlementValue GetCustomerEntitlementValueV2(ctx, customerIdOrKey, entitlementIdOrFeatureKey).Time(time).Execute()

Get customer entitlement value



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
	customerIdOrKey := *openapiclient.NewULIDOrExternalKey() // ULIDOrExternalKey | 
	entitlementIdOrFeatureKey := "entitlementIdOrFeatureKey_example" // string | 
	time := time.Now() // time.Time |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.GetCustomerEntitlementValueV2(context.Background(), customerIdOrKey, entitlementIdOrFeatureKey).Time(time).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.GetCustomerEntitlementValueV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomerEntitlementValueV2`: EntitlementValue
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.GetCustomerEntitlementValueV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerIdOrKey** | [**ULIDOrExternalKey**](.md) |  | 
**entitlementIdOrFeatureKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerEntitlementValueV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **time** | **time.Time** |  | 

### Return type

[**EntitlementValue**](EntitlementValue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntitlement

> Entitlement GetEntitlement(ctx, subjectIdOrKey, entitlementId).Execute()

Get subject entitlement



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
	subjectIdOrKey := "subjectIdOrKey_example" // string | 
	entitlementId := "entitlementId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.GetEntitlement(context.Background(), subjectIdOrKey, entitlementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.GetEntitlement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEntitlement`: Entitlement
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.GetEntitlement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectIdOrKey** | **string** |  | 
**entitlementId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Entitlement**](Entitlement.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntitlementById

> Entitlement GetEntitlementById(ctx, entitlementId).Execute()

Get entitlement by id



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
	entitlementId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.GetEntitlementById(context.Background(), entitlementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.GetEntitlementById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEntitlementById`: Entitlement
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.GetEntitlementById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entitlementId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntitlementByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Entitlement**](Entitlement.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntitlementByIdV2

> EntitlementV2 GetEntitlementByIdV2(ctx, entitlementId).Execute()

Get entitlement by id



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
	entitlementId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.GetEntitlementByIdV2(context.Background(), entitlementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.GetEntitlementByIdV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEntitlementByIdV2`: EntitlementV2
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.GetEntitlementByIdV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entitlementId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntitlementByIdV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EntitlementV2**](EntitlementV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntitlementHistory

> WindowedBalanceHistory GetEntitlementHistory(ctx, subjectIdOrKey, entitlementId).WindowSize(windowSize).From(from).To(to).WindowTimeZone(windowTimeZone).Execute()

Get subject entitlement history



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
	subjectIdOrKey := "subjectIdOrKey_example" // string | 
	entitlementId := "entitlementId_example" // string | 
	windowSize := openapiclient.WindowSize("MINUTE") // WindowSize | Windowsize
	from := time.Now() // time.Time | Start of time range to query entitlement: date-time in RFC 3339 format. Defaults to the last reset. Gets truncated to the granularity of the underlying meter. (optional)
	to := time.Now() // time.Time | End of time range to query entitlement: date-time in RFC 3339 format. Defaults to now. If not now then gets truncated to the granularity of the underlying meter. (optional)
	windowTimeZone := "windowTimeZone_example" // string | The timezone used when calculating the windows. (optional) (default to "UTC")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.GetEntitlementHistory(context.Background(), subjectIdOrKey, entitlementId).WindowSize(windowSize).From(from).To(to).WindowTimeZone(windowTimeZone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.GetEntitlementHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEntitlementHistory`: WindowedBalanceHistory
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.GetEntitlementHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectIdOrKey** | **string** |  | 
**entitlementId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntitlementHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **windowSize** | [**WindowSize**](WindowSize.md) | Windowsize | 
 **from** | **time.Time** | Start of time range to query entitlement: date-time in RFC 3339 format. Defaults to the last reset. Gets truncated to the granularity of the underlying meter. | 
 **to** | **time.Time** | End of time range to query entitlement: date-time in RFC 3339 format. Defaults to now. If not now then gets truncated to the granularity of the underlying meter. | 
 **windowTimeZone** | **string** | The timezone used when calculating the windows. | [default to &quot;UTC&quot;]

### Return type

[**WindowedBalanceHistory**](WindowedBalanceHistory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntitlementValue

> EntitlementValue GetEntitlementValue(ctx, subjectIdOrKey, entitlementIdOrFeatureKey).Time(time).Execute()

Get subject entitlement value



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
	subjectIdOrKey := "subjectIdOrKey_example" // string | 
	entitlementIdOrFeatureKey := "entitlementIdOrFeatureKey_example" // string | 
	time := time.Now() // time.Time |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.GetEntitlementValue(context.Background(), subjectIdOrKey, entitlementIdOrFeatureKey).Time(time).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.GetEntitlementValue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEntitlementValue`: EntitlementValue
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.GetEntitlementValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectIdOrKey** | **string** |  | 
**entitlementIdOrFeatureKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntitlementValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **time** | **time.Time** |  | 

### Return type

[**EntitlementValue**](EntitlementValue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCustomerEntitlementGrantsV2

> GrantV2PaginatedResponse ListCustomerEntitlementGrantsV2(ctx, customerIdOrKey, entitlementIdOrFeatureKey).IncludeDeleted(includeDeleted).Page(page).PageSize(pageSize).Offset(offset).Limit(limit).Order(order).OrderBy(orderBy).Execute()

List customer entitlement grants



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
	customerIdOrKey := *openapiclient.NewULIDOrExternalKey() // ULIDOrExternalKey | 
	entitlementIdOrFeatureKey := "entitlementIdOrFeatureKey_example" // string | 
	includeDeleted := true // bool |  (optional) (default to false)
	page := int32(56) // int32 | Page index.  Default is 1. (optional) (default to 1)
	pageSize := int32(56) // int32 | The maximum number of items per page.  Default is 100. (optional) (default to 100)
	offset := int32(56) // int32 | Number of items to skip.  Default is 0. (optional) (default to 0)
	limit := int32(56) // int32 | Number of items to return.  Default is 100. (optional) (default to 100)
	order := openapiclient.SortOrder("ASC") // SortOrder | The order direction. (optional)
	orderBy := openapiclient.GrantOrderBy("id") // GrantOrderBy | The order by field. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.ListCustomerEntitlementGrantsV2(context.Background(), customerIdOrKey, entitlementIdOrFeatureKey).IncludeDeleted(includeDeleted).Page(page).PageSize(pageSize).Offset(offset).Limit(limit).Order(order).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.ListCustomerEntitlementGrantsV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCustomerEntitlementGrantsV2`: GrantV2PaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.ListCustomerEntitlementGrantsV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerIdOrKey** | [**ULIDOrExternalKey**](.md) |  | 
**entitlementIdOrFeatureKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCustomerEntitlementGrantsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeDeleted** | **bool** |  | [default to false]
 **page** | **int32** | Page index.  Default is 1. | [default to 1]
 **pageSize** | **int32** | The maximum number of items per page.  Default is 100. | [default to 100]
 **offset** | **int32** | Number of items to skip.  Default is 0. | [default to 0]
 **limit** | **int32** | Number of items to return.  Default is 100. | [default to 100]
 **order** | [**SortOrder**](SortOrder.md) | The order direction. | 
 **orderBy** | [**GrantOrderBy**](GrantOrderBy.md) | The order by field. | 

### Return type

[**GrantV2PaginatedResponse**](GrantV2PaginatedResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCustomerEntitlementsV2

> EntitlementV2PaginatedResponse ListCustomerEntitlementsV2(ctx, customerIdOrKey).IncludeDeleted(includeDeleted).Page(page).PageSize(pageSize).Order(order).OrderBy(orderBy).Execute()

List customer entitlements



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
	customerIdOrKey := *openapiclient.NewULIDOrExternalKey() // ULIDOrExternalKey | 
	includeDeleted := true // bool |  (optional) (default to false)
	page := int32(56) // int32 | Page index.  Default is 1. (optional) (default to 1)
	pageSize := int32(56) // int32 | The maximum number of items per page.  Default is 100. (optional) (default to 100)
	order := openapiclient.SortOrder("ASC") // SortOrder | The order direction. (optional)
	orderBy := openapiclient.EntitlementOrderBy("createdAt") // EntitlementOrderBy | The order by field. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.ListCustomerEntitlementsV2(context.Background(), customerIdOrKey).IncludeDeleted(includeDeleted).Page(page).PageSize(pageSize).Order(order).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.ListCustomerEntitlementsV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCustomerEntitlementsV2`: EntitlementV2PaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.ListCustomerEntitlementsV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerIdOrKey** | [**ULIDOrExternalKey**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCustomerEntitlementsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeDeleted** | **bool** |  | [default to false]
 **page** | **int32** | Page index.  Default is 1. | [default to 1]
 **pageSize** | **int32** | The maximum number of items per page.  Default is 100. | [default to 100]
 **order** | [**SortOrder**](SortOrder.md) | The order direction. | 
 **orderBy** | [**EntitlementOrderBy**](EntitlementOrderBy.md) | The order by field. | 

### Return type

[**EntitlementV2PaginatedResponse**](EntitlementV2PaginatedResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEntitlementGrants

> []EntitlementGrant ListEntitlementGrants(ctx, subjectIdOrKey, entitlementIdOrFeatureKey).IncludeDeleted(includeDeleted).OrderBy(orderBy).Execute()

List subject entitlement grants



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
	subjectIdOrKey := "subjectIdOrKey_example" // string | 
	entitlementIdOrFeatureKey := "entitlementIdOrFeatureKey_example" // string | 
	includeDeleted := true // bool |  (optional) (default to false)
	orderBy := openapiclient.GrantOrderBy("id") // GrantOrderBy |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.ListEntitlementGrants(context.Background(), subjectIdOrKey, entitlementIdOrFeatureKey).IncludeDeleted(includeDeleted).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.ListEntitlementGrants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEntitlementGrants`: []EntitlementGrant
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.ListEntitlementGrants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectIdOrKey** | **string** |  | 
**entitlementIdOrFeatureKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEntitlementGrantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeDeleted** | **bool** |  | [default to false]
 **orderBy** | [**GrantOrderBy**](GrantOrderBy.md) |  | 

### Return type

[**[]EntitlementGrant**](EntitlementGrant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEntitlements

> ListEntitlementsResult ListEntitlements(ctx).Feature(feature).Subject(subject).EntitlementType(entitlementType).ExcludeInactive(excludeInactive).Page(page).PageSize(pageSize).Offset(offset).Limit(limit).Order(order).OrderBy(orderBy).Execute()

List all entitlements



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
	feature := []string{"Inner_example"} // []string | Filtering by multiple features.  Usage: `?feature=feature-1&feature=feature-2` (optional)
	subject := []string{"Inner_example"} // []string | Filtering by multiple subjects.  Usage: `?subject=customer-1&subject=customer-2` (optional)
	entitlementType := []openapiclient.EntitlementType{openapiclient.EntitlementType("metered")} // []EntitlementType | Filtering by multiple entitlement types.  Usage: `?entitlementType=metered&entitlementType=boolean` (optional)
	excludeInactive := true // bool | Exclude inactive entitlements in the response (those scheduled for later or earlier) (optional) (default to false)
	page := int32(56) // int32 | Page index.  Default is 1. (optional) (default to 1)
	pageSize := int32(56) // int32 | The maximum number of items per page.  Default is 100. (optional) (default to 100)
	offset := int32(56) // int32 | Number of items to skip.  Default is 0. (optional) (default to 0)
	limit := int32(56) // int32 | Number of items to return.  Default is 100. (optional) (default to 100)
	order := openapiclient.SortOrder("ASC") // SortOrder | The order direction. (optional)
	orderBy := openapiclient.EntitlementOrderBy("createdAt") // EntitlementOrderBy | The order by field. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.ListEntitlements(context.Background()).Feature(feature).Subject(subject).EntitlementType(entitlementType).ExcludeInactive(excludeInactive).Page(page).PageSize(pageSize).Offset(offset).Limit(limit).Order(order).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.ListEntitlements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEntitlements`: ListEntitlementsResult
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.ListEntitlements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **feature** | **[]string** | Filtering by multiple features.  Usage: &#x60;?feature&#x3D;feature-1&amp;feature&#x3D;feature-2&#x60; | 
 **subject** | **[]string** | Filtering by multiple subjects.  Usage: &#x60;?subject&#x3D;customer-1&amp;subject&#x3D;customer-2&#x60; | 
 **entitlementType** | [**[]EntitlementType**](EntitlementType.md) | Filtering by multiple entitlement types.  Usage: &#x60;?entitlementType&#x3D;metered&amp;entitlementType&#x3D;boolean&#x60; | 
 **excludeInactive** | **bool** | Exclude inactive entitlements in the response (those scheduled for later or earlier) | [default to false]
 **page** | **int32** | Page index.  Default is 1. | [default to 1]
 **pageSize** | **int32** | The maximum number of items per page.  Default is 100. | [default to 100]
 **offset** | **int32** | Number of items to skip.  Default is 0. | [default to 0]
 **limit** | **int32** | Number of items to return.  Default is 100. | [default to 100]
 **order** | [**SortOrder**](SortOrder.md) | The order direction. | 
 **orderBy** | [**EntitlementOrderBy**](EntitlementOrderBy.md) | The order by field. | 

### Return type

[**ListEntitlementsResult**](ListEntitlementsResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEntitlementsV2

> EntitlementV2PaginatedResponse ListEntitlementsV2(ctx).Feature(feature).CustomerKeys(customerKeys).CustomerIds(customerIds).EntitlementType(entitlementType).ExcludeInactive(excludeInactive).Page(page).PageSize(pageSize).Offset(offset).Limit(limit).Order(order).OrderBy(orderBy).Execute()

List all entitlements



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
	feature := []string{"Inner_example"} // []string | Filtering by multiple features.  Usage: `?feature=feature-1&feature=feature-2` (optional)
	customerKeys := []string{"Inner_example"} // []string | Filtering by multiple customers.  Usage: `?customerKeys=customer-1&customerKeys=customer-3` (optional)
	customerIds := []string{"Inner_example"} // []string | Filtering by multiple customers.  Usage: `?customerIds=01K4WAQ0J99ZZ0MD75HXR112H8&customerIds=01K4WAQ0J99ZZ0MD75HXR112H9` (optional)
	entitlementType := []openapiclient.EntitlementType{openapiclient.EntitlementType("metered")} // []EntitlementType | Filtering by multiple entitlement types.  Usage: `?entitlementType=metered&entitlementType=boolean` (optional)
	excludeInactive := true // bool | Exclude inactive entitlements in the response (those scheduled for later or earlier) (optional) (default to false)
	page := int32(56) // int32 | Page index.  Default is 1. (optional) (default to 1)
	pageSize := int32(56) // int32 | The maximum number of items per page.  Default is 100. (optional) (default to 100)
	offset := int32(56) // int32 | Number of items to skip.  Default is 0. (optional) (default to 0)
	limit := int32(56) // int32 | Number of items to return.  Default is 100. (optional) (default to 100)
	order := openapiclient.SortOrder("ASC") // SortOrder | The order direction. (optional)
	orderBy := openapiclient.EntitlementOrderBy("createdAt") // EntitlementOrderBy | The order by field. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.ListEntitlementsV2(context.Background()).Feature(feature).CustomerKeys(customerKeys).CustomerIds(customerIds).EntitlementType(entitlementType).ExcludeInactive(excludeInactive).Page(page).PageSize(pageSize).Offset(offset).Limit(limit).Order(order).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.ListEntitlementsV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEntitlementsV2`: EntitlementV2PaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.ListEntitlementsV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEntitlementsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **feature** | **[]string** | Filtering by multiple features.  Usage: &#x60;?feature&#x3D;feature-1&amp;feature&#x3D;feature-2&#x60; | 
 **customerKeys** | **[]string** | Filtering by multiple customers.  Usage: &#x60;?customerKeys&#x3D;customer-1&amp;customerKeys&#x3D;customer-3&#x60; | 
 **customerIds** | **[]string** | Filtering by multiple customers.  Usage: &#x60;?customerIds&#x3D;01K4WAQ0J99ZZ0MD75HXR112H8&amp;customerIds&#x3D;01K4WAQ0J99ZZ0MD75HXR112H9&#x60; | 
 **entitlementType** | [**[]EntitlementType**](EntitlementType.md) | Filtering by multiple entitlement types.  Usage: &#x60;?entitlementType&#x3D;metered&amp;entitlementType&#x3D;boolean&#x60; | 
 **excludeInactive** | **bool** | Exclude inactive entitlements in the response (those scheduled for later or earlier) | [default to false]
 **page** | **int32** | Page index.  Default is 1. | [default to 1]
 **pageSize** | **int32** | The maximum number of items per page.  Default is 100. | [default to 100]
 **offset** | **int32** | Number of items to skip.  Default is 0. | [default to 0]
 **limit** | **int32** | Number of items to return.  Default is 100. | [default to 100]
 **order** | [**SortOrder**](SortOrder.md) | The order direction. | 
 **orderBy** | [**EntitlementOrderBy**](EntitlementOrderBy.md) | The order by field. | 

### Return type

[**EntitlementV2PaginatedResponse**](EntitlementV2PaginatedResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGrantsGet

> ListGrantsGet200Response ListGrantsGet(ctx).Feature(feature).Subject(subject).IncludeDeleted(includeDeleted).Page(page).PageSize(pageSize).Offset(offset).Limit(limit).Order(order).OrderBy(orderBy).Execute()

List grants



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
	feature := []string{"Inner_example"} // []string | Filtering by multiple features.  Usage: `?feature=feature-1&feature=feature-2` (optional)
	subject := []string{"Inner_example"} // []string | Filtering by multiple subjects.  Usage: `?subject=customer-1&subject=customer-2` (optional)
	includeDeleted := true // bool | Include deleted (optional) (default to false)
	page := int32(56) // int32 | Page index.  Default is 1. (optional) (default to 1)
	pageSize := int32(56) // int32 | The maximum number of items per page.  Default is 100. (optional) (default to 100)
	offset := int32(56) // int32 | Number of items to skip.  Default is 0. (optional) (default to 0)
	limit := int32(56) // int32 | Number of items to return.  Default is 100. (optional) (default to 100)
	order := openapiclient.SortOrder("ASC") // SortOrder | The order direction. (optional)
	orderBy := openapiclient.GrantOrderBy("id") // GrantOrderBy | The order by field. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.ListGrantsGet(context.Background()).Feature(feature).Subject(subject).IncludeDeleted(includeDeleted).Page(page).PageSize(pageSize).Offset(offset).Limit(limit).Order(order).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.ListGrantsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGrantsGet`: ListGrantsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.ListGrantsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGrantsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **feature** | **[]string** | Filtering by multiple features.  Usage: &#x60;?feature&#x3D;feature-1&amp;feature&#x3D;feature-2&#x60; | 
 **subject** | **[]string** | Filtering by multiple subjects.  Usage: &#x60;?subject&#x3D;customer-1&amp;subject&#x3D;customer-2&#x60; | 
 **includeDeleted** | **bool** | Include deleted | [default to false]
 **page** | **int32** | Page index.  Default is 1. | [default to 1]
 **pageSize** | **int32** | The maximum number of items per page.  Default is 100. | [default to 100]
 **offset** | **int32** | Number of items to skip.  Default is 0. | [default to 0]
 **limit** | **int32** | Number of items to return.  Default is 100. | [default to 100]
 **order** | [**SortOrder**](SortOrder.md) | The order direction. | 
 **orderBy** | [**GrantOrderBy**](GrantOrderBy.md) | The order by field. | 

### Return type

[**ListGrantsGet200Response**](ListGrantsGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGrantsV2

> GrantV2PaginatedResponse ListGrantsV2(ctx).Feature(feature).Customer(customer).IncludeDeleted(includeDeleted).Page(page).PageSize(pageSize).Offset(offset).Limit(limit).Order(order).OrderBy(orderBy).Execute()

List grants



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
	feature := []string{"Inner_example"} // []string | Filtering by multiple features.  Usage: `?feature=feature-1&feature=feature-2` (optional)
	customer := []openapiclient.ULIDOrExternalKey{*openapiclient.NewULIDOrExternalKey()} // []ULIDOrExternalKey | Filtering by multiple customers (either by ID or key).  Usage: `?customer=customer-1&customer=customer-2` (optional)
	includeDeleted := true // bool | Include deleted (optional) (default to false)
	page := int32(56) // int32 | Page index.  Default is 1. (optional) (default to 1)
	pageSize := int32(56) // int32 | The maximum number of items per page.  Default is 100. (optional) (default to 100)
	offset := int32(56) // int32 | Number of items to skip.  Default is 0. (optional) (default to 0)
	limit := int32(56) // int32 | Number of items to return.  Default is 100. (optional) (default to 100)
	order := openapiclient.SortOrder("ASC") // SortOrder | The order direction. (optional)
	orderBy := openapiclient.GrantOrderBy("id") // GrantOrderBy | The order by field. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.ListGrantsV2(context.Background()).Feature(feature).Customer(customer).IncludeDeleted(includeDeleted).Page(page).PageSize(pageSize).Offset(offset).Limit(limit).Order(order).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.ListGrantsV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGrantsV2`: GrantV2PaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.ListGrantsV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGrantsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **feature** | **[]string** | Filtering by multiple features.  Usage: &#x60;?feature&#x3D;feature-1&amp;feature&#x3D;feature-2&#x60; | 
 **customer** | [**[]ULIDOrExternalKey**](ULIDOrExternalKey.md) | Filtering by multiple customers (either by ID or key).  Usage: &#x60;?customer&#x3D;customer-1&amp;customer&#x3D;customer-2&#x60; | 
 **includeDeleted** | **bool** | Include deleted | [default to false]
 **page** | **int32** | Page index.  Default is 1. | [default to 1]
 **pageSize** | **int32** | The maximum number of items per page.  Default is 100. | [default to 100]
 **offset** | **int32** | Number of items to skip.  Default is 0. | [default to 0]
 **limit** | **int32** | Number of items to return.  Default is 100. | [default to 100]
 **order** | [**SortOrder**](SortOrder.md) | The order direction. | 
 **orderBy** | [**GrantOrderBy**](GrantOrderBy.md) | The order by field. | 

### Return type

[**GrantV2PaginatedResponse**](GrantV2PaginatedResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubjectEntitlements

> []Entitlement ListSubjectEntitlements(ctx, subjectIdOrKey).IncludeDeleted(includeDeleted).Execute()

List subject entitlements



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
	subjectIdOrKey := "subjectIdOrKey_example" // string | 
	includeDeleted := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.ListSubjectEntitlements(context.Background(), subjectIdOrKey).IncludeDeleted(includeDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.ListSubjectEntitlements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSubjectEntitlements`: []Entitlement
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.ListSubjectEntitlements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectIdOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSubjectEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeDeleted** | **bool** |  | [default to false]

### Return type

[**[]Entitlement**](Entitlement.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OverrideCustomerEntitlementV2

> EntitlementV2 OverrideCustomerEntitlementV2(ctx, customerIdOrKey, entitlementIdOrFeatureKey).EntitlementV2CreateInputs(entitlementV2CreateInputs).Execute()

Override customer entitlement



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
	customerIdOrKey := *openapiclient.NewULIDOrExternalKey() // ULIDOrExternalKey | 
	entitlementIdOrFeatureKey := *openapiclient.NewULIDOrExternalKey() // ULIDOrExternalKey | 
	entitlementV2CreateInputs := openapiclient.EntitlementV2CreateInputs{EntitlementBooleanCreateInputs: openapiclient.NewEntitlementBooleanCreateInputs("Type_example")} // EntitlementV2CreateInputs | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.OverrideCustomerEntitlementV2(context.Background(), customerIdOrKey, entitlementIdOrFeatureKey).EntitlementV2CreateInputs(entitlementV2CreateInputs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.OverrideCustomerEntitlementV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OverrideCustomerEntitlementV2`: EntitlementV2
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.OverrideCustomerEntitlementV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerIdOrKey** | [**ULIDOrExternalKey**](.md) |  | 
**entitlementIdOrFeatureKey** | [**ULIDOrExternalKey**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOverrideCustomerEntitlementV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **entitlementV2CreateInputs** | [**EntitlementV2CreateInputs**](EntitlementV2CreateInputs.md) |  | 

### Return type

[**EntitlementV2**](EntitlementV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OverrideEntitlement

> Entitlement OverrideEntitlement(ctx, subjectIdOrKey, entitlementIdOrFeatureKey).EntitlementCreateInputs(entitlementCreateInputs).Execute()

Override subject entitlement



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
	subjectIdOrKey := "subjectIdOrKey_example" // string | 
	entitlementIdOrFeatureKey := "entitlementIdOrFeatureKey_example" // string | 
	entitlementCreateInputs := openapiclient.EntitlementCreateInputs{EntitlementBooleanCreateInputs: openapiclient.NewEntitlementBooleanCreateInputs("Type_example")} // EntitlementCreateInputs | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.OverrideEntitlement(context.Background(), subjectIdOrKey, entitlementIdOrFeatureKey).EntitlementCreateInputs(entitlementCreateInputs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.OverrideEntitlement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OverrideEntitlement`: Entitlement
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.OverrideEntitlement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectIdOrKey** | **string** |  | 
**entitlementIdOrFeatureKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOverrideEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **entitlementCreateInputs** | [**EntitlementCreateInputs**](EntitlementCreateInputs.md) |  | 

### Return type

[**Entitlement**](Entitlement.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetCustomerEntitlementUsageV2

> ResetCustomerEntitlementUsageV2(ctx, customerIdOrKey, entitlementIdOrFeatureKey).ResetEntitlementUsageInput(resetEntitlementUsageInput).Execute()

Reset customer entitlement



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
	customerIdOrKey := *openapiclient.NewULIDOrExternalKey() // ULIDOrExternalKey | 
	entitlementIdOrFeatureKey := "entitlementIdOrFeatureKey_example" // string | 
	resetEntitlementUsageInput := *openapiclient.NewResetEntitlementUsageInput() // ResetEntitlementUsageInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EntitlementsAPI.ResetCustomerEntitlementUsageV2(context.Background(), customerIdOrKey, entitlementIdOrFeatureKey).ResetEntitlementUsageInput(resetEntitlementUsageInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.ResetCustomerEntitlementUsageV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerIdOrKey** | [**ULIDOrExternalKey**](.md) |  | 
**entitlementIdOrFeatureKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetCustomerEntitlementUsageV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resetEntitlementUsageInput** | [**ResetEntitlementUsageInput**](ResetEntitlementUsageInput.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetEntitlementUsage

> ResetEntitlementUsage(ctx, subjectIdOrKey, entitlementId).ResetEntitlementUsageInput(resetEntitlementUsageInput).Execute()

Reset subject entitlement



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
	subjectIdOrKey := "subjectIdOrKey_example" // string | 
	entitlementId := "entitlementId_example" // string | 
	resetEntitlementUsageInput := *openapiclient.NewResetEntitlementUsageInput() // ResetEntitlementUsageInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EntitlementsAPI.ResetEntitlementUsage(context.Background(), subjectIdOrKey, entitlementId).ResetEntitlementUsageInput(resetEntitlementUsageInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.ResetEntitlementUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectIdOrKey** | **string** |  | 
**entitlementId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetEntitlementUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resetEntitlementUsageInput** | [**ResetEntitlementUsageInput**](ResetEntitlementUsageInput.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VoidGrantDelete

> VoidGrantDelete(ctx, grantId).Execute()

Void grant



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
	grantId := "grantId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EntitlementsAPI.VoidGrantDelete(context.Background(), grantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.VoidGrantDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**grantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVoidGrantDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

