# AppApiV1AutoTopupRouterTriggerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Triggered** | **bool** |  | 
**Reason** | Pointer to **string** |  | [optional] 
**RuleId** | Pointer to **string** |  | [optional] 
**WalletId** | Pointer to **string** |  | [optional] 
**TopupAmountMicros** | Pointer to **int32** |  | [optional] 
**PaymentMethodRef** | Pointer to **string** |  | [optional] 
**InvoiceMode** | Pointer to **string** |  | [optional] 
**OutboxEventId** | Pointer to **string** |  | [optional] 
**CooldownUntil** | Pointer to **time.Time** |  | [optional] 
**WalletState** | Pointer to **string** |  | [optional] 
**AvailableBalanceMicros** | Pointer to **int32** |  | [optional] 

## Methods

### NewAppApiV1AutoTopupRouterTriggerResponse

`func NewAppApiV1AutoTopupRouterTriggerResponse(triggered bool, ) *AppApiV1AutoTopupRouterTriggerResponse`

NewAppApiV1AutoTopupRouterTriggerResponse instantiates a new AppApiV1AutoTopupRouterTriggerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppApiV1AutoTopupRouterTriggerResponseWithDefaults

`func NewAppApiV1AutoTopupRouterTriggerResponseWithDefaults() *AppApiV1AutoTopupRouterTriggerResponse`

NewAppApiV1AutoTopupRouterTriggerResponseWithDefaults instantiates a new AppApiV1AutoTopupRouterTriggerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTriggered

`func (o *AppApiV1AutoTopupRouterTriggerResponse) GetTriggered() bool`

GetTriggered returns the Triggered field if non-nil, zero value otherwise.

### GetTriggeredOk

`func (o *AppApiV1AutoTopupRouterTriggerResponse) GetTriggeredOk() (*bool, bool)`

GetTriggeredOk returns a tuple with the Triggered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggered

`func (o *AppApiV1AutoTopupRouterTriggerResponse) SetTriggered(v bool)`

SetTriggered sets Triggered field to given value.


### GetReason

`func (o *AppApiV1AutoTopupRouterTriggerResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AppApiV1AutoTopupRouterTriggerResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AppApiV1AutoTopupRouterTriggerResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AppApiV1AutoTopupRouterTriggerResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRuleId

`func (o *AppApiV1AutoTopupRouterTriggerResponse) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *AppApiV1AutoTopupRouterTriggerResponse) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *AppApiV1AutoTopupRouterTriggerResponse) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *AppApiV1AutoTopupRouterTriggerResponse) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetWalletId

`func (o *AppApiV1AutoTopupRouterTriggerResponse) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *AppApiV1AutoTopupRouterTriggerResponse) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *AppApiV1AutoTopupRouterTriggerResponse) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *AppApiV1AutoTopupRouterTriggerResponse) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.

### GetTopupAmountMicros

`func (o *AppApiV1AutoTopupRouterTriggerResponse) GetTopupAmountMicros() int32`

GetTopupAmountMicros returns the TopupAmountMicros field if non-nil, zero value otherwise.

### GetTopupAmountMicrosOk

`func (o *AppApiV1AutoTopupRouterTriggerResponse) GetTopupAmountMicrosOk() (*int32, bool)`

GetTopupAmountMicrosOk returns a tuple with the TopupAmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopupAmountMicros

`func (o *AppApiV1AutoTopupRouterTriggerResponse) SetTopupAmountMicros(v int32)`

SetTopupAmountMicros sets TopupAmountMicros field to given value.

### HasTopupAmountMicros

`func (o *AppApiV1AutoTopupRouterTriggerResponse) HasTopupAmountMicros() bool`

HasTopupAmountMicros returns a boolean if a field has been set.

### GetPaymentMethodRef

`func (o *AppApiV1AutoTopupRouterTriggerResponse) GetPaymentMethodRef() string`

GetPaymentMethodRef returns the PaymentMethodRef field if non-nil, zero value otherwise.

### GetPaymentMethodRefOk

`func (o *AppApiV1AutoTopupRouterTriggerResponse) GetPaymentMethodRefOk() (*string, bool)`

GetPaymentMethodRefOk returns a tuple with the PaymentMethodRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodRef

`func (o *AppApiV1AutoTopupRouterTriggerResponse) SetPaymentMethodRef(v string)`

SetPaymentMethodRef sets PaymentMethodRef field to given value.

### HasPaymentMethodRef

`func (o *AppApiV1AutoTopupRouterTriggerResponse) HasPaymentMethodRef() bool`

HasPaymentMethodRef returns a boolean if a field has been set.

### GetInvoiceMode

`func (o *AppApiV1AutoTopupRouterTriggerResponse) GetInvoiceMode() string`

GetInvoiceMode returns the InvoiceMode field if non-nil, zero value otherwise.

### GetInvoiceModeOk

`func (o *AppApiV1AutoTopupRouterTriggerResponse) GetInvoiceModeOk() (*string, bool)`

GetInvoiceModeOk returns a tuple with the InvoiceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceMode

`func (o *AppApiV1AutoTopupRouterTriggerResponse) SetInvoiceMode(v string)`

SetInvoiceMode sets InvoiceMode field to given value.

### HasInvoiceMode

`func (o *AppApiV1AutoTopupRouterTriggerResponse) HasInvoiceMode() bool`

HasInvoiceMode returns a boolean if a field has been set.

### GetOutboxEventId

`func (o *AppApiV1AutoTopupRouterTriggerResponse) GetOutboxEventId() string`

GetOutboxEventId returns the OutboxEventId field if non-nil, zero value otherwise.

### GetOutboxEventIdOk

`func (o *AppApiV1AutoTopupRouterTriggerResponse) GetOutboxEventIdOk() (*string, bool)`

GetOutboxEventIdOk returns a tuple with the OutboxEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboxEventId

`func (o *AppApiV1AutoTopupRouterTriggerResponse) SetOutboxEventId(v string)`

SetOutboxEventId sets OutboxEventId field to given value.

### HasOutboxEventId

`func (o *AppApiV1AutoTopupRouterTriggerResponse) HasOutboxEventId() bool`

HasOutboxEventId returns a boolean if a field has been set.

### GetCooldownUntil

`func (o *AppApiV1AutoTopupRouterTriggerResponse) GetCooldownUntil() time.Time`

GetCooldownUntil returns the CooldownUntil field if non-nil, zero value otherwise.

### GetCooldownUntilOk

`func (o *AppApiV1AutoTopupRouterTriggerResponse) GetCooldownUntilOk() (*time.Time, bool)`

GetCooldownUntilOk returns a tuple with the CooldownUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCooldownUntil

`func (o *AppApiV1AutoTopupRouterTriggerResponse) SetCooldownUntil(v time.Time)`

SetCooldownUntil sets CooldownUntil field to given value.

### HasCooldownUntil

`func (o *AppApiV1AutoTopupRouterTriggerResponse) HasCooldownUntil() bool`

HasCooldownUntil returns a boolean if a field has been set.

### GetWalletState

`func (o *AppApiV1AutoTopupRouterTriggerResponse) GetWalletState() string`

GetWalletState returns the WalletState field if non-nil, zero value otherwise.

### GetWalletStateOk

`func (o *AppApiV1AutoTopupRouterTriggerResponse) GetWalletStateOk() (*string, bool)`

GetWalletStateOk returns a tuple with the WalletState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletState

`func (o *AppApiV1AutoTopupRouterTriggerResponse) SetWalletState(v string)`

SetWalletState sets WalletState field to given value.

### HasWalletState

`func (o *AppApiV1AutoTopupRouterTriggerResponse) HasWalletState() bool`

HasWalletState returns a boolean if a field has been set.

### GetAvailableBalanceMicros

`func (o *AppApiV1AutoTopupRouterTriggerResponse) GetAvailableBalanceMicros() int32`

GetAvailableBalanceMicros returns the AvailableBalanceMicros field if non-nil, zero value otherwise.

### GetAvailableBalanceMicrosOk

`func (o *AppApiV1AutoTopupRouterTriggerResponse) GetAvailableBalanceMicrosOk() (*int32, bool)`

GetAvailableBalanceMicrosOk returns a tuple with the AvailableBalanceMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBalanceMicros

`func (o *AppApiV1AutoTopupRouterTriggerResponse) SetAvailableBalanceMicros(v int32)`

SetAvailableBalanceMicros sets AvailableBalanceMicros field to given value.

### HasAvailableBalanceMicros

`func (o *AppApiV1AutoTopupRouterTriggerResponse) HasAvailableBalanceMicros() bool`

HasAvailableBalanceMicros returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


