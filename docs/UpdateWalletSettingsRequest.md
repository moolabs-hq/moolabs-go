# UpdateWalletSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentWalletId** | Pointer to **string** | Parent wallet ID. Set to null to detach from hierarchy. | [optional] 
**LocalHardCapMicros** | Pointer to **int32** | Local hard cap (micros). Set to null to remove override. | [optional] 
**Policy** | Pointer to [**WalletPolicy**](WalletPolicy.md) | Wallet policy (SOFT_BORROW, HARD_BUDGET, NOTIFY_ONLY) | [optional] 
**AuditNote** | Pointer to **string** | Audit note for this change | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


