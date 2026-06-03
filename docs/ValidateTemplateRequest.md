# ValidateTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** |  | [optional] [default to "email"]
**SubjectTemplate** | **string** |  | 
**BodyTemplate** | **string** |  | 
**PreviewMode** | Pointer to **string** |  | [optional] 
**CaseId** | Pointer to **string** |  | [optional] 

## Methods

### NewValidateTemplateRequest

`func NewValidateTemplateRequest(subjectTemplate string, bodyTemplate string, ) *ValidateTemplateRequest`

NewValidateTemplateRequest instantiates a new ValidateTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateTemplateRequestWithDefaults

`func NewValidateTemplateRequestWithDefaults() *ValidateTemplateRequest`

NewValidateTemplateRequestWithDefaults instantiates a new ValidateTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *ValidateTemplateRequest) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ValidateTemplateRequest) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ValidateTemplateRequest) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ValidateTemplateRequest) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetSubjectTemplate

`func (o *ValidateTemplateRequest) GetSubjectTemplate() string`

GetSubjectTemplate returns the SubjectTemplate field if non-nil, zero value otherwise.

### GetSubjectTemplateOk

`func (o *ValidateTemplateRequest) GetSubjectTemplateOk() (*string, bool)`

GetSubjectTemplateOk returns a tuple with the SubjectTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectTemplate

`func (o *ValidateTemplateRequest) SetSubjectTemplate(v string)`

SetSubjectTemplate sets SubjectTemplate field to given value.


### GetBodyTemplate

`func (o *ValidateTemplateRequest) GetBodyTemplate() string`

GetBodyTemplate returns the BodyTemplate field if non-nil, zero value otherwise.

### GetBodyTemplateOk

`func (o *ValidateTemplateRequest) GetBodyTemplateOk() (*string, bool)`

GetBodyTemplateOk returns a tuple with the BodyTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyTemplate

`func (o *ValidateTemplateRequest) SetBodyTemplate(v string)`

SetBodyTemplate sets BodyTemplate field to given value.


### GetPreviewMode

`func (o *ValidateTemplateRequest) GetPreviewMode() string`

GetPreviewMode returns the PreviewMode field if non-nil, zero value otherwise.

### GetPreviewModeOk

`func (o *ValidateTemplateRequest) GetPreviewModeOk() (*string, bool)`

GetPreviewModeOk returns a tuple with the PreviewMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewMode

`func (o *ValidateTemplateRequest) SetPreviewMode(v string)`

SetPreviewMode sets PreviewMode field to given value.

### HasPreviewMode

`func (o *ValidateTemplateRequest) HasPreviewMode() bool`

HasPreviewMode returns a boolean if a field has been set.

### GetCaseId

`func (o *ValidateTemplateRequest) GetCaseId() string`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *ValidateTemplateRequest) GetCaseIdOk() (*string, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *ValidateTemplateRequest) SetCaseId(v string)`

SetCaseId sets CaseId field to given value.

### HasCaseId

`func (o *ValidateTemplateRequest) HasCaseId() bool`

HasCaseId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


