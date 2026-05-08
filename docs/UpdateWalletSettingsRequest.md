# UpdateWalletSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentWalletId** | Pointer to **NullableString** |  | [optional] 
**LocalHardCapMicros** | Pointer to **NullableInt32** |  | [optional] 
**Policy** | Pointer to [**NullableWalletPolicy**](WalletPolicy.md) |  | [optional] 
**AuditNote** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdateWalletSettingsRequest

`func NewUpdateWalletSettingsRequest() *UpdateWalletSettingsRequest`

NewUpdateWalletSettingsRequest instantiates a new UpdateWalletSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWalletSettingsRequestWithDefaults

`func NewUpdateWalletSettingsRequestWithDefaults() *UpdateWalletSettingsRequest`

NewUpdateWalletSettingsRequestWithDefaults instantiates a new UpdateWalletSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentWalletId

`func (o *UpdateWalletSettingsRequest) GetParentWalletId() string`

GetParentWalletId returns the ParentWalletId field if non-nil, zero value otherwise.

### GetParentWalletIdOk

`func (o *UpdateWalletSettingsRequest) GetParentWalletIdOk() (*string, bool)`

GetParentWalletIdOk returns a tuple with the ParentWalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentWalletId

`func (o *UpdateWalletSettingsRequest) SetParentWalletId(v string)`

SetParentWalletId sets ParentWalletId field to given value.

### HasParentWalletId

`func (o *UpdateWalletSettingsRequest) HasParentWalletId() bool`

HasParentWalletId returns a boolean if a field has been set.

### SetParentWalletIdNil

`func (o *UpdateWalletSettingsRequest) SetParentWalletIdNil(b bool)`

 SetParentWalletIdNil sets the value for ParentWalletId to be an explicit nil

### UnsetParentWalletId
`func (o *UpdateWalletSettingsRequest) UnsetParentWalletId()`

UnsetParentWalletId ensures that no value is present for ParentWalletId, not even an explicit nil
### GetLocalHardCapMicros

`func (o *UpdateWalletSettingsRequest) GetLocalHardCapMicros() int32`

GetLocalHardCapMicros returns the LocalHardCapMicros field if non-nil, zero value otherwise.

### GetLocalHardCapMicrosOk

`func (o *UpdateWalletSettingsRequest) GetLocalHardCapMicrosOk() (*int32, bool)`

GetLocalHardCapMicrosOk returns a tuple with the LocalHardCapMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalHardCapMicros

`func (o *UpdateWalletSettingsRequest) SetLocalHardCapMicros(v int32)`

SetLocalHardCapMicros sets LocalHardCapMicros field to given value.

### HasLocalHardCapMicros

`func (o *UpdateWalletSettingsRequest) HasLocalHardCapMicros() bool`

HasLocalHardCapMicros returns a boolean if a field has been set.

### SetLocalHardCapMicrosNil

`func (o *UpdateWalletSettingsRequest) SetLocalHardCapMicrosNil(b bool)`

 SetLocalHardCapMicrosNil sets the value for LocalHardCapMicros to be an explicit nil

### UnsetLocalHardCapMicros
`func (o *UpdateWalletSettingsRequest) UnsetLocalHardCapMicros()`

UnsetLocalHardCapMicros ensures that no value is present for LocalHardCapMicros, not even an explicit nil
### GetPolicy

`func (o *UpdateWalletSettingsRequest) GetPolicy() WalletPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *UpdateWalletSettingsRequest) GetPolicyOk() (*WalletPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *UpdateWalletSettingsRequest) SetPolicy(v WalletPolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *UpdateWalletSettingsRequest) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### SetPolicyNil

`func (o *UpdateWalletSettingsRequest) SetPolicyNil(b bool)`

 SetPolicyNil sets the value for Policy to be an explicit nil

### UnsetPolicy
`func (o *UpdateWalletSettingsRequest) UnsetPolicy()`

UnsetPolicy ensures that no value is present for Policy, not even an explicit nil
### GetAuditNote

`func (o *UpdateWalletSettingsRequest) GetAuditNote() string`

GetAuditNote returns the AuditNote field if non-nil, zero value otherwise.

### GetAuditNoteOk

`func (o *UpdateWalletSettingsRequest) GetAuditNoteOk() (*string, bool)`

GetAuditNoteOk returns a tuple with the AuditNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditNote

`func (o *UpdateWalletSettingsRequest) SetAuditNote(v string)`

SetAuditNote sets AuditNote field to given value.

### HasAuditNote

`func (o *UpdateWalletSettingsRequest) HasAuditNote() bool`

HasAuditNote returns a boolean if a field has been set.

### SetAuditNoteNil

`func (o *UpdateWalletSettingsRequest) SetAuditNoteNil(b bool)`

 SetAuditNoteNil sets the value for AuditNote to be an explicit nil

### UnsetAuditNote
`func (o *UpdateWalletSettingsRequest) UnsetAuditNote()`

UnsetAuditNote ensures that no value is present for AuditNote, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


