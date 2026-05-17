# TieredPricingPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **string** |  | 
**Tiers** | [**[]PriceTierPayload**](PriceTierPayload.md) |  | 

## Methods

### NewTieredPricingPayload

`func NewTieredPricingPayload(mode string, tiers []PriceTierPayload, ) *TieredPricingPayload`

NewTieredPricingPayload instantiates a new TieredPricingPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTieredPricingPayloadWithDefaults

`func NewTieredPricingPayloadWithDefaults() *TieredPricingPayload`

NewTieredPricingPayloadWithDefaults instantiates a new TieredPricingPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *TieredPricingPayload) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *TieredPricingPayload) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *TieredPricingPayload) SetMode(v string)`

SetMode sets Mode field to given value.


### GetTiers

`func (o *TieredPricingPayload) GetTiers() []PriceTierPayload`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *TieredPricingPayload) GetTiersOk() (*[]PriceTierPayload, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *TieredPricingPayload) SetTiers(v []PriceTierPayload)`

SetTiers sets Tiers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


