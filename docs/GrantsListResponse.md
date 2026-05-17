# GrantsListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]GrantResponse**](GrantResponse.md) |  | 
**Total** | **int32** |  | 
**TotalRemaining** | Pointer to **float32** |  | [optional] 

## Methods

### NewGrantsListResponse

`func NewGrantsListResponse(items []GrantResponse, total int32, ) *GrantsListResponse`

NewGrantsListResponse instantiates a new GrantsListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantsListResponseWithDefaults

`func NewGrantsListResponseWithDefaults() *GrantsListResponse`

NewGrantsListResponseWithDefaults instantiates a new GrantsListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *GrantsListResponse) GetItems() []GrantResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GrantsListResponse) GetItemsOk() (*[]GrantResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GrantsListResponse) SetItems(v []GrantResponse)`

SetItems sets Items field to given value.


### GetTotal

`func (o *GrantsListResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GrantsListResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GrantsListResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetTotalRemaining

`func (o *GrantsListResponse) GetTotalRemaining() float32`

GetTotalRemaining returns the TotalRemaining field if non-nil, zero value otherwise.

### GetTotalRemainingOk

`func (o *GrantsListResponse) GetTotalRemainingOk() (*float32, bool)`

GetTotalRemainingOk returns a tuple with the TotalRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRemaining

`func (o *GrantsListResponse) SetTotalRemaining(v float32)`

SetTotalRemaining sets TotalRemaining field to given value.

### HasTotalRemaining

`func (o *GrantsListResponse) HasTotalRemaining() bool`

HasTotalRemaining returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


