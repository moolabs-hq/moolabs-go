# AllocationRuleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**RuleType** | **string** |  | 
**SourceFilter** | **map[string]interface{}** |  | 
**TargetType** | **string** |  | 
**AllocationKeys** | **[]interface{}** |  | 
**AmortizationMonths** | **int32** |  | 
**IsActive** | **bool** |  | 
**Priority** | **int32** |  | 
**CreatedBy** | **string** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 

## Methods

### NewAllocationRuleResponse

`func NewAllocationRuleResponse(id string, tenantId string, name string, description string, ruleType string, sourceFilter map[string]interface{}, targetType string, allocationKeys []interface{}, amortizationMonths int32, isActive bool, priority int32, createdBy string, createdAt string, updatedAt string, ) *AllocationRuleResponse`

NewAllocationRuleResponse instantiates a new AllocationRuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllocationRuleResponseWithDefaults

`func NewAllocationRuleResponseWithDefaults() *AllocationRuleResponse`

NewAllocationRuleResponseWithDefaults instantiates a new AllocationRuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AllocationRuleResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AllocationRuleResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AllocationRuleResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *AllocationRuleResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AllocationRuleResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AllocationRuleResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetName

`func (o *AllocationRuleResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AllocationRuleResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AllocationRuleResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AllocationRuleResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AllocationRuleResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AllocationRuleResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetRuleType

`func (o *AllocationRuleResponse) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *AllocationRuleResponse) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *AllocationRuleResponse) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.


### GetSourceFilter

`func (o *AllocationRuleResponse) GetSourceFilter() map[string]interface{}`

GetSourceFilter returns the SourceFilter field if non-nil, zero value otherwise.

### GetSourceFilterOk

`func (o *AllocationRuleResponse) GetSourceFilterOk() (*map[string]interface{}, bool)`

GetSourceFilterOk returns a tuple with the SourceFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFilter

`func (o *AllocationRuleResponse) SetSourceFilter(v map[string]interface{})`

SetSourceFilter sets SourceFilter field to given value.


### GetTargetType

`func (o *AllocationRuleResponse) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *AllocationRuleResponse) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *AllocationRuleResponse) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.


### GetAllocationKeys

`func (o *AllocationRuleResponse) GetAllocationKeys() []interface{}`

GetAllocationKeys returns the AllocationKeys field if non-nil, zero value otherwise.

### GetAllocationKeysOk

`func (o *AllocationRuleResponse) GetAllocationKeysOk() (*[]interface{}, bool)`

GetAllocationKeysOk returns a tuple with the AllocationKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationKeys

`func (o *AllocationRuleResponse) SetAllocationKeys(v []interface{})`

SetAllocationKeys sets AllocationKeys field to given value.


### GetAmortizationMonths

`func (o *AllocationRuleResponse) GetAmortizationMonths() int32`

GetAmortizationMonths returns the AmortizationMonths field if non-nil, zero value otherwise.

### GetAmortizationMonthsOk

`func (o *AllocationRuleResponse) GetAmortizationMonthsOk() (*int32, bool)`

GetAmortizationMonthsOk returns a tuple with the AmortizationMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmortizationMonths

`func (o *AllocationRuleResponse) SetAmortizationMonths(v int32)`

SetAmortizationMonths sets AmortizationMonths field to given value.


### GetIsActive

`func (o *AllocationRuleResponse) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *AllocationRuleResponse) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *AllocationRuleResponse) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetPriority

`func (o *AllocationRuleResponse) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *AllocationRuleResponse) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *AllocationRuleResponse) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetCreatedBy

`func (o *AllocationRuleResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AllocationRuleResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AllocationRuleResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatedAt

`func (o *AllocationRuleResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AllocationRuleResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AllocationRuleResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AllocationRuleResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AllocationRuleResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AllocationRuleResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


