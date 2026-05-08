# StripeCustomerAppDataCreateOrUpdateItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The app ID. If not provided, it will use the global default for the app type. | [optional] 
**Type** | **string** | The app name. | 
**StripeCustomerId** | **string** | The Stripe customer ID. | 
**StripeDefaultPaymentMethodId** | Pointer to **string** | The Stripe default payment method ID. | [optional] 

## Methods

### NewStripeCustomerAppDataCreateOrUpdateItem

`func NewStripeCustomerAppDataCreateOrUpdateItem(type_ string, stripeCustomerId string, ) *StripeCustomerAppDataCreateOrUpdateItem`

NewStripeCustomerAppDataCreateOrUpdateItem instantiates a new StripeCustomerAppDataCreateOrUpdateItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStripeCustomerAppDataCreateOrUpdateItemWithDefaults

`func NewStripeCustomerAppDataCreateOrUpdateItemWithDefaults() *StripeCustomerAppDataCreateOrUpdateItem`

NewStripeCustomerAppDataCreateOrUpdateItemWithDefaults instantiates a new StripeCustomerAppDataCreateOrUpdateItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StripeCustomerAppDataCreateOrUpdateItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StripeCustomerAppDataCreateOrUpdateItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StripeCustomerAppDataCreateOrUpdateItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StripeCustomerAppDataCreateOrUpdateItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *StripeCustomerAppDataCreateOrUpdateItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StripeCustomerAppDataCreateOrUpdateItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StripeCustomerAppDataCreateOrUpdateItem) SetType(v string)`

SetType sets Type field to given value.


### GetStripeCustomerId

`func (o *StripeCustomerAppDataCreateOrUpdateItem) GetStripeCustomerId() string`

GetStripeCustomerId returns the StripeCustomerId field if non-nil, zero value otherwise.

### GetStripeCustomerIdOk

`func (o *StripeCustomerAppDataCreateOrUpdateItem) GetStripeCustomerIdOk() (*string, bool)`

GetStripeCustomerIdOk returns a tuple with the StripeCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeCustomerId

`func (o *StripeCustomerAppDataCreateOrUpdateItem) SetStripeCustomerId(v string)`

SetStripeCustomerId sets StripeCustomerId field to given value.


### GetStripeDefaultPaymentMethodId

`func (o *StripeCustomerAppDataCreateOrUpdateItem) GetStripeDefaultPaymentMethodId() string`

GetStripeDefaultPaymentMethodId returns the StripeDefaultPaymentMethodId field if non-nil, zero value otherwise.

### GetStripeDefaultPaymentMethodIdOk

`func (o *StripeCustomerAppDataCreateOrUpdateItem) GetStripeDefaultPaymentMethodIdOk() (*string, bool)`

GetStripeDefaultPaymentMethodIdOk returns a tuple with the StripeDefaultPaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeDefaultPaymentMethodId

`func (o *StripeCustomerAppDataCreateOrUpdateItem) SetStripeDefaultPaymentMethodId(v string)`

SetStripeDefaultPaymentMethodId sets StripeDefaultPaymentMethodId field to given value.

### HasStripeDefaultPaymentMethodId

`func (o *StripeCustomerAppDataCreateOrUpdateItem) HasStripeDefaultPaymentMethodId() bool`

HasStripeDefaultPaymentMethodId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


