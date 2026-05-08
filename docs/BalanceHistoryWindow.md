# BalanceHistoryWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Period** | [**Period**](Period.md) |  | 
**Usage** | **float64** | The total usage of the feature in the period. | [readonly] 
**BalanceAtStart** | **float64** | The entitlement balance at the start of the period. | [readonly] 

## Methods

### NewBalanceHistoryWindow

`func NewBalanceHistoryWindow(period Period, usage float64, balanceAtStart float64, ) *BalanceHistoryWindow`

NewBalanceHistoryWindow instantiates a new BalanceHistoryWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceHistoryWindowWithDefaults

`func NewBalanceHistoryWindowWithDefaults() *BalanceHistoryWindow`

NewBalanceHistoryWindowWithDefaults instantiates a new BalanceHistoryWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriod

`func (o *BalanceHistoryWindow) GetPeriod() Period`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *BalanceHistoryWindow) GetPeriodOk() (*Period, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *BalanceHistoryWindow) SetPeriod(v Period)`

SetPeriod sets Period field to given value.


### GetUsage

`func (o *BalanceHistoryWindow) GetUsage() float64`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *BalanceHistoryWindow) GetUsageOk() (*float64, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *BalanceHistoryWindow) SetUsage(v float64)`

SetUsage sets Usage field to given value.


### GetBalanceAtStart

`func (o *BalanceHistoryWindow) GetBalanceAtStart() float64`

GetBalanceAtStart returns the BalanceAtStart field if non-nil, zero value otherwise.

### GetBalanceAtStartOk

`func (o *BalanceHistoryWindow) GetBalanceAtStartOk() (*float64, bool)`

GetBalanceAtStartOk returns a tuple with the BalanceAtStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAtStart

`func (o *BalanceHistoryWindow) SetBalanceAtStart(v float64)`

SetBalanceAtStart sets BalanceAtStart field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


