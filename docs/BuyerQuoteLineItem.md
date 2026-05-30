# BuyerQuoteLineItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LineNo** | **int32** |  | 
**SkuRef** | **string** |  | 
**FeatureKey** | Pointer to **string** |  | [optional] 
**Quantity** | **string** |  | 
**QuantityUnit** | Pointer to **string** |  | [optional] 
**UnitPriceMicros** | Pointer to **int32** |  | [optional] 
**SubtotalMicros** | **int32** |  | 
**DiscountMicros** | **int32** |  | 
**TotalMicros** | **int32** |  | 

## Methods

### NewBuyerQuoteLineItem

`func NewBuyerQuoteLineItem(lineNo int32, skuRef string, quantity string, subtotalMicros int32, discountMicros int32, totalMicros int32, ) *BuyerQuoteLineItem`

NewBuyerQuoteLineItem instantiates a new BuyerQuoteLineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuyerQuoteLineItemWithDefaults

`func NewBuyerQuoteLineItemWithDefaults() *BuyerQuoteLineItem`

NewBuyerQuoteLineItemWithDefaults instantiates a new BuyerQuoteLineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLineNo

`func (o *BuyerQuoteLineItem) GetLineNo() int32`

GetLineNo returns the LineNo field if non-nil, zero value otherwise.

### GetLineNoOk

`func (o *BuyerQuoteLineItem) GetLineNoOk() (*int32, bool)`

GetLineNoOk returns a tuple with the LineNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNo

`func (o *BuyerQuoteLineItem) SetLineNo(v int32)`

SetLineNo sets LineNo field to given value.


### GetSkuRef

`func (o *BuyerQuoteLineItem) GetSkuRef() string`

GetSkuRef returns the SkuRef field if non-nil, zero value otherwise.

### GetSkuRefOk

`func (o *BuyerQuoteLineItem) GetSkuRefOk() (*string, bool)`

GetSkuRefOk returns a tuple with the SkuRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuRef

`func (o *BuyerQuoteLineItem) SetSkuRef(v string)`

SetSkuRef sets SkuRef field to given value.


### GetFeatureKey

`func (o *BuyerQuoteLineItem) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *BuyerQuoteLineItem) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *BuyerQuoteLineItem) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.

### HasFeatureKey

`func (o *BuyerQuoteLineItem) HasFeatureKey() bool`

HasFeatureKey returns a boolean if a field has been set.

### GetQuantity

`func (o *BuyerQuoteLineItem) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *BuyerQuoteLineItem) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *BuyerQuoteLineItem) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.


### GetQuantityUnit

`func (o *BuyerQuoteLineItem) GetQuantityUnit() string`

GetQuantityUnit returns the QuantityUnit field if non-nil, zero value otherwise.

### GetQuantityUnitOk

`func (o *BuyerQuoteLineItem) GetQuantityUnitOk() (*string, bool)`

GetQuantityUnitOk returns a tuple with the QuantityUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityUnit

`func (o *BuyerQuoteLineItem) SetQuantityUnit(v string)`

SetQuantityUnit sets QuantityUnit field to given value.

### HasQuantityUnit

`func (o *BuyerQuoteLineItem) HasQuantityUnit() bool`

HasQuantityUnit returns a boolean if a field has been set.

### GetUnitPriceMicros

`func (o *BuyerQuoteLineItem) GetUnitPriceMicros() int32`

GetUnitPriceMicros returns the UnitPriceMicros field if non-nil, zero value otherwise.

### GetUnitPriceMicrosOk

`func (o *BuyerQuoteLineItem) GetUnitPriceMicrosOk() (*int32, bool)`

GetUnitPriceMicrosOk returns a tuple with the UnitPriceMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPriceMicros

`func (o *BuyerQuoteLineItem) SetUnitPriceMicros(v int32)`

SetUnitPriceMicros sets UnitPriceMicros field to given value.

### HasUnitPriceMicros

`func (o *BuyerQuoteLineItem) HasUnitPriceMicros() bool`

HasUnitPriceMicros returns a boolean if a field has been set.

### GetSubtotalMicros

`func (o *BuyerQuoteLineItem) GetSubtotalMicros() int32`

GetSubtotalMicros returns the SubtotalMicros field if non-nil, zero value otherwise.

### GetSubtotalMicrosOk

`func (o *BuyerQuoteLineItem) GetSubtotalMicrosOk() (*int32, bool)`

GetSubtotalMicrosOk returns a tuple with the SubtotalMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalMicros

`func (o *BuyerQuoteLineItem) SetSubtotalMicros(v int32)`

SetSubtotalMicros sets SubtotalMicros field to given value.


### GetDiscountMicros

`func (o *BuyerQuoteLineItem) GetDiscountMicros() int32`

GetDiscountMicros returns the DiscountMicros field if non-nil, zero value otherwise.

### GetDiscountMicrosOk

`func (o *BuyerQuoteLineItem) GetDiscountMicrosOk() (*int32, bool)`

GetDiscountMicrosOk returns a tuple with the DiscountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountMicros

`func (o *BuyerQuoteLineItem) SetDiscountMicros(v int32)`

SetDiscountMicros sets DiscountMicros field to given value.


### GetTotalMicros

`func (o *BuyerQuoteLineItem) GetTotalMicros() int32`

GetTotalMicros returns the TotalMicros field if non-nil, zero value otherwise.

### GetTotalMicrosOk

`func (o *BuyerQuoteLineItem) GetTotalMicrosOk() (*int32, bool)`

GetTotalMicrosOk returns a tuple with the TotalMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMicros

`func (o *BuyerQuoteLineItem) SetTotalMicros(v int32)`

SetTotalMicros sets TotalMicros field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


