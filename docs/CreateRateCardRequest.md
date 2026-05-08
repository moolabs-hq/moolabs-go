# CreateRateCardRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Tenant identifier | 
**PlanId** | Pointer to **NullableString** |  | [optional] 
**FeatureKey** | **string** | Feature key | 
**Currency** | Pointer to **string** | Currency code (ISO 4217) | [optional] [default to "USD"]
**Version** | **string** | Rate card version | 
**EffectiveFrom** | **time.Time** | Effective from timestamp | 
**EffectiveTo** | Pointer to **NullableTime** |  | [optional] 
**PricingModel** | **map[string]interface{}** | Pricing model (JSON object) | 

## Methods

### NewCreateRateCardRequest

`func NewCreateRateCardRequest(tenantId string, featureKey string, version string, effectiveFrom time.Time, pricingModel map[string]interface{}, ) *CreateRateCardRequest`

NewCreateRateCardRequest instantiates a new CreateRateCardRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRateCardRequestWithDefaults

`func NewCreateRateCardRequestWithDefaults() *CreateRateCardRequest`

NewCreateRateCardRequestWithDefaults instantiates a new CreateRateCardRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *CreateRateCardRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CreateRateCardRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CreateRateCardRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetPlanId

`func (o *CreateRateCardRequest) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *CreateRateCardRequest) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *CreateRateCardRequest) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *CreateRateCardRequest) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### SetPlanIdNil

`func (o *CreateRateCardRequest) SetPlanIdNil(b bool)`

 SetPlanIdNil sets the value for PlanId to be an explicit nil

### UnsetPlanId
`func (o *CreateRateCardRequest) UnsetPlanId()`

UnsetPlanId ensures that no value is present for PlanId, not even an explicit nil
### GetFeatureKey

`func (o *CreateRateCardRequest) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *CreateRateCardRequest) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *CreateRateCardRequest) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.


### GetCurrency

`func (o *CreateRateCardRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateRateCardRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateRateCardRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreateRateCardRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetVersion

`func (o *CreateRateCardRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CreateRateCardRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CreateRateCardRequest) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetEffectiveFrom

`func (o *CreateRateCardRequest) GetEffectiveFrom() time.Time`

GetEffectiveFrom returns the EffectiveFrom field if non-nil, zero value otherwise.

### GetEffectiveFromOk

`func (o *CreateRateCardRequest) GetEffectiveFromOk() (*time.Time, bool)`

GetEffectiveFromOk returns a tuple with the EffectiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveFrom

`func (o *CreateRateCardRequest) SetEffectiveFrom(v time.Time)`

SetEffectiveFrom sets EffectiveFrom field to given value.


### GetEffectiveTo

`func (o *CreateRateCardRequest) GetEffectiveTo() time.Time`

GetEffectiveTo returns the EffectiveTo field if non-nil, zero value otherwise.

### GetEffectiveToOk

`func (o *CreateRateCardRequest) GetEffectiveToOk() (*time.Time, bool)`

GetEffectiveToOk returns a tuple with the EffectiveTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveTo

`func (o *CreateRateCardRequest) SetEffectiveTo(v time.Time)`

SetEffectiveTo sets EffectiveTo field to given value.

### HasEffectiveTo

`func (o *CreateRateCardRequest) HasEffectiveTo() bool`

HasEffectiveTo returns a boolean if a field has been set.

### SetEffectiveToNil

`func (o *CreateRateCardRequest) SetEffectiveToNil(b bool)`

 SetEffectiveToNil sets the value for EffectiveTo to be an explicit nil

### UnsetEffectiveTo
`func (o *CreateRateCardRequest) UnsetEffectiveTo()`

UnsetEffectiveTo ensures that no value is present for EffectiveTo, not even an explicit nil
### GetPricingModel

`func (o *CreateRateCardRequest) GetPricingModel() map[string]interface{}`

GetPricingModel returns the PricingModel field if non-nil, zero value otherwise.

### GetPricingModelOk

`func (o *CreateRateCardRequest) GetPricingModelOk() (*map[string]interface{}, bool)`

GetPricingModelOk returns a tuple with the PricingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingModel

`func (o *CreateRateCardRequest) SetPricingModel(v map[string]interface{})`

SetPricingModel sets PricingModel field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


