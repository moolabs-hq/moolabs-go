# RecurringPeriodV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interval** | [**RecurringPeriodInterval**](RecurringPeriodInterval.md) | The unit of time for the interval. Heuristically maps ISO duraitons to enum values or returns the ISO duration. | 
**Anchor** | **time.Time** | A date-time anchor to base the recurring period on. | 

## Methods

### NewRecurringPeriodV2

`func NewRecurringPeriodV2(interval RecurringPeriodInterval, anchor time.Time, ) *RecurringPeriodV2`

NewRecurringPeriodV2 instantiates a new RecurringPeriodV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecurringPeriodV2WithDefaults

`func NewRecurringPeriodV2WithDefaults() *RecurringPeriodV2`

NewRecurringPeriodV2WithDefaults instantiates a new RecurringPeriodV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterval

`func (o *RecurringPeriodV2) GetInterval() RecurringPeriodInterval`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *RecurringPeriodV2) GetIntervalOk() (*RecurringPeriodInterval, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *RecurringPeriodV2) SetInterval(v RecurringPeriodInterval)`

SetInterval sets Interval field to given value.


### GetAnchor

`func (o *RecurringPeriodV2) GetAnchor() time.Time`

GetAnchor returns the Anchor field if non-nil, zero value otherwise.

### GetAnchorOk

`func (o *RecurringPeriodV2) GetAnchorOk() (*time.Time, bool)`

GetAnchorOk returns a tuple with the Anchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchor

`func (o *RecurringPeriodV2) SetAnchor(v time.Time)`

SetAnchor sets Anchor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


