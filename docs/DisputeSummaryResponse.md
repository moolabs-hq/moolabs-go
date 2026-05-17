# DisputeSummaryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**AsOf** | **time.Time** |  | 
**TotalDisputes** | **int32** |  | 
**OpenDisputes** | **int32** |  | 
**TotalDisputedMicros** | **int32** |  | 
**CurrencyCode** | Pointer to **string** |  | [optional] [default to "USD"]
**ResolvedValid** | **int32** |  | 
**ResolvedInvalid** | **int32** |  | 
**AvgResolutionDays** | Pointer to **float32** |  | [optional] 

## Methods

### NewDisputeSummaryResponse

`func NewDisputeSummaryResponse(tenantId string, asOf time.Time, totalDisputes int32, openDisputes int32, totalDisputedMicros int32, resolvedValid int32, resolvedInvalid int32, ) *DisputeSummaryResponse`

NewDisputeSummaryResponse instantiates a new DisputeSummaryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisputeSummaryResponseWithDefaults

`func NewDisputeSummaryResponseWithDefaults() *DisputeSummaryResponse`

NewDisputeSummaryResponseWithDefaults instantiates a new DisputeSummaryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *DisputeSummaryResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DisputeSummaryResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DisputeSummaryResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetAsOf

`func (o *DisputeSummaryResponse) GetAsOf() time.Time`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *DisputeSummaryResponse) GetAsOfOk() (*time.Time, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *DisputeSummaryResponse) SetAsOf(v time.Time)`

SetAsOf sets AsOf field to given value.


### GetTotalDisputes

`func (o *DisputeSummaryResponse) GetTotalDisputes() int32`

GetTotalDisputes returns the TotalDisputes field if non-nil, zero value otherwise.

### GetTotalDisputesOk

`func (o *DisputeSummaryResponse) GetTotalDisputesOk() (*int32, bool)`

GetTotalDisputesOk returns a tuple with the TotalDisputes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDisputes

`func (o *DisputeSummaryResponse) SetTotalDisputes(v int32)`

SetTotalDisputes sets TotalDisputes field to given value.


### GetOpenDisputes

`func (o *DisputeSummaryResponse) GetOpenDisputes() int32`

GetOpenDisputes returns the OpenDisputes field if non-nil, zero value otherwise.

### GetOpenDisputesOk

`func (o *DisputeSummaryResponse) GetOpenDisputesOk() (*int32, bool)`

GetOpenDisputesOk returns a tuple with the OpenDisputes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenDisputes

`func (o *DisputeSummaryResponse) SetOpenDisputes(v int32)`

SetOpenDisputes sets OpenDisputes field to given value.


### GetTotalDisputedMicros

`func (o *DisputeSummaryResponse) GetTotalDisputedMicros() int32`

GetTotalDisputedMicros returns the TotalDisputedMicros field if non-nil, zero value otherwise.

### GetTotalDisputedMicrosOk

`func (o *DisputeSummaryResponse) GetTotalDisputedMicrosOk() (*int32, bool)`

GetTotalDisputedMicrosOk returns a tuple with the TotalDisputedMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDisputedMicros

`func (o *DisputeSummaryResponse) SetTotalDisputedMicros(v int32)`

SetTotalDisputedMicros sets TotalDisputedMicros field to given value.


### GetCurrencyCode

`func (o *DisputeSummaryResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *DisputeSummaryResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *DisputeSummaryResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *DisputeSummaryResponse) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetResolvedValid

`func (o *DisputeSummaryResponse) GetResolvedValid() int32`

GetResolvedValid returns the ResolvedValid field if non-nil, zero value otherwise.

### GetResolvedValidOk

`func (o *DisputeSummaryResponse) GetResolvedValidOk() (*int32, bool)`

GetResolvedValidOk returns a tuple with the ResolvedValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedValid

`func (o *DisputeSummaryResponse) SetResolvedValid(v int32)`

SetResolvedValid sets ResolvedValid field to given value.


### GetResolvedInvalid

`func (o *DisputeSummaryResponse) GetResolvedInvalid() int32`

GetResolvedInvalid returns the ResolvedInvalid field if non-nil, zero value otherwise.

### GetResolvedInvalidOk

`func (o *DisputeSummaryResponse) GetResolvedInvalidOk() (*int32, bool)`

GetResolvedInvalidOk returns a tuple with the ResolvedInvalid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedInvalid

`func (o *DisputeSummaryResponse) SetResolvedInvalid(v int32)`

SetResolvedInvalid sets ResolvedInvalid field to given value.


### GetAvgResolutionDays

`func (o *DisputeSummaryResponse) GetAvgResolutionDays() float32`

GetAvgResolutionDays returns the AvgResolutionDays field if non-nil, zero value otherwise.

### GetAvgResolutionDaysOk

`func (o *DisputeSummaryResponse) GetAvgResolutionDaysOk() (*float32, bool)`

GetAvgResolutionDaysOk returns a tuple with the AvgResolutionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgResolutionDays

`func (o *DisputeSummaryResponse) SetAvgResolutionDays(v float32)`

SetAvgResolutionDays sets AvgResolutionDays field to given value.

### HasAvgResolutionDays

`func (o *DisputeSummaryResponse) HasAvgResolutionDays() bool`

HasAvgResolutionDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


