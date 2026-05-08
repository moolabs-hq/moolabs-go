# GrantBurnDownHistorySegment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Period** | [**Period**](Period.md) | The period of the segment. | 
**Usage** | **float64** | The total usage of the grant in the period. | [readonly] 
**Overage** | **float64** | Overuse that wasn&#39;t covered by grants. | [readonly] 
**BalanceAtStart** | **float64** | entitlement balance at the start of the period. | [readonly] 
**GrantBalancesAtStart** | **map[string]float64** | The balance breakdown of each active grant at the start of the period: GrantID: Balance | [readonly] 
**BalanceAtEnd** | **float64** | The entitlement balance at the end of the period. | [readonly] 
**GrantBalancesAtEnd** | **map[string]float64** | The balance breakdown of each active grant at the end of the period: GrantID: Balance | [readonly] 
**GrantUsages** | [**[]GrantUsageRecord**](GrantUsageRecord.md) | Which grants were actually burnt down in the period and by what amount. | 

## Methods

### NewGrantBurnDownHistorySegment

`func NewGrantBurnDownHistorySegment(period Period, usage float64, overage float64, balanceAtStart float64, grantBalancesAtStart map[string]float64, balanceAtEnd float64, grantBalancesAtEnd map[string]float64, grantUsages []GrantUsageRecord, ) *GrantBurnDownHistorySegment`

NewGrantBurnDownHistorySegment instantiates a new GrantBurnDownHistorySegment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantBurnDownHistorySegmentWithDefaults

`func NewGrantBurnDownHistorySegmentWithDefaults() *GrantBurnDownHistorySegment`

NewGrantBurnDownHistorySegmentWithDefaults instantiates a new GrantBurnDownHistorySegment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriod

`func (o *GrantBurnDownHistorySegment) GetPeriod() Period`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *GrantBurnDownHistorySegment) GetPeriodOk() (*Period, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *GrantBurnDownHistorySegment) SetPeriod(v Period)`

SetPeriod sets Period field to given value.


### GetUsage

`func (o *GrantBurnDownHistorySegment) GetUsage() float64`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *GrantBurnDownHistorySegment) GetUsageOk() (*float64, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *GrantBurnDownHistorySegment) SetUsage(v float64)`

SetUsage sets Usage field to given value.


### GetOverage

`func (o *GrantBurnDownHistorySegment) GetOverage() float64`

GetOverage returns the Overage field if non-nil, zero value otherwise.

### GetOverageOk

`func (o *GrantBurnDownHistorySegment) GetOverageOk() (*float64, bool)`

GetOverageOk returns a tuple with the Overage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverage

`func (o *GrantBurnDownHistorySegment) SetOverage(v float64)`

SetOverage sets Overage field to given value.


### GetBalanceAtStart

`func (o *GrantBurnDownHistorySegment) GetBalanceAtStart() float64`

GetBalanceAtStart returns the BalanceAtStart field if non-nil, zero value otherwise.

### GetBalanceAtStartOk

`func (o *GrantBurnDownHistorySegment) GetBalanceAtStartOk() (*float64, bool)`

GetBalanceAtStartOk returns a tuple with the BalanceAtStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAtStart

`func (o *GrantBurnDownHistorySegment) SetBalanceAtStart(v float64)`

SetBalanceAtStart sets BalanceAtStart field to given value.


### GetGrantBalancesAtStart

`func (o *GrantBurnDownHistorySegment) GetGrantBalancesAtStart() map[string]float64`

GetGrantBalancesAtStart returns the GrantBalancesAtStart field if non-nil, zero value otherwise.

### GetGrantBalancesAtStartOk

`func (o *GrantBurnDownHistorySegment) GetGrantBalancesAtStartOk() (*map[string]float64, bool)`

GetGrantBalancesAtStartOk returns a tuple with the GrantBalancesAtStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantBalancesAtStart

`func (o *GrantBurnDownHistorySegment) SetGrantBalancesAtStart(v map[string]float64)`

SetGrantBalancesAtStart sets GrantBalancesAtStart field to given value.


### GetBalanceAtEnd

`func (o *GrantBurnDownHistorySegment) GetBalanceAtEnd() float64`

GetBalanceAtEnd returns the BalanceAtEnd field if non-nil, zero value otherwise.

### GetBalanceAtEndOk

`func (o *GrantBurnDownHistorySegment) GetBalanceAtEndOk() (*float64, bool)`

GetBalanceAtEndOk returns a tuple with the BalanceAtEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAtEnd

`func (o *GrantBurnDownHistorySegment) SetBalanceAtEnd(v float64)`

SetBalanceAtEnd sets BalanceAtEnd field to given value.


### GetGrantBalancesAtEnd

`func (o *GrantBurnDownHistorySegment) GetGrantBalancesAtEnd() map[string]float64`

GetGrantBalancesAtEnd returns the GrantBalancesAtEnd field if non-nil, zero value otherwise.

### GetGrantBalancesAtEndOk

`func (o *GrantBurnDownHistorySegment) GetGrantBalancesAtEndOk() (*map[string]float64, bool)`

GetGrantBalancesAtEndOk returns a tuple with the GrantBalancesAtEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantBalancesAtEnd

`func (o *GrantBurnDownHistorySegment) SetGrantBalancesAtEnd(v map[string]float64)`

SetGrantBalancesAtEnd sets GrantBalancesAtEnd field to given value.


### GetGrantUsages

`func (o *GrantBurnDownHistorySegment) GetGrantUsages() []GrantUsageRecord`

GetGrantUsages returns the GrantUsages field if non-nil, zero value otherwise.

### GetGrantUsagesOk

`func (o *GrantBurnDownHistorySegment) GetGrantUsagesOk() (*[]GrantUsageRecord, bool)`

GetGrantUsagesOk returns a tuple with the GrantUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantUsages

`func (o *GrantBurnDownHistorySegment) SetGrantUsages(v []GrantUsageRecord)`

SetGrantUsages sets GrantUsages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


