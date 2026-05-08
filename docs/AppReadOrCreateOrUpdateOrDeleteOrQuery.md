# AppReadOrCreateOrUpdateOrDeleteOrQuery

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
**Type** | **string** | The app&#39;s type is Stripe. | 
**StripeAccountId** | **string** | The Stripe account ID. | [readonly] 
**Livemode** | **bool** | Livemode, true if the app is in production mode. | [readonly] 
**MaskedAPIKey** | **string** | The masked API key. Only shows the first 8 and last 3 characters. | [readonly] 
**SecretAPIKey** | Pointer to **string** | The Stripe API key. | [optional] 
**EnableDraftSyncHook** | **bool** | Enable draft.sync hook.  If the hook is not enabled, the invoice will be progressed to the next state automatically. | 
**EnableIssuingSyncHook** | **bool** | Enable issuing.sync hook.  If the hook is not enabled, the invoice will be progressed to the next state automatically. | 

## Methods

### NewAppReadOrCreateOrUpdateOrDeleteOrQuery

`func NewAppReadOrCreateOrUpdateOrDeleteOrQuery(id string, name string, createdAt time.Time, updatedAt time.Time, listing MarketplaceListing, status AppStatus, type_ string, stripeAccountId string, livemode bool, maskedAPIKey string, enableDraftSyncHook bool, enableIssuingSyncHook bool, ) *AppReadOrCreateOrUpdateOrDeleteOrQuery`

NewAppReadOrCreateOrUpdateOrDeleteOrQuery instantiates a new AppReadOrCreateOrUpdateOrDeleteOrQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppReadOrCreateOrUpdateOrDeleteOrQueryWithDefaults

`func NewAppReadOrCreateOrUpdateOrDeleteOrQueryWithDefaults() *AppReadOrCreateOrUpdateOrDeleteOrQuery`

NewAppReadOrCreateOrUpdateOrDeleteOrQueryWithDefaults instantiates a new AppReadOrCreateOrUpdateOrDeleteOrQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetListing

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) GetListing() MarketplaceListing`

GetListing returns the Listing field if non-nil, zero value otherwise.

### GetListingOk

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) GetListingOk() (*MarketplaceListing, bool)`

GetListingOk returns a tuple with the Listing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListing

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) SetListing(v MarketplaceListing)`

SetListing sets Listing field to given value.


### GetStatus

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) GetStatus() AppStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) GetStatusOk() (*AppStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) SetStatus(v AppStatus)`

SetStatus sets Status field to given value.


### GetType

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) SetType(v string)`

SetType sets Type field to given value.


### GetStripeAccountId

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) GetStripeAccountId() string`

GetStripeAccountId returns the StripeAccountId field if non-nil, zero value otherwise.

### GetStripeAccountIdOk

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) GetStripeAccountIdOk() (*string, bool)`

GetStripeAccountIdOk returns a tuple with the StripeAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeAccountId

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) SetStripeAccountId(v string)`

SetStripeAccountId sets StripeAccountId field to given value.


### GetLivemode

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.


### GetMaskedAPIKey

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) GetMaskedAPIKey() string`

GetMaskedAPIKey returns the MaskedAPIKey field if non-nil, zero value otherwise.

### GetMaskedAPIKeyOk

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) GetMaskedAPIKeyOk() (*string, bool)`

GetMaskedAPIKeyOk returns a tuple with the MaskedAPIKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskedAPIKey

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) SetMaskedAPIKey(v string)`

SetMaskedAPIKey sets MaskedAPIKey field to given value.


### GetSecretAPIKey

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) GetSecretAPIKey() string`

GetSecretAPIKey returns the SecretAPIKey field if non-nil, zero value otherwise.

### GetSecretAPIKeyOk

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) GetSecretAPIKeyOk() (*string, bool)`

GetSecretAPIKeyOk returns a tuple with the SecretAPIKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAPIKey

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) SetSecretAPIKey(v string)`

SetSecretAPIKey sets SecretAPIKey field to given value.

### HasSecretAPIKey

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) HasSecretAPIKey() bool`

HasSecretAPIKey returns a boolean if a field has been set.

### GetEnableDraftSyncHook

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) GetEnableDraftSyncHook() bool`

GetEnableDraftSyncHook returns the EnableDraftSyncHook field if non-nil, zero value otherwise.

### GetEnableDraftSyncHookOk

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) GetEnableDraftSyncHookOk() (*bool, bool)`

GetEnableDraftSyncHookOk returns a tuple with the EnableDraftSyncHook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDraftSyncHook

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) SetEnableDraftSyncHook(v bool)`

SetEnableDraftSyncHook sets EnableDraftSyncHook field to given value.


### GetEnableIssuingSyncHook

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) GetEnableIssuingSyncHook() bool`

GetEnableIssuingSyncHook returns the EnableIssuingSyncHook field if non-nil, zero value otherwise.

### GetEnableIssuingSyncHookOk

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) GetEnableIssuingSyncHookOk() (*bool, bool)`

GetEnableIssuingSyncHookOk returns a tuple with the EnableIssuingSyncHook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIssuingSyncHook

`func (o *AppReadOrCreateOrUpdateOrDeleteOrQuery) SetEnableIssuingSyncHook(v bool)`

SetEnableIssuingSyncHook sets EnableIssuingSyncHook field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


