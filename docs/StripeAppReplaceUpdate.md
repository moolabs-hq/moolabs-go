# StripeAppReplaceUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Human-readable name for the resource. Between 1 and 256 characters. | 
**Description** | Pointer to **string** | Optional description of the resource. Maximum 1024 characters. | [optional] 
**Metadata** | Pointer to **map[string]string** | Additional metadata for the resource. | [optional] 
**Type** | **string** | The app&#39;s type is Stripe. | 
**SecretAPIKey** | Pointer to **string** | The Stripe API key. | [optional] 

## Methods

### NewStripeAppReplaceUpdate

`func NewStripeAppReplaceUpdate(name string, type_ string, ) *StripeAppReplaceUpdate`

NewStripeAppReplaceUpdate instantiates a new StripeAppReplaceUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStripeAppReplaceUpdateWithDefaults

`func NewStripeAppReplaceUpdateWithDefaults() *StripeAppReplaceUpdate`

NewStripeAppReplaceUpdateWithDefaults instantiates a new StripeAppReplaceUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StripeAppReplaceUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StripeAppReplaceUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StripeAppReplaceUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *StripeAppReplaceUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StripeAppReplaceUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StripeAppReplaceUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StripeAppReplaceUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *StripeAppReplaceUpdate) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *StripeAppReplaceUpdate) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *StripeAppReplaceUpdate) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *StripeAppReplaceUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetType

`func (o *StripeAppReplaceUpdate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StripeAppReplaceUpdate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StripeAppReplaceUpdate) SetType(v string)`

SetType sets Type field to given value.


### GetSecretAPIKey

`func (o *StripeAppReplaceUpdate) GetSecretAPIKey() string`

GetSecretAPIKey returns the SecretAPIKey field if non-nil, zero value otherwise.

### GetSecretAPIKeyOk

`func (o *StripeAppReplaceUpdate) GetSecretAPIKeyOk() (*string, bool)`

GetSecretAPIKeyOk returns a tuple with the SecretAPIKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAPIKey

`func (o *StripeAppReplaceUpdate) SetSecretAPIKey(v string)`

SetSecretAPIKey sets SecretAPIKey field to given value.

### HasSecretAPIKey

`func (o *StripeAppReplaceUpdate) HasSecretAPIKey() bool`

HasSecretAPIKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


