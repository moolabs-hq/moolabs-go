# RateCatalogCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**Provider** | **string** |  | 
**Model** | **string** |  | 
**MetricType** | **string** |  | 
**Tier** | Pointer to **string** |  | [optional] [default to "default"]
**RatePerUnit** | [**RatePerUnit**](RatePerUnit.md) |  | 
**RateUnitScale** | Pointer to **int32** |  | [optional] [default to 1000000]
**Currency** | Pointer to **string** |  | [optional] [default to "USD"]
**Source** | Pointer to **string** |  | [optional] [default to "manual_override"]
**SourceRef** | Pointer to **string** |  | [optional] 
**EffectiveFrom** | **time.Time** |  | 
**EffectiveTo** | Pointer to **time.Time** |  | [optional] 
**PricingRules** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewRateCatalogCreate

`func NewRateCatalogCreate(tenantId string, provider string, model string, metricType string, ratePerUnit RatePerUnit, effectiveFrom time.Time, ) *RateCatalogCreate`

NewRateCatalogCreate instantiates a new RateCatalogCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateCatalogCreateWithDefaults

`func NewRateCatalogCreateWithDefaults() *RateCatalogCreate`

NewRateCatalogCreateWithDefaults instantiates a new RateCatalogCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *RateCatalogCreate) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *RateCatalogCreate) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *RateCatalogCreate) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetProvider

`func (o *RateCatalogCreate) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *RateCatalogCreate) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *RateCatalogCreate) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetModel

`func (o *RateCatalogCreate) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *RateCatalogCreate) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *RateCatalogCreate) SetModel(v string)`

SetModel sets Model field to given value.


### GetMetricType

`func (o *RateCatalogCreate) GetMetricType() string`

GetMetricType returns the MetricType field if non-nil, zero value otherwise.

### GetMetricTypeOk

`func (o *RateCatalogCreate) GetMetricTypeOk() (*string, bool)`

GetMetricTypeOk returns a tuple with the MetricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricType

`func (o *RateCatalogCreate) SetMetricType(v string)`

SetMetricType sets MetricType field to given value.


### GetTier

`func (o *RateCatalogCreate) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *RateCatalogCreate) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *RateCatalogCreate) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *RateCatalogCreate) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetRatePerUnit

`func (o *RateCatalogCreate) GetRatePerUnit() RatePerUnit`

GetRatePerUnit returns the RatePerUnit field if non-nil, zero value otherwise.

### GetRatePerUnitOk

`func (o *RateCatalogCreate) GetRatePerUnitOk() (*RatePerUnit, bool)`

GetRatePerUnitOk returns a tuple with the RatePerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePerUnit

`func (o *RateCatalogCreate) SetRatePerUnit(v RatePerUnit)`

SetRatePerUnit sets RatePerUnit field to given value.


### GetRateUnitScale

`func (o *RateCatalogCreate) GetRateUnitScale() int32`

GetRateUnitScale returns the RateUnitScale field if non-nil, zero value otherwise.

### GetRateUnitScaleOk

`func (o *RateCatalogCreate) GetRateUnitScaleOk() (*int32, bool)`

GetRateUnitScaleOk returns a tuple with the RateUnitScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateUnitScale

`func (o *RateCatalogCreate) SetRateUnitScale(v int32)`

SetRateUnitScale sets RateUnitScale field to given value.

### HasRateUnitScale

`func (o *RateCatalogCreate) HasRateUnitScale() bool`

HasRateUnitScale returns a boolean if a field has been set.

### GetCurrency

`func (o *RateCatalogCreate) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *RateCatalogCreate) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *RateCatalogCreate) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *RateCatalogCreate) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetSource

`func (o *RateCatalogCreate) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RateCatalogCreate) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RateCatalogCreate) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *RateCatalogCreate) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceRef

`func (o *RateCatalogCreate) GetSourceRef() string`

GetSourceRef returns the SourceRef field if non-nil, zero value otherwise.

### GetSourceRefOk

`func (o *RateCatalogCreate) GetSourceRefOk() (*string, bool)`

GetSourceRefOk returns a tuple with the SourceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRef

`func (o *RateCatalogCreate) SetSourceRef(v string)`

SetSourceRef sets SourceRef field to given value.

### HasSourceRef

`func (o *RateCatalogCreate) HasSourceRef() bool`

HasSourceRef returns a boolean if a field has been set.

### GetEffectiveFrom

`func (o *RateCatalogCreate) GetEffectiveFrom() time.Time`

GetEffectiveFrom returns the EffectiveFrom field if non-nil, zero value otherwise.

### GetEffectiveFromOk

`func (o *RateCatalogCreate) GetEffectiveFromOk() (*time.Time, bool)`

GetEffectiveFromOk returns a tuple with the EffectiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveFrom

`func (o *RateCatalogCreate) SetEffectiveFrom(v time.Time)`

SetEffectiveFrom sets EffectiveFrom field to given value.


### GetEffectiveTo

`func (o *RateCatalogCreate) GetEffectiveTo() time.Time`

GetEffectiveTo returns the EffectiveTo field if non-nil, zero value otherwise.

### GetEffectiveToOk

`func (o *RateCatalogCreate) GetEffectiveToOk() (*time.Time, bool)`

GetEffectiveToOk returns a tuple with the EffectiveTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveTo

`func (o *RateCatalogCreate) SetEffectiveTo(v time.Time)`

SetEffectiveTo sets EffectiveTo field to given value.

### HasEffectiveTo

`func (o *RateCatalogCreate) HasEffectiveTo() bool`

HasEffectiveTo returns a boolean if a field has been set.

### GetPricingRules

`func (o *RateCatalogCreate) GetPricingRules() map[string]interface{}`

GetPricingRules returns the PricingRules field if non-nil, zero value otherwise.

### GetPricingRulesOk

`func (o *RateCatalogCreate) GetPricingRulesOk() (*map[string]interface{}, bool)`

GetPricingRulesOk returns a tuple with the PricingRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingRules

`func (o *RateCatalogCreate) SetPricingRules(v map[string]interface{})`

SetPricingRules sets PricingRules field to given value.

### HasPricingRules

`func (o *RateCatalogCreate) HasPricingRules() bool`

HasPricingRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


