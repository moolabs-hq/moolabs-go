# ProcessValueRecognitionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Tenant identifier | 
**PoolId** | **string** | Pool identifier | 
**JournalEntryId** | **string** | Journal entry ID to process | 
**EffectiveAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewProcessValueRecognitionRequest

`func NewProcessValueRecognitionRequest(tenantId string, poolId string, journalEntryId string, ) *ProcessValueRecognitionRequest`

NewProcessValueRecognitionRequest instantiates a new ProcessValueRecognitionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessValueRecognitionRequestWithDefaults

`func NewProcessValueRecognitionRequestWithDefaults() *ProcessValueRecognitionRequest`

NewProcessValueRecognitionRequestWithDefaults instantiates a new ProcessValueRecognitionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *ProcessValueRecognitionRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ProcessValueRecognitionRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ProcessValueRecognitionRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetPoolId

`func (o *ProcessValueRecognitionRequest) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *ProcessValueRecognitionRequest) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *ProcessValueRecognitionRequest) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.


### GetJournalEntryId

`func (o *ProcessValueRecognitionRequest) GetJournalEntryId() string`

GetJournalEntryId returns the JournalEntryId field if non-nil, zero value otherwise.

### GetJournalEntryIdOk

`func (o *ProcessValueRecognitionRequest) GetJournalEntryIdOk() (*string, bool)`

GetJournalEntryIdOk returns a tuple with the JournalEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalEntryId

`func (o *ProcessValueRecognitionRequest) SetJournalEntryId(v string)`

SetJournalEntryId sets JournalEntryId field to given value.


### GetEffectiveAt

`func (o *ProcessValueRecognitionRequest) GetEffectiveAt() time.Time`

GetEffectiveAt returns the EffectiveAt field if non-nil, zero value otherwise.

### GetEffectiveAtOk

`func (o *ProcessValueRecognitionRequest) GetEffectiveAtOk() (*time.Time, bool)`

GetEffectiveAtOk returns a tuple with the EffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAt

`func (o *ProcessValueRecognitionRequest) SetEffectiveAt(v time.Time)`

SetEffectiveAt sets EffectiveAt field to given value.

### HasEffectiveAt

`func (o *ProcessValueRecognitionRequest) HasEffectiveAt() bool`

HasEffectiveAt returns a boolean if a field has been set.

### SetEffectiveAtNil

`func (o *ProcessValueRecognitionRequest) SetEffectiveAtNil(b bool)`

 SetEffectiveAtNil sets the value for EffectiveAt to be an explicit nil

### UnsetEffectiveAt
`func (o *ProcessValueRecognitionRequest) UnsetEffectiveAt()`

UnsetEffectiveAt ensures that no value is present for EffectiveAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


