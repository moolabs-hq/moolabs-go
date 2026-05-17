# AppApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | **string** |  | 
**TenantId** | **string** |  | 
**SubscriptionId** | **string** |  | 
**CustomerId** | **string** |  | 
**PlanRef** | [**PlanRefPayload**](PlanRefPayload.md) |  | 
**EffectiveAt** | Pointer to **time.Time** |  | [optional] 
**Subscription** | **map[string]interface{}** |  | 
**CommercialOverrides** | Pointer to [**CommercialOverridesPayload**](CommercialOverridesPayload.md) |  | [optional] 

## Methods

### NewAppApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest

`func NewAppApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest(operation string, tenantId string, subscriptionId string, customerId string, planRef PlanRefPayload, subscription map[string]interface{}, ) *AppApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest`

NewAppApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest instantiates a new AppApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequestWithDefaults

`func NewAppApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequestWithDefaults() *AppApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest`

NewAppApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequestWithDefaults instantiates a new AppApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *AppApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *AppApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *AppApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetTenantId

`func (o *AppApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AppApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AppApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetSubscriptionId

`func (o *AppApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AppApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AppApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetCustomerId

`func (o *AppApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *AppApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *AppApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetPlanRef

`func (o *AppApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest) GetPlanRef() PlanRefPayload`

GetPlanRef returns the PlanRef field if non-nil, zero value otherwise.

### GetPlanRefOk

`func (o *AppApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest) GetPlanRefOk() (*PlanRefPayload, bool)`

GetPlanRefOk returns a tuple with the PlanRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanRef

`func (o *AppApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest) SetPlanRef(v PlanRefPayload)`

SetPlanRef sets PlanRef field to given value.


### GetEffectiveAt

`func (o *AppApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest) GetEffectiveAt() time.Time`

GetEffectiveAt returns the EffectiveAt field if non-nil, zero value otherwise.

### GetEffectiveAtOk

`func (o *AppApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest) GetEffectiveAtOk() (*time.Time, bool)`

GetEffectiveAtOk returns a tuple with the EffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAt

`func (o *AppApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest) SetEffectiveAt(v time.Time)`

SetEffectiveAt sets EffectiveAt field to given value.

### HasEffectiveAt

`func (o *AppApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest) HasEffectiveAt() bool`

HasEffectiveAt returns a boolean if a field has been set.

### GetSubscription

`func (o *AppApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest) GetSubscription() map[string]interface{}`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *AppApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest) GetSubscriptionOk() (*map[string]interface{}, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *AppApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest) SetSubscription(v map[string]interface{})`

SetSubscription sets Subscription field to given value.


### GetCommercialOverrides

`func (o *AppApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest) GetCommercialOverrides() CommercialOverridesPayload`

GetCommercialOverrides returns the CommercialOverrides field if non-nil, zero value otherwise.

### GetCommercialOverridesOk

`func (o *AppApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest) GetCommercialOverridesOk() (*CommercialOverridesPayload, bool)`

GetCommercialOverridesOk returns a tuple with the CommercialOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommercialOverrides

`func (o *AppApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest) SetCommercialOverrides(v CommercialOverridesPayload)`

SetCommercialOverrides sets CommercialOverrides field to given value.

### HasCommercialOverrides

`func (o *AppApiV1InternalSubscriptionSyncSchemasSubscriptionSyncRequest) HasCommercialOverrides() bool`

HasCommercialOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


