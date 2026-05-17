# PaginatedCustomers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**Items** | [**[]CustomerItem**](CustomerItem.md) |  | 
**NextCursor** | Pointer to **string** |  | [optional] 

## Methods

### NewPaginatedCustomers

`func NewPaginatedCustomers(tenantId string, items []CustomerItem, ) *PaginatedCustomers`

NewPaginatedCustomers instantiates a new PaginatedCustomers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedCustomersWithDefaults

`func NewPaginatedCustomersWithDefaults() *PaginatedCustomers`

NewPaginatedCustomersWithDefaults instantiates a new PaginatedCustomers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *PaginatedCustomers) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *PaginatedCustomers) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *PaginatedCustomers) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetItems

`func (o *PaginatedCustomers) GetItems() []CustomerItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PaginatedCustomers) GetItemsOk() (*[]CustomerItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PaginatedCustomers) SetItems(v []CustomerItem)`

SetItems sets Items field to given value.


### GetNextCursor

`func (o *PaginatedCustomers) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *PaginatedCustomers) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *PaginatedCustomers) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *PaginatedCustomers) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


