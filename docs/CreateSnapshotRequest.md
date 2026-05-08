# CreateSnapshotRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Tenant identifier | 
**PoolId** | **string** | Pool identifier | 
**WalletId** | **string** | Wallet identifier | 
**AsOfEffectiveAt** | **time.Time** | Effective timestamp for balance calculation | 
**AsOfRecordedAt** | Pointer to **NullableTime** |  | [optional] 
**IsolationLevel** | Pointer to **string** | Transaction isolation level | [optional] [default to "REPEATABLE READ"]

## Methods

### NewCreateSnapshotRequest

`func NewCreateSnapshotRequest(tenantId string, poolId string, walletId string, asOfEffectiveAt time.Time, ) *CreateSnapshotRequest`

NewCreateSnapshotRequest instantiates a new CreateSnapshotRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSnapshotRequestWithDefaults

`func NewCreateSnapshotRequestWithDefaults() *CreateSnapshotRequest`

NewCreateSnapshotRequestWithDefaults instantiates a new CreateSnapshotRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *CreateSnapshotRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CreateSnapshotRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CreateSnapshotRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetPoolId

`func (o *CreateSnapshotRequest) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *CreateSnapshotRequest) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *CreateSnapshotRequest) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.


### GetWalletId

`func (o *CreateSnapshotRequest) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *CreateSnapshotRequest) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *CreateSnapshotRequest) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetAsOfEffectiveAt

`func (o *CreateSnapshotRequest) GetAsOfEffectiveAt() time.Time`

GetAsOfEffectiveAt returns the AsOfEffectiveAt field if non-nil, zero value otherwise.

### GetAsOfEffectiveAtOk

`func (o *CreateSnapshotRequest) GetAsOfEffectiveAtOk() (*time.Time, bool)`

GetAsOfEffectiveAtOk returns a tuple with the AsOfEffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOfEffectiveAt

`func (o *CreateSnapshotRequest) SetAsOfEffectiveAt(v time.Time)`

SetAsOfEffectiveAt sets AsOfEffectiveAt field to given value.


### GetAsOfRecordedAt

`func (o *CreateSnapshotRequest) GetAsOfRecordedAt() time.Time`

GetAsOfRecordedAt returns the AsOfRecordedAt field if non-nil, zero value otherwise.

### GetAsOfRecordedAtOk

`func (o *CreateSnapshotRequest) GetAsOfRecordedAtOk() (*time.Time, bool)`

GetAsOfRecordedAtOk returns a tuple with the AsOfRecordedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOfRecordedAt

`func (o *CreateSnapshotRequest) SetAsOfRecordedAt(v time.Time)`

SetAsOfRecordedAt sets AsOfRecordedAt field to given value.

### HasAsOfRecordedAt

`func (o *CreateSnapshotRequest) HasAsOfRecordedAt() bool`

HasAsOfRecordedAt returns a boolean if a field has been set.

### SetAsOfRecordedAtNil

`func (o *CreateSnapshotRequest) SetAsOfRecordedAtNil(b bool)`

 SetAsOfRecordedAtNil sets the value for AsOfRecordedAt to be an explicit nil

### UnsetAsOfRecordedAt
`func (o *CreateSnapshotRequest) UnsetAsOfRecordedAt()`

UnsetAsOfRecordedAt ensures that no value is present for AsOfRecordedAt, not even an explicit nil
### GetIsolationLevel

`func (o *CreateSnapshotRequest) GetIsolationLevel() string`

GetIsolationLevel returns the IsolationLevel field if non-nil, zero value otherwise.

### GetIsolationLevelOk

`func (o *CreateSnapshotRequest) GetIsolationLevelOk() (*string, bool)`

GetIsolationLevelOk returns a tuple with the IsolationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolationLevel

`func (o *CreateSnapshotRequest) SetIsolationLevel(v string)`

SetIsolationLevel sets IsolationLevel field to given value.

### HasIsolationLevel

`func (o *CreateSnapshotRequest) HasIsolationLevel() bool`

HasIsolationLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


