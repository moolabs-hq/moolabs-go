# WalletResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**PoolId** | **string** |  | 
**OwnerType** | [**WalletOwnerType**](WalletOwnerType.md) |  | 
**OwnerId** | **string** |  | 
**ParentWalletId** | **NullableString** |  | 
**Depth** | **int32** |  | 
**Policy** | [**WalletPolicy**](WalletPolicy.md) |  | 
**LocalHardCapMicros** | **NullableInt32** |  | 
**LocalSoftThresholdMicros** | **NullableInt32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewWalletResponse

`func NewWalletResponse(id string, tenantId string, poolId string, ownerType WalletOwnerType, ownerId string, parentWalletId NullableString, depth int32, policy WalletPolicy, localHardCapMicros NullableInt32, localSoftThresholdMicros NullableInt32, createdAt time.Time, updatedAt time.Time, ) *WalletResponse`

NewWalletResponse instantiates a new WalletResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletResponseWithDefaults

`func NewWalletResponseWithDefaults() *WalletResponse`

NewWalletResponseWithDefaults instantiates a new WalletResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WalletResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WalletResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WalletResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *WalletResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *WalletResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *WalletResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetPoolId

`func (o *WalletResponse) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *WalletResponse) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *WalletResponse) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.


### GetOwnerType

`func (o *WalletResponse) GetOwnerType() WalletOwnerType`

GetOwnerType returns the OwnerType field if non-nil, zero value otherwise.

### GetOwnerTypeOk

`func (o *WalletResponse) GetOwnerTypeOk() (*WalletOwnerType, bool)`

GetOwnerTypeOk returns a tuple with the OwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerType

`func (o *WalletResponse) SetOwnerType(v WalletOwnerType)`

SetOwnerType sets OwnerType field to given value.


### GetOwnerId

`func (o *WalletResponse) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *WalletResponse) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *WalletResponse) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetParentWalletId

`func (o *WalletResponse) GetParentWalletId() string`

GetParentWalletId returns the ParentWalletId field if non-nil, zero value otherwise.

### GetParentWalletIdOk

`func (o *WalletResponse) GetParentWalletIdOk() (*string, bool)`

GetParentWalletIdOk returns a tuple with the ParentWalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentWalletId

`func (o *WalletResponse) SetParentWalletId(v string)`

SetParentWalletId sets ParentWalletId field to given value.


### SetParentWalletIdNil

`func (o *WalletResponse) SetParentWalletIdNil(b bool)`

 SetParentWalletIdNil sets the value for ParentWalletId to be an explicit nil

### UnsetParentWalletId
`func (o *WalletResponse) UnsetParentWalletId()`

UnsetParentWalletId ensures that no value is present for ParentWalletId, not even an explicit nil
### GetDepth

`func (o *WalletResponse) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *WalletResponse) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *WalletResponse) SetDepth(v int32)`

SetDepth sets Depth field to given value.


### GetPolicy

`func (o *WalletResponse) GetPolicy() WalletPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *WalletResponse) GetPolicyOk() (*WalletPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *WalletResponse) SetPolicy(v WalletPolicy)`

SetPolicy sets Policy field to given value.


### GetLocalHardCapMicros

`func (o *WalletResponse) GetLocalHardCapMicros() int32`

GetLocalHardCapMicros returns the LocalHardCapMicros field if non-nil, zero value otherwise.

### GetLocalHardCapMicrosOk

`func (o *WalletResponse) GetLocalHardCapMicrosOk() (*int32, bool)`

GetLocalHardCapMicrosOk returns a tuple with the LocalHardCapMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalHardCapMicros

`func (o *WalletResponse) SetLocalHardCapMicros(v int32)`

SetLocalHardCapMicros sets LocalHardCapMicros field to given value.


### SetLocalHardCapMicrosNil

`func (o *WalletResponse) SetLocalHardCapMicrosNil(b bool)`

 SetLocalHardCapMicrosNil sets the value for LocalHardCapMicros to be an explicit nil

### UnsetLocalHardCapMicros
`func (o *WalletResponse) UnsetLocalHardCapMicros()`

UnsetLocalHardCapMicros ensures that no value is present for LocalHardCapMicros, not even an explicit nil
### GetLocalSoftThresholdMicros

`func (o *WalletResponse) GetLocalSoftThresholdMicros() int32`

GetLocalSoftThresholdMicros returns the LocalSoftThresholdMicros field if non-nil, zero value otherwise.

### GetLocalSoftThresholdMicrosOk

`func (o *WalletResponse) GetLocalSoftThresholdMicrosOk() (*int32, bool)`

GetLocalSoftThresholdMicrosOk returns a tuple with the LocalSoftThresholdMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSoftThresholdMicros

`func (o *WalletResponse) SetLocalSoftThresholdMicros(v int32)`

SetLocalSoftThresholdMicros sets LocalSoftThresholdMicros field to given value.


### SetLocalSoftThresholdMicrosNil

`func (o *WalletResponse) SetLocalSoftThresholdMicrosNil(b bool)`

 SetLocalSoftThresholdMicrosNil sets the value for LocalSoftThresholdMicros to be an explicit nil

### UnsetLocalSoftThresholdMicros
`func (o *WalletResponse) UnsetLocalSoftThresholdMicros()`

UnsetLocalSoftThresholdMicros ensures that no value is present for LocalSoftThresholdMicros, not even an explicit nil
### GetCreatedAt

`func (o *WalletResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WalletResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WalletResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *WalletResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WalletResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WalletResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


