# ImportPromisesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Promises** | [**[]ImportPromiseItem**](ImportPromiseItem.md) |  | 

## Methods

### NewImportPromisesRequest

`func NewImportPromisesRequest(promises []ImportPromiseItem, ) *ImportPromisesRequest`

NewImportPromisesRequest instantiates a new ImportPromisesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportPromisesRequestWithDefaults

`func NewImportPromisesRequestWithDefaults() *ImportPromisesRequest`

NewImportPromisesRequestWithDefaults instantiates a new ImportPromisesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPromises

`func (o *ImportPromisesRequest) GetPromises() []ImportPromiseItem`

GetPromises returns the Promises field if non-nil, zero value otherwise.

### GetPromisesOk

`func (o *ImportPromisesRequest) GetPromisesOk() (*[]ImportPromiseItem, bool)`

GetPromisesOk returns a tuple with the Promises field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromises

`func (o *ImportPromisesRequest) SetPromises(v []ImportPromiseItem)`

SetPromises sets Promises field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


