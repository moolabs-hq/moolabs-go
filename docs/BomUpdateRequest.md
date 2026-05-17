# BomUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**WorkflowType** | Pointer to **string** |  | [optional] 
**Components** | Pointer to [**[]BomComponentIn**](BomComponentIn.md) |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewBomUpdateRequest

`func NewBomUpdateRequest() *BomUpdateRequest`

NewBomUpdateRequest instantiates a new BomUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBomUpdateRequestWithDefaults

`func NewBomUpdateRequestWithDefaults() *BomUpdateRequest`

NewBomUpdateRequestWithDefaults instantiates a new BomUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BomUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BomUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BomUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BomUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *BomUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BomUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BomUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BomUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetWorkflowType

`func (o *BomUpdateRequest) GetWorkflowType() string`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *BomUpdateRequest) GetWorkflowTypeOk() (*string, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowType

`func (o *BomUpdateRequest) SetWorkflowType(v string)`

SetWorkflowType sets WorkflowType field to given value.

### HasWorkflowType

`func (o *BomUpdateRequest) HasWorkflowType() bool`

HasWorkflowType returns a boolean if a field has been set.

### GetComponents

`func (o *BomUpdateRequest) GetComponents() []BomComponentIn`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *BomUpdateRequest) GetComponentsOk() (*[]BomComponentIn, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *BomUpdateRequest) SetComponents(v []BomComponentIn)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *BomUpdateRequest) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetCreatedBy

`func (o *BomUpdateRequest) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BomUpdateRequest) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BomUpdateRequest) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *BomUpdateRequest) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


