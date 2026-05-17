# VarianceSummaryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PeriodStart** | **string** |  | 
**PeriodEnd** | **string** |  | 
**Items** | [**[]VarianceDecompositionOut**](VarianceDecompositionOut.md) |  | 
**TotalItems** | **int32** |  | 

## Methods

### NewVarianceSummaryResponse

`func NewVarianceSummaryResponse(periodStart string, periodEnd string, items []VarianceDecompositionOut, totalItems int32, ) *VarianceSummaryResponse`

NewVarianceSummaryResponse instantiates a new VarianceSummaryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVarianceSummaryResponseWithDefaults

`func NewVarianceSummaryResponseWithDefaults() *VarianceSummaryResponse`

NewVarianceSummaryResponseWithDefaults instantiates a new VarianceSummaryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriodStart

`func (o *VarianceSummaryResponse) GetPeriodStart() string`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *VarianceSummaryResponse) GetPeriodStartOk() (*string, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *VarianceSummaryResponse) SetPeriodStart(v string)`

SetPeriodStart sets PeriodStart field to given value.


### GetPeriodEnd

`func (o *VarianceSummaryResponse) GetPeriodEnd() string`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *VarianceSummaryResponse) GetPeriodEndOk() (*string, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *VarianceSummaryResponse) SetPeriodEnd(v string)`

SetPeriodEnd sets PeriodEnd field to given value.


### GetItems

`func (o *VarianceSummaryResponse) GetItems() []VarianceDecompositionOut`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *VarianceSummaryResponse) GetItemsOk() (*[]VarianceDecompositionOut, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *VarianceSummaryResponse) SetItems(v []VarianceDecompositionOut)`

SetItems sets Items field to given value.


### GetTotalItems

`func (o *VarianceSummaryResponse) GetTotalItems() int32`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *VarianceSummaryResponse) GetTotalItemsOk() (*int32, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *VarianceSummaryResponse) SetTotalItems(v int32)`

SetTotalItems sets TotalItems field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


