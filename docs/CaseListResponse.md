# CaseListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]CaseListItemResponse**](CaseListItemResponse.md) |  | 
**Total** | **int32** |  | 
**Page** | **int32** |  | 
**PageSize** | **int32** |  | 

## Methods

### NewCaseListResponse

`func NewCaseListResponse(items []CaseListItemResponse, total int32, page int32, pageSize int32, ) *CaseListResponse`

NewCaseListResponse instantiates a new CaseListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaseListResponseWithDefaults

`func NewCaseListResponseWithDefaults() *CaseListResponse`

NewCaseListResponseWithDefaults instantiates a new CaseListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *CaseListResponse) GetItems() []CaseListItemResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CaseListResponse) GetItemsOk() (*[]CaseListItemResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CaseListResponse) SetItems(v []CaseListItemResponse)`

SetItems sets Items field to given value.


### GetTotal

`func (o *CaseListResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CaseListResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CaseListResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetPage

`func (o *CaseListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *CaseListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *CaseListResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *CaseListResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *CaseListResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *CaseListResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


