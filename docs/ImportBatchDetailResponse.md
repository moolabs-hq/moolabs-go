# ImportBatchDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImportBatchId** | **string** |  | 
**CloudProvider** | **string** |  | 
**BillingPeriodStart** | **string** |  | 
**BillingPeriodEnd** | **string** |  | 
**BatchStatus** | **string** |  | 
**Rows** | [**[]ImportBatchDetailRow**](ImportBatchDetailRow.md) |  | 

## Methods

### NewImportBatchDetailResponse

`func NewImportBatchDetailResponse(importBatchId string, cloudProvider string, billingPeriodStart string, billingPeriodEnd string, batchStatus string, rows []ImportBatchDetailRow, ) *ImportBatchDetailResponse`

NewImportBatchDetailResponse instantiates a new ImportBatchDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportBatchDetailResponseWithDefaults

`func NewImportBatchDetailResponseWithDefaults() *ImportBatchDetailResponse`

NewImportBatchDetailResponseWithDefaults instantiates a new ImportBatchDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImportBatchId

`func (o *ImportBatchDetailResponse) GetImportBatchId() string`

GetImportBatchId returns the ImportBatchId field if non-nil, zero value otherwise.

### GetImportBatchIdOk

`func (o *ImportBatchDetailResponse) GetImportBatchIdOk() (*string, bool)`

GetImportBatchIdOk returns a tuple with the ImportBatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportBatchId

`func (o *ImportBatchDetailResponse) SetImportBatchId(v string)`

SetImportBatchId sets ImportBatchId field to given value.


### GetCloudProvider

`func (o *ImportBatchDetailResponse) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ImportBatchDetailResponse) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ImportBatchDetailResponse) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetBillingPeriodStart

`func (o *ImportBatchDetailResponse) GetBillingPeriodStart() string`

GetBillingPeriodStart returns the BillingPeriodStart field if non-nil, zero value otherwise.

### GetBillingPeriodStartOk

`func (o *ImportBatchDetailResponse) GetBillingPeriodStartOk() (*string, bool)`

GetBillingPeriodStartOk returns a tuple with the BillingPeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriodStart

`func (o *ImportBatchDetailResponse) SetBillingPeriodStart(v string)`

SetBillingPeriodStart sets BillingPeriodStart field to given value.


### GetBillingPeriodEnd

`func (o *ImportBatchDetailResponse) GetBillingPeriodEnd() string`

GetBillingPeriodEnd returns the BillingPeriodEnd field if non-nil, zero value otherwise.

### GetBillingPeriodEndOk

`func (o *ImportBatchDetailResponse) GetBillingPeriodEndOk() (*string, bool)`

GetBillingPeriodEndOk returns a tuple with the BillingPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriodEnd

`func (o *ImportBatchDetailResponse) SetBillingPeriodEnd(v string)`

SetBillingPeriodEnd sets BillingPeriodEnd field to given value.


### GetBatchStatus

`func (o *ImportBatchDetailResponse) GetBatchStatus() string`

GetBatchStatus returns the BatchStatus field if non-nil, zero value otherwise.

### GetBatchStatusOk

`func (o *ImportBatchDetailResponse) GetBatchStatusOk() (*string, bool)`

GetBatchStatusOk returns a tuple with the BatchStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchStatus

`func (o *ImportBatchDetailResponse) SetBatchStatus(v string)`

SetBatchStatus sets BatchStatus field to given value.


### GetRows

`func (o *ImportBatchDetailResponse) GetRows() []ImportBatchDetailRow`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *ImportBatchDetailResponse) GetRowsOk() (*[]ImportBatchDetailRow, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *ImportBatchDetailResponse) SetRows(v []ImportBatchDetailRow)`

SetRows sets Rows field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


