# RateCardDiscountInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "percentage"]
**Value** | **float32** |  | 
**Scope** | Pointer to **string** |  | [optional] [default to "all"]
**Duration** | Pointer to **string** |  | [optional] [default to "entire"]
**Cycles** | Pointer to **int32** |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **string** |  | [optional] 
**SelectedItems** | Pointer to **[]string** |  | [optional] 
**Note** | Pointer to **string** |  | [optional] 

## Methods

### NewRateCardDiscountInput

`func NewRateCardDiscountInput(value float32, ) *RateCardDiscountInput`

NewRateCardDiscountInput instantiates a new RateCardDiscountInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateCardDiscountInputWithDefaults

`func NewRateCardDiscountInputWithDefaults() *RateCardDiscountInput`

NewRateCardDiscountInputWithDefaults instantiates a new RateCardDiscountInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RateCardDiscountInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RateCardDiscountInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RateCardDiscountInput) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RateCardDiscountInput) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *RateCardDiscountInput) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RateCardDiscountInput) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RateCardDiscountInput) SetValue(v float32)`

SetValue sets Value field to given value.


### GetScope

`func (o *RateCardDiscountInput) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *RateCardDiscountInput) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *RateCardDiscountInput) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *RateCardDiscountInput) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetDuration

`func (o *RateCardDiscountInput) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *RateCardDiscountInput) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *RateCardDiscountInput) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *RateCardDiscountInput) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetCycles

`func (o *RateCardDiscountInput) GetCycles() int32`

GetCycles returns the Cycles field if non-nil, zero value otherwise.

### GetCyclesOk

`func (o *RateCardDiscountInput) GetCyclesOk() (*int32, bool)`

GetCyclesOk returns a tuple with the Cycles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycles

`func (o *RateCardDiscountInput) SetCycles(v int32)`

SetCycles sets Cycles field to given value.

### HasCycles

`func (o *RateCardDiscountInput) HasCycles() bool`

HasCycles returns a boolean if a field has been set.

### GetStartDate

`func (o *RateCardDiscountInput) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *RateCardDiscountInput) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *RateCardDiscountInput) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *RateCardDiscountInput) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *RateCardDiscountInput) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *RateCardDiscountInput) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *RateCardDiscountInput) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *RateCardDiscountInput) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetSelectedItems

`func (o *RateCardDiscountInput) GetSelectedItems() []string`

GetSelectedItems returns the SelectedItems field if non-nil, zero value otherwise.

### GetSelectedItemsOk

`func (o *RateCardDiscountInput) GetSelectedItemsOk() (*[]string, bool)`

GetSelectedItemsOk returns a tuple with the SelectedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedItems

`func (o *RateCardDiscountInput) SetSelectedItems(v []string)`

SetSelectedItems sets SelectedItems field to given value.

### HasSelectedItems

`func (o *RateCardDiscountInput) HasSelectedItems() bool`

HasSelectedItems returns a boolean if a field has been set.

### GetNote

`func (o *RateCardDiscountInput) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *RateCardDiscountInput) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *RateCardDiscountInput) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *RateCardDiscountInput) HasNote() bool`

HasNote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


