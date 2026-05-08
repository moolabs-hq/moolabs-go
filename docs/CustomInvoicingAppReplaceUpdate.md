# CustomInvoicingAppReplaceUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Human-readable name for the resource. Between 1 and 256 characters. | 
**Description** | Pointer to **string** | Optional description of the resource. Maximum 1024 characters. | [optional] 
**Metadata** | Pointer to **map[string]string** | Additional metadata for the resource. | [optional] 
**Type** | **string** | The app&#39;s type is CustomInvoicing. | 
**EnableDraftSyncHook** | **bool** | Enable draft.sync hook.  If the hook is not enabled, the invoice will be progressed to the next state automatically. | 
**EnableIssuingSyncHook** | **bool** | Enable issuing.sync hook.  If the hook is not enabled, the invoice will be progressed to the next state automatically. | 

## Methods

### NewCustomInvoicingAppReplaceUpdate

`func NewCustomInvoicingAppReplaceUpdate(name string, type_ string, enableDraftSyncHook bool, enableIssuingSyncHook bool, ) *CustomInvoicingAppReplaceUpdate`

NewCustomInvoicingAppReplaceUpdate instantiates a new CustomInvoicingAppReplaceUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomInvoicingAppReplaceUpdateWithDefaults

`func NewCustomInvoicingAppReplaceUpdateWithDefaults() *CustomInvoicingAppReplaceUpdate`

NewCustomInvoicingAppReplaceUpdateWithDefaults instantiates a new CustomInvoicingAppReplaceUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CustomInvoicingAppReplaceUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomInvoicingAppReplaceUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomInvoicingAppReplaceUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CustomInvoicingAppReplaceUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomInvoicingAppReplaceUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomInvoicingAppReplaceUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomInvoicingAppReplaceUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *CustomInvoicingAppReplaceUpdate) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CustomInvoicingAppReplaceUpdate) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CustomInvoicingAppReplaceUpdate) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CustomInvoicingAppReplaceUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetType

`func (o *CustomInvoicingAppReplaceUpdate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomInvoicingAppReplaceUpdate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomInvoicingAppReplaceUpdate) SetType(v string)`

SetType sets Type field to given value.


### GetEnableDraftSyncHook

`func (o *CustomInvoicingAppReplaceUpdate) GetEnableDraftSyncHook() bool`

GetEnableDraftSyncHook returns the EnableDraftSyncHook field if non-nil, zero value otherwise.

### GetEnableDraftSyncHookOk

`func (o *CustomInvoicingAppReplaceUpdate) GetEnableDraftSyncHookOk() (*bool, bool)`

GetEnableDraftSyncHookOk returns a tuple with the EnableDraftSyncHook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDraftSyncHook

`func (o *CustomInvoicingAppReplaceUpdate) SetEnableDraftSyncHook(v bool)`

SetEnableDraftSyncHook sets EnableDraftSyncHook field to given value.


### GetEnableIssuingSyncHook

`func (o *CustomInvoicingAppReplaceUpdate) GetEnableIssuingSyncHook() bool`

GetEnableIssuingSyncHook returns the EnableIssuingSyncHook field if non-nil, zero value otherwise.

### GetEnableIssuingSyncHookOk

`func (o *CustomInvoicingAppReplaceUpdate) GetEnableIssuingSyncHookOk() (*bool, bool)`

GetEnableIssuingSyncHookOk returns a tuple with the EnableIssuingSyncHook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIssuingSyncHook

`func (o *CustomInvoicingAppReplaceUpdate) SetEnableIssuingSyncHook(v bool)`

SetEnableIssuingSyncHook sets EnableIssuingSyncHook field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


