# InvoiceStatusDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Immutable** | **bool** | Is the invoice editable? | [readonly] 
**Failed** | **bool** | Is the invoice in a failed state? | [readonly] 
**ExtendedStatus** | **string** | Extended status information for the invoice. | [readonly] 
**AvailableActions** | [**InvoiceAvailableActions**](InvoiceAvailableActions.md) | The actions that can be performed on the invoice. | 

## Methods

### NewInvoiceStatusDetails

`func NewInvoiceStatusDetails(immutable bool, failed bool, extendedStatus string, availableActions InvoiceAvailableActions, ) *InvoiceStatusDetails`

NewInvoiceStatusDetails instantiates a new InvoiceStatusDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceStatusDetailsWithDefaults

`func NewInvoiceStatusDetailsWithDefaults() *InvoiceStatusDetails`

NewInvoiceStatusDetailsWithDefaults instantiates a new InvoiceStatusDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImmutable

`func (o *InvoiceStatusDetails) GetImmutable() bool`

GetImmutable returns the Immutable field if non-nil, zero value otherwise.

### GetImmutableOk

`func (o *InvoiceStatusDetails) GetImmutableOk() (*bool, bool)`

GetImmutableOk returns a tuple with the Immutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmutable

`func (o *InvoiceStatusDetails) SetImmutable(v bool)`

SetImmutable sets Immutable field to given value.


### GetFailed

`func (o *InvoiceStatusDetails) GetFailed() bool`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *InvoiceStatusDetails) GetFailedOk() (*bool, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *InvoiceStatusDetails) SetFailed(v bool)`

SetFailed sets Failed field to given value.


### GetExtendedStatus

`func (o *InvoiceStatusDetails) GetExtendedStatus() string`

GetExtendedStatus returns the ExtendedStatus field if non-nil, zero value otherwise.

### GetExtendedStatusOk

`func (o *InvoiceStatusDetails) GetExtendedStatusOk() (*string, bool)`

GetExtendedStatusOk returns a tuple with the ExtendedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedStatus

`func (o *InvoiceStatusDetails) SetExtendedStatus(v string)`

SetExtendedStatus sets ExtendedStatus field to given value.


### GetAvailableActions

`func (o *InvoiceStatusDetails) GetAvailableActions() InvoiceAvailableActions`

GetAvailableActions returns the AvailableActions field if non-nil, zero value otherwise.

### GetAvailableActionsOk

`func (o *InvoiceStatusDetails) GetAvailableActionsOk() (*InvoiceAvailableActions, bool)`

GetAvailableActionsOk returns a tuple with the AvailableActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableActions

`func (o *InvoiceStatusDetails) SetAvailableActions(v InvoiceAvailableActions)`

SetAvailableActions sets AvailableActions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


