# MeterValidationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | The path to the field. | [readonly] 
**Code** | **string** | The machine readable description of the error. | [readonly] 
**Message** | **string** | The human readable description of the error. | [readonly] 
**Attributes** | Pointer to **map[string]interface{}** | Additional attributes. | [optional] [readonly] 

## Methods

### NewMeterValidationError

`func NewMeterValidationError(field string, code string, message string, ) *MeterValidationError`

NewMeterValidationError instantiates a new MeterValidationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeterValidationErrorWithDefaults

`func NewMeterValidationErrorWithDefaults() *MeterValidationError`

NewMeterValidationErrorWithDefaults instantiates a new MeterValidationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *MeterValidationError) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *MeterValidationError) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *MeterValidationError) SetField(v string)`

SetField sets Field field to given value.


### GetCode

`func (o *MeterValidationError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *MeterValidationError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *MeterValidationError) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *MeterValidationError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MeterValidationError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MeterValidationError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetAttributes

`func (o *MeterValidationError) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MeterValidationError) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MeterValidationError) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *MeterValidationError) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


