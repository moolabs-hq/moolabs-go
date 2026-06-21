# ContractInboundMessageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**ThreadId** | **string** |  | 
**ProviderMessageId** | **string** |  | 
**FromEmail** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**QuarantineReason** | Pointer to **string** |  | [optional] 
**SafeMetadata** | **map[string]interface{}** |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewContractInboundMessageResponse

`func NewContractInboundMessageResponse(id string, tenantId string, threadId string, providerMessageId string, status string, safeMetadata map[string]interface{}, ) *ContractInboundMessageResponse`

NewContractInboundMessageResponse instantiates a new ContractInboundMessageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractInboundMessageResponseWithDefaults

`func NewContractInboundMessageResponseWithDefaults() *ContractInboundMessageResponse`

NewContractInboundMessageResponseWithDefaults instantiates a new ContractInboundMessageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContractInboundMessageResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContractInboundMessageResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContractInboundMessageResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *ContractInboundMessageResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ContractInboundMessageResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ContractInboundMessageResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetThreadId

`func (o *ContractInboundMessageResponse) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *ContractInboundMessageResponse) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *ContractInboundMessageResponse) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.


### GetProviderMessageId

`func (o *ContractInboundMessageResponse) GetProviderMessageId() string`

GetProviderMessageId returns the ProviderMessageId field if non-nil, zero value otherwise.

### GetProviderMessageIdOk

`func (o *ContractInboundMessageResponse) GetProviderMessageIdOk() (*string, bool)`

GetProviderMessageIdOk returns a tuple with the ProviderMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderMessageId

`func (o *ContractInboundMessageResponse) SetProviderMessageId(v string)`

SetProviderMessageId sets ProviderMessageId field to given value.


### GetFromEmail

`func (o *ContractInboundMessageResponse) GetFromEmail() string`

GetFromEmail returns the FromEmail field if non-nil, zero value otherwise.

### GetFromEmailOk

`func (o *ContractInboundMessageResponse) GetFromEmailOk() (*string, bool)`

GetFromEmailOk returns a tuple with the FromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEmail

`func (o *ContractInboundMessageResponse) SetFromEmail(v string)`

SetFromEmail sets FromEmail field to given value.

### HasFromEmail

`func (o *ContractInboundMessageResponse) HasFromEmail() bool`

HasFromEmail returns a boolean if a field has been set.

### GetSubject

`func (o *ContractInboundMessageResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ContractInboundMessageResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ContractInboundMessageResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *ContractInboundMessageResponse) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetStatus

`func (o *ContractInboundMessageResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ContractInboundMessageResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ContractInboundMessageResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetQuarantineReason

`func (o *ContractInboundMessageResponse) GetQuarantineReason() string`

GetQuarantineReason returns the QuarantineReason field if non-nil, zero value otherwise.

### GetQuarantineReasonOk

`func (o *ContractInboundMessageResponse) GetQuarantineReasonOk() (*string, bool)`

GetQuarantineReasonOk returns a tuple with the QuarantineReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarantineReason

`func (o *ContractInboundMessageResponse) SetQuarantineReason(v string)`

SetQuarantineReason sets QuarantineReason field to given value.

### HasQuarantineReason

`func (o *ContractInboundMessageResponse) HasQuarantineReason() bool`

HasQuarantineReason returns a boolean if a field has been set.

### GetSafeMetadata

`func (o *ContractInboundMessageResponse) GetSafeMetadata() map[string]interface{}`

GetSafeMetadata returns the SafeMetadata field if non-nil, zero value otherwise.

### GetSafeMetadataOk

`func (o *ContractInboundMessageResponse) GetSafeMetadataOk() (*map[string]interface{}, bool)`

GetSafeMetadataOk returns a tuple with the SafeMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeMetadata

`func (o *ContractInboundMessageResponse) SetSafeMetadata(v map[string]interface{})`

SetSafeMetadata sets SafeMetadata field to given value.


### GetCreatedAt

`func (o *ContractInboundMessageResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ContractInboundMessageResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ContractInboundMessageResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ContractInboundMessageResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


