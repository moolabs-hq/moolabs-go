# NotificationEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier of the notification event. | [readonly] 
**Type** | [**NotificationEventType**](NotificationEventType.md) | Type of the notification event. | [readonly] 
**CreatedAt** | **time.Time** | Timestamp when the notification event was created in RFC 3339 format. | [readonly] 
**Rule** | [**NotificationRule**](NotificationRule.md) | The nnotification rule which generated this event. | [readonly] 
**DeliveryStatus** | [**[]NotificationEventDeliveryStatus**](NotificationEventDeliveryStatus.md) | The delivery status of the notification event. | 
**Payload** | [**NotificationEventPayload**](NotificationEventPayload.md) | Timestamp when the notification event was created in RFC 3339 format. | [readonly] 
**Annotations** | Pointer to **map[string]interface{}** | Set of key-value pairs managed by the system. Cannot be modified by user. | [optional] [readonly] 

## Methods

### NewNotificationEvent

`func NewNotificationEvent(id string, type_ NotificationEventType, createdAt time.Time, rule NotificationRule, deliveryStatus []NotificationEventDeliveryStatus, payload NotificationEventPayload, ) *NotificationEvent`

NewNotificationEvent instantiates a new NotificationEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationEventWithDefaults

`func NewNotificationEventWithDefaults() *NotificationEvent`

NewNotificationEventWithDefaults instantiates a new NotificationEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotificationEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationEvent) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *NotificationEvent) GetType() NotificationEventType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationEvent) GetTypeOk() (*NotificationEventType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationEvent) SetType(v NotificationEventType)`

SetType sets Type field to given value.


### GetCreatedAt

`func (o *NotificationEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NotificationEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NotificationEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetRule

`func (o *NotificationEvent) GetRule() NotificationRule`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *NotificationEvent) GetRuleOk() (*NotificationRule, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *NotificationEvent) SetRule(v NotificationRule)`

SetRule sets Rule field to given value.


### GetDeliveryStatus

`func (o *NotificationEvent) GetDeliveryStatus() []NotificationEventDeliveryStatus`

GetDeliveryStatus returns the DeliveryStatus field if non-nil, zero value otherwise.

### GetDeliveryStatusOk

`func (o *NotificationEvent) GetDeliveryStatusOk() (*[]NotificationEventDeliveryStatus, bool)`

GetDeliveryStatusOk returns a tuple with the DeliveryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryStatus

`func (o *NotificationEvent) SetDeliveryStatus(v []NotificationEventDeliveryStatus)`

SetDeliveryStatus sets DeliveryStatus field to given value.


### GetPayload

`func (o *NotificationEvent) GetPayload() NotificationEventPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *NotificationEvent) GetPayloadOk() (*NotificationEventPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *NotificationEvent) SetPayload(v NotificationEventPayload)`

SetPayload sets Payload field to given value.


### GetAnnotations

`func (o *NotificationEvent) GetAnnotations() map[string]interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *NotificationEvent) GetAnnotationsOk() (*map[string]interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *NotificationEvent) SetAnnotations(v map[string]interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *NotificationEvent) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


