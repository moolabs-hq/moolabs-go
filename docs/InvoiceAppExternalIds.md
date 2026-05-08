# InvoiceAppExternalIds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invoicing** | Pointer to **string** | The external ID of the invoice in the invoicing app if available. | [optional] [readonly] 
**Tax** | Pointer to **string** | The external ID of the invoice in the tax app if available. | [optional] [readonly] 
**Payment** | Pointer to **string** | The external ID of the invoice in the payment app if available. | [optional] [readonly] 

## Methods

### NewInvoiceAppExternalIds

`func NewInvoiceAppExternalIds() *InvoiceAppExternalIds`

NewInvoiceAppExternalIds instantiates a new InvoiceAppExternalIds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceAppExternalIdsWithDefaults

`func NewInvoiceAppExternalIdsWithDefaults() *InvoiceAppExternalIds`

NewInvoiceAppExternalIdsWithDefaults instantiates a new InvoiceAppExternalIds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoicing

`func (o *InvoiceAppExternalIds) GetInvoicing() string`

GetInvoicing returns the Invoicing field if non-nil, zero value otherwise.

### GetInvoicingOk

`func (o *InvoiceAppExternalIds) GetInvoicingOk() (*string, bool)`

GetInvoicingOk returns a tuple with the Invoicing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicing

`func (o *InvoiceAppExternalIds) SetInvoicing(v string)`

SetInvoicing sets Invoicing field to given value.

### HasInvoicing

`func (o *InvoiceAppExternalIds) HasInvoicing() bool`

HasInvoicing returns a boolean if a field has been set.

### GetTax

`func (o *InvoiceAppExternalIds) GetTax() string`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *InvoiceAppExternalIds) GetTaxOk() (*string, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *InvoiceAppExternalIds) SetTax(v string)`

SetTax sets Tax field to given value.

### HasTax

`func (o *InvoiceAppExternalIds) HasTax() bool`

HasTax returns a boolean if a field has been set.

### GetPayment

`func (o *InvoiceAppExternalIds) GetPayment() string`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *InvoiceAppExternalIds) GetPaymentOk() (*string, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *InvoiceAppExternalIds) SetPayment(v string)`

SetPayment sets Payment field to given value.

### HasPayment

`func (o *InvoiceAppExternalIds) HasPayment() bool`

HasPayment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


