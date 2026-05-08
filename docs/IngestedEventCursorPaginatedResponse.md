# IngestedEventCursorPaginatedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]IngestedEvent**](IngestedEvent.md) | The items in the response. | 
**NextCursor** | Pointer to **string** | The cursor of the last item in the list. | [optional] 

## Methods

### NewIngestedEventCursorPaginatedResponse

`func NewIngestedEventCursorPaginatedResponse(items []IngestedEvent, ) *IngestedEventCursorPaginatedResponse`

NewIngestedEventCursorPaginatedResponse instantiates a new IngestedEventCursorPaginatedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestedEventCursorPaginatedResponseWithDefaults

`func NewIngestedEventCursorPaginatedResponseWithDefaults() *IngestedEventCursorPaginatedResponse`

NewIngestedEventCursorPaginatedResponseWithDefaults instantiates a new IngestedEventCursorPaginatedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *IngestedEventCursorPaginatedResponse) GetItems() []IngestedEvent`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *IngestedEventCursorPaginatedResponse) GetItemsOk() (*[]IngestedEvent, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *IngestedEventCursorPaginatedResponse) SetItems(v []IngestedEvent)`

SetItems sets Items field to given value.


### GetNextCursor

`func (o *IngestedEventCursorPaginatedResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *IngestedEventCursorPaginatedResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *IngestedEventCursorPaginatedResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *IngestedEventCursorPaginatedResponse) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


