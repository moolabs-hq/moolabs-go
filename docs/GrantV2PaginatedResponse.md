# GrantV2PaginatedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** | The total number of items. | 
**Page** | **int32** | The page index. | 
**PageSize** | **int32** | The maximum number of items per page. | 
**Items** | [**[]EntitlementGrantV2**](EntitlementGrantV2.md) | The items in the current page. | 

## Methods

### NewGrantV2PaginatedResponse

`func NewGrantV2PaginatedResponse(totalCount int32, page int32, pageSize int32, items []EntitlementGrantV2, ) *GrantV2PaginatedResponse`

NewGrantV2PaginatedResponse instantiates a new GrantV2PaginatedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantV2PaginatedResponseWithDefaults

`func NewGrantV2PaginatedResponseWithDefaults() *GrantV2PaginatedResponse`

NewGrantV2PaginatedResponseWithDefaults instantiates a new GrantV2PaginatedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *GrantV2PaginatedResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *GrantV2PaginatedResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *GrantV2PaginatedResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetPage

`func (o *GrantV2PaginatedResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GrantV2PaginatedResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GrantV2PaginatedResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *GrantV2PaginatedResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *GrantV2PaginatedResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *GrantV2PaginatedResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetItems

`func (o *GrantV2PaginatedResponse) GetItems() []EntitlementGrantV2`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GrantV2PaginatedResponse) GetItemsOk() (*[]EntitlementGrantV2, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GrantV2PaginatedResponse) SetItems(v []EntitlementGrantV2)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


