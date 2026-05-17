# DSOResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**AsOf** | **time.Time** |  | 
**DsoDays** | **float32** |  | 
**TotalOutstandingMicros** | **int32** |  | 
**TotalRevenueMicros** | **int32** |  | 
**PeriodDays** | Pointer to **int32** |  | [optional] [default to 90]
**CurrencyCode** | Pointer to **string** |  | [optional] [default to "USD"]

## Methods

### NewDSOResponse

`func NewDSOResponse(tenantId string, asOf time.Time, dsoDays float32, totalOutstandingMicros int32, totalRevenueMicros int32, ) *DSOResponse`

NewDSOResponse instantiates a new DSOResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDSOResponseWithDefaults

`func NewDSOResponseWithDefaults() *DSOResponse`

NewDSOResponseWithDefaults instantiates a new DSOResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *DSOResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DSOResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DSOResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetAsOf

`func (o *DSOResponse) GetAsOf() time.Time`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *DSOResponse) GetAsOfOk() (*time.Time, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *DSOResponse) SetAsOf(v time.Time)`

SetAsOf sets AsOf field to given value.


### GetDsoDays

`func (o *DSOResponse) GetDsoDays() float32`

GetDsoDays returns the DsoDays field if non-nil, zero value otherwise.

### GetDsoDaysOk

`func (o *DSOResponse) GetDsoDaysOk() (*float32, bool)`

GetDsoDaysOk returns a tuple with the DsoDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsoDays

`func (o *DSOResponse) SetDsoDays(v float32)`

SetDsoDays sets DsoDays field to given value.


### GetTotalOutstandingMicros

`func (o *DSOResponse) GetTotalOutstandingMicros() int32`

GetTotalOutstandingMicros returns the TotalOutstandingMicros field if non-nil, zero value otherwise.

### GetTotalOutstandingMicrosOk

`func (o *DSOResponse) GetTotalOutstandingMicrosOk() (*int32, bool)`

GetTotalOutstandingMicrosOk returns a tuple with the TotalOutstandingMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOutstandingMicros

`func (o *DSOResponse) SetTotalOutstandingMicros(v int32)`

SetTotalOutstandingMicros sets TotalOutstandingMicros field to given value.


### GetTotalRevenueMicros

`func (o *DSOResponse) GetTotalRevenueMicros() int32`

GetTotalRevenueMicros returns the TotalRevenueMicros field if non-nil, zero value otherwise.

### GetTotalRevenueMicrosOk

`func (o *DSOResponse) GetTotalRevenueMicrosOk() (*int32, bool)`

GetTotalRevenueMicrosOk returns a tuple with the TotalRevenueMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRevenueMicros

`func (o *DSOResponse) SetTotalRevenueMicros(v int32)`

SetTotalRevenueMicros sets TotalRevenueMicros field to given value.


### GetPeriodDays

`func (o *DSOResponse) GetPeriodDays() int32`

GetPeriodDays returns the PeriodDays field if non-nil, zero value otherwise.

### GetPeriodDaysOk

`func (o *DSOResponse) GetPeriodDaysOk() (*int32, bool)`

GetPeriodDaysOk returns a tuple with the PeriodDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodDays

`func (o *DSOResponse) SetPeriodDays(v int32)`

SetPeriodDays sets PeriodDays field to given value.

### HasPeriodDays

`func (o *DSOResponse) HasPeriodDays() bool`

HasPeriodDays returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *DSOResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *DSOResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *DSOResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *DSOResponse) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


