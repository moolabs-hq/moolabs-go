# AcuteValidationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Loc** | [**[]ValidationErrorLocInner**](ValidationErrorLocInner.md) |  | 
**Msg** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewAcuteValidationError

`func NewAcuteValidationError(loc []ValidationErrorLocInner, msg string, type_ string, ) *AcuteValidationError`

NewAcuteValidationError instantiates a new AcuteValidationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcuteValidationErrorWithDefaults

`func NewAcuteValidationErrorWithDefaults() *AcuteValidationError`

NewAcuteValidationErrorWithDefaults instantiates a new AcuteValidationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoc

`func (o *AcuteValidationError) GetLoc() []ValidationErrorLocInner`

GetLoc returns the Loc field if non-nil, zero value otherwise.

### GetLocOk

`func (o *AcuteValidationError) GetLocOk() (*[]ValidationErrorLocInner, bool)`

GetLocOk returns a tuple with the Loc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoc

`func (o *AcuteValidationError) SetLoc(v []ValidationErrorLocInner)`

SetLoc sets Loc field to given value.


### GetMsg

`func (o *AcuteValidationError) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *AcuteValidationError) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *AcuteValidationError) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetType

`func (o *AcuteValidationError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AcuteValidationError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AcuteValidationError) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


