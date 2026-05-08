# CreateStripeCheckoutSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **string** | If not provided, the default Stripe app is used if any. | [optional] 
**Customer** | [**CreateStripeCheckoutSessionRequestCustomer**](CreateStripeCheckoutSessionRequestCustomer.md) |  | 
**StripeCustomerId** | Pointer to **string** | Stripe customer ID. If not provided OpenMeter creates a new Stripe customer or uses the OpenMeter customer&#39;s default Stripe customer ID. | [optional] 
**Options** | [**CreateStripeCheckoutSessionRequestOptions**](CreateStripeCheckoutSessionRequestOptions.md) | Options passed to Stripe when creating the checkout session. | 

## Methods

### NewCreateStripeCheckoutSessionRequest

`func NewCreateStripeCheckoutSessionRequest(customer CreateStripeCheckoutSessionRequestCustomer, options CreateStripeCheckoutSessionRequestOptions, ) *CreateStripeCheckoutSessionRequest`

NewCreateStripeCheckoutSessionRequest instantiates a new CreateStripeCheckoutSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStripeCheckoutSessionRequestWithDefaults

`func NewCreateStripeCheckoutSessionRequestWithDefaults() *CreateStripeCheckoutSessionRequest`

NewCreateStripeCheckoutSessionRequestWithDefaults instantiates a new CreateStripeCheckoutSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *CreateStripeCheckoutSessionRequest) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *CreateStripeCheckoutSessionRequest) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *CreateStripeCheckoutSessionRequest) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *CreateStripeCheckoutSessionRequest) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetCustomer

`func (o *CreateStripeCheckoutSessionRequest) GetCustomer() CreateStripeCheckoutSessionRequestCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *CreateStripeCheckoutSessionRequest) GetCustomerOk() (*CreateStripeCheckoutSessionRequestCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *CreateStripeCheckoutSessionRequest) SetCustomer(v CreateStripeCheckoutSessionRequestCustomer)`

SetCustomer sets Customer field to given value.


### GetStripeCustomerId

`func (o *CreateStripeCheckoutSessionRequest) GetStripeCustomerId() string`

GetStripeCustomerId returns the StripeCustomerId field if non-nil, zero value otherwise.

### GetStripeCustomerIdOk

`func (o *CreateStripeCheckoutSessionRequest) GetStripeCustomerIdOk() (*string, bool)`

GetStripeCustomerIdOk returns a tuple with the StripeCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeCustomerId

`func (o *CreateStripeCheckoutSessionRequest) SetStripeCustomerId(v string)`

SetStripeCustomerId sets StripeCustomerId field to given value.

### HasStripeCustomerId

`func (o *CreateStripeCheckoutSessionRequest) HasStripeCustomerId() bool`

HasStripeCustomerId returns a boolean if a field has been set.

### GetOptions

`func (o *CreateStripeCheckoutSessionRequest) GetOptions() CreateStripeCheckoutSessionRequestOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CreateStripeCheckoutSessionRequest) GetOptionsOk() (*CreateStripeCheckoutSessionRequestOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CreateStripeCheckoutSessionRequest) SetOptions(v CreateStripeCheckoutSessionRequestOptions)`

SetOptions sets Options field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


