# ErrorExtension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | The path to the field. | [readonly] 
**Code** | **string** | The machine readable description of the error. | [readonly] 
**Message** | **string** | The human readable description of the error. | [readonly] 

## Methods

### NewErrorExtension

`func NewErrorExtension(field string, code string, message string, ) *ErrorExtension`

NewErrorExtension instantiates a new ErrorExtension object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorExtensionWithDefaults

`func NewErrorExtensionWithDefaults() *ErrorExtension`

NewErrorExtensionWithDefaults instantiates a new ErrorExtension object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *ErrorExtension) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *ErrorExtension) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *ErrorExtension) SetField(v string)`

SetField sets Field field to given value.


### GetCode

`func (o *ErrorExtension) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorExtension) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorExtension) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *ErrorExtension) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorExtension) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorExtension) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


