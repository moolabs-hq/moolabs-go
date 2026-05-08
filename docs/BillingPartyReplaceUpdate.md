# BillingPartyReplaceUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | An optional unique key of the party (if available) | [optional] 
**Name** | Pointer to **string** | Legal name or representation of the organization. | [optional] 
**TaxId** | Pointer to [**BillingPartyTaxIdentity**](BillingPartyTaxIdentity.md) | The entity&#39;s legal ID code used for tax purposes. They may have other numbers, but we&#39;re only interested in those valid for tax purposes. | [optional] 
**Addresses** | Pointer to [**[]Address**](Address.md) | Regular post addresses for where information should be sent if needed. | [optional] 

## Methods

### NewBillingPartyReplaceUpdate

`func NewBillingPartyReplaceUpdate() *BillingPartyReplaceUpdate`

NewBillingPartyReplaceUpdate instantiates a new BillingPartyReplaceUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingPartyReplaceUpdateWithDefaults

`func NewBillingPartyReplaceUpdateWithDefaults() *BillingPartyReplaceUpdate`

NewBillingPartyReplaceUpdateWithDefaults instantiates a new BillingPartyReplaceUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *BillingPartyReplaceUpdate) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *BillingPartyReplaceUpdate) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *BillingPartyReplaceUpdate) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *BillingPartyReplaceUpdate) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *BillingPartyReplaceUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingPartyReplaceUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingPartyReplaceUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BillingPartyReplaceUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTaxId

`func (o *BillingPartyReplaceUpdate) GetTaxId() BillingPartyTaxIdentity`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *BillingPartyReplaceUpdate) GetTaxIdOk() (*BillingPartyTaxIdentity, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *BillingPartyReplaceUpdate) SetTaxId(v BillingPartyTaxIdentity)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *BillingPartyReplaceUpdate) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### GetAddresses

`func (o *BillingPartyReplaceUpdate) GetAddresses() []Address`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *BillingPartyReplaceUpdate) GetAddressesOk() (*[]Address, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *BillingPartyReplaceUpdate) SetAddresses(v []Address)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *BillingPartyReplaceUpdate) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


