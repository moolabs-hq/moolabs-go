# CreateStripeCheckoutSessionRequestOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingAddressCollection** | Pointer to [**CreateStripeCheckoutSessionBillingAddressCollection**](CreateStripeCheckoutSessionBillingAddressCollection.md) | Specify whether Checkout should collect the customer’s billing address. Defaults to auto. | [optional] 
**CancelURL** | Pointer to **string** | If set, Checkout displays a back button and customers will be directed to this URL if they decide to cancel payment and return to your website. This parameter is not allowed if ui_mode is embedded. | [optional] 
**ClientReferenceID** | Pointer to **string** | A unique string to reference the Checkout Session. This can be a customer ID, a cart ID, or similar, and can be used to reconcile the session with your internal systems. | [optional] 
**CustomerUpdate** | Pointer to [**CreateStripeCheckoutSessionCustomerUpdate**](CreateStripeCheckoutSessionCustomerUpdate.md) | Controls what fields on Customer can be updated by the Checkout Session. | [optional] 
**ConsentCollection** | Pointer to [**CreateStripeCheckoutSessionConsentCollection**](CreateStripeCheckoutSessionConsentCollection.md) | Configure fields for the Checkout Session to gather active consent from customers. | [optional] 
**Currency** | Pointer to **string** | Three-letter ISO currency code, in lowercase. | [optional] 
**CustomText** | Pointer to [**CheckoutSessionCustomTextAfterSubmitParams**](CheckoutSessionCustomTextAfterSubmitParams.md) | Display additional text for your customers using custom text. | [optional] 
**ExpiresAt** | Pointer to **int64** | The Epoch time in seconds at which the Checkout Session will expire. It can be anywhere from 30 minutes to 24 hours after Checkout Session creation. By default, this value is 24 hours from creation. | [optional] 
**Locale** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** | Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to metadata. | [optional] 
**ReturnURL** | Pointer to **string** | The URL to redirect your customer back to after they authenticate or cancel their payment on the payment method’s app or site. This parameter is required if ui_mode is embedded and redirect-based payment methods are enabled on the session. | [optional] 
**SuccessURL** | Pointer to **string** | The URL to which Stripe should send customers when payment or setup is complete. This parameter is not allowed if ui_mode is embedded. If you’d like to use information from the successful Checkout Session on your page, read the guide on customizing your success page: https://docs.stripe.com/payments/checkout/custom-success-page | [optional] 
**UiMode** | Pointer to [**CheckoutSessionUIMode**](CheckoutSessionUIMode.md) | The UI mode of the Session. Defaults to hosted. | [optional] 
**PaymentMethodTypes** | Pointer to **[]string** | A list of the types of payment methods (e.g., card) this Checkout Session can accept. | [optional] 
**RedirectOnCompletion** | Pointer to [**CreateStripeCheckoutSessionRedirectOnCompletion**](CreateStripeCheckoutSessionRedirectOnCompletion.md) | This parameter applies to ui_mode: embedded. Defaults to always. Learn more about the redirect behavior of embedded sessions at https://docs.stripe.com/payments/checkout/custom-success-page?payment-ui&#x3D;embedded-form | [optional] 
**TaxIdCollection** | Pointer to [**CreateCheckoutSessionTaxIdCollection**](CreateCheckoutSessionTaxIdCollection.md) | Controls tax ID collection during checkout. | [optional] 

## Methods

### NewCreateStripeCheckoutSessionRequestOptions

`func NewCreateStripeCheckoutSessionRequestOptions() *CreateStripeCheckoutSessionRequestOptions`

NewCreateStripeCheckoutSessionRequestOptions instantiates a new CreateStripeCheckoutSessionRequestOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStripeCheckoutSessionRequestOptionsWithDefaults

`func NewCreateStripeCheckoutSessionRequestOptionsWithDefaults() *CreateStripeCheckoutSessionRequestOptions`

NewCreateStripeCheckoutSessionRequestOptionsWithDefaults instantiates a new CreateStripeCheckoutSessionRequestOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingAddressCollection

`func (o *CreateStripeCheckoutSessionRequestOptions) GetBillingAddressCollection() CreateStripeCheckoutSessionBillingAddressCollection`

GetBillingAddressCollection returns the BillingAddressCollection field if non-nil, zero value otherwise.

### GetBillingAddressCollectionOk

`func (o *CreateStripeCheckoutSessionRequestOptions) GetBillingAddressCollectionOk() (*CreateStripeCheckoutSessionBillingAddressCollection, bool)`

GetBillingAddressCollectionOk returns a tuple with the BillingAddressCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddressCollection

`func (o *CreateStripeCheckoutSessionRequestOptions) SetBillingAddressCollection(v CreateStripeCheckoutSessionBillingAddressCollection)`

SetBillingAddressCollection sets BillingAddressCollection field to given value.

### HasBillingAddressCollection

`func (o *CreateStripeCheckoutSessionRequestOptions) HasBillingAddressCollection() bool`

HasBillingAddressCollection returns a boolean if a field has been set.

### GetCancelURL

`func (o *CreateStripeCheckoutSessionRequestOptions) GetCancelURL() string`

GetCancelURL returns the CancelURL field if non-nil, zero value otherwise.

### GetCancelURLOk

`func (o *CreateStripeCheckoutSessionRequestOptions) GetCancelURLOk() (*string, bool)`

GetCancelURLOk returns a tuple with the CancelURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelURL

`func (o *CreateStripeCheckoutSessionRequestOptions) SetCancelURL(v string)`

SetCancelURL sets CancelURL field to given value.

### HasCancelURL

`func (o *CreateStripeCheckoutSessionRequestOptions) HasCancelURL() bool`

HasCancelURL returns a boolean if a field has been set.

### GetClientReferenceID

`func (o *CreateStripeCheckoutSessionRequestOptions) GetClientReferenceID() string`

GetClientReferenceID returns the ClientReferenceID field if non-nil, zero value otherwise.

### GetClientReferenceIDOk

`func (o *CreateStripeCheckoutSessionRequestOptions) GetClientReferenceIDOk() (*string, bool)`

GetClientReferenceIDOk returns a tuple with the ClientReferenceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientReferenceID

`func (o *CreateStripeCheckoutSessionRequestOptions) SetClientReferenceID(v string)`

SetClientReferenceID sets ClientReferenceID field to given value.

### HasClientReferenceID

`func (o *CreateStripeCheckoutSessionRequestOptions) HasClientReferenceID() bool`

HasClientReferenceID returns a boolean if a field has been set.

### GetCustomerUpdate

`func (o *CreateStripeCheckoutSessionRequestOptions) GetCustomerUpdate() CreateStripeCheckoutSessionCustomerUpdate`

GetCustomerUpdate returns the CustomerUpdate field if non-nil, zero value otherwise.

### GetCustomerUpdateOk

`func (o *CreateStripeCheckoutSessionRequestOptions) GetCustomerUpdateOk() (*CreateStripeCheckoutSessionCustomerUpdate, bool)`

GetCustomerUpdateOk returns a tuple with the CustomerUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUpdate

`func (o *CreateStripeCheckoutSessionRequestOptions) SetCustomerUpdate(v CreateStripeCheckoutSessionCustomerUpdate)`

SetCustomerUpdate sets CustomerUpdate field to given value.

### HasCustomerUpdate

`func (o *CreateStripeCheckoutSessionRequestOptions) HasCustomerUpdate() bool`

HasCustomerUpdate returns a boolean if a field has been set.

### GetConsentCollection

`func (o *CreateStripeCheckoutSessionRequestOptions) GetConsentCollection() CreateStripeCheckoutSessionConsentCollection`

GetConsentCollection returns the ConsentCollection field if non-nil, zero value otherwise.

### GetConsentCollectionOk

`func (o *CreateStripeCheckoutSessionRequestOptions) GetConsentCollectionOk() (*CreateStripeCheckoutSessionConsentCollection, bool)`

GetConsentCollectionOk returns a tuple with the ConsentCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentCollection

`func (o *CreateStripeCheckoutSessionRequestOptions) SetConsentCollection(v CreateStripeCheckoutSessionConsentCollection)`

SetConsentCollection sets ConsentCollection field to given value.

### HasConsentCollection

`func (o *CreateStripeCheckoutSessionRequestOptions) HasConsentCollection() bool`

HasConsentCollection returns a boolean if a field has been set.

### GetCurrency

`func (o *CreateStripeCheckoutSessionRequestOptions) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateStripeCheckoutSessionRequestOptions) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateStripeCheckoutSessionRequestOptions) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreateStripeCheckoutSessionRequestOptions) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomText

`func (o *CreateStripeCheckoutSessionRequestOptions) GetCustomText() CheckoutSessionCustomTextAfterSubmitParams`

GetCustomText returns the CustomText field if non-nil, zero value otherwise.

### GetCustomTextOk

`func (o *CreateStripeCheckoutSessionRequestOptions) GetCustomTextOk() (*CheckoutSessionCustomTextAfterSubmitParams, bool)`

GetCustomTextOk returns a tuple with the CustomText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText

`func (o *CreateStripeCheckoutSessionRequestOptions) SetCustomText(v CheckoutSessionCustomTextAfterSubmitParams)`

SetCustomText sets CustomText field to given value.

### HasCustomText

`func (o *CreateStripeCheckoutSessionRequestOptions) HasCustomText() bool`

HasCustomText returns a boolean if a field has been set.

### GetExpiresAt

`func (o *CreateStripeCheckoutSessionRequestOptions) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CreateStripeCheckoutSessionRequestOptions) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CreateStripeCheckoutSessionRequestOptions) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CreateStripeCheckoutSessionRequestOptions) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetLocale

`func (o *CreateStripeCheckoutSessionRequestOptions) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *CreateStripeCheckoutSessionRequestOptions) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *CreateStripeCheckoutSessionRequestOptions) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *CreateStripeCheckoutSessionRequestOptions) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateStripeCheckoutSessionRequestOptions) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateStripeCheckoutSessionRequestOptions) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateStripeCheckoutSessionRequestOptions) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateStripeCheckoutSessionRequestOptions) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetReturnURL

`func (o *CreateStripeCheckoutSessionRequestOptions) GetReturnURL() string`

GetReturnURL returns the ReturnURL field if non-nil, zero value otherwise.

### GetReturnURLOk

`func (o *CreateStripeCheckoutSessionRequestOptions) GetReturnURLOk() (*string, bool)`

GetReturnURLOk returns a tuple with the ReturnURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnURL

`func (o *CreateStripeCheckoutSessionRequestOptions) SetReturnURL(v string)`

SetReturnURL sets ReturnURL field to given value.

### HasReturnURL

`func (o *CreateStripeCheckoutSessionRequestOptions) HasReturnURL() bool`

HasReturnURL returns a boolean if a field has been set.

### GetSuccessURL

`func (o *CreateStripeCheckoutSessionRequestOptions) GetSuccessURL() string`

GetSuccessURL returns the SuccessURL field if non-nil, zero value otherwise.

### GetSuccessURLOk

`func (o *CreateStripeCheckoutSessionRequestOptions) GetSuccessURLOk() (*string, bool)`

GetSuccessURLOk returns a tuple with the SuccessURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessURL

`func (o *CreateStripeCheckoutSessionRequestOptions) SetSuccessURL(v string)`

SetSuccessURL sets SuccessURL field to given value.

### HasSuccessURL

`func (o *CreateStripeCheckoutSessionRequestOptions) HasSuccessURL() bool`

HasSuccessURL returns a boolean if a field has been set.

### GetUiMode

`func (o *CreateStripeCheckoutSessionRequestOptions) GetUiMode() CheckoutSessionUIMode`

GetUiMode returns the UiMode field if non-nil, zero value otherwise.

### GetUiModeOk

`func (o *CreateStripeCheckoutSessionRequestOptions) GetUiModeOk() (*CheckoutSessionUIMode, bool)`

GetUiModeOk returns a tuple with the UiMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiMode

`func (o *CreateStripeCheckoutSessionRequestOptions) SetUiMode(v CheckoutSessionUIMode)`

SetUiMode sets UiMode field to given value.

### HasUiMode

`func (o *CreateStripeCheckoutSessionRequestOptions) HasUiMode() bool`

HasUiMode returns a boolean if a field has been set.

### GetPaymentMethodTypes

`func (o *CreateStripeCheckoutSessionRequestOptions) GetPaymentMethodTypes() []string`

GetPaymentMethodTypes returns the PaymentMethodTypes field if non-nil, zero value otherwise.

### GetPaymentMethodTypesOk

`func (o *CreateStripeCheckoutSessionRequestOptions) GetPaymentMethodTypesOk() (*[]string, bool)`

GetPaymentMethodTypesOk returns a tuple with the PaymentMethodTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodTypes

`func (o *CreateStripeCheckoutSessionRequestOptions) SetPaymentMethodTypes(v []string)`

SetPaymentMethodTypes sets PaymentMethodTypes field to given value.

### HasPaymentMethodTypes

`func (o *CreateStripeCheckoutSessionRequestOptions) HasPaymentMethodTypes() bool`

HasPaymentMethodTypes returns a boolean if a field has been set.

### GetRedirectOnCompletion

`func (o *CreateStripeCheckoutSessionRequestOptions) GetRedirectOnCompletion() CreateStripeCheckoutSessionRedirectOnCompletion`

GetRedirectOnCompletion returns the RedirectOnCompletion field if non-nil, zero value otherwise.

### GetRedirectOnCompletionOk

`func (o *CreateStripeCheckoutSessionRequestOptions) GetRedirectOnCompletionOk() (*CreateStripeCheckoutSessionRedirectOnCompletion, bool)`

GetRedirectOnCompletionOk returns a tuple with the RedirectOnCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectOnCompletion

`func (o *CreateStripeCheckoutSessionRequestOptions) SetRedirectOnCompletion(v CreateStripeCheckoutSessionRedirectOnCompletion)`

SetRedirectOnCompletion sets RedirectOnCompletion field to given value.

### HasRedirectOnCompletion

`func (o *CreateStripeCheckoutSessionRequestOptions) HasRedirectOnCompletion() bool`

HasRedirectOnCompletion returns a boolean if a field has been set.

### GetTaxIdCollection

`func (o *CreateStripeCheckoutSessionRequestOptions) GetTaxIdCollection() CreateCheckoutSessionTaxIdCollection`

GetTaxIdCollection returns the TaxIdCollection field if non-nil, zero value otherwise.

### GetTaxIdCollectionOk

`func (o *CreateStripeCheckoutSessionRequestOptions) GetTaxIdCollectionOk() (*CreateCheckoutSessionTaxIdCollection, bool)`

GetTaxIdCollectionOk returns a tuple with the TaxIdCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIdCollection

`func (o *CreateStripeCheckoutSessionRequestOptions) SetTaxIdCollection(v CreateCheckoutSessionTaxIdCollection)`

SetTaxIdCollection sets TaxIdCollection field to given value.

### HasTaxIdCollection

`func (o *CreateStripeCheckoutSessionRequestOptions) HasTaxIdCollection() bool`

HasTaxIdCollection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


