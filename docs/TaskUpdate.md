# TaskUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignedTo** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **string** |  | [optional] 
**DueAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewTaskUpdate

`func NewTaskUpdate() *TaskUpdate`

NewTaskUpdate instantiates a new TaskUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskUpdateWithDefaults

`func NewTaskUpdateWithDefaults() *TaskUpdate`

NewTaskUpdateWithDefaults instantiates a new TaskUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignedTo

`func (o *TaskUpdate) GetAssignedTo() string`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *TaskUpdate) GetAssignedToOk() (*string, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTo

`func (o *TaskUpdate) SetAssignedTo(v string)`

SetAssignedTo sets AssignedTo field to given value.

### HasAssignedTo

`func (o *TaskUpdate) HasAssignedTo() bool`

HasAssignedTo returns a boolean if a field has been set.

### GetPriority

`func (o *TaskUpdate) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TaskUpdate) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TaskUpdate) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *TaskUpdate) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetDueAt

`func (o *TaskUpdate) GetDueAt() time.Time`

GetDueAt returns the DueAt field if non-nil, zero value otherwise.

### GetDueAtOk

`func (o *TaskUpdate) GetDueAtOk() (*time.Time, bool)`

GetDueAtOk returns a tuple with the DueAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueAt

`func (o *TaskUpdate) SetDueAt(v time.Time)`

SetDueAt sets DueAt field to given value.

### HasDueAt

`func (o *TaskUpdate) HasDueAt() bool`

HasDueAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


