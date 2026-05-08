# PlanAddonCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **map[string]string** | Additional metadata for the resource. | [optional] 
**FromPlanPhase** | **string** | The key of the plan phase from the add-on becomes available for purchase. | 
**MaxQuantity** | Pointer to **int32** | The maximum number of times the add-on can be purchased for the plan. It is not applicable for add-ons with single instance type. | [optional] 
**AddonId** | **string** | The add-on unique identifier in ULID format. | 

## Methods

### NewPlanAddonCreate

`func NewPlanAddonCreate(fromPlanPhase string, addonId string, ) *PlanAddonCreate`

NewPlanAddonCreate instantiates a new PlanAddonCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanAddonCreateWithDefaults

`func NewPlanAddonCreateWithDefaults() *PlanAddonCreate`

NewPlanAddonCreateWithDefaults instantiates a new PlanAddonCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *PlanAddonCreate) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PlanAddonCreate) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PlanAddonCreate) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PlanAddonCreate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetFromPlanPhase

`func (o *PlanAddonCreate) GetFromPlanPhase() string`

GetFromPlanPhase returns the FromPlanPhase field if non-nil, zero value otherwise.

### GetFromPlanPhaseOk

`func (o *PlanAddonCreate) GetFromPlanPhaseOk() (*string, bool)`

GetFromPlanPhaseOk returns a tuple with the FromPlanPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromPlanPhase

`func (o *PlanAddonCreate) SetFromPlanPhase(v string)`

SetFromPlanPhase sets FromPlanPhase field to given value.


### GetMaxQuantity

`func (o *PlanAddonCreate) GetMaxQuantity() int32`

GetMaxQuantity returns the MaxQuantity field if non-nil, zero value otherwise.

### GetMaxQuantityOk

`func (o *PlanAddonCreate) GetMaxQuantityOk() (*int32, bool)`

GetMaxQuantityOk returns a tuple with the MaxQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQuantity

`func (o *PlanAddonCreate) SetMaxQuantity(v int32)`

SetMaxQuantity sets MaxQuantity field to given value.

### HasMaxQuantity

`func (o *PlanAddonCreate) HasMaxQuantity() bool`

HasMaxQuantity returns a boolean if a field has been set.

### GetAddonId

`func (o *PlanAddonCreate) GetAddonId() string`

GetAddonId returns the AddonId field if non-nil, zero value otherwise.

### GetAddonIdOk

`func (o *PlanAddonCreate) GetAddonIdOk() (*string, bool)`

GetAddonIdOk returns a tuple with the AddonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonId

`func (o *PlanAddonCreate) SetAddonId(v string)`

SetAddonId sets AddonId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


