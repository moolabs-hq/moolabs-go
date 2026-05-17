# EscalationCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | **string** |  | 
**EscalatedTo** | Pointer to **string** |  | [optional] 
**EscalatedFrom** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **string** | low, medium, high, critical | [optional] [default to "medium"]
**SlaDeadline** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewEscalationCreate

`func NewEscalationCreate(reason string, ) *EscalationCreate`

NewEscalationCreate instantiates a new EscalationCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEscalationCreateWithDefaults

`func NewEscalationCreateWithDefaults() *EscalationCreate`

NewEscalationCreateWithDefaults instantiates a new EscalationCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *EscalationCreate) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *EscalationCreate) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *EscalationCreate) SetReason(v string)`

SetReason sets Reason field to given value.


### GetEscalatedTo

`func (o *EscalationCreate) GetEscalatedTo() string`

GetEscalatedTo returns the EscalatedTo field if non-nil, zero value otherwise.

### GetEscalatedToOk

`func (o *EscalationCreate) GetEscalatedToOk() (*string, bool)`

GetEscalatedToOk returns a tuple with the EscalatedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEscalatedTo

`func (o *EscalationCreate) SetEscalatedTo(v string)`

SetEscalatedTo sets EscalatedTo field to given value.

### HasEscalatedTo

`func (o *EscalationCreate) HasEscalatedTo() bool`

HasEscalatedTo returns a boolean if a field has been set.

### GetEscalatedFrom

`func (o *EscalationCreate) GetEscalatedFrom() string`

GetEscalatedFrom returns the EscalatedFrom field if non-nil, zero value otherwise.

### GetEscalatedFromOk

`func (o *EscalationCreate) GetEscalatedFromOk() (*string, bool)`

GetEscalatedFromOk returns a tuple with the EscalatedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEscalatedFrom

`func (o *EscalationCreate) SetEscalatedFrom(v string)`

SetEscalatedFrom sets EscalatedFrom field to given value.

### HasEscalatedFrom

`func (o *EscalationCreate) HasEscalatedFrom() bool`

HasEscalatedFrom returns a boolean if a field has been set.

### GetPriority

`func (o *EscalationCreate) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *EscalationCreate) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *EscalationCreate) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *EscalationCreate) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetSlaDeadline

`func (o *EscalationCreate) GetSlaDeadline() time.Time`

GetSlaDeadline returns the SlaDeadline field if non-nil, zero value otherwise.

### GetSlaDeadlineOk

`func (o *EscalationCreate) GetSlaDeadlineOk() (*time.Time, bool)`

GetSlaDeadlineOk returns a tuple with the SlaDeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaDeadline

`func (o *EscalationCreate) SetSlaDeadline(v time.Time)`

SetSlaDeadline sets SlaDeadline field to given value.

### HasSlaDeadline

`func (o *EscalationCreate) HasSlaDeadline() bool`

HasSlaDeadline returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


