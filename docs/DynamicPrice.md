# DynamicPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the price. | 
**Multiplier** | Pointer to **string** | The multiplier to apply to the base price to get the dynamic price.  Examples: - 0.0: the price is zero - 0.5: the price is 50% of the base price - 1.0: the price is the same as the base price - 1.5: the price is 150% of the base price | [optional] 

## Methods

### NewDynamicPrice

`func NewDynamicPrice(type_ string, ) *DynamicPrice`

NewDynamicPrice instantiates a new DynamicPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicPriceWithDefaults

`func NewDynamicPriceWithDefaults() *DynamicPrice`

NewDynamicPriceWithDefaults instantiates a new DynamicPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DynamicPrice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DynamicPrice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DynamicPrice) SetType(v string)`

SetType sets Type field to given value.


### GetMultiplier

`func (o *DynamicPrice) GetMultiplier() string`

GetMultiplier returns the Multiplier field if non-nil, zero value otherwise.

### GetMultiplierOk

`func (o *DynamicPrice) GetMultiplierOk() (*string, bool)`

GetMultiplierOk returns a tuple with the Multiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplier

`func (o *DynamicPrice) SetMultiplier(v string)`

SetMultiplier sets Multiplier field to given value.

### HasMultiplier

`func (o *DynamicPrice) HasMultiplier() bool`

HasMultiplier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


