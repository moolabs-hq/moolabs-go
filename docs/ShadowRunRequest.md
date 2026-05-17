# ShadowRunRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Tenant UUID | 
**PeriodStart** | **time.Time** |  | 
**PeriodEnd** | **time.Time** |  | 

## Methods

### NewShadowRunRequest

`func NewShadowRunRequest(tenantId string, periodStart time.Time, periodEnd time.Time, ) *ShadowRunRequest`

NewShadowRunRequest instantiates a new ShadowRunRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShadowRunRequestWithDefaults

`func NewShadowRunRequestWithDefaults() *ShadowRunRequest`

NewShadowRunRequestWithDefaults instantiates a new ShadowRunRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *ShadowRunRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ShadowRunRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ShadowRunRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetPeriodStart

`func (o *ShadowRunRequest) GetPeriodStart() time.Time`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *ShadowRunRequest) GetPeriodStartOk() (*time.Time, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *ShadowRunRequest) SetPeriodStart(v time.Time)`

SetPeriodStart sets PeriodStart field to given value.


### GetPeriodEnd

`func (o *ShadowRunRequest) GetPeriodEnd() time.Time`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *ShadowRunRequest) GetPeriodEndOk() (*time.Time, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *ShadowRunRequest) SetPeriodEnd(v time.Time)`

SetPeriodEnd sets PeriodEnd field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


