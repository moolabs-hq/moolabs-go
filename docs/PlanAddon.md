# PlanAddon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | Timestamp of when the resource was created. | [readonly] 
**UpdatedAt** | **time.Time** | Timestamp of when the resource was last updated. | [readonly] 
**DeletedAt** | Pointer to **time.Time** | Timestamp of when the resource was permanently deleted. | [optional] [readonly] 
**Annotations** | Pointer to **map[string]interface{}** | Set of key-value pairs managed by the system. Cannot be modified by user. | [optional] [readonly] 
**Metadata** | Pointer to **map[string]string** | Additional metadata for the resource. | [optional] 
**Addon** | [**Addon**](Addon.md) | Add-on object. | [readonly] 
**FromPlanPhase** | **string** | The key of the plan phase from the add-on becomes available for purchase. | 
**MaxQuantity** | Pointer to **int32** | The maximum number of times the add-on can be purchased for the plan. It is not applicable for add-ons with single instance type. | [optional] 
**ValidationErrors** | [**[]MeterValidationError**](MeterValidationError.md) | List of validation errors. | 

## Methods

### NewPlanAddon

`func NewPlanAddon(createdAt time.Time, updatedAt time.Time, addon Addon, fromPlanPhase string, validationErrors []MeterValidationError, ) *PlanAddon`

NewPlanAddon instantiates a new PlanAddon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanAddonWithDefaults

`func NewPlanAddonWithDefaults() *PlanAddon`

NewPlanAddonWithDefaults instantiates a new PlanAddon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PlanAddon) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PlanAddon) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PlanAddon) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PlanAddon) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PlanAddon) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PlanAddon) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *PlanAddon) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *PlanAddon) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *PlanAddon) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *PlanAddon) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetAnnotations

`func (o *PlanAddon) GetAnnotations() map[string]interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *PlanAddon) GetAnnotationsOk() (*map[string]interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *PlanAddon) SetAnnotations(v map[string]interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *PlanAddon) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetMetadata

`func (o *PlanAddon) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PlanAddon) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PlanAddon) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PlanAddon) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetAddon

`func (o *PlanAddon) GetAddon() Addon`

GetAddon returns the Addon field if non-nil, zero value otherwise.

### GetAddonOk

`func (o *PlanAddon) GetAddonOk() (*Addon, bool)`

GetAddonOk returns a tuple with the Addon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddon

`func (o *PlanAddon) SetAddon(v Addon)`

SetAddon sets Addon field to given value.


### GetFromPlanPhase

`func (o *PlanAddon) GetFromPlanPhase() string`

GetFromPlanPhase returns the FromPlanPhase field if non-nil, zero value otherwise.

### GetFromPlanPhaseOk

`func (o *PlanAddon) GetFromPlanPhaseOk() (*string, bool)`

GetFromPlanPhaseOk returns a tuple with the FromPlanPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromPlanPhase

`func (o *PlanAddon) SetFromPlanPhase(v string)`

SetFromPlanPhase sets FromPlanPhase field to given value.


### GetMaxQuantity

`func (o *PlanAddon) GetMaxQuantity() int32`

GetMaxQuantity returns the MaxQuantity field if non-nil, zero value otherwise.

### GetMaxQuantityOk

`func (o *PlanAddon) GetMaxQuantityOk() (*int32, bool)`

GetMaxQuantityOk returns a tuple with the MaxQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQuantity

`func (o *PlanAddon) SetMaxQuantity(v int32)`

SetMaxQuantity sets MaxQuantity field to given value.

### HasMaxQuantity

`func (o *PlanAddon) HasMaxQuantity() bool`

HasMaxQuantity returns a boolean if a field has been set.

### GetValidationErrors

`func (o *PlanAddon) GetValidationErrors() []MeterValidationError`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *PlanAddon) GetValidationErrorsOk() (*[]MeterValidationError, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *PlanAddon) SetValidationErrors(v []MeterValidationError)`

SetValidationErrors sets ValidationErrors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


