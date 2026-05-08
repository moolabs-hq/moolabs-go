# AutoTopupRuleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**PoolId** | **string** |  | 
**WalletId** | **string** |  | 
**TriggerType** | **string** |  | 
**TriggerValue** | **string** |  | 
**TopupAmountMicros** | **int32** |  | 
**TopupCooldownSeconds** | **int32** |  | 
**MaxTopupsPerDay** | **int32** |  | 
**PaymentMethodRef** | **NullableString** |  | 
**InvoiceMode** | **string** |  | 
**Enabled** | Pointer to **NullableBool** |  | [optional] 
**LastTriggeredAt** | Pointer to **NullableTime** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewAutoTopupRuleResponse

`func NewAutoTopupRuleResponse(id string, tenantId string, poolId string, walletId string, triggerType string, triggerValue string, topupAmountMicros int32, topupCooldownSeconds int32, maxTopupsPerDay int32, paymentMethodRef NullableString, invoiceMode string, ) *AutoTopupRuleResponse`

NewAutoTopupRuleResponse instantiates a new AutoTopupRuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoTopupRuleResponseWithDefaults

`func NewAutoTopupRuleResponseWithDefaults() *AutoTopupRuleResponse`

NewAutoTopupRuleResponseWithDefaults instantiates a new AutoTopupRuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutoTopupRuleResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoTopupRuleResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoTopupRuleResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *AutoTopupRuleResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AutoTopupRuleResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AutoTopupRuleResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetPoolId

`func (o *AutoTopupRuleResponse) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *AutoTopupRuleResponse) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *AutoTopupRuleResponse) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.


### GetWalletId

`func (o *AutoTopupRuleResponse) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *AutoTopupRuleResponse) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *AutoTopupRuleResponse) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetTriggerType

`func (o *AutoTopupRuleResponse) GetTriggerType() string`

GetTriggerType returns the TriggerType field if non-nil, zero value otherwise.

### GetTriggerTypeOk

`func (o *AutoTopupRuleResponse) GetTriggerTypeOk() (*string, bool)`

GetTriggerTypeOk returns a tuple with the TriggerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerType

`func (o *AutoTopupRuleResponse) SetTriggerType(v string)`

SetTriggerType sets TriggerType field to given value.


### GetTriggerValue

`func (o *AutoTopupRuleResponse) GetTriggerValue() string`

GetTriggerValue returns the TriggerValue field if non-nil, zero value otherwise.

### GetTriggerValueOk

`func (o *AutoTopupRuleResponse) GetTriggerValueOk() (*string, bool)`

GetTriggerValueOk returns a tuple with the TriggerValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerValue

`func (o *AutoTopupRuleResponse) SetTriggerValue(v string)`

SetTriggerValue sets TriggerValue field to given value.


### GetTopupAmountMicros

`func (o *AutoTopupRuleResponse) GetTopupAmountMicros() int32`

GetTopupAmountMicros returns the TopupAmountMicros field if non-nil, zero value otherwise.

### GetTopupAmountMicrosOk

`func (o *AutoTopupRuleResponse) GetTopupAmountMicrosOk() (*int32, bool)`

GetTopupAmountMicrosOk returns a tuple with the TopupAmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopupAmountMicros

`func (o *AutoTopupRuleResponse) SetTopupAmountMicros(v int32)`

SetTopupAmountMicros sets TopupAmountMicros field to given value.


### GetTopupCooldownSeconds

`func (o *AutoTopupRuleResponse) GetTopupCooldownSeconds() int32`

GetTopupCooldownSeconds returns the TopupCooldownSeconds field if non-nil, zero value otherwise.

### GetTopupCooldownSecondsOk

`func (o *AutoTopupRuleResponse) GetTopupCooldownSecondsOk() (*int32, bool)`

GetTopupCooldownSecondsOk returns a tuple with the TopupCooldownSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopupCooldownSeconds

`func (o *AutoTopupRuleResponse) SetTopupCooldownSeconds(v int32)`

SetTopupCooldownSeconds sets TopupCooldownSeconds field to given value.


### GetMaxTopupsPerDay

`func (o *AutoTopupRuleResponse) GetMaxTopupsPerDay() int32`

GetMaxTopupsPerDay returns the MaxTopupsPerDay field if non-nil, zero value otherwise.

### GetMaxTopupsPerDayOk

`func (o *AutoTopupRuleResponse) GetMaxTopupsPerDayOk() (*int32, bool)`

GetMaxTopupsPerDayOk returns a tuple with the MaxTopupsPerDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTopupsPerDay

`func (o *AutoTopupRuleResponse) SetMaxTopupsPerDay(v int32)`

SetMaxTopupsPerDay sets MaxTopupsPerDay field to given value.


### GetPaymentMethodRef

`func (o *AutoTopupRuleResponse) GetPaymentMethodRef() string`

GetPaymentMethodRef returns the PaymentMethodRef field if non-nil, zero value otherwise.

### GetPaymentMethodRefOk

`func (o *AutoTopupRuleResponse) GetPaymentMethodRefOk() (*string, bool)`

GetPaymentMethodRefOk returns a tuple with the PaymentMethodRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodRef

`func (o *AutoTopupRuleResponse) SetPaymentMethodRef(v string)`

SetPaymentMethodRef sets PaymentMethodRef field to given value.


### SetPaymentMethodRefNil

`func (o *AutoTopupRuleResponse) SetPaymentMethodRefNil(b bool)`

 SetPaymentMethodRefNil sets the value for PaymentMethodRef to be an explicit nil

### UnsetPaymentMethodRef
`func (o *AutoTopupRuleResponse) UnsetPaymentMethodRef()`

UnsetPaymentMethodRef ensures that no value is present for PaymentMethodRef, not even an explicit nil
### GetInvoiceMode

`func (o *AutoTopupRuleResponse) GetInvoiceMode() string`

GetInvoiceMode returns the InvoiceMode field if non-nil, zero value otherwise.

### GetInvoiceModeOk

`func (o *AutoTopupRuleResponse) GetInvoiceModeOk() (*string, bool)`

GetInvoiceModeOk returns a tuple with the InvoiceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceMode

`func (o *AutoTopupRuleResponse) SetInvoiceMode(v string)`

SetInvoiceMode sets InvoiceMode field to given value.


### GetEnabled

`func (o *AutoTopupRuleResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AutoTopupRuleResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AutoTopupRuleResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AutoTopupRuleResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *AutoTopupRuleResponse) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *AutoTopupRuleResponse) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetLastTriggeredAt

`func (o *AutoTopupRuleResponse) GetLastTriggeredAt() time.Time`

GetLastTriggeredAt returns the LastTriggeredAt field if non-nil, zero value otherwise.

### GetLastTriggeredAtOk

`func (o *AutoTopupRuleResponse) GetLastTriggeredAtOk() (*time.Time, bool)`

GetLastTriggeredAtOk returns a tuple with the LastTriggeredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTriggeredAt

`func (o *AutoTopupRuleResponse) SetLastTriggeredAt(v time.Time)`

SetLastTriggeredAt sets LastTriggeredAt field to given value.

### HasLastTriggeredAt

`func (o *AutoTopupRuleResponse) HasLastTriggeredAt() bool`

HasLastTriggeredAt returns a boolean if a field has been set.

### SetLastTriggeredAtNil

`func (o *AutoTopupRuleResponse) SetLastTriggeredAtNil(b bool)`

 SetLastTriggeredAtNil sets the value for LastTriggeredAt to be an explicit nil

### UnsetLastTriggeredAt
`func (o *AutoTopupRuleResponse) UnsetLastTriggeredAt()`

UnsetLastTriggeredAt ensures that no value is present for LastTriggeredAt, not even an explicit nil
### GetCreatedAt

`func (o *AutoTopupRuleResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AutoTopupRuleResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AutoTopupRuleResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AutoTopupRuleResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *AutoTopupRuleResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *AutoTopupRuleResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


