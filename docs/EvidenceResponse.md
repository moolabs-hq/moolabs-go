# EvidenceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**DisputeId** | **string** |  | 
**EvidenceType** | Pointer to **string** |  | [optional] 
**SourceTable** | **string** |  | 
**SourceId** | **string** |  | 
**Summary** | Pointer to **string** |  | [optional] 
**AmountMicros** | Pointer to **int32** |  | [optional] 
**ResolvedAmountMicros** | Pointer to **int32** |  | [optional] 
**ResolutionStatus** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewEvidenceResponse

`func NewEvidenceResponse(id string, disputeId string, sourceTable string, sourceId string, ) *EvidenceResponse`

NewEvidenceResponse instantiates a new EvidenceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvidenceResponseWithDefaults

`func NewEvidenceResponseWithDefaults() *EvidenceResponse`

NewEvidenceResponseWithDefaults instantiates a new EvidenceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EvidenceResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EvidenceResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EvidenceResponse) SetId(v string)`

SetId sets Id field to given value.


### GetDisputeId

`func (o *EvidenceResponse) GetDisputeId() string`

GetDisputeId returns the DisputeId field if non-nil, zero value otherwise.

### GetDisputeIdOk

`func (o *EvidenceResponse) GetDisputeIdOk() (*string, bool)`

GetDisputeIdOk returns a tuple with the DisputeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputeId

`func (o *EvidenceResponse) SetDisputeId(v string)`

SetDisputeId sets DisputeId field to given value.


### GetEvidenceType

`func (o *EvidenceResponse) GetEvidenceType() string`

GetEvidenceType returns the EvidenceType field if non-nil, zero value otherwise.

### GetEvidenceTypeOk

`func (o *EvidenceResponse) GetEvidenceTypeOk() (*string, bool)`

GetEvidenceTypeOk returns a tuple with the EvidenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceType

`func (o *EvidenceResponse) SetEvidenceType(v string)`

SetEvidenceType sets EvidenceType field to given value.

### HasEvidenceType

`func (o *EvidenceResponse) HasEvidenceType() bool`

HasEvidenceType returns a boolean if a field has been set.

### GetSourceTable

`func (o *EvidenceResponse) GetSourceTable() string`

GetSourceTable returns the SourceTable field if non-nil, zero value otherwise.

### GetSourceTableOk

`func (o *EvidenceResponse) GetSourceTableOk() (*string, bool)`

GetSourceTableOk returns a tuple with the SourceTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTable

`func (o *EvidenceResponse) SetSourceTable(v string)`

SetSourceTable sets SourceTable field to given value.


### GetSourceId

`func (o *EvidenceResponse) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *EvidenceResponse) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *EvidenceResponse) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetSummary

`func (o *EvidenceResponse) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *EvidenceResponse) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *EvidenceResponse) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *EvidenceResponse) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetAmountMicros

`func (o *EvidenceResponse) GetAmountMicros() int32`

GetAmountMicros returns the AmountMicros field if non-nil, zero value otherwise.

### GetAmountMicrosOk

`func (o *EvidenceResponse) GetAmountMicrosOk() (*int32, bool)`

GetAmountMicrosOk returns a tuple with the AmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMicros

`func (o *EvidenceResponse) SetAmountMicros(v int32)`

SetAmountMicros sets AmountMicros field to given value.

### HasAmountMicros

`func (o *EvidenceResponse) HasAmountMicros() bool`

HasAmountMicros returns a boolean if a field has been set.

### GetResolvedAmountMicros

`func (o *EvidenceResponse) GetResolvedAmountMicros() int32`

GetResolvedAmountMicros returns the ResolvedAmountMicros field if non-nil, zero value otherwise.

### GetResolvedAmountMicrosOk

`func (o *EvidenceResponse) GetResolvedAmountMicrosOk() (*int32, bool)`

GetResolvedAmountMicrosOk returns a tuple with the ResolvedAmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedAmountMicros

`func (o *EvidenceResponse) SetResolvedAmountMicros(v int32)`

SetResolvedAmountMicros sets ResolvedAmountMicros field to given value.

### HasResolvedAmountMicros

`func (o *EvidenceResponse) HasResolvedAmountMicros() bool`

HasResolvedAmountMicros returns a boolean if a field has been set.

### GetResolutionStatus

`func (o *EvidenceResponse) GetResolutionStatus() string`

GetResolutionStatus returns the ResolutionStatus field if non-nil, zero value otherwise.

### GetResolutionStatusOk

`func (o *EvidenceResponse) GetResolutionStatusOk() (*string, bool)`

GetResolutionStatusOk returns a tuple with the ResolutionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionStatus

`func (o *EvidenceResponse) SetResolutionStatus(v string)`

SetResolutionStatus sets ResolutionStatus field to given value.

### HasResolutionStatus

`func (o *EvidenceResponse) HasResolutionStatus() bool`

HasResolutionStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EvidenceResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EvidenceResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EvidenceResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EvidenceResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


