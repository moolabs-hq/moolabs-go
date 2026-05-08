# VoidInvoiceActionCreateItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Percentage** | **float64** | How much of the total line items to be voided? (e.g. 100% means all charges are voided) | 
**Action** | [**VoidInvoiceLineActionCreateItem**](VoidInvoiceLineActionCreateItem.md) | The action to take on the line items. | 

## Methods

### NewVoidInvoiceActionCreateItem

`func NewVoidInvoiceActionCreateItem(percentage float64, action VoidInvoiceLineActionCreateItem, ) *VoidInvoiceActionCreateItem`

NewVoidInvoiceActionCreateItem instantiates a new VoidInvoiceActionCreateItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoidInvoiceActionCreateItemWithDefaults

`func NewVoidInvoiceActionCreateItemWithDefaults() *VoidInvoiceActionCreateItem`

NewVoidInvoiceActionCreateItemWithDefaults instantiates a new VoidInvoiceActionCreateItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPercentage

`func (o *VoidInvoiceActionCreateItem) GetPercentage() float64`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *VoidInvoiceActionCreateItem) GetPercentageOk() (*float64, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *VoidInvoiceActionCreateItem) SetPercentage(v float64)`

SetPercentage sets Percentage field to given value.


### GetAction

`func (o *VoidInvoiceActionCreateItem) GetAction() VoidInvoiceLineActionCreateItem`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *VoidInvoiceActionCreateItem) GetActionOk() (*VoidInvoiceLineActionCreateItem, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *VoidInvoiceActionCreateItem) SetAction(v VoidInvoiceLineActionCreateItem)`

SetAction sets Action field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


