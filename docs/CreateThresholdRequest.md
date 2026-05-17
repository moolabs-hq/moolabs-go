# CreateThresholdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThresholdType** | **string** | PERCENT or ABSOLUTE | 
**ThresholdValue** | **int32** | Threshold value (0-100 for PERCENT, micros for ABSOLUTE) | 
**NotifyChannels** | Pointer to [**[]NotifyChannel**](NotifyChannel.md) | Structured notification channels, e.g. [{\&quot;type\&quot;: \&quot;email\&quot;, \&quot;address\&quot;: \&quot;ops@example.com\&quot;}] | [optional] 

## Methods

### NewCreateThresholdRequest

`func NewCreateThresholdRequest(thresholdType string, thresholdValue int32, ) *CreateThresholdRequest`

NewCreateThresholdRequest instantiates a new CreateThresholdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateThresholdRequestWithDefaults

`func NewCreateThresholdRequestWithDefaults() *CreateThresholdRequest`

NewCreateThresholdRequestWithDefaults instantiates a new CreateThresholdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThresholdType

`func (o *CreateThresholdRequest) GetThresholdType() string`

GetThresholdType returns the ThresholdType field if non-nil, zero value otherwise.

### GetThresholdTypeOk

`func (o *CreateThresholdRequest) GetThresholdTypeOk() (*string, bool)`

GetThresholdTypeOk returns a tuple with the ThresholdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdType

`func (o *CreateThresholdRequest) SetThresholdType(v string)`

SetThresholdType sets ThresholdType field to given value.


### GetThresholdValue

`func (o *CreateThresholdRequest) GetThresholdValue() int32`

GetThresholdValue returns the ThresholdValue field if non-nil, zero value otherwise.

### GetThresholdValueOk

`func (o *CreateThresholdRequest) GetThresholdValueOk() (*int32, bool)`

GetThresholdValueOk returns a tuple with the ThresholdValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdValue

`func (o *CreateThresholdRequest) SetThresholdValue(v int32)`

SetThresholdValue sets ThresholdValue field to given value.


### GetNotifyChannels

`func (o *CreateThresholdRequest) GetNotifyChannels() []NotifyChannel`

GetNotifyChannels returns the NotifyChannels field if non-nil, zero value otherwise.

### GetNotifyChannelsOk

`func (o *CreateThresholdRequest) GetNotifyChannelsOk() (*[]NotifyChannel, bool)`

GetNotifyChannelsOk returns a tuple with the NotifyChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyChannels

`func (o *CreateThresholdRequest) SetNotifyChannels(v []NotifyChannel)`

SetNotifyChannels sets NotifyChannels field to given value.

### HasNotifyChannels

`func (o *CreateThresholdRequest) HasNotifyChannels() bool`

HasNotifyChannels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


