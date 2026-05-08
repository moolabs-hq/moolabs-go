# CreateCheckoutSessionTaxIdCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Enable tax ID collection during checkout. Defaults to false. | 
**Required** | Pointer to [**CreateCheckoutSessionTaxIdCollectionRequired**](CreateCheckoutSessionTaxIdCollectionRequired.md) | Describes whether a tax ID is required during checkout. Defaults to never. | [optional] 

## Methods

### NewCreateCheckoutSessionTaxIdCollection

`func NewCreateCheckoutSessionTaxIdCollection(enabled bool, ) *CreateCheckoutSessionTaxIdCollection`

NewCreateCheckoutSessionTaxIdCollection instantiates a new CreateCheckoutSessionTaxIdCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCheckoutSessionTaxIdCollectionWithDefaults

`func NewCreateCheckoutSessionTaxIdCollectionWithDefaults() *CreateCheckoutSessionTaxIdCollection`

NewCreateCheckoutSessionTaxIdCollectionWithDefaults instantiates a new CreateCheckoutSessionTaxIdCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *CreateCheckoutSessionTaxIdCollection) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateCheckoutSessionTaxIdCollection) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateCheckoutSessionTaxIdCollection) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetRequired

`func (o *CreateCheckoutSessionTaxIdCollection) GetRequired() CreateCheckoutSessionTaxIdCollectionRequired`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *CreateCheckoutSessionTaxIdCollection) GetRequiredOk() (*CreateCheckoutSessionTaxIdCollectionRequired, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *CreateCheckoutSessionTaxIdCollection) SetRequired(v CreateCheckoutSessionTaxIdCollectionRequired)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *CreateCheckoutSessionTaxIdCollection) HasRequired() bool`

HasRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


