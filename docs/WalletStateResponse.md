# WalletStateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletId** | **string** |  | 
**PoolId** | **string** |  | 
**BalanceMicros** | **int32** |  | 
**BalanceDecimal** | **float32** |  | 
**EffectiveAvailableMicros** | Pointer to **int32** |  | [optional] 
**State** | **string** |  | 
**AsOf** | **string** |  | 
**Consistency** | **string** |  | 
**LastProjectedAt** | Pointer to **string** |  | [optional] 
**ProjectedFromRecordedCut** | Pointer to **string** |  | [optional] 

## Methods

### NewWalletStateResponse

`func NewWalletStateResponse(walletId string, poolId string, balanceMicros int32, balanceDecimal float32, state string, asOf string, consistency string, ) *WalletStateResponse`

NewWalletStateResponse instantiates a new WalletStateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletStateResponseWithDefaults

`func NewWalletStateResponseWithDefaults() *WalletStateResponse`

NewWalletStateResponseWithDefaults instantiates a new WalletStateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletId

`func (o *WalletStateResponse) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *WalletStateResponse) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *WalletStateResponse) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetPoolId

`func (o *WalletStateResponse) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *WalletStateResponse) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *WalletStateResponse) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.


### GetBalanceMicros

`func (o *WalletStateResponse) GetBalanceMicros() int32`

GetBalanceMicros returns the BalanceMicros field if non-nil, zero value otherwise.

### GetBalanceMicrosOk

`func (o *WalletStateResponse) GetBalanceMicrosOk() (*int32, bool)`

GetBalanceMicrosOk returns a tuple with the BalanceMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceMicros

`func (o *WalletStateResponse) SetBalanceMicros(v int32)`

SetBalanceMicros sets BalanceMicros field to given value.


### GetBalanceDecimal

`func (o *WalletStateResponse) GetBalanceDecimal() float32`

GetBalanceDecimal returns the BalanceDecimal field if non-nil, zero value otherwise.

### GetBalanceDecimalOk

`func (o *WalletStateResponse) GetBalanceDecimalOk() (*float32, bool)`

GetBalanceDecimalOk returns a tuple with the BalanceDecimal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceDecimal

`func (o *WalletStateResponse) SetBalanceDecimal(v float32)`

SetBalanceDecimal sets BalanceDecimal field to given value.


### GetEffectiveAvailableMicros

`func (o *WalletStateResponse) GetEffectiveAvailableMicros() int32`

GetEffectiveAvailableMicros returns the EffectiveAvailableMicros field if non-nil, zero value otherwise.

### GetEffectiveAvailableMicrosOk

`func (o *WalletStateResponse) GetEffectiveAvailableMicrosOk() (*int32, bool)`

GetEffectiveAvailableMicrosOk returns a tuple with the EffectiveAvailableMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAvailableMicros

`func (o *WalletStateResponse) SetEffectiveAvailableMicros(v int32)`

SetEffectiveAvailableMicros sets EffectiveAvailableMicros field to given value.

### HasEffectiveAvailableMicros

`func (o *WalletStateResponse) HasEffectiveAvailableMicros() bool`

HasEffectiveAvailableMicros returns a boolean if a field has been set.

### GetState

`func (o *WalletStateResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *WalletStateResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *WalletStateResponse) SetState(v string)`

SetState sets State field to given value.


### GetAsOf

`func (o *WalletStateResponse) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *WalletStateResponse) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *WalletStateResponse) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.


### GetConsistency

`func (o *WalletStateResponse) GetConsistency() string`

GetConsistency returns the Consistency field if non-nil, zero value otherwise.

### GetConsistencyOk

`func (o *WalletStateResponse) GetConsistencyOk() (*string, bool)`

GetConsistencyOk returns a tuple with the Consistency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsistency

`func (o *WalletStateResponse) SetConsistency(v string)`

SetConsistency sets Consistency field to given value.


### GetLastProjectedAt

`func (o *WalletStateResponse) GetLastProjectedAt() string`

GetLastProjectedAt returns the LastProjectedAt field if non-nil, zero value otherwise.

### GetLastProjectedAtOk

`func (o *WalletStateResponse) GetLastProjectedAtOk() (*string, bool)`

GetLastProjectedAtOk returns a tuple with the LastProjectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastProjectedAt

`func (o *WalletStateResponse) SetLastProjectedAt(v string)`

SetLastProjectedAt sets LastProjectedAt field to given value.

### HasLastProjectedAt

`func (o *WalletStateResponse) HasLastProjectedAt() bool`

HasLastProjectedAt returns a boolean if a field has been set.

### GetProjectedFromRecordedCut

`func (o *WalletStateResponse) GetProjectedFromRecordedCut() string`

GetProjectedFromRecordedCut returns the ProjectedFromRecordedCut field if non-nil, zero value otherwise.

### GetProjectedFromRecordedCutOk

`func (o *WalletStateResponse) GetProjectedFromRecordedCutOk() (*string, bool)`

GetProjectedFromRecordedCutOk returns a tuple with the ProjectedFromRecordedCut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectedFromRecordedCut

`func (o *WalletStateResponse) SetProjectedFromRecordedCut(v string)`

SetProjectedFromRecordedCut sets ProjectedFromRecordedCut field to given value.

### HasProjectedFromRecordedCut

`func (o *WalletStateResponse) HasProjectedFromRecordedCut() bool`

HasProjectedFromRecordedCut returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


