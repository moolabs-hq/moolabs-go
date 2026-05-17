# AllocationsListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]PaymentAllocationResponse**](PaymentAllocationResponse.md) |  | 
**Total** | **int32** |  | 

## Methods

### NewAllocationsListResponse

`func NewAllocationsListResponse(items []PaymentAllocationResponse, total int32, ) *AllocationsListResponse`

NewAllocationsListResponse instantiates a new AllocationsListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllocationsListResponseWithDefaults

`func NewAllocationsListResponseWithDefaults() *AllocationsListResponse`

NewAllocationsListResponseWithDefaults instantiates a new AllocationsListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *AllocationsListResponse) GetItems() []PaymentAllocationResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AllocationsListResponse) GetItemsOk() (*[]PaymentAllocationResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AllocationsListResponse) SetItems(v []PaymentAllocationResponse)`

SetItems sets Items field to given value.


### GetTotal

`func (o *AllocationsListResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AllocationsListResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AllocationsListResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


