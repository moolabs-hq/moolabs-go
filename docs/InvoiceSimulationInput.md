# InvoiceSimulationInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **string** | The number of the invoice. | [optional] 
**Currency** | **string** | Currency for all invoice line items.  Multi currency invoices are not supported yet. | 
**Lines** | [**[]InvoiceSimulationLine**](InvoiceSimulationLine.md) | Lines to be included in the generated invoice. | 

## Methods

### NewInvoiceSimulationInput

`func NewInvoiceSimulationInput(currency string, lines []InvoiceSimulationLine, ) *InvoiceSimulationInput`

NewInvoiceSimulationInput instantiates a new InvoiceSimulationInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceSimulationInputWithDefaults

`func NewInvoiceSimulationInputWithDefaults() *InvoiceSimulationInput`

NewInvoiceSimulationInputWithDefaults instantiates a new InvoiceSimulationInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *InvoiceSimulationInput) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *InvoiceSimulationInput) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *InvoiceSimulationInput) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *InvoiceSimulationInput) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetCurrency

`func (o *InvoiceSimulationInput) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InvoiceSimulationInput) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InvoiceSimulationInput) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetLines

`func (o *InvoiceSimulationInput) GetLines() []InvoiceSimulationLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *InvoiceSimulationInput) GetLinesOk() (*[]InvoiceSimulationLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *InvoiceSimulationInput) SetLines(v []InvoiceSimulationLine)`

SetLines sets Lines field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


