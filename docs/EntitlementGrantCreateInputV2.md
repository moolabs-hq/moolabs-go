# EntitlementGrantCreateInputV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float64** | The amount to grant. Should be a positive number. | 
**Priority** | Pointer to **int32** | The priority of the grant. Grants with higher priority are applied first. Priority is a positive decimal numbers. With lower numbers indicating higher importance. For example, a priority of 1 is more urgent than a priority of 2. When there are several grants available for the same subject, the system selects the grant with the highest priority. In cases where grants share the same priority level, the grant closest to its expiration will be used first. In the case of two grants have identical priorities and expiration dates, the system will use the grant that was created first. | [optional] 
**EffectiveAt** | **time.Time** | Effective date for grants and anchor for recurring grants. Provided value will be ceiled to metering windowSize (minute). | 
**MinRolloverAmount** | Pointer to **float64** | Grants are rolled over at reset, after which they can have a different balance compared to what they had before the reset. Balance after the reset is calculated as: Balance_After_Reset &#x3D; MIN(MaxRolloverAmount, MAX(Balance_Before_Reset, MinRolloverAmount)) | [optional] [default to 0]
**Metadata** | Pointer to **map[string]string** | The grant metadata. | [optional] 
**Recurrence** | Pointer to [**RecurringPeriodCreateInput**](RecurringPeriodCreateInput.md) | The subject of the grant. | [optional] 
**MaxRolloverAmount** | Pointer to **float64** | Grants are rolled over at reset, after which they can have a different balance compared to what they had before the reset. The default value equals grant amount. Balance after the reset is calculated as: Balance_After_Reset &#x3D; MIN(MaxRolloverAmount, MAX(Balance_Before_Reset, MinRolloverAmount)) | [optional] 
**Expiration** | Pointer to [**ExpirationPeriod**](ExpirationPeriod.md) | The grant expiration definition. If no expiration is provided, the grant can be active indefinitely. | [optional] 
**Annotations** | Pointer to **map[string]interface{}** | Grant annotations | [optional] 

## Methods

### NewEntitlementGrantCreateInputV2

`func NewEntitlementGrantCreateInputV2(amount float64, effectiveAt time.Time, ) *EntitlementGrantCreateInputV2`

NewEntitlementGrantCreateInputV2 instantiates a new EntitlementGrantCreateInputV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementGrantCreateInputV2WithDefaults

`func NewEntitlementGrantCreateInputV2WithDefaults() *EntitlementGrantCreateInputV2`

NewEntitlementGrantCreateInputV2WithDefaults instantiates a new EntitlementGrantCreateInputV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *EntitlementGrantCreateInputV2) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *EntitlementGrantCreateInputV2) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *EntitlementGrantCreateInputV2) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetPriority

`func (o *EntitlementGrantCreateInputV2) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *EntitlementGrantCreateInputV2) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *EntitlementGrantCreateInputV2) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *EntitlementGrantCreateInputV2) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetEffectiveAt

`func (o *EntitlementGrantCreateInputV2) GetEffectiveAt() time.Time`

GetEffectiveAt returns the EffectiveAt field if non-nil, zero value otherwise.

### GetEffectiveAtOk

`func (o *EntitlementGrantCreateInputV2) GetEffectiveAtOk() (*time.Time, bool)`

GetEffectiveAtOk returns a tuple with the EffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAt

`func (o *EntitlementGrantCreateInputV2) SetEffectiveAt(v time.Time)`

SetEffectiveAt sets EffectiveAt field to given value.


### GetMinRolloverAmount

`func (o *EntitlementGrantCreateInputV2) GetMinRolloverAmount() float64`

GetMinRolloverAmount returns the MinRolloverAmount field if non-nil, zero value otherwise.

### GetMinRolloverAmountOk

`func (o *EntitlementGrantCreateInputV2) GetMinRolloverAmountOk() (*float64, bool)`

GetMinRolloverAmountOk returns a tuple with the MinRolloverAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRolloverAmount

`func (o *EntitlementGrantCreateInputV2) SetMinRolloverAmount(v float64)`

SetMinRolloverAmount sets MinRolloverAmount field to given value.

### HasMinRolloverAmount

`func (o *EntitlementGrantCreateInputV2) HasMinRolloverAmount() bool`

HasMinRolloverAmount returns a boolean if a field has been set.

### GetMetadata

`func (o *EntitlementGrantCreateInputV2) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EntitlementGrantCreateInputV2) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EntitlementGrantCreateInputV2) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EntitlementGrantCreateInputV2) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRecurrence

`func (o *EntitlementGrantCreateInputV2) GetRecurrence() RecurringPeriodCreateInput`

GetRecurrence returns the Recurrence field if non-nil, zero value otherwise.

### GetRecurrenceOk

`func (o *EntitlementGrantCreateInputV2) GetRecurrenceOk() (*RecurringPeriodCreateInput, bool)`

GetRecurrenceOk returns a tuple with the Recurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrence

`func (o *EntitlementGrantCreateInputV2) SetRecurrence(v RecurringPeriodCreateInput)`

SetRecurrence sets Recurrence field to given value.

### HasRecurrence

`func (o *EntitlementGrantCreateInputV2) HasRecurrence() bool`

HasRecurrence returns a boolean if a field has been set.

### GetMaxRolloverAmount

`func (o *EntitlementGrantCreateInputV2) GetMaxRolloverAmount() float64`

GetMaxRolloverAmount returns the MaxRolloverAmount field if non-nil, zero value otherwise.

### GetMaxRolloverAmountOk

`func (o *EntitlementGrantCreateInputV2) GetMaxRolloverAmountOk() (*float64, bool)`

GetMaxRolloverAmountOk returns a tuple with the MaxRolloverAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRolloverAmount

`func (o *EntitlementGrantCreateInputV2) SetMaxRolloverAmount(v float64)`

SetMaxRolloverAmount sets MaxRolloverAmount field to given value.

### HasMaxRolloverAmount

`func (o *EntitlementGrantCreateInputV2) HasMaxRolloverAmount() bool`

HasMaxRolloverAmount returns a boolean if a field has been set.

### GetExpiration

`func (o *EntitlementGrantCreateInputV2) GetExpiration() ExpirationPeriod`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *EntitlementGrantCreateInputV2) GetExpirationOk() (*ExpirationPeriod, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *EntitlementGrantCreateInputV2) SetExpiration(v ExpirationPeriod)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *EntitlementGrantCreateInputV2) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetAnnotations

`func (o *EntitlementGrantCreateInputV2) GetAnnotations() map[string]interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *EntitlementGrantCreateInputV2) GetAnnotationsOk() (*map[string]interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *EntitlementGrantCreateInputV2) SetAnnotations(v map[string]interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *EntitlementGrantCreateInputV2) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


