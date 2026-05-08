# CustomInvoicingApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier for the resource. | [readonly] 
**Name** | **string** | Human-readable name for the resource. Between 1 and 256 characters. | 
**Description** | Pointer to **string** | Optional description of the resource. Maximum 1024 characters. | [optional] 
**Metadata** | Pointer to **map[string]string** | Additional metadata for the resource. | [optional] 
**CreatedAt** | **time.Time** | Timestamp of when the resource was created. | [readonly] 
**UpdatedAt** | **time.Time** | Timestamp of when the resource was last updated. | [readonly] 
**DeletedAt** | Pointer to **time.Time** | Timestamp of when the resource was permanently deleted. | [optional] [readonly] 
**Listing** | [**MarketplaceListing**](MarketplaceListing.md) | The marketplace listing that this installed app is based on. | [readonly] 
**Status** | [**AppStatus**](AppStatus.md) | Status of the app connection. | [readonly] 
**Type** | **string** | The app&#39;s type is CustomInvoicing. | 
**EnableDraftSyncHook** | **bool** | Enable draft.sync hook.  If the hook is not enabled, the invoice will be progressed to the next state automatically. | 
**EnableIssuingSyncHook** | **bool** | Enable issuing.sync hook.  If the hook is not enabled, the invoice will be progressed to the next state automatically. | 

## Methods

### NewCustomInvoicingApp

`func NewCustomInvoicingApp(id string, name string, createdAt time.Time, updatedAt time.Time, listing MarketplaceListing, status AppStatus, type_ string, enableDraftSyncHook bool, enableIssuingSyncHook bool, ) *CustomInvoicingApp`

NewCustomInvoicingApp instantiates a new CustomInvoicingApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomInvoicingAppWithDefaults

`func NewCustomInvoicingAppWithDefaults() *CustomInvoicingApp`

NewCustomInvoicingAppWithDefaults instantiates a new CustomInvoicingApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomInvoicingApp) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomInvoicingApp) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomInvoicingApp) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CustomInvoicingApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomInvoicingApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomInvoicingApp) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CustomInvoicingApp) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomInvoicingApp) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomInvoicingApp) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomInvoicingApp) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *CustomInvoicingApp) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CustomInvoicingApp) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CustomInvoicingApp) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CustomInvoicingApp) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CustomInvoicingApp) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomInvoicingApp) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomInvoicingApp) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CustomInvoicingApp) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CustomInvoicingApp) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CustomInvoicingApp) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *CustomInvoicingApp) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *CustomInvoicingApp) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *CustomInvoicingApp) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *CustomInvoicingApp) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetListing

`func (o *CustomInvoicingApp) GetListing() MarketplaceListing`

GetListing returns the Listing field if non-nil, zero value otherwise.

### GetListingOk

`func (o *CustomInvoicingApp) GetListingOk() (*MarketplaceListing, bool)`

GetListingOk returns a tuple with the Listing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListing

`func (o *CustomInvoicingApp) SetListing(v MarketplaceListing)`

SetListing sets Listing field to given value.


### GetStatus

`func (o *CustomInvoicingApp) GetStatus() AppStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CustomInvoicingApp) GetStatusOk() (*AppStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CustomInvoicingApp) SetStatus(v AppStatus)`

SetStatus sets Status field to given value.


### GetType

`func (o *CustomInvoicingApp) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomInvoicingApp) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomInvoicingApp) SetType(v string)`

SetType sets Type field to given value.


### GetEnableDraftSyncHook

`func (o *CustomInvoicingApp) GetEnableDraftSyncHook() bool`

GetEnableDraftSyncHook returns the EnableDraftSyncHook field if non-nil, zero value otherwise.

### GetEnableDraftSyncHookOk

`func (o *CustomInvoicingApp) GetEnableDraftSyncHookOk() (*bool, bool)`

GetEnableDraftSyncHookOk returns a tuple with the EnableDraftSyncHook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDraftSyncHook

`func (o *CustomInvoicingApp) SetEnableDraftSyncHook(v bool)`

SetEnableDraftSyncHook sets EnableDraftSyncHook field to given value.


### GetEnableIssuingSyncHook

`func (o *CustomInvoicingApp) GetEnableIssuingSyncHook() bool`

GetEnableIssuingSyncHook returns the EnableIssuingSyncHook field if non-nil, zero value otherwise.

### GetEnableIssuingSyncHookOk

`func (o *CustomInvoicingApp) GetEnableIssuingSyncHookOk() (*bool, bool)`

GetEnableIssuingSyncHookOk returns a tuple with the EnableIssuingSyncHook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIssuingSyncHook

`func (o *CustomInvoicingApp) SetEnableIssuingSyncHook(v bool)`

SetEnableIssuingSyncHook sets EnableIssuingSyncHook field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


