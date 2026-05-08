# StripeCustomerAppDataBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StripeCustomerId** | **string** | The Stripe customer ID. | 
**StripeDefaultPaymentMethodId** | Pointer to **string** | The Stripe default payment method ID. | [optional] 

## Methods

### NewStripeCustomerAppDataBase

`func NewStripeCustomerAppDataBase(stripeCustomerId string, ) *StripeCustomerAppDataBase`

NewStripeCustomerAppDataBase instantiates a new StripeCustomerAppDataBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStripeCustomerAppDataBaseWithDefaults

`func NewStripeCustomerAppDataBaseWithDefaults() *StripeCustomerAppDataBase`

NewStripeCustomerAppDataBaseWithDefaults instantiates a new StripeCustomerAppDataBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStripeCustomerId

`func (o *StripeCustomerAppDataBase) GetStripeCustomerId() string`

GetStripeCustomerId returns the StripeCustomerId field if non-nil, zero value otherwise.

### GetStripeCustomerIdOk

`func (o *StripeCustomerAppDataBase) GetStripeCustomerIdOk() (*string, bool)`

GetStripeCustomerIdOk returns a tuple with the StripeCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeCustomerId

`func (o *StripeCustomerAppDataBase) SetStripeCustomerId(v string)`

SetStripeCustomerId sets StripeCustomerId field to given value.


### GetStripeDefaultPaymentMethodId

`func (o *StripeCustomerAppDataBase) GetStripeDefaultPaymentMethodId() string`

GetStripeDefaultPaymentMethodId returns the StripeDefaultPaymentMethodId field if non-nil, zero value otherwise.

### GetStripeDefaultPaymentMethodIdOk

`func (o *StripeCustomerAppDataBase) GetStripeDefaultPaymentMethodIdOk() (*string, bool)`

GetStripeDefaultPaymentMethodIdOk returns a tuple with the StripeDefaultPaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeDefaultPaymentMethodId

`func (o *StripeCustomerAppDataBase) SetStripeDefaultPaymentMethodId(v string)`

SetStripeDefaultPaymentMethodId sets StripeDefaultPaymentMethodId field to given value.

### HasStripeDefaultPaymentMethodId

`func (o *StripeCustomerAppDataBase) HasStripeDefaultPaymentMethodId() bool`

HasStripeDefaultPaymentMethodId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


