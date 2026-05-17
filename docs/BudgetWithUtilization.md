# BudgetWithUtilization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**BudgetName** | **string** |  | 
**ScopeType** | **string** |  | 
**ScopeId** | **string** |  | 
**PeriodType** | **string** |  | 
**BudgetAmount** | **string** |  | 
**Currency** | **string** |  | 
**AlertThresholds** | **[]int32** |  | 
**NotifyChannels** | **[]interface{}** |  | 
**IsActive** | **bool** |  | 
**CreatedBy** | **string** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**CurrentSpend** | Pointer to **string** |  | [optional] 
**UtilizationPct** | Pointer to **string** |  | [optional] 
**LoadedMarginAvailable** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewBudgetWithUtilization

`func NewBudgetWithUtilization(id string, tenantId string, budgetName string, scopeType string, scopeId string, periodType string, budgetAmount string, currency string, alertThresholds []int32, notifyChannels []interface{}, isActive bool, createdBy string, createdAt string, updatedAt string, ) *BudgetWithUtilization`

NewBudgetWithUtilization instantiates a new BudgetWithUtilization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBudgetWithUtilizationWithDefaults

`func NewBudgetWithUtilizationWithDefaults() *BudgetWithUtilization`

NewBudgetWithUtilizationWithDefaults instantiates a new BudgetWithUtilization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BudgetWithUtilization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BudgetWithUtilization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BudgetWithUtilization) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *BudgetWithUtilization) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *BudgetWithUtilization) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *BudgetWithUtilization) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetBudgetName

`func (o *BudgetWithUtilization) GetBudgetName() string`

GetBudgetName returns the BudgetName field if non-nil, zero value otherwise.

### GetBudgetNameOk

`func (o *BudgetWithUtilization) GetBudgetNameOk() (*string, bool)`

GetBudgetNameOk returns a tuple with the BudgetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetName

`func (o *BudgetWithUtilization) SetBudgetName(v string)`

SetBudgetName sets BudgetName field to given value.


### GetScopeType

`func (o *BudgetWithUtilization) GetScopeType() string`

GetScopeType returns the ScopeType field if non-nil, zero value otherwise.

### GetScopeTypeOk

`func (o *BudgetWithUtilization) GetScopeTypeOk() (*string, bool)`

GetScopeTypeOk returns a tuple with the ScopeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeType

`func (o *BudgetWithUtilization) SetScopeType(v string)`

SetScopeType sets ScopeType field to given value.


### GetScopeId

`func (o *BudgetWithUtilization) GetScopeId() string`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *BudgetWithUtilization) GetScopeIdOk() (*string, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *BudgetWithUtilization) SetScopeId(v string)`

SetScopeId sets ScopeId field to given value.


### GetPeriodType

`func (o *BudgetWithUtilization) GetPeriodType() string`

GetPeriodType returns the PeriodType field if non-nil, zero value otherwise.

### GetPeriodTypeOk

`func (o *BudgetWithUtilization) GetPeriodTypeOk() (*string, bool)`

GetPeriodTypeOk returns a tuple with the PeriodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodType

`func (o *BudgetWithUtilization) SetPeriodType(v string)`

SetPeriodType sets PeriodType field to given value.


### GetBudgetAmount

`func (o *BudgetWithUtilization) GetBudgetAmount() string`

GetBudgetAmount returns the BudgetAmount field if non-nil, zero value otherwise.

### GetBudgetAmountOk

`func (o *BudgetWithUtilization) GetBudgetAmountOk() (*string, bool)`

GetBudgetAmountOk returns a tuple with the BudgetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetAmount

`func (o *BudgetWithUtilization) SetBudgetAmount(v string)`

SetBudgetAmount sets BudgetAmount field to given value.


### GetCurrency

`func (o *BudgetWithUtilization) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BudgetWithUtilization) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BudgetWithUtilization) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetAlertThresholds

`func (o *BudgetWithUtilization) GetAlertThresholds() []int32`

GetAlertThresholds returns the AlertThresholds field if non-nil, zero value otherwise.

### GetAlertThresholdsOk

`func (o *BudgetWithUtilization) GetAlertThresholdsOk() (*[]int32, bool)`

GetAlertThresholdsOk returns a tuple with the AlertThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertThresholds

`func (o *BudgetWithUtilization) SetAlertThresholds(v []int32)`

SetAlertThresholds sets AlertThresholds field to given value.


### GetNotifyChannels

`func (o *BudgetWithUtilization) GetNotifyChannels() []interface{}`

GetNotifyChannels returns the NotifyChannels field if non-nil, zero value otherwise.

### GetNotifyChannelsOk

`func (o *BudgetWithUtilization) GetNotifyChannelsOk() (*[]interface{}, bool)`

GetNotifyChannelsOk returns a tuple with the NotifyChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyChannels

`func (o *BudgetWithUtilization) SetNotifyChannels(v []interface{})`

SetNotifyChannels sets NotifyChannels field to given value.


### GetIsActive

`func (o *BudgetWithUtilization) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *BudgetWithUtilization) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *BudgetWithUtilization) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetCreatedBy

`func (o *BudgetWithUtilization) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BudgetWithUtilization) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BudgetWithUtilization) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatedAt

`func (o *BudgetWithUtilization) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BudgetWithUtilization) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BudgetWithUtilization) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BudgetWithUtilization) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BudgetWithUtilization) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BudgetWithUtilization) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCurrentSpend

`func (o *BudgetWithUtilization) GetCurrentSpend() string`

GetCurrentSpend returns the CurrentSpend field if non-nil, zero value otherwise.

### GetCurrentSpendOk

`func (o *BudgetWithUtilization) GetCurrentSpendOk() (*string, bool)`

GetCurrentSpendOk returns a tuple with the CurrentSpend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSpend

`func (o *BudgetWithUtilization) SetCurrentSpend(v string)`

SetCurrentSpend sets CurrentSpend field to given value.

### HasCurrentSpend

`func (o *BudgetWithUtilization) HasCurrentSpend() bool`

HasCurrentSpend returns a boolean if a field has been set.

### GetUtilizationPct

`func (o *BudgetWithUtilization) GetUtilizationPct() string`

GetUtilizationPct returns the UtilizationPct field if non-nil, zero value otherwise.

### GetUtilizationPctOk

`func (o *BudgetWithUtilization) GetUtilizationPctOk() (*string, bool)`

GetUtilizationPctOk returns a tuple with the UtilizationPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizationPct

`func (o *BudgetWithUtilization) SetUtilizationPct(v string)`

SetUtilizationPct sets UtilizationPct field to given value.

### HasUtilizationPct

`func (o *BudgetWithUtilization) HasUtilizationPct() bool`

HasUtilizationPct returns a boolean if a field has been set.

### GetLoadedMarginAvailable

`func (o *BudgetWithUtilization) GetLoadedMarginAvailable() bool`

GetLoadedMarginAvailable returns the LoadedMarginAvailable field if non-nil, zero value otherwise.

### GetLoadedMarginAvailableOk

`func (o *BudgetWithUtilization) GetLoadedMarginAvailableOk() (*bool, bool)`

GetLoadedMarginAvailableOk returns a tuple with the LoadedMarginAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadedMarginAvailable

`func (o *BudgetWithUtilization) SetLoadedMarginAvailable(v bool)`

SetLoadedMarginAvailable sets LoadedMarginAvailable field to given value.

### HasLoadedMarginAvailable

`func (o *BudgetWithUtilization) HasLoadedMarginAvailable() bool`

HasLoadedMarginAvailable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


