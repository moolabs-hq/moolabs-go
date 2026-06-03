# CrmQuoteDraftRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | **string** |  | 
**RecordId** | **string** |  | 
**QuoteId** | **string** |  | 
**QuoteVersion** | **int32** |  | 
**CandidateLines** | Pointer to **[]map[string]interface{}** |  | [optional] 
**IdempotencyKey** | Pointer to **string** |  | [optional] 

## Methods

### NewCrmQuoteDraftRequest

`func NewCrmQuoteDraftRequest(recordType string, recordId string, quoteId string, quoteVersion int32, ) *CrmQuoteDraftRequest`

NewCrmQuoteDraftRequest instantiates a new CrmQuoteDraftRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrmQuoteDraftRequestWithDefaults

`func NewCrmQuoteDraftRequestWithDefaults() *CrmQuoteDraftRequest`

NewCrmQuoteDraftRequestWithDefaults instantiates a new CrmQuoteDraftRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *CrmQuoteDraftRequest) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *CrmQuoteDraftRequest) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *CrmQuoteDraftRequest) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.


### GetRecordId

`func (o *CrmQuoteDraftRequest) GetRecordId() string`

GetRecordId returns the RecordId field if non-nil, zero value otherwise.

### GetRecordIdOk

`func (o *CrmQuoteDraftRequest) GetRecordIdOk() (*string, bool)`

GetRecordIdOk returns a tuple with the RecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordId

`func (o *CrmQuoteDraftRequest) SetRecordId(v string)`

SetRecordId sets RecordId field to given value.


### GetQuoteId

`func (o *CrmQuoteDraftRequest) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *CrmQuoteDraftRequest) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *CrmQuoteDraftRequest) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.


### GetQuoteVersion

`func (o *CrmQuoteDraftRequest) GetQuoteVersion() int32`

GetQuoteVersion returns the QuoteVersion field if non-nil, zero value otherwise.

### GetQuoteVersionOk

`func (o *CrmQuoteDraftRequest) GetQuoteVersionOk() (*int32, bool)`

GetQuoteVersionOk returns a tuple with the QuoteVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteVersion

`func (o *CrmQuoteDraftRequest) SetQuoteVersion(v int32)`

SetQuoteVersion sets QuoteVersion field to given value.


### GetCandidateLines

`func (o *CrmQuoteDraftRequest) GetCandidateLines() []map[string]interface{}`

GetCandidateLines returns the CandidateLines field if non-nil, zero value otherwise.

### GetCandidateLinesOk

`func (o *CrmQuoteDraftRequest) GetCandidateLinesOk() (*[]map[string]interface{}, bool)`

GetCandidateLinesOk returns a tuple with the CandidateLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidateLines

`func (o *CrmQuoteDraftRequest) SetCandidateLines(v []map[string]interface{})`

SetCandidateLines sets CandidateLines field to given value.

### HasCandidateLines

`func (o *CrmQuoteDraftRequest) HasCandidateLines() bool`

HasCandidateLines returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *CrmQuoteDraftRequest) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *CrmQuoteDraftRequest) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *CrmQuoteDraftRequest) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *CrmQuoteDraftRequest) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


