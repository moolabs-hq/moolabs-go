# CashAppConfigUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FuzzyMatchTolerancePct** | Pointer to **float32** | Fuzzy match tolerance as decimal (e.g. 0.01 &#x3D; ±1%) | [optional] 
**ExactMatchAutoApply** | Pointer to **bool** |  | [optional] 
**FuzzyMatchThreshold** | Pointer to **float32** | Confidence threshold for fuzzy name matching | [optional] 
**MultiInvoiceMatchEnabled** | Pointer to **bool** |  | [optional] 
**MaxAutoAllocationAmountMicros** | Pointer to **int32** |  | [optional] 

## Methods

### NewCashAppConfigUpdateRequest

`func NewCashAppConfigUpdateRequest() *CashAppConfigUpdateRequest`

NewCashAppConfigUpdateRequest instantiates a new CashAppConfigUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCashAppConfigUpdateRequestWithDefaults

`func NewCashAppConfigUpdateRequestWithDefaults() *CashAppConfigUpdateRequest`

NewCashAppConfigUpdateRequestWithDefaults instantiates a new CashAppConfigUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFuzzyMatchTolerancePct

`func (o *CashAppConfigUpdateRequest) GetFuzzyMatchTolerancePct() float32`

GetFuzzyMatchTolerancePct returns the FuzzyMatchTolerancePct field if non-nil, zero value otherwise.

### GetFuzzyMatchTolerancePctOk

`func (o *CashAppConfigUpdateRequest) GetFuzzyMatchTolerancePctOk() (*float32, bool)`

GetFuzzyMatchTolerancePctOk returns a tuple with the FuzzyMatchTolerancePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuzzyMatchTolerancePct

`func (o *CashAppConfigUpdateRequest) SetFuzzyMatchTolerancePct(v float32)`

SetFuzzyMatchTolerancePct sets FuzzyMatchTolerancePct field to given value.

### HasFuzzyMatchTolerancePct

`func (o *CashAppConfigUpdateRequest) HasFuzzyMatchTolerancePct() bool`

HasFuzzyMatchTolerancePct returns a boolean if a field has been set.

### GetExactMatchAutoApply

`func (o *CashAppConfigUpdateRequest) GetExactMatchAutoApply() bool`

GetExactMatchAutoApply returns the ExactMatchAutoApply field if non-nil, zero value otherwise.

### GetExactMatchAutoApplyOk

`func (o *CashAppConfigUpdateRequest) GetExactMatchAutoApplyOk() (*bool, bool)`

GetExactMatchAutoApplyOk returns a tuple with the ExactMatchAutoApply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExactMatchAutoApply

`func (o *CashAppConfigUpdateRequest) SetExactMatchAutoApply(v bool)`

SetExactMatchAutoApply sets ExactMatchAutoApply field to given value.

### HasExactMatchAutoApply

`func (o *CashAppConfigUpdateRequest) HasExactMatchAutoApply() bool`

HasExactMatchAutoApply returns a boolean if a field has been set.

### GetFuzzyMatchThreshold

`func (o *CashAppConfigUpdateRequest) GetFuzzyMatchThreshold() float32`

GetFuzzyMatchThreshold returns the FuzzyMatchThreshold field if non-nil, zero value otherwise.

### GetFuzzyMatchThresholdOk

`func (o *CashAppConfigUpdateRequest) GetFuzzyMatchThresholdOk() (*float32, bool)`

GetFuzzyMatchThresholdOk returns a tuple with the FuzzyMatchThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuzzyMatchThreshold

`func (o *CashAppConfigUpdateRequest) SetFuzzyMatchThreshold(v float32)`

SetFuzzyMatchThreshold sets FuzzyMatchThreshold field to given value.

### HasFuzzyMatchThreshold

`func (o *CashAppConfigUpdateRequest) HasFuzzyMatchThreshold() bool`

HasFuzzyMatchThreshold returns a boolean if a field has been set.

### GetMultiInvoiceMatchEnabled

`func (o *CashAppConfigUpdateRequest) GetMultiInvoiceMatchEnabled() bool`

GetMultiInvoiceMatchEnabled returns the MultiInvoiceMatchEnabled field if non-nil, zero value otherwise.

### GetMultiInvoiceMatchEnabledOk

`func (o *CashAppConfigUpdateRequest) GetMultiInvoiceMatchEnabledOk() (*bool, bool)`

GetMultiInvoiceMatchEnabledOk returns a tuple with the MultiInvoiceMatchEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiInvoiceMatchEnabled

`func (o *CashAppConfigUpdateRequest) SetMultiInvoiceMatchEnabled(v bool)`

SetMultiInvoiceMatchEnabled sets MultiInvoiceMatchEnabled field to given value.

### HasMultiInvoiceMatchEnabled

`func (o *CashAppConfigUpdateRequest) HasMultiInvoiceMatchEnabled() bool`

HasMultiInvoiceMatchEnabled returns a boolean if a field has been set.

### GetMaxAutoAllocationAmountMicros

`func (o *CashAppConfigUpdateRequest) GetMaxAutoAllocationAmountMicros() int32`

GetMaxAutoAllocationAmountMicros returns the MaxAutoAllocationAmountMicros field if non-nil, zero value otherwise.

### GetMaxAutoAllocationAmountMicrosOk

`func (o *CashAppConfigUpdateRequest) GetMaxAutoAllocationAmountMicrosOk() (*int32, bool)`

GetMaxAutoAllocationAmountMicrosOk returns a tuple with the MaxAutoAllocationAmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAutoAllocationAmountMicros

`func (o *CashAppConfigUpdateRequest) SetMaxAutoAllocationAmountMicros(v int32)`

SetMaxAutoAllocationAmountMicros sets MaxAutoAllocationAmountMicros field to given value.

### HasMaxAutoAllocationAmountMicros

`func (o *CashAppConfigUpdateRequest) HasMaxAutoAllocationAmountMicros() bool`

HasMaxAutoAllocationAmountMicros returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


