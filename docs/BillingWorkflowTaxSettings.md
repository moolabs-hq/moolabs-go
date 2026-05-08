# BillingWorkflowTaxSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enable automatic tax calculation when tax is supported by the app. For example, with Stripe Invoicing when enabled, tax is calculated via Stripe Tax. | [optional] [default to true]
**Enforced** | Pointer to **bool** | Enforce tax calculation when tax is supported by the app. When enabled, OpenMeter will not allow to create an invoice without tax calculation. Enforcement is different per apps, for example, Stripe app requires customer to have a tax location when starting a paid subscription. | [optional] [default to false]

## Methods

### NewBillingWorkflowTaxSettings

`func NewBillingWorkflowTaxSettings() *BillingWorkflowTaxSettings`

NewBillingWorkflowTaxSettings instantiates a new BillingWorkflowTaxSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingWorkflowTaxSettingsWithDefaults

`func NewBillingWorkflowTaxSettingsWithDefaults() *BillingWorkflowTaxSettings`

NewBillingWorkflowTaxSettingsWithDefaults instantiates a new BillingWorkflowTaxSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *BillingWorkflowTaxSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BillingWorkflowTaxSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BillingWorkflowTaxSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BillingWorkflowTaxSettings) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEnforced

`func (o *BillingWorkflowTaxSettings) GetEnforced() bool`

GetEnforced returns the Enforced field if non-nil, zero value otherwise.

### GetEnforcedOk

`func (o *BillingWorkflowTaxSettings) GetEnforcedOk() (*bool, bool)`

GetEnforcedOk returns a tuple with the Enforced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforced

`func (o *BillingWorkflowTaxSettings) SetEnforced(v bool)`

SetEnforced sets Enforced field to given value.

### HasEnforced

`func (o *BillingWorkflowTaxSettings) HasEnforced() bool`

HasEnforced returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


