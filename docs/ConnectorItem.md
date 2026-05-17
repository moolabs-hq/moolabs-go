# ConnectorItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Category** | **string** |  | 
**Status** | **string** |  | 
**AuthStatus** | **string** |  | 
**LastSync** | **string** |  | 
**MappingInfo** | Pointer to **string** |  | [optional] 
**SyncDirection** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Config** | **map[string]interface{}** |  | 

## Methods

### NewConnectorItem

`func NewConnectorItem(id string, name string, category string, status string, authStatus string, lastSync string, config map[string]interface{}, ) *ConnectorItem`

NewConnectorItem instantiates a new ConnectorItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorItemWithDefaults

`func NewConnectorItemWithDefaults() *ConnectorItem`

NewConnectorItemWithDefaults instantiates a new ConnectorItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConnectorItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectorItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectorItem) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ConnectorItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectorItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectorItem) SetName(v string)`

SetName sets Name field to given value.


### GetCategory

`func (o *ConnectorItem) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ConnectorItem) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ConnectorItem) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetStatus

`func (o *ConnectorItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConnectorItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConnectorItem) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAuthStatus

`func (o *ConnectorItem) GetAuthStatus() string`

GetAuthStatus returns the AuthStatus field if non-nil, zero value otherwise.

### GetAuthStatusOk

`func (o *ConnectorItem) GetAuthStatusOk() (*string, bool)`

GetAuthStatusOk returns a tuple with the AuthStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthStatus

`func (o *ConnectorItem) SetAuthStatus(v string)`

SetAuthStatus sets AuthStatus field to given value.


### GetLastSync

`func (o *ConnectorItem) GetLastSync() string`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *ConnectorItem) GetLastSyncOk() (*string, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *ConnectorItem) SetLastSync(v string)`

SetLastSync sets LastSync field to given value.


### GetMappingInfo

`func (o *ConnectorItem) GetMappingInfo() string`

GetMappingInfo returns the MappingInfo field if non-nil, zero value otherwise.

### GetMappingInfoOk

`func (o *ConnectorItem) GetMappingInfoOk() (*string, bool)`

GetMappingInfoOk returns a tuple with the MappingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingInfo

`func (o *ConnectorItem) SetMappingInfo(v string)`

SetMappingInfo sets MappingInfo field to given value.

### HasMappingInfo

`func (o *ConnectorItem) HasMappingInfo() bool`

HasMappingInfo returns a boolean if a field has been set.

### GetSyncDirection

`func (o *ConnectorItem) GetSyncDirection() string`

GetSyncDirection returns the SyncDirection field if non-nil, zero value otherwise.

### GetSyncDirectionOk

`func (o *ConnectorItem) GetSyncDirectionOk() (*string, bool)`

GetSyncDirectionOk returns a tuple with the SyncDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncDirection

`func (o *ConnectorItem) SetSyncDirection(v string)`

SetSyncDirection sets SyncDirection field to given value.

### HasSyncDirection

`func (o *ConnectorItem) HasSyncDirection() bool`

HasSyncDirection returns a boolean if a field has been set.

### GetOwner

`func (o *ConnectorItem) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ConnectorItem) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ConnectorItem) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ConnectorItem) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetConfig

`func (o *ConnectorItem) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ConnectorItem) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ConnectorItem) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


