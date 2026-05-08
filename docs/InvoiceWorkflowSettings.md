# InvoiceWorkflowSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apps** | Pointer to [**BillingProfileAppsOrReference**](BillingProfileAppsOrReference.md) | The apps that will be used to orchestrate the invoice&#39;s workflow. | [optional] [readonly] 
**SourceBillingProfileId** | **string** | sourceBillingProfileID is the billing profile on which the workflow was based on.  The profile is snapshotted on invoice creation, after which it can be altered independently of the profile itself. | [readonly] 
**Workflow** | [**BillingWorkflow**](BillingWorkflow.md) | The workflow details used by this invoice. | 

## Methods

### NewInvoiceWorkflowSettings

`func NewInvoiceWorkflowSettings(sourceBillingProfileId string, workflow BillingWorkflow, ) *InvoiceWorkflowSettings`

NewInvoiceWorkflowSettings instantiates a new InvoiceWorkflowSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceWorkflowSettingsWithDefaults

`func NewInvoiceWorkflowSettingsWithDefaults() *InvoiceWorkflowSettings`

NewInvoiceWorkflowSettingsWithDefaults instantiates a new InvoiceWorkflowSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApps

`func (o *InvoiceWorkflowSettings) GetApps() BillingProfileAppsOrReference`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *InvoiceWorkflowSettings) GetAppsOk() (*BillingProfileAppsOrReference, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *InvoiceWorkflowSettings) SetApps(v BillingProfileAppsOrReference)`

SetApps sets Apps field to given value.

### HasApps

`func (o *InvoiceWorkflowSettings) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetSourceBillingProfileId

`func (o *InvoiceWorkflowSettings) GetSourceBillingProfileId() string`

GetSourceBillingProfileId returns the SourceBillingProfileId field if non-nil, zero value otherwise.

### GetSourceBillingProfileIdOk

`func (o *InvoiceWorkflowSettings) GetSourceBillingProfileIdOk() (*string, bool)`

GetSourceBillingProfileIdOk returns a tuple with the SourceBillingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceBillingProfileId

`func (o *InvoiceWorkflowSettings) SetSourceBillingProfileId(v string)`

SetSourceBillingProfileId sets SourceBillingProfileId field to given value.


### GetWorkflow

`func (o *InvoiceWorkflowSettings) GetWorkflow() BillingWorkflow`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *InvoiceWorkflowSettings) GetWorkflowOk() (*BillingWorkflow, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *InvoiceWorkflowSettings) SetWorkflow(v BillingWorkflow)`

SetWorkflow sets Workflow field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


