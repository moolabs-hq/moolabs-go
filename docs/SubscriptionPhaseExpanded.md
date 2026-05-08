# SubscriptionPhaseExpanded

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
**Key** | **string** | A locally unique identifier for the resource. | 
**Discounts** | Pointer to [**Discounts**](Discounts.md) | The discounts on the plan. | [optional] 
**ActiveFrom** | **time.Time** | The time from which the phase is active. | 
**ActiveTo** | Pointer to **time.Time** | The until which the Phase is active. | [optional] 
**Items** | [**[]SubscriptionItem**](SubscriptionItem.md) | The items of the phase. The structure is flattened to better conform to the Plan API. The timelines are flattened according to the following rules: - for the current phase, the &#x60;items&#x60; contains only the active item for each key - for past phases, the &#x60;items&#x60; contains only the last item for each key - for future phases, the &#x60;items&#x60; contains only the first version of the item for each key | 
**ItemTimelines** | [**map[string][]SubscriptionItem**](array.md) | Includes all versions of the items on each key, including all edits, scheduled changes, etc... | 

## Methods

### NewSubscriptionPhaseExpanded

`func NewSubscriptionPhaseExpanded(id string, name string, createdAt time.Time, updatedAt time.Time, key string, activeFrom time.Time, items []SubscriptionItem, itemTimelines map[string][]SubscriptionItem, ) *SubscriptionPhaseExpanded`

NewSubscriptionPhaseExpanded instantiates a new SubscriptionPhaseExpanded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionPhaseExpandedWithDefaults

`func NewSubscriptionPhaseExpandedWithDefaults() *SubscriptionPhaseExpanded`

NewSubscriptionPhaseExpandedWithDefaults instantiates a new SubscriptionPhaseExpanded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubscriptionPhaseExpanded) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionPhaseExpanded) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionPhaseExpanded) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SubscriptionPhaseExpanded) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubscriptionPhaseExpanded) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubscriptionPhaseExpanded) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SubscriptionPhaseExpanded) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SubscriptionPhaseExpanded) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SubscriptionPhaseExpanded) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SubscriptionPhaseExpanded) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *SubscriptionPhaseExpanded) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SubscriptionPhaseExpanded) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SubscriptionPhaseExpanded) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SubscriptionPhaseExpanded) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SubscriptionPhaseExpanded) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubscriptionPhaseExpanded) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubscriptionPhaseExpanded) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *SubscriptionPhaseExpanded) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SubscriptionPhaseExpanded) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SubscriptionPhaseExpanded) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *SubscriptionPhaseExpanded) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *SubscriptionPhaseExpanded) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *SubscriptionPhaseExpanded) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *SubscriptionPhaseExpanded) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetKey

`func (o *SubscriptionPhaseExpanded) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SubscriptionPhaseExpanded) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SubscriptionPhaseExpanded) SetKey(v string)`

SetKey sets Key field to given value.


### GetDiscounts

`func (o *SubscriptionPhaseExpanded) GetDiscounts() Discounts`

GetDiscounts returns the Discounts field if non-nil, zero value otherwise.

### GetDiscountsOk

`func (o *SubscriptionPhaseExpanded) GetDiscountsOk() (*Discounts, bool)`

GetDiscountsOk returns a tuple with the Discounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscounts

`func (o *SubscriptionPhaseExpanded) SetDiscounts(v Discounts)`

SetDiscounts sets Discounts field to given value.

### HasDiscounts

`func (o *SubscriptionPhaseExpanded) HasDiscounts() bool`

HasDiscounts returns a boolean if a field has been set.

### GetActiveFrom

`func (o *SubscriptionPhaseExpanded) GetActiveFrom() time.Time`

GetActiveFrom returns the ActiveFrom field if non-nil, zero value otherwise.

### GetActiveFromOk

`func (o *SubscriptionPhaseExpanded) GetActiveFromOk() (*time.Time, bool)`

GetActiveFromOk returns a tuple with the ActiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveFrom

`func (o *SubscriptionPhaseExpanded) SetActiveFrom(v time.Time)`

SetActiveFrom sets ActiveFrom field to given value.


### GetActiveTo

`func (o *SubscriptionPhaseExpanded) GetActiveTo() time.Time`

GetActiveTo returns the ActiveTo field if non-nil, zero value otherwise.

### GetActiveToOk

`func (o *SubscriptionPhaseExpanded) GetActiveToOk() (*time.Time, bool)`

GetActiveToOk returns a tuple with the ActiveTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTo

`func (o *SubscriptionPhaseExpanded) SetActiveTo(v time.Time)`

SetActiveTo sets ActiveTo field to given value.

### HasActiveTo

`func (o *SubscriptionPhaseExpanded) HasActiveTo() bool`

HasActiveTo returns a boolean if a field has been set.

### GetItems

`func (o *SubscriptionPhaseExpanded) GetItems() []SubscriptionItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SubscriptionPhaseExpanded) GetItemsOk() (*[]SubscriptionItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SubscriptionPhaseExpanded) SetItems(v []SubscriptionItem)`

SetItems sets Items field to given value.


### GetItemTimelines

`func (o *SubscriptionPhaseExpanded) GetItemTimelines() map[string][]SubscriptionItem`

GetItemTimelines returns the ItemTimelines field if non-nil, zero value otherwise.

### GetItemTimelinesOk

`func (o *SubscriptionPhaseExpanded) GetItemTimelinesOk() (*map[string][]SubscriptionItem, bool)`

GetItemTimelinesOk returns a tuple with the ItemTimelines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemTimelines

`func (o *SubscriptionPhaseExpanded) SetItemTimelines(v map[string][]SubscriptionItem)`

SetItemTimelines sets ItemTimelines field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


