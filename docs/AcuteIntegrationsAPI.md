# \AcuteIntegrationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BillingImport**](AcuteIntegrationsAPI.md#BillingImport) | **Post** /api/v1/cost/integrations/billing/import | Manual Tier 5 aggregate billing import
[**ConfigureHelicone**](AcuteIntegrationsAPI.md#ConfigureHelicone) | **Post** /api/v1/cost/integrations/helicone | Configure Helicone pull connector
[**ConfigureLangfuse**](AcuteIntegrationsAPI.md#ConfigureLangfuse) | **Post** /api/v1/cost/integrations/langfuse | Configure Langfuse pull connector
[**HeliconeWebhook**](AcuteIntegrationsAPI.md#HeliconeWebhook) | **Post** /api/v1/cost/integrations/helicone/webhook | Receive Helicone webhook push events
[**IngestOtel**](AcuteIntegrationsAPI.md#IngestOtel) | **Post** /api/v1/cost/ingest/otel | OTLP/HTTP span ingestion (Tier 2)
[**ListConnectors**](AcuteIntegrationsAPI.md#ListConnectors) | **Get** /api/v1/cost/integrations/connectors | List connector health
[**ListProviders**](AcuteIntegrationsAPI.md#ListProviders) | **Get** /api/v1/cost/integrations/providers | List providers configured in rate catalog
[**RotateKeyApi**](AcuteIntegrationsAPI.md#RotateKeyApi) | **Put** /api/v1/cost/integrations/{connector_type}/rotate-key | Rotate connector API key (24h grace period for previous key)



## BillingImport

> BillingImportResponse BillingImport(ctx).BillingImportRequest(billingImportRequest).XTenantID(xTenantID).Execute()

Manual Tier 5 aggregate billing import



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
	billingImportRequest := *openapiclient.NewBillingImportRequest("Provider_example", time.Now(), time.Now(), *openapiclient.NewTotalCost()) // BillingImportRequest | 
	xTenantID := "xTenantID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AcuteIntegrationsAPI.BillingImport(context.Background()).BillingImportRequest(billingImportRequest).XTenantID(xTenantID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AcuteIntegrationsAPI.BillingImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingImport`: BillingImportResponse
	fmt.Fprintf(os.Stdout, "Response from `AcuteIntegrationsAPI.BillingImport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBillingImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **billingImportRequest** | [**BillingImportRequest**](BillingImportRequest.md) |  | 
 **xTenantID** | **string** |  | 

### Return type

[**BillingImportResponse**](BillingImportResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigureHelicone

> ConnectorConfigResponse ConfigureHelicone(ctx).HeliconeConfigRequest(heliconeConfigRequest).XTenantID(xTenantID).Execute()

Configure Helicone pull connector



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
	heliconeConfigRequest := *openapiclient.NewHeliconeConfigRequest("ApiKey_example") // HeliconeConfigRequest | 
	xTenantID := "xTenantID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AcuteIntegrationsAPI.ConfigureHelicone(context.Background()).HeliconeConfigRequest(heliconeConfigRequest).XTenantID(xTenantID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AcuteIntegrationsAPI.ConfigureHelicone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfigureHelicone`: ConnectorConfigResponse
	fmt.Fprintf(os.Stdout, "Response from `AcuteIntegrationsAPI.ConfigureHelicone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConfigureHeliconeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **heliconeConfigRequest** | [**HeliconeConfigRequest**](HeliconeConfigRequest.md) |  | 
 **xTenantID** | **string** |  | 

### Return type

[**ConnectorConfigResponse**](ConnectorConfigResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigureLangfuse

> ConnectorConfigResponse ConfigureLangfuse(ctx).LangfuseConfigRequest(langfuseConfigRequest).XTenantID(xTenantID).Execute()

Configure Langfuse pull connector



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
	langfuseConfigRequest := *openapiclient.NewLangfuseConfigRequest("PublicKey_example", "SecretKey_example") // LangfuseConfigRequest | 
	xTenantID := "xTenantID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AcuteIntegrationsAPI.ConfigureLangfuse(context.Background()).LangfuseConfigRequest(langfuseConfigRequest).XTenantID(xTenantID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AcuteIntegrationsAPI.ConfigureLangfuse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfigureLangfuse`: ConnectorConfigResponse
	fmt.Fprintf(os.Stdout, "Response from `AcuteIntegrationsAPI.ConfigureLangfuse`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConfigureLangfuseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **langfuseConfigRequest** | [**LangfuseConfigRequest**](LangfuseConfigRequest.md) |  | 
 **xTenantID** | **string** |  | 

### Return type

[**ConnectorConfigResponse**](ConnectorConfigResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeliconeWebhook

> OTLPIngestResponse HeliconeWebhook(ctx).HeliconeSignature(heliconeSignature).XTenantID(xTenantID).Execute()

Receive Helicone webhook push events



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
	heliconeSignature := "heliconeSignature_example" // string |  (optional)
	xTenantID := "xTenantID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AcuteIntegrationsAPI.HeliconeWebhook(context.Background()).HeliconeSignature(heliconeSignature).XTenantID(xTenantID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AcuteIntegrationsAPI.HeliconeWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HeliconeWebhook`: OTLPIngestResponse
	fmt.Fprintf(os.Stdout, "Response from `AcuteIntegrationsAPI.HeliconeWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHeliconeWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **heliconeSignature** | **string** |  | 
 **xTenantID** | **string** |  | 

### Return type

[**OTLPIngestResponse**](OTLPIngestResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IngestOtel

> OTLPIngestResponse IngestOtel(ctx).RequestBody(requestBody).XTenantID(xTenantID).Execute()

OTLP/HTTP span ingestion (Tier 2)



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
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 
	xTenantID := "xTenantID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AcuteIntegrationsAPI.IngestOtel(context.Background()).RequestBody(requestBody).XTenantID(xTenantID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AcuteIntegrationsAPI.IngestOtel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IngestOtel`: OTLPIngestResponse
	fmt.Fprintf(os.Stdout, "Response from `AcuteIntegrationsAPI.IngestOtel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIngestOtelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]interface{}** |  | 
 **xTenantID** | **string** |  | 

### Return type

[**OTLPIngestResponse**](OTLPIngestResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectors

> []ConnectorHealthResponse ListConnectors(ctx).XTenantID(xTenantID).Execute()

List connector health



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
	xTenantID := "xTenantID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AcuteIntegrationsAPI.ListConnectors(context.Background()).XTenantID(xTenantID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AcuteIntegrationsAPI.ListConnectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListConnectors`: []ConnectorHealthResponse
	fmt.Fprintf(os.Stdout, "Response from `AcuteIntegrationsAPI.ListConnectors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListConnectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTenantID** | **string** |  | 

### Return type

[**[]ConnectorHealthResponse**](ConnectorHealthResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProviders

> []ProviderSummary ListProviders(ctx).XTenantID(xTenantID).Execute()

List providers configured in rate catalog



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
	xTenantID := "xTenantID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AcuteIntegrationsAPI.ListProviders(context.Background()).XTenantID(xTenantID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AcuteIntegrationsAPI.ListProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProviders`: []ProviderSummary
	fmt.Fprintf(os.Stdout, "Response from `AcuteIntegrationsAPI.ListProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTenantID** | **string** |  | 

### Return type

[**[]ProviderSummary**](ProviderSummary.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RotateKeyApi

> RotateKeyResponse RotateKeyApi(ctx, connectorType).RotateKeyRequest(rotateKeyRequest).XTenantID(xTenantID).Execute()

Rotate connector API key (24h grace period for previous key)



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
	connectorType := "connectorType_example" // string | 
	rotateKeyRequest := *openapiclient.NewRotateKeyRequest("NewApiKey_example") // RotateKeyRequest | 
	xTenantID := "xTenantID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AcuteIntegrationsAPI.RotateKeyApi(context.Background(), connectorType).RotateKeyRequest(rotateKeyRequest).XTenantID(xTenantID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AcuteIntegrationsAPI.RotateKeyApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RotateKeyApi`: RotateKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `AcuteIntegrationsAPI.RotateKeyApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRotateKeyApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rotateKeyRequest** | [**RotateKeyRequest**](RotateKeyRequest.md) |  | 
 **xTenantID** | **string** |  | 

### Return type

[**RotateKeyResponse**](RotateKeyResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

