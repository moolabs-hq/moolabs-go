# QuoteLineItemInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LineNo** | Pointer to **int32** |  | [optional] 
**SkuRef** | **string** |  | 
**FeatureKey** | Pointer to **string** |  | [optional] 
**SkuVersion** | Pointer to **string** |  | [optional] 
**RateCardRef** | Pointer to **string** |  | [optional] 
**PlanId** | Pointer to **string** |  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 
**PlanVersion** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **string** |  | [optional] [default to "1"]
**QuantityUnit** | Pointer to **string** |  | [optional] [default to "unit"]
**PricingBasis** | Pointer to **string** |  | [optional] [default to "recurring"]
**RequestedDiscount** | Pointer to **map[string]interface{}** |  | [optional] 
**FeatureOverrides** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewQuoteLineItemInput

`func NewQuoteLineItemInput(skuRef string, ) *QuoteLineItemInput`

NewQuoteLineItemInput instantiates a new QuoteLineItemInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteLineItemInputWithDefaults

`func NewQuoteLineItemInputWithDefaults() *QuoteLineItemInput`

NewQuoteLineItemInputWithDefaults instantiates a new QuoteLineItemInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLineNo

`func (o *QuoteLineItemInput) GetLineNo() int32`

GetLineNo returns the LineNo field if non-nil, zero value otherwise.

### GetLineNoOk

`func (o *QuoteLineItemInput) GetLineNoOk() (*int32, bool)`

GetLineNoOk returns a tuple with the LineNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNo

`func (o *QuoteLineItemInput) SetLineNo(v int32)`

SetLineNo sets LineNo field to given value.

### HasLineNo

`func (o *QuoteLineItemInput) HasLineNo() bool`

HasLineNo returns a boolean if a field has been set.

### GetSkuRef

`func (o *QuoteLineItemInput) GetSkuRef() string`

GetSkuRef returns the SkuRef field if non-nil, zero value otherwise.

### GetSkuRefOk

`func (o *QuoteLineItemInput) GetSkuRefOk() (*string, bool)`

GetSkuRefOk returns a tuple with the SkuRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuRef

`func (o *QuoteLineItemInput) SetSkuRef(v string)`

SetSkuRef sets SkuRef field to given value.


### GetFeatureKey

`func (o *QuoteLineItemInput) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *QuoteLineItemInput) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *QuoteLineItemInput) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.

### HasFeatureKey

`func (o *QuoteLineItemInput) HasFeatureKey() bool`

HasFeatureKey returns a boolean if a field has been set.

### GetSkuVersion

`func (o *QuoteLineItemInput) GetSkuVersion() string`

GetSkuVersion returns the SkuVersion field if non-nil, zero value otherwise.

### GetSkuVersionOk

`func (o *QuoteLineItemInput) GetSkuVersionOk() (*string, bool)`

GetSkuVersionOk returns a tuple with the SkuVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuVersion

`func (o *QuoteLineItemInput) SetSkuVersion(v string)`

SetSkuVersion sets SkuVersion field to given value.

### HasSkuVersion

`func (o *QuoteLineItemInput) HasSkuVersion() bool`

HasSkuVersion returns a boolean if a field has been set.

### GetRateCardRef

`func (o *QuoteLineItemInput) GetRateCardRef() string`

GetRateCardRef returns the RateCardRef field if non-nil, zero value otherwise.

### GetRateCardRefOk

`func (o *QuoteLineItemInput) GetRateCardRefOk() (*string, bool)`

GetRateCardRefOk returns a tuple with the RateCardRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateCardRef

`func (o *QuoteLineItemInput) SetRateCardRef(v string)`

SetRateCardRef sets RateCardRef field to given value.

### HasRateCardRef

`func (o *QuoteLineItemInput) HasRateCardRef() bool`

HasRateCardRef returns a boolean if a field has been set.

### GetPlanId

`func (o *QuoteLineItemInput) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *QuoteLineItemInput) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *QuoteLineItemInput) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *QuoteLineItemInput) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *QuoteLineItemInput) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *QuoteLineItemInput) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *QuoteLineItemInput) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *QuoteLineItemInput) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetPlanVersion

`func (o *QuoteLineItemInput) GetPlanVersion() string`

GetPlanVersion returns the PlanVersion field if non-nil, zero value otherwise.

### GetPlanVersionOk

`func (o *QuoteLineItemInput) GetPlanVersionOk() (*string, bool)`

GetPlanVersionOk returns a tuple with the PlanVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanVersion

`func (o *QuoteLineItemInput) SetPlanVersion(v string)`

SetPlanVersion sets PlanVersion field to given value.

### HasPlanVersion

`func (o *QuoteLineItemInput) HasPlanVersion() bool`

HasPlanVersion returns a boolean if a field has been set.

### GetQuantity

`func (o *QuoteLineItemInput) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *QuoteLineItemInput) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *QuoteLineItemInput) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *QuoteLineItemInput) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetQuantityUnit

`func (o *QuoteLineItemInput) GetQuantityUnit() string`

GetQuantityUnit returns the QuantityUnit field if non-nil, zero value otherwise.

### GetQuantityUnitOk

`func (o *QuoteLineItemInput) GetQuantityUnitOk() (*string, bool)`

GetQuantityUnitOk returns a tuple with the QuantityUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityUnit

`func (o *QuoteLineItemInput) SetQuantityUnit(v string)`

SetQuantityUnit sets QuantityUnit field to given value.

### HasQuantityUnit

`func (o *QuoteLineItemInput) HasQuantityUnit() bool`

HasQuantityUnit returns a boolean if a field has been set.

### GetPricingBasis

`func (o *QuoteLineItemInput) GetPricingBasis() string`

GetPricingBasis returns the PricingBasis field if non-nil, zero value otherwise.

### GetPricingBasisOk

`func (o *QuoteLineItemInput) GetPricingBasisOk() (*string, bool)`

GetPricingBasisOk returns a tuple with the PricingBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingBasis

`func (o *QuoteLineItemInput) SetPricingBasis(v string)`

SetPricingBasis sets PricingBasis field to given value.

### HasPricingBasis

`func (o *QuoteLineItemInput) HasPricingBasis() bool`

HasPricingBasis returns a boolean if a field has been set.

### GetRequestedDiscount

`func (o *QuoteLineItemInput) GetRequestedDiscount() map[string]interface{}`

GetRequestedDiscount returns the RequestedDiscount field if non-nil, zero value otherwise.

### GetRequestedDiscountOk

`func (o *QuoteLineItemInput) GetRequestedDiscountOk() (*map[string]interface{}, bool)`

GetRequestedDiscountOk returns a tuple with the RequestedDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedDiscount

`func (o *QuoteLineItemInput) SetRequestedDiscount(v map[string]interface{})`

SetRequestedDiscount sets RequestedDiscount field to given value.

### HasRequestedDiscount

`func (o *QuoteLineItemInput) HasRequestedDiscount() bool`

HasRequestedDiscount returns a boolean if a field has been set.

### GetFeatureOverrides

`func (o *QuoteLineItemInput) GetFeatureOverrides() []map[string]interface{}`

GetFeatureOverrides returns the FeatureOverrides field if non-nil, zero value otherwise.

### GetFeatureOverridesOk

`func (o *QuoteLineItemInput) GetFeatureOverridesOk() (*[]map[string]interface{}, bool)`

GetFeatureOverridesOk returns a tuple with the FeatureOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureOverrides

`func (o *QuoteLineItemInput) SetFeatureOverrides(v []map[string]interface{})`

SetFeatureOverrides sets FeatureOverrides field to given value.

### HasFeatureOverrides

`func (o *QuoteLineItemInput) HasFeatureOverrides() bool`

HasFeatureOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


