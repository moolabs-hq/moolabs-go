# EntitlementGrantCreateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float64** | The amount to grant. Should be a positive number. | 
**Priority** | Pointer to **int32** | The priority of the grant. Grants with higher priority are applied first. Priority is a positive decimal numbers. With lower numbers indicating higher importance. For example, a priority of 1 is more urgent than a priority of 2. When there are several grants available for the same subject, the system selects the grant with the highest priority. In cases where grants share the same priority level, the grant closest to its expiration will be used first. In the case of two grants have identical priorities and expiration dates, the system will use the grant that was created first. | [optional] 
**EffectiveAt** | **time.Time** | Effective date for grants and anchor for recurring grants. Provided value will be ceiled to metering windowSize (minute). | 
**Expiration** | [**ExpirationPeriod**](ExpirationPeriod.md) | The grant expiration definition | 
**MaxRolloverAmount** | Pointer to **float64** | Grants are rolled over at reset, after which they can have a different balance compared to what they had before the reset. Balance after the reset is calculated as: Balance_After_Reset &#x3D; MIN(MaxRolloverAmount, MAX(Balance_Before_Reset, MinRolloverAmount)) | [optional] [default to 0]
**MinRolloverAmount** | Pointer to **float64** | Grants are rolled over at reset, after which they can have a different balance compared to what they had before the reset. Balance after the reset is calculated as: Balance_After_Reset &#x3D; MIN(MaxRolloverAmount, MAX(Balance_Before_Reset, MinRolloverAmount)) | [optional] [default to 0]
**Metadata** | Pointer to **map[string]string** | The grant metadata. | [optional] 
**Recurrence** | Pointer to [**RecurringPeriodCreateInput**](RecurringPeriodCreateInput.md) | The subject of the grant. | [optional] 

## Methods

### NewEntitlementGrantCreateInput

`func NewEntitlementGrantCreateInput(amount float64, effectiveAt time.Time, expiration ExpirationPeriod, ) *EntitlementGrantCreateInput`

NewEntitlementGrantCreateInput instantiates a new EntitlementGrantCreateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementGrantCreateInputWithDefaults

`func NewEntitlementGrantCreateInputWithDefaults() *EntitlementGrantCreateInput`

NewEntitlementGrantCreateInputWithDefaults instantiates a new EntitlementGrantCreateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *EntitlementGrantCreateInput) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *EntitlementGrantCreateInput) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *EntitlementGrantCreateInput) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetPriority

`func (o *EntitlementGrantCreateInput) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *EntitlementGrantCreateInput) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *EntitlementGrantCreateInput) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *EntitlementGrantCreateInput) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetEffectiveAt

`func (o *EntitlementGrantCreateInput) GetEffectiveAt() time.Time`

GetEffectiveAt returns the EffectiveAt field if non-nil, zero value otherwise.

### GetEffectiveAtOk

`func (o *EntitlementGrantCreateInput) GetEffectiveAtOk() (*time.Time, bool)`

GetEffectiveAtOk returns a tuple with the EffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAt

`func (o *EntitlementGrantCreateInput) SetEffectiveAt(v time.Time)`

SetEffectiveAt sets EffectiveAt field to given value.


### GetExpiration

`func (o *EntitlementGrantCreateInput) GetExpiration() ExpirationPeriod`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *EntitlementGrantCreateInput) GetExpirationOk() (*ExpirationPeriod, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *EntitlementGrantCreateInput) SetExpiration(v ExpirationPeriod)`

SetExpiration sets Expiration field to given value.


### GetMaxRolloverAmount

`func (o *EntitlementGrantCreateInput) GetMaxRolloverAmount() float64`

GetMaxRolloverAmount returns the MaxRolloverAmount field if non-nil, zero value otherwise.

### GetMaxRolloverAmountOk

`func (o *EntitlementGrantCreateInput) GetMaxRolloverAmountOk() (*float64, bool)`

GetMaxRolloverAmountOk returns a tuple with the MaxRolloverAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRolloverAmount

`func (o *EntitlementGrantCreateInput) SetMaxRolloverAmount(v float64)`

SetMaxRolloverAmount sets MaxRolloverAmount field to given value.

### HasMaxRolloverAmount

`func (o *EntitlementGrantCreateInput) HasMaxRolloverAmount() bool`

HasMaxRolloverAmount returns a boolean if a field has been set.

### GetMinRolloverAmount

`func (o *EntitlementGrantCreateInput) GetMinRolloverAmount() float64`

GetMinRolloverAmount returns the MinRolloverAmount field if non-nil, zero value otherwise.

### GetMinRolloverAmountOk

`func (o *EntitlementGrantCreateInput) GetMinRolloverAmountOk() (*float64, bool)`

GetMinRolloverAmountOk returns a tuple with the MinRolloverAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRolloverAmount

`func (o *EntitlementGrantCreateInput) SetMinRolloverAmount(v float64)`

SetMinRolloverAmount sets MinRolloverAmount field to given value.

### HasMinRolloverAmount

`func (o *EntitlementGrantCreateInput) HasMinRolloverAmount() bool`

HasMinRolloverAmount returns a boolean if a field has been set.

### GetMetadata

`func (o *EntitlementGrantCreateInput) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EntitlementGrantCreateInput) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EntitlementGrantCreateInput) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EntitlementGrantCreateInput) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRecurrence

`func (o *EntitlementGrantCreateInput) GetRecurrence() RecurringPeriodCreateInput`

GetRecurrence returns the Recurrence field if non-nil, zero value otherwise.

### GetRecurrenceOk

`func (o *EntitlementGrantCreateInput) GetRecurrenceOk() (*RecurringPeriodCreateInput, bool)`

GetRecurrenceOk returns a tuple with the Recurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrence

`func (o *EntitlementGrantCreateInput) SetRecurrence(v RecurringPeriodCreateInput)`

SetRecurrence sets Recurrence field to given value.

### HasRecurrence

`func (o *EntitlementGrantCreateInput) HasRecurrence() bool`

HasRecurrence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


