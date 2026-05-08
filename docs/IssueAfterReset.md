# IssueAfterReset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float64** | The initial grant amount | 
**Priority** | Pointer to **int32** | The priority of the issue after reset | [optional] [default to 1]

## Methods

### NewIssueAfterReset

`func NewIssueAfterReset(amount float64, ) *IssueAfterReset`

NewIssueAfterReset instantiates a new IssueAfterReset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueAfterResetWithDefaults

`func NewIssueAfterResetWithDefaults() *IssueAfterReset`

NewIssueAfterResetWithDefaults instantiates a new IssueAfterReset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *IssueAfterReset) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *IssueAfterReset) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *IssueAfterReset) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetPriority

`func (o *IssueAfterReset) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *IssueAfterReset) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *IssueAfterReset) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *IssueAfterReset) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


