# ImportDisputesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disputes** | [**[]ImportDisputeItem**](ImportDisputeItem.md) |  | 

## Methods

### NewImportDisputesRequest

`func NewImportDisputesRequest(disputes []ImportDisputeItem, ) *ImportDisputesRequest`

NewImportDisputesRequest instantiates a new ImportDisputesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportDisputesRequestWithDefaults

`func NewImportDisputesRequestWithDefaults() *ImportDisputesRequest`

NewImportDisputesRequestWithDefaults instantiates a new ImportDisputesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisputes

`func (o *ImportDisputesRequest) GetDisputes() []ImportDisputeItem`

GetDisputes returns the Disputes field if non-nil, zero value otherwise.

### GetDisputesOk

`func (o *ImportDisputesRequest) GetDisputesOk() (*[]ImportDisputeItem, bool)`

GetDisputesOk returns a tuple with the Disputes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputes

`func (o *ImportDisputesRequest) SetDisputes(v []ImportDisputeItem)`

SetDisputes sets Disputes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


