# TemplatePreviewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Enabled** | **bool** |  | 
**Subject** | **string** |  | 
**Body** | **string** |  | 
**UsedFallback** | **bool** |  | 
**ExpectedVariables** | **[]string** |  | 
**MissingVariables** | **[]string** |  | 
**SampleContext** | **map[string]interface{}** |  | 

## Methods

### NewTemplatePreviewResponse

`func NewTemplatePreviewResponse(id string, enabled bool, subject string, body string, usedFallback bool, expectedVariables []string, missingVariables []string, sampleContext map[string]interface{}, ) *TemplatePreviewResponse`

NewTemplatePreviewResponse instantiates a new TemplatePreviewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplatePreviewResponseWithDefaults

`func NewTemplatePreviewResponseWithDefaults() *TemplatePreviewResponse`

NewTemplatePreviewResponseWithDefaults instantiates a new TemplatePreviewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TemplatePreviewResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplatePreviewResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplatePreviewResponse) SetId(v string)`

SetId sets Id field to given value.


### GetEnabled

`func (o *TemplatePreviewResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TemplatePreviewResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TemplatePreviewResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSubject

`func (o *TemplatePreviewResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *TemplatePreviewResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *TemplatePreviewResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetBody

`func (o *TemplatePreviewResponse) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *TemplatePreviewResponse) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *TemplatePreviewResponse) SetBody(v string)`

SetBody sets Body field to given value.


### GetUsedFallback

`func (o *TemplatePreviewResponse) GetUsedFallback() bool`

GetUsedFallback returns the UsedFallback field if non-nil, zero value otherwise.

### GetUsedFallbackOk

`func (o *TemplatePreviewResponse) GetUsedFallbackOk() (*bool, bool)`

GetUsedFallbackOk returns a tuple with the UsedFallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedFallback

`func (o *TemplatePreviewResponse) SetUsedFallback(v bool)`

SetUsedFallback sets UsedFallback field to given value.


### GetExpectedVariables

`func (o *TemplatePreviewResponse) GetExpectedVariables() []string`

GetExpectedVariables returns the ExpectedVariables field if non-nil, zero value otherwise.

### GetExpectedVariablesOk

`func (o *TemplatePreviewResponse) GetExpectedVariablesOk() (*[]string, bool)`

GetExpectedVariablesOk returns a tuple with the ExpectedVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedVariables

`func (o *TemplatePreviewResponse) SetExpectedVariables(v []string)`

SetExpectedVariables sets ExpectedVariables field to given value.


### GetMissingVariables

`func (o *TemplatePreviewResponse) GetMissingVariables() []string`

GetMissingVariables returns the MissingVariables field if non-nil, zero value otherwise.

### GetMissingVariablesOk

`func (o *TemplatePreviewResponse) GetMissingVariablesOk() (*[]string, bool)`

GetMissingVariablesOk returns a tuple with the MissingVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingVariables

`func (o *TemplatePreviewResponse) SetMissingVariables(v []string)`

SetMissingVariables sets MissingVariables field to given value.


### GetSampleContext

`func (o *TemplatePreviewResponse) GetSampleContext() map[string]interface{}`

GetSampleContext returns the SampleContext field if non-nil, zero value otherwise.

### GetSampleContextOk

`func (o *TemplatePreviewResponse) GetSampleContextOk() (*map[string]interface{}, bool)`

GetSampleContextOk returns a tuple with the SampleContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleContext

`func (o *TemplatePreviewResponse) SetSampleContext(v map[string]interface{})`

SetSampleContext sets SampleContext field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


