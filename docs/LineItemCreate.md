# LineItemCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricType** | **string** |  | 
**Quantity** | [**Quantity**](Quantity.md) |  | 
**Unit** | **string** |  | 
**RateSnapshotId** | Pointer to **string** |  | [optional] 
**UnitCost** | [**UnitCost**](UnitCost.md) |  | 
**LineTotal** | [**LineTotal**](LineTotal.md) |  | 
**ReportingLineTotal** | [**ReportingLineTotal**](ReportingLineTotal.md) |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewLineItemCreate

`func NewLineItemCreate(metricType string, quantity Quantity, unit string, unitCost UnitCost, lineTotal LineTotal, reportingLineTotal ReportingLineTotal, ) *LineItemCreate`

NewLineItemCreate instantiates a new LineItemCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineItemCreateWithDefaults

`func NewLineItemCreateWithDefaults() *LineItemCreate`

NewLineItemCreateWithDefaults instantiates a new LineItemCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricType

`func (o *LineItemCreate) GetMetricType() string`

GetMetricType returns the MetricType field if non-nil, zero value otherwise.

### GetMetricTypeOk

`func (o *LineItemCreate) GetMetricTypeOk() (*string, bool)`

GetMetricTypeOk returns a tuple with the MetricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricType

`func (o *LineItemCreate) SetMetricType(v string)`

SetMetricType sets MetricType field to given value.


### GetQuantity

`func (o *LineItemCreate) GetQuantity() Quantity`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *LineItemCreate) GetQuantityOk() (*Quantity, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *LineItemCreate) SetQuantity(v Quantity)`

SetQuantity sets Quantity field to given value.


### GetUnit

`func (o *LineItemCreate) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *LineItemCreate) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *LineItemCreate) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetRateSnapshotId

`func (o *LineItemCreate) GetRateSnapshotId() string`

GetRateSnapshotId returns the RateSnapshotId field if non-nil, zero value otherwise.

### GetRateSnapshotIdOk

`func (o *LineItemCreate) GetRateSnapshotIdOk() (*string, bool)`

GetRateSnapshotIdOk returns a tuple with the RateSnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateSnapshotId

`func (o *LineItemCreate) SetRateSnapshotId(v string)`

SetRateSnapshotId sets RateSnapshotId field to given value.

### HasRateSnapshotId

`func (o *LineItemCreate) HasRateSnapshotId() bool`

HasRateSnapshotId returns a boolean if a field has been set.

### GetUnitCost

`func (o *LineItemCreate) GetUnitCost() UnitCost`

GetUnitCost returns the UnitCost field if non-nil, zero value otherwise.

### GetUnitCostOk

`func (o *LineItemCreate) GetUnitCostOk() (*UnitCost, bool)`

GetUnitCostOk returns a tuple with the UnitCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitCost

`func (o *LineItemCreate) SetUnitCost(v UnitCost)`

SetUnitCost sets UnitCost field to given value.


### GetLineTotal

`func (o *LineItemCreate) GetLineTotal() LineTotal`

GetLineTotal returns the LineTotal field if non-nil, zero value otherwise.

### GetLineTotalOk

`func (o *LineItemCreate) GetLineTotalOk() (*LineTotal, bool)`

GetLineTotalOk returns a tuple with the LineTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineTotal

`func (o *LineItemCreate) SetLineTotal(v LineTotal)`

SetLineTotal sets LineTotal field to given value.


### GetReportingLineTotal

`func (o *LineItemCreate) GetReportingLineTotal() ReportingLineTotal`

GetReportingLineTotal returns the ReportingLineTotal field if non-nil, zero value otherwise.

### GetReportingLineTotalOk

`func (o *LineItemCreate) GetReportingLineTotalOk() (*ReportingLineTotal, bool)`

GetReportingLineTotalOk returns a tuple with the ReportingLineTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingLineTotal

`func (o *LineItemCreate) SetReportingLineTotal(v ReportingLineTotal)`

SetReportingLineTotal sets ReportingLineTotal field to given value.


### GetMetadata

`func (o *LineItemCreate) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *LineItemCreate) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *LineItemCreate) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *LineItemCreate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


