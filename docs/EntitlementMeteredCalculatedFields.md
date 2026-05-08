# EntitlementMeteredCalculatedFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastReset** | **time.Time** | The time the last reset happened. | [readonly] 
**CurrentUsagePeriod** | [**Period**](Period.md) | The current usage period. | [readonly] 
**MeasureUsageFrom** | **time.Time** | The time from which usage is measured. If not specified on creation, defaults to entitlement creation time. | [readonly] 
**UsagePeriod** | [**RecurringPeriod**](RecurringPeriod.md) | THe usage period of the entitlement. | [readonly] 

## Methods

### NewEntitlementMeteredCalculatedFields

`func NewEntitlementMeteredCalculatedFields(lastReset time.Time, currentUsagePeriod Period, measureUsageFrom time.Time, usagePeriod RecurringPeriod, ) *EntitlementMeteredCalculatedFields`

NewEntitlementMeteredCalculatedFields instantiates a new EntitlementMeteredCalculatedFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementMeteredCalculatedFieldsWithDefaults

`func NewEntitlementMeteredCalculatedFieldsWithDefaults() *EntitlementMeteredCalculatedFields`

NewEntitlementMeteredCalculatedFieldsWithDefaults instantiates a new EntitlementMeteredCalculatedFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastReset

`func (o *EntitlementMeteredCalculatedFields) GetLastReset() time.Time`

GetLastReset returns the LastReset field if non-nil, zero value otherwise.

### GetLastResetOk

`func (o *EntitlementMeteredCalculatedFields) GetLastResetOk() (*time.Time, bool)`

GetLastResetOk returns a tuple with the LastReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReset

`func (o *EntitlementMeteredCalculatedFields) SetLastReset(v time.Time)`

SetLastReset sets LastReset field to given value.


### GetCurrentUsagePeriod

`func (o *EntitlementMeteredCalculatedFields) GetCurrentUsagePeriod() Period`

GetCurrentUsagePeriod returns the CurrentUsagePeriod field if non-nil, zero value otherwise.

### GetCurrentUsagePeriodOk

`func (o *EntitlementMeteredCalculatedFields) GetCurrentUsagePeriodOk() (*Period, bool)`

GetCurrentUsagePeriodOk returns a tuple with the CurrentUsagePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUsagePeriod

`func (o *EntitlementMeteredCalculatedFields) SetCurrentUsagePeriod(v Period)`

SetCurrentUsagePeriod sets CurrentUsagePeriod field to given value.


### GetMeasureUsageFrom

`func (o *EntitlementMeteredCalculatedFields) GetMeasureUsageFrom() time.Time`

GetMeasureUsageFrom returns the MeasureUsageFrom field if non-nil, zero value otherwise.

### GetMeasureUsageFromOk

`func (o *EntitlementMeteredCalculatedFields) GetMeasureUsageFromOk() (*time.Time, bool)`

GetMeasureUsageFromOk returns a tuple with the MeasureUsageFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasureUsageFrom

`func (o *EntitlementMeteredCalculatedFields) SetMeasureUsageFrom(v time.Time)`

SetMeasureUsageFrom sets MeasureUsageFrom field to given value.


### GetUsagePeriod

`func (o *EntitlementMeteredCalculatedFields) GetUsagePeriod() RecurringPeriod`

GetUsagePeriod returns the UsagePeriod field if non-nil, zero value otherwise.

### GetUsagePeriodOk

`func (o *EntitlementMeteredCalculatedFields) GetUsagePeriodOk() (*RecurringPeriod, bool)`

GetUsagePeriodOk returns a tuple with the UsagePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsagePeriod

`func (o *EntitlementMeteredCalculatedFields) SetUsagePeriod(v RecurringPeriod)`

SetUsagePeriod sets UsagePeriod field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


