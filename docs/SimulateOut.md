# SimulateOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BomId** | **string** |  | 
**BomVersion** | **int32** |  | 
**FeatureKey** | **string** |  | 
**BillingEventCount** | **int32** |  | 
**Components** | **[]map[string]interface{}** |  | 
**TotalProjectedCost** | **float32** |  | 
**TotalStandardCost** | **float32** |  | 
**VsStandard** | **map[string]interface{}** |  | 
**RateFreshness** | **string** |  | 

## Methods

### NewSimulateOut

`func NewSimulateOut(bomId string, bomVersion int32, featureKey string, billingEventCount int32, components []map[string]interface{}, totalProjectedCost float32, totalStandardCost float32, vsStandard map[string]interface{}, rateFreshness string, ) *SimulateOut`

NewSimulateOut instantiates a new SimulateOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimulateOutWithDefaults

`func NewSimulateOutWithDefaults() *SimulateOut`

NewSimulateOutWithDefaults instantiates a new SimulateOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBomId

`func (o *SimulateOut) GetBomId() string`

GetBomId returns the BomId field if non-nil, zero value otherwise.

### GetBomIdOk

`func (o *SimulateOut) GetBomIdOk() (*string, bool)`

GetBomIdOk returns a tuple with the BomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBomId

`func (o *SimulateOut) SetBomId(v string)`

SetBomId sets BomId field to given value.


### GetBomVersion

`func (o *SimulateOut) GetBomVersion() int32`

GetBomVersion returns the BomVersion field if non-nil, zero value otherwise.

### GetBomVersionOk

`func (o *SimulateOut) GetBomVersionOk() (*int32, bool)`

GetBomVersionOk returns a tuple with the BomVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBomVersion

`func (o *SimulateOut) SetBomVersion(v int32)`

SetBomVersion sets BomVersion field to given value.


### GetFeatureKey

`func (o *SimulateOut) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *SimulateOut) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *SimulateOut) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.


### GetBillingEventCount

`func (o *SimulateOut) GetBillingEventCount() int32`

GetBillingEventCount returns the BillingEventCount field if non-nil, zero value otherwise.

### GetBillingEventCountOk

`func (o *SimulateOut) GetBillingEventCountOk() (*int32, bool)`

GetBillingEventCountOk returns a tuple with the BillingEventCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEventCount

`func (o *SimulateOut) SetBillingEventCount(v int32)`

SetBillingEventCount sets BillingEventCount field to given value.


### GetComponents

`func (o *SimulateOut) GetComponents() []map[string]interface{}`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *SimulateOut) GetComponentsOk() (*[]map[string]interface{}, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *SimulateOut) SetComponents(v []map[string]interface{})`

SetComponents sets Components field to given value.


### GetTotalProjectedCost

`func (o *SimulateOut) GetTotalProjectedCost() float32`

GetTotalProjectedCost returns the TotalProjectedCost field if non-nil, zero value otherwise.

### GetTotalProjectedCostOk

`func (o *SimulateOut) GetTotalProjectedCostOk() (*float32, bool)`

GetTotalProjectedCostOk returns a tuple with the TotalProjectedCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProjectedCost

`func (o *SimulateOut) SetTotalProjectedCost(v float32)`

SetTotalProjectedCost sets TotalProjectedCost field to given value.


### GetTotalStandardCost

`func (o *SimulateOut) GetTotalStandardCost() float32`

GetTotalStandardCost returns the TotalStandardCost field if non-nil, zero value otherwise.

### GetTotalStandardCostOk

`func (o *SimulateOut) GetTotalStandardCostOk() (*float32, bool)`

GetTotalStandardCostOk returns a tuple with the TotalStandardCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStandardCost

`func (o *SimulateOut) SetTotalStandardCost(v float32)`

SetTotalStandardCost sets TotalStandardCost field to given value.


### GetVsStandard

`func (o *SimulateOut) GetVsStandard() map[string]interface{}`

GetVsStandard returns the VsStandard field if non-nil, zero value otherwise.

### GetVsStandardOk

`func (o *SimulateOut) GetVsStandardOk() (*map[string]interface{}, bool)`

GetVsStandardOk returns a tuple with the VsStandard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsStandard

`func (o *SimulateOut) SetVsStandard(v map[string]interface{})`

SetVsStandard sets VsStandard field to given value.


### GetRateFreshness

`func (o *SimulateOut) GetRateFreshness() string`

GetRateFreshness returns the RateFreshness field if non-nil, zero value otherwise.

### GetRateFreshnessOk

`func (o *SimulateOut) GetRateFreshnessOk() (*string, bool)`

GetRateFreshnessOk returns a tuple with the RateFreshness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateFreshness

`func (o *SimulateOut) SetRateFreshness(v string)`

SetRateFreshness sets RateFreshness field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


