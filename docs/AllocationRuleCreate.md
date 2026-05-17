# AllocationRuleCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**RuleType** | **string** | &#39;proportional_usage&#39;, &#39;fixed_percentage&#39;, &#39;equal_split&#39;, &#39;amortization&#39; | 
**SourceFilter** | Pointer to **map[string]interface{}** |  | [optional] 
**TargetType** | **string** | &#39;customer&#39;, &#39;feature&#39;, or &#39;plan&#39; | 
**AllocationKeys** | Pointer to [**[]AllocationKey**](AllocationKey.md) |  | [optional] 
**AmortizationMonths** | Pointer to **int32** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] [default to true]
**Priority** | Pointer to **int32** |  | [optional] [default to 100]
**CreatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewAllocationRuleCreate

`func NewAllocationRuleCreate(name string, ruleType string, targetType string, ) *AllocationRuleCreate`

NewAllocationRuleCreate instantiates a new AllocationRuleCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllocationRuleCreateWithDefaults

`func NewAllocationRuleCreateWithDefaults() *AllocationRuleCreate`

NewAllocationRuleCreateWithDefaults instantiates a new AllocationRuleCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AllocationRuleCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AllocationRuleCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AllocationRuleCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AllocationRuleCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AllocationRuleCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AllocationRuleCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AllocationRuleCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRuleType

`func (o *AllocationRuleCreate) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *AllocationRuleCreate) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *AllocationRuleCreate) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.


### GetSourceFilter

`func (o *AllocationRuleCreate) GetSourceFilter() map[string]interface{}`

GetSourceFilter returns the SourceFilter field if non-nil, zero value otherwise.

### GetSourceFilterOk

`func (o *AllocationRuleCreate) GetSourceFilterOk() (*map[string]interface{}, bool)`

GetSourceFilterOk returns a tuple with the SourceFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFilter

`func (o *AllocationRuleCreate) SetSourceFilter(v map[string]interface{})`

SetSourceFilter sets SourceFilter field to given value.

### HasSourceFilter

`func (o *AllocationRuleCreate) HasSourceFilter() bool`

HasSourceFilter returns a boolean if a field has been set.

### GetTargetType

`func (o *AllocationRuleCreate) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *AllocationRuleCreate) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *AllocationRuleCreate) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.


### GetAllocationKeys

`func (o *AllocationRuleCreate) GetAllocationKeys() []AllocationKey`

GetAllocationKeys returns the AllocationKeys field if non-nil, zero value otherwise.

### GetAllocationKeysOk

`func (o *AllocationRuleCreate) GetAllocationKeysOk() (*[]AllocationKey, bool)`

GetAllocationKeysOk returns a tuple with the AllocationKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationKeys

`func (o *AllocationRuleCreate) SetAllocationKeys(v []AllocationKey)`

SetAllocationKeys sets AllocationKeys field to given value.

### HasAllocationKeys

`func (o *AllocationRuleCreate) HasAllocationKeys() bool`

HasAllocationKeys returns a boolean if a field has been set.

### GetAmortizationMonths

`func (o *AllocationRuleCreate) GetAmortizationMonths() int32`

GetAmortizationMonths returns the AmortizationMonths field if non-nil, zero value otherwise.

### GetAmortizationMonthsOk

`func (o *AllocationRuleCreate) GetAmortizationMonthsOk() (*int32, bool)`

GetAmortizationMonthsOk returns a tuple with the AmortizationMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmortizationMonths

`func (o *AllocationRuleCreate) SetAmortizationMonths(v int32)`

SetAmortizationMonths sets AmortizationMonths field to given value.

### HasAmortizationMonths

`func (o *AllocationRuleCreate) HasAmortizationMonths() bool`

HasAmortizationMonths returns a boolean if a field has been set.

### GetIsActive

`func (o *AllocationRuleCreate) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *AllocationRuleCreate) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *AllocationRuleCreate) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *AllocationRuleCreate) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetPriority

`func (o *AllocationRuleCreate) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *AllocationRuleCreate) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *AllocationRuleCreate) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *AllocationRuleCreate) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetCreatedBy

`func (o *AllocationRuleCreate) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AllocationRuleCreate) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AllocationRuleCreate) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AllocationRuleCreate) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


