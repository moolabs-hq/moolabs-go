# PreviewTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VersionId** | **string** |  | 
**PreviewMode** | Pointer to **string** |  | [optional] 
**CaseId** | Pointer to **string** |  | [optional] 

## Methods

### NewPreviewTemplateRequest

`func NewPreviewTemplateRequest(versionId string, ) *PreviewTemplateRequest`

NewPreviewTemplateRequest instantiates a new PreviewTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreviewTemplateRequestWithDefaults

`func NewPreviewTemplateRequestWithDefaults() *PreviewTemplateRequest`

NewPreviewTemplateRequestWithDefaults instantiates a new PreviewTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersionId

`func (o *PreviewTemplateRequest) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *PreviewTemplateRequest) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *PreviewTemplateRequest) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetPreviewMode

`func (o *PreviewTemplateRequest) GetPreviewMode() string`

GetPreviewMode returns the PreviewMode field if non-nil, zero value otherwise.

### GetPreviewModeOk

`func (o *PreviewTemplateRequest) GetPreviewModeOk() (*string, bool)`

GetPreviewModeOk returns a tuple with the PreviewMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewMode

`func (o *PreviewTemplateRequest) SetPreviewMode(v string)`

SetPreviewMode sets PreviewMode field to given value.

### HasPreviewMode

`func (o *PreviewTemplateRequest) HasPreviewMode() bool`

HasPreviewMode returns a boolean if a field has been set.

### GetCaseId

`func (o *PreviewTemplateRequest) GetCaseId() string`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *PreviewTemplateRequest) GetCaseIdOk() (*string, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *PreviewTemplateRequest) SetCaseId(v string)`

SetCaseId sets CaseId field to given value.

### HasCaseId

`func (o *PreviewTemplateRequest) HasCaseId() bool`

HasCaseId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


