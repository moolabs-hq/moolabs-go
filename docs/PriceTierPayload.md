# PriceTierPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpTo** | Pointer to [**NullableUpto**](Upto.md) |  | [optional] 
**UnitAmount** | [**Unitamount**](Unitamount.md) |  | 
**FlatAmount** | Pointer to [**NullableFlatamount**](Flatamount.md) |  | [optional] 

## Methods

### NewPriceTierPayload

`func NewPriceTierPayload(unitAmount Unitamount, ) *PriceTierPayload`

NewPriceTierPayload instantiates a new PriceTierPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceTierPayloadWithDefaults

`func NewPriceTierPayloadWithDefaults() *PriceTierPayload`

NewPriceTierPayloadWithDefaults instantiates a new PriceTierPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpTo

`func (o *PriceTierPayload) GetUpTo() Upto`

GetUpTo returns the UpTo field if non-nil, zero value otherwise.

### GetUpToOk

`func (o *PriceTierPayload) GetUpToOk() (*Upto, bool)`

GetUpToOk returns a tuple with the UpTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpTo

`func (o *PriceTierPayload) SetUpTo(v Upto)`

SetUpTo sets UpTo field to given value.

### HasUpTo

`func (o *PriceTierPayload) HasUpTo() bool`

HasUpTo returns a boolean if a field has been set.

### SetUpToNil

`func (o *PriceTierPayload) SetUpToNil(b bool)`

 SetUpToNil sets the value for UpTo to be an explicit nil

### UnsetUpTo
`func (o *PriceTierPayload) UnsetUpTo()`

UnsetUpTo ensures that no value is present for UpTo, not even an explicit nil
### GetUnitAmount

`func (o *PriceTierPayload) GetUnitAmount() Unitamount`

GetUnitAmount returns the UnitAmount field if non-nil, zero value otherwise.

### GetUnitAmountOk

`func (o *PriceTierPayload) GetUnitAmountOk() (*Unitamount, bool)`

GetUnitAmountOk returns a tuple with the UnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmount

`func (o *PriceTierPayload) SetUnitAmount(v Unitamount)`

SetUnitAmount sets UnitAmount field to given value.


### GetFlatAmount

`func (o *PriceTierPayload) GetFlatAmount() Flatamount`

GetFlatAmount returns the FlatAmount field if non-nil, zero value otherwise.

### GetFlatAmountOk

`func (o *PriceTierPayload) GetFlatAmountOk() (*Flatamount, bool)`

GetFlatAmountOk returns a tuple with the FlatAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlatAmount

`func (o *PriceTierPayload) SetFlatAmount(v Flatamount)`

SetFlatAmount sets FlatAmount field to given value.

### HasFlatAmount

`func (o *PriceTierPayload) HasFlatAmount() bool`

HasFlatAmount returns a boolean if a field has been set.

### SetFlatAmountNil

`func (o *PriceTierPayload) SetFlatAmountNil(b bool)`

 SetFlatAmountNil sets the value for FlatAmount to be an explicit nil

### UnsetFlatAmount
`func (o *PriceTierPayload) UnsetFlatAmount()`

UnsetFlatAmount ensures that no value is present for FlatAmount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


