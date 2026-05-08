# CreateStripeCheckoutSessionRequestCustomer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ULID (Universally Unique Lexicographically Sortable Identifier). | 
**Key** | **string** | An optional unique key of the customer. Useful to reference the customer in external systems. For example, your database ID. | 
**Name** | **string** | Human-readable name for the resource. Between 1 and 256 characters. | 
**Description** | Pointer to **string** | Optional description of the resource. Maximum 1024 characters. | [optional] 
**Metadata** | Pointer to **map[string]string** | Additional metadata for the resource. | [optional] 
**UsageAttribution** | [**CustomerUsageAttribution**](CustomerUsageAttribution.md) | Mapping to attribute metered usage to the customer | 
**PrimaryEmail** | Pointer to **string** | The primary email address of the customer. | [optional] 
**Currency** | Pointer to **string** | Currency of the customer. Used for billing, tax and invoicing. | [optional] 
**BillingAddress** | Pointer to [**Address**](Address.md) | The billing address of the customer. Used for tax and invoicing. | [optional] 

## Methods

### NewCreateStripeCheckoutSessionRequestCustomer

`func NewCreateStripeCheckoutSessionRequestCustomer(id string, key string, name string, usageAttribution CustomerUsageAttribution, ) *CreateStripeCheckoutSessionRequestCustomer`

NewCreateStripeCheckoutSessionRequestCustomer instantiates a new CreateStripeCheckoutSessionRequestCustomer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStripeCheckoutSessionRequestCustomerWithDefaults

`func NewCreateStripeCheckoutSessionRequestCustomerWithDefaults() *CreateStripeCheckoutSessionRequestCustomer`

NewCreateStripeCheckoutSessionRequestCustomerWithDefaults instantiates a new CreateStripeCheckoutSessionRequestCustomer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateStripeCheckoutSessionRequestCustomer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateStripeCheckoutSessionRequestCustomer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateStripeCheckoutSessionRequestCustomer) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *CreateStripeCheckoutSessionRequestCustomer) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateStripeCheckoutSessionRequestCustomer) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateStripeCheckoutSessionRequestCustomer) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *CreateStripeCheckoutSessionRequestCustomer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateStripeCheckoutSessionRequestCustomer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateStripeCheckoutSessionRequestCustomer) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateStripeCheckoutSessionRequestCustomer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateStripeCheckoutSessionRequestCustomer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateStripeCheckoutSessionRequestCustomer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateStripeCheckoutSessionRequestCustomer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateStripeCheckoutSessionRequestCustomer) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateStripeCheckoutSessionRequestCustomer) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateStripeCheckoutSessionRequestCustomer) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateStripeCheckoutSessionRequestCustomer) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetUsageAttribution

`func (o *CreateStripeCheckoutSessionRequestCustomer) GetUsageAttribution() CustomerUsageAttribution`

GetUsageAttribution returns the UsageAttribution field if non-nil, zero value otherwise.

### GetUsageAttributionOk

`func (o *CreateStripeCheckoutSessionRequestCustomer) GetUsageAttributionOk() (*CustomerUsageAttribution, bool)`

GetUsageAttributionOk returns a tuple with the UsageAttribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageAttribution

`func (o *CreateStripeCheckoutSessionRequestCustomer) SetUsageAttribution(v CustomerUsageAttribution)`

SetUsageAttribution sets UsageAttribution field to given value.


### GetPrimaryEmail

`func (o *CreateStripeCheckoutSessionRequestCustomer) GetPrimaryEmail() string`

GetPrimaryEmail returns the PrimaryEmail field if non-nil, zero value otherwise.

### GetPrimaryEmailOk

`func (o *CreateStripeCheckoutSessionRequestCustomer) GetPrimaryEmailOk() (*string, bool)`

GetPrimaryEmailOk returns a tuple with the PrimaryEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryEmail

`func (o *CreateStripeCheckoutSessionRequestCustomer) SetPrimaryEmail(v string)`

SetPrimaryEmail sets PrimaryEmail field to given value.

### HasPrimaryEmail

`func (o *CreateStripeCheckoutSessionRequestCustomer) HasPrimaryEmail() bool`

HasPrimaryEmail returns a boolean if a field has been set.

### GetCurrency

`func (o *CreateStripeCheckoutSessionRequestCustomer) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateStripeCheckoutSessionRequestCustomer) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateStripeCheckoutSessionRequestCustomer) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreateStripeCheckoutSessionRequestCustomer) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetBillingAddress

`func (o *CreateStripeCheckoutSessionRequestCustomer) GetBillingAddress() Address`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *CreateStripeCheckoutSessionRequestCustomer) GetBillingAddressOk() (*Address, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *CreateStripeCheckoutSessionRequestCustomer) SetBillingAddress(v Address)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *CreateStripeCheckoutSessionRequestCustomer) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


