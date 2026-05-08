# CustomerAppData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The app ID. If not provided, it will use the global default for the app type. | [optional] 
**Type** | **string** | The app name. | 
**StripeCustomerId** | **string** | The Stripe customer ID. | 
**StripeDefaultPaymentMethodId** | Pointer to **string** | The Stripe default payment method ID. | [optional] 
**App** | Pointer to [**CustomInvoicingApp**](CustomInvoicingApp.md) | The installed custom invoicing app this data belongs to. | [optional] [readonly] 
**Metadata** | Pointer to **map[string]string** | Metadata to be used by the custom invoicing provider. | [optional] 

## Methods

### NewCustomerAppData

`func NewCustomerAppData(type_ string, stripeCustomerId string, ) *CustomerAppData`

NewCustomerAppData instantiates a new CustomerAppData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerAppDataWithDefaults

`func NewCustomerAppDataWithDefaults() *CustomerAppData`

NewCustomerAppDataWithDefaults instantiates a new CustomerAppData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomerAppData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerAppData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerAppData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomerAppData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CustomerAppData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomerAppData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomerAppData) SetType(v string)`

SetType sets Type field to given value.


### GetStripeCustomerId

`func (o *CustomerAppData) GetStripeCustomerId() string`

GetStripeCustomerId returns the StripeCustomerId field if non-nil, zero value otherwise.

### GetStripeCustomerIdOk

`func (o *CustomerAppData) GetStripeCustomerIdOk() (*string, bool)`

GetStripeCustomerIdOk returns a tuple with the StripeCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeCustomerId

`func (o *CustomerAppData) SetStripeCustomerId(v string)`

SetStripeCustomerId sets StripeCustomerId field to given value.


### GetStripeDefaultPaymentMethodId

`func (o *CustomerAppData) GetStripeDefaultPaymentMethodId() string`

GetStripeDefaultPaymentMethodId returns the StripeDefaultPaymentMethodId field if non-nil, zero value otherwise.

### GetStripeDefaultPaymentMethodIdOk

`func (o *CustomerAppData) GetStripeDefaultPaymentMethodIdOk() (*string, bool)`

GetStripeDefaultPaymentMethodIdOk returns a tuple with the StripeDefaultPaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeDefaultPaymentMethodId

`func (o *CustomerAppData) SetStripeDefaultPaymentMethodId(v string)`

SetStripeDefaultPaymentMethodId sets StripeDefaultPaymentMethodId field to given value.

### HasStripeDefaultPaymentMethodId

`func (o *CustomerAppData) HasStripeDefaultPaymentMethodId() bool`

HasStripeDefaultPaymentMethodId returns a boolean if a field has been set.

### GetApp

`func (o *CustomerAppData) GetApp() CustomInvoicingApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *CustomerAppData) GetAppOk() (*CustomInvoicingApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *CustomerAppData) SetApp(v CustomInvoicingApp)`

SetApp sets App field to given value.

### HasApp

`func (o *CustomerAppData) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetMetadata

`func (o *CustomerAppData) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CustomerAppData) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CustomerAppData) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CustomerAppData) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


