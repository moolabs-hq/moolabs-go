# NotificationEventBalanceThresholdPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier for the notification event the payload belongs to. | [readonly] 
**Type** | **string** | Type of the notification event. | [readonly] 
**Timestamp** | **time.Time** | Timestamp when the notification event was created in RFC 3339 format. | [readonly] 
**Data** | [**NotificationEventBalanceThresholdPayloadData**](NotificationEventBalanceThresholdPayloadData.md) | The data of the payload. | [readonly] 

## Methods

### NewNotificationEventBalanceThresholdPayload

`func NewNotificationEventBalanceThresholdPayload(id string, type_ string, timestamp time.Time, data NotificationEventBalanceThresholdPayloadData, ) *NotificationEventBalanceThresholdPayload`

NewNotificationEventBalanceThresholdPayload instantiates a new NotificationEventBalanceThresholdPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationEventBalanceThresholdPayloadWithDefaults

`func NewNotificationEventBalanceThresholdPayloadWithDefaults() *NotificationEventBalanceThresholdPayload`

NewNotificationEventBalanceThresholdPayloadWithDefaults instantiates a new NotificationEventBalanceThresholdPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotificationEventBalanceThresholdPayload) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationEventBalanceThresholdPayload) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationEventBalanceThresholdPayload) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *NotificationEventBalanceThresholdPayload) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationEventBalanceThresholdPayload) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationEventBalanceThresholdPayload) SetType(v string)`

SetType sets Type field to given value.


### GetTimestamp

`func (o *NotificationEventBalanceThresholdPayload) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *NotificationEventBalanceThresholdPayload) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *NotificationEventBalanceThresholdPayload) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetData

`func (o *NotificationEventBalanceThresholdPayload) GetData() NotificationEventBalanceThresholdPayloadData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NotificationEventBalanceThresholdPayload) GetDataOk() (*NotificationEventBalanceThresholdPayloadData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NotificationEventBalanceThresholdPayload) SetData(v NotificationEventBalanceThresholdPayloadData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


