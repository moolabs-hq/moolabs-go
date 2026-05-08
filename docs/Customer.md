# Customer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier for the resource. | [readonly] 
**Name** | **string** | Human-readable name for the resource. Between 1 and 256 characters. | 
**Description** | Pointer to **string** | Optional description of the resource. Maximum 1024 characters. | [optional] 
**Metadata** | Pointer to **map[string]string** | Additional metadata for the resource. | [optional] 
**CreatedAt** | **time.Time** | Timestamp of when the resource was created. | [readonly] 
**UpdatedAt** | **time.Time** | Timestamp of when the resource was last updated. | [readonly] 
**DeletedAt** | Pointer to **time.Time** | Timestamp of when the resource was permanently deleted. | [optional] [readonly] 
**Key** | Pointer to **string** | An optional unique key of the customer. Useful to reference the customer in external systems. For example, your database ID. | [optional] 
**UsageAttribution** | [**CustomerUsageAttribution**](CustomerUsageAttribution.md) | Mapping to attribute metered usage to the customer | 
**PrimaryEmail** | Pointer to **string** | The primary email address of the customer. | [optional] 
**Currency** | Pointer to **string** | Currency of the customer. Used for billing, tax and invoicing. | [optional] 
**BillingAddress** | Pointer to [**Address**](Address.md) | The billing address of the customer. Used for tax and invoicing. | [optional] 
**CurrentSubscriptionId** | Pointer to **string** | The ID of the Subscription if the customer has one. | [optional] [readonly] 
**Subscriptions** | Pointer to [**[]Subscription**](Subscription.md) | The subscriptions of the customer. Only with the &#x60;subscriptions&#x60; expand option. | [optional] 
**Annotations** | Pointer to **map[string]interface{}** | Set of key-value pairs managed by the system. Cannot be modified by user. | [optional] [readonly] 

## Methods

### NewCustomer

`func NewCustomer(id string, name string, createdAt time.Time, updatedAt time.Time, usageAttribution CustomerUsageAttribution, ) *Customer`

NewCustomer instantiates a new Customer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerWithDefaults

`func NewCustomerWithDefaults() *Customer`

NewCustomerWithDefaults instantiates a new Customer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Customer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Customer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Customer) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Customer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Customer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Customer) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Customer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Customer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Customer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Customer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *Customer) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Customer) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Customer) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Customer) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Customer) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Customer) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Customer) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Customer) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Customer) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Customer) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *Customer) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Customer) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Customer) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Customer) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetKey

`func (o *Customer) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Customer) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Customer) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *Customer) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetUsageAttribution

`func (o *Customer) GetUsageAttribution() CustomerUsageAttribution`

GetUsageAttribution returns the UsageAttribution field if non-nil, zero value otherwise.

### GetUsageAttributionOk

`func (o *Customer) GetUsageAttributionOk() (*CustomerUsageAttribution, bool)`

GetUsageAttributionOk returns a tuple with the UsageAttribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageAttribution

`func (o *Customer) SetUsageAttribution(v CustomerUsageAttribution)`

SetUsageAttribution sets UsageAttribution field to given value.


### GetPrimaryEmail

`func (o *Customer) GetPrimaryEmail() string`

GetPrimaryEmail returns the PrimaryEmail field if non-nil, zero value otherwise.

### GetPrimaryEmailOk

`func (o *Customer) GetPrimaryEmailOk() (*string, bool)`

GetPrimaryEmailOk returns a tuple with the PrimaryEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryEmail

`func (o *Customer) SetPrimaryEmail(v string)`

SetPrimaryEmail sets PrimaryEmail field to given value.

### HasPrimaryEmail

`func (o *Customer) HasPrimaryEmail() bool`

HasPrimaryEmail returns a boolean if a field has been set.

### GetCurrency

`func (o *Customer) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Customer) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Customer) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Customer) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetBillingAddress

`func (o *Customer) GetBillingAddress() Address`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *Customer) GetBillingAddressOk() (*Address, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *Customer) SetBillingAddress(v Address)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *Customer) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetCurrentSubscriptionId

`func (o *Customer) GetCurrentSubscriptionId() string`

GetCurrentSubscriptionId returns the CurrentSubscriptionId field if non-nil, zero value otherwise.

### GetCurrentSubscriptionIdOk

`func (o *Customer) GetCurrentSubscriptionIdOk() (*string, bool)`

GetCurrentSubscriptionIdOk returns a tuple with the CurrentSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSubscriptionId

`func (o *Customer) SetCurrentSubscriptionId(v string)`

SetCurrentSubscriptionId sets CurrentSubscriptionId field to given value.

### HasCurrentSubscriptionId

`func (o *Customer) HasCurrentSubscriptionId() bool`

HasCurrentSubscriptionId returns a boolean if a field has been set.

### GetSubscriptions

`func (o *Customer) GetSubscriptions() []Subscription`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *Customer) GetSubscriptionsOk() (*[]Subscription, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *Customer) SetSubscriptions(v []Subscription)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *Customer) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetAnnotations

`func (o *Customer) GetAnnotations() map[string]interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *Customer) GetAnnotationsOk() (*map[string]interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *Customer) SetAnnotations(v map[string]interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *Customer) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


