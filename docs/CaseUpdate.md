# CaseUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RiskTier** | Pointer to **string** |  | [optional] 
**AssignedAgent** | Pointer to **string** |  | [optional] 
**AssignedHuman** | Pointer to **string** |  | [optional] 
**CaseOwnerTeam** | Pointer to **string** |  | [optional] 
**AutonomyLevel** | Pointer to **int32** | Per-case autonomy override (1&#x3D;inform-only, 5&#x3D;fully autonomous). Capped by the orchestrator agent policy&#39;s max_autonomy. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Status** | Pointer to **string** | Target case status. Goes through the FSM in case_state.transition_case so an invalid transition (e.g. closed → open without a reason) returns 422. | [optional] 
**StatusReason** | Pointer to **string** |  | [optional] 

## Methods

### NewCaseUpdate

`func NewCaseUpdate() *CaseUpdate`

NewCaseUpdate instantiates a new CaseUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaseUpdateWithDefaults

`func NewCaseUpdateWithDefaults() *CaseUpdate`

NewCaseUpdateWithDefaults instantiates a new CaseUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRiskTier

`func (o *CaseUpdate) GetRiskTier() string`

GetRiskTier returns the RiskTier field if non-nil, zero value otherwise.

### GetRiskTierOk

`func (o *CaseUpdate) GetRiskTierOk() (*string, bool)`

GetRiskTierOk returns a tuple with the RiskTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskTier

`func (o *CaseUpdate) SetRiskTier(v string)`

SetRiskTier sets RiskTier field to given value.

### HasRiskTier

`func (o *CaseUpdate) HasRiskTier() bool`

HasRiskTier returns a boolean if a field has been set.

### GetAssignedAgent

`func (o *CaseUpdate) GetAssignedAgent() string`

GetAssignedAgent returns the AssignedAgent field if non-nil, zero value otherwise.

### GetAssignedAgentOk

`func (o *CaseUpdate) GetAssignedAgentOk() (*string, bool)`

GetAssignedAgentOk returns a tuple with the AssignedAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedAgent

`func (o *CaseUpdate) SetAssignedAgent(v string)`

SetAssignedAgent sets AssignedAgent field to given value.

### HasAssignedAgent

`func (o *CaseUpdate) HasAssignedAgent() bool`

HasAssignedAgent returns a boolean if a field has been set.

### GetAssignedHuman

`func (o *CaseUpdate) GetAssignedHuman() string`

GetAssignedHuman returns the AssignedHuman field if non-nil, zero value otherwise.

### GetAssignedHumanOk

`func (o *CaseUpdate) GetAssignedHumanOk() (*string, bool)`

GetAssignedHumanOk returns a tuple with the AssignedHuman field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedHuman

`func (o *CaseUpdate) SetAssignedHuman(v string)`

SetAssignedHuman sets AssignedHuman field to given value.

### HasAssignedHuman

`func (o *CaseUpdate) HasAssignedHuman() bool`

HasAssignedHuman returns a boolean if a field has been set.

### GetCaseOwnerTeam

`func (o *CaseUpdate) GetCaseOwnerTeam() string`

GetCaseOwnerTeam returns the CaseOwnerTeam field if non-nil, zero value otherwise.

### GetCaseOwnerTeamOk

`func (o *CaseUpdate) GetCaseOwnerTeamOk() (*string, bool)`

GetCaseOwnerTeamOk returns a tuple with the CaseOwnerTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseOwnerTeam

`func (o *CaseUpdate) SetCaseOwnerTeam(v string)`

SetCaseOwnerTeam sets CaseOwnerTeam field to given value.

### HasCaseOwnerTeam

`func (o *CaseUpdate) HasCaseOwnerTeam() bool`

HasCaseOwnerTeam returns a boolean if a field has been set.

### GetAutonomyLevel

`func (o *CaseUpdate) GetAutonomyLevel() int32`

GetAutonomyLevel returns the AutonomyLevel field if non-nil, zero value otherwise.

### GetAutonomyLevelOk

`func (o *CaseUpdate) GetAutonomyLevelOk() (*int32, bool)`

GetAutonomyLevelOk returns a tuple with the AutonomyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutonomyLevel

`func (o *CaseUpdate) SetAutonomyLevel(v int32)`

SetAutonomyLevel sets AutonomyLevel field to given value.

### HasAutonomyLevel

`func (o *CaseUpdate) HasAutonomyLevel() bool`

HasAutonomyLevel returns a boolean if a field has been set.

### GetMetadata

`func (o *CaseUpdate) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CaseUpdate) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CaseUpdate) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CaseUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetStatus

`func (o *CaseUpdate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CaseUpdate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CaseUpdate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CaseUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusReason

`func (o *CaseUpdate) GetStatusReason() string`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *CaseUpdate) GetStatusReasonOk() (*string, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *CaseUpdate) SetStatusReason(v string)`

SetStatusReason sets StatusReason field to given value.

### HasStatusReason

`func (o *CaseUpdate) HasStatusReason() bool`

HasStatusReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


