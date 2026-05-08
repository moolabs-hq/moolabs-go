# NotificationChannelWebhookCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Notification channel type. | 
**Name** | **string** | User friendly name of the channel. | 
**Disabled** | Pointer to **bool** | Whether the channel is disabled or not. | [optional] [default to false]
**Metadata** | Pointer to **map[string]string** | Additional metadata for the resource. | [optional] 
**Url** | **string** | Webhook URL where the notification is sent. | 
**CustomHeaders** | Pointer to **map[string]string** | Custom HTTP headers sent as part of the webhook request. | [optional] 
**SigningSecret** | Pointer to **string** | Signing secret used for webhook request validation on the receiving end.  Format: &#x60;base64&#x60; encoded random bytes optionally prefixed with &#x60;whsec_&#x60;. Recommended size: 24 | [optional] 

## Methods

### NewNotificationChannelWebhookCreateRequest

`func NewNotificationChannelWebhookCreateRequest(type_ string, name string, url string, ) *NotificationChannelWebhookCreateRequest`

NewNotificationChannelWebhookCreateRequest instantiates a new NotificationChannelWebhookCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationChannelWebhookCreateRequestWithDefaults

`func NewNotificationChannelWebhookCreateRequestWithDefaults() *NotificationChannelWebhookCreateRequest`

NewNotificationChannelWebhookCreateRequestWithDefaults instantiates a new NotificationChannelWebhookCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NotificationChannelWebhookCreateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationChannelWebhookCreateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationChannelWebhookCreateRequest) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *NotificationChannelWebhookCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationChannelWebhookCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationChannelWebhookCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDisabled

`func (o *NotificationChannelWebhookCreateRequest) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *NotificationChannelWebhookCreateRequest) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *NotificationChannelWebhookCreateRequest) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *NotificationChannelWebhookCreateRequest) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetMetadata

`func (o *NotificationChannelWebhookCreateRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NotificationChannelWebhookCreateRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NotificationChannelWebhookCreateRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NotificationChannelWebhookCreateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetUrl

`func (o *NotificationChannelWebhookCreateRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NotificationChannelWebhookCreateRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NotificationChannelWebhookCreateRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetCustomHeaders

`func (o *NotificationChannelWebhookCreateRequest) GetCustomHeaders() map[string]string`

GetCustomHeaders returns the CustomHeaders field if non-nil, zero value otherwise.

### GetCustomHeadersOk

`func (o *NotificationChannelWebhookCreateRequest) GetCustomHeadersOk() (*map[string]string, bool)`

GetCustomHeadersOk returns a tuple with the CustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaders

`func (o *NotificationChannelWebhookCreateRequest) SetCustomHeaders(v map[string]string)`

SetCustomHeaders sets CustomHeaders field to given value.

### HasCustomHeaders

`func (o *NotificationChannelWebhookCreateRequest) HasCustomHeaders() bool`

HasCustomHeaders returns a boolean if a field has been set.

### GetSigningSecret

`func (o *NotificationChannelWebhookCreateRequest) GetSigningSecret() string`

GetSigningSecret returns the SigningSecret field if non-nil, zero value otherwise.

### GetSigningSecretOk

`func (o *NotificationChannelWebhookCreateRequest) GetSigningSecretOk() (*string, bool)`

GetSigningSecretOk returns a tuple with the SigningSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningSecret

`func (o *NotificationChannelWebhookCreateRequest) SetSigningSecret(v string)`

SetSigningSecret sets SigningSecret field to given value.

### HasSigningSecret

`func (o *NotificationChannelWebhookCreateRequest) HasSigningSecret() bool`

HasSigningSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


