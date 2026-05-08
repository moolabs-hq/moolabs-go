# StripeCustomerAppData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The app ID. If not provided, it will use the global default for the app type. | [optional] 
**Type** | **string** | The app name. | 
**StripeCustomerId** | **string** | The Stripe customer ID. | 
**StripeDefaultPaymentMethodId** | Pointer to **string** | The Stripe default payment method ID. | [optional] 
**App** | Pointer to [**StripeApp**](StripeApp.md) | The installed stripe app this data belongs to. | [optional] [readonly] 

## Methods

### NewStripeCustomerAppData

`func NewStripeCustomerAppData(type_ string, stripeCustomerId string, ) *StripeCustomerAppData`

NewStripeCustomerAppData instantiates a new StripeCustomerAppData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStripeCustomerAppDataWithDefaults

`func NewStripeCustomerAppDataWithDefaults() *StripeCustomerAppData`

NewStripeCustomerAppDataWithDefaults instantiates a new StripeCustomerAppData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StripeCustomerAppData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StripeCustomerAppData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StripeCustomerAppData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StripeCustomerAppData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *StripeCustomerAppData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StripeCustomerAppData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StripeCustomerAppData) SetType(v string)`

SetType sets Type field to given value.


### GetStripeCustomerId

`func (o *StripeCustomerAppData) GetStripeCustomerId() string`

GetStripeCustomerId returns the StripeCustomerId field if non-nil, zero value otherwise.

### GetStripeCustomerIdOk

`func (o *StripeCustomerAppData) GetStripeCustomerIdOk() (*string, bool)`

GetStripeCustomerIdOk returns a tuple with the StripeCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeCustomerId

`func (o *StripeCustomerAppData) SetStripeCustomerId(v string)`

SetStripeCustomerId sets StripeCustomerId field to given value.


### GetStripeDefaultPaymentMethodId

`func (o *StripeCustomerAppData) GetStripeDefaultPaymentMethodId() string`

GetStripeDefaultPaymentMethodId returns the StripeDefaultPaymentMethodId field if non-nil, zero value otherwise.

### GetStripeDefaultPaymentMethodIdOk

`func (o *StripeCustomerAppData) GetStripeDefaultPaymentMethodIdOk() (*string, bool)`

GetStripeDefaultPaymentMethodIdOk returns a tuple with the StripeDefaultPaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeDefaultPaymentMethodId

`func (o *StripeCustomerAppData) SetStripeDefaultPaymentMethodId(v string)`

SetStripeDefaultPaymentMethodId sets StripeDefaultPaymentMethodId field to given value.

### HasStripeDefaultPaymentMethodId

`func (o *StripeCustomerAppData) HasStripeDefaultPaymentMethodId() bool`

HasStripeDefaultPaymentMethodId returns a boolean if a field has been set.

### GetApp

`func (o *StripeCustomerAppData) GetApp() StripeApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *StripeCustomerAppData) GetAppOk() (*StripeApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *StripeCustomerAppData) SetApp(v StripeApp)`

SetApp sets App field to given value.

### HasApp

`func (o *StripeCustomerAppData) HasApp() bool`

HasApp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


