# LifecycleEventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | **string** | Subscription ID | 
**EventType** | **string** | Event type: ACTIVATED, UPDATED, CANCELLED, PAUSED, RESUMED | 
**EventId** | **string** | Unique event ID (for idempotency) | 
**EffectiveAt** | **string** | When the event takes effect (ISO 8601) | 
**SubscriptionData** | Pointer to **map[string]interface{}** |  | [optional] 
**PayloadJson** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewLifecycleEventRequest

`func NewLifecycleEventRequest(subscriptionId string, eventType string, eventId string, effectiveAt string, ) *LifecycleEventRequest`

NewLifecycleEventRequest instantiates a new LifecycleEventRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLifecycleEventRequestWithDefaults

`func NewLifecycleEventRequestWithDefaults() *LifecycleEventRequest`

NewLifecycleEventRequestWithDefaults instantiates a new LifecycleEventRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *LifecycleEventRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *LifecycleEventRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *LifecycleEventRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetEventType

`func (o *LifecycleEventRequest) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *LifecycleEventRequest) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *LifecycleEventRequest) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetEventId

`func (o *LifecycleEventRequest) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *LifecycleEventRequest) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *LifecycleEventRequest) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetEffectiveAt

`func (o *LifecycleEventRequest) GetEffectiveAt() string`

GetEffectiveAt returns the EffectiveAt field if non-nil, zero value otherwise.

### GetEffectiveAtOk

`func (o *LifecycleEventRequest) GetEffectiveAtOk() (*string, bool)`

GetEffectiveAtOk returns a tuple with the EffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAt

`func (o *LifecycleEventRequest) SetEffectiveAt(v string)`

SetEffectiveAt sets EffectiveAt field to given value.


### GetSubscriptionData

`func (o *LifecycleEventRequest) GetSubscriptionData() map[string]interface{}`

GetSubscriptionData returns the SubscriptionData field if non-nil, zero value otherwise.

### GetSubscriptionDataOk

`func (o *LifecycleEventRequest) GetSubscriptionDataOk() (*map[string]interface{}, bool)`

GetSubscriptionDataOk returns a tuple with the SubscriptionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionData

`func (o *LifecycleEventRequest) SetSubscriptionData(v map[string]interface{})`

SetSubscriptionData sets SubscriptionData field to given value.

### HasSubscriptionData

`func (o *LifecycleEventRequest) HasSubscriptionData() bool`

HasSubscriptionData returns a boolean if a field has been set.

### SetSubscriptionDataNil

`func (o *LifecycleEventRequest) SetSubscriptionDataNil(b bool)`

 SetSubscriptionDataNil sets the value for SubscriptionData to be an explicit nil

### UnsetSubscriptionData
`func (o *LifecycleEventRequest) UnsetSubscriptionData()`

UnsetSubscriptionData ensures that no value is present for SubscriptionData, not even an explicit nil
### GetPayloadJson

`func (o *LifecycleEventRequest) GetPayloadJson() string`

GetPayloadJson returns the PayloadJson field if non-nil, zero value otherwise.

### GetPayloadJsonOk

`func (o *LifecycleEventRequest) GetPayloadJsonOk() (*string, bool)`

GetPayloadJsonOk returns a tuple with the PayloadJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadJson

`func (o *LifecycleEventRequest) SetPayloadJson(v string)`

SetPayloadJson sets PayloadJson field to given value.

### HasPayloadJson

`func (o *LifecycleEventRequest) HasPayloadJson() bool`

HasPayloadJson returns a boolean if a field has been set.

### SetPayloadJsonNil

`func (o *LifecycleEventRequest) SetPayloadJsonNil(b bool)`

 SetPayloadJsonNil sets the value for PayloadJson to be an explicit nil

### UnsetPayloadJson
`func (o *LifecycleEventRequest) UnsetPayloadJson()`

UnsetPayloadJson ensures that no value is present for PayloadJson, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


