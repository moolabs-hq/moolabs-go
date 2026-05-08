# VoidInvoiceActionLineOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LineId** | **string** | The line item ID to override. | 
**Action** | [**VoidInvoiceActionCreateItem**](VoidInvoiceActionCreateItem.md) | The action to take on the line item. | 

## Methods

### NewVoidInvoiceActionLineOverride

`func NewVoidInvoiceActionLineOverride(lineId string, action VoidInvoiceActionCreateItem, ) *VoidInvoiceActionLineOverride`

NewVoidInvoiceActionLineOverride instantiates a new VoidInvoiceActionLineOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoidInvoiceActionLineOverrideWithDefaults

`func NewVoidInvoiceActionLineOverrideWithDefaults() *VoidInvoiceActionLineOverride`

NewVoidInvoiceActionLineOverrideWithDefaults instantiates a new VoidInvoiceActionLineOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLineId

`func (o *VoidInvoiceActionLineOverride) GetLineId() string`

GetLineId returns the LineId field if non-nil, zero value otherwise.

### GetLineIdOk

`func (o *VoidInvoiceActionLineOverride) GetLineIdOk() (*string, bool)`

GetLineIdOk returns a tuple with the LineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineId

`func (o *VoidInvoiceActionLineOverride) SetLineId(v string)`

SetLineId sets LineId field to given value.


### GetAction

`func (o *VoidInvoiceActionLineOverride) GetAction() VoidInvoiceActionCreateItem`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *VoidInvoiceActionLineOverride) GetActionOk() (*VoidInvoiceActionCreateItem, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *VoidInvoiceActionLineOverride) SetAction(v VoidInvoiceActionCreateItem)`

SetAction sets Action field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


