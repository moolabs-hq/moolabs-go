# MappingRuleUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**SourceType** | Pointer to **string** |  | [optional] 
**SourceField** | Pointer to **string** |  | [optional] 
**TransformType** | Pointer to **string** |  | [optional] 
**TransformPattern** | Pointer to **string** |  | [optional] 
**TransformMap** | Pointer to **map[string]interface{}** |  | [optional] 
**TargetField** | Pointer to **string** |  | [optional] 
**PiiCheckEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewMappingRuleUpdate

`func NewMappingRuleUpdate() *MappingRuleUpdate`

NewMappingRuleUpdate instantiates a new MappingRuleUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMappingRuleUpdateWithDefaults

`func NewMappingRuleUpdateWithDefaults() *MappingRuleUpdate`

NewMappingRuleUpdateWithDefaults instantiates a new MappingRuleUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *MappingRuleUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MappingRuleUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MappingRuleUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MappingRuleUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPriority

`func (o *MappingRuleUpdate) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *MappingRuleUpdate) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *MappingRuleUpdate) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *MappingRuleUpdate) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetSourceType

`func (o *MappingRuleUpdate) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *MappingRuleUpdate) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *MappingRuleUpdate) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *MappingRuleUpdate) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetSourceField

`func (o *MappingRuleUpdate) GetSourceField() string`

GetSourceField returns the SourceField field if non-nil, zero value otherwise.

### GetSourceFieldOk

`func (o *MappingRuleUpdate) GetSourceFieldOk() (*string, bool)`

GetSourceFieldOk returns a tuple with the SourceField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceField

`func (o *MappingRuleUpdate) SetSourceField(v string)`

SetSourceField sets SourceField field to given value.

### HasSourceField

`func (o *MappingRuleUpdate) HasSourceField() bool`

HasSourceField returns a boolean if a field has been set.

### GetTransformType

`func (o *MappingRuleUpdate) GetTransformType() string`

GetTransformType returns the TransformType field if non-nil, zero value otherwise.

### GetTransformTypeOk

`func (o *MappingRuleUpdate) GetTransformTypeOk() (*string, bool)`

GetTransformTypeOk returns a tuple with the TransformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformType

`func (o *MappingRuleUpdate) SetTransformType(v string)`

SetTransformType sets TransformType field to given value.

### HasTransformType

`func (o *MappingRuleUpdate) HasTransformType() bool`

HasTransformType returns a boolean if a field has been set.

### GetTransformPattern

`func (o *MappingRuleUpdate) GetTransformPattern() string`

GetTransformPattern returns the TransformPattern field if non-nil, zero value otherwise.

### GetTransformPatternOk

`func (o *MappingRuleUpdate) GetTransformPatternOk() (*string, bool)`

GetTransformPatternOk returns a tuple with the TransformPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformPattern

`func (o *MappingRuleUpdate) SetTransformPattern(v string)`

SetTransformPattern sets TransformPattern field to given value.

### HasTransformPattern

`func (o *MappingRuleUpdate) HasTransformPattern() bool`

HasTransformPattern returns a boolean if a field has been set.

### GetTransformMap

`func (o *MappingRuleUpdate) GetTransformMap() map[string]interface{}`

GetTransformMap returns the TransformMap field if non-nil, zero value otherwise.

### GetTransformMapOk

`func (o *MappingRuleUpdate) GetTransformMapOk() (*map[string]interface{}, bool)`

GetTransformMapOk returns a tuple with the TransformMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformMap

`func (o *MappingRuleUpdate) SetTransformMap(v map[string]interface{})`

SetTransformMap sets TransformMap field to given value.

### HasTransformMap

`func (o *MappingRuleUpdate) HasTransformMap() bool`

HasTransformMap returns a boolean if a field has been set.

### GetTargetField

`func (o *MappingRuleUpdate) GetTargetField() string`

GetTargetField returns the TargetField field if non-nil, zero value otherwise.

### GetTargetFieldOk

`func (o *MappingRuleUpdate) GetTargetFieldOk() (*string, bool)`

GetTargetFieldOk returns a tuple with the TargetField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetField

`func (o *MappingRuleUpdate) SetTargetField(v string)`

SetTargetField sets TargetField field to given value.

### HasTargetField

`func (o *MappingRuleUpdate) HasTargetField() bool`

HasTargetField returns a boolean if a field has been set.

### GetPiiCheckEnabled

`func (o *MappingRuleUpdate) GetPiiCheckEnabled() bool`

GetPiiCheckEnabled returns the PiiCheckEnabled field if non-nil, zero value otherwise.

### GetPiiCheckEnabledOk

`func (o *MappingRuleUpdate) GetPiiCheckEnabledOk() (*bool, bool)`

GetPiiCheckEnabledOk returns a tuple with the PiiCheckEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPiiCheckEnabled

`func (o *MappingRuleUpdate) SetPiiCheckEnabled(v bool)`

SetPiiCheckEnabled sets PiiCheckEnabled field to given value.

### HasPiiCheckEnabled

`func (o *MappingRuleUpdate) HasPiiCheckEnabled() bool`

HasPiiCheckEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


