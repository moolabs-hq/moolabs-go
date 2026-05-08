# BillingWorkflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collection** | Pointer to [**BillingWorkflowCollectionSettings**](BillingWorkflowCollectionSettings.md) | The collection settings for this workflow | [optional] 
**Invoicing** | Pointer to [**BillingWorkflowInvoicingSettings**](BillingWorkflowInvoicingSettings.md) | The invoicing settings for this workflow | [optional] 
**Payment** | Pointer to [**BillingWorkflowPaymentSettings**](BillingWorkflowPaymentSettings.md) | The payment settings for this workflow | [optional] 
**Tax** | Pointer to [**BillingWorkflowTaxSettings**](BillingWorkflowTaxSettings.md) | The tax settings for this workflow | [optional] 

## Methods

### NewBillingWorkflow

`func NewBillingWorkflow() *BillingWorkflow`

NewBillingWorkflow instantiates a new BillingWorkflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingWorkflowWithDefaults

`func NewBillingWorkflowWithDefaults() *BillingWorkflow`

NewBillingWorkflowWithDefaults instantiates a new BillingWorkflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollection

`func (o *BillingWorkflow) GetCollection() BillingWorkflowCollectionSettings`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *BillingWorkflow) GetCollectionOk() (*BillingWorkflowCollectionSettings, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *BillingWorkflow) SetCollection(v BillingWorkflowCollectionSettings)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *BillingWorkflow) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetInvoicing

`func (o *BillingWorkflow) GetInvoicing() BillingWorkflowInvoicingSettings`

GetInvoicing returns the Invoicing field if non-nil, zero value otherwise.

### GetInvoicingOk

`func (o *BillingWorkflow) GetInvoicingOk() (*BillingWorkflowInvoicingSettings, bool)`

GetInvoicingOk returns a tuple with the Invoicing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicing

`func (o *BillingWorkflow) SetInvoicing(v BillingWorkflowInvoicingSettings)`

SetInvoicing sets Invoicing field to given value.

### HasInvoicing

`func (o *BillingWorkflow) HasInvoicing() bool`

HasInvoicing returns a boolean if a field has been set.

### GetPayment

`func (o *BillingWorkflow) GetPayment() BillingWorkflowPaymentSettings`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *BillingWorkflow) GetPaymentOk() (*BillingWorkflowPaymentSettings, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *BillingWorkflow) SetPayment(v BillingWorkflowPaymentSettings)`

SetPayment sets Payment field to given value.

### HasPayment

`func (o *BillingWorkflow) HasPayment() bool`

HasPayment returns a boolean if a field has been set.

### GetTax

`func (o *BillingWorkflow) GetTax() BillingWorkflowTaxSettings`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *BillingWorkflow) GetTaxOk() (*BillingWorkflowTaxSettings, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *BillingWorkflow) SetTax(v BillingWorkflowTaxSettings)`

SetTax sets Tax field to given value.

### HasTax

`func (o *BillingWorkflow) HasTax() bool`

HasTax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


