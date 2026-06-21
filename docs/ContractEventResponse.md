# ContractEventResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**ThreadId** | **string** |  | 
**QuoteId** | **string** |  | 
**EventType** | **string** |  | 
**ActorType** | **string** |  | 
**ActorId** | Pointer to **string** |  | [optional] 
**ParticipantId** | Pointer to **string** |  | [optional] 
**DocumentVersionId** | Pointer to **string** |  | [optional] 
**SafeMetadata** | **map[string]interface{}** |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewContractEventResponse

`func NewContractEventResponse(id string, tenantId string, threadId string, quoteId string, eventType string, actorType string, safeMetadata map[string]interface{}, ) *ContractEventResponse`

NewContractEventResponse instantiates a new ContractEventResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractEventResponseWithDefaults

`func NewContractEventResponseWithDefaults() *ContractEventResponse`

NewContractEventResponseWithDefaults instantiates a new ContractEventResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContractEventResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContractEventResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContractEventResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *ContractEventResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ContractEventResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ContractEventResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetThreadId

`func (o *ContractEventResponse) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *ContractEventResponse) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *ContractEventResponse) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.


### GetQuoteId

`func (o *ContractEventResponse) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *ContractEventResponse) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *ContractEventResponse) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.


### GetEventType

`func (o *ContractEventResponse) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ContractEventResponse) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ContractEventResponse) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetActorType

`func (o *ContractEventResponse) GetActorType() string`

GetActorType returns the ActorType field if non-nil, zero value otherwise.

### GetActorTypeOk

`func (o *ContractEventResponse) GetActorTypeOk() (*string, bool)`

GetActorTypeOk returns a tuple with the ActorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorType

`func (o *ContractEventResponse) SetActorType(v string)`

SetActorType sets ActorType field to given value.


### GetActorId

`func (o *ContractEventResponse) GetActorId() string`

GetActorId returns the ActorId field if non-nil, zero value otherwise.

### GetActorIdOk

`func (o *ContractEventResponse) GetActorIdOk() (*string, bool)`

GetActorIdOk returns a tuple with the ActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorId

`func (o *ContractEventResponse) SetActorId(v string)`

SetActorId sets ActorId field to given value.

### HasActorId

`func (o *ContractEventResponse) HasActorId() bool`

HasActorId returns a boolean if a field has been set.

### GetParticipantId

`func (o *ContractEventResponse) GetParticipantId() string`

GetParticipantId returns the ParticipantId field if non-nil, zero value otherwise.

### GetParticipantIdOk

`func (o *ContractEventResponse) GetParticipantIdOk() (*string, bool)`

GetParticipantIdOk returns a tuple with the ParticipantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantId

`func (o *ContractEventResponse) SetParticipantId(v string)`

SetParticipantId sets ParticipantId field to given value.

### HasParticipantId

`func (o *ContractEventResponse) HasParticipantId() bool`

HasParticipantId returns a boolean if a field has been set.

### GetDocumentVersionId

`func (o *ContractEventResponse) GetDocumentVersionId() string`

GetDocumentVersionId returns the DocumentVersionId field if non-nil, zero value otherwise.

### GetDocumentVersionIdOk

`func (o *ContractEventResponse) GetDocumentVersionIdOk() (*string, bool)`

GetDocumentVersionIdOk returns a tuple with the DocumentVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentVersionId

`func (o *ContractEventResponse) SetDocumentVersionId(v string)`

SetDocumentVersionId sets DocumentVersionId field to given value.

### HasDocumentVersionId

`func (o *ContractEventResponse) HasDocumentVersionId() bool`

HasDocumentVersionId returns a boolean if a field has been set.

### GetSafeMetadata

`func (o *ContractEventResponse) GetSafeMetadata() map[string]interface{}`

GetSafeMetadata returns the SafeMetadata field if non-nil, zero value otherwise.

### GetSafeMetadataOk

`func (o *ContractEventResponse) GetSafeMetadataOk() (*map[string]interface{}, bool)`

GetSafeMetadataOk returns a tuple with the SafeMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeMetadata

`func (o *ContractEventResponse) SetSafeMetadata(v map[string]interface{})`

SetSafeMetadata sets SafeMetadata field to given value.


### GetCreatedAt

`func (o *ContractEventResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ContractEventResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ContractEventResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ContractEventResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


