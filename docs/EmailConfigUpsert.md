# EmailConfigUpsert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SenderDomain** | **string** |  | 
**FromAddress** | Pointer to **string** |  | [optional] 
**ReplyDomain** | Pointer to **string** |  | [optional] 
**InternalContactDomains** | Pointer to **[]string** |  | [optional] 
**InboundSecret** | Pointer to **string** |  | [optional] 

## Methods

### NewEmailConfigUpsert

`func NewEmailConfigUpsert(senderDomain string, ) *EmailConfigUpsert`

NewEmailConfigUpsert instantiates a new EmailConfigUpsert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailConfigUpsertWithDefaults

`func NewEmailConfigUpsertWithDefaults() *EmailConfigUpsert`

NewEmailConfigUpsertWithDefaults instantiates a new EmailConfigUpsert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSenderDomain

`func (o *EmailConfigUpsert) GetSenderDomain() string`

GetSenderDomain returns the SenderDomain field if non-nil, zero value otherwise.

### GetSenderDomainOk

`func (o *EmailConfigUpsert) GetSenderDomainOk() (*string, bool)`

GetSenderDomainOk returns a tuple with the SenderDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderDomain

`func (o *EmailConfigUpsert) SetSenderDomain(v string)`

SetSenderDomain sets SenderDomain field to given value.


### GetFromAddress

`func (o *EmailConfigUpsert) GetFromAddress() string`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *EmailConfigUpsert) GetFromAddressOk() (*string, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *EmailConfigUpsert) SetFromAddress(v string)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *EmailConfigUpsert) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetReplyDomain

`func (o *EmailConfigUpsert) GetReplyDomain() string`

GetReplyDomain returns the ReplyDomain field if non-nil, zero value otherwise.

### GetReplyDomainOk

`func (o *EmailConfigUpsert) GetReplyDomainOk() (*string, bool)`

GetReplyDomainOk returns a tuple with the ReplyDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyDomain

`func (o *EmailConfigUpsert) SetReplyDomain(v string)`

SetReplyDomain sets ReplyDomain field to given value.

### HasReplyDomain

`func (o *EmailConfigUpsert) HasReplyDomain() bool`

HasReplyDomain returns a boolean if a field has been set.

### GetInternalContactDomains

`func (o *EmailConfigUpsert) GetInternalContactDomains() []string`

GetInternalContactDomains returns the InternalContactDomains field if non-nil, zero value otherwise.

### GetInternalContactDomainsOk

`func (o *EmailConfigUpsert) GetInternalContactDomainsOk() (*[]string, bool)`

GetInternalContactDomainsOk returns a tuple with the InternalContactDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalContactDomains

`func (o *EmailConfigUpsert) SetInternalContactDomains(v []string)`

SetInternalContactDomains sets InternalContactDomains field to given value.

### HasInternalContactDomains

`func (o *EmailConfigUpsert) HasInternalContactDomains() bool`

HasInternalContactDomains returns a boolean if a field has been set.

### GetInboundSecret

`func (o *EmailConfigUpsert) GetInboundSecret() string`

GetInboundSecret returns the InboundSecret field if non-nil, zero value otherwise.

### GetInboundSecretOk

`func (o *EmailConfigUpsert) GetInboundSecretOk() (*string, bool)`

GetInboundSecretOk returns a tuple with the InboundSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundSecret

`func (o *EmailConfigUpsert) SetInboundSecret(v string)`

SetInboundSecret sets InboundSecret field to given value.

### HasInboundSecret

`func (o *EmailConfigUpsert) HasInboundSecret() bool`

HasInboundSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


