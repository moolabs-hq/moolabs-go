# DeriveRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Tenant UUID | 
**FeatureKey** | **string** |  | 
**PeriodStart** | **time.Time** |  | 
**PeriodEnd** | **time.Time** |  | 

## Methods

### NewDeriveRequest

`func NewDeriveRequest(tenantId string, featureKey string, periodStart time.Time, periodEnd time.Time, ) *DeriveRequest`

NewDeriveRequest instantiates a new DeriveRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeriveRequestWithDefaults

`func NewDeriveRequestWithDefaults() *DeriveRequest`

NewDeriveRequestWithDefaults instantiates a new DeriveRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *DeriveRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DeriveRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DeriveRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetFeatureKey

`func (o *DeriveRequest) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *DeriveRequest) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *DeriveRequest) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.


### GetPeriodStart

`func (o *DeriveRequest) GetPeriodStart() time.Time`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *DeriveRequest) GetPeriodStartOk() (*time.Time, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *DeriveRequest) SetPeriodStart(v time.Time)`

SetPeriodStart sets PeriodStart field to given value.


### GetPeriodEnd

`func (o *DeriveRequest) GetPeriodEnd() time.Time`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *DeriveRequest) GetPeriodEndOk() (*time.Time, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *DeriveRequest) SetPeriodEnd(v time.Time)`

SetPeriodEnd sets PeriodEnd field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


