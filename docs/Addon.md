# Addon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier for the resource. | [readonly] 
**Name** | **string** | Human-readable name for the resource. Between 1 and 256 characters. | 
**Description** | Pointer to **string** | Optional description of the resource. Maximum 1024 characters. | [optional] 
**Metadata** | Pointer to **map[string]string** | Additional metadata for the resource. | [optional] 
**CreatedAt** | **time.Time** | Timestamp of when the resource was created. | [readonly] 
**UpdatedAt** | **time.Time** | Timestamp of when the resource was last updated. | [readonly] 
**DeletedAt** | Pointer to **time.Time** | Timestamp of when the resource was permanently deleted. | [optional] [readonly] 
**Key** | **string** | A semi-unique identifier for the resource. | 
**Annotations** | Pointer to **map[string]interface{}** | Set of key-value pairs managed by the system. Cannot be modified by user. | [optional] [readonly] 
**Version** | **int32** | Version of the add-on. Incremented when the add-on is updated. | [readonly] [default to 1]
**InstanceType** | [**AddonInstanceType**](AddonInstanceType.md) | The instanceType of the add-ons. Can be \&quot;single\&quot; or \&quot;multiple\&quot;. | 
**Currency** | **string** | The currency code of the add-on. | 
**EffectiveFrom** | Pointer to **time.Time** | The date and time when the add-on becomes effective. When not specified, the add-on is a draft. | [optional] [readonly] 
**EffectiveTo** | Pointer to **time.Time** | The date and time when the add-on is no longer effective. When not specified, the add-on is effective indefinitely. | [optional] [readonly] 
**Status** | [**AddonStatus**](AddonStatus.md) | The status of the add-on. Computed based on the effective start and end dates: - draft &#x3D; no effectiveFrom - active &#x3D; effectiveFrom &lt;&#x3D; now &lt; effectiveTo - archived  &#x3D; effectiveTo &lt;&#x3D; now | [readonly] 
**RateCards** | [**[]RateCard**](RateCard.md) | The rate cards of the add-on. | 
**ValidationErrors** | [**[]MeterValidationError**](MeterValidationError.md) | List of validation errors. | 

## Methods

### NewAddon

`func NewAddon(id string, name string, createdAt time.Time, updatedAt time.Time, key string, version int32, instanceType AddonInstanceType, currency string, status AddonStatus, rateCards []RateCard, validationErrors []MeterValidationError, ) *Addon`

NewAddon instantiates a new Addon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddonWithDefaults

`func NewAddonWithDefaults() *Addon`

NewAddonWithDefaults instantiates a new Addon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Addon) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Addon) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Addon) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Addon) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Addon) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Addon) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Addon) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Addon) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Addon) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Addon) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *Addon) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Addon) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Addon) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Addon) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Addon) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Addon) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Addon) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Addon) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Addon) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Addon) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *Addon) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Addon) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Addon) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Addon) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetKey

`func (o *Addon) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Addon) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Addon) SetKey(v string)`

SetKey sets Key field to given value.


### GetAnnotations

`func (o *Addon) GetAnnotations() map[string]interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *Addon) GetAnnotationsOk() (*map[string]interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *Addon) SetAnnotations(v map[string]interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *Addon) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetVersion

`func (o *Addon) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Addon) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Addon) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetInstanceType

`func (o *Addon) GetInstanceType() AddonInstanceType`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *Addon) GetInstanceTypeOk() (*AddonInstanceType, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *Addon) SetInstanceType(v AddonInstanceType)`

SetInstanceType sets InstanceType field to given value.


### GetCurrency

`func (o *Addon) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Addon) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Addon) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetEffectiveFrom

`func (o *Addon) GetEffectiveFrom() time.Time`

GetEffectiveFrom returns the EffectiveFrom field if non-nil, zero value otherwise.

### GetEffectiveFromOk

`func (o *Addon) GetEffectiveFromOk() (*time.Time, bool)`

GetEffectiveFromOk returns a tuple with the EffectiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveFrom

`func (o *Addon) SetEffectiveFrom(v time.Time)`

SetEffectiveFrom sets EffectiveFrom field to given value.

### HasEffectiveFrom

`func (o *Addon) HasEffectiveFrom() bool`

HasEffectiveFrom returns a boolean if a field has been set.

### GetEffectiveTo

`func (o *Addon) GetEffectiveTo() time.Time`

GetEffectiveTo returns the EffectiveTo field if non-nil, zero value otherwise.

### GetEffectiveToOk

`func (o *Addon) GetEffectiveToOk() (*time.Time, bool)`

GetEffectiveToOk returns a tuple with the EffectiveTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveTo

`func (o *Addon) SetEffectiveTo(v time.Time)`

SetEffectiveTo sets EffectiveTo field to given value.

### HasEffectiveTo

`func (o *Addon) HasEffectiveTo() bool`

HasEffectiveTo returns a boolean if a field has been set.

### GetStatus

`func (o *Addon) GetStatus() AddonStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Addon) GetStatusOk() (*AddonStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Addon) SetStatus(v AddonStatus)`

SetStatus sets Status field to given value.


### GetRateCards

`func (o *Addon) GetRateCards() []RateCard`

GetRateCards returns the RateCards field if non-nil, zero value otherwise.

### GetRateCardsOk

`func (o *Addon) GetRateCardsOk() (*[]RateCard, bool)`

GetRateCardsOk returns a tuple with the RateCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateCards

`func (o *Addon) SetRateCards(v []RateCard)`

SetRateCards sets RateCards field to given value.


### GetValidationErrors

`func (o *Addon) GetValidationErrors() []MeterValidationError`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *Addon) GetValidationErrorsOk() (*[]MeterValidationError, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *Addon) SetValidationErrors(v []MeterValidationError)`

SetValidationErrors sets ValidationErrors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


