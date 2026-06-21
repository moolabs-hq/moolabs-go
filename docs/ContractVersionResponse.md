# ContractVersionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**ThreadId** | **string** |  | 
**DocumentId** | **string** |  | 
**QuoteId** | **string** |  | 
**VersionNo** | **int32** |  | 
**Source** | **string** |  | 
**ParticipantId** | Pointer to **string** |  | [optional] 
**InboundMessageId** | Pointer to **string** |  | [optional] 
**Filename** | **string** |  | 
**ContentType** | **string** |  | 
**SizeBytes** | **int32** |  | 
**ArtifactHash** | **string** |  | 
**ScanStatus** | **string** |  | 
**ExtractedText** | **string** |  | 
**Structure** | **map[string]interface{}** |  | 
**ProviderMessageId** | Pointer to **string** |  | [optional] 
**EmailSubject** | Pointer to **string** |  | [optional] 
**ReceivedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedByActorId** | **string** |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewContractVersionResponse

`func NewContractVersionResponse(id string, tenantId string, threadId string, documentId string, quoteId string, versionNo int32, source string, filename string, contentType string, sizeBytes int32, artifactHash string, scanStatus string, extractedText string, structure map[string]interface{}, createdByActorId string, ) *ContractVersionResponse`

NewContractVersionResponse instantiates a new ContractVersionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractVersionResponseWithDefaults

`func NewContractVersionResponseWithDefaults() *ContractVersionResponse`

NewContractVersionResponseWithDefaults instantiates a new ContractVersionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContractVersionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContractVersionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContractVersionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *ContractVersionResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ContractVersionResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ContractVersionResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetThreadId

`func (o *ContractVersionResponse) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *ContractVersionResponse) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *ContractVersionResponse) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.


### GetDocumentId

`func (o *ContractVersionResponse) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *ContractVersionResponse) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *ContractVersionResponse) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.


### GetQuoteId

`func (o *ContractVersionResponse) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *ContractVersionResponse) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *ContractVersionResponse) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.


### GetVersionNo

`func (o *ContractVersionResponse) GetVersionNo() int32`

GetVersionNo returns the VersionNo field if non-nil, zero value otherwise.

### GetVersionNoOk

`func (o *ContractVersionResponse) GetVersionNoOk() (*int32, bool)`

GetVersionNoOk returns a tuple with the VersionNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionNo

`func (o *ContractVersionResponse) SetVersionNo(v int32)`

SetVersionNo sets VersionNo field to given value.


### GetSource

`func (o *ContractVersionResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ContractVersionResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ContractVersionResponse) SetSource(v string)`

SetSource sets Source field to given value.


### GetParticipantId

`func (o *ContractVersionResponse) GetParticipantId() string`

GetParticipantId returns the ParticipantId field if non-nil, zero value otherwise.

### GetParticipantIdOk

`func (o *ContractVersionResponse) GetParticipantIdOk() (*string, bool)`

GetParticipantIdOk returns a tuple with the ParticipantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantId

`func (o *ContractVersionResponse) SetParticipantId(v string)`

SetParticipantId sets ParticipantId field to given value.

### HasParticipantId

`func (o *ContractVersionResponse) HasParticipantId() bool`

HasParticipantId returns a boolean if a field has been set.

### GetInboundMessageId

`func (o *ContractVersionResponse) GetInboundMessageId() string`

GetInboundMessageId returns the InboundMessageId field if non-nil, zero value otherwise.

### GetInboundMessageIdOk

`func (o *ContractVersionResponse) GetInboundMessageIdOk() (*string, bool)`

GetInboundMessageIdOk returns a tuple with the InboundMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundMessageId

`func (o *ContractVersionResponse) SetInboundMessageId(v string)`

SetInboundMessageId sets InboundMessageId field to given value.

### HasInboundMessageId

`func (o *ContractVersionResponse) HasInboundMessageId() bool`

HasInboundMessageId returns a boolean if a field has been set.

### GetFilename

`func (o *ContractVersionResponse) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *ContractVersionResponse) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *ContractVersionResponse) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetContentType

`func (o *ContractVersionResponse) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ContractVersionResponse) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ContractVersionResponse) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetSizeBytes

`func (o *ContractVersionResponse) GetSizeBytes() int32`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *ContractVersionResponse) GetSizeBytesOk() (*int32, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *ContractVersionResponse) SetSizeBytes(v int32)`

SetSizeBytes sets SizeBytes field to given value.


### GetArtifactHash

`func (o *ContractVersionResponse) GetArtifactHash() string`

GetArtifactHash returns the ArtifactHash field if non-nil, zero value otherwise.

### GetArtifactHashOk

`func (o *ContractVersionResponse) GetArtifactHashOk() (*string, bool)`

GetArtifactHashOk returns a tuple with the ArtifactHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactHash

`func (o *ContractVersionResponse) SetArtifactHash(v string)`

SetArtifactHash sets ArtifactHash field to given value.


### GetScanStatus

`func (o *ContractVersionResponse) GetScanStatus() string`

GetScanStatus returns the ScanStatus field if non-nil, zero value otherwise.

### GetScanStatusOk

`func (o *ContractVersionResponse) GetScanStatusOk() (*string, bool)`

GetScanStatusOk returns a tuple with the ScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanStatus

`func (o *ContractVersionResponse) SetScanStatus(v string)`

SetScanStatus sets ScanStatus field to given value.


### GetExtractedText

`func (o *ContractVersionResponse) GetExtractedText() string`

GetExtractedText returns the ExtractedText field if non-nil, zero value otherwise.

### GetExtractedTextOk

`func (o *ContractVersionResponse) GetExtractedTextOk() (*string, bool)`

GetExtractedTextOk returns a tuple with the ExtractedText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractedText

`func (o *ContractVersionResponse) SetExtractedText(v string)`

SetExtractedText sets ExtractedText field to given value.


### GetStructure

`func (o *ContractVersionResponse) GetStructure() map[string]interface{}`

GetStructure returns the Structure field if non-nil, zero value otherwise.

### GetStructureOk

`func (o *ContractVersionResponse) GetStructureOk() (*map[string]interface{}, bool)`

GetStructureOk returns a tuple with the Structure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructure

`func (o *ContractVersionResponse) SetStructure(v map[string]interface{})`

SetStructure sets Structure field to given value.


### GetProviderMessageId

`func (o *ContractVersionResponse) GetProviderMessageId() string`

GetProviderMessageId returns the ProviderMessageId field if non-nil, zero value otherwise.

### GetProviderMessageIdOk

`func (o *ContractVersionResponse) GetProviderMessageIdOk() (*string, bool)`

GetProviderMessageIdOk returns a tuple with the ProviderMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderMessageId

`func (o *ContractVersionResponse) SetProviderMessageId(v string)`

SetProviderMessageId sets ProviderMessageId field to given value.

### HasProviderMessageId

`func (o *ContractVersionResponse) HasProviderMessageId() bool`

HasProviderMessageId returns a boolean if a field has been set.

### GetEmailSubject

`func (o *ContractVersionResponse) GetEmailSubject() string`

GetEmailSubject returns the EmailSubject field if non-nil, zero value otherwise.

### GetEmailSubjectOk

`func (o *ContractVersionResponse) GetEmailSubjectOk() (*string, bool)`

GetEmailSubjectOk returns a tuple with the EmailSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubject

`func (o *ContractVersionResponse) SetEmailSubject(v string)`

SetEmailSubject sets EmailSubject field to given value.

### HasEmailSubject

`func (o *ContractVersionResponse) HasEmailSubject() bool`

HasEmailSubject returns a boolean if a field has been set.

### GetReceivedAt

`func (o *ContractVersionResponse) GetReceivedAt() time.Time`

GetReceivedAt returns the ReceivedAt field if non-nil, zero value otherwise.

### GetReceivedAtOk

`func (o *ContractVersionResponse) GetReceivedAtOk() (*time.Time, bool)`

GetReceivedAtOk returns a tuple with the ReceivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedAt

`func (o *ContractVersionResponse) SetReceivedAt(v time.Time)`

SetReceivedAt sets ReceivedAt field to given value.

### HasReceivedAt

`func (o *ContractVersionResponse) HasReceivedAt() bool`

HasReceivedAt returns a boolean if a field has been set.

### GetCreatedByActorId

`func (o *ContractVersionResponse) GetCreatedByActorId() string`

GetCreatedByActorId returns the CreatedByActorId field if non-nil, zero value otherwise.

### GetCreatedByActorIdOk

`func (o *ContractVersionResponse) GetCreatedByActorIdOk() (*string, bool)`

GetCreatedByActorIdOk returns a tuple with the CreatedByActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByActorId

`func (o *ContractVersionResponse) SetCreatedByActorId(v string)`

SetCreatedByActorId sets CreatedByActorId field to given value.


### GetCreatedAt

`func (o *ContractVersionResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ContractVersionResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ContractVersionResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ContractVersionResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


