# AllocateCreditsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceWalletId** | **string** |  | 
**TargetWalletId** | **string** |  | 
**AmountMicros** | **int32** | Amount in micros to transfer | 
**SubjectId** | **string** | Subject performing the allocation (for audit) | 
**Reason** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAllocateCreditsRequest

`func NewAllocateCreditsRequest(sourceWalletId string, targetWalletId string, amountMicros int32, subjectId string, ) *AllocateCreditsRequest`

NewAllocateCreditsRequest instantiates a new AllocateCreditsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllocateCreditsRequestWithDefaults

`func NewAllocateCreditsRequestWithDefaults() *AllocateCreditsRequest`

NewAllocateCreditsRequestWithDefaults instantiates a new AllocateCreditsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceWalletId

`func (o *AllocateCreditsRequest) GetSourceWalletId() string`

GetSourceWalletId returns the SourceWalletId field if non-nil, zero value otherwise.

### GetSourceWalletIdOk

`func (o *AllocateCreditsRequest) GetSourceWalletIdOk() (*string, bool)`

GetSourceWalletIdOk returns a tuple with the SourceWalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceWalletId

`func (o *AllocateCreditsRequest) SetSourceWalletId(v string)`

SetSourceWalletId sets SourceWalletId field to given value.


### GetTargetWalletId

`func (o *AllocateCreditsRequest) GetTargetWalletId() string`

GetTargetWalletId returns the TargetWalletId field if non-nil, zero value otherwise.

### GetTargetWalletIdOk

`func (o *AllocateCreditsRequest) GetTargetWalletIdOk() (*string, bool)`

GetTargetWalletIdOk returns a tuple with the TargetWalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetWalletId

`func (o *AllocateCreditsRequest) SetTargetWalletId(v string)`

SetTargetWalletId sets TargetWalletId field to given value.


### GetAmountMicros

`func (o *AllocateCreditsRequest) GetAmountMicros() int32`

GetAmountMicros returns the AmountMicros field if non-nil, zero value otherwise.

### GetAmountMicrosOk

`func (o *AllocateCreditsRequest) GetAmountMicrosOk() (*int32, bool)`

GetAmountMicrosOk returns a tuple with the AmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMicros

`func (o *AllocateCreditsRequest) SetAmountMicros(v int32)`

SetAmountMicros sets AmountMicros field to given value.


### GetSubjectId

`func (o *AllocateCreditsRequest) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *AllocateCreditsRequest) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *AllocateCreditsRequest) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.


### GetReason

`func (o *AllocateCreditsRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AllocateCreditsRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AllocateCreditsRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AllocateCreditsRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *AllocateCreditsRequest) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *AllocateCreditsRequest) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


