# InboundEmailPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recipient** | **string** |  | 
**FromEmail** | **string** |  | 
**MessageId** | **string** |  | 
**Subject** | Pointer to **string** |  | [optional] 
**Attachments** | Pointer to [**[]InboundEmailAttachment**](InboundEmailAttachment.md) |  | [optional] 

## Methods

### NewInboundEmailPayload

`func NewInboundEmailPayload(recipient string, fromEmail string, messageId string, ) *InboundEmailPayload`

NewInboundEmailPayload instantiates a new InboundEmailPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInboundEmailPayloadWithDefaults

`func NewInboundEmailPayloadWithDefaults() *InboundEmailPayload`

NewInboundEmailPayloadWithDefaults instantiates a new InboundEmailPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipient

`func (o *InboundEmailPayload) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *InboundEmailPayload) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *InboundEmailPayload) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.


### GetFromEmail

`func (o *InboundEmailPayload) GetFromEmail() string`

GetFromEmail returns the FromEmail field if non-nil, zero value otherwise.

### GetFromEmailOk

`func (o *InboundEmailPayload) GetFromEmailOk() (*string, bool)`

GetFromEmailOk returns a tuple with the FromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEmail

`func (o *InboundEmailPayload) SetFromEmail(v string)`

SetFromEmail sets FromEmail field to given value.


### GetMessageId

`func (o *InboundEmailPayload) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *InboundEmailPayload) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *InboundEmailPayload) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.


### GetSubject

`func (o *InboundEmailPayload) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *InboundEmailPayload) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *InboundEmailPayload) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *InboundEmailPayload) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetAttachments

`func (o *InboundEmailPayload) GetAttachments() []InboundEmailAttachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *InboundEmailPayload) GetAttachmentsOk() (*[]InboundEmailAttachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *InboundEmailPayload) SetAttachments(v []InboundEmailAttachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *InboundEmailPayload) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


