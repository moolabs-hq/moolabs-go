# CreateQuoteRedlineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuoteVersion** | **int32** |  | 
**Source** | **string** |  | 
**ProposedTerms** | Pointer to **map[string]interface{}** |  | [optional] 
**RedlineText** | **string** |  | 
**EvidenceMetadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCreateQuoteRedlineRequest

`func NewCreateQuoteRedlineRequest(quoteVersion int32, source string, redlineText string, ) *CreateQuoteRedlineRequest`

NewCreateQuoteRedlineRequest instantiates a new CreateQuoteRedlineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateQuoteRedlineRequestWithDefaults

`func NewCreateQuoteRedlineRequestWithDefaults() *CreateQuoteRedlineRequest`

NewCreateQuoteRedlineRequestWithDefaults instantiates a new CreateQuoteRedlineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuoteVersion

`func (o *CreateQuoteRedlineRequest) GetQuoteVersion() int32`

GetQuoteVersion returns the QuoteVersion field if non-nil, zero value otherwise.

### GetQuoteVersionOk

`func (o *CreateQuoteRedlineRequest) GetQuoteVersionOk() (*int32, bool)`

GetQuoteVersionOk returns a tuple with the QuoteVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteVersion

`func (o *CreateQuoteRedlineRequest) SetQuoteVersion(v int32)`

SetQuoteVersion sets QuoteVersion field to given value.


### GetSource

`func (o *CreateQuoteRedlineRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CreateQuoteRedlineRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CreateQuoteRedlineRequest) SetSource(v string)`

SetSource sets Source field to given value.


### GetProposedTerms

`func (o *CreateQuoteRedlineRequest) GetProposedTerms() map[string]interface{}`

GetProposedTerms returns the ProposedTerms field if non-nil, zero value otherwise.

### GetProposedTermsOk

`func (o *CreateQuoteRedlineRequest) GetProposedTermsOk() (*map[string]interface{}, bool)`

GetProposedTermsOk returns a tuple with the ProposedTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedTerms

`func (o *CreateQuoteRedlineRequest) SetProposedTerms(v map[string]interface{})`

SetProposedTerms sets ProposedTerms field to given value.

### HasProposedTerms

`func (o *CreateQuoteRedlineRequest) HasProposedTerms() bool`

HasProposedTerms returns a boolean if a field has been set.

### GetRedlineText

`func (o *CreateQuoteRedlineRequest) GetRedlineText() string`

GetRedlineText returns the RedlineText field if non-nil, zero value otherwise.

### GetRedlineTextOk

`func (o *CreateQuoteRedlineRequest) GetRedlineTextOk() (*string, bool)`

GetRedlineTextOk returns a tuple with the RedlineText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedlineText

`func (o *CreateQuoteRedlineRequest) SetRedlineText(v string)`

SetRedlineText sets RedlineText field to given value.


### GetEvidenceMetadata

`func (o *CreateQuoteRedlineRequest) GetEvidenceMetadata() map[string]interface{}`

GetEvidenceMetadata returns the EvidenceMetadata field if non-nil, zero value otherwise.

### GetEvidenceMetadataOk

`func (o *CreateQuoteRedlineRequest) GetEvidenceMetadataOk() (*map[string]interface{}, bool)`

GetEvidenceMetadataOk returns a tuple with the EvidenceMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceMetadata

`func (o *CreateQuoteRedlineRequest) SetEvidenceMetadata(v map[string]interface{})`

SetEvidenceMetadata sets EvidenceMetadata field to given value.

### HasEvidenceMetadata

`func (o *CreateQuoteRedlineRequest) HasEvidenceMetadata() bool`

HasEvidenceMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


