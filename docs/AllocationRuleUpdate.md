# AllocationRuleUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**RuleType** | Pointer to **string** |  | [optional] 
**SourceFilter** | Pointer to **map[string]interface{}** |  | [optional] 
**TargetType** | Pointer to **string** |  | [optional] 
**AllocationKeys** | Pointer to [**[]AllocationKey**](AllocationKey.md) |  | [optional] 
**AmortizationMonths** | Pointer to **int32** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 

## Methods

### NewAllocationRuleUpdate

`func NewAllocationRuleUpdate() *AllocationRuleUpdate`

NewAllocationRuleUpdate instantiates a new AllocationRuleUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllocationRuleUpdateWithDefaults

`func NewAllocationRuleUpdateWithDefaults() *AllocationRuleUpdate`

NewAllocationRuleUpdateWithDefaults instantiates a new AllocationRuleUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AllocationRuleUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AllocationRuleUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AllocationRuleUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AllocationRuleUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AllocationRuleUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AllocationRuleUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AllocationRuleUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AllocationRuleUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRuleType

`func (o *AllocationRuleUpdate) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *AllocationRuleUpdate) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *AllocationRuleUpdate) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.

### HasRuleType

`func (o *AllocationRuleUpdate) HasRuleType() bool`

HasRuleType returns a boolean if a field has been set.

### GetSourceFilter

`func (o *AllocationRuleUpdate) GetSourceFilter() map[string]interface{}`

GetSourceFilter returns the SourceFilter field if non-nil, zero value otherwise.

### GetSourceFilterOk

`func (o *AllocationRuleUpdate) GetSourceFilterOk() (*map[string]interface{}, bool)`

GetSourceFilterOk returns a tuple with the SourceFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFilter

`func (o *AllocationRuleUpdate) SetSourceFilter(v map[string]interface{})`

SetSourceFilter sets SourceFilter field to given value.

### HasSourceFilter

`func (o *AllocationRuleUpdate) HasSourceFilter() bool`

HasSourceFilter returns a boolean if a field has been set.

### GetTargetType

`func (o *AllocationRuleUpdate) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *AllocationRuleUpdate) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *AllocationRuleUpdate) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *AllocationRuleUpdate) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### GetAllocationKeys

`func (o *AllocationRuleUpdate) GetAllocationKeys() []AllocationKey`

GetAllocationKeys returns the AllocationKeys field if non-nil, zero value otherwise.

### GetAllocationKeysOk

`func (o *AllocationRuleUpdate) GetAllocationKeysOk() (*[]AllocationKey, bool)`

GetAllocationKeysOk returns a tuple with the AllocationKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationKeys

`func (o *AllocationRuleUpdate) SetAllocationKeys(v []AllocationKey)`

SetAllocationKeys sets AllocationKeys field to given value.

### HasAllocationKeys

`func (o *AllocationRuleUpdate) HasAllocationKeys() bool`

HasAllocationKeys returns a boolean if a field has been set.

### GetAmortizationMonths

`func (o *AllocationRuleUpdate) GetAmortizationMonths() int32`

GetAmortizationMonths returns the AmortizationMonths field if non-nil, zero value otherwise.

### GetAmortizationMonthsOk

`func (o *AllocationRuleUpdate) GetAmortizationMonthsOk() (*int32, bool)`

GetAmortizationMonthsOk returns a tuple with the AmortizationMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmortizationMonths

`func (o *AllocationRuleUpdate) SetAmortizationMonths(v int32)`

SetAmortizationMonths sets AmortizationMonths field to given value.

### HasAmortizationMonths

`func (o *AllocationRuleUpdate) HasAmortizationMonths() bool`

HasAmortizationMonths returns a boolean if a field has been set.

### GetIsActive

`func (o *AllocationRuleUpdate) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *AllocationRuleUpdate) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *AllocationRuleUpdate) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *AllocationRuleUpdate) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetPriority

`func (o *AllocationRuleUpdate) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *AllocationRuleUpdate) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *AllocationRuleUpdate) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *AllocationRuleUpdate) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


