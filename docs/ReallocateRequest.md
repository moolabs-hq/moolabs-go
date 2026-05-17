# ReallocateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PeriodStart** | **time.Time** |  | 
**PeriodEnd** | **time.Time** |  | 

## Methods

### NewReallocateRequest

`func NewReallocateRequest(periodStart time.Time, periodEnd time.Time, ) *ReallocateRequest`

NewReallocateRequest instantiates a new ReallocateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReallocateRequestWithDefaults

`func NewReallocateRequestWithDefaults() *ReallocateRequest`

NewReallocateRequestWithDefaults instantiates a new ReallocateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriodStart

`func (o *ReallocateRequest) GetPeriodStart() time.Time`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *ReallocateRequest) GetPeriodStartOk() (*time.Time, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *ReallocateRequest) SetPeriodStart(v time.Time)`

SetPeriodStart sets PeriodStart field to given value.


### GetPeriodEnd

`func (o *ReallocateRequest) GetPeriodEnd() time.Time`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *ReallocateRequest) GetPeriodEndOk() (*time.Time, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *ReallocateRequest) SetPeriodEnd(v time.Time)`

SetPeriodEnd sets PeriodEnd field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


