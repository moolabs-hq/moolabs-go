# CaseEscalateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | **string** |  | 
**EscalateTo** | Pointer to **string** |  | [optional] 

## Methods

### NewCaseEscalateRequest

`func NewCaseEscalateRequest(reason string, ) *CaseEscalateRequest`

NewCaseEscalateRequest instantiates a new CaseEscalateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaseEscalateRequestWithDefaults

`func NewCaseEscalateRequestWithDefaults() *CaseEscalateRequest`

NewCaseEscalateRequestWithDefaults instantiates a new CaseEscalateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *CaseEscalateRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CaseEscalateRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CaseEscalateRequest) SetReason(v string)`

SetReason sets Reason field to given value.


### GetEscalateTo

`func (o *CaseEscalateRequest) GetEscalateTo() string`

GetEscalateTo returns the EscalateTo field if non-nil, zero value otherwise.

### GetEscalateToOk

`func (o *CaseEscalateRequest) GetEscalateToOk() (*string, bool)`

GetEscalateToOk returns a tuple with the EscalateTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEscalateTo

`func (o *CaseEscalateRequest) SetEscalateTo(v string)`

SetEscalateTo sets EscalateTo field to given value.

### HasEscalateTo

`func (o *CaseEscalateRequest) HasEscalateTo() bool`

HasEscalateTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


