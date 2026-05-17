# ScenarioCompareResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalScenarios** | **int32** |  | 
**Scenarios** | [**[]ScenarioItem**](ScenarioItem.md) |  | 

## Methods

### NewScenarioCompareResponse

`func NewScenarioCompareResponse(totalScenarios int32, scenarios []ScenarioItem, ) *ScenarioCompareResponse`

NewScenarioCompareResponse instantiates a new ScenarioCompareResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScenarioCompareResponseWithDefaults

`func NewScenarioCompareResponseWithDefaults() *ScenarioCompareResponse`

NewScenarioCompareResponseWithDefaults instantiates a new ScenarioCompareResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalScenarios

`func (o *ScenarioCompareResponse) GetTotalScenarios() int32`

GetTotalScenarios returns the TotalScenarios field if non-nil, zero value otherwise.

### GetTotalScenariosOk

`func (o *ScenarioCompareResponse) GetTotalScenariosOk() (*int32, bool)`

GetTotalScenariosOk returns a tuple with the TotalScenarios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalScenarios

`func (o *ScenarioCompareResponse) SetTotalScenarios(v int32)`

SetTotalScenarios sets TotalScenarios field to given value.


### GetScenarios

`func (o *ScenarioCompareResponse) GetScenarios() []ScenarioItem`

GetScenarios returns the Scenarios field if non-nil, zero value otherwise.

### GetScenariosOk

`func (o *ScenarioCompareResponse) GetScenariosOk() (*[]ScenarioItem, bool)`

GetScenariosOk returns a tuple with the Scenarios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenarios

`func (o *ScenarioCompareResponse) SetScenarios(v []ScenarioItem)`

SetScenarios sets Scenarios field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


