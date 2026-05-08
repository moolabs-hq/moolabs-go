# ExpirationPeriod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | [**ExpirationDuration**](ExpirationDuration.md) | The unit of time for the expiration period. | 
**Count** | **int32** | The number of time units in the expiration period. | 

## Methods

### NewExpirationPeriod

`func NewExpirationPeriod(duration ExpirationDuration, count int32, ) *ExpirationPeriod`

NewExpirationPeriod instantiates a new ExpirationPeriod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpirationPeriodWithDefaults

`func NewExpirationPeriodWithDefaults() *ExpirationPeriod`

NewExpirationPeriodWithDefaults instantiates a new ExpirationPeriod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *ExpirationPeriod) GetDuration() ExpirationDuration`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ExpirationPeriod) GetDurationOk() (*ExpirationDuration, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ExpirationPeriod) SetDuration(v ExpirationDuration)`

SetDuration sets Duration field to given value.


### GetCount

`func (o *ExpirationPeriod) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ExpirationPeriod) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ExpirationPeriod) SetCount(v int32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


