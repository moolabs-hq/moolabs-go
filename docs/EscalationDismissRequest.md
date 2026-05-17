# EscalationDismissRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | **string** |  | 
**Hard** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewEscalationDismissRequest

`func NewEscalationDismissRequest(reason string, ) *EscalationDismissRequest`

NewEscalationDismissRequest instantiates a new EscalationDismissRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEscalationDismissRequestWithDefaults

`func NewEscalationDismissRequestWithDefaults() *EscalationDismissRequest`

NewEscalationDismissRequestWithDefaults instantiates a new EscalationDismissRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *EscalationDismissRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *EscalationDismissRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *EscalationDismissRequest) SetReason(v string)`

SetReason sets Reason field to given value.


### GetHard

`func (o *EscalationDismissRequest) GetHard() bool`

GetHard returns the Hard field if non-nil, zero value otherwise.

### GetHardOk

`func (o *EscalationDismissRequest) GetHardOk() (*bool, bool)`

GetHardOk returns a tuple with the Hard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHard

`func (o *EscalationDismissRequest) SetHard(v bool)`

SetHard sets Hard field to given value.

### HasHard

`func (o *EscalationDismissRequest) HasHard() bool`

HasHard returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


