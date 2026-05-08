# CreateStripeCheckoutSessionCustomerUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**CreateStripeCheckoutSessionCustomerUpdateBehavior**](CreateStripeCheckoutSessionCustomerUpdateBehavior.md) | Describes whether Checkout saves the billing address onto customer.address. To always collect a full billing address, use billing_address_collection. Defaults to never. | [optional] 
**Name** | Pointer to [**CreateStripeCheckoutSessionCustomerUpdateBehavior**](CreateStripeCheckoutSessionCustomerUpdateBehavior.md) | Describes whether Checkout saves the name onto customer.name. Defaults to never. | [optional] 
**Shipping** | Pointer to [**CreateStripeCheckoutSessionCustomerUpdateBehavior**](CreateStripeCheckoutSessionCustomerUpdateBehavior.md) | Describes whether Checkout saves shipping information onto customer.shipping. To collect shipping information, use shipping_address_collection. Defaults to never. | [optional] 

## Methods

### NewCreateStripeCheckoutSessionCustomerUpdate

`func NewCreateStripeCheckoutSessionCustomerUpdate() *CreateStripeCheckoutSessionCustomerUpdate`

NewCreateStripeCheckoutSessionCustomerUpdate instantiates a new CreateStripeCheckoutSessionCustomerUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStripeCheckoutSessionCustomerUpdateWithDefaults

`func NewCreateStripeCheckoutSessionCustomerUpdateWithDefaults() *CreateStripeCheckoutSessionCustomerUpdate`

NewCreateStripeCheckoutSessionCustomerUpdateWithDefaults instantiates a new CreateStripeCheckoutSessionCustomerUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *CreateStripeCheckoutSessionCustomerUpdate) GetAddress() CreateStripeCheckoutSessionCustomerUpdateBehavior`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CreateStripeCheckoutSessionCustomerUpdate) GetAddressOk() (*CreateStripeCheckoutSessionCustomerUpdateBehavior, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CreateStripeCheckoutSessionCustomerUpdate) SetAddress(v CreateStripeCheckoutSessionCustomerUpdateBehavior)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CreateStripeCheckoutSessionCustomerUpdate) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetName

`func (o *CreateStripeCheckoutSessionCustomerUpdate) GetName() CreateStripeCheckoutSessionCustomerUpdateBehavior`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateStripeCheckoutSessionCustomerUpdate) GetNameOk() (*CreateStripeCheckoutSessionCustomerUpdateBehavior, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateStripeCheckoutSessionCustomerUpdate) SetName(v CreateStripeCheckoutSessionCustomerUpdateBehavior)`

SetName sets Name field to given value.

### HasName

`func (o *CreateStripeCheckoutSessionCustomerUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShipping

`func (o *CreateStripeCheckoutSessionCustomerUpdate) GetShipping() CreateStripeCheckoutSessionCustomerUpdateBehavior`

GetShipping returns the Shipping field if non-nil, zero value otherwise.

### GetShippingOk

`func (o *CreateStripeCheckoutSessionCustomerUpdate) GetShippingOk() (*CreateStripeCheckoutSessionCustomerUpdateBehavior, bool)`

GetShippingOk returns a tuple with the Shipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipping

`func (o *CreateStripeCheckoutSessionCustomerUpdate) SetShipping(v CreateStripeCheckoutSessionCustomerUpdateBehavior)`

SetShipping sets Shipping field to given value.

### HasShipping

`func (o *CreateStripeCheckoutSessionCustomerUpdate) HasShipping() bool`

HasShipping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


