# Discounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Percentage** | Pointer to [**DiscountPercentage**](DiscountPercentage.md) | The percentage discount. | [optional] 
**Usage** | Pointer to [**DiscountUsage**](DiscountUsage.md) | The usage discount. | [optional] 

## Methods

### NewDiscounts

`func NewDiscounts() *Discounts`

NewDiscounts instantiates a new Discounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscountsWithDefaults

`func NewDiscountsWithDefaults() *Discounts`

NewDiscountsWithDefaults instantiates a new Discounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPercentage

`func (o *Discounts) GetPercentage() DiscountPercentage`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *Discounts) GetPercentageOk() (*DiscountPercentage, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *Discounts) SetPercentage(v DiscountPercentage)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *Discounts) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### GetUsage

`func (o *Discounts) GetUsage() DiscountUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *Discounts) GetUsageOk() (*DiscountUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *Discounts) SetUsage(v DiscountUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *Discounts) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


