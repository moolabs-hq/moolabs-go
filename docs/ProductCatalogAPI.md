# \ProductCatalogAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveAddon**](ProductCatalogAPI.md#ArchiveAddon) | **Post** /api/v1/addons/{addonId}/archive | Archive add-on version
[**ArchivePlan**](ProductCatalogAPI.md#ArchivePlan) | **Post** /api/v1/plans/{planId}/archive | Archive plan version
[**CreateAddon**](ProductCatalogAPI.md#CreateAddon) | **Post** /api/v1/addons | Create an add-on
[**CreateFeature**](ProductCatalogAPI.md#CreateFeature) | **Post** /api/v1/features | Create feature
[**CreatePlan**](ProductCatalogAPI.md#CreatePlan) | **Post** /api/v1/plans | Create a plan
[**CreatePlanAddon**](ProductCatalogAPI.md#CreatePlanAddon) | **Post** /api/v1/plans/{planId}/addons | Create new add-on assignment for plan
[**DeleteAddon**](ProductCatalogAPI.md#DeleteAddon) | **Delete** /api/v1/addons/{addonId} | Delete add-on
[**DeleteFeature**](ProductCatalogAPI.md#DeleteFeature) | **Delete** /api/v1/features/{featureId} | Delete feature
[**DeletePlan**](ProductCatalogAPI.md#DeletePlan) | **Delete** /api/v1/plans/{planId} | Delete plan
[**DeletePlanAddon**](ProductCatalogAPI.md#DeletePlanAddon) | **Delete** /api/v1/plans/{planId}/addons/{planAddonId} | Delete add-on assignment for plan
[**GetAddon**](ProductCatalogAPI.md#GetAddon) | **Get** /api/v1/addons/{addonId} | Get add-on
[**GetFeature**](ProductCatalogAPI.md#GetFeature) | **Get** /api/v1/features/{featureId} | Get feature
[**GetPlan**](ProductCatalogAPI.md#GetPlan) | **Get** /api/v1/plans/{planId} | Get plan
[**GetPlanAddon**](ProductCatalogAPI.md#GetPlanAddon) | **Get** /api/v1/plans/{planId}/addons/{planAddonId} | Get add-on assignment for plan
[**ListAddons**](ProductCatalogAPI.md#ListAddons) | **Get** /api/v1/addons | List add-ons
[**ListFeatures**](ProductCatalogAPI.md#ListFeatures) | **Get** /api/v1/features | List features
[**ListPlanAddons**](ProductCatalogAPI.md#ListPlanAddons) | **Get** /api/v1/plans/{planId}/addons | List all available add-ons for plan
[**ListPlans**](ProductCatalogAPI.md#ListPlans) | **Get** /api/v1/plans | List plans
[**NextPlan**](ProductCatalogAPI.md#NextPlan) | **Post** /api/v1/plans/{planIdOrKey}/next | New draft plan
[**PublishAddon**](ProductCatalogAPI.md#PublishAddon) | **Post** /api/v1/addons/{addonId}/publish | Publish add-on
[**PublishPlan**](ProductCatalogAPI.md#PublishPlan) | **Post** /api/v1/plans/{planId}/publish | Publish plan
[**UpdateAddon**](ProductCatalogAPI.md#UpdateAddon) | **Put** /api/v1/addons/{addonId} | Update add-on
[**UpdatePlan**](ProductCatalogAPI.md#UpdatePlan) | **Put** /api/v1/plans/{planId} | Update a plan
[**UpdatePlanAddon**](ProductCatalogAPI.md#UpdatePlanAddon) | **Put** /api/v1/plans/{planId}/addons/{planAddonId} | Update add-on assignment for plan



## ArchiveAddon

> Addon ArchiveAddon(ctx, addonId).Execute()

Archive add-on version



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
	addonId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductCatalogAPI.ArchiveAddon(context.Background(), addonId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductCatalogAPI.ArchiveAddon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArchiveAddon`: Addon
	fmt.Fprintf(os.Stdout, "Response from `ProductCatalogAPI.ArchiveAddon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchiveAddonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Addon**](Addon.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArchivePlan

> Plan ArchivePlan(ctx, planId).Execute()

Archive plan version



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
	planId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductCatalogAPI.ArchivePlan(context.Background(), planId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductCatalogAPI.ArchivePlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArchivePlan`: Plan
	fmt.Fprintf(os.Stdout, "Response from `ProductCatalogAPI.ArchivePlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchivePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Plan**](Plan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAddon

> Addon CreateAddon(ctx).AddonCreate(addonCreate).Execute()

Create an add-on



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
	addonCreate := *openapiclient.NewAddonCreate("Name_example", "Key_example", openapiclient.AddonInstanceType("single"), "USD", []openapiclient.RateCard{openapiclient.RateCard{RateCardFlatFee: openapiclient.NewRateCardFlatFee("Type_example", "Key_example", "Name_example", "BillingCadence_example", *openapiclient.NewFlatPriceWithPaymentTerm("Type_example", "Amount_example"))}}) // AddonCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductCatalogAPI.CreateAddon(context.Background()).AddonCreate(addonCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductCatalogAPI.CreateAddon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAddon`: Addon
	fmt.Fprintf(os.Stdout, "Response from `ProductCatalogAPI.CreateAddon`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAddonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addonCreate** | [**AddonCreate**](AddonCreate.md) |  | 

### Return type

[**Addon**](Addon.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFeature

> Feature CreateFeature(ctx).FeatureCreateInputs(featureCreateInputs).Execute()

Create feature



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
	featureCreateInputs := *openapiclient.NewFeatureCreateInputs("Key_example", "Name_example") // FeatureCreateInputs | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductCatalogAPI.CreateFeature(context.Background()).FeatureCreateInputs(featureCreateInputs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductCatalogAPI.CreateFeature``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFeature`: Feature
	fmt.Fprintf(os.Stdout, "Response from `ProductCatalogAPI.CreateFeature`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **featureCreateInputs** | [**FeatureCreateInputs**](FeatureCreateInputs.md) |  | 

### Return type

[**Feature**](Feature.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePlan

> Plan CreatePlan(ctx).PlanCreate(planCreate).Execute()

Create a plan



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
	planCreate := *openapiclient.NewPlanCreate("Name_example", "Key_example", "USD", "P1M", []openapiclient.PlanPhase{*openapiclient.NewPlanPhase("Key_example", "Name_example", "P1Y", []openapiclient.RateCard{openapiclient.RateCard{RateCardFlatFee: openapiclient.NewRateCardFlatFee("Type_example", "Key_example", "Name_example", "BillingCadence_example", *openapiclient.NewFlatPriceWithPaymentTerm("Type_example", "Amount_example"))}})}) // PlanCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductCatalogAPI.CreatePlan(context.Background()).PlanCreate(planCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductCatalogAPI.CreatePlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePlan`: Plan
	fmt.Fprintf(os.Stdout, "Response from `ProductCatalogAPI.CreatePlan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **planCreate** | [**PlanCreate**](PlanCreate.md) |  | 

### Return type

[**Plan**](Plan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePlanAddon

> PlanAddon CreatePlanAddon(ctx, planId).PlanAddonCreate(planAddonCreate).Execute()

Create new add-on assignment for plan



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
	planId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 
	planAddonCreate := *openapiclient.NewPlanAddonCreate("FromPlanPhase_example", "01G65Z755AFWAKHE12NY0CQ9FH") // PlanAddonCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductCatalogAPI.CreatePlanAddon(context.Background(), planId).PlanAddonCreate(planAddonCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductCatalogAPI.CreatePlanAddon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePlanAddon`: PlanAddon
	fmt.Fprintf(os.Stdout, "Response from `ProductCatalogAPI.CreatePlanAddon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePlanAddonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **planAddonCreate** | [**PlanAddonCreate**](PlanAddonCreate.md) |  | 

### Return type

[**PlanAddon**](PlanAddon.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAddon

> DeleteAddon(ctx, addonId).Execute()

Delete add-on



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
	addonId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProductCatalogAPI.DeleteAddon(context.Background(), addonId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductCatalogAPI.DeleteAddon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAddonRequest struct via the builder pattern


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


## DeleteFeature

> DeleteFeature(ctx, featureId).Execute()

Delete feature



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
	featureId := "featureId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProductCatalogAPI.DeleteFeature(context.Background(), featureId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductCatalogAPI.DeleteFeature``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFeatureRequest struct via the builder pattern


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


## DeletePlan

> DeletePlan(ctx, planId).Execute()

Delete plan



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
	planId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProductCatalogAPI.DeletePlan(context.Background(), planId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductCatalogAPI.DeletePlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePlanRequest struct via the builder pattern


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


## DeletePlanAddon

> DeletePlanAddon(ctx, planId, planAddonId).Execute()

Delete add-on assignment for plan



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
	planId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 
	planAddonId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProductCatalogAPI.DeletePlanAddon(context.Background(), planId, planAddonId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductCatalogAPI.DeletePlanAddon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **string** |  | 
**planAddonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePlanAddonRequest struct via the builder pattern


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


## GetAddon

> Addon GetAddon(ctx, addonId).IncludeLatest(includeLatest).Execute()

Get add-on



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
	addonId := "addonId_example" // string | 
	includeLatest := true // bool | Include latest version of the add-on instead of the version in active state.  Usage: `?includeLatest=true` (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductCatalogAPI.GetAddon(context.Background(), addonId).IncludeLatest(includeLatest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductCatalogAPI.GetAddon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAddon`: Addon
	fmt.Fprintf(os.Stdout, "Response from `ProductCatalogAPI.GetAddon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeLatest** | **bool** | Include latest version of the add-on instead of the version in active state.  Usage: &#x60;?includeLatest&#x3D;true&#x60; | [default to false]

### Return type

[**Addon**](Addon.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeature

> Feature GetFeature(ctx, featureId).Execute()

Get feature



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
	featureId := "featureId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductCatalogAPI.GetFeature(context.Background(), featureId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductCatalogAPI.GetFeature``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFeature`: Feature
	fmt.Fprintf(os.Stdout, "Response from `ProductCatalogAPI.GetFeature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Feature**](Feature.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlan

> Plan GetPlan(ctx, planId).IncludeLatest(includeLatest).Execute()

Get plan



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
	planId := "planId_example" // string | 
	includeLatest := true // bool | Include latest version of the Plan instead of the version in active state.  Usage: `?includeLatest=true` (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductCatalogAPI.GetPlan(context.Background(), planId).IncludeLatest(includeLatest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductCatalogAPI.GetPlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlan`: Plan
	fmt.Fprintf(os.Stdout, "Response from `ProductCatalogAPI.GetPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeLatest** | **bool** | Include latest version of the Plan instead of the version in active state.  Usage: &#x60;?includeLatest&#x3D;true&#x60; | [default to false]

### Return type

[**Plan**](Plan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlanAddon

> PlanAddon GetPlanAddon(ctx, planId, planAddonId).Execute()

Get add-on assignment for plan



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
	planId := "planId_example" // string | 
	planAddonId := "planAddonId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductCatalogAPI.GetPlanAddon(context.Background(), planId, planAddonId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductCatalogAPI.GetPlanAddon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlanAddon`: PlanAddon
	fmt.Fprintf(os.Stdout, "Response from `ProductCatalogAPI.GetPlanAddon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **string** |  | 
**planAddonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanAddonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PlanAddon**](PlanAddon.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAddons

> AddonPaginatedResponse ListAddons(ctx).IncludeDeleted(includeDeleted).Id(id).Key(key).KeyVersion(keyVersion).Status(status).Currency(currency).Page(page).PageSize(pageSize).Order(order).OrderBy(orderBy).Execute()

List add-ons



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
	includeDeleted := true // bool | Include deleted add-ons in response.  Usage: `?includeDeleted=true` (optional) (default to false)
	id := []string{"01G65Z755AFWAKHE12NY0CQ9FH"} // []string | Filter by addon.id attribute (optional)
	key := []string{"Inner_example"} // []string | Filter by addon.key attribute (optional)
	keyVersion := map[string][]int32{"key": map[string][]int32{"key": []int32{int32(123)}}} // map[string][]int32 | Filter by addon.key and addon.version attributes (optional)
	status := []openapiclient.AddonStatus{openapiclient.AddonStatus("draft")} // []AddonStatus | Only return add-ons with the given status.  Usage: - `?status=active`: return only the currently active add-ons - `?status=draft`: return only the draft add-ons - `?status=archived`: return only the archived add-ons (optional)
	currency := []string{"USD"} // []string | Filter by addon.currency attribute (optional)
	page := int32(56) // int32 | Page index.  Default is 1. (optional) (default to 1)
	pageSize := int32(56) // int32 | The maximum number of items per page.  Default is 100. (optional) (default to 100)
	order := openapiclient.SortOrder("ASC") // SortOrder | The order direction. (optional)
	orderBy := openapiclient.AddonOrderBy("id") // AddonOrderBy | The order by field. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductCatalogAPI.ListAddons(context.Background()).IncludeDeleted(includeDeleted).Id(id).Key(key).KeyVersion(keyVersion).Status(status).Currency(currency).Page(page).PageSize(pageSize).Order(order).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductCatalogAPI.ListAddons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAddons`: AddonPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductCatalogAPI.ListAddons`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAddonsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeDeleted** | **bool** | Include deleted add-ons in response.  Usage: &#x60;?includeDeleted&#x3D;true&#x60; | [default to false]
 **id** | **[]string** | Filter by addon.id attribute | 
 **key** | **[]string** | Filter by addon.key attribute | 
 **keyVersion** | [**map[string]map[string][]int32**](map.md) | Filter by addon.key and addon.version attributes | 
 **status** | [**[]AddonStatus**](AddonStatus.md) | Only return add-ons with the given status.  Usage: - &#x60;?status&#x3D;active&#x60;: return only the currently active add-ons - &#x60;?status&#x3D;draft&#x60;: return only the draft add-ons - &#x60;?status&#x3D;archived&#x60;: return only the archived add-ons | 
 **currency** | **[]string** | Filter by addon.currency attribute | 
 **page** | **int32** | Page index.  Default is 1. | [default to 1]
 **pageSize** | **int32** | The maximum number of items per page.  Default is 100. | [default to 100]
 **order** | [**SortOrder**](SortOrder.md) | The order direction. | 
 **orderBy** | [**AddonOrderBy**](AddonOrderBy.md) | The order by field. | 

### Return type

[**AddonPaginatedResponse**](AddonPaginatedResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFeatures

> ListFeaturesResult ListFeatures(ctx).MeterSlug(meterSlug).IncludeArchived(includeArchived).Page(page).PageSize(pageSize).Offset(offset).Limit(limit).Order(order).OrderBy(orderBy).Execute()

List features



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
	meterSlug := []string{"Inner_example"} // []string | Filter by meterSlug (optional)
	includeArchived := true // bool | Include archived features in response. (optional) (default to false)
	page := int32(56) // int32 | Page index.  Default is 1. (optional) (default to 1)
	pageSize := int32(56) // int32 | The maximum number of items per page.  Default is 100. (optional) (default to 100)
	offset := int32(56) // int32 | Number of items to skip.  Default is 0. (optional) (default to 0)
	limit := int32(56) // int32 | Number of items to return.  Default is 100. (optional) (default to 100)
	order := openapiclient.SortOrder("ASC") // SortOrder | The order direction. (optional)
	orderBy := openapiclient.FeatureOrderBy("id") // FeatureOrderBy | The order by field. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductCatalogAPI.ListFeatures(context.Background()).MeterSlug(meterSlug).IncludeArchived(includeArchived).Page(page).PageSize(pageSize).Offset(offset).Limit(limit).Order(order).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductCatalogAPI.ListFeatures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFeatures`: ListFeaturesResult
	fmt.Fprintf(os.Stdout, "Response from `ProductCatalogAPI.ListFeatures`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFeaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **meterSlug** | **[]string** | Filter by meterSlug | 
 **includeArchived** | **bool** | Include archived features in response. | [default to false]
 **page** | **int32** | Page index.  Default is 1. | [default to 1]
 **pageSize** | **int32** | The maximum number of items per page.  Default is 100. | [default to 100]
 **offset** | **int32** | Number of items to skip.  Default is 0. | [default to 0]
 **limit** | **int32** | Number of items to return.  Default is 100. | [default to 100]
 **order** | [**SortOrder**](SortOrder.md) | The order direction. | 
 **orderBy** | [**FeatureOrderBy**](FeatureOrderBy.md) | The order by field. | 

### Return type

[**ListFeaturesResult**](ListFeaturesResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPlanAddons

> PlanAddonPaginatedResponse ListPlanAddons(ctx, planId).IncludeDeleted(includeDeleted).Id(id).Key(key).KeyVersion(keyVersion).Page(page).PageSize(pageSize).Order(order).OrderBy(orderBy).Execute()

List all available add-ons for plan



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
	planId := "planId_example" // string | 
	includeDeleted := true // bool | Include deleted plan add-on assignments.  Usage: `?includeDeleted=true` (optional) (default to false)
	id := []string{"01G65Z755AFWAKHE12NY0CQ9FH"} // []string | Filter by addon.id attribute. (optional)
	key := []string{"Inner_example"} // []string | Filter by addon.key attribute. (optional)
	keyVersion := map[string][]int32{"key": map[string][]int32{"key": []int32{int32(123)}}} // map[string][]int32 | Filter by addon.key and addon.version attributes. (optional)
	page := int32(56) // int32 | Page index.  Default is 1. (optional) (default to 1)
	pageSize := int32(56) // int32 | The maximum number of items per page.  Default is 100. (optional) (default to 100)
	order := openapiclient.SortOrder("ASC") // SortOrder | The order direction. (optional)
	orderBy := openapiclient.PlanAddonOrderBy("id") // PlanAddonOrderBy | The order by field. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductCatalogAPI.ListPlanAddons(context.Background(), planId).IncludeDeleted(includeDeleted).Id(id).Key(key).KeyVersion(keyVersion).Page(page).PageSize(pageSize).Order(order).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductCatalogAPI.ListPlanAddons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPlanAddons`: PlanAddonPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductCatalogAPI.ListPlanAddons`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPlanAddonsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeDeleted** | **bool** | Include deleted plan add-on assignments.  Usage: &#x60;?includeDeleted&#x3D;true&#x60; | [default to false]
 **id** | **[]string** | Filter by addon.id attribute. | 
 **key** | **[]string** | Filter by addon.key attribute. | 
 **keyVersion** | [**map[string]map[string][]int32**](map.md) | Filter by addon.key and addon.version attributes. | 
 **page** | **int32** | Page index.  Default is 1. | [default to 1]
 **pageSize** | **int32** | The maximum number of items per page.  Default is 100. | [default to 100]
 **order** | [**SortOrder**](SortOrder.md) | The order direction. | 
 **orderBy** | [**PlanAddonOrderBy**](PlanAddonOrderBy.md) | The order by field. | 

### Return type

[**PlanAddonPaginatedResponse**](PlanAddonPaginatedResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPlans

> PlanPaginatedResponse ListPlans(ctx).IncludeDeleted(includeDeleted).Id(id).Key(key).KeyVersion(keyVersion).Status(status).Currency(currency).Page(page).PageSize(pageSize).Order(order).OrderBy(orderBy).Execute()

List plans



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
	includeDeleted := true // bool | Include deleted plans in response.  Usage: `?includeDeleted=true` (optional) (default to false)
	id := []string{"01G65Z755AFWAKHE12NY0CQ9FH"} // []string | Filter by plan.id attribute (optional)
	key := []string{"Inner_example"} // []string | Filter by plan.key attribute (optional)
	keyVersion := map[string][]int32{"key": map[string][]int32{"key": []int32{int32(123)}}} // map[string][]int32 | Filter by plan.key and plan.version attributes (optional)
	status := []openapiclient.PlanStatus{openapiclient.PlanStatus("draft")} // []PlanStatus | Only return plans with the given status.  Usage: - `?status=active`: return only the currently active plan - `?status=draft`: return only the draft plan - `?status=archived`: return only the archived plans (optional)
	currency := []string{"USD"} // []string | Filter by plan.currency attribute (optional)
	page := int32(56) // int32 | Page index.  Default is 1. (optional) (default to 1)
	pageSize := int32(56) // int32 | The maximum number of items per page.  Default is 100. (optional) (default to 100)
	order := openapiclient.SortOrder("ASC") // SortOrder | The order direction. (optional)
	orderBy := openapiclient.PlanOrderBy("id") // PlanOrderBy | The order by field. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductCatalogAPI.ListPlans(context.Background()).IncludeDeleted(includeDeleted).Id(id).Key(key).KeyVersion(keyVersion).Status(status).Currency(currency).Page(page).PageSize(pageSize).Order(order).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductCatalogAPI.ListPlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPlans`: PlanPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductCatalogAPI.ListPlans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeDeleted** | **bool** | Include deleted plans in response.  Usage: &#x60;?includeDeleted&#x3D;true&#x60; | [default to false]
 **id** | **[]string** | Filter by plan.id attribute | 
 **key** | **[]string** | Filter by plan.key attribute | 
 **keyVersion** | [**map[string]map[string][]int32**](map.md) | Filter by plan.key and plan.version attributes | 
 **status** | [**[]PlanStatus**](PlanStatus.md) | Only return plans with the given status.  Usage: - &#x60;?status&#x3D;active&#x60;: return only the currently active plan - &#x60;?status&#x3D;draft&#x60;: return only the draft plan - &#x60;?status&#x3D;archived&#x60;: return only the archived plans | 
 **currency** | **[]string** | Filter by plan.currency attribute | 
 **page** | **int32** | Page index.  Default is 1. | [default to 1]
 **pageSize** | **int32** | The maximum number of items per page.  Default is 100. | [default to 100]
 **order** | [**SortOrder**](SortOrder.md) | The order direction. | 
 **orderBy** | [**PlanOrderBy**](PlanOrderBy.md) | The order by field. | 

### Return type

[**PlanPaginatedResponse**](PlanPaginatedResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NextPlan

> Plan NextPlan(ctx, planIdOrKey).Execute()

New draft plan



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
	planIdOrKey := "planIdOrKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductCatalogAPI.NextPlan(context.Background(), planIdOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductCatalogAPI.NextPlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NextPlan`: Plan
	fmt.Fprintf(os.Stdout, "Response from `ProductCatalogAPI.NextPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planIdOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNextPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Plan**](Plan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublishAddon

> Addon PublishAddon(ctx, addonId).Execute()

Publish add-on



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
	addonId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductCatalogAPI.PublishAddon(context.Background(), addonId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductCatalogAPI.PublishAddon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublishAddon`: Addon
	fmt.Fprintf(os.Stdout, "Response from `ProductCatalogAPI.PublishAddon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublishAddonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Addon**](Addon.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublishPlan

> Plan PublishPlan(ctx, planId).Execute()

Publish plan



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
	planId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductCatalogAPI.PublishPlan(context.Background(), planId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductCatalogAPI.PublishPlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublishPlan`: Plan
	fmt.Fprintf(os.Stdout, "Response from `ProductCatalogAPI.PublishPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublishPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Plan**](Plan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAddon

> Addon UpdateAddon(ctx, addonId).AddonReplaceUpdate(addonReplaceUpdate).Execute()

Update add-on



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
	addonId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 
	addonReplaceUpdate := *openapiclient.NewAddonReplaceUpdate("Name_example", openapiclient.AddonInstanceType("single"), []openapiclient.RateCard{openapiclient.RateCard{RateCardFlatFee: openapiclient.NewRateCardFlatFee("Type_example", "Key_example", "Name_example", "BillingCadence_example", *openapiclient.NewFlatPriceWithPaymentTerm("Type_example", "Amount_example"))}}) // AddonReplaceUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductCatalogAPI.UpdateAddon(context.Background(), addonId).AddonReplaceUpdate(addonReplaceUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductCatalogAPI.UpdateAddon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAddon`: Addon
	fmt.Fprintf(os.Stdout, "Response from `ProductCatalogAPI.UpdateAddon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAddonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addonReplaceUpdate** | [**AddonReplaceUpdate**](AddonReplaceUpdate.md) |  | 

### Return type

[**Addon**](Addon.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePlan

> Plan UpdatePlan(ctx, planId).PlanReplaceUpdate(planReplaceUpdate).Execute()

Update a plan



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
	planId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 
	planReplaceUpdate := *openapiclient.NewPlanReplaceUpdate("Name_example", "P1M", []openapiclient.PlanPhase{*openapiclient.NewPlanPhase("Key_example", "Name_example", "P1Y", []openapiclient.RateCard{openapiclient.RateCard{RateCardFlatFee: openapiclient.NewRateCardFlatFee("Type_example", "Key_example", "Name_example", "BillingCadence_example", *openapiclient.NewFlatPriceWithPaymentTerm("Type_example", "Amount_example"))}})}) // PlanReplaceUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductCatalogAPI.UpdatePlan(context.Background(), planId).PlanReplaceUpdate(planReplaceUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductCatalogAPI.UpdatePlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePlan`: Plan
	fmt.Fprintf(os.Stdout, "Response from `ProductCatalogAPI.UpdatePlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **planReplaceUpdate** | [**PlanReplaceUpdate**](PlanReplaceUpdate.md) |  | 

### Return type

[**Plan**](Plan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePlanAddon

> PlanAddon UpdatePlanAddon(ctx, planId, planAddonId).PlanAddonReplaceUpdate(planAddonReplaceUpdate).Execute()

Update add-on assignment for plan



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
	planId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 
	planAddonId := "01G65Z755AFWAKHE12NY0CQ9FH" // string | 
	planAddonReplaceUpdate := *openapiclient.NewPlanAddonReplaceUpdate("FromPlanPhase_example") // PlanAddonReplaceUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductCatalogAPI.UpdatePlanAddon(context.Background(), planId, planAddonId).PlanAddonReplaceUpdate(planAddonReplaceUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductCatalogAPI.UpdatePlanAddon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePlanAddon`: PlanAddon
	fmt.Fprintf(os.Stdout, "Response from `ProductCatalogAPI.UpdatePlanAddon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **string** |  | 
**planAddonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePlanAddonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **planAddonReplaceUpdate** | [**PlanAddonReplaceUpdate**](PlanAddonReplaceUpdate.md) |  | 

### Return type

[**PlanAddon**](PlanAddon.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

