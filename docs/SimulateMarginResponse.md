# SimulateMarginResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureKey** | **string** |  | 
**SimulationPeriodMonths** | **int32** |  | 
**ProjectedVolume** | **int32** |  | 
**ProposedPricePerEvent** | **float32** |  | 
**ProjectedRevenue** | **float32** |  | 
**ProjectedCost** | **float32** |  | 
**ProjectedMargin** | **float32** |  | 
**ProjectedMarginPct** | **float32** |  | 
**BomId** | **string** |  | 
**BomVersion** | **int32** |  | 
**ComponentBreakdown** | **[]map[string]interface{}** |  | 
**RiskFlags** | **[]map[string]interface{}** |  | 
**ReconciliationBasis** | [**ReconciliationBasisOut**](ReconciliationBasisOut.md) |  | 

## Methods

### NewSimulateMarginResponse

`func NewSimulateMarginResponse(featureKey string, simulationPeriodMonths int32, projectedVolume int32, proposedPricePerEvent float32, projectedRevenue float32, projectedCost float32, projectedMargin float32, projectedMarginPct float32, bomId string, bomVersion int32, componentBreakdown []map[string]interface{}, riskFlags []map[string]interface{}, reconciliationBasis ReconciliationBasisOut, ) *SimulateMarginResponse`

NewSimulateMarginResponse instantiates a new SimulateMarginResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimulateMarginResponseWithDefaults

`func NewSimulateMarginResponseWithDefaults() *SimulateMarginResponse`

NewSimulateMarginResponseWithDefaults instantiates a new SimulateMarginResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatureKey

`func (o *SimulateMarginResponse) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *SimulateMarginResponse) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *SimulateMarginResponse) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.


### GetSimulationPeriodMonths

`func (o *SimulateMarginResponse) GetSimulationPeriodMonths() int32`

GetSimulationPeriodMonths returns the SimulationPeriodMonths field if non-nil, zero value otherwise.

### GetSimulationPeriodMonthsOk

`func (o *SimulateMarginResponse) GetSimulationPeriodMonthsOk() (*int32, bool)`

GetSimulationPeriodMonthsOk returns a tuple with the SimulationPeriodMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimulationPeriodMonths

`func (o *SimulateMarginResponse) SetSimulationPeriodMonths(v int32)`

SetSimulationPeriodMonths sets SimulationPeriodMonths field to given value.


### GetProjectedVolume

`func (o *SimulateMarginResponse) GetProjectedVolume() int32`

GetProjectedVolume returns the ProjectedVolume field if non-nil, zero value otherwise.

### GetProjectedVolumeOk

`func (o *SimulateMarginResponse) GetProjectedVolumeOk() (*int32, bool)`

GetProjectedVolumeOk returns a tuple with the ProjectedVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectedVolume

`func (o *SimulateMarginResponse) SetProjectedVolume(v int32)`

SetProjectedVolume sets ProjectedVolume field to given value.


### GetProposedPricePerEvent

`func (o *SimulateMarginResponse) GetProposedPricePerEvent() float32`

GetProposedPricePerEvent returns the ProposedPricePerEvent field if non-nil, zero value otherwise.

### GetProposedPricePerEventOk

`func (o *SimulateMarginResponse) GetProposedPricePerEventOk() (*float32, bool)`

GetProposedPricePerEventOk returns a tuple with the ProposedPricePerEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedPricePerEvent

`func (o *SimulateMarginResponse) SetProposedPricePerEvent(v float32)`

SetProposedPricePerEvent sets ProposedPricePerEvent field to given value.


### GetProjectedRevenue

`func (o *SimulateMarginResponse) GetProjectedRevenue() float32`

GetProjectedRevenue returns the ProjectedRevenue field if non-nil, zero value otherwise.

### GetProjectedRevenueOk

`func (o *SimulateMarginResponse) GetProjectedRevenueOk() (*float32, bool)`

GetProjectedRevenueOk returns a tuple with the ProjectedRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectedRevenue

`func (o *SimulateMarginResponse) SetProjectedRevenue(v float32)`

SetProjectedRevenue sets ProjectedRevenue field to given value.


### GetProjectedCost

`func (o *SimulateMarginResponse) GetProjectedCost() float32`

GetProjectedCost returns the ProjectedCost field if non-nil, zero value otherwise.

### GetProjectedCostOk

`func (o *SimulateMarginResponse) GetProjectedCostOk() (*float32, bool)`

GetProjectedCostOk returns a tuple with the ProjectedCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectedCost

`func (o *SimulateMarginResponse) SetProjectedCost(v float32)`

SetProjectedCost sets ProjectedCost field to given value.


### GetProjectedMargin

`func (o *SimulateMarginResponse) GetProjectedMargin() float32`

GetProjectedMargin returns the ProjectedMargin field if non-nil, zero value otherwise.

### GetProjectedMarginOk

`func (o *SimulateMarginResponse) GetProjectedMarginOk() (*float32, bool)`

GetProjectedMarginOk returns a tuple with the ProjectedMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectedMargin

`func (o *SimulateMarginResponse) SetProjectedMargin(v float32)`

SetProjectedMargin sets ProjectedMargin field to given value.


### GetProjectedMarginPct

`func (o *SimulateMarginResponse) GetProjectedMarginPct() float32`

GetProjectedMarginPct returns the ProjectedMarginPct field if non-nil, zero value otherwise.

### GetProjectedMarginPctOk

`func (o *SimulateMarginResponse) GetProjectedMarginPctOk() (*float32, bool)`

GetProjectedMarginPctOk returns a tuple with the ProjectedMarginPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectedMarginPct

`func (o *SimulateMarginResponse) SetProjectedMarginPct(v float32)`

SetProjectedMarginPct sets ProjectedMarginPct field to given value.


### GetBomId

`func (o *SimulateMarginResponse) GetBomId() string`

GetBomId returns the BomId field if non-nil, zero value otherwise.

### GetBomIdOk

`func (o *SimulateMarginResponse) GetBomIdOk() (*string, bool)`

GetBomIdOk returns a tuple with the BomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBomId

`func (o *SimulateMarginResponse) SetBomId(v string)`

SetBomId sets BomId field to given value.


### GetBomVersion

`func (o *SimulateMarginResponse) GetBomVersion() int32`

GetBomVersion returns the BomVersion field if non-nil, zero value otherwise.

### GetBomVersionOk

`func (o *SimulateMarginResponse) GetBomVersionOk() (*int32, bool)`

GetBomVersionOk returns a tuple with the BomVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBomVersion

`func (o *SimulateMarginResponse) SetBomVersion(v int32)`

SetBomVersion sets BomVersion field to given value.


### GetComponentBreakdown

`func (o *SimulateMarginResponse) GetComponentBreakdown() []map[string]interface{}`

GetComponentBreakdown returns the ComponentBreakdown field if non-nil, zero value otherwise.

### GetComponentBreakdownOk

`func (o *SimulateMarginResponse) GetComponentBreakdownOk() (*[]map[string]interface{}, bool)`

GetComponentBreakdownOk returns a tuple with the ComponentBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentBreakdown

`func (o *SimulateMarginResponse) SetComponentBreakdown(v []map[string]interface{})`

SetComponentBreakdown sets ComponentBreakdown field to given value.


### GetRiskFlags

`func (o *SimulateMarginResponse) GetRiskFlags() []map[string]interface{}`

GetRiskFlags returns the RiskFlags field if non-nil, zero value otherwise.

### GetRiskFlagsOk

`func (o *SimulateMarginResponse) GetRiskFlagsOk() (*[]map[string]interface{}, bool)`

GetRiskFlagsOk returns a tuple with the RiskFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskFlags

`func (o *SimulateMarginResponse) SetRiskFlags(v []map[string]interface{})`

SetRiskFlags sets RiskFlags field to given value.


### GetReconciliationBasis

`func (o *SimulateMarginResponse) GetReconciliationBasis() ReconciliationBasisOut`

GetReconciliationBasis returns the ReconciliationBasis field if non-nil, zero value otherwise.

### GetReconciliationBasisOk

`func (o *SimulateMarginResponse) GetReconciliationBasisOk() (*ReconciliationBasisOut, bool)`

GetReconciliationBasisOk returns a tuple with the ReconciliationBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconciliationBasis

`func (o *SimulateMarginResponse) SetReconciliationBasis(v ReconciliationBasisOut)`

SetReconciliationBasis sets ReconciliationBasis field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


