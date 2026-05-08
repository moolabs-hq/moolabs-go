# NotificationChannelMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifies the notification channel. | [readonly] 
**Type** | [**NotificationChannelType**](NotificationChannelType.md) | Notification channel type. | 

## Methods

### NewNotificationChannelMeta

`func NewNotificationChannelMeta(id string, type_ NotificationChannelType, ) *NotificationChannelMeta`

NewNotificationChannelMeta instantiates a new NotificationChannelMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationChannelMetaWithDefaults

`func NewNotificationChannelMetaWithDefaults() *NotificationChannelMeta`

NewNotificationChannelMetaWithDefaults instantiates a new NotificationChannelMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotificationChannelMeta) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationChannelMeta) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationChannelMeta) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *NotificationChannelMeta) GetType() NotificationChannelType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationChannelMeta) GetTypeOk() (*NotificationChannelType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationChannelMeta) SetType(v NotificationChannelType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


