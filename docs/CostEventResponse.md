# CostEventResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**IdempotencyKey** | **string** |  | 
**RequestId** | Pointer to **string** |  | [optional] 
**CustomerId** | Pointer to **string** |  | [optional] 
**FeatureId** | Pointer to **string** |  | [optional] 
**Provider** | **string** |  | 
**ModelRequested** | Pointer to **string** |  | [optional] 
**ModelResponded** | Pointer to **string** |  | [optional] 
**ObservedTotalCost** | **string** |  | 
**Currency** | **string** |  | 
**ReportingTotalCost** | **string** |  | 
**ReportingCurrency** | **string** |  | 
**Status** | **string** |  | 
**Tags** | Pointer to **map[string]interface{}** |  | [optional] 
**EventTimestamp** | **time.Time** |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewCostEventResponse

`func NewCostEventResponse(id string, tenantId string, idempotencyKey string, provider string, observedTotalCost string, currency string, reportingTotalCost string, reportingCurrency string, status string, eventTimestamp time.Time, createdAt time.Time, ) *CostEventResponse`

NewCostEventResponse instantiates a new CostEventResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostEventResponseWithDefaults

`func NewCostEventResponseWithDefaults() *CostEventResponse`

NewCostEventResponseWithDefaults instantiates a new CostEventResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CostEventResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CostEventResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CostEventResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *CostEventResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CostEventResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CostEventResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetIdempotencyKey

`func (o *CostEventResponse) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *CostEventResponse) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *CostEventResponse) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.


### GetRequestId

`func (o *CostEventResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CostEventResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CostEventResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *CostEventResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetCustomerId

`func (o *CostEventResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CostEventResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CostEventResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *CostEventResponse) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetFeatureId

`func (o *CostEventResponse) GetFeatureId() string`

GetFeatureId returns the FeatureId field if non-nil, zero value otherwise.

### GetFeatureIdOk

`func (o *CostEventResponse) GetFeatureIdOk() (*string, bool)`

GetFeatureIdOk returns a tuple with the FeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureId

`func (o *CostEventResponse) SetFeatureId(v string)`

SetFeatureId sets FeatureId field to given value.

### HasFeatureId

`func (o *CostEventResponse) HasFeatureId() bool`

HasFeatureId returns a boolean if a field has been set.

### GetProvider

`func (o *CostEventResponse) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CostEventResponse) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CostEventResponse) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetModelRequested

`func (o *CostEventResponse) GetModelRequested() string`

GetModelRequested returns the ModelRequested field if non-nil, zero value otherwise.

### GetModelRequestedOk

`func (o *CostEventResponse) GetModelRequestedOk() (*string, bool)`

GetModelRequestedOk returns a tuple with the ModelRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelRequested

`func (o *CostEventResponse) SetModelRequested(v string)`

SetModelRequested sets ModelRequested field to given value.

### HasModelRequested

`func (o *CostEventResponse) HasModelRequested() bool`

HasModelRequested returns a boolean if a field has been set.

### GetModelResponded

`func (o *CostEventResponse) GetModelResponded() string`

GetModelResponded returns the ModelResponded field if non-nil, zero value otherwise.

### GetModelRespondedOk

`func (o *CostEventResponse) GetModelRespondedOk() (*string, bool)`

GetModelRespondedOk returns a tuple with the ModelResponded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelResponded

`func (o *CostEventResponse) SetModelResponded(v string)`

SetModelResponded sets ModelResponded field to given value.

### HasModelResponded

`func (o *CostEventResponse) HasModelResponded() bool`

HasModelResponded returns a boolean if a field has been set.

### GetObservedTotalCost

`func (o *CostEventResponse) GetObservedTotalCost() string`

GetObservedTotalCost returns the ObservedTotalCost field if non-nil, zero value otherwise.

### GetObservedTotalCostOk

`func (o *CostEventResponse) GetObservedTotalCostOk() (*string, bool)`

GetObservedTotalCostOk returns a tuple with the ObservedTotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedTotalCost

`func (o *CostEventResponse) SetObservedTotalCost(v string)`

SetObservedTotalCost sets ObservedTotalCost field to given value.


### GetCurrency

`func (o *CostEventResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CostEventResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CostEventResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetReportingTotalCost

`func (o *CostEventResponse) GetReportingTotalCost() string`

GetReportingTotalCost returns the ReportingTotalCost field if non-nil, zero value otherwise.

### GetReportingTotalCostOk

`func (o *CostEventResponse) GetReportingTotalCostOk() (*string, bool)`

GetReportingTotalCostOk returns a tuple with the ReportingTotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingTotalCost

`func (o *CostEventResponse) SetReportingTotalCost(v string)`

SetReportingTotalCost sets ReportingTotalCost field to given value.


### GetReportingCurrency

`func (o *CostEventResponse) GetReportingCurrency() string`

GetReportingCurrency returns the ReportingCurrency field if non-nil, zero value otherwise.

### GetReportingCurrencyOk

`func (o *CostEventResponse) GetReportingCurrencyOk() (*string, bool)`

GetReportingCurrencyOk returns a tuple with the ReportingCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingCurrency

`func (o *CostEventResponse) SetReportingCurrency(v string)`

SetReportingCurrency sets ReportingCurrency field to given value.


### GetStatus

`func (o *CostEventResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CostEventResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CostEventResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTags

`func (o *CostEventResponse) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CostEventResponse) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CostEventResponse) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *CostEventResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEventTimestamp

`func (o *CostEventResponse) GetEventTimestamp() time.Time`

GetEventTimestamp returns the EventTimestamp field if non-nil, zero value otherwise.

### GetEventTimestampOk

`func (o *CostEventResponse) GetEventTimestampOk() (*time.Time, bool)`

GetEventTimestampOk returns a tuple with the EventTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTimestamp

`func (o *CostEventResponse) SetEventTimestamp(v time.Time)`

SetEventTimestamp sets EventTimestamp field to given value.


### GetCreatedAt

`func (o *CostEventResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CostEventResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CostEventResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


