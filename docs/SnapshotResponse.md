# SnapshotResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**PoolId** | **string** |  | 
**WalletId** | **string** |  | 
**AsOfEffectiveAt** | **time.Time** |  | 
**AsOfRecordedAt** | **time.Time** |  | 
**BalanceMicros** | **int32** |  | 
**SnapshotFingerprint** | **string** |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewSnapshotResponse

`func NewSnapshotResponse(id string, tenantId string, poolId string, walletId string, asOfEffectiveAt time.Time, asOfRecordedAt time.Time, balanceMicros int32, snapshotFingerprint string, createdAt time.Time, ) *SnapshotResponse`

NewSnapshotResponse instantiates a new SnapshotResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotResponseWithDefaults

`func NewSnapshotResponseWithDefaults() *SnapshotResponse`

NewSnapshotResponseWithDefaults instantiates a new SnapshotResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SnapshotResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SnapshotResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SnapshotResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *SnapshotResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SnapshotResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SnapshotResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetPoolId

`func (o *SnapshotResponse) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *SnapshotResponse) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *SnapshotResponse) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.


### GetWalletId

`func (o *SnapshotResponse) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *SnapshotResponse) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *SnapshotResponse) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetAsOfEffectiveAt

`func (o *SnapshotResponse) GetAsOfEffectiveAt() time.Time`

GetAsOfEffectiveAt returns the AsOfEffectiveAt field if non-nil, zero value otherwise.

### GetAsOfEffectiveAtOk

`func (o *SnapshotResponse) GetAsOfEffectiveAtOk() (*time.Time, bool)`

GetAsOfEffectiveAtOk returns a tuple with the AsOfEffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOfEffectiveAt

`func (o *SnapshotResponse) SetAsOfEffectiveAt(v time.Time)`

SetAsOfEffectiveAt sets AsOfEffectiveAt field to given value.


### GetAsOfRecordedAt

`func (o *SnapshotResponse) GetAsOfRecordedAt() time.Time`

GetAsOfRecordedAt returns the AsOfRecordedAt field if non-nil, zero value otherwise.

### GetAsOfRecordedAtOk

`func (o *SnapshotResponse) GetAsOfRecordedAtOk() (*time.Time, bool)`

GetAsOfRecordedAtOk returns a tuple with the AsOfRecordedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOfRecordedAt

`func (o *SnapshotResponse) SetAsOfRecordedAt(v time.Time)`

SetAsOfRecordedAt sets AsOfRecordedAt field to given value.


### GetBalanceMicros

`func (o *SnapshotResponse) GetBalanceMicros() int32`

GetBalanceMicros returns the BalanceMicros field if non-nil, zero value otherwise.

### GetBalanceMicrosOk

`func (o *SnapshotResponse) GetBalanceMicrosOk() (*int32, bool)`

GetBalanceMicrosOk returns a tuple with the BalanceMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceMicros

`func (o *SnapshotResponse) SetBalanceMicros(v int32)`

SetBalanceMicros sets BalanceMicros field to given value.


### GetSnapshotFingerprint

`func (o *SnapshotResponse) GetSnapshotFingerprint() string`

GetSnapshotFingerprint returns the SnapshotFingerprint field if non-nil, zero value otherwise.

### GetSnapshotFingerprintOk

`func (o *SnapshotResponse) GetSnapshotFingerprintOk() (*string, bool)`

GetSnapshotFingerprintOk returns a tuple with the SnapshotFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotFingerprint

`func (o *SnapshotResponse) SetSnapshotFingerprint(v string)`

SetSnapshotFingerprint sets SnapshotFingerprint field to given value.


### GetCreatedAt

`func (o *SnapshotResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SnapshotResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SnapshotResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


