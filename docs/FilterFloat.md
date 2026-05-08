# FilterFloat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Eq** | Pointer to **float64** | The field must be equal to the provided value. | [optional] 
**Ne** | Pointer to **float64** | The field must not be equal to the provided value. | [optional] 
**Gt** | Pointer to **float64** | The field must be greater than the provided value. | [optional] 
**Gte** | Pointer to **float64** | The field must be greater than or equal to the provided value. | [optional] 
**Lt** | Pointer to **float64** | The field must be less than the provided value. | [optional] 
**Lte** | Pointer to **float64** | The field must be less than or equal to the provided value. | [optional] 
**And** | Pointer to [**[]FilterFloat**](FilterFloat.md) | Provide a list of filters to be combined with a logical AND. | [optional] 
**Or** | Pointer to [**[]FilterFloat**](FilterFloat.md) | Provide a list of filters to be combined with a logical OR. | [optional] 

## Methods

### NewFilterFloat

`func NewFilterFloat() *FilterFloat`

NewFilterFloat instantiates a new FilterFloat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterFloatWithDefaults

`func NewFilterFloatWithDefaults() *FilterFloat`

NewFilterFloatWithDefaults instantiates a new FilterFloat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEq

`func (o *FilterFloat) GetEq() float64`

GetEq returns the Eq field if non-nil, zero value otherwise.

### GetEqOk

`func (o *FilterFloat) GetEqOk() (*float64, bool)`

GetEqOk returns a tuple with the Eq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEq

`func (o *FilterFloat) SetEq(v float64)`

SetEq sets Eq field to given value.

### HasEq

`func (o *FilterFloat) HasEq() bool`

HasEq returns a boolean if a field has been set.

### GetNe

`func (o *FilterFloat) GetNe() float64`

GetNe returns the Ne field if non-nil, zero value otherwise.

### GetNeOk

`func (o *FilterFloat) GetNeOk() (*float64, bool)`

GetNeOk returns a tuple with the Ne field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNe

`func (o *FilterFloat) SetNe(v float64)`

SetNe sets Ne field to given value.

### HasNe

`func (o *FilterFloat) HasNe() bool`

HasNe returns a boolean if a field has been set.

### GetGt

`func (o *FilterFloat) GetGt() float64`

GetGt returns the Gt field if non-nil, zero value otherwise.

### GetGtOk

`func (o *FilterFloat) GetGtOk() (*float64, bool)`

GetGtOk returns a tuple with the Gt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGt

`func (o *FilterFloat) SetGt(v float64)`

SetGt sets Gt field to given value.

### HasGt

`func (o *FilterFloat) HasGt() bool`

HasGt returns a boolean if a field has been set.

### GetGte

`func (o *FilterFloat) GetGte() float64`

GetGte returns the Gte field if non-nil, zero value otherwise.

### GetGteOk

`func (o *FilterFloat) GetGteOk() (*float64, bool)`

GetGteOk returns a tuple with the Gte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGte

`func (o *FilterFloat) SetGte(v float64)`

SetGte sets Gte field to given value.

### HasGte

`func (o *FilterFloat) HasGte() bool`

HasGte returns a boolean if a field has been set.

### GetLt

`func (o *FilterFloat) GetLt() float64`

GetLt returns the Lt field if non-nil, zero value otherwise.

### GetLtOk

`func (o *FilterFloat) GetLtOk() (*float64, bool)`

GetLtOk returns a tuple with the Lt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLt

`func (o *FilterFloat) SetLt(v float64)`

SetLt sets Lt field to given value.

### HasLt

`func (o *FilterFloat) HasLt() bool`

HasLt returns a boolean if a field has been set.

### GetLte

`func (o *FilterFloat) GetLte() float64`

GetLte returns the Lte field if non-nil, zero value otherwise.

### GetLteOk

`func (o *FilterFloat) GetLteOk() (*float64, bool)`

GetLteOk returns a tuple with the Lte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLte

`func (o *FilterFloat) SetLte(v float64)`

SetLte sets Lte field to given value.

### HasLte

`func (o *FilterFloat) HasLte() bool`

HasLte returns a boolean if a field has been set.

### GetAnd

`func (o *FilterFloat) GetAnd() []FilterFloat`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *FilterFloat) GetAndOk() (*[]FilterFloat, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *FilterFloat) SetAnd(v []FilterFloat)`

SetAnd sets And field to given value.

### HasAnd

`func (o *FilterFloat) HasAnd() bool`

HasAnd returns a boolean if a field has been set.

### GetOr

`func (o *FilterFloat) GetOr() []FilterFloat`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *FilterFloat) GetOrOk() (*[]FilterFloat, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *FilterFloat) SetOr(v []FilterFloat)`

SetOr sets Or field to given value.

### HasOr

`func (o *FilterFloat) HasOr() bool`

HasOr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


