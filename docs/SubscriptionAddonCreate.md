# SubscriptionAddonCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Human-readable name for the resource. Between 1 and 256 characters. | 
**Description** | Pointer to **string** | Optional description of the resource. Maximum 1024 characters. | [optional] 
**Metadata** | Pointer to **map[string]string** | Additional metadata for the resource. | [optional] 
**Quantity** | **int32** | The quantity of the add-on. Always 1 for single instance add-ons. | 
**Timing** | [**SubscriptionTiming**](SubscriptionTiming.md) | The timing of the operation. After the create or update, a new entry will be created in the timeline. | 
**Addon** | [**Addon2**](Addon2.md) |  | 

## Methods

### NewSubscriptionAddonCreate

`func NewSubscriptionAddonCreate(name string, quantity int32, timing SubscriptionTiming, addon Addon2, ) *SubscriptionAddonCreate`

NewSubscriptionAddonCreate instantiates a new SubscriptionAddonCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionAddonCreateWithDefaults

`func NewSubscriptionAddonCreateWithDefaults() *SubscriptionAddonCreate`

NewSubscriptionAddonCreateWithDefaults instantiates a new SubscriptionAddonCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SubscriptionAddonCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubscriptionAddonCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubscriptionAddonCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SubscriptionAddonCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SubscriptionAddonCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SubscriptionAddonCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SubscriptionAddonCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *SubscriptionAddonCreate) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SubscriptionAddonCreate) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SubscriptionAddonCreate) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SubscriptionAddonCreate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetQuantity

`func (o *SubscriptionAddonCreate) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *SubscriptionAddonCreate) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *SubscriptionAddonCreate) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetTiming

`func (o *SubscriptionAddonCreate) GetTiming() SubscriptionTiming`

GetTiming returns the Timing field if non-nil, zero value otherwise.

### GetTimingOk

`func (o *SubscriptionAddonCreate) GetTimingOk() (*SubscriptionTiming, bool)`

GetTimingOk returns a tuple with the Timing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiming

`func (o *SubscriptionAddonCreate) SetTiming(v SubscriptionTiming)`

SetTiming sets Timing field to given value.


### GetAddon

`func (o *SubscriptionAddonCreate) GetAddon() Addon2`

GetAddon returns the Addon field if non-nil, zero value otherwise.

### GetAddonOk

`func (o *SubscriptionAddonCreate) GetAddonOk() (*Addon2, bool)`

GetAddonOk returns a tuple with the Addon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddon

`func (o *SubscriptionAddonCreate) SetAddon(v Addon2)`

SetAddon sets Addon field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


