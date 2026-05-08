# InvoiceWorkflowInvoicingSettingsReplaceUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoAdvance** | Pointer to **bool** | Whether to automatically issue the invoice after the draftPeriod has passed. | [optional] [default to true]
**DraftPeriod** | Pointer to **string** | The period for the invoice to be kept in draft status for manual reviews. | [optional] [default to "P0D"]
**DueAfter** | Pointer to **string** | The period after which the invoice is due. With some payment solutions it&#39;s only applicable for manual collection method. | [optional] [default to "P30D"]
**DefaultTaxConfig** | Pointer to [**TaxConfig**](TaxConfig.md) | Default tax configuration to apply to the invoices. | [optional] 

## Methods

### NewInvoiceWorkflowInvoicingSettingsReplaceUpdate

`func NewInvoiceWorkflowInvoicingSettingsReplaceUpdate() *InvoiceWorkflowInvoicingSettingsReplaceUpdate`

NewInvoiceWorkflowInvoicingSettingsReplaceUpdate instantiates a new InvoiceWorkflowInvoicingSettingsReplaceUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceWorkflowInvoicingSettingsReplaceUpdateWithDefaults

`func NewInvoiceWorkflowInvoicingSettingsReplaceUpdateWithDefaults() *InvoiceWorkflowInvoicingSettingsReplaceUpdate`

NewInvoiceWorkflowInvoicingSettingsReplaceUpdateWithDefaults instantiates a new InvoiceWorkflowInvoicingSettingsReplaceUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoAdvance

`func (o *InvoiceWorkflowInvoicingSettingsReplaceUpdate) GetAutoAdvance() bool`

GetAutoAdvance returns the AutoAdvance field if non-nil, zero value otherwise.

### GetAutoAdvanceOk

`func (o *InvoiceWorkflowInvoicingSettingsReplaceUpdate) GetAutoAdvanceOk() (*bool, bool)`

GetAutoAdvanceOk returns a tuple with the AutoAdvance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAdvance

`func (o *InvoiceWorkflowInvoicingSettingsReplaceUpdate) SetAutoAdvance(v bool)`

SetAutoAdvance sets AutoAdvance field to given value.

### HasAutoAdvance

`func (o *InvoiceWorkflowInvoicingSettingsReplaceUpdate) HasAutoAdvance() bool`

HasAutoAdvance returns a boolean if a field has been set.

### GetDraftPeriod

`func (o *InvoiceWorkflowInvoicingSettingsReplaceUpdate) GetDraftPeriod() string`

GetDraftPeriod returns the DraftPeriod field if non-nil, zero value otherwise.

### GetDraftPeriodOk

`func (o *InvoiceWorkflowInvoicingSettingsReplaceUpdate) GetDraftPeriodOk() (*string, bool)`

GetDraftPeriodOk returns a tuple with the DraftPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftPeriod

`func (o *InvoiceWorkflowInvoicingSettingsReplaceUpdate) SetDraftPeriod(v string)`

SetDraftPeriod sets DraftPeriod field to given value.

### HasDraftPeriod

`func (o *InvoiceWorkflowInvoicingSettingsReplaceUpdate) HasDraftPeriod() bool`

HasDraftPeriod returns a boolean if a field has been set.

### GetDueAfter

`func (o *InvoiceWorkflowInvoicingSettingsReplaceUpdate) GetDueAfter() string`

GetDueAfter returns the DueAfter field if non-nil, zero value otherwise.

### GetDueAfterOk

`func (o *InvoiceWorkflowInvoicingSettingsReplaceUpdate) GetDueAfterOk() (*string, bool)`

GetDueAfterOk returns a tuple with the DueAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueAfter

`func (o *InvoiceWorkflowInvoicingSettingsReplaceUpdate) SetDueAfter(v string)`

SetDueAfter sets DueAfter field to given value.

### HasDueAfter

`func (o *InvoiceWorkflowInvoicingSettingsReplaceUpdate) HasDueAfter() bool`

HasDueAfter returns a boolean if a field has been set.

### GetDefaultTaxConfig

`func (o *InvoiceWorkflowInvoicingSettingsReplaceUpdate) GetDefaultTaxConfig() TaxConfig`

GetDefaultTaxConfig returns the DefaultTaxConfig field if non-nil, zero value otherwise.

### GetDefaultTaxConfigOk

`func (o *InvoiceWorkflowInvoicingSettingsReplaceUpdate) GetDefaultTaxConfigOk() (*TaxConfig, bool)`

GetDefaultTaxConfigOk returns a tuple with the DefaultTaxConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTaxConfig

`func (o *InvoiceWorkflowInvoicingSettingsReplaceUpdate) SetDefaultTaxConfig(v TaxConfig)`

SetDefaultTaxConfig sets DefaultTaxConfig field to given value.

### HasDefaultTaxConfig

`func (o *InvoiceWorkflowInvoicingSettingsReplaceUpdate) HasDefaultTaxConfig() bool`

HasDefaultTaxConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


