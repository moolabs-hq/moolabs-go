# ModelCostItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** |  | 
**Provider** | **string** |  | 
**TotalCost** | **float32** |  | 
**RequestCount** | **int32** |  | 
**PctOfTotal** | **float32** |  | 

## Methods

### NewModelCostItem

`func NewModelCostItem(model string, provider string, totalCost float32, requestCount int32, pctOfTotal float32, ) *ModelCostItem`

NewModelCostItem instantiates a new ModelCostItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelCostItemWithDefaults

`func NewModelCostItemWithDefaults() *ModelCostItem`

NewModelCostItemWithDefaults instantiates a new ModelCostItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *ModelCostItem) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ModelCostItem) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ModelCostItem) SetModel(v string)`

SetModel sets Model field to given value.


### GetProvider

`func (o *ModelCostItem) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ModelCostItem) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ModelCostItem) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetTotalCost

`func (o *ModelCostItem) GetTotalCost() float32`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *ModelCostItem) GetTotalCostOk() (*float32, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *ModelCostItem) SetTotalCost(v float32)`

SetTotalCost sets TotalCost field to given value.


### GetRequestCount

`func (o *ModelCostItem) GetRequestCount() int32`

GetRequestCount returns the RequestCount field if non-nil, zero value otherwise.

### GetRequestCountOk

`func (o *ModelCostItem) GetRequestCountOk() (*int32, bool)`

GetRequestCountOk returns a tuple with the RequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCount

`func (o *ModelCostItem) SetRequestCount(v int32)`

SetRequestCount sets RequestCount field to given value.


### GetPctOfTotal

`func (o *ModelCostItem) GetPctOfTotal() float32`

GetPctOfTotal returns the PctOfTotal field if non-nil, zero value otherwise.

### GetPctOfTotalOk

`func (o *ModelCostItem) GetPctOfTotalOk() (*float32, bool)`

GetPctOfTotalOk returns a tuple with the PctOfTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPctOfTotal

`func (o *ModelCostItem) SetPctOfTotal(v float32)`

SetPctOfTotal sets PctOfTotal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


