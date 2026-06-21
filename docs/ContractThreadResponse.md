# ContractThreadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**QuoteId** | **string** |  | 
**QuoteVersion** | Pointer to **int32** |  | [optional] 
**Status** | **string** |  | 
**Subject** | Pointer to **string** |  | [optional] 
**ReplyAlias** | **string** |  | 
**CreatedByActorId** | **string** |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**ClosedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewContractThreadResponse

`func NewContractThreadResponse(id string, tenantId string, quoteId string, status string, replyAlias string, createdByActorId string, ) *ContractThreadResponse`

NewContractThreadResponse instantiates a new ContractThreadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractThreadResponseWithDefaults

`func NewContractThreadResponseWithDefaults() *ContractThreadResponse`

NewContractThreadResponseWithDefaults instantiates a new ContractThreadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContractThreadResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContractThreadResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContractThreadResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *ContractThreadResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ContractThreadResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ContractThreadResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetQuoteId

`func (o *ContractThreadResponse) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *ContractThreadResponse) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *ContractThreadResponse) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.


### GetQuoteVersion

`func (o *ContractThreadResponse) GetQuoteVersion() int32`

GetQuoteVersion returns the QuoteVersion field if non-nil, zero value otherwise.

### GetQuoteVersionOk

`func (o *ContractThreadResponse) GetQuoteVersionOk() (*int32, bool)`

GetQuoteVersionOk returns a tuple with the QuoteVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteVersion

`func (o *ContractThreadResponse) SetQuoteVersion(v int32)`

SetQuoteVersion sets QuoteVersion field to given value.

### HasQuoteVersion

`func (o *ContractThreadResponse) HasQuoteVersion() bool`

HasQuoteVersion returns a boolean if a field has been set.

### GetStatus

`func (o *ContractThreadResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ContractThreadResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ContractThreadResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSubject

`func (o *ContractThreadResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ContractThreadResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ContractThreadResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *ContractThreadResponse) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetReplyAlias

`func (o *ContractThreadResponse) GetReplyAlias() string`

GetReplyAlias returns the ReplyAlias field if non-nil, zero value otherwise.

### GetReplyAliasOk

`func (o *ContractThreadResponse) GetReplyAliasOk() (*string, bool)`

GetReplyAliasOk returns a tuple with the ReplyAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyAlias

`func (o *ContractThreadResponse) SetReplyAlias(v string)`

SetReplyAlias sets ReplyAlias field to given value.


### GetCreatedByActorId

`func (o *ContractThreadResponse) GetCreatedByActorId() string`

GetCreatedByActorId returns the CreatedByActorId field if non-nil, zero value otherwise.

### GetCreatedByActorIdOk

`func (o *ContractThreadResponse) GetCreatedByActorIdOk() (*string, bool)`

GetCreatedByActorIdOk returns a tuple with the CreatedByActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByActorId

`func (o *ContractThreadResponse) SetCreatedByActorId(v string)`

SetCreatedByActorId sets CreatedByActorId field to given value.


### GetCreatedAt

`func (o *ContractThreadResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ContractThreadResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ContractThreadResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ContractThreadResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ContractThreadResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ContractThreadResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ContractThreadResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ContractThreadResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetClosedAt

`func (o *ContractThreadResponse) GetClosedAt() time.Time`

GetClosedAt returns the ClosedAt field if non-nil, zero value otherwise.

### GetClosedAtOk

`func (o *ContractThreadResponse) GetClosedAtOk() (*time.Time, bool)`

GetClosedAtOk returns a tuple with the ClosedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedAt

`func (o *ContractThreadResponse) SetClosedAt(v time.Time)`

SetClosedAt sets ClosedAt field to given value.

### HasClosedAt

`func (o *ContractThreadResponse) HasClosedAt() bool`

HasClosedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


