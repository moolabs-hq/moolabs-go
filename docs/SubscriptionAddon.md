# SubscriptionAddon

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
**ActiveFrom** | **time.Time** | The cadence start of the resource. | [readonly] 
**ActiveTo** | Pointer to **time.Time** | The cadence end of the resource. | [optional] [readonly] 
**Addon** | [**Addon1**](Addon1.md) |  | 
**QuantityAt** | **time.Time** | For which point in time the quantity was resolved to. | [readonly] 
**Quantity** | **int32** | The quantity of the add-on. Always 1 for single instance add-ons. | 
**Timeline** | [**[]SubscriptionAddonTimelineSegment**](SubscriptionAddonTimelineSegment.md) | The timeline of the add-on. The returned periods are sorted and continuous. | 
**SubscriptionId** | **string** | The ID of the subscription. | [readonly] 
**RateCards** | [**[]SubscriptionAddonRateCard**](SubscriptionAddonRateCard.md) | The rate cards of the add-on. | 

## Methods

### NewSubscriptionAddon

`func NewSubscriptionAddon(id string, name string, createdAt time.Time, updatedAt time.Time, activeFrom time.Time, addon Addon1, quantityAt time.Time, quantity int32, timeline []SubscriptionAddonTimelineSegment, subscriptionId string, rateCards []SubscriptionAddonRateCard, ) *SubscriptionAddon`

NewSubscriptionAddon instantiates a new SubscriptionAddon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionAddonWithDefaults

`func NewSubscriptionAddonWithDefaults() *SubscriptionAddon`

NewSubscriptionAddonWithDefaults instantiates a new SubscriptionAddon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubscriptionAddon) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionAddon) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionAddon) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SubscriptionAddon) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubscriptionAddon) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubscriptionAddon) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SubscriptionAddon) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SubscriptionAddon) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SubscriptionAddon) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SubscriptionAddon) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *SubscriptionAddon) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SubscriptionAddon) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SubscriptionAddon) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SubscriptionAddon) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SubscriptionAddon) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubscriptionAddon) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubscriptionAddon) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *SubscriptionAddon) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SubscriptionAddon) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SubscriptionAddon) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *SubscriptionAddon) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *SubscriptionAddon) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *SubscriptionAddon) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *SubscriptionAddon) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetActiveFrom

`func (o *SubscriptionAddon) GetActiveFrom() time.Time`

GetActiveFrom returns the ActiveFrom field if non-nil, zero value otherwise.

### GetActiveFromOk

`func (o *SubscriptionAddon) GetActiveFromOk() (*time.Time, bool)`

GetActiveFromOk returns a tuple with the ActiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveFrom

`func (o *SubscriptionAddon) SetActiveFrom(v time.Time)`

SetActiveFrom sets ActiveFrom field to given value.


### GetActiveTo

`func (o *SubscriptionAddon) GetActiveTo() time.Time`

GetActiveTo returns the ActiveTo field if non-nil, zero value otherwise.

### GetActiveToOk

`func (o *SubscriptionAddon) GetActiveToOk() (*time.Time, bool)`

GetActiveToOk returns a tuple with the ActiveTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTo

`func (o *SubscriptionAddon) SetActiveTo(v time.Time)`

SetActiveTo sets ActiveTo field to given value.

### HasActiveTo

`func (o *SubscriptionAddon) HasActiveTo() bool`

HasActiveTo returns a boolean if a field has been set.

### GetAddon

`func (o *SubscriptionAddon) GetAddon() Addon1`

GetAddon returns the Addon field if non-nil, zero value otherwise.

### GetAddonOk

`func (o *SubscriptionAddon) GetAddonOk() (*Addon1, bool)`

GetAddonOk returns a tuple with the Addon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddon

`func (o *SubscriptionAddon) SetAddon(v Addon1)`

SetAddon sets Addon field to given value.


### GetQuantityAt

`func (o *SubscriptionAddon) GetQuantityAt() time.Time`

GetQuantityAt returns the QuantityAt field if non-nil, zero value otherwise.

### GetQuantityAtOk

`func (o *SubscriptionAddon) GetQuantityAtOk() (*time.Time, bool)`

GetQuantityAtOk returns a tuple with the QuantityAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityAt

`func (o *SubscriptionAddon) SetQuantityAt(v time.Time)`

SetQuantityAt sets QuantityAt field to given value.


### GetQuantity

`func (o *SubscriptionAddon) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *SubscriptionAddon) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *SubscriptionAddon) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetTimeline

`func (o *SubscriptionAddon) GetTimeline() []SubscriptionAddonTimelineSegment`

GetTimeline returns the Timeline field if non-nil, zero value otherwise.

### GetTimelineOk

`func (o *SubscriptionAddon) GetTimelineOk() (*[]SubscriptionAddonTimelineSegment, bool)`

GetTimelineOk returns a tuple with the Timeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeline

`func (o *SubscriptionAddon) SetTimeline(v []SubscriptionAddonTimelineSegment)`

SetTimeline sets Timeline field to given value.


### GetSubscriptionId

`func (o *SubscriptionAddon) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *SubscriptionAddon) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *SubscriptionAddon) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetRateCards

`func (o *SubscriptionAddon) GetRateCards() []SubscriptionAddonRateCard`

GetRateCards returns the RateCards field if non-nil, zero value otherwise.

### GetRateCardsOk

`func (o *SubscriptionAddon) GetRateCardsOk() (*[]SubscriptionAddonRateCard, bool)`

GetRateCardsOk returns a tuple with the RateCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateCards

`func (o *SubscriptionAddon) SetRateCards(v []SubscriptionAddonRateCard)`

SetRateCards sets RateCards field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


