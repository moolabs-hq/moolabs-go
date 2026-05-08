# TaxConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Behavior** | Pointer to [**TaxBehavior**](TaxBehavior.md) | Tax behavior.  If not specified the billing profile is used to determine the tax behavior. If not specified in the billing profile, the provider&#39;s default behavior is used. | [optional] 
**Stripe** | Pointer to [**StripeTaxConfig**](StripeTaxConfig.md) | Stripe tax config. | [optional] 
**CustomInvoicing** | Pointer to [**CustomInvoicingTaxConfig**](CustomInvoicingTaxConfig.md) | Custom invoicing tax config. | [optional] 

## Methods

### NewTaxConfig

`func NewTaxConfig() *TaxConfig`

NewTaxConfig instantiates a new TaxConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxConfigWithDefaults

`func NewTaxConfigWithDefaults() *TaxConfig`

NewTaxConfigWithDefaults instantiates a new TaxConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBehavior

`func (o *TaxConfig) GetBehavior() TaxBehavior`

GetBehavior returns the Behavior field if non-nil, zero value otherwise.

### GetBehaviorOk

`func (o *TaxConfig) GetBehaviorOk() (*TaxBehavior, bool)`

GetBehaviorOk returns a tuple with the Behavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehavior

`func (o *TaxConfig) SetBehavior(v TaxBehavior)`

SetBehavior sets Behavior field to given value.

### HasBehavior

`func (o *TaxConfig) HasBehavior() bool`

HasBehavior returns a boolean if a field has been set.

### GetStripe

`func (o *TaxConfig) GetStripe() StripeTaxConfig`

GetStripe returns the Stripe field if non-nil, zero value otherwise.

### GetStripeOk

`func (o *TaxConfig) GetStripeOk() (*StripeTaxConfig, bool)`

GetStripeOk returns a tuple with the Stripe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripe

`func (o *TaxConfig) SetStripe(v StripeTaxConfig)`

SetStripe sets Stripe field to given value.

### HasStripe

`func (o *TaxConfig) HasStripe() bool`

HasStripe returns a boolean if a field has been set.

### GetCustomInvoicing

`func (o *TaxConfig) GetCustomInvoicing() CustomInvoicingTaxConfig`

GetCustomInvoicing returns the CustomInvoicing field if non-nil, zero value otherwise.

### GetCustomInvoicingOk

`func (o *TaxConfig) GetCustomInvoicingOk() (*CustomInvoicingTaxConfig, bool)`

GetCustomInvoicingOk returns a tuple with the CustomInvoicing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomInvoicing

`func (o *TaxConfig) SetCustomInvoicing(v CustomInvoicingTaxConfig)`

SetCustomInvoicing sets CustomInvoicing field to given value.

### HasCustomInvoicing

`func (o *TaxConfig) HasCustomInvoicing() bool`

HasCustomInvoicing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


