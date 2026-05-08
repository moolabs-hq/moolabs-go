# MeterQueryRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **float64** | The aggregated value. | 
**WindowStart** | **time.Time** | The start of the window the value is aggregated over. | 
**WindowEnd** | **time.Time** | The end of the window the value is aggregated over. | 
**Subject** | **string** | The subject the value is aggregated over. If not specified, the value is aggregated over all subjects. | 
**CustomerId** | Pointer to **string** | The customer ID the value is aggregated over. | [optional] 
**GroupBy** | **map[string]string** | The group by values the value is aggregated over. | 

## Methods

### NewMeterQueryRow

`func NewMeterQueryRow(value float64, windowStart time.Time, windowEnd time.Time, subject string, groupBy map[string]string, ) *MeterQueryRow`

NewMeterQueryRow instantiates a new MeterQueryRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeterQueryRowWithDefaults

`func NewMeterQueryRowWithDefaults() *MeterQueryRow`

NewMeterQueryRowWithDefaults instantiates a new MeterQueryRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *MeterQueryRow) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MeterQueryRow) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MeterQueryRow) SetValue(v float64)`

SetValue sets Value field to given value.


### GetWindowStart

`func (o *MeterQueryRow) GetWindowStart() time.Time`

GetWindowStart returns the WindowStart field if non-nil, zero value otherwise.

### GetWindowStartOk

`func (o *MeterQueryRow) GetWindowStartOk() (*time.Time, bool)`

GetWindowStartOk returns a tuple with the WindowStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowStart

`func (o *MeterQueryRow) SetWindowStart(v time.Time)`

SetWindowStart sets WindowStart field to given value.


### GetWindowEnd

`func (o *MeterQueryRow) GetWindowEnd() time.Time`

GetWindowEnd returns the WindowEnd field if non-nil, zero value otherwise.

### GetWindowEndOk

`func (o *MeterQueryRow) GetWindowEndOk() (*time.Time, bool)`

GetWindowEndOk returns a tuple with the WindowEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowEnd

`func (o *MeterQueryRow) SetWindowEnd(v time.Time)`

SetWindowEnd sets WindowEnd field to given value.


### GetSubject

`func (o *MeterQueryRow) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *MeterQueryRow) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *MeterQueryRow) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetCustomerId

`func (o *MeterQueryRow) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *MeterQueryRow) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *MeterQueryRow) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *MeterQueryRow) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetGroupBy

`func (o *MeterQueryRow) GetGroupBy() map[string]string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *MeterQueryRow) GetGroupByOk() (*map[string]string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *MeterQueryRow) SetGroupBy(v map[string]string)`

SetGroupBy sets GroupBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


