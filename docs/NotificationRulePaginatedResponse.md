# NotificationRulePaginatedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** | The total number of items. | 
**Page** | **int32** | The page index. | 
**PageSize** | **int32** | The maximum number of items per page. | 
**Items** | [**[]NotificationRule**](NotificationRule.md) | The items in the current page. | 

## Methods

### NewNotificationRulePaginatedResponse

`func NewNotificationRulePaginatedResponse(totalCount int32, page int32, pageSize int32, items []NotificationRule, ) *NotificationRulePaginatedResponse`

NewNotificationRulePaginatedResponse instantiates a new NotificationRulePaginatedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationRulePaginatedResponseWithDefaults

`func NewNotificationRulePaginatedResponseWithDefaults() *NotificationRulePaginatedResponse`

NewNotificationRulePaginatedResponseWithDefaults instantiates a new NotificationRulePaginatedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *NotificationRulePaginatedResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *NotificationRulePaginatedResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *NotificationRulePaginatedResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetPage

`func (o *NotificationRulePaginatedResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *NotificationRulePaginatedResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *NotificationRulePaginatedResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *NotificationRulePaginatedResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *NotificationRulePaginatedResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *NotificationRulePaginatedResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetItems

`func (o *NotificationRulePaginatedResponse) GetItems() []NotificationRule`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *NotificationRulePaginatedResponse) GetItemsOk() (*[]NotificationRule, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *NotificationRulePaginatedResponse) SetItems(v []NotificationRule)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


