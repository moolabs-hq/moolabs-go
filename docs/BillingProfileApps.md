# BillingProfileApps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tax** | [**App**](App.md) | The tax app used for this workflow | [readonly] 
**Invoicing** | [**App**](App.md) | The invoicing app used for this workflow | [readonly] 
**Payment** | [**App**](App.md) | The payment app used for this workflow | [readonly] 

## Methods

### NewBillingProfileApps

`func NewBillingProfileApps(tax App, invoicing App, payment App, ) *BillingProfileApps`

NewBillingProfileApps instantiates a new BillingProfileApps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingProfileAppsWithDefaults

`func NewBillingProfileAppsWithDefaults() *BillingProfileApps`

NewBillingProfileAppsWithDefaults instantiates a new BillingProfileApps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTax

`func (o *BillingProfileApps) GetTax() App`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *BillingProfileApps) GetTaxOk() (*App, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *BillingProfileApps) SetTax(v App)`

SetTax sets Tax field to given value.


### GetInvoicing

`func (o *BillingProfileApps) GetInvoicing() App`

GetInvoicing returns the Invoicing field if non-nil, zero value otherwise.

### GetInvoicingOk

`func (o *BillingProfileApps) GetInvoicingOk() (*App, bool)`

GetInvoicingOk returns a tuple with the Invoicing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicing

`func (o *BillingProfileApps) SetInvoicing(v App)`

SetInvoicing sets Invoicing field to given value.


### GetPayment

`func (o *BillingProfileApps) GetPayment() App`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *BillingProfileApps) GetPaymentOk() (*App, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *BillingProfileApps) SetPayment(v App)`

SetPayment sets Payment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


