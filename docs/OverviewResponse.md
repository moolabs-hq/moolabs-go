# OverviewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** |  | 
**TotalSpend** | **float32** |  | 
**Requests** | **int32** |  | 
**ReconciledPercent** | **float32** |  | 
**ObservationStatus** | **string** |  | 
**IngestHealth** | [**IngestHealth**](IngestHealth.md) |  | 

## Methods

### NewOverviewResponse

`func NewOverviewResponse(currency string, totalSpend float32, requests int32, reconciledPercent float32, observationStatus string, ingestHealth IngestHealth, ) *OverviewResponse`

NewOverviewResponse instantiates a new OverviewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverviewResponseWithDefaults

`func NewOverviewResponseWithDefaults() *OverviewResponse`

NewOverviewResponseWithDefaults instantiates a new OverviewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *OverviewResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *OverviewResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *OverviewResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetTotalSpend

`func (o *OverviewResponse) GetTotalSpend() float32`

GetTotalSpend returns the TotalSpend field if non-nil, zero value otherwise.

### GetTotalSpendOk

`func (o *OverviewResponse) GetTotalSpendOk() (*float32, bool)`

GetTotalSpendOk returns a tuple with the TotalSpend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSpend

`func (o *OverviewResponse) SetTotalSpend(v float32)`

SetTotalSpend sets TotalSpend field to given value.


### GetRequests

`func (o *OverviewResponse) GetRequests() int32`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *OverviewResponse) GetRequestsOk() (*int32, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *OverviewResponse) SetRequests(v int32)`

SetRequests sets Requests field to given value.


### GetReconciledPercent

`func (o *OverviewResponse) GetReconciledPercent() float32`

GetReconciledPercent returns the ReconciledPercent field if non-nil, zero value otherwise.

### GetReconciledPercentOk

`func (o *OverviewResponse) GetReconciledPercentOk() (*float32, bool)`

GetReconciledPercentOk returns a tuple with the ReconciledPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconciledPercent

`func (o *OverviewResponse) SetReconciledPercent(v float32)`

SetReconciledPercent sets ReconciledPercent field to given value.


### GetObservationStatus

`func (o *OverviewResponse) GetObservationStatus() string`

GetObservationStatus returns the ObservationStatus field if non-nil, zero value otherwise.

### GetObservationStatusOk

`func (o *OverviewResponse) GetObservationStatusOk() (*string, bool)`

GetObservationStatusOk returns a tuple with the ObservationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationStatus

`func (o *OverviewResponse) SetObservationStatus(v string)`

SetObservationStatus sets ObservationStatus field to given value.


### GetIngestHealth

`func (o *OverviewResponse) GetIngestHealth() IngestHealth`

GetIngestHealth returns the IngestHealth field if non-nil, zero value otherwise.

### GetIngestHealthOk

`func (o *OverviewResponse) GetIngestHealthOk() (*IngestHealth, bool)`

GetIngestHealthOk returns a tuple with the IngestHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestHealth

`func (o *OverviewResponse) SetIngestHealth(v IngestHealth)`

SetIngestHealth sets IngestHealth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


