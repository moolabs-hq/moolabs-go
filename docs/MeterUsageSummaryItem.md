# MeterUsageSummaryItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MeterSlug** | **string** | The meter slug. | 
**Description** | Pointer to **string** | Human-readable meter description. | [optional] 
**Aggregation** | **string** | The meter&#39;s aggregation type (SUM, COUNT, AVG, etc). | 
**Value** | **float64** | Total aggregated value for the requested time range. | 
**Windows** | Pointer to [**[]MeterQueryRow**](MeterQueryRow.md) | Per-window breakdown (only present when windowSize is specified). | [optional] 

## Methods

### NewMeterUsageSummaryItem

`func NewMeterUsageSummaryItem(meterSlug string, aggregation string, value float64, ) *MeterUsageSummaryItem`

NewMeterUsageSummaryItem instantiates a new MeterUsageSummaryItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeterUsageSummaryItemWithDefaults

`func NewMeterUsageSummaryItemWithDefaults() *MeterUsageSummaryItem`

NewMeterUsageSummaryItemWithDefaults instantiates a new MeterUsageSummaryItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeterSlug

`func (o *MeterUsageSummaryItem) GetMeterSlug() string`

GetMeterSlug returns the MeterSlug field if non-nil, zero value otherwise.

### GetMeterSlugOk

`func (o *MeterUsageSummaryItem) GetMeterSlugOk() (*string, bool)`

GetMeterSlugOk returns a tuple with the MeterSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterSlug

`func (o *MeterUsageSummaryItem) SetMeterSlug(v string)`

SetMeterSlug sets MeterSlug field to given value.


### GetDescription

`func (o *MeterUsageSummaryItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MeterUsageSummaryItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MeterUsageSummaryItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MeterUsageSummaryItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAggregation

`func (o *MeterUsageSummaryItem) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *MeterUsageSummaryItem) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *MeterUsageSummaryItem) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.


### GetValue

`func (o *MeterUsageSummaryItem) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MeterUsageSummaryItem) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MeterUsageSummaryItem) SetValue(v float64)`

SetValue sets Value field to given value.


### GetWindows

`func (o *MeterUsageSummaryItem) GetWindows() []MeterQueryRow`

GetWindows returns the Windows field if non-nil, zero value otherwise.

### GetWindowsOk

`func (o *MeterUsageSummaryItem) GetWindowsOk() (*[]MeterQueryRow, bool)`

GetWindowsOk returns a tuple with the Windows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindows

`func (o *MeterUsageSummaryItem) SetWindows(v []MeterQueryRow)`

SetWindows sets Windows field to given value.

### HasWindows

`func (o *MeterUsageSummaryItem) HasWindows() bool`

HasWindows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


