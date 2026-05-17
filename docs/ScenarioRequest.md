# ScenarioRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scenarios** | [**[]SimulateMarginRequest**](SimulateMarginRequest.md) | List of simulation scenarios for side-by-side comparison | 

## Methods

### NewScenarioRequest

`func NewScenarioRequest(scenarios []SimulateMarginRequest, ) *ScenarioRequest`

NewScenarioRequest instantiates a new ScenarioRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScenarioRequestWithDefaults

`func NewScenarioRequestWithDefaults() *ScenarioRequest`

NewScenarioRequestWithDefaults instantiates a new ScenarioRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScenarios

`func (o *ScenarioRequest) GetScenarios() []SimulateMarginRequest`

GetScenarios returns the Scenarios field if non-nil, zero value otherwise.

### GetScenariosOk

`func (o *ScenarioRequest) GetScenariosOk() (*[]SimulateMarginRequest, bool)`

GetScenariosOk returns a tuple with the Scenarios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenarios

`func (o *ScenarioRequest) SetScenarios(v []SimulateMarginRequest)`

SetScenarios sets Scenarios field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


