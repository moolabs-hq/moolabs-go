# SubscriptionEdit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customizations** | [**[]SubscriptionEditOperation**](SubscriptionEditOperation.md) | Batch processing commands for manipulating running subscriptions. The key format is &#x60;/phases/{phaseKey}&#x60; or &#x60;/phases/{phaseKey}/items/{itemKey}&#x60;. | 
**Timing** | Pointer to [**SubscriptionTiming**](SubscriptionTiming.md) | Whether the billing period should be restarted.Timing configuration to allow for the changes to take effect at different times. | [optional] 

## Methods

### NewSubscriptionEdit

`func NewSubscriptionEdit(customizations []SubscriptionEditOperation, ) *SubscriptionEdit`

NewSubscriptionEdit instantiates a new SubscriptionEdit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionEditWithDefaults

`func NewSubscriptionEditWithDefaults() *SubscriptionEdit`

NewSubscriptionEditWithDefaults instantiates a new SubscriptionEdit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomizations

`func (o *SubscriptionEdit) GetCustomizations() []SubscriptionEditOperation`

GetCustomizations returns the Customizations field if non-nil, zero value otherwise.

### GetCustomizationsOk

`func (o *SubscriptionEdit) GetCustomizationsOk() (*[]SubscriptionEditOperation, bool)`

GetCustomizationsOk returns a tuple with the Customizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomizations

`func (o *SubscriptionEdit) SetCustomizations(v []SubscriptionEditOperation)`

SetCustomizations sets Customizations field to given value.


### GetTiming

`func (o *SubscriptionEdit) GetTiming() SubscriptionTiming`

GetTiming returns the Timing field if non-nil, zero value otherwise.

### GetTimingOk

`func (o *SubscriptionEdit) GetTimingOk() (*SubscriptionTiming, bool)`

GetTimingOk returns a tuple with the Timing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiming

`func (o *SubscriptionEdit) SetTiming(v SubscriptionTiming)`

SetTiming sets Timing field to given value.

### HasTiming

`func (o *SubscriptionEdit) HasTiming() bool`

HasTiming returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


