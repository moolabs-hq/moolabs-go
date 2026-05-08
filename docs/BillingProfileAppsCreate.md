# BillingProfileAppsCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tax** | **string** | The tax app used for this workflow | 
**Invoicing** | **string** | The invoicing app used for this workflow | 
**Payment** | **string** | The payment app used for this workflow | 

## Methods

### NewBillingProfileAppsCreate

`func NewBillingProfileAppsCreate(tax string, invoicing string, payment string, ) *BillingProfileAppsCreate`

NewBillingProfileAppsCreate instantiates a new BillingProfileAppsCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingProfileAppsCreateWithDefaults

`func NewBillingProfileAppsCreateWithDefaults() *BillingProfileAppsCreate`

NewBillingProfileAppsCreateWithDefaults instantiates a new BillingProfileAppsCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTax

`func (o *BillingProfileAppsCreate) GetTax() string`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *BillingProfileAppsCreate) GetTaxOk() (*string, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *BillingProfileAppsCreate) SetTax(v string)`

SetTax sets Tax field to given value.


### GetInvoicing

`func (o *BillingProfileAppsCreate) GetInvoicing() string`

GetInvoicing returns the Invoicing field if non-nil, zero value otherwise.

### GetInvoicingOk

`func (o *BillingProfileAppsCreate) GetInvoicingOk() (*string, bool)`

GetInvoicingOk returns a tuple with the Invoicing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicing

`func (o *BillingProfileAppsCreate) SetInvoicing(v string)`

SetInvoicing sets Invoicing field to given value.


### GetPayment

`func (o *BillingProfileAppsCreate) GetPayment() string`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *BillingProfileAppsCreate) GetPaymentOk() (*string, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *BillingProfileAppsCreate) SetPayment(v string)`

SetPayment sets Payment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


