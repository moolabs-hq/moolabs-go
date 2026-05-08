# CreateValueRecognitionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Tenant identifier | 
**PoolId** | **string** | Pool identifier | 
**WalletId** | **string** | Wallet identifier | 
**SubjectId** | **string** | Subject identifier | 
**UsdAmountMicros** | **int32** | USD amount in micros | 
**EffectiveAt** | **time.Time** | Effective timestamp | 
**UsageEventId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateValueRecognitionRequest

`func NewCreateValueRecognitionRequest(tenantId string, poolId string, walletId string, subjectId string, usdAmountMicros int32, effectiveAt time.Time, ) *CreateValueRecognitionRequest`

NewCreateValueRecognitionRequest instantiates a new CreateValueRecognitionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateValueRecognitionRequestWithDefaults

`func NewCreateValueRecognitionRequestWithDefaults() *CreateValueRecognitionRequest`

NewCreateValueRecognitionRequestWithDefaults instantiates a new CreateValueRecognitionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *CreateValueRecognitionRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CreateValueRecognitionRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CreateValueRecognitionRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetPoolId

`func (o *CreateValueRecognitionRequest) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *CreateValueRecognitionRequest) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *CreateValueRecognitionRequest) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.


### GetWalletId

`func (o *CreateValueRecognitionRequest) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *CreateValueRecognitionRequest) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *CreateValueRecognitionRequest) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetSubjectId

`func (o *CreateValueRecognitionRequest) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *CreateValueRecognitionRequest) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *CreateValueRecognitionRequest) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.


### GetUsdAmountMicros

`func (o *CreateValueRecognitionRequest) GetUsdAmountMicros() int32`

GetUsdAmountMicros returns the UsdAmountMicros field if non-nil, zero value otherwise.

### GetUsdAmountMicrosOk

`func (o *CreateValueRecognitionRequest) GetUsdAmountMicrosOk() (*int32, bool)`

GetUsdAmountMicrosOk returns a tuple with the UsdAmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsdAmountMicros

`func (o *CreateValueRecognitionRequest) SetUsdAmountMicros(v int32)`

SetUsdAmountMicros sets UsdAmountMicros field to given value.


### GetEffectiveAt

`func (o *CreateValueRecognitionRequest) GetEffectiveAt() time.Time`

GetEffectiveAt returns the EffectiveAt field if non-nil, zero value otherwise.

### GetEffectiveAtOk

`func (o *CreateValueRecognitionRequest) GetEffectiveAtOk() (*time.Time, bool)`

GetEffectiveAtOk returns a tuple with the EffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAt

`func (o *CreateValueRecognitionRequest) SetEffectiveAt(v time.Time)`

SetEffectiveAt sets EffectiveAt field to given value.


### GetUsageEventId

`func (o *CreateValueRecognitionRequest) GetUsageEventId() string`

GetUsageEventId returns the UsageEventId field if non-nil, zero value otherwise.

### GetUsageEventIdOk

`func (o *CreateValueRecognitionRequest) GetUsageEventIdOk() (*string, bool)`

GetUsageEventIdOk returns a tuple with the UsageEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageEventId

`func (o *CreateValueRecognitionRequest) SetUsageEventId(v string)`

SetUsageEventId sets UsageEventId field to given value.

### HasUsageEventId

`func (o *CreateValueRecognitionRequest) HasUsageEventId() bool`

HasUsageEventId returns a boolean if a field has been set.

### SetUsageEventIdNil

`func (o *CreateValueRecognitionRequest) SetUsageEventIdNil(b bool)`

 SetUsageEventIdNil sets the value for UsageEventId to be an explicit nil

### UnsetUsageEventId
`func (o *CreateValueRecognitionRequest) UnsetUsageEventId()`

UnsetUsageEventId ensures that no value is present for UsageEventId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


