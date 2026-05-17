# BudgetUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BudgetName** | Pointer to **string** |  | [optional] 
**ScopeType** | Pointer to **string** |  | [optional] 
**ScopeId** | Pointer to **string** |  | [optional] 
**PeriodType** | Pointer to **string** |  | [optional] 
**BudgetAmount** | Pointer to [**NullableBudgetAmount1**](BudgetAmount1.md) |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**AlertThresholds** | Pointer to **[]int32** |  | [optional] 
**NotifyChannels** | Pointer to **[]string** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 

## Methods

### NewBudgetUpdate

`func NewBudgetUpdate() *BudgetUpdate`

NewBudgetUpdate instantiates a new BudgetUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBudgetUpdateWithDefaults

`func NewBudgetUpdateWithDefaults() *BudgetUpdate`

NewBudgetUpdateWithDefaults instantiates a new BudgetUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudgetName

`func (o *BudgetUpdate) GetBudgetName() string`

GetBudgetName returns the BudgetName field if non-nil, zero value otherwise.

### GetBudgetNameOk

`func (o *BudgetUpdate) GetBudgetNameOk() (*string, bool)`

GetBudgetNameOk returns a tuple with the BudgetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetName

`func (o *BudgetUpdate) SetBudgetName(v string)`

SetBudgetName sets BudgetName field to given value.

### HasBudgetName

`func (o *BudgetUpdate) HasBudgetName() bool`

HasBudgetName returns a boolean if a field has been set.

### GetScopeType

`func (o *BudgetUpdate) GetScopeType() string`

GetScopeType returns the ScopeType field if non-nil, zero value otherwise.

### GetScopeTypeOk

`func (o *BudgetUpdate) GetScopeTypeOk() (*string, bool)`

GetScopeTypeOk returns a tuple with the ScopeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeType

`func (o *BudgetUpdate) SetScopeType(v string)`

SetScopeType sets ScopeType field to given value.

### HasScopeType

`func (o *BudgetUpdate) HasScopeType() bool`

HasScopeType returns a boolean if a field has been set.

### GetScopeId

`func (o *BudgetUpdate) GetScopeId() string`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *BudgetUpdate) GetScopeIdOk() (*string, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *BudgetUpdate) SetScopeId(v string)`

SetScopeId sets ScopeId field to given value.

### HasScopeId

`func (o *BudgetUpdate) HasScopeId() bool`

HasScopeId returns a boolean if a field has been set.

### GetPeriodType

`func (o *BudgetUpdate) GetPeriodType() string`

GetPeriodType returns the PeriodType field if non-nil, zero value otherwise.

### GetPeriodTypeOk

`func (o *BudgetUpdate) GetPeriodTypeOk() (*string, bool)`

GetPeriodTypeOk returns a tuple with the PeriodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodType

`func (o *BudgetUpdate) SetPeriodType(v string)`

SetPeriodType sets PeriodType field to given value.

### HasPeriodType

`func (o *BudgetUpdate) HasPeriodType() bool`

HasPeriodType returns a boolean if a field has been set.

### GetBudgetAmount

`func (o *BudgetUpdate) GetBudgetAmount() BudgetAmount1`

GetBudgetAmount returns the BudgetAmount field if non-nil, zero value otherwise.

### GetBudgetAmountOk

`func (o *BudgetUpdate) GetBudgetAmountOk() (*BudgetAmount1, bool)`

GetBudgetAmountOk returns a tuple with the BudgetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetAmount

`func (o *BudgetUpdate) SetBudgetAmount(v BudgetAmount1)`

SetBudgetAmount sets BudgetAmount field to given value.

### HasBudgetAmount

`func (o *BudgetUpdate) HasBudgetAmount() bool`

HasBudgetAmount returns a boolean if a field has been set.

### SetBudgetAmountNil

`func (o *BudgetUpdate) SetBudgetAmountNil(b bool)`

 SetBudgetAmountNil sets the value for BudgetAmount to be an explicit nil

### UnsetBudgetAmount
`func (o *BudgetUpdate) UnsetBudgetAmount()`

UnsetBudgetAmount ensures that no value is present for BudgetAmount, not even an explicit nil
### GetCurrency

`func (o *BudgetUpdate) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BudgetUpdate) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BudgetUpdate) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BudgetUpdate) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetAlertThresholds

`func (o *BudgetUpdate) GetAlertThresholds() []int32`

GetAlertThresholds returns the AlertThresholds field if non-nil, zero value otherwise.

### GetAlertThresholdsOk

`func (o *BudgetUpdate) GetAlertThresholdsOk() (*[]int32, bool)`

GetAlertThresholdsOk returns a tuple with the AlertThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertThresholds

`func (o *BudgetUpdate) SetAlertThresholds(v []int32)`

SetAlertThresholds sets AlertThresholds field to given value.

### HasAlertThresholds

`func (o *BudgetUpdate) HasAlertThresholds() bool`

HasAlertThresholds returns a boolean if a field has been set.

### GetNotifyChannels

`func (o *BudgetUpdate) GetNotifyChannels() []string`

GetNotifyChannels returns the NotifyChannels field if non-nil, zero value otherwise.

### GetNotifyChannelsOk

`func (o *BudgetUpdate) GetNotifyChannelsOk() (*[]string, bool)`

GetNotifyChannelsOk returns a tuple with the NotifyChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyChannels

`func (o *BudgetUpdate) SetNotifyChannels(v []string)`

SetNotifyChannels sets NotifyChannels field to given value.

### HasNotifyChannels

`func (o *BudgetUpdate) HasNotifyChannels() bool`

HasNotifyChannels returns a boolean if a field has been set.

### GetIsActive

`func (o *BudgetUpdate) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *BudgetUpdate) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *BudgetUpdate) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *BudgetUpdate) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


