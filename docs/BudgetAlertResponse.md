# BudgetAlertResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**BudgetId** | **string** |  | 
**ThresholdPct** | **int32** |  | 
**CurrentSpend** | **string** |  | 
**UtilizationPct** | **string** |  | 
**AlertAction** | **string** |  | 
**PeriodStart** | **string** |  | 
**PeriodEnd** | **string** |  | 
**FiredAt** | **string** |  | 
**AcknowledgedAt** | **string** |  | 

## Methods

### NewBudgetAlertResponse

`func NewBudgetAlertResponse(id string, tenantId string, budgetId string, thresholdPct int32, currentSpend string, utilizationPct string, alertAction string, periodStart string, periodEnd string, firedAt string, acknowledgedAt string, ) *BudgetAlertResponse`

NewBudgetAlertResponse instantiates a new BudgetAlertResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBudgetAlertResponseWithDefaults

`func NewBudgetAlertResponseWithDefaults() *BudgetAlertResponse`

NewBudgetAlertResponseWithDefaults instantiates a new BudgetAlertResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BudgetAlertResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BudgetAlertResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BudgetAlertResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *BudgetAlertResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *BudgetAlertResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *BudgetAlertResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetBudgetId

`func (o *BudgetAlertResponse) GetBudgetId() string`

GetBudgetId returns the BudgetId field if non-nil, zero value otherwise.

### GetBudgetIdOk

`func (o *BudgetAlertResponse) GetBudgetIdOk() (*string, bool)`

GetBudgetIdOk returns a tuple with the BudgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetId

`func (o *BudgetAlertResponse) SetBudgetId(v string)`

SetBudgetId sets BudgetId field to given value.


### GetThresholdPct

`func (o *BudgetAlertResponse) GetThresholdPct() int32`

GetThresholdPct returns the ThresholdPct field if non-nil, zero value otherwise.

### GetThresholdPctOk

`func (o *BudgetAlertResponse) GetThresholdPctOk() (*int32, bool)`

GetThresholdPctOk returns a tuple with the ThresholdPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdPct

`func (o *BudgetAlertResponse) SetThresholdPct(v int32)`

SetThresholdPct sets ThresholdPct field to given value.


### GetCurrentSpend

`func (o *BudgetAlertResponse) GetCurrentSpend() string`

GetCurrentSpend returns the CurrentSpend field if non-nil, zero value otherwise.

### GetCurrentSpendOk

`func (o *BudgetAlertResponse) GetCurrentSpendOk() (*string, bool)`

GetCurrentSpendOk returns a tuple with the CurrentSpend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSpend

`func (o *BudgetAlertResponse) SetCurrentSpend(v string)`

SetCurrentSpend sets CurrentSpend field to given value.


### GetUtilizationPct

`func (o *BudgetAlertResponse) GetUtilizationPct() string`

GetUtilizationPct returns the UtilizationPct field if non-nil, zero value otherwise.

### GetUtilizationPctOk

`func (o *BudgetAlertResponse) GetUtilizationPctOk() (*string, bool)`

GetUtilizationPctOk returns a tuple with the UtilizationPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizationPct

`func (o *BudgetAlertResponse) SetUtilizationPct(v string)`

SetUtilizationPct sets UtilizationPct field to given value.


### GetAlertAction

`func (o *BudgetAlertResponse) GetAlertAction() string`

GetAlertAction returns the AlertAction field if non-nil, zero value otherwise.

### GetAlertActionOk

`func (o *BudgetAlertResponse) GetAlertActionOk() (*string, bool)`

GetAlertActionOk returns a tuple with the AlertAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertAction

`func (o *BudgetAlertResponse) SetAlertAction(v string)`

SetAlertAction sets AlertAction field to given value.


### GetPeriodStart

`func (o *BudgetAlertResponse) GetPeriodStart() string`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *BudgetAlertResponse) GetPeriodStartOk() (*string, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *BudgetAlertResponse) SetPeriodStart(v string)`

SetPeriodStart sets PeriodStart field to given value.


### GetPeriodEnd

`func (o *BudgetAlertResponse) GetPeriodEnd() string`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *BudgetAlertResponse) GetPeriodEndOk() (*string, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *BudgetAlertResponse) SetPeriodEnd(v string)`

SetPeriodEnd sets PeriodEnd field to given value.


### GetFiredAt

`func (o *BudgetAlertResponse) GetFiredAt() string`

GetFiredAt returns the FiredAt field if non-nil, zero value otherwise.

### GetFiredAtOk

`func (o *BudgetAlertResponse) GetFiredAtOk() (*string, bool)`

GetFiredAtOk returns a tuple with the FiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiredAt

`func (o *BudgetAlertResponse) SetFiredAt(v string)`

SetFiredAt sets FiredAt field to given value.


### GetAcknowledgedAt

`func (o *BudgetAlertResponse) GetAcknowledgedAt() string`

GetAcknowledgedAt returns the AcknowledgedAt field if non-nil, zero value otherwise.

### GetAcknowledgedAtOk

`func (o *BudgetAlertResponse) GetAcknowledgedAtOk() (*string, bool)`

GetAcknowledgedAtOk returns a tuple with the AcknowledgedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgedAt

`func (o *BudgetAlertResponse) SetAcknowledgedAt(v string)`

SetAcknowledgedAt sets AcknowledgedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


