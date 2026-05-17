# ImportBatchSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImportBatchId** | **string** |  | 
**CloudProvider** | **string** |  | 
**BillingPeriodStart** | **string** |  | 
**BillingPeriodEnd** | **string** |  | 
**BatchStatus** | **string** |  | 
**RowCount** | **int32** |  | 
**TotalReportingCost** | **string** |  | 
**ReportingCurrency** | **string** |  | 
**ImportedAt** | **string** |  | 
**ClickhouseSyncStatus** | Pointer to **string** |  | [optional] [default to "pending"]
**ClickhouseSyncedAt** | Pointer to **string** |  | [optional] 
**ClickhouseSyncError** | Pointer to **string** |  | [optional] 
**ClickhouseSyncAttempts** | Pointer to **int32** |  | [optional] [default to 0]

## Methods

### NewImportBatchSummary

`func NewImportBatchSummary(importBatchId string, cloudProvider string, billingPeriodStart string, billingPeriodEnd string, batchStatus string, rowCount int32, totalReportingCost string, reportingCurrency string, importedAt string, ) *ImportBatchSummary`

NewImportBatchSummary instantiates a new ImportBatchSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportBatchSummaryWithDefaults

`func NewImportBatchSummaryWithDefaults() *ImportBatchSummary`

NewImportBatchSummaryWithDefaults instantiates a new ImportBatchSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImportBatchId

`func (o *ImportBatchSummary) GetImportBatchId() string`

GetImportBatchId returns the ImportBatchId field if non-nil, zero value otherwise.

### GetImportBatchIdOk

`func (o *ImportBatchSummary) GetImportBatchIdOk() (*string, bool)`

GetImportBatchIdOk returns a tuple with the ImportBatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportBatchId

`func (o *ImportBatchSummary) SetImportBatchId(v string)`

SetImportBatchId sets ImportBatchId field to given value.


### GetCloudProvider

`func (o *ImportBatchSummary) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ImportBatchSummary) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ImportBatchSummary) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetBillingPeriodStart

`func (o *ImportBatchSummary) GetBillingPeriodStart() string`

GetBillingPeriodStart returns the BillingPeriodStart field if non-nil, zero value otherwise.

### GetBillingPeriodStartOk

`func (o *ImportBatchSummary) GetBillingPeriodStartOk() (*string, bool)`

GetBillingPeriodStartOk returns a tuple with the BillingPeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriodStart

`func (o *ImportBatchSummary) SetBillingPeriodStart(v string)`

SetBillingPeriodStart sets BillingPeriodStart field to given value.


### GetBillingPeriodEnd

`func (o *ImportBatchSummary) GetBillingPeriodEnd() string`

GetBillingPeriodEnd returns the BillingPeriodEnd field if non-nil, zero value otherwise.

### GetBillingPeriodEndOk

`func (o *ImportBatchSummary) GetBillingPeriodEndOk() (*string, bool)`

GetBillingPeriodEndOk returns a tuple with the BillingPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriodEnd

`func (o *ImportBatchSummary) SetBillingPeriodEnd(v string)`

SetBillingPeriodEnd sets BillingPeriodEnd field to given value.


### GetBatchStatus

`func (o *ImportBatchSummary) GetBatchStatus() string`

GetBatchStatus returns the BatchStatus field if non-nil, zero value otherwise.

### GetBatchStatusOk

`func (o *ImportBatchSummary) GetBatchStatusOk() (*string, bool)`

GetBatchStatusOk returns a tuple with the BatchStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchStatus

`func (o *ImportBatchSummary) SetBatchStatus(v string)`

SetBatchStatus sets BatchStatus field to given value.


### GetRowCount

`func (o *ImportBatchSummary) GetRowCount() int32`

GetRowCount returns the RowCount field if non-nil, zero value otherwise.

### GetRowCountOk

`func (o *ImportBatchSummary) GetRowCountOk() (*int32, bool)`

GetRowCountOk returns a tuple with the RowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowCount

`func (o *ImportBatchSummary) SetRowCount(v int32)`

SetRowCount sets RowCount field to given value.


### GetTotalReportingCost

`func (o *ImportBatchSummary) GetTotalReportingCost() string`

GetTotalReportingCost returns the TotalReportingCost field if non-nil, zero value otherwise.

### GetTotalReportingCostOk

`func (o *ImportBatchSummary) GetTotalReportingCostOk() (*string, bool)`

GetTotalReportingCostOk returns a tuple with the TotalReportingCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalReportingCost

`func (o *ImportBatchSummary) SetTotalReportingCost(v string)`

SetTotalReportingCost sets TotalReportingCost field to given value.


### GetReportingCurrency

`func (o *ImportBatchSummary) GetReportingCurrency() string`

GetReportingCurrency returns the ReportingCurrency field if non-nil, zero value otherwise.

### GetReportingCurrencyOk

`func (o *ImportBatchSummary) GetReportingCurrencyOk() (*string, bool)`

GetReportingCurrencyOk returns a tuple with the ReportingCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingCurrency

`func (o *ImportBatchSummary) SetReportingCurrency(v string)`

SetReportingCurrency sets ReportingCurrency field to given value.


### GetImportedAt

`func (o *ImportBatchSummary) GetImportedAt() string`

GetImportedAt returns the ImportedAt field if non-nil, zero value otherwise.

### GetImportedAtOk

`func (o *ImportBatchSummary) GetImportedAtOk() (*string, bool)`

GetImportedAtOk returns a tuple with the ImportedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedAt

`func (o *ImportBatchSummary) SetImportedAt(v string)`

SetImportedAt sets ImportedAt field to given value.


### GetClickhouseSyncStatus

`func (o *ImportBatchSummary) GetClickhouseSyncStatus() string`

GetClickhouseSyncStatus returns the ClickhouseSyncStatus field if non-nil, zero value otherwise.

### GetClickhouseSyncStatusOk

`func (o *ImportBatchSummary) GetClickhouseSyncStatusOk() (*string, bool)`

GetClickhouseSyncStatusOk returns a tuple with the ClickhouseSyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickhouseSyncStatus

`func (o *ImportBatchSummary) SetClickhouseSyncStatus(v string)`

SetClickhouseSyncStatus sets ClickhouseSyncStatus field to given value.

### HasClickhouseSyncStatus

`func (o *ImportBatchSummary) HasClickhouseSyncStatus() bool`

HasClickhouseSyncStatus returns a boolean if a field has been set.

### GetClickhouseSyncedAt

`func (o *ImportBatchSummary) GetClickhouseSyncedAt() string`

GetClickhouseSyncedAt returns the ClickhouseSyncedAt field if non-nil, zero value otherwise.

### GetClickhouseSyncedAtOk

`func (o *ImportBatchSummary) GetClickhouseSyncedAtOk() (*string, bool)`

GetClickhouseSyncedAtOk returns a tuple with the ClickhouseSyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickhouseSyncedAt

`func (o *ImportBatchSummary) SetClickhouseSyncedAt(v string)`

SetClickhouseSyncedAt sets ClickhouseSyncedAt field to given value.

### HasClickhouseSyncedAt

`func (o *ImportBatchSummary) HasClickhouseSyncedAt() bool`

HasClickhouseSyncedAt returns a boolean if a field has been set.

### GetClickhouseSyncError

`func (o *ImportBatchSummary) GetClickhouseSyncError() string`

GetClickhouseSyncError returns the ClickhouseSyncError field if non-nil, zero value otherwise.

### GetClickhouseSyncErrorOk

`func (o *ImportBatchSummary) GetClickhouseSyncErrorOk() (*string, bool)`

GetClickhouseSyncErrorOk returns a tuple with the ClickhouseSyncError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickhouseSyncError

`func (o *ImportBatchSummary) SetClickhouseSyncError(v string)`

SetClickhouseSyncError sets ClickhouseSyncError field to given value.

### HasClickhouseSyncError

`func (o *ImportBatchSummary) HasClickhouseSyncError() bool`

HasClickhouseSyncError returns a boolean if a field has been set.

### GetClickhouseSyncAttempts

`func (o *ImportBatchSummary) GetClickhouseSyncAttempts() int32`

GetClickhouseSyncAttempts returns the ClickhouseSyncAttempts field if non-nil, zero value otherwise.

### GetClickhouseSyncAttemptsOk

`func (o *ImportBatchSummary) GetClickhouseSyncAttemptsOk() (*int32, bool)`

GetClickhouseSyncAttemptsOk returns a tuple with the ClickhouseSyncAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickhouseSyncAttempts

`func (o *ImportBatchSummary) SetClickhouseSyncAttempts(v int32)`

SetClickhouseSyncAttempts sets ClickhouseSyncAttempts field to given value.

### HasClickhouseSyncAttempts

`func (o *ImportBatchSummary) HasClickhouseSyncAttempts() bool`

HasClickhouseSyncAttempts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


