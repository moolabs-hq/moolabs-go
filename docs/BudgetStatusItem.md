# BudgetStatusItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BudgetId** | **string** |  | 
**BudgetName** | **string** |  | 
**ScopeType** | **string** |  | 
**ScopeId** | **string** |  | 
**BudgetAmount** | **string** |  | 
**Currency** | **string** |  | 
**CurrentSpend** | **string** |  | 
**UtilizationPct** | **string** |  | 
**IsActive** | **bool** |  | 
**ActiveAlerts** | **int32** |  | 

## Methods

### NewBudgetStatusItem

`func NewBudgetStatusItem(budgetId string, budgetName string, scopeType string, scopeId string, budgetAmount string, currency string, currentSpend string, utilizationPct string, isActive bool, activeAlerts int32, ) *BudgetStatusItem`

NewBudgetStatusItem instantiates a new BudgetStatusItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBudgetStatusItemWithDefaults

`func NewBudgetStatusItemWithDefaults() *BudgetStatusItem`

NewBudgetStatusItemWithDefaults instantiates a new BudgetStatusItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudgetId

`func (o *BudgetStatusItem) GetBudgetId() string`

GetBudgetId returns the BudgetId field if non-nil, zero value otherwise.

### GetBudgetIdOk

`func (o *BudgetStatusItem) GetBudgetIdOk() (*string, bool)`

GetBudgetIdOk returns a tuple with the BudgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetId

`func (o *BudgetStatusItem) SetBudgetId(v string)`

SetBudgetId sets BudgetId field to given value.


### GetBudgetName

`func (o *BudgetStatusItem) GetBudgetName() string`

GetBudgetName returns the BudgetName field if non-nil, zero value otherwise.

### GetBudgetNameOk

`func (o *BudgetStatusItem) GetBudgetNameOk() (*string, bool)`

GetBudgetNameOk returns a tuple with the BudgetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetName

`func (o *BudgetStatusItem) SetBudgetName(v string)`

SetBudgetName sets BudgetName field to given value.


### GetScopeType

`func (o *BudgetStatusItem) GetScopeType() string`

GetScopeType returns the ScopeType field if non-nil, zero value otherwise.

### GetScopeTypeOk

`func (o *BudgetStatusItem) GetScopeTypeOk() (*string, bool)`

GetScopeTypeOk returns a tuple with the ScopeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeType

`func (o *BudgetStatusItem) SetScopeType(v string)`

SetScopeType sets ScopeType field to given value.


### GetScopeId

`func (o *BudgetStatusItem) GetScopeId() string`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *BudgetStatusItem) GetScopeIdOk() (*string, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *BudgetStatusItem) SetScopeId(v string)`

SetScopeId sets ScopeId field to given value.


### GetBudgetAmount

`func (o *BudgetStatusItem) GetBudgetAmount() string`

GetBudgetAmount returns the BudgetAmount field if non-nil, zero value otherwise.

### GetBudgetAmountOk

`func (o *BudgetStatusItem) GetBudgetAmountOk() (*string, bool)`

GetBudgetAmountOk returns a tuple with the BudgetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetAmount

`func (o *BudgetStatusItem) SetBudgetAmount(v string)`

SetBudgetAmount sets BudgetAmount field to given value.


### GetCurrency

`func (o *BudgetStatusItem) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BudgetStatusItem) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BudgetStatusItem) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCurrentSpend

`func (o *BudgetStatusItem) GetCurrentSpend() string`

GetCurrentSpend returns the CurrentSpend field if non-nil, zero value otherwise.

### GetCurrentSpendOk

`func (o *BudgetStatusItem) GetCurrentSpendOk() (*string, bool)`

GetCurrentSpendOk returns a tuple with the CurrentSpend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSpend

`func (o *BudgetStatusItem) SetCurrentSpend(v string)`

SetCurrentSpend sets CurrentSpend field to given value.


### GetUtilizationPct

`func (o *BudgetStatusItem) GetUtilizationPct() string`

GetUtilizationPct returns the UtilizationPct field if non-nil, zero value otherwise.

### GetUtilizationPctOk

`func (o *BudgetStatusItem) GetUtilizationPctOk() (*string, bool)`

GetUtilizationPctOk returns a tuple with the UtilizationPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizationPct

`func (o *BudgetStatusItem) SetUtilizationPct(v string)`

SetUtilizationPct sets UtilizationPct field to given value.


### GetIsActive

`func (o *BudgetStatusItem) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *BudgetStatusItem) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *BudgetStatusItem) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetActiveAlerts

`func (o *BudgetStatusItem) GetActiveAlerts() int32`

GetActiveAlerts returns the ActiveAlerts field if non-nil, zero value otherwise.

### GetActiveAlertsOk

`func (o *BudgetStatusItem) GetActiveAlertsOk() (*int32, bool)`

GetActiveAlertsOk returns a tuple with the ActiveAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveAlerts

`func (o *BudgetStatusItem) SetActiveAlerts(v int32)`

SetActiveAlerts sets ActiveAlerts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


