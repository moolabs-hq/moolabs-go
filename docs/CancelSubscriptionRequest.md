# CancelSubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timing** | Pointer to [**SubscriptionTiming**](SubscriptionTiming.md) | If not provided the subscription is canceled immediately. | [optional] 

## Methods

### NewCancelSubscriptionRequest

`func NewCancelSubscriptionRequest() *CancelSubscriptionRequest`

NewCancelSubscriptionRequest instantiates a new CancelSubscriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelSubscriptionRequestWithDefaults

`func NewCancelSubscriptionRequestWithDefaults() *CancelSubscriptionRequest`

NewCancelSubscriptionRequestWithDefaults instantiates a new CancelSubscriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTiming

`func (o *CancelSubscriptionRequest) GetTiming() SubscriptionTiming`

GetTiming returns the Timing field if non-nil, zero value otherwise.

### GetTimingOk

`func (o *CancelSubscriptionRequest) GetTimingOk() (*SubscriptionTiming, bool)`

GetTimingOk returns a tuple with the Timing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiming

`func (o *CancelSubscriptionRequest) SetTiming(v SubscriptionTiming)`

SetTiming sets Timing field to given value.

### HasTiming

`func (o *CancelSubscriptionRequest) HasTiming() bool`

HasTiming returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


