# InvoicePendingLineCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lines** | [**[]InvoiceLine**](InvoiceLine.md) | The lines that were created. | 
**Invoice** | [**Invoice**](Invoice.md) | The invoice containing the created lines. | [readonly] 
**IsInvoiceNew** | **bool** | Whether the invoice was newly created. | [readonly] 

## Methods

### NewInvoicePendingLineCreateResponse

`func NewInvoicePendingLineCreateResponse(lines []InvoiceLine, invoice Invoice, isInvoiceNew bool, ) *InvoicePendingLineCreateResponse`

NewInvoicePendingLineCreateResponse instantiates a new InvoicePendingLineCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoicePendingLineCreateResponseWithDefaults

`func NewInvoicePendingLineCreateResponseWithDefaults() *InvoicePendingLineCreateResponse`

NewInvoicePendingLineCreateResponseWithDefaults instantiates a new InvoicePendingLineCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLines

`func (o *InvoicePendingLineCreateResponse) GetLines() []InvoiceLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *InvoicePendingLineCreateResponse) GetLinesOk() (*[]InvoiceLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *InvoicePendingLineCreateResponse) SetLines(v []InvoiceLine)`

SetLines sets Lines field to given value.


### GetInvoice

`func (o *InvoicePendingLineCreateResponse) GetInvoice() Invoice`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *InvoicePendingLineCreateResponse) GetInvoiceOk() (*Invoice, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *InvoicePendingLineCreateResponse) SetInvoice(v Invoice)`

SetInvoice sets Invoice field to given value.


### GetIsInvoiceNew

`func (o *InvoicePendingLineCreateResponse) GetIsInvoiceNew() bool`

GetIsInvoiceNew returns the IsInvoiceNew field if non-nil, zero value otherwise.

### GetIsInvoiceNewOk

`func (o *InvoicePendingLineCreateResponse) GetIsInvoiceNewOk() (*bool, bool)`

GetIsInvoiceNewOk returns a tuple with the IsInvoiceNew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInvoiceNew

`func (o *InvoicePendingLineCreateResponse) SetIsInvoiceNew(v bool)`

SetIsInvoiceNew sets IsInvoiceNew field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


