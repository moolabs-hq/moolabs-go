# ShadowRunSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShadowRunId** | **string** |  | 
**RowsProcessed** | **int32** |  | 
**RowsAttributed** | **int32** |  | 
**CoveragePct** | **float32** |  | 
**AlgorithmsUsed** | **map[string]int32** |  | 

## Methods

### NewShadowRunSummary

`func NewShadowRunSummary(shadowRunId string, rowsProcessed int32, rowsAttributed int32, coveragePct float32, algorithmsUsed map[string]int32, ) *ShadowRunSummary`

NewShadowRunSummary instantiates a new ShadowRunSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShadowRunSummaryWithDefaults

`func NewShadowRunSummaryWithDefaults() *ShadowRunSummary`

NewShadowRunSummaryWithDefaults instantiates a new ShadowRunSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShadowRunId

`func (o *ShadowRunSummary) GetShadowRunId() string`

GetShadowRunId returns the ShadowRunId field if non-nil, zero value otherwise.

### GetShadowRunIdOk

`func (o *ShadowRunSummary) GetShadowRunIdOk() (*string, bool)`

GetShadowRunIdOk returns a tuple with the ShadowRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShadowRunId

`func (o *ShadowRunSummary) SetShadowRunId(v string)`

SetShadowRunId sets ShadowRunId field to given value.


### GetRowsProcessed

`func (o *ShadowRunSummary) GetRowsProcessed() int32`

GetRowsProcessed returns the RowsProcessed field if non-nil, zero value otherwise.

### GetRowsProcessedOk

`func (o *ShadowRunSummary) GetRowsProcessedOk() (*int32, bool)`

GetRowsProcessedOk returns a tuple with the RowsProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowsProcessed

`func (o *ShadowRunSummary) SetRowsProcessed(v int32)`

SetRowsProcessed sets RowsProcessed field to given value.


### GetRowsAttributed

`func (o *ShadowRunSummary) GetRowsAttributed() int32`

GetRowsAttributed returns the RowsAttributed field if non-nil, zero value otherwise.

### GetRowsAttributedOk

`func (o *ShadowRunSummary) GetRowsAttributedOk() (*int32, bool)`

GetRowsAttributedOk returns a tuple with the RowsAttributed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowsAttributed

`func (o *ShadowRunSummary) SetRowsAttributed(v int32)`

SetRowsAttributed sets RowsAttributed field to given value.


### GetCoveragePct

`func (o *ShadowRunSummary) GetCoveragePct() float32`

GetCoveragePct returns the CoveragePct field if non-nil, zero value otherwise.

### GetCoveragePctOk

`func (o *ShadowRunSummary) GetCoveragePctOk() (*float32, bool)`

GetCoveragePctOk returns a tuple with the CoveragePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoveragePct

`func (o *ShadowRunSummary) SetCoveragePct(v float32)`

SetCoveragePct sets CoveragePct field to given value.


### GetAlgorithmsUsed

`func (o *ShadowRunSummary) GetAlgorithmsUsed() map[string]int32`

GetAlgorithmsUsed returns the AlgorithmsUsed field if non-nil, zero value otherwise.

### GetAlgorithmsUsedOk

`func (o *ShadowRunSummary) GetAlgorithmsUsedOk() (*map[string]int32, bool)`

GetAlgorithmsUsedOk returns a tuple with the AlgorithmsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithmsUsed

`func (o *ShadowRunSummary) SetAlgorithmsUsed(v map[string]int32)`

SetAlgorithmsUsed sets AlgorithmsUsed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


