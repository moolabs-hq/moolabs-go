# TieredPriceWithCommitments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the price.  One of: flat, unit, or tiered. | 
**Mode** | [**TieredPriceMode**](TieredPriceMode.md) | Defines if the tiering mode is volume-based or graduated: - In &#x60;volume&#x60;-based tiering, the maximum quantity within a period determines the per unit price. - In &#x60;graduated&#x60; tiering, pricing can change as the quantity grows. | 
**Tiers** | [**[]PriceTier**](PriceTier.md) | The tiers of the tiered price. At least one price component is required in each tier. | 
**MinimumAmount** | Pointer to **string** | The customer is committed to spend at least the amount. | [optional] 
**MaximumAmount** | Pointer to **string** | The customer is limited to spend at most the amount. | [optional] 

## Methods

### NewTieredPriceWithCommitments

`func NewTieredPriceWithCommitments(type_ string, mode TieredPriceMode, tiers []PriceTier, ) *TieredPriceWithCommitments`

NewTieredPriceWithCommitments instantiates a new TieredPriceWithCommitments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTieredPriceWithCommitmentsWithDefaults

`func NewTieredPriceWithCommitmentsWithDefaults() *TieredPriceWithCommitments`

NewTieredPriceWithCommitmentsWithDefaults instantiates a new TieredPriceWithCommitments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TieredPriceWithCommitments) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TieredPriceWithCommitments) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TieredPriceWithCommitments) SetType(v string)`

SetType sets Type field to given value.


### GetMode

`func (o *TieredPriceWithCommitments) GetMode() TieredPriceMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *TieredPriceWithCommitments) GetModeOk() (*TieredPriceMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *TieredPriceWithCommitments) SetMode(v TieredPriceMode)`

SetMode sets Mode field to given value.


### GetTiers

`func (o *TieredPriceWithCommitments) GetTiers() []PriceTier`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *TieredPriceWithCommitments) GetTiersOk() (*[]PriceTier, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *TieredPriceWithCommitments) SetTiers(v []PriceTier)`

SetTiers sets Tiers field to given value.


### GetMinimumAmount

`func (o *TieredPriceWithCommitments) GetMinimumAmount() string`

GetMinimumAmount returns the MinimumAmount field if non-nil, zero value otherwise.

### GetMinimumAmountOk

`func (o *TieredPriceWithCommitments) GetMinimumAmountOk() (*string, bool)`

GetMinimumAmountOk returns a tuple with the MinimumAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumAmount

`func (o *TieredPriceWithCommitments) SetMinimumAmount(v string)`

SetMinimumAmount sets MinimumAmount field to given value.

### HasMinimumAmount

`func (o *TieredPriceWithCommitments) HasMinimumAmount() bool`

HasMinimumAmount returns a boolean if a field has been set.

### GetMaximumAmount

`func (o *TieredPriceWithCommitments) GetMaximumAmount() string`

GetMaximumAmount returns the MaximumAmount field if non-nil, zero value otherwise.

### GetMaximumAmountOk

`func (o *TieredPriceWithCommitments) GetMaximumAmountOk() (*string, bool)`

GetMaximumAmountOk returns a tuple with the MaximumAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumAmount

`func (o *TieredPriceWithCommitments) SetMaximumAmount(v string)`

SetMaximumAmount sets MaximumAmount field to given value.

### HasMaximumAmount

`func (o *TieredPriceWithCommitments) HasMaximumAmount() bool`

HasMaximumAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


