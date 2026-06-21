# SaveRedlineEditRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocKey** | **string** |  | 
**BlockId** | **int32** |  | 
**Op** | **string** |  | 
**Before** | **string** |  | 
**After** | **string** |  | 

## Methods

### NewSaveRedlineEditRequest

`func NewSaveRedlineEditRequest(docKey string, blockId int32, op string, before string, after string, ) *SaveRedlineEditRequest`

NewSaveRedlineEditRequest instantiates a new SaveRedlineEditRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveRedlineEditRequestWithDefaults

`func NewSaveRedlineEditRequestWithDefaults() *SaveRedlineEditRequest`

NewSaveRedlineEditRequestWithDefaults instantiates a new SaveRedlineEditRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocKey

`func (o *SaveRedlineEditRequest) GetDocKey() string`

GetDocKey returns the DocKey field if non-nil, zero value otherwise.

### GetDocKeyOk

`func (o *SaveRedlineEditRequest) GetDocKeyOk() (*string, bool)`

GetDocKeyOk returns a tuple with the DocKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocKey

`func (o *SaveRedlineEditRequest) SetDocKey(v string)`

SetDocKey sets DocKey field to given value.


### GetBlockId

`func (o *SaveRedlineEditRequest) GetBlockId() int32`

GetBlockId returns the BlockId field if non-nil, zero value otherwise.

### GetBlockIdOk

`func (o *SaveRedlineEditRequest) GetBlockIdOk() (*int32, bool)`

GetBlockIdOk returns a tuple with the BlockId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockId

`func (o *SaveRedlineEditRequest) SetBlockId(v int32)`

SetBlockId sets BlockId field to given value.


### GetOp

`func (o *SaveRedlineEditRequest) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *SaveRedlineEditRequest) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *SaveRedlineEditRequest) SetOp(v string)`

SetOp sets Op field to given value.


### GetBefore

`func (o *SaveRedlineEditRequest) GetBefore() string`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *SaveRedlineEditRequest) GetBeforeOk() (*string, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *SaveRedlineEditRequest) SetBefore(v string)`

SetBefore sets Before field to given value.


### GetAfter

`func (o *SaveRedlineEditRequest) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *SaveRedlineEditRequest) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *SaveRedlineEditRequest) SetAfter(v string)`

SetAfter sets After field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


