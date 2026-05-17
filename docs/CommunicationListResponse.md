# CommunicationListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]CommunicationResponse**](CommunicationResponse.md) |  | 
**Total** | **int32** |  | 
**Page** | **int32** |  | 
**PageSize** | **int32** |  | 

## Methods

### NewCommunicationListResponse

`func NewCommunicationListResponse(items []CommunicationResponse, total int32, page int32, pageSize int32, ) *CommunicationListResponse`

NewCommunicationListResponse instantiates a new CommunicationListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommunicationListResponseWithDefaults

`func NewCommunicationListResponseWithDefaults() *CommunicationListResponse`

NewCommunicationListResponseWithDefaults instantiates a new CommunicationListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *CommunicationListResponse) GetItems() []CommunicationResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CommunicationListResponse) GetItemsOk() (*[]CommunicationResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CommunicationListResponse) SetItems(v []CommunicationResponse)`

SetItems sets Items field to given value.


### GetTotal

`func (o *CommunicationListResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CommunicationListResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CommunicationListResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetPage

`func (o *CommunicationListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *CommunicationListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *CommunicationListResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *CommunicationListResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *CommunicationListResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *CommunicationListResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


