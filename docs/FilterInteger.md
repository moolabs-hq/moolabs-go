# FilterInteger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Eq** | Pointer to **int32** | The field must be equal to the provided value. | [optional] 
**Ne** | Pointer to **int32** | The field must not be equal to the provided value. | [optional] 
**Gt** | Pointer to **int32** | The field must be greater than the provided value. | [optional] 
**Gte** | Pointer to **int32** | The field must be greater than or equal to the provided value. | [optional] 
**Lt** | Pointer to **int32** | The field must be less than the provided value. | [optional] 
**Lte** | Pointer to **int32** | The field must be less than or equal to the provided value. | [optional] 
**And** | Pointer to [**[]FilterInteger**](FilterInteger.md) | Provide a list of filters to be combined with a logical AND. | [optional] 
**Or** | Pointer to [**[]FilterInteger**](FilterInteger.md) | Provide a list of filters to be combined with a logical OR. | [optional] 

## Methods

### NewFilterInteger

`func NewFilterInteger() *FilterInteger`

NewFilterInteger instantiates a new FilterInteger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterIntegerWithDefaults

`func NewFilterIntegerWithDefaults() *FilterInteger`

NewFilterIntegerWithDefaults instantiates a new FilterInteger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEq

`func (o *FilterInteger) GetEq() int32`

GetEq returns the Eq field if non-nil, zero value otherwise.

### GetEqOk

`func (o *FilterInteger) GetEqOk() (*int32, bool)`

GetEqOk returns a tuple with the Eq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEq

`func (o *FilterInteger) SetEq(v int32)`

SetEq sets Eq field to given value.

### HasEq

`func (o *FilterInteger) HasEq() bool`

HasEq returns a boolean if a field has been set.

### GetNe

`func (o *FilterInteger) GetNe() int32`

GetNe returns the Ne field if non-nil, zero value otherwise.

### GetNeOk

`func (o *FilterInteger) GetNeOk() (*int32, bool)`

GetNeOk returns a tuple with the Ne field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNe

`func (o *FilterInteger) SetNe(v int32)`

SetNe sets Ne field to given value.

### HasNe

`func (o *FilterInteger) HasNe() bool`

HasNe returns a boolean if a field has been set.

### GetGt

`func (o *FilterInteger) GetGt() int32`

GetGt returns the Gt field if non-nil, zero value otherwise.

### GetGtOk

`func (o *FilterInteger) GetGtOk() (*int32, bool)`

GetGtOk returns a tuple with the Gt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGt

`func (o *FilterInteger) SetGt(v int32)`

SetGt sets Gt field to given value.

### HasGt

`func (o *FilterInteger) HasGt() bool`

HasGt returns a boolean if a field has been set.

### GetGte

`func (o *FilterInteger) GetGte() int32`

GetGte returns the Gte field if non-nil, zero value otherwise.

### GetGteOk

`func (o *FilterInteger) GetGteOk() (*int32, bool)`

GetGteOk returns a tuple with the Gte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGte

`func (o *FilterInteger) SetGte(v int32)`

SetGte sets Gte field to given value.

### HasGte

`func (o *FilterInteger) HasGte() bool`

HasGte returns a boolean if a field has been set.

### GetLt

`func (o *FilterInteger) GetLt() int32`

GetLt returns the Lt field if non-nil, zero value otherwise.

### GetLtOk

`func (o *FilterInteger) GetLtOk() (*int32, bool)`

GetLtOk returns a tuple with the Lt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLt

`func (o *FilterInteger) SetLt(v int32)`

SetLt sets Lt field to given value.

### HasLt

`func (o *FilterInteger) HasLt() bool`

HasLt returns a boolean if a field has been set.

### GetLte

`func (o *FilterInteger) GetLte() int32`

GetLte returns the Lte field if non-nil, zero value otherwise.

### GetLteOk

`func (o *FilterInteger) GetLteOk() (*int32, bool)`

GetLteOk returns a tuple with the Lte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLte

`func (o *FilterInteger) SetLte(v int32)`

SetLte sets Lte field to given value.

### HasLte

`func (o *FilterInteger) HasLte() bool`

HasLte returns a boolean if a field has been set.

### GetAnd

`func (o *FilterInteger) GetAnd() []FilterInteger`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *FilterInteger) GetAndOk() (*[]FilterInteger, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *FilterInteger) SetAnd(v []FilterInteger)`

SetAnd sets And field to given value.

### HasAnd

`func (o *FilterInteger) HasAnd() bool`

HasAnd returns a boolean if a field has been set.

### GetOr

`func (o *FilterInteger) GetOr() []FilterInteger`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *FilterInteger) GetOrOk() (*[]FilterInteger, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *FilterInteger) SetOr(v []FilterInteger)`

SetOr sets Or field to given value.

### HasOr

`func (o *FilterInteger) HasOr() bool`

HasOr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


