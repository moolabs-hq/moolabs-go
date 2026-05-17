# GrantActivationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**SubscriptionId** | **string** |  | 
**CustomerId** | **string** |  | 
**PlanId** | **string** |  | 
**PlanVersion** | **int32** |  | 
**EffectiveAt** | **time.Time** |  | 
**CreditConfig** | [**CreditConfigPayload**](CreditConfigPayload.md) |  | 
**CommercialOverrides** | Pointer to [**CommercialOverridesPayload**](CommercialOverridesPayload.md) |  | [optional] 
**Entitlements** | [**[]EntitlementPayload**](EntitlementPayload.md) |  | 
**SubjectKeys** | **[]string** |  | 

## Methods

### NewGrantActivationRequest

`func NewGrantActivationRequest(tenantId string, subscriptionId string, customerId string, planId string, planVersion int32, effectiveAt time.Time, creditConfig CreditConfigPayload, entitlements []EntitlementPayload, subjectKeys []string, ) *GrantActivationRequest`

NewGrantActivationRequest instantiates a new GrantActivationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantActivationRequestWithDefaults

`func NewGrantActivationRequestWithDefaults() *GrantActivationRequest`

NewGrantActivationRequestWithDefaults instantiates a new GrantActivationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *GrantActivationRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *GrantActivationRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *GrantActivationRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetSubscriptionId

`func (o *GrantActivationRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *GrantActivationRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *GrantActivationRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetCustomerId

`func (o *GrantActivationRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *GrantActivationRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *GrantActivationRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetPlanId

`func (o *GrantActivationRequest) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *GrantActivationRequest) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *GrantActivationRequest) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.


### GetPlanVersion

`func (o *GrantActivationRequest) GetPlanVersion() int32`

GetPlanVersion returns the PlanVersion field if non-nil, zero value otherwise.

### GetPlanVersionOk

`func (o *GrantActivationRequest) GetPlanVersionOk() (*int32, bool)`

GetPlanVersionOk returns a tuple with the PlanVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanVersion

`func (o *GrantActivationRequest) SetPlanVersion(v int32)`

SetPlanVersion sets PlanVersion field to given value.


### GetEffectiveAt

`func (o *GrantActivationRequest) GetEffectiveAt() time.Time`

GetEffectiveAt returns the EffectiveAt field if non-nil, zero value otherwise.

### GetEffectiveAtOk

`func (o *GrantActivationRequest) GetEffectiveAtOk() (*time.Time, bool)`

GetEffectiveAtOk returns a tuple with the EffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAt

`func (o *GrantActivationRequest) SetEffectiveAt(v time.Time)`

SetEffectiveAt sets EffectiveAt field to given value.


### GetCreditConfig

`func (o *GrantActivationRequest) GetCreditConfig() CreditConfigPayload`

GetCreditConfig returns the CreditConfig field if non-nil, zero value otherwise.

### GetCreditConfigOk

`func (o *GrantActivationRequest) GetCreditConfigOk() (*CreditConfigPayload, bool)`

GetCreditConfigOk returns a tuple with the CreditConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditConfig

`func (o *GrantActivationRequest) SetCreditConfig(v CreditConfigPayload)`

SetCreditConfig sets CreditConfig field to given value.


### GetCommercialOverrides

`func (o *GrantActivationRequest) GetCommercialOverrides() CommercialOverridesPayload`

GetCommercialOverrides returns the CommercialOverrides field if non-nil, zero value otherwise.

### GetCommercialOverridesOk

`func (o *GrantActivationRequest) GetCommercialOverridesOk() (*CommercialOverridesPayload, bool)`

GetCommercialOverridesOk returns a tuple with the CommercialOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommercialOverrides

`func (o *GrantActivationRequest) SetCommercialOverrides(v CommercialOverridesPayload)`

SetCommercialOverrides sets CommercialOverrides field to given value.

### HasCommercialOverrides

`func (o *GrantActivationRequest) HasCommercialOverrides() bool`

HasCommercialOverrides returns a boolean if a field has been set.

### GetEntitlements

`func (o *GrantActivationRequest) GetEntitlements() []EntitlementPayload`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *GrantActivationRequest) GetEntitlementsOk() (*[]EntitlementPayload, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *GrantActivationRequest) SetEntitlements(v []EntitlementPayload)`

SetEntitlements sets Entitlements field to given value.


### GetSubjectKeys

`func (o *GrantActivationRequest) GetSubjectKeys() []string`

GetSubjectKeys returns the SubjectKeys field if non-nil, zero value otherwise.

### GetSubjectKeysOk

`func (o *GrantActivationRequest) GetSubjectKeysOk() (*[]string, bool)`

GetSubjectKeysOk returns a tuple with the SubjectKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectKeys

`func (o *GrantActivationRequest) SetSubjectKeys(v []string)`

SetSubjectKeys sets SubjectKeys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


