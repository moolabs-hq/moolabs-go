# CreateFxRateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Tenant identifier | 
**PoolId** | **string** | Pool identifier | 
**CreditsPerUsdMicros** | **int32** | Credits per USD in micros | 
**EffectiveAt** | **time.Time** | Effective timestamp for this rate | 
**RateVersion** | Pointer to **string** | Rate version identifier (auto-generated if not provided) | [optional] 
**UsdMicros** | Pointer to **int32** | USD amount in micros (default 1 USD) | [optional] [default to 1000000]
**CreatedBy** | Pointer to **string** | Email or identifier of the user creating this rate | [optional] 

## Methods

### NewCreateFxRateRequest

`func NewCreateFxRateRequest(tenantId string, poolId string, creditsPerUsdMicros int32, effectiveAt time.Time, ) *CreateFxRateRequest`

NewCreateFxRateRequest instantiates a new CreateFxRateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFxRateRequestWithDefaults

`func NewCreateFxRateRequestWithDefaults() *CreateFxRateRequest`

NewCreateFxRateRequestWithDefaults instantiates a new CreateFxRateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *CreateFxRateRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CreateFxRateRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CreateFxRateRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetPoolId

`func (o *CreateFxRateRequest) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *CreateFxRateRequest) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *CreateFxRateRequest) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.


### GetCreditsPerUsdMicros

`func (o *CreateFxRateRequest) GetCreditsPerUsdMicros() int32`

GetCreditsPerUsdMicros returns the CreditsPerUsdMicros field if non-nil, zero value otherwise.

### GetCreditsPerUsdMicrosOk

`func (o *CreateFxRateRequest) GetCreditsPerUsdMicrosOk() (*int32, bool)`

GetCreditsPerUsdMicrosOk returns a tuple with the CreditsPerUsdMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsPerUsdMicros

`func (o *CreateFxRateRequest) SetCreditsPerUsdMicros(v int32)`

SetCreditsPerUsdMicros sets CreditsPerUsdMicros field to given value.


### GetEffectiveAt

`func (o *CreateFxRateRequest) GetEffectiveAt() time.Time`

GetEffectiveAt returns the EffectiveAt field if non-nil, zero value otherwise.

### GetEffectiveAtOk

`func (o *CreateFxRateRequest) GetEffectiveAtOk() (*time.Time, bool)`

GetEffectiveAtOk returns a tuple with the EffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAt

`func (o *CreateFxRateRequest) SetEffectiveAt(v time.Time)`

SetEffectiveAt sets EffectiveAt field to given value.


### GetRateVersion

`func (o *CreateFxRateRequest) GetRateVersion() string`

GetRateVersion returns the RateVersion field if non-nil, zero value otherwise.

### GetRateVersionOk

`func (o *CreateFxRateRequest) GetRateVersionOk() (*string, bool)`

GetRateVersionOk returns a tuple with the RateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateVersion

`func (o *CreateFxRateRequest) SetRateVersion(v string)`

SetRateVersion sets RateVersion field to given value.

### HasRateVersion

`func (o *CreateFxRateRequest) HasRateVersion() bool`

HasRateVersion returns a boolean if a field has been set.

### GetUsdMicros

`func (o *CreateFxRateRequest) GetUsdMicros() int32`

GetUsdMicros returns the UsdMicros field if non-nil, zero value otherwise.

### GetUsdMicrosOk

`func (o *CreateFxRateRequest) GetUsdMicrosOk() (*int32, bool)`

GetUsdMicrosOk returns a tuple with the UsdMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsdMicros

`func (o *CreateFxRateRequest) SetUsdMicros(v int32)`

SetUsdMicros sets UsdMicros field to given value.

### HasUsdMicros

`func (o *CreateFxRateRequest) HasUsdMicros() bool`

HasUsdMicros returns a boolean if a field has been set.

### GetCreatedBy

`func (o *CreateFxRateRequest) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CreateFxRateRequest) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *CreateFxRateRequest) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *CreateFxRateRequest) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


