# RecordUsageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Tenant identifier | 
**SubjectId** | **string** | Subject identifier | 
**PoolId** | **string** | Credit pool identifier | 
**WalletId** | Pointer to **NullableString** |  | [optional] 
**SubscriptionId** | Pointer to **NullableString** |  | [optional] 
**FeatureKey** | **string** | Feature key | 
**MeterSlug** | Pointer to **NullableString** |  | [optional] 
**UsageEventId** | **string** | Usage event ID (idempotency key) | 
**RequestId** | Pointer to **NullableString** |  | [optional] 
**JobId** | Pointer to **NullableString** |  | [optional] 
**EffectiveAt** | **time.Time** | Effective timestamp | 
**UsageVector** | **map[string]interface{}** | Usage vector (JSON object) | 

## Methods

### NewRecordUsageRequest

`func NewRecordUsageRequest(tenantId string, subjectId string, poolId string, featureKey string, usageEventId string, effectiveAt time.Time, usageVector map[string]interface{}, ) *RecordUsageRequest`

NewRecordUsageRequest instantiates a new RecordUsageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordUsageRequestWithDefaults

`func NewRecordUsageRequestWithDefaults() *RecordUsageRequest`

NewRecordUsageRequestWithDefaults instantiates a new RecordUsageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *RecordUsageRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *RecordUsageRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *RecordUsageRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetSubjectId

`func (o *RecordUsageRequest) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *RecordUsageRequest) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *RecordUsageRequest) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.


### GetPoolId

`func (o *RecordUsageRequest) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *RecordUsageRequest) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *RecordUsageRequest) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.


### GetWalletId

`func (o *RecordUsageRequest) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *RecordUsageRequest) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *RecordUsageRequest) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *RecordUsageRequest) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.

### SetWalletIdNil

`func (o *RecordUsageRequest) SetWalletIdNil(b bool)`

 SetWalletIdNil sets the value for WalletId to be an explicit nil

### UnsetWalletId
`func (o *RecordUsageRequest) UnsetWalletId()`

UnsetWalletId ensures that no value is present for WalletId, not even an explicit nil
### GetSubscriptionId

`func (o *RecordUsageRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *RecordUsageRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *RecordUsageRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *RecordUsageRequest) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *RecordUsageRequest) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *RecordUsageRequest) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetFeatureKey

`func (o *RecordUsageRequest) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *RecordUsageRequest) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *RecordUsageRequest) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.


### GetMeterSlug

`func (o *RecordUsageRequest) GetMeterSlug() string`

GetMeterSlug returns the MeterSlug field if non-nil, zero value otherwise.

### GetMeterSlugOk

`func (o *RecordUsageRequest) GetMeterSlugOk() (*string, bool)`

GetMeterSlugOk returns a tuple with the MeterSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterSlug

`func (o *RecordUsageRequest) SetMeterSlug(v string)`

SetMeterSlug sets MeterSlug field to given value.

### HasMeterSlug

`func (o *RecordUsageRequest) HasMeterSlug() bool`

HasMeterSlug returns a boolean if a field has been set.

### SetMeterSlugNil

`func (o *RecordUsageRequest) SetMeterSlugNil(b bool)`

 SetMeterSlugNil sets the value for MeterSlug to be an explicit nil

### UnsetMeterSlug
`func (o *RecordUsageRequest) UnsetMeterSlug()`

UnsetMeterSlug ensures that no value is present for MeterSlug, not even an explicit nil
### GetUsageEventId

`func (o *RecordUsageRequest) GetUsageEventId() string`

GetUsageEventId returns the UsageEventId field if non-nil, zero value otherwise.

### GetUsageEventIdOk

`func (o *RecordUsageRequest) GetUsageEventIdOk() (*string, bool)`

GetUsageEventIdOk returns a tuple with the UsageEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageEventId

`func (o *RecordUsageRequest) SetUsageEventId(v string)`

SetUsageEventId sets UsageEventId field to given value.


### GetRequestId

`func (o *RecordUsageRequest) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *RecordUsageRequest) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *RecordUsageRequest) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *RecordUsageRequest) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### SetRequestIdNil

`func (o *RecordUsageRequest) SetRequestIdNil(b bool)`

 SetRequestIdNil sets the value for RequestId to be an explicit nil

### UnsetRequestId
`func (o *RecordUsageRequest) UnsetRequestId()`

UnsetRequestId ensures that no value is present for RequestId, not even an explicit nil
### GetJobId

`func (o *RecordUsageRequest) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *RecordUsageRequest) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *RecordUsageRequest) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *RecordUsageRequest) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### SetJobIdNil

`func (o *RecordUsageRequest) SetJobIdNil(b bool)`

 SetJobIdNil sets the value for JobId to be an explicit nil

### UnsetJobId
`func (o *RecordUsageRequest) UnsetJobId()`

UnsetJobId ensures that no value is present for JobId, not even an explicit nil
### GetEffectiveAt

`func (o *RecordUsageRequest) GetEffectiveAt() time.Time`

GetEffectiveAt returns the EffectiveAt field if non-nil, zero value otherwise.

### GetEffectiveAtOk

`func (o *RecordUsageRequest) GetEffectiveAtOk() (*time.Time, bool)`

GetEffectiveAtOk returns a tuple with the EffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAt

`func (o *RecordUsageRequest) SetEffectiveAt(v time.Time)`

SetEffectiveAt sets EffectiveAt field to given value.


### GetUsageVector

`func (o *RecordUsageRequest) GetUsageVector() map[string]interface{}`

GetUsageVector returns the UsageVector field if non-nil, zero value otherwise.

### GetUsageVectorOk

`func (o *RecordUsageRequest) GetUsageVectorOk() (*map[string]interface{}, bool)`

GetUsageVectorOk returns a tuple with the UsageVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageVector

`func (o *RecordUsageRequest) SetUsageVector(v map[string]interface{})`

SetUsageVector sets UsageVector field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


