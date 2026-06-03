# DraftLockResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VersionId** | **string** |  | 
**LockId** | **string** |  | 
**LockToken** | **string** |  | 
**LockOwnerActorId** | **string** |  | 
**LockOwnerDisplay** | Pointer to **string** |  | [optional] 
**LockExpiresAt** | **string** |  | 

## Methods

### NewDraftLockResponse

`func NewDraftLockResponse(versionId string, lockId string, lockToken string, lockOwnerActorId string, lockExpiresAt string, ) *DraftLockResponse`

NewDraftLockResponse instantiates a new DraftLockResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDraftLockResponseWithDefaults

`func NewDraftLockResponseWithDefaults() *DraftLockResponse`

NewDraftLockResponseWithDefaults instantiates a new DraftLockResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersionId

`func (o *DraftLockResponse) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *DraftLockResponse) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *DraftLockResponse) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetLockId

`func (o *DraftLockResponse) GetLockId() string`

GetLockId returns the LockId field if non-nil, zero value otherwise.

### GetLockIdOk

`func (o *DraftLockResponse) GetLockIdOk() (*string, bool)`

GetLockIdOk returns a tuple with the LockId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockId

`func (o *DraftLockResponse) SetLockId(v string)`

SetLockId sets LockId field to given value.


### GetLockToken

`func (o *DraftLockResponse) GetLockToken() string`

GetLockToken returns the LockToken field if non-nil, zero value otherwise.

### GetLockTokenOk

`func (o *DraftLockResponse) GetLockTokenOk() (*string, bool)`

GetLockTokenOk returns a tuple with the LockToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockToken

`func (o *DraftLockResponse) SetLockToken(v string)`

SetLockToken sets LockToken field to given value.


### GetLockOwnerActorId

`func (o *DraftLockResponse) GetLockOwnerActorId() string`

GetLockOwnerActorId returns the LockOwnerActorId field if non-nil, zero value otherwise.

### GetLockOwnerActorIdOk

`func (o *DraftLockResponse) GetLockOwnerActorIdOk() (*string, bool)`

GetLockOwnerActorIdOk returns a tuple with the LockOwnerActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockOwnerActorId

`func (o *DraftLockResponse) SetLockOwnerActorId(v string)`

SetLockOwnerActorId sets LockOwnerActorId field to given value.


### GetLockOwnerDisplay

`func (o *DraftLockResponse) GetLockOwnerDisplay() string`

GetLockOwnerDisplay returns the LockOwnerDisplay field if non-nil, zero value otherwise.

### GetLockOwnerDisplayOk

`func (o *DraftLockResponse) GetLockOwnerDisplayOk() (*string, bool)`

GetLockOwnerDisplayOk returns a tuple with the LockOwnerDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockOwnerDisplay

`func (o *DraftLockResponse) SetLockOwnerDisplay(v string)`

SetLockOwnerDisplay sets LockOwnerDisplay field to given value.

### HasLockOwnerDisplay

`func (o *DraftLockResponse) HasLockOwnerDisplay() bool`

HasLockOwnerDisplay returns a boolean if a field has been set.

### GetLockExpiresAt

`func (o *DraftLockResponse) GetLockExpiresAt() string`

GetLockExpiresAt returns the LockExpiresAt field if non-nil, zero value otherwise.

### GetLockExpiresAtOk

`func (o *DraftLockResponse) GetLockExpiresAtOk() (*string, bool)`

GetLockExpiresAtOk returns a tuple with the LockExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockExpiresAt

`func (o *DraftLockResponse) SetLockExpiresAt(v string)`

SetLockExpiresAt sets LockExpiresAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


