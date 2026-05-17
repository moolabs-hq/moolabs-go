# EscalationListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]EscalationResponse**](EscalationResponse.md) |  | 
**Total** | **int32** |  | 
**Page** | **int32** |  | 
**PageSize** | **int32** |  | 

## Methods

### NewEscalationListResponse

`func NewEscalationListResponse(items []EscalationResponse, total int32, page int32, pageSize int32, ) *EscalationListResponse`

NewEscalationListResponse instantiates a new EscalationListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEscalationListResponseWithDefaults

`func NewEscalationListResponseWithDefaults() *EscalationListResponse`

NewEscalationListResponseWithDefaults instantiates a new EscalationListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *EscalationListResponse) GetItems() []EscalationResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *EscalationListResponse) GetItemsOk() (*[]EscalationResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *EscalationListResponse) SetItems(v []EscalationResponse)`

SetItems sets Items field to given value.


### GetTotal

`func (o *EscalationListResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *EscalationListResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *EscalationListResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetPage

`func (o *EscalationListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *EscalationListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *EscalationListResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *EscalationListResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *EscalationListResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *EscalationListResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


