# VoidInvoiceActionInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**VoidInvoiceActionCreate**](VoidInvoiceActionCreate.md) | The action to take on the voided line items. | 
**Reason** | **string** | The reason for voiding the invoice. | 
**Overrides** | Pointer to [**[]VoidInvoiceActionLineOverride**](VoidInvoiceActionLineOverride.md) | Per line item overrides for the action.  If not specified, the &#x60;action&#x60; will be applied to all line items. | [optional] 

## Methods

### NewVoidInvoiceActionInput

`func NewVoidInvoiceActionInput(action VoidInvoiceActionCreate, reason string, ) *VoidInvoiceActionInput`

NewVoidInvoiceActionInput instantiates a new VoidInvoiceActionInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoidInvoiceActionInputWithDefaults

`func NewVoidInvoiceActionInputWithDefaults() *VoidInvoiceActionInput`

NewVoidInvoiceActionInputWithDefaults instantiates a new VoidInvoiceActionInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *VoidInvoiceActionInput) GetAction() VoidInvoiceActionCreate`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *VoidInvoiceActionInput) GetActionOk() (*VoidInvoiceActionCreate, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *VoidInvoiceActionInput) SetAction(v VoidInvoiceActionCreate)`

SetAction sets Action field to given value.


### GetReason

`func (o *VoidInvoiceActionInput) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *VoidInvoiceActionInput) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *VoidInvoiceActionInput) SetReason(v string)`

SetReason sets Reason field to given value.


### GetOverrides

`func (o *VoidInvoiceActionInput) GetOverrides() []VoidInvoiceActionLineOverride`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *VoidInvoiceActionInput) GetOverridesOk() (*[]VoidInvoiceActionLineOverride, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *VoidInvoiceActionInput) SetOverrides(v []VoidInvoiceActionLineOverride)`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *VoidInvoiceActionInput) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


