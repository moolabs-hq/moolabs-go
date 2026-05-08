# InvoiceReplaceUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Optional description of the resource. Maximum 1024 characters. | [optional] 
**Metadata** | Pointer to **map[string]string** | Additional metadata for the resource. | [optional] 
**Supplier** | [**BillingPartyReplaceUpdate**](BillingPartyReplaceUpdate.md) | The supplier of the lines included in the invoice. | 
**Customer** | [**BillingPartyReplaceUpdate**](BillingPartyReplaceUpdate.md) | The customer the invoice is sent to. | 
**Lines** | [**[]InvoiceLineReplaceUpdate**](InvoiceLineReplaceUpdate.md) | The lines included in the invoice. | 
**Workflow** | [**InvoiceWorkflowReplaceUpdate**](InvoiceWorkflowReplaceUpdate.md) | The workflow settings for the invoice. | 

## Methods

### NewInvoiceReplaceUpdate

`func NewInvoiceReplaceUpdate(supplier BillingPartyReplaceUpdate, customer BillingPartyReplaceUpdate, lines []InvoiceLineReplaceUpdate, workflow InvoiceWorkflowReplaceUpdate, ) *InvoiceReplaceUpdate`

NewInvoiceReplaceUpdate instantiates a new InvoiceReplaceUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceReplaceUpdateWithDefaults

`func NewInvoiceReplaceUpdateWithDefaults() *InvoiceReplaceUpdate`

NewInvoiceReplaceUpdateWithDefaults instantiates a new InvoiceReplaceUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *InvoiceReplaceUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InvoiceReplaceUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InvoiceReplaceUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InvoiceReplaceUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *InvoiceReplaceUpdate) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InvoiceReplaceUpdate) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InvoiceReplaceUpdate) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InvoiceReplaceUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSupplier

`func (o *InvoiceReplaceUpdate) GetSupplier() BillingPartyReplaceUpdate`

GetSupplier returns the Supplier field if non-nil, zero value otherwise.

### GetSupplierOk

`func (o *InvoiceReplaceUpdate) GetSupplierOk() (*BillingPartyReplaceUpdate, bool)`

GetSupplierOk returns a tuple with the Supplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplier

`func (o *InvoiceReplaceUpdate) SetSupplier(v BillingPartyReplaceUpdate)`

SetSupplier sets Supplier field to given value.


### GetCustomer

`func (o *InvoiceReplaceUpdate) GetCustomer() BillingPartyReplaceUpdate`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *InvoiceReplaceUpdate) GetCustomerOk() (*BillingPartyReplaceUpdate, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *InvoiceReplaceUpdate) SetCustomer(v BillingPartyReplaceUpdate)`

SetCustomer sets Customer field to given value.


### GetLines

`func (o *InvoiceReplaceUpdate) GetLines() []InvoiceLineReplaceUpdate`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *InvoiceReplaceUpdate) GetLinesOk() (*[]InvoiceLineReplaceUpdate, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *InvoiceReplaceUpdate) SetLines(v []InvoiceLineReplaceUpdate)`

SetLines sets Lines field to given value.


### GetWorkflow

`func (o *InvoiceReplaceUpdate) GetWorkflow() InvoiceWorkflowReplaceUpdate`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *InvoiceReplaceUpdate) GetWorkflowOk() (*InvoiceWorkflowReplaceUpdate, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *InvoiceReplaceUpdate) SetWorkflow(v InvoiceWorkflowReplaceUpdate)`

SetWorkflow sets Workflow field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


