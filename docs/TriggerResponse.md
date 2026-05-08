# TriggerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Triggered** | **bool** |  | 
**Reason** | Pointer to **NullableString** |  | [optional] 
**RuleId** | Pointer to **NullableString** |  | [optional] 
**WalletId** | Pointer to **NullableString** |  | [optional] 
**TopupAmountMicros** | Pointer to **NullableInt32** |  | [optional] 
**PaymentMethodRef** | Pointer to **NullableString** |  | [optional] 
**InvoiceMode** | Pointer to **NullableString** |  | [optional] 
**OutboxEventId** | Pointer to **NullableString** |  | [optional] 
**CooldownUntil** | Pointer to **NullableTime** |  | [optional] 
**WalletState** | Pointer to **NullableString** |  | [optional] 
**AvailableBalanceMicros** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewTriggerResponse

`func NewTriggerResponse(triggered bool, ) *TriggerResponse`

NewTriggerResponse instantiates a new TriggerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerResponseWithDefaults

`func NewTriggerResponseWithDefaults() *TriggerResponse`

NewTriggerResponseWithDefaults instantiates a new TriggerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTriggered

`func (o *TriggerResponse) GetTriggered() bool`

GetTriggered returns the Triggered field if non-nil, zero value otherwise.

### GetTriggeredOk

`func (o *TriggerResponse) GetTriggeredOk() (*bool, bool)`

GetTriggeredOk returns a tuple with the Triggered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggered

`func (o *TriggerResponse) SetTriggered(v bool)`

SetTriggered sets Triggered field to given value.


### GetReason

`func (o *TriggerResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *TriggerResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *TriggerResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *TriggerResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *TriggerResponse) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *TriggerResponse) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetRuleId

`func (o *TriggerResponse) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *TriggerResponse) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *TriggerResponse) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *TriggerResponse) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### SetRuleIdNil

`func (o *TriggerResponse) SetRuleIdNil(b bool)`

 SetRuleIdNil sets the value for RuleId to be an explicit nil

### UnsetRuleId
`func (o *TriggerResponse) UnsetRuleId()`

UnsetRuleId ensures that no value is present for RuleId, not even an explicit nil
### GetWalletId

`func (o *TriggerResponse) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *TriggerResponse) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *TriggerResponse) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *TriggerResponse) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.

### SetWalletIdNil

`func (o *TriggerResponse) SetWalletIdNil(b bool)`

 SetWalletIdNil sets the value for WalletId to be an explicit nil

### UnsetWalletId
`func (o *TriggerResponse) UnsetWalletId()`

UnsetWalletId ensures that no value is present for WalletId, not even an explicit nil
### GetTopupAmountMicros

`func (o *TriggerResponse) GetTopupAmountMicros() int32`

GetTopupAmountMicros returns the TopupAmountMicros field if non-nil, zero value otherwise.

### GetTopupAmountMicrosOk

`func (o *TriggerResponse) GetTopupAmountMicrosOk() (*int32, bool)`

GetTopupAmountMicrosOk returns a tuple with the TopupAmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopupAmountMicros

`func (o *TriggerResponse) SetTopupAmountMicros(v int32)`

SetTopupAmountMicros sets TopupAmountMicros field to given value.

### HasTopupAmountMicros

`func (o *TriggerResponse) HasTopupAmountMicros() bool`

HasTopupAmountMicros returns a boolean if a field has been set.

### SetTopupAmountMicrosNil

`func (o *TriggerResponse) SetTopupAmountMicrosNil(b bool)`

 SetTopupAmountMicrosNil sets the value for TopupAmountMicros to be an explicit nil

### UnsetTopupAmountMicros
`func (o *TriggerResponse) UnsetTopupAmountMicros()`

UnsetTopupAmountMicros ensures that no value is present for TopupAmountMicros, not even an explicit nil
### GetPaymentMethodRef

`func (o *TriggerResponse) GetPaymentMethodRef() string`

GetPaymentMethodRef returns the PaymentMethodRef field if non-nil, zero value otherwise.

### GetPaymentMethodRefOk

`func (o *TriggerResponse) GetPaymentMethodRefOk() (*string, bool)`

GetPaymentMethodRefOk returns a tuple with the PaymentMethodRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodRef

`func (o *TriggerResponse) SetPaymentMethodRef(v string)`

SetPaymentMethodRef sets PaymentMethodRef field to given value.

### HasPaymentMethodRef

`func (o *TriggerResponse) HasPaymentMethodRef() bool`

HasPaymentMethodRef returns a boolean if a field has been set.

### SetPaymentMethodRefNil

`func (o *TriggerResponse) SetPaymentMethodRefNil(b bool)`

 SetPaymentMethodRefNil sets the value for PaymentMethodRef to be an explicit nil

### UnsetPaymentMethodRef
`func (o *TriggerResponse) UnsetPaymentMethodRef()`

UnsetPaymentMethodRef ensures that no value is present for PaymentMethodRef, not even an explicit nil
### GetInvoiceMode

`func (o *TriggerResponse) GetInvoiceMode() string`

GetInvoiceMode returns the InvoiceMode field if non-nil, zero value otherwise.

### GetInvoiceModeOk

`func (o *TriggerResponse) GetInvoiceModeOk() (*string, bool)`

GetInvoiceModeOk returns a tuple with the InvoiceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceMode

`func (o *TriggerResponse) SetInvoiceMode(v string)`

SetInvoiceMode sets InvoiceMode field to given value.

### HasInvoiceMode

`func (o *TriggerResponse) HasInvoiceMode() bool`

HasInvoiceMode returns a boolean if a field has been set.

### SetInvoiceModeNil

`func (o *TriggerResponse) SetInvoiceModeNil(b bool)`

 SetInvoiceModeNil sets the value for InvoiceMode to be an explicit nil

### UnsetInvoiceMode
`func (o *TriggerResponse) UnsetInvoiceMode()`

UnsetInvoiceMode ensures that no value is present for InvoiceMode, not even an explicit nil
### GetOutboxEventId

`func (o *TriggerResponse) GetOutboxEventId() string`

GetOutboxEventId returns the OutboxEventId field if non-nil, zero value otherwise.

### GetOutboxEventIdOk

`func (o *TriggerResponse) GetOutboxEventIdOk() (*string, bool)`

GetOutboxEventIdOk returns a tuple with the OutboxEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboxEventId

`func (o *TriggerResponse) SetOutboxEventId(v string)`

SetOutboxEventId sets OutboxEventId field to given value.

### HasOutboxEventId

`func (o *TriggerResponse) HasOutboxEventId() bool`

HasOutboxEventId returns a boolean if a field has been set.

### SetOutboxEventIdNil

`func (o *TriggerResponse) SetOutboxEventIdNil(b bool)`

 SetOutboxEventIdNil sets the value for OutboxEventId to be an explicit nil

### UnsetOutboxEventId
`func (o *TriggerResponse) UnsetOutboxEventId()`

UnsetOutboxEventId ensures that no value is present for OutboxEventId, not even an explicit nil
### GetCooldownUntil

`func (o *TriggerResponse) GetCooldownUntil() time.Time`

GetCooldownUntil returns the CooldownUntil field if non-nil, zero value otherwise.

### GetCooldownUntilOk

`func (o *TriggerResponse) GetCooldownUntilOk() (*time.Time, bool)`

GetCooldownUntilOk returns a tuple with the CooldownUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCooldownUntil

`func (o *TriggerResponse) SetCooldownUntil(v time.Time)`

SetCooldownUntil sets CooldownUntil field to given value.

### HasCooldownUntil

`func (o *TriggerResponse) HasCooldownUntil() bool`

HasCooldownUntil returns a boolean if a field has been set.

### SetCooldownUntilNil

`func (o *TriggerResponse) SetCooldownUntilNil(b bool)`

 SetCooldownUntilNil sets the value for CooldownUntil to be an explicit nil

### UnsetCooldownUntil
`func (o *TriggerResponse) UnsetCooldownUntil()`

UnsetCooldownUntil ensures that no value is present for CooldownUntil, not even an explicit nil
### GetWalletState

`func (o *TriggerResponse) GetWalletState() string`

GetWalletState returns the WalletState field if non-nil, zero value otherwise.

### GetWalletStateOk

`func (o *TriggerResponse) GetWalletStateOk() (*string, bool)`

GetWalletStateOk returns a tuple with the WalletState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletState

`func (o *TriggerResponse) SetWalletState(v string)`

SetWalletState sets WalletState field to given value.

### HasWalletState

`func (o *TriggerResponse) HasWalletState() bool`

HasWalletState returns a boolean if a field has been set.

### SetWalletStateNil

`func (o *TriggerResponse) SetWalletStateNil(b bool)`

 SetWalletStateNil sets the value for WalletState to be an explicit nil

### UnsetWalletState
`func (o *TriggerResponse) UnsetWalletState()`

UnsetWalletState ensures that no value is present for WalletState, not even an explicit nil
### GetAvailableBalanceMicros

`func (o *TriggerResponse) GetAvailableBalanceMicros() int32`

GetAvailableBalanceMicros returns the AvailableBalanceMicros field if non-nil, zero value otherwise.

### GetAvailableBalanceMicrosOk

`func (o *TriggerResponse) GetAvailableBalanceMicrosOk() (*int32, bool)`

GetAvailableBalanceMicrosOk returns a tuple with the AvailableBalanceMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBalanceMicros

`func (o *TriggerResponse) SetAvailableBalanceMicros(v int32)`

SetAvailableBalanceMicros sets AvailableBalanceMicros field to given value.

### HasAvailableBalanceMicros

`func (o *TriggerResponse) HasAvailableBalanceMicros() bool`

HasAvailableBalanceMicros returns a boolean if a field has been set.

### SetAvailableBalanceMicrosNil

`func (o *TriggerResponse) SetAvailableBalanceMicrosNil(b bool)`

 SetAvailableBalanceMicrosNil sets the value for AvailableBalanceMicros to be an explicit nil

### UnsetAvailableBalanceMicros
`func (o *TriggerResponse) UnsetAvailableBalanceMicros()`

UnsetAvailableBalanceMicros ensures that no value is present for AvailableBalanceMicros, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


