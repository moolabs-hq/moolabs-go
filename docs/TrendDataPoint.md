# TrendDataPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** |  | 
**TotalCost** | **float32** |  | 
**RequestCount** | **int32** |  | 

## Methods

### NewTrendDataPoint

`func NewTrendDataPoint(date string, totalCost float32, requestCount int32, ) *TrendDataPoint`

NewTrendDataPoint instantiates a new TrendDataPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrendDataPointWithDefaults

`func NewTrendDataPointWithDefaults() *TrendDataPoint`

NewTrendDataPointWithDefaults instantiates a new TrendDataPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *TrendDataPoint) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *TrendDataPoint) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *TrendDataPoint) SetDate(v string)`

SetDate sets Date field to given value.


### GetTotalCost

`func (o *TrendDataPoint) GetTotalCost() float32`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *TrendDataPoint) GetTotalCostOk() (*float32, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *TrendDataPoint) SetTotalCost(v float32)`

SetTotalCost sets TotalCost field to given value.


### GetRequestCount

`func (o *TrendDataPoint) GetRequestCount() int32`

GetRequestCount returns the RequestCount field if non-nil, zero value otherwise.

### GetRequestCountOk

`func (o *TrendDataPoint) GetRequestCountOk() (*int32, bool)`

GetRequestCountOk returns a tuple with the RequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCount

`func (o *TrendDataPoint) SetRequestCount(v int32)`

SetRequestCount sets RequestCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


