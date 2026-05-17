# MarginSnapshotResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PeriodType** | **string** |  | 
**PeriodStart** | **string** |  | 
**PeriodEnd** | **string** |  | 
**TotalBillingEvents** | **int32** |  | 
**FullyReconciledCount** | **int32** |  | 
**PartiallyReconciledCount** | **int32** |  | 
**UnreconciledCount** | **int32** |  | 
**TotalValuedBurn** | **string** |  | 
**TotalCost** | **string** |  | 
**DirectMargin** | **string** |  | 
**DirectMarginPct** | **string** |  | 
**MarginIncludesGrade** | **string** |  | 
**ReconciliationCompleteness** | **string** |  | 
**LoadedMargin** | Pointer to **string** |  | [optional] 
**LoadedMarginPct** | Pointer to **string** |  | [optional] 
**LoadedMarginAvailable** | Pointer to **bool** |  | [optional] [default to false]
**StandardCost** | Pointer to **string** |  | [optional] 
**StandardMargin** | Pointer to **string** |  | [optional] 
**StandardMarginPct** | Pointer to **string** |  | [optional] 
**ObservationStatus** | Pointer to **string** |  | [optional] 
**ReconciliationBasis** | [**ReconciliationBasis**](ReconciliationBasis.md) |  | 

## Methods

### NewMarginSnapshotResponse

`func NewMarginSnapshotResponse(periodType string, periodStart string, periodEnd string, totalBillingEvents int32, fullyReconciledCount int32, partiallyReconciledCount int32, unreconciledCount int32, totalValuedBurn string, totalCost string, directMargin string, directMarginPct string, marginIncludesGrade string, reconciliationCompleteness string, reconciliationBasis ReconciliationBasis, ) *MarginSnapshotResponse`

NewMarginSnapshotResponse instantiates a new MarginSnapshotResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarginSnapshotResponseWithDefaults

`func NewMarginSnapshotResponseWithDefaults() *MarginSnapshotResponse`

NewMarginSnapshotResponseWithDefaults instantiates a new MarginSnapshotResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriodType

`func (o *MarginSnapshotResponse) GetPeriodType() string`

GetPeriodType returns the PeriodType field if non-nil, zero value otherwise.

### GetPeriodTypeOk

`func (o *MarginSnapshotResponse) GetPeriodTypeOk() (*string, bool)`

GetPeriodTypeOk returns a tuple with the PeriodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodType

`func (o *MarginSnapshotResponse) SetPeriodType(v string)`

SetPeriodType sets PeriodType field to given value.


### GetPeriodStart

`func (o *MarginSnapshotResponse) GetPeriodStart() string`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *MarginSnapshotResponse) GetPeriodStartOk() (*string, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *MarginSnapshotResponse) SetPeriodStart(v string)`

SetPeriodStart sets PeriodStart field to given value.


### GetPeriodEnd

`func (o *MarginSnapshotResponse) GetPeriodEnd() string`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *MarginSnapshotResponse) GetPeriodEndOk() (*string, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *MarginSnapshotResponse) SetPeriodEnd(v string)`

SetPeriodEnd sets PeriodEnd field to given value.


### GetTotalBillingEvents

`func (o *MarginSnapshotResponse) GetTotalBillingEvents() int32`

GetTotalBillingEvents returns the TotalBillingEvents field if non-nil, zero value otherwise.

### GetTotalBillingEventsOk

`func (o *MarginSnapshotResponse) GetTotalBillingEventsOk() (*int32, bool)`

GetTotalBillingEventsOk returns a tuple with the TotalBillingEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBillingEvents

`func (o *MarginSnapshotResponse) SetTotalBillingEvents(v int32)`

SetTotalBillingEvents sets TotalBillingEvents field to given value.


### GetFullyReconciledCount

`func (o *MarginSnapshotResponse) GetFullyReconciledCount() int32`

GetFullyReconciledCount returns the FullyReconciledCount field if non-nil, zero value otherwise.

### GetFullyReconciledCountOk

`func (o *MarginSnapshotResponse) GetFullyReconciledCountOk() (*int32, bool)`

GetFullyReconciledCountOk returns a tuple with the FullyReconciledCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyReconciledCount

`func (o *MarginSnapshotResponse) SetFullyReconciledCount(v int32)`

SetFullyReconciledCount sets FullyReconciledCount field to given value.


### GetPartiallyReconciledCount

`func (o *MarginSnapshotResponse) GetPartiallyReconciledCount() int32`

GetPartiallyReconciledCount returns the PartiallyReconciledCount field if non-nil, zero value otherwise.

### GetPartiallyReconciledCountOk

`func (o *MarginSnapshotResponse) GetPartiallyReconciledCountOk() (*int32, bool)`

GetPartiallyReconciledCountOk returns a tuple with the PartiallyReconciledCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartiallyReconciledCount

`func (o *MarginSnapshotResponse) SetPartiallyReconciledCount(v int32)`

SetPartiallyReconciledCount sets PartiallyReconciledCount field to given value.


### GetUnreconciledCount

`func (o *MarginSnapshotResponse) GetUnreconciledCount() int32`

GetUnreconciledCount returns the UnreconciledCount field if non-nil, zero value otherwise.

### GetUnreconciledCountOk

`func (o *MarginSnapshotResponse) GetUnreconciledCountOk() (*int32, bool)`

GetUnreconciledCountOk returns a tuple with the UnreconciledCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnreconciledCount

`func (o *MarginSnapshotResponse) SetUnreconciledCount(v int32)`

SetUnreconciledCount sets UnreconciledCount field to given value.


### GetTotalValuedBurn

`func (o *MarginSnapshotResponse) GetTotalValuedBurn() string`

GetTotalValuedBurn returns the TotalValuedBurn field if non-nil, zero value otherwise.

### GetTotalValuedBurnOk

`func (o *MarginSnapshotResponse) GetTotalValuedBurnOk() (*string, bool)`

GetTotalValuedBurnOk returns a tuple with the TotalValuedBurn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalValuedBurn

`func (o *MarginSnapshotResponse) SetTotalValuedBurn(v string)`

SetTotalValuedBurn sets TotalValuedBurn field to given value.


### GetTotalCost

`func (o *MarginSnapshotResponse) GetTotalCost() string`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *MarginSnapshotResponse) GetTotalCostOk() (*string, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *MarginSnapshotResponse) SetTotalCost(v string)`

SetTotalCost sets TotalCost field to given value.


### GetDirectMargin

`func (o *MarginSnapshotResponse) GetDirectMargin() string`

GetDirectMargin returns the DirectMargin field if non-nil, zero value otherwise.

### GetDirectMarginOk

`func (o *MarginSnapshotResponse) GetDirectMarginOk() (*string, bool)`

GetDirectMarginOk returns a tuple with the DirectMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectMargin

`func (o *MarginSnapshotResponse) SetDirectMargin(v string)`

SetDirectMargin sets DirectMargin field to given value.


### GetDirectMarginPct

`func (o *MarginSnapshotResponse) GetDirectMarginPct() string`

GetDirectMarginPct returns the DirectMarginPct field if non-nil, zero value otherwise.

### GetDirectMarginPctOk

`func (o *MarginSnapshotResponse) GetDirectMarginPctOk() (*string, bool)`

GetDirectMarginPctOk returns a tuple with the DirectMarginPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectMarginPct

`func (o *MarginSnapshotResponse) SetDirectMarginPct(v string)`

SetDirectMarginPct sets DirectMarginPct field to given value.


### GetMarginIncludesGrade

`func (o *MarginSnapshotResponse) GetMarginIncludesGrade() string`

GetMarginIncludesGrade returns the MarginIncludesGrade field if non-nil, zero value otherwise.

### GetMarginIncludesGradeOk

`func (o *MarginSnapshotResponse) GetMarginIncludesGradeOk() (*string, bool)`

GetMarginIncludesGradeOk returns a tuple with the MarginIncludesGrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginIncludesGrade

`func (o *MarginSnapshotResponse) SetMarginIncludesGrade(v string)`

SetMarginIncludesGrade sets MarginIncludesGrade field to given value.


### GetReconciliationCompleteness

`func (o *MarginSnapshotResponse) GetReconciliationCompleteness() string`

GetReconciliationCompleteness returns the ReconciliationCompleteness field if non-nil, zero value otherwise.

### GetReconciliationCompletenessOk

`func (o *MarginSnapshotResponse) GetReconciliationCompletenessOk() (*string, bool)`

GetReconciliationCompletenessOk returns a tuple with the ReconciliationCompleteness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconciliationCompleteness

`func (o *MarginSnapshotResponse) SetReconciliationCompleteness(v string)`

SetReconciliationCompleteness sets ReconciliationCompleteness field to given value.


### GetLoadedMargin

`func (o *MarginSnapshotResponse) GetLoadedMargin() string`

GetLoadedMargin returns the LoadedMargin field if non-nil, zero value otherwise.

### GetLoadedMarginOk

`func (o *MarginSnapshotResponse) GetLoadedMarginOk() (*string, bool)`

GetLoadedMarginOk returns a tuple with the LoadedMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadedMargin

`func (o *MarginSnapshotResponse) SetLoadedMargin(v string)`

SetLoadedMargin sets LoadedMargin field to given value.

### HasLoadedMargin

`func (o *MarginSnapshotResponse) HasLoadedMargin() bool`

HasLoadedMargin returns a boolean if a field has been set.

### GetLoadedMarginPct

`func (o *MarginSnapshotResponse) GetLoadedMarginPct() string`

GetLoadedMarginPct returns the LoadedMarginPct field if non-nil, zero value otherwise.

### GetLoadedMarginPctOk

`func (o *MarginSnapshotResponse) GetLoadedMarginPctOk() (*string, bool)`

GetLoadedMarginPctOk returns a tuple with the LoadedMarginPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadedMarginPct

`func (o *MarginSnapshotResponse) SetLoadedMarginPct(v string)`

SetLoadedMarginPct sets LoadedMarginPct field to given value.

### HasLoadedMarginPct

`func (o *MarginSnapshotResponse) HasLoadedMarginPct() bool`

HasLoadedMarginPct returns a boolean if a field has been set.

### GetLoadedMarginAvailable

`func (o *MarginSnapshotResponse) GetLoadedMarginAvailable() bool`

GetLoadedMarginAvailable returns the LoadedMarginAvailable field if non-nil, zero value otherwise.

### GetLoadedMarginAvailableOk

`func (o *MarginSnapshotResponse) GetLoadedMarginAvailableOk() (*bool, bool)`

GetLoadedMarginAvailableOk returns a tuple with the LoadedMarginAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadedMarginAvailable

`func (o *MarginSnapshotResponse) SetLoadedMarginAvailable(v bool)`

SetLoadedMarginAvailable sets LoadedMarginAvailable field to given value.

### HasLoadedMarginAvailable

`func (o *MarginSnapshotResponse) HasLoadedMarginAvailable() bool`

HasLoadedMarginAvailable returns a boolean if a field has been set.

### GetStandardCost

`func (o *MarginSnapshotResponse) GetStandardCost() string`

GetStandardCost returns the StandardCost field if non-nil, zero value otherwise.

### GetStandardCostOk

`func (o *MarginSnapshotResponse) GetStandardCostOk() (*string, bool)`

GetStandardCostOk returns a tuple with the StandardCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardCost

`func (o *MarginSnapshotResponse) SetStandardCost(v string)`

SetStandardCost sets StandardCost field to given value.

### HasStandardCost

`func (o *MarginSnapshotResponse) HasStandardCost() bool`

HasStandardCost returns a boolean if a field has been set.

### GetStandardMargin

`func (o *MarginSnapshotResponse) GetStandardMargin() string`

GetStandardMargin returns the StandardMargin field if non-nil, zero value otherwise.

### GetStandardMarginOk

`func (o *MarginSnapshotResponse) GetStandardMarginOk() (*string, bool)`

GetStandardMarginOk returns a tuple with the StandardMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardMargin

`func (o *MarginSnapshotResponse) SetStandardMargin(v string)`

SetStandardMargin sets StandardMargin field to given value.

### HasStandardMargin

`func (o *MarginSnapshotResponse) HasStandardMargin() bool`

HasStandardMargin returns a boolean if a field has been set.

### GetStandardMarginPct

`func (o *MarginSnapshotResponse) GetStandardMarginPct() string`

GetStandardMarginPct returns the StandardMarginPct field if non-nil, zero value otherwise.

### GetStandardMarginPctOk

`func (o *MarginSnapshotResponse) GetStandardMarginPctOk() (*string, bool)`

GetStandardMarginPctOk returns a tuple with the StandardMarginPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardMarginPct

`func (o *MarginSnapshotResponse) SetStandardMarginPct(v string)`

SetStandardMarginPct sets StandardMarginPct field to given value.

### HasStandardMarginPct

`func (o *MarginSnapshotResponse) HasStandardMarginPct() bool`

HasStandardMarginPct returns a boolean if a field has been set.

### GetObservationStatus

`func (o *MarginSnapshotResponse) GetObservationStatus() string`

GetObservationStatus returns the ObservationStatus field if non-nil, zero value otherwise.

### GetObservationStatusOk

`func (o *MarginSnapshotResponse) GetObservationStatusOk() (*string, bool)`

GetObservationStatusOk returns a tuple with the ObservationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationStatus

`func (o *MarginSnapshotResponse) SetObservationStatus(v string)`

SetObservationStatus sets ObservationStatus field to given value.

### HasObservationStatus

`func (o *MarginSnapshotResponse) HasObservationStatus() bool`

HasObservationStatus returns a boolean if a field has been set.

### GetReconciliationBasis

`func (o *MarginSnapshotResponse) GetReconciliationBasis() ReconciliationBasis`

GetReconciliationBasis returns the ReconciliationBasis field if non-nil, zero value otherwise.

### GetReconciliationBasisOk

`func (o *MarginSnapshotResponse) GetReconciliationBasisOk() (*ReconciliationBasis, bool)`

GetReconciliationBasisOk returns a tuple with the ReconciliationBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconciliationBasis

`func (o *MarginSnapshotResponse) SetReconciliationBasis(v ReconciliationBasis)`

SetReconciliationBasis sets ReconciliationBasis field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


