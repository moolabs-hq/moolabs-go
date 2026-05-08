# CreateStripeCheckoutSessionConsentCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentMethodReuseAgreement** | Pointer to [**CreateStripeCheckoutSessionConsentCollectionPaymentMethodReuseAgreement**](CreateStripeCheckoutSessionConsentCollectionPaymentMethodReuseAgreement.md) | Determines the position and visibility of the payment method reuse agreement in the UI. When set to auto, Stripe’s defaults will be used. When set to hidden, the payment method reuse agreement text will always be hidden in the UI. | [optional] 
**Promotions** | Pointer to [**CreateStripeCheckoutSessionConsentCollectionPromotions**](CreateStripeCheckoutSessionConsentCollectionPromotions.md) | If set to auto, enables the collection of customer consent for promotional communications. The Checkout Session will determine whether to display an option to opt into promotional communication from the merchant depending on the customer’s locale. Only available to US merchants. | [optional] 
**TermsOfService** | Pointer to [**CreateStripeCheckoutSessionConsentCollectionTermsOfService**](CreateStripeCheckoutSessionConsentCollectionTermsOfService.md) | If set to required, it requires customers to check a terms of service checkbox before being able to pay. There must be a valid terms of service URL set in your Stripe Dashboard settings. https://dashboard.stripe.com/settings/public | [optional] 

## Methods

### NewCreateStripeCheckoutSessionConsentCollection

`func NewCreateStripeCheckoutSessionConsentCollection() *CreateStripeCheckoutSessionConsentCollection`

NewCreateStripeCheckoutSessionConsentCollection instantiates a new CreateStripeCheckoutSessionConsentCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStripeCheckoutSessionConsentCollectionWithDefaults

`func NewCreateStripeCheckoutSessionConsentCollectionWithDefaults() *CreateStripeCheckoutSessionConsentCollection`

NewCreateStripeCheckoutSessionConsentCollectionWithDefaults instantiates a new CreateStripeCheckoutSessionConsentCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentMethodReuseAgreement

`func (o *CreateStripeCheckoutSessionConsentCollection) GetPaymentMethodReuseAgreement() CreateStripeCheckoutSessionConsentCollectionPaymentMethodReuseAgreement`

GetPaymentMethodReuseAgreement returns the PaymentMethodReuseAgreement field if non-nil, zero value otherwise.

### GetPaymentMethodReuseAgreementOk

`func (o *CreateStripeCheckoutSessionConsentCollection) GetPaymentMethodReuseAgreementOk() (*CreateStripeCheckoutSessionConsentCollectionPaymentMethodReuseAgreement, bool)`

GetPaymentMethodReuseAgreementOk returns a tuple with the PaymentMethodReuseAgreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodReuseAgreement

`func (o *CreateStripeCheckoutSessionConsentCollection) SetPaymentMethodReuseAgreement(v CreateStripeCheckoutSessionConsentCollectionPaymentMethodReuseAgreement)`

SetPaymentMethodReuseAgreement sets PaymentMethodReuseAgreement field to given value.

### HasPaymentMethodReuseAgreement

`func (o *CreateStripeCheckoutSessionConsentCollection) HasPaymentMethodReuseAgreement() bool`

HasPaymentMethodReuseAgreement returns a boolean if a field has been set.

### GetPromotions

`func (o *CreateStripeCheckoutSessionConsentCollection) GetPromotions() CreateStripeCheckoutSessionConsentCollectionPromotions`

GetPromotions returns the Promotions field if non-nil, zero value otherwise.

### GetPromotionsOk

`func (o *CreateStripeCheckoutSessionConsentCollection) GetPromotionsOk() (*CreateStripeCheckoutSessionConsentCollectionPromotions, bool)`

GetPromotionsOk returns a tuple with the Promotions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotions

`func (o *CreateStripeCheckoutSessionConsentCollection) SetPromotions(v CreateStripeCheckoutSessionConsentCollectionPromotions)`

SetPromotions sets Promotions field to given value.

### HasPromotions

`func (o *CreateStripeCheckoutSessionConsentCollection) HasPromotions() bool`

HasPromotions returns a boolean if a field has been set.

### GetTermsOfService

`func (o *CreateStripeCheckoutSessionConsentCollection) GetTermsOfService() CreateStripeCheckoutSessionConsentCollectionTermsOfService`

GetTermsOfService returns the TermsOfService field if non-nil, zero value otherwise.

### GetTermsOfServiceOk

`func (o *CreateStripeCheckoutSessionConsentCollection) GetTermsOfServiceOk() (*CreateStripeCheckoutSessionConsentCollectionTermsOfService, bool)`

GetTermsOfServiceOk returns a tuple with the TermsOfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfService

`func (o *CreateStripeCheckoutSessionConsentCollection) SetTermsOfService(v CreateStripeCheckoutSessionConsentCollectionTermsOfService)`

SetTermsOfService sets TermsOfService field to given value.

### HasTermsOfService

`func (o *CreateStripeCheckoutSessionConsentCollection) HasTermsOfService() bool`

HasTermsOfService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


