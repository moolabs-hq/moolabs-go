# NotificationRuleMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifies the notification rule. | [readonly] 
**Type** | [**NotificationEventType**](NotificationEventType.md) | Notification rule type. | [readonly] 

## Methods

### NewNotificationRuleMeta

`func NewNotificationRuleMeta(id string, type_ NotificationEventType, ) *NotificationRuleMeta`

NewNotificationRuleMeta instantiates a new NotificationRuleMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationRuleMetaWithDefaults

`func NewNotificationRuleMetaWithDefaults() *NotificationRuleMeta`

NewNotificationRuleMetaWithDefaults instantiates a new NotificationRuleMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotificationRuleMeta) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationRuleMeta) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationRuleMeta) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *NotificationRuleMeta) GetType() NotificationEventType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationRuleMeta) GetTypeOk() (*NotificationEventType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationRuleMeta) SetType(v NotificationEventType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


