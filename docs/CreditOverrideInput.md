# CreditOverrideInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **string** |  | 
**Amount** | Pointer to **float32** |  | [optional] [default to 0.0]
**Unit** | Pointer to **string** |  | [optional] [default to "credits"]
**Expiry** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **string** |  | [optional] 
**Note** | Pointer to **string** |  | [optional] 
**PoolOverrides** | Pointer to **map[string]float32** |  | [optional] 

## Methods

### NewCreditOverrideInput

`func NewCreditOverrideInput(mode string, ) *CreditOverrideInput`

NewCreditOverrideInput instantiates a new CreditOverrideInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditOverrideInputWithDefaults

`func NewCreditOverrideInputWithDefaults() *CreditOverrideInput`

NewCreditOverrideInputWithDefaults instantiates a new CreditOverrideInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *CreditOverrideInput) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CreditOverrideInput) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CreditOverrideInput) SetMode(v string)`

SetMode sets Mode field to given value.


### GetAmount

`func (o *CreditOverrideInput) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreditOverrideInput) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreditOverrideInput) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CreditOverrideInput) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetUnit

`func (o *CreditOverrideInput) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *CreditOverrideInput) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *CreditOverrideInput) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *CreditOverrideInput) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetExpiry

`func (o *CreditOverrideInput) GetExpiry() string`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *CreditOverrideInput) GetExpiryOk() (*string, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *CreditOverrideInput) SetExpiry(v string)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *CreditOverrideInput) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetStartDate

`func (o *CreditOverrideInput) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CreditOverrideInput) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CreditOverrideInput) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CreditOverrideInput) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *CreditOverrideInput) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CreditOverrideInput) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CreditOverrideInput) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CreditOverrideInput) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetNote

`func (o *CreditOverrideInput) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *CreditOverrideInput) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *CreditOverrideInput) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *CreditOverrideInput) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetPoolOverrides

`func (o *CreditOverrideInput) GetPoolOverrides() map[string]float32`

GetPoolOverrides returns the PoolOverrides field if non-nil, zero value otherwise.

### GetPoolOverridesOk

`func (o *CreditOverrideInput) GetPoolOverridesOk() (*map[string]float32, bool)`

GetPoolOverridesOk returns a tuple with the PoolOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolOverrides

`func (o *CreditOverrideInput) SetPoolOverrides(v map[string]float32)`

SetPoolOverrides sets PoolOverrides field to given value.

### HasPoolOverrides

`func (o *CreditOverrideInput) HasPoolOverrides() bool`

HasPoolOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


