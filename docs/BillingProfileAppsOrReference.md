# BillingProfileAppsOrReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tax** | [**AppReference**](AppReference.md) | The tax app used for this workflow | [readonly] 
**Invoicing** | [**AppReference**](AppReference.md) | The invoicing app used for this workflow | [readonly] 
**Payment** | [**AppReference**](AppReference.md) | The payment app used for this workflow | [readonly] 

## Methods

### NewBillingProfileAppsOrReference

`func NewBillingProfileAppsOrReference(tax AppReference, invoicing AppReference, payment AppReference, ) *BillingProfileAppsOrReference`

NewBillingProfileAppsOrReference instantiates a new BillingProfileAppsOrReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingProfileAppsOrReferenceWithDefaults

`func NewBillingProfileAppsOrReferenceWithDefaults() *BillingProfileAppsOrReference`

NewBillingProfileAppsOrReferenceWithDefaults instantiates a new BillingProfileAppsOrReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTax

`func (o *BillingProfileAppsOrReference) GetTax() AppReference`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *BillingProfileAppsOrReference) GetTaxOk() (*AppReference, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *BillingProfileAppsOrReference) SetTax(v AppReference)`

SetTax sets Tax field to given value.


### GetInvoicing

`func (o *BillingProfileAppsOrReference) GetInvoicing() AppReference`

GetInvoicing returns the Invoicing field if non-nil, zero value otherwise.

### GetInvoicingOk

`func (o *BillingProfileAppsOrReference) GetInvoicingOk() (*AppReference, bool)`

GetInvoicingOk returns a tuple with the Invoicing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicing

`func (o *BillingProfileAppsOrReference) SetInvoicing(v AppReference)`

SetInvoicing sets Invoicing field to given value.


### GetPayment

`func (o *BillingProfileAppsOrReference) GetPayment() AppReference`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *BillingProfileAppsOrReference) GetPaymentOk() (*AppReference, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *BillingProfileAppsOrReference) SetPayment(v AppReference)`

SetPayment sets Payment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


