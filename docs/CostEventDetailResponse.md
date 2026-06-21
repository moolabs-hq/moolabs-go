# CostEventDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**IdempotencyKey** | **string** |  | 
**RequestId** | Pointer to **string** |  | [optional] 
**CustomerId** | **string** |  | 
**FeatureId** | Pointer to **string** |  | [optional] 
**Provider** | **string** |  | 
**ModelRequested** | **string** |  | 
**ModelResponded** | **string** |  | 
**ObservedTotalCost** | **string** |  | 
**Currency** | **string** |  | 
**ReportingTotalCost** | **string** |  | 
**ReportingCurrency** | **string** |  | 
**Status** | **string** |  | 
**Tags** | **map[string]interface{}** |  | 
**EventTimestamp** | **time.Time** |  | 
**CreatedAt** | **time.Time** |  | 
**TraceId** | **string** |  | 
**SpanId** | **string** |  | 
**LatencyMs** | **int32** |  | 
**LineItems** | **[]map[string]interface{}** |  | 
**ReconciliationStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewCostEventDetailResponse

`func NewCostEventDetailResponse(id string, tenantId string, idempotencyKey string, customerId string, provider string, modelRequested string, modelResponded string, observedTotalCost string, currency string, reportingTotalCost string, reportingCurrency string, status string, tags map[string]interface{}, eventTimestamp time.Time, createdAt time.Time, traceId string, spanId string, latencyMs int32, lineItems []map[string]interface{}, ) *CostEventDetailResponse`

NewCostEventDetailResponse instantiates a new CostEventDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostEventDetailResponseWithDefaults

`func NewCostEventDetailResponseWithDefaults() *CostEventDetailResponse`

NewCostEventDetailResponseWithDefaults instantiates a new CostEventDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CostEventDetailResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CostEventDetailResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CostEventDetailResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *CostEventDetailResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CostEventDetailResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CostEventDetailResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetIdempotencyKey

`func (o *CostEventDetailResponse) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *CostEventDetailResponse) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *CostEventDetailResponse) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.


### GetRequestId

`func (o *CostEventDetailResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CostEventDetailResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CostEventDetailResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *CostEventDetailResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetCustomerId

`func (o *CostEventDetailResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CostEventDetailResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CostEventDetailResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetFeatureId

`func (o *CostEventDetailResponse) GetFeatureId() string`

GetFeatureId returns the FeatureId field if non-nil, zero value otherwise.

### GetFeatureIdOk

`func (o *CostEventDetailResponse) GetFeatureIdOk() (*string, bool)`

GetFeatureIdOk returns a tuple with the FeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureId

`func (o *CostEventDetailResponse) SetFeatureId(v string)`

SetFeatureId sets FeatureId field to given value.

### HasFeatureId

`func (o *CostEventDetailResponse) HasFeatureId() bool`

HasFeatureId returns a boolean if a field has been set.

### GetProvider

`func (o *CostEventDetailResponse) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CostEventDetailResponse) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CostEventDetailResponse) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetModelRequested

`func (o *CostEventDetailResponse) GetModelRequested() string`

GetModelRequested returns the ModelRequested field if non-nil, zero value otherwise.

### GetModelRequestedOk

`func (o *CostEventDetailResponse) GetModelRequestedOk() (*string, bool)`

GetModelRequestedOk returns a tuple with the ModelRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelRequested

`func (o *CostEventDetailResponse) SetModelRequested(v string)`

SetModelRequested sets ModelRequested field to given value.


### GetModelResponded

`func (o *CostEventDetailResponse) GetModelResponded() string`

GetModelResponded returns the ModelResponded field if non-nil, zero value otherwise.

### GetModelRespondedOk

`func (o *CostEventDetailResponse) GetModelRespondedOk() (*string, bool)`

GetModelRespondedOk returns a tuple with the ModelResponded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelResponded

`func (o *CostEventDetailResponse) SetModelResponded(v string)`

SetModelResponded sets ModelResponded field to given value.


### GetObservedTotalCost

`func (o *CostEventDetailResponse) GetObservedTotalCost() string`

GetObservedTotalCost returns the ObservedTotalCost field if non-nil, zero value otherwise.

### GetObservedTotalCostOk

`func (o *CostEventDetailResponse) GetObservedTotalCostOk() (*string, bool)`

GetObservedTotalCostOk returns a tuple with the ObservedTotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedTotalCost

`func (o *CostEventDetailResponse) SetObservedTotalCost(v string)`

SetObservedTotalCost sets ObservedTotalCost field to given value.


### GetCurrency

`func (o *CostEventDetailResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CostEventDetailResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CostEventDetailResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetReportingTotalCost

`func (o *CostEventDetailResponse) GetReportingTotalCost() string`

GetReportingTotalCost returns the ReportingTotalCost field if non-nil, zero value otherwise.

### GetReportingTotalCostOk

`func (o *CostEventDetailResponse) GetReportingTotalCostOk() (*string, bool)`

GetReportingTotalCostOk returns a tuple with the ReportingTotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingTotalCost

`func (o *CostEventDetailResponse) SetReportingTotalCost(v string)`

SetReportingTotalCost sets ReportingTotalCost field to given value.


### GetReportingCurrency

`func (o *CostEventDetailResponse) GetReportingCurrency() string`

GetReportingCurrency returns the ReportingCurrency field if non-nil, zero value otherwise.

### GetReportingCurrencyOk

`func (o *CostEventDetailResponse) GetReportingCurrencyOk() (*string, bool)`

GetReportingCurrencyOk returns a tuple with the ReportingCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingCurrency

`func (o *CostEventDetailResponse) SetReportingCurrency(v string)`

SetReportingCurrency sets ReportingCurrency field to given value.


### GetStatus

`func (o *CostEventDetailResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CostEventDetailResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CostEventDetailResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTags

`func (o *CostEventDetailResponse) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CostEventDetailResponse) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CostEventDetailResponse) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.


### GetEventTimestamp

`func (o *CostEventDetailResponse) GetEventTimestamp() time.Time`

GetEventTimestamp returns the EventTimestamp field if non-nil, zero value otherwise.

### GetEventTimestampOk

`func (o *CostEventDetailResponse) GetEventTimestampOk() (*time.Time, bool)`

GetEventTimestampOk returns a tuple with the EventTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTimestamp

`func (o *CostEventDetailResponse) SetEventTimestamp(v time.Time)`

SetEventTimestamp sets EventTimestamp field to given value.


### GetCreatedAt

`func (o *CostEventDetailResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CostEventDetailResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CostEventDetailResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetTraceId

`func (o *CostEventDetailResponse) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *CostEventDetailResponse) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *CostEventDetailResponse) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.


### GetSpanId

`func (o *CostEventDetailResponse) GetSpanId() string`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *CostEventDetailResponse) GetSpanIdOk() (*string, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *CostEventDetailResponse) SetSpanId(v string)`

SetSpanId sets SpanId field to given value.


### GetLatencyMs

`func (o *CostEventDetailResponse) GetLatencyMs() int32`

GetLatencyMs returns the LatencyMs field if non-nil, zero value otherwise.

### GetLatencyMsOk

`func (o *CostEventDetailResponse) GetLatencyMsOk() (*int32, bool)`

GetLatencyMsOk returns a tuple with the LatencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyMs

`func (o *CostEventDetailResponse) SetLatencyMs(v int32)`

SetLatencyMs sets LatencyMs field to given value.


### GetLineItems

`func (o *CostEventDetailResponse) GetLineItems() []map[string]interface{}`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *CostEventDetailResponse) GetLineItemsOk() (*[]map[string]interface{}, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *CostEventDetailResponse) SetLineItems(v []map[string]interface{})`

SetLineItems sets LineItems field to given value.


### GetReconciliationStatus

`func (o *CostEventDetailResponse) GetReconciliationStatus() string`

GetReconciliationStatus returns the ReconciliationStatus field if non-nil, zero value otherwise.

### GetReconciliationStatusOk

`func (o *CostEventDetailResponse) GetReconciliationStatusOk() (*string, bool)`

GetReconciliationStatusOk returns a tuple with the ReconciliationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconciliationStatus

`func (o *CostEventDetailResponse) SetReconciliationStatus(v string)`

SetReconciliationStatus sets ReconciliationStatus field to given value.

### HasReconciliationStatus

`func (o *CostEventDetailResponse) HasReconciliationStatus() bool`

HasReconciliationStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


