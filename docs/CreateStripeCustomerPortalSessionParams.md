# CreateStripeCustomerPortalSessionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationId** | Pointer to **string** | The ID of an existing configuration to use for this session, describing its functionality and features. If not specified, the session uses the default configuration.  See https://docs.stripe.com/api/customer_portal/sessions/create#create_portal_session-configuration | [optional] 
**Locale** | Pointer to **string** | The IETF language tag of the locale customer portal is displayed in. If blank or auto, the customer’s preferred_locales or browser’s locale is used.  See: https://docs.stripe.com/api/customer_portal/sessions/create#create_portal_session-locale | [optional] 
**ReturnUrl** | Pointer to **string** | The URL to redirect the customer to after they have completed their requested actions.  See: https://docs.stripe.com/api/customer_portal/sessions/create#create_portal_session-return_url | [optional] 

## Methods

### NewCreateStripeCustomerPortalSessionParams

`func NewCreateStripeCustomerPortalSessionParams() *CreateStripeCustomerPortalSessionParams`

NewCreateStripeCustomerPortalSessionParams instantiates a new CreateStripeCustomerPortalSessionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStripeCustomerPortalSessionParamsWithDefaults

`func NewCreateStripeCustomerPortalSessionParamsWithDefaults() *CreateStripeCustomerPortalSessionParams`

NewCreateStripeCustomerPortalSessionParamsWithDefaults instantiates a new CreateStripeCustomerPortalSessionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurationId

`func (o *CreateStripeCustomerPortalSessionParams) GetConfigurationId() string`

GetConfigurationId returns the ConfigurationId field if non-nil, zero value otherwise.

### GetConfigurationIdOk

`func (o *CreateStripeCustomerPortalSessionParams) GetConfigurationIdOk() (*string, bool)`

GetConfigurationIdOk returns a tuple with the ConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationId

`func (o *CreateStripeCustomerPortalSessionParams) SetConfigurationId(v string)`

SetConfigurationId sets ConfigurationId field to given value.

### HasConfigurationId

`func (o *CreateStripeCustomerPortalSessionParams) HasConfigurationId() bool`

HasConfigurationId returns a boolean if a field has been set.

### GetLocale

`func (o *CreateStripeCustomerPortalSessionParams) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *CreateStripeCustomerPortalSessionParams) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *CreateStripeCustomerPortalSessionParams) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *CreateStripeCustomerPortalSessionParams) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetReturnUrl

`func (o *CreateStripeCustomerPortalSessionParams) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *CreateStripeCustomerPortalSessionParams) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *CreateStripeCustomerPortalSessionParams) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *CreateStripeCustomerPortalSessionParams) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


