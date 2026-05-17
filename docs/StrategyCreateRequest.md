# StrategyCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**RiskTier** | **string** |  | 
**Stages** | [**[]StageSchema**](StageSchema.md) |  | 
**ChannelPolicy** | Pointer to [**map[string]ChannelPolicySchema**](ChannelPolicySchema.md) |  | [optional] 
**SegmentFilters** | Pointer to **map[string]interface{}** |  | [optional] 
**EscalationPolicy** | Pointer to **map[string]interface{}** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] [default to 100]
**IsDefault** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewStrategyCreateRequest

`func NewStrategyCreateRequest(name string, riskTier string, stages []StageSchema, ) *StrategyCreateRequest`

NewStrategyCreateRequest instantiates a new StrategyCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStrategyCreateRequestWithDefaults

`func NewStrategyCreateRequestWithDefaults() *StrategyCreateRequest`

NewStrategyCreateRequestWithDefaults instantiates a new StrategyCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StrategyCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StrategyCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StrategyCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRiskTier

`func (o *StrategyCreateRequest) GetRiskTier() string`

GetRiskTier returns the RiskTier field if non-nil, zero value otherwise.

### GetRiskTierOk

`func (o *StrategyCreateRequest) GetRiskTierOk() (*string, bool)`

GetRiskTierOk returns a tuple with the RiskTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskTier

`func (o *StrategyCreateRequest) SetRiskTier(v string)`

SetRiskTier sets RiskTier field to given value.


### GetStages

`func (o *StrategyCreateRequest) GetStages() []StageSchema`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *StrategyCreateRequest) GetStagesOk() (*[]StageSchema, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *StrategyCreateRequest) SetStages(v []StageSchema)`

SetStages sets Stages field to given value.


### GetChannelPolicy

`func (o *StrategyCreateRequest) GetChannelPolicy() map[string]ChannelPolicySchema`

GetChannelPolicy returns the ChannelPolicy field if non-nil, zero value otherwise.

### GetChannelPolicyOk

`func (o *StrategyCreateRequest) GetChannelPolicyOk() (*map[string]ChannelPolicySchema, bool)`

GetChannelPolicyOk returns a tuple with the ChannelPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelPolicy

`func (o *StrategyCreateRequest) SetChannelPolicy(v map[string]ChannelPolicySchema)`

SetChannelPolicy sets ChannelPolicy field to given value.

### HasChannelPolicy

`func (o *StrategyCreateRequest) HasChannelPolicy() bool`

HasChannelPolicy returns a boolean if a field has been set.

### GetSegmentFilters

`func (o *StrategyCreateRequest) GetSegmentFilters() map[string]interface{}`

GetSegmentFilters returns the SegmentFilters field if non-nil, zero value otherwise.

### GetSegmentFiltersOk

`func (o *StrategyCreateRequest) GetSegmentFiltersOk() (*map[string]interface{}, bool)`

GetSegmentFiltersOk returns a tuple with the SegmentFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentFilters

`func (o *StrategyCreateRequest) SetSegmentFilters(v map[string]interface{})`

SetSegmentFilters sets SegmentFilters field to given value.

### HasSegmentFilters

`func (o *StrategyCreateRequest) HasSegmentFilters() bool`

HasSegmentFilters returns a boolean if a field has been set.

### GetEscalationPolicy

`func (o *StrategyCreateRequest) GetEscalationPolicy() map[string]interface{}`

GetEscalationPolicy returns the EscalationPolicy field if non-nil, zero value otherwise.

### GetEscalationPolicyOk

`func (o *StrategyCreateRequest) GetEscalationPolicyOk() (*map[string]interface{}, bool)`

GetEscalationPolicyOk returns a tuple with the EscalationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEscalationPolicy

`func (o *StrategyCreateRequest) SetEscalationPolicy(v map[string]interface{})`

SetEscalationPolicy sets EscalationPolicy field to given value.

### HasEscalationPolicy

`func (o *StrategyCreateRequest) HasEscalationPolicy() bool`

HasEscalationPolicy returns a boolean if a field has been set.

### GetPriority

`func (o *StrategyCreateRequest) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *StrategyCreateRequest) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *StrategyCreateRequest) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *StrategyCreateRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetIsDefault

`func (o *StrategyCreateRequest) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *StrategyCreateRequest) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *StrategyCreateRequest) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *StrategyCreateRequest) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


