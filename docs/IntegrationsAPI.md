# \IntegrationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomerUpsertV1**](IntegrationsAPI.md#CustomerUpsertV1) | **Post** /v1/integrations/netsuite-sim/events/customer | Customer Upsert
[**EnqueueArcTransactionV1Integrations**](IntegrationsAPI.md#EnqueueArcTransactionV1Integrations) | **Post** /v1/integrations/netsuite-sim/arc-transactions | Enqueue Arc Transaction
[**InvoiceUpsertV1**](IntegrationsAPI.md#InvoiceUpsertV1) | **Post** /v1/integrations/netsuite-sim/events/invoice | Invoice Upsert
[**OpenmeterWebhook**](IntegrationsAPI.md#OpenmeterWebhook) | **Post** /v1/integrations/moometer/webhooks/moometer | Openmeter Webhook
[**OpenmeterWebhookBatch**](IntegrationsAPI.md#OpenmeterWebhookBatch) | **Post** /v1/integrations/moometer/webhooks/moometer/batch | Openmeter Webhook Batch
[**ProcessPendingV1Integrations**](IntegrationsAPI.md#ProcessPendingV1Integrations) | **Post** /v1/integrations/netsuite-sim/process-pending | Process Pending
[**ReadinessV1**](IntegrationsAPI.md#ReadinessV1) | **Get** /v1/integrations/netsuite-sim/readiness | Readiness



## CustomerUpsertV1

> interface{} CustomerUpsertV1(ctx).CustomerUpsertRequest(customerUpsertRequest).XIdempotencyKey(xIdempotencyKey).Execute()

Customer Upsert

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/moolabs/moolabs-go"
)

func main() {
	customerUpsertRequest := *openapiclient.NewCustomerUpsertRequest("NetsuiteCustomerId_example", "MoometerCustomerId_example", "CompanyName_example") // CustomerUpsertRequest | 
	xIdempotencyKey := "xIdempotencyKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.CustomerUpsertV1(context.Background()).CustomerUpsertRequest(customerUpsertRequest).XIdempotencyKey(xIdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.CustomerUpsertV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerUpsertV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.CustomerUpsertV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerUpsertV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerUpsertRequest** | [**CustomerUpsertRequest**](CustomerUpsertRequest.md) |  | 
 **xIdempotencyKey** | **string** |  | 

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


## EnqueueArcTransactionV1Integrations

> interface{} EnqueueArcTransactionV1Integrations(ctx).ArcTransactionRequest(arcTransactionRequest).XIdempotencyKey(xIdempotencyKey).Execute()

Enqueue Arc Transaction

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
	arcTransactionRequest := *openapiclient.NewArcTransactionRequest("ArcTxnId_example", "TxnType_example", "MoometerCustomerId_example", int32(123), "CurrencyCode_example", time.Now()) // ArcTransactionRequest | 
	xIdempotencyKey := "xIdempotencyKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.EnqueueArcTransactionV1Integrations(context.Background()).ArcTransactionRequest(arcTransactionRequest).XIdempotencyKey(xIdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.EnqueueArcTransactionV1Integrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnqueueArcTransactionV1Integrations`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.EnqueueArcTransactionV1Integrations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnqueueArcTransactionV1IntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arcTransactionRequest** | [**ArcTransactionRequest**](ArcTransactionRequest.md) |  | 
 **xIdempotencyKey** | **string** |  | 

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


## InvoiceUpsertV1

> interface{} InvoiceUpsertV1(ctx).InvoiceUpsertRequest(invoiceUpsertRequest).XIdempotencyKey(xIdempotencyKey).Execute()

Invoice Upsert

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
	invoiceUpsertRequest := *openapiclient.NewInvoiceUpsertRequest("NetsuiteInvoiceId_example", "MoometerInvoiceId_example", "NetsuiteCustomerId_example", "MoometerCustomerId_example", "InvoiceNumber_example", time.Now(), time.Now(), "CurrencyCode_example", int32(123), int32(123)) // InvoiceUpsertRequest | 
	xIdempotencyKey := "xIdempotencyKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.InvoiceUpsertV1(context.Background()).InvoiceUpsertRequest(invoiceUpsertRequest).XIdempotencyKey(xIdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.InvoiceUpsertV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceUpsertV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.InvoiceUpsertV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceUpsertV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invoiceUpsertRequest** | [**InvoiceUpsertRequest**](InvoiceUpsertRequest.md) |  | 
 **xIdempotencyKey** | **string** |  | 

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
	resp, r, err := apiClient.IntegrationsAPI.OpenmeterWebhook(context.Background()).TenantId(tenantId).PoolId(poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.OpenmeterWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpenmeterWebhook`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.OpenmeterWebhook`: %v\n", resp)
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
	resp, r, err := apiClient.IntegrationsAPI.OpenmeterWebhookBatch(context.Background()).TenantId(tenantId).PoolId(poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.OpenmeterWebhookBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpenmeterWebhookBatch`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.OpenmeterWebhookBatch`: %v\n", resp)
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


## ProcessPendingV1Integrations

> interface{} ProcessPendingV1Integrations(ctx).Limit(limit).Execute()

Process Pending

### Example

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.ProcessPendingV1Integrations(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.ProcessPendingV1Integrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProcessPendingV1Integrations`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.ProcessPendingV1Integrations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProcessPendingV1IntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 100]

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


## ReadinessV1

> interface{} ReadinessV1(ctx).Execute()

Readiness

### Example

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
	resp, r, err := apiClient.IntegrationsAPI.ReadinessV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.ReadinessV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadinessV1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.ReadinessV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReadinessV1Request struct via the builder pattern


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

