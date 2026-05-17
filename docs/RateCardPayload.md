# RateCardPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureKey** | **string** |  | 
**PricingType** | **string** |  | 
**Currency** | **string** |  | 
**Amount** | Pointer to [**NullableAmount1**](Amount1.md) |  | [optional] 
**TieredPricing** | Pointer to [**TieredPricingPayload**](TieredPricingPayload.md) |  | [optional] 

## Methods

### NewRateCardPayload

`func NewRateCardPayload(featureKey string, pricingType string, currency string, ) *RateCardPayload`

NewRateCardPayload instantiates a new RateCardPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateCardPayloadWithDefaults

`func NewRateCardPayloadWithDefaults() *RateCardPayload`

NewRateCardPayloadWithDefaults instantiates a new RateCardPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatureKey

`func (o *RateCardPayload) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *RateCardPayload) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *RateCardPayload) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.


### GetPricingType

`func (o *RateCardPayload) GetPricingType() string`

GetPricingType returns the PricingType field if non-nil, zero value otherwise.

### GetPricingTypeOk

`func (o *RateCardPayload) GetPricingTypeOk() (*string, bool)`

GetPricingTypeOk returns a tuple with the PricingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingType

`func (o *RateCardPayload) SetPricingType(v string)`

SetPricingType sets PricingType field to given value.


### GetCurrency

`func (o *RateCardPayload) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *RateCardPayload) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *RateCardPayload) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetAmount

`func (o *RateCardPayload) GetAmount() Amount1`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RateCardPayload) GetAmountOk() (*Amount1, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RateCardPayload) SetAmount(v Amount1)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *RateCardPayload) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *RateCardPayload) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *RateCardPayload) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetTieredPricing

`func (o *RateCardPayload) GetTieredPricing() TieredPricingPayload`

GetTieredPricing returns the TieredPricing field if non-nil, zero value otherwise.

### GetTieredPricingOk

`func (o *RateCardPayload) GetTieredPricingOk() (*TieredPricingPayload, bool)`

GetTieredPricingOk returns a tuple with the TieredPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTieredPricing

`func (o *RateCardPayload) SetTieredPricing(v TieredPricingPayload)`

SetTieredPricing sets TieredPricing field to given value.

### HasTieredPricing

`func (o *RateCardPayload) HasTieredPricing() bool`

HasTieredPricing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


