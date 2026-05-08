# MeterQueryResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **time.Time** | The start of the period the usage is queried from. If not specified, the usage is queried from the beginning of time. | [optional] 
**To** | Pointer to **time.Time** | The end of the period the usage is queried to. If not specified, the usage is queried up to the current time. | [optional] 
**WindowSize** | Pointer to [**WindowSize**](WindowSize.md) | The window size that the usage is aggregated. If not specified, the usage is aggregated over the entire period. | [optional] 
**Data** | [**[]MeterQueryRow**](MeterQueryRow.md) | The usage data. If no data is available, an empty array is returned. | 

## Methods

### NewMeterQueryResult

`func NewMeterQueryResult(data []MeterQueryRow, ) *MeterQueryResult`

NewMeterQueryResult instantiates a new MeterQueryResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeterQueryResultWithDefaults

`func NewMeterQueryResultWithDefaults() *MeterQueryResult`

NewMeterQueryResultWithDefaults instantiates a new MeterQueryResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *MeterQueryResult) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *MeterQueryResult) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *MeterQueryResult) SetFrom(v time.Time)`

SetFrom sets From field to given value.

### HasFrom

`func (o *MeterQueryResult) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *MeterQueryResult) GetTo() time.Time`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *MeterQueryResult) GetToOk() (*time.Time, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *MeterQueryResult) SetTo(v time.Time)`

SetTo sets To field to given value.

### HasTo

`func (o *MeterQueryResult) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetWindowSize

`func (o *MeterQueryResult) GetWindowSize() WindowSize`

GetWindowSize returns the WindowSize field if non-nil, zero value otherwise.

### GetWindowSizeOk

`func (o *MeterQueryResult) GetWindowSizeOk() (*WindowSize, bool)`

GetWindowSizeOk returns a tuple with the WindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSize

`func (o *MeterQueryResult) SetWindowSize(v WindowSize)`

SetWindowSize sets WindowSize field to given value.

### HasWindowSize

`func (o *MeterQueryResult) HasWindowSize() bool`

HasWindowSize returns a boolean if a field has been set.

### GetData

`func (o *MeterQueryResult) GetData() []MeterQueryRow`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MeterQueryResult) GetDataOk() (*[]MeterQueryRow, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MeterQueryResult) SetData(v []MeterQueryRow)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


