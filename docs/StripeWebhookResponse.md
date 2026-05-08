# StripeWebhookResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NamespaceId** | **string** | ULID (Universally Unique Lexicographically Sortable Identifier). | 
**AppId** | **string** | ULID (Universally Unique Lexicographically Sortable Identifier). | 
**CustomerId** | Pointer to **string** | ULID (Universally Unique Lexicographically Sortable Identifier). | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewStripeWebhookResponse

`func NewStripeWebhookResponse(namespaceId string, appId string, ) *StripeWebhookResponse`

NewStripeWebhookResponse instantiates a new StripeWebhookResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStripeWebhookResponseWithDefaults

`func NewStripeWebhookResponseWithDefaults() *StripeWebhookResponse`

NewStripeWebhookResponseWithDefaults instantiates a new StripeWebhookResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespaceId

`func (o *StripeWebhookResponse) GetNamespaceId() string`

GetNamespaceId returns the NamespaceId field if non-nil, zero value otherwise.

### GetNamespaceIdOk

`func (o *StripeWebhookResponse) GetNamespaceIdOk() (*string, bool)`

GetNamespaceIdOk returns a tuple with the NamespaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceId

`func (o *StripeWebhookResponse) SetNamespaceId(v string)`

SetNamespaceId sets NamespaceId field to given value.


### GetAppId

`func (o *StripeWebhookResponse) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *StripeWebhookResponse) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *StripeWebhookResponse) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetCustomerId

`func (o *StripeWebhookResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *StripeWebhookResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *StripeWebhookResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *StripeWebhookResponse) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetMessage

`func (o *StripeWebhookResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *StripeWebhookResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *StripeWebhookResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *StripeWebhookResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


