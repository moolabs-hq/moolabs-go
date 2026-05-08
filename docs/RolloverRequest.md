# RolloverRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Tenant identifier | 
**SubscriptionId** | **string** | Subscription identifier | 
**PeriodEnd** | Pointer to **NullableTime** |  | [optional] 
**GraceWindowSeconds** | Pointer to **int32** | Grace window in seconds | [optional] [default to 3600]

## Methods

### NewRolloverRequest

`func NewRolloverRequest(tenantId string, subscriptionId string, ) *RolloverRequest`

NewRolloverRequest instantiates a new RolloverRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolloverRequestWithDefaults

`func NewRolloverRequestWithDefaults() *RolloverRequest`

NewRolloverRequestWithDefaults instantiates a new RolloverRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *RolloverRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *RolloverRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *RolloverRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetSubscriptionId

`func (o *RolloverRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *RolloverRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *RolloverRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetPeriodEnd

`func (o *RolloverRequest) GetPeriodEnd() time.Time`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *RolloverRequest) GetPeriodEndOk() (*time.Time, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *RolloverRequest) SetPeriodEnd(v time.Time)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *RolloverRequest) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### SetPeriodEndNil

`func (o *RolloverRequest) SetPeriodEndNil(b bool)`

 SetPeriodEndNil sets the value for PeriodEnd to be an explicit nil

### UnsetPeriodEnd
`func (o *RolloverRequest) UnsetPeriodEnd()`

UnsetPeriodEnd ensures that no value is present for PeriodEnd, not even an explicit nil
### GetGraceWindowSeconds

`func (o *RolloverRequest) GetGraceWindowSeconds() int32`

GetGraceWindowSeconds returns the GraceWindowSeconds field if non-nil, zero value otherwise.

### GetGraceWindowSecondsOk

`func (o *RolloverRequest) GetGraceWindowSecondsOk() (*int32, bool)`

GetGraceWindowSecondsOk returns a tuple with the GraceWindowSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraceWindowSeconds

`func (o *RolloverRequest) SetGraceWindowSeconds(v int32)`

SetGraceWindowSeconds sets GraceWindowSeconds field to given value.

### HasGraceWindowSeconds

`func (o *RolloverRequest) HasGraceWindowSeconds() bool`

HasGraceWindowSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


