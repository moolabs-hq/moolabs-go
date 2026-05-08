# SubscriptionAddonUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Human-readable name for the resource. Between 1 and 256 characters. | [optional] 
**Description** | Pointer to **string** | Optional description of the resource. Maximum 1024 characters. | [optional] 
**Metadata** | Pointer to **map[string]string** | Additional metadata for the resource. | [optional] 
**Quantity** | Pointer to **int32** | The quantity of the add-on. Always 1 for single instance add-ons. | [optional] 
**Timing** | Pointer to [**SubscriptionTiming**](SubscriptionTiming.md) | The timing of the operation. After the create or update, a new entry will be created in the timeline. | [optional] 

## Methods

### NewSubscriptionAddonUpdate

`func NewSubscriptionAddonUpdate() *SubscriptionAddonUpdate`

NewSubscriptionAddonUpdate instantiates a new SubscriptionAddonUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionAddonUpdateWithDefaults

`func NewSubscriptionAddonUpdateWithDefaults() *SubscriptionAddonUpdate`

NewSubscriptionAddonUpdateWithDefaults instantiates a new SubscriptionAddonUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SubscriptionAddonUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubscriptionAddonUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubscriptionAddonUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SubscriptionAddonUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *SubscriptionAddonUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SubscriptionAddonUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SubscriptionAddonUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SubscriptionAddonUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *SubscriptionAddonUpdate) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SubscriptionAddonUpdate) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SubscriptionAddonUpdate) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SubscriptionAddonUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetQuantity

`func (o *SubscriptionAddonUpdate) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *SubscriptionAddonUpdate) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *SubscriptionAddonUpdate) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *SubscriptionAddonUpdate) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetTiming

`func (o *SubscriptionAddonUpdate) GetTiming() SubscriptionTiming`

GetTiming returns the Timing field if non-nil, zero value otherwise.

### GetTimingOk

`func (o *SubscriptionAddonUpdate) GetTimingOk() (*SubscriptionTiming, bool)`

GetTimingOk returns a tuple with the Timing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiming

`func (o *SubscriptionAddonUpdate) SetTiming(v SubscriptionTiming)`

SetTiming sets Timing field to given value.

### HasTiming

`func (o *SubscriptionAddonUpdate) HasTiming() bool`

HasTiming returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


