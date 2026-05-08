# WindowedBalanceHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WindowedHistory** | [**[]BalanceHistoryWindow**](BalanceHistoryWindow.md) | The windowed balance history. - It only returns rows for windows where there was usage. - The windows are inclusive at their start and exclusive at their end. - The last window may be smaller than the window size and is inclusive at both ends. | 
**BurndownHistory** | [**[]GrantBurnDownHistorySegment**](GrantBurnDownHistorySegment.md) | Grant burndown history. | 

## Methods

### NewWindowedBalanceHistory

`func NewWindowedBalanceHistory(windowedHistory []BalanceHistoryWindow, burndownHistory []GrantBurnDownHistorySegment, ) *WindowedBalanceHistory`

NewWindowedBalanceHistory instantiates a new WindowedBalanceHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWindowedBalanceHistoryWithDefaults

`func NewWindowedBalanceHistoryWithDefaults() *WindowedBalanceHistory`

NewWindowedBalanceHistoryWithDefaults instantiates a new WindowedBalanceHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWindowedHistory

`func (o *WindowedBalanceHistory) GetWindowedHistory() []BalanceHistoryWindow`

GetWindowedHistory returns the WindowedHistory field if non-nil, zero value otherwise.

### GetWindowedHistoryOk

`func (o *WindowedBalanceHistory) GetWindowedHistoryOk() (*[]BalanceHistoryWindow, bool)`

GetWindowedHistoryOk returns a tuple with the WindowedHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowedHistory

`func (o *WindowedBalanceHistory) SetWindowedHistory(v []BalanceHistoryWindow)`

SetWindowedHistory sets WindowedHistory field to given value.


### GetBurndownHistory

`func (o *WindowedBalanceHistory) GetBurndownHistory() []GrantBurnDownHistorySegment`

GetBurndownHistory returns the BurndownHistory field if non-nil, zero value otherwise.

### GetBurndownHistoryOk

`func (o *WindowedBalanceHistory) GetBurndownHistoryOk() (*[]GrantBurnDownHistorySegment, bool)`

GetBurndownHistoryOk returns a tuple with the BurndownHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurndownHistory

`func (o *WindowedBalanceHistory) SetBurndownHistory(v []GrantBurnDownHistorySegment)`

SetBurndownHistory sets BurndownHistory field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


