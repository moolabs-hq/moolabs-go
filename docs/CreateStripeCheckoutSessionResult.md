# CreateStripeCheckoutSessionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** | The OpenMeter customer ID. | 
**StripeCustomerId** | **string** | The Stripe customer ID. | 
**SessionId** | **string** | The checkout session ID. | 
**SetupIntentId** | **string** | The checkout session setup intent ID. | 
**ClientSecret** | Pointer to **string** | The client secret of the checkout session. This can be used to initialize Stripe.js for your client-side implementation. | [optional] 
**ClientReferenceId** | Pointer to **string** | A unique string to reference the Checkout Session. This can be a customer ID, a cart ID, or similar, and can be used to reconcile the session with your internal systems. | [optional] 
**CustomerEmail** | Pointer to **string** | Customer&#39;s email address provided to Stripe. | [optional] 
**Currency** | Pointer to **string** | Three-letter ISO currency code, in lowercase. | [optional] 
**CreatedAt** | **time.Time** | Timestamp at which the checkout session was created. | 
**ExpiresAt** | Pointer to **time.Time** | Timestamp at which the checkout session will expire. | [optional] 
**Metadata** | Pointer to **map[string]string** | Set of key-value pairs attached to the checkout session. | [optional] 
**Status** | Pointer to **string** | The status of the checkout session. | [optional] 
**Url** | Pointer to **string** | URL to show the checkout session. | [optional] 
**Mode** | [**StripeCheckoutSessionMode**](StripeCheckoutSessionMode.md) | Mode Always &#x60;setup&#x60; for now. | 
**CancelURL** | Pointer to **string** | Cancel URL. | [optional] 
**SuccessURL** | Pointer to **string** | Success URL. | [optional] 
**ReturnURL** | Pointer to **string** | Return URL. | [optional] 

## Methods

### NewCreateStripeCheckoutSessionResult

`func NewCreateStripeCheckoutSessionResult(customerId string, stripeCustomerId string, sessionId string, setupIntentId string, createdAt time.Time, mode StripeCheckoutSessionMode, ) *CreateStripeCheckoutSessionResult`

NewCreateStripeCheckoutSessionResult instantiates a new CreateStripeCheckoutSessionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStripeCheckoutSessionResultWithDefaults

`func NewCreateStripeCheckoutSessionResultWithDefaults() *CreateStripeCheckoutSessionResult`

NewCreateStripeCheckoutSessionResultWithDefaults instantiates a new CreateStripeCheckoutSessionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *CreateStripeCheckoutSessionResult) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CreateStripeCheckoutSessionResult) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CreateStripeCheckoutSessionResult) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetStripeCustomerId

`func (o *CreateStripeCheckoutSessionResult) GetStripeCustomerId() string`

GetStripeCustomerId returns the StripeCustomerId field if non-nil, zero value otherwise.

### GetStripeCustomerIdOk

`func (o *CreateStripeCheckoutSessionResult) GetStripeCustomerIdOk() (*string, bool)`

GetStripeCustomerIdOk returns a tuple with the StripeCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeCustomerId

`func (o *CreateStripeCheckoutSessionResult) SetStripeCustomerId(v string)`

SetStripeCustomerId sets StripeCustomerId field to given value.


### GetSessionId

`func (o *CreateStripeCheckoutSessionResult) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *CreateStripeCheckoutSessionResult) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *CreateStripeCheckoutSessionResult) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetSetupIntentId

`func (o *CreateStripeCheckoutSessionResult) GetSetupIntentId() string`

GetSetupIntentId returns the SetupIntentId field if non-nil, zero value otherwise.

### GetSetupIntentIdOk

`func (o *CreateStripeCheckoutSessionResult) GetSetupIntentIdOk() (*string, bool)`

GetSetupIntentIdOk returns a tuple with the SetupIntentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupIntentId

`func (o *CreateStripeCheckoutSessionResult) SetSetupIntentId(v string)`

SetSetupIntentId sets SetupIntentId field to given value.


### GetClientSecret

`func (o *CreateStripeCheckoutSessionResult) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *CreateStripeCheckoutSessionResult) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *CreateStripeCheckoutSessionResult) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *CreateStripeCheckoutSessionResult) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetClientReferenceId

`func (o *CreateStripeCheckoutSessionResult) GetClientReferenceId() string`

GetClientReferenceId returns the ClientReferenceId field if non-nil, zero value otherwise.

### GetClientReferenceIdOk

`func (o *CreateStripeCheckoutSessionResult) GetClientReferenceIdOk() (*string, bool)`

GetClientReferenceIdOk returns a tuple with the ClientReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientReferenceId

`func (o *CreateStripeCheckoutSessionResult) SetClientReferenceId(v string)`

SetClientReferenceId sets ClientReferenceId field to given value.

### HasClientReferenceId

`func (o *CreateStripeCheckoutSessionResult) HasClientReferenceId() bool`

HasClientReferenceId returns a boolean if a field has been set.

### GetCustomerEmail

`func (o *CreateStripeCheckoutSessionResult) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *CreateStripeCheckoutSessionResult) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *CreateStripeCheckoutSessionResult) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *CreateStripeCheckoutSessionResult) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### GetCurrency

`func (o *CreateStripeCheckoutSessionResult) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateStripeCheckoutSessionResult) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateStripeCheckoutSessionResult) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreateStripeCheckoutSessionResult) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CreateStripeCheckoutSessionResult) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateStripeCheckoutSessionResult) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateStripeCheckoutSessionResult) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetExpiresAt

`func (o *CreateStripeCheckoutSessionResult) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CreateStripeCheckoutSessionResult) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CreateStripeCheckoutSessionResult) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CreateStripeCheckoutSessionResult) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateStripeCheckoutSessionResult) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateStripeCheckoutSessionResult) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateStripeCheckoutSessionResult) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateStripeCheckoutSessionResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetStatus

`func (o *CreateStripeCheckoutSessionResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateStripeCheckoutSessionResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateStripeCheckoutSessionResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateStripeCheckoutSessionResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUrl

`func (o *CreateStripeCheckoutSessionResult) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateStripeCheckoutSessionResult) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateStripeCheckoutSessionResult) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CreateStripeCheckoutSessionResult) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetMode

`func (o *CreateStripeCheckoutSessionResult) GetMode() StripeCheckoutSessionMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CreateStripeCheckoutSessionResult) GetModeOk() (*StripeCheckoutSessionMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CreateStripeCheckoutSessionResult) SetMode(v StripeCheckoutSessionMode)`

SetMode sets Mode field to given value.


### GetCancelURL

`func (o *CreateStripeCheckoutSessionResult) GetCancelURL() string`

GetCancelURL returns the CancelURL field if non-nil, zero value otherwise.

### GetCancelURLOk

`func (o *CreateStripeCheckoutSessionResult) GetCancelURLOk() (*string, bool)`

GetCancelURLOk returns a tuple with the CancelURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelURL

`func (o *CreateStripeCheckoutSessionResult) SetCancelURL(v string)`

SetCancelURL sets CancelURL field to given value.

### HasCancelURL

`func (o *CreateStripeCheckoutSessionResult) HasCancelURL() bool`

HasCancelURL returns a boolean if a field has been set.

### GetSuccessURL

`func (o *CreateStripeCheckoutSessionResult) GetSuccessURL() string`

GetSuccessURL returns the SuccessURL field if non-nil, zero value otherwise.

### GetSuccessURLOk

`func (o *CreateStripeCheckoutSessionResult) GetSuccessURLOk() (*string, bool)`

GetSuccessURLOk returns a tuple with the SuccessURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessURL

`func (o *CreateStripeCheckoutSessionResult) SetSuccessURL(v string)`

SetSuccessURL sets SuccessURL field to given value.

### HasSuccessURL

`func (o *CreateStripeCheckoutSessionResult) HasSuccessURL() bool`

HasSuccessURL returns a boolean if a field has been set.

### GetReturnURL

`func (o *CreateStripeCheckoutSessionResult) GetReturnURL() string`

GetReturnURL returns the ReturnURL field if non-nil, zero value otherwise.

### GetReturnURLOk

`func (o *CreateStripeCheckoutSessionResult) GetReturnURLOk() (*string, bool)`

GetReturnURLOk returns a tuple with the ReturnURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnURL

`func (o *CreateStripeCheckoutSessionResult) SetReturnURL(v string)`

SetReturnURL sets ReturnURL field to given value.

### HasReturnURL

`func (o *CreateStripeCheckoutSessionResult) HasReturnURL() bool`

HasReturnURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


