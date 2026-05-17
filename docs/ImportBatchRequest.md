# ImportBatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | **string** | &#39;aws&#39;, &#39;gcp&#39;, or &#39;azure&#39; | 
**BillingPeriodStart** | **time.Time** |  | 
**BillingPeriodEnd** | **time.Time** |  | 
**Rows** | [**[]CloudCostRowInput**](CloudCostRowInput.md) |  | 
**ReportingCurrency** | Pointer to **string** |  | [optional] [default to "USD"]

## Methods

### NewImportBatchRequest

`func NewImportBatchRequest(cloudProvider string, billingPeriodStart time.Time, billingPeriodEnd time.Time, rows []CloudCostRowInput, ) *ImportBatchRequest`

NewImportBatchRequest instantiates a new ImportBatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportBatchRequestWithDefaults

`func NewImportBatchRequestWithDefaults() *ImportBatchRequest`

NewImportBatchRequestWithDefaults instantiates a new ImportBatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *ImportBatchRequest) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ImportBatchRequest) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ImportBatchRequest) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetBillingPeriodStart

`func (o *ImportBatchRequest) GetBillingPeriodStart() time.Time`

GetBillingPeriodStart returns the BillingPeriodStart field if non-nil, zero value otherwise.

### GetBillingPeriodStartOk

`func (o *ImportBatchRequest) GetBillingPeriodStartOk() (*time.Time, bool)`

GetBillingPeriodStartOk returns a tuple with the BillingPeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriodStart

`func (o *ImportBatchRequest) SetBillingPeriodStart(v time.Time)`

SetBillingPeriodStart sets BillingPeriodStart field to given value.


### GetBillingPeriodEnd

`func (o *ImportBatchRequest) GetBillingPeriodEnd() time.Time`

GetBillingPeriodEnd returns the BillingPeriodEnd field if non-nil, zero value otherwise.

### GetBillingPeriodEndOk

`func (o *ImportBatchRequest) GetBillingPeriodEndOk() (*time.Time, bool)`

GetBillingPeriodEndOk returns a tuple with the BillingPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriodEnd

`func (o *ImportBatchRequest) SetBillingPeriodEnd(v time.Time)`

SetBillingPeriodEnd sets BillingPeriodEnd field to given value.


### GetRows

`func (o *ImportBatchRequest) GetRows() []CloudCostRowInput`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *ImportBatchRequest) GetRowsOk() (*[]CloudCostRowInput, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *ImportBatchRequest) SetRows(v []CloudCostRowInput)`

SetRows sets Rows field to given value.


### GetReportingCurrency

`func (o *ImportBatchRequest) GetReportingCurrency() string`

GetReportingCurrency returns the ReportingCurrency field if non-nil, zero value otherwise.

### GetReportingCurrencyOk

`func (o *ImportBatchRequest) GetReportingCurrencyOk() (*string, bool)`

GetReportingCurrencyOk returns a tuple with the ReportingCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingCurrency

`func (o *ImportBatchRequest) SetReportingCurrency(v string)`

SetReportingCurrency sets ReportingCurrency field to given value.

### HasReportingCurrency

`func (o *ImportBatchRequest) HasReportingCurrency() bool`

HasReportingCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


