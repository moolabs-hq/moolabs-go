# EscalateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | **string** |  | 
**Priority** | Pointer to **string** |  | [optional] [default to "medium"]

## Methods

### NewEscalateRequest

`func NewEscalateRequest(reason string, ) *EscalateRequest`

NewEscalateRequest instantiates a new EscalateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEscalateRequestWithDefaults

`func NewEscalateRequestWithDefaults() *EscalateRequest`

NewEscalateRequestWithDefaults instantiates a new EscalateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *EscalateRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *EscalateRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *EscalateRequest) SetReason(v string)`

SetReason sets Reason field to given value.


### GetPriority

`func (o *EscalateRequest) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *EscalateRequest) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *EscalateRequest) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *EscalateRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


