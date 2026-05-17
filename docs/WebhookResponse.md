# WebhookResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**AuthType** | **string** |  | 
**HasSecret** | **bool** |  | 
**SubscribedEvents** | **[]string** |  | 
**OwnerEmail** | Pointer to **string** |  | [optional] 

## Methods

### NewWebhookResponse

`func NewWebhookResponse(url string, authType string, hasSecret bool, subscribedEvents []string, ) *WebhookResponse`

NewWebhookResponse instantiates a new WebhookResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookResponseWithDefaults

`func NewWebhookResponseWithDefaults() *WebhookResponse`

NewWebhookResponseWithDefaults instantiates a new WebhookResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *WebhookResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookResponse) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetAuthType

`func (o *WebhookResponse) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *WebhookResponse) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *WebhookResponse) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.


### GetHasSecret

`func (o *WebhookResponse) GetHasSecret() bool`

GetHasSecret returns the HasSecret field if non-nil, zero value otherwise.

### GetHasSecretOk

`func (o *WebhookResponse) GetHasSecretOk() (*bool, bool)`

GetHasSecretOk returns a tuple with the HasSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSecret

`func (o *WebhookResponse) SetHasSecret(v bool)`

SetHasSecret sets HasSecret field to given value.


### GetSubscribedEvents

`func (o *WebhookResponse) GetSubscribedEvents() []string`

GetSubscribedEvents returns the SubscribedEvents field if non-nil, zero value otherwise.

### GetSubscribedEventsOk

`func (o *WebhookResponse) GetSubscribedEventsOk() (*[]string, bool)`

GetSubscribedEventsOk returns a tuple with the SubscribedEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedEvents

`func (o *WebhookResponse) SetSubscribedEvents(v []string)`

SetSubscribedEvents sets SubscribedEvents field to given value.


### GetOwnerEmail

`func (o *WebhookResponse) GetOwnerEmail() string`

GetOwnerEmail returns the OwnerEmail field if non-nil, zero value otherwise.

### GetOwnerEmailOk

`func (o *WebhookResponse) GetOwnerEmailOk() (*string, bool)`

GetOwnerEmailOk returns a tuple with the OwnerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerEmail

`func (o *WebhookResponse) SetOwnerEmail(v string)`

SetOwnerEmail sets OwnerEmail field to given value.

### HasOwnerEmail

`func (o *WebhookResponse) HasOwnerEmail() bool`

HasOwnerEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


