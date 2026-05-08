# Progress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **int32** | Success is the number of items that succeeded | 
**Failed** | **int32** | Failed is the number of items that failed | 
**Total** | **int32** | The total number of items to process | 
**UpdatedAt** | **time.Time** | The time the progress was last updated | 

## Methods

### NewProgress

`func NewProgress(success int32, failed int32, total int32, updatedAt time.Time, ) *Progress`

NewProgress instantiates a new Progress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProgressWithDefaults

`func NewProgressWithDefaults() *Progress`

NewProgressWithDefaults instantiates a new Progress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *Progress) GetSuccess() int32`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *Progress) GetSuccessOk() (*int32, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *Progress) SetSuccess(v int32)`

SetSuccess sets Success field to given value.


### GetFailed

`func (o *Progress) GetFailed() int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *Progress) GetFailedOk() (*int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *Progress) SetFailed(v int32)`

SetFailed sets Failed field to given value.


### GetTotal

`func (o *Progress) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Progress) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Progress) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetUpdatedAt

`func (o *Progress) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Progress) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Progress) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


