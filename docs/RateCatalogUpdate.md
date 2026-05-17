# RateCatalogUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RatePerUnit** | Pointer to [**NullableRatePerUnit1**](RatePerUnit1.md) |  | [optional] 
**RateUnitScale** | Pointer to **int32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**SourceRef** | Pointer to **string** |  | [optional] 
**EffectiveFrom** | Pointer to **time.Time** |  | [optional] 
**EffectiveTo** | Pointer to **time.Time** |  | [optional] 
**PricingRules** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewRateCatalogUpdate

`func NewRateCatalogUpdate() *RateCatalogUpdate`

NewRateCatalogUpdate instantiates a new RateCatalogUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateCatalogUpdateWithDefaults

`func NewRateCatalogUpdateWithDefaults() *RateCatalogUpdate`

NewRateCatalogUpdateWithDefaults instantiates a new RateCatalogUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRatePerUnit

`func (o *RateCatalogUpdate) GetRatePerUnit() RatePerUnit1`

GetRatePerUnit returns the RatePerUnit field if non-nil, zero value otherwise.

### GetRatePerUnitOk

`func (o *RateCatalogUpdate) GetRatePerUnitOk() (*RatePerUnit1, bool)`

GetRatePerUnitOk returns a tuple with the RatePerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePerUnit

`func (o *RateCatalogUpdate) SetRatePerUnit(v RatePerUnit1)`

SetRatePerUnit sets RatePerUnit field to given value.

### HasRatePerUnit

`func (o *RateCatalogUpdate) HasRatePerUnit() bool`

HasRatePerUnit returns a boolean if a field has been set.

### SetRatePerUnitNil

`func (o *RateCatalogUpdate) SetRatePerUnitNil(b bool)`

 SetRatePerUnitNil sets the value for RatePerUnit to be an explicit nil

### UnsetRatePerUnit
`func (o *RateCatalogUpdate) UnsetRatePerUnit()`

UnsetRatePerUnit ensures that no value is present for RatePerUnit, not even an explicit nil
### GetRateUnitScale

`func (o *RateCatalogUpdate) GetRateUnitScale() int32`

GetRateUnitScale returns the RateUnitScale field if non-nil, zero value otherwise.

### GetRateUnitScaleOk

`func (o *RateCatalogUpdate) GetRateUnitScaleOk() (*int32, bool)`

GetRateUnitScaleOk returns a tuple with the RateUnitScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateUnitScale

`func (o *RateCatalogUpdate) SetRateUnitScale(v int32)`

SetRateUnitScale sets RateUnitScale field to given value.

### HasRateUnitScale

`func (o *RateCatalogUpdate) HasRateUnitScale() bool`

HasRateUnitScale returns a boolean if a field has been set.

### GetCurrency

`func (o *RateCatalogUpdate) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *RateCatalogUpdate) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *RateCatalogUpdate) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *RateCatalogUpdate) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetSource

`func (o *RateCatalogUpdate) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RateCatalogUpdate) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RateCatalogUpdate) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *RateCatalogUpdate) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceRef

`func (o *RateCatalogUpdate) GetSourceRef() string`

GetSourceRef returns the SourceRef field if non-nil, zero value otherwise.

### GetSourceRefOk

`func (o *RateCatalogUpdate) GetSourceRefOk() (*string, bool)`

GetSourceRefOk returns a tuple with the SourceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRef

`func (o *RateCatalogUpdate) SetSourceRef(v string)`

SetSourceRef sets SourceRef field to given value.

### HasSourceRef

`func (o *RateCatalogUpdate) HasSourceRef() bool`

HasSourceRef returns a boolean if a field has been set.

### GetEffectiveFrom

`func (o *RateCatalogUpdate) GetEffectiveFrom() time.Time`

GetEffectiveFrom returns the EffectiveFrom field if non-nil, zero value otherwise.

### GetEffectiveFromOk

`func (o *RateCatalogUpdate) GetEffectiveFromOk() (*time.Time, bool)`

GetEffectiveFromOk returns a tuple with the EffectiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveFrom

`func (o *RateCatalogUpdate) SetEffectiveFrom(v time.Time)`

SetEffectiveFrom sets EffectiveFrom field to given value.

### HasEffectiveFrom

`func (o *RateCatalogUpdate) HasEffectiveFrom() bool`

HasEffectiveFrom returns a boolean if a field has been set.

### GetEffectiveTo

`func (o *RateCatalogUpdate) GetEffectiveTo() time.Time`

GetEffectiveTo returns the EffectiveTo field if non-nil, zero value otherwise.

### GetEffectiveToOk

`func (o *RateCatalogUpdate) GetEffectiveToOk() (*time.Time, bool)`

GetEffectiveToOk returns a tuple with the EffectiveTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveTo

`func (o *RateCatalogUpdate) SetEffectiveTo(v time.Time)`

SetEffectiveTo sets EffectiveTo field to given value.

### HasEffectiveTo

`func (o *RateCatalogUpdate) HasEffectiveTo() bool`

HasEffectiveTo returns a boolean if a field has been set.

### GetPricingRules

`func (o *RateCatalogUpdate) GetPricingRules() map[string]interface{}`

GetPricingRules returns the PricingRules field if non-nil, zero value otherwise.

### GetPricingRulesOk

`func (o *RateCatalogUpdate) GetPricingRulesOk() (*map[string]interface{}, bool)`

GetPricingRulesOk returns a tuple with the PricingRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingRules

`func (o *RateCatalogUpdate) SetPricingRules(v map[string]interface{})`

SetPricingRules sets PricingRules field to given value.

### HasPricingRules

`func (o *RateCatalogUpdate) HasPricingRules() bool`

HasPricingRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


