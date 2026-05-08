# BillingProfileCustomerWorkflowOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collection** | Pointer to [**BillingWorkflowCollectionSettings**](BillingWorkflowCollectionSettings.md) | The collection settings for this workflow | [optional] 
**Invoicing** | Pointer to [**BillingWorkflowInvoicingSettings**](BillingWorkflowInvoicingSettings.md) | The invoicing settings for this workflow | [optional] 
**Payment** | Pointer to [**BillingWorkflowPaymentSettings**](BillingWorkflowPaymentSettings.md) | The payment settings for this workflow | [optional] 
**Tax** | Pointer to [**BillingWorkflowTaxSettings**](BillingWorkflowTaxSettings.md) | The tax settings for this workflow | [optional] 
**TaxApp** | [**AppReadOrCreateOrUpdateOrDeleteOrQuery**](AppReadOrCreateOrUpdateOrDeleteOrQuery.md) | The tax app used for this workflow | [readonly] 
**InvoicingApp** | [**AppReadOrCreateOrUpdateOrDeleteOrQuery**](AppReadOrCreateOrUpdateOrDeleteOrQuery.md) | The invoicing app used for this workflow | [readonly] 
**PaymentApp** | [**AppReadOrCreateOrUpdateOrDeleteOrQuery**](AppReadOrCreateOrUpdateOrDeleteOrQuery.md) | The payment app used for this workflow | [readonly] 

## Methods

### NewBillingProfileCustomerWorkflowOverride

`func NewBillingProfileCustomerWorkflowOverride(taxApp AppReadOrCreateOrUpdateOrDeleteOrQuery, invoicingApp AppReadOrCreateOrUpdateOrDeleteOrQuery, paymentApp AppReadOrCreateOrUpdateOrDeleteOrQuery, ) *BillingProfileCustomerWorkflowOverride`

NewBillingProfileCustomerWorkflowOverride instantiates a new BillingProfileCustomerWorkflowOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingProfileCustomerWorkflowOverrideWithDefaults

`func NewBillingProfileCustomerWorkflowOverrideWithDefaults() *BillingProfileCustomerWorkflowOverride`

NewBillingProfileCustomerWorkflowOverrideWithDefaults instantiates a new BillingProfileCustomerWorkflowOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollection

`func (o *BillingProfileCustomerWorkflowOverride) GetCollection() BillingWorkflowCollectionSettings`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *BillingProfileCustomerWorkflowOverride) GetCollectionOk() (*BillingWorkflowCollectionSettings, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *BillingProfileCustomerWorkflowOverride) SetCollection(v BillingWorkflowCollectionSettings)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *BillingProfileCustomerWorkflowOverride) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetInvoicing

`func (o *BillingProfileCustomerWorkflowOverride) GetInvoicing() BillingWorkflowInvoicingSettings`

GetInvoicing returns the Invoicing field if non-nil, zero value otherwise.

### GetInvoicingOk

`func (o *BillingProfileCustomerWorkflowOverride) GetInvoicingOk() (*BillingWorkflowInvoicingSettings, bool)`

GetInvoicingOk returns a tuple with the Invoicing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicing

`func (o *BillingProfileCustomerWorkflowOverride) SetInvoicing(v BillingWorkflowInvoicingSettings)`

SetInvoicing sets Invoicing field to given value.

### HasInvoicing

`func (o *BillingProfileCustomerWorkflowOverride) HasInvoicing() bool`

HasInvoicing returns a boolean if a field has been set.

### GetPayment

`func (o *BillingProfileCustomerWorkflowOverride) GetPayment() BillingWorkflowPaymentSettings`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *BillingProfileCustomerWorkflowOverride) GetPaymentOk() (*BillingWorkflowPaymentSettings, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *BillingProfileCustomerWorkflowOverride) SetPayment(v BillingWorkflowPaymentSettings)`

SetPayment sets Payment field to given value.

### HasPayment

`func (o *BillingProfileCustomerWorkflowOverride) HasPayment() bool`

HasPayment returns a boolean if a field has been set.

### GetTax

`func (o *BillingProfileCustomerWorkflowOverride) GetTax() BillingWorkflowTaxSettings`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *BillingProfileCustomerWorkflowOverride) GetTaxOk() (*BillingWorkflowTaxSettings, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *BillingProfileCustomerWorkflowOverride) SetTax(v BillingWorkflowTaxSettings)`

SetTax sets Tax field to given value.

### HasTax

`func (o *BillingProfileCustomerWorkflowOverride) HasTax() bool`

HasTax returns a boolean if a field has been set.

### GetTaxApp

`func (o *BillingProfileCustomerWorkflowOverride) GetTaxApp() AppReadOrCreateOrUpdateOrDeleteOrQuery`

GetTaxApp returns the TaxApp field if non-nil, zero value otherwise.

### GetTaxAppOk

`func (o *BillingProfileCustomerWorkflowOverride) GetTaxAppOk() (*AppReadOrCreateOrUpdateOrDeleteOrQuery, bool)`

GetTaxAppOk returns a tuple with the TaxApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxApp

`func (o *BillingProfileCustomerWorkflowOverride) SetTaxApp(v AppReadOrCreateOrUpdateOrDeleteOrQuery)`

SetTaxApp sets TaxApp field to given value.


### GetInvoicingApp

`func (o *BillingProfileCustomerWorkflowOverride) GetInvoicingApp() AppReadOrCreateOrUpdateOrDeleteOrQuery`

GetInvoicingApp returns the InvoicingApp field if non-nil, zero value otherwise.

### GetInvoicingAppOk

`func (o *BillingProfileCustomerWorkflowOverride) GetInvoicingAppOk() (*AppReadOrCreateOrUpdateOrDeleteOrQuery, bool)`

GetInvoicingAppOk returns a tuple with the InvoicingApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicingApp

`func (o *BillingProfileCustomerWorkflowOverride) SetInvoicingApp(v AppReadOrCreateOrUpdateOrDeleteOrQuery)`

SetInvoicingApp sets InvoicingApp field to given value.


### GetPaymentApp

`func (o *BillingProfileCustomerWorkflowOverride) GetPaymentApp() AppReadOrCreateOrUpdateOrDeleteOrQuery`

GetPaymentApp returns the PaymentApp field if non-nil, zero value otherwise.

### GetPaymentAppOk

`func (o *BillingProfileCustomerWorkflowOverride) GetPaymentAppOk() (*AppReadOrCreateOrUpdateOrDeleteOrQuery, bool)`

GetPaymentAppOk returns a tuple with the PaymentApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentApp

`func (o *BillingProfileCustomerWorkflowOverride) SetPaymentApp(v AppReadOrCreateOrUpdateOrDeleteOrQuery)`

SetPaymentApp sets PaymentApp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


