# CasePauseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | **string** |  | 
**ResumeAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCasePauseRequest

`func NewCasePauseRequest(reason string, ) *CasePauseRequest`

NewCasePauseRequest instantiates a new CasePauseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCasePauseRequestWithDefaults

`func NewCasePauseRequestWithDefaults() *CasePauseRequest`

NewCasePauseRequestWithDefaults instantiates a new CasePauseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *CasePauseRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CasePauseRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CasePauseRequest) SetReason(v string)`

SetReason sets Reason field to given value.


### GetResumeAt

`func (o *CasePauseRequest) GetResumeAt() time.Time`

GetResumeAt returns the ResumeAt field if non-nil, zero value otherwise.

### GetResumeAtOk

`func (o *CasePauseRequest) GetResumeAtOk() (*time.Time, bool)`

GetResumeAtOk returns a tuple with the ResumeAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeAt

`func (o *CasePauseRequest) SetResumeAt(v time.Time)`

SetResumeAt sets ResumeAt field to given value.

### HasResumeAt

`func (o *CasePauseRequest) HasResumeAt() bool`

HasResumeAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


