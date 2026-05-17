# BulkReplayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventIds** | Pointer to **[]string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**MaxCount** | Pointer to **int32** |  | [optional] [default to 100]

## Methods

### NewBulkReplayRequest

`func NewBulkReplayRequest() *BulkReplayRequest`

NewBulkReplayRequest instantiates a new BulkReplayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkReplayRequestWithDefaults

`func NewBulkReplayRequestWithDefaults() *BulkReplayRequest`

NewBulkReplayRequestWithDefaults instantiates a new BulkReplayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventIds

`func (o *BulkReplayRequest) GetEventIds() []string`

GetEventIds returns the EventIds field if non-nil, zero value otherwise.

### GetEventIdsOk

`func (o *BulkReplayRequest) GetEventIdsOk() (*[]string, bool)`

GetEventIdsOk returns a tuple with the EventIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIds

`func (o *BulkReplayRequest) SetEventIds(v []string)`

SetEventIds sets EventIds field to given value.

### HasEventIds

`func (o *BulkReplayRequest) HasEventIds() bool`

HasEventIds returns a boolean if a field has been set.

### GetSource

`func (o *BulkReplayRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *BulkReplayRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *BulkReplayRequest) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *BulkReplayRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetMaxCount

`func (o *BulkReplayRequest) GetMaxCount() int32`

GetMaxCount returns the MaxCount field if non-nil, zero value otherwise.

### GetMaxCountOk

`func (o *BulkReplayRequest) GetMaxCountOk() (*int32, bool)`

GetMaxCountOk returns a tuple with the MaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCount

`func (o *BulkReplayRequest) SetMaxCount(v int32)`

SetMaxCount sets MaxCount field to given value.

### HasMaxCount

`func (o *BulkReplayRequest) HasMaxCount() bool`

HasMaxCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


