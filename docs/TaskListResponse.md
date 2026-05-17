# TaskListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]TaskResponse**](TaskResponse.md) |  | 
**Total** | **int32** |  | 
**Page** | **int32** |  | 
**PageSize** | **int32** |  | 

## Methods

### NewTaskListResponse

`func NewTaskListResponse(items []TaskResponse, total int32, page int32, pageSize int32, ) *TaskListResponse`

NewTaskListResponse instantiates a new TaskListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskListResponseWithDefaults

`func NewTaskListResponseWithDefaults() *TaskListResponse`

NewTaskListResponseWithDefaults instantiates a new TaskListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *TaskListResponse) GetItems() []TaskResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TaskListResponse) GetItemsOk() (*[]TaskResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TaskListResponse) SetItems(v []TaskResponse)`

SetItems sets Items field to given value.


### GetTotal

`func (o *TaskListResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *TaskListResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *TaskListResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetPage

`func (o *TaskListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *TaskListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *TaskListResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *TaskListResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *TaskListResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *TaskListResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


