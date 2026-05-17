# VarianceDecompositionOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureKey** | **string** |  | 
**PeriodStart** | **string** |  | 
**PeriodEnd** | **string** |  | 
**StandardCost** | **float32** |  | 
**ActualCost** | **float32** |  | 
**TotalVariance** | **float32** |  | 
**PriceVariance** | **float32** |  | 
**UsageVariance** | **float32** |  | 
**VolumeVariance** | **float32** |  | 
**MixVariance** | **float32** |  | 
**AttributionTier** | **string** |  | 
**HeuristicAlert** | **bool** |  | 
**BomVersion** | **int32** |  | 
**BomTemplateId** | **string** |  | 
**ComputedAt** | **string** |  | 
**AlertNote** | Pointer to **string** |  | [optional] 

## Methods

### NewVarianceDecompositionOut

`func NewVarianceDecompositionOut(featureKey string, periodStart string, periodEnd string, standardCost float32, actualCost float32, totalVariance float32, priceVariance float32, usageVariance float32, volumeVariance float32, mixVariance float32, attributionTier string, heuristicAlert bool, bomVersion int32, bomTemplateId string, computedAt string, ) *VarianceDecompositionOut`

NewVarianceDecompositionOut instantiates a new VarianceDecompositionOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVarianceDecompositionOutWithDefaults

`func NewVarianceDecompositionOutWithDefaults() *VarianceDecompositionOut`

NewVarianceDecompositionOutWithDefaults instantiates a new VarianceDecompositionOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatureKey

`func (o *VarianceDecompositionOut) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *VarianceDecompositionOut) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *VarianceDecompositionOut) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.


### GetPeriodStart

`func (o *VarianceDecompositionOut) GetPeriodStart() string`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *VarianceDecompositionOut) GetPeriodStartOk() (*string, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *VarianceDecompositionOut) SetPeriodStart(v string)`

SetPeriodStart sets PeriodStart field to given value.


### GetPeriodEnd

`func (o *VarianceDecompositionOut) GetPeriodEnd() string`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *VarianceDecompositionOut) GetPeriodEndOk() (*string, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *VarianceDecompositionOut) SetPeriodEnd(v string)`

SetPeriodEnd sets PeriodEnd field to given value.


### GetStandardCost

`func (o *VarianceDecompositionOut) GetStandardCost() float32`

GetStandardCost returns the StandardCost field if non-nil, zero value otherwise.

### GetStandardCostOk

`func (o *VarianceDecompositionOut) GetStandardCostOk() (*float32, bool)`

GetStandardCostOk returns a tuple with the StandardCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardCost

`func (o *VarianceDecompositionOut) SetStandardCost(v float32)`

SetStandardCost sets StandardCost field to given value.


### GetActualCost

`func (o *VarianceDecompositionOut) GetActualCost() float32`

GetActualCost returns the ActualCost field if non-nil, zero value otherwise.

### GetActualCostOk

`func (o *VarianceDecompositionOut) GetActualCostOk() (*float32, bool)`

GetActualCostOk returns a tuple with the ActualCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualCost

`func (o *VarianceDecompositionOut) SetActualCost(v float32)`

SetActualCost sets ActualCost field to given value.


### GetTotalVariance

`func (o *VarianceDecompositionOut) GetTotalVariance() float32`

GetTotalVariance returns the TotalVariance field if non-nil, zero value otherwise.

### GetTotalVarianceOk

`func (o *VarianceDecompositionOut) GetTotalVarianceOk() (*float32, bool)`

GetTotalVarianceOk returns a tuple with the TotalVariance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVariance

`func (o *VarianceDecompositionOut) SetTotalVariance(v float32)`

SetTotalVariance sets TotalVariance field to given value.


### GetPriceVariance

`func (o *VarianceDecompositionOut) GetPriceVariance() float32`

GetPriceVariance returns the PriceVariance field if non-nil, zero value otherwise.

### GetPriceVarianceOk

`func (o *VarianceDecompositionOut) GetPriceVarianceOk() (*float32, bool)`

GetPriceVarianceOk returns a tuple with the PriceVariance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceVariance

`func (o *VarianceDecompositionOut) SetPriceVariance(v float32)`

SetPriceVariance sets PriceVariance field to given value.


### GetUsageVariance

`func (o *VarianceDecompositionOut) GetUsageVariance() float32`

GetUsageVariance returns the UsageVariance field if non-nil, zero value otherwise.

### GetUsageVarianceOk

`func (o *VarianceDecompositionOut) GetUsageVarianceOk() (*float32, bool)`

GetUsageVarianceOk returns a tuple with the UsageVariance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageVariance

`func (o *VarianceDecompositionOut) SetUsageVariance(v float32)`

SetUsageVariance sets UsageVariance field to given value.


### GetVolumeVariance

`func (o *VarianceDecompositionOut) GetVolumeVariance() float32`

GetVolumeVariance returns the VolumeVariance field if non-nil, zero value otherwise.

### GetVolumeVarianceOk

`func (o *VarianceDecompositionOut) GetVolumeVarianceOk() (*float32, bool)`

GetVolumeVarianceOk returns a tuple with the VolumeVariance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeVariance

`func (o *VarianceDecompositionOut) SetVolumeVariance(v float32)`

SetVolumeVariance sets VolumeVariance field to given value.


### GetMixVariance

`func (o *VarianceDecompositionOut) GetMixVariance() float32`

GetMixVariance returns the MixVariance field if non-nil, zero value otherwise.

### GetMixVarianceOk

`func (o *VarianceDecompositionOut) GetMixVarianceOk() (*float32, bool)`

GetMixVarianceOk returns a tuple with the MixVariance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMixVariance

`func (o *VarianceDecompositionOut) SetMixVariance(v float32)`

SetMixVariance sets MixVariance field to given value.


### GetAttributionTier

`func (o *VarianceDecompositionOut) GetAttributionTier() string`

GetAttributionTier returns the AttributionTier field if non-nil, zero value otherwise.

### GetAttributionTierOk

`func (o *VarianceDecompositionOut) GetAttributionTierOk() (*string, bool)`

GetAttributionTierOk returns a tuple with the AttributionTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributionTier

`func (o *VarianceDecompositionOut) SetAttributionTier(v string)`

SetAttributionTier sets AttributionTier field to given value.


### GetHeuristicAlert

`func (o *VarianceDecompositionOut) GetHeuristicAlert() bool`

GetHeuristicAlert returns the HeuristicAlert field if non-nil, zero value otherwise.

### GetHeuristicAlertOk

`func (o *VarianceDecompositionOut) GetHeuristicAlertOk() (*bool, bool)`

GetHeuristicAlertOk returns a tuple with the HeuristicAlert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeuristicAlert

`func (o *VarianceDecompositionOut) SetHeuristicAlert(v bool)`

SetHeuristicAlert sets HeuristicAlert field to given value.


### GetBomVersion

`func (o *VarianceDecompositionOut) GetBomVersion() int32`

GetBomVersion returns the BomVersion field if non-nil, zero value otherwise.

### GetBomVersionOk

`func (o *VarianceDecompositionOut) GetBomVersionOk() (*int32, bool)`

GetBomVersionOk returns a tuple with the BomVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBomVersion

`func (o *VarianceDecompositionOut) SetBomVersion(v int32)`

SetBomVersion sets BomVersion field to given value.


### GetBomTemplateId

`func (o *VarianceDecompositionOut) GetBomTemplateId() string`

GetBomTemplateId returns the BomTemplateId field if non-nil, zero value otherwise.

### GetBomTemplateIdOk

`func (o *VarianceDecompositionOut) GetBomTemplateIdOk() (*string, bool)`

GetBomTemplateIdOk returns a tuple with the BomTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBomTemplateId

`func (o *VarianceDecompositionOut) SetBomTemplateId(v string)`

SetBomTemplateId sets BomTemplateId field to given value.


### GetComputedAt

`func (o *VarianceDecompositionOut) GetComputedAt() string`

GetComputedAt returns the ComputedAt field if non-nil, zero value otherwise.

### GetComputedAtOk

`func (o *VarianceDecompositionOut) GetComputedAtOk() (*string, bool)`

GetComputedAtOk returns a tuple with the ComputedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedAt

`func (o *VarianceDecompositionOut) SetComputedAt(v string)`

SetComputedAt sets ComputedAt field to given value.


### GetAlertNote

`func (o *VarianceDecompositionOut) GetAlertNote() string`

GetAlertNote returns the AlertNote field if non-nil, zero value otherwise.

### GetAlertNoteOk

`func (o *VarianceDecompositionOut) GetAlertNoteOk() (*string, bool)`

GetAlertNoteOk returns a tuple with the AlertNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertNote

`func (o *VarianceDecompositionOut) SetAlertNote(v string)`

SetAlertNote sets AlertNote field to given value.

### HasAlertNote

`func (o *VarianceDecompositionOut) HasAlertNote() bool`

HasAlertNote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


