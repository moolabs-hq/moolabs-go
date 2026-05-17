# BudgetResponse

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

## Methods

### NewBudgetResponse

`func NewBudgetResponse(id string, tenantId string, budgetName string, scopeType string, scopeId string, periodType string, budgetAmount string, currency string, alertThresholds []int32, notifyChannels []interface{}, isActive bool, createdBy string, createdAt string, updatedAt string, ) *BudgetResponse`

NewBudgetResponse instantiates a new BudgetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBudgetResponseWithDefaults

`func NewBudgetResponseWithDefaults() *BudgetResponse`

NewBudgetResponseWithDefaults instantiates a new BudgetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BudgetResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BudgetResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BudgetResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *BudgetResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *BudgetResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *BudgetResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetBudgetName

`func (o *BudgetResponse) GetBudgetName() string`

GetBudgetName returns the BudgetName field if non-nil, zero value otherwise.

### GetBudgetNameOk

`func (o *BudgetResponse) GetBudgetNameOk() (*string, bool)`

GetBudgetNameOk returns a tuple with the BudgetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetName

`func (o *BudgetResponse) SetBudgetName(v string)`

SetBudgetName sets BudgetName field to given value.


### GetScopeType

`func (o *BudgetResponse) GetScopeType() string`

GetScopeType returns the ScopeType field if non-nil, zero value otherwise.

### GetScopeTypeOk

`func (o *BudgetResponse) GetScopeTypeOk() (*string, bool)`

GetScopeTypeOk returns a tuple with the ScopeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeType

`func (o *BudgetResponse) SetScopeType(v string)`

SetScopeType sets ScopeType field to given value.


### GetScopeId

`func (o *BudgetResponse) GetScopeId() string`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *BudgetResponse) GetScopeIdOk() (*string, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *BudgetResponse) SetScopeId(v string)`

SetScopeId sets ScopeId field to given value.


### GetPeriodType

`func (o *BudgetResponse) GetPeriodType() string`

GetPeriodType returns the PeriodType field if non-nil, zero value otherwise.

### GetPeriodTypeOk

`func (o *BudgetResponse) GetPeriodTypeOk() (*string, bool)`

GetPeriodTypeOk returns a tuple with the PeriodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodType

`func (o *BudgetResponse) SetPeriodType(v string)`

SetPeriodType sets PeriodType field to given value.


### GetBudgetAmount

`func (o *BudgetResponse) GetBudgetAmount() string`

GetBudgetAmount returns the BudgetAmount field if non-nil, zero value otherwise.

### GetBudgetAmountOk

`func (o *BudgetResponse) GetBudgetAmountOk() (*string, bool)`

GetBudgetAmountOk returns a tuple with the BudgetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetAmount

`func (o *BudgetResponse) SetBudgetAmount(v string)`

SetBudgetAmount sets BudgetAmount field to given value.


### GetCurrency

`func (o *BudgetResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BudgetResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BudgetResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetAlertThresholds

`func (o *BudgetResponse) GetAlertThresholds() []int32`

GetAlertThresholds returns the AlertThresholds field if non-nil, zero value otherwise.

### GetAlertThresholdsOk

`func (o *BudgetResponse) GetAlertThresholdsOk() (*[]int32, bool)`

GetAlertThresholdsOk returns a tuple with the AlertThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertThresholds

`func (o *BudgetResponse) SetAlertThresholds(v []int32)`

SetAlertThresholds sets AlertThresholds field to given value.


### GetNotifyChannels

`func (o *BudgetResponse) GetNotifyChannels() []interface{}`

GetNotifyChannels returns the NotifyChannels field if non-nil, zero value otherwise.

### GetNotifyChannelsOk

`func (o *BudgetResponse) GetNotifyChannelsOk() (*[]interface{}, bool)`

GetNotifyChannelsOk returns a tuple with the NotifyChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyChannels

`func (o *BudgetResponse) SetNotifyChannels(v []interface{})`

SetNotifyChannels sets NotifyChannels field to given value.


### GetIsActive

`func (o *BudgetResponse) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *BudgetResponse) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *BudgetResponse) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetCreatedBy

`func (o *BudgetResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BudgetResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BudgetResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatedAt

`func (o *BudgetResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BudgetResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BudgetResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BudgetResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BudgetResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BudgetResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


