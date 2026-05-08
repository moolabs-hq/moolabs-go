# BillingInvoiceCustomerExtendedDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the party (if available) | [optional] [readonly] 
**Key** | Pointer to **string** | An optional unique key of the party (if available) | [optional] 
**Name** | Pointer to **string** | Legal name or representation of the organization. | [optional] 
**TaxId** | Pointer to [**BillingPartyTaxIdentity**](BillingPartyTaxIdentity.md) | The entity&#39;s legal ID code used for tax purposes. They may have other numbers, but we&#39;re only interested in those valid for tax purposes. | [optional] 
**Addresses** | Pointer to [**[]Address**](Address.md) | Regular post addresses for where information should be sent if needed. | [optional] 
**UsageAttribution** | [**CustomerUsageAttribution**](CustomerUsageAttribution.md) | Mapping to attribute metered usage to the customer | 

## Methods

### NewBillingInvoiceCustomerExtendedDetails

`func NewBillingInvoiceCustomerExtendedDetails(usageAttribution CustomerUsageAttribution, ) *BillingInvoiceCustomerExtendedDetails`

NewBillingInvoiceCustomerExtendedDetails instantiates a new BillingInvoiceCustomerExtendedDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingInvoiceCustomerExtendedDetailsWithDefaults

`func NewBillingInvoiceCustomerExtendedDetailsWithDefaults() *BillingInvoiceCustomerExtendedDetails`

NewBillingInvoiceCustomerExtendedDetailsWithDefaults instantiates a new BillingInvoiceCustomerExtendedDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BillingInvoiceCustomerExtendedDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingInvoiceCustomerExtendedDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingInvoiceCustomerExtendedDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillingInvoiceCustomerExtendedDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *BillingInvoiceCustomerExtendedDetails) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *BillingInvoiceCustomerExtendedDetails) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *BillingInvoiceCustomerExtendedDetails) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *BillingInvoiceCustomerExtendedDetails) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *BillingInvoiceCustomerExtendedDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingInvoiceCustomerExtendedDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingInvoiceCustomerExtendedDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BillingInvoiceCustomerExtendedDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTaxId

`func (o *BillingInvoiceCustomerExtendedDetails) GetTaxId() BillingPartyTaxIdentity`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *BillingInvoiceCustomerExtendedDetails) GetTaxIdOk() (*BillingPartyTaxIdentity, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *BillingInvoiceCustomerExtendedDetails) SetTaxId(v BillingPartyTaxIdentity)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *BillingInvoiceCustomerExtendedDetails) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### GetAddresses

`func (o *BillingInvoiceCustomerExtendedDetails) GetAddresses() []Address`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *BillingInvoiceCustomerExtendedDetails) GetAddressesOk() (*[]Address, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *BillingInvoiceCustomerExtendedDetails) SetAddresses(v []Address)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *BillingInvoiceCustomerExtendedDetails) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetUsageAttribution

`func (o *BillingInvoiceCustomerExtendedDetails) GetUsageAttribution() CustomerUsageAttribution`

GetUsageAttribution returns the UsageAttribution field if non-nil, zero value otherwise.

### GetUsageAttributionOk

`func (o *BillingInvoiceCustomerExtendedDetails) GetUsageAttributionOk() (*CustomerUsageAttribution, bool)`

GetUsageAttributionOk returns a tuple with the UsageAttribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageAttribution

`func (o *BillingInvoiceCustomerExtendedDetails) SetUsageAttribution(v CustomerUsageAttribution)`

SetUsageAttribution sets UsageAttribution field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


