# RedlineSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ByVerdict** | **map[string]int32** |  | 
**Total** | **int32** |  | 

## Methods

### NewRedlineSummary

`func NewRedlineSummary(byVerdict map[string]int32, total int32, ) *RedlineSummary`

NewRedlineSummary instantiates a new RedlineSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedlineSummaryWithDefaults

`func NewRedlineSummaryWithDefaults() *RedlineSummary`

NewRedlineSummaryWithDefaults instantiates a new RedlineSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetByVerdict

`func (o *RedlineSummary) GetByVerdict() map[string]int32`

GetByVerdict returns the ByVerdict field if non-nil, zero value otherwise.

### GetByVerdictOk

`func (o *RedlineSummary) GetByVerdictOk() (*map[string]int32, bool)`

GetByVerdictOk returns a tuple with the ByVerdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByVerdict

`func (o *RedlineSummary) SetByVerdict(v map[string]int32)`

SetByVerdict sets ByVerdict field to given value.


### GetTotal

`func (o *RedlineSummary) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *RedlineSummary) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *RedlineSummary) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


