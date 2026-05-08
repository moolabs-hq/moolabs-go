# RateCardResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**PlanId** | Pointer to **NullableString** |  | [optional] 
**FeatureKey** | **string** |  | 
**Currency** | **string** |  | 
**Version** | **string** |  | 
**EffectiveFrom** | **time.Time** |  | 
**EffectiveTo** | **NullableTime** |  | 
**PricingFingerprint** | **string** |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewRateCardResponse

`func NewRateCardResponse(id string, tenantId string, featureKey string, currency string, version string, effectiveFrom time.Time, effectiveTo NullableTime, pricingFingerprint string, createdAt time.Time, ) *RateCardResponse`

NewRateCardResponse instantiates a new RateCardResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateCardResponseWithDefaults

`func NewRateCardResponseWithDefaults() *RateCardResponse`

NewRateCardResponseWithDefaults instantiates a new RateCardResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RateCardResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RateCardResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RateCardResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *RateCardResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *RateCardResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *RateCardResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetPlanId

`func (o *RateCardResponse) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *RateCardResponse) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *RateCardResponse) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *RateCardResponse) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### SetPlanIdNil

`func (o *RateCardResponse) SetPlanIdNil(b bool)`

 SetPlanIdNil sets the value for PlanId to be an explicit nil

### UnsetPlanId
`func (o *RateCardResponse) UnsetPlanId()`

UnsetPlanId ensures that no value is present for PlanId, not even an explicit nil
### GetFeatureKey

`func (o *RateCardResponse) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *RateCardResponse) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *RateCardResponse) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.


### GetCurrency

`func (o *RateCardResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *RateCardResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *RateCardResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetVersion

`func (o *RateCardResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RateCardResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RateCardResponse) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetEffectiveFrom

`func (o *RateCardResponse) GetEffectiveFrom() time.Time`

GetEffectiveFrom returns the EffectiveFrom field if non-nil, zero value otherwise.

### GetEffectiveFromOk

`func (o *RateCardResponse) GetEffectiveFromOk() (*time.Time, bool)`

GetEffectiveFromOk returns a tuple with the EffectiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveFrom

`func (o *RateCardResponse) SetEffectiveFrom(v time.Time)`

SetEffectiveFrom sets EffectiveFrom field to given value.


### GetEffectiveTo

`func (o *RateCardResponse) GetEffectiveTo() time.Time`

GetEffectiveTo returns the EffectiveTo field if non-nil, zero value otherwise.

### GetEffectiveToOk

`func (o *RateCardResponse) GetEffectiveToOk() (*time.Time, bool)`

GetEffectiveToOk returns a tuple with the EffectiveTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveTo

`func (o *RateCardResponse) SetEffectiveTo(v time.Time)`

SetEffectiveTo sets EffectiveTo field to given value.


### SetEffectiveToNil

`func (o *RateCardResponse) SetEffectiveToNil(b bool)`

 SetEffectiveToNil sets the value for EffectiveTo to be an explicit nil

### UnsetEffectiveTo
`func (o *RateCardResponse) UnsetEffectiveTo()`

UnsetEffectiveTo ensures that no value is present for EffectiveTo, not even an explicit nil
### GetPricingFingerprint

`func (o *RateCardResponse) GetPricingFingerprint() string`

GetPricingFingerprint returns the PricingFingerprint field if non-nil, zero value otherwise.

### GetPricingFingerprintOk

`func (o *RateCardResponse) GetPricingFingerprintOk() (*string, bool)`

GetPricingFingerprintOk returns a tuple with the PricingFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingFingerprint

`func (o *RateCardResponse) SetPricingFingerprint(v string)`

SetPricingFingerprint sets PricingFingerprint field to given value.


### GetCreatedAt

`func (o *RateCardResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RateCardResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RateCardResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


