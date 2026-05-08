# InvoicePaginatedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** | The total number of items. | 
**Page** | **int32** | The page index. | 
**PageSize** | **int32** | The maximum number of items per page. | 
**Items** | [**[]Invoice**](Invoice.md) | The items in the current page. | 

## Methods

### NewInvoicePaginatedResponse

`func NewInvoicePaginatedResponse(totalCount int32, page int32, pageSize int32, items []Invoice, ) *InvoicePaginatedResponse`

NewInvoicePaginatedResponse instantiates a new InvoicePaginatedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoicePaginatedResponseWithDefaults

`func NewInvoicePaginatedResponseWithDefaults() *InvoicePaginatedResponse`

NewInvoicePaginatedResponseWithDefaults instantiates a new InvoicePaginatedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *InvoicePaginatedResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *InvoicePaginatedResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *InvoicePaginatedResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetPage

`func (o *InvoicePaginatedResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *InvoicePaginatedResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *InvoicePaginatedResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *InvoicePaginatedResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *InvoicePaginatedResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *InvoicePaginatedResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetItems

`func (o *InvoicePaginatedResponse) GetItems() []Invoice`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InvoicePaginatedResponse) GetItemsOk() (*[]Invoice, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InvoicePaginatedResponse) SetItems(v []Invoice)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


