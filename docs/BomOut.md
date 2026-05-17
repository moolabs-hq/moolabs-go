# BomOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**FeatureKey** | **string** |  | 
**Version** | **int32** |  | 
**Status** | **string** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**WorkflowType** | **string** |  | 
**DerivedFrom** | **string** |  | 
**ApprovedBy** | **string** |  | 
**ApprovedAt** | **string** |  | 
**SupersededById** | **string** |  | 
**SupersededAt** | **string** |  | 
**CreatedBy** | **string** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**Components** | Pointer to [**[]BomComponentOut**](BomComponentOut.md) |  | [optional] 

## Methods

### NewBomOut

`func NewBomOut(id string, tenantId string, featureKey string, version int32, status string, name string, description string, workflowType string, derivedFrom string, approvedBy string, approvedAt string, supersededById string, supersededAt string, createdBy string, createdAt string, updatedAt string, ) *BomOut`

NewBomOut instantiates a new BomOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBomOutWithDefaults

`func NewBomOutWithDefaults() *BomOut`

NewBomOutWithDefaults instantiates a new BomOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BomOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BomOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BomOut) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *BomOut) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *BomOut) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *BomOut) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetFeatureKey

`func (o *BomOut) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *BomOut) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *BomOut) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.


### GetVersion

`func (o *BomOut) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BomOut) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BomOut) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetStatus

`func (o *BomOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BomOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BomOut) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetName

`func (o *BomOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BomOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BomOut) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BomOut) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BomOut) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BomOut) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetWorkflowType

`func (o *BomOut) GetWorkflowType() string`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *BomOut) GetWorkflowTypeOk() (*string, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowType

`func (o *BomOut) SetWorkflowType(v string)`

SetWorkflowType sets WorkflowType field to given value.


### GetDerivedFrom

`func (o *BomOut) GetDerivedFrom() string`

GetDerivedFrom returns the DerivedFrom field if non-nil, zero value otherwise.

### GetDerivedFromOk

`func (o *BomOut) GetDerivedFromOk() (*string, bool)`

GetDerivedFromOk returns a tuple with the DerivedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedFrom

`func (o *BomOut) SetDerivedFrom(v string)`

SetDerivedFrom sets DerivedFrom field to given value.


### GetApprovedBy

`func (o *BomOut) GetApprovedBy() string`

GetApprovedBy returns the ApprovedBy field if non-nil, zero value otherwise.

### GetApprovedByOk

`func (o *BomOut) GetApprovedByOk() (*string, bool)`

GetApprovedByOk returns a tuple with the ApprovedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedBy

`func (o *BomOut) SetApprovedBy(v string)`

SetApprovedBy sets ApprovedBy field to given value.


### GetApprovedAt

`func (o *BomOut) GetApprovedAt() string`

GetApprovedAt returns the ApprovedAt field if non-nil, zero value otherwise.

### GetApprovedAtOk

`func (o *BomOut) GetApprovedAtOk() (*string, bool)`

GetApprovedAtOk returns a tuple with the ApprovedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAt

`func (o *BomOut) SetApprovedAt(v string)`

SetApprovedAt sets ApprovedAt field to given value.


### GetSupersededById

`func (o *BomOut) GetSupersededById() string`

GetSupersededById returns the SupersededById field if non-nil, zero value otherwise.

### GetSupersededByIdOk

`func (o *BomOut) GetSupersededByIdOk() (*string, bool)`

GetSupersededByIdOk returns a tuple with the SupersededById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupersededById

`func (o *BomOut) SetSupersededById(v string)`

SetSupersededById sets SupersededById field to given value.


### GetSupersededAt

`func (o *BomOut) GetSupersededAt() string`

GetSupersededAt returns the SupersededAt field if non-nil, zero value otherwise.

### GetSupersededAtOk

`func (o *BomOut) GetSupersededAtOk() (*string, bool)`

GetSupersededAtOk returns a tuple with the SupersededAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupersededAt

`func (o *BomOut) SetSupersededAt(v string)`

SetSupersededAt sets SupersededAt field to given value.


### GetCreatedBy

`func (o *BomOut) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BomOut) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BomOut) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatedAt

`func (o *BomOut) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BomOut) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BomOut) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BomOut) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BomOut) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BomOut) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetComponents

`func (o *BomOut) GetComponents() []BomComponentOut`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *BomOut) GetComponentsOk() (*[]BomComponentOut, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *BomOut) SetComponents(v []BomComponentOut)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *BomOut) HasComponents() bool`

HasComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


