# RecurringPeriod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interval** | [**RecurringPeriodInterval**](RecurringPeriodInterval.md) | The unit of time for the interval. Heuristically maps ISO duraitons to enum values or returns the ISO duration. | 
**Anchor** | **time.Time** | A date-time anchor to base the recurring period on. | 
**IntervalISO** | **string** | The unit of time for the interval in ISO8601 format. | 

## Methods

### NewRecurringPeriod

`func NewRecurringPeriod(interval RecurringPeriodInterval, anchor time.Time, intervalISO string, ) *RecurringPeriod`

NewRecurringPeriod instantiates a new RecurringPeriod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecurringPeriodWithDefaults

`func NewRecurringPeriodWithDefaults() *RecurringPeriod`

NewRecurringPeriodWithDefaults instantiates a new RecurringPeriod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterval

`func (o *RecurringPeriod) GetInterval() RecurringPeriodInterval`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *RecurringPeriod) GetIntervalOk() (*RecurringPeriodInterval, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *RecurringPeriod) SetInterval(v RecurringPeriodInterval)`

SetInterval sets Interval field to given value.


### GetAnchor

`func (o *RecurringPeriod) GetAnchor() time.Time`

GetAnchor returns the Anchor field if non-nil, zero value otherwise.

### GetAnchorOk

`func (o *RecurringPeriod) GetAnchorOk() (*time.Time, bool)`

GetAnchorOk returns a tuple with the Anchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchor

`func (o *RecurringPeriod) SetAnchor(v time.Time)`

SetAnchor sets Anchor field to given value.


### GetIntervalISO

`func (o *RecurringPeriod) GetIntervalISO() string`

GetIntervalISO returns the IntervalISO field if non-nil, zero value otherwise.

### GetIntervalISOOk

`func (o *RecurringPeriod) GetIntervalISOOk() (*string, bool)`

GetIntervalISOOk returns a tuple with the IntervalISO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalISO

`func (o *RecurringPeriod) SetIntervalISO(v string)`

SetIntervalISO sets IntervalISO field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


