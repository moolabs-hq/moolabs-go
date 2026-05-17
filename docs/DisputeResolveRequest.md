# DisputeResolveRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resolution** | **string** | resolved_valid or resolved_invalid | 
**ResolvedAmountMicros** | Pointer to **int32** |  | [optional] 
**ResolutionNotes** | Pointer to **string** |  | [optional] 

## Methods

### NewDisputeResolveRequest

`func NewDisputeResolveRequest(resolution string, ) *DisputeResolveRequest`

NewDisputeResolveRequest instantiates a new DisputeResolveRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisputeResolveRequestWithDefaults

`func NewDisputeResolveRequestWithDefaults() *DisputeResolveRequest`

NewDisputeResolveRequestWithDefaults instantiates a new DisputeResolveRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResolution

`func (o *DisputeResolveRequest) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *DisputeResolveRequest) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *DisputeResolveRequest) SetResolution(v string)`

SetResolution sets Resolution field to given value.


### GetResolvedAmountMicros

`func (o *DisputeResolveRequest) GetResolvedAmountMicros() int32`

GetResolvedAmountMicros returns the ResolvedAmountMicros field if non-nil, zero value otherwise.

### GetResolvedAmountMicrosOk

`func (o *DisputeResolveRequest) GetResolvedAmountMicrosOk() (*int32, bool)`

GetResolvedAmountMicrosOk returns a tuple with the ResolvedAmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedAmountMicros

`func (o *DisputeResolveRequest) SetResolvedAmountMicros(v int32)`

SetResolvedAmountMicros sets ResolvedAmountMicros field to given value.

### HasResolvedAmountMicros

`func (o *DisputeResolveRequest) HasResolvedAmountMicros() bool`

HasResolvedAmountMicros returns a boolean if a field has been set.

### GetResolutionNotes

`func (o *DisputeResolveRequest) GetResolutionNotes() string`

GetResolutionNotes returns the ResolutionNotes field if non-nil, zero value otherwise.

### GetResolutionNotesOk

`func (o *DisputeResolveRequest) GetResolutionNotesOk() (*string, bool)`

GetResolutionNotesOk returns a tuple with the ResolutionNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionNotes

`func (o *DisputeResolveRequest) SetResolutionNotes(v string)`

SetResolutionNotes sets ResolutionNotes field to given value.

### HasResolutionNotes

`func (o *DisputeResolveRequest) HasResolutionNotes() bool`

HasResolutionNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


