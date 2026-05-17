# PolicyUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxAutonomy** | Pointer to **int32** |  | [optional] 
**AllowedActions** | Pointer to **map[string]interface{}** |  | [optional] 
**RequiresHumanFor** | Pointer to **map[string]interface{}** |  | [optional] 
**PolicyVersion** | Pointer to **string** |  | [optional] 

## Methods

### NewPolicyUpdateRequest

`func NewPolicyUpdateRequest() *PolicyUpdateRequest`

NewPolicyUpdateRequest instantiates a new PolicyUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyUpdateRequestWithDefaults

`func NewPolicyUpdateRequestWithDefaults() *PolicyUpdateRequest`

NewPolicyUpdateRequestWithDefaults instantiates a new PolicyUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxAutonomy

`func (o *PolicyUpdateRequest) GetMaxAutonomy() int32`

GetMaxAutonomy returns the MaxAutonomy field if non-nil, zero value otherwise.

### GetMaxAutonomyOk

`func (o *PolicyUpdateRequest) GetMaxAutonomyOk() (*int32, bool)`

GetMaxAutonomyOk returns a tuple with the MaxAutonomy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAutonomy

`func (o *PolicyUpdateRequest) SetMaxAutonomy(v int32)`

SetMaxAutonomy sets MaxAutonomy field to given value.

### HasMaxAutonomy

`func (o *PolicyUpdateRequest) HasMaxAutonomy() bool`

HasMaxAutonomy returns a boolean if a field has been set.

### GetAllowedActions

`func (o *PolicyUpdateRequest) GetAllowedActions() map[string]interface{}`

GetAllowedActions returns the AllowedActions field if non-nil, zero value otherwise.

### GetAllowedActionsOk

`func (o *PolicyUpdateRequest) GetAllowedActionsOk() (*map[string]interface{}, bool)`

GetAllowedActionsOk returns a tuple with the AllowedActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedActions

`func (o *PolicyUpdateRequest) SetAllowedActions(v map[string]interface{})`

SetAllowedActions sets AllowedActions field to given value.

### HasAllowedActions

`func (o *PolicyUpdateRequest) HasAllowedActions() bool`

HasAllowedActions returns a boolean if a field has been set.

### GetRequiresHumanFor

`func (o *PolicyUpdateRequest) GetRequiresHumanFor() map[string]interface{}`

GetRequiresHumanFor returns the RequiresHumanFor field if non-nil, zero value otherwise.

### GetRequiresHumanForOk

`func (o *PolicyUpdateRequest) GetRequiresHumanForOk() (*map[string]interface{}, bool)`

GetRequiresHumanForOk returns a tuple with the RequiresHumanFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresHumanFor

`func (o *PolicyUpdateRequest) SetRequiresHumanFor(v map[string]interface{})`

SetRequiresHumanFor sets RequiresHumanFor field to given value.

### HasRequiresHumanFor

`func (o *PolicyUpdateRequest) HasRequiresHumanFor() bool`

HasRequiresHumanFor returns a boolean if a field has been set.

### GetPolicyVersion

`func (o *PolicyUpdateRequest) GetPolicyVersion() string`

GetPolicyVersion returns the PolicyVersion field if non-nil, zero value otherwise.

### GetPolicyVersionOk

`func (o *PolicyUpdateRequest) GetPolicyVersionOk() (*string, bool)`

GetPolicyVersionOk returns a tuple with the PolicyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyVersion

`func (o *PolicyUpdateRequest) SetPolicyVersion(v string)`

SetPolicyVersion sets PolicyVersion field to given value.

### HasPolicyVersion

`func (o *PolicyUpdateRequest) HasPolicyVersion() bool`

HasPolicyVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


