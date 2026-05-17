# SimulateMarginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Tenant UUID | 
**FeatureKey** | **string** | Feature key to simulate | 
**ProposedPricePerEvent** | **float32** | Proposed price per billing event | 
**ProjectedVolume** | **int32** | Billing events per month | 
**ProposedModelMix** | Pointer to **map[string]float32** | Model mix fractions — must sum to 1.0. E.g. {&#39;claude-sonnet-4.6&#39;: 0.7, &#39;claude-haiku-3.5&#39;: 0.3} | [optional] 
**SimulationPeriodMonths** | Pointer to **int32** | Number of months to simulate | [optional] [default to 1]
**Label** | Pointer to **string** | Optional label for this scenario | [optional] 

## Methods

### NewSimulateMarginRequest

`func NewSimulateMarginRequest(tenantId string, featureKey string, proposedPricePerEvent float32, projectedVolume int32, ) *SimulateMarginRequest`

NewSimulateMarginRequest instantiates a new SimulateMarginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimulateMarginRequestWithDefaults

`func NewSimulateMarginRequestWithDefaults() *SimulateMarginRequest`

NewSimulateMarginRequestWithDefaults instantiates a new SimulateMarginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *SimulateMarginRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SimulateMarginRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SimulateMarginRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetFeatureKey

`func (o *SimulateMarginRequest) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *SimulateMarginRequest) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *SimulateMarginRequest) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.


### GetProposedPricePerEvent

`func (o *SimulateMarginRequest) GetProposedPricePerEvent() float32`

GetProposedPricePerEvent returns the ProposedPricePerEvent field if non-nil, zero value otherwise.

### GetProposedPricePerEventOk

`func (o *SimulateMarginRequest) GetProposedPricePerEventOk() (*float32, bool)`

GetProposedPricePerEventOk returns a tuple with the ProposedPricePerEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedPricePerEvent

`func (o *SimulateMarginRequest) SetProposedPricePerEvent(v float32)`

SetProposedPricePerEvent sets ProposedPricePerEvent field to given value.


### GetProjectedVolume

`func (o *SimulateMarginRequest) GetProjectedVolume() int32`

GetProjectedVolume returns the ProjectedVolume field if non-nil, zero value otherwise.

### GetProjectedVolumeOk

`func (o *SimulateMarginRequest) GetProjectedVolumeOk() (*int32, bool)`

GetProjectedVolumeOk returns a tuple with the ProjectedVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectedVolume

`func (o *SimulateMarginRequest) SetProjectedVolume(v int32)`

SetProjectedVolume sets ProjectedVolume field to given value.


### GetProposedModelMix

`func (o *SimulateMarginRequest) GetProposedModelMix() map[string]float32`

GetProposedModelMix returns the ProposedModelMix field if non-nil, zero value otherwise.

### GetProposedModelMixOk

`func (o *SimulateMarginRequest) GetProposedModelMixOk() (*map[string]float32, bool)`

GetProposedModelMixOk returns a tuple with the ProposedModelMix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedModelMix

`func (o *SimulateMarginRequest) SetProposedModelMix(v map[string]float32)`

SetProposedModelMix sets ProposedModelMix field to given value.

### HasProposedModelMix

`func (o *SimulateMarginRequest) HasProposedModelMix() bool`

HasProposedModelMix returns a boolean if a field has been set.

### GetSimulationPeriodMonths

`func (o *SimulateMarginRequest) GetSimulationPeriodMonths() int32`

GetSimulationPeriodMonths returns the SimulationPeriodMonths field if non-nil, zero value otherwise.

### GetSimulationPeriodMonthsOk

`func (o *SimulateMarginRequest) GetSimulationPeriodMonthsOk() (*int32, bool)`

GetSimulationPeriodMonthsOk returns a tuple with the SimulationPeriodMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimulationPeriodMonths

`func (o *SimulateMarginRequest) SetSimulationPeriodMonths(v int32)`

SetSimulationPeriodMonths sets SimulationPeriodMonths field to given value.

### HasSimulationPeriodMonths

`func (o *SimulateMarginRequest) HasSimulationPeriodMonths() bool`

HasSimulationPeriodMonths returns a boolean if a field has been set.

### GetLabel

`func (o *SimulateMarginRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SimulateMarginRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SimulateMarginRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *SimulateMarginRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


