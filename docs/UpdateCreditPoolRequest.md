# UpdateCreditPoolRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Pool name | [optional] 
**Currency** | Pointer to **string** | Currency code (ISO 4217) | [optional] 
**NotificationEmails** | Pointer to **[]string** | Email addresses for pool-level alerts | [optional] 
**AtCapBehavior** | Pointer to **string** | Default wallet policy when balance hits cap: SOFT_BORROW | HARD_BUDGET | NOTIFY_ONLY | [optional] 
**ThresholdPolicy** | Pointer to **string** | Named alert-threshold preset (e.g. Conservative, Standard, Aggressive) | [optional] 
**CreditLabel** | Pointer to **string** | Credit unit label (e.g. credits, tokens, points) | [optional] 

## Methods

### NewUpdateCreditPoolRequest

`func NewUpdateCreditPoolRequest() *UpdateCreditPoolRequest`

NewUpdateCreditPoolRequest instantiates a new UpdateCreditPoolRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCreditPoolRequestWithDefaults

`func NewUpdateCreditPoolRequestWithDefaults() *UpdateCreditPoolRequest`

NewUpdateCreditPoolRequestWithDefaults instantiates a new UpdateCreditPoolRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateCreditPoolRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCreditPoolRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCreditPoolRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateCreditPoolRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCurrency

`func (o *UpdateCreditPoolRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UpdateCreditPoolRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UpdateCreditPoolRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UpdateCreditPoolRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetNotificationEmails

`func (o *UpdateCreditPoolRequest) GetNotificationEmails() []string`

GetNotificationEmails returns the NotificationEmails field if non-nil, zero value otherwise.

### GetNotificationEmailsOk

`func (o *UpdateCreditPoolRequest) GetNotificationEmailsOk() (*[]string, bool)`

GetNotificationEmailsOk returns a tuple with the NotificationEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationEmails

`func (o *UpdateCreditPoolRequest) SetNotificationEmails(v []string)`

SetNotificationEmails sets NotificationEmails field to given value.

### HasNotificationEmails

`func (o *UpdateCreditPoolRequest) HasNotificationEmails() bool`

HasNotificationEmails returns a boolean if a field has been set.

### GetAtCapBehavior

`func (o *UpdateCreditPoolRequest) GetAtCapBehavior() string`

GetAtCapBehavior returns the AtCapBehavior field if non-nil, zero value otherwise.

### GetAtCapBehaviorOk

`func (o *UpdateCreditPoolRequest) GetAtCapBehaviorOk() (*string, bool)`

GetAtCapBehaviorOk returns a tuple with the AtCapBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtCapBehavior

`func (o *UpdateCreditPoolRequest) SetAtCapBehavior(v string)`

SetAtCapBehavior sets AtCapBehavior field to given value.

### HasAtCapBehavior

`func (o *UpdateCreditPoolRequest) HasAtCapBehavior() bool`

HasAtCapBehavior returns a boolean if a field has been set.

### GetThresholdPolicy

`func (o *UpdateCreditPoolRequest) GetThresholdPolicy() string`

GetThresholdPolicy returns the ThresholdPolicy field if non-nil, zero value otherwise.

### GetThresholdPolicyOk

`func (o *UpdateCreditPoolRequest) GetThresholdPolicyOk() (*string, bool)`

GetThresholdPolicyOk returns a tuple with the ThresholdPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdPolicy

`func (o *UpdateCreditPoolRequest) SetThresholdPolicy(v string)`

SetThresholdPolicy sets ThresholdPolicy field to given value.

### HasThresholdPolicy

`func (o *UpdateCreditPoolRequest) HasThresholdPolicy() bool`

HasThresholdPolicy returns a boolean if a field has been set.

### GetCreditLabel

`func (o *UpdateCreditPoolRequest) GetCreditLabel() string`

GetCreditLabel returns the CreditLabel field if non-nil, zero value otherwise.

### GetCreditLabelOk

`func (o *UpdateCreditPoolRequest) GetCreditLabelOk() (*string, bool)`

GetCreditLabelOk returns a tuple with the CreditLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLabel

`func (o *UpdateCreditPoolRequest) SetCreditLabel(v string)`

SetCreditLabel sets CreditLabel field to given value.

### HasCreditLabel

`func (o *UpdateCreditPoolRequest) HasCreditLabel() bool`

HasCreditLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


