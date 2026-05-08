# AddonReplaceUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Human-readable name for the resource. Between 1 and 256 characters. | 
**Description** | Pointer to **string** | Optional description of the resource. Maximum 1024 characters. | [optional] 
**Metadata** | Pointer to **map[string]string** | Additional metadata for the resource. | [optional] 
**InstanceType** | [**AddonInstanceType**](AddonInstanceType.md) | The instanceType of the add-ons. Can be \&quot;single\&quot; or \&quot;multiple\&quot;. | 
**RateCards** | [**[]RateCard**](RateCard.md) | The rate cards of the add-on. | 

## Methods

### NewAddonReplaceUpdate

`func NewAddonReplaceUpdate(name string, instanceType AddonInstanceType, rateCards []RateCard, ) *AddonReplaceUpdate`

NewAddonReplaceUpdate instantiates a new AddonReplaceUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddonReplaceUpdateWithDefaults

`func NewAddonReplaceUpdateWithDefaults() *AddonReplaceUpdate`

NewAddonReplaceUpdateWithDefaults instantiates a new AddonReplaceUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddonReplaceUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddonReplaceUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddonReplaceUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AddonReplaceUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddonReplaceUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddonReplaceUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddonReplaceUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *AddonReplaceUpdate) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AddonReplaceUpdate) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AddonReplaceUpdate) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AddonReplaceUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetInstanceType

`func (o *AddonReplaceUpdate) GetInstanceType() AddonInstanceType`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *AddonReplaceUpdate) GetInstanceTypeOk() (*AddonInstanceType, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *AddonReplaceUpdate) SetInstanceType(v AddonInstanceType)`

SetInstanceType sets InstanceType field to given value.


### GetRateCards

`func (o *AddonReplaceUpdate) GetRateCards() []RateCard`

GetRateCards returns the RateCards field if non-nil, zero value otherwise.

### GetRateCardsOk

`func (o *AddonReplaceUpdate) GetRateCardsOk() (*[]RateCard, bool)`

GetRateCardsOk returns a tuple with the RateCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateCards

`func (o *AddonReplaceUpdate) SetRateCards(v []RateCard)`

SetRateCards sets RateCards field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


