# BomVarianceOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BomTemplateId** | **string** |  | 
**BomVersion** | **int32** |  | 
**FeatureKey** | **string** |  | 
**PeriodStart** | **string** |  | 
**PeriodEnd** | **string** |  | 
**StandardTotal** | **float32** |  | 
**ActualTotal** | **float32** |  | 
**TotalVariance** | **float32** |  | 
**PriceVariance** | **float32** |  | 
**UsageVariance** | **float32** |  | 
**VolumeVariance** | **float32** |  | 
**MixVariance** | **float32** |  | 
**AttributionTier** | **string** |  | 
**HeuristicAlert** | **bool** |  | 
**Components** | **[]map[string]interface{}** |  | 

## Methods

### NewBomVarianceOut

`func NewBomVarianceOut(bomTemplateId string, bomVersion int32, featureKey string, periodStart string, periodEnd string, standardTotal float32, actualTotal float32, totalVariance float32, priceVariance float32, usageVariance float32, volumeVariance float32, mixVariance float32, attributionTier string, heuristicAlert bool, components []map[string]interface{}, ) *BomVarianceOut`

NewBomVarianceOut instantiates a new BomVarianceOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBomVarianceOutWithDefaults

`func NewBomVarianceOutWithDefaults() *BomVarianceOut`

NewBomVarianceOutWithDefaults instantiates a new BomVarianceOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBomTemplateId

`func (o *BomVarianceOut) GetBomTemplateId() string`

GetBomTemplateId returns the BomTemplateId field if non-nil, zero value otherwise.

### GetBomTemplateIdOk

`func (o *BomVarianceOut) GetBomTemplateIdOk() (*string, bool)`

GetBomTemplateIdOk returns a tuple with the BomTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBomTemplateId

`func (o *BomVarianceOut) SetBomTemplateId(v string)`

SetBomTemplateId sets BomTemplateId field to given value.


### GetBomVersion

`func (o *BomVarianceOut) GetBomVersion() int32`

GetBomVersion returns the BomVersion field if non-nil, zero value otherwise.

### GetBomVersionOk

`func (o *BomVarianceOut) GetBomVersionOk() (*int32, bool)`

GetBomVersionOk returns a tuple with the BomVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBomVersion

`func (o *BomVarianceOut) SetBomVersion(v int32)`

SetBomVersion sets BomVersion field to given value.


### GetFeatureKey

`func (o *BomVarianceOut) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *BomVarianceOut) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *BomVarianceOut) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.


### GetPeriodStart

`func (o *BomVarianceOut) GetPeriodStart() string`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *BomVarianceOut) GetPeriodStartOk() (*string, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *BomVarianceOut) SetPeriodStart(v string)`

SetPeriodStart sets PeriodStart field to given value.


### GetPeriodEnd

`func (o *BomVarianceOut) GetPeriodEnd() string`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *BomVarianceOut) GetPeriodEndOk() (*string, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *BomVarianceOut) SetPeriodEnd(v string)`

SetPeriodEnd sets PeriodEnd field to given value.


### GetStandardTotal

`func (o *BomVarianceOut) GetStandardTotal() float32`

GetStandardTotal returns the StandardTotal field if non-nil, zero value otherwise.

### GetStandardTotalOk

`func (o *BomVarianceOut) GetStandardTotalOk() (*float32, bool)`

GetStandardTotalOk returns a tuple with the StandardTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardTotal

`func (o *BomVarianceOut) SetStandardTotal(v float32)`

SetStandardTotal sets StandardTotal field to given value.


### GetActualTotal

`func (o *BomVarianceOut) GetActualTotal() float32`

GetActualTotal returns the ActualTotal field if non-nil, zero value otherwise.

### GetActualTotalOk

`func (o *BomVarianceOut) GetActualTotalOk() (*float32, bool)`

GetActualTotalOk returns a tuple with the ActualTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualTotal

`func (o *BomVarianceOut) SetActualTotal(v float32)`

SetActualTotal sets ActualTotal field to given value.


### GetTotalVariance

`func (o *BomVarianceOut) GetTotalVariance() float32`

GetTotalVariance returns the TotalVariance field if non-nil, zero value otherwise.

### GetTotalVarianceOk

`func (o *BomVarianceOut) GetTotalVarianceOk() (*float32, bool)`

GetTotalVarianceOk returns a tuple with the TotalVariance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVariance

`func (o *BomVarianceOut) SetTotalVariance(v float32)`

SetTotalVariance sets TotalVariance field to given value.


### GetPriceVariance

`func (o *BomVarianceOut) GetPriceVariance() float32`

GetPriceVariance returns the PriceVariance field if non-nil, zero value otherwise.

### GetPriceVarianceOk

`func (o *BomVarianceOut) GetPriceVarianceOk() (*float32, bool)`

GetPriceVarianceOk returns a tuple with the PriceVariance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceVariance

`func (o *BomVarianceOut) SetPriceVariance(v float32)`

SetPriceVariance sets PriceVariance field to given value.


### GetUsageVariance

`func (o *BomVarianceOut) GetUsageVariance() float32`

GetUsageVariance returns the UsageVariance field if non-nil, zero value otherwise.

### GetUsageVarianceOk

`func (o *BomVarianceOut) GetUsageVarianceOk() (*float32, bool)`

GetUsageVarianceOk returns a tuple with the UsageVariance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageVariance

`func (o *BomVarianceOut) SetUsageVariance(v float32)`

SetUsageVariance sets UsageVariance field to given value.


### GetVolumeVariance

`func (o *BomVarianceOut) GetVolumeVariance() float32`

GetVolumeVariance returns the VolumeVariance field if non-nil, zero value otherwise.

### GetVolumeVarianceOk

`func (o *BomVarianceOut) GetVolumeVarianceOk() (*float32, bool)`

GetVolumeVarianceOk returns a tuple with the VolumeVariance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeVariance

`func (o *BomVarianceOut) SetVolumeVariance(v float32)`

SetVolumeVariance sets VolumeVariance field to given value.


### GetMixVariance

`func (o *BomVarianceOut) GetMixVariance() float32`

GetMixVariance returns the MixVariance field if non-nil, zero value otherwise.

### GetMixVarianceOk

`func (o *BomVarianceOut) GetMixVarianceOk() (*float32, bool)`

GetMixVarianceOk returns a tuple with the MixVariance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMixVariance

`func (o *BomVarianceOut) SetMixVariance(v float32)`

SetMixVariance sets MixVariance field to given value.


### GetAttributionTier

`func (o *BomVarianceOut) GetAttributionTier() string`

GetAttributionTier returns the AttributionTier field if non-nil, zero value otherwise.

### GetAttributionTierOk

`func (o *BomVarianceOut) GetAttributionTierOk() (*string, bool)`

GetAttributionTierOk returns a tuple with the AttributionTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributionTier

`func (o *BomVarianceOut) SetAttributionTier(v string)`

SetAttributionTier sets AttributionTier field to given value.


### GetHeuristicAlert

`func (o *BomVarianceOut) GetHeuristicAlert() bool`

GetHeuristicAlert returns the HeuristicAlert field if non-nil, zero value otherwise.

### GetHeuristicAlertOk

`func (o *BomVarianceOut) GetHeuristicAlertOk() (*bool, bool)`

GetHeuristicAlertOk returns a tuple with the HeuristicAlert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeuristicAlert

`func (o *BomVarianceOut) SetHeuristicAlert(v bool)`

SetHeuristicAlert sets HeuristicAlert field to given value.


### GetComponents

`func (o *BomVarianceOut) GetComponents() []map[string]interface{}`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *BomVarianceOut) GetComponentsOk() (*[]map[string]interface{}, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *BomVarianceOut) SetComponents(v []map[string]interface{})`

SetComponents sets Components field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


