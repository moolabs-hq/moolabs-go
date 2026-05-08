# InvoiceWorkflowSettingsReplaceUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invoicing** | [**InvoiceWorkflowInvoicingSettingsReplaceUpdate**](InvoiceWorkflowInvoicingSettingsReplaceUpdate.md) | The invoicing settings for this workflow | 
**Payment** | [**BillingWorkflowPaymentSettings**](BillingWorkflowPaymentSettings.md) | The payment settings for this workflow | 

## Methods

### NewInvoiceWorkflowSettingsReplaceUpdate

`func NewInvoiceWorkflowSettingsReplaceUpdate(invoicing InvoiceWorkflowInvoicingSettingsReplaceUpdate, payment BillingWorkflowPaymentSettings, ) *InvoiceWorkflowSettingsReplaceUpdate`

NewInvoiceWorkflowSettingsReplaceUpdate instantiates a new InvoiceWorkflowSettingsReplaceUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceWorkflowSettingsReplaceUpdateWithDefaults

`func NewInvoiceWorkflowSettingsReplaceUpdateWithDefaults() *InvoiceWorkflowSettingsReplaceUpdate`

NewInvoiceWorkflowSettingsReplaceUpdateWithDefaults instantiates a new InvoiceWorkflowSettingsReplaceUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoicing

`func (o *InvoiceWorkflowSettingsReplaceUpdate) GetInvoicing() InvoiceWorkflowInvoicingSettingsReplaceUpdate`

GetInvoicing returns the Invoicing field if non-nil, zero value otherwise.

### GetInvoicingOk

`func (o *InvoiceWorkflowSettingsReplaceUpdate) GetInvoicingOk() (*InvoiceWorkflowInvoicingSettingsReplaceUpdate, bool)`

GetInvoicingOk returns a tuple with the Invoicing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicing

`func (o *InvoiceWorkflowSettingsReplaceUpdate) SetInvoicing(v InvoiceWorkflowInvoicingSettingsReplaceUpdate)`

SetInvoicing sets Invoicing field to given value.


### GetPayment

`func (o *InvoiceWorkflowSettingsReplaceUpdate) GetPayment() BillingWorkflowPaymentSettings`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *InvoiceWorkflowSettingsReplaceUpdate) GetPaymentOk() (*BillingWorkflowPaymentSettings, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *InvoiceWorkflowSettingsReplaceUpdate) SetPayment(v BillingWorkflowPaymentSettings)`

SetPayment sets Payment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


