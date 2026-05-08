# BillingProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier for the resource. | [readonly] 
**Name** | **string** | Human-readable name for the resource. Between 1 and 256 characters. | 
**Description** | Pointer to **string** | Optional description of the resource. Maximum 1024 characters. | [optional] 
**Metadata** | Pointer to **map[string]string** | Additional metadata for the resource. | [optional] 
**CreatedAt** | **time.Time** | Timestamp of when the resource was created. | [readonly] 
**UpdatedAt** | **time.Time** | Timestamp of when the resource was last updated. | [readonly] 
**DeletedAt** | Pointer to **time.Time** | Timestamp of when the resource was permanently deleted. | [optional] [readonly] 
**Supplier** | [**BillingParty**](BillingParty.md) | The name and contact information for the supplier this billing profile represents | 
**Workflow** | [**BillingWorkflow**](BillingWorkflow.md) | The billing workflow settings for this profile | [readonly] 
**Apps** | [**BillingProfileAppsOrReference**](BillingProfileAppsOrReference.md) | The applications used by this billing profile.  Expand settings govern if this includes the whole app object or just the ID references. | [readonly] 
**Default** | **bool** | Is this the default profile? | 

## Methods

### NewBillingProfile

`func NewBillingProfile(id string, name string, createdAt time.Time, updatedAt time.Time, supplier BillingParty, workflow BillingWorkflow, apps BillingProfileAppsOrReference, default_ bool, ) *BillingProfile`

NewBillingProfile instantiates a new BillingProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingProfileWithDefaults

`func NewBillingProfileWithDefaults() *BillingProfile`

NewBillingProfileWithDefaults instantiates a new BillingProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BillingProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingProfile) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BillingProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingProfile) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BillingProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BillingProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BillingProfile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BillingProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *BillingProfile) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BillingProfile) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BillingProfile) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BillingProfile) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BillingProfile) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BillingProfile) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BillingProfile) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BillingProfile) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BillingProfile) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BillingProfile) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *BillingProfile) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *BillingProfile) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *BillingProfile) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *BillingProfile) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetSupplier

`func (o *BillingProfile) GetSupplier() BillingParty`

GetSupplier returns the Supplier field if non-nil, zero value otherwise.

### GetSupplierOk

`func (o *BillingProfile) GetSupplierOk() (*BillingParty, bool)`

GetSupplierOk returns a tuple with the Supplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplier

`func (o *BillingProfile) SetSupplier(v BillingParty)`

SetSupplier sets Supplier field to given value.


### GetWorkflow

`func (o *BillingProfile) GetWorkflow() BillingWorkflow`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *BillingProfile) GetWorkflowOk() (*BillingWorkflow, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *BillingProfile) SetWorkflow(v BillingWorkflow)`

SetWorkflow sets Workflow field to given value.


### GetApps

`func (o *BillingProfile) GetApps() BillingProfileAppsOrReference`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *BillingProfile) GetAppsOk() (*BillingProfileAppsOrReference, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *BillingProfile) SetApps(v BillingProfileAppsOrReference)`

SetApps sets Apps field to given value.


### GetDefault

`func (o *BillingProfile) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *BillingProfile) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *BillingProfile) SetDefault(v bool)`

SetDefault sets Default field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


