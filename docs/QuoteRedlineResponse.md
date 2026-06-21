# QuoteRedlineResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**QuoteId** | **string** |  | 
**QuoteVersion** | **int32** |  | 
**Source** | **string** |  | 
**Status** | **string** |  | 
**ProposedTerms** | **map[string]interface{}** |  | 
**RedlineText** | **string** |  | 
**EvidenceMetadata** | **map[string]interface{}** |  | 
**CreatedByActorType** | **string** |  | 
**CreatedByActorId** | **string** |  | 
**AcceptedByActorId** | Pointer to **string** |  | [optional] 
**AcceptedAt** | Pointer to **string** |  | [optional] 
**ResultingQuoteVersion** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**ClauseFamily** | Pointer to **string** |  | [optional] 
**Verdict** | Pointer to **string** |  | [optional] 
**MatchedTier** | Pointer to **string** |  | [optional] 
**CitedSpan** | Pointer to **map[string]interface{}** |  | [optional] 
**Confidence** | Pointer to **float32** |  | [optional] 

## Methods

### NewQuoteRedlineResponse

`func NewQuoteRedlineResponse(id string, tenantId string, quoteId string, quoteVersion int32, source string, status string, proposedTerms map[string]interface{}, redlineText string, evidenceMetadata map[string]interface{}, createdByActorType string, createdByActorId string, ) *QuoteRedlineResponse`

NewQuoteRedlineResponse instantiates a new QuoteRedlineResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteRedlineResponseWithDefaults

`func NewQuoteRedlineResponseWithDefaults() *QuoteRedlineResponse`

NewQuoteRedlineResponseWithDefaults instantiates a new QuoteRedlineResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuoteRedlineResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuoteRedlineResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuoteRedlineResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *QuoteRedlineResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *QuoteRedlineResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *QuoteRedlineResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetQuoteId

`func (o *QuoteRedlineResponse) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *QuoteRedlineResponse) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *QuoteRedlineResponse) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.


### GetQuoteVersion

`func (o *QuoteRedlineResponse) GetQuoteVersion() int32`

GetQuoteVersion returns the QuoteVersion field if non-nil, zero value otherwise.

### GetQuoteVersionOk

`func (o *QuoteRedlineResponse) GetQuoteVersionOk() (*int32, bool)`

GetQuoteVersionOk returns a tuple with the QuoteVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteVersion

`func (o *QuoteRedlineResponse) SetQuoteVersion(v int32)`

SetQuoteVersion sets QuoteVersion field to given value.


### GetSource

`func (o *QuoteRedlineResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *QuoteRedlineResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *QuoteRedlineResponse) SetSource(v string)`

SetSource sets Source field to given value.


### GetStatus

`func (o *QuoteRedlineResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QuoteRedlineResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QuoteRedlineResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetProposedTerms

`func (o *QuoteRedlineResponse) GetProposedTerms() map[string]interface{}`

GetProposedTerms returns the ProposedTerms field if non-nil, zero value otherwise.

### GetProposedTermsOk

`func (o *QuoteRedlineResponse) GetProposedTermsOk() (*map[string]interface{}, bool)`

GetProposedTermsOk returns a tuple with the ProposedTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedTerms

`func (o *QuoteRedlineResponse) SetProposedTerms(v map[string]interface{})`

SetProposedTerms sets ProposedTerms field to given value.


### GetRedlineText

`func (o *QuoteRedlineResponse) GetRedlineText() string`

GetRedlineText returns the RedlineText field if non-nil, zero value otherwise.

### GetRedlineTextOk

`func (o *QuoteRedlineResponse) GetRedlineTextOk() (*string, bool)`

GetRedlineTextOk returns a tuple with the RedlineText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedlineText

`func (o *QuoteRedlineResponse) SetRedlineText(v string)`

SetRedlineText sets RedlineText field to given value.


### GetEvidenceMetadata

`func (o *QuoteRedlineResponse) GetEvidenceMetadata() map[string]interface{}`

GetEvidenceMetadata returns the EvidenceMetadata field if non-nil, zero value otherwise.

### GetEvidenceMetadataOk

`func (o *QuoteRedlineResponse) GetEvidenceMetadataOk() (*map[string]interface{}, bool)`

GetEvidenceMetadataOk returns a tuple with the EvidenceMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceMetadata

`func (o *QuoteRedlineResponse) SetEvidenceMetadata(v map[string]interface{})`

SetEvidenceMetadata sets EvidenceMetadata field to given value.


### GetCreatedByActorType

`func (o *QuoteRedlineResponse) GetCreatedByActorType() string`

GetCreatedByActorType returns the CreatedByActorType field if non-nil, zero value otherwise.

### GetCreatedByActorTypeOk

`func (o *QuoteRedlineResponse) GetCreatedByActorTypeOk() (*string, bool)`

GetCreatedByActorTypeOk returns a tuple with the CreatedByActorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByActorType

`func (o *QuoteRedlineResponse) SetCreatedByActorType(v string)`

SetCreatedByActorType sets CreatedByActorType field to given value.


### GetCreatedByActorId

`func (o *QuoteRedlineResponse) GetCreatedByActorId() string`

GetCreatedByActorId returns the CreatedByActorId field if non-nil, zero value otherwise.

### GetCreatedByActorIdOk

`func (o *QuoteRedlineResponse) GetCreatedByActorIdOk() (*string, bool)`

GetCreatedByActorIdOk returns a tuple with the CreatedByActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByActorId

`func (o *QuoteRedlineResponse) SetCreatedByActorId(v string)`

SetCreatedByActorId sets CreatedByActorId field to given value.


### GetAcceptedByActorId

`func (o *QuoteRedlineResponse) GetAcceptedByActorId() string`

GetAcceptedByActorId returns the AcceptedByActorId field if non-nil, zero value otherwise.

### GetAcceptedByActorIdOk

`func (o *QuoteRedlineResponse) GetAcceptedByActorIdOk() (*string, bool)`

GetAcceptedByActorIdOk returns a tuple with the AcceptedByActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedByActorId

`func (o *QuoteRedlineResponse) SetAcceptedByActorId(v string)`

SetAcceptedByActorId sets AcceptedByActorId field to given value.

### HasAcceptedByActorId

`func (o *QuoteRedlineResponse) HasAcceptedByActorId() bool`

HasAcceptedByActorId returns a boolean if a field has been set.

### GetAcceptedAt

`func (o *QuoteRedlineResponse) GetAcceptedAt() string`

GetAcceptedAt returns the AcceptedAt field if non-nil, zero value otherwise.

### GetAcceptedAtOk

`func (o *QuoteRedlineResponse) GetAcceptedAtOk() (*string, bool)`

GetAcceptedAtOk returns a tuple with the AcceptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedAt

`func (o *QuoteRedlineResponse) SetAcceptedAt(v string)`

SetAcceptedAt sets AcceptedAt field to given value.

### HasAcceptedAt

`func (o *QuoteRedlineResponse) HasAcceptedAt() bool`

HasAcceptedAt returns a boolean if a field has been set.

### GetResultingQuoteVersion

`func (o *QuoteRedlineResponse) GetResultingQuoteVersion() int32`

GetResultingQuoteVersion returns the ResultingQuoteVersion field if non-nil, zero value otherwise.

### GetResultingQuoteVersionOk

`func (o *QuoteRedlineResponse) GetResultingQuoteVersionOk() (*int32, bool)`

GetResultingQuoteVersionOk returns a tuple with the ResultingQuoteVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultingQuoteVersion

`func (o *QuoteRedlineResponse) SetResultingQuoteVersion(v int32)`

SetResultingQuoteVersion sets ResultingQuoteVersion field to given value.

### HasResultingQuoteVersion

`func (o *QuoteRedlineResponse) HasResultingQuoteVersion() bool`

HasResultingQuoteVersion returns a boolean if a field has been set.

### GetCreatedAt

`func (o *QuoteRedlineResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QuoteRedlineResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QuoteRedlineResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *QuoteRedlineResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetClauseFamily

`func (o *QuoteRedlineResponse) GetClauseFamily() string`

GetClauseFamily returns the ClauseFamily field if non-nil, zero value otherwise.

### GetClauseFamilyOk

`func (o *QuoteRedlineResponse) GetClauseFamilyOk() (*string, bool)`

GetClauseFamilyOk returns a tuple with the ClauseFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClauseFamily

`func (o *QuoteRedlineResponse) SetClauseFamily(v string)`

SetClauseFamily sets ClauseFamily field to given value.

### HasClauseFamily

`func (o *QuoteRedlineResponse) HasClauseFamily() bool`

HasClauseFamily returns a boolean if a field has been set.

### GetVerdict

`func (o *QuoteRedlineResponse) GetVerdict() string`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *QuoteRedlineResponse) GetVerdictOk() (*string, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *QuoteRedlineResponse) SetVerdict(v string)`

SetVerdict sets Verdict field to given value.

### HasVerdict

`func (o *QuoteRedlineResponse) HasVerdict() bool`

HasVerdict returns a boolean if a field has been set.

### GetMatchedTier

`func (o *QuoteRedlineResponse) GetMatchedTier() string`

GetMatchedTier returns the MatchedTier field if non-nil, zero value otherwise.

### GetMatchedTierOk

`func (o *QuoteRedlineResponse) GetMatchedTierOk() (*string, bool)`

GetMatchedTierOk returns a tuple with the MatchedTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedTier

`func (o *QuoteRedlineResponse) SetMatchedTier(v string)`

SetMatchedTier sets MatchedTier field to given value.

### HasMatchedTier

`func (o *QuoteRedlineResponse) HasMatchedTier() bool`

HasMatchedTier returns a boolean if a field has been set.

### GetCitedSpan

`func (o *QuoteRedlineResponse) GetCitedSpan() map[string]interface{}`

GetCitedSpan returns the CitedSpan field if non-nil, zero value otherwise.

### GetCitedSpanOk

`func (o *QuoteRedlineResponse) GetCitedSpanOk() (*map[string]interface{}, bool)`

GetCitedSpanOk returns a tuple with the CitedSpan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitedSpan

`func (o *QuoteRedlineResponse) SetCitedSpan(v map[string]interface{})`

SetCitedSpan sets CitedSpan field to given value.

### HasCitedSpan

`func (o *QuoteRedlineResponse) HasCitedSpan() bool`

HasCitedSpan returns a boolean if a field has been set.

### GetConfidence

`func (o *QuoteRedlineResponse) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *QuoteRedlineResponse) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *QuoteRedlineResponse) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *QuoteRedlineResponse) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


