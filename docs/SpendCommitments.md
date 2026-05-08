# SpendCommitments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinimumAmount** | Pointer to **string** | The customer is committed to spend at least the amount. | [optional] 
**MaximumAmount** | Pointer to **string** | The customer is limited to spend at most the amount. | [optional] 

## Methods

### NewSpendCommitments

`func NewSpendCommitments() *SpendCommitments`

NewSpendCommitments instantiates a new SpendCommitments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpendCommitmentsWithDefaults

`func NewSpendCommitmentsWithDefaults() *SpendCommitments`

NewSpendCommitmentsWithDefaults instantiates a new SpendCommitments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinimumAmount

`func (o *SpendCommitments) GetMinimumAmount() string`

GetMinimumAmount returns the MinimumAmount field if non-nil, zero value otherwise.

### GetMinimumAmountOk

`func (o *SpendCommitments) GetMinimumAmountOk() (*string, bool)`

GetMinimumAmountOk returns a tuple with the MinimumAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumAmount

`func (o *SpendCommitments) SetMinimumAmount(v string)`

SetMinimumAmount sets MinimumAmount field to given value.

### HasMinimumAmount

`func (o *SpendCommitments) HasMinimumAmount() bool`

HasMinimumAmount returns a boolean if a field has been set.

### GetMaximumAmount

`func (o *SpendCommitments) GetMaximumAmount() string`

GetMaximumAmount returns the MaximumAmount field if non-nil, zero value otherwise.

### GetMaximumAmountOk

`func (o *SpendCommitments) GetMaximumAmountOk() (*string, bool)`

GetMaximumAmountOk returns a tuple with the MaximumAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumAmount

`func (o *SpendCommitments) SetMaximumAmount(v string)`

SetMaximumAmount sets MaximumAmount field to given value.

### HasMaximumAmount

`func (o *SpendCommitments) HasMaximumAmount() bool`

HasMaximumAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


