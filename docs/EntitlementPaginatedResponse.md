# EntitlementPaginatedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** | The total number of items. | 
**Page** | **int32** | The page index. | 
**PageSize** | **int32** | The maximum number of items per page. | 
**Items** | [**[]Entitlement**](Entitlement.md) | The items in the current page. | 

## Methods

### NewEntitlementPaginatedResponse

`func NewEntitlementPaginatedResponse(totalCount int32, page int32, pageSize int32, items []Entitlement, ) *EntitlementPaginatedResponse`

NewEntitlementPaginatedResponse instantiates a new EntitlementPaginatedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementPaginatedResponseWithDefaults

`func NewEntitlementPaginatedResponseWithDefaults() *EntitlementPaginatedResponse`

NewEntitlementPaginatedResponseWithDefaults instantiates a new EntitlementPaginatedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *EntitlementPaginatedResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *EntitlementPaginatedResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *EntitlementPaginatedResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetPage

`func (o *EntitlementPaginatedResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *EntitlementPaginatedResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *EntitlementPaginatedResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *EntitlementPaginatedResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *EntitlementPaginatedResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *EntitlementPaginatedResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetItems

`func (o *EntitlementPaginatedResponse) GetItems() []Entitlement`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *EntitlementPaginatedResponse) GetItemsOk() (*[]Entitlement, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *EntitlementPaginatedResponse) SetItems(v []Entitlement)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


