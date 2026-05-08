# AddonCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Human-readable name for the resource. Between 1 and 256 characters. | 
**Description** | Pointer to **string** | Optional description of the resource. Maximum 1024 characters. | [optional] 
**Metadata** | Pointer to **map[string]string** | Additional metadata for the resource. | [optional] 
**Key** | **string** | A semi-unique identifier for the resource. | 
**InstanceType** | [**AddonInstanceType**](AddonInstanceType.md) | The instanceType of the add-ons. Can be \&quot;single\&quot; or \&quot;multiple\&quot;. | 
**Currency** | **string** | The currency code of the add-on. | 
**RateCards** | [**[]RateCard**](RateCard.md) | The rate cards of the add-on. | 

## Methods

### NewAddonCreate

`func NewAddonCreate(name string, key string, instanceType AddonInstanceType, currency string, rateCards []RateCard, ) *AddonCreate`

NewAddonCreate instantiates a new AddonCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddonCreateWithDefaults

`func NewAddonCreateWithDefaults() *AddonCreate`

NewAddonCreateWithDefaults instantiates a new AddonCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddonCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddonCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddonCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AddonCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddonCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddonCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddonCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *AddonCreate) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AddonCreate) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AddonCreate) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AddonCreate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetKey

`func (o *AddonCreate) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AddonCreate) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AddonCreate) SetKey(v string)`

SetKey sets Key field to given value.


### GetInstanceType

`func (o *AddonCreate) GetInstanceType() AddonInstanceType`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *AddonCreate) GetInstanceTypeOk() (*AddonInstanceType, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *AddonCreate) SetInstanceType(v AddonInstanceType)`

SetInstanceType sets InstanceType field to given value.


### GetCurrency

`func (o *AddonCreate) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AddonCreate) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AddonCreate) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetRateCards

`func (o *AddonCreate) GetRateCards() []RateCard`

GetRateCards returns the RateCards field if non-nil, zero value otherwise.

### GetRateCardsOk

`func (o *AddonCreate) GetRateCardsOk() (*[]RateCard, bool)`

GetRateCardsOk returns a tuple with the RateCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateCards

`func (o *AddonCreate) SetRateCards(v []RateCard)`

SetRateCards sets RateCards field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


