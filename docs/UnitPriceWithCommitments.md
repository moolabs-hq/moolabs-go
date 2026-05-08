# UnitPriceWithCommitments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the price. | 
**Amount** | **string** | The amount of the unit price. | 
**MinimumAmount** | Pointer to **string** | The customer is committed to spend at least the amount. | [optional] 
**MaximumAmount** | Pointer to **string** | The customer is limited to spend at most the amount. | [optional] 

## Methods

### NewUnitPriceWithCommitments

`func NewUnitPriceWithCommitments(type_ string, amount string, ) *UnitPriceWithCommitments`

NewUnitPriceWithCommitments instantiates a new UnitPriceWithCommitments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnitPriceWithCommitmentsWithDefaults

`func NewUnitPriceWithCommitmentsWithDefaults() *UnitPriceWithCommitments`

NewUnitPriceWithCommitmentsWithDefaults instantiates a new UnitPriceWithCommitments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UnitPriceWithCommitments) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnitPriceWithCommitments) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnitPriceWithCommitments) SetType(v string)`

SetType sets Type field to given value.


### GetAmount

`func (o *UnitPriceWithCommitments) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UnitPriceWithCommitments) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UnitPriceWithCommitments) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetMinimumAmount

`func (o *UnitPriceWithCommitments) GetMinimumAmount() string`

GetMinimumAmount returns the MinimumAmount field if non-nil, zero value otherwise.

### GetMinimumAmountOk

`func (o *UnitPriceWithCommitments) GetMinimumAmountOk() (*string, bool)`

GetMinimumAmountOk returns a tuple with the MinimumAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumAmount

`func (o *UnitPriceWithCommitments) SetMinimumAmount(v string)`

SetMinimumAmount sets MinimumAmount field to given value.

### HasMinimumAmount

`func (o *UnitPriceWithCommitments) HasMinimumAmount() bool`

HasMinimumAmount returns a boolean if a field has been set.

### GetMaximumAmount

`func (o *UnitPriceWithCommitments) GetMaximumAmount() string`

GetMaximumAmount returns the MaximumAmount field if non-nil, zero value otherwise.

### GetMaximumAmountOk

`func (o *UnitPriceWithCommitments) GetMaximumAmountOk() (*string, bool)`

GetMaximumAmountOk returns a tuple with the MaximumAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumAmount

`func (o *UnitPriceWithCommitments) SetMaximumAmount(v string)`

SetMaximumAmount sets MaximumAmount field to given value.

### HasMaximumAmount

`func (o *UnitPriceWithCommitments) HasMaximumAmount() bool`

HasMaximumAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


