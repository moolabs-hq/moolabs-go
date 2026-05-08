# ListEntitlementsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** | The total number of items. | 
**Page** | **int32** | The page index. | 
**PageSize** | **int32** | The maximum number of items per page. | 
**Items** | [**[]Entitlement**](Entitlement.md) | The items in the current page. | 

## Methods

### NewListEntitlementsResult

`func NewListEntitlementsResult(totalCount int32, page int32, pageSize int32, items []Entitlement, ) *ListEntitlementsResult`

NewListEntitlementsResult instantiates a new ListEntitlementsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEntitlementsResultWithDefaults

`func NewListEntitlementsResultWithDefaults() *ListEntitlementsResult`

NewListEntitlementsResultWithDefaults instantiates a new ListEntitlementsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *ListEntitlementsResult) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListEntitlementsResult) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListEntitlementsResult) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetPage

`func (o *ListEntitlementsResult) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListEntitlementsResult) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListEntitlementsResult) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *ListEntitlementsResult) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListEntitlementsResult) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListEntitlementsResult) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetItems

`func (o *ListEntitlementsResult) GetItems() []Entitlement`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListEntitlementsResult) GetItemsOk() (*[]Entitlement, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListEntitlementsResult) SetItems(v []Entitlement)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


