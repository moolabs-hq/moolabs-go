# EvidenceCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EvidenceType** | **string** |  | 
**SourceTable** | Pointer to **string** |  | [optional] 
**SourceId** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**AmountMicros** | Pointer to **int32** |  | [optional] 

## Methods

### NewEvidenceCreateRequest

`func NewEvidenceCreateRequest(evidenceType string, ) *EvidenceCreateRequest`

NewEvidenceCreateRequest instantiates a new EvidenceCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvidenceCreateRequestWithDefaults

`func NewEvidenceCreateRequestWithDefaults() *EvidenceCreateRequest`

NewEvidenceCreateRequestWithDefaults instantiates a new EvidenceCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvidenceType

`func (o *EvidenceCreateRequest) GetEvidenceType() string`

GetEvidenceType returns the EvidenceType field if non-nil, zero value otherwise.

### GetEvidenceTypeOk

`func (o *EvidenceCreateRequest) GetEvidenceTypeOk() (*string, bool)`

GetEvidenceTypeOk returns a tuple with the EvidenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceType

`func (o *EvidenceCreateRequest) SetEvidenceType(v string)`

SetEvidenceType sets EvidenceType field to given value.


### GetSourceTable

`func (o *EvidenceCreateRequest) GetSourceTable() string`

GetSourceTable returns the SourceTable field if non-nil, zero value otherwise.

### GetSourceTableOk

`func (o *EvidenceCreateRequest) GetSourceTableOk() (*string, bool)`

GetSourceTableOk returns a tuple with the SourceTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTable

`func (o *EvidenceCreateRequest) SetSourceTable(v string)`

SetSourceTable sets SourceTable field to given value.

### HasSourceTable

`func (o *EvidenceCreateRequest) HasSourceTable() bool`

HasSourceTable returns a boolean if a field has been set.

### GetSourceId

`func (o *EvidenceCreateRequest) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *EvidenceCreateRequest) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *EvidenceCreateRequest) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *EvidenceCreateRequest) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetSummary

`func (o *EvidenceCreateRequest) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *EvidenceCreateRequest) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *EvidenceCreateRequest) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *EvidenceCreateRequest) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetAmountMicros

`func (o *EvidenceCreateRequest) GetAmountMicros() int32`

GetAmountMicros returns the AmountMicros field if non-nil, zero value otherwise.

### GetAmountMicrosOk

`func (o *EvidenceCreateRequest) GetAmountMicrosOk() (*int32, bool)`

GetAmountMicrosOk returns a tuple with the AmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMicros

`func (o *EvidenceCreateRequest) SetAmountMicros(v int32)`

SetAmountMicros sets AmountMicros field to given value.

### HasAmountMicros

`func (o *EvidenceCreateRequest) HasAmountMicros() bool`

HasAmountMicros returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


