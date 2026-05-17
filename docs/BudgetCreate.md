# BudgetCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BudgetName** | **string** |  | 
**ScopeType** | **string** | &#39;tenant&#39;, &#39;customer&#39;, &#39;feature&#39;, &#39;cloud_provider&#39; | 
**ScopeId** | Pointer to **string** |  | [optional] 
**PeriodType** | Pointer to **string** |  | [optional] [default to "monthly"]
**BudgetAmount** | [**BudgetAmount**](BudgetAmount.md) |  | 
**Currency** | Pointer to **string** |  | [optional] [default to "USD"]
**AlertThresholds** | Pointer to **[]int32** |  | [optional] 
**NotifyChannels** | Pointer to **[]string** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] [default to true]
**CreatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewBudgetCreate

`func NewBudgetCreate(budgetName string, scopeType string, budgetAmount BudgetAmount, ) *BudgetCreate`

NewBudgetCreate instantiates a new BudgetCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBudgetCreateWithDefaults

`func NewBudgetCreateWithDefaults() *BudgetCreate`

NewBudgetCreateWithDefaults instantiates a new BudgetCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudgetName

`func (o *BudgetCreate) GetBudgetName() string`

GetBudgetName returns the BudgetName field if non-nil, zero value otherwise.

### GetBudgetNameOk

`func (o *BudgetCreate) GetBudgetNameOk() (*string, bool)`

GetBudgetNameOk returns a tuple with the BudgetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetName

`func (o *BudgetCreate) SetBudgetName(v string)`

SetBudgetName sets BudgetName field to given value.


### GetScopeType

`func (o *BudgetCreate) GetScopeType() string`

GetScopeType returns the ScopeType field if non-nil, zero value otherwise.

### GetScopeTypeOk

`func (o *BudgetCreate) GetScopeTypeOk() (*string, bool)`

GetScopeTypeOk returns a tuple with the ScopeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeType

`func (o *BudgetCreate) SetScopeType(v string)`

SetScopeType sets ScopeType field to given value.


### GetScopeId

`func (o *BudgetCreate) GetScopeId() string`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *BudgetCreate) GetScopeIdOk() (*string, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *BudgetCreate) SetScopeId(v string)`

SetScopeId sets ScopeId field to given value.

### HasScopeId

`func (o *BudgetCreate) HasScopeId() bool`

HasScopeId returns a boolean if a field has been set.

### GetPeriodType

`func (o *BudgetCreate) GetPeriodType() string`

GetPeriodType returns the PeriodType field if non-nil, zero value otherwise.

### GetPeriodTypeOk

`func (o *BudgetCreate) GetPeriodTypeOk() (*string, bool)`

GetPeriodTypeOk returns a tuple with the PeriodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodType

`func (o *BudgetCreate) SetPeriodType(v string)`

SetPeriodType sets PeriodType field to given value.

### HasPeriodType

`func (o *BudgetCreate) HasPeriodType() bool`

HasPeriodType returns a boolean if a field has been set.

### GetBudgetAmount

`func (o *BudgetCreate) GetBudgetAmount() BudgetAmount`

GetBudgetAmount returns the BudgetAmount field if non-nil, zero value otherwise.

### GetBudgetAmountOk

`func (o *BudgetCreate) GetBudgetAmountOk() (*BudgetAmount, bool)`

GetBudgetAmountOk returns a tuple with the BudgetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetAmount

`func (o *BudgetCreate) SetBudgetAmount(v BudgetAmount)`

SetBudgetAmount sets BudgetAmount field to given value.


### GetCurrency

`func (o *BudgetCreate) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BudgetCreate) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BudgetCreate) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BudgetCreate) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetAlertThresholds

`func (o *BudgetCreate) GetAlertThresholds() []int32`

GetAlertThresholds returns the AlertThresholds field if non-nil, zero value otherwise.

### GetAlertThresholdsOk

`func (o *BudgetCreate) GetAlertThresholdsOk() (*[]int32, bool)`

GetAlertThresholdsOk returns a tuple with the AlertThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertThresholds

`func (o *BudgetCreate) SetAlertThresholds(v []int32)`

SetAlertThresholds sets AlertThresholds field to given value.

### HasAlertThresholds

`func (o *BudgetCreate) HasAlertThresholds() bool`

HasAlertThresholds returns a boolean if a field has been set.

### GetNotifyChannels

`func (o *BudgetCreate) GetNotifyChannels() []string`

GetNotifyChannels returns the NotifyChannels field if non-nil, zero value otherwise.

### GetNotifyChannelsOk

`func (o *BudgetCreate) GetNotifyChannelsOk() (*[]string, bool)`

GetNotifyChannelsOk returns a tuple with the NotifyChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyChannels

`func (o *BudgetCreate) SetNotifyChannels(v []string)`

SetNotifyChannels sets NotifyChannels field to given value.

### HasNotifyChannels

`func (o *BudgetCreate) HasNotifyChannels() bool`

HasNotifyChannels returns a boolean if a field has been set.

### GetIsActive

`func (o *BudgetCreate) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *BudgetCreate) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *BudgetCreate) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *BudgetCreate) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetCreatedBy

`func (o *BudgetCreate) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BudgetCreate) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BudgetCreate) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *BudgetCreate) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


