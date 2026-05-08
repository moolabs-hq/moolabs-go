# InvoiceDetailedLineRateCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaxConfig** | Pointer to [**TaxConfig**](TaxConfig.md) | The tax config of the rate card. When undefined, the tax config of the feature or the default tax config of the plan is used. | [optional] 
**Price** | [**FlatPriceWithPaymentTerm**](FlatPriceWithPaymentTerm.md) | The price of the rate card. When null, the feature or service is free. | 
**Quantity** | Pointer to **string** | Quantity of the item being sold.  Default: 1 | [optional] 
**Discounts** | Pointer to [**BillingDiscounts**](BillingDiscounts.md) | The discounts that are applied to the line. | [optional] 

## Methods

### NewInvoiceDetailedLineRateCard

`func NewInvoiceDetailedLineRateCard(price FlatPriceWithPaymentTerm, ) *InvoiceDetailedLineRateCard`

NewInvoiceDetailedLineRateCard instantiates a new InvoiceDetailedLineRateCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceDetailedLineRateCardWithDefaults

`func NewInvoiceDetailedLineRateCardWithDefaults() *InvoiceDetailedLineRateCard`

NewInvoiceDetailedLineRateCardWithDefaults instantiates a new InvoiceDetailedLineRateCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaxConfig

`func (o *InvoiceDetailedLineRateCard) GetTaxConfig() TaxConfig`

GetTaxConfig returns the TaxConfig field if non-nil, zero value otherwise.

### GetTaxConfigOk

`func (o *InvoiceDetailedLineRateCard) GetTaxConfigOk() (*TaxConfig, bool)`

GetTaxConfigOk returns a tuple with the TaxConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxConfig

`func (o *InvoiceDetailedLineRateCard) SetTaxConfig(v TaxConfig)`

SetTaxConfig sets TaxConfig field to given value.

### HasTaxConfig

`func (o *InvoiceDetailedLineRateCard) HasTaxConfig() bool`

HasTaxConfig returns a boolean if a field has been set.

### GetPrice

`func (o *InvoiceDetailedLineRateCard) GetPrice() FlatPriceWithPaymentTerm`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *InvoiceDetailedLineRateCard) GetPriceOk() (*FlatPriceWithPaymentTerm, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *InvoiceDetailedLineRateCard) SetPrice(v FlatPriceWithPaymentTerm)`

SetPrice sets Price field to given value.


### GetQuantity

`func (o *InvoiceDetailedLineRateCard) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *InvoiceDetailedLineRateCard) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *InvoiceDetailedLineRateCard) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *InvoiceDetailedLineRateCard) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetDiscounts

`func (o *InvoiceDetailedLineRateCard) GetDiscounts() BillingDiscounts`

GetDiscounts returns the Discounts field if non-nil, zero value otherwise.

### GetDiscountsOk

`func (o *InvoiceDetailedLineRateCard) GetDiscountsOk() (*BillingDiscounts, bool)`

GetDiscountsOk returns a tuple with the Discounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscounts

`func (o *InvoiceDetailedLineRateCard) SetDiscounts(v BillingDiscounts)`

SetDiscounts sets Discounts field to given value.

### HasDiscounts

`func (o *InvoiceDetailedLineRateCard) HasDiscounts() bool`

HasDiscounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


