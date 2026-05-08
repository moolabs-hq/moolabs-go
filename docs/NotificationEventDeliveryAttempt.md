# NotificationEventDeliveryAttempt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | [**NotificationEventDeliveryStatusState**](NotificationEventDeliveryStatusState.md) | State of teh delivery attempt. | [readonly] 
**Response** | [**EventDeliveryAttemptResponse**](EventDeliveryAttemptResponse.md) | Response returned by the notification event recipient. | [readonly] 
**Timestamp** | **time.Time** | Timestamp of the delivery attempt. | [readonly] 

## Methods

### NewNotificationEventDeliveryAttempt

`func NewNotificationEventDeliveryAttempt(state NotificationEventDeliveryStatusState, response EventDeliveryAttemptResponse, timestamp time.Time, ) *NotificationEventDeliveryAttempt`

NewNotificationEventDeliveryAttempt instantiates a new NotificationEventDeliveryAttempt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationEventDeliveryAttemptWithDefaults

`func NewNotificationEventDeliveryAttemptWithDefaults() *NotificationEventDeliveryAttempt`

NewNotificationEventDeliveryAttemptWithDefaults instantiates a new NotificationEventDeliveryAttempt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *NotificationEventDeliveryAttempt) GetState() NotificationEventDeliveryStatusState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NotificationEventDeliveryAttempt) GetStateOk() (*NotificationEventDeliveryStatusState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NotificationEventDeliveryAttempt) SetState(v NotificationEventDeliveryStatusState)`

SetState sets State field to given value.


### GetResponse

`func (o *NotificationEventDeliveryAttempt) GetResponse() EventDeliveryAttemptResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *NotificationEventDeliveryAttempt) GetResponseOk() (*EventDeliveryAttemptResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *NotificationEventDeliveryAttempt) SetResponse(v EventDeliveryAttemptResponse)`

SetResponse sets Response field to given value.


### GetTimestamp

`func (o *NotificationEventDeliveryAttempt) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *NotificationEventDeliveryAttempt) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *NotificationEventDeliveryAttempt) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


