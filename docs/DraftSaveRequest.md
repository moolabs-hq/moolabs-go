# DraftSaveRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** |  | [optional] [default to "email"]
**SubjectTemplate** | **string** |  | 
**BodyTemplate** | **string** |  | 
**ChangeReason** | **string** |  | 
**PromptSummary** | Pointer to **string** |  | [optional] 
**IdempotencyKey** | Pointer to **string** |  | [optional] 
**BaseVersionId** | Pointer to **string** |  | [optional] 
**LockId** | Pointer to **string** |  | [optional] 
**LockToken** | Pointer to **string** |  | [optional] 

## Methods

### NewDraftSaveRequest

`func NewDraftSaveRequest(subjectTemplate string, bodyTemplate string, changeReason string, ) *DraftSaveRequest`

NewDraftSaveRequest instantiates a new DraftSaveRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDraftSaveRequestWithDefaults

`func NewDraftSaveRequestWithDefaults() *DraftSaveRequest`

NewDraftSaveRequestWithDefaults instantiates a new DraftSaveRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *DraftSaveRequest) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *DraftSaveRequest) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *DraftSaveRequest) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *DraftSaveRequest) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetSubjectTemplate

`func (o *DraftSaveRequest) GetSubjectTemplate() string`

GetSubjectTemplate returns the SubjectTemplate field if non-nil, zero value otherwise.

### GetSubjectTemplateOk

`func (o *DraftSaveRequest) GetSubjectTemplateOk() (*string, bool)`

GetSubjectTemplateOk returns a tuple with the SubjectTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectTemplate

`func (o *DraftSaveRequest) SetSubjectTemplate(v string)`

SetSubjectTemplate sets SubjectTemplate field to given value.


### GetBodyTemplate

`func (o *DraftSaveRequest) GetBodyTemplate() string`

GetBodyTemplate returns the BodyTemplate field if non-nil, zero value otherwise.

### GetBodyTemplateOk

`func (o *DraftSaveRequest) GetBodyTemplateOk() (*string, bool)`

GetBodyTemplateOk returns a tuple with the BodyTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyTemplate

`func (o *DraftSaveRequest) SetBodyTemplate(v string)`

SetBodyTemplate sets BodyTemplate field to given value.


### GetChangeReason

`func (o *DraftSaveRequest) GetChangeReason() string`

GetChangeReason returns the ChangeReason field if non-nil, zero value otherwise.

### GetChangeReasonOk

`func (o *DraftSaveRequest) GetChangeReasonOk() (*string, bool)`

GetChangeReasonOk returns a tuple with the ChangeReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeReason

`func (o *DraftSaveRequest) SetChangeReason(v string)`

SetChangeReason sets ChangeReason field to given value.


### GetPromptSummary

`func (o *DraftSaveRequest) GetPromptSummary() string`

GetPromptSummary returns the PromptSummary field if non-nil, zero value otherwise.

### GetPromptSummaryOk

`func (o *DraftSaveRequest) GetPromptSummaryOk() (*string, bool)`

GetPromptSummaryOk returns a tuple with the PromptSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptSummary

`func (o *DraftSaveRequest) SetPromptSummary(v string)`

SetPromptSummary sets PromptSummary field to given value.

### HasPromptSummary

`func (o *DraftSaveRequest) HasPromptSummary() bool`

HasPromptSummary returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *DraftSaveRequest) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *DraftSaveRequest) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *DraftSaveRequest) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *DraftSaveRequest) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetBaseVersionId

`func (o *DraftSaveRequest) GetBaseVersionId() string`

GetBaseVersionId returns the BaseVersionId field if non-nil, zero value otherwise.

### GetBaseVersionIdOk

`func (o *DraftSaveRequest) GetBaseVersionIdOk() (*string, bool)`

GetBaseVersionIdOk returns a tuple with the BaseVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseVersionId

`func (o *DraftSaveRequest) SetBaseVersionId(v string)`

SetBaseVersionId sets BaseVersionId field to given value.

### HasBaseVersionId

`func (o *DraftSaveRequest) HasBaseVersionId() bool`

HasBaseVersionId returns a boolean if a field has been set.

### GetLockId

`func (o *DraftSaveRequest) GetLockId() string`

GetLockId returns the LockId field if non-nil, zero value otherwise.

### GetLockIdOk

`func (o *DraftSaveRequest) GetLockIdOk() (*string, bool)`

GetLockIdOk returns a tuple with the LockId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockId

`func (o *DraftSaveRequest) SetLockId(v string)`

SetLockId sets LockId field to given value.

### HasLockId

`func (o *DraftSaveRequest) HasLockId() bool`

HasLockId returns a boolean if a field has been set.

### GetLockToken

`func (o *DraftSaveRequest) GetLockToken() string`

GetLockToken returns the LockToken field if non-nil, zero value otherwise.

### GetLockTokenOk

`func (o *DraftSaveRequest) GetLockTokenOk() (*string, bool)`

GetLockTokenOk returns a tuple with the LockToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockToken

`func (o *DraftSaveRequest) SetLockToken(v string)`

SetLockToken sets LockToken field to given value.

### HasLockToken

`func (o *DraftSaveRequest) HasLockToken() bool`

HasLockToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


