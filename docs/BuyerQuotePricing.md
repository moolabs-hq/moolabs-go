# BuyerQuotePricing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** |  | 
**SubtotalMicros** | **int32** |  | 
**DiscountMicros** | **int32** |  | 
**TotalMicros** | **int32** |  | 

## Methods

### NewBuyerQuotePricing

`func NewBuyerQuotePricing(currency string, subtotalMicros int32, discountMicros int32, totalMicros int32, ) *BuyerQuotePricing`

NewBuyerQuotePricing instantiates a new BuyerQuotePricing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuyerQuotePricingWithDefaults

`func NewBuyerQuotePricingWithDefaults() *BuyerQuotePricing`

NewBuyerQuotePricingWithDefaults instantiates a new BuyerQuotePricing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *BuyerQuotePricing) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BuyerQuotePricing) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BuyerQuotePricing) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetSubtotalMicros

`func (o *BuyerQuotePricing) GetSubtotalMicros() int32`

GetSubtotalMicros returns the SubtotalMicros field if non-nil, zero value otherwise.

### GetSubtotalMicrosOk

`func (o *BuyerQuotePricing) GetSubtotalMicrosOk() (*int32, bool)`

GetSubtotalMicrosOk returns a tuple with the SubtotalMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalMicros

`func (o *BuyerQuotePricing) SetSubtotalMicros(v int32)`

SetSubtotalMicros sets SubtotalMicros field to given value.


### GetDiscountMicros

`func (o *BuyerQuotePricing) GetDiscountMicros() int32`

GetDiscountMicros returns the DiscountMicros field if non-nil, zero value otherwise.

### GetDiscountMicrosOk

`func (o *BuyerQuotePricing) GetDiscountMicrosOk() (*int32, bool)`

GetDiscountMicrosOk returns a tuple with the DiscountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountMicros

`func (o *BuyerQuotePricing) SetDiscountMicros(v int32)`

SetDiscountMicros sets DiscountMicros field to given value.


### GetTotalMicros

`func (o *BuyerQuotePricing) GetTotalMicros() int32`

GetTotalMicros returns the TotalMicros field if non-nil, zero value otherwise.

### GetTotalMicrosOk

`func (o *BuyerQuotePricing) GetTotalMicrosOk() (*int32, bool)`

GetTotalMicrosOk returns a tuple with the TotalMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMicros

`func (o *BuyerQuotePricing) SetTotalMicros(v int32)`

SetTotalMicros sets TotalMicros field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


