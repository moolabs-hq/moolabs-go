# ReconcileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | Pointer to **string** | Deprecated. Tenant is derived from API key. If provided, must match authenticated tenant. | [optional] 
**CostEventIds** | Pointer to **[]string** | Specific cost_event IDs to reconcile. If omitted, runs full pass. | [optional] 

## Methods

### NewReconcileRequest

`func NewReconcileRequest() *ReconcileRequest`

NewReconcileRequest instantiates a new ReconcileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReconcileRequestWithDefaults

`func NewReconcileRequestWithDefaults() *ReconcileRequest`

NewReconcileRequestWithDefaults instantiates a new ReconcileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *ReconcileRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ReconcileRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ReconcileRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ReconcileRequest) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetCostEventIds

`func (o *ReconcileRequest) GetCostEventIds() []string`

GetCostEventIds returns the CostEventIds field if non-nil, zero value otherwise.

### GetCostEventIdsOk

`func (o *ReconcileRequest) GetCostEventIdsOk() (*[]string, bool)`

GetCostEventIdsOk returns a tuple with the CostEventIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostEventIds

`func (o *ReconcileRequest) SetCostEventIds(v []string)`

SetCostEventIds sets CostEventIds field to given value.

### HasCostEventIds

`func (o *ReconcileRequest) HasCostEventIds() bool`

HasCostEventIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


