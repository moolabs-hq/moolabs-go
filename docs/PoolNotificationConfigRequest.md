# PoolNotificationConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationEmails** | **[]string** | List of email addresses to notify on pool-level alerts | 

## Methods

### NewPoolNotificationConfigRequest

`func NewPoolNotificationConfigRequest(notificationEmails []string, ) *PoolNotificationConfigRequest`

NewPoolNotificationConfigRequest instantiates a new PoolNotificationConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolNotificationConfigRequestWithDefaults

`func NewPoolNotificationConfigRequestWithDefaults() *PoolNotificationConfigRequest`

NewPoolNotificationConfigRequestWithDefaults instantiates a new PoolNotificationConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationEmails

`func (o *PoolNotificationConfigRequest) GetNotificationEmails() []string`

GetNotificationEmails returns the NotificationEmails field if non-nil, zero value otherwise.

### GetNotificationEmailsOk

`func (o *PoolNotificationConfigRequest) GetNotificationEmailsOk() (*[]string, bool)`

GetNotificationEmailsOk returns a tuple with the NotificationEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationEmails

`func (o *PoolNotificationConfigRequest) SetNotificationEmails(v []string)`

SetNotificationEmails sets NotificationEmails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


