# CreateAutoTopupRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Tenant identifier | 
**PoolId** | **string** | Pool identifier | 
**WalletId** | **string** | Wallet identifier | 
**TriggerType** | **string** | Trigger type (STATE or THRESHOLD) | 
**TriggerValue** | **string** | Trigger value (e.g., &#39;LOW&#39;, &#39;OVERDRAFT&#39;, &#39;ABS:-100000&#39;) | 
**TopupAmountMicros** | **int32** | Top-up amount in micros | 
**TopupCooldownSeconds** | Pointer to **int32** | Cooldown period in seconds | [optional] [default to 3600]
**MaxTopupsPerDay** | Pointer to **int32** | Maximum top-ups per day | [optional] [default to 5]
**PaymentMethodRef** | Pointer to **NullableString** |  | [optional] 
**InvoiceMode** | Pointer to **string** | Invoice mode | [optional] [default to "INVOICE_NOW"]

## Methods

### NewCreateAutoTopupRuleRequest

`func NewCreateAutoTopupRuleRequest(tenantId string, poolId string, walletId string, triggerType string, triggerValue string, topupAmountMicros int32, ) *CreateAutoTopupRuleRequest`

NewCreateAutoTopupRuleRequest instantiates a new CreateAutoTopupRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAutoTopupRuleRequestWithDefaults

`func NewCreateAutoTopupRuleRequestWithDefaults() *CreateAutoTopupRuleRequest`

NewCreateAutoTopupRuleRequestWithDefaults instantiates a new CreateAutoTopupRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *CreateAutoTopupRuleRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CreateAutoTopupRuleRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CreateAutoTopupRuleRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetPoolId

`func (o *CreateAutoTopupRuleRequest) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *CreateAutoTopupRuleRequest) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *CreateAutoTopupRuleRequest) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.


### GetWalletId

`func (o *CreateAutoTopupRuleRequest) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *CreateAutoTopupRuleRequest) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *CreateAutoTopupRuleRequest) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetTriggerType

`func (o *CreateAutoTopupRuleRequest) GetTriggerType() string`

GetTriggerType returns the TriggerType field if non-nil, zero value otherwise.

### GetTriggerTypeOk

`func (o *CreateAutoTopupRuleRequest) GetTriggerTypeOk() (*string, bool)`

GetTriggerTypeOk returns a tuple with the TriggerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerType

`func (o *CreateAutoTopupRuleRequest) SetTriggerType(v string)`

SetTriggerType sets TriggerType field to given value.


### GetTriggerValue

`func (o *CreateAutoTopupRuleRequest) GetTriggerValue() string`

GetTriggerValue returns the TriggerValue field if non-nil, zero value otherwise.

### GetTriggerValueOk

`func (o *CreateAutoTopupRuleRequest) GetTriggerValueOk() (*string, bool)`

GetTriggerValueOk returns a tuple with the TriggerValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerValue

`func (o *CreateAutoTopupRuleRequest) SetTriggerValue(v string)`

SetTriggerValue sets TriggerValue field to given value.


### GetTopupAmountMicros

`func (o *CreateAutoTopupRuleRequest) GetTopupAmountMicros() int32`

GetTopupAmountMicros returns the TopupAmountMicros field if non-nil, zero value otherwise.

### GetTopupAmountMicrosOk

`func (o *CreateAutoTopupRuleRequest) GetTopupAmountMicrosOk() (*int32, bool)`

GetTopupAmountMicrosOk returns a tuple with the TopupAmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopupAmountMicros

`func (o *CreateAutoTopupRuleRequest) SetTopupAmountMicros(v int32)`

SetTopupAmountMicros sets TopupAmountMicros field to given value.


### GetTopupCooldownSeconds

`func (o *CreateAutoTopupRuleRequest) GetTopupCooldownSeconds() int32`

GetTopupCooldownSeconds returns the TopupCooldownSeconds field if non-nil, zero value otherwise.

### GetTopupCooldownSecondsOk

`func (o *CreateAutoTopupRuleRequest) GetTopupCooldownSecondsOk() (*int32, bool)`

GetTopupCooldownSecondsOk returns a tuple with the TopupCooldownSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopupCooldownSeconds

`func (o *CreateAutoTopupRuleRequest) SetTopupCooldownSeconds(v int32)`

SetTopupCooldownSeconds sets TopupCooldownSeconds field to given value.

### HasTopupCooldownSeconds

`func (o *CreateAutoTopupRuleRequest) HasTopupCooldownSeconds() bool`

HasTopupCooldownSeconds returns a boolean if a field has been set.

### GetMaxTopupsPerDay

`func (o *CreateAutoTopupRuleRequest) GetMaxTopupsPerDay() int32`

GetMaxTopupsPerDay returns the MaxTopupsPerDay field if non-nil, zero value otherwise.

### GetMaxTopupsPerDayOk

`func (o *CreateAutoTopupRuleRequest) GetMaxTopupsPerDayOk() (*int32, bool)`

GetMaxTopupsPerDayOk returns a tuple with the MaxTopupsPerDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTopupsPerDay

`func (o *CreateAutoTopupRuleRequest) SetMaxTopupsPerDay(v int32)`

SetMaxTopupsPerDay sets MaxTopupsPerDay field to given value.

### HasMaxTopupsPerDay

`func (o *CreateAutoTopupRuleRequest) HasMaxTopupsPerDay() bool`

HasMaxTopupsPerDay returns a boolean if a field has been set.

### GetPaymentMethodRef

`func (o *CreateAutoTopupRuleRequest) GetPaymentMethodRef() string`

GetPaymentMethodRef returns the PaymentMethodRef field if non-nil, zero value otherwise.

### GetPaymentMethodRefOk

`func (o *CreateAutoTopupRuleRequest) GetPaymentMethodRefOk() (*string, bool)`

GetPaymentMethodRefOk returns a tuple with the PaymentMethodRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodRef

`func (o *CreateAutoTopupRuleRequest) SetPaymentMethodRef(v string)`

SetPaymentMethodRef sets PaymentMethodRef field to given value.

### HasPaymentMethodRef

`func (o *CreateAutoTopupRuleRequest) HasPaymentMethodRef() bool`

HasPaymentMethodRef returns a boolean if a field has been set.

### SetPaymentMethodRefNil

`func (o *CreateAutoTopupRuleRequest) SetPaymentMethodRefNil(b bool)`

 SetPaymentMethodRefNil sets the value for PaymentMethodRef to be an explicit nil

### UnsetPaymentMethodRef
`func (o *CreateAutoTopupRuleRequest) UnsetPaymentMethodRef()`

UnsetPaymentMethodRef ensures that no value is present for PaymentMethodRef, not even an explicit nil
### GetInvoiceMode

`func (o *CreateAutoTopupRuleRequest) GetInvoiceMode() string`

GetInvoiceMode returns the InvoiceMode field if non-nil, zero value otherwise.

### GetInvoiceModeOk

`func (o *CreateAutoTopupRuleRequest) GetInvoiceModeOk() (*string, bool)`

GetInvoiceModeOk returns a tuple with the InvoiceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceMode

`func (o *CreateAutoTopupRuleRequest) SetInvoiceMode(v string)`

SetInvoiceMode sets InvoiceMode field to given value.

### HasInvoiceMode

`func (o *CreateAutoTopupRuleRequest) HasInvoiceMode() bool`

HasInvoiceMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


