# MappingRuleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**RuleName** | **string** |  | 
**Description** | **string** |  | 
**IsEnabled** | **bool** |  | 
**Priority** | **int32** |  | 
**SourceType** | **string** |  | 
**SourceField** | **string** |  | 
**TransformType** | **string** |  | 
**TransformPattern** | **string** |  | 
**TransformMap** | **map[string]interface{}** |  | 
**TargetField** | **string** |  | 
**PiiCheckEnabled** | **bool** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**CreatedBy** | **string** |  | 

## Methods

### NewMappingRuleResponse

`func NewMappingRuleResponse(id string, tenantId string, ruleName string, description string, isEnabled bool, priority int32, sourceType string, sourceField string, transformType string, transformPattern string, transformMap map[string]interface{}, targetField string, piiCheckEnabled bool, createdAt time.Time, updatedAt time.Time, createdBy string, ) *MappingRuleResponse`

NewMappingRuleResponse instantiates a new MappingRuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMappingRuleResponseWithDefaults

`func NewMappingRuleResponseWithDefaults() *MappingRuleResponse`

NewMappingRuleResponseWithDefaults instantiates a new MappingRuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MappingRuleResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MappingRuleResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MappingRuleResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *MappingRuleResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *MappingRuleResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *MappingRuleResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetRuleName

`func (o *MappingRuleResponse) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *MappingRuleResponse) GetRuleNameOk() (*string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleName

`func (o *MappingRuleResponse) SetRuleName(v string)`

SetRuleName sets RuleName field to given value.


### GetDescription

`func (o *MappingRuleResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MappingRuleResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MappingRuleResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIsEnabled

`func (o *MappingRuleResponse) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *MappingRuleResponse) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *MappingRuleResponse) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetPriority

`func (o *MappingRuleResponse) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *MappingRuleResponse) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *MappingRuleResponse) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetSourceType

`func (o *MappingRuleResponse) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *MappingRuleResponse) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *MappingRuleResponse) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetSourceField

`func (o *MappingRuleResponse) GetSourceField() string`

GetSourceField returns the SourceField field if non-nil, zero value otherwise.

### GetSourceFieldOk

`func (o *MappingRuleResponse) GetSourceFieldOk() (*string, bool)`

GetSourceFieldOk returns a tuple with the SourceField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceField

`func (o *MappingRuleResponse) SetSourceField(v string)`

SetSourceField sets SourceField field to given value.


### GetTransformType

`func (o *MappingRuleResponse) GetTransformType() string`

GetTransformType returns the TransformType field if non-nil, zero value otherwise.

### GetTransformTypeOk

`func (o *MappingRuleResponse) GetTransformTypeOk() (*string, bool)`

GetTransformTypeOk returns a tuple with the TransformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformType

`func (o *MappingRuleResponse) SetTransformType(v string)`

SetTransformType sets TransformType field to given value.


### GetTransformPattern

`func (o *MappingRuleResponse) GetTransformPattern() string`

GetTransformPattern returns the TransformPattern field if non-nil, zero value otherwise.

### GetTransformPatternOk

`func (o *MappingRuleResponse) GetTransformPatternOk() (*string, bool)`

GetTransformPatternOk returns a tuple with the TransformPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformPattern

`func (o *MappingRuleResponse) SetTransformPattern(v string)`

SetTransformPattern sets TransformPattern field to given value.


### GetTransformMap

`func (o *MappingRuleResponse) GetTransformMap() map[string]interface{}`

GetTransformMap returns the TransformMap field if non-nil, zero value otherwise.

### GetTransformMapOk

`func (o *MappingRuleResponse) GetTransformMapOk() (*map[string]interface{}, bool)`

GetTransformMapOk returns a tuple with the TransformMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformMap

`func (o *MappingRuleResponse) SetTransformMap(v map[string]interface{})`

SetTransformMap sets TransformMap field to given value.


### GetTargetField

`func (o *MappingRuleResponse) GetTargetField() string`

GetTargetField returns the TargetField field if non-nil, zero value otherwise.

### GetTargetFieldOk

`func (o *MappingRuleResponse) GetTargetFieldOk() (*string, bool)`

GetTargetFieldOk returns a tuple with the TargetField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetField

`func (o *MappingRuleResponse) SetTargetField(v string)`

SetTargetField sets TargetField field to given value.


### GetPiiCheckEnabled

`func (o *MappingRuleResponse) GetPiiCheckEnabled() bool`

GetPiiCheckEnabled returns the PiiCheckEnabled field if non-nil, zero value otherwise.

### GetPiiCheckEnabledOk

`func (o *MappingRuleResponse) GetPiiCheckEnabledOk() (*bool, bool)`

GetPiiCheckEnabledOk returns a tuple with the PiiCheckEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPiiCheckEnabled

`func (o *MappingRuleResponse) SetPiiCheckEnabled(v bool)`

SetPiiCheckEnabled sets PiiCheckEnabled field to given value.


### GetCreatedAt

`func (o *MappingRuleResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MappingRuleResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MappingRuleResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *MappingRuleResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MappingRuleResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MappingRuleResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedBy

`func (o *MappingRuleResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MappingRuleResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *MappingRuleResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


