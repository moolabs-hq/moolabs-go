# EvidenceListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]EvidenceResponse**](EvidenceResponse.md) |  | 
**Total** | **int32** |  | 

## Methods

### NewEvidenceListResponse

`func NewEvidenceListResponse(items []EvidenceResponse, total int32, ) *EvidenceListResponse`

NewEvidenceListResponse instantiates a new EvidenceListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvidenceListResponseWithDefaults

`func NewEvidenceListResponseWithDefaults() *EvidenceListResponse`

NewEvidenceListResponseWithDefaults instantiates a new EvidenceListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *EvidenceListResponse) GetItems() []EvidenceResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *EvidenceListResponse) GetItemsOk() (*[]EvidenceResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *EvidenceListResponse) SetItems(v []EvidenceResponse)`

SetItems sets Items field to given value.


### GetTotal

`func (o *EvidenceListResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *EvidenceListResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *EvidenceListResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


