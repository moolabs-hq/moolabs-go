# RateCardCostRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**FeatureKey** | **string** |  | 
**PlanId** | Pointer to **string** |  | [optional] 
**Currency** | **string** |  | 
**UnitPriceMicros** | Pointer to **int32** |  | [optional] 
**CostUnitMicros** | Pointer to **int32** |  | [optional] 
**CostSource** | **string** |  | 
**MarginPct** | Pointer to **float32** |  | [optional] 
**PricingShape** | Pointer to **string** |  | [optional] 

## Methods

### NewRateCardCostRow

`func NewRateCardCostRow(id string, featureKey string, currency string, costSource string, ) *RateCardCostRow`

NewRateCardCostRow instantiates a new RateCardCostRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateCardCostRowWithDefaults

`func NewRateCardCostRowWithDefaults() *RateCardCostRow`

NewRateCardCostRowWithDefaults instantiates a new RateCardCostRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RateCardCostRow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RateCardCostRow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RateCardCostRow) SetId(v string)`

SetId sets Id field to given value.


### GetFeatureKey

`func (o *RateCardCostRow) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *RateCardCostRow) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *RateCardCostRow) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.


### GetPlanId

`func (o *RateCardCostRow) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *RateCardCostRow) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *RateCardCostRow) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *RateCardCostRow) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetCurrency

`func (o *RateCardCostRow) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *RateCardCostRow) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *RateCardCostRow) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetUnitPriceMicros

`func (o *RateCardCostRow) GetUnitPriceMicros() int32`

GetUnitPriceMicros returns the UnitPriceMicros field if non-nil, zero value otherwise.

### GetUnitPriceMicrosOk

`func (o *RateCardCostRow) GetUnitPriceMicrosOk() (*int32, bool)`

GetUnitPriceMicrosOk returns a tuple with the UnitPriceMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPriceMicros

`func (o *RateCardCostRow) SetUnitPriceMicros(v int32)`

SetUnitPriceMicros sets UnitPriceMicros field to given value.

### HasUnitPriceMicros

`func (o *RateCardCostRow) HasUnitPriceMicros() bool`

HasUnitPriceMicros returns a boolean if a field has been set.

### GetCostUnitMicros

`func (o *RateCardCostRow) GetCostUnitMicros() int32`

GetCostUnitMicros returns the CostUnitMicros field if non-nil, zero value otherwise.

### GetCostUnitMicrosOk

`func (o *RateCardCostRow) GetCostUnitMicrosOk() (*int32, bool)`

GetCostUnitMicrosOk returns a tuple with the CostUnitMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostUnitMicros

`func (o *RateCardCostRow) SetCostUnitMicros(v int32)`

SetCostUnitMicros sets CostUnitMicros field to given value.

### HasCostUnitMicros

`func (o *RateCardCostRow) HasCostUnitMicros() bool`

HasCostUnitMicros returns a boolean if a field has been set.

### GetCostSource

`func (o *RateCardCostRow) GetCostSource() string`

GetCostSource returns the CostSource field if non-nil, zero value otherwise.

### GetCostSourceOk

`func (o *RateCardCostRow) GetCostSourceOk() (*string, bool)`

GetCostSourceOk returns a tuple with the CostSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostSource

`func (o *RateCardCostRow) SetCostSource(v string)`

SetCostSource sets CostSource field to given value.


### GetMarginPct

`func (o *RateCardCostRow) GetMarginPct() float32`

GetMarginPct returns the MarginPct field if non-nil, zero value otherwise.

### GetMarginPctOk

`func (o *RateCardCostRow) GetMarginPctOk() (*float32, bool)`

GetMarginPctOk returns a tuple with the MarginPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginPct

`func (o *RateCardCostRow) SetMarginPct(v float32)`

SetMarginPct sets MarginPct field to given value.

### HasMarginPct

`func (o *RateCardCostRow) HasMarginPct() bool`

HasMarginPct returns a boolean if a field has been set.

### GetPricingShape

`func (o *RateCardCostRow) GetPricingShape() string`

GetPricingShape returns the PricingShape field if non-nil, zero value otherwise.

### GetPricingShapeOk

`func (o *RateCardCostRow) GetPricingShapeOk() (*string, bool)`

GetPricingShapeOk returns a tuple with the PricingShape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingShape

`func (o *RateCardCostRow) SetPricingShape(v string)`

SetPricingShape sets PricingShape field to given value.

### HasPricingShape

`func (o *RateCardCostRow) HasPricingShape() bool`

HasPricingShape returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


