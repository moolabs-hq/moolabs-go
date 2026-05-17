# ObservationOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureKey** | **string** |  | 
**BomId** | **string** |  | 
**BomVersion** | **int32** |  | 
**TotalBillingEvents** | **int32** |  | 
**FullyObservedCount** | **int32** |  | 
**PartiallyObservedCount** | **int32** |  | 
**ObservationUnknownCount** | **int32** |  | 
**ObservationCoveragePct** | **float32** |  | 
**PerWorkflowType** | **[]map[string]interface{}** |  | 

## Methods

### NewObservationOut

`func NewObservationOut(featureKey string, bomId string, bomVersion int32, totalBillingEvents int32, fullyObservedCount int32, partiallyObservedCount int32, observationUnknownCount int32, observationCoveragePct float32, perWorkflowType []map[string]interface{}, ) *ObservationOut`

NewObservationOut instantiates a new ObservationOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObservationOutWithDefaults

`func NewObservationOutWithDefaults() *ObservationOut`

NewObservationOutWithDefaults instantiates a new ObservationOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatureKey

`func (o *ObservationOut) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *ObservationOut) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *ObservationOut) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.


### GetBomId

`func (o *ObservationOut) GetBomId() string`

GetBomId returns the BomId field if non-nil, zero value otherwise.

### GetBomIdOk

`func (o *ObservationOut) GetBomIdOk() (*string, bool)`

GetBomIdOk returns a tuple with the BomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBomId

`func (o *ObservationOut) SetBomId(v string)`

SetBomId sets BomId field to given value.


### GetBomVersion

`func (o *ObservationOut) GetBomVersion() int32`

GetBomVersion returns the BomVersion field if non-nil, zero value otherwise.

### GetBomVersionOk

`func (o *ObservationOut) GetBomVersionOk() (*int32, bool)`

GetBomVersionOk returns a tuple with the BomVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBomVersion

`func (o *ObservationOut) SetBomVersion(v int32)`

SetBomVersion sets BomVersion field to given value.


### GetTotalBillingEvents

`func (o *ObservationOut) GetTotalBillingEvents() int32`

GetTotalBillingEvents returns the TotalBillingEvents field if non-nil, zero value otherwise.

### GetTotalBillingEventsOk

`func (o *ObservationOut) GetTotalBillingEventsOk() (*int32, bool)`

GetTotalBillingEventsOk returns a tuple with the TotalBillingEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBillingEvents

`func (o *ObservationOut) SetTotalBillingEvents(v int32)`

SetTotalBillingEvents sets TotalBillingEvents field to given value.


### GetFullyObservedCount

`func (o *ObservationOut) GetFullyObservedCount() int32`

GetFullyObservedCount returns the FullyObservedCount field if non-nil, zero value otherwise.

### GetFullyObservedCountOk

`func (o *ObservationOut) GetFullyObservedCountOk() (*int32, bool)`

GetFullyObservedCountOk returns a tuple with the FullyObservedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyObservedCount

`func (o *ObservationOut) SetFullyObservedCount(v int32)`

SetFullyObservedCount sets FullyObservedCount field to given value.


### GetPartiallyObservedCount

`func (o *ObservationOut) GetPartiallyObservedCount() int32`

GetPartiallyObservedCount returns the PartiallyObservedCount field if non-nil, zero value otherwise.

### GetPartiallyObservedCountOk

`func (o *ObservationOut) GetPartiallyObservedCountOk() (*int32, bool)`

GetPartiallyObservedCountOk returns a tuple with the PartiallyObservedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartiallyObservedCount

`func (o *ObservationOut) SetPartiallyObservedCount(v int32)`

SetPartiallyObservedCount sets PartiallyObservedCount field to given value.


### GetObservationUnknownCount

`func (o *ObservationOut) GetObservationUnknownCount() int32`

GetObservationUnknownCount returns the ObservationUnknownCount field if non-nil, zero value otherwise.

### GetObservationUnknownCountOk

`func (o *ObservationOut) GetObservationUnknownCountOk() (*int32, bool)`

GetObservationUnknownCountOk returns a tuple with the ObservationUnknownCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationUnknownCount

`func (o *ObservationOut) SetObservationUnknownCount(v int32)`

SetObservationUnknownCount sets ObservationUnknownCount field to given value.


### GetObservationCoveragePct

`func (o *ObservationOut) GetObservationCoveragePct() float32`

GetObservationCoveragePct returns the ObservationCoveragePct field if non-nil, zero value otherwise.

### GetObservationCoveragePctOk

`func (o *ObservationOut) GetObservationCoveragePctOk() (*float32, bool)`

GetObservationCoveragePctOk returns a tuple with the ObservationCoveragePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationCoveragePct

`func (o *ObservationOut) SetObservationCoveragePct(v float32)`

SetObservationCoveragePct sets ObservationCoveragePct field to given value.


### GetPerWorkflowType

`func (o *ObservationOut) GetPerWorkflowType() []map[string]interface{}`

GetPerWorkflowType returns the PerWorkflowType field if non-nil, zero value otherwise.

### GetPerWorkflowTypeOk

`func (o *ObservationOut) GetPerWorkflowTypeOk() (*[]map[string]interface{}, bool)`

GetPerWorkflowTypeOk returns a tuple with the PerWorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerWorkflowType

`func (o *ObservationOut) SetPerWorkflowType(v []map[string]interface{})`

SetPerWorkflowType sets PerWorkflowType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


