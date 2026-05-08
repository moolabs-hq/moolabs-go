# EntitlementGrantV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | Timestamp of when the resource was created. | [readonly] 
**UpdatedAt** | **time.Time** | Timestamp of when the resource was last updated. | [readonly] 
**DeletedAt** | Pointer to **time.Time** | Timestamp of when the resource was permanently deleted. | [optional] [readonly] 
**Amount** | **float64** | The amount to grant. Should be a positive number. | 
**Priority** | Pointer to **int32** | The priority of the grant. Grants with higher priority are applied first. Priority is a positive decimal numbers. With lower numbers indicating higher importance. For example, a priority of 1 is more urgent than a priority of 2. When there are several grants available for the same subject, the system selects the grant with the highest priority. In cases where grants share the same priority level, the grant closest to its expiration will be used first. In the case of two grants have identical priorities and expiration dates, the system will use the grant that was created first. | [optional] 
**EffectiveAt** | **time.Time** | Effective date for grants and anchor for recurring grants. Provided value will be ceiled to metering windowSize (minute). | 
**MinRolloverAmount** | Pointer to **float64** | Grants are rolled over at reset, after which they can have a different balance compared to what they had before the reset. Balance after the reset is calculated as: Balance_After_Reset &#x3D; MIN(MaxRolloverAmount, MAX(Balance_Before_Reset, MinRolloverAmount)) | [optional] [default to 0]
**Metadata** | Pointer to **map[string]string** | The grant metadata. | [optional] 
**MaxRolloverAmount** | Pointer to **float64** | Grants are rolled over at reset, after which they can have a different balance compared to what they had before the reset. The default value equals grant amount. Balance after the reset is calculated as: Balance_After_Reset &#x3D; MIN(MaxRolloverAmount, MAX(Balance_Before_Reset, MinRolloverAmount)) | [optional] 
**Expiration** | Pointer to [**ExpirationPeriod**](ExpirationPeriod.md) | The grant expiration definition. If no expiration is provided, the grant can be active indefinitely. | [optional] 
**Annotations** | Pointer to **map[string]interface{}** | Grant annotations | [optional] 
**Id** | **string** | Readonly unique ULID identifier. | [readonly] 
**EntitlementId** | **string** | The unique entitlement ULID that the grant is associated with. | [readonly] 
**NextRecurrence** | Pointer to **time.Time** | The next time the grant will recurr. | [optional] 
**ExpiresAt** | Pointer to **time.Time** | The time the grant expires. | [optional] [readonly] 
**VoidedAt** | Pointer to **time.Time** | The time the grant was voided. | [optional] 
**Recurrence** | Pointer to [**RecurringPeriod**](RecurringPeriod.md) | The recurrence period of the grant. | [optional] 

## Methods

### NewEntitlementGrantV2

`func NewEntitlementGrantV2(createdAt time.Time, updatedAt time.Time, amount float64, effectiveAt time.Time, id string, entitlementId string, ) *EntitlementGrantV2`

NewEntitlementGrantV2 instantiates a new EntitlementGrantV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementGrantV2WithDefaults

`func NewEntitlementGrantV2WithDefaults() *EntitlementGrantV2`

NewEntitlementGrantV2WithDefaults instantiates a new EntitlementGrantV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *EntitlementGrantV2) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EntitlementGrantV2) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EntitlementGrantV2) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *EntitlementGrantV2) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EntitlementGrantV2) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EntitlementGrantV2) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *EntitlementGrantV2) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *EntitlementGrantV2) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *EntitlementGrantV2) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *EntitlementGrantV2) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetAmount

`func (o *EntitlementGrantV2) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *EntitlementGrantV2) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *EntitlementGrantV2) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetPriority

`func (o *EntitlementGrantV2) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *EntitlementGrantV2) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *EntitlementGrantV2) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *EntitlementGrantV2) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetEffectiveAt

`func (o *EntitlementGrantV2) GetEffectiveAt() time.Time`

GetEffectiveAt returns the EffectiveAt field if non-nil, zero value otherwise.

### GetEffectiveAtOk

`func (o *EntitlementGrantV2) GetEffectiveAtOk() (*time.Time, bool)`

GetEffectiveAtOk returns a tuple with the EffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAt

`func (o *EntitlementGrantV2) SetEffectiveAt(v time.Time)`

SetEffectiveAt sets EffectiveAt field to given value.


### GetMinRolloverAmount

`func (o *EntitlementGrantV2) GetMinRolloverAmount() float64`

GetMinRolloverAmount returns the MinRolloverAmount field if non-nil, zero value otherwise.

### GetMinRolloverAmountOk

`func (o *EntitlementGrantV2) GetMinRolloverAmountOk() (*float64, bool)`

GetMinRolloverAmountOk returns a tuple with the MinRolloverAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRolloverAmount

`func (o *EntitlementGrantV2) SetMinRolloverAmount(v float64)`

SetMinRolloverAmount sets MinRolloverAmount field to given value.

### HasMinRolloverAmount

`func (o *EntitlementGrantV2) HasMinRolloverAmount() bool`

HasMinRolloverAmount returns a boolean if a field has been set.

### GetMetadata

`func (o *EntitlementGrantV2) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EntitlementGrantV2) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EntitlementGrantV2) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EntitlementGrantV2) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMaxRolloverAmount

`func (o *EntitlementGrantV2) GetMaxRolloverAmount() float64`

GetMaxRolloverAmount returns the MaxRolloverAmount field if non-nil, zero value otherwise.

### GetMaxRolloverAmountOk

`func (o *EntitlementGrantV2) GetMaxRolloverAmountOk() (*float64, bool)`

GetMaxRolloverAmountOk returns a tuple with the MaxRolloverAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRolloverAmount

`func (o *EntitlementGrantV2) SetMaxRolloverAmount(v float64)`

SetMaxRolloverAmount sets MaxRolloverAmount field to given value.

### HasMaxRolloverAmount

`func (o *EntitlementGrantV2) HasMaxRolloverAmount() bool`

HasMaxRolloverAmount returns a boolean if a field has been set.

### GetExpiration

`func (o *EntitlementGrantV2) GetExpiration() ExpirationPeriod`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *EntitlementGrantV2) GetExpirationOk() (*ExpirationPeriod, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *EntitlementGrantV2) SetExpiration(v ExpirationPeriod)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *EntitlementGrantV2) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetAnnotations

`func (o *EntitlementGrantV2) GetAnnotations() map[string]interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *EntitlementGrantV2) GetAnnotationsOk() (*map[string]interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *EntitlementGrantV2) SetAnnotations(v map[string]interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *EntitlementGrantV2) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetId

`func (o *EntitlementGrantV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntitlementGrantV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntitlementGrantV2) SetId(v string)`

SetId sets Id field to given value.


### GetEntitlementId

`func (o *EntitlementGrantV2) GetEntitlementId() string`

GetEntitlementId returns the EntitlementId field if non-nil, zero value otherwise.

### GetEntitlementIdOk

`func (o *EntitlementGrantV2) GetEntitlementIdOk() (*string, bool)`

GetEntitlementIdOk returns a tuple with the EntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementId

`func (o *EntitlementGrantV2) SetEntitlementId(v string)`

SetEntitlementId sets EntitlementId field to given value.


### GetNextRecurrence

`func (o *EntitlementGrantV2) GetNextRecurrence() time.Time`

GetNextRecurrence returns the NextRecurrence field if non-nil, zero value otherwise.

### GetNextRecurrenceOk

`func (o *EntitlementGrantV2) GetNextRecurrenceOk() (*time.Time, bool)`

GetNextRecurrenceOk returns a tuple with the NextRecurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRecurrence

`func (o *EntitlementGrantV2) SetNextRecurrence(v time.Time)`

SetNextRecurrence sets NextRecurrence field to given value.

### HasNextRecurrence

`func (o *EntitlementGrantV2) HasNextRecurrence() bool`

HasNextRecurrence returns a boolean if a field has been set.

### GetExpiresAt

`func (o *EntitlementGrantV2) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *EntitlementGrantV2) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *EntitlementGrantV2) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *EntitlementGrantV2) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetVoidedAt

`func (o *EntitlementGrantV2) GetVoidedAt() time.Time`

GetVoidedAt returns the VoidedAt field if non-nil, zero value otherwise.

### GetVoidedAtOk

`func (o *EntitlementGrantV2) GetVoidedAtOk() (*time.Time, bool)`

GetVoidedAtOk returns a tuple with the VoidedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidedAt

`func (o *EntitlementGrantV2) SetVoidedAt(v time.Time)`

SetVoidedAt sets VoidedAt field to given value.

### HasVoidedAt

`func (o *EntitlementGrantV2) HasVoidedAt() bool`

HasVoidedAt returns a boolean if a field has been set.

### GetRecurrence

`func (o *EntitlementGrantV2) GetRecurrence() RecurringPeriod`

GetRecurrence returns the Recurrence field if non-nil, zero value otherwise.

### GetRecurrenceOk

`func (o *EntitlementGrantV2) GetRecurrenceOk() (*RecurringPeriod, bool)`

GetRecurrenceOk returns a tuple with the Recurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrence

`func (o *EntitlementGrantV2) SetRecurrence(v RecurringPeriod)`

SetRecurrence sets Recurrence field to given value.

### HasRecurrence

`func (o *EntitlementGrantV2) HasRecurrence() bool`

HasRecurrence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


