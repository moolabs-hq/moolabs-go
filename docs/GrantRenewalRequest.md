# GrantRenewalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**SubscriptionId** | **string** |  | 
**CustomerId** | **string** |  | 
**PlanId** | **string** |  | 
**PlanVersion** | **int32** |  | 
**BillingPeriodEnd** | **time.Time** |  | 
**CreditConfig** | [**CreditConfigPayload**](CreditConfigPayload.md) |  | 
**CommercialOverrides** | Pointer to [**CommercialOverridesPayload**](CommercialOverridesPayload.md) |  | [optional] 
**Entitlements** | [**[]EntitlementPayload**](EntitlementPayload.md) |  | 
**SubjectKeys** | **[]string** |  | 

## Methods

### NewGrantRenewalRequest

`func NewGrantRenewalRequest(tenantId string, subscriptionId string, customerId string, planId string, planVersion int32, billingPeriodEnd time.Time, creditConfig CreditConfigPayload, entitlements []EntitlementPayload, subjectKeys []string, ) *GrantRenewalRequest`

NewGrantRenewalRequest instantiates a new GrantRenewalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantRenewalRequestWithDefaults

`func NewGrantRenewalRequestWithDefaults() *GrantRenewalRequest`

NewGrantRenewalRequestWithDefaults instantiates a new GrantRenewalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *GrantRenewalRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *GrantRenewalRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *GrantRenewalRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetSubscriptionId

`func (o *GrantRenewalRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *GrantRenewalRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *GrantRenewalRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetCustomerId

`func (o *GrantRenewalRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *GrantRenewalRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *GrantRenewalRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetPlanId

`func (o *GrantRenewalRequest) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *GrantRenewalRequest) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *GrantRenewalRequest) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.


### GetPlanVersion

`func (o *GrantRenewalRequest) GetPlanVersion() int32`

GetPlanVersion returns the PlanVersion field if non-nil, zero value otherwise.

### GetPlanVersionOk

`func (o *GrantRenewalRequest) GetPlanVersionOk() (*int32, bool)`

GetPlanVersionOk returns a tuple with the PlanVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanVersion

`func (o *GrantRenewalRequest) SetPlanVersion(v int32)`

SetPlanVersion sets PlanVersion field to given value.


### GetBillingPeriodEnd

`func (o *GrantRenewalRequest) GetBillingPeriodEnd() time.Time`

GetBillingPeriodEnd returns the BillingPeriodEnd field if non-nil, zero value otherwise.

### GetBillingPeriodEndOk

`func (o *GrantRenewalRequest) GetBillingPeriodEndOk() (*time.Time, bool)`

GetBillingPeriodEndOk returns a tuple with the BillingPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriodEnd

`func (o *GrantRenewalRequest) SetBillingPeriodEnd(v time.Time)`

SetBillingPeriodEnd sets BillingPeriodEnd field to given value.


### GetCreditConfig

`func (o *GrantRenewalRequest) GetCreditConfig() CreditConfigPayload`

GetCreditConfig returns the CreditConfig field if non-nil, zero value otherwise.

### GetCreditConfigOk

`func (o *GrantRenewalRequest) GetCreditConfigOk() (*CreditConfigPayload, bool)`

GetCreditConfigOk returns a tuple with the CreditConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditConfig

`func (o *GrantRenewalRequest) SetCreditConfig(v CreditConfigPayload)`

SetCreditConfig sets CreditConfig field to given value.


### GetCommercialOverrides

`func (o *GrantRenewalRequest) GetCommercialOverrides() CommercialOverridesPayload`

GetCommercialOverrides returns the CommercialOverrides field if non-nil, zero value otherwise.

### GetCommercialOverridesOk

`func (o *GrantRenewalRequest) GetCommercialOverridesOk() (*CommercialOverridesPayload, bool)`

GetCommercialOverridesOk returns a tuple with the CommercialOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommercialOverrides

`func (o *GrantRenewalRequest) SetCommercialOverrides(v CommercialOverridesPayload)`

SetCommercialOverrides sets CommercialOverrides field to given value.

### HasCommercialOverrides

`func (o *GrantRenewalRequest) HasCommercialOverrides() bool`

HasCommercialOverrides returns a boolean if a field has been set.

### GetEntitlements

`func (o *GrantRenewalRequest) GetEntitlements() []EntitlementPayload`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *GrantRenewalRequest) GetEntitlementsOk() (*[]EntitlementPayload, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *GrantRenewalRequest) SetEntitlements(v []EntitlementPayload)`

SetEntitlements sets Entitlements field to given value.


### GetSubjectKeys

`func (o *GrantRenewalRequest) GetSubjectKeys() []string`

GetSubjectKeys returns the SubjectKeys field if non-nil, zero value otherwise.

### GetSubjectKeysOk

`func (o *GrantRenewalRequest) GetSubjectKeysOk() (*[]string, bool)`

GetSubjectKeysOk returns a tuple with the SubjectKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectKeys

`func (o *GrantRenewalRequest) SetSubjectKeys(v []string)`

SetSubjectKeys sets SubjectKeys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


