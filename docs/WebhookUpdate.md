# WebhookUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** |  | [optional] 
**AuthType** | Pointer to **string** |  | [optional] [default to "none"]
**Secret** | Pointer to **string** |  | [optional] 
**SubscribedEvents** | Pointer to **[]string** |  | [optional] 
**OwnerEmail** | Pointer to **string** |  | [optional] 

## Methods

### NewWebhookUpdate

`func NewWebhookUpdate() *WebhookUpdate`

NewWebhookUpdate instantiates a new WebhookUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookUpdateWithDefaults

`func NewWebhookUpdateWithDefaults() *WebhookUpdate`

NewWebhookUpdateWithDefaults instantiates a new WebhookUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *WebhookUpdate) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookUpdate) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookUpdate) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WebhookUpdate) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetAuthType

`func (o *WebhookUpdate) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *WebhookUpdate) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *WebhookUpdate) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *WebhookUpdate) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetSecret

`func (o *WebhookUpdate) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *WebhookUpdate) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *WebhookUpdate) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *WebhookUpdate) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetSubscribedEvents

`func (o *WebhookUpdate) GetSubscribedEvents() []string`

GetSubscribedEvents returns the SubscribedEvents field if non-nil, zero value otherwise.

### GetSubscribedEventsOk

`func (o *WebhookUpdate) GetSubscribedEventsOk() (*[]string, bool)`

GetSubscribedEventsOk returns a tuple with the SubscribedEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedEvents

`func (o *WebhookUpdate) SetSubscribedEvents(v []string)`

SetSubscribedEvents sets SubscribedEvents field to given value.

### HasSubscribedEvents

`func (o *WebhookUpdate) HasSubscribedEvents() bool`

HasSubscribedEvents returns a boolean if a field has been set.

### GetOwnerEmail

`func (o *WebhookUpdate) GetOwnerEmail() string`

GetOwnerEmail returns the OwnerEmail field if non-nil, zero value otherwise.

### GetOwnerEmailOk

`func (o *WebhookUpdate) GetOwnerEmailOk() (*string, bool)`

GetOwnerEmailOk returns a tuple with the OwnerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerEmail

`func (o *WebhookUpdate) SetOwnerEmail(v string)`

SetOwnerEmail sets OwnerEmail field to given value.

### HasOwnerEmail

`func (o *WebhookUpdate) HasOwnerEmail() bool`

HasOwnerEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


