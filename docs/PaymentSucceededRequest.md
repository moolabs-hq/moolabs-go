# PaymentSucceededRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Tenant identifier | 
**PoolId** | **string** | Pool identifier | 
**WalletId** | **string** | Wallet identifier | 
**PaymentId** | **string** | Payment identifier (idempotency key) | 
**AmountMicros** | **int32** | Amount in micros | 
**FxRateVersion** | Pointer to **string** | FX rate version | [optional] 
**CreditsPerUsdMicros** | Pointer to **int32** | Credits per USD micros | [optional] 
**EffectiveAt** | Pointer to **time.Time** | Effective timestamp | [optional] 

## Methods

### NewPaymentSucceededRequest

`func NewPaymentSucceededRequest(tenantId string, poolId string, walletId string, paymentId string, amountMicros int32, ) *PaymentSucceededRequest`

NewPaymentSucceededRequest instantiates a new PaymentSucceededRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentSucceededRequestWithDefaults

`func NewPaymentSucceededRequestWithDefaults() *PaymentSucceededRequest`

NewPaymentSucceededRequestWithDefaults instantiates a new PaymentSucceededRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *PaymentSucceededRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *PaymentSucceededRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *PaymentSucceededRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetPoolId

`func (o *PaymentSucceededRequest) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *PaymentSucceededRequest) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *PaymentSucceededRequest) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.


### GetWalletId

`func (o *PaymentSucceededRequest) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *PaymentSucceededRequest) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *PaymentSucceededRequest) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetPaymentId

`func (o *PaymentSucceededRequest) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PaymentSucceededRequest) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PaymentSucceededRequest) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### GetAmountMicros

`func (o *PaymentSucceededRequest) GetAmountMicros() int32`

GetAmountMicros returns the AmountMicros field if non-nil, zero value otherwise.

### GetAmountMicrosOk

`func (o *PaymentSucceededRequest) GetAmountMicrosOk() (*int32, bool)`

GetAmountMicrosOk returns a tuple with the AmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMicros

`func (o *PaymentSucceededRequest) SetAmountMicros(v int32)`

SetAmountMicros sets AmountMicros field to given value.


### GetFxRateVersion

`func (o *PaymentSucceededRequest) GetFxRateVersion() string`

GetFxRateVersion returns the FxRateVersion field if non-nil, zero value otherwise.

### GetFxRateVersionOk

`func (o *PaymentSucceededRequest) GetFxRateVersionOk() (*string, bool)`

GetFxRateVersionOk returns a tuple with the FxRateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFxRateVersion

`func (o *PaymentSucceededRequest) SetFxRateVersion(v string)`

SetFxRateVersion sets FxRateVersion field to given value.

### HasFxRateVersion

`func (o *PaymentSucceededRequest) HasFxRateVersion() bool`

HasFxRateVersion returns a boolean if a field has been set.

### GetCreditsPerUsdMicros

`func (o *PaymentSucceededRequest) GetCreditsPerUsdMicros() int32`

GetCreditsPerUsdMicros returns the CreditsPerUsdMicros field if non-nil, zero value otherwise.

### GetCreditsPerUsdMicrosOk

`func (o *PaymentSucceededRequest) GetCreditsPerUsdMicrosOk() (*int32, bool)`

GetCreditsPerUsdMicrosOk returns a tuple with the CreditsPerUsdMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsPerUsdMicros

`func (o *PaymentSucceededRequest) SetCreditsPerUsdMicros(v int32)`

SetCreditsPerUsdMicros sets CreditsPerUsdMicros field to given value.

### HasCreditsPerUsdMicros

`func (o *PaymentSucceededRequest) HasCreditsPerUsdMicros() bool`

HasCreditsPerUsdMicros returns a boolean if a field has been set.

### GetEffectiveAt

`func (o *PaymentSucceededRequest) GetEffectiveAt() time.Time`

GetEffectiveAt returns the EffectiveAt field if non-nil, zero value otherwise.

### GetEffectiveAtOk

`func (o *PaymentSucceededRequest) GetEffectiveAtOk() (*time.Time, bool)`

GetEffectiveAtOk returns a tuple with the EffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAt

`func (o *PaymentSucceededRequest) SetEffectiveAt(v time.Time)`

SetEffectiveAt sets EffectiveAt field to given value.

### HasEffectiveAt

`func (o *PaymentSucceededRequest) HasEffectiveAt() bool`

HasEffectiveAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


