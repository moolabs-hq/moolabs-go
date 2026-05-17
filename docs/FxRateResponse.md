# FxRateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**PoolId** | **string** |  | 
**CreditsPerUsdMicros** | **int32** |  | 
**UsdMicros** | **int32** |  | 
**EffectiveAt** | **time.Time** |  | 
**RateVersion** | **string** |  | 
**RecordedAt** | **time.Time** |  | 
**CreatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewFxRateResponse

`func NewFxRateResponse(id string, tenantId string, poolId string, creditsPerUsdMicros int32, usdMicros int32, effectiveAt time.Time, rateVersion string, recordedAt time.Time, ) *FxRateResponse`

NewFxRateResponse instantiates a new FxRateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFxRateResponseWithDefaults

`func NewFxRateResponseWithDefaults() *FxRateResponse`

NewFxRateResponseWithDefaults instantiates a new FxRateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FxRateResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FxRateResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FxRateResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *FxRateResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *FxRateResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *FxRateResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetPoolId

`func (o *FxRateResponse) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *FxRateResponse) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *FxRateResponse) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.


### GetCreditsPerUsdMicros

`func (o *FxRateResponse) GetCreditsPerUsdMicros() int32`

GetCreditsPerUsdMicros returns the CreditsPerUsdMicros field if non-nil, zero value otherwise.

### GetCreditsPerUsdMicrosOk

`func (o *FxRateResponse) GetCreditsPerUsdMicrosOk() (*int32, bool)`

GetCreditsPerUsdMicrosOk returns a tuple with the CreditsPerUsdMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsPerUsdMicros

`func (o *FxRateResponse) SetCreditsPerUsdMicros(v int32)`

SetCreditsPerUsdMicros sets CreditsPerUsdMicros field to given value.


### GetUsdMicros

`func (o *FxRateResponse) GetUsdMicros() int32`

GetUsdMicros returns the UsdMicros field if non-nil, zero value otherwise.

### GetUsdMicrosOk

`func (o *FxRateResponse) GetUsdMicrosOk() (*int32, bool)`

GetUsdMicrosOk returns a tuple with the UsdMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsdMicros

`func (o *FxRateResponse) SetUsdMicros(v int32)`

SetUsdMicros sets UsdMicros field to given value.


### GetEffectiveAt

`func (o *FxRateResponse) GetEffectiveAt() time.Time`

GetEffectiveAt returns the EffectiveAt field if non-nil, zero value otherwise.

### GetEffectiveAtOk

`func (o *FxRateResponse) GetEffectiveAtOk() (*time.Time, bool)`

GetEffectiveAtOk returns a tuple with the EffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAt

`func (o *FxRateResponse) SetEffectiveAt(v time.Time)`

SetEffectiveAt sets EffectiveAt field to given value.


### GetRateVersion

`func (o *FxRateResponse) GetRateVersion() string`

GetRateVersion returns the RateVersion field if non-nil, zero value otherwise.

### GetRateVersionOk

`func (o *FxRateResponse) GetRateVersionOk() (*string, bool)`

GetRateVersionOk returns a tuple with the RateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateVersion

`func (o *FxRateResponse) SetRateVersion(v string)`

SetRateVersion sets RateVersion field to given value.


### GetRecordedAt

`func (o *FxRateResponse) GetRecordedAt() time.Time`

GetRecordedAt returns the RecordedAt field if non-nil, zero value otherwise.

### GetRecordedAtOk

`func (o *FxRateResponse) GetRecordedAtOk() (*time.Time, bool)`

GetRecordedAtOk returns a tuple with the RecordedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordedAt

`func (o *FxRateResponse) SetRecordedAt(v time.Time)`

SetRecordedAt sets RecordedAt field to given value.


### GetCreatedBy

`func (o *FxRateResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *FxRateResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *FxRateResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *FxRateResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


