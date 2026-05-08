# BillingWorkflowCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collection** | Pointer to [**BillingWorkflowCollectionSettings**](BillingWorkflowCollectionSettings.md) | The collection settings for this workflow | [optional] 
**Invoicing** | Pointer to [**BillingWorkflowInvoicingSettings**](BillingWorkflowInvoicingSettings.md) | The invoicing settings for this workflow | [optional] 
**Payment** | Pointer to [**BillingWorkflowPaymentSettings**](BillingWorkflowPaymentSettings.md) | The payment settings for this workflow | [optional] 
**Tax** | Pointer to [**BillingWorkflowTaxSettings**](BillingWorkflowTaxSettings.md) | The tax settings for this workflow | [optional] 

## Methods

### NewBillingWorkflowCreate

`func NewBillingWorkflowCreate() *BillingWorkflowCreate`

NewBillingWorkflowCreate instantiates a new BillingWorkflowCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingWorkflowCreateWithDefaults

`func NewBillingWorkflowCreateWithDefaults() *BillingWorkflowCreate`

NewBillingWorkflowCreateWithDefaults instantiates a new BillingWorkflowCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollection

`func (o *BillingWorkflowCreate) GetCollection() BillingWorkflowCollectionSettings`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *BillingWorkflowCreate) GetCollectionOk() (*BillingWorkflowCollectionSettings, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *BillingWorkflowCreate) SetCollection(v BillingWorkflowCollectionSettings)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *BillingWorkflowCreate) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetInvoicing

`func (o *BillingWorkflowCreate) GetInvoicing() BillingWorkflowInvoicingSettings`

GetInvoicing returns the Invoicing field if non-nil, zero value otherwise.

### GetInvoicingOk

`func (o *BillingWorkflowCreate) GetInvoicingOk() (*BillingWorkflowInvoicingSettings, bool)`

GetInvoicingOk returns a tuple with the Invoicing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicing

`func (o *BillingWorkflowCreate) SetInvoicing(v BillingWorkflowInvoicingSettings)`

SetInvoicing sets Invoicing field to given value.

### HasInvoicing

`func (o *BillingWorkflowCreate) HasInvoicing() bool`

HasInvoicing returns a boolean if a field has been set.

### GetPayment

`func (o *BillingWorkflowCreate) GetPayment() BillingWorkflowPaymentSettings`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *BillingWorkflowCreate) GetPaymentOk() (*BillingWorkflowPaymentSettings, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *BillingWorkflowCreate) SetPayment(v BillingWorkflowPaymentSettings)`

SetPayment sets Payment field to given value.

### HasPayment

`func (o *BillingWorkflowCreate) HasPayment() bool`

HasPayment returns a boolean if a field has been set.

### GetTax

`func (o *BillingWorkflowCreate) GetTax() BillingWorkflowTaxSettings`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *BillingWorkflowCreate) GetTaxOk() (*BillingWorkflowTaxSettings, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *BillingWorkflowCreate) SetTax(v BillingWorkflowTaxSettings)`

SetTax sets Tax field to given value.

### HasTax

`func (o *BillingWorkflowCreate) HasTax() bool`

HasTax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


