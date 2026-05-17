# PaymentFailedRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Tenant identifier | 
**PoolId** | **string** | Pool identifier | 
**WalletId** | **string** | Wallet identifier | 
**PaymentId** | **string** | Payment identifier (idempotency key) | 
**FailureReason** | **string** | Failure reason | 
**AmountMicros** | Pointer to **int32** | Amount in micros | [optional] 
**FailedAt** | Pointer to **time.Time** | Failure timestamp | [optional] 

## Methods

### NewPaymentFailedRequest

`func NewPaymentFailedRequest(tenantId string, poolId string, walletId string, paymentId string, failureReason string, ) *PaymentFailedRequest`

NewPaymentFailedRequest instantiates a new PaymentFailedRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentFailedRequestWithDefaults

`func NewPaymentFailedRequestWithDefaults() *PaymentFailedRequest`

NewPaymentFailedRequestWithDefaults instantiates a new PaymentFailedRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *PaymentFailedRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *PaymentFailedRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *PaymentFailedRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetPoolId

`func (o *PaymentFailedRequest) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *PaymentFailedRequest) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *PaymentFailedRequest) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.


### GetWalletId

`func (o *PaymentFailedRequest) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *PaymentFailedRequest) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *PaymentFailedRequest) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetPaymentId

`func (o *PaymentFailedRequest) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PaymentFailedRequest) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PaymentFailedRequest) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### GetFailureReason

`func (o *PaymentFailedRequest) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *PaymentFailedRequest) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *PaymentFailedRequest) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.


### GetAmountMicros

`func (o *PaymentFailedRequest) GetAmountMicros() int32`

GetAmountMicros returns the AmountMicros field if non-nil, zero value otherwise.

### GetAmountMicrosOk

`func (o *PaymentFailedRequest) GetAmountMicrosOk() (*int32, bool)`

GetAmountMicrosOk returns a tuple with the AmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMicros

`func (o *PaymentFailedRequest) SetAmountMicros(v int32)`

SetAmountMicros sets AmountMicros field to given value.

### HasAmountMicros

`func (o *PaymentFailedRequest) HasAmountMicros() bool`

HasAmountMicros returns a boolean if a field has been set.

### GetFailedAt

`func (o *PaymentFailedRequest) GetFailedAt() time.Time`

GetFailedAt returns the FailedAt field if non-nil, zero value otherwise.

### GetFailedAtOk

`func (o *PaymentFailedRequest) GetFailedAtOk() (*time.Time, bool)`

GetFailedAtOk returns a tuple with the FailedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAt

`func (o *PaymentFailedRequest) SetFailedAt(v time.Time)`

SetFailedAt sets FailedAt field to given value.

### HasFailedAt

`func (o *PaymentFailedRequest) HasFailedAt() bool`

HasFailedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


