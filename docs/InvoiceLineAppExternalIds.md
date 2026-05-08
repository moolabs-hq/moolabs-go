# InvoiceLineAppExternalIds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invoicing** | Pointer to **string** | The external ID of the invoice in the invoicing app if available. | [optional] [readonly] 
**Tax** | Pointer to **string** | The external ID of the invoice in the tax app if available. | [optional] [readonly] 

## Methods

### NewInvoiceLineAppExternalIds

`func NewInvoiceLineAppExternalIds() *InvoiceLineAppExternalIds`

NewInvoiceLineAppExternalIds instantiates a new InvoiceLineAppExternalIds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceLineAppExternalIdsWithDefaults

`func NewInvoiceLineAppExternalIdsWithDefaults() *InvoiceLineAppExternalIds`

NewInvoiceLineAppExternalIdsWithDefaults instantiates a new InvoiceLineAppExternalIds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoicing

`func (o *InvoiceLineAppExternalIds) GetInvoicing() string`

GetInvoicing returns the Invoicing field if non-nil, zero value otherwise.

### GetInvoicingOk

`func (o *InvoiceLineAppExternalIds) GetInvoicingOk() (*string, bool)`

GetInvoicingOk returns a tuple with the Invoicing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicing

`func (o *InvoiceLineAppExternalIds) SetInvoicing(v string)`

SetInvoicing sets Invoicing field to given value.

### HasInvoicing

`func (o *InvoiceLineAppExternalIds) HasInvoicing() bool`

HasInvoicing returns a boolean if a field has been set.

### GetTax

`func (o *InvoiceLineAppExternalIds) GetTax() string`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *InvoiceLineAppExternalIds) GetTaxOk() (*string, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *InvoiceLineAppExternalIds) SetTax(v string)`

SetTax sets Tax field to given value.

### HasTax

`func (o *InvoiceLineAppExternalIds) HasTax() bool`

HasTax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


