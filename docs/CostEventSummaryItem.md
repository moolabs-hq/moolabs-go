# CostEventSummaryItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsageEventId** | **string** |  | 
**Status** | **string** |  | 
**ReportingTotalCost** | Pointer to **string** |  | [optional] 
**ReportingCurrency** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**LineItemCount** | Pointer to **int32** |  | [optional] [default to 0]
**PricedLineItemCount** | Pointer to **int32** |  | [optional] [default to 0]
**CostEventCount** | Pointer to **int32** |  | [optional] [default to 0]

## Methods

### NewCostEventSummaryItem

`func NewCostEventSummaryItem(usageEventId string, status string, ) *CostEventSummaryItem`

NewCostEventSummaryItem instantiates a new CostEventSummaryItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostEventSummaryItemWithDefaults

`func NewCostEventSummaryItemWithDefaults() *CostEventSummaryItem`

NewCostEventSummaryItemWithDefaults instantiates a new CostEventSummaryItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsageEventId

`func (o *CostEventSummaryItem) GetUsageEventId() string`

GetUsageEventId returns the UsageEventId field if non-nil, zero value otherwise.

### GetUsageEventIdOk

`func (o *CostEventSummaryItem) GetUsageEventIdOk() (*string, bool)`

GetUsageEventIdOk returns a tuple with the UsageEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageEventId

`func (o *CostEventSummaryItem) SetUsageEventId(v string)`

SetUsageEventId sets UsageEventId field to given value.


### GetStatus

`func (o *CostEventSummaryItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CostEventSummaryItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CostEventSummaryItem) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetReportingTotalCost

`func (o *CostEventSummaryItem) GetReportingTotalCost() string`

GetReportingTotalCost returns the ReportingTotalCost field if non-nil, zero value otherwise.

### GetReportingTotalCostOk

`func (o *CostEventSummaryItem) GetReportingTotalCostOk() (*string, bool)`

GetReportingTotalCostOk returns a tuple with the ReportingTotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingTotalCost

`func (o *CostEventSummaryItem) SetReportingTotalCost(v string)`

SetReportingTotalCost sets ReportingTotalCost field to given value.

### HasReportingTotalCost

`func (o *CostEventSummaryItem) HasReportingTotalCost() bool`

HasReportingTotalCost returns a boolean if a field has been set.

### GetReportingCurrency

`func (o *CostEventSummaryItem) GetReportingCurrency() string`

GetReportingCurrency returns the ReportingCurrency field if non-nil, zero value otherwise.

### GetReportingCurrencyOk

`func (o *CostEventSummaryItem) GetReportingCurrencyOk() (*string, bool)`

GetReportingCurrencyOk returns a tuple with the ReportingCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingCurrency

`func (o *CostEventSummaryItem) SetReportingCurrency(v string)`

SetReportingCurrency sets ReportingCurrency field to given value.

### HasReportingCurrency

`func (o *CostEventSummaryItem) HasReportingCurrency() bool`

HasReportingCurrency returns a boolean if a field has been set.

### GetProvider

`func (o *CostEventSummaryItem) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CostEventSummaryItem) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CostEventSummaryItem) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CostEventSummaryItem) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetModel

`func (o *CostEventSummaryItem) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CostEventSummaryItem) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CostEventSummaryItem) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *CostEventSummaryItem) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetLineItemCount

`func (o *CostEventSummaryItem) GetLineItemCount() int32`

GetLineItemCount returns the LineItemCount field if non-nil, zero value otherwise.

### GetLineItemCountOk

`func (o *CostEventSummaryItem) GetLineItemCountOk() (*int32, bool)`

GetLineItemCountOk returns a tuple with the LineItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItemCount

`func (o *CostEventSummaryItem) SetLineItemCount(v int32)`

SetLineItemCount sets LineItemCount field to given value.

### HasLineItemCount

`func (o *CostEventSummaryItem) HasLineItemCount() bool`

HasLineItemCount returns a boolean if a field has been set.

### GetPricedLineItemCount

`func (o *CostEventSummaryItem) GetPricedLineItemCount() int32`

GetPricedLineItemCount returns the PricedLineItemCount field if non-nil, zero value otherwise.

### GetPricedLineItemCountOk

`func (o *CostEventSummaryItem) GetPricedLineItemCountOk() (*int32, bool)`

GetPricedLineItemCountOk returns a tuple with the PricedLineItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricedLineItemCount

`func (o *CostEventSummaryItem) SetPricedLineItemCount(v int32)`

SetPricedLineItemCount sets PricedLineItemCount field to given value.

### HasPricedLineItemCount

`func (o *CostEventSummaryItem) HasPricedLineItemCount() bool`

HasPricedLineItemCount returns a boolean if a field has been set.

### GetCostEventCount

`func (o *CostEventSummaryItem) GetCostEventCount() int32`

GetCostEventCount returns the CostEventCount field if non-nil, zero value otherwise.

### GetCostEventCountOk

`func (o *CostEventSummaryItem) GetCostEventCountOk() (*int32, bool)`

GetCostEventCountOk returns a tuple with the CostEventCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostEventCount

`func (o *CostEventSummaryItem) SetCostEventCount(v int32)`

SetCostEventCount sets CostEventCount field to given value.

### HasCostEventCount

`func (o *CostEventSummaryItem) HasCostEventCount() bool`

HasCostEventCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


