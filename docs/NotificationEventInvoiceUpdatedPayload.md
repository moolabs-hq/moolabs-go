# NotificationEventInvoiceUpdatedPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier for the notification event the payload belongs to. | [readonly] 
**Type** | **string** | Type of the notification event. | [readonly] 
**Timestamp** | **time.Time** | Timestamp when the notification event was created in RFC 3339 format. | [readonly] 
**Data** | [**Invoice**](Invoice.md) | The data of the payload. | [readonly] 

## Methods

### NewNotificationEventInvoiceUpdatedPayload

`func NewNotificationEventInvoiceUpdatedPayload(id string, type_ string, timestamp time.Time, data Invoice, ) *NotificationEventInvoiceUpdatedPayload`

NewNotificationEventInvoiceUpdatedPayload instantiates a new NotificationEventInvoiceUpdatedPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationEventInvoiceUpdatedPayloadWithDefaults

`func NewNotificationEventInvoiceUpdatedPayloadWithDefaults() *NotificationEventInvoiceUpdatedPayload`

NewNotificationEventInvoiceUpdatedPayloadWithDefaults instantiates a new NotificationEventInvoiceUpdatedPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotificationEventInvoiceUpdatedPayload) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationEventInvoiceUpdatedPayload) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationEventInvoiceUpdatedPayload) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *NotificationEventInvoiceUpdatedPayload) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationEventInvoiceUpdatedPayload) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationEventInvoiceUpdatedPayload) SetType(v string)`

SetType sets Type field to given value.


### GetTimestamp

`func (o *NotificationEventInvoiceUpdatedPayload) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *NotificationEventInvoiceUpdatedPayload) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *NotificationEventInvoiceUpdatedPayload) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetData

`func (o *NotificationEventInvoiceUpdatedPayload) GetData() Invoice`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NotificationEventInvoiceUpdatedPayload) GetDataOk() (*Invoice, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NotificationEventInvoiceUpdatedPayload) SetData(v Invoice)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


