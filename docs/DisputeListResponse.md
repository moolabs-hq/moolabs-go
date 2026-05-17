# DisputeListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]DisputeResponse**](DisputeResponse.md) |  | 
**Total** | **int32** |  | 
**Page** | **int32** |  | 
**PageSize** | **int32** |  | 

## Methods

### NewDisputeListResponse

`func NewDisputeListResponse(items []DisputeResponse, total int32, page int32, pageSize int32, ) *DisputeListResponse`

NewDisputeListResponse instantiates a new DisputeListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisputeListResponseWithDefaults

`func NewDisputeListResponseWithDefaults() *DisputeListResponse`

NewDisputeListResponseWithDefaults instantiates a new DisputeListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *DisputeListResponse) GetItems() []DisputeResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *DisputeListResponse) GetItemsOk() (*[]DisputeResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *DisputeListResponse) SetItems(v []DisputeResponse)`

SetItems sets Items field to given value.


### GetTotal

`func (o *DisputeListResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *DisputeListResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *DisputeListResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetPage

`func (o *DisputeListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *DisputeListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *DisputeListResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *DisputeListResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DisputeListResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DisputeListResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


