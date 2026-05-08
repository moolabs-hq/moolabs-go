# EntitlementGrant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | Timestamp of when the resource was created. | [readonly] 
**UpdatedAt** | **time.Time** | Timestamp of when the resource was last updated. | [readonly] 
**DeletedAt** | Pointer to **time.Time** | Timestamp of when the resource was permanently deleted. | [optional] [readonly] 
**Amount** | **float64** | The amount to grant. Should be a positive number. | 
**Priority** | Pointer to **int32** | The priority of the grant. Grants with higher priority are applied first. Priority is a positive decimal numbers. With lower numbers indicating higher importance. For example, a priority of 1 is more urgent than a priority of 2. When there are several grants available for the same subject, the system selects the grant with the highest priority. In cases where grants share the same priority level, the grant closest to its expiration will be used first. In the case of two grants have identical priorities and expiration dates, the system will use the grant that was created first. | [optional] 
**EffectiveAt** | **time.Time** | Effective date for grants and anchor for recurring grants. Provided value will be ceiled to metering windowSize (minute). | 
**Expiration** | [**ExpirationPeriod**](ExpirationPeriod.md) | The grant expiration definition | 
**MaxRolloverAmount** | Pointer to **float64** | Grants are rolled over at reset, after which they can have a different balance compared to what they had before the reset. Balance after the reset is calculated as: Balance_After_Reset &#x3D; MIN(MaxRolloverAmount, MAX(Balance_Before_Reset, MinRolloverAmount)) | [optional] [default to 0]
**MinRolloverAmount** | Pointer to **float64** | Grants are rolled over at reset, after which they can have a different balance compared to what they had before the reset. Balance after the reset is calculated as: Balance_After_Reset &#x3D; MIN(MaxRolloverAmount, MAX(Balance_Before_Reset, MinRolloverAmount)) | [optional] [default to 0]
**Metadata** | Pointer to **map[string]string** | The grant metadata. | [optional] 
**Id** | **string** | Readonly unique ULID identifier. | [readonly] 
**EntitlementId** | **string** | The unique entitlement ULID that the grant is associated with. | [readonly] 
**NextRecurrence** | Pointer to **time.Time** | The next time the grant will recurr. | [optional] 
**ExpiresAt** | Pointer to **time.Time** | The time the grant expires. | [optional] [readonly] 
**VoidedAt** | Pointer to **time.Time** | The time the grant was voided. | [optional] 
**Recurrence** | Pointer to [**RecurringPeriod**](RecurringPeriod.md) | The recurrence period of the grant. | [optional] 
**Annotations** | Pointer to **map[string]interface{}** | Grant annotations | [optional] 

## Methods

### NewEntitlementGrant

`func NewEntitlementGrant(createdAt time.Time, updatedAt time.Time, amount float64, effectiveAt time.Time, expiration ExpirationPeriod, id string, entitlementId string, ) *EntitlementGrant`

NewEntitlementGrant instantiates a new EntitlementGrant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementGrantWithDefaults

`func NewEntitlementGrantWithDefaults() *EntitlementGrant`

NewEntitlementGrantWithDefaults instantiates a new EntitlementGrant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *EntitlementGrant) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EntitlementGrant) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EntitlementGrant) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *EntitlementGrant) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EntitlementGrant) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EntitlementGrant) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *EntitlementGrant) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *EntitlementGrant) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *EntitlementGrant) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *EntitlementGrant) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetAmount

`func (o *EntitlementGrant) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *EntitlementGrant) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *EntitlementGrant) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetPriority

`func (o *EntitlementGrant) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *EntitlementGrant) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *EntitlementGrant) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *EntitlementGrant) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetEffectiveAt

`func (o *EntitlementGrant) GetEffectiveAt() time.Time`

GetEffectiveAt returns the EffectiveAt field if non-nil, zero value otherwise.

### GetEffectiveAtOk

`func (o *EntitlementGrant) GetEffectiveAtOk() (*time.Time, bool)`

GetEffectiveAtOk returns a tuple with the EffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAt

`func (o *EntitlementGrant) SetEffectiveAt(v time.Time)`

SetEffectiveAt sets EffectiveAt field to given value.


### GetExpiration

`func (o *EntitlementGrant) GetExpiration() ExpirationPeriod`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *EntitlementGrant) GetExpirationOk() (*ExpirationPeriod, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *EntitlementGrant) SetExpiration(v ExpirationPeriod)`

SetExpiration sets Expiration field to given value.


### GetMaxRolloverAmount

`func (o *EntitlementGrant) GetMaxRolloverAmount() float64`

GetMaxRolloverAmount returns the MaxRolloverAmount field if non-nil, zero value otherwise.

### GetMaxRolloverAmountOk

`func (o *EntitlementGrant) GetMaxRolloverAmountOk() (*float64, bool)`

GetMaxRolloverAmountOk returns a tuple with the MaxRolloverAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRolloverAmount

`func (o *EntitlementGrant) SetMaxRolloverAmount(v float64)`

SetMaxRolloverAmount sets MaxRolloverAmount field to given value.

### HasMaxRolloverAmount

`func (o *EntitlementGrant) HasMaxRolloverAmount() bool`

HasMaxRolloverAmount returns a boolean if a field has been set.

### GetMinRolloverAmount

`func (o *EntitlementGrant) GetMinRolloverAmount() float64`

GetMinRolloverAmount returns the MinRolloverAmount field if non-nil, zero value otherwise.

### GetMinRolloverAmountOk

`func (o *EntitlementGrant) GetMinRolloverAmountOk() (*float64, bool)`

GetMinRolloverAmountOk returns a tuple with the MinRolloverAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRolloverAmount

`func (o *EntitlementGrant) SetMinRolloverAmount(v float64)`

SetMinRolloverAmount sets MinRolloverAmount field to given value.

### HasMinRolloverAmount

`func (o *EntitlementGrant) HasMinRolloverAmount() bool`

HasMinRolloverAmount returns a boolean if a field has been set.

### GetMetadata

`func (o *EntitlementGrant) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EntitlementGrant) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EntitlementGrant) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EntitlementGrant) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetId

`func (o *EntitlementGrant) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntitlementGrant) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntitlementGrant) SetId(v string)`

SetId sets Id field to given value.


### GetEntitlementId

`func (o *EntitlementGrant) GetEntitlementId() string`

GetEntitlementId returns the EntitlementId field if non-nil, zero value otherwise.

### GetEntitlementIdOk

`func (o *EntitlementGrant) GetEntitlementIdOk() (*string, bool)`

GetEntitlementIdOk returns a tuple with the EntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementId

`func (o *EntitlementGrant) SetEntitlementId(v string)`

SetEntitlementId sets EntitlementId field to given value.


### GetNextRecurrence

`func (o *EntitlementGrant) GetNextRecurrence() time.Time`

GetNextRecurrence returns the NextRecurrence field if non-nil, zero value otherwise.

### GetNextRecurrenceOk

`func (o *EntitlementGrant) GetNextRecurrenceOk() (*time.Time, bool)`

GetNextRecurrenceOk returns a tuple with the NextRecurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRecurrence

`func (o *EntitlementGrant) SetNextRecurrence(v time.Time)`

SetNextRecurrence sets NextRecurrence field to given value.

### HasNextRecurrence

`func (o *EntitlementGrant) HasNextRecurrence() bool`

HasNextRecurrence returns a boolean if a field has been set.

### GetExpiresAt

`func (o *EntitlementGrant) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *EntitlementGrant) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *EntitlementGrant) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *EntitlementGrant) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetVoidedAt

`func (o *EntitlementGrant) GetVoidedAt() time.Time`

GetVoidedAt returns the VoidedAt field if non-nil, zero value otherwise.

### GetVoidedAtOk

`func (o *EntitlementGrant) GetVoidedAtOk() (*time.Time, bool)`

GetVoidedAtOk returns a tuple with the VoidedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidedAt

`func (o *EntitlementGrant) SetVoidedAt(v time.Time)`

SetVoidedAt sets VoidedAt field to given value.

### HasVoidedAt

`func (o *EntitlementGrant) HasVoidedAt() bool`

HasVoidedAt returns a boolean if a field has been set.

### GetRecurrence

`func (o *EntitlementGrant) GetRecurrence() RecurringPeriod`

GetRecurrence returns the Recurrence field if non-nil, zero value otherwise.

### GetRecurrenceOk

`func (o *EntitlementGrant) GetRecurrenceOk() (*RecurringPeriod, bool)`

GetRecurrenceOk returns a tuple with the Recurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrence

`func (o *EntitlementGrant) SetRecurrence(v RecurringPeriod)`

SetRecurrence sets Recurrence field to given value.

### HasRecurrence

`func (o *EntitlementGrant) HasRecurrence() bool`

HasRecurrence returns a boolean if a field has been set.

### GetAnnotations

`func (o *EntitlementGrant) GetAnnotations() map[string]interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *EntitlementGrant) GetAnnotationsOk() (*map[string]interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *EntitlementGrant) SetAnnotations(v map[string]interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *EntitlementGrant) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


