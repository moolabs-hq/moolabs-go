# CheckTriggerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsOf** | Pointer to **time.Time** | As-of timestamp (defaults to now) | [optional] 

## Methods

### NewCheckTriggerRequest

`func NewCheckTriggerRequest() *CheckTriggerRequest`

NewCheckTriggerRequest instantiates a new CheckTriggerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckTriggerRequestWithDefaults

`func NewCheckTriggerRequestWithDefaults() *CheckTriggerRequest`

NewCheckTriggerRequestWithDefaults instantiates a new CheckTriggerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsOf

`func (o *CheckTriggerRequest) GetAsOf() time.Time`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *CheckTriggerRequest) GetAsOfOk() (*time.Time, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *CheckTriggerRequest) SetAsOf(v time.Time)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *CheckTriggerRequest) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


