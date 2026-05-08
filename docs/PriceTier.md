# PriceTier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpToAmount** | Pointer to **string** | Up to and including to this quantity will be contained in the tier. If null, the tier is open-ended. | [optional] 
**FlatPrice** | [**FlatPrice**](FlatPrice.md) | The flat price component of the tier. | 
**UnitPrice** | [**UnitPrice**](UnitPrice.md) | The unit price component of the tier. | 

## Methods

### NewPriceTier

`func NewPriceTier(flatPrice FlatPrice, unitPrice UnitPrice, ) *PriceTier`

NewPriceTier instantiates a new PriceTier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceTierWithDefaults

`func NewPriceTierWithDefaults() *PriceTier`

NewPriceTierWithDefaults instantiates a new PriceTier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpToAmount

`func (o *PriceTier) GetUpToAmount() string`

GetUpToAmount returns the UpToAmount field if non-nil, zero value otherwise.

### GetUpToAmountOk

`func (o *PriceTier) GetUpToAmountOk() (*string, bool)`

GetUpToAmountOk returns a tuple with the UpToAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToAmount

`func (o *PriceTier) SetUpToAmount(v string)`

SetUpToAmount sets UpToAmount field to given value.

### HasUpToAmount

`func (o *PriceTier) HasUpToAmount() bool`

HasUpToAmount returns a boolean if a field has been set.

### GetFlatPrice

`func (o *PriceTier) GetFlatPrice() FlatPrice`

GetFlatPrice returns the FlatPrice field if non-nil, zero value otherwise.

### GetFlatPriceOk

`func (o *PriceTier) GetFlatPriceOk() (*FlatPrice, bool)`

GetFlatPriceOk returns a tuple with the FlatPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlatPrice

`func (o *PriceTier) SetFlatPrice(v FlatPrice)`

SetFlatPrice sets FlatPrice field to given value.


### GetUnitPrice

`func (o *PriceTier) GetUnitPrice() UnitPrice`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *PriceTier) GetUnitPriceOk() (*UnitPrice, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *PriceTier) SetUnitPrice(v UnitPrice)`

SetUnitPrice sets UnitPrice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


