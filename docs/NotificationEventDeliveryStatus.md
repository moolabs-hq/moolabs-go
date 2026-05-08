# NotificationEventDeliveryStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | [**NotificationEventDeliveryStatusState**](NotificationEventDeliveryStatusState.md) | Delivery state of the notification event to the channel. | [readonly] 
**Reason** | **string** | The reason of the last deliverry state update. | [readonly] 
**UpdatedAt** | **time.Time** | Timestamp of when the status was last updated in RFC 3339 format. | [readonly] 
**Channel** | [**NotificationChannelMeta**](NotificationChannelMeta.md) | Notification channel the delivery status associated with. | [readonly] 
**Annotations** | Pointer to **map[string]interface{}** | Set of key-value pairs managed by the system. Cannot be modified by user. | [optional] [readonly] 
**NextAttempt** | Pointer to **time.Time** | Timestamp of the next delivery attempt. If null it means there will be no more delivery attempts. | [optional] [readonly] 
**Attempts** | [**[]NotificationEventDeliveryAttempt**](NotificationEventDeliveryAttempt.md) | List of delivery attempts. | 

## Methods

### NewNotificationEventDeliveryStatus

`func NewNotificationEventDeliveryStatus(state NotificationEventDeliveryStatusState, reason string, updatedAt time.Time, channel NotificationChannelMeta, attempts []NotificationEventDeliveryAttempt, ) *NotificationEventDeliveryStatus`

NewNotificationEventDeliveryStatus instantiates a new NotificationEventDeliveryStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationEventDeliveryStatusWithDefaults

`func NewNotificationEventDeliveryStatusWithDefaults() *NotificationEventDeliveryStatus`

NewNotificationEventDeliveryStatusWithDefaults instantiates a new NotificationEventDeliveryStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *NotificationEventDeliveryStatus) GetState() NotificationEventDeliveryStatusState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NotificationEventDeliveryStatus) GetStateOk() (*NotificationEventDeliveryStatusState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NotificationEventDeliveryStatus) SetState(v NotificationEventDeliveryStatusState)`

SetState sets State field to given value.


### GetReason

`func (o *NotificationEventDeliveryStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *NotificationEventDeliveryStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *NotificationEventDeliveryStatus) SetReason(v string)`

SetReason sets Reason field to given value.


### GetUpdatedAt

`func (o *NotificationEventDeliveryStatus) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NotificationEventDeliveryStatus) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NotificationEventDeliveryStatus) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetChannel

`func (o *NotificationEventDeliveryStatus) GetChannel() NotificationChannelMeta`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *NotificationEventDeliveryStatus) GetChannelOk() (*NotificationChannelMeta, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *NotificationEventDeliveryStatus) SetChannel(v NotificationChannelMeta)`

SetChannel sets Channel field to given value.


### GetAnnotations

`func (o *NotificationEventDeliveryStatus) GetAnnotations() map[string]interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *NotificationEventDeliveryStatus) GetAnnotationsOk() (*map[string]interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *NotificationEventDeliveryStatus) SetAnnotations(v map[string]interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *NotificationEventDeliveryStatus) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetNextAttempt

`func (o *NotificationEventDeliveryStatus) GetNextAttempt() time.Time`

GetNextAttempt returns the NextAttempt field if non-nil, zero value otherwise.

### GetNextAttemptOk

`func (o *NotificationEventDeliveryStatus) GetNextAttemptOk() (*time.Time, bool)`

GetNextAttemptOk returns a tuple with the NextAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAttempt

`func (o *NotificationEventDeliveryStatus) SetNextAttempt(v time.Time)`

SetNextAttempt sets NextAttempt field to given value.

### HasNextAttempt

`func (o *NotificationEventDeliveryStatus) HasNextAttempt() bool`

HasNextAttempt returns a boolean if a field has been set.

### GetAttempts

`func (o *NotificationEventDeliveryStatus) GetAttempts() []NotificationEventDeliveryAttempt`

GetAttempts returns the Attempts field if non-nil, zero value otherwise.

### GetAttemptsOk

`func (o *NotificationEventDeliveryStatus) GetAttemptsOk() (*[]NotificationEventDeliveryAttempt, bool)`

GetAttemptsOk returns a tuple with the Attempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempts

`func (o *NotificationEventDeliveryStatus) SetAttempts(v []NotificationEventDeliveryAttempt)`

SetAttempts sets Attempts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


