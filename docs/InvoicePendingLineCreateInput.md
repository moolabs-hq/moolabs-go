# InvoicePendingLineCreateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** | The currency of the lines to be created. | 
**Lines** | [**[]InvoicePendingLineCreate**](InvoicePendingLineCreate.md) | The lines to be created. | 

## Methods

### NewInvoicePendingLineCreateInput

`func NewInvoicePendingLineCreateInput(currency string, lines []InvoicePendingLineCreate, ) *InvoicePendingLineCreateInput`

NewInvoicePendingLineCreateInput instantiates a new InvoicePendingLineCreateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoicePendingLineCreateInputWithDefaults

`func NewInvoicePendingLineCreateInputWithDefaults() *InvoicePendingLineCreateInput`

NewInvoicePendingLineCreateInputWithDefaults instantiates a new InvoicePendingLineCreateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *InvoicePendingLineCreateInput) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InvoicePendingLineCreateInput) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InvoicePendingLineCreateInput) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetLines

`func (o *InvoicePendingLineCreateInput) GetLines() []InvoicePendingLineCreate`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *InvoicePendingLineCreateInput) GetLinesOk() (*[]InvoicePendingLineCreate, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *InvoicePendingLineCreateInput) SetLines(v []InvoicePendingLineCreate)`

SetLines sets Lines field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


