# GrantResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**PoolId** | **string** |  | 
**WalletId** | **string** |  | 
**AmountMicros** | **int32** |  | 
**RemainingMicros** | **int32** |  | 
**Priority** | **int32** |  | 
**ValidFrom** | **time.Time** |  | 
**ExpiresAt** | **NullableTime** |  | 
**SourceType** | **string** |  | 
**SourceId** | **string** |  | 
**SubscriptionId** | **NullableString** |  | 
**ScopeType** | [**ScopeType**](ScopeType.md) |  | 
**RolloverMode** | [**RolloverMode**](RolloverMode.md) |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewGrantResponse

`func NewGrantResponse(id string, tenantId string, poolId string, walletId string, amountMicros int32, remainingMicros int32, priority int32, validFrom time.Time, expiresAt NullableTime, sourceType string, sourceId string, subscriptionId NullableString, scopeType ScopeType, rolloverMode RolloverMode, createdAt time.Time, updatedAt time.Time, ) *GrantResponse`

NewGrantResponse instantiates a new GrantResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantResponseWithDefaults

`func NewGrantResponseWithDefaults() *GrantResponse`

NewGrantResponseWithDefaults instantiates a new GrantResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GrantResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GrantResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GrantResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *GrantResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *GrantResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *GrantResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetPoolId

`func (o *GrantResponse) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *GrantResponse) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *GrantResponse) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.


### GetWalletId

`func (o *GrantResponse) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *GrantResponse) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *GrantResponse) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetAmountMicros

`func (o *GrantResponse) GetAmountMicros() int32`

GetAmountMicros returns the AmountMicros field if non-nil, zero value otherwise.

### GetAmountMicrosOk

`func (o *GrantResponse) GetAmountMicrosOk() (*int32, bool)`

GetAmountMicrosOk returns a tuple with the AmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMicros

`func (o *GrantResponse) SetAmountMicros(v int32)`

SetAmountMicros sets AmountMicros field to given value.


### GetRemainingMicros

`func (o *GrantResponse) GetRemainingMicros() int32`

GetRemainingMicros returns the RemainingMicros field if non-nil, zero value otherwise.

### GetRemainingMicrosOk

`func (o *GrantResponse) GetRemainingMicrosOk() (*int32, bool)`

GetRemainingMicrosOk returns a tuple with the RemainingMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingMicros

`func (o *GrantResponse) SetRemainingMicros(v int32)`

SetRemainingMicros sets RemainingMicros field to given value.


### GetPriority

`func (o *GrantResponse) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *GrantResponse) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *GrantResponse) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetValidFrom

`func (o *GrantResponse) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *GrantResponse) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *GrantResponse) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.


### GetExpiresAt

`func (o *GrantResponse) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GrantResponse) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GrantResponse) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### SetExpiresAtNil

`func (o *GrantResponse) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *GrantResponse) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetSourceType

`func (o *GrantResponse) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *GrantResponse) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *GrantResponse) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetSourceId

`func (o *GrantResponse) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *GrantResponse) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *GrantResponse) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetSubscriptionId

`func (o *GrantResponse) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *GrantResponse) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *GrantResponse) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### SetSubscriptionIdNil

`func (o *GrantResponse) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *GrantResponse) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetScopeType

`func (o *GrantResponse) GetScopeType() ScopeType`

GetScopeType returns the ScopeType field if non-nil, zero value otherwise.

### GetScopeTypeOk

`func (o *GrantResponse) GetScopeTypeOk() (*ScopeType, bool)`

GetScopeTypeOk returns a tuple with the ScopeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeType

`func (o *GrantResponse) SetScopeType(v ScopeType)`

SetScopeType sets ScopeType field to given value.


### GetRolloverMode

`func (o *GrantResponse) GetRolloverMode() RolloverMode`

GetRolloverMode returns the RolloverMode field if non-nil, zero value otherwise.

### GetRolloverModeOk

`func (o *GrantResponse) GetRolloverModeOk() (*RolloverMode, bool)`

GetRolloverModeOk returns a tuple with the RolloverMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloverMode

`func (o *GrantResponse) SetRolloverMode(v RolloverMode)`

SetRolloverMode sets RolloverMode field to given value.


### GetCreatedAt

`func (o *GrantResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GrantResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GrantResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *GrantResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GrantResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GrantResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


