# CostEventIngest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**IdempotencyKey** | **string** |  | 
**TraceId** | Pointer to **string** |  | [optional] 
**RootSpanId** | Pointer to **string** |  | [optional] 
**SpanId** | Pointer to **string** |  | [optional] 
**ParentSpanId** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**CustomerId** | Pointer to **string** |  | [optional] 
**FeatureId** | Pointer to **string** |  | [optional] 
**PlanId** | Pointer to **string** |  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 
**WorkflowId** | Pointer to **string** |  | [optional] 
**WorkflowVersion** | Pointer to **string** |  | [optional] 
**Provider** | **string** |  | 
**ModelRequested** | Pointer to **string** |  | [optional] 
**ModelResponded** | Pointer to **string** |  | [optional] 
**OperationType** | Pointer to **string** |  | [optional] 
**ComponentType** | Pointer to **string** |  | [optional] 
**ObservedTotalCost** | [**ObservedTotalCost**](ObservedTotalCost.md) |  | 
**Currency** | **string** |  | 
**ReportingTotalCost** | [**ReportingTotalCost**](ReportingTotalCost.md) |  | 
**ReportingCurrency** | Pointer to **string** |  | [optional] [default to "USD"]
**FxRateToReporting** | Pointer to [**NullableFxRateToReporting**](FxRateToReporting.md) |  | [optional] 
**LatencyMs** | Pointer to **int32** |  | [optional] 
**CacheHit** | Pointer to **bool** |  | [optional] [default to false]
**Environment** | Pointer to **string** |  | [optional] [default to "production"]
**UsageEventId** | Pointer to **string** |  | [optional] 
**IngestionPath** | Pointer to **string** |  | [optional] [default to "sdk"]
**IngestionSource** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]interface{}** |  | [optional] 
**ProviderMetadata** | Pointer to **map[string]interface{}** |  | [optional] 
**OtelAttributes** | Pointer to **map[string]interface{}** |  | [optional] 
**EventTimestamp** | **time.Time** |  | 
**LineItems** | Pointer to [**[]LineItemCreate**](LineItemCreate.md) |  | [optional] 

## Methods

### NewCostEventIngest

`func NewCostEventIngest(tenantId string, idempotencyKey string, provider string, observedTotalCost ObservedTotalCost, currency string, reportingTotalCost ReportingTotalCost, eventTimestamp time.Time, ) *CostEventIngest`

NewCostEventIngest instantiates a new CostEventIngest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostEventIngestWithDefaults

`func NewCostEventIngestWithDefaults() *CostEventIngest`

NewCostEventIngestWithDefaults instantiates a new CostEventIngest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *CostEventIngest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CostEventIngest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CostEventIngest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetIdempotencyKey

`func (o *CostEventIngest) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *CostEventIngest) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *CostEventIngest) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.


### GetTraceId

`func (o *CostEventIngest) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *CostEventIngest) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *CostEventIngest) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *CostEventIngest) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### GetRootSpanId

`func (o *CostEventIngest) GetRootSpanId() string`

GetRootSpanId returns the RootSpanId field if non-nil, zero value otherwise.

### GetRootSpanIdOk

`func (o *CostEventIngest) GetRootSpanIdOk() (*string, bool)`

GetRootSpanIdOk returns a tuple with the RootSpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootSpanId

`func (o *CostEventIngest) SetRootSpanId(v string)`

SetRootSpanId sets RootSpanId field to given value.

### HasRootSpanId

`func (o *CostEventIngest) HasRootSpanId() bool`

HasRootSpanId returns a boolean if a field has been set.

### GetSpanId

`func (o *CostEventIngest) GetSpanId() string`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *CostEventIngest) GetSpanIdOk() (*string, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *CostEventIngest) SetSpanId(v string)`

SetSpanId sets SpanId field to given value.

### HasSpanId

`func (o *CostEventIngest) HasSpanId() bool`

HasSpanId returns a boolean if a field has been set.

### GetParentSpanId

`func (o *CostEventIngest) GetParentSpanId() string`

GetParentSpanId returns the ParentSpanId field if non-nil, zero value otherwise.

### GetParentSpanIdOk

`func (o *CostEventIngest) GetParentSpanIdOk() (*string, bool)`

GetParentSpanIdOk returns a tuple with the ParentSpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSpanId

`func (o *CostEventIngest) SetParentSpanId(v string)`

SetParentSpanId sets ParentSpanId field to given value.

### HasParentSpanId

`func (o *CostEventIngest) HasParentSpanId() bool`

HasParentSpanId returns a boolean if a field has been set.

### GetRequestId

`func (o *CostEventIngest) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CostEventIngest) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CostEventIngest) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *CostEventIngest) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetCustomerId

`func (o *CostEventIngest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CostEventIngest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CostEventIngest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *CostEventIngest) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetFeatureId

`func (o *CostEventIngest) GetFeatureId() string`

GetFeatureId returns the FeatureId field if non-nil, zero value otherwise.

### GetFeatureIdOk

`func (o *CostEventIngest) GetFeatureIdOk() (*string, bool)`

GetFeatureIdOk returns a tuple with the FeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureId

`func (o *CostEventIngest) SetFeatureId(v string)`

SetFeatureId sets FeatureId field to given value.

### HasFeatureId

`func (o *CostEventIngest) HasFeatureId() bool`

HasFeatureId returns a boolean if a field has been set.

### GetPlanId

`func (o *CostEventIngest) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *CostEventIngest) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *CostEventIngest) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *CostEventIngest) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *CostEventIngest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CostEventIngest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CostEventIngest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *CostEventIngest) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetWorkflowId

`func (o *CostEventIngest) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *CostEventIngest) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *CostEventIngest) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *CostEventIngest) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### GetWorkflowVersion

`func (o *CostEventIngest) GetWorkflowVersion() string`

GetWorkflowVersion returns the WorkflowVersion field if non-nil, zero value otherwise.

### GetWorkflowVersionOk

`func (o *CostEventIngest) GetWorkflowVersionOk() (*string, bool)`

GetWorkflowVersionOk returns a tuple with the WorkflowVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowVersion

`func (o *CostEventIngest) SetWorkflowVersion(v string)`

SetWorkflowVersion sets WorkflowVersion field to given value.

### HasWorkflowVersion

`func (o *CostEventIngest) HasWorkflowVersion() bool`

HasWorkflowVersion returns a boolean if a field has been set.

### GetProvider

`func (o *CostEventIngest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CostEventIngest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CostEventIngest) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetModelRequested

`func (o *CostEventIngest) GetModelRequested() string`

GetModelRequested returns the ModelRequested field if non-nil, zero value otherwise.

### GetModelRequestedOk

`func (o *CostEventIngest) GetModelRequestedOk() (*string, bool)`

GetModelRequestedOk returns a tuple with the ModelRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelRequested

`func (o *CostEventIngest) SetModelRequested(v string)`

SetModelRequested sets ModelRequested field to given value.

### HasModelRequested

`func (o *CostEventIngest) HasModelRequested() bool`

HasModelRequested returns a boolean if a field has been set.

### GetModelResponded

`func (o *CostEventIngest) GetModelResponded() string`

GetModelResponded returns the ModelResponded field if non-nil, zero value otherwise.

### GetModelRespondedOk

`func (o *CostEventIngest) GetModelRespondedOk() (*string, bool)`

GetModelRespondedOk returns a tuple with the ModelResponded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelResponded

`func (o *CostEventIngest) SetModelResponded(v string)`

SetModelResponded sets ModelResponded field to given value.

### HasModelResponded

`func (o *CostEventIngest) HasModelResponded() bool`

HasModelResponded returns a boolean if a field has been set.

### GetOperationType

`func (o *CostEventIngest) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *CostEventIngest) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *CostEventIngest) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *CostEventIngest) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### GetComponentType

`func (o *CostEventIngest) GetComponentType() string`

GetComponentType returns the ComponentType field if non-nil, zero value otherwise.

### GetComponentTypeOk

`func (o *CostEventIngest) GetComponentTypeOk() (*string, bool)`

GetComponentTypeOk returns a tuple with the ComponentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentType

`func (o *CostEventIngest) SetComponentType(v string)`

SetComponentType sets ComponentType field to given value.

### HasComponentType

`func (o *CostEventIngest) HasComponentType() bool`

HasComponentType returns a boolean if a field has been set.

### GetObservedTotalCost

`func (o *CostEventIngest) GetObservedTotalCost() ObservedTotalCost`

GetObservedTotalCost returns the ObservedTotalCost field if non-nil, zero value otherwise.

### GetObservedTotalCostOk

`func (o *CostEventIngest) GetObservedTotalCostOk() (*ObservedTotalCost, bool)`

GetObservedTotalCostOk returns a tuple with the ObservedTotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedTotalCost

`func (o *CostEventIngest) SetObservedTotalCost(v ObservedTotalCost)`

SetObservedTotalCost sets ObservedTotalCost field to given value.


### GetCurrency

`func (o *CostEventIngest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CostEventIngest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CostEventIngest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetReportingTotalCost

`func (o *CostEventIngest) GetReportingTotalCost() ReportingTotalCost`

GetReportingTotalCost returns the ReportingTotalCost field if non-nil, zero value otherwise.

### GetReportingTotalCostOk

`func (o *CostEventIngest) GetReportingTotalCostOk() (*ReportingTotalCost, bool)`

GetReportingTotalCostOk returns a tuple with the ReportingTotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingTotalCost

`func (o *CostEventIngest) SetReportingTotalCost(v ReportingTotalCost)`

SetReportingTotalCost sets ReportingTotalCost field to given value.


### GetReportingCurrency

`func (o *CostEventIngest) GetReportingCurrency() string`

GetReportingCurrency returns the ReportingCurrency field if non-nil, zero value otherwise.

### GetReportingCurrencyOk

`func (o *CostEventIngest) GetReportingCurrencyOk() (*string, bool)`

GetReportingCurrencyOk returns a tuple with the ReportingCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingCurrency

`func (o *CostEventIngest) SetReportingCurrency(v string)`

SetReportingCurrency sets ReportingCurrency field to given value.

### HasReportingCurrency

`func (o *CostEventIngest) HasReportingCurrency() bool`

HasReportingCurrency returns a boolean if a field has been set.

### GetFxRateToReporting

`func (o *CostEventIngest) GetFxRateToReporting() FxRateToReporting`

GetFxRateToReporting returns the FxRateToReporting field if non-nil, zero value otherwise.

### GetFxRateToReportingOk

`func (o *CostEventIngest) GetFxRateToReportingOk() (*FxRateToReporting, bool)`

GetFxRateToReportingOk returns a tuple with the FxRateToReporting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFxRateToReporting

`func (o *CostEventIngest) SetFxRateToReporting(v FxRateToReporting)`

SetFxRateToReporting sets FxRateToReporting field to given value.

### HasFxRateToReporting

`func (o *CostEventIngest) HasFxRateToReporting() bool`

HasFxRateToReporting returns a boolean if a field has been set.

### SetFxRateToReportingNil

`func (o *CostEventIngest) SetFxRateToReportingNil(b bool)`

 SetFxRateToReportingNil sets the value for FxRateToReporting to be an explicit nil

### UnsetFxRateToReporting
`func (o *CostEventIngest) UnsetFxRateToReporting()`

UnsetFxRateToReporting ensures that no value is present for FxRateToReporting, not even an explicit nil
### GetLatencyMs

`func (o *CostEventIngest) GetLatencyMs() int32`

GetLatencyMs returns the LatencyMs field if non-nil, zero value otherwise.

### GetLatencyMsOk

`func (o *CostEventIngest) GetLatencyMsOk() (*int32, bool)`

GetLatencyMsOk returns a tuple with the LatencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyMs

`func (o *CostEventIngest) SetLatencyMs(v int32)`

SetLatencyMs sets LatencyMs field to given value.

### HasLatencyMs

`func (o *CostEventIngest) HasLatencyMs() bool`

HasLatencyMs returns a boolean if a field has been set.

### GetCacheHit

`func (o *CostEventIngest) GetCacheHit() bool`

GetCacheHit returns the CacheHit field if non-nil, zero value otherwise.

### GetCacheHitOk

`func (o *CostEventIngest) GetCacheHitOk() (*bool, bool)`

GetCacheHitOk returns a tuple with the CacheHit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheHit

`func (o *CostEventIngest) SetCacheHit(v bool)`

SetCacheHit sets CacheHit field to given value.

### HasCacheHit

`func (o *CostEventIngest) HasCacheHit() bool`

HasCacheHit returns a boolean if a field has been set.

### GetEnvironment

`func (o *CostEventIngest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CostEventIngest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CostEventIngest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CostEventIngest) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetUsageEventId

`func (o *CostEventIngest) GetUsageEventId() string`

GetUsageEventId returns the UsageEventId field if non-nil, zero value otherwise.

### GetUsageEventIdOk

`func (o *CostEventIngest) GetUsageEventIdOk() (*string, bool)`

GetUsageEventIdOk returns a tuple with the UsageEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageEventId

`func (o *CostEventIngest) SetUsageEventId(v string)`

SetUsageEventId sets UsageEventId field to given value.

### HasUsageEventId

`func (o *CostEventIngest) HasUsageEventId() bool`

HasUsageEventId returns a boolean if a field has been set.

### GetIngestionPath

`func (o *CostEventIngest) GetIngestionPath() string`

GetIngestionPath returns the IngestionPath field if non-nil, zero value otherwise.

### GetIngestionPathOk

`func (o *CostEventIngest) GetIngestionPathOk() (*string, bool)`

GetIngestionPathOk returns a tuple with the IngestionPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestionPath

`func (o *CostEventIngest) SetIngestionPath(v string)`

SetIngestionPath sets IngestionPath field to given value.

### HasIngestionPath

`func (o *CostEventIngest) HasIngestionPath() bool`

HasIngestionPath returns a boolean if a field has been set.

### GetIngestionSource

`func (o *CostEventIngest) GetIngestionSource() string`

GetIngestionSource returns the IngestionSource field if non-nil, zero value otherwise.

### GetIngestionSourceOk

`func (o *CostEventIngest) GetIngestionSourceOk() (*string, bool)`

GetIngestionSourceOk returns a tuple with the IngestionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestionSource

`func (o *CostEventIngest) SetIngestionSource(v string)`

SetIngestionSource sets IngestionSource field to given value.

### HasIngestionSource

`func (o *CostEventIngest) HasIngestionSource() bool`

HasIngestionSource returns a boolean if a field has been set.

### GetTags

`func (o *CostEventIngest) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CostEventIngest) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CostEventIngest) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *CostEventIngest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetProviderMetadata

`func (o *CostEventIngest) GetProviderMetadata() map[string]interface{}`

GetProviderMetadata returns the ProviderMetadata field if non-nil, zero value otherwise.

### GetProviderMetadataOk

`func (o *CostEventIngest) GetProviderMetadataOk() (*map[string]interface{}, bool)`

GetProviderMetadataOk returns a tuple with the ProviderMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderMetadata

`func (o *CostEventIngest) SetProviderMetadata(v map[string]interface{})`

SetProviderMetadata sets ProviderMetadata field to given value.

### HasProviderMetadata

`func (o *CostEventIngest) HasProviderMetadata() bool`

HasProviderMetadata returns a boolean if a field has been set.

### GetOtelAttributes

`func (o *CostEventIngest) GetOtelAttributes() map[string]interface{}`

GetOtelAttributes returns the OtelAttributes field if non-nil, zero value otherwise.

### GetOtelAttributesOk

`func (o *CostEventIngest) GetOtelAttributesOk() (*map[string]interface{}, bool)`

GetOtelAttributesOk returns a tuple with the OtelAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtelAttributes

`func (o *CostEventIngest) SetOtelAttributes(v map[string]interface{})`

SetOtelAttributes sets OtelAttributes field to given value.

### HasOtelAttributes

`func (o *CostEventIngest) HasOtelAttributes() bool`

HasOtelAttributes returns a boolean if a field has been set.

### GetEventTimestamp

`func (o *CostEventIngest) GetEventTimestamp() time.Time`

GetEventTimestamp returns the EventTimestamp field if non-nil, zero value otherwise.

### GetEventTimestampOk

`func (o *CostEventIngest) GetEventTimestampOk() (*time.Time, bool)`

GetEventTimestampOk returns a tuple with the EventTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTimestamp

`func (o *CostEventIngest) SetEventTimestamp(v time.Time)`

SetEventTimestamp sets EventTimestamp field to given value.


### GetLineItems

`func (o *CostEventIngest) GetLineItems() []LineItemCreate`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *CostEventIngest) GetLineItemsOk() (*[]LineItemCreate, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *CostEventIngest) SetLineItems(v []LineItemCreate)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *CostEventIngest) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


