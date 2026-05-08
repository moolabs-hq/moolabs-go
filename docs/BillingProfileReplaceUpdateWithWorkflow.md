# BillingProfileReplaceUpdateWithWorkflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Human-readable name for the resource. Between 1 and 256 characters. | 
**Description** | Pointer to **string** | Optional description of the resource. Maximum 1024 characters. | [optional] 
**Metadata** | Pointer to **map[string]string** | Additional metadata for the resource. | [optional] 
**Supplier** | [**BillingParty**](BillingParty.md) | The name and contact information for the supplier this billing profile represents | 
**Default** | **bool** | Is this the default profile? | 
**Workflow** | [**BillingWorkflow**](BillingWorkflow.md) | The billing workflow settings for this profile. | 

## Methods

### NewBillingProfileReplaceUpdateWithWorkflow

`func NewBillingProfileReplaceUpdateWithWorkflow(name string, supplier BillingParty, default_ bool, workflow BillingWorkflow, ) *BillingProfileReplaceUpdateWithWorkflow`

NewBillingProfileReplaceUpdateWithWorkflow instantiates a new BillingProfileReplaceUpdateWithWorkflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingProfileReplaceUpdateWithWorkflowWithDefaults

`func NewBillingProfileReplaceUpdateWithWorkflowWithDefaults() *BillingProfileReplaceUpdateWithWorkflow`

NewBillingProfileReplaceUpdateWithWorkflowWithDefaults instantiates a new BillingProfileReplaceUpdateWithWorkflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BillingProfileReplaceUpdateWithWorkflow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingProfileReplaceUpdateWithWorkflow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingProfileReplaceUpdateWithWorkflow) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BillingProfileReplaceUpdateWithWorkflow) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BillingProfileReplaceUpdateWithWorkflow) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BillingProfileReplaceUpdateWithWorkflow) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BillingProfileReplaceUpdateWithWorkflow) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *BillingProfileReplaceUpdateWithWorkflow) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BillingProfileReplaceUpdateWithWorkflow) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BillingProfileReplaceUpdateWithWorkflow) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BillingProfileReplaceUpdateWithWorkflow) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSupplier

`func (o *BillingProfileReplaceUpdateWithWorkflow) GetSupplier() BillingParty`

GetSupplier returns the Supplier field if non-nil, zero value otherwise.

### GetSupplierOk

`func (o *BillingProfileReplaceUpdateWithWorkflow) GetSupplierOk() (*BillingParty, bool)`

GetSupplierOk returns a tuple with the Supplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplier

`func (o *BillingProfileReplaceUpdateWithWorkflow) SetSupplier(v BillingParty)`

SetSupplier sets Supplier field to given value.


### GetDefault

`func (o *BillingProfileReplaceUpdateWithWorkflow) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *BillingProfileReplaceUpdateWithWorkflow) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *BillingProfileReplaceUpdateWithWorkflow) SetDefault(v bool)`

SetDefault sets Default field to given value.


### GetWorkflow

`func (o *BillingProfileReplaceUpdateWithWorkflow) GetWorkflow() BillingWorkflow`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *BillingProfileReplaceUpdateWithWorkflow) GetWorkflowOk() (*BillingWorkflow, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *BillingProfileReplaceUpdateWithWorkflow) SetWorkflow(v BillingWorkflow)`

SetWorkflow sets Workflow field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


