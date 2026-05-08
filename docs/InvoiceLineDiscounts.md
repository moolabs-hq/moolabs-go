# InvoiceLineDiscounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to [**[]InvoiceLineAmountDiscount**](InvoiceLineAmountDiscount.md) | Amount based discounts applied to the line.  Amount based discounts are deduced from the total price of the line. | [optional] 
**Usage** | Pointer to [**[]InvoiceLineUsageDiscount**](InvoiceLineUsageDiscount.md) | Usage based discounts applied to the line.  Usage based discounts are deduced from the usage of the line before price calculations are applied. | [optional] 

## Methods

### NewInvoiceLineDiscounts

`func NewInvoiceLineDiscounts() *InvoiceLineDiscounts`

NewInvoiceLineDiscounts instantiates a new InvoiceLineDiscounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceLineDiscountsWithDefaults

`func NewInvoiceLineDiscountsWithDefaults() *InvoiceLineDiscounts`

NewInvoiceLineDiscountsWithDefaults instantiates a new InvoiceLineDiscounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *InvoiceLineDiscounts) GetAmount() []InvoiceLineAmountDiscount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InvoiceLineDiscounts) GetAmountOk() (*[]InvoiceLineAmountDiscount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InvoiceLineDiscounts) SetAmount(v []InvoiceLineAmountDiscount)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *InvoiceLineDiscounts) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetUsage

`func (o *InvoiceLineDiscounts) GetUsage() []InvoiceLineUsageDiscount`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *InvoiceLineDiscounts) GetUsageOk() (*[]InvoiceLineUsageDiscount, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *InvoiceLineDiscounts) SetUsage(v []InvoiceLineUsageDiscount)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *InvoiceLineDiscounts) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


