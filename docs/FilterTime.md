# FilterTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gt** | Pointer to **time.Time** | The field must be greater than the provided value. | [optional] 
**Gte** | Pointer to **time.Time** | The field must be greater than or equal to the provided value. | [optional] 
**Lt** | Pointer to **time.Time** | The field must be less than the provided value. | [optional] 
**Lte** | Pointer to **time.Time** | The field must be less than or equal to the provided value. | [optional] 
**And** | Pointer to [**[]FilterTime**](FilterTime.md) | Provide a list of filters to be combined with a logical AND. | [optional] 
**Or** | Pointer to [**[]FilterTime**](FilterTime.md) | Provide a list of filters to be combined with a logical OR. | [optional] 

## Methods

### NewFilterTime

`func NewFilterTime() *FilterTime`

NewFilterTime instantiates a new FilterTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterTimeWithDefaults

`func NewFilterTimeWithDefaults() *FilterTime`

NewFilterTimeWithDefaults instantiates a new FilterTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGt

`func (o *FilterTime) GetGt() time.Time`

GetGt returns the Gt field if non-nil, zero value otherwise.

### GetGtOk

`func (o *FilterTime) GetGtOk() (*time.Time, bool)`

GetGtOk returns a tuple with the Gt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGt

`func (o *FilterTime) SetGt(v time.Time)`

SetGt sets Gt field to given value.

### HasGt

`func (o *FilterTime) HasGt() bool`

HasGt returns a boolean if a field has been set.

### GetGte

`func (o *FilterTime) GetGte() time.Time`

GetGte returns the Gte field if non-nil, zero value otherwise.

### GetGteOk

`func (o *FilterTime) GetGteOk() (*time.Time, bool)`

GetGteOk returns a tuple with the Gte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGte

`func (o *FilterTime) SetGte(v time.Time)`

SetGte sets Gte field to given value.

### HasGte

`func (o *FilterTime) HasGte() bool`

HasGte returns a boolean if a field has been set.

### GetLt

`func (o *FilterTime) GetLt() time.Time`

GetLt returns the Lt field if non-nil, zero value otherwise.

### GetLtOk

`func (o *FilterTime) GetLtOk() (*time.Time, bool)`

GetLtOk returns a tuple with the Lt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLt

`func (o *FilterTime) SetLt(v time.Time)`

SetLt sets Lt field to given value.

### HasLt

`func (o *FilterTime) HasLt() bool`

HasLt returns a boolean if a field has been set.

### GetLte

`func (o *FilterTime) GetLte() time.Time`

GetLte returns the Lte field if non-nil, zero value otherwise.

### GetLteOk

`func (o *FilterTime) GetLteOk() (*time.Time, bool)`

GetLteOk returns a tuple with the Lte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLte

`func (o *FilterTime) SetLte(v time.Time)`

SetLte sets Lte field to given value.

### HasLte

`func (o *FilterTime) HasLte() bool`

HasLte returns a boolean if a field has been set.

### GetAnd

`func (o *FilterTime) GetAnd() []FilterTime`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *FilterTime) GetAndOk() (*[]FilterTime, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *FilterTime) SetAnd(v []FilterTime)`

SetAnd sets And field to given value.

### HasAnd

`func (o *FilterTime) HasAnd() bool`

HasAnd returns a boolean if a field has been set.

### GetOr

`func (o *FilterTime) GetOr() []FilterTime`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *FilterTime) GetOrOk() (*[]FilterTime, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *FilterTime) SetOr(v []FilterTime)`

SetOr sets Or field to given value.

### HasOr

`func (o *FilterTime) HasOr() bool`

HasOr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


