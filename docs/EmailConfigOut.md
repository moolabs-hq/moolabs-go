# EmailConfigOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**SenderDomain** | **string** |  | 
**FromAddress** | **string** |  | 
**ResendDomainId** | **string** |  | 
**VerificationStatus** | **string** |  | 
**InboundSecretLast4** | **string** |  | 
**DnsRecords** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewEmailConfigOut

`func NewEmailConfigOut(tenantId string, senderDomain string, fromAddress string, resendDomainId string, verificationStatus string, inboundSecretLast4 string, ) *EmailConfigOut`

NewEmailConfigOut instantiates a new EmailConfigOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailConfigOutWithDefaults

`func NewEmailConfigOutWithDefaults() *EmailConfigOut`

NewEmailConfigOutWithDefaults instantiates a new EmailConfigOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *EmailConfigOut) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *EmailConfigOut) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *EmailConfigOut) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetSenderDomain

`func (o *EmailConfigOut) GetSenderDomain() string`

GetSenderDomain returns the SenderDomain field if non-nil, zero value otherwise.

### GetSenderDomainOk

`func (o *EmailConfigOut) GetSenderDomainOk() (*string, bool)`

GetSenderDomainOk returns a tuple with the SenderDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderDomain

`func (o *EmailConfigOut) SetSenderDomain(v string)`

SetSenderDomain sets SenderDomain field to given value.


### GetFromAddress

`func (o *EmailConfigOut) GetFromAddress() string`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *EmailConfigOut) GetFromAddressOk() (*string, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *EmailConfigOut) SetFromAddress(v string)`

SetFromAddress sets FromAddress field to given value.


### GetResendDomainId

`func (o *EmailConfigOut) GetResendDomainId() string`

GetResendDomainId returns the ResendDomainId field if non-nil, zero value otherwise.

### GetResendDomainIdOk

`func (o *EmailConfigOut) GetResendDomainIdOk() (*string, bool)`

GetResendDomainIdOk returns a tuple with the ResendDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResendDomainId

`func (o *EmailConfigOut) SetResendDomainId(v string)`

SetResendDomainId sets ResendDomainId field to given value.


### GetVerificationStatus

`func (o *EmailConfigOut) GetVerificationStatus() string`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *EmailConfigOut) GetVerificationStatusOk() (*string, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *EmailConfigOut) SetVerificationStatus(v string)`

SetVerificationStatus sets VerificationStatus field to given value.


### GetInboundSecretLast4

`func (o *EmailConfigOut) GetInboundSecretLast4() string`

GetInboundSecretLast4 returns the InboundSecretLast4 field if non-nil, zero value otherwise.

### GetInboundSecretLast4Ok

`func (o *EmailConfigOut) GetInboundSecretLast4Ok() (*string, bool)`

GetInboundSecretLast4Ok returns a tuple with the InboundSecretLast4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundSecretLast4

`func (o *EmailConfigOut) SetInboundSecretLast4(v string)`

SetInboundSecretLast4 sets InboundSecretLast4 field to given value.


### GetDnsRecords

`func (o *EmailConfigOut) GetDnsRecords() []map[string]interface{}`

GetDnsRecords returns the DnsRecords field if non-nil, zero value otherwise.

### GetDnsRecordsOk

`func (o *EmailConfigOut) GetDnsRecordsOk() (*[]map[string]interface{}, bool)`

GetDnsRecordsOk returns a tuple with the DnsRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRecords

`func (o *EmailConfigOut) SetDnsRecords(v []map[string]interface{})`

SetDnsRecords sets DnsRecords field to given value.

### HasDnsRecords

`func (o *EmailConfigOut) HasDnsRecords() bool`

HasDnsRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


