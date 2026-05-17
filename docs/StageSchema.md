# StageSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stage** | **int32** | Stage sequence number | 
**Name** | **string** |  | 
**DaysOverdueTrigger** | **int32** |  | 
**Channel** | **string** |  | 
**Tone** | Pointer to **string** |  | [optional] [default to "professional"]
**AutonomyMin** | Pointer to **int32** |  | [optional] [default to 1]
**TemplateKey** | **string** |  | 

## Methods

### NewStageSchema

`func NewStageSchema(stage int32, name string, daysOverdueTrigger int32, channel string, templateKey string, ) *StageSchema`

NewStageSchema instantiates a new StageSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStageSchemaWithDefaults

`func NewStageSchemaWithDefaults() *StageSchema`

NewStageSchemaWithDefaults instantiates a new StageSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStage

`func (o *StageSchema) GetStage() int32`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *StageSchema) GetStageOk() (*int32, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *StageSchema) SetStage(v int32)`

SetStage sets Stage field to given value.


### GetName

`func (o *StageSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StageSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StageSchema) SetName(v string)`

SetName sets Name field to given value.


### GetDaysOverdueTrigger

`func (o *StageSchema) GetDaysOverdueTrigger() int32`

GetDaysOverdueTrigger returns the DaysOverdueTrigger field if non-nil, zero value otherwise.

### GetDaysOverdueTriggerOk

`func (o *StageSchema) GetDaysOverdueTriggerOk() (*int32, bool)`

GetDaysOverdueTriggerOk returns a tuple with the DaysOverdueTrigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysOverdueTrigger

`func (o *StageSchema) SetDaysOverdueTrigger(v int32)`

SetDaysOverdueTrigger sets DaysOverdueTrigger field to given value.


### GetChannel

`func (o *StageSchema) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *StageSchema) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *StageSchema) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetTone

`func (o *StageSchema) GetTone() string`

GetTone returns the Tone field if non-nil, zero value otherwise.

### GetToneOk

`func (o *StageSchema) GetToneOk() (*string, bool)`

GetToneOk returns a tuple with the Tone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTone

`func (o *StageSchema) SetTone(v string)`

SetTone sets Tone field to given value.

### HasTone

`func (o *StageSchema) HasTone() bool`

HasTone returns a boolean if a field has been set.

### GetAutonomyMin

`func (o *StageSchema) GetAutonomyMin() int32`

GetAutonomyMin returns the AutonomyMin field if non-nil, zero value otherwise.

### GetAutonomyMinOk

`func (o *StageSchema) GetAutonomyMinOk() (*int32, bool)`

GetAutonomyMinOk returns a tuple with the AutonomyMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutonomyMin

`func (o *StageSchema) SetAutonomyMin(v int32)`

SetAutonomyMin sets AutonomyMin field to given value.

### HasAutonomyMin

`func (o *StageSchema) HasAutonomyMin() bool`

HasAutonomyMin returns a boolean if a field has been set.

### GetTemplateKey

`func (o *StageSchema) GetTemplateKey() string`

GetTemplateKey returns the TemplateKey field if non-nil, zero value otherwise.

### GetTemplateKeyOk

`func (o *StageSchema) GetTemplateKeyOk() (*string, bool)`

GetTemplateKeyOk returns a tuple with the TemplateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateKey

`func (o *StageSchema) SetTemplateKey(v string)`

SetTemplateKey sets TemplateKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


