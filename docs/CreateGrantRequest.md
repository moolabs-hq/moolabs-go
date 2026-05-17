# CreateGrantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Tenant identifier | 
**PoolId** | **string** | Credit pool identifier | 
**WalletId** | **string** | Wallet identifier | 
**AmountMicros** | **int32** | Grant amount in micros | 
**Priority** | Pointer to **int32** | Grant priority (lower numbers &#x3D; higher priority, priority 0 is used before priority 100) | [optional] [default to 100]
**ValidFrom** | **time.Time** | Grant valid from timestamp | 
**ExpiresAt** | Pointer to **time.Time** | Grant expiration timestamp | [optional] 
**SourceType** | [**GrantSourceType**](GrantSourceType.md) | Grant source type | 
**SourceId** | **string** | Source identifier (idempotency key) | 
**SubscriptionId** | Pointer to **string** | Subscription identifier | [optional] 
**ScopeType** | Pointer to [**ScopeType**](ScopeType.md) | Grant scope type | [optional] 
**ScopeValues** | Pointer to **[]string** | Scope values (feature keys or meter slugs) | [optional] 
**RolloverMode** | Pointer to [**RolloverMode**](RolloverMode.md) | Rollover mode | [optional] 
**RolloverPercent** | Pointer to **int32** | Rollover percentage (0-100) | [optional] 
**RolloverCapMicros** | Pointer to **int32** | Rollover cap (micros) | [optional] 
**RolloverExpiresAfterDays** | Pointer to **int32** | Rollover expiration (days) | [optional] 
**RolloverPriorityDelta** | Pointer to **int32** | Priority delta for rolled-over grants | [optional] 
**PaidAmountUsdMicros** | Pointer to **int32** | Paid amount in USD micros | [optional] 
**LockedFxRateVersion** | Pointer to **string** | FX rate version for paid grants (required if paid_amount_usd_micros is set) | [optional] 
**LockedCreditsPerUsdMicros** | Pointer to **int32** | Locked credits per USD micros (for paid grants) | [optional] 

## Methods

### NewCreateGrantRequest

`func NewCreateGrantRequest(tenantId string, poolId string, walletId string, amountMicros int32, validFrom time.Time, sourceType GrantSourceType, sourceId string, ) *CreateGrantRequest`

NewCreateGrantRequest instantiates a new CreateGrantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGrantRequestWithDefaults

`func NewCreateGrantRequestWithDefaults() *CreateGrantRequest`

NewCreateGrantRequestWithDefaults instantiates a new CreateGrantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *CreateGrantRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CreateGrantRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CreateGrantRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetPoolId

`func (o *CreateGrantRequest) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *CreateGrantRequest) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *CreateGrantRequest) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.


### GetWalletId

`func (o *CreateGrantRequest) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *CreateGrantRequest) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *CreateGrantRequest) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetAmountMicros

`func (o *CreateGrantRequest) GetAmountMicros() int32`

GetAmountMicros returns the AmountMicros field if non-nil, zero value otherwise.

### GetAmountMicrosOk

`func (o *CreateGrantRequest) GetAmountMicrosOk() (*int32, bool)`

GetAmountMicrosOk returns a tuple with the AmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMicros

`func (o *CreateGrantRequest) SetAmountMicros(v int32)`

SetAmountMicros sets AmountMicros field to given value.


### GetPriority

`func (o *CreateGrantRequest) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CreateGrantRequest) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CreateGrantRequest) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CreateGrantRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetValidFrom

`func (o *CreateGrantRequest) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *CreateGrantRequest) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *CreateGrantRequest) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.


### GetExpiresAt

`func (o *CreateGrantRequest) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CreateGrantRequest) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CreateGrantRequest) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CreateGrantRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetSourceType

`func (o *CreateGrantRequest) GetSourceType() GrantSourceType`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *CreateGrantRequest) GetSourceTypeOk() (*GrantSourceType, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *CreateGrantRequest) SetSourceType(v GrantSourceType)`

SetSourceType sets SourceType field to given value.


### GetSourceId

`func (o *CreateGrantRequest) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *CreateGrantRequest) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *CreateGrantRequest) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetSubscriptionId

`func (o *CreateGrantRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CreateGrantRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CreateGrantRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *CreateGrantRequest) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetScopeType

`func (o *CreateGrantRequest) GetScopeType() ScopeType`

GetScopeType returns the ScopeType field if non-nil, zero value otherwise.

### GetScopeTypeOk

`func (o *CreateGrantRequest) GetScopeTypeOk() (*ScopeType, bool)`

GetScopeTypeOk returns a tuple with the ScopeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeType

`func (o *CreateGrantRequest) SetScopeType(v ScopeType)`

SetScopeType sets ScopeType field to given value.

### HasScopeType

`func (o *CreateGrantRequest) HasScopeType() bool`

HasScopeType returns a boolean if a field has been set.

### GetScopeValues

`func (o *CreateGrantRequest) GetScopeValues() []string`

GetScopeValues returns the ScopeValues field if non-nil, zero value otherwise.

### GetScopeValuesOk

`func (o *CreateGrantRequest) GetScopeValuesOk() (*[]string, bool)`

GetScopeValuesOk returns a tuple with the ScopeValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeValues

`func (o *CreateGrantRequest) SetScopeValues(v []string)`

SetScopeValues sets ScopeValues field to given value.

### HasScopeValues

`func (o *CreateGrantRequest) HasScopeValues() bool`

HasScopeValues returns a boolean if a field has been set.

### GetRolloverMode

`func (o *CreateGrantRequest) GetRolloverMode() RolloverMode`

GetRolloverMode returns the RolloverMode field if non-nil, zero value otherwise.

### GetRolloverModeOk

`func (o *CreateGrantRequest) GetRolloverModeOk() (*RolloverMode, bool)`

GetRolloverModeOk returns a tuple with the RolloverMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloverMode

`func (o *CreateGrantRequest) SetRolloverMode(v RolloverMode)`

SetRolloverMode sets RolloverMode field to given value.

### HasRolloverMode

`func (o *CreateGrantRequest) HasRolloverMode() bool`

HasRolloverMode returns a boolean if a field has been set.

### GetRolloverPercent

`func (o *CreateGrantRequest) GetRolloverPercent() int32`

GetRolloverPercent returns the RolloverPercent field if non-nil, zero value otherwise.

### GetRolloverPercentOk

`func (o *CreateGrantRequest) GetRolloverPercentOk() (*int32, bool)`

GetRolloverPercentOk returns a tuple with the RolloverPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloverPercent

`func (o *CreateGrantRequest) SetRolloverPercent(v int32)`

SetRolloverPercent sets RolloverPercent field to given value.

### HasRolloverPercent

`func (o *CreateGrantRequest) HasRolloverPercent() bool`

HasRolloverPercent returns a boolean if a field has been set.

### GetRolloverCapMicros

`func (o *CreateGrantRequest) GetRolloverCapMicros() int32`

GetRolloverCapMicros returns the RolloverCapMicros field if non-nil, zero value otherwise.

### GetRolloverCapMicrosOk

`func (o *CreateGrantRequest) GetRolloverCapMicrosOk() (*int32, bool)`

GetRolloverCapMicrosOk returns a tuple with the RolloverCapMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloverCapMicros

`func (o *CreateGrantRequest) SetRolloverCapMicros(v int32)`

SetRolloverCapMicros sets RolloverCapMicros field to given value.

### HasRolloverCapMicros

`func (o *CreateGrantRequest) HasRolloverCapMicros() bool`

HasRolloverCapMicros returns a boolean if a field has been set.

### GetRolloverExpiresAfterDays

`func (o *CreateGrantRequest) GetRolloverExpiresAfterDays() int32`

GetRolloverExpiresAfterDays returns the RolloverExpiresAfterDays field if non-nil, zero value otherwise.

### GetRolloverExpiresAfterDaysOk

`func (o *CreateGrantRequest) GetRolloverExpiresAfterDaysOk() (*int32, bool)`

GetRolloverExpiresAfterDaysOk returns a tuple with the RolloverExpiresAfterDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloverExpiresAfterDays

`func (o *CreateGrantRequest) SetRolloverExpiresAfterDays(v int32)`

SetRolloverExpiresAfterDays sets RolloverExpiresAfterDays field to given value.

### HasRolloverExpiresAfterDays

`func (o *CreateGrantRequest) HasRolloverExpiresAfterDays() bool`

HasRolloverExpiresAfterDays returns a boolean if a field has been set.

### GetRolloverPriorityDelta

`func (o *CreateGrantRequest) GetRolloverPriorityDelta() int32`

GetRolloverPriorityDelta returns the RolloverPriorityDelta field if non-nil, zero value otherwise.

### GetRolloverPriorityDeltaOk

`func (o *CreateGrantRequest) GetRolloverPriorityDeltaOk() (*int32, bool)`

GetRolloverPriorityDeltaOk returns a tuple with the RolloverPriorityDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloverPriorityDelta

`func (o *CreateGrantRequest) SetRolloverPriorityDelta(v int32)`

SetRolloverPriorityDelta sets RolloverPriorityDelta field to given value.

### HasRolloverPriorityDelta

`func (o *CreateGrantRequest) HasRolloverPriorityDelta() bool`

HasRolloverPriorityDelta returns a boolean if a field has been set.

### GetPaidAmountUsdMicros

`func (o *CreateGrantRequest) GetPaidAmountUsdMicros() int32`

GetPaidAmountUsdMicros returns the PaidAmountUsdMicros field if non-nil, zero value otherwise.

### GetPaidAmountUsdMicrosOk

`func (o *CreateGrantRequest) GetPaidAmountUsdMicrosOk() (*int32, bool)`

GetPaidAmountUsdMicrosOk returns a tuple with the PaidAmountUsdMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidAmountUsdMicros

`func (o *CreateGrantRequest) SetPaidAmountUsdMicros(v int32)`

SetPaidAmountUsdMicros sets PaidAmountUsdMicros field to given value.

### HasPaidAmountUsdMicros

`func (o *CreateGrantRequest) HasPaidAmountUsdMicros() bool`

HasPaidAmountUsdMicros returns a boolean if a field has been set.

### GetLockedFxRateVersion

`func (o *CreateGrantRequest) GetLockedFxRateVersion() string`

GetLockedFxRateVersion returns the LockedFxRateVersion field if non-nil, zero value otherwise.

### GetLockedFxRateVersionOk

`func (o *CreateGrantRequest) GetLockedFxRateVersionOk() (*string, bool)`

GetLockedFxRateVersionOk returns a tuple with the LockedFxRateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedFxRateVersion

`func (o *CreateGrantRequest) SetLockedFxRateVersion(v string)`

SetLockedFxRateVersion sets LockedFxRateVersion field to given value.

### HasLockedFxRateVersion

`func (o *CreateGrantRequest) HasLockedFxRateVersion() bool`

HasLockedFxRateVersion returns a boolean if a field has been set.

### GetLockedCreditsPerUsdMicros

`func (o *CreateGrantRequest) GetLockedCreditsPerUsdMicros() int32`

GetLockedCreditsPerUsdMicros returns the LockedCreditsPerUsdMicros field if non-nil, zero value otherwise.

### GetLockedCreditsPerUsdMicrosOk

`func (o *CreateGrantRequest) GetLockedCreditsPerUsdMicrosOk() (*int32, bool)`

GetLockedCreditsPerUsdMicrosOk returns a tuple with the LockedCreditsPerUsdMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedCreditsPerUsdMicros

`func (o *CreateGrantRequest) SetLockedCreditsPerUsdMicros(v int32)`

SetLockedCreditsPerUsdMicros sets LockedCreditsPerUsdMicros field to given value.

### HasLockedCreditsPerUsdMicros

`func (o *CreateGrantRequest) HasLockedCreditsPerUsdMicros() bool`

HasLockedCreditsPerUsdMicros returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


