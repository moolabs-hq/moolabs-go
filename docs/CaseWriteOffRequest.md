# CaseWriteOffRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | **string** |  | 
**AmountMicros** | Pointer to **int32** | Partial write-off amount; null &#x3D; full | [optional] 

## Methods

### NewCaseWriteOffRequest

`func NewCaseWriteOffRequest(reason string, ) *CaseWriteOffRequest`

NewCaseWriteOffRequest instantiates a new CaseWriteOffRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaseWriteOffRequestWithDefaults

`func NewCaseWriteOffRequestWithDefaults() *CaseWriteOffRequest`

NewCaseWriteOffRequestWithDefaults instantiates a new CaseWriteOffRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *CaseWriteOffRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CaseWriteOffRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CaseWriteOffRequest) SetReason(v string)`

SetReason sets Reason field to given value.


### GetAmountMicros

`func (o *CaseWriteOffRequest) GetAmountMicros() int32`

GetAmountMicros returns the AmountMicros field if non-nil, zero value otherwise.

### GetAmountMicrosOk

`func (o *CaseWriteOffRequest) GetAmountMicrosOk() (*int32, bool)`

GetAmountMicrosOk returns a tuple with the AmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMicros

`func (o *CaseWriteOffRequest) SetAmountMicros(v int32)`

SetAmountMicros sets AmountMicros field to given value.

### HasAmountMicros

`func (o *CaseWriteOffRequest) HasAmountMicros() bool`

HasAmountMicros returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


