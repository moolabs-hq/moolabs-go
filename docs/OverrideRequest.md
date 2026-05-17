# OverrideRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CorrectedValues** | **map[string]interface{}** |  | 
**CompletionNote** | Pointer to **string** |  | [optional] 

## Methods

### NewOverrideRequest

`func NewOverrideRequest(correctedValues map[string]interface{}, ) *OverrideRequest`

NewOverrideRequest instantiates a new OverrideRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverrideRequestWithDefaults

`func NewOverrideRequestWithDefaults() *OverrideRequest`

NewOverrideRequestWithDefaults instantiates a new OverrideRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorrectedValues

`func (o *OverrideRequest) GetCorrectedValues() map[string]interface{}`

GetCorrectedValues returns the CorrectedValues field if non-nil, zero value otherwise.

### GetCorrectedValuesOk

`func (o *OverrideRequest) GetCorrectedValuesOk() (*map[string]interface{}, bool)`

GetCorrectedValuesOk returns a tuple with the CorrectedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrectedValues

`func (o *OverrideRequest) SetCorrectedValues(v map[string]interface{})`

SetCorrectedValues sets CorrectedValues field to given value.


### GetCompletionNote

`func (o *OverrideRequest) GetCompletionNote() string`

GetCompletionNote returns the CompletionNote field if non-nil, zero value otherwise.

### GetCompletionNoteOk

`func (o *OverrideRequest) GetCompletionNoteOk() (*string, bool)`

GetCompletionNoteOk returns a tuple with the CompletionNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionNote

`func (o *OverrideRequest) SetCompletionNote(v string)`

SetCompletionNote sets CompletionNote field to given value.

### HasCompletionNote

`func (o *OverrideRequest) HasCompletionNote() bool`

HasCompletionNote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


