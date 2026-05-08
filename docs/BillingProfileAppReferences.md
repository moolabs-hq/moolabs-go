# BillingProfileAppReferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tax** | [**AppReference**](AppReference.md) | The tax app used for this workflow | [readonly] 
**Invoicing** | [**AppReference**](AppReference.md) | The invoicing app used for this workflow | [readonly] 
**Payment** | [**AppReference**](AppReference.md) | The payment app used for this workflow | [readonly] 

## Methods

### NewBillingProfileAppReferences

`func NewBillingProfileAppReferences(tax AppReference, invoicing AppReference, payment AppReference, ) *BillingProfileAppReferences`

NewBillingProfileAppReferences instantiates a new BillingProfileAppReferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingProfileAppReferencesWithDefaults

`func NewBillingProfileAppReferencesWithDefaults() *BillingProfileAppReferences`

NewBillingProfileAppReferencesWithDefaults instantiates a new BillingProfileAppReferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTax

`func (o *BillingProfileAppReferences) GetTax() AppReference`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *BillingProfileAppReferences) GetTaxOk() (*AppReference, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *BillingProfileAppReferences) SetTax(v AppReference)`

SetTax sets Tax field to given value.


### GetInvoicing

`func (o *BillingProfileAppReferences) GetInvoicing() AppReference`

GetInvoicing returns the Invoicing field if non-nil, zero value otherwise.

### GetInvoicingOk

`func (o *BillingProfileAppReferences) GetInvoicingOk() (*AppReference, bool)`

GetInvoicingOk returns a tuple with the Invoicing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicing

`func (o *BillingProfileAppReferences) SetInvoicing(v AppReference)`

SetInvoicing sets Invoicing field to given value.


### GetPayment

`func (o *BillingProfileAppReferences) GetPayment() AppReference`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *BillingProfileAppReferences) GetPaymentOk() (*AppReference, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *BillingProfileAppReferences) SetPayment(v AppReference)`

SetPayment sets Payment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


