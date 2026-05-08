# CustomerReplaceUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Human-readable name for the resource. Between 1 and 256 characters. | 
**Description** | Pointer to **string** | Optional description of the resource. Maximum 1024 characters. | [optional] 
**Metadata** | Pointer to **map[string]string** | Additional metadata for the resource. | [optional] 
**Key** | Pointer to **string** | An optional unique key of the customer. Useful to reference the customer in external systems. For example, your database ID. | [optional] 
**UsageAttribution** | [**CustomerUsageAttribution**](CustomerUsageAttribution.md) | Mapping to attribute metered usage to the customer | 
**PrimaryEmail** | Pointer to **string** | The primary email address of the customer. | [optional] 
**Currency** | Pointer to **string** | Currency of the customer. Used for billing, tax and invoicing. | [optional] 
**BillingAddress** | Pointer to [**Address**](Address.md) | The billing address of the customer. Used for tax and invoicing. | [optional] 

## Methods

### NewCustomerReplaceUpdate

`func NewCustomerReplaceUpdate(name string, usageAttribution CustomerUsageAttribution, ) *CustomerReplaceUpdate`

NewCustomerReplaceUpdate instantiates a new CustomerReplaceUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerReplaceUpdateWithDefaults

`func NewCustomerReplaceUpdateWithDefaults() *CustomerReplaceUpdate`

NewCustomerReplaceUpdateWithDefaults instantiates a new CustomerReplaceUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CustomerReplaceUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomerReplaceUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomerReplaceUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CustomerReplaceUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomerReplaceUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomerReplaceUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomerReplaceUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *CustomerReplaceUpdate) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CustomerReplaceUpdate) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CustomerReplaceUpdate) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CustomerReplaceUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetKey

`func (o *CustomerReplaceUpdate) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CustomerReplaceUpdate) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CustomerReplaceUpdate) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CustomerReplaceUpdate) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetUsageAttribution

`func (o *CustomerReplaceUpdate) GetUsageAttribution() CustomerUsageAttribution`

GetUsageAttribution returns the UsageAttribution field if non-nil, zero value otherwise.

### GetUsageAttributionOk

`func (o *CustomerReplaceUpdate) GetUsageAttributionOk() (*CustomerUsageAttribution, bool)`

GetUsageAttributionOk returns a tuple with the UsageAttribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageAttribution

`func (o *CustomerReplaceUpdate) SetUsageAttribution(v CustomerUsageAttribution)`

SetUsageAttribution sets UsageAttribution field to given value.


### GetPrimaryEmail

`func (o *CustomerReplaceUpdate) GetPrimaryEmail() string`

GetPrimaryEmail returns the PrimaryEmail field if non-nil, zero value otherwise.

### GetPrimaryEmailOk

`func (o *CustomerReplaceUpdate) GetPrimaryEmailOk() (*string, bool)`

GetPrimaryEmailOk returns a tuple with the PrimaryEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryEmail

`func (o *CustomerReplaceUpdate) SetPrimaryEmail(v string)`

SetPrimaryEmail sets PrimaryEmail field to given value.

### HasPrimaryEmail

`func (o *CustomerReplaceUpdate) HasPrimaryEmail() bool`

HasPrimaryEmail returns a boolean if a field has been set.

### GetCurrency

`func (o *CustomerReplaceUpdate) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CustomerReplaceUpdate) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CustomerReplaceUpdate) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CustomerReplaceUpdate) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetBillingAddress

`func (o *CustomerReplaceUpdate) GetBillingAddress() Address`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *CustomerReplaceUpdate) GetBillingAddressOk() (*Address, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *CustomerReplaceUpdate) SetBillingAddress(v Address)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *CustomerReplaceUpdate) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


