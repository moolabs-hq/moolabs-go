# InvoiceAvailableActionDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultingState** | **string** | The state the invoice will reach if the action is activated and all intermediate steps are successful.  For example advancing a draft_created invoice will result in a draft_manual_approval_needed invoice. | [readonly] 

## Methods

### NewInvoiceAvailableActionDetails

`func NewInvoiceAvailableActionDetails(resultingState string, ) *InvoiceAvailableActionDetails`

NewInvoiceAvailableActionDetails instantiates a new InvoiceAvailableActionDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceAvailableActionDetailsWithDefaults

`func NewInvoiceAvailableActionDetailsWithDefaults() *InvoiceAvailableActionDetails`

NewInvoiceAvailableActionDetailsWithDefaults instantiates a new InvoiceAvailableActionDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultingState

`func (o *InvoiceAvailableActionDetails) GetResultingState() string`

GetResultingState returns the ResultingState field if non-nil, zero value otherwise.

### GetResultingStateOk

`func (o *InvoiceAvailableActionDetails) GetResultingStateOk() (*string, bool)`

GetResultingStateOk returns a tuple with the ResultingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultingState

`func (o *InvoiceAvailableActionDetails) SetResultingState(v string)`

SetResultingState sets ResultingState field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


