# BomComponentOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ComponentType** | **string** |  | 
**Provider** | **string** |  | 
**Model** | **string** |  | 
**MetricType** | **string** |  | 
**ExpectedQuantity** | **float32** |  | 
**ExpectedSpanCount** | **int32** |  | 
**UnitCost** | **float32** |  | 
**StandardCost** | **float32** |  | 
**Currency** | **string** |  | 
**SortOrder** | **int32** |  | 
**Notes** | **string** |  | 

## Methods

### NewBomComponentOut

`func NewBomComponentOut(id string, componentType string, provider string, model string, metricType string, expectedQuantity float32, expectedSpanCount int32, unitCost float32, standardCost float32, currency string, sortOrder int32, notes string, ) *BomComponentOut`

NewBomComponentOut instantiates a new BomComponentOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBomComponentOutWithDefaults

`func NewBomComponentOutWithDefaults() *BomComponentOut`

NewBomComponentOutWithDefaults instantiates a new BomComponentOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BomComponentOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BomComponentOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BomComponentOut) SetId(v string)`

SetId sets Id field to given value.


### GetComponentType

`func (o *BomComponentOut) GetComponentType() string`

GetComponentType returns the ComponentType field if non-nil, zero value otherwise.

### GetComponentTypeOk

`func (o *BomComponentOut) GetComponentTypeOk() (*string, bool)`

GetComponentTypeOk returns a tuple with the ComponentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentType

`func (o *BomComponentOut) SetComponentType(v string)`

SetComponentType sets ComponentType field to given value.


### GetProvider

`func (o *BomComponentOut) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *BomComponentOut) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *BomComponentOut) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetModel

`func (o *BomComponentOut) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *BomComponentOut) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *BomComponentOut) SetModel(v string)`

SetModel sets Model field to given value.


### GetMetricType

`func (o *BomComponentOut) GetMetricType() string`

GetMetricType returns the MetricType field if non-nil, zero value otherwise.

### GetMetricTypeOk

`func (o *BomComponentOut) GetMetricTypeOk() (*string, bool)`

GetMetricTypeOk returns a tuple with the MetricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricType

`func (o *BomComponentOut) SetMetricType(v string)`

SetMetricType sets MetricType field to given value.


### GetExpectedQuantity

`func (o *BomComponentOut) GetExpectedQuantity() float32`

GetExpectedQuantity returns the ExpectedQuantity field if non-nil, zero value otherwise.

### GetExpectedQuantityOk

`func (o *BomComponentOut) GetExpectedQuantityOk() (*float32, bool)`

GetExpectedQuantityOk returns a tuple with the ExpectedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedQuantity

`func (o *BomComponentOut) SetExpectedQuantity(v float32)`

SetExpectedQuantity sets ExpectedQuantity field to given value.


### GetExpectedSpanCount

`func (o *BomComponentOut) GetExpectedSpanCount() int32`

GetExpectedSpanCount returns the ExpectedSpanCount field if non-nil, zero value otherwise.

### GetExpectedSpanCountOk

`func (o *BomComponentOut) GetExpectedSpanCountOk() (*int32, bool)`

GetExpectedSpanCountOk returns a tuple with the ExpectedSpanCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedSpanCount

`func (o *BomComponentOut) SetExpectedSpanCount(v int32)`

SetExpectedSpanCount sets ExpectedSpanCount field to given value.


### GetUnitCost

`func (o *BomComponentOut) GetUnitCost() float32`

GetUnitCost returns the UnitCost field if non-nil, zero value otherwise.

### GetUnitCostOk

`func (o *BomComponentOut) GetUnitCostOk() (*float32, bool)`

GetUnitCostOk returns a tuple with the UnitCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitCost

`func (o *BomComponentOut) SetUnitCost(v float32)`

SetUnitCost sets UnitCost field to given value.


### GetStandardCost

`func (o *BomComponentOut) GetStandardCost() float32`

GetStandardCost returns the StandardCost field if non-nil, zero value otherwise.

### GetStandardCostOk

`func (o *BomComponentOut) GetStandardCostOk() (*float32, bool)`

GetStandardCostOk returns a tuple with the StandardCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardCost

`func (o *BomComponentOut) SetStandardCost(v float32)`

SetStandardCost sets StandardCost field to given value.


### GetCurrency

`func (o *BomComponentOut) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BomComponentOut) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BomComponentOut) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetSortOrder

`func (o *BomComponentOut) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *BomComponentOut) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *BomComponentOut) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.


### GetNotes

`func (o *BomComponentOut) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *BomComponentOut) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *BomComponentOut) SetNotes(v string)`

SetNotes sets Notes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


