# BuyerRequestChangesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Note** | **string** |  | 
**RequestedChanges** | Pointer to [**BuyerRequestedChanges**](BuyerRequestedChanges.md) |  | [optional] 

## Methods

### NewBuyerRequestChangesRequest

`func NewBuyerRequestChangesRequest(note string, ) *BuyerRequestChangesRequest`

NewBuyerRequestChangesRequest instantiates a new BuyerRequestChangesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuyerRequestChangesRequestWithDefaults

`func NewBuyerRequestChangesRequestWithDefaults() *BuyerRequestChangesRequest`

NewBuyerRequestChangesRequestWithDefaults instantiates a new BuyerRequestChangesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNote

`func (o *BuyerRequestChangesRequest) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *BuyerRequestChangesRequest) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *BuyerRequestChangesRequest) SetNote(v string)`

SetNote sets Note field to given value.


### GetRequestedChanges

`func (o *BuyerRequestChangesRequest) GetRequestedChanges() BuyerRequestedChanges`

GetRequestedChanges returns the RequestedChanges field if non-nil, zero value otherwise.

### GetRequestedChangesOk

`func (o *BuyerRequestChangesRequest) GetRequestedChangesOk() (*BuyerRequestedChanges, bool)`

GetRequestedChangesOk returns a tuple with the RequestedChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedChanges

`func (o *BuyerRequestChangesRequest) SetRequestedChanges(v BuyerRequestedChanges)`

SetRequestedChanges sets RequestedChanges field to given value.

### HasRequestedChanges

`func (o *BuyerRequestChangesRequest) HasRequestedChanges() bool`

HasRequestedChanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


