# EvidenceReviewRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResolutionStatus** | **string** | upheld, denied, or partial | 
**ResolvedAmountMicros** | **int32** |  | 

## Methods

### NewEvidenceReviewRequest

`func NewEvidenceReviewRequest(resolutionStatus string, resolvedAmountMicros int32, ) *EvidenceReviewRequest`

NewEvidenceReviewRequest instantiates a new EvidenceReviewRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvidenceReviewRequestWithDefaults

`func NewEvidenceReviewRequestWithDefaults() *EvidenceReviewRequest`

NewEvidenceReviewRequestWithDefaults instantiates a new EvidenceReviewRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResolutionStatus

`func (o *EvidenceReviewRequest) GetResolutionStatus() string`

GetResolutionStatus returns the ResolutionStatus field if non-nil, zero value otherwise.

### GetResolutionStatusOk

`func (o *EvidenceReviewRequest) GetResolutionStatusOk() (*string, bool)`

GetResolutionStatusOk returns a tuple with the ResolutionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionStatus

`func (o *EvidenceReviewRequest) SetResolutionStatus(v string)`

SetResolutionStatus sets ResolutionStatus field to given value.


### GetResolvedAmountMicros

`func (o *EvidenceReviewRequest) GetResolvedAmountMicros() int32`

GetResolvedAmountMicros returns the ResolvedAmountMicros field if non-nil, zero value otherwise.

### GetResolvedAmountMicrosOk

`func (o *EvidenceReviewRequest) GetResolvedAmountMicrosOk() (*int32, bool)`

GetResolvedAmountMicrosOk returns a tuple with the ResolvedAmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedAmountMicros

`func (o *EvidenceReviewRequest) SetResolvedAmountMicros(v int32)`

SetResolvedAmountMicros sets ResolvedAmountMicros field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


