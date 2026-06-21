# SetRateCardCostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CostUnitMicros** | Pointer to **int32** | Per-unit cost in micros (&gt;&#x3D; 0). Null clears the cost. | [optional] 
**CostSource** | Pointer to **string** | Provenance: none|manual|acute (acute treated as manual for now). | [optional] [default to "manual"]

## Methods

### NewSetRateCardCostRequest

`func NewSetRateCardCostRequest() *SetRateCardCostRequest`

NewSetRateCardCostRequest instantiates a new SetRateCardCostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetRateCardCostRequestWithDefaults

`func NewSetRateCardCostRequestWithDefaults() *SetRateCardCostRequest`

NewSetRateCardCostRequestWithDefaults instantiates a new SetRateCardCostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCostUnitMicros

`func (o *SetRateCardCostRequest) GetCostUnitMicros() int32`

GetCostUnitMicros returns the CostUnitMicros field if non-nil, zero value otherwise.

### GetCostUnitMicrosOk

`func (o *SetRateCardCostRequest) GetCostUnitMicrosOk() (*int32, bool)`

GetCostUnitMicrosOk returns a tuple with the CostUnitMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostUnitMicros

`func (o *SetRateCardCostRequest) SetCostUnitMicros(v int32)`

SetCostUnitMicros sets CostUnitMicros field to given value.

### HasCostUnitMicros

`func (o *SetRateCardCostRequest) HasCostUnitMicros() bool`

HasCostUnitMicros returns a boolean if a field has been set.

### GetCostSource

`func (o *SetRateCardCostRequest) GetCostSource() string`

GetCostSource returns the CostSource field if non-nil, zero value otherwise.

### GetCostSourceOk

`func (o *SetRateCardCostRequest) GetCostSourceOk() (*string, bool)`

GetCostSourceOk returns a tuple with the CostSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostSource

`func (o *SetRateCardCostRequest) SetCostSource(v string)`

SetCostSource sets CostSource field to given value.

### HasCostSource

`func (o *SetRateCardCostRequest) HasCostSource() bool`

HasCostSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


