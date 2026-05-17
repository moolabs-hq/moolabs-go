# TaskCompleteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resolution** | Pointer to **string** |  | [optional] 
**Payload** | Pointer to **map[string]interface{}** |  | [optional] 
**DispatchPayload** | Pointer to **map[string]interface{}** | Structured input consumed by app.services.task_dispatch — shape depends on task_type (see task_dispatch handlers). | [optional] 

## Methods

### NewTaskCompleteRequest

`func NewTaskCompleteRequest() *TaskCompleteRequest`

NewTaskCompleteRequest instantiates a new TaskCompleteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskCompleteRequestWithDefaults

`func NewTaskCompleteRequestWithDefaults() *TaskCompleteRequest`

NewTaskCompleteRequestWithDefaults instantiates a new TaskCompleteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResolution

`func (o *TaskCompleteRequest) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *TaskCompleteRequest) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *TaskCompleteRequest) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *TaskCompleteRequest) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetPayload

`func (o *TaskCompleteRequest) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *TaskCompleteRequest) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *TaskCompleteRequest) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *TaskCompleteRequest) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetDispatchPayload

`func (o *TaskCompleteRequest) GetDispatchPayload() map[string]interface{}`

GetDispatchPayload returns the DispatchPayload field if non-nil, zero value otherwise.

### GetDispatchPayloadOk

`func (o *TaskCompleteRequest) GetDispatchPayloadOk() (*map[string]interface{}, bool)`

GetDispatchPayloadOk returns a tuple with the DispatchPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispatchPayload

`func (o *TaskCompleteRequest) SetDispatchPayload(v map[string]interface{})`

SetDispatchPayload sets DispatchPayload field to given value.

### HasDispatchPayload

`func (o *TaskCompleteRequest) HasDispatchPayload() bool`

HasDispatchPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


