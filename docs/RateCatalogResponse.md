# RateCatalogResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**Provider** | **string** |  | 
**Model** | **string** |  | 
**MetricType** | **string** |  | 
**Tier** | **string** |  | 
**RatePerUnit** | **float32** |  | 
**RateUnitScale** | **int32** |  | 
**Currency** | **string** |  | 
**Source** | **string** |  | 
**SourceRef** | **string** |  | 
**EffectiveFrom** | **time.Time** |  | 
**EffectiveTo** | **time.Time** |  | 
**PricingRules** | **map[string]interface{}** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewRateCatalogResponse

`func NewRateCatalogResponse(id string, tenantId string, provider string, model string, metricType string, tier string, ratePerUnit float32, rateUnitScale int32, currency string, source string, sourceRef string, effectiveFrom time.Time, effectiveTo time.Time, pricingRules map[string]interface{}, createdAt time.Time, updatedAt time.Time, ) *RateCatalogResponse`

NewRateCatalogResponse instantiates a new RateCatalogResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateCatalogResponseWithDefaults

`func NewRateCatalogResponseWithDefaults() *RateCatalogResponse`

NewRateCatalogResponseWithDefaults instantiates a new RateCatalogResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RateCatalogResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RateCatalogResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RateCatalogResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *RateCatalogResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *RateCatalogResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *RateCatalogResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetProvider

`func (o *RateCatalogResponse) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *RateCatalogResponse) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *RateCatalogResponse) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetModel

`func (o *RateCatalogResponse) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *RateCatalogResponse) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *RateCatalogResponse) SetModel(v string)`

SetModel sets Model field to given value.


### GetMetricType

`func (o *RateCatalogResponse) GetMetricType() string`

GetMetricType returns the MetricType field if non-nil, zero value otherwise.

### GetMetricTypeOk

`func (o *RateCatalogResponse) GetMetricTypeOk() (*string, bool)`

GetMetricTypeOk returns a tuple with the MetricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricType

`func (o *RateCatalogResponse) SetMetricType(v string)`

SetMetricType sets MetricType field to given value.


### GetTier

`func (o *RateCatalogResponse) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *RateCatalogResponse) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *RateCatalogResponse) SetTier(v string)`

SetTier sets Tier field to given value.


### GetRatePerUnit

`func (o *RateCatalogResponse) GetRatePerUnit() float32`

GetRatePerUnit returns the RatePerUnit field if non-nil, zero value otherwise.

### GetRatePerUnitOk

`func (o *RateCatalogResponse) GetRatePerUnitOk() (*float32, bool)`

GetRatePerUnitOk returns a tuple with the RatePerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePerUnit

`func (o *RateCatalogResponse) SetRatePerUnit(v float32)`

SetRatePerUnit sets RatePerUnit field to given value.


### GetRateUnitScale

`func (o *RateCatalogResponse) GetRateUnitScale() int32`

GetRateUnitScale returns the RateUnitScale field if non-nil, zero value otherwise.

### GetRateUnitScaleOk

`func (o *RateCatalogResponse) GetRateUnitScaleOk() (*int32, bool)`

GetRateUnitScaleOk returns a tuple with the RateUnitScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateUnitScale

`func (o *RateCatalogResponse) SetRateUnitScale(v int32)`

SetRateUnitScale sets RateUnitScale field to given value.


### GetCurrency

`func (o *RateCatalogResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *RateCatalogResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *RateCatalogResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetSource

`func (o *RateCatalogResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RateCatalogResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RateCatalogResponse) SetSource(v string)`

SetSource sets Source field to given value.


### GetSourceRef

`func (o *RateCatalogResponse) GetSourceRef() string`

GetSourceRef returns the SourceRef field if non-nil, zero value otherwise.

### GetSourceRefOk

`func (o *RateCatalogResponse) GetSourceRefOk() (*string, bool)`

GetSourceRefOk returns a tuple with the SourceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRef

`func (o *RateCatalogResponse) SetSourceRef(v string)`

SetSourceRef sets SourceRef field to given value.


### GetEffectiveFrom

`func (o *RateCatalogResponse) GetEffectiveFrom() time.Time`

GetEffectiveFrom returns the EffectiveFrom field if non-nil, zero value otherwise.

### GetEffectiveFromOk

`func (o *RateCatalogResponse) GetEffectiveFromOk() (*time.Time, bool)`

GetEffectiveFromOk returns a tuple with the EffectiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveFrom

`func (o *RateCatalogResponse) SetEffectiveFrom(v time.Time)`

SetEffectiveFrom sets EffectiveFrom field to given value.


### GetEffectiveTo

`func (o *RateCatalogResponse) GetEffectiveTo() time.Time`

GetEffectiveTo returns the EffectiveTo field if non-nil, zero value otherwise.

### GetEffectiveToOk

`func (o *RateCatalogResponse) GetEffectiveToOk() (*time.Time, bool)`

GetEffectiveToOk returns a tuple with the EffectiveTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveTo

`func (o *RateCatalogResponse) SetEffectiveTo(v time.Time)`

SetEffectiveTo sets EffectiveTo field to given value.


### GetPricingRules

`func (o *RateCatalogResponse) GetPricingRules() map[string]interface{}`

GetPricingRules returns the PricingRules field if non-nil, zero value otherwise.

### GetPricingRulesOk

`func (o *RateCatalogResponse) GetPricingRulesOk() (*map[string]interface{}, bool)`

GetPricingRulesOk returns a tuple with the PricingRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingRules

`func (o *RateCatalogResponse) SetPricingRules(v map[string]interface{})`

SetPricingRules sets PricingRules field to given value.


### GetCreatedAt

`func (o *RateCatalogResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RateCatalogResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RateCatalogResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *RateCatalogResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RateCatalogResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RateCatalogResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


