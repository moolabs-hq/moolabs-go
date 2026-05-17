# MappingRuleCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleName** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] [default to 100]
**SourceType** | **string** |  | 
**SourceField** | **string** |  | 
**TransformType** | Pointer to **string** |  | [optional] [default to "direct"]
**TransformPattern** | Pointer to **string** |  | [optional] 
**TransformMap** | Pointer to **map[string]interface{}** |  | [optional] 
**TargetField** | **string** |  | 
**PiiCheckEnabled** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewMappingRuleCreate

`func NewMappingRuleCreate(ruleName string, sourceType string, sourceField string, targetField string, ) *MappingRuleCreate`

NewMappingRuleCreate instantiates a new MappingRuleCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMappingRuleCreateWithDefaults

`func NewMappingRuleCreateWithDefaults() *MappingRuleCreate`

NewMappingRuleCreateWithDefaults instantiates a new MappingRuleCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleName

`func (o *MappingRuleCreate) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *MappingRuleCreate) GetRuleNameOk() (*string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleName

`func (o *MappingRuleCreate) SetRuleName(v string)`

SetRuleName sets RuleName field to given value.


### GetDescription

`func (o *MappingRuleCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MappingRuleCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MappingRuleCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MappingRuleCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPriority

`func (o *MappingRuleCreate) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *MappingRuleCreate) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *MappingRuleCreate) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *MappingRuleCreate) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetSourceType

`func (o *MappingRuleCreate) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *MappingRuleCreate) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *MappingRuleCreate) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetSourceField

`func (o *MappingRuleCreate) GetSourceField() string`

GetSourceField returns the SourceField field if non-nil, zero value otherwise.

### GetSourceFieldOk

`func (o *MappingRuleCreate) GetSourceFieldOk() (*string, bool)`

GetSourceFieldOk returns a tuple with the SourceField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceField

`func (o *MappingRuleCreate) SetSourceField(v string)`

SetSourceField sets SourceField field to given value.


### GetTransformType

`func (o *MappingRuleCreate) GetTransformType() string`

GetTransformType returns the TransformType field if non-nil, zero value otherwise.

### GetTransformTypeOk

`func (o *MappingRuleCreate) GetTransformTypeOk() (*string, bool)`

GetTransformTypeOk returns a tuple with the TransformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformType

`func (o *MappingRuleCreate) SetTransformType(v string)`

SetTransformType sets TransformType field to given value.

### HasTransformType

`func (o *MappingRuleCreate) HasTransformType() bool`

HasTransformType returns a boolean if a field has been set.

### GetTransformPattern

`func (o *MappingRuleCreate) GetTransformPattern() string`

GetTransformPattern returns the TransformPattern field if non-nil, zero value otherwise.

### GetTransformPatternOk

`func (o *MappingRuleCreate) GetTransformPatternOk() (*string, bool)`

GetTransformPatternOk returns a tuple with the TransformPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformPattern

`func (o *MappingRuleCreate) SetTransformPattern(v string)`

SetTransformPattern sets TransformPattern field to given value.

### HasTransformPattern

`func (o *MappingRuleCreate) HasTransformPattern() bool`

HasTransformPattern returns a boolean if a field has been set.

### GetTransformMap

`func (o *MappingRuleCreate) GetTransformMap() map[string]interface{}`

GetTransformMap returns the TransformMap field if non-nil, zero value otherwise.

### GetTransformMapOk

`func (o *MappingRuleCreate) GetTransformMapOk() (*map[string]interface{}, bool)`

GetTransformMapOk returns a tuple with the TransformMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformMap

`func (o *MappingRuleCreate) SetTransformMap(v map[string]interface{})`

SetTransformMap sets TransformMap field to given value.

### HasTransformMap

`func (o *MappingRuleCreate) HasTransformMap() bool`

HasTransformMap returns a boolean if a field has been set.

### GetTargetField

`func (o *MappingRuleCreate) GetTargetField() string`

GetTargetField returns the TargetField field if non-nil, zero value otherwise.

### GetTargetFieldOk

`func (o *MappingRuleCreate) GetTargetFieldOk() (*string, bool)`

GetTargetFieldOk returns a tuple with the TargetField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetField

`func (o *MappingRuleCreate) SetTargetField(v string)`

SetTargetField sets TargetField field to given value.


### GetPiiCheckEnabled

`func (o *MappingRuleCreate) GetPiiCheckEnabled() bool`

GetPiiCheckEnabled returns the PiiCheckEnabled field if non-nil, zero value otherwise.

### GetPiiCheckEnabledOk

`func (o *MappingRuleCreate) GetPiiCheckEnabledOk() (*bool, bool)`

GetPiiCheckEnabledOk returns a tuple with the PiiCheckEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPiiCheckEnabled

`func (o *MappingRuleCreate) SetPiiCheckEnabled(v bool)`

SetPiiCheckEnabled sets PiiCheckEnabled field to given value.

### HasPiiCheckEnabled

`func (o *MappingRuleCreate) HasPiiCheckEnabled() bool`

HasPiiCheckEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


