# StrategyUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**RiskTier** | Pointer to **string** |  | [optional] 
**Stages** | Pointer to [**[]StageSchema**](StageSchema.md) |  | [optional] 
**ChannelPolicy** | Pointer to [**map[string]ChannelPolicySchema**](ChannelPolicySchema.md) |  | [optional] 
**SegmentFilters** | Pointer to **map[string]interface{}** |  | [optional] 
**EscalationPolicy** | Pointer to **map[string]interface{}** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 

## Methods

### NewStrategyUpdateRequest

`func NewStrategyUpdateRequest() *StrategyUpdateRequest`

NewStrategyUpdateRequest instantiates a new StrategyUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStrategyUpdateRequestWithDefaults

`func NewStrategyUpdateRequestWithDefaults() *StrategyUpdateRequest`

NewStrategyUpdateRequestWithDefaults instantiates a new StrategyUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StrategyUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StrategyUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StrategyUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StrategyUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRiskTier

`func (o *StrategyUpdateRequest) GetRiskTier() string`

GetRiskTier returns the RiskTier field if non-nil, zero value otherwise.

### GetRiskTierOk

`func (o *StrategyUpdateRequest) GetRiskTierOk() (*string, bool)`

GetRiskTierOk returns a tuple with the RiskTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskTier

`func (o *StrategyUpdateRequest) SetRiskTier(v string)`

SetRiskTier sets RiskTier field to given value.

### HasRiskTier

`func (o *StrategyUpdateRequest) HasRiskTier() bool`

HasRiskTier returns a boolean if a field has been set.

### GetStages

`func (o *StrategyUpdateRequest) GetStages() []StageSchema`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *StrategyUpdateRequest) GetStagesOk() (*[]StageSchema, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *StrategyUpdateRequest) SetStages(v []StageSchema)`

SetStages sets Stages field to given value.

### HasStages

`func (o *StrategyUpdateRequest) HasStages() bool`

HasStages returns a boolean if a field has been set.

### GetChannelPolicy

`func (o *StrategyUpdateRequest) GetChannelPolicy() map[string]ChannelPolicySchema`

GetChannelPolicy returns the ChannelPolicy field if non-nil, zero value otherwise.

### GetChannelPolicyOk

`func (o *StrategyUpdateRequest) GetChannelPolicyOk() (*map[string]ChannelPolicySchema, bool)`

GetChannelPolicyOk returns a tuple with the ChannelPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelPolicy

`func (o *StrategyUpdateRequest) SetChannelPolicy(v map[string]ChannelPolicySchema)`

SetChannelPolicy sets ChannelPolicy field to given value.

### HasChannelPolicy

`func (o *StrategyUpdateRequest) HasChannelPolicy() bool`

HasChannelPolicy returns a boolean if a field has been set.

### GetSegmentFilters

`func (o *StrategyUpdateRequest) GetSegmentFilters() map[string]interface{}`

GetSegmentFilters returns the SegmentFilters field if non-nil, zero value otherwise.

### GetSegmentFiltersOk

`func (o *StrategyUpdateRequest) GetSegmentFiltersOk() (*map[string]interface{}, bool)`

GetSegmentFiltersOk returns a tuple with the SegmentFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentFilters

`func (o *StrategyUpdateRequest) SetSegmentFilters(v map[string]interface{})`

SetSegmentFilters sets SegmentFilters field to given value.

### HasSegmentFilters

`func (o *StrategyUpdateRequest) HasSegmentFilters() bool`

HasSegmentFilters returns a boolean if a field has been set.

### GetEscalationPolicy

`func (o *StrategyUpdateRequest) GetEscalationPolicy() map[string]interface{}`

GetEscalationPolicy returns the EscalationPolicy field if non-nil, zero value otherwise.

### GetEscalationPolicyOk

`func (o *StrategyUpdateRequest) GetEscalationPolicyOk() (*map[string]interface{}, bool)`

GetEscalationPolicyOk returns a tuple with the EscalationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEscalationPolicy

`func (o *StrategyUpdateRequest) SetEscalationPolicy(v map[string]interface{})`

SetEscalationPolicy sets EscalationPolicy field to given value.

### HasEscalationPolicy

`func (o *StrategyUpdateRequest) HasEscalationPolicy() bool`

HasEscalationPolicy returns a boolean if a field has been set.

### GetPriority

`func (o *StrategyUpdateRequest) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *StrategyUpdateRequest) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *StrategyUpdateRequest) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *StrategyUpdateRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetIsDefault

`func (o *StrategyUpdateRequest) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *StrategyUpdateRequest) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *StrategyUpdateRequest) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *StrategyUpdateRequest) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


