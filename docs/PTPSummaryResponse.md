# PTPSummaryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**AsOf** | **time.Time** |  | 
**OpenCount** | **int32** |  | 
**OpenAmountMicros** | **int32** |  | 
**CurrencyCode** | Pointer to **string** |  | [optional] [default to "USD"]
**DueWithin7Days** | **int32** |  | 
**OverdueCount** | **int32** |  | 
**BrokenCount** | **int32** |  | 
**FulfilledCount** | **int32** |  | 

## Methods

### NewPTPSummaryResponse

`func NewPTPSummaryResponse(tenantId string, asOf time.Time, openCount int32, openAmountMicros int32, dueWithin7Days int32, overdueCount int32, brokenCount int32, fulfilledCount int32, ) *PTPSummaryResponse`

NewPTPSummaryResponse instantiates a new PTPSummaryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPTPSummaryResponseWithDefaults

`func NewPTPSummaryResponseWithDefaults() *PTPSummaryResponse`

NewPTPSummaryResponseWithDefaults instantiates a new PTPSummaryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *PTPSummaryResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *PTPSummaryResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *PTPSummaryResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetAsOf

`func (o *PTPSummaryResponse) GetAsOf() time.Time`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *PTPSummaryResponse) GetAsOfOk() (*time.Time, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *PTPSummaryResponse) SetAsOf(v time.Time)`

SetAsOf sets AsOf field to given value.


### GetOpenCount

`func (o *PTPSummaryResponse) GetOpenCount() int32`

GetOpenCount returns the OpenCount field if non-nil, zero value otherwise.

### GetOpenCountOk

`func (o *PTPSummaryResponse) GetOpenCountOk() (*int32, bool)`

GetOpenCountOk returns a tuple with the OpenCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenCount

`func (o *PTPSummaryResponse) SetOpenCount(v int32)`

SetOpenCount sets OpenCount field to given value.


### GetOpenAmountMicros

`func (o *PTPSummaryResponse) GetOpenAmountMicros() int32`

GetOpenAmountMicros returns the OpenAmountMicros field if non-nil, zero value otherwise.

### GetOpenAmountMicrosOk

`func (o *PTPSummaryResponse) GetOpenAmountMicrosOk() (*int32, bool)`

GetOpenAmountMicrosOk returns a tuple with the OpenAmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenAmountMicros

`func (o *PTPSummaryResponse) SetOpenAmountMicros(v int32)`

SetOpenAmountMicros sets OpenAmountMicros field to given value.


### GetCurrencyCode

`func (o *PTPSummaryResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *PTPSummaryResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *PTPSummaryResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *PTPSummaryResponse) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetDueWithin7Days

`func (o *PTPSummaryResponse) GetDueWithin7Days() int32`

GetDueWithin7Days returns the DueWithin7Days field if non-nil, zero value otherwise.

### GetDueWithin7DaysOk

`func (o *PTPSummaryResponse) GetDueWithin7DaysOk() (*int32, bool)`

GetDueWithin7DaysOk returns a tuple with the DueWithin7Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueWithin7Days

`func (o *PTPSummaryResponse) SetDueWithin7Days(v int32)`

SetDueWithin7Days sets DueWithin7Days field to given value.


### GetOverdueCount

`func (o *PTPSummaryResponse) GetOverdueCount() int32`

GetOverdueCount returns the OverdueCount field if non-nil, zero value otherwise.

### GetOverdueCountOk

`func (o *PTPSummaryResponse) GetOverdueCountOk() (*int32, bool)`

GetOverdueCountOk returns a tuple with the OverdueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverdueCount

`func (o *PTPSummaryResponse) SetOverdueCount(v int32)`

SetOverdueCount sets OverdueCount field to given value.


### GetBrokenCount

`func (o *PTPSummaryResponse) GetBrokenCount() int32`

GetBrokenCount returns the BrokenCount field if non-nil, zero value otherwise.

### GetBrokenCountOk

`func (o *PTPSummaryResponse) GetBrokenCountOk() (*int32, bool)`

GetBrokenCountOk returns a tuple with the BrokenCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokenCount

`func (o *PTPSummaryResponse) SetBrokenCount(v int32)`

SetBrokenCount sets BrokenCount field to given value.


### GetFulfilledCount

`func (o *PTPSummaryResponse) GetFulfilledCount() int32`

GetFulfilledCount returns the FulfilledCount field if non-nil, zero value otherwise.

### GetFulfilledCountOk

`func (o *PTPSummaryResponse) GetFulfilledCountOk() (*int32, bool)`

GetFulfilledCountOk returns a tuple with the FulfilledCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfilledCount

`func (o *PTPSummaryResponse) SetFulfilledCount(v int32)`

SetFulfilledCount sets FulfilledCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


