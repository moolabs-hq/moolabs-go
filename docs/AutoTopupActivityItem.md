# AutoTopupActivityItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**RuleId** | Pointer to **string** |  | [optional] 
**WalletId** | **string** |  | 
**TriggeredAt** | **time.Time** |  | 
**TriggeredBalanceMicros** | Pointer to **int32** |  | [optional] 
**TopupAmountMicros** | **int32** |  | 
**Status** | **string** |  | 
**ReferenceId** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 

## Methods

### NewAutoTopupActivityItem

`func NewAutoTopupActivityItem(id string, walletId string, triggeredAt time.Time, topupAmountMicros int32, status string, ) *AutoTopupActivityItem`

NewAutoTopupActivityItem instantiates a new AutoTopupActivityItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoTopupActivityItemWithDefaults

`func NewAutoTopupActivityItemWithDefaults() *AutoTopupActivityItem`

NewAutoTopupActivityItemWithDefaults instantiates a new AutoTopupActivityItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutoTopupActivityItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoTopupActivityItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoTopupActivityItem) SetId(v string)`

SetId sets Id field to given value.


### GetRuleId

`func (o *AutoTopupActivityItem) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *AutoTopupActivityItem) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *AutoTopupActivityItem) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *AutoTopupActivityItem) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetWalletId

`func (o *AutoTopupActivityItem) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *AutoTopupActivityItem) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *AutoTopupActivityItem) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetTriggeredAt

`func (o *AutoTopupActivityItem) GetTriggeredAt() time.Time`

GetTriggeredAt returns the TriggeredAt field if non-nil, zero value otherwise.

### GetTriggeredAtOk

`func (o *AutoTopupActivityItem) GetTriggeredAtOk() (*time.Time, bool)`

GetTriggeredAtOk returns a tuple with the TriggeredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredAt

`func (o *AutoTopupActivityItem) SetTriggeredAt(v time.Time)`

SetTriggeredAt sets TriggeredAt field to given value.


### GetTriggeredBalanceMicros

`func (o *AutoTopupActivityItem) GetTriggeredBalanceMicros() int32`

GetTriggeredBalanceMicros returns the TriggeredBalanceMicros field if non-nil, zero value otherwise.

### GetTriggeredBalanceMicrosOk

`func (o *AutoTopupActivityItem) GetTriggeredBalanceMicrosOk() (*int32, bool)`

GetTriggeredBalanceMicrosOk returns a tuple with the TriggeredBalanceMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredBalanceMicros

`func (o *AutoTopupActivityItem) SetTriggeredBalanceMicros(v int32)`

SetTriggeredBalanceMicros sets TriggeredBalanceMicros field to given value.

### HasTriggeredBalanceMicros

`func (o *AutoTopupActivityItem) HasTriggeredBalanceMicros() bool`

HasTriggeredBalanceMicros returns a boolean if a field has been set.

### GetTopupAmountMicros

`func (o *AutoTopupActivityItem) GetTopupAmountMicros() int32`

GetTopupAmountMicros returns the TopupAmountMicros field if non-nil, zero value otherwise.

### GetTopupAmountMicrosOk

`func (o *AutoTopupActivityItem) GetTopupAmountMicrosOk() (*int32, bool)`

GetTopupAmountMicrosOk returns a tuple with the TopupAmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopupAmountMicros

`func (o *AutoTopupActivityItem) SetTopupAmountMicros(v int32)`

SetTopupAmountMicros sets TopupAmountMicros field to given value.


### GetStatus

`func (o *AutoTopupActivityItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AutoTopupActivityItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AutoTopupActivityItem) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetReferenceId

`func (o *AutoTopupActivityItem) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *AutoTopupActivityItem) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *AutoTopupActivityItem) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *AutoTopupActivityItem) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetErrorMessage

`func (o *AutoTopupActivityItem) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *AutoTopupActivityItem) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *AutoTopupActivityItem) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *AutoTopupActivityItem) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


