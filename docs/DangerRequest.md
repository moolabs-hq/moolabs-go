# DangerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuditNote** | **string** |  | 
**Confirm** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewDangerRequest

`func NewDangerRequest(auditNote string, ) *DangerRequest`

NewDangerRequest instantiates a new DangerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDangerRequestWithDefaults

`func NewDangerRequestWithDefaults() *DangerRequest`

NewDangerRequestWithDefaults instantiates a new DangerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuditNote

`func (o *DangerRequest) GetAuditNote() string`

GetAuditNote returns the AuditNote field if non-nil, zero value otherwise.

### GetAuditNoteOk

`func (o *DangerRequest) GetAuditNoteOk() (*string, bool)`

GetAuditNoteOk returns a tuple with the AuditNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditNote

`func (o *DangerRequest) SetAuditNote(v string)`

SetAuditNote sets AuditNote field to given value.


### GetConfirm

`func (o *DangerRequest) GetConfirm() bool`

GetConfirm returns the Confirm field if non-nil, zero value otherwise.

### GetConfirmOk

`func (o *DangerRequest) GetConfirmOk() (*bool, bool)`

GetConfirmOk returns a tuple with the Confirm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirm

`func (o *DangerRequest) SetConfirm(v bool)`

SetConfirm sets Confirm field to given value.

### HasConfirm

`func (o *DangerRequest) HasConfirm() bool`

HasConfirm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


