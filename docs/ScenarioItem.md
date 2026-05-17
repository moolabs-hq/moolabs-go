# ScenarioItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScenarioIndex** | **int32** |  | 
**ScenarioLabel** | **string** |  | 
**Error** | Pointer to **string** |  | [optional] 
**FeatureKey** | Pointer to **string** |  | [optional] 
**SimulationPeriodMonths** | Pointer to **int32** |  | [optional] 
**ProjectedVolume** | Pointer to **int32** |  | [optional] 
**ProposedPricePerEvent** | Pointer to **float32** |  | [optional] 
**ProjectedRevenue** | Pointer to **float32** |  | [optional] 
**ProjectedCost** | Pointer to **float32** |  | [optional] 
**ProjectedMargin** | Pointer to **float32** |  | [optional] 
**ProjectedMarginPct** | Pointer to **float32** |  | [optional] 
**BomId** | Pointer to **string** |  | [optional] 
**BomVersion** | Pointer to **int32** |  | [optional] 
**ComponentBreakdown** | Pointer to **[]map[string]interface{}** |  | [optional] 
**RiskFlags** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ReconciliationBasis** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewScenarioItem

`func NewScenarioItem(scenarioIndex int32, scenarioLabel string, ) *ScenarioItem`

NewScenarioItem instantiates a new ScenarioItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScenarioItemWithDefaults

`func NewScenarioItemWithDefaults() *ScenarioItem`

NewScenarioItemWithDefaults instantiates a new ScenarioItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScenarioIndex

`func (o *ScenarioItem) GetScenarioIndex() int32`

GetScenarioIndex returns the ScenarioIndex field if non-nil, zero value otherwise.

### GetScenarioIndexOk

`func (o *ScenarioItem) GetScenarioIndexOk() (*int32, bool)`

GetScenarioIndexOk returns a tuple with the ScenarioIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenarioIndex

`func (o *ScenarioItem) SetScenarioIndex(v int32)`

SetScenarioIndex sets ScenarioIndex field to given value.


### GetScenarioLabel

`func (o *ScenarioItem) GetScenarioLabel() string`

GetScenarioLabel returns the ScenarioLabel field if non-nil, zero value otherwise.

### GetScenarioLabelOk

`func (o *ScenarioItem) GetScenarioLabelOk() (*string, bool)`

GetScenarioLabelOk returns a tuple with the ScenarioLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenarioLabel

`func (o *ScenarioItem) SetScenarioLabel(v string)`

SetScenarioLabel sets ScenarioLabel field to given value.


### GetError

`func (o *ScenarioItem) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ScenarioItem) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ScenarioItem) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ScenarioItem) HasError() bool`

HasError returns a boolean if a field has been set.

### GetFeatureKey

`func (o *ScenarioItem) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *ScenarioItem) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *ScenarioItem) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.

### HasFeatureKey

`func (o *ScenarioItem) HasFeatureKey() bool`

HasFeatureKey returns a boolean if a field has been set.

### GetSimulationPeriodMonths

`func (o *ScenarioItem) GetSimulationPeriodMonths() int32`

GetSimulationPeriodMonths returns the SimulationPeriodMonths field if non-nil, zero value otherwise.

### GetSimulationPeriodMonthsOk

`func (o *ScenarioItem) GetSimulationPeriodMonthsOk() (*int32, bool)`

GetSimulationPeriodMonthsOk returns a tuple with the SimulationPeriodMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimulationPeriodMonths

`func (o *ScenarioItem) SetSimulationPeriodMonths(v int32)`

SetSimulationPeriodMonths sets SimulationPeriodMonths field to given value.

### HasSimulationPeriodMonths

`func (o *ScenarioItem) HasSimulationPeriodMonths() bool`

HasSimulationPeriodMonths returns a boolean if a field has been set.

### GetProjectedVolume

`func (o *ScenarioItem) GetProjectedVolume() int32`

GetProjectedVolume returns the ProjectedVolume field if non-nil, zero value otherwise.

### GetProjectedVolumeOk

`func (o *ScenarioItem) GetProjectedVolumeOk() (*int32, bool)`

GetProjectedVolumeOk returns a tuple with the ProjectedVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectedVolume

`func (o *ScenarioItem) SetProjectedVolume(v int32)`

SetProjectedVolume sets ProjectedVolume field to given value.

### HasProjectedVolume

`func (o *ScenarioItem) HasProjectedVolume() bool`

HasProjectedVolume returns a boolean if a field has been set.

### GetProposedPricePerEvent

`func (o *ScenarioItem) GetProposedPricePerEvent() float32`

GetProposedPricePerEvent returns the ProposedPricePerEvent field if non-nil, zero value otherwise.

### GetProposedPricePerEventOk

`func (o *ScenarioItem) GetProposedPricePerEventOk() (*float32, bool)`

GetProposedPricePerEventOk returns a tuple with the ProposedPricePerEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedPricePerEvent

`func (o *ScenarioItem) SetProposedPricePerEvent(v float32)`

SetProposedPricePerEvent sets ProposedPricePerEvent field to given value.

### HasProposedPricePerEvent

`func (o *ScenarioItem) HasProposedPricePerEvent() bool`

HasProposedPricePerEvent returns a boolean if a field has been set.

### GetProjectedRevenue

`func (o *ScenarioItem) GetProjectedRevenue() float32`

GetProjectedRevenue returns the ProjectedRevenue field if non-nil, zero value otherwise.

### GetProjectedRevenueOk

`func (o *ScenarioItem) GetProjectedRevenueOk() (*float32, bool)`

GetProjectedRevenueOk returns a tuple with the ProjectedRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectedRevenue

`func (o *ScenarioItem) SetProjectedRevenue(v float32)`

SetProjectedRevenue sets ProjectedRevenue field to given value.

### HasProjectedRevenue

`func (o *ScenarioItem) HasProjectedRevenue() bool`

HasProjectedRevenue returns a boolean if a field has been set.

### GetProjectedCost

`func (o *ScenarioItem) GetProjectedCost() float32`

GetProjectedCost returns the ProjectedCost field if non-nil, zero value otherwise.

### GetProjectedCostOk

`func (o *ScenarioItem) GetProjectedCostOk() (*float32, bool)`

GetProjectedCostOk returns a tuple with the ProjectedCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectedCost

`func (o *ScenarioItem) SetProjectedCost(v float32)`

SetProjectedCost sets ProjectedCost field to given value.

### HasProjectedCost

`func (o *ScenarioItem) HasProjectedCost() bool`

HasProjectedCost returns a boolean if a field has been set.

### GetProjectedMargin

`func (o *ScenarioItem) GetProjectedMargin() float32`

GetProjectedMargin returns the ProjectedMargin field if non-nil, zero value otherwise.

### GetProjectedMarginOk

`func (o *ScenarioItem) GetProjectedMarginOk() (*float32, bool)`

GetProjectedMarginOk returns a tuple with the ProjectedMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectedMargin

`func (o *ScenarioItem) SetProjectedMargin(v float32)`

SetProjectedMargin sets ProjectedMargin field to given value.

### HasProjectedMargin

`func (o *ScenarioItem) HasProjectedMargin() bool`

HasProjectedMargin returns a boolean if a field has been set.

### GetProjectedMarginPct

`func (o *ScenarioItem) GetProjectedMarginPct() float32`

GetProjectedMarginPct returns the ProjectedMarginPct field if non-nil, zero value otherwise.

### GetProjectedMarginPctOk

`func (o *ScenarioItem) GetProjectedMarginPctOk() (*float32, bool)`

GetProjectedMarginPctOk returns a tuple with the ProjectedMarginPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectedMarginPct

`func (o *ScenarioItem) SetProjectedMarginPct(v float32)`

SetProjectedMarginPct sets ProjectedMarginPct field to given value.

### HasProjectedMarginPct

`func (o *ScenarioItem) HasProjectedMarginPct() bool`

HasProjectedMarginPct returns a boolean if a field has been set.

### GetBomId

`func (o *ScenarioItem) GetBomId() string`

GetBomId returns the BomId field if non-nil, zero value otherwise.

### GetBomIdOk

`func (o *ScenarioItem) GetBomIdOk() (*string, bool)`

GetBomIdOk returns a tuple with the BomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBomId

`func (o *ScenarioItem) SetBomId(v string)`

SetBomId sets BomId field to given value.

### HasBomId

`func (o *ScenarioItem) HasBomId() bool`

HasBomId returns a boolean if a field has been set.

### GetBomVersion

`func (o *ScenarioItem) GetBomVersion() int32`

GetBomVersion returns the BomVersion field if non-nil, zero value otherwise.

### GetBomVersionOk

`func (o *ScenarioItem) GetBomVersionOk() (*int32, bool)`

GetBomVersionOk returns a tuple with the BomVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBomVersion

`func (o *ScenarioItem) SetBomVersion(v int32)`

SetBomVersion sets BomVersion field to given value.

### HasBomVersion

`func (o *ScenarioItem) HasBomVersion() bool`

HasBomVersion returns a boolean if a field has been set.

### GetComponentBreakdown

`func (o *ScenarioItem) GetComponentBreakdown() []map[string]interface{}`

GetComponentBreakdown returns the ComponentBreakdown field if non-nil, zero value otherwise.

### GetComponentBreakdownOk

`func (o *ScenarioItem) GetComponentBreakdownOk() (*[]map[string]interface{}, bool)`

GetComponentBreakdownOk returns a tuple with the ComponentBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentBreakdown

`func (o *ScenarioItem) SetComponentBreakdown(v []map[string]interface{})`

SetComponentBreakdown sets ComponentBreakdown field to given value.

### HasComponentBreakdown

`func (o *ScenarioItem) HasComponentBreakdown() bool`

HasComponentBreakdown returns a boolean if a field has been set.

### GetRiskFlags

`func (o *ScenarioItem) GetRiskFlags() []map[string]interface{}`

GetRiskFlags returns the RiskFlags field if non-nil, zero value otherwise.

### GetRiskFlagsOk

`func (o *ScenarioItem) GetRiskFlagsOk() (*[]map[string]interface{}, bool)`

GetRiskFlagsOk returns a tuple with the RiskFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskFlags

`func (o *ScenarioItem) SetRiskFlags(v []map[string]interface{})`

SetRiskFlags sets RiskFlags field to given value.

### HasRiskFlags

`func (o *ScenarioItem) HasRiskFlags() bool`

HasRiskFlags returns a boolean if a field has been set.

### GetReconciliationBasis

`func (o *ScenarioItem) GetReconciliationBasis() map[string]interface{}`

GetReconciliationBasis returns the ReconciliationBasis field if non-nil, zero value otherwise.

### GetReconciliationBasisOk

`func (o *ScenarioItem) GetReconciliationBasisOk() (*map[string]interface{}, bool)`

GetReconciliationBasisOk returns a tuple with the ReconciliationBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconciliationBasis

`func (o *ScenarioItem) SetReconciliationBasis(v map[string]interface{})`

SetReconciliationBasis sets ReconciliationBasis field to given value.

### HasReconciliationBasis

`func (o *ScenarioItem) HasReconciliationBasis() bool`

HasReconciliationBasis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


