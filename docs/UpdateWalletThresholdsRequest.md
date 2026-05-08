# UpdateWalletThresholdsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalSoftThresholdMicros** | Pointer to **NullableInt32** |  | [optional] 
**LocalHardCapMicros** | Pointer to **NullableInt32** |  | [optional] 
**AuditNote** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdateWalletThresholdsRequest

`func NewUpdateWalletThresholdsRequest() *UpdateWalletThresholdsRequest`

NewUpdateWalletThresholdsRequest instantiates a new UpdateWalletThresholdsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWalletThresholdsRequestWithDefaults

`func NewUpdateWalletThresholdsRequestWithDefaults() *UpdateWalletThresholdsRequest`

NewUpdateWalletThresholdsRequestWithDefaults instantiates a new UpdateWalletThresholdsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalSoftThresholdMicros

`func (o *UpdateWalletThresholdsRequest) GetLocalSoftThresholdMicros() int32`

GetLocalSoftThresholdMicros returns the LocalSoftThresholdMicros field if non-nil, zero value otherwise.

### GetLocalSoftThresholdMicrosOk

`func (o *UpdateWalletThresholdsRequest) GetLocalSoftThresholdMicrosOk() (*int32, bool)`

GetLocalSoftThresholdMicrosOk returns a tuple with the LocalSoftThresholdMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSoftThresholdMicros

`func (o *UpdateWalletThresholdsRequest) SetLocalSoftThresholdMicros(v int32)`

SetLocalSoftThresholdMicros sets LocalSoftThresholdMicros field to given value.

### HasLocalSoftThresholdMicros

`func (o *UpdateWalletThresholdsRequest) HasLocalSoftThresholdMicros() bool`

HasLocalSoftThresholdMicros returns a boolean if a field has been set.

### SetLocalSoftThresholdMicrosNil

`func (o *UpdateWalletThresholdsRequest) SetLocalSoftThresholdMicrosNil(b bool)`

 SetLocalSoftThresholdMicrosNil sets the value for LocalSoftThresholdMicros to be an explicit nil

### UnsetLocalSoftThresholdMicros
`func (o *UpdateWalletThresholdsRequest) UnsetLocalSoftThresholdMicros()`

UnsetLocalSoftThresholdMicros ensures that no value is present for LocalSoftThresholdMicros, not even an explicit nil
### GetLocalHardCapMicros

`func (o *UpdateWalletThresholdsRequest) GetLocalHardCapMicros() int32`

GetLocalHardCapMicros returns the LocalHardCapMicros field if non-nil, zero value otherwise.

### GetLocalHardCapMicrosOk

`func (o *UpdateWalletThresholdsRequest) GetLocalHardCapMicrosOk() (*int32, bool)`

GetLocalHardCapMicrosOk returns a tuple with the LocalHardCapMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalHardCapMicros

`func (o *UpdateWalletThresholdsRequest) SetLocalHardCapMicros(v int32)`

SetLocalHardCapMicros sets LocalHardCapMicros field to given value.

### HasLocalHardCapMicros

`func (o *UpdateWalletThresholdsRequest) HasLocalHardCapMicros() bool`

HasLocalHardCapMicros returns a boolean if a field has been set.

### SetLocalHardCapMicrosNil

`func (o *UpdateWalletThresholdsRequest) SetLocalHardCapMicrosNil(b bool)`

 SetLocalHardCapMicrosNil sets the value for LocalHardCapMicros to be an explicit nil

### UnsetLocalHardCapMicros
`func (o *UpdateWalletThresholdsRequest) UnsetLocalHardCapMicros()`

UnsetLocalHardCapMicros ensures that no value is present for LocalHardCapMicros, not even an explicit nil
### GetAuditNote

`func (o *UpdateWalletThresholdsRequest) GetAuditNote() string`

GetAuditNote returns the AuditNote field if non-nil, zero value otherwise.

### GetAuditNoteOk

`func (o *UpdateWalletThresholdsRequest) GetAuditNoteOk() (*string, bool)`

GetAuditNoteOk returns a tuple with the AuditNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditNote

`func (o *UpdateWalletThresholdsRequest) SetAuditNote(v string)`

SetAuditNote sets AuditNote field to given value.

### HasAuditNote

`func (o *UpdateWalletThresholdsRequest) HasAuditNote() bool`

HasAuditNote returns a boolean if a field has been set.

### SetAuditNoteNil

`func (o *UpdateWalletThresholdsRequest) SetAuditNoteNil(b bool)`

 SetAuditNoteNil sets the value for AuditNote to be an explicit nil

### UnsetAuditNote
`func (o *UpdateWalletThresholdsRequest) UnsetAuditNote()`

UnsetAuditNote ensures that no value is present for AuditNote, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


