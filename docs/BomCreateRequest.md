# BomCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Tenant UUID | 
**FeatureKey** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**WorkflowType** | Pointer to **string** |  | [optional] 
**Components** | [**[]BomComponentIn**](BomComponentIn.md) |  | 
**CreatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewBomCreateRequest

`func NewBomCreateRequest(tenantId string, featureKey string, name string, components []BomComponentIn, ) *BomCreateRequest`

NewBomCreateRequest instantiates a new BomCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBomCreateRequestWithDefaults

`func NewBomCreateRequestWithDefaults() *BomCreateRequest`

NewBomCreateRequestWithDefaults instantiates a new BomCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *BomCreateRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *BomCreateRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *BomCreateRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetFeatureKey

`func (o *BomCreateRequest) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *BomCreateRequest) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *BomCreateRequest) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.


### GetName

`func (o *BomCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BomCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BomCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BomCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BomCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BomCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BomCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetWorkflowType

`func (o *BomCreateRequest) GetWorkflowType() string`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *BomCreateRequest) GetWorkflowTypeOk() (*string, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowType

`func (o *BomCreateRequest) SetWorkflowType(v string)`

SetWorkflowType sets WorkflowType field to given value.

### HasWorkflowType

`func (o *BomCreateRequest) HasWorkflowType() bool`

HasWorkflowType returns a boolean if a field has been set.

### GetComponents

`func (o *BomCreateRequest) GetComponents() []BomComponentIn`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *BomCreateRequest) GetComponentsOk() (*[]BomComponentIn, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *BomCreateRequest) SetComponents(v []BomComponentIn)`

SetComponents sets Components field to given value.


### GetCreatedBy

`func (o *BomCreateRequest) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BomCreateRequest) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BomCreateRequest) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *BomCreateRequest) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


