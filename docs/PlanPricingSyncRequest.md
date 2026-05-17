# PlanPricingSyncRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**Operation** | **string** |  | 
**PlanId** | **string** |  | 
**Plan** | Pointer to **map[string]interface{}** |  | [optional] 
**EffectiveAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPlanPricingSyncRequest

`func NewPlanPricingSyncRequest(tenantId string, operation string, planId string, ) *PlanPricingSyncRequest`

NewPlanPricingSyncRequest instantiates a new PlanPricingSyncRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanPricingSyncRequestWithDefaults

`func NewPlanPricingSyncRequestWithDefaults() *PlanPricingSyncRequest`

NewPlanPricingSyncRequestWithDefaults instantiates a new PlanPricingSyncRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *PlanPricingSyncRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *PlanPricingSyncRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *PlanPricingSyncRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetOperation

`func (o *PlanPricingSyncRequest) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *PlanPricingSyncRequest) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *PlanPricingSyncRequest) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetPlanId

`func (o *PlanPricingSyncRequest) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *PlanPricingSyncRequest) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *PlanPricingSyncRequest) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.


### GetPlan

`func (o *PlanPricingSyncRequest) GetPlan() map[string]interface{}`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *PlanPricingSyncRequest) GetPlanOk() (*map[string]interface{}, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *PlanPricingSyncRequest) SetPlan(v map[string]interface{})`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *PlanPricingSyncRequest) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetEffectiveAt

`func (o *PlanPricingSyncRequest) GetEffectiveAt() time.Time`

GetEffectiveAt returns the EffectiveAt field if non-nil, zero value otherwise.

### GetEffectiveAtOk

`func (o *PlanPricingSyncRequest) GetEffectiveAtOk() (*time.Time, bool)`

GetEffectiveAtOk returns a tuple with the EffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAt

`func (o *PlanPricingSyncRequest) SetEffectiveAt(v time.Time)`

SetEffectiveAt sets EffectiveAt field to given value.

### HasEffectiveAt

`func (o *PlanPricingSyncRequest) HasEffectiveAt() bool`

HasEffectiveAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


