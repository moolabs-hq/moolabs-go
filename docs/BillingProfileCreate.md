# BillingProfileCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Human-readable name for the resource. Between 1 and 256 characters. | 
**Description** | Pointer to **string** | Optional description of the resource. Maximum 1024 characters. | [optional] 
**Metadata** | Pointer to **map[string]string** | Additional metadata for the resource. | [optional] 
**Supplier** | [**BillingParty**](BillingParty.md) | The name and contact information for the supplier this billing profile represents | 
**Default** | **bool** | Is this the default profile? | 
**Workflow** | [**BillingWorkflowCreate**](BillingWorkflowCreate.md) | The billing workflow settings for this profile. | 
**Apps** | [**BillingProfileAppsCreate**](BillingProfileAppsCreate.md) | The apps used by this billing profile. | 

## Methods

### NewBillingProfileCreate

`func NewBillingProfileCreate(name string, supplier BillingParty, default_ bool, workflow BillingWorkflowCreate, apps BillingProfileAppsCreate, ) *BillingProfileCreate`

NewBillingProfileCreate instantiates a new BillingProfileCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingProfileCreateWithDefaults

`func NewBillingProfileCreateWithDefaults() *BillingProfileCreate`

NewBillingProfileCreateWithDefaults instantiates a new BillingProfileCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BillingProfileCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingProfileCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingProfileCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BillingProfileCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BillingProfileCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BillingProfileCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BillingProfileCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *BillingProfileCreate) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BillingProfileCreate) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BillingProfileCreate) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BillingProfileCreate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSupplier

`func (o *BillingProfileCreate) GetSupplier() BillingParty`

GetSupplier returns the Supplier field if non-nil, zero value otherwise.

### GetSupplierOk

`func (o *BillingProfileCreate) GetSupplierOk() (*BillingParty, bool)`

GetSupplierOk returns a tuple with the Supplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplier

`func (o *BillingProfileCreate) SetSupplier(v BillingParty)`

SetSupplier sets Supplier field to given value.


### GetDefault

`func (o *BillingProfileCreate) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *BillingProfileCreate) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *BillingProfileCreate) SetDefault(v bool)`

SetDefault sets Default field to given value.


### GetWorkflow

`func (o *BillingProfileCreate) GetWorkflow() BillingWorkflowCreate`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *BillingProfileCreate) GetWorkflowOk() (*BillingWorkflowCreate, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *BillingProfileCreate) SetWorkflow(v BillingWorkflowCreate)`

SetWorkflow sets Workflow field to given value.


### GetApps

`func (o *BillingProfileCreate) GetApps() BillingProfileAppsCreate`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *BillingProfileCreate) GetAppsOk() (*BillingProfileAppsCreate, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *BillingProfileCreate) SetApps(v BillingProfileAppsCreate)`

SetApps sets Apps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


