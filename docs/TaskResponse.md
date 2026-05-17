# TaskResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**CaseId** | Pointer to **string** |  | [optional] 
**TaskType** | **string** |  | 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**AssignedTo** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**Priority** | **string** |  | 
**DueAt** | Pointer to **time.Time** |  | [optional] 
**Context** | Pointer to **map[string]interface{}** |  | [optional] 
**Payload** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**ParentTaskId** | Pointer to **string** |  | [optional] 
**DraftVersion** | Pointer to **int32** |  | [optional] [default to 1]
**LockVersion** | Pointer to **int32** |  | [optional] [default to 1]
**CompletionKind** | Pointer to **string** |  | [optional] 
**PriorDrafts** | Pointer to **[]interface{}** |  | [optional] [default to []]
**AssociatedActivityId** | Pointer to **string** |  | [optional] 
**AssociatedActivityType** | Pointer to **string** |  | [optional] 
**CustomerName** | Pointer to **string** |  | [optional] 

## Methods

### NewTaskResponse

`func NewTaskResponse(id string, tenantId string, taskType string, status string, priority string, createdAt time.Time, updatedAt time.Time, ) *TaskResponse`

NewTaskResponse instantiates a new TaskResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskResponseWithDefaults

`func NewTaskResponseWithDefaults() *TaskResponse`

NewTaskResponseWithDefaults instantiates a new TaskResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaskResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *TaskResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TaskResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TaskResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCaseId

`func (o *TaskResponse) GetCaseId() string`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *TaskResponse) GetCaseIdOk() (*string, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *TaskResponse) SetCaseId(v string)`

SetCaseId sets CaseId field to given value.

### HasCaseId

`func (o *TaskResponse) HasCaseId() bool`

HasCaseId returns a boolean if a field has been set.

### GetTaskType

`func (o *TaskResponse) GetTaskType() string`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *TaskResponse) GetTaskTypeOk() (*string, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *TaskResponse) SetTaskType(v string)`

SetTaskType sets TaskType field to given value.


### GetTitle

`func (o *TaskResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TaskResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TaskResponse) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TaskResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *TaskResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TaskResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TaskResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TaskResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAssignedTo

`func (o *TaskResponse) GetAssignedTo() string`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *TaskResponse) GetAssignedToOk() (*string, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTo

`func (o *TaskResponse) SetAssignedTo(v string)`

SetAssignedTo sets AssignedTo field to given value.

### HasAssignedTo

`func (o *TaskResponse) HasAssignedTo() bool`

HasAssignedTo returns a boolean if a field has been set.

### GetStatus

`func (o *TaskResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPriority

`func (o *TaskResponse) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TaskResponse) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TaskResponse) SetPriority(v string)`

SetPriority sets Priority field to given value.


### GetDueAt

`func (o *TaskResponse) GetDueAt() time.Time`

GetDueAt returns the DueAt field if non-nil, zero value otherwise.

### GetDueAtOk

`func (o *TaskResponse) GetDueAtOk() (*time.Time, bool)`

GetDueAtOk returns a tuple with the DueAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueAt

`func (o *TaskResponse) SetDueAt(v time.Time)`

SetDueAt sets DueAt field to given value.

### HasDueAt

`func (o *TaskResponse) HasDueAt() bool`

HasDueAt returns a boolean if a field has been set.

### GetContext

`func (o *TaskResponse) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *TaskResponse) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *TaskResponse) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *TaskResponse) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetPayload

`func (o *TaskResponse) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *TaskResponse) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *TaskResponse) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *TaskResponse) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TaskResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TaskResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TaskResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *TaskResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TaskResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TaskResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetParentTaskId

`func (o *TaskResponse) GetParentTaskId() string`

GetParentTaskId returns the ParentTaskId field if non-nil, zero value otherwise.

### GetParentTaskIdOk

`func (o *TaskResponse) GetParentTaskIdOk() (*string, bool)`

GetParentTaskIdOk returns a tuple with the ParentTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTaskId

`func (o *TaskResponse) SetParentTaskId(v string)`

SetParentTaskId sets ParentTaskId field to given value.

### HasParentTaskId

`func (o *TaskResponse) HasParentTaskId() bool`

HasParentTaskId returns a boolean if a field has been set.

### GetDraftVersion

`func (o *TaskResponse) GetDraftVersion() int32`

GetDraftVersion returns the DraftVersion field if non-nil, zero value otherwise.

### GetDraftVersionOk

`func (o *TaskResponse) GetDraftVersionOk() (*int32, bool)`

GetDraftVersionOk returns a tuple with the DraftVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftVersion

`func (o *TaskResponse) SetDraftVersion(v int32)`

SetDraftVersion sets DraftVersion field to given value.

### HasDraftVersion

`func (o *TaskResponse) HasDraftVersion() bool`

HasDraftVersion returns a boolean if a field has been set.

### GetLockVersion

`func (o *TaskResponse) GetLockVersion() int32`

GetLockVersion returns the LockVersion field if non-nil, zero value otherwise.

### GetLockVersionOk

`func (o *TaskResponse) GetLockVersionOk() (*int32, bool)`

GetLockVersionOk returns a tuple with the LockVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockVersion

`func (o *TaskResponse) SetLockVersion(v int32)`

SetLockVersion sets LockVersion field to given value.

### HasLockVersion

`func (o *TaskResponse) HasLockVersion() bool`

HasLockVersion returns a boolean if a field has been set.

### GetCompletionKind

`func (o *TaskResponse) GetCompletionKind() string`

GetCompletionKind returns the CompletionKind field if non-nil, zero value otherwise.

### GetCompletionKindOk

`func (o *TaskResponse) GetCompletionKindOk() (*string, bool)`

GetCompletionKindOk returns a tuple with the CompletionKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionKind

`func (o *TaskResponse) SetCompletionKind(v string)`

SetCompletionKind sets CompletionKind field to given value.

### HasCompletionKind

`func (o *TaskResponse) HasCompletionKind() bool`

HasCompletionKind returns a boolean if a field has been set.

### GetPriorDrafts

`func (o *TaskResponse) GetPriorDrafts() []interface{}`

GetPriorDrafts returns the PriorDrafts field if non-nil, zero value otherwise.

### GetPriorDraftsOk

`func (o *TaskResponse) GetPriorDraftsOk() (*[]interface{}, bool)`

GetPriorDraftsOk returns a tuple with the PriorDrafts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorDrafts

`func (o *TaskResponse) SetPriorDrafts(v []interface{})`

SetPriorDrafts sets PriorDrafts field to given value.

### HasPriorDrafts

`func (o *TaskResponse) HasPriorDrafts() bool`

HasPriorDrafts returns a boolean if a field has been set.

### GetAssociatedActivityId

`func (o *TaskResponse) GetAssociatedActivityId() string`

GetAssociatedActivityId returns the AssociatedActivityId field if non-nil, zero value otherwise.

### GetAssociatedActivityIdOk

`func (o *TaskResponse) GetAssociatedActivityIdOk() (*string, bool)`

GetAssociatedActivityIdOk returns a tuple with the AssociatedActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedActivityId

`func (o *TaskResponse) SetAssociatedActivityId(v string)`

SetAssociatedActivityId sets AssociatedActivityId field to given value.

### HasAssociatedActivityId

`func (o *TaskResponse) HasAssociatedActivityId() bool`

HasAssociatedActivityId returns a boolean if a field has been set.

### GetAssociatedActivityType

`func (o *TaskResponse) GetAssociatedActivityType() string`

GetAssociatedActivityType returns the AssociatedActivityType field if non-nil, zero value otherwise.

### GetAssociatedActivityTypeOk

`func (o *TaskResponse) GetAssociatedActivityTypeOk() (*string, bool)`

GetAssociatedActivityTypeOk returns a tuple with the AssociatedActivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedActivityType

`func (o *TaskResponse) SetAssociatedActivityType(v string)`

SetAssociatedActivityType sets AssociatedActivityType field to given value.

### HasAssociatedActivityType

`func (o *TaskResponse) HasAssociatedActivityType() bool`

HasAssociatedActivityType returns a boolean if a field has been set.

### GetCustomerName

`func (o *TaskResponse) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *TaskResponse) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *TaskResponse) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.

### HasCustomerName

`func (o *TaskResponse) HasCustomerName() bool`

HasCustomerName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


