# BillingDiscounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Percentage** | Pointer to [**BillingDiscountPercentage**](BillingDiscountPercentage.md) | The percentage discount. | [optional] 
**Usage** | Pointer to [**BillingDiscountUsage**](BillingDiscountUsage.md) | The usage discount. | [optional] 

## Methods

### NewBillingDiscounts

`func NewBillingDiscounts() *BillingDiscounts`

NewBillingDiscounts instantiates a new BillingDiscounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingDiscountsWithDefaults

`func NewBillingDiscountsWithDefaults() *BillingDiscounts`

NewBillingDiscountsWithDefaults instantiates a new BillingDiscounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPercentage

`func (o *BillingDiscounts) GetPercentage() BillingDiscountPercentage`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *BillingDiscounts) GetPercentageOk() (*BillingDiscountPercentage, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *BillingDiscounts) SetPercentage(v BillingDiscountPercentage)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *BillingDiscounts) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### GetUsage

`func (o *BillingDiscounts) GetUsage() BillingDiscountUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *BillingDiscounts) GetUsageOk() (*BillingDiscountUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *BillingDiscounts) SetUsage(v BillingDiscountUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *BillingDiscounts) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


