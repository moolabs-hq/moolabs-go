# SyncRunItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**Status** | **string** |  | 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**FinishedAt** | Pointer to **time.Time** |  | [optional] 
**ErrorSummary** | Pointer to **map[string]interface{}** |  | [optional] 
**RowCounts** | Pointer to **map[string]interface{}** |  | [optional] 
**IsBootstrap** | **bool** |  | 

## Methods

### NewSyncRunItem

`func NewSyncRunItem(id string, tenantId string, status string, isBootstrap bool, ) *SyncRunItem`

NewSyncRunItem instantiates a new SyncRunItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncRunItemWithDefaults

`func NewSyncRunItemWithDefaults() *SyncRunItem`

NewSyncRunItemWithDefaults instantiates a new SyncRunItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SyncRunItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SyncRunItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SyncRunItem) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *SyncRunItem) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SyncRunItem) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SyncRunItem) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetStatus

`func (o *SyncRunItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SyncRunItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SyncRunItem) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStartedAt

`func (o *SyncRunItem) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *SyncRunItem) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *SyncRunItem) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *SyncRunItem) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetFinishedAt

`func (o *SyncRunItem) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *SyncRunItem) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *SyncRunItem) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *SyncRunItem) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetErrorSummary

`func (o *SyncRunItem) GetErrorSummary() map[string]interface{}`

GetErrorSummary returns the ErrorSummary field if non-nil, zero value otherwise.

### GetErrorSummaryOk

`func (o *SyncRunItem) GetErrorSummaryOk() (*map[string]interface{}, bool)`

GetErrorSummaryOk returns a tuple with the ErrorSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorSummary

`func (o *SyncRunItem) SetErrorSummary(v map[string]interface{})`

SetErrorSummary sets ErrorSummary field to given value.

### HasErrorSummary

`func (o *SyncRunItem) HasErrorSummary() bool`

HasErrorSummary returns a boolean if a field has been set.

### GetRowCounts

`func (o *SyncRunItem) GetRowCounts() map[string]interface{}`

GetRowCounts returns the RowCounts field if non-nil, zero value otherwise.

### GetRowCountsOk

`func (o *SyncRunItem) GetRowCountsOk() (*map[string]interface{}, bool)`

GetRowCountsOk returns a tuple with the RowCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowCounts

`func (o *SyncRunItem) SetRowCounts(v map[string]interface{})`

SetRowCounts sets RowCounts field to given value.

### HasRowCounts

`func (o *SyncRunItem) HasRowCounts() bool`

HasRowCounts returns a boolean if a field has been set.

### GetIsBootstrap

`func (o *SyncRunItem) GetIsBootstrap() bool`

GetIsBootstrap returns the IsBootstrap field if non-nil, zero value otherwise.

### GetIsBootstrapOk

`func (o *SyncRunItem) GetIsBootstrapOk() (*bool, bool)`

GetIsBootstrapOk returns a tuple with the IsBootstrap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBootstrap

`func (o *SyncRunItem) SetIsBootstrap(v bool)`

SetIsBootstrap sets IsBootstrap field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


