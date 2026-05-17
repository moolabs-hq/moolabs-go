# ShadowRunMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShadowRunId** | **string** |  | 
**PeriodStart** | **string** |  | 
**PeriodEnd** | **string** |  | 
**Wape** | **float32** |  | 
**CoveragePct** | **float32** |  | 
**TotalImportRows** | **int32** |  | 
**AttributedRows** | **int32** |  | 
**AlgorithmBreakdown** | [**map[string]AlgorithmBreakdownItem**](AlgorithmBreakdownItem.md) |  | 

## Methods

### NewShadowRunMetrics

`func NewShadowRunMetrics(shadowRunId string, periodStart string, periodEnd string, wape float32, coveragePct float32, totalImportRows int32, attributedRows int32, algorithmBreakdown map[string]AlgorithmBreakdownItem, ) *ShadowRunMetrics`

NewShadowRunMetrics instantiates a new ShadowRunMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShadowRunMetricsWithDefaults

`func NewShadowRunMetricsWithDefaults() *ShadowRunMetrics`

NewShadowRunMetricsWithDefaults instantiates a new ShadowRunMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShadowRunId

`func (o *ShadowRunMetrics) GetShadowRunId() string`

GetShadowRunId returns the ShadowRunId field if non-nil, zero value otherwise.

### GetShadowRunIdOk

`func (o *ShadowRunMetrics) GetShadowRunIdOk() (*string, bool)`

GetShadowRunIdOk returns a tuple with the ShadowRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShadowRunId

`func (o *ShadowRunMetrics) SetShadowRunId(v string)`

SetShadowRunId sets ShadowRunId field to given value.


### GetPeriodStart

`func (o *ShadowRunMetrics) GetPeriodStart() string`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *ShadowRunMetrics) GetPeriodStartOk() (*string, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *ShadowRunMetrics) SetPeriodStart(v string)`

SetPeriodStart sets PeriodStart field to given value.


### GetPeriodEnd

`func (o *ShadowRunMetrics) GetPeriodEnd() string`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *ShadowRunMetrics) GetPeriodEndOk() (*string, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *ShadowRunMetrics) SetPeriodEnd(v string)`

SetPeriodEnd sets PeriodEnd field to given value.


### GetWape

`func (o *ShadowRunMetrics) GetWape() float32`

GetWape returns the Wape field if non-nil, zero value otherwise.

### GetWapeOk

`func (o *ShadowRunMetrics) GetWapeOk() (*float32, bool)`

GetWapeOk returns a tuple with the Wape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWape

`func (o *ShadowRunMetrics) SetWape(v float32)`

SetWape sets Wape field to given value.


### GetCoveragePct

`func (o *ShadowRunMetrics) GetCoveragePct() float32`

GetCoveragePct returns the CoveragePct field if non-nil, zero value otherwise.

### GetCoveragePctOk

`func (o *ShadowRunMetrics) GetCoveragePctOk() (*float32, bool)`

GetCoveragePctOk returns a tuple with the CoveragePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoveragePct

`func (o *ShadowRunMetrics) SetCoveragePct(v float32)`

SetCoveragePct sets CoveragePct field to given value.


### GetTotalImportRows

`func (o *ShadowRunMetrics) GetTotalImportRows() int32`

GetTotalImportRows returns the TotalImportRows field if non-nil, zero value otherwise.

### GetTotalImportRowsOk

`func (o *ShadowRunMetrics) GetTotalImportRowsOk() (*int32, bool)`

GetTotalImportRowsOk returns a tuple with the TotalImportRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalImportRows

`func (o *ShadowRunMetrics) SetTotalImportRows(v int32)`

SetTotalImportRows sets TotalImportRows field to given value.


### GetAttributedRows

`func (o *ShadowRunMetrics) GetAttributedRows() int32`

GetAttributedRows returns the AttributedRows field if non-nil, zero value otherwise.

### GetAttributedRowsOk

`func (o *ShadowRunMetrics) GetAttributedRowsOk() (*int32, bool)`

GetAttributedRowsOk returns a tuple with the AttributedRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributedRows

`func (o *ShadowRunMetrics) SetAttributedRows(v int32)`

SetAttributedRows sets AttributedRows field to given value.


### GetAlgorithmBreakdown

`func (o *ShadowRunMetrics) GetAlgorithmBreakdown() map[string]AlgorithmBreakdownItem`

GetAlgorithmBreakdown returns the AlgorithmBreakdown field if non-nil, zero value otherwise.

### GetAlgorithmBreakdownOk

`func (o *ShadowRunMetrics) GetAlgorithmBreakdownOk() (*map[string]AlgorithmBreakdownItem, bool)`

GetAlgorithmBreakdownOk returns a tuple with the AlgorithmBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithmBreakdown

`func (o *ShadowRunMetrics) SetAlgorithmBreakdown(v map[string]AlgorithmBreakdownItem)`

SetAlgorithmBreakdown sets AlgorithmBreakdown field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


