# FilterString

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Eq** | Pointer to **string** | The field must be equal to the provided value. | [optional] 
**Ne** | Pointer to **string** | The field must not be equal to the provided value. | [optional] 
**In** | Pointer to **[]string** | The field must be in the provided list of values. | [optional] 
**Nin** | Pointer to **[]string** | The field must not be in the provided list of values. | [optional] 
**Like** | Pointer to **string** | The field must match the provided value. | [optional] 
**Nlike** | Pointer to **string** | The field must not match the provided value. | [optional] 
**Ilike** | Pointer to **string** | The field must match the provided value, ignoring case. | [optional] 
**Nilike** | Pointer to **string** | The field must not match the provided value, ignoring case. | [optional] 
**Gt** | Pointer to **string** | The field must be greater than the provided value. | [optional] 
**Gte** | Pointer to **string** | The field must be greater than or equal to the provided value. | [optional] 
**Lt** | Pointer to **string** | The field must be less than the provided value. | [optional] 
**Lte** | Pointer to **string** | The field must be less than or equal to the provided value. | [optional] 
**And** | Pointer to [**[]FilterString**](FilterString.md) | Provide a list of filters to be combined with a logical AND. | [optional] 
**Or** | Pointer to [**[]FilterString**](FilterString.md) | Provide a list of filters to be combined with a logical OR. | [optional] 

## Methods

### NewFilterString

`func NewFilterString() *FilterString`

NewFilterString instantiates a new FilterString object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterStringWithDefaults

`func NewFilterStringWithDefaults() *FilterString`

NewFilterStringWithDefaults instantiates a new FilterString object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEq

`func (o *FilterString) GetEq() string`

GetEq returns the Eq field if non-nil, zero value otherwise.

### GetEqOk

`func (o *FilterString) GetEqOk() (*string, bool)`

GetEqOk returns a tuple with the Eq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEq

`func (o *FilterString) SetEq(v string)`

SetEq sets Eq field to given value.

### HasEq

`func (o *FilterString) HasEq() bool`

HasEq returns a boolean if a field has been set.

### GetNe

`func (o *FilterString) GetNe() string`

GetNe returns the Ne field if non-nil, zero value otherwise.

### GetNeOk

`func (o *FilterString) GetNeOk() (*string, bool)`

GetNeOk returns a tuple with the Ne field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNe

`func (o *FilterString) SetNe(v string)`

SetNe sets Ne field to given value.

### HasNe

`func (o *FilterString) HasNe() bool`

HasNe returns a boolean if a field has been set.

### GetIn

`func (o *FilterString) GetIn() []string`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *FilterString) GetInOk() (*[]string, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *FilterString) SetIn(v []string)`

SetIn sets In field to given value.

### HasIn

`func (o *FilterString) HasIn() bool`

HasIn returns a boolean if a field has been set.

### GetNin

`func (o *FilterString) GetNin() []string`

GetNin returns the Nin field if non-nil, zero value otherwise.

### GetNinOk

`func (o *FilterString) GetNinOk() (*[]string, bool)`

GetNinOk returns a tuple with the Nin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNin

`func (o *FilterString) SetNin(v []string)`

SetNin sets Nin field to given value.

### HasNin

`func (o *FilterString) HasNin() bool`

HasNin returns a boolean if a field has been set.

### GetLike

`func (o *FilterString) GetLike() string`

GetLike returns the Like field if non-nil, zero value otherwise.

### GetLikeOk

`func (o *FilterString) GetLikeOk() (*string, bool)`

GetLikeOk returns a tuple with the Like field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLike

`func (o *FilterString) SetLike(v string)`

SetLike sets Like field to given value.

### HasLike

`func (o *FilterString) HasLike() bool`

HasLike returns a boolean if a field has been set.

### GetNlike

`func (o *FilterString) GetNlike() string`

GetNlike returns the Nlike field if non-nil, zero value otherwise.

### GetNlikeOk

`func (o *FilterString) GetNlikeOk() (*string, bool)`

GetNlikeOk returns a tuple with the Nlike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNlike

`func (o *FilterString) SetNlike(v string)`

SetNlike sets Nlike field to given value.

### HasNlike

`func (o *FilterString) HasNlike() bool`

HasNlike returns a boolean if a field has been set.

### GetIlike

`func (o *FilterString) GetIlike() string`

GetIlike returns the Ilike field if non-nil, zero value otherwise.

### GetIlikeOk

`func (o *FilterString) GetIlikeOk() (*string, bool)`

GetIlikeOk returns a tuple with the Ilike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIlike

`func (o *FilterString) SetIlike(v string)`

SetIlike sets Ilike field to given value.

### HasIlike

`func (o *FilterString) HasIlike() bool`

HasIlike returns a boolean if a field has been set.

### GetNilike

`func (o *FilterString) GetNilike() string`

GetNilike returns the Nilike field if non-nil, zero value otherwise.

### GetNilikeOk

`func (o *FilterString) GetNilikeOk() (*string, bool)`

GetNilikeOk returns a tuple with the Nilike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNilike

`func (o *FilterString) SetNilike(v string)`

SetNilike sets Nilike field to given value.

### HasNilike

`func (o *FilterString) HasNilike() bool`

HasNilike returns a boolean if a field has been set.

### GetGt

`func (o *FilterString) GetGt() string`

GetGt returns the Gt field if non-nil, zero value otherwise.

### GetGtOk

`func (o *FilterString) GetGtOk() (*string, bool)`

GetGtOk returns a tuple with the Gt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGt

`func (o *FilterString) SetGt(v string)`

SetGt sets Gt field to given value.

### HasGt

`func (o *FilterString) HasGt() bool`

HasGt returns a boolean if a field has been set.

### GetGte

`func (o *FilterString) GetGte() string`

GetGte returns the Gte field if non-nil, zero value otherwise.

### GetGteOk

`func (o *FilterString) GetGteOk() (*string, bool)`

GetGteOk returns a tuple with the Gte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGte

`func (o *FilterString) SetGte(v string)`

SetGte sets Gte field to given value.

### HasGte

`func (o *FilterString) HasGte() bool`

HasGte returns a boolean if a field has been set.

### GetLt

`func (o *FilterString) GetLt() string`

GetLt returns the Lt field if non-nil, zero value otherwise.

### GetLtOk

`func (o *FilterString) GetLtOk() (*string, bool)`

GetLtOk returns a tuple with the Lt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLt

`func (o *FilterString) SetLt(v string)`

SetLt sets Lt field to given value.

### HasLt

`func (o *FilterString) HasLt() bool`

HasLt returns a boolean if a field has been set.

### GetLte

`func (o *FilterString) GetLte() string`

GetLte returns the Lte field if non-nil, zero value otherwise.

### GetLteOk

`func (o *FilterString) GetLteOk() (*string, bool)`

GetLteOk returns a tuple with the Lte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLte

`func (o *FilterString) SetLte(v string)`

SetLte sets Lte field to given value.

### HasLte

`func (o *FilterString) HasLte() bool`

HasLte returns a boolean if a field has been set.

### GetAnd

`func (o *FilterString) GetAnd() []FilterString`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *FilterString) GetAndOk() (*[]FilterString, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *FilterString) SetAnd(v []FilterString)`

SetAnd sets And field to given value.

### HasAnd

`func (o *FilterString) HasAnd() bool`

HasAnd returns a boolean if a field has been set.

### GetOr

`func (o *FilterString) GetOr() []FilterString`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *FilterString) GetOrOk() (*[]FilterString, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *FilterString) SetOr(v []FilterString)`

SetOr sets Or field to given value.

### HasOr

`func (o *FilterString) HasOr() bool`

HasOr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


