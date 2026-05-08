# StripeCustomerPortalSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the customer portal session.  See: https://docs.stripe.com/api/customer_portal/sessions/object#portal_session_object-id | 
**StripeCustomerId** | **string** | The ID of the stripe customer. | 
**ConfigurationId** | **string** | Configuration used to customize the customer portal.  See: https://docs.stripe.com/api/customer_portal/sessions/object#portal_session_object-configuration | 
**Livemode** | **bool** | Livemode.  See: https://docs.stripe.com/api/customer_portal/sessions/object#portal_session_object-livemode | 
**CreatedAt** | **time.Time** | Created at.  See: https://docs.stripe.com/api/customer_portal/sessions/object#portal_session_object-created | 
**ReturnUrl** | **string** | Return URL.  See: https://docs.stripe.com/api/customer_portal/sessions/object#portal_session_object-return_url | 
**Locale** | **string** | Status.   /_** The IETF language tag of the locale customer portal is displayed in.  See: https://docs.stripe.com/api/customer_portal/sessions/object#portal_session_object-locale | 
**Url** | **string** | /_** The ID of the customer.The URL to redirect the customer to after they have completed their requested actions. | 

## Methods

### NewStripeCustomerPortalSession

`func NewStripeCustomerPortalSession(id string, stripeCustomerId string, configurationId string, livemode bool, createdAt time.Time, returnUrl string, locale string, url string, ) *StripeCustomerPortalSession`

NewStripeCustomerPortalSession instantiates a new StripeCustomerPortalSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStripeCustomerPortalSessionWithDefaults

`func NewStripeCustomerPortalSessionWithDefaults() *StripeCustomerPortalSession`

NewStripeCustomerPortalSessionWithDefaults instantiates a new StripeCustomerPortalSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StripeCustomerPortalSession) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StripeCustomerPortalSession) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StripeCustomerPortalSession) SetId(v string)`

SetId sets Id field to given value.


### GetStripeCustomerId

`func (o *StripeCustomerPortalSession) GetStripeCustomerId() string`

GetStripeCustomerId returns the StripeCustomerId field if non-nil, zero value otherwise.

### GetStripeCustomerIdOk

`func (o *StripeCustomerPortalSession) GetStripeCustomerIdOk() (*string, bool)`

GetStripeCustomerIdOk returns a tuple with the StripeCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeCustomerId

`func (o *StripeCustomerPortalSession) SetStripeCustomerId(v string)`

SetStripeCustomerId sets StripeCustomerId field to given value.


### GetConfigurationId

`func (o *StripeCustomerPortalSession) GetConfigurationId() string`

GetConfigurationId returns the ConfigurationId field if non-nil, zero value otherwise.

### GetConfigurationIdOk

`func (o *StripeCustomerPortalSession) GetConfigurationIdOk() (*string, bool)`

GetConfigurationIdOk returns a tuple with the ConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationId

`func (o *StripeCustomerPortalSession) SetConfigurationId(v string)`

SetConfigurationId sets ConfigurationId field to given value.


### GetLivemode

`func (o *StripeCustomerPortalSession) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *StripeCustomerPortalSession) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *StripeCustomerPortalSession) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.


### GetCreatedAt

`func (o *StripeCustomerPortalSession) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StripeCustomerPortalSession) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StripeCustomerPortalSession) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetReturnUrl

`func (o *StripeCustomerPortalSession) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *StripeCustomerPortalSession) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *StripeCustomerPortalSession) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.


### GetLocale

`func (o *StripeCustomerPortalSession) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *StripeCustomerPortalSession) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *StripeCustomerPortalSession) SetLocale(v string)`

SetLocale sets Locale field to given value.


### GetUrl

`func (o *StripeCustomerPortalSession) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *StripeCustomerPortalSession) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *StripeCustomerPortalSession) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


