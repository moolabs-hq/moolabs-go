# InvoiceTotals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | The total value of the line before taxes, discounts and commitments. | [readonly] 
**ChargesTotal** | **string** | The amount of value of the line that are due to additional charges. | [readonly] 
**DiscountsTotal** | **string** | The amount of value of the line that are due to discounts. | [readonly] 
**TaxesInclusiveTotal** | **string** | The total amount of taxes that are included in the line. | [readonly] 
**TaxesExclusiveTotal** | **string** | The total amount of taxes that are added on top of amount from the line. | [readonly] 
**TaxesTotal** | **string** | The total amount of taxes for this line. | [readonly] 
**Total** | **string** | The total amount value of the line after taxes, discounts and commitments. | [readonly] 

## Methods

### NewInvoiceTotals

`func NewInvoiceTotals(amount string, chargesTotal string, discountsTotal string, taxesInclusiveTotal string, taxesExclusiveTotal string, taxesTotal string, total string, ) *InvoiceTotals`

NewInvoiceTotals instantiates a new InvoiceTotals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceTotalsWithDefaults

`func NewInvoiceTotalsWithDefaults() *InvoiceTotals`

NewInvoiceTotalsWithDefaults instantiates a new InvoiceTotals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *InvoiceTotals) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InvoiceTotals) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InvoiceTotals) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetChargesTotal

`func (o *InvoiceTotals) GetChargesTotal() string`

GetChargesTotal returns the ChargesTotal field if non-nil, zero value otherwise.

### GetChargesTotalOk

`func (o *InvoiceTotals) GetChargesTotalOk() (*string, bool)`

GetChargesTotalOk returns a tuple with the ChargesTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargesTotal

`func (o *InvoiceTotals) SetChargesTotal(v string)`

SetChargesTotal sets ChargesTotal field to given value.


### GetDiscountsTotal

`func (o *InvoiceTotals) GetDiscountsTotal() string`

GetDiscountsTotal returns the DiscountsTotal field if non-nil, zero value otherwise.

### GetDiscountsTotalOk

`func (o *InvoiceTotals) GetDiscountsTotalOk() (*string, bool)`

GetDiscountsTotalOk returns a tuple with the DiscountsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountsTotal

`func (o *InvoiceTotals) SetDiscountsTotal(v string)`

SetDiscountsTotal sets DiscountsTotal field to given value.


### GetTaxesInclusiveTotal

`func (o *InvoiceTotals) GetTaxesInclusiveTotal() string`

GetTaxesInclusiveTotal returns the TaxesInclusiveTotal field if non-nil, zero value otherwise.

### GetTaxesInclusiveTotalOk

`func (o *InvoiceTotals) GetTaxesInclusiveTotalOk() (*string, bool)`

GetTaxesInclusiveTotalOk returns a tuple with the TaxesInclusiveTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxesInclusiveTotal

`func (o *InvoiceTotals) SetTaxesInclusiveTotal(v string)`

SetTaxesInclusiveTotal sets TaxesInclusiveTotal field to given value.


### GetTaxesExclusiveTotal

`func (o *InvoiceTotals) GetTaxesExclusiveTotal() string`

GetTaxesExclusiveTotal returns the TaxesExclusiveTotal field if non-nil, zero value otherwise.

### GetTaxesExclusiveTotalOk

`func (o *InvoiceTotals) GetTaxesExclusiveTotalOk() (*string, bool)`

GetTaxesExclusiveTotalOk returns a tuple with the TaxesExclusiveTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxesExclusiveTotal

`func (o *InvoiceTotals) SetTaxesExclusiveTotal(v string)`

SetTaxesExclusiveTotal sets TaxesExclusiveTotal field to given value.


### GetTaxesTotal

`func (o *InvoiceTotals) GetTaxesTotal() string`

GetTaxesTotal returns the TaxesTotal field if non-nil, zero value otherwise.

### GetTaxesTotalOk

`func (o *InvoiceTotals) GetTaxesTotalOk() (*string, bool)`

GetTaxesTotalOk returns a tuple with the TaxesTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxesTotal

`func (o *InvoiceTotals) SetTaxesTotal(v string)`

SetTaxesTotal sets TaxesTotal field to given value.


### GetTotal

`func (o *InvoiceTotals) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InvoiceTotals) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InvoiceTotals) SetTotal(v string)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


