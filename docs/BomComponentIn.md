# BomComponentIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentType** | Pointer to **string** | &#39;llm_call&#39;, &#39;embedding&#39;, etc. | [optional] [default to "llm_call"]
**Provider** | **string** |  | 
**Model** | **string** |  | 
**MetricType** | **string** |  | 
**ExpectedQuantity** | **float32** | Expected units per billing event | 
**ExpectedSpanCount** | Pointer to **int32** | Expected cost_event spans per billing event | [optional] [default to 1]
**UnitCost** | Pointer to **float32** | Override from rate_catalog if provided | [optional] 
**Currency** | Pointer to **string** |  | [optional] [default to "USD"]
**SortOrder** | Pointer to **int32** |  | [optional] [default to 0]
**Notes** | Pointer to **string** |  | [optional] 

## Methods

### NewBomComponentIn

`func NewBomComponentIn(provider string, model string, metricType string, expectedQuantity float32, ) *BomComponentIn`

NewBomComponentIn instantiates a new BomComponentIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBomComponentInWithDefaults

`func NewBomComponentInWithDefaults() *BomComponentIn`

NewBomComponentInWithDefaults instantiates a new BomComponentIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentType

`func (o *BomComponentIn) GetComponentType() string`

GetComponentType returns the ComponentType field if non-nil, zero value otherwise.

### GetComponentTypeOk

`func (o *BomComponentIn) GetComponentTypeOk() (*string, bool)`

GetComponentTypeOk returns a tuple with the ComponentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentType

`func (o *BomComponentIn) SetComponentType(v string)`

SetComponentType sets ComponentType field to given value.

### HasComponentType

`func (o *BomComponentIn) HasComponentType() bool`

HasComponentType returns a boolean if a field has been set.

### GetProvider

`func (o *BomComponentIn) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *BomComponentIn) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *BomComponentIn) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetModel

`func (o *BomComponentIn) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *BomComponentIn) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *BomComponentIn) SetModel(v string)`

SetModel sets Model field to given value.


### GetMetricType

`func (o *BomComponentIn) GetMetricType() string`

GetMetricType returns the MetricType field if non-nil, zero value otherwise.

### GetMetricTypeOk

`func (o *BomComponentIn) GetMetricTypeOk() (*string, bool)`

GetMetricTypeOk returns a tuple with the MetricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricType

`func (o *BomComponentIn) SetMetricType(v string)`

SetMetricType sets MetricType field to given value.


### GetExpectedQuantity

`func (o *BomComponentIn) GetExpectedQuantity() float32`

GetExpectedQuantity returns the ExpectedQuantity field if non-nil, zero value otherwise.

### GetExpectedQuantityOk

`func (o *BomComponentIn) GetExpectedQuantityOk() (*float32, bool)`

GetExpectedQuantityOk returns a tuple with the ExpectedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedQuantity

`func (o *BomComponentIn) SetExpectedQuantity(v float32)`

SetExpectedQuantity sets ExpectedQuantity field to given value.


### GetExpectedSpanCount

`func (o *BomComponentIn) GetExpectedSpanCount() int32`

GetExpectedSpanCount returns the ExpectedSpanCount field if non-nil, zero value otherwise.

### GetExpectedSpanCountOk

`func (o *BomComponentIn) GetExpectedSpanCountOk() (*int32, bool)`

GetExpectedSpanCountOk returns a tuple with the ExpectedSpanCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedSpanCount

`func (o *BomComponentIn) SetExpectedSpanCount(v int32)`

SetExpectedSpanCount sets ExpectedSpanCount field to given value.

### HasExpectedSpanCount

`func (o *BomComponentIn) HasExpectedSpanCount() bool`

HasExpectedSpanCount returns a boolean if a field has been set.

### GetUnitCost

`func (o *BomComponentIn) GetUnitCost() float32`

GetUnitCost returns the UnitCost field if non-nil, zero value otherwise.

### GetUnitCostOk

`func (o *BomComponentIn) GetUnitCostOk() (*float32, bool)`

GetUnitCostOk returns a tuple with the UnitCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitCost

`func (o *BomComponentIn) SetUnitCost(v float32)`

SetUnitCost sets UnitCost field to given value.

### HasUnitCost

`func (o *BomComponentIn) HasUnitCost() bool`

HasUnitCost returns a boolean if a field has been set.

### GetCurrency

`func (o *BomComponentIn) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BomComponentIn) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BomComponentIn) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BomComponentIn) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetSortOrder

`func (o *BomComponentIn) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *BomComponentIn) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *BomComponentIn) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *BomComponentIn) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetNotes

`func (o *BomComponentIn) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *BomComponentIn) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *BomComponentIn) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *BomComponentIn) HasNotes() bool`

HasNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


