# SubscriptionChangeResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Current** | [**Subscription**](Subscription.md) | The current subscription before the change. | 
**Next** | [**SubscriptionExpanded**](SubscriptionExpanded.md) | The new state of the subscription after the change. | 

## Methods

### NewSubscriptionChangeResponseBody

`func NewSubscriptionChangeResponseBody(current Subscription, next SubscriptionExpanded, ) *SubscriptionChangeResponseBody`

NewSubscriptionChangeResponseBody instantiates a new SubscriptionChangeResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionChangeResponseBodyWithDefaults

`func NewSubscriptionChangeResponseBodyWithDefaults() *SubscriptionChangeResponseBody`

NewSubscriptionChangeResponseBodyWithDefaults instantiates a new SubscriptionChangeResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrent

`func (o *SubscriptionChangeResponseBody) GetCurrent() Subscription`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *SubscriptionChangeResponseBody) GetCurrentOk() (*Subscription, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *SubscriptionChangeResponseBody) SetCurrent(v Subscription)`

SetCurrent sets Current field to given value.


### GetNext

`func (o *SubscriptionChangeResponseBody) GetNext() SubscriptionExpanded`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *SubscriptionChangeResponseBody) GetNextOk() (*SubscriptionExpanded, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *SubscriptionChangeResponseBody) SetNext(v SubscriptionExpanded)`

SetNext sets Next field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


