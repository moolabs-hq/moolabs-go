# RecurringPeriodCreateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interval** | [**RecurringPeriodInterval**](RecurringPeriodInterval.md) | The unit of time for the interval. | 
**Anchor** | Pointer to **time.Time** | A date-time anchor to base the recurring period on. | [optional] 

## Methods

### NewRecurringPeriodCreateInput

`func NewRecurringPeriodCreateInput(interval RecurringPeriodInterval, ) *RecurringPeriodCreateInput`

NewRecurringPeriodCreateInput instantiates a new RecurringPeriodCreateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecurringPeriodCreateInputWithDefaults

`func NewRecurringPeriodCreateInputWithDefaults() *RecurringPeriodCreateInput`

NewRecurringPeriodCreateInputWithDefaults instantiates a new RecurringPeriodCreateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterval

`func (o *RecurringPeriodCreateInput) GetInterval() RecurringPeriodInterval`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *RecurringPeriodCreateInput) GetIntervalOk() (*RecurringPeriodInterval, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *RecurringPeriodCreateInput) SetInterval(v RecurringPeriodInterval)`

SetInterval sets Interval field to given value.


### GetAnchor

`func (o *RecurringPeriodCreateInput) GetAnchor() time.Time`

GetAnchor returns the Anchor field if non-nil, zero value otherwise.

### GetAnchorOk

`func (o *RecurringPeriodCreateInput) GetAnchorOk() (*time.Time, bool)`

GetAnchorOk returns a tuple with the Anchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchor

`func (o *RecurringPeriodCreateInput) SetAnchor(v time.Time)`

SetAnchor sets Anchor field to given value.

### HasAnchor

`func (o *RecurringPeriodCreateInput) HasAnchor() bool`

HasAnchor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


