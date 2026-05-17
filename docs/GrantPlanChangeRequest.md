# GrantPlanChangeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**SubscriptionId** | **string** |  | 
**CustomerId** | **string** |  | 
**OldPlanId** | **string** |  | 
**OldPlanVersion** | **int32** |  | 
**NewPlanId** | **string** |  | 
**NewPlanVersion** | **int32** |  | 
**EffectiveAt** | **time.Time** |  | 
**NewCreditConfig** | [**CreditConfigPayload**](CreditConfigPayload.md) |  | 
**CommercialOverrides** | Pointer to [**CommercialOverridesPayload**](CommercialOverridesPayload.md) |  | [optional] 
**Entitlements** | [**[]EntitlementPayload**](EntitlementPayload.md) |  | 
**SubjectKeys** | **[]string** |  | 

## Methods

### NewGrantPlanChangeRequest

`func NewGrantPlanChangeRequest(tenantId string, subscriptionId string, customerId string, oldPlanId string, oldPlanVersion int32, newPlanId string, newPlanVersion int32, effectiveAt time.Time, newCreditConfig CreditConfigPayload, entitlements []EntitlementPayload, subjectKeys []string, ) *GrantPlanChangeRequest`

NewGrantPlanChangeRequest instantiates a new GrantPlanChangeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantPlanChangeRequestWithDefaults

`func NewGrantPlanChangeRequestWithDefaults() *GrantPlanChangeRequest`

NewGrantPlanChangeRequestWithDefaults instantiates a new GrantPlanChangeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *GrantPlanChangeRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *GrantPlanChangeRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *GrantPlanChangeRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetSubscriptionId

`func (o *GrantPlanChangeRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *GrantPlanChangeRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *GrantPlanChangeRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetCustomerId

`func (o *GrantPlanChangeRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *GrantPlanChangeRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *GrantPlanChangeRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetOldPlanId

`func (o *GrantPlanChangeRequest) GetOldPlanId() string`

GetOldPlanId returns the OldPlanId field if non-nil, zero value otherwise.

### GetOldPlanIdOk

`func (o *GrantPlanChangeRequest) GetOldPlanIdOk() (*string, bool)`

GetOldPlanIdOk returns a tuple with the OldPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPlanId

`func (o *GrantPlanChangeRequest) SetOldPlanId(v string)`

SetOldPlanId sets OldPlanId field to given value.


### GetOldPlanVersion

`func (o *GrantPlanChangeRequest) GetOldPlanVersion() int32`

GetOldPlanVersion returns the OldPlanVersion field if non-nil, zero value otherwise.

### GetOldPlanVersionOk

`func (o *GrantPlanChangeRequest) GetOldPlanVersionOk() (*int32, bool)`

GetOldPlanVersionOk returns a tuple with the OldPlanVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPlanVersion

`func (o *GrantPlanChangeRequest) SetOldPlanVersion(v int32)`

SetOldPlanVersion sets OldPlanVersion field to given value.


### GetNewPlanId

`func (o *GrantPlanChangeRequest) GetNewPlanId() string`

GetNewPlanId returns the NewPlanId field if non-nil, zero value otherwise.

### GetNewPlanIdOk

`func (o *GrantPlanChangeRequest) GetNewPlanIdOk() (*string, bool)`

GetNewPlanIdOk returns a tuple with the NewPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPlanId

`func (o *GrantPlanChangeRequest) SetNewPlanId(v string)`

SetNewPlanId sets NewPlanId field to given value.


### GetNewPlanVersion

`func (o *GrantPlanChangeRequest) GetNewPlanVersion() int32`

GetNewPlanVersion returns the NewPlanVersion field if non-nil, zero value otherwise.

### GetNewPlanVersionOk

`func (o *GrantPlanChangeRequest) GetNewPlanVersionOk() (*int32, bool)`

GetNewPlanVersionOk returns a tuple with the NewPlanVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPlanVersion

`func (o *GrantPlanChangeRequest) SetNewPlanVersion(v int32)`

SetNewPlanVersion sets NewPlanVersion field to given value.


### GetEffectiveAt

`func (o *GrantPlanChangeRequest) GetEffectiveAt() time.Time`

GetEffectiveAt returns the EffectiveAt field if non-nil, zero value otherwise.

### GetEffectiveAtOk

`func (o *GrantPlanChangeRequest) GetEffectiveAtOk() (*time.Time, bool)`

GetEffectiveAtOk returns a tuple with the EffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAt

`func (o *GrantPlanChangeRequest) SetEffectiveAt(v time.Time)`

SetEffectiveAt sets EffectiveAt field to given value.


### GetNewCreditConfig

`func (o *GrantPlanChangeRequest) GetNewCreditConfig() CreditConfigPayload`

GetNewCreditConfig returns the NewCreditConfig field if non-nil, zero value otherwise.

### GetNewCreditConfigOk

`func (o *GrantPlanChangeRequest) GetNewCreditConfigOk() (*CreditConfigPayload, bool)`

GetNewCreditConfigOk returns a tuple with the NewCreditConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewCreditConfig

`func (o *GrantPlanChangeRequest) SetNewCreditConfig(v CreditConfigPayload)`

SetNewCreditConfig sets NewCreditConfig field to given value.


### GetCommercialOverrides

`func (o *GrantPlanChangeRequest) GetCommercialOverrides() CommercialOverridesPayload`

GetCommercialOverrides returns the CommercialOverrides field if non-nil, zero value otherwise.

### GetCommercialOverridesOk

`func (o *GrantPlanChangeRequest) GetCommercialOverridesOk() (*CommercialOverridesPayload, bool)`

GetCommercialOverridesOk returns a tuple with the CommercialOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommercialOverrides

`func (o *GrantPlanChangeRequest) SetCommercialOverrides(v CommercialOverridesPayload)`

SetCommercialOverrides sets CommercialOverrides field to given value.

### HasCommercialOverrides

`func (o *GrantPlanChangeRequest) HasCommercialOverrides() bool`

HasCommercialOverrides returns a boolean if a field has been set.

### GetEntitlements

`func (o *GrantPlanChangeRequest) GetEntitlements() []EntitlementPayload`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *GrantPlanChangeRequest) GetEntitlementsOk() (*[]EntitlementPayload, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *GrantPlanChangeRequest) SetEntitlements(v []EntitlementPayload)`

SetEntitlements sets Entitlements field to given value.


### GetSubjectKeys

`func (o *GrantPlanChangeRequest) GetSubjectKeys() []string`

GetSubjectKeys returns the SubjectKeys field if non-nil, zero value otherwise.

### GetSubjectKeysOk

`func (o *GrantPlanChangeRequest) GetSubjectKeysOk() (*[]string, bool)`

GetSubjectKeysOk returns a tuple with the SubjectKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectKeys

`func (o *GrantPlanChangeRequest) SetSubjectKeys(v []string)`

SetSubjectKeys sets SubjectKeys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


