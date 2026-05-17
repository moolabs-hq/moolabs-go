# CommercialOverridesInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | **string** |  | 
**CustomerId** | **string** |  | 
**TenantId** | Pointer to **string** |  | [optional] 
**EffectiveFrom** | **string** |  | 
**EntitlementOverrides** | Pointer to [**[]EntitlementOverrideInput**](EntitlementOverrideInput.md) |  | [optional] 
**CreditOverride** | Pointer to [**CreditOverrideInput**](CreditOverrideInput.md) |  | [optional] 
**RateCardDiscount** | Pointer to [**RateCardDiscountInput**](RateCardDiscountInput.md) |  | [optional] 

## Methods

### NewCommercialOverridesInput

`func NewCommercialOverridesInput(subscriptionId string, customerId string, effectiveFrom string, ) *CommercialOverridesInput`

NewCommercialOverridesInput instantiates a new CommercialOverridesInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommercialOverridesInputWithDefaults

`func NewCommercialOverridesInputWithDefaults() *CommercialOverridesInput`

NewCommercialOverridesInputWithDefaults instantiates a new CommercialOverridesInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *CommercialOverridesInput) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CommercialOverridesInput) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CommercialOverridesInput) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetCustomerId

`func (o *CommercialOverridesInput) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CommercialOverridesInput) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CommercialOverridesInput) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetTenantId

`func (o *CommercialOverridesInput) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CommercialOverridesInput) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CommercialOverridesInput) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CommercialOverridesInput) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetEffectiveFrom

`func (o *CommercialOverridesInput) GetEffectiveFrom() string`

GetEffectiveFrom returns the EffectiveFrom field if non-nil, zero value otherwise.

### GetEffectiveFromOk

`func (o *CommercialOverridesInput) GetEffectiveFromOk() (*string, bool)`

GetEffectiveFromOk returns a tuple with the EffectiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveFrom

`func (o *CommercialOverridesInput) SetEffectiveFrom(v string)`

SetEffectiveFrom sets EffectiveFrom field to given value.


### GetEntitlementOverrides

`func (o *CommercialOverridesInput) GetEntitlementOverrides() []EntitlementOverrideInput`

GetEntitlementOverrides returns the EntitlementOverrides field if non-nil, zero value otherwise.

### GetEntitlementOverridesOk

`func (o *CommercialOverridesInput) GetEntitlementOverridesOk() (*[]EntitlementOverrideInput, bool)`

GetEntitlementOverridesOk returns a tuple with the EntitlementOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementOverrides

`func (o *CommercialOverridesInput) SetEntitlementOverrides(v []EntitlementOverrideInput)`

SetEntitlementOverrides sets EntitlementOverrides field to given value.

### HasEntitlementOverrides

`func (o *CommercialOverridesInput) HasEntitlementOverrides() bool`

HasEntitlementOverrides returns a boolean if a field has been set.

### GetCreditOverride

`func (o *CommercialOverridesInput) GetCreditOverride() CreditOverrideInput`

GetCreditOverride returns the CreditOverride field if non-nil, zero value otherwise.

### GetCreditOverrideOk

`func (o *CommercialOverridesInput) GetCreditOverrideOk() (*CreditOverrideInput, bool)`

GetCreditOverrideOk returns a tuple with the CreditOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditOverride

`func (o *CommercialOverridesInput) SetCreditOverride(v CreditOverrideInput)`

SetCreditOverride sets CreditOverride field to given value.

### HasCreditOverride

`func (o *CommercialOverridesInput) HasCreditOverride() bool`

HasCreditOverride returns a boolean if a field has been set.

### GetRateCardDiscount

`func (o *CommercialOverridesInput) GetRateCardDiscount() RateCardDiscountInput`

GetRateCardDiscount returns the RateCardDiscount field if non-nil, zero value otherwise.

### GetRateCardDiscountOk

`func (o *CommercialOverridesInput) GetRateCardDiscountOk() (*RateCardDiscountInput, bool)`

GetRateCardDiscountOk returns a tuple with the RateCardDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateCardDiscount

`func (o *CommercialOverridesInput) SetRateCardDiscount(v RateCardDiscountInput)`

SetRateCardDiscount sets RateCardDiscount field to given value.

### HasRateCardDiscount

`func (o *CommercialOverridesInput) HasRateCardDiscount() bool`

HasRateCardDiscount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


