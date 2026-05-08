# TieredPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the price.  One of: flat, unit, or tiered. | 
**Mode** | [**TieredPriceMode**](TieredPriceMode.md) | Defines if the tiering mode is volume-based or graduated: - In &#x60;volume&#x60;-based tiering, the maximum quantity within a period determines the per unit price. - In &#x60;graduated&#x60; tiering, pricing can change as the quantity grows. | 
**Tiers** | [**[]PriceTier**](PriceTier.md) | The tiers of the tiered price. At least one price component is required in each tier. | 

## Methods

### NewTieredPrice

`func NewTieredPrice(type_ string, mode TieredPriceMode, tiers []PriceTier, ) *TieredPrice`

NewTieredPrice instantiates a new TieredPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTieredPriceWithDefaults

`func NewTieredPriceWithDefaults() *TieredPrice`

NewTieredPriceWithDefaults instantiates a new TieredPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TieredPrice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TieredPrice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TieredPrice) SetType(v string)`

SetType sets Type field to given value.


### GetMode

`func (o *TieredPrice) GetMode() TieredPriceMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *TieredPrice) GetModeOk() (*TieredPriceMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *TieredPrice) SetMode(v TieredPriceMode)`

SetMode sets Mode field to given value.


### GetTiers

`func (o *TieredPrice) GetTiers() []PriceTier`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *TieredPrice) GetTiersOk() (*[]PriceTier, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *TieredPrice) SetTiers(v []PriceTier)`

SetTiers sets Tiers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


