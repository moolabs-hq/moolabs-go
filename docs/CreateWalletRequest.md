# CreateWalletRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Tenant identifier | 
**PoolId** | **string** | Credit pool identifier | 
**OwnerType** | [**WalletOwnerType**](WalletOwnerType.md) | Owner type (ORG, TEAM, USER, etc.) | 
**OwnerId** | **string** | Owner identifier | 
**ParentWalletId** | Pointer to **NullableString** |  | [optional] 
**Depth** | Pointer to **int32** | Hierarchy depth (0-3) | [optional] [default to 0]
**Policy** | Pointer to [**WalletPolicy**](WalletPolicy.md) | Wallet policy | [optional] 
**LocalHardCapMicros** | Pointer to **NullableInt32** |  | [optional] 
**LocalSoftThresholdMicros** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCreateWalletRequest

`func NewCreateWalletRequest(tenantId string, poolId string, ownerType WalletOwnerType, ownerId string, ) *CreateWalletRequest`

NewCreateWalletRequest instantiates a new CreateWalletRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWalletRequestWithDefaults

`func NewCreateWalletRequestWithDefaults() *CreateWalletRequest`

NewCreateWalletRequestWithDefaults instantiates a new CreateWalletRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *CreateWalletRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CreateWalletRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CreateWalletRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetPoolId

`func (o *CreateWalletRequest) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *CreateWalletRequest) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *CreateWalletRequest) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.


### GetOwnerType

`func (o *CreateWalletRequest) GetOwnerType() WalletOwnerType`

GetOwnerType returns the OwnerType field if non-nil, zero value otherwise.

### GetOwnerTypeOk

`func (o *CreateWalletRequest) GetOwnerTypeOk() (*WalletOwnerType, bool)`

GetOwnerTypeOk returns a tuple with the OwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerType

`func (o *CreateWalletRequest) SetOwnerType(v WalletOwnerType)`

SetOwnerType sets OwnerType field to given value.


### GetOwnerId

`func (o *CreateWalletRequest) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *CreateWalletRequest) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *CreateWalletRequest) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetParentWalletId

`func (o *CreateWalletRequest) GetParentWalletId() string`

GetParentWalletId returns the ParentWalletId field if non-nil, zero value otherwise.

### GetParentWalletIdOk

`func (o *CreateWalletRequest) GetParentWalletIdOk() (*string, bool)`

GetParentWalletIdOk returns a tuple with the ParentWalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentWalletId

`func (o *CreateWalletRequest) SetParentWalletId(v string)`

SetParentWalletId sets ParentWalletId field to given value.

### HasParentWalletId

`func (o *CreateWalletRequest) HasParentWalletId() bool`

HasParentWalletId returns a boolean if a field has been set.

### SetParentWalletIdNil

`func (o *CreateWalletRequest) SetParentWalletIdNil(b bool)`

 SetParentWalletIdNil sets the value for ParentWalletId to be an explicit nil

### UnsetParentWalletId
`func (o *CreateWalletRequest) UnsetParentWalletId()`

UnsetParentWalletId ensures that no value is present for ParentWalletId, not even an explicit nil
### GetDepth

`func (o *CreateWalletRequest) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *CreateWalletRequest) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *CreateWalletRequest) SetDepth(v int32)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *CreateWalletRequest) HasDepth() bool`

HasDepth returns a boolean if a field has been set.

### GetPolicy

`func (o *CreateWalletRequest) GetPolicy() WalletPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *CreateWalletRequest) GetPolicyOk() (*WalletPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *CreateWalletRequest) SetPolicy(v WalletPolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *CreateWalletRequest) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetLocalHardCapMicros

`func (o *CreateWalletRequest) GetLocalHardCapMicros() int32`

GetLocalHardCapMicros returns the LocalHardCapMicros field if non-nil, zero value otherwise.

### GetLocalHardCapMicrosOk

`func (o *CreateWalletRequest) GetLocalHardCapMicrosOk() (*int32, bool)`

GetLocalHardCapMicrosOk returns a tuple with the LocalHardCapMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalHardCapMicros

`func (o *CreateWalletRequest) SetLocalHardCapMicros(v int32)`

SetLocalHardCapMicros sets LocalHardCapMicros field to given value.

### HasLocalHardCapMicros

`func (o *CreateWalletRequest) HasLocalHardCapMicros() bool`

HasLocalHardCapMicros returns a boolean if a field has been set.

### SetLocalHardCapMicrosNil

`func (o *CreateWalletRequest) SetLocalHardCapMicrosNil(b bool)`

 SetLocalHardCapMicrosNil sets the value for LocalHardCapMicros to be an explicit nil

### UnsetLocalHardCapMicros
`func (o *CreateWalletRequest) UnsetLocalHardCapMicros()`

UnsetLocalHardCapMicros ensures that no value is present for LocalHardCapMicros, not even an explicit nil
### GetLocalSoftThresholdMicros

`func (o *CreateWalletRequest) GetLocalSoftThresholdMicros() int32`

GetLocalSoftThresholdMicros returns the LocalSoftThresholdMicros field if non-nil, zero value otherwise.

### GetLocalSoftThresholdMicrosOk

`func (o *CreateWalletRequest) GetLocalSoftThresholdMicrosOk() (*int32, bool)`

GetLocalSoftThresholdMicrosOk returns a tuple with the LocalSoftThresholdMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSoftThresholdMicros

`func (o *CreateWalletRequest) SetLocalSoftThresholdMicros(v int32)`

SetLocalSoftThresholdMicros sets LocalSoftThresholdMicros field to given value.

### HasLocalSoftThresholdMicros

`func (o *CreateWalletRequest) HasLocalSoftThresholdMicros() bool`

HasLocalSoftThresholdMicros returns a boolean if a field has been set.

### SetLocalSoftThresholdMicrosNil

`func (o *CreateWalletRequest) SetLocalSoftThresholdMicrosNil(b bool)`

 SetLocalSoftThresholdMicrosNil sets the value for LocalSoftThresholdMicros to be an explicit nil

### UnsetLocalSoftThresholdMicros
`func (o *CreateWalletRequest) UnsetLocalSoftThresholdMicros()`

UnsetLocalSoftThresholdMicros ensures that no value is present for LocalSoftThresholdMicros, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


