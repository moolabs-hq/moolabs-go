# PaginatedCreditMemos

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]CreditMemoItem**](CreditMemoItem.md) |  | 
**NextCursor** | Pointer to **string** |  | [optional] 

## Methods

### NewPaginatedCreditMemos

`func NewPaginatedCreditMemos(items []CreditMemoItem, ) *PaginatedCreditMemos`

NewPaginatedCreditMemos instantiates a new PaginatedCreditMemos object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedCreditMemosWithDefaults

`func NewPaginatedCreditMemosWithDefaults() *PaginatedCreditMemos`

NewPaginatedCreditMemosWithDefaults instantiates a new PaginatedCreditMemos object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *PaginatedCreditMemos) GetItems() []CreditMemoItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PaginatedCreditMemos) GetItemsOk() (*[]CreditMemoItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PaginatedCreditMemos) SetItems(v []CreditMemoItem)`

SetItems sets Items field to given value.


### GetNextCursor

`func (o *PaginatedCreditMemos) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *PaginatedCreditMemos) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *PaginatedCreditMemos) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *PaginatedCreditMemos) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


