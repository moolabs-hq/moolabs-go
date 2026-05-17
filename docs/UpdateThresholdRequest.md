# UpdateThresholdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThresholdValue** | Pointer to **int32** | New threshold value | [optional] 
**NotifyChannels** | Pointer to [**[]NotifyChannel**](NotifyChannel.md) | New notification channels | [optional] 

## Methods

### NewUpdateThresholdRequest

`func NewUpdateThresholdRequest() *UpdateThresholdRequest`

NewUpdateThresholdRequest instantiates a new UpdateThresholdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateThresholdRequestWithDefaults

`func NewUpdateThresholdRequestWithDefaults() *UpdateThresholdRequest`

NewUpdateThresholdRequestWithDefaults instantiates a new UpdateThresholdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThresholdValue

`func (o *UpdateThresholdRequest) GetThresholdValue() int32`

GetThresholdValue returns the ThresholdValue field if non-nil, zero value otherwise.

### GetThresholdValueOk

`func (o *UpdateThresholdRequest) GetThresholdValueOk() (*int32, bool)`

GetThresholdValueOk returns a tuple with the ThresholdValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdValue

`func (o *UpdateThresholdRequest) SetThresholdValue(v int32)`

SetThresholdValue sets ThresholdValue field to given value.

### HasThresholdValue

`func (o *UpdateThresholdRequest) HasThresholdValue() bool`

HasThresholdValue returns a boolean if a field has been set.

### GetNotifyChannels

`func (o *UpdateThresholdRequest) GetNotifyChannels() []NotifyChannel`

GetNotifyChannels returns the NotifyChannels field if non-nil, zero value otherwise.

### GetNotifyChannelsOk

`func (o *UpdateThresholdRequest) GetNotifyChannelsOk() (*[]NotifyChannel, bool)`

GetNotifyChannelsOk returns a tuple with the NotifyChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyChannels

`func (o *UpdateThresholdRequest) SetNotifyChannels(v []NotifyChannel)`

SetNotifyChannels sets NotifyChannels field to given value.

### HasNotifyChannels

`func (o *UpdateThresholdRequest) HasNotifyChannels() bool`

HasNotifyChannels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


