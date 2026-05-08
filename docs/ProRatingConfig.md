# ProRatingConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Whether pro-rating is enabled for this plan. | [default to true]
**Mode** | [**ProRatingMode**](ProRatingMode.md) | How to handle pro-rating for billing period changes. | 

## Methods

### NewProRatingConfig

`func NewProRatingConfig(enabled bool, mode ProRatingMode, ) *ProRatingConfig`

NewProRatingConfig instantiates a new ProRatingConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProRatingConfigWithDefaults

`func NewProRatingConfigWithDefaults() *ProRatingConfig`

NewProRatingConfigWithDefaults instantiates a new ProRatingConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ProRatingConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ProRatingConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ProRatingConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMode

`func (o *ProRatingConfig) GetMode() ProRatingMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ProRatingConfig) GetModeOk() (*ProRatingMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ProRatingConfig) SetMode(v ProRatingMode)`

SetMode sets Mode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


