# RateCardUsageBasedPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the price. | 
**Amount** | **string** | The price of one package. | 
**PaymentTerm** | Pointer to [**PricePaymentTerm**](PricePaymentTerm.md) | The payment term of the flat price. Defaults to in advance. | [optional] 
**MinimumAmount** | Pointer to **string** | The customer is committed to spend at least the amount. | [optional] 
**MaximumAmount** | Pointer to **string** | The customer is limited to spend at most the amount. | [optional] 
**Mode** | [**TieredPriceMode**](TieredPriceMode.md) | Defines if the tiering mode is volume-based or graduated: - In &#x60;volume&#x60;-based tiering, the maximum quantity within a period determines the per unit price. - In &#x60;graduated&#x60; tiering, pricing can change as the quantity grows. | 
**Tiers** | [**[]PriceTier**](PriceTier.md) | The tiers of the tiered price. At least one price component is required in each tier. | 
**Multiplier** | Pointer to **string** | The multiplier to apply to the base price to get the dynamic price.  Examples: - 0.0: the price is zero - 0.5: the price is 50% of the base price - 1.0: the price is the same as the base price - 1.5: the price is 150% of the base price | [optional] 
**QuantityPerPackage** | **string** | The quantity per package. | 

## Methods

### NewRateCardUsageBasedPrice

`func NewRateCardUsageBasedPrice(type_ string, amount string, mode TieredPriceMode, tiers []PriceTier, quantityPerPackage string, ) *RateCardUsageBasedPrice`

NewRateCardUsageBasedPrice instantiates a new RateCardUsageBasedPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateCardUsageBasedPriceWithDefaults

`func NewRateCardUsageBasedPriceWithDefaults() *RateCardUsageBasedPrice`

NewRateCardUsageBasedPriceWithDefaults instantiates a new RateCardUsageBasedPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RateCardUsageBasedPrice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RateCardUsageBasedPrice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RateCardUsageBasedPrice) SetType(v string)`

SetType sets Type field to given value.


### GetAmount

`func (o *RateCardUsageBasedPrice) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RateCardUsageBasedPrice) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RateCardUsageBasedPrice) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetPaymentTerm

`func (o *RateCardUsageBasedPrice) GetPaymentTerm() PricePaymentTerm`

GetPaymentTerm returns the PaymentTerm field if non-nil, zero value otherwise.

### GetPaymentTermOk

`func (o *RateCardUsageBasedPrice) GetPaymentTermOk() (*PricePaymentTerm, bool)`

GetPaymentTermOk returns a tuple with the PaymentTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTerm

`func (o *RateCardUsageBasedPrice) SetPaymentTerm(v PricePaymentTerm)`

SetPaymentTerm sets PaymentTerm field to given value.

### HasPaymentTerm

`func (o *RateCardUsageBasedPrice) HasPaymentTerm() bool`

HasPaymentTerm returns a boolean if a field has been set.

### GetMinimumAmount

`func (o *RateCardUsageBasedPrice) GetMinimumAmount() string`

GetMinimumAmount returns the MinimumAmount field if non-nil, zero value otherwise.

### GetMinimumAmountOk

`func (o *RateCardUsageBasedPrice) GetMinimumAmountOk() (*string, bool)`

GetMinimumAmountOk returns a tuple with the MinimumAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumAmount

`func (o *RateCardUsageBasedPrice) SetMinimumAmount(v string)`

SetMinimumAmount sets MinimumAmount field to given value.

### HasMinimumAmount

`func (o *RateCardUsageBasedPrice) HasMinimumAmount() bool`

HasMinimumAmount returns a boolean if a field has been set.

### GetMaximumAmount

`func (o *RateCardUsageBasedPrice) GetMaximumAmount() string`

GetMaximumAmount returns the MaximumAmount field if non-nil, zero value otherwise.

### GetMaximumAmountOk

`func (o *RateCardUsageBasedPrice) GetMaximumAmountOk() (*string, bool)`

GetMaximumAmountOk returns a tuple with the MaximumAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumAmount

`func (o *RateCardUsageBasedPrice) SetMaximumAmount(v string)`

SetMaximumAmount sets MaximumAmount field to given value.

### HasMaximumAmount

`func (o *RateCardUsageBasedPrice) HasMaximumAmount() bool`

HasMaximumAmount returns a boolean if a field has been set.

### GetMode

`func (o *RateCardUsageBasedPrice) GetMode() TieredPriceMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *RateCardUsageBasedPrice) GetModeOk() (*TieredPriceMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *RateCardUsageBasedPrice) SetMode(v TieredPriceMode)`

SetMode sets Mode field to given value.


### GetTiers

`func (o *RateCardUsageBasedPrice) GetTiers() []PriceTier`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *RateCardUsageBasedPrice) GetTiersOk() (*[]PriceTier, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *RateCardUsageBasedPrice) SetTiers(v []PriceTier)`

SetTiers sets Tiers field to given value.


### GetMultiplier

`func (o *RateCardUsageBasedPrice) GetMultiplier() string`

GetMultiplier returns the Multiplier field if non-nil, zero value otherwise.

### GetMultiplierOk

`func (o *RateCardUsageBasedPrice) GetMultiplierOk() (*string, bool)`

GetMultiplierOk returns a tuple with the Multiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplier

`func (o *RateCardUsageBasedPrice) SetMultiplier(v string)`

SetMultiplier sets Multiplier field to given value.

### HasMultiplier

`func (o *RateCardUsageBasedPrice) HasMultiplier() bool`

HasMultiplier returns a boolean if a field has been set.

### GetQuantityPerPackage

`func (o *RateCardUsageBasedPrice) GetQuantityPerPackage() string`

GetQuantityPerPackage returns the QuantityPerPackage field if non-nil, zero value otherwise.

### GetQuantityPerPackageOk

`func (o *RateCardUsageBasedPrice) GetQuantityPerPackageOk() (*string, bool)`

GetQuantityPerPackageOk returns a tuple with the QuantityPerPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityPerPackage

`func (o *RateCardUsageBasedPrice) SetQuantityPerPackage(v string)`

SetQuantityPerPackage sets QuantityPerPackage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


