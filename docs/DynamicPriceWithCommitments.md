# DynamicPriceWithCommitments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the price. | 
**Multiplier** | Pointer to **string** | The multiplier to apply to the base price to get the dynamic price.  Examples: - 0.0: the price is zero - 0.5: the price is 50% of the base price - 1.0: the price is the same as the base price - 1.5: the price is 150% of the base price | [optional] 
**MinimumAmount** | Pointer to **string** | The customer is committed to spend at least the amount. | [optional] 
**MaximumAmount** | Pointer to **string** | The customer is limited to spend at most the amount. | [optional] 

## Methods

### NewDynamicPriceWithCommitments

`func NewDynamicPriceWithCommitments(type_ string, ) *DynamicPriceWithCommitments`

NewDynamicPriceWithCommitments instantiates a new DynamicPriceWithCommitments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicPriceWithCommitmentsWithDefaults

`func NewDynamicPriceWithCommitmentsWithDefaults() *DynamicPriceWithCommitments`

NewDynamicPriceWithCommitmentsWithDefaults instantiates a new DynamicPriceWithCommitments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DynamicPriceWithCommitments) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DynamicPriceWithCommitments) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DynamicPriceWithCommitments) SetType(v string)`

SetType sets Type field to given value.


### GetMultiplier

`func (o *DynamicPriceWithCommitments) GetMultiplier() string`

GetMultiplier returns the Multiplier field if non-nil, zero value otherwise.

### GetMultiplierOk

`func (o *DynamicPriceWithCommitments) GetMultiplierOk() (*string, bool)`

GetMultiplierOk returns a tuple with the Multiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplier

`func (o *DynamicPriceWithCommitments) SetMultiplier(v string)`

SetMultiplier sets Multiplier field to given value.

### HasMultiplier

`func (o *DynamicPriceWithCommitments) HasMultiplier() bool`

HasMultiplier returns a boolean if a field has been set.

### GetMinimumAmount

`func (o *DynamicPriceWithCommitments) GetMinimumAmount() string`

GetMinimumAmount returns the MinimumAmount field if non-nil, zero value otherwise.

### GetMinimumAmountOk

`func (o *DynamicPriceWithCommitments) GetMinimumAmountOk() (*string, bool)`

GetMinimumAmountOk returns a tuple with the MinimumAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumAmount

`func (o *DynamicPriceWithCommitments) SetMinimumAmount(v string)`

SetMinimumAmount sets MinimumAmount field to given value.

### HasMinimumAmount

`func (o *DynamicPriceWithCommitments) HasMinimumAmount() bool`

HasMinimumAmount returns a boolean if a field has been set.

### GetMaximumAmount

`func (o *DynamicPriceWithCommitments) GetMaximumAmount() string`

GetMaximumAmount returns the MaximumAmount field if non-nil, zero value otherwise.

### GetMaximumAmountOk

`func (o *DynamicPriceWithCommitments) GetMaximumAmountOk() (*string, bool)`

GetMaximumAmountOk returns a tuple with the MaximumAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumAmount

`func (o *DynamicPriceWithCommitments) SetMaximumAmount(v string)`

SetMaximumAmount sets MaximumAmount field to given value.

### HasMaximumAmount

`func (o *DynamicPriceWithCommitments) HasMaximumAmount() bool`

HasMaximumAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


