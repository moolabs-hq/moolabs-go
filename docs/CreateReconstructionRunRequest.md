# CreateReconstructionRunRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Tenant identifier | 
**PoolId** | **string** | Pool identifier | 
**RootWalletId** | **string** | Root wallet ID | 
**FromEffectiveAt** | **time.Time** | Start of effective time window | 
**ToEffectiveAt** | **time.Time** | End of effective time window | 
**TriggerEventId** | **string** | Trigger event ID (LATE event) | 
**LateThresholdSeconds** | Pointer to **int32** | LATE event threshold (seconds) | [optional] [default to 3600]

## Methods

### NewCreateReconstructionRunRequest

`func NewCreateReconstructionRunRequest(tenantId string, poolId string, rootWalletId string, fromEffectiveAt time.Time, toEffectiveAt time.Time, triggerEventId string, ) *CreateReconstructionRunRequest`

NewCreateReconstructionRunRequest instantiates a new CreateReconstructionRunRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReconstructionRunRequestWithDefaults

`func NewCreateReconstructionRunRequestWithDefaults() *CreateReconstructionRunRequest`

NewCreateReconstructionRunRequestWithDefaults instantiates a new CreateReconstructionRunRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *CreateReconstructionRunRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CreateReconstructionRunRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CreateReconstructionRunRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetPoolId

`func (o *CreateReconstructionRunRequest) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *CreateReconstructionRunRequest) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *CreateReconstructionRunRequest) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.


### GetRootWalletId

`func (o *CreateReconstructionRunRequest) GetRootWalletId() string`

GetRootWalletId returns the RootWalletId field if non-nil, zero value otherwise.

### GetRootWalletIdOk

`func (o *CreateReconstructionRunRequest) GetRootWalletIdOk() (*string, bool)`

GetRootWalletIdOk returns a tuple with the RootWalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootWalletId

`func (o *CreateReconstructionRunRequest) SetRootWalletId(v string)`

SetRootWalletId sets RootWalletId field to given value.


### GetFromEffectiveAt

`func (o *CreateReconstructionRunRequest) GetFromEffectiveAt() time.Time`

GetFromEffectiveAt returns the FromEffectiveAt field if non-nil, zero value otherwise.

### GetFromEffectiveAtOk

`func (o *CreateReconstructionRunRequest) GetFromEffectiveAtOk() (*time.Time, bool)`

GetFromEffectiveAtOk returns a tuple with the FromEffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEffectiveAt

`func (o *CreateReconstructionRunRequest) SetFromEffectiveAt(v time.Time)`

SetFromEffectiveAt sets FromEffectiveAt field to given value.


### GetToEffectiveAt

`func (o *CreateReconstructionRunRequest) GetToEffectiveAt() time.Time`

GetToEffectiveAt returns the ToEffectiveAt field if non-nil, zero value otherwise.

### GetToEffectiveAtOk

`func (o *CreateReconstructionRunRequest) GetToEffectiveAtOk() (*time.Time, bool)`

GetToEffectiveAtOk returns a tuple with the ToEffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToEffectiveAt

`func (o *CreateReconstructionRunRequest) SetToEffectiveAt(v time.Time)`

SetToEffectiveAt sets ToEffectiveAt field to given value.


### GetTriggerEventId

`func (o *CreateReconstructionRunRequest) GetTriggerEventId() string`

GetTriggerEventId returns the TriggerEventId field if non-nil, zero value otherwise.

### GetTriggerEventIdOk

`func (o *CreateReconstructionRunRequest) GetTriggerEventIdOk() (*string, bool)`

GetTriggerEventIdOk returns a tuple with the TriggerEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerEventId

`func (o *CreateReconstructionRunRequest) SetTriggerEventId(v string)`

SetTriggerEventId sets TriggerEventId field to given value.


### GetLateThresholdSeconds

`func (o *CreateReconstructionRunRequest) GetLateThresholdSeconds() int32`

GetLateThresholdSeconds returns the LateThresholdSeconds field if non-nil, zero value otherwise.

### GetLateThresholdSecondsOk

`func (o *CreateReconstructionRunRequest) GetLateThresholdSecondsOk() (*int32, bool)`

GetLateThresholdSecondsOk returns a tuple with the LateThresholdSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLateThresholdSeconds

`func (o *CreateReconstructionRunRequest) SetLateThresholdSeconds(v int32)`

SetLateThresholdSeconds sets LateThresholdSeconds field to given value.

### HasLateThresholdSeconds

`func (o *CreateReconstructionRunRequest) HasLateThresholdSeconds() bool`

HasLateThresholdSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


