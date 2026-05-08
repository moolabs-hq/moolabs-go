# Price

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the price. | 
**Amount** | **string** | The price of one package. | 
**Mode** | [**TieredPriceMode**](TieredPriceMode.md) | Defines if the tiering mode is volume-based or graduated: - In &#x60;volume&#x60;-based tiering, the maximum quantity within a period determines the per unit price. - In &#x60;graduated&#x60; tiering, pricing can change as the quantity grows. | 
**Tiers** | [**[]PriceTier**](PriceTier.md) | The tiers of the tiered price. At least one price component is required in each tier. | 
**Multiplier** | Pointer to **string** | The multiplier to apply to the base price to get the dynamic price.  Examples: - 0.0: the price is zero - 0.5: the price is 50% of the base price - 1.0: the price is the same as the base price - 1.5: the price is 150% of the base price | [optional] 
**QuantityPerPackage** | **string** | The quantity per package. | 

## Methods

### NewPrice

`func NewPrice(type_ string, amount string, mode TieredPriceMode, tiers []PriceTier, quantityPerPackage string, ) *Price`

NewPrice instantiates a new Price object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceWithDefaults

`func NewPriceWithDefaults() *Price`

NewPriceWithDefaults instantiates a new Price object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Price) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Price) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Price) SetType(v string)`

SetType sets Type field to given value.


### GetAmount

`func (o *Price) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Price) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Price) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetMode

`func (o *Price) GetMode() TieredPriceMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *Price) GetModeOk() (*TieredPriceMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *Price) SetMode(v TieredPriceMode)`

SetMode sets Mode field to given value.


### GetTiers

`func (o *Price) GetTiers() []PriceTier`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *Price) GetTiersOk() (*[]PriceTier, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *Price) SetTiers(v []PriceTier)`

SetTiers sets Tiers field to given value.


### GetMultiplier

`func (o *Price) GetMultiplier() string`

GetMultiplier returns the Multiplier field if non-nil, zero value otherwise.

### GetMultiplierOk

`func (o *Price) GetMultiplierOk() (*string, bool)`

GetMultiplierOk returns a tuple with the Multiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplier

`func (o *Price) SetMultiplier(v string)`

SetMultiplier sets Multiplier field to given value.

### HasMultiplier

`func (o *Price) HasMultiplier() bool`

HasMultiplier returns a boolean if a field has been set.

### GetQuantityPerPackage

`func (o *Price) GetQuantityPerPackage() string`

GetQuantityPerPackage returns the QuantityPerPackage field if non-nil, zero value otherwise.

### GetQuantityPerPackageOk

`func (o *Price) GetQuantityPerPackageOk() (*string, bool)`

GetQuantityPerPackageOk returns a tuple with the QuantityPerPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityPerPackage

`func (o *Price) SetQuantityPerPackage(v string)`

SetQuantityPerPackage sets QuantityPerPackage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


