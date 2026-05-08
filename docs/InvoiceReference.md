# InvoiceReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the invoice. | [readonly] 
**Number** | Pointer to **string** | The number of the invoice. | [optional] [readonly] 

## Methods

### NewInvoiceReference

`func NewInvoiceReference(id string, ) *InvoiceReference`

NewInvoiceReference instantiates a new InvoiceReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceReferenceWithDefaults

`func NewInvoiceReferenceWithDefaults() *InvoiceReference`

NewInvoiceReferenceWithDefaults instantiates a new InvoiceReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceReference) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceReference) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceReference) SetId(v string)`

SetId sets Id field to given value.


### GetNumber

`func (o *InvoiceReference) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *InvoiceReference) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *InvoiceReference) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *InvoiceReference) HasNumber() bool`

HasNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


