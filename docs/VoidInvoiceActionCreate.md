# VoidInvoiceActionCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Percentage** | **float64** | How much of the total line items to be voided? (e.g. 100% means all charges are voided) | 
**Action** | [**VoidInvoiceLineActionCreate**](VoidInvoiceLineActionCreate.md) | The action to take on the line items. | 

## Methods

### NewVoidInvoiceActionCreate

`func NewVoidInvoiceActionCreate(percentage float64, action VoidInvoiceLineActionCreate, ) *VoidInvoiceActionCreate`

NewVoidInvoiceActionCreate instantiates a new VoidInvoiceActionCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoidInvoiceActionCreateWithDefaults

`func NewVoidInvoiceActionCreateWithDefaults() *VoidInvoiceActionCreate`

NewVoidInvoiceActionCreateWithDefaults instantiates a new VoidInvoiceActionCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPercentage

`func (o *VoidInvoiceActionCreate) GetPercentage() float64`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *VoidInvoiceActionCreate) GetPercentageOk() (*float64, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *VoidInvoiceActionCreate) SetPercentage(v float64)`

SetPercentage sets Percentage field to given value.


### GetAction

`func (o *VoidInvoiceActionCreate) GetAction() VoidInvoiceLineActionCreate`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *VoidInvoiceActionCreate) GetActionOk() (*VoidInvoiceLineActionCreate, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *VoidInvoiceActionCreate) SetAction(v VoidInvoiceLineActionCreate)`

SetAction sets Action field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


