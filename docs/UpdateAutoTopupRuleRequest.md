# UpdateAutoTopupRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TriggerType** | Pointer to **string** |  | [optional] 
**TriggerValue** | Pointer to **string** |  | [optional] 
**TopupAmountMicros** | Pointer to **int32** |  | [optional] 
**TopupCooldownSeconds** | Pointer to **int32** |  | [optional] 
**MaxTopupsPerDay** | Pointer to **int32** |  | [optional] 
**PaymentMethodRef** | Pointer to **string** |  | [optional] 
**InvoiceMode** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateAutoTopupRuleRequest

`func NewUpdateAutoTopupRuleRequest() *UpdateAutoTopupRuleRequest`

NewUpdateAutoTopupRuleRequest instantiates a new UpdateAutoTopupRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAutoTopupRuleRequestWithDefaults

`func NewUpdateAutoTopupRuleRequestWithDefaults() *UpdateAutoTopupRuleRequest`

NewUpdateAutoTopupRuleRequestWithDefaults instantiates a new UpdateAutoTopupRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTriggerType

`func (o *UpdateAutoTopupRuleRequest) GetTriggerType() string`

GetTriggerType returns the TriggerType field if non-nil, zero value otherwise.

### GetTriggerTypeOk

`func (o *UpdateAutoTopupRuleRequest) GetTriggerTypeOk() (*string, bool)`

GetTriggerTypeOk returns a tuple with the TriggerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerType

`func (o *UpdateAutoTopupRuleRequest) SetTriggerType(v string)`

SetTriggerType sets TriggerType field to given value.

### HasTriggerType

`func (o *UpdateAutoTopupRuleRequest) HasTriggerType() bool`

HasTriggerType returns a boolean if a field has been set.

### GetTriggerValue

`func (o *UpdateAutoTopupRuleRequest) GetTriggerValue() string`

GetTriggerValue returns the TriggerValue field if non-nil, zero value otherwise.

### GetTriggerValueOk

`func (o *UpdateAutoTopupRuleRequest) GetTriggerValueOk() (*string, bool)`

GetTriggerValueOk returns a tuple with the TriggerValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerValue

`func (o *UpdateAutoTopupRuleRequest) SetTriggerValue(v string)`

SetTriggerValue sets TriggerValue field to given value.

### HasTriggerValue

`func (o *UpdateAutoTopupRuleRequest) HasTriggerValue() bool`

HasTriggerValue returns a boolean if a field has been set.

### GetTopupAmountMicros

`func (o *UpdateAutoTopupRuleRequest) GetTopupAmountMicros() int32`

GetTopupAmountMicros returns the TopupAmountMicros field if non-nil, zero value otherwise.

### GetTopupAmountMicrosOk

`func (o *UpdateAutoTopupRuleRequest) GetTopupAmountMicrosOk() (*int32, bool)`

GetTopupAmountMicrosOk returns a tuple with the TopupAmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopupAmountMicros

`func (o *UpdateAutoTopupRuleRequest) SetTopupAmountMicros(v int32)`

SetTopupAmountMicros sets TopupAmountMicros field to given value.

### HasTopupAmountMicros

`func (o *UpdateAutoTopupRuleRequest) HasTopupAmountMicros() bool`

HasTopupAmountMicros returns a boolean if a field has been set.

### GetTopupCooldownSeconds

`func (o *UpdateAutoTopupRuleRequest) GetTopupCooldownSeconds() int32`

GetTopupCooldownSeconds returns the TopupCooldownSeconds field if non-nil, zero value otherwise.

### GetTopupCooldownSecondsOk

`func (o *UpdateAutoTopupRuleRequest) GetTopupCooldownSecondsOk() (*int32, bool)`

GetTopupCooldownSecondsOk returns a tuple with the TopupCooldownSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopupCooldownSeconds

`func (o *UpdateAutoTopupRuleRequest) SetTopupCooldownSeconds(v int32)`

SetTopupCooldownSeconds sets TopupCooldownSeconds field to given value.

### HasTopupCooldownSeconds

`func (o *UpdateAutoTopupRuleRequest) HasTopupCooldownSeconds() bool`

HasTopupCooldownSeconds returns a boolean if a field has been set.

### GetMaxTopupsPerDay

`func (o *UpdateAutoTopupRuleRequest) GetMaxTopupsPerDay() int32`

GetMaxTopupsPerDay returns the MaxTopupsPerDay field if non-nil, zero value otherwise.

### GetMaxTopupsPerDayOk

`func (o *UpdateAutoTopupRuleRequest) GetMaxTopupsPerDayOk() (*int32, bool)`

GetMaxTopupsPerDayOk returns a tuple with the MaxTopupsPerDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTopupsPerDay

`func (o *UpdateAutoTopupRuleRequest) SetMaxTopupsPerDay(v int32)`

SetMaxTopupsPerDay sets MaxTopupsPerDay field to given value.

### HasMaxTopupsPerDay

`func (o *UpdateAutoTopupRuleRequest) HasMaxTopupsPerDay() bool`

HasMaxTopupsPerDay returns a boolean if a field has been set.

### GetPaymentMethodRef

`func (o *UpdateAutoTopupRuleRequest) GetPaymentMethodRef() string`

GetPaymentMethodRef returns the PaymentMethodRef field if non-nil, zero value otherwise.

### GetPaymentMethodRefOk

`func (o *UpdateAutoTopupRuleRequest) GetPaymentMethodRefOk() (*string, bool)`

GetPaymentMethodRefOk returns a tuple with the PaymentMethodRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodRef

`func (o *UpdateAutoTopupRuleRequest) SetPaymentMethodRef(v string)`

SetPaymentMethodRef sets PaymentMethodRef field to given value.

### HasPaymentMethodRef

`func (o *UpdateAutoTopupRuleRequest) HasPaymentMethodRef() bool`

HasPaymentMethodRef returns a boolean if a field has been set.

### GetInvoiceMode

`func (o *UpdateAutoTopupRuleRequest) GetInvoiceMode() string`

GetInvoiceMode returns the InvoiceMode field if non-nil, zero value otherwise.

### GetInvoiceModeOk

`func (o *UpdateAutoTopupRuleRequest) GetInvoiceModeOk() (*string, bool)`

GetInvoiceModeOk returns a tuple with the InvoiceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceMode

`func (o *UpdateAutoTopupRuleRequest) SetInvoiceMode(v string)`

SetInvoiceMode sets InvoiceMode field to given value.

### HasInvoiceMode

`func (o *UpdateAutoTopupRuleRequest) HasInvoiceMode() bool`

HasInvoiceMode returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateAutoTopupRuleRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateAutoTopupRuleRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateAutoTopupRuleRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateAutoTopupRuleRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


