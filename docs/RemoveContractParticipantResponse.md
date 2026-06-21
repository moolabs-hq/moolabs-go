# RemoveContractParticipantResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**ThreadId** | **string** |  | 
**Side** | **string** |  | 
**Role** | **string** |  | 
**Email** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**AddedByActorId** | **string** |  | 
**RemovedByActorId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewRemoveContractParticipantResponse

`func NewRemoveContractParticipantResponse(id string, tenantId string, threadId string, side string, role string, status string, addedByActorId string, ) *RemoveContractParticipantResponse`

NewRemoveContractParticipantResponse instantiates a new RemoveContractParticipantResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveContractParticipantResponseWithDefaults

`func NewRemoveContractParticipantResponseWithDefaults() *RemoveContractParticipantResponse`

NewRemoveContractParticipantResponseWithDefaults instantiates a new RemoveContractParticipantResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RemoveContractParticipantResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RemoveContractParticipantResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RemoveContractParticipantResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *RemoveContractParticipantResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *RemoveContractParticipantResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *RemoveContractParticipantResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetThreadId

`func (o *RemoveContractParticipantResponse) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *RemoveContractParticipantResponse) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *RemoveContractParticipantResponse) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.


### GetSide

`func (o *RemoveContractParticipantResponse) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *RemoveContractParticipantResponse) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *RemoveContractParticipantResponse) SetSide(v string)`

SetSide sets Side field to given value.


### GetRole

`func (o *RemoveContractParticipantResponse) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *RemoveContractParticipantResponse) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *RemoveContractParticipantResponse) SetRole(v string)`

SetRole sets Role field to given value.


### GetEmail

`func (o *RemoveContractParticipantResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *RemoveContractParticipantResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *RemoveContractParticipantResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *RemoveContractParticipantResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDisplayName

`func (o *RemoveContractParticipantResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *RemoveContractParticipantResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *RemoveContractParticipantResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *RemoveContractParticipantResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetStatus

`func (o *RemoveContractParticipantResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RemoveContractParticipantResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RemoveContractParticipantResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAddedByActorId

`func (o *RemoveContractParticipantResponse) GetAddedByActorId() string`

GetAddedByActorId returns the AddedByActorId field if non-nil, zero value otherwise.

### GetAddedByActorIdOk

`func (o *RemoveContractParticipantResponse) GetAddedByActorIdOk() (*string, bool)`

GetAddedByActorIdOk returns a tuple with the AddedByActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedByActorId

`func (o *RemoveContractParticipantResponse) SetAddedByActorId(v string)`

SetAddedByActorId sets AddedByActorId field to given value.


### GetRemovedByActorId

`func (o *RemoveContractParticipantResponse) GetRemovedByActorId() string`

GetRemovedByActorId returns the RemovedByActorId field if non-nil, zero value otherwise.

### GetRemovedByActorIdOk

`func (o *RemoveContractParticipantResponse) GetRemovedByActorIdOk() (*string, bool)`

GetRemovedByActorIdOk returns a tuple with the RemovedByActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedByActorId

`func (o *RemoveContractParticipantResponse) SetRemovedByActorId(v string)`

SetRemovedByActorId sets RemovedByActorId field to given value.

### HasRemovedByActorId

`func (o *RemoveContractParticipantResponse) HasRemovedByActorId() bool`

HasRemovedByActorId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RemoveContractParticipantResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RemoveContractParticipantResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RemoveContractParticipantResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RemoveContractParticipantResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RemoveContractParticipantResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RemoveContractParticipantResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RemoveContractParticipantResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RemoveContractParticipantResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


