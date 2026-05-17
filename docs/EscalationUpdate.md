# EscalationUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EscalatedTo** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **string** |  | [optional] 
**SlaDeadline** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewEscalationUpdate

`func NewEscalationUpdate() *EscalationUpdate`

NewEscalationUpdate instantiates a new EscalationUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEscalationUpdateWithDefaults

`func NewEscalationUpdateWithDefaults() *EscalationUpdate`

NewEscalationUpdateWithDefaults instantiates a new EscalationUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEscalatedTo

`func (o *EscalationUpdate) GetEscalatedTo() string`

GetEscalatedTo returns the EscalatedTo field if non-nil, zero value otherwise.

### GetEscalatedToOk

`func (o *EscalationUpdate) GetEscalatedToOk() (*string, bool)`

GetEscalatedToOk returns a tuple with the EscalatedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEscalatedTo

`func (o *EscalationUpdate) SetEscalatedTo(v string)`

SetEscalatedTo sets EscalatedTo field to given value.

### HasEscalatedTo

`func (o *EscalationUpdate) HasEscalatedTo() bool`

HasEscalatedTo returns a boolean if a field has been set.

### GetPriority

`func (o *EscalationUpdate) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *EscalationUpdate) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *EscalationUpdate) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *EscalationUpdate) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetSlaDeadline

`func (o *EscalationUpdate) GetSlaDeadline() time.Time`

GetSlaDeadline returns the SlaDeadline field if non-nil, zero value otherwise.

### GetSlaDeadlineOk

`func (o *EscalationUpdate) GetSlaDeadlineOk() (*time.Time, bool)`

GetSlaDeadlineOk returns a tuple with the SlaDeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaDeadline

`func (o *EscalationUpdate) SetSlaDeadline(v time.Time)`

SetSlaDeadline sets SlaDeadline field to given value.

### HasSlaDeadline

`func (o *EscalationUpdate) HasSlaDeadline() bool`

HasSlaDeadline returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


