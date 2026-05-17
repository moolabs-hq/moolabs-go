# DryRunResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DerivedFields** | **map[string]string** |  | 
**MatchGrade** | **string** |  | 
**RulesApplied** | **[]string** |  | 
**RulesSkipped** | **[]interface{}** |  | 
**Tier3Enabled** | **bool** |  | 
**Note** | **string** |  | 

## Methods

### NewDryRunResponse

`func NewDryRunResponse(derivedFields map[string]string, matchGrade string, rulesApplied []string, rulesSkipped []interface{}, tier3Enabled bool, note string, ) *DryRunResponse`

NewDryRunResponse instantiates a new DryRunResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDryRunResponseWithDefaults

`func NewDryRunResponseWithDefaults() *DryRunResponse`

NewDryRunResponseWithDefaults instantiates a new DryRunResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDerivedFields

`func (o *DryRunResponse) GetDerivedFields() map[string]string`

GetDerivedFields returns the DerivedFields field if non-nil, zero value otherwise.

### GetDerivedFieldsOk

`func (o *DryRunResponse) GetDerivedFieldsOk() (*map[string]string, bool)`

GetDerivedFieldsOk returns a tuple with the DerivedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedFields

`func (o *DryRunResponse) SetDerivedFields(v map[string]string)`

SetDerivedFields sets DerivedFields field to given value.


### GetMatchGrade

`func (o *DryRunResponse) GetMatchGrade() string`

GetMatchGrade returns the MatchGrade field if non-nil, zero value otherwise.

### GetMatchGradeOk

`func (o *DryRunResponse) GetMatchGradeOk() (*string, bool)`

GetMatchGradeOk returns a tuple with the MatchGrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchGrade

`func (o *DryRunResponse) SetMatchGrade(v string)`

SetMatchGrade sets MatchGrade field to given value.


### GetRulesApplied

`func (o *DryRunResponse) GetRulesApplied() []string`

GetRulesApplied returns the RulesApplied field if non-nil, zero value otherwise.

### GetRulesAppliedOk

`func (o *DryRunResponse) GetRulesAppliedOk() (*[]string, bool)`

GetRulesAppliedOk returns a tuple with the RulesApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesApplied

`func (o *DryRunResponse) SetRulesApplied(v []string)`

SetRulesApplied sets RulesApplied field to given value.


### GetRulesSkipped

`func (o *DryRunResponse) GetRulesSkipped() []interface{}`

GetRulesSkipped returns the RulesSkipped field if non-nil, zero value otherwise.

### GetRulesSkippedOk

`func (o *DryRunResponse) GetRulesSkippedOk() (*[]interface{}, bool)`

GetRulesSkippedOk returns a tuple with the RulesSkipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesSkipped

`func (o *DryRunResponse) SetRulesSkipped(v []interface{})`

SetRulesSkipped sets RulesSkipped field to given value.


### GetTier3Enabled

`func (o *DryRunResponse) GetTier3Enabled() bool`

GetTier3Enabled returns the Tier3Enabled field if non-nil, zero value otherwise.

### GetTier3EnabledOk

`func (o *DryRunResponse) GetTier3EnabledOk() (*bool, bool)`

GetTier3EnabledOk returns a tuple with the Tier3Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier3Enabled

`func (o *DryRunResponse) SetTier3Enabled(v bool)`

SetTier3Enabled sets Tier3Enabled field to given value.


### GetNote

`func (o *DryRunResponse) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *DryRunResponse) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *DryRunResponse) SetNote(v string)`

SetNote sets Note field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


