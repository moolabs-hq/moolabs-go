# NotificationEventResendRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channels** | Pointer to **[]string** | Notification channels to which the event should be re-sent. | [optional] 

## Methods

### NewNotificationEventResendRequest

`func NewNotificationEventResendRequest() *NotificationEventResendRequest`

NewNotificationEventResendRequest instantiates a new NotificationEventResendRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationEventResendRequestWithDefaults

`func NewNotificationEventResendRequestWithDefaults() *NotificationEventResendRequest`

NewNotificationEventResendRequestWithDefaults instantiates a new NotificationEventResendRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannels

`func (o *NotificationEventResendRequest) GetChannels() []string`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *NotificationEventResendRequest) GetChannelsOk() (*[]string, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *NotificationEventResendRequest) SetChannels(v []string)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *NotificationEventResendRequest) HasChannels() bool`

HasChannels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


