# AppReplaceUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Human-readable name for the resource. Between 1 and 256 characters. | 
**Description** | Pointer to **string** | Optional description of the resource. Maximum 1024 characters. | [optional] 
**Metadata** | Pointer to **map[string]string** | Additional metadata for the resource. | [optional] 
**Type** | **string** | The app&#39;s type is Stripe. | 
**SecretAPIKey** | Pointer to **string** | The Stripe API key. | [optional] 
**EnableDraftSyncHook** | **bool** | Enable draft.sync hook.  If the hook is not enabled, the invoice will be progressed to the next state automatically. | 
**EnableIssuingSyncHook** | **bool** | Enable issuing.sync hook.  If the hook is not enabled, the invoice will be progressed to the next state automatically. | 

## Methods

### NewAppReplaceUpdate

`func NewAppReplaceUpdate(name string, type_ string, enableDraftSyncHook bool, enableIssuingSyncHook bool, ) *AppReplaceUpdate`

NewAppReplaceUpdate instantiates a new AppReplaceUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppReplaceUpdateWithDefaults

`func NewAppReplaceUpdateWithDefaults() *AppReplaceUpdate`

NewAppReplaceUpdateWithDefaults instantiates a new AppReplaceUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AppReplaceUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppReplaceUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppReplaceUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AppReplaceUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppReplaceUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppReplaceUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppReplaceUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *AppReplaceUpdate) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AppReplaceUpdate) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AppReplaceUpdate) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AppReplaceUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetType

`func (o *AppReplaceUpdate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppReplaceUpdate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppReplaceUpdate) SetType(v string)`

SetType sets Type field to given value.


### GetSecretAPIKey

`func (o *AppReplaceUpdate) GetSecretAPIKey() string`

GetSecretAPIKey returns the SecretAPIKey field if non-nil, zero value otherwise.

### GetSecretAPIKeyOk

`func (o *AppReplaceUpdate) GetSecretAPIKeyOk() (*string, bool)`

GetSecretAPIKeyOk returns a tuple with the SecretAPIKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAPIKey

`func (o *AppReplaceUpdate) SetSecretAPIKey(v string)`

SetSecretAPIKey sets SecretAPIKey field to given value.

### HasSecretAPIKey

`func (o *AppReplaceUpdate) HasSecretAPIKey() bool`

HasSecretAPIKey returns a boolean if a field has been set.

### GetEnableDraftSyncHook

`func (o *AppReplaceUpdate) GetEnableDraftSyncHook() bool`

GetEnableDraftSyncHook returns the EnableDraftSyncHook field if non-nil, zero value otherwise.

### GetEnableDraftSyncHookOk

`func (o *AppReplaceUpdate) GetEnableDraftSyncHookOk() (*bool, bool)`

GetEnableDraftSyncHookOk returns a tuple with the EnableDraftSyncHook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDraftSyncHook

`func (o *AppReplaceUpdate) SetEnableDraftSyncHook(v bool)`

SetEnableDraftSyncHook sets EnableDraftSyncHook field to given value.


### GetEnableIssuingSyncHook

`func (o *AppReplaceUpdate) GetEnableIssuingSyncHook() bool`

GetEnableIssuingSyncHook returns the EnableIssuingSyncHook field if non-nil, zero value otherwise.

### GetEnableIssuingSyncHookOk

`func (o *AppReplaceUpdate) GetEnableIssuingSyncHookOk() (*bool, bool)`

GetEnableIssuingSyncHookOk returns a tuple with the EnableIssuingSyncHook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIssuingSyncHook

`func (o *AppReplaceUpdate) SetEnableIssuingSyncHook(v bool)`

SetEnableIssuingSyncHook sets EnableIssuingSyncHook field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


