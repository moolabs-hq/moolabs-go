# ContractParticipantResponse

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

### NewContractParticipantResponse

`func NewContractParticipantResponse(id string, tenantId string, threadId string, side string, role string, status string, addedByActorId string, ) *ContractParticipantResponse`

NewContractParticipantResponse instantiates a new ContractParticipantResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractParticipantResponseWithDefaults

`func NewContractParticipantResponseWithDefaults() *ContractParticipantResponse`

NewContractParticipantResponseWithDefaults instantiates a new ContractParticipantResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContractParticipantResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContractParticipantResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContractParticipantResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *ContractParticipantResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ContractParticipantResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ContractParticipantResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetThreadId

`func (o *ContractParticipantResponse) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *ContractParticipantResponse) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *ContractParticipantResponse) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.


### GetSide

`func (o *ContractParticipantResponse) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *ContractParticipantResponse) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *ContractParticipantResponse) SetSide(v string)`

SetSide sets Side field to given value.


### GetRole

`func (o *ContractParticipantResponse) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ContractParticipantResponse) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ContractParticipantResponse) SetRole(v string)`

SetRole sets Role field to given value.


### GetEmail

`func (o *ContractParticipantResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ContractParticipantResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ContractParticipantResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ContractParticipantResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDisplayName

`func (o *ContractParticipantResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ContractParticipantResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ContractParticipantResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ContractParticipantResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetStatus

`func (o *ContractParticipantResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ContractParticipantResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ContractParticipantResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAddedByActorId

`func (o *ContractParticipantResponse) GetAddedByActorId() string`

GetAddedByActorId returns the AddedByActorId field if non-nil, zero value otherwise.

### GetAddedByActorIdOk

`func (o *ContractParticipantResponse) GetAddedByActorIdOk() (*string, bool)`

GetAddedByActorIdOk returns a tuple with the AddedByActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedByActorId

`func (o *ContractParticipantResponse) SetAddedByActorId(v string)`

SetAddedByActorId sets AddedByActorId field to given value.


### GetRemovedByActorId

`func (o *ContractParticipantResponse) GetRemovedByActorId() string`

GetRemovedByActorId returns the RemovedByActorId field if non-nil, zero value otherwise.

### GetRemovedByActorIdOk

`func (o *ContractParticipantResponse) GetRemovedByActorIdOk() (*string, bool)`

GetRemovedByActorIdOk returns a tuple with the RemovedByActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedByActorId

`func (o *ContractParticipantResponse) SetRemovedByActorId(v string)`

SetRemovedByActorId sets RemovedByActorId field to given value.

### HasRemovedByActorId

`func (o *ContractParticipantResponse) HasRemovedByActorId() bool`

HasRemovedByActorId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ContractParticipantResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ContractParticipantResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ContractParticipantResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ContractParticipantResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ContractParticipantResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ContractParticipantResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ContractParticipantResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ContractParticipantResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


